// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package sec

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

func easyjsonC80ae7adDecodeGithubComPixelguy95Sec(in *jlexer.Lexer, out *Unit) {
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
		case "val":
			out.Value = float64(in.Float64())
		case "start":
			out.Start = string(in.String())
		case "end":
			out.End = string(in.String())
		case "fy":
			out.FiscalYear = uint16(in.Uint16())
		case "fp":
			out.FiscalPeriod = string(in.String())
		case "accn":
			out.Account = string(in.String())
		case "form":
			out.Form = string(in.String())
		case "filed":
			out.FiledOn = string(in.String())
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
func easyjsonC80ae7adEncodeGithubComPixelguy95Sec(out *jwriter.Writer, in Unit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"val\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Value))
	}
	{
		const prefix string = ",\"start\":"
		out.RawString(prefix)
		out.String(string(in.Start))
	}
	{
		const prefix string = ",\"end\":"
		out.RawString(prefix)
		out.String(string(in.End))
	}
	{
		const prefix string = ",\"fy\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.FiscalYear))
	}
	{
		const prefix string = ",\"fp\":"
		out.RawString(prefix)
		out.String(string(in.FiscalPeriod))
	}
	{
		const prefix string = ",\"accn\":"
		out.RawString(prefix)
		out.String(string(in.Account))
	}
	{
		const prefix string = ",\"form\":"
		out.RawString(prefix)
		out.String(string(in.Form))
	}
	{
		const prefix string = ",\"filed\":"
		out.RawString(prefix)
		out.String(string(in.FiledOn))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Unit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComPixelguy95Sec(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Unit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComPixelguy95Sec(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Unit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComPixelguy95Sec(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Unit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComPixelguy95Sec(l, v)
}
func easyjsonC80ae7adDecodeGithubComPixelguy95Sec1(in *jlexer.Lexer, out *Facts) {
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
		case "dei":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.DEI = make(map[string]Fact)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 Fact
					(v1).UnmarshalEasyJSON(in)
					(out.DEI)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "us-gaap":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.UsGAAP = make(map[string]Fact)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 Fact
					(v2).UnmarshalEasyJSON(in)
					(out.UsGAAP)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonC80ae7adEncodeGithubComPixelguy95Sec1(out *jwriter.Writer, in Facts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dei\":"
		out.RawString(prefix[1:])
		if in.DEI == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.DEI {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				(v3Value).MarshalEasyJSON(out)
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"us-gaap\":"
		out.RawString(prefix)
		if in.UsGAAP == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.UsGAAP {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				(v4Value).MarshalEasyJSON(out)
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Facts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComPixelguy95Sec1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Facts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComPixelguy95Sec1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Facts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComPixelguy95Sec1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Facts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComPixelguy95Sec1(l, v)
}
func easyjsonC80ae7adDecodeGithubComPixelguy95Sec2(in *jlexer.Lexer, out *Fact) {
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
		case "label":
			out.Label = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "units":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Units = make(map[string][]Unit)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 []Unit
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						in.Delim('[')
						if v5 == nil {
							if !in.IsDelim(']') {
								v5 = make([]Unit, 0, 0)
							} else {
								v5 = []Unit{}
							}
						} else {
							v5 = (v5)[:0]
						}
						for !in.IsDelim(']') {
							var v6 Unit
							(v6).UnmarshalEasyJSON(in)
							v5 = append(v5, v6)
							in.WantComma()
						}
						in.Delim(']')
					}
					(out.Units)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonC80ae7adEncodeGithubComPixelguy95Sec2(out *jwriter.Writer, in Fact) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"label\":"
		out.RawString(prefix[1:])
		out.String(string(in.Label))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"units\":"
		out.RawString(prefix)
		if in.Units == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v7First := true
			for v7Name, v7Value := range in.Units {
				if v7First {
					v7First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v7Name))
				out.RawByte(':')
				if v7Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v8, v9 := range v7Value {
						if v8 > 0 {
							out.RawByte(',')
						}
						(v9).MarshalEasyJSON(out)
					}
					out.RawByte(']')
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Fact) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComPixelguy95Sec2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Fact) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComPixelguy95Sec2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Fact) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComPixelguy95Sec2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Fact) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComPixelguy95Sec2(l, v)
}
func easyjsonC80ae7adDecodeGithubComPixelguy95Sec3(in *jlexer.Lexer, out *CompanyFacts) {
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
		case "cik":
			out.CIK = uint64(in.Uint64())
		case "entityName":
			out.EntityName = string(in.String())
		case "facts":
			(out.Facts).UnmarshalEasyJSON(in)
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
func easyjsonC80ae7adEncodeGithubComPixelguy95Sec3(out *jwriter.Writer, in CompanyFacts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cik\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.CIK))
	}
	{
		const prefix string = ",\"entityName\":"
		out.RawString(prefix)
		out.String(string(in.EntityName))
	}
	{
		const prefix string = ",\"facts\":"
		out.RawString(prefix)
		(in.Facts).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CompanyFacts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComPixelguy95Sec3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CompanyFacts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComPixelguy95Sec3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CompanyFacts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComPixelguy95Sec3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CompanyFacts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComPixelguy95Sec3(l, v)
}
