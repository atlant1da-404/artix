package document

import (
	"errors"
	"fmt"
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

func (d *Document) Create(data []byte) (interface{}, error) {

	if data == nil {
		return data, errors.New(documentSizeErr)
	}

	uuid := hash.RandStringBytesMaskImprSrc(100)
	d.values[uuid] = data

	return data, nil
}

func (d *Document) GetAll(name string, buck []byte, data interface{}) error {

	fmt.Println(buck)
	return nil
}
