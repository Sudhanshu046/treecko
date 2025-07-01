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
```
import "github.com/Sudhanshu046/treecko/hashmap"   
```

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

🧪 Requirements
---------------

*   Go 1.18 or later (for generics support)
