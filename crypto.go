package gecdhm

import (
	"crypto/elliptic"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
)

//Package errors related to handling of data
var (
	ErrMarshal   = errors.New("Unexpected error during data marshalling")
	ErrUnmarshal = errors.New("Unexpected error during data unmarshalling")
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

//Marshal transforms a Point into binary data
func (p Point) Marshal() (data []byte) {
	lengthBytesBuffer := make([]byte, 10)
	xBytes := p.x.Bytes()
	xByteLength := int64(len(xBytes))
	if p.x.Sign() == -1 {
		xByteLength = -xByteLength
		fmt.Println("xlenght", xByteLength)
	}
	n := binary.PutVarint(lengthBytesBuffer, xByteLength)
	data = append(data, lengthBytesBuffer[:n]...) //length
	data = append(data, xBytes...)

	yBytes := p.y.Bytes()
	yByteLength := int64(len(yBytes))
	if p.y.Sign() == -1 {
		yByteLength = -yByteLength
	}
	binary.PutVarint(lengthBytesBuffer, yByteLength)
	data = append(data, lengthBytesBuffer[:n]...)
	data = append(data, yBytes...)
	return
}

//Unmarshal transforms binary data into a Point
func (p *Point) Unmarshal(data []byte) (err error) {
	var i = int64(0)
	xByteLength, n := binary.Varint(data)
	xNegative := false
	if xByteLength < 0 {
		xByteLength = -xByteLength
		xNegative = true
	}
	i += int64(n)
	p.x = new(big.Int).SetBytes(data[i : i+xByteLength])
	if xNegative {
		p.x.Neg(p.x)
	}
	i += xByteLength
	yByteLength, n := binary.Varint(data[i:])
	yNegative := false
	if yByteLength < 0 {
		yByteLength = -yByteLength
		yNegative = true
	}
	i += int64(n)
	p.y = new(big.Int).SetBytes(data[i : i+yByteLength])
	if yNegative {
		p.y.Neg(p.y)
	}
	if recover() != nil {
		err = ErrUnmarshal
		p.x = nil
		p.y = nil
	}
	return
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
