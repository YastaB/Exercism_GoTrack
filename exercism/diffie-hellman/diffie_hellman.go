package diffiehellman

import (
	"math"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {

}

func PublicKey(A, p *big.Int, g int64) *big.Int {
	
	return big.Int.DivMod(g, A, p)

}

func SecretKey(a, B, p *big.Int) *big.Int {

}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {

}
