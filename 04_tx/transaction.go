package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

const subsidy = 10

// Transaction represents a Bitcoin transaction
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
	Text string
}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

// SetID sets ID of a transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
	Text      string
}

// TXOutput represents a transaction output
type TXOutput struct {
	Value        int
	ScriptPubKey string
	Text         string
}

// CanUnlockOutputWith checks whether the address initiated the transaction
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	fmt.Printf("    - CanUnlockOutputWith: pk(%s) = unlockdata(%s)\n", in.ScriptSig, unlockingData)
	return in.ScriptSig == unlockingData
}

// CanBeUnlockedWith checks if the output can be unlocked with the provided data
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	fmt.Printf("    - CanBeUnlockedWith: pk(%s) = unlockdata(%s)\n", out.ScriptPubKey, unlockingData)
	return out.ScriptPubKey == unlockingData
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data, fmt.Sprintf("to: %s, data: %s", to, data)}
	txout := TXOutput{subsidy, to, fmt.Sprintf("to: %s, data: %s", to, data)}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}, ""}
	tx.SetID()

	return &tx
}

// NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		//spew.Dump(validOutputs)
		//log.Panic("ERROR: Not enough funds")
		fmt.Println("not enough funds")
		os.Exit(1)
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from, fmt.Sprintf("(fron=%s,account=%d)", from, acc)}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to, fmt.Sprintf("(to=%s,amount=%d)", to, amount)})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from, fmt.Sprintf("(formular=%d-%d,from=%s)", acc, amount, from)}) // a change
	}

	msg := fmt.Sprintf("%s sends %d to %s", from, amount, to)
	tx := Transaction{nil, inputs, outputs, msg}
	//inJson, _ := json.Marshal(inputs)
	//outJson, _ := json.Marshal(outputs)
	tx.SetID()

	return &tx
}
