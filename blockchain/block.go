package blockchain

import  (
  "fmt"
  "bytes"
  "crypto/sha256"
)

type BlockChain struct {
  Blocks []*Block
}

type Block struct {
  Hash      []byte
  Data      []byte
  PrevHash  []byte
}

func CreateBlock(data string, prevHash []byte) *Block {
  block := &Block{[]byte{}, []byte(data), prevHash, 0}
  pow := NewProof(block)
  nonce, hash := pow.Run()
  block.Hash  = hash[:]
  block.Nonce = nonce
  return block
}

func (chain *BlockChain) AddBlock(data string) {
  prevBlock := chain.Blocks[len(chain.Blocks)-1]
  new := CreateBlock(data, prevBlock.Hash)
  chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
  return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
  fmt.Println("Init BlockChain")
  return &BlockChain{[]*Block{Genesis()}}
}
