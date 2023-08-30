package pokeapi

type LocationAreaRes struct {
	count    int
	next     *string
	previous *string
	results  []struct {
		name string
		url  string
	}
}
