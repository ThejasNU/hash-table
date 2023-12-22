package hashtable

type item struct {
	key   string
	value string
}

type table struct {
	size       int
	itemsCount int
	items      []*item
}
