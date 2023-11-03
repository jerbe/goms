package container

import "sync"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/24 01:00
  @describe :
*/

// NewSet 返回一个新的集合
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

// Set 集合
type Set[T comparable] struct {
	m       map[T]struct{}
	rwMutex sync.RWMutex
}

// Add 添加元素
func (s *Set[T]) Add(item T) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.m[item] = struct{}{}
}

// Contains 判断是否包含元素
func (s *Set[T]) Contains(item T) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Remove 移除指定元素
func (s *Set[T]) Remove(item T) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	delete(s.m, item)
}

// Size 获取Set的大小
func (s *Set[T]) Size() int {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return len(s.m)
}

// Elements 获取所有元素
func (s *Set[T]) Elements() []T {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	elems := make([]T, 0)
	for item, _ := range s.m {
		elems = append(elems, item)
	}
	return elems
}
