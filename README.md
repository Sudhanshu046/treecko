# 🧠 Treecko - A Go based Hashmap

A **generic** and **resizable** hashmap implementation in Go, using Go 1.18+ generics. Built for performance, usability, and educational clarity.

> 📦 Package Path: `github.com/Sudhanshu046/treecko/hashmap`

---

## 📥 Installation

To install the package, run:

```bash
go get github.com/Sudhanshu046/treecko/hashmap
```

Getting Started
------------------
`import "github.com/Sudhanshu046/treecko/hashmap"   `

Features
----------

*   🧠 Generic support using Go 1.18+ type parameters
    
*   📈 Automatic resizing based on load factor
    
*   🔍 Fast lookup, insertion, and deletion
    
*   🔄 Utility methods for cloning, clearing, checking keys/values, etc.
    

Usage Example
-----------------
``` 
package main

import (
	"fmt"
	"github.com/<yourusername>/treecko/hashmap"
)

func main() {
	// Create a new hashmap with initial size 8
	hm := hashmap.NewHashMap

	hm.Set("name", "Treecko")
	hm.Set("type", "Grass")

	// Get a value
	val, err := hm.Get("name")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Name:", val)
	}

	// Check existence
	fmt.Println("Contains key 'type'? →", hm.ContainsKey("type"))

	// Show all keys
	fmt.Println("Keys:", hm.Keys())

	// Clone the map
	clone := hm.Clone()
	fmt.Println("Cloned Keys:", clone.Keys())
}
```

API Reference
----------------

### ✅ NewHashMap\[T any\](size int) \*HashMap\[T\]

Creates a new hash map of the given initial size.

### ✅ Set(key string, value T)

Adds or updates a key-value pair.

### ✅ Get(key string) (T, error)

Retrieves the value for a given key. Returns an error if the key doesn't exist.

### ✅ Put(key string, value T) (T, error)

Updates a key’s value and returns the **old value** if the key existed, else returns the new value.

### ✅ Delete(key string) error

Deletes the key-value pair. Returns an error if the key doesn't exist.

### ✅ Keys() \[\]string

Returns a list of all keys in the hashmap.

### ✅ Values() \[\]T

Returns a list of all values in the hashmap.

### ✅ Len() int

Returns the number of key-value pairs in the hashmap.

### ✅ IsEmpty() bool

Returns true if the hashmap has no elements.

### ✅ Clear()

Removes all key-value pairs from the hashmap.

### ✅ Clone() \*HashMap\[T\]

Creates a shallow copy of the hashmap.

### ✅ ContainsKey(key string) bool

Checks whether the key exists in the hashmap.

### ✅ ContainsValue(value T) bool

Checks whether the value exists in the hashmap.

🧪 Requirements
---------------

*   Go 1.18 or later (for generics support)


