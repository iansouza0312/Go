package adress

import "testing"

type testingCase struct {
	recievedAdress string
	expectedReturn string
}

func TestAdressType(t *testing.T) {

	testingCases := []testingCase{
		{  "Rua dos passaros", "Rua" },
		{  "Avenida dos Astronautas", "Avenida" },
		{  "Rodovia dos Tamoios", "Rodovia" },
		{  "Rodovia Presidente Dutra", "Rodovia" },
		{  "Serra do mar", "Serra" },
		{  "Rua 12 de março", "Rua" },
		{  "Curva do S ", "Tipo de endereço invalido." },
	}

	for _, cenario := range testingCases {
		adressType := AdressType(cenario.recievedAdress)
		if adressType != cenario.expectedReturn {
			t.Errorf("O tipo recebido %s e diferente do esperado %s", adressType, cenario.expectedReturn)
		}
	}
}