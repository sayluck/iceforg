package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func GenerateUUID(n int) string {
	return GetRandomString(n)
}

func HashCode(s string) string {
	v := crc32.ChecksumIEEE([]byte(s))
	return strconv.FormatUint(uint64(v), 10)
}

func HashCodeUUID(n int) string {
	s := GenerateUUID(n)
	v := crc32.ChecksumIEEE([]byte(s))
	return strconv.FormatUint(uint64(v), 10)
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func PrettyJsonPrint(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	fmt.Println(out.String())
	return out.String()
}

func TramsStruct(source interface{}, target interface{}) error {
	sKind := reflect.ValueOf(source).Elem().Kind()
	tKind := reflect.ValueOf(target).Elem().Kind()
	if sKind == reflect.Ptr {
		sKind = reflect.ValueOf(source).Elem().Elem().Kind()
	}
	if tKind == reflect.Ptr {
		tKind = reflect.ValueOf(target).Elem().Elem().Kind()
	}

	if sKind != reflect.Struct ||
		tKind != reflect.Struct {
		err := fmt.Sprintf("source[%v] or target[%v]'s kind must be a struct", sKind, tKind)
		return errors.New(err)
	}

	data, err := json.Marshal(source)
	if err != nil {
		return errors.New("tramsStruct error," + err.Error())
	}
	return json.Unmarshal(data, target)
}
