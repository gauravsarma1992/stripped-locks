package stlocks

import (
	"testing"
)

func TestNewLockH(t *testing.T) {
	lockH := NewLockStore()

	if lockH == nil {
		t.Fatal("NewLockStore returned nil")
	}

	if lockH.lockCount != 2 {
		t.Errorf("Expected lockCount to be 2, got %d", lockH.lockCount)
	}

	expectedLocks := []LockName{LockOne, LockTwo}
	for _, lockName := range expectedLocks {
		if lockH.hash[uint8(lockName)] == nil {
			t.Errorf("Expected lock %d to be initialized", lockName)
		}
	}
}

func TestSetup(t *testing.T) {
	lockH := &LockStore{
		hash:      [32]*Lock{},
		lockCount: 0,
	}

	err := lockH.setup()
	if err != nil {
		t.Fatalf("Unexpected error in setup: %v", err)
	}

	if lockH.lockCount != 2 {
		t.Errorf("Expected lockCount to be 2, got %d", lockH.lockCount)
	}

	expectedLocks := []LockName{LockOne, LockTwo}
	for _, lockName := range expectedLocks {
		if lockH.hash[uint8(lockName)] == nil {
			t.Errorf("Expected lock %d to be initialized", lockName)
		}
	}
}

func TestAddLock(t *testing.T) {
	lockH := NewLockStore()

	// Test adding a new lock
	newLockName := LockName(3)
	lock, err := lockH.AddLock(newLockName)
	if err != nil {
		t.Fatalf("Unexpected error adding new lock: %v", err)
	}
	if lock == nil {
		t.Fatal("addLock returned nil lock")
	}
	if lockH.lockCount != 3 {
		t.Errorf("Expected lockCount to be 3, got %d", lockH.lockCount)
	}

	// Test adding an existing lock
	_, err = lockH.AddLock(LockOne)
	if err == nil {
		t.Error("Expected error when adding existing lock, got nil")
	}
}

func TestGetLock(t *testing.T) {
	lockH := NewLockStore()

	// Test getting an existing lock
	lock, err := lockH.GetLock(LockOne)
	if err != nil {
		t.Fatalf("Unexpected error getting existing lock: %v", err)
	}
	if lock == nil {
		t.Fatal("getLock returned nil for existing lock")
	}

	// Test getting a non-existent lock
	nonExistentLock := LockName(5)
	_, err = lockH.GetLock(nonExistentLock)
	if err == nil {
		t.Error("Expected error when getting non-existent lock, got nil")
	}
}

func TestRemoveLock(t *testing.T) {
	lockH := NewLockStore()

	// Test removing an existing lock
	err := lockH.RemoveLock(LockOne)
	if err != nil {
		t.Fatalf("Unexpected error removing existing lock: %v", err)
	}
	if lockH.lockCount != 1 {
		t.Errorf("Expected lockCount to be 1, got %d", lockH.lockCount)
	}

	// Test removing the same lock again
	err = lockH.RemoveLock(LockOne)
	if err == nil {
		t.Error("Expected error when removing non-existent lock, got nil")
	}

	// Test removing a non-existent lock
	nonExistentLock := LockName(5)
	err = lockH.RemoveLock(nonExistentLock)
	if err == nil {
		t.Error("Expected error when removing non-existent lock, got nil")
	}
}
