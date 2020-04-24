package valcheck

import (
	"fmt"

	config "github.com/kaustubhkapatral/monitoring-bot/config"
	types "github.com/kaustubhkapatral/monitoring-bot/types"
	utils "github.com/kaustubhkapatral/monitoring-bot/utils"
)

func HexCheck() (string, error) {
	hex := config.NewApp.Hex
	lcd := config.NewApp.LCD

	blocksUrl := utils.GetBlockHeightURL(lcd)
	blockInfo := &types.BlockData{}
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
	return blockInfo.BlockMeta.Header.Height, nil
}

func JailCheck() (bool, error) {
	valAddr := config.NewApp.ValAddr
	lcd := config.NewApp.LCD
	validatorUrl := utils.GetValidatorUrl(lcd, valAddr)
	valInfo := &types.Validator{}
	if err := utils.GetJSON(validatorUrl, valInfo); err != nil {
		fmt.Println("Unable to query validator endpoint", err)
		return false, err
	}
	return valInfo.Result.Jailed, nil
}
