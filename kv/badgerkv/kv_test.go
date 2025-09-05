package badgerkv_test

import (
	"testing"

	"github.com/dalloriam/ocl/kv/badgerkv"
)

func TestBadgerKV_Init(t *testing.T) {
	tmpDir := t.TempDir()
	kvStore, err := badgerkv.NewBadgerKV(tmpDir)
	if err != nil {
		t.Fatalf("Failed to initialize BadgerKV: %v", err)
	}
	defer kvStore.Close()
}

func TestBadgerKV_SetGet(t *testing.T) {
	tmpDir := t.TempDir()
	kvStore, err := badgerkv.NewBadgerKV(tmpDir)
	if err != nil {
		t.Fatalf("Failed to initialize BadgerKV: %v", err)
	}
	defer kvStore.Close()

	key := "testKey"
	value := []byte("testValue")

	// Test Set
	if err := kvStore.Set(key, value); err != nil {
		t.Fatalf("Failed to set key-value: %v", err)
	}

	// Test Get
	retrievedValue, err := kvStore.Get(key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}
	if string(retrievedValue) != string(value) {
		t.Fatalf("Expected value %s, got %s", value, retrievedValue)
	}
}

func TestBadgerKV_GetNonExistentKey(t *testing.T) {
	tmpDir := t.TempDir()
	kvStore, err := badgerkv.NewBadgerKV(tmpDir)
	if err != nil {
		t.Fatalf("Failed to initialize BadgerKV: %v", err)
	}
	defer kvStore.Close()

	_, err = kvStore.Get("nonExistentKey")
	if err == nil {
		t.Fatalf("Expected error when getting non-existent key, got nil")
	}
}

func TestBadgerKV_Delete(t *testing.T) {
	tmpDir := t.TempDir()
	kvStore, err := badgerkv.NewBadgerKV(tmpDir)
	if err != nil {
		t.Fatalf("Failed to initialize BadgerKV: %v", err)
	}
	defer kvStore.Close()

	key := "testKey"
	value := []byte("testValue")

	// Set a key-value pair
	if err := kvStore.Set(key, value); err != nil {
		t.Fatalf("Failed to set key-value: %v", err)
	}

	// Delete the key
	if err := kvStore.Delete(key); err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}

	// Try to get the deleted key
	_, err = kvStore.Get(key)
	if err == nil {
		t.Fatalf("Expected error when getting deleted key, got nil")
	}
}

func TestBadgerKV_Clear(t *testing.T) {
	tmpDir := t.TempDir()
	kvStore, err := badgerkv.NewBadgerKV(tmpDir)
	if err != nil {
		t.Fatalf("Failed to initialize BadgerKV: %v", err)
	}
	defer kvStore.Close()

	key := "testKey"
	value := []byte("testValue")

	// Set a key-value pair
	if err := kvStore.Set(key, value); err != nil {
		t.Fatalf("Failed to set key-value: %v", err)
	}

	// Clear the store
	if err := kvStore.Clear(); err != nil {
		t.Fatalf("Failed to clear store: %v", err)
	}

	// Try to get the cleared key
	_, err = kvStore.Get(key)
	if err == nil {
		t.Fatalf("Expected error when getting key after clear, got nil")
	}
}
