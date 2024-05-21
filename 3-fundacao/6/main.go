package main

import "fmt"

// Slices por baixo dos panos são arrays.
func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // colchete em branco: sabe-se que é slice (poderia estar até vazio)
	fmt.Println("fmt.Printf(len=%d cap=%d %v, len(s), cap(s), s)")
	fmt.Printf("len=%d cap=%d %v\n\n", len(s), cap(s), s)

	fmt.Println("fmt.Printf(len=%d cap=%d %v, len(s[:0]), cap(s[:0]), s[:0])")
	fmt.Printf("len=%d cap=%d %v\n\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Println("fmt.Printf(len=%d cap=%d %v, len(s[:4]), cap(s[:4]), s[:4])")
	fmt.Printf("len=%d cap=%d %v\n\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Println("fmt.Printf(len=%d cap=%d %v, len(s[2:]), cap(s[2:]), s[2:])")
	fmt.Printf("len=%d cap=%d %v\n\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)

	fmt.Println("Após append do valor 110: fmt.Printf(len=%d cap=%d %v, len(s), cap(s), s)")
	fmt.Printf("len=%d cap=%d %v\n\n", len(s), cap(s), s)
}

// output
// fmt.Printf(len=%d cap=%d %v, len(s), cap(s), s)
// len=10 cap=10 [10 20 30 40 50 60 70 80 90 100]

// fmt.Printf(len=%d cap=%d %v, len(s[:0]), cap(s[:0]), s[:0])
// len=0 cap=10 []

// fmt.Printf(len=%d cap=%d %v, len(s[:4]), cap(s[:4]), s[:4])
// len=4 cap=10 [10 20 30 40]

// fmt.Printf(len=%d cap=%d %v, len(s[2:]), cap(s[2:]), s[2:])
// len=8 cap=8 [30 40 50 60 70 80 90 100]

// Após append do valor 110: fmt.Printf(len=%d cap=%d %v, len(s), cap(s), s)
// No append, por baixo dos panos o go dobra a capacidade, ficar ligado
// len=11 cap=20 [10 20 30 40 50 60 70 80 90 100 110]
