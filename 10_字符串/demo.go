package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"token-contract/mytoken"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `
{
  "address": "1a9ec3b0b807464e6d3398a59d6b0a369bf422fa",
  "crypto": {
    "cipher": "aes-128-ctr",
    "ciphertext": "a471054846fb03e3e271339204420806334d1f09d6da40605a1a152e0d8e35f3",
    "cipherparams": {
      "iv": "44c5095dc698392c55a65aae46e0b5d9"
    },
    "kdf": "scrypt",
    "kdfparams": {
      "dklen": 32,
      "n": 262144,
      "p": 1,
      "r": 8,
      "salt": "e0a5fbaecaa3e75e20bccf61ee175141f3597d3b1bae6a28fe09f3507e63545e"
    },
    "mac": "cb3f62975cf6e7dfb454c2973bdd4a59f87262956d5534cdc87fb35703364043"
  },
  "id": "e08301fb-a263-4643-9c2b-d28959f66d6a",
  "version": 3
}
`

func main() {
    conn, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }
    auth, err := bind.NewTransactor(strings.NewReader(key), "123")
    if err != nil {
        log.Fatalf("Failed to create authorized transactor: %v", err)
    }
    address, tx, token, err := mytoken.DeployMyToken(auth, conn, big.NewInt(9651), "Contracts in Go!!!", 0, "Go!")
    if err != nil {
        log.Fatalf("Failed to deploy new token contract: %v", err)
    }
    fmt.Printf("Contract pending deploy: 0x%x\n", address)
    fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
    startTime := time.Now()
    fmt.Printf("TX start @:%s", time.Now())
    ctx := context.Background()
    addressAfterMined, err := bind.WaitDeployed(ctx, conn, tx)
    if err != nil {
        log.Fatalf("failed to deploy contact when mining :%v", err)
    }
    fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
    if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
        log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
    }
    name, err := token.Name(&bind.CallOpts{Pending: true})
    if err != nil {
        log.Fatalf("Failed to retrieve pending name: %v", err)
    }
    fmt.Println("Pending name:", name)
}