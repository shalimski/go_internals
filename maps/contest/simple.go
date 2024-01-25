package contest

type SimpleMapSearcher struct {
	stores             []Store
	storesByName       map[string][]Store
	storesByGeo        map[Geo][]Store
	storesByCategories map[string][]int
	storeOccurrences   []int
}

func NewSimpleMapSearcher(stores []Store) *SimpleMapSearcher {
	s := &SimpleMapSearcher{
		stores:             stores,
		storesByName:       make(map[string][]Store, len(stores)),
		storesByGeo:        make(map[Geo][]Store, len(stores)),
		storesByCategories: make(map[string][]int, len(stores)),
		storeOccurrences:   make([]int, len(stores)),
	}

	for i := 0; i < len(stores); i++ {
		s.storesByName[stores[i].Name] = append(s.storesByName[stores[i].Name], stores[i])
		s.storesByGeo[stores[i].Geo] = append(s.storesByGeo[stores[i].Geo], stores[i])
		for j := 0; j < len(stores[i].Categories); j++ {
			s.storesByCategories[stores[i].Categories[j]] = append(s.storesByCategories[stores[i].Categories[j]], i)
		}
	}

	return s
}

func (s *SimpleMapSearcher) GetStoreByName(name string) ([]Store, error) {
	if stores, ok := s.storesByName[name]; ok {
		return stores, nil
	}
	return nil, ErrNotFound
}

func (s *SimpleMapSearcher) GetStoreByGeo(geo Geo) ([]Store, error) {
	if stores, ok := s.storesByGeo[geo]; ok {
		return stores, nil
	}
	return nil, ErrNotFound
}

func (s *SimpleMapSearcher) GetStoreByCategories(categories []string) ([]Store, error) {
	clear(s.storeOccurrences)
	var stores []Store

	for _, category := range categories {
		if storesByCategory, ok := s.storesByCategories[category]; ok {
			for _, storeIdx := range storesByCategory {
				s.storeOccurrences[storeIdx]++
				if s.storeOccurrences[storeIdx] == len(categories) {
					stores = append(stores, s.stores[storeIdx])
				}
			}
		}
	}

	if len(stores) > 0 {
		return stores, nil
	}
	return nil, ErrNotFound
}
