package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenido UNIR!")
		fmt.Printf("\n Bienvenido UNIR! en \"/\" \n")
	})

	http.HandleFunc("/suma", func(w http.ResponseWriter, r *http.Request) {

		x := rand.Intn(100)
		y := rand.Intn(100)
		suma := x + y

		fmt.Fprintf(w, " sumar %d + %d es igual a %d", x, y, suma)
		fmt.Printf(" sumar %d + %d es igual a %d, texto plano \n", x, y, suma)
	})

	http.HandleFunc("/sumaj", func(w http.ResponseWriter, r *http.Request) {

		x := rand.Intn(100)
		y := rand.Intn(100)
		suma := x + y

		fmt.Printf(" sumar %d + %d es igual a %d, JSON \n", x, y, suma)

		type Datos struct {
			Numerox   int `json:"numerox"`
			Numeroy   int `json:"numeroy"`
			Resultado int `json:"resultado"`
		}

		type datosFinal struct {
			Lasuma []Datos `json:"lasuma"`
		}

		losDatos := &datosFinal{
			Lasuma: []Datos{{Numerox: x, Numeroy: y, Resultado: suma}},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(losDatos)

	})

	http.ListenAndServe(":8383", nil)
}
