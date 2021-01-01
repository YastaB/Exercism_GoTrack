package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey generates private key
func PrivateKey(p *big.Int) *big.Int {
	rand.Seed(time.Now().UnixNano())
	val := &big.Int{}
	rnd := big.NewInt(rand.Int63())
	val = val.Mod(rnd, p) // p = p - 1
	if val.Cmp(big.NewInt(0)) == 0 || val.Cmp(big.NewInt(1)) == 0 {
		val = big.NewInt(2)
	}
	return val
}

// PublicKey generates public key
func PublicKey(A, p *big.Int, g int64) *big.Int {
	result := big.Int{}
	return result.Exp(big.NewInt(g), A, p)
}

// SecretKey generates received public and own private
func SecretKey(a, B, p *big.Int) *big.Int {
	secret := &big.Int{}
	return secret.Exp(B, a, p)
}

// NewPair generates private and public key pairs
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
	pub := PublicKey(priv, p, g)
	return priv, pub
}
