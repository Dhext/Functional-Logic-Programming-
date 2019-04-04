package main

import(
  "fmt"
)

func mcd(a int, b int) int {
  t := 0
  for a > 0 {
    t = a
    a = b % a
    b = t
  }

  return b
}

func main(){
  a := 0
  b := 0

  fmt.Print("Ingrese el primer numero: ")
  fmt.Scanf("%d\n", &a)
  fmt.Print("Ingrese el segundo numero: ")
  fmt.Scanf("%d\n", &b)

  fmt.Printf("\nEl CMD de %d y %d es: %d", a, b, mcd(a, b))

}
