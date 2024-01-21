package contest

import "errors"

var ErrNotFound = errors.New("not found")

type Store struct {
	Name       string
	Geo        Geo
	Categories []string
}

type Geo struct {
	Latitude  float32
	Longitude float32
}

type Searcher interface {
	GetStoreByName(string) ([]Store, error)
	GetStoreByGeo(Geo) ([]Store, error)
	GetStoreByCategories([]string) ([]Store, error)
}
