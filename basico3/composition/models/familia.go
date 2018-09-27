package models

import "fmt"

// Familia é uma interface
type Familia interface {
	Dados() string
}

// Pai é o pai
type Pai struct {
	Nome  string
	Idade int
}

// Filho é o filho do Pai
type Filho struct {
	Pai
	Email string
}

// Dados mostra datos do pai
func (p Pai) Dados() string {
	return fmt.Sprintf("Nome: %s, Idade: %d", p.Nome, p.Idade)
}

// Dados mostra datos do filho
func (f Filho) Dados() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Email: %s", f.Nome, f.Idade, f.Email)
}
