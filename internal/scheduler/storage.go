package scheduler

type Storage interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
	Del(string) error
}

type mapStorage struct {
	mp map[string]interface{}
}

func NewMapStorage() Storage {
	return &mapStorage{mp: make(map[string]interface{})}
}

func (ms *mapStorage) Get(key string) (interface{}, error) {
	return ms.mp[key], nil
}

func (ms *mapStorage) Set(key string, value interface{}) error {
	ms.mp[key] = value
	return nil
}

func (ms *mapStorage) Del(key string) error {
	delete(ms.mp, key)
	return nil
}
