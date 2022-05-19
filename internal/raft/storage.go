package raft

import (
	"log"
	"reflect"
	"strings"
)

type Storage interface {
	Get(string) (string, error)
	Set(string, string)
	Del(string) error
	Load(interface{})
	Copy() interface{}
	PrefixAll(string) *map[string]string
	LoadPrefix(string, map[string]string)
	RemovePrefix(string)
	StorageType() reflect.Type
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

func (s *rfStorage) StorageType() reflect.Type {
	return reflect.TypeOf(s.storage)
}

func (s *rfStorage) Copy() interface{} {
	return s.storage
}

func (s *rfStorage) Load(data interface{}) {
	s.storage = data.(map[string]string)
}

func (s *rfStorage) Get(key string) (string, error) {
	v, ok := s.storage[key]
	if !ok {
		return "", KeyNotFoundErr
	}
	return v, nil
}

func (s *rfStorage) PrefixAll(prefix string) *map[string]string {
	m := map[string]string{}
	for k, v := range s.storage {
		if strings.HasPrefix(k, prefix) {
			m[k] = v
		}
	}
	return &m
}

func (s *rfStorage) LoadPrefix(prefix string, m map[string]string) {
	for key, value := range m {
		s.storage[key] = value
	}
}

func (s *rfStorage) RemovePrefix(prefix string) {
	for key, _ := range s.storage {
		if prefix == key {
			delete(s.storage, key)
		}
	}
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

type namespaceStorage struct {
	storage map[string]map[string]string
}

func NewNamespaceStorage() Storage {
	s := namespaceStorage{storage: make(map[string]map[string]string)}
	defaultMap := make(map[string]string)
	s.storage["default"] = defaultMap
	return &s
}

func (s *namespaceStorage) StorageType() reflect.Type {
	return reflect.TypeOf(s.storage)
}

func (s *namespaceStorage) Copy() interface{} {
	return s.storage
}

func (s *namespaceStorage) Load(data interface{}) {
	s.storage = data.(map[string]map[string]string)
}

func (s *namespaceStorage) Get(key string) (string, error) {
	prefix := s.getPrefix(key)
	var m map[string]string
	var ok bool
	if m, ok = s.storage[prefix]; !ok {
		return "", KeyNotFoundErr
	}
	v, ok := m[key]
	if !ok {
		return "", KeyNotFoundErr
	}
	return v, nil
}

func (s *namespaceStorage) PrefixAll(prefix string) *map[string]string {
	if m, ok := s.storage[prefix]; ok {
		return &m
	}
	m := make(map[string]string)
	return &m
}

func (s *namespaceStorage) LoadPrefix(prefix string, m map[string]string) {
	s.storage[prefix] = m
}

func (s *namespaceStorage) RemovePrefix(prefix string) {
	if _, ok := s.storage[prefix]; ok {
		delete(s.storage, prefix)
	}
}

func (s *namespaceStorage) Set(key string, value string) {
	prefix := s.getPrefix(key)
	log.Printf("prefix: %s\n", prefix)
	var m map[string]string
	var ok bool
	if m, ok = s.storage[prefix]; !ok {
		m = make(map[string]string)
	}
	m[key] = value
	s.storage[prefix] = m
}

func (s *namespaceStorage) Del(key string) error {
	prefix := s.getPrefix(key)
	var m map[string]string
	var ok bool
	if m, ok = s.storage[prefix]; !ok {
		return KeyNotFoundErr
	}
	if _, ok := m[key]; !ok {
		return KeyNotFoundErr
	}
	delete(m, key)
	s.storage[prefix] = m
	return nil
}

func (s *namespaceStorage) getPrefix(key string) string {
	strArr := strings.Split(key, ".")
	if len(strArr) > 1 {
		return strArr[0]
	}
	return "default"
}
