// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package parser

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

func easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser(in *jlexer.Lexer, out *ObjectResults) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "results":
			if in.IsNull() {
				in.Skip()
				out.Results = nil
			} else {
				if out.Results == nil {
					out.Results = new(map[string]int)
				}
				if in.IsNull() {
					in.Skip()
				} else {
					in.Delim('{')
					if !in.IsDelim('}') {
						*out.Results = make(map[string]int)
					} else {
						*out.Results = nil
					}
					for !in.IsDelim('}') {
						key := string(in.String())
						in.WantColon()
						var v1 int
						v1 = int(in.Int())
						(*out.Results)[key] = v1
						in.WantComma()
					}
					in.Delim('}')
				}
			}
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
func easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser(out *jwriter.Writer, in ObjectResults) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"results\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Results == nil {
			out.RawString("null")
		} else {
			if *in.Results == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
				out.RawString(`null`)
			} else {
				out.RawByte('{')
				v2First := true
				for v2Name, v2Value := range *in.Results {
					if v2First {
						v2First = false
					} else {
						out.RawByte(',')
					}
					out.String(string(v2Name))
					out.RawByte(':')
					out.Int(int(v2Value))
				}
				out.RawByte('}')
			}
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ObjectResults) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ObjectResults) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ObjectResults) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ObjectResults) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser(l, v)
}
func easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser1(in *jlexer.Lexer, out *IntResults) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "results":
			out.Results = int(in.Int())
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
func easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser1(out *jwriter.Writer, in IntResults) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"results\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Results))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IntResults) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IntResults) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IntResults) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IntResults) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser1(l, v)
}
func easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser2(in *jlexer.Lexer, out *ArrayResults) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "results":
			if in.IsNull() {
				in.Skip()
				out.Results = nil
			} else {
				if out.Results == nil {
					out.Results = new([]string)
				}
				if in.IsNull() {
					in.Skip()
					*out.Results = nil
				} else {
					in.Delim('[')
					if *out.Results == nil {
						if !in.IsDelim(']') {
							*out.Results = make([]string, 0, 4)
						} else {
							*out.Results = []string{}
						}
					} else {
						*out.Results = (*out.Results)[:0]
					}
					for !in.IsDelim(']') {
						var v3 string
						v3 = string(in.String())
						*out.Results = append(*out.Results, v3)
						in.WantComma()
					}
					in.Delim(']')
				}
			}
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
func easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser2(out *jwriter.Writer, in ArrayResults) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"results\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Results == nil {
			out.RawString("null")
		} else {
			if *in.Results == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
				out.RawString("null")
			} else {
				out.RawByte('[')
				for v4, v5 := range *in.Results {
					if v4 > 0 {
						out.RawByte(',')
					}
					out.String(string(v5))
				}
				out.RawByte(']')
			}
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ArrayResults) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ArrayResults) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF59a38b1EncodeGithubCombusbudTidalwaveParser2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ArrayResults) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ArrayResults) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF59a38b1DecodeGithubCombusbudTidalwaveParser2(l, v)
}
