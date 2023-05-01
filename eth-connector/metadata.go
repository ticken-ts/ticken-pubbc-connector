package eth_connector

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"embed"
)


const metadataBuildFilename string = "TickenEvent.json"

//go:embed TickenEvent.json
var metadataFile embed.FS

type rawEthContractMetadata struct {
	ABI []interface{} `json:"abi"`
	Bin string        `json:"bytecode"`
}

func ReadMetadata() (*bind.MetaData, error) {
	// thanks ChatGPT for this =)!
	fileContent, err := metadataFile.ReadFile(metadataBuildFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}

	var rawMetadata rawEthContractMetadata

	if err := json.Unmarshal(fileContent, &rawMetadata); err != nil {
		return nil, fmt.Errorf("failed to load metada: %s", err.Error())
	}

	var stringABIBuffer []string
	for _, abiItem := range rawMetadata.ABI {
		abiItemContent, err := json.Marshal(abiItem)
		if err != nil {
			return nil, fmt.Errorf("failed to load metada: %s", err.Error())
		}
		stringABIBuffer = append(stringABIBuffer, string(abiItemContent))
	}

	return &bind.MetaData{
		ABI: fmt.Sprintf("[%s]", strings.Join(stringABIBuffer[:], ",")),
		Bin: rawMetadata.Bin,
	}, nil
}
