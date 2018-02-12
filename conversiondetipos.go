package main

import(
	 "fmt"
	 "strconv"
)

func main() {
	edad :=22

	edad_str:= strconv.Itoa(edad)

	//una fincion puede retornar multipes valores
	edad_int,_:= strconv.Atoi("22") //el operador _ recibe el valor pero lo desecha

	fmt.Println("Mi edad es "+edad_str)
	fmt.Print("Mi edad es ")
	fmt.Print(edad_int+10)
}