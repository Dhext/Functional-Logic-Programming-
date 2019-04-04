package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"
)

func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}

func palindromo(clave string)(bool){
	var remplace string
	var pib1 string;
	var pib2 string;
	remplace = Reverse(clave)
	pib1=strings.ToLower(strings.Replace(clave," ","",-1))
	pib2=strings.ToLower( strings.Replace(remplace," ","",-1))
	if pib1 == pib2 {
		return true
	}
	return false
}

func main()  {
	var clave string ;
	fmt.Println("introduce una palabra: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		clave = scanner.Text()
	}
	if  palindromo(clave){
		println(clave + " Es palindormo")
	}else{
		println(clave + " No es palindormo")
	}
}