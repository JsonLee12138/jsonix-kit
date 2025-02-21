// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package core

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson163c17a9DecodeJsonixKitCore(in *jlexer.Lexer, out *BaseEntityWithUuid) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "createdBy":
			out.CreatedBy = string(in.String())
		case "updatedBy":
			out.UpdatedBy = string(in.String())
		case "deletedBy":
			out.DeletedBy = string(in.StringIntern())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson163c17a9EncodeJsonixKitCore(out *jwriter.Writer, in BaseEntityWithUuid) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"createdBy\":"
		out.RawString(prefix)
		out.String(string(in.CreatedBy))
	}
	{
		const prefix string = ",\"updatedBy\":"
		out.RawString(prefix)
		out.String(string(in.UpdatedBy))
	}
	{
		const prefix string = ",\"deletedBy\":"
		out.RawString(prefix)
		out.String(string(in.DeletedBy))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BaseEntityWithUuid) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeJsonixKitCore(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseEntityWithUuid) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeJsonixKitCore(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BaseEntityWithUuid) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeJsonixKitCore(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseEntityWithUuid) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeJsonixKitCore(l, v)
}
func easyjson163c17a9DecodeJsonixKitCore1(in *jlexer.Lexer, out *BaseEntityWithID) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint(in.Uint())
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "createdBy":
			out.CreatedBy = string(in.String())
		case "updatedBy":
			out.UpdatedBy = string(in.String())
		case "deletedBy":
			out.DeletedBy = string(in.StringIntern())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson163c17a9EncodeJsonixKitCore1(out *jwriter.Writer, in BaseEntityWithID) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"createdBy\":"
		out.RawString(prefix)
		out.String(string(in.CreatedBy))
	}
	{
		const prefix string = ",\"updatedBy\":"
		out.RawString(prefix)
		out.String(string(in.UpdatedBy))
	}
	{
		const prefix string = ",\"deletedBy\":"
		out.RawString(prefix)
		out.String(string(in.DeletedBy))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BaseEntityWithID) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeJsonixKitCore1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseEntityWithID) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeJsonixKitCore1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BaseEntityWithID) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeJsonixKitCore1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseEntityWithID) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeJsonixKitCore1(l, v)
}
func easyjson163c17a9DecodeJsonixKitCore2(in *jlexer.Lexer, out *BaseEntity) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "createdBy":
			out.CreatedBy = string(in.String())
		case "updatedBy":
			out.UpdatedBy = string(in.String())
		case "deletedBy":
			out.DeletedBy = string(in.StringIntern())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson163c17a9EncodeJsonixKitCore2(out *jwriter.Writer, in BaseEntity) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix[1:])
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"createdBy\":"
		out.RawString(prefix)
		out.String(string(in.CreatedBy))
	}
	{
		const prefix string = ",\"updatedBy\":"
		out.RawString(prefix)
		out.String(string(in.UpdatedBy))
	}
	{
		const prefix string = ",\"deletedBy\":"
		out.RawString(prefix)
		out.String(string(in.DeletedBy))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BaseEntity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson163c17a9EncodeJsonixKitCore2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseEntity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson163c17a9EncodeJsonixKitCore2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BaseEntity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson163c17a9DecodeJsonixKitCore2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseEntity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson163c17a9DecodeJsonixKitCore2(l, v)
}
