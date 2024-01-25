package contest

import (
	"reflect"
	"testing"
)

func TestSimpleMapSearcher_GetStoreByName(t *testing.T) {

	tests := []struct {
		name      string
		stores    []Store
		storeName string
		want      []Store
		wantErr   error
	}{
		{
			name: "test 1",
			stores: []Store{
				{
					Name: "store 1",
					Geo: Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
				{
					Name: "store 2",
					Geo: Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404"},
				},
			},
			storeName: "store 1",
			want: []Store{
				{
					Name: "store 1",
					Geo: Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSimpleMapSearcher(tt.stores)
			got, err := s.GetStoreByName(tt.storeName)

			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestSimpleMapSearcher_GetStoreByGeo(t *testing.T) {

	tests := []struct {
		name    string
		stores  []Store
		geo     Geo
		want    []Store
		wantErr error
	}{
		{
			name: "test 1",
			stores: []Store{
				{
					Name: "store 1",
					Geo: Geo{
						Latitude:  0.6,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
				{
					Name: "store 2",
					Geo: Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404"},
				},
			},
			geo: Geo{
				Latitude:  0.5,
				Longitude: 0.5,
			},
			want: []Store{
				{
					Name: "store 2",
					Geo: Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSimpleMapSearcher(tt.stores)
			got, err := s.GetStoreByGeo(tt.geo)

			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestSimpleMapSearcher_GetStoreByCategories(t *testing.T) {

	tests := []struct {
		name       string
		stores     []Store
		categories []string
		want       []Store
		wantErr    error
	}{
		{
			name: "test 1",
			stores: []Store{
				{
					Name: "store 1",
					Geo: Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
				{
					Name: "store 2",
					Geo: Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 404"},
				},
			},
			categories: []string{"category 1", "category 2"},
			want: []Store{
				{
					Name: "store 1",
					Geo: Geo{
						Latitude:  0.5,
						Longitude: 0.5,
					},
					Categories: []string{"category 1", "category 2"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSimpleMapSearcher(tt.stores)
			got, err := s.GetStoreByCategories(tt.categories)

			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleMapSearcher.GetStoreByCategories() = %v, want %v", got, tt.want)
			}

		})
	}
}
