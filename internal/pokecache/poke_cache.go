package pokecache

import (
	"sync"
	"time"
)

type PokeCacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type PokeCache struct {
	pokeCacheMap map[string]PokeCacheEntry
	mutex        sync.Mutex
}

func NewCache(interval time.Duration) *PokeCache {
	pokeCache := PokeCache{
		pokeCacheMap: make(map[string]PokeCacheEntry),
		mutex:        sync.Mutex{},
	}

	go pokeCache.ReapLoop(interval)
	return &pokeCache
}

func (p *PokeCache) Add(key string, value []byte) {
	p.mutex.Lock()
	p.pokeCacheMap[key] = PokeCacheEntry{
		CreatedAt: time.Now(),
		Val:       value,
	}
	p.mutex.Unlock()
}

func (p *PokeCache) Get(key string) ([]byte, bool) {
	p.mutex.Lock()

	val, bol := p.pokeCacheMap[key]

	p.mutex.Unlock()
	return val.Val, bol
}

func (p *PokeCache) ReapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		p.mutex.Lock()
		for key, entry := range p.pokeCacheMap {
			if time.Since(entry.CreatedAt) > duration {
				delete(p.pokeCacheMap, key)
			}
		}
		p.mutex.Unlock()
	}
}
