package models_test

import (
	"testing"

	"github.com/viniciusfeitosa/Treinamento-GO/basico3/append_struct/models"
)

func TestDadosDoPai(t *testing.T) {
	expected := "Nome: teste, Idade: 10"

	f := models.Familia{
		Pai: models.Pai{Nome: "teste", Idade: 10},
	}

	if f.DadosDoPai() != expected {
		t.Errorf("Expected: %s, Result: %s", expected, f.DadosDoPai())
	}
}
