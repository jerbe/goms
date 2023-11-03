package utils

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/24 01:24
  @describe :
*/

type Comparable interface {
	Equals(obj any) bool
}

// Equals 进行Equals对比
func Equals(a, b any) bool {
	if a == b {
		return true
	}

	if o, ok := a.(Comparable); ok && o.Equals(b) {
		return true
	}

	o, ok := b.(Comparable)
	return ok && o.Equals(a)
}

// Triple 三元对象
type Triple[X any, Y any, Z any] struct {
	x X
	y Y
	z Z
}

// Left 获取左边数据
func (t *Triple[X, Y, Z]) Left() X {
	return t.x
}

// Mid 获取中间数据
func (t *Triple[X, Y, Z]) Mid() Y {
	return t.y
}

// Right 获取右边数据
func (t *Triple[X, Y, Z]) Right() Z {
	return t.z
}
