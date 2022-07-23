package bucket

import (
	"encoding/json"
	"errors"
)

type (
	Bucket struct {
		bucket map[string][]byte
	}
)

func New() *Bucket {
	return &Bucket{
		bucket: make(map[string][]byte),
	}
}

func (b *Bucket) Insert(uuid string, data []byte) error {

	if len(data) == 0 {
		return errors.New(dataNotFound)
	}

	var test interface{}

	if err := json.Unmarshal([]byte(data), &test); err != nil {
		return err
	}

	b.bucket[uuid] = data
	return nil
}

func (b Bucket) Get(uuid string, data interface{}) error {
	if val, ok := b.bucket[uuid]; ok {
		return json.Unmarshal(val, data)
	}
	return errors.New(bucketNotFound)
}

func (b *Bucket) Update(uuid string, data []byte) error {
	if _, ok := b.bucket[uuid]; ok {

		var test interface{}

		if err := json.Unmarshal([]byte(data), &test); err != nil {
			return err
		}

		b.bucket[uuid] = data
		return nil
	}
	return errors.New(bucketNotFound)
}

func (b *Bucket) Delete(uuid string) error {
	if _, ok := b.bucket[uuid]; ok {
		delete(b.bucket, uuid)
		return nil
	}
	return errors.New(bucketNotFound)
}
