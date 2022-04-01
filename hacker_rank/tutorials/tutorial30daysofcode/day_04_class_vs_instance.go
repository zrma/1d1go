package tutorial30daysofcode

import "fmt"

type simplePerson struct {
	age int
}

const (
	msgOfInvalid  = "Age is not valid, setting age to 0."
	msgOfYoung    = "You are young."
	msgOfTeenager = "You are a teenager."
	msgOfOld      = "You are old."
)

func (p simplePerson) NewPerson(initialAge int) simplePerson {
	if initialAge < 0 {
		fmt.Println(msgOfInvalid)
		initialAge = 0
	}
	res := p
	res.age = initialAge
	return res
}

func (p simplePerson) amIOld() {
	switch {
	case p.age < 13:
		fmt.Println(msgOfYoung)
	case p.age < 18:
		fmt.Println(msgOfTeenager)
	default:
		fmt.Println(msgOfOld)
	}
}

func (p simplePerson) yearPasses() simplePerson {
	p.age++

	return p
}

func classAndInstance(age int) {
	p := simplePerson{age: age}
	p = p.NewPerson(age)
	p.amIOld()
	for i := 0; i < 3; i++ {
		p = p.yearPasses()
	}
	p.amIOld()
	fmt.Println()
}
