package main

import "fmt"

func main() {
	slice:=make([]int,3,5)//tipo,longitud,capacidad

	slice = append(slice,2)//agregar dato

	fmt.Println(slice ," slice")
	fmt.Println(len(slice) ," longitud")
	fmt.Println(cap(slice) ," capacidad")
}