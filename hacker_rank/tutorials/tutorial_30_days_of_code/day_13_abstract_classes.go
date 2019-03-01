package tutorial_30_days_of_code

import "fmt"

type Viewer interface {
	display()
}

type Book struct {
	title  string
	author string

	Viewer
}

func NewBook(title, author string) *Book {
	return &Book{title, author, nil}
}

type MyBook struct {
	*Book

	price int
}

func NewMyBook(title, author string, price int) *MyBook {
	return &MyBook{
		NewBook(title, author), price,
	}
}

func (b MyBook) display() {
	fmt.Println("Title:", b.title)
	fmt.Println("Author:", b.author)
	fmt.Println("Price:", b.price)
}

func AbstractClasses(title, author string, price int) {
	var v Viewer
	v = NewMyBook(title, author, price)
	v.display()
}
