package set

import (
	"sync"
	"sync/atomic"
)

// Set 集合
type Set struct {
	m   sync.Map
	num int32
}

// 创建
func New() *Set {
	return &Set{
		m:   sync.Map{},
		num: 0,
	}
}

// 添加
func (s *Set) Add(item interface{}) {
	s.m.Store(item, true)
	atomic.AddInt32(&s.num, 1)
}

// 删除
func (s *Set) Remove(item interface{}) {
	s.m.Delete(item)
	atomic.AddInt32(&s.num, -1)
}

// 判断是否存在
func (s *Set) Has(item interface{}) (ok bool) {
	_, ok = s.m.Load(item)
	return
}

// 获取集合大小
func (s *Set) Len() int {
	return int(s.num)
}

// 清除
func (s *Set) Clear() {
	s.m.Range(func(k, v interface{}) bool {
		s.m.Delete(k)
		atomic.AddInt32(&s.num, -1)
		return true
	})
}

// 判断是否为空
func (s *Set) IsEmpty() bool {
	return s.num == 0
}

// 转切片输出
func (s *Set) List() (list []interface{}) {
	s.m.Range(func(k, v interface{}) bool {
		list = append(list, k)
		return true
	})

	return list
}

// 遍历
func (s *Set) Range(f func(key interface{}) bool) {
	s.m.Range(func(k, v interface{}) bool {
		return f(k)
	})
}
