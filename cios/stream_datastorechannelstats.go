/*
 * Collection utility of DataStoreChannelStats Struct
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

type DataStoreChannelStatsStream []DataStoreChannelStats

func DataStoreChannelStatsStreamOf(arg ...DataStoreChannelStats) DataStoreChannelStatsStream {
	return arg
}
func DataStoreChannelStatsStreamFrom(arg []DataStoreChannelStats) DataStoreChannelStatsStream {
	return arg
}
func CreateDataStoreChannelStatsStream(arg ...DataStoreChannelStats) *DataStoreChannelStatsStream {
	tmp := DataStoreChannelStatsStreamOf(arg...)
	return &tmp
}
func GenerateDataStoreChannelStatsStream(arg []DataStoreChannelStats) *DataStoreChannelStatsStream {
	tmp := DataStoreChannelStatsStreamFrom(arg)
	return &tmp
}

func (self *DataStoreChannelStatsStream) Add(arg DataStoreChannelStats) *DataStoreChannelStatsStream {
	return self.AddAll(arg)
}
func (self *DataStoreChannelStatsStream) AddAll(arg ...DataStoreChannelStats) *DataStoreChannelStatsStream {
	*self = append(*self, arg...)
	return self
}
func (self *DataStoreChannelStatsStream) AddSafe(arg *DataStoreChannelStats) *DataStoreChannelStatsStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *DataStoreChannelStatsStream) Aggregate(fn func(DataStoreChannelStats, DataStoreChannelStats) DataStoreChannelStats) *DataStoreChannelStatsStream {
	result := DataStoreChannelStatsStreamOf()
	self.ForEach(func(v DataStoreChannelStats, i int) {
		if i == 0 {
			result.Add(fn(DataStoreChannelStats{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *DataStoreChannelStatsStream) AllMatch(fn func(DataStoreChannelStats, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *DataStoreChannelStatsStream) AnyMatch(fn func(DataStoreChannelStats, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *DataStoreChannelStatsStream) Clone() *DataStoreChannelStatsStream {
	temp := make([]DataStoreChannelStats, self.Len())
	copy(temp, *self)
	return (*DataStoreChannelStatsStream)(&temp)
}
func (self *DataStoreChannelStatsStream) Copy() *DataStoreChannelStatsStream {
	return self.Clone()
}
func (self *DataStoreChannelStatsStream) Concat(arg []DataStoreChannelStats) *DataStoreChannelStatsStream {
	return self.AddAll(arg...)
}
func (self *DataStoreChannelStatsStream) Contains(arg DataStoreChannelStats) bool {
	return self.FindIndex(func(_arg DataStoreChannelStats, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *DataStoreChannelStatsStream) Clean() *DataStoreChannelStatsStream {
	*self = DataStoreChannelStatsStreamOf()
	return self
}
func (self *DataStoreChannelStatsStream) Delete(index int) *DataStoreChannelStatsStream {
	return self.DeleteRange(index, index)
}
func (self *DataStoreChannelStatsStream) DeleteRange(startIndex, endIndex int) *DataStoreChannelStatsStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *DataStoreChannelStatsStream) Distinct() *DataStoreChannelStatsStream {
	caches := map[string]bool{}
	result := DataStoreChannelStatsStreamOf()
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
func (self *DataStoreChannelStatsStream) Each(fn func(DataStoreChannelStats)) *DataStoreChannelStatsStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *DataStoreChannelStatsStream) EachRight(fn func(DataStoreChannelStats)) *DataStoreChannelStatsStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *DataStoreChannelStatsStream) Equals(arr []DataStoreChannelStats) bool {
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
func (self *DataStoreChannelStatsStream) Filter(fn func(DataStoreChannelStats, int) bool) *DataStoreChannelStatsStream {
	result := DataStoreChannelStatsStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DataStoreChannelStatsStream) FilterSlim(fn func(DataStoreChannelStats, int) bool) *DataStoreChannelStatsStream {
	result := DataStoreChannelStatsStreamOf()
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
func (self *DataStoreChannelStatsStream) Find(fn func(DataStoreChannelStats, int) bool) *DataStoreChannelStats {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *DataStoreChannelStatsStream) FindOr(fn func(DataStoreChannelStats, int) bool, or DataStoreChannelStats) DataStoreChannelStats {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *DataStoreChannelStatsStream) FindIndex(fn func(DataStoreChannelStats, int) bool) int {
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
func (self *DataStoreChannelStatsStream) First() *DataStoreChannelStats {
	return self.Get(0)
}
func (self *DataStoreChannelStatsStream) FirstOr(arg DataStoreChannelStats) DataStoreChannelStats {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *DataStoreChannelStatsStream) ForEach(fn func(DataStoreChannelStats, int)) *DataStoreChannelStatsStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *DataStoreChannelStatsStream) ForEachRight(fn func(DataStoreChannelStats, int)) *DataStoreChannelStatsStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *DataStoreChannelStatsStream) GroupBy(fn func(DataStoreChannelStats, int) string) map[string][]DataStoreChannelStats {
	m := map[string][]DataStoreChannelStats{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *DataStoreChannelStatsStream) GroupByValues(fn func(DataStoreChannelStats, int) string) [][]DataStoreChannelStats {
	var tmp [][]DataStoreChannelStats
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *DataStoreChannelStatsStream) IndexOf(arg DataStoreChannelStats) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *DataStoreChannelStatsStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *DataStoreChannelStatsStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *DataStoreChannelStatsStream) Last() *DataStoreChannelStats {
	return self.Get(self.Len() - 1)
}
func (self *DataStoreChannelStatsStream) LastOr(arg DataStoreChannelStats) DataStoreChannelStats {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *DataStoreChannelStatsStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *DataStoreChannelStatsStream) Limit(limit int) *DataStoreChannelStatsStream {
	self.Slice(0, limit)
	return self
}

func (self *DataStoreChannelStatsStream) Map(fn func(DataStoreChannelStats, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Map2Int(fn func(DataStoreChannelStats, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Map2Int32(fn func(DataStoreChannelStats, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Map2Int64(fn func(DataStoreChannelStats, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Map2Float32(fn func(DataStoreChannelStats, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Map2Float64(fn func(DataStoreChannelStats, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Map2Bool(fn func(DataStoreChannelStats, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Map2Bytes(fn func(DataStoreChannelStats, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Map2String(fn func(DataStoreChannelStats, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Max(fn func(DataStoreChannelStats, int) float64) *DataStoreChannelStats {
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
func (self *DataStoreChannelStatsStream) Min(fn func(DataStoreChannelStats, int) float64) *DataStoreChannelStats {
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
func (self *DataStoreChannelStatsStream) NoneMatch(fn func(DataStoreChannelStats, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *DataStoreChannelStatsStream) Get(index int) *DataStoreChannelStats {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *DataStoreChannelStatsStream) GetOr(index int, arg DataStoreChannelStats) DataStoreChannelStats {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *DataStoreChannelStatsStream) Peek(fn func(*DataStoreChannelStats, int)) *DataStoreChannelStatsStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *DataStoreChannelStatsStream) Reduce(fn func(DataStoreChannelStats, DataStoreChannelStats, int) DataStoreChannelStats) *DataStoreChannelStatsStream {
	return self.ReduceInit(fn, DataStoreChannelStats{})
}
func (self *DataStoreChannelStatsStream) ReduceInit(fn func(DataStoreChannelStats, DataStoreChannelStats, int) DataStoreChannelStats, initialValue DataStoreChannelStats) *DataStoreChannelStatsStream {
	result := DataStoreChannelStatsStreamOf()
	self.ForEach(func(v DataStoreChannelStats, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *DataStoreChannelStatsStream) ReduceInterface(fn func(interface{}, DataStoreChannelStats, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(DataStoreChannelStats{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *DataStoreChannelStatsStream) ReduceString(fn func(string, DataStoreChannelStats, int) string) []string {
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
func (self *DataStoreChannelStatsStream) ReduceInt(fn func(int, DataStoreChannelStats, int) int) []int {
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
func (self *DataStoreChannelStatsStream) ReduceInt32(fn func(int32, DataStoreChannelStats, int) int32) []int32 {
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
func (self *DataStoreChannelStatsStream) ReduceInt64(fn func(int64, DataStoreChannelStats, int) int64) []int64 {
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
func (self *DataStoreChannelStatsStream) ReduceFloat32(fn func(float32, DataStoreChannelStats, int) float32) []float32 {
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
func (self *DataStoreChannelStatsStream) ReduceFloat64(fn func(float64, DataStoreChannelStats, int) float64) []float64 {
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
func (self *DataStoreChannelStatsStream) ReduceBool(fn func(bool, DataStoreChannelStats, int) bool) []bool {
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
func (self *DataStoreChannelStatsStream) Reverse() *DataStoreChannelStatsStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *DataStoreChannelStatsStream) Replace(fn func(DataStoreChannelStats, int) DataStoreChannelStats) *DataStoreChannelStatsStream {
	return self.ForEach(func(v DataStoreChannelStats, i int) { self.Set(i, fn(v, i)) })
}
func (self *DataStoreChannelStatsStream) Select(fn func(DataStoreChannelStats) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *DataStoreChannelStatsStream) Set(index int, val DataStoreChannelStats) *DataStoreChannelStatsStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *DataStoreChannelStatsStream) Skip(skip int) *DataStoreChannelStatsStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *DataStoreChannelStatsStream) SkippingEach(fn func(DataStoreChannelStats, int) int) *DataStoreChannelStatsStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *DataStoreChannelStatsStream) Slice(startIndex, n int) *DataStoreChannelStatsStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []DataStoreChannelStats{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *DataStoreChannelStatsStream) Sort(fn func(i, j int) bool) *DataStoreChannelStatsStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *DataStoreChannelStatsStream) Tail() *DataStoreChannelStats {
	return self.Last()
}
func (self *DataStoreChannelStatsStream) TailOr(arg DataStoreChannelStats) DataStoreChannelStats {
	return self.LastOr(arg)
}
func (self *DataStoreChannelStatsStream) ToList() []DataStoreChannelStats {
	return self.Val()
}
func (self *DataStoreChannelStatsStream) Unique() *DataStoreChannelStatsStream {
	return self.Distinct()
}
func (self *DataStoreChannelStatsStream) Val() []DataStoreChannelStats {
	if self == nil {
		return []DataStoreChannelStats{}
	}
	return *self.Copy()
}
func (self *DataStoreChannelStatsStream) While(fn func(DataStoreChannelStats, int) bool) *DataStoreChannelStatsStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *DataStoreChannelStatsStream) Where(fn func(DataStoreChannelStats) bool) *DataStoreChannelStatsStream {
	result := DataStoreChannelStatsStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *DataStoreChannelStatsStream) WhereSlim(fn func(DataStoreChannelStats) bool) *DataStoreChannelStatsStream {
	result := DataStoreChannelStatsStreamOf()
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
