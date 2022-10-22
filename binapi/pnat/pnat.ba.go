// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.06-rc0~150-g7bb77ca6f
// source: /usr/share/vpp/api/plugins/pnat.api.json

// Package pnat contains generated bindings for API file pnat.api.
//
// Contents:
//   2 enums
//   2 structs
//  14 messages
//
package pnat

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "pnat"
	APIVersion = "0.1.1"
	VersionCrc = 0x108d3b87
)

// PnatAttachmentPoint defines enum 'pnat_attachment_point'.
type PnatAttachmentPoint uint32

const (
	PNAT_IP4_INPUT            PnatAttachmentPoint = 0
	PNAT_IP4_OUTPUT           PnatAttachmentPoint = 1
	PNAT_ATTACHMENT_POINT_MAX PnatAttachmentPoint = 2
)

var (
	PnatAttachmentPoint_name = map[uint32]string{
		0: "PNAT_IP4_INPUT",
		1: "PNAT_IP4_OUTPUT",
		2: "PNAT_ATTACHMENT_POINT_MAX",
	}
	PnatAttachmentPoint_value = map[string]uint32{
		"PNAT_IP4_INPUT":            0,
		"PNAT_IP4_OUTPUT":           1,
		"PNAT_ATTACHMENT_POINT_MAX": 2,
	}
)

func (x PnatAttachmentPoint) String() string {
	s, ok := PnatAttachmentPoint_name[uint32(x)]
	if ok {
		return s
	}
	return "PnatAttachmentPoint(" + strconv.Itoa(int(x)) + ")"
}

// PnatMask defines enum 'pnat_mask'.
type PnatMask uint32

const (
	PNAT_SA         PnatMask = 1
	PNAT_DA         PnatMask = 2
	PNAT_SPORT      PnatMask = 4
	PNAT_DPORT      PnatMask = 8
	PNAT_COPY_BYTE  PnatMask = 16
	PNAT_CLEAR_BYTE PnatMask = 32
)

var (
	PnatMask_name = map[uint32]string{
		1:  "PNAT_SA",
		2:  "PNAT_DA",
		4:  "PNAT_SPORT",
		8:  "PNAT_DPORT",
		16: "PNAT_COPY_BYTE",
		32: "PNAT_CLEAR_BYTE",
	}
	PnatMask_value = map[string]uint32{
		"PNAT_SA":         1,
		"PNAT_DA":         2,
		"PNAT_SPORT":      4,
		"PNAT_DPORT":      8,
		"PNAT_COPY_BYTE":  16,
		"PNAT_CLEAR_BYTE": 32,
	}
)

func (x PnatMask) String() string {
	s, ok := PnatMask_name[uint32(x)]
	if ok {
		return s
	}
	return "PnatMask(" + strconv.Itoa(int(x)) + ")"
}

// PnatMatchTuple defines type 'pnat_match_tuple'.
type PnatMatchTuple struct {
	Src   ip_types.IP4Address `binapi:"ip4_address,name=src" json:"src,omitempty"`
	Dst   ip_types.IP4Address `binapi:"ip4_address,name=dst" json:"dst,omitempty"`
	Proto ip_types.IPProto    `binapi:"ip_proto,name=proto" json:"proto,omitempty"`
	Sport uint16              `binapi:"u16,name=sport" json:"sport,omitempty"`
	Dport uint16              `binapi:"u16,name=dport" json:"dport,omitempty"`
	Mask  PnatMask            `binapi:"pnat_mask,name=mask" json:"mask,omitempty"`
}

// PnatRewriteTuple defines type 'pnat_rewrite_tuple'.
type PnatRewriteTuple struct {
	Src         ip_types.IP4Address `binapi:"ip4_address,name=src" json:"src,omitempty"`
	Dst         ip_types.IP4Address `binapi:"ip4_address,name=dst" json:"dst,omitempty"`
	Sport       uint16              `binapi:"u16,name=sport" json:"sport,omitempty"`
	Dport       uint16              `binapi:"u16,name=dport" json:"dport,omitempty"`
	Mask        PnatMask            `binapi:"pnat_mask,name=mask" json:"mask,omitempty"`
	FromOffset  uint8               `binapi:"u8,name=from_offset" json:"from_offset,omitempty"`
	ToOffset    uint8               `binapi:"u8,name=to_offset" json:"to_offset,omitempty"`
	ClearOffset uint8               `binapi:"u8,name=clear_offset" json:"clear_offset,omitempty"`
}

// PnatBindingAdd defines message 'pnat_binding_add'.
// InProgress: the message form may change in the future versions
type PnatBindingAdd struct {
	Match   PnatMatchTuple   `binapi:"pnat_match_tuple,name=match" json:"match,omitempty"`
	Rewrite PnatRewriteTuple `binapi:"pnat_rewrite_tuple,name=rewrite" json:"rewrite,omitempty"`
}

func (m *PnatBindingAdd) Reset()               { *m = PnatBindingAdd{} }
func (*PnatBindingAdd) GetMessageName() string { return "pnat_binding_add" }
func (*PnatBindingAdd) GetCrcString() string   { return "f00f79aa" }
func (*PnatBindingAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PnatBindingAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 4 // m.Match.Src
	size += 1 * 4 // m.Match.Dst
	size += 1     // m.Match.Proto
	size += 2     // m.Match.Sport
	size += 2     // m.Match.Dport
	size += 4     // m.Match.Mask
	size += 1 * 4 // m.Rewrite.Src
	size += 1 * 4 // m.Rewrite.Dst
	size += 2     // m.Rewrite.Sport
	size += 2     // m.Rewrite.Dport
	size += 4     // m.Rewrite.Mask
	size += 1     // m.Rewrite.FromOffset
	size += 1     // m.Rewrite.ToOffset
	size += 1     // m.Rewrite.ClearOffset
	return size
}
func (m *PnatBindingAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Match.Src[:], 4)
	buf.EncodeBytes(m.Match.Dst[:], 4)
	buf.EncodeUint8(uint8(m.Match.Proto))
	buf.EncodeUint16(m.Match.Sport)
	buf.EncodeUint16(m.Match.Dport)
	buf.EncodeUint32(uint32(m.Match.Mask))
	buf.EncodeBytes(m.Rewrite.Src[:], 4)
	buf.EncodeBytes(m.Rewrite.Dst[:], 4)
	buf.EncodeUint16(m.Rewrite.Sport)
	buf.EncodeUint16(m.Rewrite.Dport)
	buf.EncodeUint32(uint32(m.Rewrite.Mask))
	buf.EncodeUint8(m.Rewrite.FromOffset)
	buf.EncodeUint8(m.Rewrite.ToOffset)
	buf.EncodeUint8(m.Rewrite.ClearOffset)
	return buf.Bytes(), nil
}
func (m *PnatBindingAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.Match.Src[:], buf.DecodeBytes(4))
	copy(m.Match.Dst[:], buf.DecodeBytes(4))
	m.Match.Proto = ip_types.IPProto(buf.DecodeUint8())
	m.Match.Sport = buf.DecodeUint16()
	m.Match.Dport = buf.DecodeUint16()
	m.Match.Mask = PnatMask(buf.DecodeUint32())
	copy(m.Rewrite.Src[:], buf.DecodeBytes(4))
	copy(m.Rewrite.Dst[:], buf.DecodeBytes(4))
	m.Rewrite.Sport = buf.DecodeUint16()
	m.Rewrite.Dport = buf.DecodeUint16()
	m.Rewrite.Mask = PnatMask(buf.DecodeUint32())
	m.Rewrite.FromOffset = buf.DecodeUint8()
	m.Rewrite.ToOffset = buf.DecodeUint8()
	m.Rewrite.ClearOffset = buf.DecodeUint8()
	return nil
}

// PnatBindingAddReply defines message 'pnat_binding_add_reply'.
// InProgress: the message form may change in the future versions
type PnatBindingAddReply struct {
	Retval       int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	BindingIndex uint32 `binapi:"u32,name=binding_index" json:"binding_index,omitempty"`
}

func (m *PnatBindingAddReply) Reset()               { *m = PnatBindingAddReply{} }
func (*PnatBindingAddReply) GetMessageName() string { return "pnat_binding_add_reply" }
func (*PnatBindingAddReply) GetCrcString() string   { return "4cd980a7" }
func (*PnatBindingAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PnatBindingAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.BindingIndex
	return size
}
func (m *PnatBindingAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.BindingIndex)
	return buf.Bytes(), nil
}
func (m *PnatBindingAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.BindingIndex = buf.DecodeUint32()
	return nil
}

// PnatBindingAttach defines message 'pnat_binding_attach'.
// InProgress: the message form may change in the future versions
type PnatBindingAttach struct {
	SwIfIndex    interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Attachment   PnatAttachmentPoint            `binapi:"pnat_attachment_point,name=attachment" json:"attachment,omitempty"`
	BindingIndex uint32                         `binapi:"u32,name=binding_index" json:"binding_index,omitempty"`
}

func (m *PnatBindingAttach) Reset()               { *m = PnatBindingAttach{} }
func (*PnatBindingAttach) GetMessageName() string { return "pnat_binding_attach" }
func (*PnatBindingAttach) GetCrcString() string   { return "6e074232" }
func (*PnatBindingAttach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PnatBindingAttach) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 4 // m.Attachment
	size += 4 // m.BindingIndex
	return size
}
func (m *PnatBindingAttach) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(uint32(m.Attachment))
	buf.EncodeUint32(m.BindingIndex)
	return buf.Bytes(), nil
}
func (m *PnatBindingAttach) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Attachment = PnatAttachmentPoint(buf.DecodeUint32())
	m.BindingIndex = buf.DecodeUint32()
	return nil
}

// PnatBindingAttachReply defines message 'pnat_binding_attach_reply'.
// InProgress: the message form may change in the future versions
type PnatBindingAttachReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PnatBindingAttachReply) Reset()               { *m = PnatBindingAttachReply{} }
func (*PnatBindingAttachReply) GetMessageName() string { return "pnat_binding_attach_reply" }
func (*PnatBindingAttachReply) GetCrcString() string   { return "e8d4e804" }
func (*PnatBindingAttachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PnatBindingAttachReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PnatBindingAttachReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PnatBindingAttachReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PnatBindingDel defines message 'pnat_binding_del'.
// InProgress: the message form may change in the future versions
type PnatBindingDel struct {
	BindingIndex uint32 `binapi:"u32,name=binding_index" json:"binding_index,omitempty"`
}

func (m *PnatBindingDel) Reset()               { *m = PnatBindingDel{} }
func (*PnatBindingDel) GetMessageName() string { return "pnat_binding_del" }
func (*PnatBindingDel) GetCrcString() string   { return "9259df7b" }
func (*PnatBindingDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PnatBindingDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.BindingIndex
	return size
}
func (m *PnatBindingDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.BindingIndex)
	return buf.Bytes(), nil
}
func (m *PnatBindingDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.BindingIndex = buf.DecodeUint32()
	return nil
}

// PnatBindingDelReply defines message 'pnat_binding_del_reply'.
// InProgress: the message form may change in the future versions
type PnatBindingDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PnatBindingDelReply) Reset()               { *m = PnatBindingDelReply{} }
func (*PnatBindingDelReply) GetMessageName() string { return "pnat_binding_del_reply" }
func (*PnatBindingDelReply) GetCrcString() string   { return "e8d4e804" }
func (*PnatBindingDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PnatBindingDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PnatBindingDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PnatBindingDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PnatBindingDetach defines message 'pnat_binding_detach'.
// InProgress: the message form may change in the future versions
type PnatBindingDetach struct {
	SwIfIndex    interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Attachment   PnatAttachmentPoint            `binapi:"pnat_attachment_point,name=attachment" json:"attachment,omitempty"`
	BindingIndex uint32                         `binapi:"u32,name=binding_index" json:"binding_index,omitempty"`
}

func (m *PnatBindingDetach) Reset()               { *m = PnatBindingDetach{} }
func (*PnatBindingDetach) GetMessageName() string { return "pnat_binding_detach" }
func (*PnatBindingDetach) GetCrcString() string   { return "6e074232" }
func (*PnatBindingDetach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PnatBindingDetach) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 4 // m.Attachment
	size += 4 // m.BindingIndex
	return size
}
func (m *PnatBindingDetach) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(uint32(m.Attachment))
	buf.EncodeUint32(m.BindingIndex)
	return buf.Bytes(), nil
}
func (m *PnatBindingDetach) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Attachment = PnatAttachmentPoint(buf.DecodeUint32())
	m.BindingIndex = buf.DecodeUint32()
	return nil
}

// PnatBindingDetachReply defines message 'pnat_binding_detach_reply'.
// InProgress: the message form may change in the future versions
type PnatBindingDetachReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PnatBindingDetachReply) Reset()               { *m = PnatBindingDetachReply{} }
func (*PnatBindingDetachReply) GetMessageName() string { return "pnat_binding_detach_reply" }
func (*PnatBindingDetachReply) GetCrcString() string   { return "e8d4e804" }
func (*PnatBindingDetachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PnatBindingDetachReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PnatBindingDetachReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PnatBindingDetachReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PnatBindingsDetails defines message 'pnat_bindings_details'.
// InProgress: the message form may change in the future versions
type PnatBindingsDetails struct {
	Match   PnatMatchTuple   `binapi:"pnat_match_tuple,name=match" json:"match,omitempty"`
	Rewrite PnatRewriteTuple `binapi:"pnat_rewrite_tuple,name=rewrite" json:"rewrite,omitempty"`
}

func (m *PnatBindingsDetails) Reset()               { *m = PnatBindingsDetails{} }
func (*PnatBindingsDetails) GetMessageName() string { return "pnat_bindings_details" }
func (*PnatBindingsDetails) GetCrcString() string   { return "78267a35" }
func (*PnatBindingsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PnatBindingsDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 4 // m.Match.Src
	size += 1 * 4 // m.Match.Dst
	size += 1     // m.Match.Proto
	size += 2     // m.Match.Sport
	size += 2     // m.Match.Dport
	size += 4     // m.Match.Mask
	size += 1 * 4 // m.Rewrite.Src
	size += 1 * 4 // m.Rewrite.Dst
	size += 2     // m.Rewrite.Sport
	size += 2     // m.Rewrite.Dport
	size += 4     // m.Rewrite.Mask
	size += 1     // m.Rewrite.FromOffset
	size += 1     // m.Rewrite.ToOffset
	size += 1     // m.Rewrite.ClearOffset
	return size
}
func (m *PnatBindingsDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Match.Src[:], 4)
	buf.EncodeBytes(m.Match.Dst[:], 4)
	buf.EncodeUint8(uint8(m.Match.Proto))
	buf.EncodeUint16(m.Match.Sport)
	buf.EncodeUint16(m.Match.Dport)
	buf.EncodeUint32(uint32(m.Match.Mask))
	buf.EncodeBytes(m.Rewrite.Src[:], 4)
	buf.EncodeBytes(m.Rewrite.Dst[:], 4)
	buf.EncodeUint16(m.Rewrite.Sport)
	buf.EncodeUint16(m.Rewrite.Dport)
	buf.EncodeUint32(uint32(m.Rewrite.Mask))
	buf.EncodeUint8(m.Rewrite.FromOffset)
	buf.EncodeUint8(m.Rewrite.ToOffset)
	buf.EncodeUint8(m.Rewrite.ClearOffset)
	return buf.Bytes(), nil
}
func (m *PnatBindingsDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.Match.Src[:], buf.DecodeBytes(4))
	copy(m.Match.Dst[:], buf.DecodeBytes(4))
	m.Match.Proto = ip_types.IPProto(buf.DecodeUint8())
	m.Match.Sport = buf.DecodeUint16()
	m.Match.Dport = buf.DecodeUint16()
	m.Match.Mask = PnatMask(buf.DecodeUint32())
	copy(m.Rewrite.Src[:], buf.DecodeBytes(4))
	copy(m.Rewrite.Dst[:], buf.DecodeBytes(4))
	m.Rewrite.Sport = buf.DecodeUint16()
	m.Rewrite.Dport = buf.DecodeUint16()
	m.Rewrite.Mask = PnatMask(buf.DecodeUint32())
	m.Rewrite.FromOffset = buf.DecodeUint8()
	m.Rewrite.ToOffset = buf.DecodeUint8()
	m.Rewrite.ClearOffset = buf.DecodeUint8()
	return nil
}

// PnatBindingsGet defines message 'pnat_bindings_get'.
// InProgress: the message form may change in the future versions
type PnatBindingsGet struct {
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *PnatBindingsGet) Reset()               { *m = PnatBindingsGet{} }
func (*PnatBindingsGet) GetMessageName() string { return "pnat_bindings_get" }
func (*PnatBindingsGet) GetCrcString() string   { return "f75ba505" }
func (*PnatBindingsGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PnatBindingsGet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Cursor
	return size
}
func (m *PnatBindingsGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *PnatBindingsGet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Cursor = buf.DecodeUint32()
	return nil
}

// PnatBindingsGetReply defines message 'pnat_bindings_get_reply'.
// InProgress: the message form may change in the future versions
type PnatBindingsGetReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *PnatBindingsGetReply) Reset()               { *m = PnatBindingsGetReply{} }
func (*PnatBindingsGetReply) GetMessageName() string { return "pnat_bindings_get_reply" }
func (*PnatBindingsGetReply) GetCrcString() string   { return "53b48f5d" }
func (*PnatBindingsGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PnatBindingsGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Cursor
	return size
}
func (m *PnatBindingsGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *PnatBindingsGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Cursor = buf.DecodeUint32()
	return nil
}

// PnatInterfacesDetails defines message 'pnat_interfaces_details'.
// InProgress: the message form may change in the future versions
type PnatInterfacesDetails struct {
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Enabled    []bool                         `binapi:"bool[2],name=enabled" json:"enabled,omitempty"`
	LookupMask [2]PnatMask                    `binapi:"pnat_mask[2],name=lookup_mask" json:"lookup_mask,omitempty"`
}

func (m *PnatInterfacesDetails) Reset()               { *m = PnatInterfacesDetails{} }
func (*PnatInterfacesDetails) GetMessageName() string { return "pnat_interfaces_details" }
func (*PnatInterfacesDetails) GetCrcString() string   { return "c7b0c4c0" }
func (*PnatInterfacesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PnatInterfacesDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.SwIfIndex
	size += 1 * 2 // m.Enabled
	for j1 := 0; j1 < 2; j1++ {
		size += 4 // m.LookupMask[j1]
	}
	return size
}
func (m *PnatInterfacesDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	for i := 0; i < 2; i++ {
		var x bool
		if i < len(m.Enabled) {
			x = bool(m.Enabled[i])
		}
		buf.EncodeBool(x)
	}
	for j0 := 0; j0 < 2; j0++ {
		buf.EncodeUint32(uint32(m.LookupMask[j0]))
	}
	return buf.Bytes(), nil
}
func (m *PnatInterfacesDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Enabled = make([]bool, 2)
	for i := 0; i < len(m.Enabled); i++ {
		m.Enabled[i] = buf.DecodeBool()
	}
	for j0 := 0; j0 < 2; j0++ {
		m.LookupMask[j0] = PnatMask(buf.DecodeUint32())
	}
	return nil
}

// PnatInterfacesGet defines message 'pnat_interfaces_get'.
// InProgress: the message form may change in the future versions
type PnatInterfacesGet struct {
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *PnatInterfacesGet) Reset()               { *m = PnatInterfacesGet{} }
func (*PnatInterfacesGet) GetMessageName() string { return "pnat_interfaces_get" }
func (*PnatInterfacesGet) GetCrcString() string   { return "f75ba505" }
func (*PnatInterfacesGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PnatInterfacesGet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Cursor
	return size
}
func (m *PnatInterfacesGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *PnatInterfacesGet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Cursor = buf.DecodeUint32()
	return nil
}

// PnatInterfacesGetReply defines message 'pnat_interfaces_get_reply'.
// InProgress: the message form may change in the future versions
type PnatInterfacesGetReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *PnatInterfacesGetReply) Reset()               { *m = PnatInterfacesGetReply{} }
func (*PnatInterfacesGetReply) GetMessageName() string { return "pnat_interfaces_get_reply" }
func (*PnatInterfacesGetReply) GetCrcString() string   { return "53b48f5d" }
func (*PnatInterfacesGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PnatInterfacesGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Cursor
	return size
}
func (m *PnatInterfacesGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *PnatInterfacesGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Cursor = buf.DecodeUint32()
	return nil
}

func init() { file_pnat_binapi_init() }
func file_pnat_binapi_init() {
	api.RegisterMessage((*PnatBindingAdd)(nil), "pnat_binding_add_f00f79aa")
	api.RegisterMessage((*PnatBindingAddReply)(nil), "pnat_binding_add_reply_4cd980a7")
	api.RegisterMessage((*PnatBindingAttach)(nil), "pnat_binding_attach_6e074232")
	api.RegisterMessage((*PnatBindingAttachReply)(nil), "pnat_binding_attach_reply_e8d4e804")
	api.RegisterMessage((*PnatBindingDel)(nil), "pnat_binding_del_9259df7b")
	api.RegisterMessage((*PnatBindingDelReply)(nil), "pnat_binding_del_reply_e8d4e804")
	api.RegisterMessage((*PnatBindingDetach)(nil), "pnat_binding_detach_6e074232")
	api.RegisterMessage((*PnatBindingDetachReply)(nil), "pnat_binding_detach_reply_e8d4e804")
	api.RegisterMessage((*PnatBindingsDetails)(nil), "pnat_bindings_details_78267a35")
	api.RegisterMessage((*PnatBindingsGet)(nil), "pnat_bindings_get_f75ba505")
	api.RegisterMessage((*PnatBindingsGetReply)(nil), "pnat_bindings_get_reply_53b48f5d")
	api.RegisterMessage((*PnatInterfacesDetails)(nil), "pnat_interfaces_details_c7b0c4c0")
	api.RegisterMessage((*PnatInterfacesGet)(nil), "pnat_interfaces_get_f75ba505")
	api.RegisterMessage((*PnatInterfacesGetReply)(nil), "pnat_interfaces_get_reply_53b48f5d")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PnatBindingAdd)(nil),
		(*PnatBindingAddReply)(nil),
		(*PnatBindingAttach)(nil),
		(*PnatBindingAttachReply)(nil),
		(*PnatBindingDel)(nil),
		(*PnatBindingDelReply)(nil),
		(*PnatBindingDetach)(nil),
		(*PnatBindingDetachReply)(nil),
		(*PnatBindingsDetails)(nil),
		(*PnatBindingsGet)(nil),
		(*PnatBindingsGetReply)(nil),
		(*PnatInterfacesDetails)(nil),
		(*PnatInterfacesGet)(nil),
		(*PnatInterfacesGetReply)(nil),
	}
}
