package main


import (
"fmt"
"html/template"
"math"
"net/http"
"strconv"
)

func Get_Post(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		a, err := strconv.ParseFloat(r.FormValue("num1"),64)
		if err != nil {


		}
		b, err := strconv.ParseFloat(r.FormValue("num2"),64)
		if err != nil {

		}
		c, err := strconv.ParseFloat(r.FormValue("num3"),64)
		if err != nil {

		}

		r1,r2:=formula(a,b,c);

		fmt.Fprintf(w,"<h1>X1= %f</h1>",r1)
		fmt.Fprintf(w,"<h1>X2= %f</h1>",r2)


	}
}

func formula(a,b,c float64)(float64,float64)  {
	var raiz1 float64
	var raiz2 float64
	var r float64
	r = math.Pow(b,2) - (4*a*c)
	r=math.Sqrt(r)
	raiz1=(-b + r)/(2*a)
	raiz2=(-b - r)/(2*a)
	return raiz1,raiz2;
}


func main(){

	http.HandleFunc("/", Get_Post)
	http.ListenAndServe(":8080", nil)

}
