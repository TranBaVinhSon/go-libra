// +build js
// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: events.proto

package pbtypes

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

// An event emitted from a smart contract
type Event struct {
	Key            []byte
	SequenceNumber uint64
	EventData      []byte
	TypeTag        []byte
}

// GetKey gets the Key of the Event.
func (m *Event) GetKey() (x []byte) {
	if m == nil {
		return x
	}
	return m.Key
}

// GetSequenceNumber gets the SequenceNumber of the Event.
func (m *Event) GetSequenceNumber() (x uint64) {
	if m == nil {
		return x
	}
	return m.SequenceNumber
}

// GetEventData gets the EventData of the Event.
func (m *Event) GetEventData() (x []byte) {
	if m == nil {
		return x
	}
	return m.EventData
}

// GetTypeTag gets the TypeTag of the Event.
func (m *Event) GetTypeTag() (x []byte) {
	if m == nil {
		return x
	}
	return m.TypeTag
}

// MarshalToWriter marshals Event to the provided writer.
func (m *Event) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Key) > 0 {
		writer.WriteBytes(1, m.Key)
	}

	if m.SequenceNumber != 0 {
		writer.WriteUint64(2, m.SequenceNumber)
	}

	if len(m.EventData) > 0 {
		writer.WriteBytes(3, m.EventData)
	}

	if len(m.TypeTag) > 0 {
		writer.WriteBytes(4, m.TypeTag)
	}

	return
}

// Marshal marshals Event to a slice of bytes.
func (m *Event) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Event from the provided reader.
func (m *Event) UnmarshalFromReader(reader jspb.Reader) *Event {
	for reader.Next() {
		if m == nil {
			m = &Event{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Key = reader.ReadBytes()
		case 2:
			m.SequenceNumber = reader.ReadUint64()
		case 3:
			m.EventData = reader.ReadBytes()
		case 4:
			m.TypeTag = reader.ReadBytes()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Event from a slice of bytes.
func (m *Event) Unmarshal(rawBytes []byte) (*Event, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// An event along with the proof for the event
type EventWithProof struct {
	TransactionVersion uint64
	EventIndex         uint64
	Event              *Event
	Proof              *EventProof
}

// GetTransactionVersion gets the TransactionVersion of the EventWithProof.
func (m *EventWithProof) GetTransactionVersion() (x uint64) {
	if m == nil {
		return x
	}
	return m.TransactionVersion
}

// GetEventIndex gets the EventIndex of the EventWithProof.
func (m *EventWithProof) GetEventIndex() (x uint64) {
	if m == nil {
		return x
	}
	return m.EventIndex
}

// GetEvent gets the Event of the EventWithProof.
func (m *EventWithProof) GetEvent() (x *Event) {
	if m == nil {
		return x
	}
	return m.Event
}

// GetProof gets the Proof of the EventWithProof.
func (m *EventWithProof) GetProof() (x *EventProof) {
	if m == nil {
		return x
	}
	return m.Proof
}

// MarshalToWriter marshals EventWithProof to the provided writer.
func (m *EventWithProof) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.TransactionVersion != 0 {
		writer.WriteUint64(1, m.TransactionVersion)
	}

	if m.EventIndex != 0 {
		writer.WriteUint64(2, m.EventIndex)
	}

	if m.Event != nil {
		writer.WriteMessage(3, func() {
			m.Event.MarshalToWriter(writer)
		})
	}

	if m.Proof != nil {
		writer.WriteMessage(4, func() {
			m.Proof.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals EventWithProof to a slice of bytes.
func (m *EventWithProof) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a EventWithProof from the provided reader.
func (m *EventWithProof) UnmarshalFromReader(reader jspb.Reader) *EventWithProof {
	for reader.Next() {
		if m == nil {
			m = &EventWithProof{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.TransactionVersion = reader.ReadUint64()
		case 2:
			m.EventIndex = reader.ReadUint64()
		case 3:
			reader.ReadMessage(func() {
				m.Event = m.Event.UnmarshalFromReader(reader)
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

// Unmarshal unmarshals a EventWithProof from a slice of bytes.
func (m *EventWithProof) Unmarshal(rawBytes []byte) (*EventWithProof, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// A list of events.
type EventsList struct {
	Events []*Event
}

// GetEvents gets the Events of the EventsList.
func (m *EventsList) GetEvents() (x []*Event) {
	if m == nil {
		return x
	}
	return m.Events
}

// MarshalToWriter marshals EventsList to the provided writer.
func (m *EventsList) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, msg := range m.Events {
		writer.WriteMessage(1, func() {
			msg.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals EventsList to a slice of bytes.
func (m *EventsList) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a EventsList from the provided reader.
func (m *EventsList) UnmarshalFromReader(reader jspb.Reader) *EventsList {
	for reader.Next() {
		if m == nil {
			m = &EventsList{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.Events = append(m.Events, new(Event).UnmarshalFromReader(reader))
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a EventsList from a slice of bytes.
func (m *EventsList) Unmarshal(rawBytes []byte) (*EventsList, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// A list of EventList's, each representing all events for a transaction.
type EventsForVersions struct {
	EventsForVersion []*EventsList
}

// GetEventsForVersion gets the EventsForVersion of the EventsForVersions.
func (m *EventsForVersions) GetEventsForVersion() (x []*EventsList) {
	if m == nil {
		return x
	}
	return m.EventsForVersion
}

// MarshalToWriter marshals EventsForVersions to the provided writer.
func (m *EventsForVersions) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	for _, msg := range m.EventsForVersion {
		writer.WriteMessage(1, func() {
			msg.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals EventsForVersions to a slice of bytes.
func (m *EventsForVersions) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a EventsForVersions from the provided reader.
func (m *EventsForVersions) UnmarshalFromReader(reader jspb.Reader) *EventsForVersions {
	for reader.Next() {
		if m == nil {
			m = &EventsForVersions{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			reader.ReadMessage(func() {
				m.EventsForVersion = append(m.EventsForVersion, new(EventsList).UnmarshalFromReader(reader))
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a EventsForVersions from a slice of bytes.
func (m *EventsForVersions) Unmarshal(rawBytes []byte) (*EventsForVersions, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
