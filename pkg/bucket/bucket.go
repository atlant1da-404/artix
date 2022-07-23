package bucket

import (
	"encoding/json"
	"errors"
)

type (
	bucket struct {
		bucket map[string][]byte
	}
)

func New() *bucket {
	return &bucket{
		bucket: make(map[string][]byte),
	}
}

func (b *bucket) Insert(uuid string, data []byte) {
	b.bucket[uuid] = data
}

func (b bucket) Get(uuid string, data interface{}) error {
	if val, ok := b.bucket[uuid]; ok {
		return json.Unmarshal(val, data)
	}
	return errors.New(bucketNotFound)
}

func (b *bucket) Update(uuid string, data []byte) error {
	if _, ok := b.bucket[uuid]; ok {
		b.bucket[uuid] = data
		return nil
	}
	return errors.New(bucketNotFound)
}

func (b *bucket) Delete(uuid string) error {
	if _, ok := b.bucket[uuid]; ok {
		delete(b.bucket, uuid)
		return nil
	}
	return errors.New(bucketNotFound)
}
