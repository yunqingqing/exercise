package day50

import "../utils/doublelinkedlist"

type LRUCache struct {
	data map[string]*doublelinkedlist.Node
	root *doublelinkedlist.List
	size int
}

func NewLRUCache(size int) *LRUCache {
	cache := new(LRUCache)
	cache.size = size
	cache.root = doublelinkedlist.NewList()
	cache.data = make(map[string]*doublelinkedlist.Node, size)
	return cache
}

func (cache *LRUCache) Put(key, value string) {
	node, ok := cache.data[key]

	// 当前key已经在缓存映射中，更新旧值并把对应节点移动到链表尾
	if ok {
		node.SetValue(value)
		cache.root.MoveToBack(node)
		return
	}

	// 链表大小到达最大值，淘汰链表头节点
	if cache.size == len(cache.data) {
		tmp := cache.root.Front()
		cache.root.Remove(tmp)
		delete(cache.data, tmp.GetKey())
	}

	node = doublelinkedlist.NewNode(key, value)
	cache.root.AppendToBack(node)
	cache.data[key] = node
}

func (cache *LRUCache) Get(key string) string {
	node, ok := cache.data[key]

	// 缓存命中，把对应节点移到链表尾部
	if ok {
		cache.root.MoveToBack(node)
		return node.GetValue()
	}
	return ""
}
