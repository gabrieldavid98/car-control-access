package errors

import "errors"

var (
	// ErrPlateLength error
	ErrPlateLength = errors.New("La placa no debe tener menos de 6 caracteres")

	// ErrNotAlphaNumericPlate error
	ErrNotAlphaNumericPlate = errors.New("La placa solo debe contener letras y núemeros")
)
