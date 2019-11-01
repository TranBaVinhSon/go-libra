// +build js
// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: account_state_blob.proto

package pbtypes

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type AccountStateBlob struct {
	Blob []byte
}

// GetBlob gets the Blob of the AccountStateBlob.
func (m *AccountStateBlob) GetBlob() (x []byte) {
	if m == nil {
		return x
	}
	return m.Blob
}

// MarshalToWriter marshals AccountStateBlob to the provided writer.
func (m *AccountStateBlob) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Blob) > 0 {
		writer.WriteBytes(1, m.Blob)
	}

	return
}

// Marshal marshals AccountStateBlob to a slice of bytes.
func (m *AccountStateBlob) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a AccountStateBlob from the provided reader.
func (m *AccountStateBlob) UnmarshalFromReader(reader jspb.Reader) *AccountStateBlob {
	for reader.Next() {
		if m == nil {
			m = &AccountStateBlob{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Blob = reader.ReadBytes()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a AccountStateBlob from a slice of bytes.
func (m *AccountStateBlob) Unmarshal(rawBytes []byte) (*AccountStateBlob, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type AccountStateWithProof struct {
	Version uint64
	Blob    *AccountStateBlob
	Proof   *AccountStateProof
}

// GetVersion gets the Version of the AccountStateWithProof.
func (m *AccountStateWithProof) GetVersion() (x uint64) {
	if m == nil {
		return x
	}
	return m.Version
}

// GetBlob gets the Blob of the AccountStateWithProof.
func (m *AccountStateWithProof) GetBlob() (x *AccountStateBlob) {
	if m == nil {
		return x
	}
	return m.Blob
}

// GetProof gets the Proof of the AccountStateWithProof.
func (m *AccountStateWithProof) GetProof() (x *AccountStateProof) {
	if m == nil {
		return x
	}
	return m.Proof
}

// MarshalToWriter marshals AccountStateWithProof to the provided writer.
func (m *AccountStateWithProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.Version != 0 {
		writer.WriteUint64(1, m.Version)
	}

	if m.Blob != nil {
		writer.WriteMessage(2, func() {
			m.Blob.MarshalToWriter(writer)
		})
	}

	if m.Proof != nil {
		writer.WriteMessage(3, func() {
			m.Proof.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals AccountStateWithProof to a slice of bytes.
func (m *AccountStateWithProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a AccountStateWithProof from the provided reader.
func (m *AccountStateWithProof) UnmarshalFromReader(reader jspb.Reader) *AccountStateWithProof {
	for reader.Next() {
		if m == nil {
			m = &AccountStateWithProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Version = reader.ReadUint64()
		case 2:
			reader.ReadMessage(func() {
				m.Blob = m.Blob.UnmarshalFromReader(reader)
			})
		case 3:
			reader.ReadMessage(func() {
				m.Proof = m.Proof.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a AccountStateWithProof from a slice of bytes.
func (m *AccountStateWithProof) Unmarshal(rawBytes []byte) (*AccountStateWithProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
