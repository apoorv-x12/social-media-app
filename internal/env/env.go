package env

import (
	"os"
	"strconv"
)

func GetStr(key string, alternative string) string {
	value,ok:=os.LookupEnv(key)
	if !ok {
		return alternative
	}
	return value
}

func GetInt(key string,alternative int) int {
	value,ok:=os.LookupEnv(key)
	if !ok {
		return alternative
	}

	valueInt,err:=strconv.Atoi(value)
	if err!=nil {
		return alternative
	}
	return valueInt
}