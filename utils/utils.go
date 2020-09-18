package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	types "github.com/kaustubhkapatral/monitoring-bot/types"
)

func GetBlockHeightURL(IP string) (url string) {
	url = IP + "/consensus/blocklastcommit?name=Oasis_Local"
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

func GetValidatorArray(blockinfo *types.Block_response) []string {
	validatorAddrArray := make([]string, 150)
	PrecommitsArray := blockinfo.Result.Signatures
	for i := 0; i < len(PrecommitsArray); i++ {
		validatorAddrArray = append(validatorAddrArray, PrecommitsArray[i].ValidatorAddress)
	}
	return validatorAddrArray
}

func GetValidatorUrl(IP string) (url string) {
	url = IP + "/scheduler/validators?name=Oasis_Local"
	return url
}

func GetValidators(valinfo *types.Validators) []string {
	validatorArray := make([]string, 150)
	IdArray := valinfo.Result

	for i := 0; i < len(IdArray); i++ {
		validatorArray = append(validatorArray, IdArray[i].ID)
	}
	return validatorArray
}
