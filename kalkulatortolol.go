package main

import "fmt"

func main() {

	var (
		angka    = 0
		hasil    = 0
		operator = ""
	)

	fmt.Println("Kalkulator--------------------")
	fmt.Print("Masukkan Angka: ")
	fmt.Scanln(&angka)

	hasil += angka

	for {

		fmt.Print("Masukkan Operator[+,-,*,/,%,=]: ")
		fmt.Scanln(&operator)

		if operator == "=" {
			break
		}

		fmt.Print("Masukkan Angka: ")
		fmt.Scanln(&angka)

		switch operator {
		case "+":
			hasil += angka
		case "-":
			hasil -= angka
		case "*":
			hasil *= angka
		case "/":
			hasil /= angka
		case "%":
			hasil %= angka
		}

	}
	fmt.Println("Hasil:", hasil)
}
