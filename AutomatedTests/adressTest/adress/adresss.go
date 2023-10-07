package adress

import "strings"

// AdressType check if adress is valid
func AdressType(adress string) string {
	validTypes := []string{"rua", "avenida", "Estrada", "serra", "rodovia"}
	formatAdress := strings.ToLower(adress)

	// Get first idx of adress ("Street xyz ...")
	firstWord := strings.Split(formatAdress, " ")[0]

	validateAdress := false
	for _, tipo := range validTypes {
		if tipo == firstWord {
			validateAdress = true
		}
	}

	if validateAdress {
		return strings.Title(firstWord)
	}

	return "Tipo de endereço invalido."
}