/*
 * Collection utility of SingleDeviceEntitiesComponent Struct
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

type SingleDeviceEntitiesComponentStream []SingleDeviceEntitiesComponent

func SingleDeviceEntitiesComponentStreamOf(arg ...SingleDeviceEntitiesComponent) SingleDeviceEntitiesComponentStream {
	return arg
}
func SingleDeviceEntitiesComponentStreamFrom(arg []SingleDeviceEntitiesComponent) SingleDeviceEntitiesComponentStream {
	return arg
}
func CreateSingleDeviceEntitiesComponentStream(arg ...SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	tmp := SingleDeviceEntitiesComponentStreamOf(arg...)
	return &tmp
}
func GenerateSingleDeviceEntitiesComponentStream(arg []SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	tmp := SingleDeviceEntitiesComponentStreamFrom(arg)
	return &tmp
}

func (self *SingleDeviceEntitiesComponentStream) Add(arg SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	return self.AddAll(arg)
}
func (self *SingleDeviceEntitiesComponentStream) AddAll(arg ...SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	*self = append(*self, arg...)
	return self
}
func (self *SingleDeviceEntitiesComponentStream) AddSafe(arg *SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Aggregate(fn func(SingleDeviceEntitiesComponent, SingleDeviceEntitiesComponent) SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	result := SingleDeviceEntitiesComponentStreamOf()
	self.ForEach(func(v SingleDeviceEntitiesComponent, i int) {
		if i == 0 {
			result.Add(fn(SingleDeviceEntitiesComponent{}, v))
		} else {
			result.Add(fn(result[i-1], v))
		}
	})
	*self = result
	return self
}
func (self *SingleDeviceEntitiesComponentStream) AllMatch(fn func(SingleDeviceEntitiesComponent, int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}
func (self *SingleDeviceEntitiesComponentStream) AnyMatch(fn func(SingleDeviceEntitiesComponent, int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *SingleDeviceEntitiesComponentStream) Clone() *SingleDeviceEntitiesComponentStream {
	temp := make([]SingleDeviceEntitiesComponent, self.Len())
	copy(temp, *self)
	return (*SingleDeviceEntitiesComponentStream)(&temp)
}
func (self *SingleDeviceEntitiesComponentStream) Copy() *SingleDeviceEntitiesComponentStream {
	return self.Clone()
}
func (self *SingleDeviceEntitiesComponentStream) Concat(arg []SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	return self.AddAll(arg...)
}
func (self *SingleDeviceEntitiesComponentStream) Contains(arg SingleDeviceEntitiesComponent) bool {
	return self.FindIndex(func(_arg SingleDeviceEntitiesComponent, index int) bool { return reflect.DeepEqual(_arg, arg) }) != -1
}
func (self *SingleDeviceEntitiesComponentStream) Clean() *SingleDeviceEntitiesComponentStream {
	*self = SingleDeviceEntitiesComponentStreamOf()
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Delete(index int) *SingleDeviceEntitiesComponentStream {
	return self.DeleteRange(index, index)
}
func (self *SingleDeviceEntitiesComponentStream) DeleteRange(startIndex, endIndex int) *SingleDeviceEntitiesComponentStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Distinct() *SingleDeviceEntitiesComponentStream {
	caches := map[string]bool{}
	result := SingleDeviceEntitiesComponentStreamOf()
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
func (self *SingleDeviceEntitiesComponentStream) Each(fn func(SingleDeviceEntitiesComponent)) *SingleDeviceEntitiesComponentStream {
	for _, v := range *self {
		fn(v)
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) EachRight(fn func(SingleDeviceEntitiesComponent)) *SingleDeviceEntitiesComponentStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i))
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Equals(arr []SingleDeviceEntitiesComponent) bool {
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
func (self *SingleDeviceEntitiesComponentStream) Filter(fn func(SingleDeviceEntitiesComponent, int) bool) *SingleDeviceEntitiesComponentStream {
	result := SingleDeviceEntitiesComponentStreamOf()
	for i, v := range *self {
		if fn(v, i) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleDeviceEntitiesComponentStream) FilterSlim(fn func(SingleDeviceEntitiesComponent, int) bool) *SingleDeviceEntitiesComponentStream {
	result := SingleDeviceEntitiesComponentStreamOf()
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
func (self *SingleDeviceEntitiesComponentStream) Find(fn func(SingleDeviceEntitiesComponent, int) bool) *SingleDeviceEntitiesComponent {
	if i := self.FindIndex(fn); -1 != i {
		tmp := (*self)[i]
		return &tmp
	}
	return nil
}
func (self *SingleDeviceEntitiesComponentStream) FindOr(fn func(SingleDeviceEntitiesComponent, int) bool, or SingleDeviceEntitiesComponent) SingleDeviceEntitiesComponent {
	if v := self.Find(fn); v != nil {
		return *v
	}
	return or
}
func (self *SingleDeviceEntitiesComponentStream) FindIndex(fn func(SingleDeviceEntitiesComponent, int) bool) int {
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
func (self *SingleDeviceEntitiesComponentStream) First() *SingleDeviceEntitiesComponent {
	return self.Get(0)
}
func (self *SingleDeviceEntitiesComponentStream) FirstOr(arg SingleDeviceEntitiesComponent) SingleDeviceEntitiesComponent {
	if v := self.Get(0); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDeviceEntitiesComponentStream) ForEach(fn func(SingleDeviceEntitiesComponent, int)) *SingleDeviceEntitiesComponentStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) ForEachRight(fn func(SingleDeviceEntitiesComponent, int)) *SingleDeviceEntitiesComponentStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) GroupBy(fn func(SingleDeviceEntitiesComponent, int) string) map[string][]SingleDeviceEntitiesComponent {
	m := map[string][]SingleDeviceEntitiesComponent{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *SingleDeviceEntitiesComponentStream) GroupByValues(fn func(SingleDeviceEntitiesComponent, int) string) [][]SingleDeviceEntitiesComponent {
	var tmp [][]SingleDeviceEntitiesComponent
	for _, v := range self.GroupBy(fn) {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *SingleDeviceEntitiesComponentStream) IndexOf(arg SingleDeviceEntitiesComponent) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *SingleDeviceEntitiesComponentStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *SingleDeviceEntitiesComponentStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *SingleDeviceEntitiesComponentStream) Last() *SingleDeviceEntitiesComponent {
	return self.Get(self.Len() - 1)
}
func (self *SingleDeviceEntitiesComponentStream) LastOr(arg SingleDeviceEntitiesComponent) SingleDeviceEntitiesComponent {
	if v := self.Last(); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDeviceEntitiesComponentStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *SingleDeviceEntitiesComponentStream) Limit(limit int) *SingleDeviceEntitiesComponentStream {
	self.Slice(0, limit)
	return self
}

func (self *SingleDeviceEntitiesComponentStream) Map(fn func(SingleDeviceEntitiesComponent, int) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Map2Int(fn func(SingleDeviceEntitiesComponent, int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Map2Int32(fn func(SingleDeviceEntitiesComponent, int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Map2Int64(fn func(SingleDeviceEntitiesComponent, int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Map2Float32(fn func(SingleDeviceEntitiesComponent, int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Map2Float64(fn func(SingleDeviceEntitiesComponent, int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Map2Bool(fn func(SingleDeviceEntitiesComponent, int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Map2Bytes(fn func(SingleDeviceEntitiesComponent, int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Map2String(fn func(SingleDeviceEntitiesComponent, int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Max(fn func(SingleDeviceEntitiesComponent, int) float64) *SingleDeviceEntitiesComponent {
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
func (self *SingleDeviceEntitiesComponentStream) Min(fn func(SingleDeviceEntitiesComponent, int) float64) *SingleDeviceEntitiesComponent {
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
func (self *SingleDeviceEntitiesComponentStream) NoneMatch(fn func(SingleDeviceEntitiesComponent, int) bool) bool {
	return !self.AnyMatch(fn)
}
func (self *SingleDeviceEntitiesComponentStream) Get(index int) *SingleDeviceEntitiesComponent {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *SingleDeviceEntitiesComponentStream) GetOr(index int, arg SingleDeviceEntitiesComponent) SingleDeviceEntitiesComponent {
	if v := self.Get(index); v != nil {
		return *v
	}
	return arg
}
func (self *SingleDeviceEntitiesComponentStream) Peek(fn func(*SingleDeviceEntitiesComponent, int)) *SingleDeviceEntitiesComponentStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}


func (self *SingleDeviceEntitiesComponentStream) Reduce(fn func(SingleDeviceEntitiesComponent, SingleDeviceEntitiesComponent, int) SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	return self.ReduceInit(fn, SingleDeviceEntitiesComponent{})
}
func (self *SingleDeviceEntitiesComponentStream) ReduceInit(fn func(SingleDeviceEntitiesComponent, SingleDeviceEntitiesComponent, int) SingleDeviceEntitiesComponent, initialValue SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	result := SingleDeviceEntitiesComponentStreamOf()
	self.ForEach(func(v SingleDeviceEntitiesComponent, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}
func (self *SingleDeviceEntitiesComponentStream) ReduceInterface(fn func(interface{}, SingleDeviceEntitiesComponent, int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(SingleDeviceEntitiesComponent{}, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *SingleDeviceEntitiesComponentStream) ReduceString(fn func(string, SingleDeviceEntitiesComponent, int) string) []string {
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
func (self *SingleDeviceEntitiesComponentStream) ReduceInt(fn func(int, SingleDeviceEntitiesComponent, int) int) []int {
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
func (self *SingleDeviceEntitiesComponentStream) ReduceInt32(fn func(int32, SingleDeviceEntitiesComponent, int) int32) []int32 {
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
func (self *SingleDeviceEntitiesComponentStream) ReduceInt64(fn func(int64, SingleDeviceEntitiesComponent, int) int64) []int64 {
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
func (self *SingleDeviceEntitiesComponentStream) ReduceFloat32(fn func(float32, SingleDeviceEntitiesComponent, int) float32) []float32 {
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
func (self *SingleDeviceEntitiesComponentStream) ReduceFloat64(fn func(float64, SingleDeviceEntitiesComponent, int) float64) []float64 {
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
func (self *SingleDeviceEntitiesComponentStream) ReduceBool(fn func(bool, SingleDeviceEntitiesComponent, int) bool) []bool {
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
func (self *SingleDeviceEntitiesComponentStream) Reverse() *SingleDeviceEntitiesComponentStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Replace(fn func(SingleDeviceEntitiesComponent, int) SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	return self.ForEach(func(v SingleDeviceEntitiesComponent, i int) { self.Set(i, fn(v, i)) })
}
func (self *SingleDeviceEntitiesComponentStream) Select(fn func(SingleDeviceEntitiesComponent) interface{}) interface{} {
	_array := make([]interface{}, 0, len(*self))
	for _, v := range *self {
		_array = append(_array, fn(v))
	}
	return _array
}
func (self *SingleDeviceEntitiesComponentStream) Set(index int, val SingleDeviceEntitiesComponent) *SingleDeviceEntitiesComponentStream {
	if len(*self) > index && index >= 0 {
		(*self)[index] = val
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Skip(skip int) *SingleDeviceEntitiesComponentStream {
	return self.Slice(skip, self.Len()-skip)
}
func (self *SingleDeviceEntitiesComponentStream) SkippingEach(fn func(SingleDeviceEntitiesComponent, int) int) *SingleDeviceEntitiesComponentStream {
	for i := 0; i < self.Len(); i++ {
		skip := fn(*self.Get(i), i)
		i += skip
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Slice(startIndex, n int) *SingleDeviceEntitiesComponentStream {
	if last := startIndex + n; len(*self)-1 < startIndex || last < 0 || startIndex < 0 {
		*self = []SingleDeviceEntitiesComponent{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Sort(fn func(i, j int) bool) *SingleDeviceEntitiesComponentStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *SingleDeviceEntitiesComponentStream) Tail() *SingleDeviceEntitiesComponent {
	return self.Last()
}
func (self *SingleDeviceEntitiesComponentStream) TailOr(arg SingleDeviceEntitiesComponent) SingleDeviceEntitiesComponent {
	return self.LastOr(arg)
}
func (self *SingleDeviceEntitiesComponentStream) ToList() []SingleDeviceEntitiesComponent {
	return self.Val()
}
func (self *SingleDeviceEntitiesComponentStream) Unique() *SingleDeviceEntitiesComponentStream {
	return self.Distinct()
}
func (self *SingleDeviceEntitiesComponentStream) Val() []SingleDeviceEntitiesComponent {
	if self == nil {
		return []SingleDeviceEntitiesComponent{}
	}
	return *self.Copy()
}
func (self *SingleDeviceEntitiesComponentStream) While(fn func(SingleDeviceEntitiesComponent, int) bool) *SingleDeviceEntitiesComponentStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *SingleDeviceEntitiesComponentStream) Where(fn func(SingleDeviceEntitiesComponent) bool) *SingleDeviceEntitiesComponentStream {
	result := SingleDeviceEntitiesComponentStreamOf()
	for _, v := range *self {
		if fn(v) {
			result.Add(v)
		}
	}
	*self = result
	return self
}
func (self *SingleDeviceEntitiesComponentStream) WhereSlim(fn func(SingleDeviceEntitiesComponent) bool) *SingleDeviceEntitiesComponentStream {
	result := SingleDeviceEntitiesComponentStreamOf()
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
