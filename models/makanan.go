package models

type Makanan struct {
	Name string
	Harga int64
}


func (m Makanan) GetName()string {return m.Name}
func (m Makanan) GetHarga()int64 {return m.Harga}

func NewMakanan(name string, harga int64)Makanan{
	return Makanan{Name: name, Harga: harga}
}
