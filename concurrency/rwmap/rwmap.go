package rwmap

import "sync"

type RWMap struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁保护下面的map 字段
	m            map[any]any
}

// NewRWMap 新建一个 RWMap
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[any]any, n),
	}

}

func (m *RWMap) Get(k any) (any, bool) { // 从map中读取一个值
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k] // 在锁的保护下从map中读取
	return v, existed
}

func (m *RWMap) Set(k any, v any) { // 设置一个键值对
	m.Lock()
	defer m.Unlock()
	m.m[k] = v

}

func (m *RWMap) Delete(k any) { //删除一个键
	m.Lock() // 锁保护
	defer m.Unlock()
	delete(m.m, k)

}

func (m *RWMap) Len() any { // map的长度
	m.RLock() // 锁保护
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v any) bool) { // 遍历map
	m.RLock() //遍历期间一直持有读锁
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
