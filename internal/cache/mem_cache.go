package cache

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

var memCache *cache.Cache

func MemCacheInit() {
	log.Println("MemCacheInit")
	memCache = cache.New(-1, -1)
}

func MemCacheSet(key string, value interface{}, duration int64) {
	memCache.Set(key, value, time.Duration(duration)*time.Second)
}

func MemCacheSetNoExpiration(key string, value interface{}) {
	memCache.Set(key, value, -1)
}

func MemCacheGet(key string) (interface{}, bool) {
	return memCache.Get(key)
}
