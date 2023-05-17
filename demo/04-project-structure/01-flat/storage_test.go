package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaveBeer(t *testing.T) {
	storage := new(StorageMemory)
	sampleBeer := Beer{
		Name:      "Test name",
		Brewery:   "Test brewery",
		Abv:       8,
		ShortDesc: "Test short description",
	}

	err := storage.SaveBeer(sampleBeer)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(storage.beers))

	beer := storage.beers[0]
	assert.NotNil(t, sampleBeer.ID)
	assert.Equal(t, sampleBeer.Name, beer.Name)
	assert.Equal(t, sampleBeer.Brewery, beer.Brewery)
	assert.Equal(t, sampleBeer.Abv, beer.Abv)
	assert.Equal(t, sampleBeer.ShortDesc, beer.ShortDesc)
	assert.NotNil(t, sampleBeer.Created)
	assert.True(t, sampleBeer.Created.Before(time.Now()))
}

func TestSaveBeerShouldErrorBeerAlreadyExists(t *testing.T) {
	storage := new(StorageMemory)
	sampleBeer := Beer{
		Name:      "Test name",
		Brewery:   "Test brewery",
		Abv:       8,
		ShortDesc: "Test short description",
	}

	err := storage.SaveBeer(sampleBeer)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(storage.beers))

	errDupe := storage.SaveBeer(sampleBeer)
	assert.NotNil(t, errDupe)
	assert.Equal(t, "Beer already exists", errDupe.Error())
}

func TestFindBeers(t *testing.T) {
	storage, _ := NewStorage()
	sampleBeer := Beer{
		Name:      "Test name",
		Brewery:   "Test brewery",
		Abv:       8,
		ShortDesc: "Test short description",
	}
	storage.SaveBeer(sampleBeer)

	beers, err := storage.FindBeers()

	assert.Nil(t, err)
	assert.NotNil(t, beers)
	assert.Equal(t, 1, len(beers))
}
