package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	conn, err := ethclient.Dial("http://54.245.138.237:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	var addr = common.HexToAddress("0x4845d11aeed99d5b94742aa4bb684eba96b20e94")
	var uaddr = common.HexToAddress("bc006b353770becc7fdecfd11eff9633a3ea651f")
	c, _ := NewContract(addr, conn)

	bid := new(big.Int)
	bid.SetString("245666", 10)
	maxGas := new(big.Int)
	maxGas.SetString("470000", 10)

	am := accounts.NewManager("keystore", accounts.StandardScryptN, accounts.StandardScryptP)
	account, err := am.Find(accounts.Account{Address: uaddr})
	if err != nil {
		log.Fatalf("Account not found: %v", err)
	}

	accountJSON, err := am.Export(account, "password01", "blockparty")
	if err != nil {
		log.Fatalf("Account JSON failed:  %v", err)
	}

	key, err := accounts.DecryptKey(accountJSON, "blockparty")
	if err != nil {
		log.Fatalf("Decrypt key failed: %v", err)
	}

	signer := bind.NewKeyedTransactor(key.PrivateKey)

	_, err = c.PlaceBid(signer, bid)
	if err != nil {
		log.Fatalf("Bid failed: %v", err)
	}
}
