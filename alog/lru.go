package alog

// LRUCache 是一个 LRU 缓存结构
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

// DLinkedNode 是双向链表节点结构
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// NewLRUCache 创建一个新的 LRU 缓存实例
func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get 从缓存中获取值
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

// Put 将键值对放入缓存
func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		node.value = value
		lru.moveToHead(node)
	} else {
		node = &DLinkedNode{key: key, value: value}
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
			lru.size--
		}
	}
}

// moveToHead 将节点移动到双向链表头部
func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// addToHead 在双向链表头部添加节点
func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// removeNode 移除双向链表中的节点
func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// removeTail 移除双向链表尾部节点
func (lru *LRUCache) removeTail() *DLinkedNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}
