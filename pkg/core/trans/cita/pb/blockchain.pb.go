// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProofType int32

const (
	ProofType_AuthorityRound ProofType = 0
	ProofType_Raft           ProofType = 1
	ProofType_Bft            ProofType = 2
)

var ProofType_name = map[int32]string{
	0: "AuthorityRound",
	1: "Raft",
	2: "Bft",
}

var ProofType_value = map[string]int32{
	"AuthorityRound": 0,
	"Raft":           1,
	"Bft":            2,
}

func (x ProofType) String() string {
	return proto.EnumName(ProofType_name, int32(x))
}

func (ProofType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{0}
}

type Crypto int32

const (
	Crypto_DEFAULT  Crypto = 0
	Crypto_RESERVED Crypto = 1
)

var Crypto_name = map[int32]string{
	0: "DEFAULT",
	1: "RESERVED",
}

var Crypto_value = map[string]int32{
	"DEFAULT":  0,
	"RESERVED": 1,
}

func (x Crypto) String() string {
	return proto.EnumName(Crypto_name, int32(x))
}

func (Crypto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{1}
}

type Proof struct {
	Content              []byte    `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Type                 ProofType `protobuf:"varint,2,opt,name=type,proto3,enum=pb.ProofType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{0}
}

func (m *Proof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proof.Unmarshal(m, b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return xxx_messageInfo_Proof.Size(m)
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Proof) GetType() ProofType {
	if m != nil {
		return m.Type
	}
	return ProofType_AuthorityRound
}

type BlockHeader struct {
	Prevhash             []byte   `protobuf:"bytes,1,opt,name=prevhash,proto3" json:"prevhash,omitempty"`
	Timestamp            uint64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Height               uint64   `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	StateRoot            []byte   `protobuf:"bytes,4,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	TransactionsRoot     []byte   `protobuf:"bytes,5,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	ReceiptsRoot         []byte   `protobuf:"bytes,6,opt,name=receipts_root,json=receiptsRoot,proto3" json:"receipts_root,omitempty"`
	QuotaUsed            uint64   `protobuf:"varint,7,opt,name=quota_used,json=quotaUsed,proto3" json:"quota_used,omitempty"`
	QuotaLimit           uint64   `protobuf:"varint,8,opt,name=quota_limit,json=quotaLimit,proto3" json:"quota_limit,omitempty"`
	Proof                *Proof   `protobuf:"bytes,9,opt,name=proof,proto3" json:"proof,omitempty"`
	Proposer             []byte   `protobuf:"bytes,10,opt,name=proposer,proto3" json:"proposer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{1}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetPrevhash() []byte {
	if m != nil {
		return m.Prevhash
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockHeader) GetTransactionsRoot() []byte {
	if m != nil {
		return m.TransactionsRoot
	}
	return nil
}

func (m *BlockHeader) GetReceiptsRoot() []byte {
	if m != nil {
		return m.ReceiptsRoot
	}
	return nil
}

func (m *BlockHeader) GetQuotaUsed() uint64 {
	if m != nil {
		return m.QuotaUsed
	}
	return 0
}

func (m *BlockHeader) GetQuotaLimit() uint64 {
	if m != nil {
		return m.QuotaLimit
	}
	return 0
}

func (m *BlockHeader) GetProof() *Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *BlockHeader) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type Status struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{2}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Status) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type AccountGasLimit struct {
	CommonQuotaLimit     uint64            `protobuf:"varint,1,opt,name=common_quota_limit,json=commonQuotaLimit,proto3" json:"common_quota_limit,omitempty"`
	SpecificQuotaLimit   map[string]uint64 `protobuf:"bytes,2,rep,name=specific_quota_limit,json=specificQuotaLimit,proto3" json:"specific_quota_limit,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AccountGasLimit) Reset()         { *m = AccountGasLimit{} }
func (m *AccountGasLimit) String() string { return proto.CompactTextString(m) }
func (*AccountGasLimit) ProtoMessage()    {}
func (*AccountGasLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{3}
}

func (m *AccountGasLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGasLimit.Unmarshal(m, b)
}
func (m *AccountGasLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGasLimit.Marshal(b, m, deterministic)
}
func (m *AccountGasLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGasLimit.Merge(m, src)
}
func (m *AccountGasLimit) XXX_Size() int {
	return xxx_messageInfo_AccountGasLimit.Size(m)
}
func (m *AccountGasLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGasLimit.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGasLimit proto.InternalMessageInfo

func (m *AccountGasLimit) GetCommonQuotaLimit() uint64 {
	if m != nil {
		return m.CommonQuotaLimit
	}
	return 0
}

func (m *AccountGasLimit) GetSpecificQuotaLimit() map[string]uint64 {
	if m != nil {
		return m.SpecificQuotaLimit
	}
	return nil
}

type RichStatus struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Nodes                [][]byte `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Interval             uint64   `protobuf:"varint,4,opt,name=interval,proto3" json:"interval,omitempty"`
	Version              uint32   `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	Validators           [][]byte `protobuf:"bytes,6,rep,name=validators,proto3" json:"validators,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RichStatus) Reset()         { *m = RichStatus{} }
func (m *RichStatus) String() string { return proto.CompactTextString(m) }
func (*RichStatus) ProtoMessage()    {}
func (*RichStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{4}
}

func (m *RichStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RichStatus.Unmarshal(m, b)
}
func (m *RichStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RichStatus.Marshal(b, m, deterministic)
}
func (m *RichStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RichStatus.Merge(m, src)
}
func (m *RichStatus) XXX_Size() int {
	return xxx_messageInfo_RichStatus.Size(m)
}
func (m *RichStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RichStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RichStatus proto.InternalMessageInfo

func (m *RichStatus) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *RichStatus) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RichStatus) GetNodes() [][]byte {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *RichStatus) GetInterval() uint64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *RichStatus) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RichStatus) GetValidators() [][]byte {
	if m != nil {
		return m.Validators
	}
	return nil
}

type Transaction struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Quota                uint64   `protobuf:"varint,3,opt,name=quota,proto3" json:"quota,omitempty"`
	ValidUntilBlock      uint64   `protobuf:"varint,4,opt,name=valid_until_block,json=validUntilBlock,proto3" json:"valid_until_block,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Value                []byte   `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	ChainId              uint32   `protobuf:"varint,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Version              uint32   `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	ToV1                 []byte   `protobuf:"bytes,9,opt,name=to_v1,json=toV1,proto3" json:"to_v1,omitempty"`
	ChainIdV1            []byte   `protobuf:"bytes,10,opt,name=chain_id_v1,json=chainIdV1,proto3" json:"chain_id_v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{5}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Transaction) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *Transaction) GetQuota() uint64 {
	if m != nil {
		return m.Quota
	}
	return 0
}

func (m *Transaction) GetValidUntilBlock() uint64 {
	if m != nil {
		return m.ValidUntilBlock
	}
	return 0
}

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Transaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Transaction) GetToV1() []byte {
	if m != nil {
		return m.ToV1
	}
	return nil
}

func (m *Transaction) GetChainIdV1() []byte {
	if m != nil {
		return m.ChainIdV1
	}
	return nil
}

type UnverifiedTransaction struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Signature            []byte       `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Crypto               Crypto       `protobuf:"varint,3,opt,name=crypto,proto3,enum=pb.Crypto" json:"crypto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UnverifiedTransaction) Reset()         { *m = UnverifiedTransaction{} }
func (m *UnverifiedTransaction) String() string { return proto.CompactTextString(m) }
func (*UnverifiedTransaction) ProtoMessage()    {}
func (*UnverifiedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{6}
}

func (m *UnverifiedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnverifiedTransaction.Unmarshal(m, b)
}
func (m *UnverifiedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnverifiedTransaction.Marshal(b, m, deterministic)
}
func (m *UnverifiedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnverifiedTransaction.Merge(m, src)
}
func (m *UnverifiedTransaction) XXX_Size() int {
	return xxx_messageInfo_UnverifiedTransaction.Size(m)
}
func (m *UnverifiedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_UnverifiedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_UnverifiedTransaction proto.InternalMessageInfo

func (m *UnverifiedTransaction) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *UnverifiedTransaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *UnverifiedTransaction) GetCrypto() Crypto {
	if m != nil {
		return m.Crypto
	}
	return Crypto_DEFAULT
}

type SignedTransaction struct {
	TransactionWithSig *UnverifiedTransaction `protobuf:"bytes,1,opt,name=transaction_with_sig,json=transactionWithSig,proto3" json:"transaction_with_sig,omitempty"`
	// SignedTransaction hash
	TxHash []byte `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	// public key
	Signer               []byte   `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedTransaction) Reset()         { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()    {}
func (*SignedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{7}
}

func (m *SignedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTransaction.Unmarshal(m, b)
}
func (m *SignedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTransaction.Marshal(b, m, deterministic)
}
func (m *SignedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTransaction.Merge(m, src)
}
func (m *SignedTransaction) XXX_Size() int {
	return xxx_messageInfo_SignedTransaction.Size(m)
}
func (m *SignedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTransaction proto.InternalMessageInfo

func (m *SignedTransaction) GetTransactionWithSig() *UnverifiedTransaction {
	if m != nil {
		return m.TransactionWithSig
	}
	return nil
}

func (m *SignedTransaction) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *SignedTransaction) GetSigner() []byte {
	if m != nil {
		return m.Signer
	}
	return nil
}

type BlockBody struct {
	Transactions         []*SignedTransaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockBody) Reset()         { *m = BlockBody{} }
func (m *BlockBody) String() string { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()    {}
func (*BlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{8}
}

func (m *BlockBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockBody.Unmarshal(m, b)
}
func (m *BlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockBody.Marshal(b, m, deterministic)
}
func (m *BlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockBody.Merge(m, src)
}
func (m *BlockBody) XXX_Size() int {
	return xxx_messageInfo_BlockBody.Size(m)
}
func (m *BlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_BlockBody proto.InternalMessageInfo

func (m *BlockBody) GetTransactions() []*SignedTransaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type CompactBlockBody struct {
	TxHashes             [][]byte `protobuf:"bytes,1,rep,name=tx_hashes,json=txHashes,proto3" json:"tx_hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompactBlockBody) Reset()         { *m = CompactBlockBody{} }
func (m *CompactBlockBody) String() string { return proto.CompactTextString(m) }
func (*CompactBlockBody) ProtoMessage()    {}
func (*CompactBlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{9}
}

func (m *CompactBlockBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompactBlockBody.Unmarshal(m, b)
}
func (m *CompactBlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompactBlockBody.Marshal(b, m, deterministic)
}
func (m *CompactBlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactBlockBody.Merge(m, src)
}
func (m *CompactBlockBody) XXX_Size() int {
	return xxx_messageInfo_CompactBlockBody.Size(m)
}
func (m *CompactBlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactBlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_CompactBlockBody proto.InternalMessageInfo

func (m *CompactBlockBody) GetTxHashes() [][]byte {
	if m != nil {
		return m.TxHashes
	}
	return nil
}

type Block struct {
	Version              uint32       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Header               *BlockHeader `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body                 *BlockBody   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{10}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type CompactBlock struct {
	Version              uint32            `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Header               *BlockHeader      `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body                 *CompactBlockBody `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CompactBlock) Reset()         { *m = CompactBlock{} }
func (m *CompactBlock) String() string { return proto.CompactTextString(m) }
func (*CompactBlock) ProtoMessage()    {}
func (*CompactBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{11}
}

func (m *CompactBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompactBlock.Unmarshal(m, b)
}
func (m *CompactBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompactBlock.Marshal(b, m, deterministic)
}
func (m *CompactBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactBlock.Merge(m, src)
}
func (m *CompactBlock) XXX_Size() int {
	return xxx_messageInfo_CompactBlock.Size(m)
}
func (m *CompactBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactBlock.DiscardUnknown(m)
}

var xxx_messageInfo_CompactBlock proto.InternalMessageInfo

func (m *CompactBlock) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CompactBlock) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CompactBlock) GetBody() *CompactBlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type BlockWithProof struct {
	Blk                  *Block   `protobuf:"bytes,1,opt,name=blk,proto3" json:"blk,omitempty"`
	Proof                *Proof   `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockWithProof) Reset()         { *m = BlockWithProof{} }
func (m *BlockWithProof) String() string { return proto.CompactTextString(m) }
func (*BlockWithProof) ProtoMessage()    {}
func (*BlockWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{12}
}

func (m *BlockWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockWithProof.Unmarshal(m, b)
}
func (m *BlockWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockWithProof.Marshal(b, m, deterministic)
}
func (m *BlockWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockWithProof.Merge(m, src)
}
func (m *BlockWithProof) XXX_Size() int {
	return xxx_messageInfo_BlockWithProof.Size(m)
}
func (m *BlockWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_BlockWithProof proto.InternalMessageInfo

func (m *BlockWithProof) GetBlk() *Block {
	if m != nil {
		return m.Blk
	}
	return nil
}

func (m *BlockWithProof) GetProof() *Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type BlockTxs struct {
	Height               uint64     `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Body                 *BlockBody `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BlockTxs) Reset()         { *m = BlockTxs{} }
func (m *BlockTxs) String() string { return proto.CompactTextString(m) }
func (*BlockTxs) ProtoMessage()    {}
func (*BlockTxs) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{13}
}

func (m *BlockTxs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTxs.Unmarshal(m, b)
}
func (m *BlockTxs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTxs.Marshal(b, m, deterministic)
}
func (m *BlockTxs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTxs.Merge(m, src)
}
func (m *BlockTxs) XXX_Size() int {
	return xxx_messageInfo_BlockTxs.Size(m)
}
func (m *BlockTxs) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTxs.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTxs proto.InternalMessageInfo

func (m *BlockTxs) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockTxs) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type BlackList struct {
	// black list of address, the account that sent the transaction does not have enough gas
	BlackList [][]byte `protobuf:"bytes,1,rep,name=black_list,json=blackList,proto3" json:"black_list,omitempty"`
	// clear list of address
	ClearList            [][]byte `protobuf:"bytes,2,rep,name=clear_list,json=clearList,proto3" json:"clear_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlackList) Reset()         { *m = BlackList{} }
func (m *BlackList) String() string { return proto.CompactTextString(m) }
func (*BlackList) ProtoMessage()    {}
func (*BlackList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{14}
}

func (m *BlackList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlackList.Unmarshal(m, b)
}
func (m *BlackList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlackList.Marshal(b, m, deterministic)
}
func (m *BlackList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlackList.Merge(m, src)
}
func (m *BlackList) XXX_Size() int {
	return xxx_messageInfo_BlackList.Size(m)
}
func (m *BlackList) XXX_DiscardUnknown() {
	xxx_messageInfo_BlackList.DiscardUnknown(m)
}

var xxx_messageInfo_BlackList proto.InternalMessageInfo

func (m *BlackList) GetBlackList() [][]byte {
	if m != nil {
		return m.BlackList
	}
	return nil
}

func (m *BlackList) GetClearList() [][]byte {
	if m != nil {
		return m.ClearList
	}
	return nil
}

// State positioning signal
type StateSignal struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateSignal) Reset()         { *m = StateSignal{} }
func (m *StateSignal) String() string { return proto.CompactTextString(m) }
func (*StateSignal) ProtoMessage()    {}
func (*StateSignal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{15}
}

func (m *StateSignal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateSignal.Unmarshal(m, b)
}
func (m *StateSignal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateSignal.Marshal(b, m, deterministic)
}
func (m *StateSignal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateSignal.Merge(m, src)
}
func (m *StateSignal) XXX_Size() int {
	return xxx_messageInfo_StateSignal.Size(m)
}
func (m *StateSignal) XXX_DiscardUnknown() {
	xxx_messageInfo_StateSignal.DiscardUnknown(m)
}

var xxx_messageInfo_StateSignal proto.InternalMessageInfo

func (m *StateSignal) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.ProofType", ProofType_name, ProofType_value)
	proto.RegisterEnum("pb.Crypto", Crypto_name, Crypto_value)
	proto.RegisterType((*Proof)(nil), "pb.Proof")
	proto.RegisterType((*BlockHeader)(nil), "pb.BlockHeader")
	proto.RegisterType((*Status)(nil), "pb.Status")
	proto.RegisterType((*AccountGasLimit)(nil), "pb.AccountGasLimit")
	proto.RegisterMapType((map[string]uint64)(nil), "pb.AccountGasLimit.SpecificQuotaLimitEntry")
	proto.RegisterType((*RichStatus)(nil), "pb.RichStatus")
	proto.RegisterType((*Transaction)(nil), "pb.Transaction")
	proto.RegisterType((*UnverifiedTransaction)(nil), "pb.UnverifiedTransaction")
	proto.RegisterType((*SignedTransaction)(nil), "pb.SignedTransaction")
	proto.RegisterType((*BlockBody)(nil), "pb.BlockBody")
	proto.RegisterType((*CompactBlockBody)(nil), "pb.CompactBlockBody")
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterType((*CompactBlock)(nil), "pb.CompactBlock")
	proto.RegisterType((*BlockWithProof)(nil), "pb.BlockWithProof")
	proto.RegisterType((*BlockTxs)(nil), "pb.BlockTxs")
	proto.RegisterType((*BlackList)(nil), "pb.BlackList")
	proto.RegisterType((*StateSignal)(nil), "pb.StateSignal")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_e9ac6287ce250c9a) }

var fileDescriptor_e9ac6287ce250c9a = []byte{
	// 1001 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x6d, 0x6f, 0x1b, 0x45,
	0x10, 0xee, 0x9d, 0x5f, 0x6f, 0xce, 0x49, 0x9c, 0x25, 0xa5, 0xd7, 0x96, 0xb6, 0xe1, 0x2a, 0x84,
	0x95, 0xa2, 0xa0, 0x18, 0x3e, 0x00, 0xdf, 0x92, 0xc6, 0xa5, 0x15, 0x11, 0x82, 0xcd, 0x0b, 0x9f,
	0xd0, 0x69, 0x7d, 0xb7, 0xb1, 0x57, 0xb1, 0x6f, 0x8f, 0xdb, 0xb5, 0x89, 0xc5, 0x2f, 0x40, 0xe2,
	0x0b, 0xff, 0x81, 0xbf, 0xc3, 0xcf, 0xe0, 0x7f, 0xa0, 0x9d, 0x5d, 0xdb, 0xe7, 0xbe, 0x48, 0x20,
	0xf5, 0x9b, 0x67, 0x9e, 0xb9, 0xd9, 0x67, 0x9e, 0x99, 0x9d, 0x35, 0x74, 0x87, 0x13, 0x99, 0xde,
	0xa4, 0x63, 0x26, 0xf2, 0xc3, 0xa2, 0x94, 0x5a, 0x12, 0xbf, 0x18, 0xc6, 0xa7, 0xd0, 0xf8, 0xa1,
	0x94, 0xf2, 0x9a, 0x44, 0xd0, 0x4a, 0x65, 0xae, 0x79, 0xae, 0x23, 0x6f, 0xdf, 0xeb, 0x75, 0xe8,
	0xd2, 0x24, 0x1f, 0x43, 0x5d, 0x2f, 0x0a, 0x1e, 0xf9, 0xfb, 0x5e, 0x6f, 0xbb, 0xbf, 0x75, 0x58,
	0x0c, 0x0f, 0xf1, 0x93, 0x8b, 0x45, 0xc1, 0x29, 0x42, 0xf1, 0xdf, 0x3e, 0x84, 0x27, 0x26, 0xfd,
	0x4b, 0xce, 0x32, 0x5e, 0x92, 0x07, 0xd0, 0x2e, 0x4a, 0x3e, 0x1f, 0x33, 0x35, 0x76, 0xd9, 0x56,
	0x36, 0xf9, 0x08, 0x02, 0x2d, 0xa6, 0x5c, 0x69, 0x36, 0x2d, 0x30, 0x67, 0x9d, 0xae, 0x1d, 0xe4,
	0x43, 0x68, 0x8e, 0xb9, 0x18, 0x8d, 0x75, 0x54, 0x43, 0xc8, 0x59, 0xe4, 0x11, 0x80, 0xd2, 0x4c,
	0xf3, 0xa4, 0x94, 0x52, 0x47, 0x75, 0xcc, 0x19, 0xa0, 0x87, 0x4a, 0xa9, 0xc9, 0x33, 0xd8, 0xd5,
	0x25, 0xcb, 0x15, 0x4b, 0xb5, 0x90, 0xb9, 0xb2, 0x51, 0x0d, 0x8c, 0xea, 0x56, 0x01, 0x0c, 0x7e,
	0x0a, 0x5b, 0x25, 0x4f, 0xb9, 0x28, 0xb4, 0x0b, 0x6c, 0x62, 0x60, 0x67, 0xe9, 0xc4, 0xa0, 0x47,
	0x00, 0xbf, 0xcc, 0xa4, 0x66, 0xc9, 0x4c, 0xf1, 0x2c, 0x6a, 0x59, 0x9e, 0xe8, 0xb9, 0x54, 0x3c,
	0x23, 0x4f, 0x20, 0xb4, 0xf0, 0x44, 0x4c, 0x85, 0x8e, 0xda, 0x88, 0xdb, 0x2f, 0xce, 0x8c, 0x87,
	0x3c, 0x81, 0x46, 0x61, 0x54, 0x8a, 0x82, 0x7d, 0xaf, 0x17, 0xf6, 0x83, 0x95, 0x6c, 0xd4, 0xfa,
	0xad, 0x46, 0xb2, 0x90, 0x8a, 0x97, 0x11, 0x2c, 0x35, 0xb2, 0x76, 0xfc, 0x25, 0x34, 0xcf, 0x35,
	0xd3, 0x33, 0x45, 0x08, 0xd4, 0x2b, 0x2a, 0xe2, 0xef, 0x8a, 0x46, 0x7e, 0x55, 0xa3, 0xf8, 0x1f,
	0x0f, 0x76, 0x8e, 0xd3, 0x54, 0xce, 0x72, 0xfd, 0x2d, 0x53, 0x96, 0xc6, 0x67, 0x40, 0x52, 0x39,
	0x9d, 0xca, 0x3c, 0xa9, 0xd2, 0xf5, 0xf0, 0xbb, 0xae, 0x45, 0x7e, 0x5c, 0x93, 0xfe, 0x19, 0xf6,
	0x54, 0xc1, 0x53, 0x71, 0x2d, 0xd2, 0x8d, 0x78, 0x7f, 0xbf, 0xd6, 0x0b, 0xfb, 0xcf, 0x4c, 0x0d,
	0xaf, 0x1d, 0x70, 0x78, 0xee, 0xe2, 0xd7, 0x59, 0x06, 0xb9, 0x2e, 0x17, 0x94, 0xa8, 0x37, 0x80,
	0x07, 0x03, 0xb8, 0xf7, 0x8e, 0x70, 0xd2, 0x85, 0xda, 0x0d, 0x5f, 0x20, 0xb1, 0x80, 0x9a, 0x9f,
	0x64, 0x0f, 0x1a, 0x73, 0x36, 0x99, 0x71, 0x57, 0xa4, 0x35, 0xbe, 0xf1, 0xbf, 0xf2, 0xe2, 0xbf,
	0x3c, 0x00, 0x2a, 0xd2, 0xf1, 0xff, 0x97, 0xc8, 0x24, 0xcd, 0x65, 0xc6, 0x55, 0x54, 0xdb, 0xaf,
	0xf5, 0x3a, 0xd4, 0x1a, 0xa6, 0x15, 0x22, 0xd7, 0xbc, 0x9c, 0xb3, 0x09, 0x8e, 0x56, 0x9d, 0xae,
	0x6c, 0x73, 0x2f, 0xe6, 0xbc, 0x54, 0x42, 0xe6, 0x38, 0x4f, 0x5b, 0x74, 0x69, 0x92, 0xc7, 0x00,
	0x73, 0x36, 0x11, 0x19, 0xd3, 0xb2, 0x54, 0x51, 0x13, 0x13, 0x56, 0x3c, 0xf1, 0xef, 0x3e, 0x84,
	0x17, 0xeb, 0xd9, 0x23, 0xdb, 0xe0, 0x6b, 0xe9, 0x2a, 0xf4, 0xb5, 0xb4, 0x5c, 0xf2, 0xd4, 0x16,
	0x18, 0x50, 0x6b, 0x18, 0x2f, 0x2a, 0xef, 0xe6, 0xdf, 0x1a, 0xe4, 0x00, 0x76, 0x31, 0x73, 0x32,
	0xcb, 0xb5, 0x98, 0x24, 0x78, 0x95, 0x1d, 0xd5, 0x1d, 0x04, 0x2e, 0x8d, 0x1f, 0xaf, 0xa0, 0xd1,
	0x23, 0x63, 0x9a, 0xb9, 0xf1, 0xc7, 0xdf, 0x6b, 0x31, 0xed, 0xa8, 0x5b, 0x83, 0xdc, 0x87, 0x36,
	0xee, 0x83, 0x44, 0xd8, 0x09, 0xdf, 0xa2, 0x2d, 0xb4, 0x5f, 0x65, 0xd5, 0xb2, 0xdb, 0x9b, 0x65,
	0x7f, 0x00, 0x0d, 0x2d, 0x93, 0xf9, 0x11, 0x0e, 0x76, 0x87, 0xd6, 0xb5, 0xbc, 0x3a, 0x22, 0x8f,
	0x21, 0x5c, 0x66, 0x32, 0x90, 0x9d, 0xe7, 0xc0, 0x25, 0xbb, 0x3a, 0x8a, 0xff, 0xf0, 0xe0, 0xee,
	0x65, 0x3e, 0xe7, 0xa5, 0xb8, 0x16, 0x3c, 0xab, 0xaa, 0x72, 0x04, 0x61, 0xe5, 0x82, 0xa2, 0x3c,
	0x61, 0x7f, 0xc7, 0x4c, 0x5a, 0x25, 0x8a, 0x56, 0x63, 0xcc, 0x06, 0x51, 0x62, 0x94, 0x33, 0x3d,
	0x2b, 0xad, 0x78, 0x66, 0x15, 0x2c, 0x1d, 0x24, 0x86, 0x66, 0x5a, 0x2e, 0x0a, 0x2d, 0x51, 0xc1,
	0xed, 0x3e, 0x98, 0x5c, 0xcf, 0xd1, 0x43, 0x1d, 0x12, 0xff, 0xe9, 0xc1, 0xee, 0xb9, 0x18, 0xe5,
	0x9b, 0x54, 0xbe, 0x83, 0xbd, 0xca, 0x31, 0xc9, 0xaf, 0x42, 0x8f, 0x13, 0x25, 0x46, 0x8e, 0xd3,
	0x7d, 0x93, 0xe7, 0xad, 0x35, 0x50, 0x52, 0xf9, 0xec, 0x27, 0xa1, 0xc7, 0xe7, 0x62, 0x44, 0xee,
	0x41, 0x4b, 0xdf, 0x26, 0x38, 0x98, 0x96, 0x62, 0x53, 0xdf, 0xbe, 0x74, 0xa3, 0x69, 0xc8, 0xf2,
	0x12, 0xf9, 0x75, 0xa8, 0xb3, 0xe2, 0x17, 0x10, 0x60, 0xff, 0x4e, 0x64, 0xb6, 0x20, 0x5f, 0x43,
	0xa7, 0xba, 0xb6, 0x22, 0x0f, 0x2f, 0xe0, 0x5d, 0x43, 0xe1, 0x0d, 0xde, 0x74, 0x23, 0x34, 0xfe,
	0x1c, 0xba, 0xcf, 0xe5, 0xb4, 0x60, 0xa9, 0x5e, 0xa7, 0x7b, 0x08, 0x81, 0x23, 0xc3, 0x6d, 0xae,
	0x0e, 0x6d, 0x5b, 0x3a, 0x5c, 0xc5, 0x53, 0x68, 0xd8, 0xc1, 0xa9, 0xf4, 0xdc, 0xdb, 0xec, 0xf9,
	0xa7, 0xe6, 0x3a, 0x99, 0xcd, 0x8e, 0xb5, 0xb8, 0xfe, 0x54, 0x16, 0x3e, 0x75, 0xb0, 0x79, 0x2b,
	0x86, 0x32, 0x5b, 0x60, 0x69, 0xa1, 0x7d, 0x2b, 0x56, 0x2c, 0x28, 0x42, 0xf1, 0x6f, 0xd0, 0xa9,
	0xf2, 0x7b, 0x1f, 0xa7, 0xf6, 0x36, 0x4e, 0xdd, 0xc3, 0x86, 0xbf, 0x26, 0x81, 0x3b, 0xfc, 0x7b,
	0xd8, 0x46, 0x97, 0xe9, 0x92, 0x7d, 0xf7, 0x1e, 0x42, 0x6d, 0x38, 0xb9, 0x71, 0x3d, 0x0e, 0x56,
	0x27, 0x50, 0xe3, 0x5d, 0x2f, 0x71, 0xff, 0xed, 0x4b, 0x3c, 0x1e, 0x40, 0x1b, 0xc3, 0x2f, 0x6e,
	0x55, 0x65, 0xe7, 0x78, 0x1b, 0x3b, 0xe7, 0x3f, 0x68, 0xf2, 0xca, 0xf4, 0x9e, 0xa5, 0x37, 0x67,
	0x42, 0xe1, 0xcb, 0x33, 0x34, 0x46, 0x32, 0x11, 0x4a, 0xbb, 0x6e, 0x05, 0xc3, 0x2a, 0x9c, 0x4e,
	0x38, 0x2b, 0x2d, 0xec, 0x5b, 0x18, 0x3d, 0x06, 0x8e, 0x3f, 0x81, 0xd0, 0xec, 0x45, 0x6e, 0xc6,
	0x84, 0x4d, 0xde, 0x45, 0xea, 0xa0, 0x0f, 0xc1, 0xea, 0x11, 0x27, 0x04, 0xb6, 0x8f, 0x67, 0x7a,
	0x2c, 0x4b, 0xa1, 0x17, 0x54, 0xce, 0xf2, 0xac, 0x7b, 0x87, 0xb4, 0xa1, 0x4e, 0xd9, 0xb5, 0xee,
	0x7a, 0xa4, 0x05, 0xb5, 0x93, 0x6b, 0xdd, 0xf5, 0x0f, 0x9e, 0x42, 0xd3, 0xde, 0x23, 0x12, 0x42,
	0xeb, 0x74, 0xf0, 0xe2, 0xf8, 0xf2, 0xec, 0xa2, 0x7b, 0x87, 0x74, 0xa0, 0x4d, 0x07, 0xe7, 0x03,
	0x7a, 0x35, 0x38, 0xed, 0x7a, 0xc3, 0x26, 0xfe, 0xb7, 0xf8, 0xe2, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x78, 0x1a, 0xec, 0xb1, 0x6f, 0x08, 0x00, 0x00,
}
