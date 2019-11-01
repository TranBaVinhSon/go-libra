// +build js
// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: transaction.proto

package pbtypes

import jspb "github.com/johanbrandhorst/protobuf/jspb"
import google_protobuf1 "github.com/johanbrandhorst/protobuf/ptypes/wrappers"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type TransactionArgument_ArgType int

const (
	TransactionArgument_U64       TransactionArgument_ArgType = 0
	TransactionArgument_ADDRESS   TransactionArgument_ArgType = 1
	TransactionArgument_STRING    TransactionArgument_ArgType = 2
	TransactionArgument_BYTEARRAY TransactionArgument_ArgType = 3
)

var TransactionArgument_ArgType_name = map[int]string{
	0: "U64",
	1: "ADDRESS",
	2: "STRING",
	3: "BYTEARRAY",
}
var TransactionArgument_ArgType_value = map[string]int{
	"U64":       0,
	"ADDRESS":   1,
	"STRING":    2,
	"BYTEARRAY": 3,
}

func (x TransactionArgument_ArgType) String() string {
	return TransactionArgument_ArgType_name[int(x)]
}

// An argument to the transaction if the transaction takes arguments
type TransactionArgument struct {
}

// MarshalToWriter marshals TransactionArgument to the provided writer.
func (m *TransactionArgument) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	return
}

// Marshal marshals TransactionArgument to a slice of bytes.
func (m *TransactionArgument) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a TransactionArgument from the provided reader.
func (m *TransactionArgument) UnmarshalFromReader(reader jspb.Reader) *TransactionArgument {
	for reader.Next() {
		if m == nil {
			m = &TransactionArgument{}
		}

		switch reader.GetFieldNumber() {
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a TransactionArgument from a slice of bytes.
func (m *TransactionArgument) Unmarshal(rawBytes []byte) (*TransactionArgument, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// A generic structure that represents signed RawTransaction
type SignedTransaction struct {
	// LCS bytes representation of a SignedTransaction.
	TxnBytes []byte
}

// GetTxnBytes gets the TxnBytes of the SignedTransaction.
func (m *SignedTransaction) GetTxnBytes() (x []byte) {
	if m == nil {
		return x
	}
	return m.TxnBytes
}

// MarshalToWriter marshals SignedTransaction to the provided writer.
func (m *SignedTransaction) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.TxnBytes) > 0 {
		writer.WriteBytes(5, m.TxnBytes)
	}

	return
}

// Marshal marshals SignedTransaction to a slice of bytes.
func (m *SignedTransaction) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a SignedTransaction from the provided reader.
func (m *SignedTransaction) UnmarshalFromReader(reader jspb.Reader) *SignedTransaction {
	for reader.Next() {
		if m == nil {
			m = &SignedTransaction{}
		}

		switch reader.GetFieldNumber() {
		case 5:
			m.TxnBytes = reader.ReadBytes()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a SignedTransaction from a slice of bytes.
func (m *SignedTransaction) Unmarshal(rawBytes []byte) (*SignedTransaction, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// A generic structure that represents a transaction, covering all possible
// variants.
type Transaction struct {
	Transaction []byte
}

// GetTransaction gets the Transaction of the Transaction.
func (m *Transaction) GetTransaction() (x []byte) {
	if m == nil {
		return x
	}
	return m.Transaction
}

// MarshalToWriter marshals Transaction to the provided writer.
func (m *Transaction) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Transaction) > 0 {
		writer.WriteBytes(1, m.Transaction)
	}

	return
}

// Marshal marshals Transaction to a slice of bytes.
func (m *Transaction) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Transaction from the provided reader.
func (m *Transaction) UnmarshalFromReader(reader jspb.Reader) *Transaction {
	for reader.Next() {
		if m == nil {
			m = &Transaction{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Transaction = reader.ReadBytes()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Transaction from a slice of bytes.
func (m *Transaction) Unmarshal(rawBytes []byte) (*Transaction, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type TransactionWithProof struct {
	// The version of the returned signed transaction.
	Version uint64
	// The transaction itself.
	Transaction *Transaction
	// The proof authenticating the transaction.
	Proof *TransactionProof
	// The events yielded by executing the transaction, if requested.
	Events *EventsList
}

// GetVersion gets the Version of the TransactionWithProof.
func (m *TransactionWithProof) GetVersion() (x uint64) {
	if m == nil {
		return x
	}
	return m.Version
}

// GetTransaction gets the Transaction of the TransactionWithProof.
func (m *TransactionWithProof) GetTransaction() (x *Transaction) {
	if m == nil {
		return x
	}
	return m.Transaction
}

// GetProof gets the Proof of the TransactionWithProof.
func (m *TransactionWithProof) GetProof() (x *TransactionProof) {
	if m == nil {
		return x
	}
	return m.Proof
}

// GetEvents gets the Events of the TransactionWithProof.
func (m *TransactionWithProof) GetEvents() (x *EventsList) {
	if m == nil {
		return x
	}
	return m.Events
}

// MarshalToWriter marshals TransactionWithProof to the provided writer.
func (m *TransactionWithProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Version != 0 {
		writer.WriteUint64(1, m.Version)
	}

	if m.Transaction != nil {
		writer.WriteMessage(2, func() {
			m.Transaction.MarshalToWriter(writer)
		})
	}

	if m.Proof != nil {
		writer.WriteMessage(3, func() {
			m.Proof.MarshalToWriter(writer)
		})
	}

	if m.Events != nil {
		writer.WriteMessage(4, func() {
			m.Events.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals TransactionWithProof to a slice of bytes.
func (m *TransactionWithProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a TransactionWithProof from the provided reader.
func (m *TransactionWithProof) UnmarshalFromReader(reader jspb.Reader) *TransactionWithProof {
	for reader.Next() {
		if m == nil {
			m = &TransactionWithProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Version = reader.ReadUint64()
		case 2:
			reader.ReadMessage(func() {
				m.Transaction = m.Transaction.UnmarshalFromReader(reader)
			})
		case 3:
			reader.ReadMessage(func() {
				m.Proof = m.Proof.UnmarshalFromReader(reader)
			})
		case 4:
			reader.ReadMessage(func() {
				m.Events = m.Events.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a TransactionWithProof from a slice of bytes.
func (m *TransactionWithProof) Unmarshal(rawBytes []byte) (*TransactionWithProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// A generic structure that represents a block of transactions originated from a
// particular validator instance.
type SignedTransactionsBlock struct {
	// Set of Signed Transactions
	Transactions []*SignedTransaction
	// Public key of the validator that created this block
	ValidatorPublicKey []byte
	// Signature of the validator that created this block
	ValidatorSignature []byte
}

// GetTransactions gets the Transactions of the SignedTransactionsBlock.
func (m *SignedTransactionsBlock) GetTransactions() (x []*SignedTransaction) {
	if m == nil {
		return x
	}
	return m.Transactions
}

// GetValidatorPublicKey gets the ValidatorPublicKey of the SignedTransactionsBlock.
func (m *SignedTransactionsBlock) GetValidatorPublicKey() (x []byte) {
	if m == nil {
		return x
	}
	return m.ValidatorPublicKey
}

// GetValidatorSignature gets the ValidatorSignature of the SignedTransactionsBlock.
func (m *SignedTransactionsBlock) GetValidatorSignature() (x []byte) {
	if m == nil {
		return x
	}
	return m.ValidatorSignature
}

// MarshalToWriter marshals SignedTransactionsBlock to the provided writer.
func (m *SignedTransactionsBlock) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, msg := range m.Transactions {
		writer.WriteMessage(1, func() {
			msg.MarshalToWriter(writer)
		})
	}

	if len(m.ValidatorPublicKey) > 0 {
		writer.WriteBytes(2, m.ValidatorPublicKey)
	}

	if len(m.ValidatorSignature) > 0 {
		writer.WriteBytes(3, m.ValidatorSignature)
	}

	return
}

// Marshal marshals SignedTransactionsBlock to a slice of bytes.
func (m *SignedTransactionsBlock) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a SignedTransactionsBlock from the provided reader.
func (m *SignedTransactionsBlock) UnmarshalFromReader(reader jspb.Reader) *SignedTransactionsBlock {
	for reader.Next() {
		if m == nil {
			m = &SignedTransactionsBlock{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.Transactions = append(m.Transactions, new(SignedTransaction).UnmarshalFromReader(reader))
			})
		case 2:
			m.ValidatorPublicKey = reader.ReadBytes()
		case 3:
			m.ValidatorSignature = reader.ReadBytes()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a SignedTransactionsBlock from a slice of bytes.
func (m *SignedTransactionsBlock) Unmarshal(rawBytes []byte) (*SignedTransactionsBlock, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Account state as a whole.
// After execution, updates to accounts are passed in this form to storage for
// persistence.
type AccountState struct {
	// Account address
	Address []byte
	// Account state blob
	Blob []byte
}

// GetAddress gets the Address of the AccountState.
func (m *AccountState) GetAddress() (x []byte) {
	if m == nil {
		return x
	}
	return m.Address
}

// GetBlob gets the Blob of the AccountState.
func (m *AccountState) GetBlob() (x []byte) {
	if m == nil {
		return x
	}
	return m.Blob
}

// MarshalToWriter marshals AccountState to the provided writer.
func (m *AccountState) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Address) > 0 {
		writer.WriteBytes(1, m.Address)
	}

	if len(m.Blob) > 0 {
		writer.WriteBytes(2, m.Blob)
	}

	return
}

// Marshal marshals AccountState to a slice of bytes.
func (m *AccountState) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a AccountState from the provided reader.
func (m *AccountState) UnmarshalFromReader(reader jspb.Reader) *AccountState {
	for reader.Next() {
		if m == nil {
			m = &AccountState{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Address = reader.ReadBytes()
		case 2:
			m.Blob = reader.ReadBytes()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a AccountState from a slice of bytes.
func (m *AccountState) Unmarshal(rawBytes []byte) (*AccountState, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Transaction struct to commit to storage
type TransactionToCommit struct {
	// The signed transaction which was executed
	Transaction *Transaction
	// State db updates
	AccountStates []*AccountState
	// Events yielded by the transaction.
	Events []*Event
	// The amount of gas used.
	GasUsed uint64
	// The major status of executing the transaction.
	MajorStatus uint64
}

// GetTransaction gets the Transaction of the TransactionToCommit.
func (m *TransactionToCommit) GetTransaction() (x *Transaction) {
	if m == nil {
		return x
	}
	return m.Transaction
}

// GetAccountStates gets the AccountStates of the TransactionToCommit.
func (m *TransactionToCommit) GetAccountStates() (x []*AccountState) {
	if m == nil {
		return x
	}
	return m.AccountStates
}

// GetEvents gets the Events of the TransactionToCommit.
func (m *TransactionToCommit) GetEvents() (x []*Event) {
	if m == nil {
		return x
	}
	return m.Events
}

// GetGasUsed gets the GasUsed of the TransactionToCommit.
func (m *TransactionToCommit) GetGasUsed() (x uint64) {
	if m == nil {
		return x
	}
	return m.GasUsed
}

// GetMajorStatus gets the MajorStatus of the TransactionToCommit.
func (m *TransactionToCommit) GetMajorStatus() (x uint64) {
	if m == nil {
		return x
	}
	return m.MajorStatus
}

// MarshalToWriter marshals TransactionToCommit to the provided writer.
func (m *TransactionToCommit) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Transaction != nil {
		writer.WriteMessage(1, func() {
			m.Transaction.MarshalToWriter(writer)
		})
	}

	for _, msg := range m.AccountStates {
		writer.WriteMessage(2, func() {
			msg.MarshalToWriter(writer)
		})
	}

	for _, msg := range m.Events {
		writer.WriteMessage(3, func() {
			msg.MarshalToWriter(writer)
		})
	}

	if m.GasUsed != 0 {
		writer.WriteUint64(4, m.GasUsed)
	}

	if m.MajorStatus != 0 {
		writer.WriteUint64(5, m.MajorStatus)
	}

	return
}

// Marshal marshals TransactionToCommit to a slice of bytes.
func (m *TransactionToCommit) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a TransactionToCommit from the provided reader.
func (m *TransactionToCommit) UnmarshalFromReader(reader jspb.Reader) *TransactionToCommit {
	for reader.Next() {
		if m == nil {
			m = &TransactionToCommit{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.Transaction = m.Transaction.UnmarshalFromReader(reader)
			})
		case 2:
			reader.ReadMessage(func() {
				m.AccountStates = append(m.AccountStates, new(AccountState).UnmarshalFromReader(reader))
			})
		case 3:
			reader.ReadMessage(func() {
				m.Events = append(m.Events, new(Event).UnmarshalFromReader(reader))
			})
		case 4:
			m.GasUsed = reader.ReadUint64()
		case 5:
			m.MajorStatus = reader.ReadUint64()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a TransactionToCommit from a slice of bytes.
func (m *TransactionToCommit) Unmarshal(rawBytes []byte) (*TransactionToCommit, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// A list of consecutive transactions with proof. This is mainly used for state
// synchronization when a validator would request a list of transactions from a
// peer, verify the proof, execute the transactions and persist them. Note that
// the transactions are supposed to belong to the same epoch E, otherwise
// verification will fail.
type TransactionListWithProof struct {
	// The list of transactions.
	Transactions []*Transaction
	// The list of corresponding Event objects (only present if fetch_events was set to true in req)
	EventsForVersions *EventsForVersions
	// If the list is not empty, the version of the first transaction.
	FirstTransactionVersion *google_protobuf1.UInt64Value
	// The proof authenticating the transactions and events.When this is used
	// for state synchronization, the validator who requests the transactions
	// will provide a version in the request and the proofs will be relative to
	// the given version. When this is returned in GetTransactionsResponse, the
	// proofs will be relative to the ledger info returned in
	// UpdateToLatestLedgerResponse.
	Proof *TransactionListProof
}

// GetTransactions gets the Transactions of the TransactionListWithProof.
func (m *TransactionListWithProof) GetTransactions() (x []*Transaction) {
	if m == nil {
		return x
	}
	return m.Transactions
}

// GetEventsForVersions gets the EventsForVersions of the TransactionListWithProof.
func (m *TransactionListWithProof) GetEventsForVersions() (x *EventsForVersions) {
	if m == nil {
		return x
	}
	return m.EventsForVersions
}

// GetFirstTransactionVersion gets the FirstTransactionVersion of the TransactionListWithProof.
func (m *TransactionListWithProof) GetFirstTransactionVersion() (x *google_protobuf1.UInt64Value) {
	if m == nil {
		return x
	}
	return m.FirstTransactionVersion
}

// GetProof gets the Proof of the TransactionListWithProof.
func (m *TransactionListWithProof) GetProof() (x *TransactionListProof) {
	if m == nil {
		return x
	}
	return m.Proof
}

// MarshalToWriter marshals TransactionListWithProof to the provided writer.
func (m *TransactionListWithProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, msg := range m.Transactions {
		writer.WriteMessage(1, func() {
			msg.MarshalToWriter(writer)
		})
	}

	if m.EventsForVersions != nil {
		writer.WriteMessage(2, func() {
			m.EventsForVersions.MarshalToWriter(writer)
		})
	}

	if m.FirstTransactionVersion != nil {
		writer.WriteMessage(3, func() {
			m.FirstTransactionVersion.MarshalToWriter(writer)
		})
	}

	if m.Proof != nil {
		writer.WriteMessage(4, func() {
			m.Proof.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals TransactionListWithProof to a slice of bytes.
func (m *TransactionListWithProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a TransactionListWithProof from the provided reader.
func (m *TransactionListWithProof) UnmarshalFromReader(reader jspb.Reader) *TransactionListWithProof {
	for reader.Next() {
		if m == nil {
			m = &TransactionListWithProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.Transactions = append(m.Transactions, new(Transaction).UnmarshalFromReader(reader))
			})
		case 2:
			reader.ReadMessage(func() {
				m.EventsForVersions = m.EventsForVersions.UnmarshalFromReader(reader)
			})
		case 3:
			reader.ReadMessage(func() {
				m.FirstTransactionVersion = m.FirstTransactionVersion.UnmarshalFromReader(reader)
			})
		case 4:
			reader.ReadMessage(func() {
				m.Proof = m.Proof.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a TransactionListWithProof from a slice of bytes.
func (m *TransactionListWithProof) Unmarshal(rawBytes []byte) (*TransactionListWithProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
