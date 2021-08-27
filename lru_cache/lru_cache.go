package lru_cache

type LRUCache struct {
}

func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

func (l *LRUCache) Get(key int) int {
	return 0
}

func (l *LRUCache) Put(key int, value int) {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
