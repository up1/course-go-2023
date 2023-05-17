package hello

type THB int

type Demo struct {
	Name  string
	Price THB
}

func (d Demo) SayHello() {
	i, err := doSth()
	if err != nil {
		println(err)
	}
	println("Hello World ", i)
}

func NewDemo(name string) Demo {
	return Demo{
		Name:  name,
		Price: THB(1000),
	}
}

func doSth() (int, error) {
	return 1, nil
}

func tryLoop() {
	for i := 0; i < 2; i++ {

	}

	i := 0
	for i < 2 {
		i++
	}

	for {

	}
}
