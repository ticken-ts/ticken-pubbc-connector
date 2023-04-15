package eth_connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const metadataBuildFilename string = "TickenEvent.json"

type rawEthContractMetadata struct {
	ABI []interface{} `json:"abi"`
	Bin string        `json:"bytecode"`
}

func ReadMetadata() (*bind.MetaData, error) {
	// get metadata filepath
	// independent  of where it is called
	_, thisFile, _, _ := runtime.Caller(0)
	thisFilePath := filepath.Dir(thisFile)

	fileContent, err := os.ReadFile(path.Join(thisFilePath, metadataBuildFilename))
	if err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}

	var rawMetadata rawEthContractMetadata

	if err := json.Unmarshal(fileContent, &rawMetadata); err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}

	var stringABIBuffer bytes.Buffer
	for _, abiItem := range rawMetadata.ABI {
		abiItemContent, err := json.Marshal(abiItem)
		if err != nil {
			return nil, fmt.Errorf("failed to load metada: %s", err.Error())
		}
		stringABIBuffer.Write(abiItemContent)
	}

	return &bind.MetaData{ABI: stringABIBuffer.String(), Bin: rawMetadata.Bin}, nil
}
