/*
 * Collection utility of ApiGetSeriesThumnbnailRequest Struct
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

type ApiGetSeriesThumnbnailRequestStream []ApiGetSeriesThumnbnailRequest

func ApiGetSeriesThumnbnailRequestStreamOf(arg ...ApiGetSeriesThumnbnailRequest) ApiGetSeriesThumnbnailRequestStream {
	return arg
}
func ApiGetSeriesThumnbnailRequestStreamFrom(arg []ApiGetSeriesThumnbnailRequest) ApiGetSeriesThumnbnailRequestStream {
	return arg
}
func CreateApiGetSeriesThumnbnailRequestStream(arg ...ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	tmp := ApiGetSeriesThumnbnailRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetSeriesThumnbnailRequestStream(arg []ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	tmp := ApiGetSeriesThumnbnailRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetSeriesThumnbnailRequestStream) Add(arg ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetSeriesThumnbnailRequestStream) AddAll(arg ...ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) AddSafe(arg *ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Aggregate(fn func(ApiGetSeriesThumnbnailRequest, ApiGetSeriesThumnbnailRequest) ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	result := ApiGetSeriesThumnbnailRequestStreamOf()
	self.ForEach(func(v ApiGetSeriesThumnbnailRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetSeriesThumnbnailRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) AllMatch(fn func(ApiGetSeriesThumnbnailRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetSeriesThumnbnailRequestStream) AnyMatch(fn func(ApiGetSeriesThumnbnailRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetSeriesThumnbnailRequestStream) Clone() *ApiGetSeriesThumnbnailRequestStream {
	temp := make([]ApiGetSeriesThumnbnailRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetSeriesThumnbnailRequestStream)(&temp)
}
func (self *ApiGetSeriesThumnbnailRequestStream) Copy() *ApiGetSeriesThumnbnailRequestStream {
	return self.Clone()
}
func (self *ApiGetSeriesThumnbnailRequestStream) Concat(arg []ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetSeriesThumnbnailRequestStream) Contains(arg ApiGetSeriesThumnbnailRequest) bool {
	return self.FindIndex(func(_arg ApiGetSeriesThumnbnailRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetSeriesThumnbnailRequestStream) Clean() *ApiGetSeriesThumnbnailRequestStream {
	*self = ApiGetSeriesThumnbnailRequestStreamOf()
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Delete(index int) *ApiGetSeriesThumnbnailRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetSeriesThumnbnailRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetSeriesThumnbnailRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Distinct() *ApiGetSeriesThumnbnailRequestStream {
	caches := map[string]bool{}
	result := ApiGetSeriesThumnbnailRequestStreamOf()
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
func (self *ApiGetSeriesThumnbnailRequestStream) Each(fn func(ApiGetSeriesThumnbnailRequest)) *ApiGetSeriesThumnbnailRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) EachRight(fn func(ApiGetSeriesThumnbnailRequest)) *ApiGetSeriesThumnbnailRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Equals(arr []ApiGetSeriesThumnbnailRequest) bool {
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
func (self *ApiGetSeriesThumnbnailRequestStream) Filter(fn func(ApiGetSeriesThumnbnailRequest, int) bool) *ApiGetSeriesThumnbnailRequestStream {
	result := ApiGetSeriesThumnbnailRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) FilterSlim(fn func(ApiGetSeriesThumnbnailRequest, int) bool) *ApiGetSeriesThumnbnailRequestStream {
	result := ApiGetSeriesThumnbnailRequestStreamOf()
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
func (self *ApiGetSeriesThumnbnailRequestStream) Find(fn func(ApiGetSeriesThumnbnailRequest, int) bool) *ApiGetSeriesThumnbnailRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetSeriesThumnbnailRequestStream) FindOr(fn func(ApiGetSeriesThumnbnailRequest, int) bool, or ApiGetSeriesThumnbnailRequest) ApiGetSeriesThumnbnailRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetSeriesThumnbnailRequestStream) FindIndex(fn func(ApiGetSeriesThumnbnailRequest, int) bool) int {
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
func (self *ApiGetSeriesThumnbnailRequestStream) First() *ApiGetSeriesThumnbnailRequest {
	return self.Get(0)
}
func (self *ApiGetSeriesThumnbnailRequestStream) FirstOr(arg ApiGetSeriesThumnbnailRequest) ApiGetSeriesThumnbnailRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetSeriesThumnbnailRequestStream) ForEach(fn func(ApiGetSeriesThumnbnailRequest, int)) *ApiGetSeriesThumnbnailRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) ForEachRight(fn func(ApiGetSeriesThumnbnailRequest, int)) *ApiGetSeriesThumnbnailRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) GroupBy(fn func(ApiGetSeriesThumnbnailRequest, int) string) map[string][]ApiGetSeriesThumnbnailRequest {
	m := map[string][]ApiGetSeriesThumnbnailRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetSeriesThumnbnailRequestStream) GroupByValues(fn func(ApiGetSeriesThumnbnailRequest, int) string) [][]ApiGetSeriesThumnbnailRequest {
	var tmp [][]ApiGetSeriesThumnbnailRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetSeriesThumnbnailRequestStream) IndexOf(arg ApiGetSeriesThumnbnailRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetSeriesThumnbnailRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetSeriesThumnbnailRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetSeriesThumnbnailRequestStream) Last() *ApiGetSeriesThumnbnailRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetSeriesThumnbnailRequestStream) LastOr(arg ApiGetSeriesThumnbnailRequest) ApiGetSeriesThumnbnailRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetSeriesThumnbnailRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetSeriesThumnbnailRequestStream) Limit(limit int) *ApiGetSeriesThumnbnailRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetSeriesThumnbnailRequestStream) Map(fn func(ApiGetSeriesThumnbnailRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Map2Int(fn func(ApiGetSeriesThumnbnailRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Map2Int32(fn func(ApiGetSeriesThumnbnailRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Map2Int64(fn func(ApiGetSeriesThumnbnailRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Map2Float32(fn func(ApiGetSeriesThumnbnailRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Map2Float64(fn func(ApiGetSeriesThumnbnailRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Map2Bool(fn func(ApiGetSeriesThumnbnailRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Map2Bytes(fn func(ApiGetSeriesThumnbnailRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Map2String(fn func(ApiGetSeriesThumnbnailRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Max(fn func(ApiGetSeriesThumnbnailRequest, int) float64) *ApiGetSeriesThumnbnailRequest {
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
func (self *ApiGetSeriesThumnbnailRequestStream) Min(fn func(ApiGetSeriesThumnbnailRequest, int) float64) *ApiGetSeriesThumnbnailRequest {
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
func (self *ApiGetSeriesThumnbnailRequestStream) NoneMatch(fn func(ApiGetSeriesThumnbnailRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetSeriesThumnbnailRequestStream) Get(index int) *ApiGetSeriesThumnbnailRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetSeriesThumnbnailRequestStream) GetOr(index int, arg ApiGetSeriesThumnbnailRequest) ApiGetSeriesThumnbnailRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetSeriesThumnbnailRequestStream) Peek(fn func(*ApiGetSeriesThumnbnailRequest, int)) *ApiGetSeriesThumnbnailRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetSeriesThumnbnailRequestStream) Reduce(fn func(ApiGetSeriesThumnbnailRequest, ApiGetSeriesThumnbnailRequest, int) ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	return self.ReduceInit(fn, ApiGetSeriesThumnbnailRequest{})
}
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceInit(fn func(ApiGetSeriesThumnbnailRequest, ApiGetSeriesThumnbnailRequest, int) ApiGetSeriesThumnbnailRequest, initialValue ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	result := ApiGetSeriesThumnbnailRequestStreamOf()
	self.ForEach(func(v ApiGetSeriesThumnbnailRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceInterface(fn func(interface{}, ApiGetSeriesThumnbnailRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetSeriesThumnbnailRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceString(fn func(string, ApiGetSeriesThumnbnailRequest, int) string) []string {
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
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceInt(fn func(int, ApiGetSeriesThumnbnailRequest, int) int) []int {
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
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceInt32(fn func(int32, ApiGetSeriesThumnbnailRequest, int) int32) []int32 {
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
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceInt64(fn func(int64, ApiGetSeriesThumnbnailRequest, int) int64) []int64 {
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
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceFloat32(fn func(float32, ApiGetSeriesThumnbnailRequest, int) float32) []float32 {
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
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceFloat64(fn func(float64, ApiGetSeriesThumnbnailRequest, int) float64) []float64 {
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
func (self *ApiGetSeriesThumnbnailRequestStream) ReduceBool(fn func(bool, ApiGetSeriesThumnbnailRequest, int) bool) []bool {
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
func (self *ApiGetSeriesThumnbnailRequestStream) Reverse() *ApiGetSeriesThumnbnailRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Replace(fn func(ApiGetSeriesThumnbnailRequest, int) ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	return self.ForEach(func(v ApiGetSeriesThumnbnailRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetSeriesThumnbnailRequestStream) Select(fn func(ApiGetSeriesThumnbnailRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetSeriesThumnbnailRequestStream) Set(index int, val ApiGetSeriesThumnbnailRequest) *ApiGetSeriesThumnbnailRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Skip(skip int) *ApiGetSeriesThumnbnailRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetSeriesThumnbnailRequestStream) SkippingEach(fn func(ApiGetSeriesThumnbnailRequest, int) int) *ApiGetSeriesThumnbnailRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Slice(startIndex, n int) *ApiGetSeriesThumnbnailRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetSeriesThumnbnailRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Sort(fn func(i, j int) bool) *ApiGetSeriesThumnbnailRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetSeriesThumnbnailRequestStream) Tail() *ApiGetSeriesThumnbnailRequest {
	return self.Last()
}
func (self *ApiGetSeriesThumnbnailRequestStream) TailOr(arg ApiGetSeriesThumnbnailRequest) ApiGetSeriesThumnbnailRequest {
	return self.LastOr(arg)
}
func (self *ApiGetSeriesThumnbnailRequestStream) ToList() []ApiGetSeriesThumnbnailRequest {
	return self.Val()
}
func (self *ApiGetSeriesThumnbnailRequestStream) Unique() *ApiGetSeriesThumnbnailRequestStream {
	return self.Distinct()
}
func (self *ApiGetSeriesThumnbnailRequestStream) Val() []ApiGetSeriesThumnbnailRequest {
	if self == nil {
		return []ApiGetSeriesThumnbnailRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetSeriesThumnbnailRequestStream) While(fn func(ApiGetSeriesThumnbnailRequest, int) bool) *ApiGetSeriesThumnbnailRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) Where(fn func(ApiGetSeriesThumnbnailRequest) bool) *ApiGetSeriesThumnbnailRequestStream {
	result := ApiGetSeriesThumnbnailRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetSeriesThumnbnailRequestStream) WhereSlim(fn func(ApiGetSeriesThumnbnailRequest) bool) *ApiGetSeriesThumnbnailRequestStream {
	result := ApiGetSeriesThumnbnailRequestStreamOf()
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
