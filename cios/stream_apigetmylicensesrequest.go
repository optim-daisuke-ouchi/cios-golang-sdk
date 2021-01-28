/*
 * Collection utility of ApiGetMyLicensesRequest Struct
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

type ApiGetMyLicensesRequestStream []ApiGetMyLicensesRequest

func ApiGetMyLicensesRequestStreamOf(arg ...ApiGetMyLicensesRequest) ApiGetMyLicensesRequestStream {
	return arg
}
func ApiGetMyLicensesRequestStreamFrom(arg []ApiGetMyLicensesRequest) ApiGetMyLicensesRequestStream {
	return arg
}
func CreateApiGetMyLicensesRequestStream(arg ...ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	tmp := ApiGetMyLicensesRequestStreamOf(arg...)
	return &tmp
}
func GenerateApiGetMyLicensesRequestStream(arg []ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	tmp := ApiGetMyLicensesRequestStreamFrom(arg)
	return &tmp
}

func (self *ApiGetMyLicensesRequestStream) Add(arg ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	return self.AddAll(arg)
}
func (self *ApiGetMyLicensesRequestStream) AddAll(arg ...ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	*self = append(*self, arg...)
	return self
}
func (self *ApiGetMyLicensesRequestStream) AddSafe(arg *ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) Aggregate(fn func(ApiGetMyLicensesRequest, ApiGetMyLicensesRequest) ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	result := ApiGetMyLicensesRequestStreamOf()
	self.ForEach(func(v ApiGetMyLicensesRequest, i int) {
		if i == 0 {
			result.Add(fn(ApiGetMyLicensesRequest{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ApiGetMyLicensesRequestStream) AllMatch(fn func(ApiGetMyLicensesRequest, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ApiGetMyLicensesRequestStream) AnyMatch(fn func(ApiGetMyLicensesRequest, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ApiGetMyLicensesRequestStream) Clone() *ApiGetMyLicensesRequestStream {
	temp := make([]ApiGetMyLicensesRequest, self.Len())
	copy(temp, *self)
	return (*ApiGetMyLicensesRequestStream)(&temp)
}
func (self *ApiGetMyLicensesRequestStream) Copy() *ApiGetMyLicensesRequestStream {
	return self.Clone()
}
func (self *ApiGetMyLicensesRequestStream) Concat(arg []ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	return self.AddAll(arg...)
}
func (self *ApiGetMyLicensesRequestStream) Contains(arg ApiGetMyLicensesRequest) bool {
	return self.FindIndex(func(_arg ApiGetMyLicensesRequest, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ApiGetMyLicensesRequestStream) Clean() *ApiGetMyLicensesRequestStream {
	*self = ApiGetMyLicensesRequestStreamOf()
	return self
}
func (self *ApiGetMyLicensesRequestStream) Delete(index int) *ApiGetMyLicensesRequestStream {
	return self.DeleteRange(index, index)
}
func (self *ApiGetMyLicensesRequestStream) DeleteRange(startIndex, endIndex int) *ApiGetMyLicensesRequestStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ApiGetMyLicensesRequestStream) Distinct() *ApiGetMyLicensesRequestStream {
	caches := map[string]bool{}
	result := ApiGetMyLicensesRequestStreamOf()
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
func (self *ApiGetMyLicensesRequestStream) Each(fn func(ApiGetMyLicensesRequest)) *ApiGetMyLicensesRequestStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) EachRight(fn func(ApiGetMyLicensesRequest)) *ApiGetMyLicensesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) Equals(arr []ApiGetMyLicensesRequest) bool {
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
func (self *ApiGetMyLicensesRequestStream) Filter(fn func(ApiGetMyLicensesRequest, int) bool) *ApiGetMyLicensesRequestStream {
	result := ApiGetMyLicensesRequestStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetMyLicensesRequestStream) FilterSlim(fn func(ApiGetMyLicensesRequest, int) bool) *ApiGetMyLicensesRequestStream {
	result := ApiGetMyLicensesRequestStreamOf()
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
func (self *ApiGetMyLicensesRequestStream) Find(fn func(ApiGetMyLicensesRequest, int) bool) *ApiGetMyLicensesRequest {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ApiGetMyLicensesRequestStream) FindOr(fn func(ApiGetMyLicensesRequest, int) bool, or ApiGetMyLicensesRequest) ApiGetMyLicensesRequest {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ApiGetMyLicensesRequestStream) FindIndex(fn func(ApiGetMyLicensesRequest, int) bool) int {
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
func (self *ApiGetMyLicensesRequestStream) First() *ApiGetMyLicensesRequest {
	return self.Get(0)
}
func (self *ApiGetMyLicensesRequestStream) FirstOr(arg ApiGetMyLicensesRequest) ApiGetMyLicensesRequest {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetMyLicensesRequestStream) ForEach(fn func(ApiGetMyLicensesRequest, int)) *ApiGetMyLicensesRequestStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) ForEachRight(fn func(ApiGetMyLicensesRequest, int)) *ApiGetMyLicensesRequestStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) GroupBy(fn func(ApiGetMyLicensesRequest, int) string) map[string][]ApiGetMyLicensesRequest {
	m := map[string][]ApiGetMyLicensesRequest{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ApiGetMyLicensesRequestStream) GroupByValues(fn func(ApiGetMyLicensesRequest, int) string) [][]ApiGetMyLicensesRequest {
	var tmp [][]ApiGetMyLicensesRequest
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ApiGetMyLicensesRequestStream) IndexOf(arg ApiGetMyLicensesRequest) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ApiGetMyLicensesRequestStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ApiGetMyLicensesRequestStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ApiGetMyLicensesRequestStream) Last() *ApiGetMyLicensesRequest {
	return self.Get(self.Len() - 1)
}
func (self *ApiGetMyLicensesRequestStream) LastOr(arg ApiGetMyLicensesRequest) ApiGetMyLicensesRequest {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetMyLicensesRequestStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ApiGetMyLicensesRequestStream) Limit(limit int) *ApiGetMyLicensesRequestStream {
	self.Slice(0, limit)
	return self
}

func (self *ApiGetMyLicensesRequestStream) Map(fn func(ApiGetMyLicensesRequest, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Map2Int(fn func(ApiGetMyLicensesRequest, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Map2Int32(fn func(ApiGetMyLicensesRequest, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Map2Int64(fn func(ApiGetMyLicensesRequest, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Map2Float32(fn func(ApiGetMyLicensesRequest, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Map2Float64(fn func(ApiGetMyLicensesRequest, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Map2Bool(fn func(ApiGetMyLicensesRequest, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Map2Bytes(fn func(ApiGetMyLicensesRequest, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Map2String(fn func(ApiGetMyLicensesRequest, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Max(fn func(ApiGetMyLicensesRequest, int) float64) *ApiGetMyLicensesRequest {
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
func (self *ApiGetMyLicensesRequestStream) Min(fn func(ApiGetMyLicensesRequest, int) float64) *ApiGetMyLicensesRequest {
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
func (self *ApiGetMyLicensesRequestStream) NoneMatch(fn func(ApiGetMyLicensesRequest, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ApiGetMyLicensesRequestStream) Get(index int) *ApiGetMyLicensesRequest {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ApiGetMyLicensesRequestStream) GetOr(index int, arg ApiGetMyLicensesRequest) ApiGetMyLicensesRequest {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ApiGetMyLicensesRequestStream) Peek(fn func(*ApiGetMyLicensesRequest, int)) *ApiGetMyLicensesRequestStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ApiGetMyLicensesRequestStream) Reduce(fn func(ApiGetMyLicensesRequest, ApiGetMyLicensesRequest, int) ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	return self.ReduceInit(fn, ApiGetMyLicensesRequest{})
}
func (self *ApiGetMyLicensesRequestStream) ReduceInit(fn func(ApiGetMyLicensesRequest, ApiGetMyLicensesRequest, int) ApiGetMyLicensesRequest, initialValue ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	result := ApiGetMyLicensesRequestStreamOf()
	self.ForEach(func(v ApiGetMyLicensesRequest, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ApiGetMyLicensesRequestStream) ReduceInterface(fn func(interface{}, ApiGetMyLicensesRequest, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ApiGetMyLicensesRequest{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ApiGetMyLicensesRequestStream) ReduceString(fn func(string, ApiGetMyLicensesRequest, int) string) []string {
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
func (self *ApiGetMyLicensesRequestStream) ReduceInt(fn func(int, ApiGetMyLicensesRequest, int) int) []int {
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
func (self *ApiGetMyLicensesRequestStream) ReduceInt32(fn func(int32, ApiGetMyLicensesRequest, int) int32) []int32 {
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
func (self *ApiGetMyLicensesRequestStream) ReduceInt64(fn func(int64, ApiGetMyLicensesRequest, int) int64) []int64 {
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
func (self *ApiGetMyLicensesRequestStream) ReduceFloat32(fn func(float32, ApiGetMyLicensesRequest, int) float32) []float32 {
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
func (self *ApiGetMyLicensesRequestStream) ReduceFloat64(fn func(float64, ApiGetMyLicensesRequest, int) float64) []float64 {
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
func (self *ApiGetMyLicensesRequestStream) ReduceBool(fn func(bool, ApiGetMyLicensesRequest, int) bool) []bool {
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
func (self *ApiGetMyLicensesRequestStream) Reverse() *ApiGetMyLicensesRequestStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) Replace(fn func(ApiGetMyLicensesRequest, int) ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	return self.ForEach(func(v ApiGetMyLicensesRequest, i int) { self.Set(i, fn(v, i)) })
}
func (self *ApiGetMyLicensesRequestStream) Select(fn func(ApiGetMyLicensesRequest) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ApiGetMyLicensesRequestStream) Set(index int, val ApiGetMyLicensesRequest) *ApiGetMyLicensesRequestStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) Skip(skip int) *ApiGetMyLicensesRequestStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ApiGetMyLicensesRequestStream) SkippingEach(fn func(ApiGetMyLicensesRequest, int) int) *ApiGetMyLicensesRequestStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) Slice(startIndex, n int) *ApiGetMyLicensesRequestStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ApiGetMyLicensesRequest{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) Sort(fn func(i, j int) bool) *ApiGetMyLicensesRequestStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ApiGetMyLicensesRequestStream) Tail() *ApiGetMyLicensesRequest {
	return self.Last()
}
func (self *ApiGetMyLicensesRequestStream) TailOr(arg ApiGetMyLicensesRequest) ApiGetMyLicensesRequest {
	return self.LastOr(arg)
}
func (self *ApiGetMyLicensesRequestStream) ToList() []ApiGetMyLicensesRequest {
	return self.Val()
}
func (self *ApiGetMyLicensesRequestStream) Unique() *ApiGetMyLicensesRequestStream {
	return self.Distinct()
}
func (self *ApiGetMyLicensesRequestStream) Val() []ApiGetMyLicensesRequest {
	if self == nil {
		return []ApiGetMyLicensesRequest{}
	}
	return *self.Copy()
}
func (self *ApiGetMyLicensesRequestStream) While(fn func(ApiGetMyLicensesRequest, int) bool) *ApiGetMyLicensesRequestStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ApiGetMyLicensesRequestStream) Where(fn func(ApiGetMyLicensesRequest) bool) *ApiGetMyLicensesRequestStream {
	result := ApiGetMyLicensesRequestStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ApiGetMyLicensesRequestStream) WhereSlim(fn func(ApiGetMyLicensesRequest) bool) *ApiGetMyLicensesRequestStream {
	result := ApiGetMyLicensesRequestStreamOf()
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
