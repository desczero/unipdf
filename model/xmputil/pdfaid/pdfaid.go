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

package pdfaid ;import (_a "fmt";_ad "github.com/trimmer-io/go-xmp/xmp";_eb "github.com/unidoc/unipdf/v3/model/xmputil/pdfaextension";);

// Can implements xmp.Model interface.
func (_f *Model )Can (nsName string )bool {return Namespace .GetName ()==nsName };func init (){_ad .Register (Namespace ,_ad .XmpMetadata );_eb .RegisterSchema (Namespace ,&Schema )};

// SetTag implements xmp.Model interface.
func (_ecd *Model )SetTag (tag ,value string )error {if _aee :=_ad .SetNativeField (_ecd ,tag ,value );_aee !=nil {return _a .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_aee );};return nil ;};

// SyncToXMP implements xmp.Model interface.
func (_ab *Model )SyncToXMP (d *_ad .Document )error {return nil };var Schema =_eb .Schema {NamespaceURI :Namespace .URI ,Prefix :Namespace .Name ,Schema :"\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061",Property :[]_eb .Property {{Category :_eb .PropertyCategoryInternal ,Description :"\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064",Name :"\u0070\u0061\u0072\u0074",ValueType :_eb .ValueTypeNameInteger },{Category :_eb .PropertyCategoryInternal ,Description :"A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064",Name :"\u0061\u006d\u0064",ValueType :_eb .ValueTypeNameText },{Category :_eb .PropertyCategoryInternal ,Description :"C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064",Name :"c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065",ValueType :_eb .ValueTypeNameText }},ValueType :nil };


// SyncFromXMP implements xmp.Model interface.
func (_b *Model )SyncFromXMP (d *_ad .Document )error {return nil };

// NewModel creates a new model.
func NewModel (name string )_ad .Model {return &Model {}};

// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel (d *_ad .Document )(*Model ,error ){_cb ,_ce :=d .MakeModel (Namespace );if _ce !=nil {return nil ,_ce ;};return _cb .(*Model ),nil ;};var _ _ad .Model =(*Model )(nil );var Namespace =_ad .NewNamespace ("\u0070\u0064\u0066\u0061\u0069\u0064","\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f",NewModel );


// SyncModel implements xmp.Model interface.
func (_ec *Model )SyncModel (d *_ad .Document )error {return nil };

// Namespaces implements xmp.Model interface.
func (_d *Model )Namespaces ()_ad .NamespaceList {return _ad .NamespaceList {Namespace }};

// Model is the XMP model for the PdfA metadata.
type Model struct{Part int `xmp:"pdfaid:part"`;Conformance string `xmp:"pdfaid:conformance"`;};

// CanTag implements xmp.Model interface.
func (_ceb *Model )CanTag (tag string )bool {_ ,_abd :=_ad .GetNativeField (_ceb ,tag );return _abd ==nil ;};

// GetTag implements xmp.Model interface.
func (_ac *Model )GetTag (tag string )(string ,error ){_ff ,_ea :=_ad .GetNativeField (_ac ,tag );if _ea !=nil {return "",_a .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_ea );};return _ff ,nil ;};