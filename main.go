package main

import  (
  "github.com/Dmitry1007/golang-blockchain/blockchain"
  "fmt"
  // "strconv"
)

func main()  {
  chain := blockchain.InitBlockChain()
  fmt.Println(
    "chain ", chain, "\n",
    "chain.Blocks ", chain.Blocks, "\n",
    "chain.Blocks[0] ", chain.Blocks[0], "\n",
    "chain.Blocks[0].Hash ", chain.Blocks[0].Hash, "\n",
    "chain.Blocks[0].Data ", chain.Blocks[0].Data, "\n",
    "chain.Blocks[0].PrevHash ", chain.Blocks[0].PrevHash, "\n",
    "chain.Blocks[0].Nonce ", chain.Blocks[0].Nonce, "\n",
  )
  // chain.AddBlock("First Block After Genesis")
  // chain.AddBlock("Second Block After Genesis")
  // chain.AddBlock("Third Block After Genesis")
  //
  // for _, block := range chain.Blocks {
  //   fmt.Printf("Previous Hash: %x\n", block.PrevHash)
  //   fmt.Printf("Data in Block: %s\n", block.Data)
  //   fmt.Printf("Hash: %x\n", block.Hash)
  //
  //   pow := blockchain.NewProof(block)
  //   fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
  //   fmt.Println()
  // }

}
