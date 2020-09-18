package valcheck

import (
	"fmt"
	"strconv"

	config "github.com/kaustubhkapatral/monitoring-bot/config"
	types "github.com/kaustubhkapatral/monitoring-bot/types"
	utils "github.com/kaustubhkapatral/monitoring-bot/utils"
)

func HexCheck() (string, error) {
	hex := config.NewApp.Hex
	lcd := config.NewApp.LCD

	blocksUrl := utils.GetBlockHeightURL(lcd)
	blockInfo := &types.Block_response{}
	if err := utils.GetJSON(blocksUrl, blockInfo); err != nil {
		fmt.Println("Unable to query block endpoint", err)
		return "", err
	}

	Precommits := utils.GetValidatorArray(blockInfo)

	for i := 0; i < len(Precommits); i++ {
		if Precommits[i] == hex {
			return "", nil
		}
	}
	return strconv.Itoa(blockInfo.Result.Height), nil
}

func JailCheck() (bool, error) {
	valAddr := config.NewApp.ValAddr
	lcd := config.NewApp.LCD
	validatorUrl := utils.GetValidatorUrl(lcd)
	valInfo := &types.Validators{}
	if err := utils.GetJSON(validatorUrl, valInfo); err != nil {
		fmt.Println("Unable to query validator endpoint", err)
		return false, err
	}
	TotalValidators := utils.GetValidators(valInfo)
	for i := 0; i < len(TotalValidators); i++ {
		if valAddr == TotalValidators[i] {
			return false, nil
		}
	}
	return true, nil
}
