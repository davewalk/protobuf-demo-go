// Code generated by protoc-gen-go.
// source: cat.proto
// DO NOT EDIT!

/*
Package demo is a generated protocol buffer package.

It is generated from these files:
	cat.proto
	dog.proto

It has these top-level messages:
	Cat
*/
package demo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Cat struct {
	Name    string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age     int32         `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Parents []*Cat_Parent `protobuf:"bytes,4,rep,name=parents" json:"parents,omitempty"`
}

func (m *Cat) Reset()         { *m = Cat{} }
func (m *Cat) String() string { return proto.CompactTextString(m) }
func (*Cat) ProtoMessage()    {}

func (m *Cat) GetParents() []*Cat_Parent {
	if m != nil {
		return m.Parents
	}
	return nil
}

type Cat_Parent struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Cat_Parent) Reset()         { *m = Cat_Parent{} }
func (m *Cat_Parent) String() string { return proto.CompactTextString(m) }
func (*Cat_Parent) ProtoMessage()    {}
