package main

import (
	"fmt"
	b "blockchain/blockchain"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	_, err1 := ioutil.ReadFile("data/privKey")
	_, err2 := ioutil.ReadFile("data/id")

	if  (err1 != nil) || (err2 != nil)  {
		fmt.Println("No user found : creating a new user")

		data, err := ioutil.ReadFile("data/users.json")
		if err != nil {
			fmt.Println("Error ! Couldn't open users.json")
			return
		}
		users := b.JsonToUsers(data)

		keyPair := b.GenKeyPair()
		users = b.AddUser(users, len(users)+1, keyPair.PublicKey)

		b.SavePrivateKey("data/privKey", keyPair.PrivateKey)
		ioutil.WriteFile("data/id", []byte(strconv.Itoa(len(users))), 0644)
		ioutil.WriteFile("data/users.json", b.UsersToJson(users), 0644)

		return
	}

	data, _ := ioutil.ReadFile("data/id")
	userId, _ := strconv.Atoi(string(data))

	data, _ = ioutil.ReadFile("data/users.json")
	users := b.JsonToUsers(data)

	data, err1 = ioutil.ReadFile("data/blockchains.json")
	if err1 != nil {
		ioutil.WriteFile("data/blockchains.json", b.BlockchainsToJson(b.CreateBlockchainsData()), 0644)
	}
	blockchainsData := b.JsonToBlockchains(data)

	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "-u":
		fmt.Println(userId)

	case "-s":
		switch len(os.Args) {
		case 2:
			for _, i := range blockchainsData {
				fmt.Println(i.Id, len(i.Data.Blocks))
			}
		case 3:
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Couldn't convert the second argument to an int")
				return
			}
			for i1, i2 := range blockchainsData {
				if id == i2.Id {
					for k1, k2 := range i2.Data.Blocks {
						fmt.Println(k1+1, k2.SignedHash, "\n")
					}
					return
				} else if i1 == len(blockchainsData)-1 {
					fmt.Println("Coudln't find the blockchain")
				}
			}
		case 4:
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Couldn't convert the second argument to an int")
				return
			}
			block, err := strconv.Atoi(os.Args[3])
			if err != nil {
				fmt.Println("Couldn't convert the third argument to an int")
				return
			}
			for i1, i2 := range blockchainsData {
				if id == i2.Id {
					for k1, k2 := range i2.Data.Blocks {
						if k1+1 == block {
							fmt.Println(k2.PreviousSignedHash, "\n", k2.SignedHash, "\n", k2.UserId, "\n", k2.OtherId, "\n", k2.Time)
							return
						} else if k1 == len(i2.Data.Blocks)-1 {
							fmt.Println("Couldn't find the block")
							return
						}
					}
				} else if i1 == len(blockchainsData)-1 {
					fmt.Println("Coudln't find the blockchain")
				}
			}
		}

	case "-n":
		blockchainsData = b.AddBlockchain(blockchainsData, len(blockchainsData)+1, *b.NewBlockchain(userId))
		ioutil.WriteFile("data/blockchains.json", b.BlockchainsToJson(blockchainsData), 0644)

	case "-b":
		if (len(os.Args) > 3) {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Couldn't convert the second argument to an int")
				return
			}
			
			for i1, i2 := range blockchainsData {
				if id == i2.Id {
					otherId, err := strconv.Atoi(os.Args[3])
					if err != nil {
						fmt.Println("Couldn't convert the third argument to an int")
						return
					}
					blockchainsData[id-1].Data.AddBlock(userId, otherId)
					ioutil.WriteFile("data/blockchains.json", b.BlockchainsToJson(blockchainsData), 0644)
					return
				} else if i1 == len(blockchainsData)-1 {
					fmt.Println("Couldn't find the blockchain")
				}
			}
		}

	case "-v":
		if (len(os.Args) == 2) {
			for _, i := range blockchainsData {
				if i.Data.Verify(users) == 1 {
					fmt.Println(i.Id, "Valid")
				} else {
					fmt.Println(i.Id, "Compromised")
				}
			}
		} else if (len(os.Args) == 3) {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Couldn't convert the second argument to an int")
				return
			}
			for _, i := range blockchainsData {
				if i.Id == id {
					if i.Data.Verify(users) == 1 {
						fmt.Println(i.Id, "Valid")
					} else {
						fmt.Println(i.Id, "Compromised")
					}
					return
				}
			}
			fmt.Println("Couldn't find the blockchain")
		}
	}
}