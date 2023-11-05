package test

import (
	"encoding/json"
	"solax-exporter/src/api"
	"testing"
)

func TestParseCloud(t *testing.T) {
	r, err := api.Parse([]byte(`{
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

	want := api.CloudAPIRespose{
		InverterStatus: api.NormalMode,
	}

	if *r != want {
		wantJson, _ := json.Marshal(want)
		gotJson, _ := json.Marshal(*r)
		t.Errorf("Incorrect response. Wanted %s, got %s", string(wantJson), string(gotJson))
	}
}
