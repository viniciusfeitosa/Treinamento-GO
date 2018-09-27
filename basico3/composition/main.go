package main

import (
	"fmt"

	"github.com/viniciusfeitosa/Treinamento-GO/basico3/composition/models"
)

func mostraDados(membro models.Familia) {
	fmt.Println(membro.Dados())
}

func main() {
	pai := new(models.Pai)

	pai.Nome = "João"
	pai.Idade = 50

	filho := new(models.Filho)
	filho.Nome = "Carlos"
	filho.Idade = 20
	filho.Email = "carlos@teste.com"

	mostraDados(pai)
	mostraDados(filho)
}
