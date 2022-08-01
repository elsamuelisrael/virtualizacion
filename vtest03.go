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

// dckr_pat_dCBBuAJcIC2OsWccZ8RmKmMStlA
// http://165.232.135.211:8383

// sudo docker build --tag testv03 .
// sudo docker run --publish 8383:8383 testv03
// sudo docker run -d -p 8383:8383 testv03
// sudo docker login -u elsamuelisrael
// sudo docker tag testv03 elsamuelisrael/testv03
// sudo docker push elsamuelisrael/testv03:latest

/*

# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /testv03

EXPOSE 8383

CMD [ "/testv03" ]

*/
