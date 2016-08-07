// Code generated by protoc-gen-go.
// source: server/pfs/db/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pfs/db/persist/persist.proto

It has these top-level messages:
	Clock
	ClockID
	BranchClock
	Repo
	BlockRef
	Diff
	Commit
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import _ "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileType int32

const (
	FileType_NONE FileType = 0
	FileType_FILE FileType = 1
	FileType_DIR  FileType = 2
)

var FileType_name = map[int32]string{
	0: "NONE",
	1: "FILE",
	2: "DIR",
}
var FileType_value = map[string]int32{
	"NONE": 0,
	"FILE": 1,
	"DIR":  2,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}
func (FileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Clock struct {
	// a document either has these two fields
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,2,opt,name=clock" json:"clock,omitempty"`
}

func (m *Clock) Reset()                    { *m = Clock{} }
func (m *Clock) String() string            { return proto.CompactTextString(m) }
func (*Clock) ProtoMessage()               {}
func (*Clock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClockID struct {
	ID     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Branch string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,4,opt,name=clock" json:"clock,omitempty"`
}

func (m *ClockID) Reset()                    { *m = ClockID{} }
func (m *ClockID) String() string            { return proto.CompactTextString(m) }
func (*ClockID) ProtoMessage()               {}
func (*ClockID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BranchClock struct {
	Clocks []*Clock `protobuf:"bytes,1,rep,name=clocks" json:"clocks,omitempty"`
}

func (m *BranchClock) Reset()                    { *m = BranchClock{} }
func (m *BranchClock) String() string            { return proto.CompactTextString(m) }
func (*BranchClock) ProtoMessage()               {}
func (*BranchClock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BranchClock) GetClocks() []*Clock {
	if m != nil {
		return m.Clocks
	}
	return nil
}

type Repo struct {
	Name    string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	Size    uint64                     `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Repo) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type BlockRef struct {
	Hash  string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Lower uint64 `protobuf:"varint,2,opt,name=lower" json:"lower,omitempty"`
	Upper uint64 `protobuf:"varint,3,opt,name=upper" json:"upper,omitempty"`
}

func (m *BlockRef) Reset()                    { *m = BlockRef{} }
func (m *BlockRef) String() string            { return proto.CompactTextString(m) }
func (*BlockRef) ProtoMessage()               {}
func (*BlockRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Diff struct {
	ID       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo     string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	CommitID string `protobuf:"bytes,3,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	Path     string `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	// block_refs and delete cannot both be set
	BlockRefs    []*BlockRef                `protobuf:"bytes,5,rep,name=block_refs,json=blockRefs" json:"block_refs,omitempty"`
	Delete       bool                       `protobuf:"varint,6,opt,name=delete" json:"delete,omitempty"`
	Size         uint64                     `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	BranchClocks []*BranchClock             `protobuf:"bytes,8,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	FileType     FileType                   `protobuf:"varint,9,opt,name=file_type,json=fileType,enum=FileType" json:"file_type,omitempty"`
	Modified     *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=modified" json:"modified,omitempty"`
}

func (m *Diff) Reset()                    { *m = Diff{} }
func (m *Diff) String() string            { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()               {}
func (*Diff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Diff) GetBlockRefs() []*BlockRef {
	if m != nil {
		return m.BlockRefs
	}
	return nil
}

func (m *Diff) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Diff) GetModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Commit struct {
	ID           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo         string                     `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	BranchClocks []*BranchClock             `protobuf:"bytes,3,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	Started      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started" json:"started,omitempty"`
	Finished     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=finished" json:"finished,omitempty"`
	Cancelled    bool                       `protobuf:"varint,6,opt,name=cancelled" json:"cancelled,omitempty"`
	Provenance   []string                   `protobuf:"bytes,7,rep,name=provenance" json:"provenance,omitempty"`
	Size         uint64                     `protobuf:"varint,8,opt,name=size" json:"size,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Commit) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Commit) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Commit) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func init() {
	proto.RegisterType((*Clock)(nil), "Clock")
	proto.RegisterType((*ClockID)(nil), "ClockID")
	proto.RegisterType((*BranchClock)(nil), "BranchClock")
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*BlockRef)(nil), "BlockRef")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Commit)(nil), "Commit")
	proto.RegisterEnum("FileType", FileType_name, FileType_value)
}

func init() { proto.RegisterFile("server/pfs/db/persist/persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0x6a, 0xdb, 0x40,
	0x10, 0xad, 0x6d, 0x59, 0x96, 0xc6, 0x69, 0x30, 0x4b, 0x29, 0x26, 0x2d, 0x69, 0x51, 0xa1, 0x35,
	0x85, 0xca, 0x34, 0xbd, 0x7c, 0x40, 0xe2, 0x04, 0x5c, 0x4a, 0x0a, 0x4b, 0xde, 0xfa, 0x60, 0x24,
	0x6b, 0x37, 0x5e, 0x2a, 0x59, 0x62, 0x57, 0x49, 0x68, 0x7f, 0xa6, 0x3f, 0xd3, 0x0f, 0xeb, 0xec,
	0x68, 0x65, 0x9b, 0x36, 0xe0, 0x3c, 0x79, 0xe6, 0xe8, 0xcc, 0xed, 0x9c, 0x35, 0xbc, 0x32, 0x42,
	0xdf, 0x0a, 0x3d, 0xad, 0xa4, 0x99, 0x66, 0xe9, 0xb4, 0x12, 0xda, 0x28, 0x53, 0xb7, 0xbf, 0x71,
	0xa5, 0xcb, 0xba, 0x3c, 0x7a, 0x71, 0x5d, 0x96, 0xd7, 0xb9, 0x98, 0x52, 0x96, 0xde, 0xc8, 0x69,
	0xad, 0x0a, 0x61, 0xea, 0xa4, 0xa8, 0x1c, 0xe1, 0xf8, 0x5f, 0xc2, 0x9d, 0x4e, 0x2a, 0xdb, 0xa3,
	0xf9, 0x1e, 0x7d, 0x82, 0xfe, 0x59, 0x5e, 0x2e, 0x7f, 0xb0, 0xa7, 0xe0, 0xa7, 0x3a, 0x59, 0x2f,
	0x57, 0xe3, 0xce, 0xcb, 0xce, 0x24, 0xe4, 0x2e, 0x63, 0x4f, 0xa0, 0xbf, 0xb4, 0x84, 0x71, 0x17,
	0x61, 0x8f, 0x37, 0x49, 0xf4, 0x1d, 0x06, 0x54, 0x36, 0x9f, 0xb1, 0x43, 0xe8, 0xaa, 0xcc, 0x15,
	0x61, 0xc4, 0x18, 0x78, 0x5a, 0x54, 0x25, 0xf1, 0x43, 0x4e, 0xf1, 0x4e, 0xf3, 0xde, 0xfd, 0xcd,
	0xbd, 0xdd, 0xe6, 0xef, 0x60, 0x78, 0x4a, 0xdf, 0x9b, 0xcd, 0x8e, 0xc1, 0x27, 0xdc, 0xe0, 0x90,
	0xde, 0x64, 0x78, 0xe2, 0xc7, 0x84, 0x73, 0x87, 0x46, 0x19, 0x78, 0xdc, 0x0e, 0xc1, 0xc1, 0xeb,
	0xa4, 0x10, 0x6e, 0x15, 0x8a, 0xd9, 0x47, 0x18, 0x2c, 0xb5, 0x48, 0x6a, 0x91, 0xd1, 0x3e, 0xc3,
	0x93, 0xa3, 0xb8, 0x11, 0x24, 0x6e, 0x05, 0x89, 0xaf, 0x5a, 0xc5, 0x78, 0x4b, 0xb5, 0x9d, 0x8c,
	0xfa, 0x25, 0x68, 0x59, 0x8f, 0x53, 0x1c, 0x7d, 0x81, 0xe0, 0x94, 0xc6, 0x0a, 0x69, 0xbf, 0xaf,
	0x12, 0xd3, 0x2a, 0x45, 0xb1, 0x3d, 0x25, 0x2f, 0xef, 0x84, 0x6e, 0x75, 0xa2, 0xc4, 0xa2, 0x37,
	0x56, 0x6e, 0xd7, 0xaa, 0x49, 0xa2, 0x3f, 0x5d, 0xf0, 0x66, 0x4a, 0xca, 0x07, 0x69, 0xf7, 0x0c,
	0xc2, 0x65, 0x59, 0x14, 0xaa, 0x5e, 0x20, 0xb5, 0x91, 0x2f, 0x68, 0x80, 0x39, 0x15, 0x54, 0x49,
	0xbd, 0x22, 0xfd, 0xb0, 0xc0, 0xc6, 0x6c, 0x02, 0x90, 0xda, 0x4d, 0x17, 0x5a, 0x48, 0x33, 0xee,
	0x93, 0x66, 0x61, 0xdc, 0x2e, 0xcf, 0xc3, 0xd4, 0x45, 0xc6, 0xda, 0x92, 0x89, 0x5c, 0xd4, 0x62,
	0xec, 0x63, 0x7d, 0xc0, 0x5d, 0xb6, 0xb9, 0x7f, 0xb0, 0xbd, 0x9f, 0xbd, 0x87, 0xc7, 0x8d, 0x69,
	0x0b, 0x67, 0x46, 0x40, 0x8d, 0x0f, 0xe2, 0x1d, 0xab, 0xf8, 0x41, 0xba, 0x4d, 0x0c, 0x7b, 0x0d,
	0xa1, 0x54, 0xb9, 0x58, 0xd4, 0x3f, 0x2b, 0x31, 0x0e, 0xb1, 0xd7, 0x21, 0xee, 0x71, 0x81, 0xc8,
	0x15, 0x02, 0x3c, 0x90, 0x2e, 0x62, 0x9f, 0x21, 0x28, 0xca, 0x4c, 0x49, 0x85, 0x2e, 0xc1, 0x5e,
	0x97, 0x36, 0xdc, 0xe8, 0x77, 0x17, 0xfc, 0x33, 0x52, 0xe2, 0x41, 0x42, 0xfe, 0x77, 0x41, 0x6f,
	0xef, 0x05, 0xf8, 0x7c, 0x70, 0xa8, 0xb6, 0xcf, 0xc7, 0xdb, 0xff, 0x7c, 0x1c, 0xd5, 0xde, 0x23,
	0xd5, 0x5a, 0x99, 0x15, 0x96, 0xf5, 0xf7, 0xdf, 0xd3, 0x72, 0xd9, 0x73, 0x74, 0x1a, 0x87, 0x8b,
	0x3c, 0xc7, 0xc2, 0xc6, 0x91, 0x2d, 0x80, 0x7f, 0x03, 0xc0, 0xea, 0x5b, 0xb1, 0xb6, 0x08, 0x5a,
	0xd3, 0xc3, 0xc3, 0x76, 0x90, 0x8d, 0x69, 0xc1, 0xd6, 0xb4, 0xb7, 0x6f, 0x20, 0x68, 0xf5, 0x66,
	0x01, 0x78, 0x97, 0xdf, 0x2e, 0xcf, 0x47, 0x8f, 0x6c, 0x74, 0x31, 0xff, 0x7a, 0x3e, 0xea, 0xb0,
	0x01, 0xf4, 0x66, 0x73, 0x3e, 0xea, 0xa6, 0x3e, 0x2d, 0xf6, 0xe1, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0x4c, 0xe5, 0x3d, 0x75, 0x04, 0x00, 0x00,
}