# Treecko - A Go based Hashmap

A **generic**, **resizable**, and fully-featured hashmap implementation in Go, powered by Go 1.18+ generics. This package provides a customizable and type-safe alternative to Go's built-in maps, with support for dynamic resizing, collision handling, and utility methods such as cloning, clearing, and key/value inspection.

> 📦 Package Path: `github.com/Sudhanshu046/treecko/hashmap`

---

## 📥 Installation

To install the package, run:

```bash
go get github.com/Sudhanshu046/treecko/hashmap
```

## 🚀 Getting Started
```
import "github.com/Sudhanshu046/treecko/hashmap"   
```

## ✨ Features

*   🧠 Generic support using Go 1.18+ type parameters
    
*   📈 Automatic resizing based on load factor
    
*   🔍 Fast lookup, insertion, and deletion
    
*   🔄 Utility methods for cloning, clearing, checking keys/values, etc.
    

## 🛠️ Usage Example
``` 
package main

import (
	"fmt"
	"github.com/Sudhanshu046/treecko/hashmap"
)

func main() {
	// Create a new hashmap with initial size 8
	hm := hashmap.NewHashMap[string](2)

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

## 🧪 Requirements

*   Go 1.18 or later (for generics support)
