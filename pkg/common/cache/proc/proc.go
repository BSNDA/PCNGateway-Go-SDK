package proc

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	NoExpiration      time.Duration = -1 //No expiration mark
	defaultExpiration time.Duration = 0  // default non-expired （expiration time in ns）
)

type Item struct {
	Object     interface{}   //data item
	Expiration int64         //expiration time of data item (0 never expires)
	IsSlide    bool          // slide out of date or not
	Dur        time.Duration //duration
}

type Cache struct {
	defaultExpiration time.Duration //use if the data item is not specified to expire
	items             map[string]Item
	mu                sync.RWMutex  //read-write lock
	gcInterval        time.Duration //gc cycle
	stopGc            chan bool     // stop gc channel ID
}

func (item Item) IsExpired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration //expire if it exceeds the time
}

//repeat gc
func (c *Cache) gcLoop() {
	ticker := time.NewTicker(c.gcInterval) //initialize a timer
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-c.stopGc:
			ticker.Stop()
			return
		}
	}
}

//delete expired cache
func (c *Cache) DeleteExpired() {
	now := time.Now().UnixNano()
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration {
			c.delete(k)
		}
	}
}

//delete
func (c *Cache) delete(k string) {
	delete(c.items, k)
}

//operate deletion
func (c *Cache) Delete(k string) {
	c.mu.Lock()
	c.delete(k)
	defer c.mu.Unlock()
}

//set the data item of cache, overwrite it if it is set, isSlide expired or not when d expires
func (c *Cache) Set(k string, v interface{}, d time.Duration, isSlide bool) {
	var e int64
	if d == defaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[k] = Item{
		Object:     v,
		Expiration: e,
		IsSlide:    isSlide,
		Dur:        d,
	}
}

//set data item, lock-free operation
func (c *Cache) set(k string, v interface{}, d time.Duration) {
	var e int64
	if d == defaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.items[k] = Item{
		Object:     v,
		Expiration: e,
	}
}

//Get the data item, to see if expired
func (c *Cache) get(k string) (interface{}, bool) {
	item, found := c.items[k]
	if !found || item.IsExpired() {
		return nil, false
	}

	if item.IsSlide {
		c.Set(k, item.Object, item.Dur, true)
	}

	return item.Object, true
}

//add new operations, if data item exists, then error
func (c *Cache) Add(k string, v interface{}, d time.Duration) error {
	c.mu.Lock()
	_, found := c.get(k)
	if found {
		c.mu.Unlock()
		return fmt.Errorf("Item %s already exists", k)
	}
	c.set(k, v, d)
	c.mu.Unlock()
	return nil
}

//Get cache
func (c *Cache) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[k]
	if !found || item.IsExpired() {
		c.mu.RUnlock()
		return nil, false
	}
	c.mu.RUnlock()

	if item.IsSlide {
		c.Set(k, item.Object, item.Dur, true)
	}

	return item.Object, true
}

//replace
func (c *Cache) Replace(k string, v interface{}, d time.Duration) error {
	c.mu.Lock()
	_, found := c.get(k)
	if !found {
		c.mu.Unlock()
		return fmt.Errorf("Item %s does't exists", k)
	}
	c.set(k, v, d)
	c.mu.Unlock()
	return nil
}

// input cache data to io.Writer
func (c *Cache) Save(w io.Writer) (err error) {
	enc := gob.NewEncoder(w)
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Error Registering item types with gob library")
		}
	}()
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, v := range c.items {
		gob.Register(v.Object)
	}
	err = enc.Encode(&c.items)
	return
}

//serialize to file
func (c *Cache) SaveToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err = c.Save(f); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

//load from io.Reader
func (c *Cache) Load(r io.Reader) error {
	dec := gob.NewDecoder(r)
	items := make(map[string]Item, 0)
	err := dec.Decode(&items)
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range items {
		obj, ok := c.items[k]
		if !ok || obj.IsExpired() {
			c.items[k] = v
		}
	}
	return nil
}

//load from file
func (c *Cache) LoadFromFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	if err = c.Load(f); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

//number of data items in the return cache
func (c *Cache) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

//clear the cache
func (c *Cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = map[string]Item{}
}

//stop gc
func (c *Cache) StopGc() {
	c.stopGc <- true
}

//create a new cache system
func NewCache(defaultExpiration, gcInterval time.Duration) (c *Cache) {
	c = &Cache{
		defaultExpiration: defaultExpiration,
		gcInterval:        gcInterval,
		items:             map[string]Item{},
		stopGc:            make(chan bool),
	}
	go c.gcLoop()
	return
}
