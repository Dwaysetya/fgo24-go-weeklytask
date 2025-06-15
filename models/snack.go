package models

type Snack struct {
	Name string
	Harga int
}

func (s Snack) GetName()string {return s.Name}
func (s Snack) GetHarga()int {return s.Harga}

func NewSnack(name string, harga int)Snack{
	return Snack{Name: name, Harga: harga}
}

var daftarSnack = []Item{
	NewSnack("Keripik Singkong", 5000),
	NewSnack("Keripik Kentang", 6000),
	NewSnack("Roti Bakar", 8000),
	NewSnack("Pisang Goreng", 6000),
	NewSnack("Pisang Coklat Keju", 9000),
	NewSnack("Cireng", 7000),
	NewSnack("Cimol", 7000),
	NewSnack("Seblak", 10000),
	NewSnack("Sosis Bakar", 8000),
	NewSnack("Kentang Goreng", 9000),
	NewSnack("Bakwan", 5000),
	NewSnack("Tahu Crispy", 6000),
	NewSnack("Otak-Otak", 8000),
	NewSnack("Singkong Keju", 7000),
	NewSnack("Tahu Bulat", 5000),
	NewSnack("Martabak Mini", 9000),
	NewSnack("Kue Cubit", 7000),
	NewSnack("Telur Gulung", 6000),
	NewSnack("Churros", 10000),
	NewSnack("Takoyaki", 11000),
}


func GetAllSnack() []Item {
	return daftarSnack
}


