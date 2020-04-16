package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadAndUnmarshal(filePath string, dst interface{}) error {
	_, getErr := os.Stat(filePath)
	if getErr != nil {
		_, getErr = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
		if getErr != nil {
			return getErr
		}
	}
	jsonContent, getErr := ioutil.ReadFile(filePath)
	if getErr != nil {
		return getErr
	}
	getErr = json.Unmarshal(jsonContent, &dst)
	if getErr != nil {
		return getErr
	}
	return nil
}

func MarshalAndSave(content interface{}, filePath string) error {
	jsonContent, err := json.Marshal(content)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, jsonContent, 0664)
	if err != nil {
		return err
	}
	return nil
}

func Struct2Map(src interface{}) map[string]interface{} {
	dst := make(map[string]interface{})
	// 原始复制
	/*key := reflect.TypeOf(src)
	value := reflect.ValueOf(src)
	for i := 0; i < key.NumField(); i++ {
		if value.Field(i).Interface() == "" {
			continue
		}
		dst[key.Field(i).Name] = value.Field(i).Interface()
	}
	return dst*/

	// 以 json 格式复制
	tmpJson, getErr := json.Marshal(src)
	if getErr != nil {
		fmt.Println(getErr)
	}
	getErr = json.Unmarshal(tmpJson, &dst)
	if getErr != nil {
		fmt.Println(getErr)
	}
	return dst
}
