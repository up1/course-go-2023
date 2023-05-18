package beers

import (
	"fmt"
	"time"
)

type StorageMemory struct {
	beers []Beer
}

func (s *StorageMemory) FindBeers() ([]Beer, error) {
	return s.beers, nil
}

func (s *StorageMemory) SaveBeer(beers ...Beer) error {
	for _, beer := range beers {
		for _, b := range s.beers {
			if beer.Abv == b.Abv &&
				beer.Brewery == b.Brewery &&
				beer.Name == b.Name {
				return fmt.Errorf("Beer already exists")
			}
		}

		beer.ID = len(s.beers) + 1
		beer.Created = time.Now()

		s.beers = append(s.beers, beer)
	}
	return nil
}
