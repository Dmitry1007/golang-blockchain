package main

import  (
  "github.com/Dmitry1007/golang-blockchain/blockchain"
  "fmt"
)

func main()  {
  chain := blockchain.InitBlockChain()

  chain.AddBlock("First Block After Genesis")
  chain.AddBlock("Second Block After Genesis")
  chain.AddBlock("Third Block After Genesis")

  for _, block := range chain.Blocks {
    fmt.Printf("Previous Hash: %x\n", block.PrevHash)
    fmt.Printf("Data in Block: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)
  }

}
