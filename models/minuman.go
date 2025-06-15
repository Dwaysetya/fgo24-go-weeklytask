package models


type Minuman struct{
	Name string
	Harga int
}

func (d Minuman) GetName()string {return d.Name}
func (d Minuman) GetHarga()int {return d.Harga}

func NewMinuman (name string, harga int)Minuman{
	return Minuman{Name: name, Harga: harga}
}

var daftarMinuman = []Item{
	NewMinuman("Es Teh Manis", 5000),
	NewMinuman("Teh Tarik", 7000),
	NewMinuman("Teh Hijau", 8000),
	NewMinuman("Es Jeruk", 6000),
	NewMinuman("Air Mineral", 4000),

	NewMinuman("Kopi Hitam", 7000),
	NewMinuman("Kopi Susu", 8000),
	NewMinuman("Cappuccino", 10000),
	NewMinuman("Latte", 10000),
	NewMinuman("Americano", 9000),
	NewMinuman("Espresso", 8500),
	NewMinuman("Mochaccino", 11000),

	NewMinuman("Es Coklat", 9000),
	NewMinuman("Coklat Hazelnut", 11000),
	NewMinuman("Thai Tea", 10000),
	NewMinuman("Green Tea Latte", 10500),
	NewMinuman("Taro Latte", 9500),
	NewMinuman("Boba Brown Sugar", 12000),
	NewMinuman("Boba Milk Tea", 11000),
	NewMinuman("Red Velvet Latte", 10500),
}


func GetAllMinuman() []Item {
	return daftarMinuman
}