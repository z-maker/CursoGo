package main

import "fmt"

func main() {

	i:=0

	for i:=1; i<=10; i++{
		fmt.Println("x")
	}

	for{
		i++
		if i%2==0{
			continue
		}
		fmt.Println(i)
		if i>20{
			break
		}
	}
	
}