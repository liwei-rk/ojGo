package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	for {
		if _, err := fmt.Scanf("%v %v %v %v %v %v", &x1, &y1, &x2, &y2, &x3, &y3); err != nil {
			return
		}
		x, y := bestPoint(x1, y1, x2, y2, x3, y3)
		fmt.Printf("%.6f %.6f\n", x, y)
	}
}

/*
已知平面上三点A(x1,y1), B(x2, y2), C(x3, y3)
求与这三点距离之和最小的点的坐标

即求三角形的费马点 —— 好吧,题目里有费马（Fermat）~
(1)若三角形ABC的3个内角均小于120°，那么3条距离连线正好平分费马点所在的周角。
(2)若三角形有一内角不小于120度，则此钝角的顶点就是距离和最小的点。

有一个与本问题等价的物理问题：在一个桌面上钻三个孔，取三根细线，它们的一
端连接在一起，另外一端分别从桌面的三个孔穿下去，并在细线下端挂三个等重的
砝码，细线与桌面以及小孔的摩擦力忽略不计。当系统达到静止平衡状态的时候，
桌面上三根线的连接点就是费马点。
参考维基百科： https://zh.wikipedia.org/wiki/%E8%B2%BB%E9%A6%AC%E9%BB%9E
*/
func bestPoint(x1, y1, x2, y2, x3, y3 float64) (float64, float64) {
	A := Pt(x1, y1)
	B := Pt(x2, y2)
	C := Pt(x3, y3)

	if A.Equal(B) {
		return A.X, A.Y
	}
	if A.Equal(C) {
		return A.X, A.Y
	}
	if B.Equal(C) {
		return B.X, B.Y
	}

	t := Triangle{A: A, B: B, C: C}
	p := t.GetFermatPoint()
	return p.X, p.Y
}

// 同时适用于三点在一条直线上：如果三点共线，会返回中间点
func (t Triangle) GetFermatPoint() Point {
	if point, yes := t.HasAngleMoreThan120(); yes {
		return point
	}
	A1 := GetAuxiliaryPoint(t.B, t.C, t.A)
	B1 := GetAuxiliaryPoint(t.A, t.C, t.B)
	return GetCrossPoint(t.A, A1, t.B, B1)
}

// 浮点数的比较~
type FloatCmp func() float64

const DefaultPrecision = 0.0000001

var cmp FloatCmp = func() float64 {
	return DefaultPrecision
}

func (c FloatCmp) Equal(a, b float64) bool {
	return math.Abs(a-b) < c()
}

func (c FloatCmp) Less(a, b float64) bool {
	return c.Equal(math.Max(a, b), b) && math.Abs(a-b) > c()
}

type Point struct {
	X float64
	Y float64
}

func Pt(x, y float64) Point {
	return Point{X: x, Y: y}
}

func (p Point) Equal(o Point) bool {
	return cmp.Equal(p.X, o.X) && cmp.Equal(p.Y, o.Y)
}

type Vector struct {
	X float64
	Y float64
}

func GenVector(from, to Point) Vector {
	return Vector{X: to.X - from.X, Y: to.Y - from.Y}
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Scale(s float64) Vector {
	return Vector{X: v.X * s, Y: v.Y * s}
}

func (v Vector) Opposite() Vector {
	return v.Scale(-1)
}

// 垂直向量，长度与原向量相等
func (v Vector) Vertical() Vector {
	return Vector{X: v.Y, Y: -v.X}
}

func (v Vector) ScaleToLen(length float64) Vector {
	return v.Scale(length / v.Len())
}

func Dot(v1, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func Cos(v1, v2 Vector) float64 {
	return Dot(v1, v2) / v1.Len() / v2.Len()
}

func (v Vector) LenSqure() float64 {
	return v.X*v.X + v.Y*v.Y
}

type Triangle struct {
	A Point
	B Point
	C Point
}

func (t Triangle) HasAngleMoreThan120() (Point, bool) {
	AB := GenVector(t.B, t.A)
	AC := GenVector(t.C, t.A)
	BC := GenVector(t.C, t.B)
	BA := AB.Opposite()
	CA := AC.Opposite()
	CB := BC.Opposite()
	if cmp.Less(Cos(AB, AC), -0.5) { // Cos(AB, AC) < -0.5 <=> angle A > 120
		return t.A, true
	}
	if cmp.Less(Cos(BC, BA), -0.5) {
		return t.B, true
	}
	if cmp.Less(Cos(CA, CB), -0.5) {
		return t.C, true
	}
	return Point{}, false
}

// 求O点关于边AB的对应点
// 对应点即AB为一边的正三角形的另一个点，与O分居AB直线两侧
func GetAuxiliaryPoint(A, B, O Point) Point {
	AB := GenVector(A, B)
	sqrt3 := math.Sqrt(3)
	verticalAB := AB.Vertical().Scale(0.5 * sqrt3)
	halfAB := AB.Scale(0.5)
	negVerticalAB := verticalAB.Opposite()

	O1 := Point{X: A.X + halfAB.X + verticalAB.X, Y: A.Y + halfAB.Y + verticalAB.Y}
	O2 := Point{X: A.X + halfAB.X + negVerticalAB.X, Y: A.Y + halfAB.Y + negVerticalAB.Y}

	if cmp.Less(GenVector(O, O1).Len(), GenVector(O, O2).Len()) {
		return O2
	}
	return O1
}

// 求直线AA1和BB1的交点
// 使用公式
func GetCrossPoint(A, A1, B, B1 Point) Point {
	m := (A.X-A1.X)*(B.Y-B1.Y) - (A.Y-A1.Y)*(B.X-B1.X)
	if cmp.Equal(m, 0) { //平行或重合
		panic("no cross point")
	}
	a := A.X*A1.Y - A.Y*A1.X
	b := B.X*B1.Y - B.Y*B1.X
	x := a*(B.X-B1.X) - (A.X-A1.X)*b
	y := a*(B.Y-B1.Y) - (A.Y-A1.Y)*b
	return Point{X: x / m, Y: y / m}
}
