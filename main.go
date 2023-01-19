package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Iniciando a Aplicação")
	fmt.Println("-----------------------------------------")
	fmt.Println(" ")
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(" ")
	fmt.Println("-----------------------------------------")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
}
