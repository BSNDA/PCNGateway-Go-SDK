package node

type Block struct {
	Version      int32          `json:"version""`
	Blockid      string         `json:"blockid""`
	PreHash      string         `json:"preHash""`
	Height       int64          `json:"height""`
	Timestamp    int64          `json:"timestamp""`
	Transactions []*Transaction `json:"transactions""`
	TxCount      int32          `json:"txCount""`
	NextHash     string         `json:"nextHash""`
}

type Transaction struct {
	Txid              string           `json:"txId""`
	Blockid           string           `json:"blockId""`
	Version           int32            `json:"version""`
	ContractRequests  []*InvokeRequest `json:"contractRequests""`
	ReceivedTimestamp int64            `json:"receivedTimestamp""`
}

type InvokeRequest struct {
	ContractName string `json:"contractName""`
	MethodName   string `json:"methodName""`
	Args         string `json:"args""`
}
