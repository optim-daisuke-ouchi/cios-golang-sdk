/*
 * Collection utility of ApiGetDeviceEntitiesRequest Struct
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

type ApiGetDeviceEntitiesRequestStream []ApiGetDeviceEntitiesRequest

func ApiGetDeviceEntitiesRequestStreamOf(arg ...ApiGetDeviceEntitiesRequest) ApiGetDeviceEntitiesRequestStream {
	return arg
}
func ApiGetDeviceEntitiesRequestStreamFrom(arg []ApiGetDeviceEntitiesRequest) ApiGetDeviceEntitiesRequestStream {
	return arg
}
func CreateApiGetDeviceEntitiesRequestStream(arg ...ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	tmp := ApiGetDeviceEntitiesRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDeviceEntitiesRequestStream(arg []ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	tmp := ApiGetDeviceEntitiesRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDeviceEntitiesRequestStream) Add(arg ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDeviceEntitiesRequestStream) AddAll(arg ...ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) AddSafe(arg *ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Aggregate(fn func(ApiGetDeviceEntitiesRequest, ApiGetDeviceEntitiesRequest) ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	result := ApiGetDeviceEntitiesRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceEntitiesRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDeviceEntitiesRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) AllMatch(fn func(ApiGetDeviceEntitiesRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDeviceEntitiesRequestStream) AnyMatch(fn func(ApiGetDeviceEntitiesRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDeviceEntitiesRequestStream) Clone() *ApiGetDeviceEntitiesRequestStream {
	temp := make([]ApiGetDeviceEntitiesRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDeviceEntitiesRequestStream)(&temp)
}
func (self *ApiGetDeviceEntitiesRequestStream) Copy() *ApiGetDeviceEntitiesRequestStream {
	return self.Clone()
}
func (self *ApiGetDeviceEntitiesRequestStream) Concat(arg []ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDeviceEntitiesRequestStream) Contains(arg ApiGetDeviceEntitiesRequest) bool {
	return self.FindIndex(func(_arg ApiGetDeviceEntitiesRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDeviceEntitiesRequestStream) Clean() *ApiGetDeviceEntitiesRequestStream {
	*self = ApiGetDeviceEntitiesRequestStreamOf()
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Delete(index int) *ApiGetDeviceEntitiesRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDeviceEntitiesRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDeviceEntitiesRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Distinct() *ApiGetDeviceEntitiesRequestStream {
	caches := map[string]bool{}
	result := ApiGetDeviceEntitiesRequestStreamOf()
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
func (self *ApiGetDeviceEntitiesRequestStream) Each(fn func(ApiGetDeviceEntitiesRequest)) *ApiGetDeviceEntitiesRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) EachRight(fn func(ApiGetDeviceEntitiesRequest)) *ApiGetDeviceEntitiesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Equals(arr []ApiGetDeviceEntitiesRequest) bool {
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
func (self *ApiGetDeviceEntitiesRequestStream) Filter(fn func(ApiGetDeviceEntitiesRequest, int) bool) *ApiGetDeviceEntitiesRequestStream {
	result := ApiGetDeviceEntitiesRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) FilterSlim(fn func(ApiGetDeviceEntitiesRequest, int) bool) *ApiGetDeviceEntitiesRequestStream {
	result := ApiGetDeviceEntitiesRequestStreamOf()
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
func (self *ApiGetDeviceEntitiesRequestStream) Find(fn func(ApiGetDeviceEntitiesRequest, int) bool) *ApiGetDeviceEntitiesRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceEntitiesRequestStream) FindOr(fn func(ApiGetDeviceEntitiesRequest, int) bool, or ApiGetDeviceEntitiesRequest) ApiGetDeviceEntitiesRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDeviceEntitiesRequestStream) FindIndex(fn func(ApiGetDeviceEntitiesRequest, int) bool) int {
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
func (self *ApiGetDeviceEntitiesRequestStream) First() *ApiGetDeviceEntitiesRequest {
	return self.Get(0)
}
func (self *ApiGetDeviceEntitiesRequestStream) FirstOr(arg ApiGetDeviceEntitiesRequest) ApiGetDeviceEntitiesRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesRequestStream) ForEach(fn func(ApiGetDeviceEntitiesRequest, int)) *ApiGetDeviceEntitiesRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) ForEachRight(fn func(ApiGetDeviceEntitiesRequest, int)) *ApiGetDeviceEntitiesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) GroupBy(fn func(ApiGetDeviceEntitiesRequest, int) string) map[string][]ApiGetDeviceEntitiesRequest {
	m := map[string][]ApiGetDeviceEntitiesRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDeviceEntitiesRequestStream) GroupByValues(fn func(ApiGetDeviceEntitiesRequest, int) string) [][]ApiGetDeviceEntitiesRequest {
	var tmp [][]ApiGetDeviceEntitiesRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDeviceEntitiesRequestStream) IndexOf(arg ApiGetDeviceEntitiesRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDeviceEntitiesRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDeviceEntitiesRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDeviceEntitiesRequestStream) Last() *ApiGetDeviceEntitiesRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDeviceEntitiesRequestStream) LastOr(arg ApiGetDeviceEntitiesRequest) ApiGetDeviceEntitiesRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDeviceEntitiesRequestStream) Limit(limit int) *ApiGetDeviceEntitiesRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDeviceEntitiesRequestStream) Map(fn func(ApiGetDeviceEntitiesRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Map2Int(fn func(ApiGetDeviceEntitiesRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Map2Int32(fn func(ApiGetDeviceEntitiesRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Map2Int64(fn func(ApiGetDeviceEntitiesRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Map2Float32(fn func(ApiGetDeviceEntitiesRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Map2Float64(fn func(ApiGetDeviceEntitiesRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Map2Bool(fn func(ApiGetDeviceEntitiesRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Map2Bytes(fn func(ApiGetDeviceEntitiesRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Map2String(fn func(ApiGetDeviceEntitiesRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Max(fn func(ApiGetDeviceEntitiesRequest, int) float64) *ApiGetDeviceEntitiesRequest {
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
func (self *ApiGetDeviceEntitiesRequestStream) Min(fn func(ApiGetDeviceEntitiesRequest, int) float64) *ApiGetDeviceEntitiesRequest {
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
func (self *ApiGetDeviceEntitiesRequestStream) NoneMatch(fn func(ApiGetDeviceEntitiesRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDeviceEntitiesRequestStream) Get(index int) *ApiGetDeviceEntitiesRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDeviceEntitiesRequestStream) GetOr(index int, arg ApiGetDeviceEntitiesRequest) ApiGetDeviceEntitiesRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDeviceEntitiesRequestStream) Peek(fn func(*ApiGetDeviceEntitiesRequest, int)) *ApiGetDeviceEntitiesRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDeviceEntitiesRequestStream) Reduce(fn func(ApiGetDeviceEntitiesRequest, ApiGetDeviceEntitiesRequest, int) ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	return self.ReduceInit(fn, ApiGetDeviceEntitiesRequest{})
}
func (self *ApiGetDeviceEntitiesRequestStream) ReduceInit(fn func(ApiGetDeviceEntitiesRequest, ApiGetDeviceEntitiesRequest, int) ApiGetDeviceEntitiesRequest, initialValue ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	result := ApiGetDeviceEntitiesRequestStreamOf()
	self.ForEach(func(v ApiGetDeviceEntitiesRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) ReduceInterface(fn func(interface{}, ApiGetDeviceEntitiesRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDeviceEntitiesRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDeviceEntitiesRequestStream) ReduceString(fn func(string, ApiGetDeviceEntitiesRequest, int) string) []string {
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
func (self *ApiGetDeviceEntitiesRequestStream) ReduceInt(fn func(int, ApiGetDeviceEntitiesRequest, int) int) []int {
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
func (self *ApiGetDeviceEntitiesRequestStream) ReduceInt32(fn func(int32, ApiGetDeviceEntitiesRequest, int) int32) []int32 {
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
func (self *ApiGetDeviceEntitiesRequestStream) ReduceInt64(fn func(int64, ApiGetDeviceEntitiesRequest, int) int64) []int64 {
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
func (self *ApiGetDeviceEntitiesRequestStream) ReduceFloat32(fn func(float32, ApiGetDeviceEntitiesRequest, int) float32) []float32 {
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
func (self *ApiGetDeviceEntitiesRequestStream) ReduceFloat64(fn func(float64, ApiGetDeviceEntitiesRequest, int) float64) []float64 {
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
func (self *ApiGetDeviceEntitiesRequestStream) ReduceBool(fn func(bool, ApiGetDeviceEntitiesRequest, int) bool) []bool {
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
func (self *ApiGetDeviceEntitiesRequestStream) Reverse() *ApiGetDeviceEntitiesRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Replace(fn func(ApiGetDeviceEntitiesRequest, int) ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	return self.ForEach(func(v ApiGetDeviceEntitiesRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDeviceEntitiesRequestStream) Select(fn func(ApiGetDeviceEntitiesRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDeviceEntitiesRequestStream) Set(index int, val ApiGetDeviceEntitiesRequest) *ApiGetDeviceEntitiesRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Skip(skip int) *ApiGetDeviceEntitiesRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDeviceEntitiesRequestStream) SkippingEach(fn func(ApiGetDeviceEntitiesRequest, int) int) *ApiGetDeviceEntitiesRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Slice(startIndex, n int) *ApiGetDeviceEntitiesRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDeviceEntitiesRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Sort(fn func(i, j int) bool) *ApiGetDeviceEntitiesRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDeviceEntitiesRequestStream) Tail() *ApiGetDeviceEntitiesRequest {
	return self.Last()
}
func (self *ApiGetDeviceEntitiesRequestStream) TailOr(arg ApiGetDeviceEntitiesRequest) ApiGetDeviceEntitiesRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDeviceEntitiesRequestStream) ToList() []ApiGetDeviceEntitiesRequest {
	return self.Val()
}
func (self *ApiGetDeviceEntitiesRequestStream) Unique() *ApiGetDeviceEntitiesRequestStream {
	return self.Distinct()
}
func (self *ApiGetDeviceEntitiesRequestStream) Val() []ApiGetDeviceEntitiesRequest {
	if self == nil {
		return []ApiGetDeviceEntitiesRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDeviceEntitiesRequestStream) While(fn func(ApiGetDeviceEntitiesRequest, int) bool) *ApiGetDeviceEntitiesRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) Where(fn func(ApiGetDeviceEntitiesRequest) bool) *ApiGetDeviceEntitiesRequestStream {
	result := ApiGetDeviceEntitiesRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDeviceEntitiesRequestStream) WhereSlim(fn func(ApiGetDeviceEntitiesRequest) bool) *ApiGetDeviceEntitiesRequestStream {
	result := ApiGetDeviceEntitiesRequestStreamOf()
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
