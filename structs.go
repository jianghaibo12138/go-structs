// Package structs
// @Description: 结构体通用操作包
package structs

import (
	"errors"
	"github.com/petersunbag/coven"
	"reflect"
)

//
// Structs
// @Description: 操作结构体
//
type Structs struct {
	IgnoreFields []string
	AliasFields  map[string]string
}

var (
	// ErrMustBePtr 待操作结构体必须以指针类型传入
	ErrMustBePtr = errors.New("src or dst must be ptr")
)

//
// New
// @Description: 实例化操作结构体
// @param ignore 忽略的结构体字段
// @return *Structs
//
func New(ignore []string) *Structs {
	return &Structs{
		IgnoreFields: ignore,
	}
}

//
// Map
// @Description: 结构体转换为map[string]interface{}
// @receiver s
// @param itf 带转换结构体
// @return map[string]interface{}
//
func (s *Structs) Map(itf interface{}) map[string]interface{} {
	v := reflect.ValueOf(itf)
	t := reflect.TypeOf(itf)
	if v.Kind() == reflect.Ptr { // 结构体指针
		v = v.Elem()
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	m := make(map[string]interface{})
	q := make([]interface{}, 0, 1)
	q = append(q, itf)

	for len(q) > 0 {
		v := reflect.ValueOf(q[0])
		if v.Kind() == reflect.Ptr { // 结构体指针
			v = v.Elem()
		}
		q = q[1:]
		tpy := v.Type()
		for i := 0; i < v.NumField(); i++ {
			vi := v.Field(i)
			if vi.Kind() == reflect.Ptr { // 内嵌指针
				vi = vi.Elem()
				if vi.Kind() == reflect.Struct { // 结构体
					q = append(q, vi.Interface())
				} else {
					fieldName := tpy.Field(i).Name
					if s.ignoreIndexOf(fieldName) == -1 {
						m[fieldName] = vi.Interface()
					}
				}
				continue
			}
			if vi.Kind() == reflect.Struct { // 内嵌结构体
				q = append(q, vi.Interface())
				continue
			}
			fieldName := tpy.Field(i).Name
			if s.ignoreIndexOf(fieldName) == -1 {
				m[fieldName] = vi.Interface()
			}
		}
	}
	return m
}

//
// Facade
// @Description: 结构体之间相互赋值
// @receiver s
// @param src 取值用结构体
// @param dst 赋值用结构体
// @return error
//
func (s *Structs) Facade(src, dst interface{}) error {
	if reflect.ValueOf(src).Kind() != reflect.Ptr || reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return ErrMustBePtr
	}

	c, err := coven.NewConverterOption(dst, src, &coven.StructOption{
		BannedFields: s.IgnoreFields,
		AliasFields:  s.AliasFields,
	})
	if err != nil {
		return err
	}
	err = c.Convert(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func (s *Structs) ignoreIndexOf(field string) int {
	for idx, f := range s.IgnoreFields {
		if f == field {
			return idx
		}
	}
	return -1
}
