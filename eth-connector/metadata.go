package eth_connector

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"path"
)

const metadataBuildFilename string = "../TickenEvent.json"

type rawEthContractMetadata struct {
	ABI string `json:"abi"`
	Bin string `json:"bytecode"`
}

func ReadMetadata() (*bind.MetaData, error) {
	// get metadata filepath
	// independent  of where it is called
	thisFile, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}
	thisFilePath := path.Dir(thisFile)

	fileContent, err := os.ReadFile(path.Join(thisFilePath, metadataBuildFilename))
	if err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}

	var rawMetadata rawEthContractMetadata

	if err := json.Unmarshal(fileContent, &rawMetadata); err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}

	return &bind.MetaData{ABI: rawMetadata.ABI, Bin: rawMetadata.Bin}, nil
}
