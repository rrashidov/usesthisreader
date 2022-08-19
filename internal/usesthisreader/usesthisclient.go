package usesthisreader

type UsesThisClient interface {
	GetLatest() (string, error)
	Update(string) error
}
