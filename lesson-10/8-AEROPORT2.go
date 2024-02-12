package main

import "fmt"

type AEROPORT1 struct {
	Samolyot_turi      string
	Parvoz_malumotlari int
	Shahardan          string
	Shaharga           string
	Parvoz_vaqti       int
}

func (b *AEROPORT1) par() string {

	if b.Parvoz_vaqti == 2 || b.Parvoz_vaqti == 3 && b.Shaharga == "Toshkent" {

		return fmt.Sprintf("%v ", b.Parvoz_malumotlari, b.Shaharga)
	}
	return "Bu parvoz to'g'ri kelmaydi !!!"
}
func main() {

	Parvoz1 := AEROPORT1{

		Samolyot_turi:      "Boing 777",
		Parvoz_malumotlari: 8,
		Shahardan:          "Rossiya",
		Shaharga:           "Toshkent",
		Parvoz_vaqti:       3,
	}
	Parvoz2 := AEROPORT1{

		Samolyot_turi:      "Boing 778",
		Parvoz_malumotlari: 2,
		Shahardan:          "Toshkent",
		Shaharga:           "Mexika",
		Parvoz_vaqti:       4,
	}
	Parvoz3 := AEROPORT1{

		Samolyot_turi:      "Boing 779",
		Parvoz_malumotlari: 6,
		Shahardan:          "Bishkek",
		Shaharga:           "Toshkent",
		Parvoz_vaqti:       2,
	}
	Parvoz4 := AEROPORT1{

		Samolyot_turi:      "Boing 780",
		Parvoz_malumotlari: 4,
		Shahardan:          "Parij",
		Shaharga:           "Toshkent",
		Parvoz_vaqti:       3,
	}
	Parvoz5 := AEROPORT1{

		Samolyot_turi:      "Boing 781",
		Parvoz_malumotlari: 7,
		Shahardan:          "Toshkent",
		Shaharga:           "Xitoy",
		Parvoz_vaqti:       7,
	}
	Parvoz6 := AEROPORT1{

		Samolyot_turi:      "Boing 782",
		Parvoz_malumotlari: 6,
		Shahardan:          "Toshkent",
		Shaharga:           "Rossiya",
		Parvoz_vaqti:       8,
	}
	Parvoz7 := AEROPORT1{

		Samolyot_turi:      "Boing 771",
		Parvoz_malumotlari: 7,
		Shahardan:          "Abudabi",
		Shaharga:           "Toshkent",
		Parvoz_vaqti:       2,
	}
	Parvoz8 := AEROPORT1{

		Samolyot_turi:      "Boing 770",
		Parvoz_malumotlari: 8,
		Shahardan:          "Toshkent",
		Shaharga:           "Tojmahal",
		Parvoz_vaqti:       7,
	}
	Parvoz9 := AEROPORT1{

		Samolyot_turi:      "Boing 707",
		Parvoz_malumotlari: 9,
		Shahardan:          "London",
		Shaharga:           "Toshkent",
		Parvoz_vaqti:       3,
	}
	Parvoz10 := AEROPORT1{

		Samolyot_turi:      "Boing 700",
		Parvoz_malumotlari: 10,
		Shahardan:          "Toshkent",
		Shaharga:           "London",
		Parvoz_vaqti:       4,
	}
	fmt.Println(Parvoz1.par())
	fmt.Println(Parvoz2.par())
	fmt.Println(Parvoz3.par())
	fmt.Println(Parvoz4.par())
	fmt.Println(Parvoz5.par())
	fmt.Println(Parvoz6.par())
	fmt.Println(Parvoz7.par())
	fmt.Println(Parvoz8.par())
	fmt.Println(Parvoz9.par())
	fmt.Println(Parvoz10.par())

}
