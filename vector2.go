package emath

import (
	"encoding/binary"
	"fmt"
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

func (v Vector2) Bytes() []byte {
	var b [16]byte
	binary.BigEndian.PutUint64(b[:8], math.Float64bits(v.X))
	binary.BigEndian.PutUint64(b[8:], math.Float64bits(v.Y))
	return b[:]
}

func (v Vector2) Abs() Vector2 {
	return Vector2{
		X: math.Abs(v.X),
		Y: math.Abs(v.Y),
	}
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector2) AddX(x float64) Vector2 {
	return Vector2{
		X: v.X + x,
		Y: v.Y,
	}
}

func (v Vector2) AddY(y float64) Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y + y,
	}
}

func (v Vector2) AddXY(x, y float64) Vector2 {
	return Vector2{
		X: v.X + x,
		Y: v.Y + y,
	}
}

func (v Vector2) Subtract(other Vector2) Vector2 {
	return Vector2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector2) Multiply(scalar float64) Vector2 {
	return Vector2{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vector2) Divide(scalar float64) Vector2 {
	return Vector2{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vector2) Dot(other Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2) Cross(other Vector2) float64 {
	return v.X*other.Y - v.Y*other.X
}

func (v Vector2) MagnitudeSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.MagnitudeSquared())
}

func (v Vector2) Length() float64 {
	return v.Magnitude()
}

func (v Vector2) DistanceTo(other Vector2) float64 {
	return other.Subtract(v).Magnitude()
}

func (v Vector2) DistanceSquaredTo(other Vector2) float64 {
	return other.Subtract(v).MagnitudeSquared()
}

func (v Vector2) WithinDistanceOf(other Vector2, distance float64) bool {
	return v.DistanceSquaredTo(other) <= distance*distance
}

func (v Vector2) Normalized() Vector2 {
	return v.Divide(v.Magnitude())
}

func (v Vector2) Equals(other Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vector2) EqualX(other Vector2) bool {
	return v.X == other.X
}

func (v Vector2) EqualY(other Vector2) bool {
	return v.Y == other.Y
}

func (v Vector2) DirectionTo(other Vector2) Vector2 {
	return other.Subtract(v).Normalized()
}
