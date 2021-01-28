/*
 * Collection utility of ApiRenameNodeRequest Struct
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

type ApiRenameNodeRequestStream []ApiRenameNodeRequest

func ApiRenameNodeRequestStreamOf(arg ...ApiRenameNodeRequest) ApiRenameNodeRequestStream {
	return arg
}
func ApiRenameNodeRequestStreamFrom(arg []ApiRenameNodeRequest) ApiRenameNodeRequestStream {
	return arg
}
func CreateApiRenameNodeRequestStream(arg ...ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	tmp := ApiRenameNodeRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiRenameNodeRequestStream(arg []ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	tmp := ApiRenameNodeRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiRenameNodeRequestStream) Add(arg ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	return self.AddAll(arg)
}
func (self *ApiRenameNodeRequestStream) AddAll(arg ...ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiRenameNodeRequestStream) AddSafe(arg *ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiRenameNodeRequestStream) Aggregate(fn func(ApiRenameNodeRequest, ApiRenameNodeRequest) ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	result := ApiRenameNodeRequestStreamOf()
	self.ForEach(func(v ApiRenameNodeRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiRenameNodeRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiRenameNodeRequestStream) AllMatch(fn func(ApiRenameNodeRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiRenameNodeRequestStream) AnyMatch(fn func(ApiRenameNodeRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiRenameNodeRequestStream) Clone() *ApiRenameNodeRequestStream {
	temp := make([]ApiRenameNodeRequest, self.Len())
	copy(temp, *self)
	return (*ApiRenameNodeRequestStream)(&temp)
}
func (self *ApiRenameNodeRequestStream) Copy() *ApiRenameNodeRequestStream {
	return self.Clone()
}
func (self *ApiRenameNodeRequestStream) Concat(arg []ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiRenameNodeRequestStream) Contains(arg ApiRenameNodeRequest) bool {
	return self.FindIndex(func(_arg ApiRenameNodeRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiRenameNodeRequestStream) Clean() *ApiRenameNodeRequestStream {
	*self = ApiRenameNodeRequestStreamOf()
	return self
}
func (self *ApiRenameNodeRequestStream) Delete(index int) *ApiRenameNodeRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiRenameNodeRequestStream) DeleteRange(startIndex, endIndex int) *ApiRenameNodeRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiRenameNodeRequestStream) Distinct() *ApiRenameNodeRequestStream {
	caches := map[string]bool{}
	result := ApiRenameNodeRequestStreamOf()
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
func (self *ApiRenameNodeRequestStream) Each(fn func(ApiRenameNodeRequest)) *ApiRenameNodeRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiRenameNodeRequestStream) EachRight(fn func(ApiRenameNodeRequest)) *ApiRenameNodeRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiRenameNodeRequestStream) Equals(arr []ApiRenameNodeRequest) bool {
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
func (self *ApiRenameNodeRequestStream) Filter(fn func(ApiRenameNodeRequest, int) bool) *ApiRenameNodeRequestStream {
	result := ApiRenameNodeRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiRenameNodeRequestStream) FilterSlim(fn func(ApiRenameNodeRequest, int) bool) *ApiRenameNodeRequestStream {
	result := ApiRenameNodeRequestStreamOf()
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
func (self *ApiRenameNodeRequestStream) Find(fn func(ApiRenameNodeRequest, int) bool) *ApiRenameNodeRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiRenameNodeRequestStream) FindOr(fn func(ApiRenameNodeRequest, int) bool, or ApiRenameNodeRequest) ApiRenameNodeRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiRenameNodeRequestStream) FindIndex(fn func(ApiRenameNodeRequest, int) bool) int {
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
func (self *ApiRenameNodeRequestStream) First() *ApiRenameNodeRequest {
	return self.Get(0)
}
func (self *ApiRenameNodeRequestStream) FirstOr(arg ApiRenameNodeRequest) ApiRenameNodeRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiRenameNodeRequestStream) ForEach(fn func(ApiRenameNodeRequest, int)) *ApiRenameNodeRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiRenameNodeRequestStream) ForEachRight(fn func(ApiRenameNodeRequest, int)) *ApiRenameNodeRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiRenameNodeRequestStream) GroupBy(fn func(ApiRenameNodeRequest, int) string) map[string][]ApiRenameNodeRequest {
	m := map[string][]ApiRenameNodeRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiRenameNodeRequestStream) GroupByValues(fn func(ApiRenameNodeRequest, int) string) [][]ApiRenameNodeRequest {
	var tmp [][]ApiRenameNodeRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiRenameNodeRequestStream) IndexOf(arg ApiRenameNodeRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiRenameNodeRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiRenameNodeRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiRenameNodeRequestStream) Last() *ApiRenameNodeRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiRenameNodeRequestStream) LastOr(arg ApiRenameNodeRequest) ApiRenameNodeRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiRenameNodeRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiRenameNodeRequestStream) Limit(limit int) *ApiRenameNodeRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiRenameNodeRequestStream) Map(fn func(ApiRenameNodeRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Map2Int(fn func(ApiRenameNodeRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Map2Int32(fn func(ApiRenameNodeRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Map2Int64(fn func(ApiRenameNodeRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Map2Float32(fn func(ApiRenameNodeRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Map2Float64(fn func(ApiRenameNodeRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Map2Bool(fn func(ApiRenameNodeRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Map2Bytes(fn func(ApiRenameNodeRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Map2String(fn func(ApiRenameNodeRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Max(fn func(ApiRenameNodeRequest, int) float64) *ApiRenameNodeRequest {
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
func (self *ApiRenameNodeRequestStream) Min(fn func(ApiRenameNodeRequest, int) float64) *ApiRenameNodeRequest {
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
func (self *ApiRenameNodeRequestStream) NoneMatch(fn func(ApiRenameNodeRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiRenameNodeRequestStream) Get(index int) *ApiRenameNodeRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiRenameNodeRequestStream) GetOr(index int, arg ApiRenameNodeRequest) ApiRenameNodeRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiRenameNodeRequestStream) Peek(fn func(*ApiRenameNodeRequest, int)) *ApiRenameNodeRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiRenameNodeRequestStream) Reduce(fn func(ApiRenameNodeRequest, ApiRenameNodeRequest, int) ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	return self.ReduceInit(fn, ApiRenameNodeRequest{})
}
func (self *ApiRenameNodeRequestStream) ReduceInit(fn func(ApiRenameNodeRequest, ApiRenameNodeRequest, int) ApiRenameNodeRequest, initialValue ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	result := ApiRenameNodeRequestStreamOf()
	self.ForEach(func(v ApiRenameNodeRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiRenameNodeRequestStream) ReduceInterface(fn func(interface{}, ApiRenameNodeRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiRenameNodeRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiRenameNodeRequestStream) ReduceString(fn func(string, ApiRenameNodeRequest, int) string) []string {
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
func (self *ApiRenameNodeRequestStream) ReduceInt(fn func(int, ApiRenameNodeRequest, int) int) []int {
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
func (self *ApiRenameNodeRequestStream) ReduceInt32(fn func(int32, ApiRenameNodeRequest, int) int32) []int32 {
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
func (self *ApiRenameNodeRequestStream) ReduceInt64(fn func(int64, ApiRenameNodeRequest, int) int64) []int64 {
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
func (self *ApiRenameNodeRequestStream) ReduceFloat32(fn func(float32, ApiRenameNodeRequest, int) float32) []float32 {
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
func (self *ApiRenameNodeRequestStream) ReduceFloat64(fn func(float64, ApiRenameNodeRequest, int) float64) []float64 {
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
func (self *ApiRenameNodeRequestStream) ReduceBool(fn func(bool, ApiRenameNodeRequest, int) bool) []bool {
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
func (self *ApiRenameNodeRequestStream) Reverse() *ApiRenameNodeRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiRenameNodeRequestStream) Replace(fn func(ApiRenameNodeRequest, int) ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	return self.ForEach(func(v ApiRenameNodeRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiRenameNodeRequestStream) Select(fn func(ApiRenameNodeRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiRenameNodeRequestStream) Set(index int, val ApiRenameNodeRequest) *ApiRenameNodeRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiRenameNodeRequestStream) Skip(skip int) *ApiRenameNodeRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiRenameNodeRequestStream) SkippingEach(fn func(ApiRenameNodeRequest, int) int) *ApiRenameNodeRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiRenameNodeRequestStream) Slice(startIndex, n int) *ApiRenameNodeRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiRenameNodeRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiRenameNodeRequestStream) Sort(fn func(i, j int) bool) *ApiRenameNodeRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiRenameNodeRequestStream) Tail() *ApiRenameNodeRequest {
	return self.Last()
}
func (self *ApiRenameNodeRequestStream) TailOr(arg ApiRenameNodeRequest) ApiRenameNodeRequest {
	return self.LastOr(arg)
}
func (self *ApiRenameNodeRequestStream) ToList() []ApiRenameNodeRequest {
	return self.Val()
}
func (self *ApiRenameNodeRequestStream) Unique() *ApiRenameNodeRequestStream {
	return self.Distinct()
}
func (self *ApiRenameNodeRequestStream) Val() []ApiRenameNodeRequest {
	if self == nil {
		return []ApiRenameNodeRequest{}
	}
	return *self.Copy()
}
func (self *ApiRenameNodeRequestStream) While(fn func(ApiRenameNodeRequest, int) bool) *ApiRenameNodeRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiRenameNodeRequestStream) Where(fn func(ApiRenameNodeRequest) bool) *ApiRenameNodeRequestStream {
	result := ApiRenameNodeRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiRenameNodeRequestStream) WhereSlim(fn func(ApiRenameNodeRequest) bool) *ApiRenameNodeRequestStream {
	result := ApiRenameNodeRequestStreamOf()
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
