// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: packet.proto
//
// versions:
//     protoc-gen-pico: (devel)
//     protoc:          v4.23.4

package pb

import (
	picobuf "storj.io/picobuf"
)

type Timestamp struct {
	Seconds int64 `json:"seconds,omitempty"`
	Nanos   int32 `json:"nanos,omitempty"`
}

func (m *Timestamp) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int64(1, &m.Seconds)
	c.Int32(2, &m.Nanos)
	return true
}

func (m *Timestamp) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int64(1, &m.Seconds)
	c.Int32(2, &m.Nanos)
}

type Tag struct {
	Key   string `json:"key,omitempty"`
	Value isTag_Value
}

func (m *Tag) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Key)
	if m, ok := m.Value.(*Tag_String_); ok {
		c.Bytes(2, &m.String_)
	}
	if m, ok := m.Value.(*Tag_Int64); ok {
		c.Int64(3, &m.Int64)
	}
	if m, ok := m.Value.(*Tag_Double); ok {
		c.Double(4, &m.Double)
	}
	if m, ok := m.Value.(*Tag_Bytes); ok {
		c.Bytes(5, &m.Bytes)
	}
	if m, ok := m.Value.(*Tag_Bool); ok {
		c.Bool(6, &m.Bool)
	}
	if m, ok := m.Value.(*Tag_DurationNs); ok {
		c.Int64(7, &m.DurationNs)
	}
	if m, ok := m.Value.(*Tag_Timestamp); ok {
		c.Message(8, m.Timestamp.Encode)
	}
	return true
}

func (m *Tag) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Key)
	if c.PendingField() == 2 {
		var x *Tag_String_
		if z, ok := m.Value.(*Tag_String_); ok {
			x = z
		} else {
			x = new(Tag_String_)
			m.Value = x
		}
		m := x
		c.Bytes(2, &m.String_)
	}
	if c.PendingField() == 3 {
		var x *Tag_Int64
		if z, ok := m.Value.(*Tag_Int64); ok {
			x = z
		} else {
			x = new(Tag_Int64)
			m.Value = x
		}
		m := x
		c.Int64(3, &m.Int64)
	}
	if c.PendingField() == 4 {
		var x *Tag_Double
		if z, ok := m.Value.(*Tag_Double); ok {
			x = z
		} else {
			x = new(Tag_Double)
			m.Value = x
		}
		m := x
		c.Double(4, &m.Double)
	}
	if c.PendingField() == 5 {
		var x *Tag_Bytes
		if z, ok := m.Value.(*Tag_Bytes); ok {
			x = z
		} else {
			x = new(Tag_Bytes)
			m.Value = x
		}
		m := x
		c.Bytes(5, &m.Bytes)
	}
	if c.PendingField() == 6 {
		var x *Tag_Bool
		if z, ok := m.Value.(*Tag_Bool); ok {
			x = z
		} else {
			x = new(Tag_Bool)
			m.Value = x
		}
		m := x
		c.Bool(6, &m.Bool)
	}
	if c.PendingField() == 7 {
		var x *Tag_DurationNs
		if z, ok := m.Value.(*Tag_DurationNs); ok {
			x = z
		} else {
			x = new(Tag_DurationNs)
			m.Value = x
		}
		m := x
		c.Int64(7, &m.DurationNs)
	}
	if c.PendingField() == 8 {
		var x *Tag_Timestamp
		if z, ok := m.Value.(*Tag_Timestamp); ok {
			x = z
		} else {
			x = new(Tag_Timestamp)
			m.Value = x
		}
		m := x
		c.Message(8, func(c *picobuf.Decoder) {
			if m.Timestamp == nil {
				m.Timestamp = new(Timestamp)
			}
			m.Timestamp.Decode(c)
		})
	}
}

type isTag_Value interface{ isTag_Value() }

type Tag_String_ struct {
	String_ []byte
}

type Tag_Int64 struct {
	Int64 int64
}

type Tag_Double struct {
	Double float64
}

type Tag_Bytes struct {
	Bytes []byte
}

type Tag_Bool struct {
	Bool bool
}

type Tag_DurationNs struct {
	DurationNs int64
}

type Tag_Timestamp struct {
	Timestamp *Timestamp
}

func (*Tag_String_) isTag_Value()    {}
func (*Tag_Int64) isTag_Value()      {}
func (*Tag_Double) isTag_Value()     {}
func (*Tag_Bytes) isTag_Value()      {}
func (*Tag_Bool) isTag_Value()       {}
func (*Tag_DurationNs) isTag_Value() {}
func (*Tag_Timestamp) isTag_Value()  {}

type Event struct {
	Name              string   `json:"name,omitempty"`
	Scope             []string `json:"scope,omitempty"`
	TimestampOffsetNs int64    `json:"timestamp_offset_ns,omitempty"`
	Tags              []*Tag   `json:"tags,omitempty"`
}

func (m *Event) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Name)
	c.RepeatedString(2, &m.Scope)
	c.Int64(3, &m.TimestampOffsetNs)
	for _, x := range m.Tags {
		c.AlwaysMessage(4, x.Encode)
	}
	return true
}

func (m *Event) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Name)
	c.RepeatedString(2, &m.Scope)
	c.Int64(3, &m.TimestampOffsetNs)
	c.RepeatedMessage(4, func(c *picobuf.Decoder) {
		x := new(Tag)
		c.Loop(x.Decode)
		m.Tags = append(m.Tags, x)
	})
}

type Packet struct {
	Application        string     `json:"application,omitempty"`
	ApplicationVersion string     `json:"application_version,omitempty"`
	Instance           string     `json:"instance,omitempty"`
	StartTimestamp     *Timestamp `json:"start_timestamp,omitempty"`
	SendOffsetNs       int64      `json:"send_offset_ns,omitempty"`
	Events             []*Event   `json:"events,omitempty"`
}

func (m *Packet) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Application)
	c.String(2, &m.ApplicationVersion)
	c.String(3, &m.Instance)
	c.Message(4, m.StartTimestamp.Encode)
	c.Int64(5, &m.SendOffsetNs)
	for _, x := range m.Events {
		c.AlwaysMessage(6, x.Encode)
	}
	return true
}

func (m *Packet) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Application)
	c.String(2, &m.ApplicationVersion)
	c.String(3, &m.Instance)
	c.Message(4, func(c *picobuf.Decoder) {
		if m.StartTimestamp == nil {
			m.StartTimestamp = new(Timestamp)
		}
		m.StartTimestamp.Decode(c)
	})
	c.Int64(5, &m.SendOffsetNs)
	c.RepeatedMessage(6, func(c *picobuf.Decoder) {
		x := new(Event)
		c.Loop(x.Decode)
		m.Events = append(m.Events, x)
	})
}

type Record struct {
	Application           string     `json:"application,omitempty"`
	ApplicationVersion    string     `json:"application_version,omitempty"`
	Instance              string     `json:"instance,omitempty"`
	SourceAddr            string     `json:"source_addr,omitempty"`
	Timestamp             *Timestamp `json:"timestamp,omitempty"`
	TimestampCorrectionNs int64      `json:"timestamp_correction_ns,omitempty"`
	Tags                  []*Tag     `json:"tags,omitempty"`
}

func (m *Record) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Application)
	c.String(2, &m.ApplicationVersion)
	c.String(3, &m.Instance)
	c.String(4, &m.SourceAddr)
	c.Message(5, m.Timestamp.Encode)
	c.Int64(6, &m.TimestampCorrectionNs)
	for _, x := range m.Tags {
		c.AlwaysMessage(7, x.Encode)
	}
	return true
}

func (m *Record) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Application)
	c.String(2, &m.ApplicationVersion)
	c.String(3, &m.Instance)
	c.String(4, &m.SourceAddr)
	c.Message(5, func(c *picobuf.Decoder) {
		if m.Timestamp == nil {
			m.Timestamp = new(Timestamp)
		}
		m.Timestamp.Decode(c)
	})
	c.Int64(6, &m.TimestampCorrectionNs)
	c.RepeatedMessage(7, func(c *picobuf.Decoder) {
		x := new(Tag)
		c.Loop(x.Decode)
		m.Tags = append(m.Tags, x)
	})
}
