package main

import (
"fmt"
"time"

)


func Primos1Incial(num1 int){

	var res int
	var nc int
	for i:=1;i<=num1;i++{
		for j:=1;j<=i;j++{
			res=i%j
			if res==0{
				nc++
			}
		}
		if nc==2{
			fmt.Println("-",i)
		}
		nc=0
	}
}
func PrimosIncialFinal(num1,num2 int){
	
	var res int
	var nc int
	for i:=num1;i<=num2;i++{
		for j:=1;j<=i;j++{
			res=i%j
			if res==0{
				nc++
			}
		}
		if nc==2{
			fmt.Println("-",i)
		}
		nc=0
	}
}

func Primos1_1m(){
	var res int
	var nc int
	for i:=1;i<=10000000;i++{
		for j:=1;j<=i;j++{
			res=i%j
			if res==0{
				nc++
			}
		}
		if nc==2{
			fmt.Println("-",i)
		}
		nc=0
	}

}

func primo(num int)(bool){
	isPrimo := true
	for i := 2; i <  num && isPrimo==true; i++ {
		if num%i==0{
			isPrimo=false
		}
	}
	return isPrimo
}

func RutPimos(){
	var limiteSuperior int
	var Inicio int
	
	fmt.Println("Ingresa El Inicio: ")
	fmt.Scanf("%d",)
	fmt.Scanf("%d",&Inicio)
	fmt.Println("Ingresa El Limite Superior: ")
	fmt.Scanf("%d",)
	fmt.Scanf("%d",&limiteSuperior)
	go PrimosIncialFinal(Inicio,limiteSuperior)
	go Primos1Incial(Inicio)
	<-time.After(time.Second * 5)
}

func menu(){
	fmt.Println("0 Calcular los numeros primos de un rango 1 al rango final usando concurrencia")
	fmt.Println("1 Calcular si un numero es primo o no")
	fmt.Println("2 calcular numeros primos del 1 al 1000000")
}

func main() {
	var opc=0;
	menu()
	fmt.Scanf("%d",&opc)
	switch opc {
	case 0:
		RutPimos()
	case 1:
		var num int
		fmt.Println("INGRESA UN NUMERO: ")
		fmt.Scanf("&d",)
		fmt.Scanf("%d",&num)
		if primo(num)==false{
			fmt.Print("EL numero ",num)
			fmt.Println(", NO es Primo")
		}else{
			fmt.Print("EL numero ",num)
			fmt.Println(", SI es Primo")
		}
	case 2:
		Primos1_1m();



		}
}



