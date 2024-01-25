package contest

import (
	"math/rand"
	"strconv"
	"testing"
)

const totalStores = 10000

func BenchmarkSimpleMapSearcher_GetStoreByName(b *testing.B) {

	stores := generateStores(totalStores)
	var found []Store
	var err error
	s := NewSimpleMapSearcher(stores)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, err = s.GetStoreByName("store 1")
	}

	_ = err

	_ = found
}

func BenchmarkSimpleMapSearcher_GetStoreByGeo(b *testing.B) {

	stores := generateStores(totalStores)
	var found []Store
	var err error
	s := NewSimpleMapSearcher(stores)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, err = s.GetStoreByGeo(Geo{Latitude: 0.5, Longitude: 0.5})
	}

	_ = err

	_ = found
}

func BenchmarkSimpleMapSearcher_GetStoreByCategories(b *testing.B) {

	stores := generateStores(totalStores)
	var found []Store
	var err error

	categories := randCategories()

	s := NewSimpleMapSearcher(stores)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, err = s.GetStoreByCategories(categories)
	}

	_ = err

	_ = found
}

func generateStores(i int) []Store {

	stores := make([]Store, i)

	for j := 0; j < i; j++ {
		stores[j] = randSore()
	}

	return stores
}

func randSore() Store {
	return Store{
		Name:       randName(),
		Geo:        randGeo(),
		Categories: randCategories(),
	}
}

func randName() string {
	return "store " + strconv.Itoa(rand.Intn(1000))
}

func randCountCategories() int {
	return rand.Intn(10) + 1
}

func randCategoryName() string {
	return "category " + strconv.Itoa(rand.Intn(1000))
}

func randCategories() []string {
	count := randCountCategories()
	categories := make([]string, count)
	for i := 0; i < count; i++ {
		categories[i] = randCategoryName()
	}
	return categories
}

func randGeo() Geo {
	return Geo{
		Latitude:  rand.Float32() + 1,
		Longitude: rand.Float32() + 3,
	}
}
