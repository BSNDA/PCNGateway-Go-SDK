package fabric

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-config/protolator"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/pkg/errors"
	"strings"
)

func ConvertToBlock(blockData string) (*common.Block, error) {

	blockBytes, err := base64.StdEncoding.DecodeString(blockData)
	if err != nil {
		return nil, errors.WithMessage(err, "convert block data has error")
	}
	block := &common.Block{}

	err = proto.Unmarshal(blockBytes, block)
	if err != nil {
		return nil, errors.WithMessage(err, "convert block bytes has error")
	}

	return block, nil
}

func ConvertBlockToJson(block *common.Block) (string, error) {

	var sb strings.Builder

	err := protolator.DeepMarshalJSON(&sb, block)
	if err != nil {
		return "", err
	}

	return sb.String(), err

}
