package verify

import (
	"fmt"
	"reflect"
)

const (
	// ValidTag struct tag
	ValidTag = "valid"

	wordsize = 32 << (^uint(0) >> 32 & 1)
)

type Error struct {
	Message, Key, Name, Field, Tmpl string
	Value                           interface{}
	LimitValue                      interface{}
}

type ValidFunc struct {
	Name   string
	Params []interface{}
}

type Verify struct {
}

func (v *Verify) Valid(obj interface{}) (bool, error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	switch {
	case isStruct(objT):
	case isStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return false, fmt.Errorf("%v must be a struct or a struct pointer", obj)
	}
	for i := 0; i < objT.NumField(); i++ {
		tag := getValidTag(objT.Field(i))
		fmt.Println(tag)
	}
	return true, nil
}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func getValidTag(f reflect.StructField) string {
	tag := f.Tag.Get(ValidTag)
	if len(tag) == 0 {
		return ""

	}
	return tag
}

//func getRegFuncs(tag, key string) (vfs []ValidFunc, str string, err error) {
//	tag = strings.TrimSpace(tag)
//	index := strings.Index(tag, "Match(/")
//	if index == -1 {
//		str = tag
//		return
//	}
//	end := strings.LastIndex(tag, "/)")
//	if end < index {
//		err = fmt.Errorf("invalid Match function")
//		return
//	}
//	reg, err := regexp.Compile(tag[index+len("Match(/") : end])
//	if err != nil {
//		return
//	}
//	vfs = []ValidFunc{{"Match", []interface{}{reg, key + ".Match"}}}
//	str = strings.TrimSpace(tag[:index]) + strings.TrimSpace(tag[end+len("/)"):])
//	return
//}

//
//func parseFunc(vfunc, key string) (v ValidFunc, err error) {
//	defer func() {
//		if r := recover(); r != nil {
//			err = fmt.Errorf("%v", r)
//		}
//	}()
//
//	vfunc = strings.TrimSpace(vfunc)
//	start := strings.Index(vfunc, "(")
//	var num int
//
//	// doesn't need parameter valid function
//	if start == -1 {
//		if num, err = numIn(vfunc); err != nil {
//			return
//		}
//		if num != 0 {
//			err = fmt.Errorf("%s require %d parameters", vfunc, num)
//			return
//		}
//		v = ValidFunc{vfunc, []interface{}{key + "." + vfunc}}
//		return
//	}
//
//	end := strings.Index(vfunc, ")")
//	if end == -1 {
//		err = fmt.Errorf("invalid valid function")
//		return
//	}
//
//	name := strings.TrimSpace(vfunc[:start])
//	if num, err = numIn(name); err != nil {
//		return
//	}
//
//	params := strings.Split(vfunc[start+1:end], ",")
//	// the num of param must be equal
//	if num != len(params) {
//		err = fmt.Errorf("%s require %d parameters", name, num)
//		return
//	}
//
//	tParams, err := trim(name, key+"."+name, params)
//	if err != nil {
//		return
//	}
//	v = ValidFunc{name, tParams}
//	return
//}
//
//func numIn(name string) (num int, err error) {
//	fn, ok := funcs[name]
//	if !ok {
//		err = fmt.Errorf("doesn't exsits %s valid function", name)
//		return
//	}
//	// sub *Validation obj and key
//	num = fn.Type().NumIn() - 3
//	return
//}
//
//func trim(name, key string, s []string) (ts []interface{}, err error) {
//	ts = make([]interface{}, len(s), len(s)+1)
//	fn, ok := funcs[name]
//	if !ok {
//		err = fmt.Errorf("doesn't exsits %s valid function", name)
//		return
//	}
//	for i := 0; i < len(s); i++ {
//		var param interface{}
//		// skip *Validation and obj params
//		if param, err = parseParam(fn.Type().In(i+2), strings.TrimSpace(s[i])); err != nil {
//			return
//		}
//		ts[i] = param
//	}
//	ts = append(ts, key)
//	return
//}
