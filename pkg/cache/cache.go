package cache

type Cache interface {
	Get(key string, target interface{}) error
	Set(key string, data interface{}) error
}

type cache struct {

}
