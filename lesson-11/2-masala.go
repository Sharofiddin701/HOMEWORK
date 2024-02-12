package main

import "fmt"

type AEROPORT struct {
	Samolyot_turi      string
	Parvoz_malumotlari int
	Shahardan          string
	Shaharga           string
	Parvoz_vaqti       int
}

func (b *AEROPORT) par() string {

	if b.Parvoz_malumotlari == 6 || b.Parvoz_malumotlari == 7 || b.Parvoz_malumotlari == 8 {

		return fmt.Sprintf("%v ", b.Parvoz_malumotlari)
	}
	return "Bu parvoz yoz oylarida emas !!!"
}
func main() {

	Parvoz1 := AEROPORT{

		Samolyot_turi:      "Boing 777",
		Parvoz_malumotlari: 8,
		Shahardan:          "Toshkent",
		Shaharga:           "Parij",
		Parvoz_vaqti:       3,
	}
	Parvoz2 := AEROPORT{

		Samolyot_turi:      "Boing 778",
		Parvoz_malumotlari: 2,
		Shahardan:          "Toshkent",
		Shaharga:           "Mexika",
		Parvoz_vaqti:       4,
	}
	Parvoz3 := AEROPORT{

		Samolyot_turi:      "Boing 779",
		Parvoz_malumotlari: 6,
		Shahardan:          "Toshkent",
		Shaharga:           "Bishken",
		Parvoz_vaqti:       6,
	}
	Parvoz4 := AEROPORT{

		Samolyot_turi:      "Boing 780",
		Parvoz_malumotlari: 4,
		Shahardan:          "Toshkent",
		Shaharga:           "Kaliforniya",
		Parvoz_vaqti:       5,
	}
	Parvoz5 := AEROPORT{

		Samolyot_turi:      "Boing 781",
		Parvoz_malumotlari: 7,
		Shahardan:          "Toshkent",
		Shaharga:           "Xitoy",
		Parvoz_vaqti:       7,
	}
	Parvoz6 := AEROPORT{

		Samolyot_turi:      "Boing 782",
		Parvoz_malumotlari: 6,
		Shahardan:          "Toshkent",
		Shaharga:           "Rossiya",
		Parvoz_vaqti:       8,
	}
	Parvoz7 := AEROPORT{

		Samolyot_turi:      "Boing 771",
		Parvoz_malumotlari: 7,
		Shahardan:          "Toshkent",
		Shaharga:           "Amerika",
		Parvoz_vaqti:       7,
	}
	Parvoz8 := AEROPORT{

		Samolyot_turi:      "Boing 770",
		Parvoz_malumotlari: 8,
		Shahardan:          "Toshkent",
		Shaharga:           "Tojmahal",
		Parvoz_vaqti:       2,
	}
	Parvoz9 := AEROPORT{

		Samolyot_turi:      "Boing 707",
		Parvoz_malumotlari: 9,
		Shahardan:          "Toshkent",
		Shaharga:           "Moskva",
		Parvoz_vaqti:       3,
	}
	Parvoz10 := AEROPORT{

		Samolyot_turi:      "Boing 700",
		Parvoz_malumotlari: 10,
		Shahardan:          "Toshkent",
		Shaharga:           "London",
		Parvoz_vaqti:       4,
	}
	fmt.Println("Yoz faslida", Parvoz1.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz2.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz3.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz4.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz5.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz6.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz7.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz8.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz9.par(), "- oyida bor")
	fmt.Println("Yoz faslida", Parvoz10.par(), "- oyida bor")

}