package main

import (
	"log"

	"github.com/gauravsarma1992/stripped-locks/stlocks"
)

func main() {
	var (
		err     error
		lockHsh *stlocks.LockHasher

		storeOne *stlocks.LockStore
		storeTwo *stlocks.LockStore
	)

	if lockHsh = stlocks.NewLockHasher(); lockHsh == nil {
		log.Fatal(err)
	}

	log.Println("Fetching lock store using hash key")
	if storeOne, err = lockHsh.GetLockStore("test-key-one"); err != nil {
		log.Fatal(err)
	}

	if storeTwo, err = lockHsh.GetLockStore("test-key-two"); err != nil {
		log.Fatal(err)
	}

	log.Println("Fetching lock from store using lock name")
	mutOne, _ := storeOne.GetLock(stlocks.LockOne)
	mutTwo, _ := storeTwo.GetLock(stlocks.LockTwo)

	mutOne.RLock()
	defer mutOne.RUnlock()

	mutTwo.WLock()
	defer mutTwo.WUnlock()

}
