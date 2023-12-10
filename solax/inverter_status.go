package solax

import (
	"solax_exporter/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

type InverterStatusCode string

const (
	UndefinedStatus    InverterStatusCode = ""
	WaitMode           InverterStatusCode = "wait_mode"
	CheckMode          InverterStatusCode = "check_mode"
	NormalMode         InverterStatusCode = "normal_mode"
	FaultMode          InverterStatusCode = "fault_mode"
	PermanentFaultMode InverterStatusCode = "permanent_fault_mode"
	UpdateMode         InverterStatusCode = "update_mode"
	EPSCheckMode       InverterStatusCode = "eps_check_mode"
	EPSMode            InverterStatusCode = "eps_mode"
	SelfTestMode       InverterStatusCode = "self_test_mode"
	IdleMode           InverterStatusCode = "idle_mode"
	StandbyMode        InverterStatusCode = "standby_mode"
	PvWakeUpBatMode    InverterStatusCode = "pv_wake_up_bat_mode"
	GenCheckMode       InverterStatusCode = "gen_check_mode"
	GenRunMode         InverterStatusCode = "gen_run_mode"
)

var (
	inverterStatusMap map[string]InverterStatusCode = map[string]InverterStatusCode{
		"100": WaitMode,
		"101": CheckMode,
		"102": NormalMode,
		"103": FaultMode,
		"104": PermanentFaultMode,
		"105": UpdateMode,
		"106": EPSCheckMode,
		"107": EPSMode,
		"108": SelfTestMode,
		"109": IdleMode,
		"110": StandbyMode,
		"111": PvWakeUpBatMode,
		"112": GenCheckMode,
		"113": GenRunMode,
	}
)

func (i InverterStatusCode) String() string {
	return string(i)
}

func InverterStatusCodeFromString(s string) InverterStatusCode {
	return inverterStatusMap[s]
}

func DescribeInverterStatusCode(name string, dynamicLabels []string, labels prometheus.Labels) *metrics.EnumCollector {
	descs := make(metrics.EnumCollector)
	for _, code := range inverterStatusMap {
		valueLabels := prometheus.Labels{
			"value": code.String(),
		}
		for k, v := range labels {
			valueLabels[k] = v
		}
		descs[code.String()] = prometheus.NewDesc(name, "Inverter status.", dynamicLabels, valueLabels)
	}
	return &descs
}
