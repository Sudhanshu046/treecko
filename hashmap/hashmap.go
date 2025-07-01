package hashmap

import (
	"fmt"
	"hash/maphash"
)

// KeyValue struct holds the key-value pair.
type KeyValue[T any] struct {
	Key   string
	Value T
}

// HashMap struct represents a hash map.
type HashMap[T any] struct {
	buckets     [][]KeyValue[T]
	size        int
	threshold   float64 // Load factor threshold for resizing
	seed        maphash.Seed
	elementCount int // Number of elements in the map
}

// NewHashMap creates a new HashMap with the specified initial size.
func NewHashMap[T any](size int) *HashMap[T] {
	return &HashMap[T]{
		buckets:     make([][]KeyValue[T], size),
		size:        size,
		seed:        maphash.MakeSeed(),
		threshold:   0.7, // Load factor threshold for resizing
		elementCount: 0,
	}
}

// hash generates a hash of the key.
func (hm *HashMap[T]) hash(key string) uint64 {
	var h maphash.Hash
	h.SetSeed(hm.seed)
	h.WriteString(key)
	return h.Sum64() % uint64(hm.size)
}

// Set inserts a key-value pair into the map. Resizes the map if necessary.
func (hm *HashMap[T]) Set(key string, value T) {
	// Resize if the load factor exceeds the threshold
	if float64(hm.elementCount) / float64(hm.size) >= hm.threshold {
		hm.resize()
	}

	index := hm.hash(key)
	bucket := hm.buckets[index]

	// Check if the key exists and update its value if found
	for i, kv := range bucket {
		if kv.Key == key {
			bucket[i].Value = value
			return
		}
	}

	// Add new key-value pair
	hm.buckets[index] = append(bucket, KeyValue[T]{Key: key, Value: value})
	hm.elementCount++
}

// Get retrieves a value by key from the map.
func (hm *HashMap[T]) Get(key string) (T, error) {
	index := hm.hash(key)
	bucket := hm.buckets[index]

	for _, kv := range bucket {
		if key == kv.Key {
			return kv.Value, nil
		}
	}

	var emptyValue T
	return emptyValue, fmt.Errorf("key not found: %s", key)
}

// Put is similar to Set, but it returns the old value if it exists.
func (hm *HashMap[T]) Put(key string, value T) (T, error) {
	index := hm.hash(key)
	bucket := hm.buckets[index]

	for i, kv := range bucket {
		if key == kv.Key {
			oldValue := bucket[i].Value
			bucket[i].Value = value
			return oldValue, nil
		}
	}

	// Add new key-value pair
	hm.buckets[index] = append(bucket, KeyValue[T]{Key: key, Value: value})
	hm.elementCount++
	return value, nil
}

// Delete removes a key-value pair from the map.
func (hm *HashMap[T]) Delete(key string) error {
	index := hm.hash(key)
	bucket := hm.buckets[index]

	for i, kv := range bucket {
		if key == kv.Key {
			hm.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			hm.elementCount--
			return nil
		}
	}

	return fmt.Errorf("key not found: %s", key)
}

// resize doubles the size of the bucket array and rehashes all elements.
func (hm *HashMap[T]) resize() {
	// New size is double the current size
	newSize := hm.size * 2
	newBuckets := make([][]KeyValue[T], newSize)

	// Rehash all key-value pairs to the new buckets
	for _, bucket := range hm.buckets {
		for _, kv := range bucket {
			index := maphash.Hash{}
			index.SetSeed(hm.seed)
			index.WriteString(kv.Key)
			hashValue := index.Sum64() % uint64(newSize)
			newBuckets[hashValue] = append(newBuckets[hashValue], kv)
		}
	}

	// Update the HashMap with the new size and buckets
	hm.size = newSize
	hm.buckets = newBuckets
}

// Keys returns a list of all the keys in the map.
func (hm *HashMap[T]) Keys() []string {
	keys := []string{}
	for _, bucket := range hm.buckets {
		for _, kv := range bucket {
			keys = append(keys, kv.Key)
		}
	}
	return keys
}

// Values returns a list of all the values in the map.
func (hm *HashMap[T]) Values() []T {
	values := []T{}
	for _, bucket := range hm.buckets {
		for _, kv := range bucket {
			values = append(values, kv.Value)
		}
	}
	return values
}

// Len returns the number of elements in the map.
func (hm *HashMap[T]) Len() int {
	return hm.elementCount
}

// IsEmpty returns true if the map has no elements.
func (hm *HashMap[T]) IsEmpty() bool {
	return hm.Len() == 0
}

// Clear removes all key-value pairs from the map.
func (hm *HashMap[T]) Clear() {
	for i := range hm.buckets {
		hm.buckets[i] = nil
	}
	hm.elementCount = 0
}

// Clone creates a new map with the same elements as the current one.
func (hm *HashMap[T]) Clone() *HashMap[T] {
	clone := NewHashMap[T](hm.size)
	for _, bucket := range hm.buckets {
		for _, kv := range bucket {
			clone.Set(kv.Key, kv.Value)
		}
	}
	return clone
}

// ContainsKey checks if a key exists in the map.
func (hm *HashMap[T]) ContainsKey(key string) bool {
	index := hm.hash(key)
	bucket := hm.buckets[index]

	for _, kv := range bucket {
		if key == kv.Key {
			return true
		}
	}
	return false
}

// ContainsValue checks if a value exists in the map.
func (hm *HashMap[T]) ContainsValue(value T) bool {
	for _, bucket := range hm.buckets {
		for _, kv := range bucket {
			if kv.Value == value {
				return true
			}
		}
	}
	return false
}
