package tools

import (
	"net/http"
)

func SendTxString(IP string, txs string) (string, error) {
	url := "http://" + IP + "/broadcast_tx_commit?tx=\"" + txs + "\""
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
