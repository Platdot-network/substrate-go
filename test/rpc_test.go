package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rjman-ljm/go-substrate-crypto/ss58"

	"github.com/Platdot-Network/substrate-go/client"
)

//const url = "wss://chainx.elara.patract.io"
//const url = "ws://127.0.0.1:7087"
// const url = "wss://kusama.elara.patract.io"

// const url = "wss://kusama.api.onfinality.io/public-ws"

//const url = "wss://polkadot.elara.patract.io"
//const url = "wss://rpc.polkadot.io"

// const url = "wss://kusama-rpc.polkadot.io"

const url = "wss://polkadot.api.onfinality.io/public-ws"

// const url = "wss://pub.elara.patract.io/polkadot"

//const url = "wss://dot.supercube.pro/ws"
//const url = "wss://chainx.supercube.pro/ws"
// const url = "ws://localhost:9944"

//const url = "wss://supercube.pro/ws"

//const startBlock = 5715297	/* polkadot.event.proxy_executed */
const startBlock = 7467638

// const endBlock = 7467638

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.PolkadotPrefix)
	//c.Name = expand.ClientNameChainX
	// for i := startBlock; i <= endBlock; i++ {
	//fmt.Printf("poll block#%v\n", i)
	//expand.SetSerDeOptions(false)
	resp, err := c.GetBlockByNumber(int64(startBlock))
	if err != nil {
		fmt.Printf("meet err: %v\n", err)
		//t.Fatal(err)
	}

	//hash, err := c.Api.RPC.Chain.GetBlockHash(uint64(startBlock))
	//if err != nil {
	//	fmt.Printf("GetBlockHash err\n")
	//}

	//block, err := c.Api.RPC.Chain.GetBlock(hash)
	//if err != nil {
	//	fmt.Printf("GetBlock err\n")
	//api, _ := gsrpc.NewSubstrateAPI(url)
	//
	//hash, err := api.RPC.Chain.GetBlockHash(4744169)
	//if err != nil {
	//	fmt.Printf("GetBlockHash err\n")
	//}
	//
	//block, err := api.RPC.Chain.GetBlock(hash)
	//if err != nil {
	//	fmt.Printf("Get Block err\n")
	//}
	//
	//if block != nil {
	//	currentBlock := int64(block.Block.Header.Number)
	//	fmt.Printf("block is %v\n", currentBlock)
	//}
	//}
	//
	//if block != nil {
	//	currentBlock := int64(block.Block.Header.Number)
	//	log.Debug("block is %v\n", currentBlock)
	//}

	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
	// }
}

func Test_GetAccountInfo(t *testing.T) {
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.PolkadotPrefix)
	ai, err := c.GetAccountInfo("1BzS3MjX9RXS2wz4tzHixasQqje5mMJcfTPPTyGQAEao3jU")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(ai)
	fmt.Println(string(d))
	fmt.Println(ai.Data.Free.String())
}
