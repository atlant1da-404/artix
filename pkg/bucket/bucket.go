package bucket

import (
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
		bucket map[string][]interface{}
		doc    *document.Document
	}
)

func New() *bucket {
	return &bucket{
		bucket: make(map[string][]interface{}),
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

func (b *bucket) InsertIntoBucket(name, data string) error {

	if _, ok := b.bucket[name]; ok {

		data, err := b.doc.Create(data)
		if err != nil {
			return err
		}

		b.bucket[name] = append(b.bucket[name], data)
		return nil
	}

	return errors.New(bucketNotFound)
}

func (b *bucket) GetAll(name string) ([]interface{}, error) {

	if data, ok := b.bucket[name]; ok {

		return data, nil
	}

	return nil, errors.New(bucketNotFound)
}
