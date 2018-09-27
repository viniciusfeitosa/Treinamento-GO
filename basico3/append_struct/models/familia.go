package models

import "fmt"

// Familia é uma estrutura
type Familia struct {
	Pai   Pai
	Filho Filho
}

// Pai é o pai
type Pai struct {
	Nome  string
	Idade int
}

// Filho é o filho do Pai
type Filho struct {
	Nome  string
	Idade int
	Email string
}

// DadosDoPai mostra datos do pai
func (f Familia) DadosDoPai() string {
	return fmt.Sprintf("Nome: %s, Idade: %d", f.Pai.Nome, f.Pai.Idade)
}

// DadosDoFilho mostra datos do filho
func (f Familia) DadosDoFilho() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Email: %s", f.Filho.Nome, f.Filho.Idade, f.Filho.Email)
}
