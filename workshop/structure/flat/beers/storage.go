package beers

type Storage interface {
	SaveBeer(...Beer) error
	FindBeers() ([]Beer, error)
}

func NewStorage() (Storage, error) {
	// return &StorageMemory{}, nil
	return new(StorageMemory), nil
}
