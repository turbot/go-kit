package helpers

import (
	"sync"
)

// UpgradeRWMutex accepta  apointer to a RWMutex which is assumed to have a read lock
// it safely upgrades the read lock to a write lock
func UpgradeRWMutex(m *sync.RWMutex) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Lock()
	}()
	m.RUnlock()
	wg.Wait()
}
