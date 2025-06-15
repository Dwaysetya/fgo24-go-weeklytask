package models

type Minuman struct{
	Name string
	Harga int64
}

func (d Minuman) GetName()string {return d.Name}
func (d Minuman) GetHarga()int64 {return d.Harga}

func NewMinuman (name string, harga int64)Minuman{
	return Minuman{Name: name, Harga: harga}
}