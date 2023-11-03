package data

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/31 20:45
  @describe :
*/

// NewVector 返回一个新的矢量
func NewVector(x, y int) *Vector {
	return &Vector{X: x, Y: y}
}

// Vector 向量
type Vector struct {
	// X 坐标X
	X int

	// Y 坐标Y
	Y int
}

// DistanceSqrt 指定X/Y位置的距离
func (v *Vector) DistanceSqrt(pointX, pointY float64) float64 {
	pointX -= float64(v.X)
	pointY -= float64(v.Y)
	return pointX*pointX + pointY*pointY
}

// DistanceSqrtPoint 指定矢量的位置的距离
func (v *Vector) DistanceSqrtPoint(vet *Vector) float64 {
	pointX := vet.X - v.X
	pointY := vet.X - v.Y

	return float64(pointX*pointX + pointY*pointY)
}

// PlusX X 坐标加上x数值并返回新的向量
func (v *Vector) PlusX(x int) *Vector {
	vet := *v
	vet.X += x
	return &vet
}

// MinusX X 坐标减去x数值并返回新的向量
func (v *Vector) MinusX(x int) *Vector {
	return v.PlusX(-x)
}

// PlusY Y 坐标加上y数值并返回新的向量
func (v *Vector) PlusY(y int) *Vector {
	vet := *v
	vet.Y += y
	return &vet
}

// MinusY Y 坐标减去y数值并返回新的向量
func (v *Vector) MinusY(y int) *Vector {
	return v.PlusY(-y)
}

// Plus  X Y 坐标分别加上x,y数值并返回新的向量
func (v *Vector) Plus(x, y int) *Vector {
	vet := *v
	vet.X += x
	vet.Y += y
	return &vet
}

// PlusVector 加上指定目标向量坐标点,并返回一个新的向量
func (v *Vector) PlusVector(vector *Vector) *Vector {
	vet := *v
	vet.X += vector.X
	vet.Y += vector.Y
	return &vet
}

// Minus  X Y 坐标分别减去x,y数值并返回新的向量
func (v *Vector) Minus(x, y int) *Vector {
	vet := *v
	vet.X -= x
	vet.Y -= y
	return &vet
}

// MinusVector 减去指定目标向量坐标点,并返回一个新的向量
func (v *Vector) MinusVector(vector *Vector) *Vector {
	vet := *v
	vet.X -= vector.X
	vet.Y -= vector.Y
	return &vet
}

// Assignment 将pVet向量的值赋值给当前向量
func (v *Vector) Assignment(vet *Vector) {
	v.X = vet.X
	v.Y = vet.Y
}
