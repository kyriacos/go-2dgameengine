package vec

type Vector2 struct {
	X, Y float32
}

// Mul performs a scalar multiplication between the vector and some constant value c.
func (v1 Vector2) Mul(c float32) Vector2 {
	return Vector2{v1.X * c, v1.Y * c}
}

// Add is equivalent to v3 := v1+v2
func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{v1.X + v2.X, v1.Y + v2.Y}
}

// Sub is equivalent to v3 := v1-v2
func (v1 Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{v1.X - v2.X, v1.Y - v2.Y}
}

// type Vector2Int32 struct {
// 	X, Y int32
// }

// // Mul performs a scalar multiplication between the vector and some constant value c.
// func (v1 Vector2Int32) Mul(c int) Vector2Int32 {
// 	return Vector2Int32{v1.X * c, v1.Y * c}
// }

// // Add is equivalent to v3 := v1+v2
// func (v1 Vector2Int32) Add(v2 Vector2) Vector2Int32 {
// 	return Vector2Int32{v1.X + v2.X, v1.Y + v2.Y}
// }

// // Sub is equivalent to v3 := v1-v2
// func (v1 Vector2Int32) Sub(v2 Vector2) Vector2Int32 {
// 	return Vector2Int32{v1.X - v2.X, v1.Y - v2.Y}
// }
