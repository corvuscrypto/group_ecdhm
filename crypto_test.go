package gecdhm

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"
)

func TestCombinationEquality(T *testing.T) {
	curve := elliptic.P256()

	//We first generate two private keys
	k, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		T.Error(err)
	}
	k2, x2, y2, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		T.Error(err)
	}

	exampleKeyA := NewPrivateKey(curve, k, x, y)
	exampleKeyB := NewPrivateKey(curve, k2, x2, y2)

	//Based on EC theory, when we point multiply the two private k's with the public
	//points from each key they should be equal to each other (a shared secret)
	sharedAB := exampleKeyA.Combine(exampleKeyB.publicPoint)
	sharedBA := exampleKeyB.Combine(exampleKeyA.publicPoint)

	if !sharedAB.IsEqual(sharedBA) {
		T.Error("The keys are not equal. Oh noes!")
	}
}

func TestPointMarshalling(T *testing.T) {
	//initialize two big ints
	testX := big.NewInt(100000)
	testY := big.NewInt(-100000)

	//make two points, one from the previously specified big ints, one that is just nil
	pointA := NewPoint(testX, testY)
	var pointB Point

	data := pointA.Marshal()
	if err := pointB.Unmarshal(data); err != nil {
		T.Error(err)
	}

	if !pointA.IsEqual(pointB) {
		T.Error("Points were unequal")
		T.Log("Xs: \n", pointA.x, "\n", pointB.x)
		T.Log("Ys: \n", pointA.y, "\n", pointB.y)
	}
}
