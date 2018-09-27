package main

import (
	"fmt"

	"github.com/viniciusfeitosa/Treinamento-GO/basico3/append_struct/models"
)

func main() {
	familia := models.Familia{
		Pai:   models.Pai{Nome: "João", Idade: 50},
		Filho: models.Filho{Nome: "Carlos", Idade: 20, Email: "carlos@teste.com"},
	}

	fmt.Println(familia.DadosDoPai())
	fmt.Println(familia.DadosDoFilho())
}
