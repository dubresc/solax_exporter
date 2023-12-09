package solax

import (
	"solax-exporter/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

type InverterTypeCode string

const (
	X1_LX             InverterTypeCode = "X1-LX"
	X_Hybrid          InverterTypeCode = "X-Hybrid"
	X1_Hybriyd_Fit    InverterTypeCode = "X1-Hybriyd_Fit"
	X1_Boost_Air_Mini InverterTypeCode = "X1-Boost_Air_Mini"
	X3_Hybriyd_Fit    InverterTypeCode = "X3-Hybriyd_Fit"
	X3_20K_30K        InverterTypeCode = "X3-20K_30K"
	X3_MIC_PRO        InverterTypeCode = "X3-MIC_PRO"
	X1_Smart          InverterTypeCode = "X1-Smart"
	X1_AC             InverterTypeCode = "X1-AC"
	A1_Hybrid         InverterTypeCode = "A1-Hybrid"
	A1_Fit            InverterTypeCode = "A1-Fit"
	A1_Grid           InverterTypeCode = "A1-Grid"
	J1_ESS            InverterTypeCode = "J1-ESS"
	X3_Hybrid_G4      InverterTypeCode = "X3-Hybrid-G4"
	X1_Hybrid_G4      InverterTypeCode = "X1-Hybrid-G4"
	X3_MIC_PRO_G2     InverterTypeCode = "X3-MIC_Pro-G2"
	X1_SPT            InverterTypeCode = "X1-SPT"
	X1_Boost_Mini_G4  InverterTypeCode = "X1-Boost_Mini-G4"
	A1_HYB_G2         InverterTypeCode = "A1-HYB-G2"
	A1_AC_G2          InverterTypeCode = "A1-AC-G2"
	A1_SMT_G2         InverterTypeCode = "A1-SMT-G2"
	X3_FTH            InverterTypeCode = "X3-FTH"
	X3_MGA_G2         InverterTypeCode = "X3-MGA-G2"
)

var (
	inverterTypeMap map[string]InverterTypeCode = map[string]InverterTypeCode{
		"1":  X1_LX,
		"2":  X_Hybrid,
		"3":  X1_Hybriyd_Fit,
		"4":  X1_Boost_Air_Mini,
		"5":  X3_Hybriyd_Fit,
		"6":  X3_20K_30K,
		"7":  X3_MIC_PRO,
		"8":  X1_Smart,
		"9":  X1_AC,
		"10": A1_Hybrid,
		"11": A1_Fit,
		"12": A1_Grid,
		"13": J1_ESS,
		"14": X3_Hybrid_G4,
		"15": X1_Hybrid_G4,
		"16": X3_MIC_PRO_G2,
		"17": X1_SPT,
		"18": X1_Boost_Mini_G4,
		"19": A1_HYB_G2,
		"20": A1_AC_G2,
		"21": A1_SMT_G2,
		"22": X3_FTH,
		"23": X3_MGA_G2,
	}
)

func (i InverterTypeCode) String() string {
	return string(i)
}

func InverterTypeFromString(s string) InverterTypeCode {
	return inverterTypeMap[s]
}

func DescribeInverterTypeCode(name string, dynamicLabels []string, labels prometheus.Labels) *metrics.EnumCollector {
	descs := make(metrics.EnumCollector)
	for _, code := range inverterTypeMap {
		valueLabels := prometheus.Labels{
			"value": code.String(),
		}
		for k, v := range labels {
			valueLabels[k] = v
		}
		descs[code.String()] = prometheus.NewDesc(name, "Inverter type.", dynamicLabels, valueLabels)
	}
	return &descs
}
