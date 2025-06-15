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

var daftarMinuman = []Item{
	NewMinuman("Es Teh Manis", 5000),
	NewMinuman("Es Jeruk", 6000),
	NewMinuman("Kopi Hitam", 7000),
	NewMinuman("Kopi Susu", 8000),
	NewMinuman("Air Mineral", 4000),
}

func GetAllMinuman() []Item {
	return daftarMinuman
}