//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package pdfaid ;import (_d "fmt";_g "github.com/trimmer-io/go-xmp/xmp";_a "github.com/unidoc/unipdf/v3/model/xmputil/pdfaextension";);

// CanTag implements xmp.Model interface.
func (_fc *Model )CanTag (tag string )bool {_ ,_ed :=_g .GetNativeField (_fc ,tag );return _ed ==nil };func init (){_g .Register (Namespace ,_g .XmpMetadata );_a .RegisterSchema (Namespace ,&Schema )};

// SyncFromXMP implements xmp.Model interface.
func (_ea *Model )SyncFromXMP (d *_g .Document )error {return nil };

// Namespaces implements xmp.Model interface.
func (_b *Model )Namespaces ()_g .NamespaceList {return _g .NamespaceList {Namespace }};

// Can implements xmp.Model interface.
func (_e *Model )Can (nsName string )bool {return Namespace .GetName ()==nsName };var _ _g .Model =(*Model )(nil );var Schema =_a .Schema {NamespaceURI :Namespace .URI ,Prefix :Namespace .Name ,Schema :"\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061",Property :[]_a .Property {{Category :_a .PropertyCategoryInternal ,Description :"\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064",Name :"\u0070\u0061\u0072\u0074",ValueType :_a .ValueTypeNameInteger },{Category :_a .PropertyCategoryInternal ,Description :"A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064",Name :"\u0061\u006d\u0064",ValueType :_a .ValueTypeNameText },{Category :_a .PropertyCategoryInternal ,Description :"C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064",Name :"c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065",ValueType :_a .ValueTypeNameText }},ValueType :nil };


// Model is the XMP model for the PdfA metadata.
type Model struct{Part int `xmp:"pdfaid:part"`;Conformance string `xmp:"pdfaid:conformance"`;};

// SetTag implements xmp.Model interface.
func (_bd *Model )SetTag (tag ,value string )error {if _dgf :=_g .SetNativeField (_bd ,tag ,value );_dgf !=nil {return _d .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_dgf );};return nil ;};

// GetTag implements xmp.Model interface.
func (_ef *Model )GetTag (tag string )(string ,error ){_dg ,_dge :=_g .GetNativeField (_ef ,tag );if _dge !=nil {return "",_d .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_dge );};return _dg ,nil ;};

// SyncModel implements xmp.Model interface.
func (_aff *Model )SyncModel (d *_g .Document )error {return nil };

// SyncToXMP implements xmp.Model interface.
func (_af *Model )SyncToXMP (d *_g .Document )error {return nil };

// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel (d *_g .Document )(*Model ,error ){_aa ,_fa :=d .MakeModel (Namespace );if _fa !=nil {return nil ,_fa ;};return _aa .(*Model ),nil ;};

// NewModel creates a new model.
func NewModel (name string )_g .Model {return &Model {}};var Namespace =_g .NewNamespace ("\u0070\u0064\u0066\u0061\u0069\u0064","\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f",NewModel );
