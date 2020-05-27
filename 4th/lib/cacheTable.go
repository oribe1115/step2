package lib

import "fmt"

// CacheTable キー対応する情報と、keyの更新順序を保持する
type CacheTable struct {
	Table   *HashTable
	History *HistoryList
}

func InitCacheTable(histryLimit int) *CacheTable {
	return &CacheTable{
		Table:   InitHashTable(),
		History: InitHistoryList(histryLimit),
	}
}

func (c *CacheTable) Exist(key string) bool {
	return c.Table.Exist(key)
}

func (c *CacheTable) Search(key string) string {
	return c.Table.Search(key)
}

func (c *CacheTable) Latest() (key, value string) {
	key = c.History.Latest.Content
	value = c.Table.Map[key]

	return key, value
}

func (c *CacheTable) Cache(key, value string) {
	c.Table.Add(key, value)
	isDelete, deletedKey := c.History.Cache(key)
	if isDelete {
		c.Table.Delete(deletedKey)
	}
}

func (c *CacheTable) Print() {
	fmt.Printf("Table: ")
	c.Table.PrintAll()
	fmt.Printf("History: ")
	c.History.Print()
}
