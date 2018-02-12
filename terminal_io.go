package main

import (
	"fmt"
)

func main() {
	edad:=22
	nombre:="angel"
	bandera:=true
	precio:=14.333

	var ed int
	var nom string


	fmt.Printf("Hola mundo %s %d %t %f\n",nombre,edad,bandera,precio)
	
	fmt.Printf("Leyendo desde consola:\n")
	fmt.Printf("Nombre:\n")
	fmt.Scanf("%s\n",&nom)
	
	fmt.Printf("Edad:\n")
	fmt.Scanf("%d\n",&ed)

	fmt.Printf("mis datos son:\nNombre:%v\nEdad:%d\n",nom,ed)
}