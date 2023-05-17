package main

func main() {
	v1 := V1{}
	process(v1)
}

func process(i IDo) {
	i.f1()
}

type IDo interface {
	f1() error
	f2() error
}
type V1 struct{}

func (v1 V1) f1() error {
	return nil
}
func (v1 V1) f2() error {
	return nil
}
