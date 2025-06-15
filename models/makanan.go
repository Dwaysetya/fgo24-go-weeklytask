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

var daftarMakanan = []Item{
	NewMakanan("Nasi Goreng", 18000),
	NewMakanan("Nasi Padang", 15000),
	NewMakanan("Nasi Pecel", 12000),
	NewMakanan("Nasi Rames", 13000),
	NewMakanan("Nasi Ayam", 17000),
	NewMakanan("Nasi Telur", 10000),
	NewMakanan("Nasi Bakar", 12000),
	NewMakanan("Bakso", 14000),
	NewMakanan("Mie Ayam", 13000),
	NewMakanan("Sate Ayam", 19000),
}

func GetAllMakanan() []Item {
	return daftarMakanan
}


