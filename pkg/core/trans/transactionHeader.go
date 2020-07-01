package trans

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

// TransactionHeader contains metadata for a transaction created by the SDK.
type TransactionHeader struct {
	id        fab.TransactionID
	creator   []byte
	nonce     []byte
	channelID string
}

// TransactionID returns the transaction's computed identifier.
func (th *TransactionHeader) TransactionID() fab.TransactionID {
	return th.id
}

// Creator returns the transaction creator's identity bytes.
func (th *TransactionHeader) Creator() []byte {
	return th.creator
}

// Nonce returns the transaction's generated nonce.
func (th *TransactionHeader) Nonce() []byte {
	return th.nonce
}

// ChannelID returns the transaction's target channel identifier.
func (th *TransactionHeader) ChannelID() string {
	return th.channelID
}
