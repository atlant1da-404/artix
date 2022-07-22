package document

import (
	"errors"
	"github.com/atlant1da-404/artik_db/pkg/hash"
)

type Document struct {
	values map[string]interface{}
}

func New() *Document {
	return &Document{
		values: make(map[string]interface{}),
	}
}

func (d *Document) Create(data interface{}) (interface{}, error) {

	if data == nil {
		return data, errors.New(documentSizeErr)
	}

	uuid := hash.RandStringBytesMaskImprSrc(100)
	d.values[uuid] = data

	return data, nil
}
