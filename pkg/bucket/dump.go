package bucket

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Dump(dumpName string, buck Bucket) error {

	file, err := json.MarshalIndent(buck.bucket, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dumpName, file, 0644)
}

func LoadDump(dumpName string, buck *Bucket) error {

	jsonFile, err := os.Open(dumpName)
	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return json.Unmarshal(byteValue, &buck.bucket)
}
