package main

import (
	"encoding/base64"
	"fmt"

	"github.com/quankori/go-cosmos/pkg/query"
)

type Transfer struct {
	From  string
	To    string
	Value string
}

func base64Decode(str string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}

func main() {
	// query.SubscribeBlock()
	var listEvents []Transfer
	res, _ := query.BlockResults()
	for _, v := range res.TxsResults {
		event := v.GetEvents()
		for _, s := range event {
			if s.GetType() == "transfer" {
				var result Transfer
				result.To, _ = base64Decode(s.GetAttributes()[0].Value)
				result.From, _ = base64Decode(s.GetAttributes()[1].Value)
				result.Value, _ = base64Decode(s.GetAttributes()[2].Value)
				listEvents = append(listEvents, result)
			}
		}
	}

	for _, s := range listEvents {
		fmt.Println(s)
	}
}
