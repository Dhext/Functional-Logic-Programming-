package main

import (
	"database/sql"
	"log"
	"net/http"
    "fmt"
	"html/template"
	_ "github.com/go-sql-driver/mysql"

)

type Tag struct {
	nomProducto   string    `json:"nomProducto"`
	precioProducto int64 `json:"precioProducto"`
}
var templ *template.Template
func delete(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("compra.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		productoNombre := r.FormValue("nomProducto")
		db, err := sql.Open("mysql", "root:admin@/carrito")
		if err != nil {
			panic(err.Error())
		}
		delForm, err := db.Prepare("DELETE FROM producto WHERE nomProducto=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(productoNombre)
		log.Println("DELETE")
		defer db.Close()
	}
}

func mostrar(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		db, err := sql.Open("mysql", "root:admin@/carrito")
		if err != nil {
			panic(err.Error())
		}
		results, err := db.Query("SELECT nomProducto, precioProducto FROM producto")
		if err != nil {
			panic(err.Error())
		}
		res := []Tag{}
		emp := Tag{}
		for results.Next() {
			var precioProducto int64
			var nomProducto string
			err = results.Scan(&nomProducto, &precioProducto)
			if err != nil {
				panic(err.Error())
			}
			emp.precioProducto = precioProducto
			emp.nomProducto = nomProducto
			res = append(res,emp)
		}
		fmt.Fprintf(w,"<!DOCTYPE html>")
		fmt.Fprintf(w,"<html lang=\"en\">")
		fmt.Fprintf(w,"<head>")
		fmt.Fprintf(w," <meta charset=\"UTF-8\">")
		fmt.Fprint(w,"<link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\" integrity=\"sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T\" crossorigin=\"anonymous\">")
		fmt.Fprintf(w,"</head>")
		fmt.Fprintf(w,"<body class=\"container\">")
		fmt.Fprintf(w,"<table class=\"table\">")
		fmt.Fprintf(w,"<tr>")
		fmt.Fprintf(w,"<th scope=\"col\">Producto</th>")
		fmt.Fprintf(w,"<th scope=\"col\">Precio</th>")
		fmt.Fprintf(w,"</tr>")
		for i:=0;i<len(res);i++{
			fmt.Fprintf(w,"<tr scope=\"row\">")
			fmt.Fprintf(w,"<td>%s</td>",res[i].nomProducto)
			fmt.Fprintf(w,"<td>%d</td>",res[i].precioProducto)
			fmt.Fprintf(w,"</tr>")
		}
		fmt.Fprintf(w,"</table>")
		fmt.Fprintf(w,"<form action=\"/delete\" method=\"post\">")
		fmt.Fprintf(w,"<div class=\"form-group\">")
		fmt.Fprintf(w,"<label for=\"nomProducto\">Nombre Producto</label>")
		fmt.Fprintf(w,"<input class=\"form-control\" type=\"text\" name=\"nomProducto\" id=\"nomProducto\" required>")
		fmt.Fprintf(w,"</div>")
		fmt.Fprintf(w,"<input type=\"submit\" value=\"Borrar\">")
		fmt.Fprintf(w,"</form>")
		fmt.Fprintf(w,"</body>")
		fmt.Fprintf(w,"</html>")
	}
}

func insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		db, err := sql.Open("mysql", "root:admin@/carrito")
		if err != nil {
			panic(err.Error())
		}
		productoNombre := r.FormValue("nomProducto")
		productoPrecio := r.FormValue("preProducto")



		ins, err := db.Prepare("INSERT INTO producto (nomProducto, precioProducto) VALUES (?,?)")
		if err != nil {
			panic(err.Error())
		}
		ins.Exec(productoNombre, productoPrecio)
		log.Println("INSERT: Codigo: " + productoNombre + " | Nombre: " + productoPrecio)

	}


}

func main() {
	http.HandleFunc("/", insert)
	http.HandleFunc("/mostrar", mostrar)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}
