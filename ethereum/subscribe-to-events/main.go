package main

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	nodeUrl := url.URL{
		Scheme: "wss",
		//Host:   "mainnet.infura.io",
		Host: "ropsten.infura.io",
		Path: "/ws",
	}
	c, _, err := websocket.DefaultDialer.Dial(nodeUrl.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer func() {
		if err := c.Close(); err != nil {
			log.Fatal("node connection close:", err)
		}
	}()

	if err := c.WriteMessage(
		websocket.TextMessage,
		[]byte(`{"id":"subscribeTransactions","method":"eth_subscribe","params":["newPendingTransactions"]}`),
	); err != nil {
		log.Fatal("subscription failed:", err)
	}

	if err := c.WriteMessage(
		websocket.TextMessage,
		[]byte(`{"id":"subscribeBlocks","method":"eth_subscribe","params":["newHeads"]}`),
	); err != nil {
		log.Fatal("subscription failed:", err)
	}

	for {
		/* mb = Message Bytes */
		_, mb, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read message:", err)
		}

		var m Message
		if err := json.Unmarshal(mb, &m); err != nil {
			log.Fatal("decode failed:", err)
		}

		log.Printf("%+v", m)

	}
}

type Params struct {
	Id     string `json:"id,omitempty"`
	Result string `json:"result,omitempty"`
	Params struct {
		Subscription string `json:"subscription"`
	} `json:"params,omitempty"`
}

type BlockMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription string `json:"subscription"`
		Result       struct {
			Number           string   `json:"number"`
			Hash             string   `json:"hash"`
			ParentHash       string   `json:"parentHash"`
			Sha3Uncles       string   `json:"sha3Uncles"`
			LogsBloom        string   `json:"logsBloom"`
			TransactionsRoot string   `json:"transactionsRoot"`
			StateRoot        string   `json:"stateRoot"`
			ReceiptsRoot     string   `json:"receiptsRoot"`
			Miner            string   `json:"miner"`
			Difficulty       string   `json:"difficulty"`
			ExtraData        string   `json:"extraData"`
			GasLimit         string   `json:"gasLimit"`
			GasUsed          string   `json:"gasUsed"`
			Timestamp        string   `json:"timestamp"`
			Nonce            string   `json:"nonce"`
			MixHash          string   `json:"mixHash"`
			Address          string   `json:"address"`
			Topics           []string `json:"topics"`
			Data             string   `json:"data"`
			BlockNumber      string   `json:"blockNumber"`
			TransactionHash  string   `json:"transactionHash"`
			TransactionIndex string   `json:"transactionIndex"`
			BlockHash        string   `json:"blockHash"`
			LogIndex         string   `json:"logIndex"`
			Removed          bool     `json:"removed"`
		} `json:"result"`
	} `json:"params"`
}

type TransactionMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription  string `json:"subscription"`
		TransactionId string `json:"result"`
	} `json:"params"`
}

type TokenInboundMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription string `json:"subscription"`
		Result       struct {
			Address          string   `json:"address"`
			Topics           []string `json:"topics"`
			Data             string   `json:"data"`
			BlockNumber      string   `json:"blockNumber"`
			TransactionHash  string   `json:"transactionHash"`
			TransactionIndex string   `json:"transactionIndex"`
			BlockHash        string   `json:"blockHash"`
			LogIndex         string   `json:"logIndex"`
			Removed          bool     `json:"removed"`
		} `json:"result"`
	} `json:"params"`
}
