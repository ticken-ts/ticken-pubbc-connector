package eth_connector

import (
	"testing"
)

func Test_ReadMetadata(t *testing.T) {
	metadata, err := ReadMetadata()
	if err != nil {
		t.Errorf(err.Error())
	}
	if metadata == nil {
		t.Errorf("metadata is nil")
	}

	if len(metadata.Bin) == 0 {
		t.Errorf("bin value of metadata is not set")
	}
	if len(metadata.ABI) == 0 {
		t.Errorf("ABI value of metadata is not set")
	}
}
