package codesprint5

func exceedingTheSpeedLimit(n int32) {
	const (
		MsgLineRemoved  = "License removed"
		MsgWarning      = "Warning"
		MsgNoPunishment = "No punishment"
	)

	switch {
	case n > 110:
		fine := (n - 90) * 500
		_, _ = funcPrintf("%d %s", fine, MsgLineRemoved)
	case n > 90:
		fine := (n - 90) * 300
		_, _ = funcPrintf("%d %s", fine, MsgWarning)
	default:
		_, _ = funcPrintf("%d %s", 0, MsgNoPunishment)
	}
}
