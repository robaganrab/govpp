// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.06-rc0~150-g7bb77ca6f
// source: /usr/share/vpp/api/core/crypto.api.json

// Package crypto contains generated bindings for API file crypto.api.
//
// Contents:
//   2 enums
//   4 messages
//
package crypto

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "crypto"
	APIVersion = "1.0.1"
	VersionCrc = 0x22355ec6
)

// CryptoDispatchMode defines enum 'crypto_dispatch_mode'.
type CryptoDispatchMode uint8

const (
	CRYPTO_ASYNC_DISPATCH_POLLING   CryptoDispatchMode = 0
	CRYPTO_ASYNC_DISPATCH_INTERRUPT CryptoDispatchMode = 1
)

var (
	CryptoDispatchMode_name = map[uint8]string{
		0: "CRYPTO_ASYNC_DISPATCH_POLLING",
		1: "CRYPTO_ASYNC_DISPATCH_INTERRUPT",
	}
	CryptoDispatchMode_value = map[string]uint8{
		"CRYPTO_ASYNC_DISPATCH_POLLING":   0,
		"CRYPTO_ASYNC_DISPATCH_INTERRUPT": 1,
	}
)

func (x CryptoDispatchMode) String() string {
	s, ok := CryptoDispatchMode_name[uint8(x)]
	if ok {
		return s
	}
	return "CryptoDispatchMode(" + strconv.Itoa(int(x)) + ")"
}

// CryptoOpClassType defines enum 'crypto_op_class_type'.
type CryptoOpClassType uint8

const (
	CRYPTO_API_OP_SIMPLE  CryptoOpClassType = 0
	CRYPTO_API_OP_CHAINED CryptoOpClassType = 1
	CRYPTO_API_OP_BOTH    CryptoOpClassType = 2
)

var (
	CryptoOpClassType_name = map[uint8]string{
		0: "CRYPTO_API_OP_SIMPLE",
		1: "CRYPTO_API_OP_CHAINED",
		2: "CRYPTO_API_OP_BOTH",
	}
	CryptoOpClassType_value = map[string]uint8{
		"CRYPTO_API_OP_SIMPLE":  0,
		"CRYPTO_API_OP_CHAINED": 1,
		"CRYPTO_API_OP_BOTH":    2,
	}
)

func (x CryptoOpClassType) String() string {
	s, ok := CryptoOpClassType_name[uint8(x)]
	if ok {
		return s
	}
	return "CryptoOpClassType(" + strconv.Itoa(int(x)) + ")"
}

// CryptoSetAsyncDispatch defines message 'crypto_set_async_dispatch'.
type CryptoSetAsyncDispatch struct {
	Mode CryptoDispatchMode `binapi:"crypto_dispatch_mode,name=mode" json:"mode,omitempty"`
}

func (m *CryptoSetAsyncDispatch) Reset()               { *m = CryptoSetAsyncDispatch{} }
func (*CryptoSetAsyncDispatch) GetMessageName() string { return "crypto_set_async_dispatch" }
func (*CryptoSetAsyncDispatch) GetCrcString() string   { return "5ca4adc0" }
func (*CryptoSetAsyncDispatch) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CryptoSetAsyncDispatch) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Mode
	return size
}
func (m *CryptoSetAsyncDispatch) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Mode))
	return buf.Bytes(), nil
}
func (m *CryptoSetAsyncDispatch) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Mode = CryptoDispatchMode(buf.DecodeUint8())
	return nil
}

// CryptoSetAsyncDispatchReply defines message 'crypto_set_async_dispatch_reply'.
type CryptoSetAsyncDispatchReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CryptoSetAsyncDispatchReply) Reset()               { *m = CryptoSetAsyncDispatchReply{} }
func (*CryptoSetAsyncDispatchReply) GetMessageName() string { return "crypto_set_async_dispatch_reply" }
func (*CryptoSetAsyncDispatchReply) GetCrcString() string   { return "e8d4e804" }
func (*CryptoSetAsyncDispatchReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CryptoSetAsyncDispatchReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CryptoSetAsyncDispatchReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CryptoSetAsyncDispatchReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// CryptoSetHandler defines message 'crypto_set_handler'.
type CryptoSetHandler struct {
	AlgName string            `binapi:"string[32],name=alg_name" json:"alg_name,omitempty"`
	Engine  string            `binapi:"string[16],name=engine" json:"engine,omitempty"`
	Oct     CryptoOpClassType `binapi:"crypto_op_class_type,name=oct" json:"oct,omitempty"`
	IsAsync uint8             `binapi:"u8,name=is_async" json:"is_async,omitempty"`
}

func (m *CryptoSetHandler) Reset()               { *m = CryptoSetHandler{} }
func (*CryptoSetHandler) GetMessageName() string { return "crypto_set_handler" }
func (*CryptoSetHandler) GetCrcString() string   { return "ce9ad00d" }
func (*CryptoSetHandler) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CryptoSetHandler) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 32 // m.AlgName
	size += 16 // m.Engine
	size += 1  // m.Oct
	size += 1  // m.IsAsync
	return size
}
func (m *CryptoSetHandler) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.AlgName, 32)
	buf.EncodeString(m.Engine, 16)
	buf.EncodeUint8(uint8(m.Oct))
	buf.EncodeUint8(m.IsAsync)
	return buf.Bytes(), nil
}
func (m *CryptoSetHandler) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.AlgName = buf.DecodeString(32)
	m.Engine = buf.DecodeString(16)
	m.Oct = CryptoOpClassType(buf.DecodeUint8())
	m.IsAsync = buf.DecodeUint8()
	return nil
}

// CryptoSetHandlerReply defines message 'crypto_set_handler_reply'.
type CryptoSetHandlerReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CryptoSetHandlerReply) Reset()               { *m = CryptoSetHandlerReply{} }
func (*CryptoSetHandlerReply) GetMessageName() string { return "crypto_set_handler_reply" }
func (*CryptoSetHandlerReply) GetCrcString() string   { return "e8d4e804" }
func (*CryptoSetHandlerReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CryptoSetHandlerReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CryptoSetHandlerReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CryptoSetHandlerReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_crypto_binapi_init() }
func file_crypto_binapi_init() {
	api.RegisterMessage((*CryptoSetAsyncDispatch)(nil), "crypto_set_async_dispatch_5ca4adc0")
	api.RegisterMessage((*CryptoSetAsyncDispatchReply)(nil), "crypto_set_async_dispatch_reply_e8d4e804")
	api.RegisterMessage((*CryptoSetHandler)(nil), "crypto_set_handler_ce9ad00d")
	api.RegisterMessage((*CryptoSetHandlerReply)(nil), "crypto_set_handler_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CryptoSetAsyncDispatch)(nil),
		(*CryptoSetAsyncDispatchReply)(nil),
		(*CryptoSetHandler)(nil),
		(*CryptoSetHandlerReply)(nil),
	}
}
