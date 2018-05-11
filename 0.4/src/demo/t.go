package demo

type T struct {
	kalasa string
	Name   string
}

func (t *T) SetName(name string) {
	t.kalasa = name
}
