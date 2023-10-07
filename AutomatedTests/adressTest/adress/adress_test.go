package adress

import "testing"

func TestAdressType(t *testing.T) {
	adressForSuccess := "Avenida dos Asstronautas"
	adressForError := "Rodovia dos Tamoios"

	expectedAdressType := "Avenida"

	recievedAdressType := AdressType(adressForSuccess)
	recievedAdressTypeError := AdressType(adressForError)

	if recievedAdressType != expectedAdressType{
		t.Errorf("Tipo recebido e diferente do esperado, esperava %s e recebeu %s", expectedAdressType, recievedAdressType)
	}

	// Error case
	if recievedAdressType != recievedAdressTypeError{
		t.Errorf("Tipo recebido e diferente do esperado, esperava %s e recebeu %s", expectedAdressType, recievedAdressTypeError)
	}
}