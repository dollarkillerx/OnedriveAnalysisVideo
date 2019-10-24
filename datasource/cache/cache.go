/**
 * @Author: DollarKiller
 * @Description: cache
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:57 2019-10-24
 */
package cache

import "github.com/bluele/gcache"

var Cache gcache.Cache

func init() {
	Cache = gcache.New(20).
		LRU().
		Build()
}

func Get(key string) (interface{}, error) {
	get, e := Cache.Get(key)
	return get, e
}

func Set(key string, data interface{}) error {
	err := Cache.Set(key, data)
	return err
}

func Exis(key string) bool {
	_, e := Cache.Get(key)
	if e != nil {
		return false
	}
	return true
}
