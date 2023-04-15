package eth_connector

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
)

const metadataBuildFilename string = "TickenEvent.json"

type rawEthContractMetadata struct {
	ABI string `json:"abi"`
	Bin string `json:"bin"`
}

func ReadMetadata() (*bind.MetaData, error) {
	fileContent, err := os.ReadFile(metadataBuildFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}

	var rawMetadata rawEthContractMetadata

	if err := json.Unmarshal(fileContent, &rawMetadata); err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}

	return &bind.MetaData{ABI: rawMetadata.ABI, Bin: rawMetadata.Bin}, nil
}
