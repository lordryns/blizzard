package storage

import "fmt"
import "sync"

var mutex sync.Mutex

type StoreObject struct {
	Key   string
	Value string
}

type UserStore struct {
	Store map[string]StoreObject
}

type Storage struct {
	Pool map[string]UserStore
}

func NewStorage() *Storage {
	var newContent = make(map[string]UserStore)
	return &Storage{Pool: newContent}
}

func (storage *Storage) Create(id string) bool {
	mutex.Lock()
	var _, ok = storage.Pool[id]
	if ok {
		mutex.Unlock()
		return false
	}

	storage.Pool[id] = UserStore{Store: make(map[string]StoreObject)}
	fmt.Println(storage)

	mutex.Unlock()
	return true
}

func (storage *Storage) Add(id string, object StoreObject) bool {
	mutex.Lock()
	var userStore, ok = storage.Pool[id]
	if !ok {
		mutex.Unlock()
		return false
	}
	userStore.Store[object.Key] = object
	mutex.Unlock()
	return true
}
