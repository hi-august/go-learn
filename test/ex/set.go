//TODO: 实现一个简单的Set
//
// 需要一个map保存数据, 同时需要一个读写锁(保证线程安全)
// 可以保存所有的类型作为Key
package main

import (
    "fmt"
    "sync"
)

type Set struct {
    m map[interface{}]bool
    sync.RWMutex // 添加锁
}

// New: 返回一个Set实例
func New() *Set {
    return &Set{
        m: map[interface{}]bool{},
    }
}

// Add: 增加一个元素
func (s *Set)Add(item interface{})  {
    s.Lock() // 写锁定
    defer s.Unlock() //写锁释放
    s.m[item] = true
}

// Remove: 移除一个元素
func (s *Set)Remove(item interface{})  {
    s.Lock()
    defer s.Unlock()
    delete(s.m, item)
}

// Has: 是否存在指定的元素
func (s *Set)Has(item interface{}) bool {
    // 允许读
    s.RLock() //读锁定
    defer s.RUnlock() //读锁释放
    _, ok := s.m[item]
    return ok
}

// List: 获取Map转化成的list
func (s *Set)List() []interface{} {
    s.RLock()
    defer s.RUnlock()
    var l []interface{}
    for value := range s.m {
        l = append(l, value)
    }
    return l
}

// Len: 返回元素个数
func (s *Set)Len() int {
    return len(s.List())
}

// Clear: 清除Set
func (s *Set) Clear() {
    s.Lock()
    defer s.Unlock()
    s.m = map[interface{}]bool{} //重置为空
}

// Empty: Set是否是空
func (s *Set) IsEmpty() bool {
    if s.Len() == 0 {
        return true
    }
    return false
}

func main() {
    a := New()
    a.Add("hi")
    a.Add(123)
    a.Add("august")
    fmt.Println(a, a.Len(), a.IsEmpty())
    fmt.Println(a.List()) //转化为list
    a.Remove(123)
    // a.Clear()
    fmt.Println(a, a.Has("hi"), a.Len())
}

func CheckErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}
