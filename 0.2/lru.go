package main

import "container/list"
import "fmt"

type Key interface{}

type entry struct {
    key Key
    value interface{}
}

type Cache struct {
    max int
    l *list.List
    cache map[interface{}]*list.Element
}

func New(num int) *Cache {
    return &Cache{
        max : num,
        l : list.New(),
        cache : make(map[interface{}]*list.Element),
    }
}

func (c *Cache) Add (key Key, value interface{}) {
    if e, ok := c.cache[key]; ok {
        c.l.MoveToFront(e)
        e.Value.(*entry).value = value
        return
    }

    ele := c.l.PushFront(&entry{key, value})
    c.cache[key] = ele

    // 删除旧的
    if c.max != 0 && c.l.Len() > c.max {
        c.RemoveOld()
    }
}

func (c *Cache) Get(key Key) (value interface{}) {
    if c.cache == nil {
        return
    }

    if ele, ok := c.cache[key]; ok {
        c.l.MoveToFront(ele)
        return ele.Value.(*entry).value
    }
    return

}

func (c *Cache) First() (key interface{}){

    if c.cache == nil {
        return c.cache
    }
    return c.l.Front().Value.(*entry).key
}

func (c *Cache) Remove(key Key) {
    if c.cache == nil {
        return
    }

    if ele, ok := c.cache[key]; ok {
        c.removeElement(ele)
    }
}

// old list last one
func (c *Cache) RemoveOld() {
    if c.cache == nil {
        return
    }

    e := c.l.Back()
    c.removeElement(e)
}

func (c *Cache) removeElement(e *list.Element) {
    c.l.Remove(e)
    kv := e.Value.(*entry)
    delete(c.cache, kv.key)
}

func main() {

    ele := New(100)
    ele.Add("demo", 100)
    ele.Add("demos", "tt")
    ele.Get("demo")
    fmt.Println(ele.First())
    fmt.Println(ele.Get("demos"))
    fmt.Println(ele.First())
}
