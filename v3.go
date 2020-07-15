package navmesh

import "math"

type V3 struct {
	X, Y, Z float64
}

func NewV3(x, y, z float64) *V3 {
	return &V3{X: x, Y: y, Z: z}
}

func (v *V3) Cross(v2 *V3) *V3 {
	return NewV3(v.Y*v2.Z-v.Z*v2.Y, v.Z*v2.X-v.X*v2.Z, v.X*v2.Y-v.Y*v2.X)
}

func (v *V3) Sub(v2 *V3) *V3 {
	return NewV3(v.X-v2.X, v.Y-v2.Y, v.Z-v2.Z)
}

func (v *V3) Distance(target *V3) float64 {
	dx := v.X - target.X
	dy := v.Y - target.Y
	return math.Sqrt(dx*dx + dy*dy)
}
