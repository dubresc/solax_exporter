package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CloudAPIRespose struct {
	InverterSN     string
	SN             string
	ACPower        float64
	YieldToday     float64
	YieldTotal     float64
	FeedInPower    float64
	FeedInEnergy   float64
	ConsumeEnergy  float64
	FeedInPowerM2  float64
	SOC            float64
	Peps1          *float64
	Peps2          *float64
	Peps3          *float64
	InverterType   InverterTypeCode
	InverterStatus InverterStatusCode
	UploadTime     time.Time
	BatPower       float64
	PowerDC1       *float64
	PowerDC2       *float64
	PowerDC3       *float64
	PowerDC4       *float64
	BatStatus      string // TODO
}

type CloudAPI struct {
	SN           string
	TokenID      string
	requestUrl   string
	lastResponse *CloudAPIRespose
}

func MakeCloudApi(sn string, token_id string) *CloudAPI {
	if len(sn) == 0 || len(token_id) == 0 {
		return nil
	}

	api := CloudAPI{
		SN:           sn,
		TokenID:      token_id,
		requestUrl:   fmt.Sprintf("https://www.solaxcloud.com/proxyApp/proxy/api/getRealtimeInfo.do?tokenId=%s&sn=%s", token_id, sn),
		lastResponse: nil,
	}

	return &api
}

type CloudApiParseError struct {
	Cause string
}

func (e CloudApiParseError) Error() string {
	return fmt.Sprintf("Error parsing json: %s", e.Cause)
}

type CloudApiError struct{}

func (CloudApiError) Error() string {
	return "API Error"
}

func GetField[K any](m map[string]interface{}, f string) (K, bool) {
	var res K

	iface, ok := m[f]

	if !ok {
		return res, false
	}

	res, ok = iface.(K)

	return res, ok
}

func GetNullableField[K any](m map[string]interface{}, f string) *K {
	var res K

	iface, ok := m[f]

	if !ok {
		return nil
	}

	res, ok = iface.(K)

	if !ok {
		return nil
	}

	return &res
}

func Parse(r []byte) (*CloudAPIRespose, error) {
	var j map[string]interface{}

	err := json.Unmarshal(r, &j)

	var res CloudAPIRespose

	if err != nil {
		return nil, err
	}

	success, ok := GetField[bool](j, "success")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"success\" field or not bool"}
	}

	if !success {
		return nil, &CloudApiError{}
	}

	result, ok := GetField[map[string]interface{}](j, "result")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"result\" field or not an object"}
	}

	res.InverterSN, ok = GetField[string](result, "inverterSN")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"inverterSN\" field or not a string"}
	}

	res.SN, ok = GetField[string](result, "sn")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"sn\" field or not a string"}
	}

	inverterStatus, ok := GetField[string](result, "inverterStatus")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"inverterStatus\" field or not a string"}
	}

	res.InverterStatus = InverterStatusCodeFromString(inverterStatus)

	res.ACPower, ok = GetField[float64](result, "acpower")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"acpower\" field or not a number"}
	}

	res.YieldToday, ok = GetField[float64](result, "yieldtoday")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"yieldtoday\" field or not a number"}
	}

	res.YieldTotal, ok = GetField[float64](result, "yieldtotal")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"yieldtotal\" field or not a number"}
	}

	res.FeedInPower, ok = GetField[float64](result, "feedinpower")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"feedinpower\" field or not a number"}
	}

	res.FeedInEnergy, ok = GetField[float64](result, "feedinenergy")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"feedinenergy\" field or not a number"}
	}

	res.ConsumeEnergy, ok = GetField[float64](result, "consumeenergy")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"consumeenergy\" field or not a number"}
	}

	res.FeedInPowerM2, ok = GetField[float64](result, "feedinpowerM2")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"feedinpowerM2\" field or not a number"}
	}

	res.SOC, ok = GetField[float64](result, "soc")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"soc\" field or not a number"}
	}

	res.Peps1 = GetNullableField[float64](result, "peps1")
	res.Peps2 = GetNullableField[float64](result, "peps2")
	res.Peps3 = GetNullableField[float64](result, "peps3")

	// TODO inverterType

	uploadTime, ok := GetField[string](result, "uploadTime")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"uploadTime\" field or not a string"}
	}

	res.UploadTime, err = time.Parse(time.DateTime, uploadTime)

	if err != nil {
		return nil, &CloudApiParseError{Cause: "Failed to parse \"uploadTime\""}
	}

	res.BatPower, ok = GetField[float64](result, "batPower")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"batPower\" field or not a number"}
	}

	res.PowerDC1 = GetNullableField[float64](result, "powerdc1")
	res.PowerDC2 = GetNullableField[float64](result, "powerdc2")
	res.PowerDC3 = GetNullableField[float64](result, "powerdc3")
	res.PowerDC4 = GetNullableField[float64](result, "powerdc4")

	// TODO: resolve to a enum
	res.BatStatus, ok = GetField[string](result, "batStatus")

	if !ok {
		return nil, &CloudApiParseError{Cause: "No \"batStatus\" field or not a string"}
	}

	return &res, nil
}

func (r CloudAPI) Request() (*CloudAPIRespose, error) {
	if r.lastResponse != nil && time.Now().Sub(r.lastResponse.UploadTime).Minutes() < 5.0 {
		return r.lastResponse, nil
	}

	res, err := http.Get(r.requestUrl)

	if err != nil {
		return nil, err
	}

	// TODO check status

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		return nil, err
	}

	r.lastResponse, err = Parse(body)

	return r.lastResponse, err
}
