package models

type Snack struct {
	Name string
	Harga int64
}

func (s Snack) GetName()string {return s.Name}
func (s Snack) GetHarga()int64 {return s.Harga}

func NewSnack(name string, harga int64)Snack{
	return Snack{Name: name, Harga: harga}
}

