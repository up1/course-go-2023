package main

func main() {
	v1 := V1{}
	process1(v1)
	process2(v1)
}

func process1(i IDo) {
	i.f1()
}

func process2(i interface{}) {
	i.(IDo).f1()
}

func process3(i any) {
	i.(IDo).f1()
}

type IDo2 interface {
	f3() error
}

type IDo interface {
	f1() error
	f2() error
	IDo2
}
type V1 struct{}

func (v1 V1) f1() error {
	return nil
}
func (v1 V1) f2() error {
	return nil
}
