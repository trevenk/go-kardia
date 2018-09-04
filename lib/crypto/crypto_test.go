package crypto

import (
	"bytes"
	"encoding/hex"
	"github.com/kardiachain/go-kardia/lib/common"
	"testing"
)

func TestKeccak256Hash(t *testing.T) {
	msg := []byte("abc")
	exp, _ := hex.DecodeString("4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45")
	verifyHash(t, "Sha3-256-array", func(in []byte) []byte { h := Keccak256Hash(in); return h[:] }, msg, exp)
}

func TestNewContractAddress(t *testing.T) {
	testAddrHex := "970e8128ab834e8eac17ab8e3812f010678cf791"
	testPrivHex := "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"

	key, _ := HexToECDSA(testPrivHex)
	addr := common.HexToAddress(testAddrHex)
	genAddr := PubkeyToAddress(key.PublicKey)
	verifyAddr(t, genAddr, addr)

	addr0 := CreateAddress(addr, 0)
	addr1 := CreateAddress(addr, 1)

	verifyAddr(t, common.HexToAddress("333c3310824b7c685133f2bedb2ca4b8b4df633d"), addr0)
	verifyAddr(t, common.HexToAddress("8bda78331c916a08481428e4b07c96d3e916d165"), addr1)

}

func verifyHash(t *testing.T, name string, f func([]byte) []byte, msg, exp []byte) {
	sum := f(msg)
	if !bytes.Equal(exp, sum) {
		t.Fatalf("hash %s mismatch: want: %x have: %x", name, exp, sum)
	}
}

func verifyAddr(t *testing.T, addr0, addr1 common.Address) {
	if addr0 != addr1 {
		t.Fatalf("address mismatch: want: %x have: %x", addr0, addr1)
	}
}
