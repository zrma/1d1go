package tutorial_30_days_of_code

import "fmt"

type person struct {
	age int
}

const (
	MsgOfInvalid  = "Age is not valid, setting age to 0."
	MsgOfYoung    = "You are young."
	MsgOfTeenager = "You are a teenager."
	MsgOfOld      = "You are old."
)

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		fmt.Println(MsgOfInvalid)
		initialAge = 0
	}

	p.age = initialAge
	return p
}

func (p person) amIOld() {
	switch {
	case p.age < 13:
		fmt.Println(MsgOfYoung)
	case p.age < 18:
		fmt.Println(MsgOfTeenager)
	default:
		fmt.Println(MsgOfOld)
	}
}

func (p person) yearPasses() person {
	p.age++

	return p
}

func ClassAndInstance(age int) {
	p := person{age: age}
	p = p.NewPerson(age)
	p.amIOld()
	for i := 0; i < 3; i++ {
		p = p.yearPasses()
	}
	p.amIOld()
	fmt.Println()
}
