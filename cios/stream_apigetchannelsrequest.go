/*
 * Collection utility of ApiGetChannelsRequest Struct
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

type ApiGetChannelsRequestStream []ApiGetChannelsRequest

func ApiGetChannelsRequestStreamOf(arg ...ApiGetChannelsRequest) ApiGetChannelsRequestStream {
	return arg
}
func ApiGetChannelsRequestStreamFrom(arg []ApiGetChannelsRequest) ApiGetChannelsRequestStream {
	return arg
}
func CreateApiGetChannelsRequestStream(arg ...ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	tmp := ApiGetChannelsRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetChannelsRequestStream(arg []ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	tmp := ApiGetChannelsRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetChannelsRequestStream) Add(arg ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetChannelsRequestStream) AddAll(arg ...ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetChannelsRequestStream) AddSafe(arg *ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetChannelsRequestStream) Aggregate(fn func(ApiGetChannelsRequest, ApiGetChannelsRequest) ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	result := ApiGetChannelsRequestStreamOf()
	self.ForEach(func(v ApiGetChannelsRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetChannelsRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetChannelsRequestStream) AllMatch(fn func(ApiGetChannelsRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetChannelsRequestStream) AnyMatch(fn func(ApiGetChannelsRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetChannelsRequestStream) Clone() *ApiGetChannelsRequestStream {
	temp := make([]ApiGetChannelsRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetChannelsRequestStream)(&temp)
}
func (self *ApiGetChannelsRequestStream) Copy() *ApiGetChannelsRequestStream {
	return self.Clone()
}
func (self *ApiGetChannelsRequestStream) Concat(arg []ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetChannelsRequestStream) Contains(arg ApiGetChannelsRequest) bool {
	return self.FindIndex(func(_arg ApiGetChannelsRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetChannelsRequestStream) Clean() *ApiGetChannelsRequestStream {
	*self = ApiGetChannelsRequestStreamOf()
	return self
}
func (self *ApiGetChannelsRequestStream) Delete(index int) *ApiGetChannelsRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetChannelsRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetChannelsRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetChannelsRequestStream) Distinct() *ApiGetChannelsRequestStream {
	caches := map[string]bool{}
	result := ApiGetChannelsRequestStreamOf()
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
func (self *ApiGetChannelsRequestStream) Each(fn func(ApiGetChannelsRequest)) *ApiGetChannelsRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetChannelsRequestStream) EachRight(fn func(ApiGetChannelsRequest)) *ApiGetChannelsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetChannelsRequestStream) Equals(arr []ApiGetChannelsRequest) bool {
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
func (self *ApiGetChannelsRequestStream) Filter(fn func(ApiGetChannelsRequest, int) bool) *ApiGetChannelsRequestStream {
	result := ApiGetChannelsRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetChannelsRequestStream) FilterSlim(fn func(ApiGetChannelsRequest, int) bool) *ApiGetChannelsRequestStream {
	result := ApiGetChannelsRequestStreamOf()
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
func (self *ApiGetChannelsRequestStream) Find(fn func(ApiGetChannelsRequest, int) bool) *ApiGetChannelsRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetChannelsRequestStream) FindOr(fn func(ApiGetChannelsRequest, int) bool, or ApiGetChannelsRequest) ApiGetChannelsRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetChannelsRequestStream) FindIndex(fn func(ApiGetChannelsRequest, int) bool) int {
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
func (self *ApiGetChannelsRequestStream) First() *ApiGetChannelsRequest {
	return self.Get(0)
}
func (self *ApiGetChannelsRequestStream) FirstOr(arg ApiGetChannelsRequest) ApiGetChannelsRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetChannelsRequestStream) ForEach(fn func(ApiGetChannelsRequest, int)) *ApiGetChannelsRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetChannelsRequestStream) ForEachRight(fn func(ApiGetChannelsRequest, int)) *ApiGetChannelsRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetChannelsRequestStream) GroupBy(fn func(ApiGetChannelsRequest, int) string) map[string][]ApiGetChannelsRequest {
	m := map[string][]ApiGetChannelsRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetChannelsRequestStream) GroupByValues(fn func(ApiGetChannelsRequest, int) string) [][]ApiGetChannelsRequest {
	var tmp [][]ApiGetChannelsRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetChannelsRequestStream) IndexOf(arg ApiGetChannelsRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetChannelsRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetChannelsRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetChannelsRequestStream) Last() *ApiGetChannelsRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetChannelsRequestStream) LastOr(arg ApiGetChannelsRequest) ApiGetChannelsRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetChannelsRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetChannelsRequestStream) Limit(limit int) *ApiGetChannelsRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetChannelsRequestStream) Map(fn func(ApiGetChannelsRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Map2Int(fn func(ApiGetChannelsRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Map2Int32(fn func(ApiGetChannelsRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Map2Int64(fn func(ApiGetChannelsRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Map2Float32(fn func(ApiGetChannelsRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Map2Float64(fn func(ApiGetChannelsRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Map2Bool(fn func(ApiGetChannelsRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Map2Bytes(fn func(ApiGetChannelsRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Map2String(fn func(ApiGetChannelsRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Max(fn func(ApiGetChannelsRequest, int) float64) *ApiGetChannelsRequest {
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
func (self *ApiGetChannelsRequestStream) Min(fn func(ApiGetChannelsRequest, int) float64) *ApiGetChannelsRequest {
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
func (self *ApiGetChannelsRequestStream) NoneMatch(fn func(ApiGetChannelsRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetChannelsRequestStream) Get(index int) *ApiGetChannelsRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetChannelsRequestStream) GetOr(index int, arg ApiGetChannelsRequest) ApiGetChannelsRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetChannelsRequestStream) Peek(fn func(*ApiGetChannelsRequest, int)) *ApiGetChannelsRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetChannelsRequestStream) Reduce(fn func(ApiGetChannelsRequest, ApiGetChannelsRequest, int) ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	return self.ReduceInit(fn, ApiGetChannelsRequest{})
}
func (self *ApiGetChannelsRequestStream) ReduceInit(fn func(ApiGetChannelsRequest, ApiGetChannelsRequest, int) ApiGetChannelsRequest, initialValue ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	result := ApiGetChannelsRequestStreamOf()
	self.ForEach(func(v ApiGetChannelsRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetChannelsRequestStream) ReduceInterface(fn func(interface{}, ApiGetChannelsRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetChannelsRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetChannelsRequestStream) ReduceString(fn func(string, ApiGetChannelsRequest, int) string) []string {
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
func (self *ApiGetChannelsRequestStream) ReduceInt(fn func(int, ApiGetChannelsRequest, int) int) []int {
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
func (self *ApiGetChannelsRequestStream) ReduceInt32(fn func(int32, ApiGetChannelsRequest, int) int32) []int32 {
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
func (self *ApiGetChannelsRequestStream) ReduceInt64(fn func(int64, ApiGetChannelsRequest, int) int64) []int64 {
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
func (self *ApiGetChannelsRequestStream) ReduceFloat32(fn func(float32, ApiGetChannelsRequest, int) float32) []float32 {
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
func (self *ApiGetChannelsRequestStream) ReduceFloat64(fn func(float64, ApiGetChannelsRequest, int) float64) []float64 {
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
func (self *ApiGetChannelsRequestStream) ReduceBool(fn func(bool, ApiGetChannelsRequest, int) bool) []bool {
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
func (self *ApiGetChannelsRequestStream) Reverse() *ApiGetChannelsRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetChannelsRequestStream) Replace(fn func(ApiGetChannelsRequest, int) ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	return self.ForEach(func(v ApiGetChannelsRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetChannelsRequestStream) Select(fn func(ApiGetChannelsRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetChannelsRequestStream) Set(index int, val ApiGetChannelsRequest) *ApiGetChannelsRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetChannelsRequestStream) Skip(skip int) *ApiGetChannelsRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetChannelsRequestStream) SkippingEach(fn func(ApiGetChannelsRequest, int) int) *ApiGetChannelsRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetChannelsRequestStream) Slice(startIndex, n int) *ApiGetChannelsRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetChannelsRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetChannelsRequestStream) Sort(fn func(i, j int) bool) *ApiGetChannelsRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetChannelsRequestStream) Tail() *ApiGetChannelsRequest {
	return self.Last()
}
func (self *ApiGetChannelsRequestStream) TailOr(arg ApiGetChannelsRequest) ApiGetChannelsRequest {
	return self.LastOr(arg)
}
func (self *ApiGetChannelsRequestStream) ToList() []ApiGetChannelsRequest {
	return self.Val()
}
func (self *ApiGetChannelsRequestStream) Unique() *ApiGetChannelsRequestStream {
	return self.Distinct()
}
func (self *ApiGetChannelsRequestStream) Val() []ApiGetChannelsRequest {
	if self == nil {
		return []ApiGetChannelsRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetChannelsRequestStream) While(fn func(ApiGetChannelsRequest, int) bool) *ApiGetChannelsRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetChannelsRequestStream) Where(fn func(ApiGetChannelsRequest) bool) *ApiGetChannelsRequestStream {
	result := ApiGetChannelsRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetChannelsRequestStream) WhereSlim(fn func(ApiGetChannelsRequest) bool) *ApiGetChannelsRequestStream {
	result := ApiGetChannelsRequestStreamOf()
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
