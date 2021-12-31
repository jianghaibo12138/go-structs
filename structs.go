package structs

import (
	"reflect"
)

type Structs struct{}

func New() *Structs {
	return &Structs{}
}

func (s *Structs) Map(itf interface{}) map[string]interface{} {
	// 当前函数只接收struct类型
	v := reflect.ValueOf(itf)
	if v.Kind() == reflect.Ptr { // 结构体指针
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	m := make(map[string]interface{})
	queue := make([]interface{}, 0, 1)
	queue = append(queue, itf)

	for len(queue) > 0 {
		v := reflect.ValueOf(queue[0])
		if v.Kind() == reflect.Ptr { // 结构体指针
			v = v.Elem()
		}
		queue = queue[1:]
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			vi := v.Field(i)
			if vi.Kind() == reflect.Ptr { // 内嵌指针
				vi = vi.Elem()
				if vi.Kind() == reflect.Struct { // 结构体
					queue = append(queue, vi.Interface())
				} else {
					ti := t.Field(i)
					if tagValue := ti.Tag.Get("json"); tagValue != "" {
						// 存入map
						m[tagValue] = vi.Interface()
					}
				}
				break
			}
			if vi.Kind() == reflect.Struct { // 内嵌结构体
				queue = append(queue, vi.Interface())
				break
			}
			// 一般字段
			ti := t.Field(i)
			if tagValue := ti.Tag.Get("json"); tagValue != "" {
				// 存入map
				m[tagValue] = vi.Interface()
			}
		}
	}
	return m
}
