package blockchain

import (
	"encoding/json"
	"crypto/rsa"
	"io/ioutil"
)

type UsersData struct {
	Id int `json:"id"`
	PubKey *rsa.PublicKey `json:"pubKey"`
}

type BlockchainData struct {
	Id int `json:"id"`
	Data Blockchain `json:"blockchain"`
}

func CreateUsers(id int, pubKey *rsa.PublicKey) []UsersData {
	return append([]UsersData{}, UsersData{id, pubKey})
}

func AddUser(users []UsersData, id int, pubKey *rsa.PublicKey) []UsersData {
	return append(users, UsersData{id, pubKey})
}

func UsersToJson(users []UsersData) []byte {
	data, _ := json.Marshal(users)

	return data
}

func JsonToUsers(data []byte) []UsersData {
	users := []UsersData{}
	json.Unmarshal(data, &users)

	return users
}

func CreateBlockchainsData() []BlockchainData {
	return []BlockchainData{}
}

func AddBlockchain(data []BlockchainData, id int, blockchain Blockchain) []BlockchainData {
	return append(data, BlockchainData{id, blockchain})
}

func BlockchainsToJson(blockchains []BlockchainData) []byte {
	data, _ := json.Marshal(blockchains)

	return data
}

func JsonToBlockchains(data []byte) []BlockchainData {
	blockchains := []BlockchainData{}
	json.Unmarshal(data, &blockchains)

	return blockchains
}

func SaveData(path string, data []byte) {
	ioutil.WriteFile(path, data, 0644)
}

func OpenData(path string) []byte {
	data, _ := ioutil.ReadFile(path)

	return data
}