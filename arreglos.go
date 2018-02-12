package main

import "fmt"

func main() {
	var arreglo_0 [20] int
	arreglo_1:= [10] int{1,2,3,4,5}

	var matriz_0 [3][2]int

	fmt.Println(arreglo_0)
	fmt.Println("le caben :",len(arreglo_0))

	fmt.Println(arreglo_1)
	fmt.Println("le caben :",len(arreglo_1))

	arreglo_1[9]=10

	for i:=0; i<len(arreglo_1); i++{
		fmt.Println(arreglo_1[i])
	}

	matriz_0[0][1]=10

	fmt.Println(matriz_0)

}