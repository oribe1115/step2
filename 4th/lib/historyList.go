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

	fmt.Printf("\n")
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
	if target == nil {
		return ""
	} else if target == l.Latest {
		l.Oldest = nil
		l.Latest = nil
	} else {
		nextOldest := target.Later
		nextOldest.Older = nil
		l.Oldest = nextOldest
	}

	return target.Content
}

func (l *HistoryList) addLatest(latest *HistoryListNode) {
	if l.Latest == nil {
		l.Latest = latest
		l.Oldest = latest
	} else {
		previous := l.Latest
		latest.Later = nil
		latest.Older = previous
		previous.Later = latest
		l.Latest = latest
	}
}

func (l *HistoryList) remove(node *HistoryListNode) *HistoryListNode {
	if node == l.Latest && node == l.Oldest {
		l.Latest = nil
		l.Oldest = nil
		return node
	}

	if node == l.Latest {
		l.Latest = node.Older
		l.Latest.Later = nil
	} else if node == l.Oldest {
		l.Oldest = node.Later
		l.Oldest.Older = nil
	} else {
		node.Later.Older = node.Older
		node.Older.Later = node.Later

	}

	node.Later = nil
	node.Older = nil

	return node
}

func (l *HistoryList) Cache(content string) {
	node := l.Search(content)
	if node == nil {
		l.addLatest(createHistoryListNode(content))
		if l.Length() > l.LengthLimit {
			l.deleteOldest()
		}
	} else {
		node = l.remove(node)
		l.addLatest(node)
	}
}
