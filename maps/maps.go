package maps

import (
	"sync"
)

//
// SafeInterfaceMap
// map[string]interface{} with locking to allow multithreaded access
type SafeInterfaceMap struct {
	data map[string]interface{}
	mux  sync.Mutex
}

func CreateSafeInterfaceMap() *SafeInterfaceMap {
	ret := &SafeInterfaceMap{}
	ret.data = make(map[string]interface{})
	return ret
}
func (safeMap *SafeInterfaceMap) Set(key string, val interface{}) {
	safeMap.mux.Lock()
	defer func() {
		safeMap.mux.Unlock()
	}()

	safeMap.data[key] = val
}

func (safeMap *SafeInterfaceMap) Get(key string) (val interface{}, ok bool) {
	safeMap.mux.Lock()
	defer func() {
		safeMap.mux.Unlock()
	}()

	val, ok = safeMap.data[key]
	return
}

func (safeMap *SafeInterfaceMap) Map() map[string]interface{} {
	safeMap.mux.Lock()
	defer func() {
		safeMap.mux.Unlock()
	}()
	return safeMap.data
}

//
// SafeStringMap
// map[string]string with locking to allow multithreaded access
type SafeStringMap struct {
	data map[string]string
	mux  sync.Mutex
}

func CreateSafeStringMap() *SafeStringMap {
	ret := &SafeStringMap{}
	ret.data = make(map[string]string)
	return ret
}

func (safeMap *SafeStringMap) Set(key string, val string) {
	safeMap.mux.Lock()
	defer func() {
		safeMap.mux.Unlock()
	}()

	safeMap.data[key] = val
}

func (safeMap *SafeStringMap) Get(key string) (val interface{}, ok bool) {
	safeMap.mux.Lock()
	defer func() {
		safeMap.mux.Unlock()
	}()

	val, ok = safeMap.data[key]
	return
}

func (safeMap *SafeStringMap) Map() map[string]string {
	safeMap.mux.Lock()
	defer func() {
		safeMap.mux.Unlock()
	}()
	return safeMap.data
}
