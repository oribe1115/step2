package lib

import "fmt"

// HashTable .
type HashTable struct {
	Map map[string]string
}

func InitHashTable() *HashTable {
	table := &HashTable{}
	table.Map = make(map[string]string, 0)

	return table
}

func (t *HashTable) Exist(key string) bool {
	_, ok := t.Map[key]

	return ok
}

func (t *HashTable) Search(key string) string {
	value, _ := t.Map[key]

	return value
}

func (t *HashTable) Add(key, value string) {
	t.Map[key] = value
}

func (t *HashTable) Delete(key string) {
	delete(t.Map, key)
}

func (t *HashTable) PrintAll() {
	for key := range t.Map {
		fmt.Printf("%s ", t.Map[key])
	}
	fmt.Printf("\n")
}
