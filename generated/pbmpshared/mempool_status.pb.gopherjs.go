// +build js
// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: mempool_status.proto

/*
	Package pbmpshared is a generated protocol buffer package.

	It is generated from these files:
		mempool_status.proto

	It has these top-level messages:
		MempoolAddTransactionStatus
*/
package pbmpshared

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type MempoolAddTransactionStatusCode int

const (
	// Transaction was sent to Mempool
	MempoolAddTransactionStatusCode_Valid MempoolAddTransactionStatusCode = 0
	// The sender does not have enough balance for the transaction.
	MempoolAddTransactionStatusCode_InsufficientBalance MempoolAddTransactionStatusCode = 1
	// Sequence number is old, etc.
	MempoolAddTransactionStatusCode_InvalidSeqNumber MempoolAddTransactionStatusCode = 2
	// Mempool is full (reached max global capacity)
	MempoolAddTransactionStatusCode_MempoolIsFull MempoolAddTransactionStatusCode = 3
	// Account reached max capacity per account
	MempoolAddTransactionStatusCode_TooManyTransactions MempoolAddTransactionStatusCode = 4
	// Invalid update. Only gas price increase is allowed
	MempoolAddTransactionStatusCode_InvalidUpdate MempoolAddTransactionStatusCode = 5
)

var MempoolAddTransactionStatusCode_name = map[int]string{
	0: "Valid",
	1: "InsufficientBalance",
	2: "InvalidSeqNumber",
	3: "MempoolIsFull",
	4: "TooManyTransactions",
	5: "InvalidUpdate",
}
var MempoolAddTransactionStatusCode_value = map[string]int{
	"Valid":               0,
	"InsufficientBalance": 1,
	"InvalidSeqNumber":    2,
	"MempoolIsFull":       3,
	"TooManyTransactions": 4,
	"InvalidUpdate":       5,
}

func (x MempoolAddTransactionStatusCode) String() string {
	return MempoolAddTransactionStatusCode_name[int(x)]
}

type MempoolAddTransactionStatus struct {
	Code    MempoolAddTransactionStatusCode
	Message string
}

// GetCode gets the Code of the MempoolAddTransactionStatus.
func (m *MempoolAddTransactionStatus) GetCode() (x MempoolAddTransactionStatusCode) {
	if m == nil {
		return x
	}
	return m.Code
}

// GetMessage gets the Message of the MempoolAddTransactionStatus.
func (m *MempoolAddTransactionStatus) GetMessage() (x string) {
	if m == nil {
		return x
	}
	return m.Message
}

// MarshalToWriter marshals MempoolAddTransactionStatus to the provided writer.
func (m *MempoolAddTransactionStatus) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if int(m.Code) != 0 {
		writer.WriteEnum(1, int(m.Code))
	}

	if len(m.Message) > 0 {
		writer.WriteString(2, m.Message)
	}

	return
}

// Marshal marshals MempoolAddTransactionStatus to a slice of bytes.
func (m *MempoolAddTransactionStatus) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a MempoolAddTransactionStatus from the provided reader.
func (m *MempoolAddTransactionStatus) UnmarshalFromReader(reader jspb.Reader) *MempoolAddTransactionStatus {
	for reader.Next() {
		if m == nil {
			m = &MempoolAddTransactionStatus{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Code = MempoolAddTransactionStatusCode(reader.ReadEnum())
		case 2:
			m.Message = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a MempoolAddTransactionStatus from a slice of bytes.
func (m *MempoolAddTransactionStatus) Unmarshal(rawBytes []byte) (*MempoolAddTransactionStatus, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}