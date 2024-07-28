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
package fjson ;import (_b "encoding/json";_g "github.com/unidoc/unipdf/v3/common";_bd "github.com/unidoc/unipdf/v3/core";_e "github.com/unidoc/unipdf/v3/model";_cb "io";_a "os";);

// FieldValues implements model.FieldValueProvider interface.
func (_gb *FieldData )FieldValues ()(map[string ]_bd .PdfObject ,error ){_dge :=make (map[string ]_bd .PdfObject );for _ ,_eaa :=range _gb ._gd {if len (_eaa .Value )> 0{_dge [_eaa .Name ]=_bd .MakeString (_eaa .Value );};};return _dge ,nil ;};

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _cb .Reader )(*FieldData ,error ){var _bb FieldData ;_ce :=_b .NewDecoder (r ).Decode (&_bb ._gd );if _ce !=nil {return nil ,_ce ;};return &_bb ,nil ;};

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_gd []fieldValue };

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_fbb *FieldData )SetImageFromFile (fieldName string ,imagePath string ,opt []string )error {_fg ,_bbc :=_a .Open (imagePath );if _bbc !=nil {return _bbc ;};defer _fg .Close ();_ccdg ,_bbc :=_e .ImageHandling .Read (_fg );if _bbc !=nil {_g .Log .Error ("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_bbc );
return _bbc ;};return _fbb .SetImage (fieldName ,_ccdg ,opt );};

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_gg ,_ddf :=_a .Open (filePath );if _ddf !=nil {return nil ,_ddf ;};defer _gg .Close ();return LoadFromPDF (_gg );};

// JSON returns the field data as a string in JSON format.
func (_bae FieldData )JSON ()(string ,error ){_dda ,_ddg :=_b .MarshalIndent (_bae ._gd ,"","\u0020\u0020\u0020\u0020");return string (_dda ),_ddg ;};

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_ee ,_cc :=_a .Open (filePath );if _cc !=nil {return nil ,_cc ;};defer _ee .Close ();return LoadFromJSON (_ee );};

// SetImage assign model.Image to a specific field identified by fieldName.
func (_fba *FieldData )SetImage (fieldName string ,img *_e .Image ,opt []string )error {_fbae :=fieldValue {Name :fieldName ,ImageValue :img ,Options :opt };_fba ._gd =append (_fba ._gd ,_fbae );return nil ;};type fieldValue struct{Name string `json:"name"`;
Value string `json:"value"`;ImageValue *_e .Image `json:"-"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _cb .ReadSeeker )(*FieldData ,error ){_ec ,_f :=_e .NewPdfReader (rs );if _f !=nil {return nil ,_f ;};if _ec .AcroForm ==nil {return nil ,nil ;};var _ecf []fieldValue ;_bdg :=_ec .AcroForm .AllFields ();for _ ,_gf :=range _bdg {var _be []string ;
_dg :=make (map[string ]struct{});_fb ,_gc :=_gf .FullName ();if _gc !=nil {return nil ,_gc ;};if _fe ,_bg :=_gf .V .(*_bd .PdfObjectString );_bg {_ecf =append (_ecf ,fieldValue {Name :_fb ,Value :_fe .Decoded ()});continue ;};var _af string ;for _ ,_ag :=range _gf .Annotations {_ge ,_ca :=_bd .GetName (_ag .AS );
if _ca {_af =_ge .String ();};_cbe ,_df :=_bd .GetDict (_ag .AP );if !_df {continue ;};_afb ,_ :=_bd .GetDict (_cbe .Get ("\u004e"));for _ ,_bee :=range _afb .Keys (){_eca :=_bee .String ();if _ ,_geg :=_dg [_eca ];!_geg {_be =append (_be ,_eca );_dg [_eca ]=struct{}{};
};};_dc ,_ :=_bd .GetDict (_cbe .Get ("\u0044"));for _ ,_cag :=range _dc .Keys (){_bde :=_cag .String ();if _ ,_ad :=_dg [_bde ];!_ad {_be =append (_be ,_bde );_dg [_bde ]=struct{}{};};};};_gcb :=fieldValue {Name :_fb ,Value :_af ,Options :_be };_ecf =append (_ecf ,_gcb );
};_dd :=FieldData {_gd :_ecf };return &_dd ,nil ;};

// FieldImageValues implements model.FieldImageProvider interface.
func (_cage *FieldData )FieldImageValues ()(map[string ]*_e .Image ,error ){_fd :=make (map[string ]*_e .Image );for _ ,_gbg :=range _cage ._gd {if _gbg .ImageValue !=nil {_fd [_gbg .Name ]=_gbg .ImageValue ;};};return _fd ,nil ;};