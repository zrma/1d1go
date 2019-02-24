package tutorial_30_days_of_code

import "fmt"

type Person struct {
	firstName string
	lastName  string
	id        int
}

func NewPerson(firstName, lastName string, identification int) *Person {
	return &Person{firstName, lastName, identification}
}

func (p Person) printPerson() string {
	return fmt.Sprintf("Name: %s, %s\nID: %d", p.lastName, p.firstName, p.id)
}

type Student struct {
	*Person

	testScores []int
}

func NewStudent(firstName, lastName string, identification int, scores []int) *Student {
	return &Student{
		NewPerson(firstName, lastName, identification),
		scores,
	}
}

func (s Student) calculate() string {
	l := len(s.testScores)
	sum := 0
	for _, s := range s.testScores {
		sum += s
	}

	avr := sum / l
	switch {
	case avr >= 90:
		return "O"
	case avr >= 80:
		return "E"
	case avr >= 70:
		return "A"
	case avr >= 55:
		return "P"
	case avr >= 40:
		return "D"
	default:
		return "T"
	}
}
