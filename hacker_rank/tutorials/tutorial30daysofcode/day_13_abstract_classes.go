package tutorial30daysofcode

type viewer interface {
	display()
}

type book struct {
	title  string
	author string

	viewer
}

func newBook(title, author string) *book {
	return &book{title, author, nil}
}

type myBook struct {
	*book

	price int
}

func newMyBook(title, author string, price int) *myBook {
	return &myBook{
		newBook(title, author), price,
	}
}

func (b myBook) display() {
	_, _ = fmtPrintln("Title:", b.title)
	_, _ = fmtPrintln("Author:", b.author)
	_, _ = fmtPrintln("Price:", b.price)
}

func abstractClasses(title, author string, price int) {
	v := newMyBook(title, author, price)
	v.display()
}
