// +build js
// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: validator_set.proto

package pbtypes

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

// Protobuf definition for the Rust struct ValidatorSet.
type ValidatorSet struct {
	ValidatorPublicKeys []*ValidatorPublicKeys
}

// GetValidatorPublicKeys gets the ValidatorPublicKeys of the ValidatorSet.
func (m *ValidatorSet) GetValidatorPublicKeys() (x []*ValidatorPublicKeys) {
	if m == nil {
		return x
	}
	return m.ValidatorPublicKeys
}

// MarshalToWriter marshals ValidatorSet to the provided writer.
func (m *ValidatorSet) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, msg := range m.ValidatorPublicKeys {
		writer.WriteMessage(1, func() {
			msg.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals ValidatorSet to a slice of bytes.
func (m *ValidatorSet) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a ValidatorSet from the provided reader.
func (m *ValidatorSet) UnmarshalFromReader(reader jspb.Reader) *ValidatorSet {
	for reader.Next() {
		if m == nil {
			m = &ValidatorSet{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.ValidatorPublicKeys = append(m.ValidatorPublicKeys, new(ValidatorPublicKeys).UnmarshalFromReader(reader))
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a ValidatorSet from a slice of bytes.
func (m *ValidatorSet) Unmarshal(rawBytes []byte) (*ValidatorSet, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}