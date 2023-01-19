package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando a Aplicação")
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
