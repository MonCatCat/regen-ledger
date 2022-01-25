package ecocreditv1alpha2

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_ClassInfo_3_list)(nil)

type _ClassInfo_3_list struct {
	list *[]string
}

func (x *_ClassInfo_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ClassInfo_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_ClassInfo_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_ClassInfo_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_ClassInfo_3_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message ClassInfo at list field Issuers as it is not of Message kind"))
}

func (x *_ClassInfo_3_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_ClassInfo_3_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_ClassInfo_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ClassInfo             protoreflect.MessageDescriptor
	fd_ClassInfo_class_id    protoreflect.FieldDescriptor
	fd_ClassInfo_admin       protoreflect.FieldDescriptor
	fd_ClassInfo_issuers     protoreflect.FieldDescriptor
	fd_ClassInfo_metadata    protoreflect.FieldDescriptor
	fd_ClassInfo_credit_type protoreflect.FieldDescriptor
	fd_ClassInfo_num_batches protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_ClassInfo = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("ClassInfo")
	fd_ClassInfo_class_id = md_ClassInfo.Fields().ByName("class_id")
	fd_ClassInfo_admin = md_ClassInfo.Fields().ByName("admin")
	fd_ClassInfo_issuers = md_ClassInfo.Fields().ByName("issuers")
	fd_ClassInfo_metadata = md_ClassInfo.Fields().ByName("metadata")
	fd_ClassInfo_credit_type = md_ClassInfo.Fields().ByName("credit_type")
	fd_ClassInfo_num_batches = md_ClassInfo.Fields().ByName("num_batches")
}

var _ protoreflect.Message = (*fastReflection_ClassInfo)(nil)

type fastReflection_ClassInfo ClassInfo

func (x *ClassInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ClassInfo)(x)
}

func (x *ClassInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ClassInfo_messageType fastReflection_ClassInfo_messageType
var _ protoreflect.MessageType = fastReflection_ClassInfo_messageType{}

type fastReflection_ClassInfo_messageType struct{}

func (x fastReflection_ClassInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ClassInfo)(nil)
}
func (x fastReflection_ClassInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_ClassInfo)
}
func (x fastReflection_ClassInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ClassInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ClassInfo) Type() protoreflect.MessageType {
	return _fastReflection_ClassInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ClassInfo) New() protoreflect.Message {
	return new(fastReflection_ClassInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ClassInfo) Interface() protoreflect.ProtoMessage {
	return (*ClassInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ClassInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_ClassInfo_class_id, value) {
			return
		}
	}
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_ClassInfo_admin, value) {
			return
		}
	}
	if len(x.Issuers) != 0 {
		value := protoreflect.ValueOfList(&_ClassInfo_3_list{list: &x.Issuers})
		if !f(fd_ClassInfo_issuers, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_ClassInfo_metadata, value) {
			return
		}
	}
	if x.CreditType != nil {
		value := protoreflect.ValueOfMessage(x.CreditType.ProtoReflect())
		if !f(fd_ClassInfo_credit_type, value) {
			return
		}
	}
	if x.NumBatches != uint64(0) {
		value := protoreflect.ValueOfUint64(x.NumBatches)
		if !f(fd_ClassInfo_num_batches, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ClassInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ClassInfo.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.ClassInfo.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha2.ClassInfo.issuers":
		return len(x.Issuers) != 0
	case "regen.ecocredit.v1alpha2.ClassInfo.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1alpha2.ClassInfo.credit_type":
		return x.CreditType != nil
	case "regen.ecocredit.v1alpha2.ClassInfo.num_batches":
		return x.NumBatches != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ClassInfo.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.ClassInfo.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha2.ClassInfo.issuers":
		x.Issuers = nil
	case "regen.ecocredit.v1alpha2.ClassInfo.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1alpha2.ClassInfo.credit_type":
		x.CreditType = nil
	case "regen.ecocredit.v1alpha2.ClassInfo.num_batches":
		x.NumBatches = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ClassInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.ClassInfo.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.ClassInfo.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.ClassInfo.issuers":
		if len(x.Issuers) == 0 {
			return protoreflect.ValueOfList(&_ClassInfo_3_list{})
		}
		listValue := &_ClassInfo_3_list{list: &x.Issuers}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha2.ClassInfo.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1alpha2.ClassInfo.credit_type":
		value := x.CreditType
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.ClassInfo.num_batches":
		value := x.NumBatches
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ClassInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ClassInfo.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.ClassInfo.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.ClassInfo.issuers":
		lv := value.List()
		clv := lv.(*_ClassInfo_3_list)
		x.Issuers = *clv.list
	case "regen.ecocredit.v1alpha2.ClassInfo.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1alpha2.ClassInfo.credit_type":
		x.CreditType = value.Message().Interface().(*CreditType)
	case "regen.ecocredit.v1alpha2.ClassInfo.num_batches":
		x.NumBatches = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ClassInfo.issuers":
		if x.Issuers == nil {
			x.Issuers = []string{}
		}
		value := &_ClassInfo_3_list{list: &x.Issuers}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.ClassInfo.credit_type":
		if x.CreditType == nil {
			x.CreditType = new(CreditType)
		}
		return protoreflect.ValueOfMessage(x.CreditType.ProtoReflect())
	case "regen.ecocredit.v1alpha2.ClassInfo.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.ClassInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.ClassInfo.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha2.ClassInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.ClassInfo.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha2.ClassInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.ClassInfo.num_batches":
		panic(fmt.Errorf("field num_batches of message regen.ecocredit.v1alpha2.ClassInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ClassInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ClassInfo.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.ClassInfo.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.ClassInfo.issuers":
		list := []string{}
		return protoreflect.ValueOfList(&_ClassInfo_3_list{list: &list})
	case "regen.ecocredit.v1alpha2.ClassInfo.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1alpha2.ClassInfo.credit_type":
		m := new(CreditType)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.ClassInfo.num_batches":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ClassInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ClassInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ClassInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.ClassInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ClassInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ClassInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ClassInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ClassInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Issuers) > 0 {
			for _, s := range x.Issuers {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.CreditType != nil {
			l = options.Size(x.CreditType)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.NumBatches != 0 {
			n += 1 + runtime.Sov(uint64(x.NumBatches))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ClassInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.NumBatches != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.NumBatches))
			i--
			dAtA[i] = 0x30
		}
		if x.CreditType != nil {
			encoded, err := options.Marshal(x.CreditType)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Issuers) > 0 {
			for iNdEx := len(x.Issuers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Issuers[iNdEx])
				copy(dAtA[i:], x.Issuers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuers[iNdEx])))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ClassInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuers", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Issuers = append(x.Issuers, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreditType", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.CreditType == nil {
					x.CreditType = &CreditType{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CreditType); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NumBatches", wireType)
				}
				x.NumBatches = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.NumBatches |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_ProjectInfo                  protoreflect.MessageDescriptor
	fd_ProjectInfo_project_id       protoreflect.FieldDescriptor
	fd_ProjectInfo_class_id         protoreflect.FieldDescriptor
	fd_ProjectInfo_issuer           protoreflect.FieldDescriptor
	fd_ProjectInfo_project_location protoreflect.FieldDescriptor
	fd_ProjectInfo_metadata         protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_ProjectInfo = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("ProjectInfo")
	fd_ProjectInfo_project_id = md_ProjectInfo.Fields().ByName("project_id")
	fd_ProjectInfo_class_id = md_ProjectInfo.Fields().ByName("class_id")
	fd_ProjectInfo_issuer = md_ProjectInfo.Fields().ByName("issuer")
	fd_ProjectInfo_project_location = md_ProjectInfo.Fields().ByName("project_location")
	fd_ProjectInfo_metadata = md_ProjectInfo.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_ProjectInfo)(nil)

type fastReflection_ProjectInfo ProjectInfo

func (x *ProjectInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ProjectInfo)(x)
}

func (x *ProjectInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ProjectInfo_messageType fastReflection_ProjectInfo_messageType
var _ protoreflect.MessageType = fastReflection_ProjectInfo_messageType{}

type fastReflection_ProjectInfo_messageType struct{}

func (x fastReflection_ProjectInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ProjectInfo)(nil)
}
func (x fastReflection_ProjectInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_ProjectInfo)
}
func (x fastReflection_ProjectInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ProjectInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ProjectInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_ProjectInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ProjectInfo) Type() protoreflect.MessageType {
	return _fastReflection_ProjectInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ProjectInfo) New() protoreflect.Message {
	return new(fastReflection_ProjectInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ProjectInfo) Interface() protoreflect.ProtoMessage {
	return (*ProjectInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ProjectInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProjectId != "" {
		value := protoreflect.ValueOfString(x.ProjectId)
		if !f(fd_ProjectInfo_project_id, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_ProjectInfo_class_id, value) {
			return
		}
	}
	if x.Issuer != "" {
		value := protoreflect.ValueOfString(x.Issuer)
		if !f(fd_ProjectInfo_issuer, value) {
			return
		}
	}
	if x.ProjectLocation != "" {
		value := protoreflect.ValueOfString(x.ProjectLocation)
		if !f(fd_ProjectInfo_project_location, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_ProjectInfo_metadata, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ProjectInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_id":
		return x.ProjectId != ""
	case "regen.ecocredit.v1alpha2.ProjectInfo.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha2.ProjectInfo.issuer":
		return x.Issuer != ""
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_location":
		return x.ProjectLocation != ""
	case "regen.ecocredit.v1alpha2.ProjectInfo.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_id":
		x.ProjectId = ""
	case "regen.ecocredit.v1alpha2.ProjectInfo.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha2.ProjectInfo.issuer":
		x.Issuer = ""
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_location":
		x.ProjectLocation = ""
	case "regen.ecocredit.v1alpha2.ProjectInfo.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ProjectInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.ProjectInfo.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.ProjectInfo.issuer":
		value := x.Issuer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_location":
		value := x.ProjectLocation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.ProjectInfo.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ProjectInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_id":
		x.ProjectId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.ProjectInfo.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.ProjectInfo.issuer":
		x.Issuer = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_location":
		x.ProjectLocation = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.ProjectInfo.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1alpha2.ProjectInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.ProjectInfo.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.ProjectInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.ProjectInfo.issuer":
		panic(fmt.Errorf("field issuer of message regen.ecocredit.v1alpha2.ProjectInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.v1alpha2.ProjectInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.ProjectInfo.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha2.ProjectInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ProjectInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.ProjectInfo.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.ProjectInfo.issuer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.ProjectInfo.project_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.ProjectInfo.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.ProjectInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.ProjectInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ProjectInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.ProjectInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ProjectInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ProjectInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ProjectInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ProjectInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ProjectId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Issuer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProjectLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ProjectInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.ProjectLocation) > 0 {
			i -= len(x.ProjectLocation)
			copy(dAtA[i:], x.ProjectLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Issuer) > 0 {
			i -= len(x.Issuer)
			copy(dAtA[i:], x.Issuer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuer)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ProjectId) > 0 {
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ProjectInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProjectInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProjectInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ClassId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Issuer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectLocation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectLocation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_BatchInfo                  protoreflect.MessageDescriptor
	fd_BatchInfo_project_id       protoreflect.FieldDescriptor
	fd_BatchInfo_batch_denom      protoreflect.FieldDescriptor
	fd_BatchInfo_total_amount     protoreflect.FieldDescriptor
	fd_BatchInfo_metadata         protoreflect.FieldDescriptor
	fd_BatchInfo_amount_cancelled protoreflect.FieldDescriptor
	fd_BatchInfo_start_date       protoreflect.FieldDescriptor
	fd_BatchInfo_end_date         protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_BatchInfo = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("BatchInfo")
	fd_BatchInfo_project_id = md_BatchInfo.Fields().ByName("project_id")
	fd_BatchInfo_batch_denom = md_BatchInfo.Fields().ByName("batch_denom")
	fd_BatchInfo_total_amount = md_BatchInfo.Fields().ByName("total_amount")
	fd_BatchInfo_metadata = md_BatchInfo.Fields().ByName("metadata")
	fd_BatchInfo_amount_cancelled = md_BatchInfo.Fields().ByName("amount_cancelled")
	fd_BatchInfo_start_date = md_BatchInfo.Fields().ByName("start_date")
	fd_BatchInfo_end_date = md_BatchInfo.Fields().ByName("end_date")
}

var _ protoreflect.Message = (*fastReflection_BatchInfo)(nil)

type fastReflection_BatchInfo BatchInfo

func (x *BatchInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BatchInfo)(x)
}

func (x *BatchInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BatchInfo_messageType fastReflection_BatchInfo_messageType
var _ protoreflect.MessageType = fastReflection_BatchInfo_messageType{}

type fastReflection_BatchInfo_messageType struct{}

func (x fastReflection_BatchInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BatchInfo)(nil)
}
func (x fastReflection_BatchInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_BatchInfo)
}
func (x fastReflection_BatchInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BatchInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BatchInfo) Type() protoreflect.MessageType {
	return _fastReflection_BatchInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BatchInfo) New() protoreflect.Message {
	return new(fastReflection_BatchInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BatchInfo) Interface() protoreflect.ProtoMessage {
	return (*BatchInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BatchInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProjectId != "" {
		value := protoreflect.ValueOfString(x.ProjectId)
		if !f(fd_BatchInfo_project_id, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_BatchInfo_batch_denom, value) {
			return
		}
	}
	if x.TotalAmount != "" {
		value := protoreflect.ValueOfString(x.TotalAmount)
		if !f(fd_BatchInfo_total_amount, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_BatchInfo_metadata, value) {
			return
		}
	}
	if x.AmountCancelled != "" {
		value := protoreflect.ValueOfString(x.AmountCancelled)
		if !f(fd_BatchInfo_amount_cancelled, value) {
			return
		}
	}
	if x.StartDate != nil {
		value := protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
		if !f(fd_BatchInfo_start_date, value) {
			return
		}
	}
	if x.EndDate != nil {
		value := protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
		if !f(fd_BatchInfo_end_date, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BatchInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BatchInfo.project_id":
		return x.ProjectId != ""
	case "regen.ecocredit.v1alpha2.BatchInfo.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.BatchInfo.total_amount":
		return x.TotalAmount != ""
	case "regen.ecocredit.v1alpha2.BatchInfo.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1alpha2.BatchInfo.amount_cancelled":
		return x.AmountCancelled != ""
	case "regen.ecocredit.v1alpha2.BatchInfo.start_date":
		return x.StartDate != nil
	case "regen.ecocredit.v1alpha2.BatchInfo.end_date":
		return x.EndDate != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BatchInfo.project_id":
		x.ProjectId = ""
	case "regen.ecocredit.v1alpha2.BatchInfo.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.BatchInfo.total_amount":
		x.TotalAmount = ""
	case "regen.ecocredit.v1alpha2.BatchInfo.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1alpha2.BatchInfo.amount_cancelled":
		x.AmountCancelled = ""
	case "regen.ecocredit.v1alpha2.BatchInfo.start_date":
		x.StartDate = nil
	case "regen.ecocredit.v1alpha2.BatchInfo.end_date":
		x.EndDate = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BatchInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.BatchInfo.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.BatchInfo.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.BatchInfo.total_amount":
		value := x.TotalAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.BatchInfo.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1alpha2.BatchInfo.amount_cancelled":
		value := x.AmountCancelled
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.BatchInfo.start_date":
		value := x.StartDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BatchInfo.end_date":
		value := x.EndDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BatchInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BatchInfo.project_id":
		x.ProjectId = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.BatchInfo.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.BatchInfo.total_amount":
		x.TotalAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.BatchInfo.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1alpha2.BatchInfo.amount_cancelled":
		x.AmountCancelled = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.BatchInfo.start_date":
		x.StartDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.v1alpha2.BatchInfo.end_date":
		x.EndDate = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BatchInfo.start_date":
		if x.StartDate == nil {
			x.StartDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BatchInfo.end_date":
		if x.EndDate == nil {
			x.EndDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BatchInfo.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1alpha2.BatchInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.BatchInfo.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.BatchInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.BatchInfo.total_amount":
		panic(fmt.Errorf("field total_amount of message regen.ecocredit.v1alpha2.BatchInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.BatchInfo.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha2.BatchInfo is not mutable"))
	case "regen.ecocredit.v1alpha2.BatchInfo.amount_cancelled":
		panic(fmt.Errorf("field amount_cancelled of message regen.ecocredit.v1alpha2.BatchInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BatchInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BatchInfo.project_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.BatchInfo.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.BatchInfo.total_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.BatchInfo.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1alpha2.BatchInfo.amount_cancelled":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.BatchInfo.start_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BatchInfo.end_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BatchInfo"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BatchInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BatchInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.BatchInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BatchInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BatchInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BatchInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BatchInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ProjectId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TotalAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AmountCancelled)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartDate != nil {
			l = options.Size(x.StartDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.EndDate != nil {
			l = options.Size(x.EndDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BatchInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.EndDate != nil {
			encoded, err := options.Marshal(x.EndDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x42
		}
		if x.StartDate != nil {
			encoded, err := options.Marshal(x.StartDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.AmountCancelled) > 0 {
			i -= len(x.AmountCancelled)
			copy(dAtA[i:], x.AmountCancelled)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AmountCancelled)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.TotalAmount) > 0 {
			i -= len(x.TotalAmount)
			copy(dAtA[i:], x.TotalAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TotalAmount)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ProjectId) > 0 {
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BatchInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TotalAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AmountCancelled", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AmountCancelled = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.StartDate == nil {
					x.StartDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.StartDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.EndDate == nil {
					x.EndDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.EndDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_Params_1_list)(nil)

type _Params_1_list struct {
	list *[]*v1beta1.Coin
}

func (x *_Params_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Params_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Params_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_Params_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Params_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Params_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Params_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Params_1_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_Params_2_list)(nil)

type _Params_2_list struct {
	list *[]string
}

func (x *_Params_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Params_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_Params_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Params_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Params_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Params at list field AllowedClassCreators as it is not of Message kind"))
}

func (x *_Params_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Params_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_Params_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_Params_4_list)(nil)

type _Params_4_list struct {
	list *[]*CreditType
}

func (x *_Params_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Params_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Params_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*CreditType)
	(*x.list)[i] = concreteValue
}

func (x *_Params_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*CreditType)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Params_4_list) AppendMutable() protoreflect.Value {
	v := new(CreditType)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Params_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Params_4_list) NewElement() protoreflect.Value {
	v := new(CreditType)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Params_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Params                        protoreflect.MessageDescriptor
	fd_Params_credit_class_fee       protoreflect.FieldDescriptor
	fd_Params_allowed_class_creators protoreflect.FieldDescriptor
	fd_Params_allowlist_enabled      protoreflect.FieldDescriptor
	fd_Params_credit_types           protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_Params = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("Params")
	fd_Params_credit_class_fee = md_Params.Fields().ByName("credit_class_fee")
	fd_Params_allowed_class_creators = md_Params.Fields().ByName("allowed_class_creators")
	fd_Params_allowlist_enabled = md_Params.Fields().ByName("allowlist_enabled")
	fd_Params_credit_types = md_Params.Fields().ByName("credit_types")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.CreditClassFee) != 0 {
		value := protoreflect.ValueOfList(&_Params_1_list{list: &x.CreditClassFee})
		if !f(fd_Params_credit_class_fee, value) {
			return
		}
	}
	if len(x.AllowedClassCreators) != 0 {
		value := protoreflect.ValueOfList(&_Params_2_list{list: &x.AllowedClassCreators})
		if !f(fd_Params_allowed_class_creators, value) {
			return
		}
	}
	if x.AllowlistEnabled != false {
		value := protoreflect.ValueOfBool(x.AllowlistEnabled)
		if !f(fd_Params_allowlist_enabled, value) {
			return
		}
	}
	if len(x.CreditTypes) != 0 {
		value := protoreflect.ValueOfList(&_Params_4_list{list: &x.CreditTypes})
		if !f(fd_Params_credit_types, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Params.credit_class_fee":
		return len(x.CreditClassFee) != 0
	case "regen.ecocredit.v1alpha2.Params.allowed_class_creators":
		return len(x.AllowedClassCreators) != 0
	case "regen.ecocredit.v1alpha2.Params.allowlist_enabled":
		return x.AllowlistEnabled != false
	case "regen.ecocredit.v1alpha2.Params.credit_types":
		return len(x.CreditTypes) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Params"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Params.credit_class_fee":
		x.CreditClassFee = nil
	case "regen.ecocredit.v1alpha2.Params.allowed_class_creators":
		x.AllowedClassCreators = nil
	case "regen.ecocredit.v1alpha2.Params.allowlist_enabled":
		x.AllowlistEnabled = false
	case "regen.ecocredit.v1alpha2.Params.credit_types":
		x.CreditTypes = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Params"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.Params.credit_class_fee":
		if len(x.CreditClassFee) == 0 {
			return protoreflect.ValueOfList(&_Params_1_list{})
		}
		listValue := &_Params_1_list{list: &x.CreditClassFee}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha2.Params.allowed_class_creators":
		if len(x.AllowedClassCreators) == 0 {
			return protoreflect.ValueOfList(&_Params_2_list{})
		}
		listValue := &_Params_2_list{list: &x.AllowedClassCreators}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha2.Params.allowlist_enabled":
		value := x.AllowlistEnabled
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.Params.credit_types":
		if len(x.CreditTypes) == 0 {
			return protoreflect.ValueOfList(&_Params_4_list{})
		}
		listValue := &_Params_4_list{list: &x.CreditTypes}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Params"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Params.credit_class_fee":
		lv := value.List()
		clv := lv.(*_Params_1_list)
		x.CreditClassFee = *clv.list
	case "regen.ecocredit.v1alpha2.Params.allowed_class_creators":
		lv := value.List()
		clv := lv.(*_Params_2_list)
		x.AllowedClassCreators = *clv.list
	case "regen.ecocredit.v1alpha2.Params.allowlist_enabled":
		x.AllowlistEnabled = value.Bool()
	case "regen.ecocredit.v1alpha2.Params.credit_types":
		lv := value.List()
		clv := lv.(*_Params_4_list)
		x.CreditTypes = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Params"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Params.credit_class_fee":
		if x.CreditClassFee == nil {
			x.CreditClassFee = []*v1beta1.Coin{}
		}
		value := &_Params_1_list{list: &x.CreditClassFee}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.Params.allowed_class_creators":
		if x.AllowedClassCreators == nil {
			x.AllowedClassCreators = []string{}
		}
		value := &_Params_2_list{list: &x.AllowedClassCreators}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.Params.credit_types":
		if x.CreditTypes == nil {
			x.CreditTypes = []*CreditType{}
		}
		value := &_Params_4_list{list: &x.CreditTypes}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha2.Params.allowlist_enabled":
		panic(fmt.Errorf("field allowlist_enabled of message regen.ecocredit.v1alpha2.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Params"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Params.credit_class_fee":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_Params_1_list{list: &list})
	case "regen.ecocredit.v1alpha2.Params.allowed_class_creators":
		list := []string{}
		return protoreflect.ValueOfList(&_Params_2_list{list: &list})
	case "regen.ecocredit.v1alpha2.Params.allowlist_enabled":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.Params.credit_types":
		list := []*CreditType{}
		return protoreflect.ValueOfList(&_Params_4_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Params"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.CreditClassFee) > 0 {
			for _, e := range x.CreditClassFee {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.AllowedClassCreators) > 0 {
			for _, s := range x.AllowedClassCreators {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.AllowlistEnabled {
			n += 2
		}
		if len(x.CreditTypes) > 0 {
			for _, e := range x.CreditTypes {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.CreditTypes) > 0 {
			for iNdEx := len(x.CreditTypes) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.CreditTypes[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x22
			}
		}
		if x.AllowlistEnabled {
			i--
			if x.AllowlistEnabled {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x18
		}
		if len(x.AllowedClassCreators) > 0 {
			for iNdEx := len(x.AllowedClassCreators) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.AllowedClassCreators[iNdEx])
				copy(dAtA[i:], x.AllowedClassCreators[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AllowedClassCreators[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.CreditClassFee) > 0 {
			for iNdEx := len(x.CreditClassFee) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.CreditClassFee[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreditClassFee", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CreditClassFee = append(x.CreditClassFee, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CreditClassFee[len(x.CreditClassFee)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AllowedClassCreators", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AllowedClassCreators = append(x.AllowedClassCreators, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AllowlistEnabled", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.AllowlistEnabled = bool(v != 0)
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreditTypes", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CreditTypes = append(x.CreditTypes, &CreditType{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CreditTypes[len(x.CreditTypes)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_CreditType              protoreflect.MessageDescriptor
	fd_CreditType_name         protoreflect.FieldDescriptor
	fd_CreditType_abbreviation protoreflect.FieldDescriptor
	fd_CreditType_unit         protoreflect.FieldDescriptor
	fd_CreditType_precision    protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_CreditType = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("CreditType")
	fd_CreditType_name = md_CreditType.Fields().ByName("name")
	fd_CreditType_abbreviation = md_CreditType.Fields().ByName("abbreviation")
	fd_CreditType_unit = md_CreditType.Fields().ByName("unit")
	fd_CreditType_precision = md_CreditType.Fields().ByName("precision")
}

var _ protoreflect.Message = (*fastReflection_CreditType)(nil)

type fastReflection_CreditType CreditType

func (x *CreditType) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CreditType)(x)
}

func (x *CreditType) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CreditType_messageType fastReflection_CreditType_messageType
var _ protoreflect.MessageType = fastReflection_CreditType_messageType{}

type fastReflection_CreditType_messageType struct{}

func (x fastReflection_CreditType_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CreditType)(nil)
}
func (x fastReflection_CreditType_messageType) New() protoreflect.Message {
	return new(fastReflection_CreditType)
}
func (x fastReflection_CreditType_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CreditType
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CreditType) Descriptor() protoreflect.MessageDescriptor {
	return md_CreditType
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CreditType) Type() protoreflect.MessageType {
	return _fastReflection_CreditType_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CreditType) New() protoreflect.Message {
	return new(fastReflection_CreditType)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CreditType) Interface() protoreflect.ProtoMessage {
	return (*CreditType)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CreditType) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_CreditType_name, value) {
			return
		}
	}
	if x.Abbreviation != "" {
		value := protoreflect.ValueOfString(x.Abbreviation)
		if !f(fd_CreditType_abbreviation, value) {
			return
		}
	}
	if x.Unit != "" {
		value := protoreflect.ValueOfString(x.Unit)
		if !f(fd_CreditType_unit, value) {
			return
		}
	}
	if x.Precision != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Precision)
		if !f(fd_CreditType_precision, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_CreditType) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditType.name":
		return x.Name != ""
	case "regen.ecocredit.v1alpha2.CreditType.abbreviation":
		return x.Abbreviation != ""
	case "regen.ecocredit.v1alpha2.CreditType.unit":
		return x.Unit != ""
	case "regen.ecocredit.v1alpha2.CreditType.precision":
		return x.Precision != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditType does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditType) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditType.name":
		x.Name = ""
	case "regen.ecocredit.v1alpha2.CreditType.abbreviation":
		x.Abbreviation = ""
	case "regen.ecocredit.v1alpha2.CreditType.unit":
		x.Unit = ""
	case "regen.ecocredit.v1alpha2.CreditType.precision":
		x.Precision = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditType does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CreditType) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.CreditType.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.CreditType.abbreviation":
		value := x.Abbreviation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.CreditType.unit":
		value := x.Unit
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.CreditType.precision":
		value := x.Precision
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditType does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditType) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditType.name":
		x.Name = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.CreditType.abbreviation":
		x.Abbreviation = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.CreditType.unit":
		x.Unit = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.CreditType.precision":
		x.Precision = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditType does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditType) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditType.name":
		panic(fmt.Errorf("field name of message regen.ecocredit.v1alpha2.CreditType is not mutable"))
	case "regen.ecocredit.v1alpha2.CreditType.abbreviation":
		panic(fmt.Errorf("field abbreviation of message regen.ecocredit.v1alpha2.CreditType is not mutable"))
	case "regen.ecocredit.v1alpha2.CreditType.unit":
		panic(fmt.Errorf("field unit of message regen.ecocredit.v1alpha2.CreditType is not mutable"))
	case "regen.ecocredit.v1alpha2.CreditType.precision":
		panic(fmt.Errorf("field precision of message regen.ecocredit.v1alpha2.CreditType is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditType does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CreditType) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditType.name":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.CreditType.abbreviation":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.CreditType.unit":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.CreditType.precision":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditType"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditType does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CreditType) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.CreditType", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CreditType) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditType) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_CreditType) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CreditType) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CreditType)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Abbreviation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Unit)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Precision != 0 {
			n += 1 + runtime.Sov(uint64(x.Precision))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*CreditType)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Precision != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Precision))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Unit) > 0 {
			i -= len(x.Unit)
			copy(dAtA[i:], x.Unit)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Unit)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Abbreviation) > 0 {
			i -= len(x.Abbreviation)
			copy(dAtA[i:], x.Abbreviation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Abbreviation)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*CreditType)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CreditType: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CreditType: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Abbreviation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Abbreviation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Unit = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
				}
				x.Precision = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Precision |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_CreditTypeSeq              protoreflect.MessageDescriptor
	fd_CreditTypeSeq_abbreviation protoreflect.FieldDescriptor
	fd_CreditTypeSeq_seq_number   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_CreditTypeSeq = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("CreditTypeSeq")
	fd_CreditTypeSeq_abbreviation = md_CreditTypeSeq.Fields().ByName("abbreviation")
	fd_CreditTypeSeq_seq_number = md_CreditTypeSeq.Fields().ByName("seq_number")
}

var _ protoreflect.Message = (*fastReflection_CreditTypeSeq)(nil)

type fastReflection_CreditTypeSeq CreditTypeSeq

func (x *CreditTypeSeq) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CreditTypeSeq)(x)
}

func (x *CreditTypeSeq) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CreditTypeSeq_messageType fastReflection_CreditTypeSeq_messageType
var _ protoreflect.MessageType = fastReflection_CreditTypeSeq_messageType{}

type fastReflection_CreditTypeSeq_messageType struct{}

func (x fastReflection_CreditTypeSeq_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CreditTypeSeq)(nil)
}
func (x fastReflection_CreditTypeSeq_messageType) New() protoreflect.Message {
	return new(fastReflection_CreditTypeSeq)
}
func (x fastReflection_CreditTypeSeq_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CreditTypeSeq
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CreditTypeSeq) Descriptor() protoreflect.MessageDescriptor {
	return md_CreditTypeSeq
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CreditTypeSeq) Type() protoreflect.MessageType {
	return _fastReflection_CreditTypeSeq_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CreditTypeSeq) New() protoreflect.Message {
	return new(fastReflection_CreditTypeSeq)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CreditTypeSeq) Interface() protoreflect.ProtoMessage {
	return (*CreditTypeSeq)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CreditTypeSeq) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Abbreviation != "" {
		value := protoreflect.ValueOfString(x.Abbreviation)
		if !f(fd_CreditTypeSeq_abbreviation, value) {
			return
		}
	}
	if x.SeqNumber != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SeqNumber)
		if !f(fd_CreditTypeSeq_seq_number, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_CreditTypeSeq) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.abbreviation":
		return x.Abbreviation != ""
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.seq_number":
		return x.SeqNumber != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditTypeSeq"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditTypeSeq does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditTypeSeq) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.abbreviation":
		x.Abbreviation = ""
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.seq_number":
		x.SeqNumber = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditTypeSeq"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditTypeSeq does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CreditTypeSeq) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.abbreviation":
		value := x.Abbreviation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.seq_number":
		value := x.SeqNumber
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditTypeSeq"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditTypeSeq does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditTypeSeq) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.abbreviation":
		x.Abbreviation = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.seq_number":
		x.SeqNumber = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditTypeSeq"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditTypeSeq does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditTypeSeq) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.abbreviation":
		panic(fmt.Errorf("field abbreviation of message regen.ecocredit.v1alpha2.CreditTypeSeq is not mutable"))
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.seq_number":
		panic(fmt.Errorf("field seq_number of message regen.ecocredit.v1alpha2.CreditTypeSeq is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditTypeSeq"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditTypeSeq does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CreditTypeSeq) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.abbreviation":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.CreditTypeSeq.seq_number":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.CreditTypeSeq"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.CreditTypeSeq does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CreditTypeSeq) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.CreditTypeSeq", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CreditTypeSeq) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CreditTypeSeq) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_CreditTypeSeq) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CreditTypeSeq) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CreditTypeSeq)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Abbreviation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.SeqNumber != 0 {
			n += 1 + runtime.Sov(uint64(x.SeqNumber))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*CreditTypeSeq)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.SeqNumber != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SeqNumber))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Abbreviation) > 0 {
			i -= len(x.Abbreviation)
			copy(dAtA[i:], x.Abbreviation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Abbreviation)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*CreditTypeSeq)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CreditTypeSeq: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CreditTypeSeq: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Abbreviation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Abbreviation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SeqNumber", wireType)
				}
				x.SeqNumber = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SeqNumber |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_SellOrder                     protoreflect.MessageDescriptor
	fd_SellOrder_order_id            protoreflect.FieldDescriptor
	fd_SellOrder_owner               protoreflect.FieldDescriptor
	fd_SellOrder_batch_denom         protoreflect.FieldDescriptor
	fd_SellOrder_quantity            protoreflect.FieldDescriptor
	fd_SellOrder_ask_price           protoreflect.FieldDescriptor
	fd_SellOrder_disable_auto_retire protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_SellOrder = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("SellOrder")
	fd_SellOrder_order_id = md_SellOrder.Fields().ByName("order_id")
	fd_SellOrder_owner = md_SellOrder.Fields().ByName("owner")
	fd_SellOrder_batch_denom = md_SellOrder.Fields().ByName("batch_denom")
	fd_SellOrder_quantity = md_SellOrder.Fields().ByName("quantity")
	fd_SellOrder_ask_price = md_SellOrder.Fields().ByName("ask_price")
	fd_SellOrder_disable_auto_retire = md_SellOrder.Fields().ByName("disable_auto_retire")
}

var _ protoreflect.Message = (*fastReflection_SellOrder)(nil)

type fastReflection_SellOrder SellOrder

func (x *SellOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SellOrder)(x)
}

func (x *SellOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SellOrder_messageType fastReflection_SellOrder_messageType
var _ protoreflect.MessageType = fastReflection_SellOrder_messageType{}

type fastReflection_SellOrder_messageType struct{}

func (x fastReflection_SellOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SellOrder)(nil)
}
func (x fastReflection_SellOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_SellOrder)
}
func (x fastReflection_SellOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SellOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SellOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_SellOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SellOrder) Type() protoreflect.MessageType {
	return _fastReflection_SellOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SellOrder) New() protoreflect.Message {
	return new(fastReflection_SellOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SellOrder) Interface() protoreflect.ProtoMessage {
	return (*SellOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SellOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.OrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.OrderId)
		if !f(fd_SellOrder_order_id, value) {
			return
		}
	}
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_SellOrder_owner, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_SellOrder_batch_denom, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_SellOrder_quantity, value) {
			return
		}
	}
	if x.AskPrice != nil {
		value := protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
		if !f(fd_SellOrder_ask_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_SellOrder_disable_auto_retire, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_SellOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.SellOrder.order_id":
		return x.OrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.SellOrder.owner":
		return x.Owner != ""
	case "regen.ecocredit.v1alpha2.SellOrder.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.SellOrder.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1alpha2.SellOrder.ask_price":
		return x.AskPrice != nil
	case "regen.ecocredit.v1alpha2.SellOrder.disable_auto_retire":
		return x.DisableAutoRetire != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.SellOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SellOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.SellOrder.order_id":
		x.OrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.SellOrder.owner":
		x.Owner = ""
	case "regen.ecocredit.v1alpha2.SellOrder.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.SellOrder.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1alpha2.SellOrder.ask_price":
		x.AskPrice = nil
	case "regen.ecocredit.v1alpha2.SellOrder.disable_auto_retire":
		x.DisableAutoRetire = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.SellOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SellOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.SellOrder.order_id":
		value := x.OrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.SellOrder.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.SellOrder.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.SellOrder.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.SellOrder.ask_price":
		value := x.AskPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.SellOrder.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.SellOrder does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SellOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.SellOrder.order_id":
		x.OrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.SellOrder.owner":
		x.Owner = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.SellOrder.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.SellOrder.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.SellOrder.ask_price":
		x.AskPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1alpha2.SellOrder.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.SellOrder does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SellOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.SellOrder.ask_price":
		if x.AskPrice == nil {
			x.AskPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.SellOrder.order_id":
		panic(fmt.Errorf("field order_id of message regen.ecocredit.v1alpha2.SellOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.SellOrder.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1alpha2.SellOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.SellOrder.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.SellOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.SellOrder.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1alpha2.SellOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.SellOrder.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.SellOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.SellOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SellOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.SellOrder.order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.SellOrder.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.SellOrder.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.SellOrder.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.SellOrder.ask_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.SellOrder.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.SellOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SellOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.SellOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SellOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SellOrder) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_SellOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SellOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SellOrder)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.OrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.OrderId))
		}
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AskPrice != nil {
			l = options.Size(x.AskPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*SellOrder)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x30
		}
		if x.AskPrice != nil {
			encoded, err := options.Marshal(x.AskPrice)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0x12
		}
		if x.OrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.OrderId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*SellOrder)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SellOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SellOrder: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
				}
				x.OrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.OrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AskPrice", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.AskPrice == nil {
					x.AskPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AskPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_BuyOrder                      protoreflect.MessageDescriptor
	fd_BuyOrder_buy_order_id         protoreflect.FieldDescriptor
	fd_BuyOrder_buyer                protoreflect.FieldDescriptor
	fd_BuyOrder_selection            protoreflect.FieldDescriptor
	fd_BuyOrder_quantity             protoreflect.FieldDescriptor
	fd_BuyOrder_bid_price            protoreflect.FieldDescriptor
	fd_BuyOrder_disable_auto_retire  protoreflect.FieldDescriptor
	fd_BuyOrder_disable_partial_fill protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_BuyOrder = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("BuyOrder")
	fd_BuyOrder_buy_order_id = md_BuyOrder.Fields().ByName("buy_order_id")
	fd_BuyOrder_buyer = md_BuyOrder.Fields().ByName("buyer")
	fd_BuyOrder_selection = md_BuyOrder.Fields().ByName("selection")
	fd_BuyOrder_quantity = md_BuyOrder.Fields().ByName("quantity")
	fd_BuyOrder_bid_price = md_BuyOrder.Fields().ByName("bid_price")
	fd_BuyOrder_disable_auto_retire = md_BuyOrder.Fields().ByName("disable_auto_retire")
	fd_BuyOrder_disable_partial_fill = md_BuyOrder.Fields().ByName("disable_partial_fill")
}

var _ protoreflect.Message = (*fastReflection_BuyOrder)(nil)

type fastReflection_BuyOrder BuyOrder

func (x *BuyOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BuyOrder)(x)
}

func (x *BuyOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BuyOrder_messageType fastReflection_BuyOrder_messageType
var _ protoreflect.MessageType = fastReflection_BuyOrder_messageType{}

type fastReflection_BuyOrder_messageType struct{}

func (x fastReflection_BuyOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BuyOrder)(nil)
}
func (x fastReflection_BuyOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_BuyOrder)
}
func (x fastReflection_BuyOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BuyOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BuyOrder) Type() protoreflect.MessageType {
	return _fastReflection_BuyOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BuyOrder) New() protoreflect.Message {
	return new(fastReflection_BuyOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BuyOrder) Interface() protoreflect.ProtoMessage {
	return (*BuyOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BuyOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BuyOrderId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BuyOrderId)
		if !f(fd_BuyOrder_buy_order_id, value) {
			return
		}
	}
	if x.Buyer != "" {
		value := protoreflect.ValueOfString(x.Buyer)
		if !f(fd_BuyOrder_buyer, value) {
			return
		}
	}
	if x.Selection != nil {
		value := protoreflect.ValueOfMessage(x.Selection.ProtoReflect())
		if !f(fd_BuyOrder_selection, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_BuyOrder_quantity, value) {
			return
		}
	}
	if x.BidPrice != nil {
		value := protoreflect.ValueOfMessage(x.BidPrice.ProtoReflect())
		if !f(fd_BuyOrder_bid_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_BuyOrder_disable_auto_retire, value) {
			return
		}
	}
	if x.DisablePartialFill != false {
		value := protoreflect.ValueOfBool(x.DisablePartialFill)
		if !f(fd_BuyOrder_disable_partial_fill, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BuyOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.buy_order_id":
		return x.BuyOrderId != uint64(0)
	case "regen.ecocredit.v1alpha2.BuyOrder.buyer":
		return x.Buyer != ""
	case "regen.ecocredit.v1alpha2.BuyOrder.selection":
		return x.Selection != nil
	case "regen.ecocredit.v1alpha2.BuyOrder.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.v1alpha2.BuyOrder.bid_price":
		return x.BidPrice != nil
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_partial_fill":
		return x.DisablePartialFill != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.buy_order_id":
		x.BuyOrderId = uint64(0)
	case "regen.ecocredit.v1alpha2.BuyOrder.buyer":
		x.Buyer = ""
	case "regen.ecocredit.v1alpha2.BuyOrder.selection":
		x.Selection = nil
	case "regen.ecocredit.v1alpha2.BuyOrder.quantity":
		x.Quantity = ""
	case "regen.ecocredit.v1alpha2.BuyOrder.bid_price":
		x.BidPrice = nil
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_partial_fill":
		x.DisablePartialFill = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BuyOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.buy_order_id":
		value := x.BuyOrderId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.v1alpha2.BuyOrder.buyer":
		value := x.Buyer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.BuyOrder.selection":
		value := x.Selection
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BuyOrder.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.BuyOrder.bid_price":
		value := x.BidPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_partial_fill":
		value := x.DisablePartialFill
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.buy_order_id":
		x.BuyOrderId = value.Uint()
	case "regen.ecocredit.v1alpha2.BuyOrder.buyer":
		x.Buyer = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.BuyOrder.selection":
		x.Selection = value.Message().Interface().(*BuyOrder_Selection)
	case "regen.ecocredit.v1alpha2.BuyOrder.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.BuyOrder.bid_price":
		x.BidPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_partial_fill":
		x.DisablePartialFill = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.selection":
		if x.Selection == nil {
			x.Selection = new(BuyOrder_Selection)
		}
		return protoreflect.ValueOfMessage(x.Selection.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BuyOrder.bid_price":
		if x.BidPrice == nil {
			x.BidPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.BidPrice.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BuyOrder.buy_order_id":
		panic(fmt.Errorf("field buy_order_id of message regen.ecocredit.v1alpha2.BuyOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.BuyOrder.buyer":
		panic(fmt.Errorf("field buyer of message regen.ecocredit.v1alpha2.BuyOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.BuyOrder.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.v1alpha2.BuyOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.v1alpha2.BuyOrder is not mutable"))
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_partial_fill":
		panic(fmt.Errorf("field disable_partial_fill of message regen.ecocredit.v1alpha2.BuyOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BuyOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.buy_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.v1alpha2.BuyOrder.buyer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.BuyOrder.selection":
		m := new(BuyOrder_Selection)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BuyOrder.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.BuyOrder.bid_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.v1alpha2.BuyOrder.disable_partial_fill":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BuyOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.BuyOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BuyOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BuyOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BuyOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BuyOrder)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BuyOrderId != 0 {
			n += 1 + runtime.Sov(uint64(x.BuyOrderId))
		}
		l = len(x.Buyer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Selection != nil {
			l = options.Size(x.Selection)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BidPrice != nil {
			l = options.Size(x.BidPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.DisablePartialFill {
			n += 2
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BuyOrder)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.DisablePartialFill {
			i--
			if x.DisablePartialFill {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x38
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x30
		}
		if x.BidPrice != nil {
			encoded, err := options.Marshal(x.BidPrice)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x22
		}
		if x.Selection != nil {
			encoded, err := options.Marshal(x.Selection)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Buyer) > 0 {
			i -= len(x.Buyer)
			copy(dAtA[i:], x.Buyer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Buyer)))
			i--
			dAtA[i] = 0x12
		}
		if x.BuyOrderId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BuyOrderId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BuyOrder)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BuyOrderId", wireType)
				}
				x.BuyOrderId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BuyOrderId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Buyer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Selection", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Selection == nil {
					x.Selection = &BuyOrder_Selection{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Selection); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BidPrice", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BidPrice == nil {
					x.BidPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BidPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisablePartialFill", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisablePartialFill = bool(v != 0)
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_BuyOrder_Selection               protoreflect.MessageDescriptor
	fd_BuyOrder_Selection_sell_order_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_BuyOrder_Selection = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("BuyOrder").Messages().ByName("Selection")
	fd_BuyOrder_Selection_sell_order_id = md_BuyOrder_Selection.Fields().ByName("sell_order_id")
}

var _ protoreflect.Message = (*fastReflection_BuyOrder_Selection)(nil)

type fastReflection_BuyOrder_Selection BuyOrder_Selection

func (x *BuyOrder_Selection) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BuyOrder_Selection)(x)
}

func (x *BuyOrder_Selection) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BuyOrder_Selection_messageType fastReflection_BuyOrder_Selection_messageType
var _ protoreflect.MessageType = fastReflection_BuyOrder_Selection_messageType{}

type fastReflection_BuyOrder_Selection_messageType struct{}

func (x fastReflection_BuyOrder_Selection_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BuyOrder_Selection)(nil)
}
func (x fastReflection_BuyOrder_Selection_messageType) New() protoreflect.Message {
	return new(fastReflection_BuyOrder_Selection)
}
func (x fastReflection_BuyOrder_Selection_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder_Selection
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BuyOrder_Selection) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder_Selection
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BuyOrder_Selection) Type() protoreflect.MessageType {
	return _fastReflection_BuyOrder_Selection_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BuyOrder_Selection) New() protoreflect.Message {
	return new(fastReflection_BuyOrder_Selection)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BuyOrder_Selection) Interface() protoreflect.ProtoMessage {
	return (*BuyOrder_Selection)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BuyOrder_Selection) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			v := o.SellOrderId
			value := protoreflect.ValueOfUint64(v)
			if !f(fd_BuyOrder_Selection_sell_order_id, value) {
				return
			}
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BuyOrder_Selection) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.Selection.sell_order_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*BuyOrder_Selection_SellOrderId); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder_Selection) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.Selection.sell_order_id":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BuyOrder_Selection) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.Selection.sell_order_id":
		if x.Sum == nil {
			return protoreflect.ValueOfUint64(uint64(0))
		} else if v, ok := x.Sum.(*BuyOrder_Selection_SellOrderId); ok {
			return protoreflect.ValueOfUint64(v.SellOrderId)
		} else {
			return protoreflect.ValueOfUint64(uint64(0))
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder.Selection does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder_Selection) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.Selection.sell_order_id":
		cv := value.Uint()
		x.Sum = &BuyOrder_Selection_SellOrderId{SellOrderId: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder_Selection) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.Selection.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.v1alpha2.BuyOrder.Selection is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BuyOrder_Selection) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.Selection.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BuyOrder_Selection) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.v1alpha2.BuyOrder.Selection.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			return x.Descriptor().Fields().ByName("sell_order_id")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.BuyOrder.Selection", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BuyOrder_Selection) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder_Selection) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BuyOrder_Selection) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BuyOrder_Selection) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BuyOrder_Selection)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		switch x := x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			if x == nil {
				break
			}
			n += 1 + runtime.Sov(uint64(x.SellOrderId))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BuyOrder_Selection)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		switch x := x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SellOrderId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BuyOrder_Selection)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder_Selection: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder_Selection: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SellOrderId", wireType)
				}
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Sum = &BuyOrder_Selection_SellOrderId{v}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_AskDenom               protoreflect.MessageDescriptor
	fd_AskDenom_denom         protoreflect.FieldDescriptor
	fd_AskDenom_display_denom protoreflect.FieldDescriptor
	fd_AskDenom_exponent      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_AskDenom = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("AskDenom")
	fd_AskDenom_denom = md_AskDenom.Fields().ByName("denom")
	fd_AskDenom_display_denom = md_AskDenom.Fields().ByName("display_denom")
	fd_AskDenom_exponent = md_AskDenom.Fields().ByName("exponent")
}

var _ protoreflect.Message = (*fastReflection_AskDenom)(nil)

type fastReflection_AskDenom AskDenom

func (x *AskDenom) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AskDenom)(x)
}

func (x *AskDenom) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AskDenom_messageType fastReflection_AskDenom_messageType
var _ protoreflect.MessageType = fastReflection_AskDenom_messageType{}

type fastReflection_AskDenom_messageType struct{}

func (x fastReflection_AskDenom_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AskDenom)(nil)
}
func (x fastReflection_AskDenom_messageType) New() protoreflect.Message {
	return new(fastReflection_AskDenom)
}
func (x fastReflection_AskDenom_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AskDenom
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AskDenom) Descriptor() protoreflect.MessageDescriptor {
	return md_AskDenom
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AskDenom) Type() protoreflect.MessageType {
	return _fastReflection_AskDenom_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AskDenom) New() protoreflect.Message {
	return new(fastReflection_AskDenom)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AskDenom) Interface() protoreflect.ProtoMessage {
	return (*AskDenom)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AskDenom) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_AskDenom_denom, value) {
			return
		}
	}
	if x.DisplayDenom != "" {
		value := protoreflect.ValueOfString(x.DisplayDenom)
		if !f(fd_AskDenom_display_denom, value) {
			return
		}
	}
	if x.Exponent != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Exponent)
		if !f(fd_AskDenom_exponent, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_AskDenom) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.AskDenom.denom":
		return x.Denom != ""
	case "regen.ecocredit.v1alpha2.AskDenom.display_denom":
		return x.DisplayDenom != ""
	case "regen.ecocredit.v1alpha2.AskDenom.exponent":
		return x.Exponent != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.AskDenom does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AskDenom) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.AskDenom.denom":
		x.Denom = ""
	case "regen.ecocredit.v1alpha2.AskDenom.display_denom":
		x.DisplayDenom = ""
	case "regen.ecocredit.v1alpha2.AskDenom.exponent":
		x.Exponent = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.AskDenom does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AskDenom) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.AskDenom.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.AskDenom.display_denom":
		value := x.DisplayDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.AskDenom.exponent":
		value := x.Exponent
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.AskDenom does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AskDenom) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.AskDenom.denom":
		x.Denom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.AskDenom.display_denom":
		x.DisplayDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.AskDenom.exponent":
		x.Exponent = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.AskDenom does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AskDenom) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.AskDenom.denom":
		panic(fmt.Errorf("field denom of message regen.ecocredit.v1alpha2.AskDenom is not mutable"))
	case "regen.ecocredit.v1alpha2.AskDenom.display_denom":
		panic(fmt.Errorf("field display_denom of message regen.ecocredit.v1alpha2.AskDenom is not mutable"))
	case "regen.ecocredit.v1alpha2.AskDenom.exponent":
		panic(fmt.Errorf("field exponent of message regen.ecocredit.v1alpha2.AskDenom is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.AskDenom does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AskDenom) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.AskDenom.denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.AskDenom.display_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.AskDenom.exponent":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.AskDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.AskDenom does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AskDenom) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.AskDenom", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AskDenom) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AskDenom) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_AskDenom) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AskDenom) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AskDenom)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DisplayDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Exponent != 0 {
			n += 1 + runtime.Sov(uint64(x.Exponent))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*AskDenom)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Exponent != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exponent))
			i--
			dAtA[i] = 0x18
		}
		if len(x.DisplayDenom) > 0 {
			i -= len(x.DisplayDenom)
			copy(dAtA[i:], x.DisplayDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DisplayDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*AskDenom)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AskDenom: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AskDenom: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisplayDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DisplayDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
				}
				x.Exponent = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Exponent |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_BasketCredit                 protoreflect.MessageDescriptor
	fd_BasketCredit_batch_denom     protoreflect.FieldDescriptor
	fd_BasketCredit_tradable_amount protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_BasketCredit = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("BasketCredit")
	fd_BasketCredit_batch_denom = md_BasketCredit.Fields().ByName("batch_denom")
	fd_BasketCredit_tradable_amount = md_BasketCredit.Fields().ByName("tradable_amount")
}

var _ protoreflect.Message = (*fastReflection_BasketCredit)(nil)

type fastReflection_BasketCredit BasketCredit

func (x *BasketCredit) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BasketCredit)(x)
}

func (x *BasketCredit) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BasketCredit_messageType fastReflection_BasketCredit_messageType
var _ protoreflect.MessageType = fastReflection_BasketCredit_messageType{}

type fastReflection_BasketCredit_messageType struct{}

func (x fastReflection_BasketCredit_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BasketCredit)(nil)
}
func (x fastReflection_BasketCredit_messageType) New() protoreflect.Message {
	return new(fastReflection_BasketCredit)
}
func (x fastReflection_BasketCredit_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BasketCredit
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BasketCredit) Descriptor() protoreflect.MessageDescriptor {
	return md_BasketCredit
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BasketCredit) Type() protoreflect.MessageType {
	return _fastReflection_BasketCredit_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BasketCredit) New() protoreflect.Message {
	return new(fastReflection_BasketCredit)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BasketCredit) Interface() protoreflect.ProtoMessage {
	return (*BasketCredit)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BasketCredit) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_BasketCredit_batch_denom, value) {
			return
		}
	}
	if x.TradableAmount != "" {
		value := protoreflect.ValueOfString(x.TradableAmount)
		if !f(fd_BasketCredit_tradable_amount, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BasketCredit) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCredit.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha2.BasketCredit.tradable_amount":
		return x.TradableAmount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCredit"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCredit does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketCredit) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCredit.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha2.BasketCredit.tradable_amount":
		x.TradableAmount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCredit"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCredit does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BasketCredit) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCredit.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha2.BasketCredit.tradable_amount":
		value := x.TradableAmount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCredit"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCredit does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketCredit) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCredit.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha2.BasketCredit.tradable_amount":
		x.TradableAmount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCredit"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCredit does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketCredit) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCredit.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.BasketCredit is not mutable"))
	case "regen.ecocredit.v1alpha2.BasketCredit.tradable_amount":
		panic(fmt.Errorf("field tradable_amount of message regen.ecocredit.v1alpha2.BasketCredit is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCredit"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCredit does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BasketCredit) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCredit.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.BasketCredit.tradable_amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCredit"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCredit does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BasketCredit) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.BasketCredit", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BasketCredit) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketCredit) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BasketCredit) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BasketCredit) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BasketCredit)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TradableAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BasketCredit)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.TradableAmount) > 0 {
			i -= len(x.TradableAmount)
			copy(dAtA[i:], x.TradableAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TradableAmount)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BasketCredit)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BasketCredit: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BasketCredit: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TradableAmount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TradableAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_BasketCriteria            protoreflect.MessageDescriptor
	fd_BasketCriteria_filter     protoreflect.FieldDescriptor
	fd_BasketCriteria_multiplier protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_BasketCriteria = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("BasketCriteria")
	fd_BasketCriteria_filter = md_BasketCriteria.Fields().ByName("filter")
	fd_BasketCriteria_multiplier = md_BasketCriteria.Fields().ByName("multiplier")
}

var _ protoreflect.Message = (*fastReflection_BasketCriteria)(nil)

type fastReflection_BasketCriteria BasketCriteria

func (x *BasketCriteria) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BasketCriteria)(x)
}

func (x *BasketCriteria) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BasketCriteria_messageType fastReflection_BasketCriteria_messageType
var _ protoreflect.MessageType = fastReflection_BasketCriteria_messageType{}

type fastReflection_BasketCriteria_messageType struct{}

func (x fastReflection_BasketCriteria_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BasketCriteria)(nil)
}
func (x fastReflection_BasketCriteria_messageType) New() protoreflect.Message {
	return new(fastReflection_BasketCriteria)
}
func (x fastReflection_BasketCriteria_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BasketCriteria
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BasketCriteria) Descriptor() protoreflect.MessageDescriptor {
	return md_BasketCriteria
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BasketCriteria) Type() protoreflect.MessageType {
	return _fastReflection_BasketCriteria_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BasketCriteria) New() protoreflect.Message {
	return new(fastReflection_BasketCriteria)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BasketCriteria) Interface() protoreflect.ProtoMessage {
	return (*BasketCriteria)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BasketCriteria) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Filter != nil {
		value := protoreflect.ValueOfMessage(x.Filter.ProtoReflect())
		if !f(fd_BasketCriteria_filter, value) {
			return
		}
	}
	if x.Multiplier != "" {
		value := protoreflect.ValueOfString(x.Multiplier)
		if !f(fd_BasketCriteria_multiplier, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BasketCriteria) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCriteria.filter":
		return x.Filter != nil
	case "regen.ecocredit.v1alpha2.BasketCriteria.multiplier":
		return x.Multiplier != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCriteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCriteria does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketCriteria) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCriteria.filter":
		x.Filter = nil
	case "regen.ecocredit.v1alpha2.BasketCriteria.multiplier":
		x.Multiplier = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCriteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCriteria does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BasketCriteria) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCriteria.filter":
		value := x.Filter
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BasketCriteria.multiplier":
		value := x.Multiplier
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCriteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCriteria does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketCriteria) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCriteria.filter":
		x.Filter = value.Message().Interface().(*Filter)
	case "regen.ecocredit.v1alpha2.BasketCriteria.multiplier":
		x.Multiplier = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCriteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCriteria does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketCriteria) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCriteria.filter":
		if x.Filter == nil {
			x.Filter = new(Filter)
		}
		return protoreflect.ValueOfMessage(x.Filter.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BasketCriteria.multiplier":
		panic(fmt.Errorf("field multiplier of message regen.ecocredit.v1alpha2.BasketCriteria is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCriteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCriteria does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BasketCriteria) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.BasketCriteria.filter":
		m := new(Filter)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.BasketCriteria.multiplier":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.BasketCriteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.BasketCriteria does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BasketCriteria) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.BasketCriteria", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BasketCriteria) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketCriteria) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BasketCriteria) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BasketCriteria) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BasketCriteria)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Filter != nil {
			l = options.Size(x.Filter)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Multiplier)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BasketCriteria)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Multiplier) > 0 {
			i -= len(x.Multiplier)
			copy(dAtA[i:], x.Multiplier)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Multiplier)))
			i--
			dAtA[i] = 0x12
		}
		if x.Filter != nil {
			encoded, err := options.Marshal(x.Filter)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BasketCriteria)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BasketCriteria: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BasketCriteria: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Filter == nil {
					x.Filter = &Filter{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Filter); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Multiplier = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Filter                  protoreflect.MessageDescriptor
	fd_Filter_and              protoreflect.FieldDescriptor
	fd_Filter_or               protoreflect.FieldDescriptor
	fd_Filter_credit_type_name protoreflect.FieldDescriptor
	fd_Filter_class_id         protoreflect.FieldDescriptor
	fd_Filter_project_id       protoreflect.FieldDescriptor
	fd_Filter_batch_denom      protoreflect.FieldDescriptor
	fd_Filter_class_admin      protoreflect.FieldDescriptor
	fd_Filter_issuer           protoreflect.FieldDescriptor
	fd_Filter_owner            protoreflect.FieldDescriptor
	fd_Filter_project_location protoreflect.FieldDescriptor
	fd_Filter_date_range       protoreflect.FieldDescriptor
	fd_Filter_tag              protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_Filter = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("Filter")
	fd_Filter_and = md_Filter.Fields().ByName("and")
	fd_Filter_or = md_Filter.Fields().ByName("or")
	fd_Filter_credit_type_name = md_Filter.Fields().ByName("credit_type_name")
	fd_Filter_class_id = md_Filter.Fields().ByName("class_id")
	fd_Filter_project_id = md_Filter.Fields().ByName("project_id")
	fd_Filter_batch_denom = md_Filter.Fields().ByName("batch_denom")
	fd_Filter_class_admin = md_Filter.Fields().ByName("class_admin")
	fd_Filter_issuer = md_Filter.Fields().ByName("issuer")
	fd_Filter_owner = md_Filter.Fields().ByName("owner")
	fd_Filter_project_location = md_Filter.Fields().ByName("project_location")
	fd_Filter_date_range = md_Filter.Fields().ByName("date_range")
	fd_Filter_tag = md_Filter.Fields().ByName("tag")
}

var _ protoreflect.Message = (*fastReflection_Filter)(nil)

type fastReflection_Filter Filter

func (x *Filter) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Filter)(x)
}

func (x *Filter) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Filter_messageType fastReflection_Filter_messageType
var _ protoreflect.MessageType = fastReflection_Filter_messageType{}

type fastReflection_Filter_messageType struct{}

func (x fastReflection_Filter_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Filter)(nil)
}
func (x fastReflection_Filter_messageType) New() protoreflect.Message {
	return new(fastReflection_Filter)
}
func (x fastReflection_Filter_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Filter) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Filter) Type() protoreflect.MessageType {
	return _fastReflection_Filter_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Filter) New() protoreflect.Message {
	return new(fastReflection_Filter)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Filter) Interface() protoreflect.ProtoMessage {
	return (*Filter)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Filter) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *Filter_And_:
			v := o.And
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Filter_and, value) {
				return
			}
		case *Filter_Or_:
			v := o.Or
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Filter_or, value) {
				return
			}
		case *Filter_CreditTypeName:
			v := o.CreditTypeName
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_credit_type_name, value) {
				return
			}
		case *Filter_ClassId:
			v := o.ClassId
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_class_id, value) {
				return
			}
		case *Filter_ProjectId:
			v := o.ProjectId
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_project_id, value) {
				return
			}
		case *Filter_BatchDenom:
			v := o.BatchDenom
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_batch_denom, value) {
				return
			}
		case *Filter_ClassAdmin:
			v := o.ClassAdmin
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_class_admin, value) {
				return
			}
		case *Filter_Issuer:
			v := o.Issuer
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_issuer, value) {
				return
			}
		case *Filter_Owner:
			v := o.Owner
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_owner, value) {
				return
			}
		case *Filter_ProjectLocation:
			v := o.ProjectLocation
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_project_location, value) {
				return
			}
		case *Filter_DateRange_:
			v := o.DateRange
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Filter_date_range, value) {
				return
			}
		case *Filter_Tag:
			v := o.Tag
			value := protoreflect.ValueOfString(v)
			if !f(fd_Filter_tag, value) {
				return
			}
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Filter) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.and":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_And_); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.or":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_Or_); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.credit_type_name":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_CreditTypeName); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.class_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_ClassId); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.project_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_ProjectId); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.batch_denom":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_BatchDenom); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.class_admin":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_ClassAdmin); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.issuer":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_Issuer); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.owner":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_Owner); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.project_location":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_ProjectLocation); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.date_range":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_DateRange_); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.v1alpha2.Filter.tag":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Filter_Tag); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.and":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.or":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.credit_type_name":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.class_id":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.project_id":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.batch_denom":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.class_admin":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.issuer":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.owner":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.project_location":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.date_range":
		x.Sum = nil
	case "regen.ecocredit.v1alpha2.Filter.tag":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Filter) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.and":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*Filter_And)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*Filter_And_); ok {
			return protoreflect.ValueOfMessage(v.And.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Filter_And)(nil).ProtoReflect())
		}
	case "regen.ecocredit.v1alpha2.Filter.or":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*Filter_Or)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*Filter_Or_); ok {
			return protoreflect.ValueOfMessage(v.Or.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Filter_Or)(nil).ProtoReflect())
		}
	case "regen.ecocredit.v1alpha2.Filter.credit_type_name":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_CreditTypeName); ok {
			return protoreflect.ValueOfString(v.CreditTypeName)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.v1alpha2.Filter.class_id":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_ClassId); ok {
			return protoreflect.ValueOfString(v.ClassId)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.v1alpha2.Filter.project_id":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_ProjectId); ok {
			return protoreflect.ValueOfString(v.ProjectId)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.v1alpha2.Filter.batch_denom":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_BatchDenom); ok {
			return protoreflect.ValueOfString(v.BatchDenom)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.v1alpha2.Filter.class_admin":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_ClassAdmin); ok {
			return protoreflect.ValueOfString(v.ClassAdmin)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.v1alpha2.Filter.issuer":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_Issuer); ok {
			return protoreflect.ValueOfString(v.Issuer)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.v1alpha2.Filter.owner":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_Owner); ok {
			return protoreflect.ValueOfString(v.Owner)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.v1alpha2.Filter.project_location":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_ProjectLocation); ok {
			return protoreflect.ValueOfString(v.ProjectLocation)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.v1alpha2.Filter.date_range":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*Filter_DateRange)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*Filter_DateRange_); ok {
			return protoreflect.ValueOfMessage(v.DateRange.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Filter_DateRange)(nil).ProtoReflect())
		}
	case "regen.ecocredit.v1alpha2.Filter.tag":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*Filter_Tag); ok {
			return protoreflect.ValueOfString(v.Tag)
		} else {
			return protoreflect.ValueOfString("")
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.and":
		cv := value.Message().Interface().(*Filter_And)
		x.Sum = &Filter_And_{And: cv}
	case "regen.ecocredit.v1alpha2.Filter.or":
		cv := value.Message().Interface().(*Filter_Or)
		x.Sum = &Filter_Or_{Or: cv}
	case "regen.ecocredit.v1alpha2.Filter.credit_type_name":
		cv := value.Interface().(string)
		x.Sum = &Filter_CreditTypeName{CreditTypeName: cv}
	case "regen.ecocredit.v1alpha2.Filter.class_id":
		cv := value.Interface().(string)
		x.Sum = &Filter_ClassId{ClassId: cv}
	case "regen.ecocredit.v1alpha2.Filter.project_id":
		cv := value.Interface().(string)
		x.Sum = &Filter_ProjectId{ProjectId: cv}
	case "regen.ecocredit.v1alpha2.Filter.batch_denom":
		cv := value.Interface().(string)
		x.Sum = &Filter_BatchDenom{BatchDenom: cv}
	case "regen.ecocredit.v1alpha2.Filter.class_admin":
		cv := value.Interface().(string)
		x.Sum = &Filter_ClassAdmin{ClassAdmin: cv}
	case "regen.ecocredit.v1alpha2.Filter.issuer":
		cv := value.Interface().(string)
		x.Sum = &Filter_Issuer{Issuer: cv}
	case "regen.ecocredit.v1alpha2.Filter.owner":
		cv := value.Interface().(string)
		x.Sum = &Filter_Owner{Owner: cv}
	case "regen.ecocredit.v1alpha2.Filter.project_location":
		cv := value.Interface().(string)
		x.Sum = &Filter_ProjectLocation{ProjectLocation: cv}
	case "regen.ecocredit.v1alpha2.Filter.date_range":
		cv := value.Message().Interface().(*Filter_DateRange)
		x.Sum = &Filter_DateRange_{DateRange: cv}
	case "regen.ecocredit.v1alpha2.Filter.tag":
		cv := value.Interface().(string)
		x.Sum = &Filter_Tag{Tag: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.and":
		if x.Sum == nil {
			value := &Filter_And{}
			oneofValue := &Filter_And_{And: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *Filter_And_:
			return protoreflect.ValueOfMessage(m.And.ProtoReflect())
		default:
			value := &Filter_And{}
			oneofValue := &Filter_And_{And: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "regen.ecocredit.v1alpha2.Filter.or":
		if x.Sum == nil {
			value := &Filter_Or{}
			oneofValue := &Filter_Or_{Or: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *Filter_Or_:
			return protoreflect.ValueOfMessage(m.Or.ProtoReflect())
		default:
			value := &Filter_Or{}
			oneofValue := &Filter_Or_{Or: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "regen.ecocredit.v1alpha2.Filter.date_range":
		if x.Sum == nil {
			value := &Filter_DateRange{}
			oneofValue := &Filter_DateRange_{DateRange: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *Filter_DateRange_:
			return protoreflect.ValueOfMessage(m.DateRange.ProtoReflect())
		default:
			value := &Filter_DateRange{}
			oneofValue := &Filter_DateRange_{DateRange: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "regen.ecocredit.v1alpha2.Filter.credit_type_name":
		panic(fmt.Errorf("field credit_type_name of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	case "regen.ecocredit.v1alpha2.Filter.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	case "regen.ecocredit.v1alpha2.Filter.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	case "regen.ecocredit.v1alpha2.Filter.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	case "regen.ecocredit.v1alpha2.Filter.class_admin":
		panic(fmt.Errorf("field class_admin of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	case "regen.ecocredit.v1alpha2.Filter.issuer":
		panic(fmt.Errorf("field issuer of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	case "regen.ecocredit.v1alpha2.Filter.owner":
		panic(fmt.Errorf("field owner of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	case "regen.ecocredit.v1alpha2.Filter.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	case "regen.ecocredit.v1alpha2.Filter.tag":
		panic(fmt.Errorf("field tag of message regen.ecocredit.v1alpha2.Filter is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Filter) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.and":
		value := &Filter_And{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.Filter.or":
		value := &Filter_Or{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.Filter.credit_type_name":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.Filter.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.Filter.project_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.Filter.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.Filter.class_admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.Filter.issuer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.Filter.owner":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.Filter.project_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha2.Filter.date_range":
		value := &Filter_DateRange{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.Filter.tag":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Filter) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *Filter_And_:
			return x.Descriptor().Fields().ByName("and")
		case *Filter_Or_:
			return x.Descriptor().Fields().ByName("or")
		case *Filter_CreditTypeName:
			return x.Descriptor().Fields().ByName("credit_type_name")
		case *Filter_ClassId:
			return x.Descriptor().Fields().ByName("class_id")
		case *Filter_ProjectId:
			return x.Descriptor().Fields().ByName("project_id")
		case *Filter_BatchDenom:
			return x.Descriptor().Fields().ByName("batch_denom")
		case *Filter_ClassAdmin:
			return x.Descriptor().Fields().ByName("class_admin")
		case *Filter_Issuer:
			return x.Descriptor().Fields().ByName("issuer")
		case *Filter_Owner:
			return x.Descriptor().Fields().ByName("owner")
		case *Filter_ProjectLocation:
			return x.Descriptor().Fields().ByName("project_location")
		case *Filter_DateRange_:
			return x.Descriptor().Fields().ByName("date_range")
		case *Filter_Tag:
			return x.Descriptor().Fields().ByName("tag")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.Filter", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Filter) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Filter) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Filter) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Filter)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		switch x := x.Sum.(type) {
		case *Filter_And_:
			if x == nil {
				break
			}
			l = options.Size(x.And)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_Or_:
			if x == nil {
				break
			}
			l = options.Size(x.Or)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_CreditTypeName:
			if x == nil {
				break
			}
			l = len(x.CreditTypeName)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_ClassId:
			if x == nil {
				break
			}
			l = len(x.ClassId)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_ProjectId:
			if x == nil {
				break
			}
			l = len(x.ProjectId)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_BatchDenom:
			if x == nil {
				break
			}
			l = len(x.BatchDenom)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_ClassAdmin:
			if x == nil {
				break
			}
			l = len(x.ClassAdmin)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_Issuer:
			if x == nil {
				break
			}
			l = len(x.Issuer)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_Owner:
			if x == nil {
				break
			}
			l = len(x.Owner)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_ProjectLocation:
			if x == nil {
				break
			}
			l = len(x.ProjectLocation)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_DateRange_:
			if x == nil {
				break
			}
			l = options.Size(x.DateRange)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_Tag:
			if x == nil {
				break
			}
			l = len(x.Tag)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Filter)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		switch x := x.Sum.(type) {
		case *Filter_And_:
			encoded, err := options.Marshal(x.And)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		case *Filter_Or_:
			encoded, err := options.Marshal(x.Or)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		case *Filter_CreditTypeName:
			i -= len(x.CreditTypeName)
			copy(dAtA[i:], x.CreditTypeName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CreditTypeName)))
			i--
			dAtA[i] = 0x1a
		case *Filter_ClassId:
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x22
		case *Filter_ProjectId:
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
			i--
			dAtA[i] = 0x2a
		case *Filter_BatchDenom:
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x32
		case *Filter_ClassAdmin:
			i -= len(x.ClassAdmin)
			copy(dAtA[i:], x.ClassAdmin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassAdmin)))
			i--
			dAtA[i] = 0x3a
		case *Filter_Issuer:
			i -= len(x.Issuer)
			copy(dAtA[i:], x.Issuer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuer)))
			i--
			dAtA[i] = 0x42
		case *Filter_Owner:
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0x4a
		case *Filter_ProjectLocation:
			i -= len(x.ProjectLocation)
			copy(dAtA[i:], x.ProjectLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectLocation)))
			i--
			dAtA[i] = 0x52
		case *Filter_DateRange_:
			encoded, err := options.Marshal(x.DateRange)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x5a
		case *Filter_Tag:
			i -= len(x.Tag)
			copy(dAtA[i:], x.Tag)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Tag)))
			i--
			dAtA[i] = 0x62
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Filter)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field And", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &Filter_And{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &Filter_And_{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Or", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &Filter_Or{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &Filter_Or_{v}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreditTypeName", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_CreditTypeName{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_ClassId{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_ProjectId{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_BatchDenom{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassAdmin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_ClassAdmin{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_Issuer{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 9:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_Owner{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 10:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectLocation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_ProjectLocation{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 11:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DateRange", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &Filter_DateRange{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &Filter_DateRange_{v}
				iNdEx = postIndex
			case 12:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sum = &Filter_Tag{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_Filter_And_1_list)(nil)

type _Filter_And_1_list struct {
	list *[]*Filter
}

func (x *_Filter_And_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Filter_And_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Filter_And_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Filter)
	(*x.list)[i] = concreteValue
}

func (x *_Filter_And_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Filter)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Filter_And_1_list) AppendMutable() protoreflect.Value {
	v := new(Filter)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Filter_And_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Filter_And_1_list) NewElement() protoreflect.Value {
	v := new(Filter)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Filter_And_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Filter_And         protoreflect.MessageDescriptor
	fd_Filter_And_filters protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_Filter_And = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("Filter").Messages().ByName("And")
	fd_Filter_And_filters = md_Filter_And.Fields().ByName("filters")
}

var _ protoreflect.Message = (*fastReflection_Filter_And)(nil)

type fastReflection_Filter_And Filter_And

func (x *Filter_And) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Filter_And)(x)
}

func (x *Filter_And) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Filter_And_messageType fastReflection_Filter_And_messageType
var _ protoreflect.MessageType = fastReflection_Filter_And_messageType{}

type fastReflection_Filter_And_messageType struct{}

func (x fastReflection_Filter_And_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Filter_And)(nil)
}
func (x fastReflection_Filter_And_messageType) New() protoreflect.Message {
	return new(fastReflection_Filter_And)
}
func (x fastReflection_Filter_And_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter_And
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Filter_And) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter_And
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Filter_And) Type() protoreflect.MessageType {
	return _fastReflection_Filter_And_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Filter_And) New() protoreflect.Message {
	return new(fastReflection_Filter_And)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Filter_And) Interface() protoreflect.ProtoMessage {
	return (*Filter_And)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Filter_And) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Filters) != 0 {
		value := protoreflect.ValueOfList(&_Filter_And_1_list{list: &x.Filters})
		if !f(fd_Filter_And_filters, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Filter_And) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.And.filters":
		return len(x.Filters) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.And"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.And does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_And) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.And.filters":
		x.Filters = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.And"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.And does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Filter_And) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.And.filters":
		if len(x.Filters) == 0 {
			return protoreflect.ValueOfList(&_Filter_And_1_list{})
		}
		listValue := &_Filter_And_1_list{list: &x.Filters}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.And"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.And does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_And) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.And.filters":
		lv := value.List()
		clv := lv.(*_Filter_And_1_list)
		x.Filters = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.And"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.And does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_And) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.And.filters":
		if x.Filters == nil {
			x.Filters = []*Filter{}
		}
		value := &_Filter_And_1_list{list: &x.Filters}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.And"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.And does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Filter_And) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.And.filters":
		list := []*Filter{}
		return protoreflect.ValueOfList(&_Filter_And_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.And"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.And does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Filter_And) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.Filter.And", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Filter_And) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_And) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Filter_And) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Filter_And) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Filter_And)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Filters) > 0 {
			for _, e := range x.Filters {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Filter_And)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Filters) > 0 {
			for iNdEx := len(x.Filters) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Filters[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Filter_And)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter_And: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter_And: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Filters = append(x.Filters, &Filter{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Filters[len(x.Filters)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_Filter_Or_1_list)(nil)

type _Filter_Or_1_list struct {
	list *[]*Filter
}

func (x *_Filter_Or_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Filter_Or_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Filter_Or_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Filter)
	(*x.list)[i] = concreteValue
}

func (x *_Filter_Or_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Filter)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Filter_Or_1_list) AppendMutable() protoreflect.Value {
	v := new(Filter)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Filter_Or_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Filter_Or_1_list) NewElement() protoreflect.Value {
	v := new(Filter)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Filter_Or_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Filter_Or         protoreflect.MessageDescriptor
	fd_Filter_Or_filters protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_Filter_Or = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("Filter").Messages().ByName("Or")
	fd_Filter_Or_filters = md_Filter_Or.Fields().ByName("filters")
}

var _ protoreflect.Message = (*fastReflection_Filter_Or)(nil)

type fastReflection_Filter_Or Filter_Or

func (x *Filter_Or) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Filter_Or)(x)
}

func (x *Filter_Or) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Filter_Or_messageType fastReflection_Filter_Or_messageType
var _ protoreflect.MessageType = fastReflection_Filter_Or_messageType{}

type fastReflection_Filter_Or_messageType struct{}

func (x fastReflection_Filter_Or_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Filter_Or)(nil)
}
func (x fastReflection_Filter_Or_messageType) New() protoreflect.Message {
	return new(fastReflection_Filter_Or)
}
func (x fastReflection_Filter_Or_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter_Or
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Filter_Or) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter_Or
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Filter_Or) Type() protoreflect.MessageType {
	return _fastReflection_Filter_Or_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Filter_Or) New() protoreflect.Message {
	return new(fastReflection_Filter_Or)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Filter_Or) Interface() protoreflect.ProtoMessage {
	return (*Filter_Or)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Filter_Or) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Filters) != 0 {
		value := protoreflect.ValueOfList(&_Filter_Or_1_list{list: &x.Filters})
		if !f(fd_Filter_Or_filters, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Filter_Or) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.Or.filters":
		return len(x.Filters) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.Or"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.Or does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_Or) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.Or.filters":
		x.Filters = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.Or"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.Or does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Filter_Or) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.Or.filters":
		if len(x.Filters) == 0 {
			return protoreflect.ValueOfList(&_Filter_Or_1_list{})
		}
		listValue := &_Filter_Or_1_list{list: &x.Filters}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.Or"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.Or does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_Or) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.Or.filters":
		lv := value.List()
		clv := lv.(*_Filter_Or_1_list)
		x.Filters = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.Or"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.Or does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_Or) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.Or.filters":
		if x.Filters == nil {
			x.Filters = []*Filter{}
		}
		value := &_Filter_Or_1_list{list: &x.Filters}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.Or"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.Or does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Filter_Or) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.Or.filters":
		list := []*Filter{}
		return protoreflect.ValueOfList(&_Filter_Or_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.Or"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.Or does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Filter_Or) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.Filter.Or", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Filter_Or) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_Or) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Filter_Or) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Filter_Or) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Filter_Or)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Filters) > 0 {
			for _, e := range x.Filters {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Filter_Or)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Filters) > 0 {
			for iNdEx := len(x.Filters) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Filters[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Filter_Or)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter_Or: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter_Or: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Filters = append(x.Filters, &Filter{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Filters[len(x.Filters)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Filter_DateRange            protoreflect.MessageDescriptor
	fd_Filter_DateRange_start_date protoreflect.FieldDescriptor
	fd_Filter_DateRange_end_date   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha2_types_proto_init()
	md_Filter_DateRange = File_regen_ecocredit_v1alpha2_types_proto.Messages().ByName("Filter").Messages().ByName("DateRange")
	fd_Filter_DateRange_start_date = md_Filter_DateRange.Fields().ByName("start_date")
	fd_Filter_DateRange_end_date = md_Filter_DateRange.Fields().ByName("end_date")
}

var _ protoreflect.Message = (*fastReflection_Filter_DateRange)(nil)

type fastReflection_Filter_DateRange Filter_DateRange

func (x *Filter_DateRange) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Filter_DateRange)(x)
}

func (x *Filter_DateRange) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Filter_DateRange_messageType fastReflection_Filter_DateRange_messageType
var _ protoreflect.MessageType = fastReflection_Filter_DateRange_messageType{}

type fastReflection_Filter_DateRange_messageType struct{}

func (x fastReflection_Filter_DateRange_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Filter_DateRange)(nil)
}
func (x fastReflection_Filter_DateRange_messageType) New() protoreflect.Message {
	return new(fastReflection_Filter_DateRange)
}
func (x fastReflection_Filter_DateRange_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter_DateRange
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Filter_DateRange) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter_DateRange
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Filter_DateRange) Type() protoreflect.MessageType {
	return _fastReflection_Filter_DateRange_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Filter_DateRange) New() protoreflect.Message {
	return new(fastReflection_Filter_DateRange)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Filter_DateRange) Interface() protoreflect.ProtoMessage {
	return (*Filter_DateRange)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Filter_DateRange) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.StartDate != nil {
		value := protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
		if !f(fd_Filter_DateRange_start_date, value) {
			return
		}
	}
	if x.EndDate != nil {
		value := protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
		if !f(fd_Filter_DateRange_end_date, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Filter_DateRange) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.DateRange.start_date":
		return x.StartDate != nil
	case "regen.ecocredit.v1alpha2.Filter.DateRange.end_date":
		return x.EndDate != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.DateRange"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.DateRange does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_DateRange) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.DateRange.start_date":
		x.StartDate = nil
	case "regen.ecocredit.v1alpha2.Filter.DateRange.end_date":
		x.EndDate = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.DateRange"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.DateRange does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Filter_DateRange) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.DateRange.start_date":
		value := x.StartDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha2.Filter.DateRange.end_date":
		value := x.EndDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.DateRange"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.DateRange does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_DateRange) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.DateRange.start_date":
		x.StartDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.v1alpha2.Filter.DateRange.end_date":
		x.EndDate = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.DateRange"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.DateRange does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_DateRange) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.DateRange.start_date":
		if x.StartDate == nil {
			x.StartDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
	case "regen.ecocredit.v1alpha2.Filter.DateRange.end_date":
		if x.EndDate == nil {
			x.EndDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.DateRange"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.DateRange does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Filter_DateRange) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha2.Filter.DateRange.start_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha2.Filter.DateRange.end_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha2.Filter.DateRange"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha2.Filter.DateRange does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Filter_DateRange) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha2.Filter.DateRange", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Filter_DateRange) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_DateRange) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Filter_DateRange) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Filter_DateRange) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Filter_DateRange)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.StartDate != nil {
			l = options.Size(x.StartDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.EndDate != nil {
			l = options.Size(x.EndDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Filter_DateRange)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.EndDate != nil {
			encoded, err := options.Marshal(x.EndDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x42
		}
		if x.StartDate != nil {
			encoded, err := options.Marshal(x.StartDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x3a
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Filter_DateRange)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter_DateRange: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter_DateRange: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.StartDate == nil {
					x.StartDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.StartDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.EndDate == nil {
					x.EndDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.EndDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: regen/ecocredit/v1alpha2/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClassInfo represents the high-level on-chain information for a credit class.
type ClassInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// class_id is the unique ID of credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// admin is the admin of the credit class.
	Admin string `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
	// issuers are the approved issuers of the credit class.
	Issuers []string `protobuf:"bytes,3,rep,name=issuers,proto3" json:"issuers,omitempty"`
	// metadata is any arbitrary metadata to attached to the credit class.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// credit_type describes the type of credit (e.g. carbon, biodiversity), as well as unit and precision.
	CreditType *CreditType `protobuf:"bytes,5,opt,name=credit_type,json=creditType,proto3" json:"credit_type,omitempty"`
	// The number of batches issued in this credit class.
	NumBatches uint64 `protobuf:"varint,6,opt,name=num_batches,json=numBatches,proto3" json:"num_batches,omitempty"`
}

func (x *ClassInfo) Reset() {
	*x = ClassInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassInfo) ProtoMessage() {}

// Deprecated: Use ClassInfo.ProtoReflect.Descriptor instead.
func (*ClassInfo) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{0}
}

func (x *ClassInfo) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *ClassInfo) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *ClassInfo) GetIssuers() []string {
	if x != nil {
		return x.Issuers
	}
	return nil
}

func (x *ClassInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ClassInfo) GetCreditType() *CreditType {
	if x != nil {
		return x.CreditType
	}
	return nil
}

func (x *ClassInfo) GetNumBatches() uint64 {
	if x != nil {
		return x.NumBatches
	}
	return 0
}

// ProjectInfo represents the high-level on-chain information for a project.
type ProjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id is the unique ID of the project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// class_id is the unique ID of credit class for this project.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// issuer is the issuer of the credit batches for this project.
	Issuer string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// project_location is the location of the project.
	// Full documentation can be found in MsgCreateProject.project_location.
	ProjectLocation string `protobuf:"bytes,4,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
	// metadata is any arbitrary metadata attached to the project.
	Metadata []byte `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ProjectInfo) Reset() {
	*x = ProjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInfo) ProtoMessage() {}

// Deprecated: Use ProjectInfo.ProtoReflect.Descriptor instead.
func (*ProjectInfo) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectInfo) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *ProjectInfo) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *ProjectInfo) GetProjectLocation() string {
	if x != nil {
		return x.ProjectLocation
	}
	return ""
}

func (x *ProjectInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// BatchInfo represents the high-level on-chain information for a credit batch.
type BatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// project_id is the unique ID of the project this batch belongs to.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// total_amount is the total number of active credits in the credit batch.
	// Some of the issued credits may be cancelled and will be removed from
	// total_amount and tracked in amount_cancelled. total_amount and
	// amount_cancelled will always sum to the original amount of credits that
	// were issued.
	TotalAmount string `protobuf:"bytes,4,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	// metadata is any arbitrary metadata attached to the credit batch.
	Metadata []byte `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// amount_cancelled is the number of credits in the batch that have been
	// cancelled, effectively undoing there issuance. The sum of total_amount and
	// amount_cancelled will always sum to the original amount of credits that
	// were issued.
	AmountCancelled string `protobuf:"bytes,6,opt,name=amount_cancelled,json=amountCancelled,proto3" json:"amount_cancelled,omitempty"`
	// start_date is the beginning of the period during which this credit batch
	// was quantified and verified.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date is the end of the period during which this credit batch was
	// quantified and verified.
	EndDate *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *BatchInfo) Reset() {
	*x = BatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchInfo) ProtoMessage() {}

// Deprecated: Use BatchInfo.ProtoReflect.Descriptor instead.
func (*BatchInfo) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{2}
}

func (x *BatchInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BatchInfo) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *BatchInfo) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

func (x *BatchInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BatchInfo) GetAmountCancelled() string {
	if x != nil {
		return x.AmountCancelled
	}
	return ""
}

func (x *BatchInfo) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *BatchInfo) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

// Params defines the updatable global parameters of the ecocredit module for
// use with the x/params module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// credit_class_fee is the fixed fee charged on creation of a new credit class
	CreditClassFee []*v1beta1.Coin `protobuf:"bytes,1,rep,name=credit_class_fee,json=creditClassFee,proto3" json:"credit_class_fee,omitempty"`
	// allowed_class_creators is an allowlist defining the addresses with
	// the required permissions to create credit classes
	AllowedClassCreators []string `protobuf:"bytes,2,rep,name=allowed_class_creators,json=allowedClassCreators,proto3" json:"allowed_class_creators,omitempty"`
	// allowlist_enabled is a param that enables/disables the allowlist for credit
	// creation
	AllowlistEnabled bool `protobuf:"varint,3,opt,name=allowlist_enabled,json=allowlistEnabled,proto3" json:"allowlist_enabled,omitempty"`
	// credit_types is a list of definitions for credit types
	CreditTypes []*CreditType `protobuf:"bytes,4,rep,name=credit_types,json=creditTypes,proto3" json:"credit_types,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{3}
}

func (x *Params) GetCreditClassFee() []*v1beta1.Coin {
	if x != nil {
		return x.CreditClassFee
	}
	return nil
}

func (x *Params) GetAllowedClassCreators() []string {
	if x != nil {
		return x.AllowedClassCreators
	}
	return nil
}

func (x *Params) GetAllowlistEnabled() bool {
	if x != nil {
		return x.AllowlistEnabled
	}
	return false
}

func (x *Params) GetCreditTypes() []*CreditType {
	if x != nil {
		return x.CreditTypes
	}
	return nil
}

// CreditType defines the measurement unit/precision of a certain credit type
// (e.g. carbon, biodiversity...)
type CreditType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the type of credit (e.g. carbon, biodiversity, etc)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// abbreviation is a 1-3 character uppercase abbreviation of the CreditType
	// name, used in batch denominations within the CreditType. It must be unique.
	Abbreviation string `protobuf:"bytes,2,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	// the measurement unit (e.g. kg, ton, etc)
	Unit string `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	// the decimal precision
	Precision uint32 `protobuf:"varint,4,opt,name=precision,proto3" json:"precision,omitempty"`
}

func (x *CreditType) Reset() {
	*x = CreditType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditType) ProtoMessage() {}

// Deprecated: Use CreditType.ProtoReflect.Descriptor instead.
func (*CreditType) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{4}
}

func (x *CreditType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreditType) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *CreditType) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CreditType) GetPrecision() uint32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

// CreditTypeSeq associates a sequence number with a credit type abbreviation.
// This represents the number of credit classes created with that credit type.
type CreditTypeSeq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The credit type abbreviation
	Abbreviation string `protobuf:"bytes,1,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	// The sequence number of classes of the credit type
	SeqNumber uint64 `protobuf:"varint,2,opt,name=seq_number,json=seqNumber,proto3" json:"seq_number,omitempty"`
}

func (x *CreditTypeSeq) Reset() {
	*x = CreditTypeSeq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditTypeSeq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditTypeSeq) ProtoMessage() {}

// Deprecated: Use CreditTypeSeq.ProtoReflect.Descriptor instead.
func (*CreditTypeSeq) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{5}
}

func (x *CreditTypeSeq) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *CreditTypeSeq) GetSeqNumber() uint64 {
	if x != nil {
		return x.SeqNumber
	}
	return 0
}

// SellOrder represents the information for a sell order.
type SellOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// order_id is the unique ID of sell order.
	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// owner is the address of the owner of the credits being sold.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// batch_denom is the credit batch being sold.
	BatchDenom string `protobuf:"bytes,3,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// quantity is the quantity of credits being sold.
	Quantity string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// ask_price is the price the seller is asking for each unit of the
	// batch_denom. Each credit unit of the batch will be sold for at least the
	// ask_price or more.
	AskPrice *v1beta1.Coin `protobuf:"bytes,5,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	// disable_auto_retire disables auto-retirement of credits which allows a
	// buyer to disable auto-retirement in their buy order enabling them to
	// resell the credits to another buyer.
	DisableAutoRetire bool `protobuf:"varint,6,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
}

func (x *SellOrder) Reset() {
	*x = SellOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellOrder) ProtoMessage() {}

// Deprecated: Use SellOrder.ProtoReflect.Descriptor instead.
func (*SellOrder) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{6}
}

func (x *SellOrder) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SellOrder) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SellOrder) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *SellOrder) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *SellOrder) GetAskPrice() *v1beta1.Coin {
	if x != nil {
		return x.AskPrice
	}
	return nil
}

func (x *SellOrder) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

// BuyOrder represents the information for a buy order.
type BuyOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// buy_order_id is the unique ID of buy order.
	BuyOrderId uint64 `protobuf:"varint,1,opt,name=buy_order_id,json=buyOrderId,proto3" json:"buy_order_id,omitempty"`
	// buyer is the address that created the buy order
	Buyer string `protobuf:"bytes,2,opt,name=buyer,proto3" json:"buyer,omitempty"`
	// selection is the buy order selection.
	Selection *BuyOrder_Selection `protobuf:"bytes,3,opt,name=selection,proto3" json:"selection,omitempty"`
	// quantity is the quantity of credits to buy. If the quantity of credits
	// available is less than this amount the order will be partially filled
	// unless disable_partial_fill is true.
	Quantity string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// bid price is the bid price for this buy order. A credit unit will be
	// settled at a purchase price that is no more than the bid price. The
	// buy order will fail if the buyer does not have enough funds available
	// to complete the purchase.
	BidPrice *v1beta1.Coin `protobuf:"bytes,5,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	// disable_auto_retire allows auto-retirement to be disabled. If it is set to true
	// the credits will not auto-retire and can be resold assuming that the
	// corresponding sell order has auto-retirement disabled. If the sell order
	// hasn't disabled auto-retirement and the buy order tries to disable it,
	// that buy order will fail.
	DisableAutoRetire bool `protobuf:"varint,6,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// disable_partial_fill disables the default behavior of partially filling
	// buy orders if the requested quantity is not available.
	DisablePartialFill bool `protobuf:"varint,7,opt,name=disable_partial_fill,json=disablePartialFill,proto3" json:"disable_partial_fill,omitempty"`
}

func (x *BuyOrder) Reset() {
	*x = BuyOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrder) ProtoMessage() {}

// Deprecated: Use BuyOrder.ProtoReflect.Descriptor instead.
func (*BuyOrder) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{7}
}

func (x *BuyOrder) GetBuyOrderId() uint64 {
	if x != nil {
		return x.BuyOrderId
	}
	return 0
}

func (x *BuyOrder) GetBuyer() string {
	if x != nil {
		return x.Buyer
	}
	return ""
}

func (x *BuyOrder) GetSelection() *BuyOrder_Selection {
	if x != nil {
		return x.Selection
	}
	return nil
}

func (x *BuyOrder) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *BuyOrder) GetBidPrice() *v1beta1.Coin {
	if x != nil {
		return x.BidPrice
	}
	return nil
}

func (x *BuyOrder) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *BuyOrder) GetDisablePartialFill() bool {
	if x != nil {
		return x.DisablePartialFill
	}
	return false
}

// AskDenom represents the information for an ask denom.
type AskDenom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// denom is the denom to allow (ex. ibc/GLKHDSG423SGS)
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// display_denom is the denom to display to the user and is informational
	DisplayDenom string `protobuf:"bytes,2,opt,name=display_denom,json=displayDenom,proto3" json:"display_denom,omitempty"`
	// exponent is the exponent that relates the denom to the display_denom and is
	// informational
	Exponent uint32 `protobuf:"varint,3,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (x *AskDenom) Reset() {
	*x = AskDenom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskDenom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskDenom) ProtoMessage() {}

// Deprecated: Use AskDenom.ProtoReflect.Descriptor instead.
func (*AskDenom) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{8}
}

func (x *AskDenom) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *AskDenom) GetDisplayDenom() string {
	if x != nil {
		return x.DisplayDenom
	}
	return ""
}

func (x *AskDenom) GetExponent() uint32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

// BasketCredit represents the information for a credit batch inside a basket.
type BasketCredit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// tradable_amount is the number of credits in this transfer that can be
	// traded by the recipient. Decimal values are acceptable within the
	// precision returned by Query/Precision.
	TradableAmount string `protobuf:"bytes,2,opt,name=tradable_amount,json=tradableAmount,proto3" json:"tradable_amount,omitempty"`
}

func (x *BasketCredit) Reset() {
	*x = BasketCredit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketCredit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketCredit) ProtoMessage() {}

// Deprecated: Use BasketCredit.ProtoReflect.Descriptor instead.
func (*BasketCredit) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{9}
}

func (x *BasketCredit) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *BasketCredit) GetTradableAmount() string {
	if x != nil {
		return x.TradableAmount
	}
	return ""
}

// BasketCriteria defines a criteria by which credits can be added to a basket.
type BasketCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// filter defines condition(s) that credits should satisfy in order to be
	// added to the basket.
	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// multiplier is an integer number which is applied to credit units when
	// converting to basket units. For example if the multiplier is 2000, then
	// 1.1 credits will result in 2200 basket tokens. If there are any fractional
	// amounts left over in this calculation when adding credits to a basket,
	// those fractional amounts will not get added to the basket.
	Multiplier string `protobuf:"bytes,2,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
}

func (x *BasketCriteria) Reset() {
	*x = BasketCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketCriteria) ProtoMessage() {}

// Deprecated: Use BasketCriteria.ProtoReflect.Descriptor instead.
func (*BasketCriteria) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{10}
}

func (x *BasketCriteria) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *BasketCriteria) GetMultiplier() string {
	if x != nil {
		return x.Multiplier
	}
	return ""
}

// Filter defines condition(s) that credits should satisfy in order to be added
// to the basket. It can handled nested conditions linked with and/or operators.
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum is the oneof that specifies the type of filter.
	//
	// Types that are assignable to Sum:
	//	*Filter_And_
	//	*Filter_Or_
	//	*Filter_CreditTypeName
	//	*Filter_ClassId
	//	*Filter_ProjectId
	//	*Filter_BatchDenom
	//	*Filter_ClassAdmin
	//	*Filter_Issuer
	//	*Filter_Owner
	//	*Filter_ProjectLocation
	//	*Filter_DateRange_
	//	*Filter_Tag
	Sum isFilter_Sum `protobuf_oneof:"sum"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{11}
}

func (x *Filter) GetSum() isFilter_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *Filter) GetAnd() *Filter_And {
	if x, ok := x.GetSum().(*Filter_And_); ok {
		return x.And
	}
	return nil
}

func (x *Filter) GetOr() *Filter_Or {
	if x, ok := x.GetSum().(*Filter_Or_); ok {
		return x.Or
	}
	return nil
}

func (x *Filter) GetCreditTypeName() string {
	if x, ok := x.GetSum().(*Filter_CreditTypeName); ok {
		return x.CreditTypeName
	}
	return ""
}

func (x *Filter) GetClassId() string {
	if x, ok := x.GetSum().(*Filter_ClassId); ok {
		return x.ClassId
	}
	return ""
}

func (x *Filter) GetProjectId() string {
	if x, ok := x.GetSum().(*Filter_ProjectId); ok {
		return x.ProjectId
	}
	return ""
}

func (x *Filter) GetBatchDenom() string {
	if x, ok := x.GetSum().(*Filter_BatchDenom); ok {
		return x.BatchDenom
	}
	return ""
}

func (x *Filter) GetClassAdmin() string {
	if x, ok := x.GetSum().(*Filter_ClassAdmin); ok {
		return x.ClassAdmin
	}
	return ""
}

func (x *Filter) GetIssuer() string {
	if x, ok := x.GetSum().(*Filter_Issuer); ok {
		return x.Issuer
	}
	return ""
}

func (x *Filter) GetOwner() string {
	if x, ok := x.GetSum().(*Filter_Owner); ok {
		return x.Owner
	}
	return ""
}

func (x *Filter) GetProjectLocation() string {
	if x, ok := x.GetSum().(*Filter_ProjectLocation); ok {
		return x.ProjectLocation
	}
	return ""
}

func (x *Filter) GetDateRange() *Filter_DateRange {
	if x, ok := x.GetSum().(*Filter_DateRange_); ok {
		return x.DateRange
	}
	return nil
}

func (x *Filter) GetTag() string {
	if x, ok := x.GetSum().(*Filter_Tag); ok {
		return x.Tag
	}
	return ""
}

type isFilter_Sum interface {
	isFilter_Sum()
}

type Filter_And_ struct {
	// and specifies a list of filters where all conditions should be satisfied.
	And *Filter_And `protobuf:"bytes,1,opt,name=and,proto3,oneof"`
}

type Filter_Or_ struct {
	// or specifies a list of filters where at least one of the conditions should be satisfied.
	Or *Filter_Or `protobuf:"bytes,2,opt,name=or,proto3,oneof"`
}

type Filter_CreditTypeName struct {
	// credit_type_name filters against credits from this credit type name.
	CreditTypeName string `protobuf:"bytes,3,opt,name=credit_type_name,json=creditTypeName,proto3,oneof"`
}

type Filter_ClassId struct {
	// class_id filters against credits from this credit class id.
	ClassId string `protobuf:"bytes,4,opt,name=class_id,json=classId,proto3,oneof"`
}

type Filter_ProjectId struct {
	// project_id filters against credits from this project.
	ProjectId string `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3,oneof"`
}

type Filter_BatchDenom struct {
	// batch_denom filters against credits from this batch.
	BatchDenom string `protobuf:"bytes,6,opt,name=batch_denom,json=batchDenom,proto3,oneof"`
}

type Filter_ClassAdmin struct {
	// class_admin filters against credits issued by this class admin.
	ClassAdmin string `protobuf:"bytes,7,opt,name=class_admin,json=classAdmin,proto3,oneof"`
}

type Filter_Issuer struct {
	// issuer filters against credits issued by this issuer address.
	Issuer string `protobuf:"bytes,8,opt,name=issuer,proto3,oneof"`
}

type Filter_Owner struct {
	// owner filters against credits currently owned by this address.
	Owner string `protobuf:"bytes,9,opt,name=owner,proto3,oneof"`
}

type Filter_ProjectLocation struct {
	// project_location can be specified in three levels of granularity:
	// country, sub-national-code, or postal code. If just country is given,
	// for instance "US" then any credits in the "US" will be matched even
	// their project location is more specific, ex. "US-NY 12345". If
	// a country, sub-national-code and postal code are all provided then
	// only projects in that postal code will match.
	ProjectLocation string `protobuf:"bytes,10,opt,name=project_location,json=projectLocation,proto3,oneof"`
}

type Filter_DateRange_ struct {
	// date_range filters against credit batch start_date and/or end_date.
	DateRange *Filter_DateRange `protobuf:"bytes,11,opt,name=date_range,json=dateRange,proto3,oneof"`
}

type Filter_Tag struct {
	// tag specifies a curation tag to match against.
	Tag string `protobuf:"bytes,12,opt,name=tag,proto3,oneof"`
}

func (*Filter_And_) isFilter_Sum() {}

func (*Filter_Or_) isFilter_Sum() {}

func (*Filter_CreditTypeName) isFilter_Sum() {}

func (*Filter_ClassId) isFilter_Sum() {}

func (*Filter_ProjectId) isFilter_Sum() {}

func (*Filter_BatchDenom) isFilter_Sum() {}

func (*Filter_ClassAdmin) isFilter_Sum() {}

func (*Filter_Issuer) isFilter_Sum() {}

func (*Filter_Owner) isFilter_Sum() {}

func (*Filter_ProjectLocation) isFilter_Sum() {}

func (*Filter_DateRange_) isFilter_Sum() {}

func (*Filter_Tag) isFilter_Sum() {}

// Selection defines a buy order selection.
type BuyOrder_Selection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum defines the type of selection.
	//
	// Types that are assignable to Sum:
	//	*BuyOrder_Selection_SellOrderId
	Sum isBuyOrder_Selection_Sum `protobuf_oneof:"sum"`
}

func (x *BuyOrder_Selection) Reset() {
	*x = BuyOrder_Selection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrder_Selection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrder_Selection) ProtoMessage() {}

// Deprecated: Use BuyOrder_Selection.ProtoReflect.Descriptor instead.
func (*BuyOrder_Selection) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{7, 0}
}

func (x *BuyOrder_Selection) GetSum() isBuyOrder_Selection_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *BuyOrder_Selection) GetSellOrderId() uint64 {
	if x, ok := x.GetSum().(*BuyOrder_Selection_SellOrderId); ok {
		return x.SellOrderId
	}
	return 0
}

type isBuyOrder_Selection_Sum interface {
	isBuyOrder_Selection_Sum()
}

type BuyOrder_Selection_SellOrderId struct {
	// sell_order_id is the sell order ID against which the buyer is trying to buy.
	// When sell_order_id is set, this is known as a direct buy order because it
	// is placed directly against a specific sell order.
	SellOrderId uint64 `protobuf:"varint,1,opt,name=sell_order_id,json=sellOrderId,proto3,oneof"`
}

func (*BuyOrder_Selection_SellOrderId) isBuyOrder_Selection_Sum() {}

// And specifies an "and" condition between the list of filters.
type Filter_And struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// filters is a list of filters where all conditions should be satisfied.
	Filters []*Filter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *Filter_And) Reset() {
	*x = Filter_And{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter_And) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter_And) ProtoMessage() {}

// Deprecated: Use Filter_And.ProtoReflect.Descriptor instead.
func (*Filter_And) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{11, 0}
}

func (x *Filter_And) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

// And specifies an "or" condition between the list of filters.
type Filter_Or struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// filters is a list of filters where at least one of the conditions should be satisfied.
	Filters []*Filter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *Filter_Or) Reset() {
	*x = Filter_Or{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter_Or) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter_Or) ProtoMessage() {}

// Deprecated: Use Filter_Or.ProtoReflect.Descriptor instead.
func (*Filter_Or) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{11, 1}
}

func (x *Filter_Or) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

// DateRange defines a period for credit batches in a basket.
type Filter_DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start_date is the beginning of the period during which this credit batch
	// was quantified and verified. If it is empty then there is no start date
	// limit.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date is the end of the period during which this credit batch was
	// quantified and verified. If it is empty then there is no end date
	// limit.
	EndDate *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *Filter_DateRange) Reset() {
	*x = Filter_DateRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha2_types_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter_DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter_DateRange) ProtoMessage() {}

// Deprecated: Use Filter_DateRange.ProtoReflect.Descriptor instead.
func (*Filter_DateRange) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP(), []int{11, 2}
}

func (x *Filter_DateRange) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Filter_DateRange) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

var File_regen_ecocredit_v1alpha2_types_proto protoreflect.FileDescriptor

var file_regen_ecocredit_v1alpha2_types_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb3, 0x02, 0x0a,
	0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x22, 0xab, 0x02, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x75, 0x0a,
	0x10, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x66, 0x65,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x73, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x46, 0x65, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x76, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70,
	0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x62, 0x62,
	0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xe1, 0x01, 0x0a,
	0x09, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65,
	0x22, 0xfe, 0x02, 0x0a, 0x08, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0c, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a,
	0x09, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x62, 0x69, 0x64,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x6c, 0x1a, 0x38, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0b, 0x73,
	0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75,
	0x6d, 0x22, 0x61, 0x0a, 0x08, 0x41, 0x73, 0x6b, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x22, 0x58, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x74, 0x72, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6a,
	0x0a, 0x0e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x12, 0x38, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x22, 0x81, 0x06, 0x0a, 0x06, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x03, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x03, 0x61, 0x6e, 0x64, 0x12,
	0x35, 0x0a, 0x02, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x72,
	0x48, 0x00, 0x52, 0x02, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65,
	0x6e, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x03, 0x74, 0x61, 0x67, 0x1a, 0x41, 0x0a, 0x03, 0x41, 0x6e, 0x64, 0x12, 0x3a, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x40, 0x0a, 0x02, 0x4f, 0x72, 0x12,
	0x3a, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x89, 0x01, 0x0a, 0x09,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x42, 0x82,
	0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x3b, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x58, 0xaa, 0x02, 0x18, 0x52, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x18, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xe2,
	0x02, 0x24, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a,
	0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_v1alpha2_types_proto_rawDescOnce sync.Once
	file_regen_ecocredit_v1alpha2_types_proto_rawDescData = file_regen_ecocredit_v1alpha2_types_proto_rawDesc
)

func file_regen_ecocredit_v1alpha2_types_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_v1alpha2_types_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_v1alpha2_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_v1alpha2_types_proto_rawDescData)
	})
	return file_regen_ecocredit_v1alpha2_types_proto_rawDescData
}

var file_regen_ecocredit_v1alpha2_types_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_regen_ecocredit_v1alpha2_types_proto_goTypes = []interface{}{
	(*ClassInfo)(nil),             // 0: regen.ecocredit.v1alpha2.ClassInfo
	(*ProjectInfo)(nil),           // 1: regen.ecocredit.v1alpha2.ProjectInfo
	(*BatchInfo)(nil),             // 2: regen.ecocredit.v1alpha2.BatchInfo
	(*Params)(nil),                // 3: regen.ecocredit.v1alpha2.Params
	(*CreditType)(nil),            // 4: regen.ecocredit.v1alpha2.CreditType
	(*CreditTypeSeq)(nil),         // 5: regen.ecocredit.v1alpha2.CreditTypeSeq
	(*SellOrder)(nil),             // 6: regen.ecocredit.v1alpha2.SellOrder
	(*BuyOrder)(nil),              // 7: regen.ecocredit.v1alpha2.BuyOrder
	(*AskDenom)(nil),              // 8: regen.ecocredit.v1alpha2.AskDenom
	(*BasketCredit)(nil),          // 9: regen.ecocredit.v1alpha2.BasketCredit
	(*BasketCriteria)(nil),        // 10: regen.ecocredit.v1alpha2.BasketCriteria
	(*Filter)(nil),                // 11: regen.ecocredit.v1alpha2.Filter
	(*BuyOrder_Selection)(nil),    // 12: regen.ecocredit.v1alpha2.BuyOrder.Selection
	(*Filter_And)(nil),            // 13: regen.ecocredit.v1alpha2.Filter.And
	(*Filter_Or)(nil),             // 14: regen.ecocredit.v1alpha2.Filter.Or
	(*Filter_DateRange)(nil),      // 15: regen.ecocredit.v1alpha2.Filter.DateRange
	(*timestamppb.Timestamp)(nil), // 16: google.protobuf.Timestamp
	(*v1beta1.Coin)(nil),          // 17: cosmos.base.v1beta1.Coin
}
var file_regen_ecocredit_v1alpha2_types_proto_depIdxs = []int32{
	4,  // 0: regen.ecocredit.v1alpha2.ClassInfo.credit_type:type_name -> regen.ecocredit.v1alpha2.CreditType
	16, // 1: regen.ecocredit.v1alpha2.BatchInfo.start_date:type_name -> google.protobuf.Timestamp
	16, // 2: regen.ecocredit.v1alpha2.BatchInfo.end_date:type_name -> google.protobuf.Timestamp
	17, // 3: regen.ecocredit.v1alpha2.Params.credit_class_fee:type_name -> cosmos.base.v1beta1.Coin
	4,  // 4: regen.ecocredit.v1alpha2.Params.credit_types:type_name -> regen.ecocredit.v1alpha2.CreditType
	17, // 5: regen.ecocredit.v1alpha2.SellOrder.ask_price:type_name -> cosmos.base.v1beta1.Coin
	12, // 6: regen.ecocredit.v1alpha2.BuyOrder.selection:type_name -> regen.ecocredit.v1alpha2.BuyOrder.Selection
	17, // 7: regen.ecocredit.v1alpha2.BuyOrder.bid_price:type_name -> cosmos.base.v1beta1.Coin
	11, // 8: regen.ecocredit.v1alpha2.BasketCriteria.filter:type_name -> regen.ecocredit.v1alpha2.Filter
	13, // 9: regen.ecocredit.v1alpha2.Filter.and:type_name -> regen.ecocredit.v1alpha2.Filter.And
	14, // 10: regen.ecocredit.v1alpha2.Filter.or:type_name -> regen.ecocredit.v1alpha2.Filter.Or
	15, // 11: regen.ecocredit.v1alpha2.Filter.date_range:type_name -> regen.ecocredit.v1alpha2.Filter.DateRange
	11, // 12: regen.ecocredit.v1alpha2.Filter.And.filters:type_name -> regen.ecocredit.v1alpha2.Filter
	11, // 13: regen.ecocredit.v1alpha2.Filter.Or.filters:type_name -> regen.ecocredit.v1alpha2.Filter
	16, // 14: regen.ecocredit.v1alpha2.Filter.DateRange.start_date:type_name -> google.protobuf.Timestamp
	16, // 15: regen.ecocredit.v1alpha2.Filter.DateRange.end_date:type_name -> google.protobuf.Timestamp
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_v1alpha2_types_proto_init() }
func file_regen_ecocredit_v1alpha2_types_proto_init() {
	if File_regen_ecocredit_v1alpha2_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditTypeSeq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskDenom); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasketCredit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasketCriteria); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrder_Selection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter_And); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter_Or); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_v1alpha2_types_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter_DateRange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_regen_ecocredit_v1alpha2_types_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*Filter_And_)(nil),
		(*Filter_Or_)(nil),
		(*Filter_CreditTypeName)(nil),
		(*Filter_ClassId)(nil),
		(*Filter_ProjectId)(nil),
		(*Filter_BatchDenom)(nil),
		(*Filter_ClassAdmin)(nil),
		(*Filter_Issuer)(nil),
		(*Filter_Owner)(nil),
		(*Filter_ProjectLocation)(nil),
		(*Filter_DateRange_)(nil),
		(*Filter_Tag)(nil),
	}
	file_regen_ecocredit_v1alpha2_types_proto_msgTypes[12].OneofWrappers = []interface{}{
		(*BuyOrder_Selection_SellOrderId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_v1alpha2_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_v1alpha2_types_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_v1alpha2_types_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_v1alpha2_types_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_v1alpha2_types_proto = out.File
	file_regen_ecocredit_v1alpha2_types_proto_rawDesc = nil
	file_regen_ecocredit_v1alpha2_types_proto_goTypes = nil
	file_regen_ecocredit_v1alpha2_types_proto_depIdxs = nil
}
