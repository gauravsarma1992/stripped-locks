package stlocks

import (
	"testing"
)

func TestNewLockHasher(t *testing.T) {
	lockHasher := NewLockHasher()

	if lockHasher != nil && lockHasher.concurrency != DefaultLockConcurrency {
		t.Errorf("Expected concurrency %d, got %d", DefaultLockConcurrency, lockHasher.concurrency)
	}

	for i, lock := range lockHasher.locks {
		if lock == nil {
			t.Errorf("Lock at index %d is nil", i)
		}
	}
}

func TestGetHashKey(t *testing.T) {
	lockHasher := NewLockHasher()

	testCases := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"Empty string", "", false},
		{"Non-empty string", "test-key", false},
		{"Long string", "this-is-a-very-long-test-key-to-ensure-it-works-with-longer-inputs", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hashSlot, err := lockHasher.GetHash(tc.input)

			if tc.expectError && err == nil {
				t.Error("Expected an error, but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if hashSlot >= lockHasher.concurrency {
				t.Errorf("Hash slot %d is out of range (concurrency: %d)", hashSlot, lockHasher.concurrency)
			}
		})
	}
}

func TestGetStore(t *testing.T) {
	lockHasher := NewLockHasher()

	testCases := []struct {
		name        string
		input       string
		expectError bool
	}{
		{"Empty string", "", false},
		{"Non-empty string", "test-key", false},
		{"Long string", "this-is-a-very-long-test-key-to-ensure-it-works-with-longer-inputs", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lockH, err := lockHasher.GetLockStore(tc.input)

			if tc.expectError && err == nil {
				t.Error("Expected an error, but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if lockH == nil {
				t.Error("Returned LockH is nil")
			}
		})
	}
}

func TestGetStoreDifferentKeys(t *testing.T) {
	lockHasher := NewLockHasher()

	store1, _ := lockHasher.GetLockStore("key1")
	store2, _ := lockHasher.GetLockStore("key2")

	if store1 == store2 {
		t.Error("Different keys should potentially return different stores")
	}
}

func TestGetStoreSameKey(t *testing.T) {
	lockHasher := NewLockHasher()

	store1, _ := lockHasher.GetLockStore("samekey")
	store2, _ := lockHasher.GetLockStore("samekey")

	if store1 != store2 {
		t.Error("Same key should return the same store")
	}
}
