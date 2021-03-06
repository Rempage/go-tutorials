/* // +build ignore */

// Copyright (c) 2012-2015 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

// Code generated from gen-helper.go.tmpl - DO NOT EDIT.

package codec

import (
	"encoding"
	"reflect"
)

// GenVersion is the current version of codecgen.
const GenVersion = 8

// This file is used to generate helper code for codecgen.
// The values here i.e. genHelper(En|De)coder are not to be used directly by
// library users. They WILL change continuously and without notice.
//
// To help enforce this, we create an unexported type with exported members.
// The only way to get the type is via the one exported type that we control (somewhat).
//
// When static codecs are created for types, they will use this value
// to perform encoding or decoding of primitives or known slice or map types.

// GenHelperEncoder is exported so that it can be used externally by codecgen.
//
// Library users: DO NOT USE IT DIRECTLY. IT WILL CHANGE CONTINOUSLY WITHOUT NOTICE.
func GenHelperEncoder(e *Encoder) (ge genHelperEncoder, ee encDriver) {
	ge = genHelperEncoder{e: e}
	ee = e.e
	return
}

// GenHelperDecoder is exported so that it can be used externally by codecgen.
//
// Library users: DO NOT USE IT DIRECTLY. IT WILL CHANGE CONTINOUSLY WITHOUT NOTICE.
func GenHelperDecoder(d *Decoder) (gd genHelperDecoder, dd decDriver) {
	gd = genHelperDecoder{d: d}
	dd = d.d
	return
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
type genHelperEncoder struct {
	e *Encoder
	F fastpathT
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
type genHelperDecoder struct {
	d *Decoder
	F fastpathT
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncBasicHandle() *BasicHandle {
	return f.e.h
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncBinary() bool {
	return f.e.cf.be // f.e.hh.isBinaryEncoding()
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncFallback(iv interface{}) {
	// println(">>>>>>>>> EncFallback")
	// f.e.encodeI(iv, false, false)
	f.e.encodeValue(reflect.ValueOf(iv), nil, false)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncTextMarshal(iv encoding.TextMarshaler) {
	bs, fnerr := iv.MarshalText()
	f.e.marshal(bs, fnerr, false, cUTF8)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncJSONMarshal(iv jsonMarshaler) {
	bs, fnerr := iv.MarshalJSON()
	f.e.marshal(bs, fnerr, true, cUTF8)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncBinaryMarshal(iv encoding.BinaryMarshaler) {
	bs, fnerr := iv.MarshalBinary()
	f.e.marshal(bs, fnerr, false, cRAW)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncRaw(iv Raw) { f.e.rawBytes(iv) }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: builtin no longer supported - so we make this method a no-op,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperEncoder) TimeRtidIfBinc() (v uintptr) { return }

// func (f genHelperEncoder) TimeRtidIfBinc() uintptr {
// 	if _, ok := f.e.hh.(*BincHandle); ok {
// 		return timeTypId
// 	}
// }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) IsJSONHandle() bool {
	return f.e.cf.js
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) I2Rtid(v interface{}) uintptr {
	return i2rtid(v)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) Extension(rtid uintptr) (xfn *extTypeTagFn) {
	return f.e.h.getExt(rtid)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperEncoder) EncExtension(v interface{}, xfFn *extTypeTagFn) {
	f.e.e.EncodeExt(v, xfFn.tag, xfFn.ext, f.e)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: No longer used,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperEncoder) HasExtensions() bool {
	return len(f.e.h.extHandle) != 0
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: No longer used,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperEncoder) EncExt(v interface{}) (r bool) {
	if xfFn := f.e.h.getExt(i2rtid(v)); xfFn != nil {
		f.e.e.EncodeExt(v, xfFn.tag, xfFn.ext, f.e)
		return true
	}
	return false
}

// ---------------- DECODER FOLLOWS -----------------

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecBasicHandle() *BasicHandle {
	return f.d.h
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecBinary() bool {
	return f.d.be // f.d.hh.isBinaryEncoding()
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecSwallow() { f.d.swallow() }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecScratchBuffer() []byte {
	return f.d.b[:]
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecFallback(iv interface{}, chkPtr bool) {
	// println(">>>>>>>>> DecFallback")
	rv := reflect.ValueOf(iv)
	if chkPtr {
		rv = f.d.ensureDecodeable(rv)
	}
	f.d.decodeValue(rv, nil, false)
	// f.d.decodeValueFallback(rv)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecSliceHelperStart() (decSliceHelper, int) {
	return f.d.decSliceHelperStart()
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecStructFieldNotFound(index int, name string) {
	f.d.structFieldNotFound(index, name)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecArrayCannotExpand(sliceLen, streamLen int) {
	f.d.arrayCannotExpand(sliceLen, streamLen)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecTextUnmarshal(tm encoding.TextUnmarshaler) {
	fnerr := tm.UnmarshalText(f.d.d.DecodeStringAsBytes())
	if fnerr != nil {
		panic(fnerr)
	}
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecJSONUnmarshal(tm jsonUnmarshaler) {
	// bs := f.dd.DecodeStringAsBytes()
	// grab the bytes to be read, as UnmarshalJSON needs the full JSON so as to unmarshal it itself.
	fnerr := tm.UnmarshalJSON(f.d.nextValueBytes())
	if fnerr != nil {
		panic(fnerr)
	}
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecBinaryUnmarshal(bm encoding.BinaryUnmarshaler) {
	fnerr := bm.UnmarshalBinary(f.d.d.DecodeBytes(nil, true))
	if fnerr != nil {
		panic(fnerr)
	}
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecRaw() []byte { return f.d.rawBytes() }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: builtin no longer supported - so we make this method a no-op,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperDecoder) TimeRtidIfBinc() (v uintptr) { return }

// func (f genHelperDecoder) TimeRtidIfBinc() uintptr {
// 	// Note: builtin is no longer supported - so make this a no-op
// 	if _, ok := f.d.hh.(*BincHandle); ok {
// 		return timeTypId
// 	}
// 	return 0
// }

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) IsJSONHandle() bool {
	return f.d.js
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) I2Rtid(v interface{}) uintptr {
	return i2rtid(v)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) Extension(rtid uintptr) (xfn *extTypeTagFn) {
	return f.d.h.getExt(rtid)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecExtension(v interface{}, xfFn *extTypeTagFn) {
	f.d.d.DecodeExt(v, xfFn.tag, xfFn.ext)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: No longer used,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperDecoder) HasExtensions() bool {
	return len(f.d.h.extHandle) != 0
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: No longer used,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperDecoder) DecExt(v interface{}) (r bool) {
	if xfFn := f.d.h.getExt(i2rtid(v)); xfFn != nil {
		f.d.d.DecodeExt(v, xfFn.tag, xfFn.ext)
		return true
	}
	return false
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
func (f genHelperDecoder) DecInferLen(clen, maxlen, unit int) (rvlen int) {
	return decInferLen(clen, maxlen, unit)
}

// FOR USE BY CODECGEN ONLY. IT *WILL* CHANGE WITHOUT NOTICE. *DO NOT USE*
//
// Deprecated: no longer used,
// but leave in-place so that old generated files continue to work without regeneration.
func (f genHelperDecoder) StringView(v []byte) string { return stringView(v) }
