/*
 * Collection utility of NullableDeviceModelsEntityModel Struct
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

type NullableDeviceModelsEntityModelStream []NullableDeviceModelsEntityModel

func NullableDeviceModelsEntityModelStreamOf(arg ...NullableDeviceModelsEntityModel) NullableDeviceModelsEntityModelStream {
	return arg
}
func NullableDeviceModelsEntityModelStreamFrom(arg []NullableDeviceModelsEntityModel) NullableDeviceModelsEntityModelStream {
	return arg
}
func CreateNullableDeviceModelsEntityModelStream(arg ...NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	tmp := NullableDeviceModelsEntityModelStreamOf(arg...)
	return &tmp
}
func GenerateNullableDeviceModelsEntityModelStream(arg []NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	tmp := NullableDeviceModelsEntityModelStreamFrom(arg)
	return &tmp
}

func (self *NullableDeviceModelsEntityModelStream) Add(arg NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	return self.AddAll(arg)
}
func (self *NullableDeviceModelsEntityModelStream) AddAll(arg ...NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	*self = append(*self, arg...)
	return self
}
func (self *NullableDeviceModelsEntityModelStream) AddSafe(arg *NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Aggregate(fn func(NullableDeviceModelsEntityModel, NullableDeviceModelsEntityModel) NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	result := NullableDeviceModelsEntityModelStreamOf()
	self.ForEach(func(v NullableDeviceModelsEntityModel, i int) {
		if i == 0 {
			result.Add(fn(NullableDeviceModelsEntityModel{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *NullableDeviceModelsEntityModelStream) AllMatch(fn func(NullableDeviceModelsEntityModel, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *NullableDeviceModelsEntityModelStream) AnyMatch(fn func(NullableDeviceModelsEntityModel, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *NullableDeviceModelsEntityModelStream) Clone() *NullableDeviceModelsEntityModelStream {
	temp := make([]NullableDeviceModelsEntityModel, self.Len())
	copy(temp, *self)
	return (*NullableDeviceModelsEntityModelStream)(&temp)
}
func (self *NullableDeviceModelsEntityModelStream) Copy() *NullableDeviceModelsEntityModelStream {
	return self.Clone()
}
func (self *NullableDeviceModelsEntityModelStream) Concat(arg []NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	return self.AddAll(arg...)
}
func (self *NullableDeviceModelsEntityModelStream) Contains(arg NullableDeviceModelsEntityModel) bool {
	return self.FindIndex(func(_arg NullableDeviceModelsEntityModel, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *NullableDeviceModelsEntityModelStream) Clean() *NullableDeviceModelsEntityModelStream {
	*self = NullableDeviceModelsEntityModelStreamOf()
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Delete(index int) *NullableDeviceModelsEntityModelStream {
	return self.DeleteRange(index, index)
}
func (self *NullableDeviceModelsEntityModelStream) DeleteRange(startIndex, endIndex int) *NullableDeviceModelsEntityModelStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Distinct() *NullableDeviceModelsEntityModelStream {
	caches := map[string]bool{}
	result := NullableDeviceModelsEntityModelStreamOf()
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
func (self *NullableDeviceModelsEntityModelStream) Each(fn func(NullableDeviceModelsEntityModel)) *NullableDeviceModelsEntityModelStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) EachRight(fn func(NullableDeviceModelsEntityModel)) *NullableDeviceModelsEntityModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Equals(arr []NullableDeviceModelsEntityModel) bool {
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
func (self *NullableDeviceModelsEntityModelStream) Filter(fn func(NullableDeviceModelsEntityModel, int) bool) *NullableDeviceModelsEntityModelStream {
	result := NullableDeviceModelsEntityModelStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDeviceModelsEntityModelStream) FilterSlim(fn func(NullableDeviceModelsEntityModel, int) bool) *NullableDeviceModelsEntityModelStream {
	result := NullableDeviceModelsEntityModelStreamOf()
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
func (self *NullableDeviceModelsEntityModelStream) Find(fn func(NullableDeviceModelsEntityModel, int) bool) *NullableDeviceModelsEntityModel {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *NullableDeviceModelsEntityModelStream) FindOr(fn func(NullableDeviceModelsEntityModel, int) bool, or NullableDeviceModelsEntityModel) NullableDeviceModelsEntityModel {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *NullableDeviceModelsEntityModelStream) FindIndex(fn func(NullableDeviceModelsEntityModel, int) bool) int {
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
func (self *NullableDeviceModelsEntityModelStream) First() *NullableDeviceModelsEntityModel {
	return self.Get(0)
}
func (self *NullableDeviceModelsEntityModelStream) FirstOr(arg NullableDeviceModelsEntityModel) NullableDeviceModelsEntityModel {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceModelsEntityModelStream) ForEach(fn func(NullableDeviceModelsEntityModel, int)) *NullableDeviceModelsEntityModelStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) ForEachRight(fn func(NullableDeviceModelsEntityModel, int)) *NullableDeviceModelsEntityModelStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) GroupBy(fn func(NullableDeviceModelsEntityModel, int) string) map[string][]NullableDeviceModelsEntityModel {
	m := map[string][]NullableDeviceModelsEntityModel{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *NullableDeviceModelsEntityModelStream) GroupByValues(fn func(NullableDeviceModelsEntityModel, int) string) [][]NullableDeviceModelsEntityModel {
	var tmp [][]NullableDeviceModelsEntityModel
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *NullableDeviceModelsEntityModelStream) IndexOf(arg NullableDeviceModelsEntityModel) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *NullableDeviceModelsEntityModelStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *NullableDeviceModelsEntityModelStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *NullableDeviceModelsEntityModelStream) Last() *NullableDeviceModelsEntityModel {
	return self.Get(self.Len() - 1)
}
func (self *NullableDeviceModelsEntityModelStream) LastOr(arg NullableDeviceModelsEntityModel) NullableDeviceModelsEntityModel {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceModelsEntityModelStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *NullableDeviceModelsEntityModelStream) Limit(limit int) *NullableDeviceModelsEntityModelStream {
	self.Slice(0, limit)
	return self
}

func (self *NullableDeviceModelsEntityModelStream) Map(fn func(NullableDeviceModelsEntityModel, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Map2Int(fn func(NullableDeviceModelsEntityModel, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Map2Int32(fn func(NullableDeviceModelsEntityModel, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Map2Int64(fn func(NullableDeviceModelsEntityModel, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Map2Float32(fn func(NullableDeviceModelsEntityModel, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Map2Float64(fn func(NullableDeviceModelsEntityModel, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Map2Bool(fn func(NullableDeviceModelsEntityModel, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Map2Bytes(fn func(NullableDeviceModelsEntityModel, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Map2String(fn func(NullableDeviceModelsEntityModel, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Max(fn func(NullableDeviceModelsEntityModel, int) float64) *NullableDeviceModelsEntityModel {
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
func (self *NullableDeviceModelsEntityModelStream) Min(fn func(NullableDeviceModelsEntityModel, int) float64) *NullableDeviceModelsEntityModel {
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
func (self *NullableDeviceModelsEntityModelStream) NoneMatch(fn func(NullableDeviceModelsEntityModel, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *NullableDeviceModelsEntityModelStream) Get(index int) *NullableDeviceModelsEntityModel {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *NullableDeviceModelsEntityModelStream) GetOr(index int, arg NullableDeviceModelsEntityModel) NullableDeviceModelsEntityModel {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *NullableDeviceModelsEntityModelStream) Peek(fn func(*NullableDeviceModelsEntityModel, int)) *NullableDeviceModelsEntityModelStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *NullableDeviceModelsEntityModelStream) Reduce(fn func(NullableDeviceModelsEntityModel, NullableDeviceModelsEntityModel, int) NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	return self.ReduceInit(fn, NullableDeviceModelsEntityModel{})
}
func (self *NullableDeviceModelsEntityModelStream) ReduceInit(fn func(NullableDeviceModelsEntityModel, NullableDeviceModelsEntityModel, int) NullableDeviceModelsEntityModel, initialValue NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	result := NullableDeviceModelsEntityModelStreamOf()
	self.ForEach(func(v NullableDeviceModelsEntityModel, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *NullableDeviceModelsEntityModelStream) ReduceInterface(fn func(interface{}, NullableDeviceModelsEntityModel, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(NullableDeviceModelsEntityModel{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *NullableDeviceModelsEntityModelStream) ReduceString(fn func(string, NullableDeviceModelsEntityModel, int) string) []string {
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
func (self *NullableDeviceModelsEntityModelStream) ReduceInt(fn func(int, NullableDeviceModelsEntityModel, int) int) []int {
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
func (self *NullableDeviceModelsEntityModelStream) ReduceInt32(fn func(int32, NullableDeviceModelsEntityModel, int) int32) []int32 {
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
func (self *NullableDeviceModelsEntityModelStream) ReduceInt64(fn func(int64, NullableDeviceModelsEntityModel, int) int64) []int64 {
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
func (self *NullableDeviceModelsEntityModelStream) ReduceFloat32(fn func(float32, NullableDeviceModelsEntityModel, int) float32) []float32 {
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
func (self *NullableDeviceModelsEntityModelStream) ReduceFloat64(fn func(float64, NullableDeviceModelsEntityModel, int) float64) []float64 {
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
func (self *NullableDeviceModelsEntityModelStream) ReduceBool(fn func(bool, NullableDeviceModelsEntityModel, int) bool) []bool {
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
func (self *NullableDeviceModelsEntityModelStream) Reverse() *NullableDeviceModelsEntityModelStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Replace(fn func(NullableDeviceModelsEntityModel, int) NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	return self.ForEach(func(v NullableDeviceModelsEntityModel, i int) { self.Set(i, fn(v, i)) })
}
func (self *NullableDeviceModelsEntityModelStream) Select(fn func(NullableDeviceModelsEntityModel) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *NullableDeviceModelsEntityModelStream) Set(index int, val NullableDeviceModelsEntityModel) *NullableDeviceModelsEntityModelStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Skip(skip int) *NullableDeviceModelsEntityModelStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *NullableDeviceModelsEntityModelStream) SkippingEach(fn func(NullableDeviceModelsEntityModel, int) int) *NullableDeviceModelsEntityModelStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Slice(startIndex, n int) *NullableDeviceModelsEntityModelStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []NullableDeviceModelsEntityModel{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Sort(fn func(i, j int) bool) *NullableDeviceModelsEntityModelStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *NullableDeviceModelsEntityModelStream) Tail() *NullableDeviceModelsEntityModel {
	return self.Last()
}
func (self *NullableDeviceModelsEntityModelStream) TailOr(arg NullableDeviceModelsEntityModel) NullableDeviceModelsEntityModel {
	return self.LastOr(arg)
}
func (self *NullableDeviceModelsEntityModelStream) ToList() []NullableDeviceModelsEntityModel {
	return self.Val()
}
func (self *NullableDeviceModelsEntityModelStream) Unique() *NullableDeviceModelsEntityModelStream {
	return self.Distinct()
}
func (self *NullableDeviceModelsEntityModelStream) Val() []NullableDeviceModelsEntityModel {
	if self == nil {
		return []NullableDeviceModelsEntityModel{}
	}
	return *self.Copy()
}
func (self *NullableDeviceModelsEntityModelStream) While(fn func(NullableDeviceModelsEntityModel, int) bool) *NullableDeviceModelsEntityModelStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *NullableDeviceModelsEntityModelStream) Where(fn func(NullableDeviceModelsEntityModel) bool) *NullableDeviceModelsEntityModelStream {
	result := NullableDeviceModelsEntityModelStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *NullableDeviceModelsEntityModelStream) WhereSlim(fn func(NullableDeviceModelsEntityModel) bool) *NullableDeviceModelsEntityModelStream {
	result := NullableDeviceModelsEntityModelStreamOf()
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
