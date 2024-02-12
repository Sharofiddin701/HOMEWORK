package main

import "fmt"

type Talaba1 struct {
	Ism        string
	Stipendiya int
	Kurs       int
}

func (b *Talaba1) kam() string {

	if len(b.Ism) <= 5 {
		return b.Ism
	}
	return " "
}
func main() {
	talaba1 := Talaba1{

		Ism:        "Baxti",
		Stipendiya: 50000,
		Kurs:       3,
	}
	talaba2 := Talaba1{

		Ism:        "Bobur",
		Stipendiya: 500000,
		Kurs:       2,
	}
	talaba3 := Talaba1{

		Ism:        "Shavkat",
		Stipendiya: 500000,
		Kurs:       2,
	}
	talaba4 := Talaba1{

		Ism:        "Erali",
		Stipendiya: 500000,
		Kurs:       2,
	}
	talaba5 := Talaba1{

		Ism:        "Farrux",
		Stipendiya: 400000,
		Kurs:       2,
	}
	talaba6 := Talaba1{

		Ism:        "Shuxrat",
		Stipendiya: 550000,
		Kurs:       4,
	}
	talaba7 := Talaba1{

		Ism:        "Bektemir",
		Stipendiya: 400000,
		Kurs:       2,
	}
	talaba8 := Talaba1{

		Ism:        "Lochin",
		Stipendiya: 700000,
		Kurs:       2,
	}
	talaba9 := Talaba1{

		Ism:        "Baxtiyor",
		Stipendiya: 50000,
		Kurs:       3,
	}
	talaba10 := Talaba1{

		Ism:        "Lola",
		Stipendiya: 600000,
		Kurs:       2,
	}
	fmt.Println(talaba1.kam())
	fmt.Println(talaba2.kam())
	fmt.Println(talaba3.kam())
	fmt.Println(talaba4.kam())
	fmt.Println(talaba5.kam())
	fmt.Println(talaba6.kam())
	fmt.Println(talaba7.kam())
	fmt.Println(talaba8.kam())
	fmt.Println(talaba9.kam())
	fmt.Println(talaba10.kam())
}
