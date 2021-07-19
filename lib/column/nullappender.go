package column

import (
	"fmt"
	"net"
	"reflect"
	"time"
)

var nullAppender = map[string]func(v interface{}, slice reflect.Value) (reflect.Value, error){
	"*int8": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(int8)
			if !ok {
				return slice, fmt.Errorf("canno assert to type int8")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *int8
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*int16": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(int16)
			if !ok {
				return slice, fmt.Errorf("cannot assert to type int16")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *int16
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*int32": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(int32)
			if !ok {
				return slice, fmt.Errorf("cannot assert to type int32")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *int32
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*int64": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(int64)
			if !ok {
				return slice, fmt.Errorf("cannot assert to type int64")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *int64
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*uint8": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(uint8)
			if !ok {
				return slice, fmt.Errorf("cannot assert to type uint8")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *uint8
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*uint16": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(uint16)
			if !ok {
				return slice, fmt.Errorf("canno assert to type uint16")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *uint16
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*uint32": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(uint32)
			if !ok {
				return slice, fmt.Errorf("canno assert to type uint32")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *uint32
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*uint64": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(uint64)
			if !ok {
				return slice, fmt.Errorf("canno assert to type uint64")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *uint64
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*float32": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(float32)
			if !ok {
				return slice, fmt.Errorf("canno assert to type float32")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *float32
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*float64": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(float64)
			if !ok {
				return slice, fmt.Errorf("canno assert to type float64")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *float64
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*string": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(string)
			if !ok {
				return slice, fmt.Errorf("canno assert to type string")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *string
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*time.Time": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(time.Time)
			if !ok {
				return slice, fmt.Errorf("canno assert to type time.Time")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *time.Time
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
	"*net.IP": func(v interface{}, slice reflect.Value) (reflect.Value, error) {
		if v != nil {
			v, ok := v.(net.IP)
			if !ok {
				return slice, fmt.Errorf("canno assert to type net.IP")
			}
			return reflect.Append(slice, reflect.ValueOf(&v)), nil
		}
		var vNil *net.IP
		return reflect.Append(slice, reflect.ValueOf(vNil)), nil
	},
}
