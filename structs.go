package structs

import (
	"reflect"
)

type Structs struct {
}

func New(s interface{}) {

}

func Map(s interface{}) map[string]interface{} {
	var m = make(map[string]interface{})
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { // 判断是否为可导出字段
			// 判断是否是嵌套结构
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					m[structField.Field(j).Name] = v.Field(i).Field(j).Interface()
				}
				continue
			}
			m[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return m
}
