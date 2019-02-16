package university_codesprint_5

import "fmt"

func exceedingTheSpeedLimit(s int32) {
	const (
		MsgLineRemoved  = "License removed"
		MsgWarning      = "Warning"
		MsgNoPunishment = "No punishment"
	)

	switch {
	case s > 110:
		fine := (s - 90) * 500
		fmt.Println(fine, MsgLineRemoved)
	case s > 90:
		fine := (s - 90) * 300
		fmt.Println(fine, MsgWarning)
	default:
		fmt.Println(0, MsgNoPunishment)
	}
}

func ExceedingTheSpeedLimit(s int32) {
	exceedingTheSpeedLimit(s)
}
