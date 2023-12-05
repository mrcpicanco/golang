package main

import "fmt"

const ebulicaoF float64 = 212.0 //Fora do code block

func main() {
	tempF := ebulicaoF            //dentro do code block
	tempC := (tempF - 32) * 5 / 9 //dentro do code block
	fmt.Println("A temperatura de ebulição da água em F é: ", tempF)
	fmt.Println("A temperatura de ebulição da água em graus celsius é:", tempC)
}
