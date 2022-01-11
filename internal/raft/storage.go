package raft

type Storage interface {
	Get(string) (string, error)
	Set(string, string)
	Del(string) error
	Load(*map[string]string)
	Copy() map[string]string
}

type rfStorage struct {
	storage map[string]string
}

func NewRaftStorage() Storage {
	s := rfStorage{
		make(map[string]string),
	}
	return &s
}

func (s *rfStorage) Copy() map[string]string {
	return s.storage
}

func (s *rfStorage) Load(data *map[string]string) {
	s.storage = *data
}

func (s *rfStorage) Get(key string) (string, error) {
	v, ok := s.storage[key]
	if !ok {
		return "", KeyNotFoundErr
	}
	return v, nil
}

func (s *rfStorage) Set(key string, value string) {
	s.storage[key] = value
}

func (s *rfStorage) Del(key string) error {
	if _, ok := s.storage[key]; !ok {
		return KeyNotFoundErr
	}
	delete(s.storage, key)
	return nil
}
