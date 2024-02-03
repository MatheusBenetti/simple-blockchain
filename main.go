package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		PrevHash:  prevHash,
		Hash:      []byte{},
	}
	block.SetHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Oscar.")
	bc.AddBlock("Send 2 more BTC to Oscar.")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash:  %x\n", block.Hash)
		fmt.Println()
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewFirstBlock() *Block {
	return NewBlock("First Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewFirstBlock()}}
}
