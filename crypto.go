package gecdhm

import (
	"crypto/elliptic"
	"math/big"
)

//Point is the representation of the x, y params created by point multiplication
//of the base point of the curve (G) with the private parameter (n)
type Point struct {
	x, y *big.Int
}

//NewPoint is a convenience function for creating a Point struct from parameters x
//and y. This makes code more readable in areas where we do point multiplication
//using the EC interface from the elliptic package
func NewPoint(x, y *big.Int) Point {
	return Point{
		x, y,
	}
}

//IsEqual checks if two points are equal
func (p Point) IsEqual(q Point) bool {
	return p.x.Cmp(q.x) == 0 && p.y.Cmp(q.y) == 0
}

//PrivateKey is the gecdhm representation of curve parameters specific to one party
type PrivateKey struct {
	curve       elliptic.Curve
	K           []byte
	publicPoint Point
}

//NewPrivateKey is a convenience function for creating a private key representation
//based on a curve, the private param k, and the x, y params
func NewPrivateKey(curve elliptic.Curve, k []byte, x, y *big.Int) (p PrivateKey) {
	return PrivateKey{
		curve, k, NewPoint(x, y),
	}
}

//Combine combines the public points of two keys into a new public point z by point
//EC multiplication.
func (p PrivateKey) Combine(q Point) (z Point) {
	return NewPoint(p.curve.ScalarMult(q.x, q.y, p.K))
}
