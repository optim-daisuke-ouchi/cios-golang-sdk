/*
 * Collection utility of NullableMultipleDeviceModelEntity Struct
 *
 * Generated by: Go Streamer
 */

package cios

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type NullableMultipleDeviceModelEntityStream []NullableMultipleDeviceModelEntity

func NullableMultipleDeviceModelEntityStreamOf(arg ...NullableMultipleDeviceModelEntity) NullableMultipleDeviceModelEntityStream {
	return arg
}
func NullableMultipleDeviceModelEntityStreamFrom(arg []NullableMultipleDeviceModelEntity) NullableMultipleDeviceModelEntityStream {
	return arg
}
func CreateNullableMultipleDeviceModelEntityStream(arg ...NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	tmp := NullableMultipleDeviceModelEntityStreamOf(arg...)
	return &tmp
}
func GenerateNullableMultipleDeviceModelEntityStream(arg []NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	tmp := NullableMultipleDeviceModelEntityStreamFrom(arg)
	return &tmp
}

func (self *NullableMultipleDeviceModelEntityStream) Add(arg NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	return self.AddAll(arg)
}
func (self *NullableMultipleDeviceModelEntityStream) AddAll(arg ...NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) AddSafe(arg *NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Aggregate(fn func(NullableMultipleDeviceModelEntity, NullableMultipleDeviceModelEntity) NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	result := NullableMultipleDeviceModelEntityStreamOf()
	self.ForEach(func(v NullableMultipleDeviceModelEntity, i int) {
		if i == 0 {
			result.Add(fn(NullableMultipleDeviceModelEntity{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) AllMatch(fn func(NullableMultipleDeviceModelEntity, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleDeviceModelEntityStream) AnyMatch(fn func(NullableMultipleDeviceModelEntity, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableMultipleDeviceModelEntityStream) Clone() *NullableMultipleDeviceModelEntityStream {
	temp := make([]NullableMultipleDeviceModelEntity, self.Len())
	copy(temp, *self)
	return (*NullableMultipleDeviceModelEntityStream)(&temp)
}
func (self *NullableMultipleDeviceModelEntityStream) Copy() *NullableMultipleDeviceModelEntityStream {
	return self.Clone()
}
func (self *NullableMultipleDeviceModelEntityStream) Concat(arg []NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	return self.AddAll(arg...)
}
func (self *NullableMultipleDeviceModelEntityStream) Contains(arg NullableMultipleDeviceModelEntity) bool {
	return self.FindIndex(func(_arg NullableMultipleDeviceModelEntity, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableMultipleDeviceModelEntityStream) Clean() *NullableMultipleDeviceModelEntityStream {
	*self = NullableMultipleDeviceModelEntityStreamOf()
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Delete(index int) *NullableMultipleDeviceModelEntityStream {
	return self.DeleteRange(index, index)
}
func (self *NullableMultipleDeviceModelEntityStream) DeleteRange(startIndex, endIndex int) *NullableMultipleDeviceModelEntityStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Distinct() *NullableMultipleDeviceModelEntityStream {
	caches := map[string]bool{}
	result := NullableMultipleDeviceModelEntityStreamOf()
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if !f {
				result = append(result, v)
			}
		} else if caches[key] = true; !f {
			result = append(result, v)
		}

	}
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Each(fn func(NullableMultipleDeviceModelEntity)) *NullableMultipleDeviceModelEntityStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) EachRight(fn func(NullableMultipleDeviceModelEntity)) *NullableMultipleDeviceModelEntityStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Equals(arr []NullableMultipleDeviceModelEntity) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if !reflect.DeepEqual((*self)[i], arr[i]) {
			return false
		}
	}
	return true
}
func (self *NullableMultipleDeviceModelEntityStream) Filter(fn func(NullableMultipleDeviceModelEntity, int) bool) *NullableMultipleDeviceModelEntityStream {
	result := NullableMultipleDeviceModelEntityStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) FilterSlim(fn func(NullableMultipleDeviceModelEntity, int) bool) *NullableMultipleDeviceModelEntityStream {
	result := NullableMultipleDeviceModelEntityStreamOf()
	caches := map[string]bool{}
	for i, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v, i); caches[key] {
			result.Add(v)

		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Find(fn func(NullableMultipleDeviceModelEntity, int) bool) *NullableMultipleDeviceModelEntity {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDeviceModelEntityStream) FindOr(fn func(NullableMultipleDeviceModelEntity, int) bool, or NullableMultipleDeviceModelEntity) NullableMultipleDeviceModelEntity {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableMultipleDeviceModelEntityStream) FindIndex(fn func(NullableMultipleDeviceModelEntity, int) bool) int {
	if self == nil {
		return -1
	}
	for i, v := range *self {
		if fn(v, i) {
			return i
		}
	}
	return -1
}
func (self *NullableMultipleDeviceModelEntityStream) First() *NullableMultipleDeviceModelEntity {
	return self.Get(0)
}
func (self *NullableMultipleDeviceModelEntityStream) FirstOr(arg NullableMultipleDeviceModelEntity) NullableMultipleDeviceModelEntity {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDeviceModelEntityStream) ForEach(fn func(NullableMultipleDeviceModelEntity, int)) *NullableMultipleDeviceModelEntityStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) ForEachRight(fn func(NullableMultipleDeviceModelEntity, int)) *NullableMultipleDeviceModelEntityStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) GroupBy(fn func(NullableMultipleDeviceModelEntity, int) string) map[string][]NullableMultipleDeviceModelEntity {
	m := map[string][]NullableMultipleDeviceModelEntity{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableMultipleDeviceModelEntityStream) GroupByValues(fn func(NullableMultipleDeviceModelEntity, int) string) [][]NullableMultipleDeviceModelEntity {
	var tmp [][]NullableMultipleDeviceModelEntity
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableMultipleDeviceModelEntityStream) IndexOf(arg NullableMultipleDeviceModelEntity) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableMultipleDeviceModelEntityStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableMultipleDeviceModelEntityStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableMultipleDeviceModelEntityStream) Last() *NullableMultipleDeviceModelEntity {
	return self.Get(self.Len() - 1)
}
func (self *NullableMultipleDeviceModelEntityStream) LastOr(arg NullableMultipleDeviceModelEntity) NullableMultipleDeviceModelEntity {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDeviceModelEntityStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableMultipleDeviceModelEntityStream) Limit(limit int) *NullableMultipleDeviceModelEntityStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableMultipleDeviceModelEntityStream) Map(fn func(NullableMultipleDeviceModelEntity, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Map2Int(fn func(NullableMultipleDeviceModelEntity, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Map2Int32(fn func(NullableMultipleDeviceModelEntity, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Map2Int64(fn func(NullableMultipleDeviceModelEntity, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Map2Float32(fn func(NullableMultipleDeviceModelEntity, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Map2Float64(fn func(NullableMultipleDeviceModelEntity, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Map2Bool(fn func(NullableMultipleDeviceModelEntity, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Map2Bytes(fn func(NullableMultipleDeviceModelEntity, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Map2String(fn func(NullableMultipleDeviceModelEntity, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Max(fn func(NullableMultipleDeviceModelEntity, int) float64) *NullableMultipleDeviceModelEntity {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Max(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *NullableMultipleDeviceModelEntityStream) Min(fn func(NullableMultipleDeviceModelEntity, int) float64) *NullableMultipleDeviceModelEntity {
	f := self.Get(0)
	if f == nil {
		return nil
	}
	m := fn(*f, 0)
	index := 0
	for i := 1; i < self.Len(); i++ {
		v := fn(*self.Get(i), i)
		m = math.Min(m, v)
		if m == v {
			index = i
		}
	}
	return self.Get(index)
}
func (self *NullableMultipleDeviceModelEntityStream) NoneMatch(fn func(NullableMultipleDeviceModelEntity, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableMultipleDeviceModelEntityStream) Get(index int) *NullableMultipleDeviceModelEntity {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableMultipleDeviceModelEntityStream) GetOr(index int, arg NullableMultipleDeviceModelEntity) NullableMultipleDeviceModelEntity {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableMultipleDeviceModelEntityStream) Peek(fn func(*NullableMultipleDeviceModelEntity, int)) *NullableMultipleDeviceModelEntityStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableMultipleDeviceModelEntityStream) Reduce(fn func(NullableMultipleDeviceModelEntity, NullableMultipleDeviceModelEntity, int) NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	return self.ReduceInit(fn, NullableMultipleDeviceModelEntity{})
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceInit(fn func(NullableMultipleDeviceModelEntity, NullableMultipleDeviceModelEntity, int) NullableMultipleDeviceModelEntity, initialValue NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	result := NullableMultipleDeviceModelEntityStreamOf()
	self.ForEach(func(v NullableMultipleDeviceModelEntity, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceInterface(fn func(interface{}, NullableMultipleDeviceModelEntity, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableMultipleDeviceModelEntity{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceString(fn func(string, NullableMultipleDeviceModelEntity, int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceInt(fn func(int, NullableMultipleDeviceModelEntity, int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceInt32(fn func(int32, NullableMultipleDeviceModelEntity, int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceInt64(fn func(int64, NullableMultipleDeviceModelEntity, int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceFloat32(fn func(float32, NullableMultipleDeviceModelEntity, int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceFloat64(fn func(float64, NullableMultipleDeviceModelEntity, int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelEntityStream) ReduceBool(fn func(bool, NullableMultipleDeviceModelEntity, int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableMultipleDeviceModelEntityStream) Reverse() *NullableMultipleDeviceModelEntityStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Replace(fn func(NullableMultipleDeviceModelEntity, int) NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	return self.ForEach(func(v NullableMultipleDeviceModelEntity, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableMultipleDeviceModelEntityStream) Select(fn func(NullableMultipleDeviceModelEntity) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableMultipleDeviceModelEntityStream) Set(index int, val NullableMultipleDeviceModelEntity) *NullableMultipleDeviceModelEntityStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Skip(skip int) *NullableMultipleDeviceModelEntityStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableMultipleDeviceModelEntityStream) SkippingEach(fn func(NullableMultipleDeviceModelEntity, int) int) *NullableMultipleDeviceModelEntityStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Slice(startIndex, n int) *NullableMultipleDeviceModelEntityStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableMultipleDeviceModelEntity{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Sort(fn func(i, j int) bool) *NullableMultipleDeviceModelEntityStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableMultipleDeviceModelEntityStream) Tail() *NullableMultipleDeviceModelEntity {
	return self.Last()
}
func (self *NullableMultipleDeviceModelEntityStream) TailOr(arg NullableMultipleDeviceModelEntity) NullableMultipleDeviceModelEntity {
	return self.LastOr(arg)
}
func (self *NullableMultipleDeviceModelEntityStream) ToList() []NullableMultipleDeviceModelEntity {
	return self.Val()
}
func (self *NullableMultipleDeviceModelEntityStream) Unique() *NullableMultipleDeviceModelEntityStream {
	return self.Distinct()
}
func (self *NullableMultipleDeviceModelEntityStream) Val() []NullableMultipleDeviceModelEntity {
	if self == nil {
		return []NullableMultipleDeviceModelEntity{}
	}
	return *self.Copy()
}
func (self *NullableMultipleDeviceModelEntityStream) While(fn func(NullableMultipleDeviceModelEntity, int) bool) *NullableMultipleDeviceModelEntityStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) Where(fn func(NullableMultipleDeviceModelEntity) bool) *NullableMultipleDeviceModelEntityStream {
	result := NullableMultipleDeviceModelEntityStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableMultipleDeviceModelEntityStream) WhereSlim(fn func(NullableMultipleDeviceModelEntity) bool) *NullableMultipleDeviceModelEntityStream {
	result := NullableMultipleDeviceModelEntityStreamOf()
	caches := map[string]bool{}
	for _, v := range *self {
		key := fmt.Sprintf("%+v", v)
		if f, ok := caches[key]; ok {
			if f {
				result.Add(v)
			}
		} else if caches[key] = fn(v); caches[key] {
			result.Add(v)
		}
	}
	*self = result
	return self
}
