# ancache

## Introduction
This is an implementation of an LRU (Least Recently Used) cache in Go. The cache has a maximum size and an expiration time for its entries. When the cache reaches its maximum size, the least recently used item will be removed to make room for a new item. If an item has been in the cache for longer than its expiration time, it will also be removed.

## Installation
To install the package, run the following command:

```
go get github.com/lmousom/ancache
```
## Usage
To use the cache, import the package and create a new cache with the Create function:

```go
import (
	"time"
	"ancache"
)

cache := ancache.Create(100, time.Hour)
```

This will create a new cache with a maximum size of 100 items and an expiration time of 1 hour.

You can then use the following methods to manipulate the cache:

```` `Set(key interface{}, value interface{})` ````: adds a key-value pair to the cache. If the key already exists in the cache, its value will be updated.
```` `Get(key interface{}) interface{}` ````: retrieves the value for the given key from the cache. If the key does not exist or the value has expired, this method will return nil.
```` `Clear()` ````: removes all items from the cache.

## Example
Here is an example of how to use the cache:

````go
package main

import (
	"time"
	"fmt"
	"ancache"
)

func main() {
	// Create a new cache with a maximum size of 100 items and an expiration time of 1 hour
	cache := ancache.Create(100, time.Hour)

	// Set a value in the cache
	cache.Set("key", "value")

	// Get the value from the cache
	value := cache.Get("key")
	fmt.Println(value) // Output: "value"

	// Wait 2 hours
	time.Sleep(2 * time.Hour)

	// Get the value from the cache again
	value = cache.Get("key")
	fmt.Println(value) // Output: nil (since the value has expired)
}
````

## License
This package is licensed under the MIT License.
