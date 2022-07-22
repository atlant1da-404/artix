package bucket

import (
	"encoding/json"
	"errors"
	"github.com/atlant1da-404/artik_db/pkg/document"
)

type (
	Bucket interface {
		CreateBucket(name string) error
		InsertIntoBucket(bucket, data string) error
	}
	bucket struct {
		Bucket
		bucket map[string][]byte
		doc    *document.Document
	}
)

func New() *bucket {
	return &bucket{
		bucket: make(map[string][]byte),
		doc:    document.New(),
	}
}

func (b *bucket) CreateBucket(name string) error {

	if name == "" {
		return errors.New(nameValidationErr)
	}

	b.bucket[name] = nil
	return nil
}

func (b *bucket) InsertIntoBucket(name string, data []byte) error {

	if _, ok := b.bucket[name]; ok {

		for _, docs := range data {
			b.bucket[name] = append(b.bucket[name], docs)
		}

		return nil
	}

	return errors.New(bucketNotFound)
}

func (b *bucket) GetAllFromBucket(name string, data interface{}) error {

	if val, ok := b.bucket[name]; ok {
		return json.Unmarshal(val, data)
	}

	return errors.New(bucketNotFound)
}
