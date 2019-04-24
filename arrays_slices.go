package main

import(

	"fmt"

)


func main(){
	//declaração de um array
	var meu_array [10]int
	fmt.Println(meu_array)

	//declaração curta do array
	array_curto := [5]int{}
	fmt.Println(array_curto)

	//declaração com inicialização do array
	array_inicializado := [5]int{0,1,2,3,4}
	fmt.Println(array_inicializado)

	//declaração com inicialização de lementos especificos
	array_especifico := [5]int{0:1,2:3,4:7}
	fmt.Println(array_especifico)

	//alterando valor de um indice
	meu_array[0]= 20
	fmt.Println(meu_array)





}