/*
 * Collection utility of ChannelUpdateProposal Struct
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

type ChannelUpdateProposalStream []ChannelUpdateProposal

func ChannelUpdateProposalStreamOf(arg ...ChannelUpdateProposal) ChannelUpdateProposalStream {
	return arg
}
func ChannelUpdateProposalStreamFrom(arg []ChannelUpdateProposal) ChannelUpdateProposalStream {
	return arg
}
func CreateChannelUpdateProposalStream(arg ...ChannelUpdateProposal) *ChannelUpdateProposalStream {
	tmp := ChannelUpdateProposalStreamOf(arg...)
	return &tmp
}
func GenerateChannelUpdateProposalStream(arg []ChannelUpdateProposal) *ChannelUpdateProposalStream {
	tmp := ChannelUpdateProposalStreamFrom(arg)
	return &tmp
}

func (self *ChannelUpdateProposalStream) Add(arg ChannelUpdateProposal) *ChannelUpdateProposalStream {
	return self.AddAll(arg)
}
func (self *ChannelUpdateProposalStream) AddAll(arg ...ChannelUpdateProposal) *ChannelUpdateProposalStream {
	*self = append(*self, arg...)
	return self
}
func (self *ChannelUpdateProposalStream) AddSafe(arg *ChannelUpdateProposal) *ChannelUpdateProposalStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *ChannelUpdateProposalStream) Aggregate(fn func(ChannelUpdateProposal, ChannelUpdateProposal) ChannelUpdateProposal) *ChannelUpdateProposalStream {
	result := ChannelUpdateProposalStreamOf()
	self.ForEach(func(v ChannelUpdateProposal, i int) {
		if i == 0 {
			result.Add(fn(ChannelUpdateProposal{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *ChannelUpdateProposalStream) AllMatch(fn func(ChannelUpdateProposal, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *ChannelUpdateProposalStream) AnyMatch(fn func(ChannelUpdateProposal, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *ChannelUpdateProposalStream) Clone() *ChannelUpdateProposalStream {
	temp := make([]ChannelUpdateProposal, self.Len())
	copy(temp, *self)
	return (*ChannelUpdateProposalStream)(&temp)
}
func (self *ChannelUpdateProposalStream) Copy() *ChannelUpdateProposalStream {
	return self.Clone()
}
func (self *ChannelUpdateProposalStream) Concat(arg []ChannelUpdateProposal) *ChannelUpdateProposalStream {
	return self.AddAll(arg...)
}
func (self *ChannelUpdateProposalStream) Contains(arg ChannelUpdateProposal) bool {
	return self.FindIndex(func(_arg ChannelUpdateProposal, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *ChannelUpdateProposalStream) Clean() *ChannelUpdateProposalStream {
	*self = ChannelUpdateProposalStreamOf()
	return self
}
func (self *ChannelUpdateProposalStream) Delete(index int) *ChannelUpdateProposalStream {
	return self.DeleteRange(index, index)
}
func (self *ChannelUpdateProposalStream) DeleteRange(startIndex, endIndex int) *ChannelUpdateProposalStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *ChannelUpdateProposalStream) Distinct() *ChannelUpdateProposalStream {
	caches := map[string]bool{}
	result := ChannelUpdateProposalStreamOf()
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
func (self *ChannelUpdateProposalStream) Each(fn func(ChannelUpdateProposal)) *ChannelUpdateProposalStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *ChannelUpdateProposalStream) EachRight(fn func(ChannelUpdateProposal)) *ChannelUpdateProposalStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *ChannelUpdateProposalStream) Equals(arr []ChannelUpdateProposal) bool {
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
func (self *ChannelUpdateProposalStream) Filter(fn func(ChannelUpdateProposal, int) bool) *ChannelUpdateProposalStream {
	result := ChannelUpdateProposalStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ChannelUpdateProposalStream) FilterSlim(fn func(ChannelUpdateProposal, int) bool) *ChannelUpdateProposalStream {
	result := ChannelUpdateProposalStreamOf()
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
func (self *ChannelUpdateProposalStream) Find(fn func(ChannelUpdateProposal, int) bool) *ChannelUpdateProposal {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *ChannelUpdateProposalStream) FindOr(fn func(ChannelUpdateProposal, int) bool, or ChannelUpdateProposal) ChannelUpdateProposal {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *ChannelUpdateProposalStream) FindIndex(fn func(ChannelUpdateProposal, int) bool) int {
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
func (self *ChannelUpdateProposalStream) First() *ChannelUpdateProposal {
	return self.Get(0)
}
func (self *ChannelUpdateProposalStream) FirstOr(arg ChannelUpdateProposal) ChannelUpdateProposal {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *ChannelUpdateProposalStream) ForEach(fn func(ChannelUpdateProposal, int)) *ChannelUpdateProposalStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *ChannelUpdateProposalStream) ForEachRight(fn func(ChannelUpdateProposal, int)) *ChannelUpdateProposalStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *ChannelUpdateProposalStream) GroupBy(fn func(ChannelUpdateProposal, int) string) map[string][]ChannelUpdateProposal {
	m := map[string][]ChannelUpdateProposal{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *ChannelUpdateProposalStream) GroupByValues(fn func(ChannelUpdateProposal, int) string) [][]ChannelUpdateProposal {
	var tmp [][]ChannelUpdateProposal
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *ChannelUpdateProposalStream) IndexOf(arg ChannelUpdateProposal) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *ChannelUpdateProposalStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *ChannelUpdateProposalStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *ChannelUpdateProposalStream) Last() *ChannelUpdateProposal {
	return self.Get(self.Len() - 1)
}
func (self *ChannelUpdateProposalStream) LastOr(arg ChannelUpdateProposal) ChannelUpdateProposal {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *ChannelUpdateProposalStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *ChannelUpdateProposalStream) Limit(limit int) *ChannelUpdateProposalStream {
	self.Slice(0, limit)
	return self
}

func (self *ChannelUpdateProposalStream) Map(fn func(ChannelUpdateProposal, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Map2Int(fn func(ChannelUpdateProposal, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Map2Int32(fn func(ChannelUpdateProposal, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Map2Int64(fn func(ChannelUpdateProposal, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Map2Float32(fn func(ChannelUpdateProposal, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Map2Float64(fn func(ChannelUpdateProposal, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Map2Bool(fn func(ChannelUpdateProposal, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Map2Bytes(fn func(ChannelUpdateProposal, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Map2String(fn func(ChannelUpdateProposal, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Max(fn func(ChannelUpdateProposal, int) float64) *ChannelUpdateProposal {
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
func (self *ChannelUpdateProposalStream) Min(fn func(ChannelUpdateProposal, int) float64) *ChannelUpdateProposal {
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
func (self *ChannelUpdateProposalStream) NoneMatch(fn func(ChannelUpdateProposal, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *ChannelUpdateProposalStream) Get(index int) *ChannelUpdateProposal {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *ChannelUpdateProposalStream) GetOr(index int, arg ChannelUpdateProposal) ChannelUpdateProposal {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *ChannelUpdateProposalStream) Peek(fn func(*ChannelUpdateProposal, int)) *ChannelUpdateProposalStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *ChannelUpdateProposalStream) Reduce(fn func(ChannelUpdateProposal, ChannelUpdateProposal, int) ChannelUpdateProposal) *ChannelUpdateProposalStream {
	return self.ReduceInit(fn, ChannelUpdateProposal{})
}
func (self *ChannelUpdateProposalStream) ReduceInit(fn func(ChannelUpdateProposal, ChannelUpdateProposal, int) ChannelUpdateProposal, initialValue ChannelUpdateProposal) *ChannelUpdateProposalStream {
	result := ChannelUpdateProposalStreamOf()
	self.ForEach(func(v ChannelUpdateProposal, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *ChannelUpdateProposalStream) ReduceInterface(fn func(interface{}, ChannelUpdateProposal, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(ChannelUpdateProposal{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *ChannelUpdateProposalStream) ReduceString(fn func(string, ChannelUpdateProposal, int) string) []string {
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
func (self *ChannelUpdateProposalStream) ReduceInt(fn func(int, ChannelUpdateProposal, int) int) []int {
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
func (self *ChannelUpdateProposalStream) ReduceInt32(fn func(int32, ChannelUpdateProposal, int) int32) []int32 {
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
func (self *ChannelUpdateProposalStream) ReduceInt64(fn func(int64, ChannelUpdateProposal, int) int64) []int64 {
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
func (self *ChannelUpdateProposalStream) ReduceFloat32(fn func(float32, ChannelUpdateProposal, int) float32) []float32 {
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
func (self *ChannelUpdateProposalStream) ReduceFloat64(fn func(float64, ChannelUpdateProposal, int) float64) []float64 {
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
func (self *ChannelUpdateProposalStream) ReduceBool(fn func(bool, ChannelUpdateProposal, int) bool) []bool {
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
func (self *ChannelUpdateProposalStream) Reverse() *ChannelUpdateProposalStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *ChannelUpdateProposalStream) Replace(fn func(ChannelUpdateProposal, int) ChannelUpdateProposal) *ChannelUpdateProposalStream {
	return self.ForEach(func(v ChannelUpdateProposal, i int) { self.Set(i, fn(v, i)) })
}
func (self *ChannelUpdateProposalStream) Select(fn func(ChannelUpdateProposal) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *ChannelUpdateProposalStream) Set(index int, val ChannelUpdateProposal) *ChannelUpdateProposalStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *ChannelUpdateProposalStream) Skip(skip int) *ChannelUpdateProposalStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *ChannelUpdateProposalStream) SkippingEach(fn func(ChannelUpdateProposal, int) int) *ChannelUpdateProposalStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *ChannelUpdateProposalStream) Slice(startIndex, n int) *ChannelUpdateProposalStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []ChannelUpdateProposal{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *ChannelUpdateProposalStream) Sort(fn func(i, j int) bool) *ChannelUpdateProposalStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *ChannelUpdateProposalStream) Tail() *ChannelUpdateProposal {
	return self.Last()
}
func (self *ChannelUpdateProposalStream) TailOr(arg ChannelUpdateProposal) ChannelUpdateProposal {
	return self.LastOr(arg)
}
func (self *ChannelUpdateProposalStream) ToList() []ChannelUpdateProposal {
	return self.Val()
}
func (self *ChannelUpdateProposalStream) Unique() *ChannelUpdateProposalStream {
	return self.Distinct()
}
func (self *ChannelUpdateProposalStream) Val() []ChannelUpdateProposal {
	if self == nil {
		return []ChannelUpdateProposal{}
	}
	return *self.Copy()
}
func (self *ChannelUpdateProposalStream) While(fn func(ChannelUpdateProposal, int) bool) *ChannelUpdateProposalStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *ChannelUpdateProposalStream) Where(fn func(ChannelUpdateProposal) bool) *ChannelUpdateProposalStream {
	result := ChannelUpdateProposalStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *ChannelUpdateProposalStream) WhereSlim(fn func(ChannelUpdateProposal) bool) *ChannelUpdateProposalStream {
	result := ChannelUpdateProposalStreamOf()
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
