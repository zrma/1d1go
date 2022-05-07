package codesprint5

import "fmt"

func exceedingTheSpeedLimit(n int32) {
	const (
		MsgLineRemoved  = "License removed"
		MsgWarning      = "Warning"
		MsgNoPunishment = "No punishment"
	)

	switch {
	case n > 110:
		fine := (n - 90) * 500
		fmt.Println(fine, MsgLineRemoved)
	case n > 90:
		fine := (n - 90) * 300
		fmt.Println(fine, MsgWarning)
	default:
		fmt.Println(0, MsgNoPunishment)
	}
}
