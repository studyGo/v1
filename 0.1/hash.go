package main

import "hash/crc32"
import "fmt"
import "sort"
import "strconv"

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

func (m *Map) Set(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}

	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	index := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })
	if index == len(m.keys) {
		index = 0
	}
	return m.hashMap[m.keys[index]]
}

func main() {
	// 一致性hash
	// uint 32  =  0 ---- (1 >> 32) - 1
	// Set 定义key 的位置 但是不是均匀分布的 所以key 越多 越趋向于均匀
	m := New(10, nil)
	m.Set("demo", "select", "test")
	fmt.Println(m.Get("ll"))
	fmt.Println(m.Get("get select key"))
	fmt.Println(m.Get("1--1-1-1"))
}
