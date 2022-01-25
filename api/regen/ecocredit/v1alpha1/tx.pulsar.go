package ecocreditv1alpha1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_MsgCreateClass_2_list)(nil)

type _MsgCreateClass_2_list struct {
	list *[]string
}

func (x *_MsgCreateClass_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateClass_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_MsgCreateClass_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateClass_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateClass_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgCreateClass at list field Issuers as it is not of Message kind"))
}

func (x *_MsgCreateClass_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateClass_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_MsgCreateClass_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateClass                  protoreflect.MessageDescriptor
	fd_MsgCreateClass_admin            protoreflect.FieldDescriptor
	fd_MsgCreateClass_issuers          protoreflect.FieldDescriptor
	fd_MsgCreateClass_metadata         protoreflect.FieldDescriptor
	fd_MsgCreateClass_credit_type_name protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgCreateClass = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgCreateClass")
	fd_MsgCreateClass_admin = md_MsgCreateClass.Fields().ByName("admin")
	fd_MsgCreateClass_issuers = md_MsgCreateClass.Fields().ByName("issuers")
	fd_MsgCreateClass_metadata = md_MsgCreateClass.Fields().ByName("metadata")
	fd_MsgCreateClass_credit_type_name = md_MsgCreateClass.Fields().ByName("credit_type_name")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateClass)(nil)

type fastReflection_MsgCreateClass MsgCreateClass

func (x *MsgCreateClass) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateClass)(x)
}

func (x *MsgCreateClass) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateClass_messageType fastReflection_MsgCreateClass_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateClass_messageType{}

type fastReflection_MsgCreateClass_messageType struct{}

func (x fastReflection_MsgCreateClass_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateClass)(nil)
}
func (x fastReflection_MsgCreateClass_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateClass)
}
func (x fastReflection_MsgCreateClass_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateClass
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateClass) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateClass
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateClass) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateClass_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateClass) New() protoreflect.Message {
	return new(fastReflection_MsgCreateClass)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateClass) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateClass)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateClass) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgCreateClass_admin, value) {
			return
		}
	}
	if len(x.Issuers) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateClass_2_list{list: &x.Issuers})
		if !f(fd_MsgCreateClass_issuers, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateClass_metadata, value) {
			return
		}
	}
	if x.CreditTypeName != "" {
		value := protoreflect.ValueOfString(x.CreditTypeName)
		if !f(fd_MsgCreateClass_credit_type_name, value) {
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
func (x *fastReflection_MsgCreateClass) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClass.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha1.MsgCreateClass.issuers":
		return len(x.Issuers) != 0
	case "regen.ecocredit.v1alpha1.MsgCreateClass.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1alpha1.MsgCreateClass.credit_type_name":
		return x.CreditTypeName != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClass does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateClass) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClass.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha1.MsgCreateClass.issuers":
		x.Issuers = nil
	case "regen.ecocredit.v1alpha1.MsgCreateClass.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1alpha1.MsgCreateClass.credit_type_name":
		x.CreditTypeName = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClass does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateClass) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClass.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgCreateClass.issuers":
		if len(x.Issuers) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateClass_2_list{})
		}
		listValue := &_MsgCreateClass_2_list{list: &x.Issuers}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha1.MsgCreateClass.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1alpha1.MsgCreateClass.credit_type_name":
		value := x.CreditTypeName
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClass does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateClass) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClass.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgCreateClass.issuers":
		lv := value.List()
		clv := lv.(*_MsgCreateClass_2_list)
		x.Issuers = *clv.list
	case "regen.ecocredit.v1alpha1.MsgCreateClass.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1alpha1.MsgCreateClass.credit_type_name":
		x.CreditTypeName = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClass does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateClass) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClass.issuers":
		if x.Issuers == nil {
			x.Issuers = []string{}
		}
		value := &_MsgCreateClass_2_list{list: &x.Issuers}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha1.MsgCreateClass.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha1.MsgCreateClass is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCreateClass.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha1.MsgCreateClass is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCreateClass.credit_type_name":
		panic(fmt.Errorf("field credit_type_name of message regen.ecocredit.v1alpha1.MsgCreateClass is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClass does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateClass) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClass.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgCreateClass.issuers":
		list := []string{}
		return protoreflect.ValueOfList(&_MsgCreateClass_2_list{list: &list})
	case "regen.ecocredit.v1alpha1.MsgCreateClass.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1alpha1.MsgCreateClass.credit_type_name":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClass"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClass does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateClass) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgCreateClass", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateClass) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateClass) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateClass) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateClass) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateClass)
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
		l = len(x.CreditTypeName)
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
		x := input.Message.Interface().(*MsgCreateClass)
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
		if len(x.CreditTypeName) > 0 {
			i -= len(x.CreditTypeName)
			copy(dAtA[i:], x.CreditTypeName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CreditTypeName)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Issuers) > 0 {
			for iNdEx := len(x.Issuers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Issuers[iNdEx])
				copy(dAtA[i:], x.Issuers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuers[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
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
		x := input.Message.Interface().(*MsgCreateClass)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateClass: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateClass: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
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
			case 3:
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
			case 4:
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
				x.CreditTypeName = string(dAtA[iNdEx:postIndex])
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
	md_MsgCreateClassResponse          protoreflect.MessageDescriptor
	fd_MsgCreateClassResponse_class_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgCreateClassResponse = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgCreateClassResponse")
	fd_MsgCreateClassResponse_class_id = md_MsgCreateClassResponse.Fields().ByName("class_id")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateClassResponse)(nil)

type fastReflection_MsgCreateClassResponse MsgCreateClassResponse

func (x *MsgCreateClassResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateClassResponse)(x)
}

func (x *MsgCreateClassResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateClassResponse_messageType fastReflection_MsgCreateClassResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateClassResponse_messageType{}

type fastReflection_MsgCreateClassResponse_messageType struct{}

func (x fastReflection_MsgCreateClassResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateClassResponse)(nil)
}
func (x fastReflection_MsgCreateClassResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateClassResponse)
}
func (x fastReflection_MsgCreateClassResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateClassResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateClassResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateClassResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateClassResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateClassResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateClassResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateClassResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateClassResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateClassResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateClassResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgCreateClassResponse_class_id, value) {
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
func (x *fastReflection_MsgCreateClassResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClassResponse.class_id":
		return x.ClassId != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClassResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateClassResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClassResponse.class_id":
		x.ClassId = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClassResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateClassResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClassResponse.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClassResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateClassResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClassResponse.class_id":
		x.ClassId = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClassResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateClassResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClassResponse.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha1.MsgCreateClassResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClassResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateClassResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateClassResponse.class_id":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateClassResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateClassResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateClassResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgCreateClassResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateClassResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateClassResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateClassResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateClassResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateClassResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateClassResponse)
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
		x := input.Message.Interface().(*MsgCreateClassResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateClassResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateClassResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_MsgCreateBatch_3_list)(nil)

type _MsgCreateBatch_3_list struct {
	list *[]*MsgCreateBatch_BatchIssuance
}

func (x *_MsgCreateBatch_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateBatch_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreateBatch_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgCreateBatch_BatchIssuance)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateBatch_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgCreateBatch_BatchIssuance)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateBatch_3_list) AppendMutable() protoreflect.Value {
	v := new(MsgCreateBatch_BatchIssuance)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateBatch_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateBatch_3_list) NewElement() protoreflect.Value {
	v := new(MsgCreateBatch_BatchIssuance)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateBatch_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateBatch                  protoreflect.MessageDescriptor
	fd_MsgCreateBatch_issuer           protoreflect.FieldDescriptor
	fd_MsgCreateBatch_class_id         protoreflect.FieldDescriptor
	fd_MsgCreateBatch_issuance         protoreflect.FieldDescriptor
	fd_MsgCreateBatch_metadata         protoreflect.FieldDescriptor
	fd_MsgCreateBatch_start_date       protoreflect.FieldDescriptor
	fd_MsgCreateBatch_end_date         protoreflect.FieldDescriptor
	fd_MsgCreateBatch_project_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgCreateBatch = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgCreateBatch")
	fd_MsgCreateBatch_issuer = md_MsgCreateBatch.Fields().ByName("issuer")
	fd_MsgCreateBatch_class_id = md_MsgCreateBatch.Fields().ByName("class_id")
	fd_MsgCreateBatch_issuance = md_MsgCreateBatch.Fields().ByName("issuance")
	fd_MsgCreateBatch_metadata = md_MsgCreateBatch.Fields().ByName("metadata")
	fd_MsgCreateBatch_start_date = md_MsgCreateBatch.Fields().ByName("start_date")
	fd_MsgCreateBatch_end_date = md_MsgCreateBatch.Fields().ByName("end_date")
	fd_MsgCreateBatch_project_location = md_MsgCreateBatch.Fields().ByName("project_location")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateBatch)(nil)

type fastReflection_MsgCreateBatch MsgCreateBatch

func (x *MsgCreateBatch) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateBatch)(x)
}

func (x *MsgCreateBatch) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateBatch_messageType fastReflection_MsgCreateBatch_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateBatch_messageType{}

type fastReflection_MsgCreateBatch_messageType struct{}

func (x fastReflection_MsgCreateBatch_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateBatch)(nil)
}
func (x fastReflection_MsgCreateBatch_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatch)
}
func (x fastReflection_MsgCreateBatch_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatch
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateBatch) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatch
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateBatch) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateBatch_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateBatch) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatch)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateBatch) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateBatch)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateBatch) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Issuer != "" {
		value := protoreflect.ValueOfString(x.Issuer)
		if !f(fd_MsgCreateBatch_issuer, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgCreateBatch_class_id, value) {
			return
		}
	}
	if len(x.Issuance) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateBatch_3_list{list: &x.Issuance})
		if !f(fd_MsgCreateBatch_issuance, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateBatch_metadata, value) {
			return
		}
	}
	if x.StartDate != nil {
		value := protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
		if !f(fd_MsgCreateBatch_start_date, value) {
			return
		}
	}
	if x.EndDate != nil {
		value := protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
		if !f(fd_MsgCreateBatch_end_date, value) {
			return
		}
	}
	if x.ProjectLocation != "" {
		value := protoreflect.ValueOfString(x.ProjectLocation)
		if !f(fd_MsgCreateBatch_project_location, value) {
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
func (x *fastReflection_MsgCreateBatch) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuer":
		return x.Issuer != ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuance":
		return len(x.Issuance) != 0
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.metadata":
		return len(x.Metadata) != 0
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.start_date":
		return x.StartDate != nil
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.end_date":
		return x.EndDate != nil
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.project_location":
		return x.ProjectLocation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatch) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuer":
		x.Issuer = ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuance":
		x.Issuance = nil
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.metadata":
		x.Metadata = nil
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.start_date":
		x.StartDate = nil
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.end_date":
		x.EndDate = nil
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.project_location":
		x.ProjectLocation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateBatch) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuer":
		value := x.Issuer
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuance":
		if len(x.Issuance) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateBatch_3_list{})
		}
		listValue := &_MsgCreateBatch_3_list{list: &x.Issuance}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.start_date":
		value := x.StartDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.end_date":
		value := x.EndDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.project_location":
		value := x.ProjectLocation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateBatch) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuer":
		x.Issuer = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuance":
		lv := value.List()
		clv := lv.(*_MsgCreateBatch_3_list)
		x.Issuance = *clv.list
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.metadata":
		x.Metadata = value.Bytes()
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.start_date":
		x.StartDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.end_date":
		x.EndDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.project_location":
		x.ProjectLocation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateBatch) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuance":
		if x.Issuance == nil {
			x.Issuance = []*MsgCreateBatch_BatchIssuance{}
		}
		value := &_MsgCreateBatch_3_list{list: &x.Issuance}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.start_date":
		if x.StartDate == nil {
			x.StartDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.StartDate.ProtoReflect())
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.end_date":
		if x.EndDate == nil {
			x.EndDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.EndDate.ProtoReflect())
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuer":
		panic(fmt.Errorf("field issuer of message regen.ecocredit.v1alpha1.MsgCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha1.MsgCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha1.MsgCreateBatch is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.v1alpha1.MsgCreateBatch is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateBatch) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuer":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.issuance":
		list := []*MsgCreateBatch_BatchIssuance{}
		return protoreflect.ValueOfList(&_MsgCreateBatch_3_list{list: &list})
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.start_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.end_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.project_location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateBatch) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgCreateBatch", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateBatch) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatch) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateBatch) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateBatch) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateBatch)
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
		l = len(x.Issuer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Issuance) > 0 {
			for _, e := range x.Issuance {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Metadata)
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
		l = len(x.ProjectLocation)
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
		x := input.Message.Interface().(*MsgCreateBatch)
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
		if len(x.ProjectLocation) > 0 {
			i -= len(x.ProjectLocation)
			copy(dAtA[i:], x.ProjectLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectLocation)))
			i--
			dAtA[i] = 0x3a
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
			dAtA[i] = 0x32
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
			dAtA[i] = 0x2a
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Issuance) > 0 {
			for iNdEx := len(x.Issuance) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Issuance[iNdEx])
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
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Issuer) > 0 {
			i -= len(x.Issuer)
			copy(dAtA[i:], x.Issuer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuer)))
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
		x := input.Message.Interface().(*MsgCreateBatch)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatch: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatch: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Issuance", wireType)
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
				x.Issuance = append(x.Issuance, &MsgCreateBatch_BatchIssuance{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Issuance[len(x.Issuance)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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
			case 6:
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
			case 7:
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
	md_MsgCreateBatch_BatchIssuance                     protoreflect.MessageDescriptor
	fd_MsgCreateBatch_BatchIssuance_recipient           protoreflect.FieldDescriptor
	fd_MsgCreateBatch_BatchIssuance_tradable_amount     protoreflect.FieldDescriptor
	fd_MsgCreateBatch_BatchIssuance_retired_amount      protoreflect.FieldDescriptor
	fd_MsgCreateBatch_BatchIssuance_retirement_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgCreateBatch_BatchIssuance = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgCreateBatch").Messages().ByName("BatchIssuance")
	fd_MsgCreateBatch_BatchIssuance_recipient = md_MsgCreateBatch_BatchIssuance.Fields().ByName("recipient")
	fd_MsgCreateBatch_BatchIssuance_tradable_amount = md_MsgCreateBatch_BatchIssuance.Fields().ByName("tradable_amount")
	fd_MsgCreateBatch_BatchIssuance_retired_amount = md_MsgCreateBatch_BatchIssuance.Fields().ByName("retired_amount")
	fd_MsgCreateBatch_BatchIssuance_retirement_location = md_MsgCreateBatch_BatchIssuance.Fields().ByName("retirement_location")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateBatch_BatchIssuance)(nil)

type fastReflection_MsgCreateBatch_BatchIssuance MsgCreateBatch_BatchIssuance

func (x *MsgCreateBatch_BatchIssuance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateBatch_BatchIssuance)(x)
}

func (x *MsgCreateBatch_BatchIssuance) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateBatch_BatchIssuance_messageType fastReflection_MsgCreateBatch_BatchIssuance_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateBatch_BatchIssuance_messageType{}

type fastReflection_MsgCreateBatch_BatchIssuance_messageType struct{}

func (x fastReflection_MsgCreateBatch_BatchIssuance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateBatch_BatchIssuance)(nil)
}
func (x fastReflection_MsgCreateBatch_BatchIssuance_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatch_BatchIssuance)
}
func (x fastReflection_MsgCreateBatch_BatchIssuance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatch_BatchIssuance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatch_BatchIssuance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateBatch_BatchIssuance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatch_BatchIssuance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateBatch_BatchIssuance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Recipient != "" {
		value := protoreflect.ValueOfString(x.Recipient)
		if !f(fd_MsgCreateBatch_BatchIssuance_recipient, value) {
			return
		}
	}
	if x.TradableAmount != "" {
		value := protoreflect.ValueOfString(x.TradableAmount)
		if !f(fd_MsgCreateBatch_BatchIssuance_tradable_amount, value) {
			return
		}
	}
	if x.RetiredAmount != "" {
		value := protoreflect.ValueOfString(x.RetiredAmount)
		if !f(fd_MsgCreateBatch_BatchIssuance_retired_amount, value) {
			return
		}
	}
	if x.RetirementLocation != "" {
		value := protoreflect.ValueOfString(x.RetirementLocation)
		if !f(fd_MsgCreateBatch_BatchIssuance_retirement_location, value) {
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
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.recipient":
		return x.Recipient != ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.tradable_amount":
		return x.TradableAmount != ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retired_amount":
		return x.RetiredAmount != ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retirement_location":
		return x.RetirementLocation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.recipient":
		x.Recipient = ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.tradable_amount":
		x.TradableAmount = ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retired_amount":
		x.RetiredAmount = ""
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retirement_location":
		x.RetirementLocation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.recipient":
		value := x.Recipient
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.tradable_amount":
		value := x.TradableAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retired_amount":
		value := x.RetiredAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retirement_location":
		value := x.RetirementLocation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.recipient":
		x.Recipient = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.tradable_amount":
		x.TradableAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retired_amount":
		x.RetiredAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retirement_location":
		x.RetirementLocation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateBatch_BatchIssuance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.recipient":
		panic(fmt.Errorf("field recipient of message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.tradable_amount":
		panic(fmt.Errorf("field tradable_amount of message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retired_amount":
		panic(fmt.Errorf("field retired_amount of message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retirement_location":
		panic(fmt.Errorf("field retirement_location of message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.recipient":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.tradable_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retired_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance.retirement_location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateBatch_BatchIssuance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateBatch_BatchIssuance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateBatch_BatchIssuance)
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
		l = len(x.Recipient)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TradableAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetiredAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetirementLocation)
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
		x := input.Message.Interface().(*MsgCreateBatch_BatchIssuance)
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
		if len(x.RetirementLocation) > 0 {
			i -= len(x.RetirementLocation)
			copy(dAtA[i:], x.RetirementLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetirementLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.RetiredAmount) > 0 {
			i -= len(x.RetiredAmount)
			copy(dAtA[i:], x.RetiredAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetiredAmount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TradableAmount) > 0 {
			i -= len(x.TradableAmount)
			copy(dAtA[i:], x.TradableAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TradableAmount)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Recipient) > 0 {
			i -= len(x.Recipient)
			copy(dAtA[i:], x.Recipient)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Recipient)))
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
		x := input.Message.Interface().(*MsgCreateBatch_BatchIssuance)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatch_BatchIssuance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatch_BatchIssuance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
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
				x.Recipient = string(dAtA[iNdEx:postIndex])
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
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetiredAmount", wireType)
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
				x.RetiredAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetirementLocation", wireType)
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
				x.RetirementLocation = string(dAtA[iNdEx:postIndex])
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
	md_MsgCreateBatchResponse             protoreflect.MessageDescriptor
	fd_MsgCreateBatchResponse_batch_denom protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgCreateBatchResponse = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgCreateBatchResponse")
	fd_MsgCreateBatchResponse_batch_denom = md_MsgCreateBatchResponse.Fields().ByName("batch_denom")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateBatchResponse)(nil)

type fastReflection_MsgCreateBatchResponse MsgCreateBatchResponse

func (x *MsgCreateBatchResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateBatchResponse)(x)
}

func (x *MsgCreateBatchResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateBatchResponse_messageType fastReflection_MsgCreateBatchResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateBatchResponse_messageType{}

type fastReflection_MsgCreateBatchResponse_messageType struct{}

func (x fastReflection_MsgCreateBatchResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateBatchResponse)(nil)
}
func (x fastReflection_MsgCreateBatchResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatchResponse)
}
func (x fastReflection_MsgCreateBatchResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatchResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateBatchResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateBatchResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateBatchResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateBatchResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateBatchResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateBatchResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateBatchResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateBatchResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateBatchResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgCreateBatchResponse_batch_denom, value) {
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
func (x *fastReflection_MsgCreateBatchResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatchResponse.batch_denom":
		return x.BatchDenom != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatchResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatchResponse.batch_denom":
		x.BatchDenom = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateBatchResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatchResponse.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatchResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCreateBatchResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatchResponse.batch_denom":
		x.BatchDenom = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCreateBatchResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatchResponse.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha1.MsgCreateBatchResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateBatchResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCreateBatchResponse.batch_denom":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCreateBatchResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCreateBatchResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateBatchResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgCreateBatchResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateBatchResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateBatchResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCreateBatchResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateBatchResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateBatchResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateBatchResponse)
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
		x := input.Message.Interface().(*MsgCreateBatchResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatchResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateBatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_MsgSend_3_list)(nil)

type _MsgSend_3_list struct {
	list *[]*MsgSend_SendCredits
}

func (x *_MsgSend_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgSend_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgSend_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgSend_SendCredits)
	(*x.list)[i] = concreteValue
}

func (x *_MsgSend_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgSend_SendCredits)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgSend_3_list) AppendMutable() protoreflect.Value {
	v := new(MsgSend_SendCredits)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSend_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgSend_3_list) NewElement() protoreflect.Value {
	v := new(MsgSend_SendCredits)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgSend_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgSend           protoreflect.MessageDescriptor
	fd_MsgSend_sender    protoreflect.FieldDescriptor
	fd_MsgSend_recipient protoreflect.FieldDescriptor
	fd_MsgSend_credits   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgSend = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgSend")
	fd_MsgSend_sender = md_MsgSend.Fields().ByName("sender")
	fd_MsgSend_recipient = md_MsgSend.Fields().ByName("recipient")
	fd_MsgSend_credits = md_MsgSend.Fields().ByName("credits")
}

var _ protoreflect.Message = (*fastReflection_MsgSend)(nil)

type fastReflection_MsgSend MsgSend

func (x *MsgSend) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSend)(x)
}

func (x *MsgSend) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSend_messageType fastReflection_MsgSend_messageType
var _ protoreflect.MessageType = fastReflection_MsgSend_messageType{}

type fastReflection_MsgSend_messageType struct{}

func (x fastReflection_MsgSend_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSend)(nil)
}
func (x fastReflection_MsgSend_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSend)
}
func (x fastReflection_MsgSend_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSend) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSend) Type() protoreflect.MessageType {
	return _fastReflection_MsgSend_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSend) New() protoreflect.Message {
	return new(fastReflection_MsgSend)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSend) Interface() protoreflect.ProtoMessage {
	return (*MsgSend)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSend) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sender != "" {
		value := protoreflect.ValueOfString(x.Sender)
		if !f(fd_MsgSend_sender, value) {
			return
		}
	}
	if x.Recipient != "" {
		value := protoreflect.ValueOfString(x.Recipient)
		if !f(fd_MsgSend_recipient, value) {
			return
		}
	}
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgSend_3_list{list: &x.Credits})
		if !f(fd_MsgSend_credits, value) {
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
func (x *fastReflection_MsgSend) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.sender":
		return x.Sender != ""
	case "regen.ecocredit.v1alpha1.MsgSend.recipient":
		return x.Recipient != ""
	case "regen.ecocredit.v1alpha1.MsgSend.credits":
		return len(x.Credits) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.sender":
		x.Sender = ""
	case "regen.ecocredit.v1alpha1.MsgSend.recipient":
		x.Recipient = ""
	case "regen.ecocredit.v1alpha1.MsgSend.credits":
		x.Credits = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSend) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.sender":
		value := x.Sender
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgSend.recipient":
		value := x.Recipient
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgSend.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgSend_3_list{})
		}
		listValue := &_MsgSend_3_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSend) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.sender":
		x.Sender = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgSend.recipient":
		x.Recipient = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgSend.credits":
		lv := value.List()
		clv := lv.(*_MsgSend_3_list)
		x.Credits = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSend) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.credits":
		if x.Credits == nil {
			x.Credits = []*MsgSend_SendCredits{}
		}
		value := &_MsgSend_3_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha1.MsgSend.sender":
		panic(fmt.Errorf("field sender of message regen.ecocredit.v1alpha1.MsgSend is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgSend.recipient":
		panic(fmt.Errorf("field recipient of message regen.ecocredit.v1alpha1.MsgSend is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSend) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.sender":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgSend.recipient":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgSend.credits":
		list := []*MsgSend_SendCredits{}
		return protoreflect.ValueOfList(&_MsgSend_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSend) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgSend", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSend) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSend) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSend) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSend)
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
		l = len(x.Sender)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Recipient)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
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
		x := input.Message.Interface().(*MsgSend)
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
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
		}
		if len(x.Recipient) > 0 {
			i -= len(x.Recipient)
			copy(dAtA[i:], x.Recipient)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Recipient)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Sender) > 0 {
			i -= len(x.Sender)
			copy(dAtA[i:], x.Sender)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Sender)))
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
		x := input.Message.Interface().(*MsgSend)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
				x.Sender = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
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
				x.Recipient = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &MsgSend_SendCredits{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
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
	md_MsgSend_SendCredits                     protoreflect.MessageDescriptor
	fd_MsgSend_SendCredits_batch_denom         protoreflect.FieldDescriptor
	fd_MsgSend_SendCredits_tradable_amount     protoreflect.FieldDescriptor
	fd_MsgSend_SendCredits_retired_amount      protoreflect.FieldDescriptor
	fd_MsgSend_SendCredits_retirement_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgSend_SendCredits = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgSend").Messages().ByName("SendCredits")
	fd_MsgSend_SendCredits_batch_denom = md_MsgSend_SendCredits.Fields().ByName("batch_denom")
	fd_MsgSend_SendCredits_tradable_amount = md_MsgSend_SendCredits.Fields().ByName("tradable_amount")
	fd_MsgSend_SendCredits_retired_amount = md_MsgSend_SendCredits.Fields().ByName("retired_amount")
	fd_MsgSend_SendCredits_retirement_location = md_MsgSend_SendCredits.Fields().ByName("retirement_location")
}

var _ protoreflect.Message = (*fastReflection_MsgSend_SendCredits)(nil)

type fastReflection_MsgSend_SendCredits MsgSend_SendCredits

func (x *MsgSend_SendCredits) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSend_SendCredits)(x)
}

func (x *MsgSend_SendCredits) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSend_SendCredits_messageType fastReflection_MsgSend_SendCredits_messageType
var _ protoreflect.MessageType = fastReflection_MsgSend_SendCredits_messageType{}

type fastReflection_MsgSend_SendCredits_messageType struct{}

func (x fastReflection_MsgSend_SendCredits_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSend_SendCredits)(nil)
}
func (x fastReflection_MsgSend_SendCredits_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSend_SendCredits)
}
func (x fastReflection_MsgSend_SendCredits_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend_SendCredits
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSend_SendCredits) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend_SendCredits
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSend_SendCredits) Type() protoreflect.MessageType {
	return _fastReflection_MsgSend_SendCredits_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSend_SendCredits) New() protoreflect.Message {
	return new(fastReflection_MsgSend_SendCredits)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSend_SendCredits) Interface() protoreflect.ProtoMessage {
	return (*MsgSend_SendCredits)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSend_SendCredits) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgSend_SendCredits_batch_denom, value) {
			return
		}
	}
	if x.TradableAmount != "" {
		value := protoreflect.ValueOfString(x.TradableAmount)
		if !f(fd_MsgSend_SendCredits_tradable_amount, value) {
			return
		}
	}
	if x.RetiredAmount != "" {
		value := protoreflect.ValueOfString(x.RetiredAmount)
		if !f(fd_MsgSend_SendCredits_retired_amount, value) {
			return
		}
	}
	if x.RetirementLocation != "" {
		value := protoreflect.ValueOfString(x.RetirementLocation)
		if !f(fd_MsgSend_SendCredits_retirement_location, value) {
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
func (x *fastReflection_MsgSend_SendCredits) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.tradable_amount":
		return x.TradableAmount != ""
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retired_amount":
		return x.RetiredAmount != ""
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retirement_location":
		return x.RetirementLocation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend.SendCredits does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend_SendCredits) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.tradable_amount":
		x.TradableAmount = ""
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retired_amount":
		x.RetiredAmount = ""
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retirement_location":
		x.RetirementLocation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend.SendCredits does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSend_SendCredits) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.tradable_amount":
		value := x.TradableAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retired_amount":
		value := x.RetiredAmount
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retirement_location":
		value := x.RetirementLocation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend.SendCredits does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSend_SendCredits) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.tradable_amount":
		x.TradableAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retired_amount":
		x.RetiredAmount = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retirement_location":
		x.RetirementLocation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend.SendCredits does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSend_SendCredits) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha1.MsgSend.SendCredits is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.tradable_amount":
		panic(fmt.Errorf("field tradable_amount of message regen.ecocredit.v1alpha1.MsgSend.SendCredits is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retired_amount":
		panic(fmt.Errorf("field retired_amount of message regen.ecocredit.v1alpha1.MsgSend.SendCredits is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retirement_location":
		panic(fmt.Errorf("field retirement_location of message regen.ecocredit.v1alpha1.MsgSend.SendCredits is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend.SendCredits does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSend_SendCredits) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.tradable_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retired_amount":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgSend.SendCredits.retirement_location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSend.SendCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSend.SendCredits does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSend_SendCredits) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgSend.SendCredits", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSend_SendCredits) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend_SendCredits) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSend_SendCredits) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSend_SendCredits) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSend_SendCredits)
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
		l = len(x.RetiredAmount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RetirementLocation)
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
		x := input.Message.Interface().(*MsgSend_SendCredits)
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
		if len(x.RetirementLocation) > 0 {
			i -= len(x.RetirementLocation)
			copy(dAtA[i:], x.RetirementLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetirementLocation)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.RetiredAmount) > 0 {
			i -= len(x.RetiredAmount)
			copy(dAtA[i:], x.RetiredAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RetiredAmount)))
			i--
			dAtA[i] = 0x1a
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
		x := input.Message.Interface().(*MsgSend_SendCredits)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend_SendCredits: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSend_SendCredits: illegal tag %d (wire type %d)", fieldNum, wire)
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
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetiredAmount", wireType)
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
				x.RetiredAmount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RetirementLocation", wireType)
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
				x.RetirementLocation = string(dAtA[iNdEx:postIndex])
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
	md_MsgSendResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgSendResponse = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgSendResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgSendResponse)(nil)

type fastReflection_MsgSendResponse MsgSendResponse

func (x *MsgSendResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSendResponse)(x)
}

func (x *MsgSendResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSendResponse_messageType fastReflection_MsgSendResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgSendResponse_messageType{}

type fastReflection_MsgSendResponse_messageType struct{}

func (x fastReflection_MsgSendResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSendResponse)(nil)
}
func (x fastReflection_MsgSendResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSendResponse)
}
func (x fastReflection_MsgSendResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSendResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSendResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSendResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSendResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgSendResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSendResponse) New() protoreflect.Message {
	return new(fastReflection_MsgSendResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSendResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgSendResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSendResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgSendResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSendResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSendResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSendResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSendResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSendResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSendResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSendResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgSendResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgSendResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSendResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgSendResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSendResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSendResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSendResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSendResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSendResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgSendResponse)
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
		x := input.Message.Interface().(*MsgSendResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSendResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgRetire_2_list)(nil)

type _MsgRetire_2_list struct {
	list *[]*MsgRetire_RetireCredits
}

func (x *_MsgRetire_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgRetire_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgRetire_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgRetire_RetireCredits)
	(*x.list)[i] = concreteValue
}

func (x *_MsgRetire_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgRetire_RetireCredits)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgRetire_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgRetire_RetireCredits)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgRetire_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgRetire_2_list) NewElement() protoreflect.Value {
	v := new(MsgRetire_RetireCredits)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgRetire_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgRetire          protoreflect.MessageDescriptor
	fd_MsgRetire_holder   protoreflect.FieldDescriptor
	fd_MsgRetire_credits  protoreflect.FieldDescriptor
	fd_MsgRetire_location protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgRetire = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgRetire")
	fd_MsgRetire_holder = md_MsgRetire.Fields().ByName("holder")
	fd_MsgRetire_credits = md_MsgRetire.Fields().ByName("credits")
	fd_MsgRetire_location = md_MsgRetire.Fields().ByName("location")
}

var _ protoreflect.Message = (*fastReflection_MsgRetire)(nil)

type fastReflection_MsgRetire MsgRetire

func (x *MsgRetire) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgRetire)(x)
}

func (x *MsgRetire) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgRetire_messageType fastReflection_MsgRetire_messageType
var _ protoreflect.MessageType = fastReflection_MsgRetire_messageType{}

type fastReflection_MsgRetire_messageType struct{}

func (x fastReflection_MsgRetire_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgRetire)(nil)
}
func (x fastReflection_MsgRetire_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgRetire)
}
func (x fastReflection_MsgRetire_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetire
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgRetire) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetire
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgRetire) Type() protoreflect.MessageType {
	return _fastReflection_MsgRetire_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgRetire) New() protoreflect.Message {
	return new(fastReflection_MsgRetire)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgRetire) Interface() protoreflect.ProtoMessage {
	return (*MsgRetire)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgRetire) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Holder != "" {
		value := protoreflect.ValueOfString(x.Holder)
		if !f(fd_MsgRetire_holder, value) {
			return
		}
	}
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgRetire_2_list{list: &x.Credits})
		if !f(fd_MsgRetire_credits, value) {
			return
		}
	}
	if x.Location != "" {
		value := protoreflect.ValueOfString(x.Location)
		if !f(fd_MsgRetire_location, value) {
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
func (x *fastReflection_MsgRetire) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.holder":
		return x.Holder != ""
	case "regen.ecocredit.v1alpha1.MsgRetire.credits":
		return len(x.Credits) != 0
	case "regen.ecocredit.v1alpha1.MsgRetire.location":
		return x.Location != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetire) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.holder":
		x.Holder = ""
	case "regen.ecocredit.v1alpha1.MsgRetire.credits":
		x.Credits = nil
	case "regen.ecocredit.v1alpha1.MsgRetire.location":
		x.Location = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgRetire) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.holder":
		value := x.Holder
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgRetire.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgRetire_2_list{})
		}
		listValue := &_MsgRetire_2_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	case "regen.ecocredit.v1alpha1.MsgRetire.location":
		value := x.Location
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgRetire) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.holder":
		x.Holder = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgRetire.credits":
		lv := value.List()
		clv := lv.(*_MsgRetire_2_list)
		x.Credits = *clv.list
	case "regen.ecocredit.v1alpha1.MsgRetire.location":
		x.Location = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgRetire) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.credits":
		if x.Credits == nil {
			x.Credits = []*MsgRetire_RetireCredits{}
		}
		value := &_MsgRetire_2_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha1.MsgRetire.holder":
		panic(fmt.Errorf("field holder of message regen.ecocredit.v1alpha1.MsgRetire is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgRetire.location":
		panic(fmt.Errorf("field location of message regen.ecocredit.v1alpha1.MsgRetire is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgRetire) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.holder":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgRetire.credits":
		list := []*MsgRetire_RetireCredits{}
		return protoreflect.ValueOfList(&_MsgRetire_2_list{list: &list})
	case "regen.ecocredit.v1alpha1.MsgRetire.location":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgRetire) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgRetire", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgRetire) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetire) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgRetire) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgRetire) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgRetire)
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
		l = len(x.Holder)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Location)
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
		x := input.Message.Interface().(*MsgRetire)
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
		if len(x.Location) > 0 {
			i -= len(x.Location)
			copy(dAtA[i:], x.Location)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Location)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
			}
		}
		if len(x.Holder) > 0 {
			i -= len(x.Holder)
			copy(dAtA[i:], x.Holder)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Holder)))
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
		x := input.Message.Interface().(*MsgRetire)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetire: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetire: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Holder", wireType)
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
				x.Holder = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &MsgRetire_RetireCredits{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
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
				x.Location = string(dAtA[iNdEx:postIndex])
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
	md_MsgRetire_RetireCredits             protoreflect.MessageDescriptor
	fd_MsgRetire_RetireCredits_batch_denom protoreflect.FieldDescriptor
	fd_MsgRetire_RetireCredits_amount      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgRetire_RetireCredits = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgRetire").Messages().ByName("RetireCredits")
	fd_MsgRetire_RetireCredits_batch_denom = md_MsgRetire_RetireCredits.Fields().ByName("batch_denom")
	fd_MsgRetire_RetireCredits_amount = md_MsgRetire_RetireCredits.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_MsgRetire_RetireCredits)(nil)

type fastReflection_MsgRetire_RetireCredits MsgRetire_RetireCredits

func (x *MsgRetire_RetireCredits) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgRetire_RetireCredits)(x)
}

func (x *MsgRetire_RetireCredits) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgRetire_RetireCredits_messageType fastReflection_MsgRetire_RetireCredits_messageType
var _ protoreflect.MessageType = fastReflection_MsgRetire_RetireCredits_messageType{}

type fastReflection_MsgRetire_RetireCredits_messageType struct{}

func (x fastReflection_MsgRetire_RetireCredits_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgRetire_RetireCredits)(nil)
}
func (x fastReflection_MsgRetire_RetireCredits_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgRetire_RetireCredits)
}
func (x fastReflection_MsgRetire_RetireCredits_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetire_RetireCredits
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgRetire_RetireCredits) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetire_RetireCredits
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgRetire_RetireCredits) Type() protoreflect.MessageType {
	return _fastReflection_MsgRetire_RetireCredits_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgRetire_RetireCredits) New() protoreflect.Message {
	return new(fastReflection_MsgRetire_RetireCredits)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgRetire_RetireCredits) Interface() protoreflect.ProtoMessage {
	return (*MsgRetire_RetireCredits)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgRetire_RetireCredits) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgRetire_RetireCredits_batch_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_MsgRetire_RetireCredits_amount, value) {
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
func (x *fastReflection_MsgRetire_RetireCredits) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.amount":
		return x.Amount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetire_RetireCredits) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.amount":
		x.Amount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgRetire_RetireCredits) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire.RetireCredits does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgRetire_RetireCredits) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.amount":
		x.Amount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgRetire_RetireCredits) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha1.MsgRetire.RetireCredits is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.amount":
		panic(fmt.Errorf("field amount of message regen.ecocredit.v1alpha1.MsgRetire.RetireCredits is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgRetire_RetireCredits) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgRetire.RetireCredits.amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetire.RetireCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetire.RetireCredits does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgRetire_RetireCredits) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgRetire.RetireCredits", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgRetire_RetireCredits) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetire_RetireCredits) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgRetire_RetireCredits) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgRetire_RetireCredits) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgRetire_RetireCredits)
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
		l = len(x.Amount)
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
		x := input.Message.Interface().(*MsgRetire_RetireCredits)
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
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
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
		x := input.Message.Interface().(*MsgRetire_RetireCredits)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetire_RetireCredits: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetire_RetireCredits: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
				x.Amount = string(dAtA[iNdEx:postIndex])
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
	md_MsgRetireResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgRetireResponse = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgRetireResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgRetireResponse)(nil)

type fastReflection_MsgRetireResponse MsgRetireResponse

func (x *MsgRetireResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgRetireResponse)(x)
}

func (x *MsgRetireResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgRetireResponse_messageType fastReflection_MsgRetireResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgRetireResponse_messageType{}

type fastReflection_MsgRetireResponse_messageType struct{}

func (x fastReflection_MsgRetireResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgRetireResponse)(nil)
}
func (x fastReflection_MsgRetireResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgRetireResponse)
}
func (x fastReflection_MsgRetireResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetireResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgRetireResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgRetireResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgRetireResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgRetireResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgRetireResponse) New() protoreflect.Message {
	return new(fastReflection_MsgRetireResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgRetireResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgRetireResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgRetireResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgRetireResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetireResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetireResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetireResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgRetireResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetireResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgRetireResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetireResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgRetireResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetireResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgRetireResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgRetireResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgRetireResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgRetireResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgRetireResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgRetireResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgRetireResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgRetireResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgRetireResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgRetireResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgRetireResponse)
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
		x := input.Message.Interface().(*MsgRetireResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetireResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgRetireResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgCancel_2_list)(nil)

type _MsgCancel_2_list struct {
	list *[]*MsgCancel_CancelCredits
}

func (x *_MsgCancel_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCancel_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCancel_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgCancel_CancelCredits)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCancel_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgCancel_CancelCredits)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCancel_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgCancel_CancelCredits)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCancel_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCancel_2_list) NewElement() protoreflect.Value {
	v := new(MsgCancel_CancelCredits)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCancel_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCancel         protoreflect.MessageDescriptor
	fd_MsgCancel_holder  protoreflect.FieldDescriptor
	fd_MsgCancel_credits protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgCancel = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgCancel")
	fd_MsgCancel_holder = md_MsgCancel.Fields().ByName("holder")
	fd_MsgCancel_credits = md_MsgCancel.Fields().ByName("credits")
}

var _ protoreflect.Message = (*fastReflection_MsgCancel)(nil)

type fastReflection_MsgCancel MsgCancel

func (x *MsgCancel) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCancel)(x)
}

func (x *MsgCancel) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCancel_messageType fastReflection_MsgCancel_messageType
var _ protoreflect.MessageType = fastReflection_MsgCancel_messageType{}

type fastReflection_MsgCancel_messageType struct{}

func (x fastReflection_MsgCancel_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCancel)(nil)
}
func (x fastReflection_MsgCancel_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCancel)
}
func (x fastReflection_MsgCancel_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancel
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCancel) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancel
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCancel) Type() protoreflect.MessageType {
	return _fastReflection_MsgCancel_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCancel) New() protoreflect.Message {
	return new(fastReflection_MsgCancel)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCancel) Interface() protoreflect.ProtoMessage {
	return (*MsgCancel)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCancel) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Holder != "" {
		value := protoreflect.ValueOfString(x.Holder)
		if !f(fd_MsgCancel_holder, value) {
			return
		}
	}
	if len(x.Credits) != 0 {
		value := protoreflect.ValueOfList(&_MsgCancel_2_list{list: &x.Credits})
		if !f(fd_MsgCancel_credits, value) {
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
func (x *fastReflection_MsgCancel) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.holder":
		return x.Holder != ""
	case "regen.ecocredit.v1alpha1.MsgCancel.credits":
		return len(x.Credits) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancel) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.holder":
		x.Holder = ""
	case "regen.ecocredit.v1alpha1.MsgCancel.credits":
		x.Credits = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCancel) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.holder":
		value := x.Holder
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgCancel.credits":
		if len(x.Credits) == 0 {
			return protoreflect.ValueOfList(&_MsgCancel_2_list{})
		}
		listValue := &_MsgCancel_2_list{list: &x.Credits}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCancel) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.holder":
		x.Holder = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgCancel.credits":
		lv := value.List()
		clv := lv.(*_MsgCancel_2_list)
		x.Credits = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCancel) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.credits":
		if x.Credits == nil {
			x.Credits = []*MsgCancel_CancelCredits{}
		}
		value := &_MsgCancel_2_list{list: &x.Credits}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha1.MsgCancel.holder":
		panic(fmt.Errorf("field holder of message regen.ecocredit.v1alpha1.MsgCancel is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCancel) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.holder":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgCancel.credits":
		list := []*MsgCancel_CancelCredits{}
		return protoreflect.ValueOfList(&_MsgCancel_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCancel) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgCancel", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCancel) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancel) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCancel) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCancel) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCancel)
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
		l = len(x.Holder)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Credits) > 0 {
			for _, e := range x.Credits {
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
		x := input.Message.Interface().(*MsgCancel)
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
		if len(x.Credits) > 0 {
			for iNdEx := len(x.Credits) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Credits[iNdEx])
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
			}
		}
		if len(x.Holder) > 0 {
			i -= len(x.Holder)
			copy(dAtA[i:], x.Holder)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Holder)))
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
		x := input.Message.Interface().(*MsgCancel)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancel: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancel: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Holder", wireType)
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
				x.Holder = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
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
				x.Credits = append(x.Credits, &MsgCancel_CancelCredits{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Credits[len(x.Credits)-1]); err != nil {
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
	md_MsgCancel_CancelCredits             protoreflect.MessageDescriptor
	fd_MsgCancel_CancelCredits_batch_denom protoreflect.FieldDescriptor
	fd_MsgCancel_CancelCredits_amount      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgCancel_CancelCredits = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgCancel").Messages().ByName("CancelCredits")
	fd_MsgCancel_CancelCredits_batch_denom = md_MsgCancel_CancelCredits.Fields().ByName("batch_denom")
	fd_MsgCancel_CancelCredits_amount = md_MsgCancel_CancelCredits.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_MsgCancel_CancelCredits)(nil)

type fastReflection_MsgCancel_CancelCredits MsgCancel_CancelCredits

func (x *MsgCancel_CancelCredits) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCancel_CancelCredits)(x)
}

func (x *MsgCancel_CancelCredits) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCancel_CancelCredits_messageType fastReflection_MsgCancel_CancelCredits_messageType
var _ protoreflect.MessageType = fastReflection_MsgCancel_CancelCredits_messageType{}

type fastReflection_MsgCancel_CancelCredits_messageType struct{}

func (x fastReflection_MsgCancel_CancelCredits_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCancel_CancelCredits)(nil)
}
func (x fastReflection_MsgCancel_CancelCredits_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCancel_CancelCredits)
}
func (x fastReflection_MsgCancel_CancelCredits_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancel_CancelCredits
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCancel_CancelCredits) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancel_CancelCredits
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCancel_CancelCredits) Type() protoreflect.MessageType {
	return _fastReflection_MsgCancel_CancelCredits_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCancel_CancelCredits) New() protoreflect.Message {
	return new(fastReflection_MsgCancel_CancelCredits)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCancel_CancelCredits) Interface() protoreflect.ProtoMessage {
	return (*MsgCancel_CancelCredits)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCancel_CancelCredits) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_MsgCancel_CancelCredits_batch_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_MsgCancel_CancelCredits_amount, value) {
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
func (x *fastReflection_MsgCancel_CancelCredits) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.amount":
		return x.Amount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancel_CancelCredits) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.amount":
		x.Amount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCancel_CancelCredits) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel.CancelCredits does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCancel_CancelCredits) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.amount":
		x.Amount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCancel_CancelCredits) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.v1alpha1.MsgCancel.CancelCredits is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.amount":
		panic(fmt.Errorf("field amount of message regen.ecocredit.v1alpha1.MsgCancel.CancelCredits is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCancel_CancelCredits) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgCancel.CancelCredits.amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancel.CancelCredits"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancel.CancelCredits does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCancel_CancelCredits) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgCancel.CancelCredits", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCancel_CancelCredits) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancel_CancelCredits) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCancel_CancelCredits) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCancel_CancelCredits) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCancel_CancelCredits)
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
		l = len(x.Amount)
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
		x := input.Message.Interface().(*MsgCancel_CancelCredits)
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
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
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
		x := input.Message.Interface().(*MsgCancel_CancelCredits)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancel_CancelCredits: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancel_CancelCredits: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
				x.Amount = string(dAtA[iNdEx:postIndex])
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
	md_MsgCancelResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgCancelResponse = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgCancelResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgCancelResponse)(nil)

type fastReflection_MsgCancelResponse MsgCancelResponse

func (x *MsgCancelResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCancelResponse)(x)
}

func (x *MsgCancelResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCancelResponse_messageType fastReflection_MsgCancelResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCancelResponse_messageType{}

type fastReflection_MsgCancelResponse_messageType struct{}

func (x fastReflection_MsgCancelResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCancelResponse)(nil)
}
func (x fastReflection_MsgCancelResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCancelResponse)
}
func (x fastReflection_MsgCancelResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancelResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCancelResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCancelResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCancelResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCancelResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCancelResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCancelResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCancelResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCancelResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCancelResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgCancelResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancelResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancelResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancelResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCancelResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancelResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgCancelResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancelResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgCancelResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancelResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCancelResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgCancelResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgCancelResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCancelResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgCancelResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCancelResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCancelResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgCancelResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCancelResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCancelResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCancelResponse)
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
		x := input.Message.Interface().(*MsgCancelResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancelResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCancelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_MsgUpdateClassAdmin           protoreflect.MessageDescriptor
	fd_MsgUpdateClassAdmin_admin     protoreflect.FieldDescriptor
	fd_MsgUpdateClassAdmin_class_id  protoreflect.FieldDescriptor
	fd_MsgUpdateClassAdmin_new_admin protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgUpdateClassAdmin = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgUpdateClassAdmin")
	fd_MsgUpdateClassAdmin_admin = md_MsgUpdateClassAdmin.Fields().ByName("admin")
	fd_MsgUpdateClassAdmin_class_id = md_MsgUpdateClassAdmin.Fields().ByName("class_id")
	fd_MsgUpdateClassAdmin_new_admin = md_MsgUpdateClassAdmin.Fields().ByName("new_admin")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassAdmin)(nil)

type fastReflection_MsgUpdateClassAdmin MsgUpdateClassAdmin

func (x *MsgUpdateClassAdmin) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassAdmin)(x)
}

func (x *MsgUpdateClassAdmin) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassAdmin_messageType fastReflection_MsgUpdateClassAdmin_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassAdmin_messageType{}

type fastReflection_MsgUpdateClassAdmin_messageType struct{}

func (x fastReflection_MsgUpdateClassAdmin_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassAdmin)(nil)
}
func (x fastReflection_MsgUpdateClassAdmin_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassAdmin)
}
func (x fastReflection_MsgUpdateClassAdmin_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassAdmin
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassAdmin) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassAdmin
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassAdmin) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassAdmin_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassAdmin) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassAdmin)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassAdmin) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassAdmin)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassAdmin) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateClassAdmin_admin, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgUpdateClassAdmin_class_id, value) {
			return
		}
	}
	if x.NewAdmin != "" {
		value := protoreflect.ValueOfString(x.NewAdmin)
		if !f(fd_MsgUpdateClassAdmin_new_admin, value) {
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
func (x *fastReflection_MsgUpdateClassAdmin) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.new_admin":
		return x.NewAdmin != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassAdmin) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.new_admin":
		x.NewAdmin = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassAdmin) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.new_admin":
		value := x.NewAdmin
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassAdmin) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.new_admin":
		x.NewAdmin = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassAdmin) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.new_admin":
		panic(fmt.Errorf("field new_admin of message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassAdmin) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgUpdateClassAdmin.new_admin":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdmin"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdmin does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassAdmin) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgUpdateClassAdmin", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassAdmin) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassAdmin) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassAdmin) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassAdmin) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassAdmin)
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
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NewAdmin)
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
		x := input.Message.Interface().(*MsgUpdateClassAdmin)
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
		if len(x.NewAdmin) > 0 {
			i -= len(x.NewAdmin)
			copy(dAtA[i:], x.NewAdmin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NewAdmin)))
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
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
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
		x := input.Message.Interface().(*MsgUpdateClassAdmin)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassAdmin: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewAdmin", wireType)
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
				x.NewAdmin = string(dAtA[iNdEx:postIndex])
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
	md_MsgUpdateClassAdminResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgUpdateClassAdminResponse = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgUpdateClassAdminResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassAdminResponse)(nil)

type fastReflection_MsgUpdateClassAdminResponse MsgUpdateClassAdminResponse

func (x *MsgUpdateClassAdminResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassAdminResponse)(x)
}

func (x *MsgUpdateClassAdminResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassAdminResponse_messageType fastReflection_MsgUpdateClassAdminResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassAdminResponse_messageType{}

type fastReflection_MsgUpdateClassAdminResponse_messageType struct{}

func (x fastReflection_MsgUpdateClassAdminResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassAdminResponse)(nil)
}
func (x fastReflection_MsgUpdateClassAdminResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassAdminResponse)
}
func (x fastReflection_MsgUpdateClassAdminResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassAdminResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassAdminResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassAdminResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassAdminResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassAdminResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassAdminResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassAdminResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassAdminResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassAdminResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassAdminResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgUpdateClassAdminResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassAdminResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassAdminResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassAdminResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassAdminResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassAdminResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassAdminResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassAdminResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassAdminResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassAdminResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassAdminResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassAdminResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateClassAdminResponse)
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
		x := input.Message.Interface().(*MsgUpdateClassAdminResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassAdminResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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

var _ protoreflect.List = (*_MsgUpdateClassIssuers_3_list)(nil)

type _MsgUpdateClassIssuers_3_list struct {
	list *[]string
}

func (x *_MsgUpdateClassIssuers_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgUpdateClassIssuers_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_MsgUpdateClassIssuers_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgUpdateClassIssuers_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgUpdateClassIssuers_3_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgUpdateClassIssuers at list field Issuers as it is not of Message kind"))
}

func (x *_MsgUpdateClassIssuers_3_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgUpdateClassIssuers_3_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_MsgUpdateClassIssuers_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgUpdateClassIssuers          protoreflect.MessageDescriptor
	fd_MsgUpdateClassIssuers_admin    protoreflect.FieldDescriptor
	fd_MsgUpdateClassIssuers_class_id protoreflect.FieldDescriptor
	fd_MsgUpdateClassIssuers_issuers  protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgUpdateClassIssuers = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgUpdateClassIssuers")
	fd_MsgUpdateClassIssuers_admin = md_MsgUpdateClassIssuers.Fields().ByName("admin")
	fd_MsgUpdateClassIssuers_class_id = md_MsgUpdateClassIssuers.Fields().ByName("class_id")
	fd_MsgUpdateClassIssuers_issuers = md_MsgUpdateClassIssuers.Fields().ByName("issuers")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassIssuers)(nil)

type fastReflection_MsgUpdateClassIssuers MsgUpdateClassIssuers

func (x *MsgUpdateClassIssuers) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassIssuers)(x)
}

func (x *MsgUpdateClassIssuers) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassIssuers_messageType fastReflection_MsgUpdateClassIssuers_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassIssuers_messageType{}

type fastReflection_MsgUpdateClassIssuers_messageType struct{}

func (x fastReflection_MsgUpdateClassIssuers_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassIssuers)(nil)
}
func (x fastReflection_MsgUpdateClassIssuers_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassIssuers)
}
func (x fastReflection_MsgUpdateClassIssuers_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassIssuers
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassIssuers) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassIssuers
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassIssuers) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassIssuers_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassIssuers) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassIssuers)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassIssuers) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassIssuers)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassIssuers) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateClassIssuers_admin, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgUpdateClassIssuers_class_id, value) {
			return
		}
	}
	if len(x.Issuers) != 0 {
		value := protoreflect.ValueOfList(&_MsgUpdateClassIssuers_3_list{list: &x.Issuers})
		if !f(fd_MsgUpdateClassIssuers_issuers, value) {
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
func (x *fastReflection_MsgUpdateClassIssuers) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.issuers":
		return len(x.Issuers) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassIssuers) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.issuers":
		x.Issuers = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassIssuers) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.issuers":
		if len(x.Issuers) == 0 {
			return protoreflect.ValueOfList(&_MsgUpdateClassIssuers_3_list{})
		}
		listValue := &_MsgUpdateClassIssuers_3_list{list: &x.Issuers}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuers does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassIssuers) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.issuers":
		lv := value.List()
		clv := lv.(*_MsgUpdateClassIssuers_3_list)
		x.Issuers = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassIssuers) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.issuers":
		if x.Issuers == nil {
			x.Issuers = []string{}
		}
		value := &_MsgUpdateClassIssuers_3_list{list: &x.Issuers}
		return protoreflect.ValueOfList(value)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha1.MsgUpdateClassIssuers is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha1.MsgUpdateClassIssuers is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassIssuers) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgUpdateClassIssuers.issuers":
		list := []string{}
		return protoreflect.ValueOfList(&_MsgUpdateClassIssuers_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuers"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuers does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassIssuers) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgUpdateClassIssuers", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassIssuers) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassIssuers) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassIssuers) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassIssuers) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassIssuers)
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
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Issuers) > 0 {
			for _, s := range x.Issuers {
				l = len(s)
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
		x := input.Message.Interface().(*MsgUpdateClassIssuers)
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
		if len(x.Issuers) > 0 {
			for iNdEx := len(x.Issuers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Issuers[iNdEx])
				copy(dAtA[i:], x.Issuers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Issuers[iNdEx])))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
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
		x := input.Message.Interface().(*MsgUpdateClassIssuers)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassIssuers: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassIssuers: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
	md_MsgUpdateClassIssuersResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgUpdateClassIssuersResponse = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgUpdateClassIssuersResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassIssuersResponse)(nil)

type fastReflection_MsgUpdateClassIssuersResponse MsgUpdateClassIssuersResponse

func (x *MsgUpdateClassIssuersResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassIssuersResponse)(x)
}

func (x *MsgUpdateClassIssuersResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassIssuersResponse_messageType fastReflection_MsgUpdateClassIssuersResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassIssuersResponse_messageType{}

type fastReflection_MsgUpdateClassIssuersResponse_messageType struct{}

func (x fastReflection_MsgUpdateClassIssuersResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassIssuersResponse)(nil)
}
func (x fastReflection_MsgUpdateClassIssuersResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassIssuersResponse)
}
func (x fastReflection_MsgUpdateClassIssuersResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassIssuersResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassIssuersResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassIssuersResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassIssuersResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassIssuersResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassIssuersResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgUpdateClassIssuersResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassIssuersResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassIssuersResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassIssuersResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassIssuersResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassIssuersResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassIssuersResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassIssuersResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassIssuersResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassIssuersResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassIssuersResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateClassIssuersResponse)
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
		x := input.Message.Interface().(*MsgUpdateClassIssuersResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassIssuersResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassIssuersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_MsgUpdateClassMetadata          protoreflect.MessageDescriptor
	fd_MsgUpdateClassMetadata_admin    protoreflect.FieldDescriptor
	fd_MsgUpdateClassMetadata_class_id protoreflect.FieldDescriptor
	fd_MsgUpdateClassMetadata_metadata protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgUpdateClassMetadata = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgUpdateClassMetadata")
	fd_MsgUpdateClassMetadata_admin = md_MsgUpdateClassMetadata.Fields().ByName("admin")
	fd_MsgUpdateClassMetadata_class_id = md_MsgUpdateClassMetadata.Fields().ByName("class_id")
	fd_MsgUpdateClassMetadata_metadata = md_MsgUpdateClassMetadata.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassMetadata)(nil)

type fastReflection_MsgUpdateClassMetadata MsgUpdateClassMetadata

func (x *MsgUpdateClassMetadata) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassMetadata)(x)
}

func (x *MsgUpdateClassMetadata) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassMetadata_messageType fastReflection_MsgUpdateClassMetadata_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassMetadata_messageType{}

type fastReflection_MsgUpdateClassMetadata_messageType struct{}

func (x fastReflection_MsgUpdateClassMetadata_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassMetadata)(nil)
}
func (x fastReflection_MsgUpdateClassMetadata_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassMetadata)
}
func (x fastReflection_MsgUpdateClassMetadata_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassMetadata
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassMetadata) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassMetadata
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassMetadata) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassMetadata_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassMetadata) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassMetadata)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassMetadata) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassMetadata)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassMetadata) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateClassMetadata_admin, value) {
			return
		}
	}
	if x.ClassId != "" {
		value := protoreflect.ValueOfString(x.ClassId)
		if !f(fd_MsgUpdateClassMetadata_class_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgUpdateClassMetadata_metadata, value) {
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
func (x *fastReflection_MsgUpdateClassMetadata) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.admin":
		return x.Admin != ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.class_id":
		return x.ClassId != ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassMetadata) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.admin":
		x.Admin = ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.class_id":
		x.ClassId = ""
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassMetadata) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.class_id":
		value := x.ClassId
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassMetadata) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.admin":
		x.Admin = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.class_id":
		x.ClassId = value.Interface().(string)
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassMetadata) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.admin":
		panic(fmt.Errorf("field admin of message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata is not mutable"))
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.metadata":
		panic(fmt.Errorf("field metadata of message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassMetadata) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.admin":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.v1alpha1.MsgUpdateClassMetadata.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadata"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadata does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassMetadata) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgUpdateClassMetadata", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassMetadata) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassMetadata) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassMetadata) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassMetadata) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassMetadata)
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
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ClassId)
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
		x := input.Message.Interface().(*MsgUpdateClassMetadata)
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
			dAtA[i] = 0x1a
		}
		if len(x.ClassId) > 0 {
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
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
		x := input.Message.Interface().(*MsgUpdateClassMetadata)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassMetadata: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
	md_MsgUpdateClassMetadataResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_ecocredit_v1alpha1_tx_proto_init()
	md_MsgUpdateClassMetadataResponse = File_regen_ecocredit_v1alpha1_tx_proto.Messages().ByName("MsgUpdateClassMetadataResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateClassMetadataResponse)(nil)

type fastReflection_MsgUpdateClassMetadataResponse MsgUpdateClassMetadataResponse

func (x *MsgUpdateClassMetadataResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassMetadataResponse)(x)
}

func (x *MsgUpdateClassMetadataResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateClassMetadataResponse_messageType fastReflection_MsgUpdateClassMetadataResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateClassMetadataResponse_messageType{}

type fastReflection_MsgUpdateClassMetadataResponse_messageType struct{}

func (x fastReflection_MsgUpdateClassMetadataResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateClassMetadataResponse)(nil)
}
func (x fastReflection_MsgUpdateClassMetadataResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassMetadataResponse)
}
func (x fastReflection_MsgUpdateClassMetadataResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassMetadataResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateClassMetadataResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateClassMetadataResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateClassMetadataResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateClassMetadataResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateClassMetadataResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgUpdateClassMetadataResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateClassMetadataResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgUpdateClassMetadataResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgUpdateClassMetadataResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateClassMetadataResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateClassMetadataResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateClassMetadataResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateClassMetadataResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgUpdateClassMetadataResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateClassMetadataResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateClassMetadataResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateClassMetadataResponse)
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
		x := input.Message.Interface().(*MsgUpdateClassMetadataResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassMetadataResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateClassMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
// source: regen/ecocredit/v1alpha1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgCreateClass is the Msg/CreateClass request type.
type MsgCreateClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the address of the account that created the credit class.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// issuers are the account addresses of the approved issuers.
	Issuers []string `protobuf:"bytes,2,rep,name=issuers,proto3" json:"issuers,omitempty"`
	// metadata is any arbitrary metadata to attached to the credit class.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// credit_type_name describes the type of credit (e.g. "carbon", "biodiversity").
	CreditTypeName string `protobuf:"bytes,4,opt,name=credit_type_name,json=creditTypeName,proto3" json:"credit_type_name,omitempty"`
}

func (x *MsgCreateClass) Reset() {
	*x = MsgCreateClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateClass) ProtoMessage() {}

// Deprecated: Use MsgCreateClass.ProtoReflect.Descriptor instead.
func (*MsgCreateClass) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreateClass) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgCreateClass) GetIssuers() []string {
	if x != nil {
		return x.Issuers
	}
	return nil
}

func (x *MsgCreateClass) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateClass) GetCreditTypeName() string {
	if x != nil {
		return x.CreditTypeName
	}
	return ""
}

// MsgCreateClassResponse is the Msg/CreateClass response type.
type MsgCreateClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// class_id is the unique ID of the newly created credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *MsgCreateClassResponse) Reset() {
	*x = MsgCreateClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateClassResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateClassResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateClassResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCreateClassResponse) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

// MsgCreateBatch is the Msg/CreateBatch request type.
type MsgCreateBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// issuer is the address of the batch issuer.
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// class_id is the unique ID of the class.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// issuance are the credits issued in the batch.
	Issuance []*MsgCreateBatch_BatchIssuance `protobuf:"bytes,3,rep,name=issuance,proto3" json:"issuance,omitempty"`
	// metadata is any arbitrary metadata attached to the credit batch.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// start_date is the beginning of the period during which this credit batch
	// was quantified and verified.
	StartDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// end_date is the end of the period during which this credit batch was
	// quantified and verified.
	EndDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// project_location is the location of the project backing the credits in this
	// batch. It is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters. country-code is required, while sub-national-code
	// and postal-code can be added for increasing precision.
	ProjectLocation string `protobuf:"bytes,7,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
}

func (x *MsgCreateBatch) Reset() {
	*x = MsgCreateBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBatch) ProtoMessage() {}

// Deprecated: Use MsgCreateBatch.ProtoReflect.Descriptor instead.
func (*MsgCreateBatch) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgCreateBatch) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *MsgCreateBatch) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *MsgCreateBatch) GetIssuance() []*MsgCreateBatch_BatchIssuance {
	if x != nil {
		return x.Issuance
	}
	return nil
}

func (x *MsgCreateBatch) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateBatch) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *MsgCreateBatch) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *MsgCreateBatch) GetProjectLocation() string {
	if x != nil {
		return x.ProjectLocation
	}
	return ""
}

// MsgCreateBatchResponse is the Msg/CreateBatch response type.
type MsgCreateBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique denomination ID of the newly created batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
}

func (x *MsgCreateBatchResponse) Reset() {
	*x = MsgCreateBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBatchResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateBatchResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateBatchResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{3}
}

func (x *MsgCreateBatchResponse) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

// MsgSend is the Msg/Send request type.
type MsgSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sender is the address of the account sending credits.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// sender is the address of the account receiving credits.
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// credits are the credits being sent.
	Credits []*MsgSend_SendCredits `protobuf:"bytes,3,rep,name=credits,proto3" json:"credits,omitempty"`
}

func (x *MsgSend) Reset() {
	*x = MsgSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSend) ProtoMessage() {}

// Deprecated: Use MsgSend.ProtoReflect.Descriptor instead.
func (*MsgSend) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{4}
}

func (x *MsgSend) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MsgSend) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *MsgSend) GetCredits() []*MsgSend_SendCredits {
	if x != nil {
		return x.Credits
	}
	return nil
}

// MsgSendResponse is the Msg/Send response type.
type MsgSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgSendResponse) Reset() {
	*x = MsgSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSendResponse) ProtoMessage() {}

// Deprecated: Use MsgSendResponse.ProtoReflect.Descriptor instead.
func (*MsgSendResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{5}
}

// MsgRetire is the Msg/Retire request type.
type MsgRetire struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// holder is the credit holder address.
	Holder string `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	// credits are the credits being retired.
	Credits []*MsgRetire_RetireCredits `protobuf:"bytes,2,rep,name=credits,proto3" json:"credits,omitempty"`
	// location is the location of the beneficiary or buyer of the retired
	// credits. It is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters.
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *MsgRetire) Reset() {
	*x = MsgRetire{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRetire) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRetire) ProtoMessage() {}

// Deprecated: Use MsgRetire.ProtoReflect.Descriptor instead.
func (*MsgRetire) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{6}
}

func (x *MsgRetire) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

func (x *MsgRetire) GetCredits() []*MsgRetire_RetireCredits {
	if x != nil {
		return x.Credits
	}
	return nil
}

func (x *MsgRetire) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

// MsgRetire is the Msg/Retire response type.
type MsgRetireResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgRetireResponse) Reset() {
	*x = MsgRetireResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRetireResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRetireResponse) ProtoMessage() {}

// Deprecated: Use MsgRetireResponse.ProtoReflect.Descriptor instead.
func (*MsgRetireResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{7}
}

// MsgCancel is the Msg/Cancel request type.
type MsgCancel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// holder is the credit holder address.
	Holder string `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	// credits are the credits being cancelled.
	Credits []*MsgCancel_CancelCredits `protobuf:"bytes,2,rep,name=credits,proto3" json:"credits,omitempty"`
}

func (x *MsgCancel) Reset() {
	*x = MsgCancel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCancel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCancel) ProtoMessage() {}

// Deprecated: Use MsgCancel.ProtoReflect.Descriptor instead.
func (*MsgCancel) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{8}
}

func (x *MsgCancel) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

func (x *MsgCancel) GetCredits() []*MsgCancel_CancelCredits {
	if x != nil {
		return x.Credits
	}
	return nil
}

// MsgCancelResponse is the Msg/Cancel response type.
type MsgCancelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgCancelResponse) Reset() {
	*x = MsgCancelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCancelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCancelResponse) ProtoMessage() {}

// Deprecated: Use MsgCancelResponse.ProtoReflect.Descriptor instead.
func (*MsgCancelResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{9}
}

// MsgUpdateClassAdmin is the Msg/UpdateClassAdmin request type.
type MsgUpdateClassAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the address of the account that is the admin of the credit class.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// class_id is the unique ID of the credit class.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// new_admin is the address of the new admin of the credit class.
	NewAdmin string `protobuf:"bytes,3,opt,name=new_admin,json=newAdmin,proto3" json:"new_admin,omitempty"`
}

func (x *MsgUpdateClassAdmin) Reset() {
	*x = MsgUpdateClassAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassAdmin) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassAdmin.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassAdmin) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{10}
}

func (x *MsgUpdateClassAdmin) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateClassAdmin) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *MsgUpdateClassAdmin) GetNewAdmin() string {
	if x != nil {
		return x.NewAdmin
	}
	return ""
}

// MsgUpdateClassAdminResponse is the MsgUpdateClassAdmin response type.
type MsgUpdateClassAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateClassAdminResponse) Reset() {
	*x = MsgUpdateClassAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassAdminResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassAdminResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassAdminResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{11}
}

// MsgUpdateClassIssuers is the Msg/UpdateClassIssuers request type.
type MsgUpdateClassIssuers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the address of the account that is the admin of the credit class.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// class_id is the unique ID of the credit class.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// issuers are the updated account addresses of the approved issuers.
	Issuers []string `protobuf:"bytes,3,rep,name=issuers,proto3" json:"issuers,omitempty"`
}

func (x *MsgUpdateClassIssuers) Reset() {
	*x = MsgUpdateClassIssuers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassIssuers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassIssuers) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassIssuers.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassIssuers) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{12}
}

func (x *MsgUpdateClassIssuers) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateClassIssuers) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *MsgUpdateClassIssuers) GetIssuers() []string {
	if x != nil {
		return x.Issuers
	}
	return nil
}

// MsgUpdateClassIssuersResponse is the MsgUpdateClassIssuers response type.
type MsgUpdateClassIssuersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateClassIssuersResponse) Reset() {
	*x = MsgUpdateClassIssuersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassIssuersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassIssuersResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassIssuersResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassIssuersResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{13}
}

// MsgUpdateClassMetadata is the Msg/UpdateClassMetadata request type.
type MsgUpdateClassMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the address of the account that is the admin of the credit class.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// class_id is the unique ID of the credit class.
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// metadata is the updated arbitrary metadata to be attached to the credit class.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MsgUpdateClassMetadata) Reset() {
	*x = MsgUpdateClassMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassMetadata) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassMetadata.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassMetadata) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{14}
}

func (x *MsgUpdateClassMetadata) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateClassMetadata) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *MsgUpdateClassMetadata) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgUpdateClassMetadataResponse is the MsgUpdateClassMetadata response type.
type MsgUpdateClassMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateClassMetadataResponse) Reset() {
	*x = MsgUpdateClassMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateClassMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateClassMetadataResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateClassMetadataResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateClassMetadataResponse) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{15}
}

// BatchIssuance represents the issuance of some credits in a batch to a
// single recipient.
type MsgCreateBatch_BatchIssuance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// recipient is the account of the recipient.
	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// tradable_amount is the number of credits in this issuance that can be
	// traded by this recipient. Decimal values are acceptable.
	TradableAmount string `protobuf:"bytes,2,opt,name=tradable_amount,json=tradableAmount,proto3" json:"tradable_amount,omitempty"`
	// retired_amount is the number of credits in this issuance that are
	// effectively retired by the issuer on receipt. Decimal values are
	// acceptable.
	RetiredAmount string `protobuf:"bytes,3,opt,name=retired_amount,json=retiredAmount,proto3" json:"retired_amount,omitempty"`
	// retirement_location is the location of the beneficiary or buyer of the
	// retired credits. This must be provided if retired_amount is positive. It
	// is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters.
	RetirementLocation string `protobuf:"bytes,4,opt,name=retirement_location,json=retirementLocation,proto3" json:"retirement_location,omitempty"`
}

func (x *MsgCreateBatch_BatchIssuance) Reset() {
	*x = MsgCreateBatch_BatchIssuance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBatch_BatchIssuance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBatch_BatchIssuance) ProtoMessage() {}

// Deprecated: Use MsgCreateBatch_BatchIssuance.ProtoReflect.Descriptor instead.
func (*MsgCreateBatch_BatchIssuance) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{2, 0}
}

func (x *MsgCreateBatch_BatchIssuance) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *MsgCreateBatch_BatchIssuance) GetTradableAmount() string {
	if x != nil {
		return x.TradableAmount
	}
	return ""
}

func (x *MsgCreateBatch_BatchIssuance) GetRetiredAmount() string {
	if x != nil {
		return x.RetiredAmount
	}
	return ""
}

func (x *MsgCreateBatch_BatchIssuance) GetRetirementLocation() string {
	if x != nil {
		return x.RetirementLocation
	}
	return ""
}

// SendCredits specifies a batch and the number of credits being transferred.
// This is split into tradable credits, which will remain tradable on receipt,
// and retired credits, which will be retired on receipt.
type MsgSend_SendCredits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// tradable_amount is the number of credits in this transfer that can be
	// traded by the recipient. Decimal values are acceptable within the
	// precision returned by Query/Precision.
	TradableAmount string `protobuf:"bytes,2,opt,name=tradable_amount,json=tradableAmount,proto3" json:"tradable_amount,omitempty"`
	// retired_amount is the number of credits in this transfer that are
	// effectively retired by the issuer on receipt. Decimal values are
	// acceptable within the precision returned by Query/Precision.
	RetiredAmount string `protobuf:"bytes,3,opt,name=retired_amount,json=retiredAmount,proto3" json:"retired_amount,omitempty"`
	// retirement_location is the location of the beneficiary or buyer of the
	// retired credits. This must be provided if retired_amount is positive. It
	// is a string of the form
	// <country-code>[-<sub-national-code>[ <postal-code>]], with the first two
	// fields conforming to ISO 3166-2, and postal-code being up to 64
	// alphanumeric characters.
	RetirementLocation string `protobuf:"bytes,4,opt,name=retirement_location,json=retirementLocation,proto3" json:"retirement_location,omitempty"`
}

func (x *MsgSend_SendCredits) Reset() {
	*x = MsgSend_SendCredits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSend_SendCredits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSend_SendCredits) ProtoMessage() {}

// Deprecated: Use MsgSend_SendCredits.ProtoReflect.Descriptor instead.
func (*MsgSend_SendCredits) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{4, 0}
}

func (x *MsgSend_SendCredits) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *MsgSend_SendCredits) GetTradableAmount() string {
	if x != nil {
		return x.TradableAmount
	}
	return ""
}

func (x *MsgSend_SendCredits) GetRetiredAmount() string {
	if x != nil {
		return x.RetiredAmount
	}
	return ""
}

func (x *MsgSend_SendCredits) GetRetirementLocation() string {
	if x != nil {
		return x.RetirementLocation
	}
	return ""
}

// RetireCredits specifies a batch and the number of credits being retired.
type MsgRetire_RetireCredits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the number of credits being retired.
	// Decimal values are acceptable within the precision returned by
	// Query/Precision.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgRetire_RetireCredits) Reset() {
	*x = MsgRetire_RetireCredits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRetire_RetireCredits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRetire_RetireCredits) ProtoMessage() {}

// Deprecated: Use MsgRetire_RetireCredits.ProtoReflect.Descriptor instead.
func (*MsgRetire_RetireCredits) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{6, 0}
}

func (x *MsgRetire_RetireCredits) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *MsgRetire_RetireCredits) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

// CancelCredits specifies a batch and the number of credits being cancelled.
type MsgCancel_CancelCredits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the number of credits being cancelled.
	// Decimal values are acceptable within the precision returned by
	// Query/Precision.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgCancel_CancelCredits) Reset() {
	*x = MsgCancel_CancelCredits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCancel_CancelCredits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCancel_CancelCredits) ProtoMessage() {}

// Deprecated: Use MsgCancel_CancelCredits.ProtoReflect.Descriptor instead.
func (*MsgCancel_CancelCredits) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP(), []int{8, 0}
}

func (x *MsgCancel_CancelCredits) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *MsgCancel_CancelCredits) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_regen_ecocredit_v1alpha1_tx_proto protoreflect.FileDescriptor

var file_regen_ecocredit_v1alpha1_tx_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a,
	0x16, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x49, 0x64, 0x22, 0x8d, 0x04, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x73, 0x73, 0x75, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0xae, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x73, 0x73, 0x75, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x64,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x74, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0xba, 0x02,
	0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x47, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53,
	0x65, 0x6e, 0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x1a, 0xaf, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x74,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x73,
	0x67, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd6, 0x01,
	0x0a, 0x09, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x48, 0x0a, 0x0d,
	0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x74,
	0x69, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x09,
	0x4d, 0x73, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x12, 0x4b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x73, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x1a, 0x48,
	0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x0a,
	0x13, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x62, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a,
	0x1e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xe7, 0x06, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x69, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x1a, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x69, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x30, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x1a, 0x29, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x23, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x74, 0x69,
	0x72, 0x65, 0x1a, 0x2b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5a, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x1a, 0x2b,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x2d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x35,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x1a, 0x37, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x38, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xff, 0x01, 0x0a, 0x1c, 0x63, 0x6f,
	0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x07, 0x54, 0x78, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x45,
	0x58, 0xaa, 0x02, 0x18, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x18, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x24, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c,
	0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1a, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_v1alpha1_tx_proto_rawDescOnce sync.Once
	file_regen_ecocredit_v1alpha1_tx_proto_rawDescData = file_regen_ecocredit_v1alpha1_tx_proto_rawDesc
)

func file_regen_ecocredit_v1alpha1_tx_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_v1alpha1_tx_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_v1alpha1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_v1alpha1_tx_proto_rawDescData)
	})
	return file_regen_ecocredit_v1alpha1_tx_proto_rawDescData
}

var file_regen_ecocredit_v1alpha1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_regen_ecocredit_v1alpha1_tx_proto_goTypes = []interface{}{
	(*MsgCreateClass)(nil),                 // 0: regen.ecocredit.v1alpha1.MsgCreateClass
	(*MsgCreateClassResponse)(nil),         // 1: regen.ecocredit.v1alpha1.MsgCreateClassResponse
	(*MsgCreateBatch)(nil),                 // 2: regen.ecocredit.v1alpha1.MsgCreateBatch
	(*MsgCreateBatchResponse)(nil),         // 3: regen.ecocredit.v1alpha1.MsgCreateBatchResponse
	(*MsgSend)(nil),                        // 4: regen.ecocredit.v1alpha1.MsgSend
	(*MsgSendResponse)(nil),                // 5: regen.ecocredit.v1alpha1.MsgSendResponse
	(*MsgRetire)(nil),                      // 6: regen.ecocredit.v1alpha1.MsgRetire
	(*MsgRetireResponse)(nil),              // 7: regen.ecocredit.v1alpha1.MsgRetireResponse
	(*MsgCancel)(nil),                      // 8: regen.ecocredit.v1alpha1.MsgCancel
	(*MsgCancelResponse)(nil),              // 9: regen.ecocredit.v1alpha1.MsgCancelResponse
	(*MsgUpdateClassAdmin)(nil),            // 10: regen.ecocredit.v1alpha1.MsgUpdateClassAdmin
	(*MsgUpdateClassAdminResponse)(nil),    // 11: regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse
	(*MsgUpdateClassIssuers)(nil),          // 12: regen.ecocredit.v1alpha1.MsgUpdateClassIssuers
	(*MsgUpdateClassIssuersResponse)(nil),  // 13: regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse
	(*MsgUpdateClassMetadata)(nil),         // 14: regen.ecocredit.v1alpha1.MsgUpdateClassMetadata
	(*MsgUpdateClassMetadataResponse)(nil), // 15: regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse
	(*MsgCreateBatch_BatchIssuance)(nil),   // 16: regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance
	(*MsgSend_SendCredits)(nil),            // 17: regen.ecocredit.v1alpha1.MsgSend.SendCredits
	(*MsgRetire_RetireCredits)(nil),        // 18: regen.ecocredit.v1alpha1.MsgRetire.RetireCredits
	(*MsgCancel_CancelCredits)(nil),        // 19: regen.ecocredit.v1alpha1.MsgCancel.CancelCredits
	(*timestamppb.Timestamp)(nil),          // 20: google.protobuf.Timestamp
}
var file_regen_ecocredit_v1alpha1_tx_proto_depIdxs = []int32{
	16, // 0: regen.ecocredit.v1alpha1.MsgCreateBatch.issuance:type_name -> regen.ecocredit.v1alpha1.MsgCreateBatch.BatchIssuance
	20, // 1: regen.ecocredit.v1alpha1.MsgCreateBatch.start_date:type_name -> google.protobuf.Timestamp
	20, // 2: regen.ecocredit.v1alpha1.MsgCreateBatch.end_date:type_name -> google.protobuf.Timestamp
	17, // 3: regen.ecocredit.v1alpha1.MsgSend.credits:type_name -> regen.ecocredit.v1alpha1.MsgSend.SendCredits
	18, // 4: regen.ecocredit.v1alpha1.MsgRetire.credits:type_name -> regen.ecocredit.v1alpha1.MsgRetire.RetireCredits
	19, // 5: regen.ecocredit.v1alpha1.MsgCancel.credits:type_name -> regen.ecocredit.v1alpha1.MsgCancel.CancelCredits
	0,  // 6: regen.ecocredit.v1alpha1.Msg.CreateClass:input_type -> regen.ecocredit.v1alpha1.MsgCreateClass
	2,  // 7: regen.ecocredit.v1alpha1.Msg.CreateBatch:input_type -> regen.ecocredit.v1alpha1.MsgCreateBatch
	4,  // 8: regen.ecocredit.v1alpha1.Msg.Send:input_type -> regen.ecocredit.v1alpha1.MsgSend
	6,  // 9: regen.ecocredit.v1alpha1.Msg.Retire:input_type -> regen.ecocredit.v1alpha1.MsgRetire
	8,  // 10: regen.ecocredit.v1alpha1.Msg.Cancel:input_type -> regen.ecocredit.v1alpha1.MsgCancel
	10, // 11: regen.ecocredit.v1alpha1.Msg.UpdateClassAdmin:input_type -> regen.ecocredit.v1alpha1.MsgUpdateClassAdmin
	12, // 12: regen.ecocredit.v1alpha1.Msg.UpdateClassIssuers:input_type -> regen.ecocredit.v1alpha1.MsgUpdateClassIssuers
	14, // 13: regen.ecocredit.v1alpha1.Msg.UpdateClassMetadata:input_type -> regen.ecocredit.v1alpha1.MsgUpdateClassMetadata
	1,  // 14: regen.ecocredit.v1alpha1.Msg.CreateClass:output_type -> regen.ecocredit.v1alpha1.MsgCreateClassResponse
	3,  // 15: regen.ecocredit.v1alpha1.Msg.CreateBatch:output_type -> regen.ecocredit.v1alpha1.MsgCreateBatchResponse
	5,  // 16: regen.ecocredit.v1alpha1.Msg.Send:output_type -> regen.ecocredit.v1alpha1.MsgSendResponse
	7,  // 17: regen.ecocredit.v1alpha1.Msg.Retire:output_type -> regen.ecocredit.v1alpha1.MsgRetireResponse
	9,  // 18: regen.ecocredit.v1alpha1.Msg.Cancel:output_type -> regen.ecocredit.v1alpha1.MsgCancelResponse
	11, // 19: regen.ecocredit.v1alpha1.Msg.UpdateClassAdmin:output_type -> regen.ecocredit.v1alpha1.MsgUpdateClassAdminResponse
	13, // 20: regen.ecocredit.v1alpha1.Msg.UpdateClassIssuers:output_type -> regen.ecocredit.v1alpha1.MsgUpdateClassIssuersResponse
	15, // 21: regen.ecocredit.v1alpha1.Msg.UpdateClassMetadata:output_type -> regen.ecocredit.v1alpha1.MsgUpdateClassMetadataResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_v1alpha1_tx_proto_init() }
func file_regen_ecocredit_v1alpha1_tx_proto_init() {
	if File_regen_ecocredit_v1alpha1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateClass); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateClassResponse); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBatch); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBatchResponse); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSend); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSendResponse); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRetire); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRetireResponse); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCancel); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCancelResponse); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassAdmin); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassAdminResponse); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassIssuers); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassIssuersResponse); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassMetadata); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateClassMetadataResponse); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBatch_BatchIssuance); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSend_SendCredits); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRetire_RetireCredits); i {
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
		file_regen_ecocredit_v1alpha1_tx_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCancel_CancelCredits); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_v1alpha1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_regen_ecocredit_v1alpha1_tx_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_v1alpha1_tx_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_v1alpha1_tx_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_v1alpha1_tx_proto = out.File
	file_regen_ecocredit_v1alpha1_tx_proto_rawDesc = nil
	file_regen_ecocredit_v1alpha1_tx_proto_goTypes = nil
	file_regen_ecocredit_v1alpha1_tx_proto_depIdxs = nil
}
