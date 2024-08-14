package stlocks

import (
	"fmt"
	"sync"
)

const (
	LockOne LockName = LockName(1)
	LockTwo LockName = LockName(2)
)

type (
	// Used to define different locks for different functionalities
	LockName uint8

	Lock struct {
		mutex    *sync.RWMutex
		name     LockName
		refCount uint32
	}

	LockStore struct {
		hash      [32]*Lock
		lockCount uint8
	}
)

func NewLockStore() (lockSt *LockStore) {
	lockSt = &LockStore{
		hash:      [32]*Lock{},
		lockCount: 0,
	}
	if err := lockSt.setup(); err != nil {
		return
	}
	return
}

func (lockSt *LockStore) setup() (err error) {
	availableLocks := []LockName{
		LockOne,
		LockTwo,
	}
	for _, lockName := range availableLocks {
		if _, err = lockSt.AddLock(lockName); err != nil {
			return
		}
	}
	return
}

func (lockSt *LockStore) AddLock(name LockName) (lock *Lock, err error) {
	lock = &Lock{
		mutex:    &sync.RWMutex{},
		name:     name,
		refCount: 0,
	}
	if lockSt.hash[uint8(lock.name)] != nil {
		err = fmt.Errorf("slot already filled for %d", lock.name)
		return
	}
	lockSt.hash[uint8(lock.name)] = lock
	lockSt.lockCount++
	return
}

func (lockSt *LockStore) GetLock(name LockName) (lock *Lock, err error) {
	if lock = lockSt.hash[uint8(name)]; lock == nil {
		err = fmt.Errorf("lock not found for %d", name)
		return
	}
	return
}

func (lockSt *LockStore) RemoveLock(name LockName) (err error) {
	if lockSt.hash[uint8(name)] == nil {
		err = fmt.Errorf("lock not found for %d", name)
		return
	}
	lockSt.hash[uint8(name)] = nil
	lockSt.lockCount--
	return
}

func (lock *Lock) WLock() {
	lock.mutex.Lock()
	lock.refCount++
}

func (lock *Lock) WUnlock() {
	lock.mutex.Unlock()
	lock.refCount--
}

func (lock *Lock) RLock() {
	lock.mutex.RLock()
	lock.refCount++
}

func (lock *Lock) RUnlock() {
	lock.mutex.RLock()
	lock.refCount--
}
