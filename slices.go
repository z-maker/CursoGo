package main

import "fmt"

func main() {
	slice := []int{1,2,3,4}

	var slice_1 []int

	arreglo:=[3]int{1,2,3}
	
	slice_2:=arreglo[1:2]//se puede hacer esto para partir el slice


	fmt.Println("Slice 0",slice)

	if slice_1 == nil{
		fmt.Println("slice 1 esta vacio")
	}else{
		fmt.Println(slice_1)
	}
	fmt.Print("de arreglo : ",arreglo)
	fmt.Println(" a slice: ",slice_2)

	/*
	slices:
	puntero al arreglo
	longitud de arreglo
	capacidad
	*/


}