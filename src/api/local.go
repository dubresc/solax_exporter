package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type LocalApiResponse struct {
	Foo string
}

type LocalApi struct {
	URL string
	SN  string
}

func (l LocalApi) Request() (*LocalApiResponse, error) {
	bodyString := fmt.Sprintf("optType=ReadRealTimeData&pwd=%s", l.SN)
	reader := bytes.NewReader([]byte(bodyString))

	res, err := http.Post(l.URL, "application/x-www-form-urlencoded", reader)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		return nil, err
	}

	// TODO

	return &LocalApiResponse{Foo: string(body)}, nil
}
