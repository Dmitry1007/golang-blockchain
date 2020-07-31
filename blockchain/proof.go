package blockchain

import (
  "fmt"
)

// Take the data from the block

// Create a counter (nonce) which starts at 0

// Create a hash of the data plus the counter

// Check the hash to see if it meets a set of requirements

// Requirements:
// The First few bytes of the hash must contain 0s

const Difficulty = 12

type ProofOfWork struct {
  Block  *Block
  Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
  target := big.NewInt(1)
  target.Lsh(target, uint(256-Difficulty))
  pow := &ProofOfWork{b, target}
  return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
  data := bytes.Join(
    [][]byte{
      pow.Block.PrevHash,
      pow.Block.Data,
      ToHex(int64(nonce)),
      ToHex(int64(Difficulty)),
    },
    []byte{},
  )
  return data
}

func ToHex(num int64) []byte {
  buff := new(bytes.Buffer)
  err  := binary.Write(buff, binary.BigEndian, num)
  if err != nil {
    log.Panic(err)
  }
  return buff.Bytes()
}
