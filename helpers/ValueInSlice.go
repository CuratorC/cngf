package helpers

import "github.com/spf13/cast"

func IntInSlice(value int, slice []int) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func UintInSlice(value uint, slice []uint) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func Int8InSlice(value int8, slice []int8) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func Uint8InSlice(value uint8, slice []uint8) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func Int16InSlice(value int16, slice []int16) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func Uint16InSlice(value uint16, slice []uint16) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func Int32InSlice(value int32, slice []int32) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func Uint32InSlice(value uint32, slice []uint32) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func Int64InSlice(value int64, slice []int64) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}

func Uint64InSlice(value uint64, slice []uint64) bool {
	if Empty(slice) {
		return false
	}
	key := cast.ToString(value)
	m := cast.ToStringMap(slice)
	if _, ok := m[key]; !ok {
		return false
	}
	return true
}
