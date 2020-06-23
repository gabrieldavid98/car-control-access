package services

// AccessService interface provides required methods
// for a payment service
type AccessService interface {
	RegisterAccess(plate string) error
	RegisterExit(plate string) error
}
