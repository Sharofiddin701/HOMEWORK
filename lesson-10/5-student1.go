package main

import "fmt"

func (b *Talaba) Summa() int {
	s := 0

	if b.Kurs == 2 {
		s = b.Stipendiya
	}
	return s
}

type Talaba struct {
	Ism        string
	Stipendiya int
	Kurs       int
}

func main() {
	talaba1 := Talaba{

		Ism:        "Baxtiyor",
		Stipendiya: 50000,
		Kurs:       3,
	}
	talaba2 := Talaba{

		Ism:        "Bobur",
		Stipendiya: 500000,
		Kurs:       2,
	}
	talaba3 := Talaba{

		Ism:        "Shavkat",
		Stipendiya: 500000,
		Kurs:       2,
	}
	talaba4 := Talaba{

		Ism:        "Erali",
		Stipendiya: 500000,
		Kurs:       2,
	}
	talaba5 := Talaba{

		Ism:        "Farrux",
		Stipendiya: 400000,
		Kurs:       2,
	}
	talaba6 := Talaba{

		Ism:        "Shuxrat",
		Stipendiya: 550000,
		Kurs:       4,
	}
	talaba7 := Talaba{

		Ism:        "Bektemir",
		Stipendiya: 400000,
		Kurs:       2,
	}
	talaba8 := Talaba{

		Ism:        "Lochin",
		Stipendiya: 700000,
		Kurs:       2,
	}
	talaba9 := Talaba{

		Ism:        "Baxtiyor",
		Stipendiya: 50000,
		Kurs:       3,
	}
	talaba10 := Talaba{

		Ism:        "Lola",
		Stipendiya: 600000,
		Kurs:       2,
	}

	var summa int

	summa = talaba1.Summa() + talaba2.Summa() + talaba3.Summa() + talaba4.Summa() + talaba5.Summa() + talaba6.Summa() + talaba7.Summa() + talaba8.Summa() + talaba9.Summa() + talaba10.Summa()
	fmt.Print(summa)

}
