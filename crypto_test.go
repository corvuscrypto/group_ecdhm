package gecdhm

import (
	"crypto/elliptic"
	"crypto/rand"
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
