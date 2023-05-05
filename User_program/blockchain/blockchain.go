package blockchain

import (
	"strconv"
	"bytes"
	"time"
)

type Block struct {
	PreviousSignedHash []byte
	SignedHash []byte
	UserId int
	OtherId int
	Time int64
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain(userId int) *Blockchain {
	return &Blockchain{[]*Block{FirstBlock(userId)}}
}

func FirstBlock(userId int) *Block {
	return NewBlock([]byte{}, userId, userId)
}

func NewBlock(previousSignedHash []byte, userId int, otherId int) *Block {
	block := &Block{previousSignedHash, []byte{}, userId, otherId, time.Now().Unix()}
	block.SetSignedHash()

	return block
}

func (block *Block) CalculateHash() []byte {
	allData := bytes.Join([][]byte{block.PreviousSignedHash, []byte(strconv.Itoa(block.UserId)), []byte(strconv.Itoa(block.OtherId)), []byte(strconv.FormatInt(block.Time, 10))}, []byte{})
	
	return HashMsg(string(allData))
}

func (block *Block) SetSignedHash() {
	key := OpenPrivateKey("data/privKey")
	block.SignedHash = Sign(block.CalculateHash(), key)
}

func (blockchain *Blockchain) AddBlock(userId int, otherId int) {
	previousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(previousBlock.SignedHash, userId, otherId)
	
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func (blockchain *Blockchain) Verify(users []UsersData) int {
	for i1, i2 := range blockchain.Blocks {
		if len(blockchain.Blocks) > i1+1 {
			if i1 != len(blockchain.Blocks)-1 {
				if (blockchain.Blocks[i1].OtherId != blockchain.Blocks[i1+1].UserId) {
					return 0
				}
			}
		}

		hash := i2.CalculateHash()
		for k1, k2 := range users {
			if k2.Id == i2.UserId {
				if VerifySignature(hash, i2.SignedHash, k2.PubKey) == 0 {
					return 0
				}
			} else if k1 == len(users)-1 {
				return 0
			}
		}
	}

	return 1
}