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

package pdfaid ;import (_ae "fmt";_e "github.com/trimmer-io/go-xmp/xmp";_b "github.com/unidoc/unipdf/v3/model/xmputil/pdfaextension";);

// Namespaces implements xmp.Model interface.
func (_gd *Model )Namespaces ()_e .NamespaceList {return _e .NamespaceList {Namespace }};func init (){_e .Register (Namespace ,_e .XmpMetadata );_b .RegisterSchema (Namespace ,&Schema )};

// SyncToXMP implements xmp.Model interface.
func (_ef *Model )SyncToXMP (d *_e .Document )error {return nil };

// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel (d *_e .Document )(*Model ,error ){_fa ,_bd :=d .MakeModel (Namespace );if _bd !=nil {return nil ,_bd ;};return _fa .(*Model ),nil ;};var Namespace =_e .NewNamespace ("\u0070\u0064\u0066\u0061\u0069\u0064","\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f",NewModel );


// SyncFromXMP implements xmp.Model interface.
func (_d *Model )SyncFromXMP (d *_e .Document )error {return nil };

// SetTag implements xmp.Model interface.
func (_bfg *Model )SetTag (tag ,value string )error {if _cb :=_e .SetNativeField (_bfg ,tag ,value );_cb !=nil {return _ae .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_cb );};return nil ;};

// NewModel creates a new model.
func NewModel (name string )_e .Model {return &Model {}};

// SyncModel implements xmp.Model interface.
func (_bc *Model )SyncModel (d *_e .Document )error {return nil };var _ _e .Model =(*Model )(nil );

// GetTag implements xmp.Model interface.
func (_ffd *Model )GetTag (tag string )(string ,error ){_gg ,_c :=_e .GetNativeField (_ffd ,tag );if _c !=nil {return "",_ae .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_c );};return _gg ,nil ;};

// Can implements xmp.Model interface.
func (_g *Model )Can (nsName string )bool {return Namespace .GetName ()==nsName };

// CanTag implements xmp.Model interface.
func (_eb *Model )CanTag (tag string )bool {_ ,_ff :=_e .GetNativeField (_eb ,tag );return _ff ==nil };

// Model is the XMP model for the PdfA metadata.
type Model struct{Part int `xmp:"pdfaid:part"`;Conformance string `xmp:"pdfaid:conformance"`;};var Schema =_b .Schema {NamespaceURI :Namespace .URI ,Prefix :Namespace .Name ,Schema :"\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061",Property :[]_b .Property {{Category :_b .PropertyCategoryInternal ,Description :"\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064",Name :"\u0070\u0061\u0072\u0074",ValueType :_b .ValueTypeNameInteger },{Category :_b .PropertyCategoryInternal ,Description :"A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064",Name :"\u0061\u006d\u0064",ValueType :_b .ValueTypeNameText },{Category :_b .PropertyCategoryInternal ,Description :"C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064",Name :"c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065",ValueType :_b .ValueTypeNameText }},ValueType :nil };
