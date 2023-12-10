package solax_test

import (
	"encoding/json"
	"reflect"
	"solax_exporter/solax"
	"testing"
	"time"
)

func TestParseCloud(t *testing.T) {
	r, err := solax.ParseCloudRespose([]byte(`{
		"success":true,
		"exception":"Query success!",
		"result":{
			"inverterSN":"NOTAREALSN4242",
			"sn":"NOTAREALSN",
			"acpower":587.0,
			"yieldtoday":4.2,
			"yieldtotal":133.7,
			"feedinpower":0.0,
			"feedinenergy":5.1,
			"consumeenergy":12.44,
			"feedinpowerM2":0.0,
			"soc":14.0,
			"peps1":0.0,
			"peps2":null,
			"peps3":null,
			"inverterType":"15",
			"inverterStatus":"102",
			"uploadTime":"2023-10-18 21:25:38",
			"batPower":-576.0,
			"powerdc1":0.0,
			"powerdc2":1.0,
			"powerdc3":null,
			"powerdc4":null,
			"batStatus":"0"
		},
		"code":0
	}`))

	if err != nil {
		t.Errorf("Parse returned error %s", err)
	}

	want := solax.CloudAPIRespose{
		InverterSN:     "NOTAREALSN4242",
		SN:             "NOTAREALSN",
		InverterStatus: solax.NormalMode,
		InverterType:   solax.X1_Hybrid_G4,
		ACPower:        587.0,
		YieldToday:     4.2,
		YieldTotal:     133.7,
		FeedInPower:    0.0,
		FeedInEnergy:   5.1,
		ConsumeEnergy:  12.44,
		FeedInPowerM2:  0.0,
		SOC:            14.0,
		Peps1:          new(float64),
		Peps2:          nil,
		Peps3:          nil,
		UploadTime:     time.Date(2023, time.October, 18, 21, 25, 38, 0, time.UTC),
		BatPower:       -576.0,
		PowerDC1:       new(float64),
		PowerDC2:       new(float64),
		PowerDC3:       nil,
		PowerDC4:       nil,
		BatStatus:      "0",
	}

	*want.Peps1 = 0.0
	*want.PowerDC1 = 0.0
	*want.PowerDC2 = 1.0

	if !reflect.DeepEqual(*r, want) {
		wantJson, _ := json.Marshal(want)
		gotJson, _ := json.Marshal(*r)
		t.Errorf("Incorrect response. Wanted %s, got %s", string(wantJson), string(gotJson))
	}
}
