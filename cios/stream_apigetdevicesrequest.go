/*
 * Collection utility of ApiGetDevicesRequest Struct
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

type ApiGetDevicesRequestStream []ApiGetDevicesRequest

func ApiGetDevicesRequestStreamOf(arg ...ApiGetDevicesRequest) ApiGetDevicesRequestStream {
	return arg
}
func ApiGetDevicesRequestStreamFrom(arg []ApiGetDevicesRequest) ApiGetDevicesRequestStream {
	return arg
}
func CreateApiGetDevicesRequestStream(arg ...ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	tmp := ApiGetDevicesRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDevicesRequestStream(arg []ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	tmp := ApiGetDevicesRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDevicesRequestStream) Add(arg ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDevicesRequestStream) AddAll(arg ...ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDevicesRequestStream) AddSafe(arg *ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDevicesRequestStream) Aggregate(fn func(ApiGetDevicesRequest, ApiGetDevicesRequest) ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	result := ApiGetDevicesRequestStreamOf()
	self.ForEach(func(v ApiGetDevicesRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDevicesRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDevicesRequestStream) AllMatch(fn func(ApiGetDevicesRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDevicesRequestStream) AnyMatch(fn func(ApiGetDevicesRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDevicesRequestStream) Clone() *ApiGetDevicesRequestStream {
	temp := make([]ApiGetDevicesRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDevicesRequestStream)(&temp)
}
func (self *ApiGetDevicesRequestStream) Copy() *ApiGetDevicesRequestStream {
	return self.Clone()
}
func (self *ApiGetDevicesRequestStream) Concat(arg []ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDevicesRequestStream) Contains(arg ApiGetDevicesRequest) bool {
	return self.FindIndex(func(_arg ApiGetDevicesRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDevicesRequestStream) Clean() *ApiGetDevicesRequestStream {
	*self = ApiGetDevicesRequestStreamOf()
	return self
}
func (self *ApiGetDevicesRequestStream) Delete(index int) *ApiGetDevicesRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDevicesRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDevicesRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDevicesRequestStream) Distinct() *ApiGetDevicesRequestStream {
	caches := map[string]bool{}
	result := ApiGetDevicesRequestStreamOf()
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
func (self *ApiGetDevicesRequestStream) Each(fn func(ApiGetDevicesRequest)) *ApiGetDevicesRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDevicesRequestStream) EachRight(fn func(ApiGetDevicesRequest)) *ApiGetDevicesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDevicesRequestStream) Equals(arr []ApiGetDevicesRequest) bool {
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
func (self *ApiGetDevicesRequestStream) Filter(fn func(ApiGetDevicesRequest, int) bool) *ApiGetDevicesRequestStream {
	result := ApiGetDevicesRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDevicesRequestStream) FilterSlim(fn func(ApiGetDevicesRequest, int) bool) *ApiGetDevicesRequestStream {
	result := ApiGetDevicesRequestStreamOf()
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
func (self *ApiGetDevicesRequestStream) Find(fn func(ApiGetDevicesRequest, int) bool) *ApiGetDevicesRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDevicesRequestStream) FindOr(fn func(ApiGetDevicesRequest, int) bool, or ApiGetDevicesRequest) ApiGetDevicesRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDevicesRequestStream) FindIndex(fn func(ApiGetDevicesRequest, int) bool) int {
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
func (self *ApiGetDevicesRequestStream) First() *ApiGetDevicesRequest {
	return self.Get(0)
}
func (self *ApiGetDevicesRequestStream) FirstOr(arg ApiGetDevicesRequest) ApiGetDevicesRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDevicesRequestStream) ForEach(fn func(ApiGetDevicesRequest, int)) *ApiGetDevicesRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDevicesRequestStream) ForEachRight(fn func(ApiGetDevicesRequest, int)) *ApiGetDevicesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDevicesRequestStream) GroupBy(fn func(ApiGetDevicesRequest, int) string) map[string][]ApiGetDevicesRequest {
	m := map[string][]ApiGetDevicesRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDevicesRequestStream) GroupByValues(fn func(ApiGetDevicesRequest, int) string) [][]ApiGetDevicesRequest {
	var tmp [][]ApiGetDevicesRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDevicesRequestStream) IndexOf(arg ApiGetDevicesRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDevicesRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDevicesRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDevicesRequestStream) Last() *ApiGetDevicesRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDevicesRequestStream) LastOr(arg ApiGetDevicesRequest) ApiGetDevicesRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDevicesRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDevicesRequestStream) Limit(limit int) *ApiGetDevicesRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDevicesRequestStream) Map(fn func(ApiGetDevicesRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Map2Int(fn func(ApiGetDevicesRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Map2Int32(fn func(ApiGetDevicesRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Map2Int64(fn func(ApiGetDevicesRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Map2Float32(fn func(ApiGetDevicesRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Map2Float64(fn func(ApiGetDevicesRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Map2Bool(fn func(ApiGetDevicesRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Map2Bytes(fn func(ApiGetDevicesRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Map2String(fn func(ApiGetDevicesRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Max(fn func(ApiGetDevicesRequest, int) float64) *ApiGetDevicesRequest {
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
func (self *ApiGetDevicesRequestStream) Min(fn func(ApiGetDevicesRequest, int) float64) *ApiGetDevicesRequest {
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
func (self *ApiGetDevicesRequestStream) NoneMatch(fn func(ApiGetDevicesRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDevicesRequestStream) Get(index int) *ApiGetDevicesRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDevicesRequestStream) GetOr(index int, arg ApiGetDevicesRequest) ApiGetDevicesRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDevicesRequestStream) Peek(fn func(*ApiGetDevicesRequest, int)) *ApiGetDevicesRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDevicesRequestStream) Reduce(fn func(ApiGetDevicesRequest, ApiGetDevicesRequest, int) ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	return self.ReduceInit(fn, ApiGetDevicesRequest{})
}
func (self *ApiGetDevicesRequestStream) ReduceInit(fn func(ApiGetDevicesRequest, ApiGetDevicesRequest, int) ApiGetDevicesRequest, initialValue ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	result := ApiGetDevicesRequestStreamOf()
	self.ForEach(func(v ApiGetDevicesRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDevicesRequestStream) ReduceInterface(fn func(interface{}, ApiGetDevicesRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDevicesRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDevicesRequestStream) ReduceString(fn func(string, ApiGetDevicesRequest, int) string) []string {
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
func (self *ApiGetDevicesRequestStream) ReduceInt(fn func(int, ApiGetDevicesRequest, int) int) []int {
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
func (self *ApiGetDevicesRequestStream) ReduceInt32(fn func(int32, ApiGetDevicesRequest, int) int32) []int32 {
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
func (self *ApiGetDevicesRequestStream) ReduceInt64(fn func(int64, ApiGetDevicesRequest, int) int64) []int64 {
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
func (self *ApiGetDevicesRequestStream) ReduceFloat32(fn func(float32, ApiGetDevicesRequest, int) float32) []float32 {
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
func (self *ApiGetDevicesRequestStream) ReduceFloat64(fn func(float64, ApiGetDevicesRequest, int) float64) []float64 {
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
func (self *ApiGetDevicesRequestStream) ReduceBool(fn func(bool, ApiGetDevicesRequest, int) bool) []bool {
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
func (self *ApiGetDevicesRequestStream) Reverse() *ApiGetDevicesRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDevicesRequestStream) Replace(fn func(ApiGetDevicesRequest, int) ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	return self.ForEach(func(v ApiGetDevicesRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDevicesRequestStream) Select(fn func(ApiGetDevicesRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDevicesRequestStream) Set(index int, val ApiGetDevicesRequest) *ApiGetDevicesRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDevicesRequestStream) Skip(skip int) *ApiGetDevicesRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDevicesRequestStream) SkippingEach(fn func(ApiGetDevicesRequest, int) int) *ApiGetDevicesRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDevicesRequestStream) Slice(startIndex, n int) *ApiGetDevicesRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDevicesRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDevicesRequestStream) Sort(fn func(i, j int) bool) *ApiGetDevicesRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDevicesRequestStream) Tail() *ApiGetDevicesRequest {
	return self.Last()
}
func (self *ApiGetDevicesRequestStream) TailOr(arg ApiGetDevicesRequest) ApiGetDevicesRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDevicesRequestStream) ToList() []ApiGetDevicesRequest {
	return self.Val()
}
func (self *ApiGetDevicesRequestStream) Unique() *ApiGetDevicesRequestStream {
	return self.Distinct()
}
func (self *ApiGetDevicesRequestStream) Val() []ApiGetDevicesRequest {
	if self == nil {
		return []ApiGetDevicesRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDevicesRequestStream) While(fn func(ApiGetDevicesRequest, int) bool) *ApiGetDevicesRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDevicesRequestStream) Where(fn func(ApiGetDevicesRequest) bool) *ApiGetDevicesRequestStream {
	result := ApiGetDevicesRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDevicesRequestStream) WhereSlim(fn func(ApiGetDevicesRequest) bool) *ApiGetDevicesRequestStream {
	result := ApiGetDevicesRequestStreamOf()
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
