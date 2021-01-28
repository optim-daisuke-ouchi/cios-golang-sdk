/*
 * Collection utility of ApiGetDataStoreObjectDataRequest Struct
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

type ApiGetDataStoreObjectDataRequestStream []ApiGetDataStoreObjectDataRequest

func ApiGetDataStoreObjectDataRequestStreamOf(arg ...ApiGetDataStoreObjectDataRequest) ApiGetDataStoreObjectDataRequestStream {
	return arg
}
func ApiGetDataStoreObjectDataRequestStreamFrom(arg []ApiGetDataStoreObjectDataRequest) ApiGetDataStoreObjectDataRequestStream {
	return arg
}
func CreateApiGetDataStoreObjectDataRequestStream(arg ...ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	tmp := ApiGetDataStoreObjectDataRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetDataStoreObjectDataRequestStream(arg []ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	tmp := ApiGetDataStoreObjectDataRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetDataStoreObjectDataRequestStream) Add(arg ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetDataStoreObjectDataRequestStream) AddAll(arg ...ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) AddSafe(arg *ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Aggregate(fn func(ApiGetDataStoreObjectDataRequest, ApiGetDataStoreObjectDataRequest) ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	result := ApiGetDataStoreObjectDataRequestStreamOf()
	self.ForEach(func(v ApiGetDataStoreObjectDataRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetDataStoreObjectDataRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) AllMatch(fn func(ApiGetDataStoreObjectDataRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetDataStoreObjectDataRequestStream) AnyMatch(fn func(ApiGetDataStoreObjectDataRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetDataStoreObjectDataRequestStream) Clone() *ApiGetDataStoreObjectDataRequestStream {
	temp := make([]ApiGetDataStoreObjectDataRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetDataStoreObjectDataRequestStream)(&temp)
}
func (self *ApiGetDataStoreObjectDataRequestStream) Copy() *ApiGetDataStoreObjectDataRequestStream {
	return self.Clone()
}
func (self *ApiGetDataStoreObjectDataRequestStream) Concat(arg []ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetDataStoreObjectDataRequestStream) Contains(arg ApiGetDataStoreObjectDataRequest) bool {
	return self.FindIndex(func(_arg ApiGetDataStoreObjectDataRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetDataStoreObjectDataRequestStream) Clean() *ApiGetDataStoreObjectDataRequestStream {
	*self = ApiGetDataStoreObjectDataRequestStreamOf()
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Delete(index int) *ApiGetDataStoreObjectDataRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetDataStoreObjectDataRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetDataStoreObjectDataRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Distinct() *ApiGetDataStoreObjectDataRequestStream {
	caches := map[string]bool{}
	result := ApiGetDataStoreObjectDataRequestStreamOf()
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
func (self *ApiGetDataStoreObjectDataRequestStream) Each(fn func(ApiGetDataStoreObjectDataRequest)) *ApiGetDataStoreObjectDataRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) EachRight(fn func(ApiGetDataStoreObjectDataRequest)) *ApiGetDataStoreObjectDataRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Equals(arr []ApiGetDataStoreObjectDataRequest) bool {
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
func (self *ApiGetDataStoreObjectDataRequestStream) Filter(fn func(ApiGetDataStoreObjectDataRequest, int) bool) *ApiGetDataStoreObjectDataRequestStream {
	result := ApiGetDataStoreObjectDataRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) FilterSlim(fn func(ApiGetDataStoreObjectDataRequest, int) bool) *ApiGetDataStoreObjectDataRequestStream {
	result := ApiGetDataStoreObjectDataRequestStreamOf()
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
func (self *ApiGetDataStoreObjectDataRequestStream) Find(fn func(ApiGetDataStoreObjectDataRequest, int) bool) *ApiGetDataStoreObjectDataRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetDataStoreObjectDataRequestStream) FindOr(fn func(ApiGetDataStoreObjectDataRequest, int) bool, or ApiGetDataStoreObjectDataRequest) ApiGetDataStoreObjectDataRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetDataStoreObjectDataRequestStream) FindIndex(fn func(ApiGetDataStoreObjectDataRequest, int) bool) int {
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
func (self *ApiGetDataStoreObjectDataRequestStream) First() *ApiGetDataStoreObjectDataRequest {
	return self.Get(0)
}
func (self *ApiGetDataStoreObjectDataRequestStream) FirstOr(arg ApiGetDataStoreObjectDataRequest) ApiGetDataStoreObjectDataRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreObjectDataRequestStream) ForEach(fn func(ApiGetDataStoreObjectDataRequest, int)) *ApiGetDataStoreObjectDataRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) ForEachRight(fn func(ApiGetDataStoreObjectDataRequest, int)) *ApiGetDataStoreObjectDataRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) GroupBy(fn func(ApiGetDataStoreObjectDataRequest, int) string) map[string][]ApiGetDataStoreObjectDataRequest {
	m := map[string][]ApiGetDataStoreObjectDataRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetDataStoreObjectDataRequestStream) GroupByValues(fn func(ApiGetDataStoreObjectDataRequest, int) string) [][]ApiGetDataStoreObjectDataRequest {
	var tmp [][]ApiGetDataStoreObjectDataRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetDataStoreObjectDataRequestStream) IndexOf(arg ApiGetDataStoreObjectDataRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetDataStoreObjectDataRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetDataStoreObjectDataRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetDataStoreObjectDataRequestStream) Last() *ApiGetDataStoreObjectDataRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetDataStoreObjectDataRequestStream) LastOr(arg ApiGetDataStoreObjectDataRequest) ApiGetDataStoreObjectDataRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreObjectDataRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetDataStoreObjectDataRequestStream) Limit(limit int) *ApiGetDataStoreObjectDataRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetDataStoreObjectDataRequestStream) Map(fn func(ApiGetDataStoreObjectDataRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Map2Int(fn func(ApiGetDataStoreObjectDataRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Map2Int32(fn func(ApiGetDataStoreObjectDataRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Map2Int64(fn func(ApiGetDataStoreObjectDataRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Map2Float32(fn func(ApiGetDataStoreObjectDataRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Map2Float64(fn func(ApiGetDataStoreObjectDataRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Map2Bool(fn func(ApiGetDataStoreObjectDataRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Map2Bytes(fn func(ApiGetDataStoreObjectDataRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Map2String(fn func(ApiGetDataStoreObjectDataRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Max(fn func(ApiGetDataStoreObjectDataRequest, int) float64) *ApiGetDataStoreObjectDataRequest {
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
func (self *ApiGetDataStoreObjectDataRequestStream) Min(fn func(ApiGetDataStoreObjectDataRequest, int) float64) *ApiGetDataStoreObjectDataRequest {
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
func (self *ApiGetDataStoreObjectDataRequestStream) NoneMatch(fn func(ApiGetDataStoreObjectDataRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetDataStoreObjectDataRequestStream) Get(index int) *ApiGetDataStoreObjectDataRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetDataStoreObjectDataRequestStream) GetOr(index int, arg ApiGetDataStoreObjectDataRequest) ApiGetDataStoreObjectDataRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetDataStoreObjectDataRequestStream) Peek(fn func(*ApiGetDataStoreObjectDataRequest, int)) *ApiGetDataStoreObjectDataRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetDataStoreObjectDataRequestStream) Reduce(fn func(ApiGetDataStoreObjectDataRequest, ApiGetDataStoreObjectDataRequest, int) ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	return self.ReduceInit(fn, ApiGetDataStoreObjectDataRequest{})
}
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceInit(fn func(ApiGetDataStoreObjectDataRequest, ApiGetDataStoreObjectDataRequest, int) ApiGetDataStoreObjectDataRequest, initialValue ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	result := ApiGetDataStoreObjectDataRequestStreamOf()
	self.ForEach(func(v ApiGetDataStoreObjectDataRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceInterface(fn func(interface{}, ApiGetDataStoreObjectDataRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetDataStoreObjectDataRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceString(fn func(string, ApiGetDataStoreObjectDataRequest, int) string) []string {
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
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceInt(fn func(int, ApiGetDataStoreObjectDataRequest, int) int) []int {
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
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceInt32(fn func(int32, ApiGetDataStoreObjectDataRequest, int) int32) []int32 {
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
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceInt64(fn func(int64, ApiGetDataStoreObjectDataRequest, int) int64) []int64 {
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
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceFloat32(fn func(float32, ApiGetDataStoreObjectDataRequest, int) float32) []float32 {
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
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceFloat64(fn func(float64, ApiGetDataStoreObjectDataRequest, int) float64) []float64 {
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
func (self *ApiGetDataStoreObjectDataRequestStream) ReduceBool(fn func(bool, ApiGetDataStoreObjectDataRequest, int) bool) []bool {
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
func (self *ApiGetDataStoreObjectDataRequestStream) Reverse() *ApiGetDataStoreObjectDataRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Replace(fn func(ApiGetDataStoreObjectDataRequest, int) ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	return self.ForEach(func(v ApiGetDataStoreObjectDataRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetDataStoreObjectDataRequestStream) Select(fn func(ApiGetDataStoreObjectDataRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetDataStoreObjectDataRequestStream) Set(index int, val ApiGetDataStoreObjectDataRequest) *ApiGetDataStoreObjectDataRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Skip(skip int) *ApiGetDataStoreObjectDataRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetDataStoreObjectDataRequestStream) SkippingEach(fn func(ApiGetDataStoreObjectDataRequest, int) int) *ApiGetDataStoreObjectDataRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Slice(startIndex, n int) *ApiGetDataStoreObjectDataRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetDataStoreObjectDataRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Sort(fn func(i, j int) bool) *ApiGetDataStoreObjectDataRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetDataStoreObjectDataRequestStream) Tail() *ApiGetDataStoreObjectDataRequest {
	return self.Last()
}
func (self *ApiGetDataStoreObjectDataRequestStream) TailOr(arg ApiGetDataStoreObjectDataRequest) ApiGetDataStoreObjectDataRequest {
	return self.LastOr(arg)
}
func (self *ApiGetDataStoreObjectDataRequestStream) ToList() []ApiGetDataStoreObjectDataRequest {
	return self.Val()
}
func (self *ApiGetDataStoreObjectDataRequestStream) Unique() *ApiGetDataStoreObjectDataRequestStream {
	return self.Distinct()
}
func (self *ApiGetDataStoreObjectDataRequestStream) Val() []ApiGetDataStoreObjectDataRequest {
	if self == nil {
		return []ApiGetDataStoreObjectDataRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetDataStoreObjectDataRequestStream) While(fn func(ApiGetDataStoreObjectDataRequest, int) bool) *ApiGetDataStoreObjectDataRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) Where(fn func(ApiGetDataStoreObjectDataRequest) bool) *ApiGetDataStoreObjectDataRequestStream {
	result := ApiGetDataStoreObjectDataRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetDataStoreObjectDataRequestStream) WhereSlim(fn func(ApiGetDataStoreObjectDataRequest) bool) *ApiGetDataStoreObjectDataRequestStream {
	result := ApiGetDataStoreObjectDataRequestStreamOf()
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
