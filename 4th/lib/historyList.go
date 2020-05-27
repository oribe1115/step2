package lib

import "fmt"

type HistoryList struct {
	Latest      *HistoryListNode
	Oldest      *HistoryListNode
	LengthLimit int
}

type HistoryListNode struct {
	Content string
	Later   *HistoryListNode
	Older   *HistoryListNode
}

func createHistoryListNode(content string) *HistoryListNode {
	node := &HistoryListNode{
		Content: content,
		Later:   nil,
		Older:   nil,
	}

	return node
}

func InitHistoryList(limit int) *HistoryList {
	list := &HistoryList{
		Latest:      nil,
		Oldest:      nil,
		LengthLimit: limit,
	}

	return list
}

func (l *HistoryList) Length() int {
	length := 0
	node := l.Latest
	for node != nil {
		length++
		node = node.Older
	}

	return length
}

func (l *HistoryList) Print() {
	node := l.Latest
	for node != nil {
		fmt.Printf("%s ", node.Content)
		node = node.Older
	}
}

func (l *HistoryList) Search(content string) *HistoryListNode {
	node := l.Latest
	for node != nil {
		if node.Content == content {
			break
		}
		node = node.Older
	}

	return node
}

func (l *HistoryList) deleteOldest() (deletedContent string) {
	target := l.Oldest
	nextOldest := target.Later
	nextOldest.Older = nil
	l.Oldest = nextOldest

	return target.Content
}

func (l *HistoryList) addLatest(content string) {
	previousLatest := l.Latest
	nextLatest := createHistoryListNode(content)
	nextLatest.Older = previousLatest
	previousLatest.Later = nextLatest
	l.Latest = nextLatest
}
