package main

import (
	"fmt"
)

func main() {
	// menampilkan nilai i : 21 fmt.Printf("%T \n", i) // contoh : fmt.Printf("%v \n", i)
	var i int = 21

	fmt.Printf("%v \n", i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)

	//menampilkan tanda %
	fmt.Printf("%% \n")

	// menampilkan nilai boolean j : true
	var j bool = true
	fmt.Printf("%t \n", j)

	// menampilkan nilai boolean j : true
	fmt.Printf("\n%b \n", i)

	// menampilkan unicode russia : Я (ya)
	fmt.Printf("%c \n", '\u042F')

	// menampilkan nilai base 10 : 21
	fmt.Printf("%d\n", i)

	// menampilkan nilai base 8 :25
	fmt.Printf("%o\n", i)

	// menampilkan nilai base 16 : f
	fmt.Printf("%x\n", 15)

	// menampilkan nilai base 16 : F 13
	fmt.Printf("%X\n", 15)

	// menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
	fmt.Printf("%U \n", 'Я')

	// menampilkan float : 123.456000
	var bil float64 = 123.456000
	fmt.Printf("%f \n", bil)

	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%E \n", bil)
}
