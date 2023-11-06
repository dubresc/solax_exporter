package api

type InverterStatusCode string

const (
	Undefined          InverterStatusCode = ""
	WaitMode           InverterStatusCode = "100"
	CheckMode          InverterStatusCode = "101"
	NormalMode         InverterStatusCode = "102"
	FaultMode          InverterStatusCode = "103"
	PermanentFaultMode InverterStatusCode = "104"
	UpdateMode         InverterStatusCode = "105"
	EPSCheckMode       InverterStatusCode = "106"
	EPSMode            InverterStatusCode = "107"
	SelfTestMode       InverterStatusCode = "108"
	IdleMode           InverterStatusCode = "109"
	StandbyMode        InverterStatusCode = "110"
	PvWakeUpBatMode    InverterStatusCode = "111"
	GenCheckMode       InverterStatusCode = "112"
	GenRunMode         InverterStatusCode = "113"
)

func (i InverterStatusCode) String() string {
	switch i {
	case WaitMode:
		return "wait_mode"
	case CheckMode:
		return "check_mode"
	case NormalMode:
		return "normal_mode"
	case FaultMode:
		return "fault_mode"
	case PermanentFaultMode:
		return "permanent_fault_mode"
	case UpdateMode:
		return "update_mode"
	case EPSCheckMode:
		return "eps_check_mode"
	case EPSMode:
		return "eps_mode"
	case SelfTestMode:
		return "self_test_mode"
	case IdleMode:
		return "idle_mode"
	case StandbyMode:
		return "standby_mode"
	case PvWakeUpBatMode:
		return "pv_wake_up_bat_mode"
	case GenCheckMode:
		return "gen_check_mode"
	case GenRunMode:
		return "gen_run_mode"
	}
	return "undefined"
}

func InverterStatusCodeFromString(s string) InverterStatusCode {
	switch s {
	case string(WaitMode):
		return WaitMode
	case string(CheckMode):
		return CheckMode
	case string(NormalMode):
		return NormalMode
	case string(FaultMode):
		return FaultMode
	case string(PermanentFaultMode):
		return PermanentFaultMode
	case string(UpdateMode):
		return UpdateMode
	case string(EPSCheckMode):
		return EPSCheckMode
	case string(EPSMode):
		return EPSMode
	case string(SelfTestMode):
		return SelfTestMode
	case string(IdleMode):
		return IdleMode
	case string(StandbyMode):
		return StandbyMode
	case string(PvWakeUpBatMode):
		return PvWakeUpBatMode
	case string(GenCheckMode):
		return GenCheckMode
	case string(GenRunMode):
		return GenRunMode
	}
	return Undefined
}
