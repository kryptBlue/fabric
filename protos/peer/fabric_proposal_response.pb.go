// Code generated by protoc-gen-go.
// source: peer/fabric_proposal_response.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response2 `protobuf:"bytes,4,opt,name=response" json:"response,omitempty"`
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement *Endorsement `protobuf:"bytes,6,opt,name=endorsement" json:"endorsement,omitempty"`
}

func (m *ProposalResponse) Reset()                    { *m = ProposalResponse{} }
func (m *ProposalResponse) String() string            { return proto.CompactTextString(m) }
func (*ProposalResponse) ProtoMessage()               {}
func (*ProposalResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *ProposalResponse) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ProposalResponse) GetResponse() *Response2 {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response2 struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Response2) Reset()                    { *m = Response2{} }
func (m *Response2) String() string            { return proto.CompactTextString(m) }
func (*Response2) ProtoMessage()               {}
func (*Response2) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte). However this implies that the hash can only be verified
	// if the entire proposal message is available when ProposalResponsePayload is
	// included in a transaction or stored in the ledger. For confidentiality
	// reasons, with chaincodes it might be undesirable to store the proposal
	// payload in the ledger.  If the type is CHAINCODE, this is handled by
	// separating the proposal's header and
	// the payload: the header is always hashed in its entirety whereas the
	// payload can either be hashed fully, or only its hash may be hashed, or
	// nothing from the payload can be hashed. The PayloadVisibility field in the
	// Header's extension controls to which extent the proposal payload is
	// "visible" in the sense that was just explained.
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposalHash,proto3" json:"proposalHash,omitempty"`
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch []byte `protobuf:"bytes,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension []byte `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *ProposalResponsePayload) Reset()                    { *m = ProposalResponsePayload{} }
func (m *ProposalResponsePayload) String() string            { return proto.CompactTextString(m) }
func (*ProposalResponsePayload) ProtoMessage()               {}
func (*ProposalResponsePayload) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	// Identity of the endorser (e.g. its certificate)
	Endorser []byte `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Endorsement) Reset()                    { *m = Endorsement{} }
func (m *Endorsement) String() string            { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()               {}
func (*Endorsement) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func init() {
	proto.RegisterType((*ProposalResponse)(nil), "protos.ProposalResponse")
	proto.RegisterType((*Response2)(nil), "protos.Response2")
	proto.RegisterType((*ProposalResponsePayload)(nil), "protos.ProposalResponsePayload")
	proto.RegisterType((*Endorsement)(nil), "protos.Endorsement")
}

func init() { proto.RegisterFile("peer/fabric_proposal_response.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0x41, 0x6b, 0xe3, 0x30,
	0x10, 0x85, 0xf1, 0xee, 0x26, 0x1b, 0x2b, 0x3e, 0xec, 0x6a, 0x97, 0xd6, 0x84, 0x42, 0x83, 0x7b,
	0x49, 0x29, 0xb5, 0x21, 0xa5, 0xd0, 0x73, 0xa1, 0xb4, 0xc7, 0x20, 0x0a, 0x85, 0x5e, 0x82, 0x9c,
	0x4c, 0x6c, 0x83, 0x6d, 0xa9, 0x1a, 0xb9, 0x34, 0x7f, 0xb8, 0xbf, 0xa3, 0x58, 0x96, 0xec, 0xa4,
	0x27, 0xf3, 0xc6, 0xa3, 0xf7, 0xbd, 0x91, 0x86, 0x5c, 0x48, 0x00, 0x95, 0xec, 0x78, 0xaa, 0x8a,
	0xcd, 0x5a, 0x2a, 0x21, 0x05, 0xf2, 0x72, 0xad, 0x00, 0xa5, 0xa8, 0x11, 0x62, 0xa9, 0x84, 0x16,
	0x74, 0x6c, 0x3e, 0x38, 0x3b, 0xcf, 0x84, 0xc8, 0x4a, 0x48, 0x8c, 0x4c, 0x9b, 0x5d, 0xa2, 0x8b,
	0x0a, 0x50, 0xf3, 0x4a, 0x76, 0x8d, 0xd1, 0xa7, 0x47, 0xfe, 0xac, 0xac, 0x09, 0xb3, 0x1e, 0x34,
	0x24, 0xbf, 0xdf, 0x41, 0x61, 0x21, 0xea, 0xd0, 0x9b, 0x7b, 0x8b, 0x11, 0x73, 0x92, 0xde, 0x11,
	0xbf, 0x77, 0x08, 0x7f, 0xcc, 0xbd, 0xc5, 0x74, 0x39, 0x8b, 0x3b, 0x46, 0xec, 0x18, 0xf1, 0xb3,
	0xeb, 0x60, 0x43, 0x33, 0xbd, 0x26, 0x13, 0x97, 0x31, 0xfc, 0x65, 0x0e, 0xfe, 0xed, 0x4e, 0x60,
	0xec, 0xb8, 0x4b, 0xd6, 0xb7, 0xb4, 0x11, 0x24, 0xdf, 0x97, 0x82, 0x6f, 0xc3, 0xd1, 0xdc, 0x5b,
	0x04, 0xcc, 0x49, 0x7a, 0x4b, 0xa6, 0x50, 0x6f, 0x85, 0x42, 0xa8, 0xa0, 0xd6, 0xe1, 0xd8, 0x78,
	0xfd, 0x73, 0x5e, 0x0f, 0xc3, 0x2f, 0x76, 0xd8, 0x17, 0xbd, 0x10, 0xbf, 0xe7, 0xd0, 0x13, 0x32,
	0x46, 0xcd, 0x75, 0x83, 0x76, 0x3e, 0xab, 0x5a, 0x6a, 0x05, 0x88, 0x3c, 0x03, 0x33, 0x9c, 0xcf,
	0x9c, 0x3c, 0xcc, 0xf3, 0xf3, 0x28, 0x4f, 0xf4, 0x46, 0x4e, 0xbf, 0x5f, 0xe0, 0xca, 0x46, 0x8d,
	0x48, 0xe0, 0x1e, 0xe8, 0x89, 0x63, 0x6e, 0x60, 0x01, 0x3b, 0xaa, 0xd1, 0xff, 0x64, 0x04, 0x52,
	0x6c, 0x72, 0x03, 0x0c, 0x58, 0x27, 0xe8, 0x19, 0xf1, 0xe1, 0x43, 0x43, 0x6d, 0xde, 0xa0, 0x03,
	0x0e, 0x85, 0xe8, 0x91, 0x4c, 0x0f, 0xe6, 0xa4, 0x33, 0x32, 0xb1, 0x93, 0x2a, 0x8b, 0xe8, 0x75,
	0x6b, 0x84, 0x45, 0x56, 0x73, 0xdd, 0x28, 0xb0, 0x88, 0xa1, 0x70, 0x7f, 0xf5, 0x7a, 0x99, 0x15,
	0x3a, 0x6f, 0xd2, 0x78, 0x23, 0xaa, 0x24, 0xdf, 0x4b, 0x50, 0x25, 0x6c, 0xb3, 0x7e, 0xbf, 0xba,
	0xbd, 0xc1, 0xa4, 0x5d, 0xb9, 0xb4, 0xdb, 0xa9, 0x9b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73,
	0xae, 0x6f, 0xc0, 0x81, 0x02, 0x00, 0x00,
}
