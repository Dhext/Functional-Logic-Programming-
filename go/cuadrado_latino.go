package main

import(
  "fmt"
)

func main(){

  numero := 0;
  fmt.Println("Ingrese un numero")
  fmt.Scanf("%d\n", &numero)

  for i := 0; i < numero; i++ {
    fmt.Println()
    for j := numero - i; j < numero ; j++ {
      fmt.Printf( "%d ", j + 1 )
    }
    for k := 0; k < numero - i; k++ {
      fmt.Printf( "%d ", k + 1 )
    }
  }

}
