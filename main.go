package main

import  (
  "github.com/Dmitry1007/golang-blockchain/blockchain"
  "fmt"
  "strconv"
)

func main()  {
  chain := blockchain.InitBlockChain()

  chain.AddBlock("Second Block After Genesis")
  chain.AddBlock("Third Block After Genesis")

  for _, block := range chain.Blocks {
    fmt.Printf("Previous Hash: %x\n", block.PrevHash)
    fmt.Printf("Data in Block: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)

    pow := blockchain.NewProof(block)
    fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
    fmt.Println()
  }

}
