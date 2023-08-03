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

package pdfaextension ;import (_g "fmt";_ff "github.com/trimmer-io/go-xmp/models/dc";_ea "github.com/trimmer-io/go-xmp/models/pdf";_gag "github.com/trimmer-io/go-xmp/models/xmp_base";_cc "github.com/trimmer-io/go-xmp/models/xmp_mm";_ga "github.com/trimmer-io/go-xmp/xmp";
_fe "reflect";_f "sort";_c "strings";_e "sync";);

// MakeModel creates or gets a model from document.
func MakeModel (d *_ga .Document )(*Model ,error ){_bf ,_feg :=d .MakeModel (Namespace );if _feg !=nil {return nil ,_feg ;};return _bf .(*Model ),nil ;};func init (){_ga .Register (Namespace ,_ga .XmpMetadata );_ga .Register (SchemaNS );_ga .Register (PropertyNS );
_ga .Register (ValueTypeNS );_ga .Register (FieldNS );};

// GetTag implements xmp.Model interface.
func (_cf *Model )GetTag (tag string )(string ,error ){_dbf ,_be :=_ga .GetNativeField (_cf ,tag );if _be !=nil {return "",_g .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_be );};return _dbf ,nil ;};

// SetTag implements xmp.Model interface.
func (_af *Model )SetTag (tag ,value string )error {if _ge :=_ga .SetNativeField (_af ,tag ,value );_ge !=nil {return _g .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_ge );};return nil ;};const (PropertyCategoryUndefined PropertyCategory =iota ;
PropertyCategoryInternal ;PropertyCategoryExternal ;);

// RegisterSchema registers schema extension definition.
func RegisterSchema (ns *_ga .Namespace ,schema *Schema ){_fcf .Lock ();defer _fcf .Unlock ();_gdb [ns .URI ]=schema ;};

// BagOfValueTypeName gets the ValueTypeName of the bag of provided value type names.
func BagOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0062\u0061\u0067\u0020"+vt };

// Schema is the pdfa extension schema.
type Schema struct{

// NamespaceURI is schema namespace URI.
NamespaceURI string `xmp:"pdfaSchema:namespaceURI"`;

// Prefix is preferred schema namespace prefix.
Prefix string `xmp:"pdfaSchema:prefix"`;

// Schema is optional description.
Schema string `xmp:"pdfaSchema:schema"`;

// Property is description of schema properties.
Property Properties `xmp:"pdfaSchema:property"`;

// ValueType is description of schema-specific value types.
ValueType ValueTypes `xmp:"pdfaSchema:valueType"`;};

// FillModel fills in the document XMP model.
func FillModel (d *_ga .Document ,extModel *Model )error {var _a []*_ga .Namespace ;for _ ,_fg :=range d .Namespaces (){_a =append (_a ,_fg );};_f .Slice (_a ,func (_gg ,_eb int )bool {return _a [_gg ].Name < _a [_eb ].Name });for _ ,_gf :=range _a {switch _gf {case Namespace ,SchemaNS ,PropertyNS ,ValueTypeNS ,FieldNS ,_gag .NsXmp ,_ff .NsDc :continue ;
default:_gfc ,_db :=GetSchema (_gf .URI );if !_db {continue ;};_de :=d .FindModel (_gf );_fge :=*_gfc ;_fge .Property =Properties {};for _ ,_ag :=range _gfc .Property {_da :=_cg (_de ,_gf ,_ag );if _da {_fge .Property =append (_fge .Property ,_ag );};};
if len (_fge .Property )==0{continue ;};var _ef bool ;for _ed ,_fd :=range extModel .Schemas {if _fd .Schema ==_fge .Schema {_ef =true ;extModel .Schemas [_ed ]=_fge ;break ;};};if !_ef {extModel .Schemas =append (extModel .Schemas ,_fge );};};};return nil ;
};

// ValueTypeName is the name of the value type.
type ValueTypeName string ;

// SeqOfValueTypeName gets a value type name of a sequence of input value type names.
func SeqOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0073\u0065\u0071\u0020"+vt };

// ValueTypes is the slice of field value types.
type ValueTypes []FieldValueType ;

// SyncToXMP implements xmp.Model interface.
func (_cge *Model )SyncToXMP (d *_ga .Document )error {return nil };

// MarshalXMP implements xmp.Marshaler interface.
func (_ab Schemas )MarshalXMP (e *_ga .Encoder ,node *_ga .Node ,m _ga .Model )error {return _ga .MarshalArray (e ,node ,_ab .Typ (),_ab );};

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_fgeg *ValueTypes )UnmarshalXMP (d *_ga .Decoder ,node *_ga .Node ,m _ga .Model )error {return _ga .UnmarshalArray (d ,node ,_fgeg .Typ (),_fgeg );};

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_ac *Schemas )UnmarshalXMP (d *_ga .Decoder ,node *_ga .Node ,m _ga .Model )error {return _ga .UnmarshalArray (d ,node ,_ac .Typ (),_ac );};

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (_aa *PropertyCategory )UnmarshalText (in []byte )error {switch string (in ){case "\u0069\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_aa =PropertyCategoryInternal ;case "\u0065\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_aa =PropertyCategoryExternal ;
default:*_aa =PropertyCategoryUndefined ;};return nil ;};func init (){_gdb =map[string ]*Schema {_cc .NsXmpMM .URI :&XmpMediaManagementSchema ,"\u0078\u006d\u0070\u0069\u0064\u0071":&XmpIDQualSchema ,_ea .NsPDF .URI :&PdfSchema ,"\u0073\u0074\u0045v\u0074":&FieldResourceEventSchema ,"\u0073\u0074\u0056e\u0072":&FieldVersionSchema };
};var (_gdb =map[string ]*Schema {};_fcf _e .RWMutex ;);var FieldVersionSchema =Schema {NamespaceURI :"\u0068\u0074\u0074p\u003a\u002f\u002f\u006e\u0073\u002e\u0061\u0064\u006f\u0062\u0065\u002e\u0063\u006f\u006d\u002f\u0078\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u0073\u0054\u0079\u0070\u0065/\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u0023",Prefix :"\u0073\u0074\u0056e\u0072",Schema :"\u0042a\u0073\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074y\u0070\u0065\u0020\u0056\u0065\u0072\u0073\u0069\u006f\u006e",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0043\u006fmm\u0065\u006e\u0074s\u0020\u0063\u006f\u006ecer\u006ein\u0067\u0020\u0077\u0068\u0061\u0074\u0020wa\u0073\u0020\u0063\u0068\u0061\u006e\u0067e\u0064",Name :"\u0063\u006f\u006d\u006d\u0065\u006e\u0074\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0048\u0069\u0067\u0068\u0020\u006c\u0065\u0076\u0065\u006c\u002c\u0020\u0066\u006f\u0072\u006d\u0061\u006c\u0020\u0064e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0077\u0068\u0061\u0074\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0068\u0065\u0020\u0075\u0073\u0065\u0072 \u0070\u0065\u0072f\u006f\u0072\u006d\u0065\u0064\u002e",Name :"\u0065\u0076\u0065n\u0074",ValueType :ValueTypeResourceEvent },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068e\u0020\u0064\u0061\u0074\u0065\u0020\u006f\u006e\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0065\u0063\u006b\u0065\u0064\u0020\u0069\u006e\u002e",Name :"\u006d\u006f\u0064\u0069\u0066\u0079\u0044\u0061\u0074\u0065",ValueType :ValueTypeNameDate },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068e\u0020\u0070\u0065\u0072s\u006f\u006e \u0077\u0068\u006f\u0020\u006d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0074\u0068\u0069\u0073\u0020\u0076\u0065\u0072s\u0069\u006f\u006e\u002e",Name :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0072",ValueType :ValueTypeNameProperName },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 n\u0065\u0077\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e",Name :"\u0076e\u0072\u0073\u0069\u006f\u006e",ValueType :ValueTypeNameText }},ValueType :nil };
var PdfSchema =Schema {NamespaceURI :"\u0068\u0074\u0074\u0070:\u002f\u002f\u006e\u0073\u002e\u0061\u0064\u006f\u0062\u0065.\u0063o\u006d\u002f\u0070\u0064\u0066\u002f\u0031.\u0033\u002f",Prefix :"\u0070\u0064\u0066",Schema :"\u0041\u0064o\u0062\u0065\u0020P\u0044\u0046\u0020\u0053\u0063\u0068\u0065\u006d\u0061",Property :Properties {{Category :PropertyCategoryInternal ,Description :"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",Name :"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0050\u0044F\u0020\u0066\u0069l\u0065\u0020\u0076e\u0072\u0073\u0069\u006f\u006e\u0020\u0028\u0066\u006f\u0072 \u0065\u0078\u0061\u006d\u0070le\u003a\u0020\u0031\u002e\u0030\u002c\u0020\u0031\u002e\u0033\u002c\u0020\u0061\u006e\u0064\u0020\u0073\u006f\u0020\u006f\u006e\u0029\u002e",Name :"\u0050\u0044\u0046\u0056\u0065\u0072\u0073\u0069\u006f\u006e",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0074\u006f\u006fl\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074.",Name :"\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",ValueType :ValueTypeNameAgentName },{Category :PropertyCategoryInternal ,Description :"\u0049\u006e\u0064\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0074\u0068\u0069s\u0020\u0069\u0073\u0020\u0061\u0020\u0072i\u0067\u0068\u0074\u0073\u002d\u006d\u0061\u006e\u0061\u0067\u0065d\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u002e\u002e",Name :"\u004d\u0061\u0072\u006b\u0065\u0064",ValueType :ValueTypeNameBoolean }}};


// ValueType is the pdfa extension value type schema.
type ValueType struct{Description string `xmp:"pdfaType:description"`;Field []FieldValueType `xmp:"pdfaType:field"`;NamespaceURI string `xmp:"pdfaType:namespaceURI"`;Prefix string `xmp:"pdfaType:prefix"`;Type string `xmp:"pdfaType:type"`;};var (Namespace =_ga .NewNamespace ("\u0070\u0064\u0066\u0061\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e","\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0077\u0077\u0077\u002e\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067\u002f\u0070\u0064\u0066\u0061/\u006e\u0073\u002f\u0065\u0078t\u0065\u006es\u0069\u006f\u006e\u002f",NewModel );
SchemaNS =_ga .NewNamespace ("\u0070\u0064\u0066\u0061\u0053\u0063\u0068\u0065\u006d\u0061","h\u0074\u0074\u0070\u003a\u002f\u002fw\u0077\u0077\u002e\u0061\u0069\u0069m\u002e\u006f\u0072\u0067\u002f\u0070\u0064f\u0061\u002f\u006e\u0073\u002f\u0073\u0063\u0068\u0065\u006da\u0023",nil );
PropertyNS =_ga .NewNamespace ("\u0070\u0064\u0066a\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0079","\u0068\u0074t\u0070\u003a\u002f\u002fw\u0077\u0077.\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067/\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0070\u0072\u006f\u0070e\u0072\u0074\u0079\u0023",nil );
ValueTypeNS =_ga .NewNamespace ("\u0070\u0064\u0066\u0061\u0054\u0079\u0070\u0065","\u0068\u0074\u0074\u0070\u003a\u002f/\u0077\u0077\u0077\u002e\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067\u002fp\u0064\u0066\u0061\u002f\u006e\u0073\u002ft\u0079\u0070\u0065\u0023",nil );
FieldNS =_ga .NewNamespace ("\u0070d\u0066\u0061\u0046\u0069\u0065\u006cd","\u0068\u0074\u0074p:\u002f\u002f\u0077\u0077\u0077\u002e\u0061\u0069\u0069m\u002eo\u0072g\u002fp\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0066\u0069\u0065\u006c\u0064\u0023",nil );
);

// ClosedChoiceValueTypeName gets the closed choice of provided value type name.
func ClosedChoiceValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0043\u006c\u006f\u0073\u0065\u0064\u0020\u0043\u0068\u006f\u0069\u0063e\u0020\u006f\u0066\u0020"+vt ;};

// Model is the pdfa extension metadata model.
type Model struct{Schemas Schemas `xmp:"pdfaExtension:schemas"`;};

// IsZero checks if the resources list has no entries.
func (_cff Schemas )IsZero ()bool {return len (_cff )==0};

// SetSchema sets the schema into given model.
func (_dec *Model )SetSchema (namespaceURI string ,s Schema ){for _fa :=0;_fa < len (_dec .Schemas );_fa ++{if _dec .Schemas [_fa ].NamespaceURI ==namespaceURI {_dec .Schemas [_fa ]=s ;return ;};};_dec .Schemas =append (_dec .Schemas ,s );};func _cg (_eda _ga .Model ,_bb *_ga .Namespace ,_ccf Property )bool {_ae :=_fe .ValueOf (_eda ).Elem ();
if _ae .Kind ()==_fe .Ptr {_ae =_ae .Elem ();};_dac :=_ae .Type ();for _gb :=0;_gb < _dac .NumField ();_gb ++{_fde :=_dac .Field (_gb );_ee :=_fde .Tag .Get ("\u0078\u006d\u0070");if _ee ==""{continue ;};if !_c .HasPrefix (_ee ,_bb .Name ){continue ;};
_ebb :=_c .IndexRune (_ee ,':');if _ebb ==-1{continue ;};_dc :=_ee [_ebb +1:];if _dc ==_ccf .Name {_agb :=_ae .Field (_gb );return !_agb .IsZero ();};};return false ;};

// AltOfValueTypeName gets the ValueTypeName of the alt of given value type names.
func AltOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0061\u006c\u0074\u0020"+vt };func (_daf Schemas )Typ ()_ga .ArrayType {return _ga .ArrayTypeUnordered };var XmpIDQualSchema =Schema {NamespaceURI :"\u0068t\u0074\u0070:\u002f\u002f\u006e\u0073.\u0061\u0064\u006fb\u0065\u002e\u0063\u006f\u006d\u002f\u0078\u006d\u0070/I\u0064\u0065\u006et\u0069\u0066i\u0065\u0072\u002f\u0071\u0075\u0061l\u002f\u0031.\u0030\u002f",Prefix :"\u0078\u006d\u0070\u0069\u0064\u0071",Schema :"X\u004dP\u0020\u0078\u006d\u0070\u0069\u0064\u0071\u0020q\u0075\u0061\u006c\u0069fi\u0065\u0072",Property :[]Property {{Category :PropertyCategoryInternal ,Name :"\u0053\u0063\u0068\u0065\u006d\u0065",Description :"\u0041\u0020\u0071\u0075\u0061\u006c\u0069\u0066\u0069\u0065\u0072\u0020\u0070\u0072o\u0076\u0069\u0064i\u006e\u0067\u0020\u0074h\u0065\u0020\u006e\u0020\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u006c \u0069\u0064\u0065n\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0065\u0020\u0075\u0073ed\u0020\u0066\u006fr\u0020\u0061\u006e\u0020\u0069\u0074\u0065\u006d \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0078\u006d\u0070\u003a\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u002e",ValueType :ValueTypeNameText }},ValueType :nil };
var XmpMediaManagementSchema =Schema {NamespaceURI :"\u0068\u0074\u0074p\u003a\u002f\u002f\u006es\u002e\u0061\u0064\u006f\u0062\u0065\u002ec\u006f\u006d\u002f\u0078\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u006d\u006d\u002f",Prefix :"\u0078\u006d\u0070M\u004d",Schema :"X\u004dP\u0020\u004d\u0065\u0064\u0069\u0061\u0020\u004da\u006e\u0061\u0067\u0065me\u006e\u0074",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0041\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006fr\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020w\u0068\u0069\u0063\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u006f\u006e\u0065\u0020i\u0073\u0020\u0064e\u0072\u0069\u0076\u0065\u0064\u002e",Name :"D\u0065\u0072\u0069\u0076\u0065\u0064\u0046\u0072\u006f\u006d",ValueType :ValueTypeNameResourceRef },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0063\u006fm\u006d\u006f\u006e\u0020\u0069d\u0065\u006e\u0074\u0069\u0066\u0069e\u0072\u0020\u0066\u006f\u0072 \u0061\u006c\u006c\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0073 \u0061\u006e\u0064\u0020\u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u0020o\u0066\u0020\u0061\u0020\u0072\u0065\u0073\u006fu\u0072\u0063\u0065",Name :"\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"UU\u0049\u0044 \u0062\u0061\u0073\u0065\u0064\u0020\u0069\u0064\u0065n\u0074\u0069\u0066\u0069\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0020\u0069\u006e\u0063\u0061\u0072\u006e\u0061\u0074i\u006fn\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u006fc\u0075m\u0065\u006et",Name :"\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0063\u006fm\u006d\u006f\u006e\u0020\u0069d\u0065\u006e\u0074\u0069\u0066\u0069e\u0072\u0020\u0066\u006f\u0072 \u0061\u006c\u006c\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0073 \u0061\u006e\u0064\u0020\u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u0020o\u0066\u0020\u0061\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074",Name :"\u004fr\u0069g\u0069\u006e\u0061\u006c\u0044o\u0063\u0075m\u0065\u006e\u0074\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0069\u0073 \u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065",Name :"\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006eC\u006c\u0061\u0073\u0073",ValueType :ValueTypeNameRenditionClass },{Category :PropertyCategoryInternal ,Description :"\u0043\u0061n\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020t\u006f\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0072\u0065\u006e\u0064\u0069t\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0074\u0068a\u0074 \u0061r\u0065\u0020\u0074o\u006f\u0020\u0063\u006f\u006d\u0070\u006c\u0065\u0078\u0020\u006f\u0072\u0020\u0076\u0065\u0072\u0062o\u0073\u0065\u0020\u0074\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0069\u006e",Name :"\u0052e\u006ed\u0069\u0074\u0069\u006f\u006e\u0050\u0061\u0072\u0061\u006d\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u0069\u0064\u0065\u006e\u0074i\u0066\u0069\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0074\u0068i\u0073\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u002e",Name :"\u0056e\u0072\u0073\u0069\u006f\u006e\u0049D",ValueType :ValueTypeNameText },{Category :PropertyCategoryExternal ,Description :"\u0054\u0068\u0065\u0020\u0076\u0065r\u0073\u0069\u006f\u006e\u0020\u0068\u0069\u0073\u0074\u006f\u0072\u0079\u0020\u0061\u0073\u0073\u006f\u0063\u0069\u0061t\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u0072e\u0073o\u0075\u0072\u0063\u0065",Name :"\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u0073",ValueType :SeqOfValueTypeName (ValueTypeNameVersion )},{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0072\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075r\u0063\u0065\u0027\u0073\u0020\u006d\u0061n\u0061\u0067\u0065\u0072",Name :"\u004da\u006e\u0061\u0067\u0065\u0072",ValueType :ValueTypeNameAgentName },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0027\u0073\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0072 \u0076\u0061\u0072\u0069\u0061\u006e\u0074\u002e",Name :"\u004d\u0061\u006e\u0061\u0067\u0065\u0072\u0056\u0061r\u0069\u0061\u006e\u0074",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0027\u0073\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0072 \u0076\u0061\u0072\u0069\u0061\u006e\u0074\u002e",Name :"\u004d\u0061\u006e\u0061\u0067\u0065\u0072\u0056\u0061r\u0069\u0061\u006e\u0074",ValueType :ValueTypeNameText }}};


// MarshalXMP implements xmp.Marshaler interface.
func (_ggc Properties )MarshalXMP (e *_ga .Encoder ,node *_ga .Node ,m _ga .Model )error {return _ga .MarshalArray (e ,node ,_ggc .Typ (),_ggc );};

// SyncFromXMP implements xmp.Model interface.
func (_edad *Model )SyncFromXMP (d *_ga .Document )error {return nil };const (ValueTypeNameBoolean ValueTypeName ="\u0042o\u006f\u006c\u0065\u0061\u006e";ValueTypeNameDate ValueTypeName ="\u0044\u0061\u0074\u0065";ValueTypeNameInteger ValueTypeName ="\u0049n\u0074\u0065\u0067\u0065\u0072";
ValueTypeNameReal ValueTypeName ="\u0052\u0065\u0061\u006c";ValueTypeNameText ValueTypeName ="\u0054\u0065\u0078\u0074";ValueTypeNameAgentName ValueTypeName ="\u0041g\u0065\u006e\u0074\u004e\u0061\u006de";ValueTypeNameProperName ValueTypeName ="\u0050\u0072\u006f\u0070\u0065\u0072\u004e\u0061\u006d\u0065";
ValueTypeNameXPath ValueTypeName ="\u0058\u0050\u0061t\u0068";ValueTypeNameGUID ValueTypeName ="\u0047\u0055\u0049\u0044";ValueTypeNameLocale ValueTypeName ="\u004c\u006f\u0063\u0061\u006c\u0065";ValueTypeNameMIMEType ValueTypeName ="\u004d\u0049\u004d\u0045\u0054\u0079\u0070\u0065";
ValueTypeNameRenditionClass ValueTypeName ="\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006eC\u006c\u0061\u0073\u0073";ValueTypeNameResourceRef ValueTypeName ="R\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0052\u0065\u0066";ValueTypeNameURL ValueTypeName ="\u0055\u0052\u004c";
ValueTypeNameURI ValueTypeName ="\u0055\u0052\u0049";ValueTypeNameVersion ValueTypeName ="\u0056e\u0072\u0073\u0069\u006f\u006e";);

// FieldValueType is a schema that describes a field in a structured type.
type FieldValueType struct{Name string `xmp:"pdfaField:description"`;Description string `xmp:"pdfaField:name"`;ValueType ValueTypeName `xmp:"pdfaField:valueType"`;};

// Typ gets array type of the field value types.
func (_fgf ValueTypes )Typ ()_ga .ArrayType {return _ga .ArrayTypeOrdered };

// Can implements xmp.Model interface.
func (_deb *Model )Can (nsName string )bool {return Namespace .GetName ()==nsName };

// SyncModel implements xmp.Model interface.
func (_efb *Model )SyncModel (d *_ga .Document )error {return nil };

// Properties is a list of properties.
type Properties []Property ;

// MarshalXMP implements xmp.Marshaler interface.
func (_abg ValueTypes )MarshalXMP (e *_ga .Encoder ,node *_ga .Node ,m _ga .Model )error {return _ga .MarshalArray (e ,node ,_abg .Typ (),_abg );};

// CanTag implements xmp.Model interface.
func (_dd *Model )CanTag (tag string )bool {_ ,_ba :=_ga .GetNativeField (_dd ,tag );return _ba ==nil };

// Typ gets the array type of properties.
func (_agc Properties )Typ ()_ga .ArrayType {return _ga .ArrayTypeOrdered };const (ValueTypeResourceEvent ValueTypeName ="\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0045\u0076\u0065\u006e\u0074";);

// NewModel creates a new pdfAExtension model.
func NewModel (name string )_ga .Model {return &Model {}};var FieldResourceEventSchema =Schema {NamespaceURI :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u006es\u002e\u0061\u0064ob\u0065\u002e\u0063\u006f\u006d\u002fx\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u0073\u0054\u0079\u0070\u0065\u002f\u0052\u0065s\u006f\u0075\u0072\u0063\u0065\u0045\u0076\u0065n\u0074\u0023",Prefix :"\u0073\u0074\u0045v\u0074",Schema :"\u0044e\u0066\u0069n\u0069\u0074\u0069\u006fn\u0020\u006f\u0066 \u0062\u0061\u0073\u0069\u0063\u0020\u0076\u0061\u006cue\u0020\u0074\u0079p\u0065\u0020R\u0065\u0073\u006f\u0075\u0072\u0063e\u0045\u0076e\u006e\u0074",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0054he\u0020a\u0063t\u0069\u006f\u006e\u0020\u0074\u0068a\u0074\u0020\u006f\u0063c\u0075\u0072\u0072\u0065\u0064\u002e\u0020\u0044\u0065\u0066\u0069\u006e\u0065d \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0061\u0072\u0065\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065d\u002c \u0063\u006f\u0070\u0069\u0065\u0064\u002c\u0020\u0063\u0072\u0065\u0061\u0074e\u0064\u002c \u0063\u0072\u006fp\u0070\u0065\u0064\u002c\u0020\u0065\u0064\u0069\u0074ed\u002c\u0020\u0066i\u006c\u002d\u0074\u0065r\u0065\u0064\u002c\u0020\u0066\u006fr\u006d\u0061t\u0074\u0065\u0064\u002c\u0020\u0076\u0065\u0072s\u0069\u006f\u006e\u005f\u0075\u0070\u0064\u0061\u0074\u0065\u0064\u002c\u0020\u0070\u0072\u0069\u006e\u0074\u0065\u0064\u002c\u0020\u0070ubli\u0073\u0068\u0065\u0064\u002c\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0064\u002c\u0020\u0070\u0072\u006f\u0064\u0075\u0063\u0065\u0064\u002c\u0020\u0072\u0065\u0073i\u007ae\u0064.\u004e\u0065\u0077\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u0076\u0065r\u0062\u0073 \u0069n\u0020\u0074\u0068\u0065\u0020\u0070a\u0073\u0074\u0020\u0074\u0065\u006e\u0073\u0065\u002e",Name :"\u0061\u0063\u0074\u0069\u006f\u006e",ValueType :"O\u0070\u0065\u006e\u0020\u0043\u0068\u006f\u0069\u0063\u0065"},{Category :PropertyCategoryInternal ,Description :"T\u0068\u0065\u0020\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0020\u0049\u0044\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u006d\u006f\u0064\u0069\u0066\u0069\u0065d \u0072\u0065\u0073o\u0075r\u0063\u0065",Name :"\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0041\u0064di\u0074\u0069\u006fn\u0061\u006c\u0020\u0064esc\u0072ip\u0074\u0069\u006f\u006e\u0020\u006f\u0066 t\u0068\u0065\u0020\u0061\u0063\u0074\u0069o\u006e",Name :"\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054h\u0065\u0020s\u006f\u0066\u0074\u0077a\u0072\u0065\u0020a\u0067\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074 p\u0065\u0072\u0066o\u0072\u006de\u0064\u0020\u0074\u0068\u0065\u0020a\u0063\u0074i\u006f\u006e",Name :"\u0073\u006f\u0066\u0074\u0077\u0061\u0072\u0065\u0041\u0067\u0065\u006e\u0074",ValueType :ValueTypeNameAgentName },{Category :PropertyCategoryInternal ,Description :"\u004f\u0070t\u0069\u006f\u006e\u0061\u006c\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u006f\u0066\u0020\u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u006f\u0063\u0063\u0075\u0072\u0072\u0065\u0064",Name :"\u0077\u0068\u0065\u006e",ValueType :ValueTypeNameDate }},ValueType :nil };


// Property is a schema that describes single property.
type Property struct{

// Category is the property category.
Category PropertyCategory `xmp:"pdfaProperty:category"`;

// Description of the property.
Description string `xmp:"pdfaProperty:description"`;

// Name is a property name.
Name string `xmp:"pdfaProperty:name"`;

// ValueType is the property value type.
ValueType ValueTypeName `xmp:"pdfaProperty:valueType"`;};

// Namespaces implements xmp.Model interface.
func (_gbb *Model )Namespaces ()_ga .NamespaceList {return _ga .NamespaceList {Namespace }};

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_eea *Properties )UnmarshalXMP (d *_ga .Decoder ,node *_ga .Node ,m _ga .Model )error {return _ga .UnmarshalArray (d ,node ,_eea .Typ (),_eea );};

// PropertyCategory is the property category enumerator.
type PropertyCategory int ;

// MarshalText implements encoding.TextMarshaler interface.
func (_geg PropertyCategory )MarshalText ()([]byte ,error ){switch _geg {case PropertyCategoryInternal :return []byte ("\u0069\u006e\u0074\u0065\u0072\u006e\u0061\u006c"),nil ;case PropertyCategoryExternal :return []byte ("\u0065\u0078\u0074\u0065\u0072\u006e\u0061\u006c"),nil ;
case PropertyCategoryUndefined :return []byte (""),nil ;default:return nil ,_g .Errorf ("\u0075\u006ed\u0065\u0066\u0069\u006ee\u0064\u0020p\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020c\u0061\u0074\u0065\u0067\u006f\u0072\u0079\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_geg );
};};

// Schemas is the array of xmp metadata extension resources.
type Schemas []Schema ;

// GetSchema for provided namespace.
func GetSchema (namespaceURI string )(*Schema ,bool ){_fcf .RLock ();defer _fcf .RUnlock ();_baf ,_afb :=_gdb [namespaceURI ];return _baf ,_afb ;};