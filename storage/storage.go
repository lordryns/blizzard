package storage

import (
	"errors"
	"fmt"
	"sync"
)

var mutex sync.Mutex

type StoreObject struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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
	fmt.Printf("Total stores: %v\n", len(storage.Pool))

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

func (storage *Storage) Get(id, key string) (map[string]StoreObject, error) {

	var userStore, ok = storage.Pool[id]
	if !ok {
		return make(map[string]StoreObject), errors.New("Store does not exist!")
	}

	var userStoreList = make(map[string]StoreObject)
	if len(key) > 0 {
		userStoreList[key] = userStore.Store[key]
		return userStoreList, nil
	} else {
		return userStore.Store, nil
	}
}

func (storage *Storage) Delete(id, key string) bool {
	mutex.Lock()
	var userStore, ok = storage.Pool[id]

	if !ok {
		mutex.Unlock()
		return false
	}

	delete(userStore.Store, key)
	mutex.Unlock()
	return true
}
