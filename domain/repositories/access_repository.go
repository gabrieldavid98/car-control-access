package repositories

// AccessRepository interface
type AccessRepository interface {
	RegisterAccess(plate string) error
	RegisterExit(plate string) error
}
