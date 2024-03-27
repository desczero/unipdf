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

// Package fjson provides support for loading PDF form field data from JSON data/files.
package fjson ;import (_a "encoding/json";_b "github.com/unidoc/unipdf/v3/common";_gf "github.com/unidoc/unipdf/v3/core";_bf "github.com/unidoc/unipdf/v3/model";_g "io";_f "os";);

// FieldImageValues implements model.FieldImageProvider interface.
func (_cb *FieldData )FieldImageValues ()(map[string ]*_bf .Image ,error ){_fe :=make (map[string ]*_bf .Image );for _ ,_afd :=range _cb ._ae {if _afd .ImageValue !=nil {_fe [_afd .Name ]=_afd .ImageValue ;};};return _fe ,nil ;};

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _g .ReadSeeker )(*FieldData ,error ){_ad ,_be :=_bf .NewPdfReader (rs );if _be !=nil {return nil ,_be ;};if _ad .AcroForm ==nil {return nil ,nil ;};var _e []fieldValue ;_gg :=_ad .AcroForm .AllFields ();for _ ,_gd :=range _gg {var _adb []string ;
_aec :=make (map[string ]struct{});_afe ,_bff :=_gd .FullName ();if _bff !=nil {return nil ,_bff ;};if _ge ,_ff :=_gd .V .(*_gf .PdfObjectString );_ff {_e =append (_e ,fieldValue {Name :_afe ,Value :_ge .Decoded ()});continue ;};var _de string ;for _ ,_dd :=range _gd .Annotations {_aeg ,_da :=_gf .GetName (_dd .AS );
if _da {_de =_aeg .String ();};_c ,_dae :=_gf .GetDict (_dd .AP );if !_dae {continue ;};_ag ,_ :=_gf .GetDict (_c .Get ("\u004e"));for _ ,_ga :=range _ag .Keys (){_ec :=_ga .String ();if _ ,_gag :=_aec [_ec ];!_gag {_adb =append (_adb ,_ec );_aec [_ec ]=struct{}{};
};};_ef ,_ :=_gf .GetDict (_c .Get ("\u0044"));for _ ,_eb :=range _ef .Keys (){_eg :=_eb .String ();if _ ,_ggc :=_aec [_eg ];!_ggc {_adb =append (_adb ,_eg );_aec [_eg ]=struct{}{};};};};_aa :=fieldValue {Name :_afe ,Value :_de ,Options :_adb };_e =append (_e ,_aa );
};_bb :=FieldData {_ae :_e };return &_bb ,nil ;};

// JSON returns the field data as a string in JSON format.
func (_fc FieldData )JSON ()(string ,error ){_aed ,_ebb :=_a .MarshalIndent (_fc ._ae ,"","\u0020\u0020\u0020\u0020");return string (_aed ),_ebb ;};

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_fg ,_dab :=_f .Open (filePath );if _dab !=nil {return nil ,_dab ;};defer _fg .Close ();return LoadFromPDF (_fg );};

// SetImage assign model.Image to a specific field identified by fieldName.
func (_ddg *FieldData )SetImage (fieldName string ,img *_bf .Image ,opt []string )error {_ggf :=fieldValue {Name :fieldName ,ImageValue :img ,Options :opt };_ddg ._ae =append (_ddg ._ae ,_ggf );return nil ;};

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_dcf *FieldData )SetImageFromFile (fieldName string ,imagePath string ,opt []string )error {_ffg ,_gae :=_f .Open (imagePath );if _gae !=nil {return _gae ;};defer _ffg .Close ();_cbc ,_gae :=_bf .ImageHandling .Read (_ffg );if _gae !=nil {_b .Log .Error ("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_gae );
return _gae ;};return _dcf .SetImage (fieldName ,_cbc ,opt );};

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_ae []fieldValue };

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_dc ,_af :=_f .Open (filePath );if _af !=nil {return nil ,_af ;};defer _dc .Close ();return LoadFromJSON (_dc );};

// FieldValues implements model.FieldValueProvider interface.
func (_agf *FieldData )FieldValues ()(map[string ]_gf .PdfObject ,error ){_egg :=make (map[string ]_gf .PdfObject );for _ ,_gcc :=range _agf ._ae {if len (_gcc .Value )> 0{_egg [_gcc .Name ]=_gf .MakeString (_gcc .Value );};};return _egg ,nil ;};type fieldValue struct{Name string `json:"name"`;
Value string `json:"value"`;ImageValue *_bf .Image `json:"-"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _g .Reader )(*FieldData ,error ){var _gc FieldData ;_bfe :=_a .NewDecoder (r ).Decode (&_gc ._ae );if _bfe !=nil {return nil ,_bfe ;};return &_gc ,nil ;};