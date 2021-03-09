package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	types "github.com/kaustubhkapatral/monitoring-bot/types"
)

func GetBlockHeightURL(IP string) (url string) {
	url = IP + "/blocks/latest"
	return url
}

func GetJSON(url string, target interface{}) error {
	// client := http.Client{Timeout: time.Second * 2}
	client := http.DefaultClient
	client.Timeout = time.Second * 5
	for i := 0; i < 5; i++ {

		// simple GET request on given URL
		r, err := client.Get(url)
		// r, err := http.Get(url)
		if err != nil {
			time.Sleep(time.Second * 2)
			continue
		} else {
			btes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return err
			}
			if err = json.Unmarshal(btes, &target); err != nil {
				return err
			}
			defer r.Body.Close()
			r.Close = true
			http.DefaultClient.CloseIdleConnections()
			break
		}
	}

	return nil
}

func GetValidatorArray(blockinfo *types.BlockData) []string {
	validatorAddrArray := make([]string, 150)
	PrecommitsArray := blockinfo.Block.LastCommit.Signatures
	for i := 0; i < len(PrecommitsArray); i++ {
		validatorAddrArray = append(validatorAddrArray, PrecommitsArray[i].ValidatorAddress)
	}
	return validatorAddrArray
}

func GetValidatorUrl(IP string, val string) (url string) {
	url = IP + "/cosmos/staking/v1beta1/validators/" + val
	return url
}
