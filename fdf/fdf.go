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

// Package fdf provides support for loading form field data from Form Field Data (FDF) files.
package fdf ;import (_dd "bufio";_cc "bytes";_d "encoding/hex";_e "errors";_ca "fmt";_cea "github.com/unidoc/unipdf/v3/common";_bg "github.com/unidoc/unipdf/v3/core";_a "io";_dc "os";_ef "regexp";_f "sort";_ce "strconv";_c "strings";);var _afc =_ef .MustCompile ("\u0025\u0025\u0045O\u0046");
func (_bbe *fdfParser )parseArray ()(*_bg .PdfObjectArray ,error ){_cfc :=_bg .MakeArray ();_bbe ._dg .ReadByte ();for {_bbe .skipSpaces ();_bbd ,_beb :=_bbe ._dg .Peek (1);if _beb !=nil {return _cfc ,_beb ;};if _bbd [0]==']'{_bbe ._dg .ReadByte ();break ;
};_bbg ,_beb :=_bbe .parseObject ();if _beb !=nil {return _cfc ,_beb ;};_cfc .Append (_bbg );};return _cfc ,nil ;};func (_fde *fdfParser )parseName ()(_bg .PdfObjectName ,error ){var _fdd _cc .Buffer ;_fgd :=false ;for {_cef ,_gdc :=_fde ._dg .Peek (1);
if _gdc ==_a .EOF {break ;};if _gdc !=nil {return _bg .PdfObjectName (_fdd .String ()),_gdc ;};if !_fgd {if _cef [0]=='/'{_fgd =true ;_fde ._dg .ReadByte ();}else if _cef [0]=='%'{_fde .readComment ();_fde .skipSpaces ();}else {_cea .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029",_cef ,_cef );
return _bg .PdfObjectName (_fdd .String ()),_ca .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029",_cef [0]);};}else {if _bg .IsWhiteSpace (_cef [0]){break ;}else if (_cef [0]=='/')||(_cef [0]=='[')||(_cef [0]=='(')||(_cef [0]==']')||(_cef [0]=='<')||(_cef [0]=='>'){break ;
}else if _cef [0]=='#'{_edfd ,_adg :=_fde ._dg .Peek (3);if _adg !=nil {return _bg .PdfObjectName (_fdd .String ()),_adg ;};_fde ._dg .Discard (3);_abb ,_adg :=_d .DecodeString (string (_edfd [1:3]));if _adg !=nil {return _bg .PdfObjectName (_fdd .String ()),_adg ;
};_fdd .Write (_abb );}else {_de ,_ :=_fde ._dg .ReadByte ();_fdd .WriteByte (_de );};};};return _bg .PdfObjectName (_fdd .String ()),nil ;};func (_aa *fdfParser )skipSpaces ()(int ,error ){_ga :=0;for {_ea ,_ccd :=_aa ._dg .ReadByte ();if _ccd !=nil {return 0,_ccd ;
};if _bg .IsWhiteSpace (_ea ){_ga ++;}else {_aa ._dg .UnreadByte ();break ;};};return _ga ,nil ;};

// Load loads FDF form data from `r`.
func Load (r _a .ReadSeeker )(*Data ,error ){_ccc ,_ee :=_cdbd (r );if _ee !=nil {return nil ,_ee ;};_cb ,_ee :=_ccc .Root ();if _ee !=nil {return nil ,_ee ;};_bc ,_fc :=_bg .GetArray (_cb .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_fc {return nil ,_e .New ("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067");
};return &Data {_cdb :_bc ,_cd :_cb },nil ;};var _afa =_ef .MustCompile ("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029");var _bef =_ef .MustCompile ("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029");
func (_bfe *fdfParser )readTextLine ()(string ,error ){var _gda _cc .Buffer ;for {_aba ,_ebg :=_bfe ._dg .Peek (1);if _ebg !=nil {_cea .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_ebg .Error ());return _gda .String (),_ebg ;};if (_aba [0]!='\r')&&(_aba [0]!='\n'){_eef ,_ :=_bfe ._dg .ReadByte ();
_gda .WriteByte (_eef );}else {break ;};};return _gda .String (),nil ;};type fdfParser struct{_fge int ;_bf int ;_db map[int64 ]_bg .PdfObject ;_dbc _a .ReadSeeker ;_dg *_dd .Reader ;_ae int64 ;_daf *_bg .PdfObjectDictionary ;};

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data )FieldDictionaries ()(map[string ]*_bg .PdfObjectDictionary ,error ){_ebf :=map[string ]*_bg .PdfObjectDictionary {};for _ced :=0;_ced < fdf ._cdb .Len ();_ced ++{_ad ,_fcc :=_bg .GetDict (fdf ._cdb .Get (_ced ));if _fcc {_bgf ,_ :=_bg .GetString (_ad .Get ("\u0054"));
if _bgf !=nil {_ebf [_bgf .Str ()]=_ad ;};};};return _ebf ,nil ;};func (_ebgc *fdfParser )parseHexString ()(*_bg .PdfObjectString ,error ){_ebgc ._dg .ReadByte ();var _ebb _cc .Buffer ;for {_bee ,_caab :=_ebgc ._dg .Peek (1);if _caab !=nil {return _bg .MakeHexString (""),_caab ;
};if _bee [0]=='>'{_ebgc ._dg .ReadByte ();break ;};_caf ,_ :=_ebgc ._dg .ReadByte ();if !_bg .IsWhiteSpace (_caf ){_ebb .WriteByte (_caf );};};if _ebb .Len ()%2==1{_ebb .WriteRune ('0');};_cbg ,_edbb :=_d .DecodeString (_ebb .String ());if _edbb !=nil {_cea .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067",_ebb .String ());
return _bg .MakeHexString (""),nil ;};return _bg .MakeHexString (string (_cbg )),nil ;};func (_aaa *fdfParser )parseBool ()(_bg .PdfObjectBool ,error ){_fbg ,_cafa :=_aaa ._dg .Peek (4);if _cafa !=nil {return _bg .PdfObjectBool (false ),_cafa ;};if (len (_fbg )>=4)&&(string (_fbg [:4])=="\u0074\u0072\u0075\u0065"){_aaa ._dg .Discard (4);
return _bg .PdfObjectBool (true ),nil ;};_fbg ,_cafa =_aaa ._dg .Peek (5);if _cafa !=nil {return _bg .PdfObjectBool (false ),_cafa ;};if (len (_fbg )>=5)&&(string (_fbg [:5])=="\u0066\u0061\u006cs\u0065"){_aaa ._dg .Discard (5);return _bg .PdfObjectBool (false ),nil ;
};return _bg .PdfObjectBool (false ),_e .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");};func _ecgc (_fdde string )(*fdfParser ,error ){_aafa :=fdfParser {};_fdf :=[]byte (_fdde );
_bgeg :=_cc .NewReader (_fdf );_aafa ._dbc =_bgeg ;_aafa ._db =map[int64 ]_bg .PdfObject {};_bfeg :=_dd .NewReader (_bgeg );_aafa ._dg =_bfeg ;_aafa ._ae =int64 (len (_fdde ));return &_aafa ,_aafa .parse ();};var _ff =_ef .MustCompile ("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029");
func _cdbd (_efe _a .ReadSeeker )(*fdfParser ,error ){_bfc :=&fdfParser {};_bfc ._dbc =_efe ;_bfc ._db =map[int64 ]_bg .PdfObject {};_fegd ,_abef ,_dcf :=_bfc .parseFdfVersion ();if _dcf !=nil {_cea .Log .Error ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076",_dcf );
return nil ,_dcf ;};_bfc ._fge =_fegd ;_bfc ._bf =_abef ;_dcf =_bfc .parse ();return _bfc ,_dcf ;};func (_df *fdfParser )setFileOffset (_bec int64 ){_df ._dbc .Seek (_bec ,_a .SeekStart );_df ._dg =_dd .NewReader (_df ._dbc );};func _bge (_bdg string )(_bg .PdfObjectReference ,error ){_edfdc :=_bg .PdfObjectReference {};
_dce :=_bd .FindStringSubmatch (_bdg );if len (_dce )< 3{_cea .Log .Debug ("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065");return _edfdc ,_e .New ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e");
};_abc ,_befb :=_ce .Atoi (_dce [1]);if _befb !=nil {_cea .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030",_dce [1]);
return _edfdc ,nil ;};_edfdc .ObjectNumber =int64 (_abc );_ffcf ,_befb :=_ce .Atoi (_dce [2]);if _befb !=nil {_cea .Log .Debug ("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030",_dce [2]);
return _edfdc ,nil ;};_edfdc .GenerationNumber =int64 (_ffcf );return _edfdc ,nil ;};func (_dbf *fdfParser )parseNumber ()(_bg .PdfObject ,error ){return _bg .ParseNumber (_dbf ._dg )};

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data )FieldValues ()(map[string ]_bg .PdfObject ,error ){_ab ,_cg :=fdf .FieldDictionaries ();if _cg !=nil {return nil ,_cg ;};var _bb []string ;for _fg :=range _ab {_bb =append (_bb ,_fg );};_f .Strings (_bb );_be :=map[string ]_bg .PdfObject {};
for _ ,_bga :=range _bb {_da :=_ab [_bga ];_fb :=_bg .TraceToDirectObject (_da .Get ("\u0056"));_be [_bga ]=_fb ;};return _be ,nil ;};

// Root returns the Root of the FDF document.
func (_fea *fdfParser )Root ()(*_bg .PdfObjectDictionary ,error ){if _fea ._daf !=nil {if _dbd ,_abcc :=_fea .trace (_fea ._daf .Get ("\u0052\u006f\u006f\u0074")).(*_bg .PdfObjectDictionary );_abcc {if _ccb ,_bcf :=_fea .trace (_dbd .Get ("\u0046\u0044\u0046")).(*_bg .PdfObjectDictionary );
_bcf {return _ccb ,nil ;};};};var _aaeg []int64 ;for _gca :=range _fea ._db {_aaeg =append (_aaeg ,_gca );};_f .Slice (_aaeg ,func (_fce ,_cdgf int )bool {return _aaeg [_fce ]< _aaeg [_cdgf ]});for _ ,_dcg :=range _aaeg {_egag :=_fea ._db [_dcg ];if _gdcc ,_bfb :=_fea .trace (_egag ).(*_bg .PdfObjectDictionary );
_bfb {if _gdf ,_ebfc :=_fea .trace (_gdcc .Get ("\u0046\u0044\u0046")).(*_bg .PdfObjectDictionary );_ebfc {return _gdf ,nil ;};};};return nil ,_e .New ("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};var _eg =_ef .MustCompile ("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a");
func (_edf *fdfParser )readComment ()(string ,error ){var _cae _cc .Buffer ;_ ,_edb :=_edf .skipSpaces ();if _edb !=nil {return _cae .String (),_edb ;};_gbe :=true ;for {_bfa ,_ebd :=_edf ._dg .Peek (1);if _ebd !=nil {_cea .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_ebd .Error ());
return _cae .String (),_ebd ;};if _gbe &&_bfa [0]!='%'{return _cae .String (),_e .New ("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025");};_gbe =false ;if (_bfa [0]!='\r')&&(_bfa [0]!='\n'){_bag ,_ :=_edf ._dg .ReadByte ();
_cae .WriteByte (_bag );}else {break ;};};return _cae .String (),nil ;};var _bd =_ef .MustCompile ("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052");func (_ebcb *fdfParser )parseNull ()(_bg .PdfObjectNull ,error ){_ ,_gea :=_ebcb ._dg .Discard (4);
return _bg .PdfObjectNull {},_gea ;};func (_gaf *fdfParser )skipComments ()error {if _ ,_fa :=_gaf .skipSpaces ();_fa !=nil {return _fa ;};_dgf :=true ;for {_bgfa ,_bgfaf :=_gaf ._dg .Peek (1);if _bgfaf !=nil {_cea .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_bgfaf .Error ());
return _bgfaf ;};if _dgf &&_bgfa [0]!='%'{return nil ;};_dgf =false ;if (_bgfa [0]!='\r')&&(_bgfa [0]!='\n'){_gaf ._dg .ReadByte ();}else {break ;};};return _gaf .skipComments ();};func (_eee *fdfParser )seekFdfVersionTopDown ()(int ,int ,error ){_eee ._dbc .Seek (0,_a .SeekStart );
_eee ._dg =_dd .NewReader (_eee ._dbc );_bcc :=20;_efcb :=make ([]byte ,_bcc );for {_gfd ,_egeb :=_eee ._dg .ReadByte ();if _egeb !=nil {if _egeb ==_a .EOF {break ;}else {return 0,0,_egeb ;};};if _bg .IsDecimalDigit (_gfd )&&_efcb [_bcc -1]=='.'&&_bg .IsDecimalDigit (_efcb [_bcc -2])&&_efcb [_bcc -3]=='-'&&_efcb [_bcc -4]=='F'&&_efcb [_bcc -5]=='D'&&_efcb [_bcc -6]=='P'{_gcd :=int (_efcb [_bcc -2]-'0');
_bdcb :=int (_gfd -'0');return _gcd ,_bdcb ,nil ;};_efcb =append (_efcb [1:_bcc ],_gfd );};return 0,0,_e .New ("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};func (_ac *fdfParser )parseObject ()(_bg .PdfObject ,error ){_cea .Log .Trace ("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074");
_ac .skipSpaces ();for {_aae ,_bda :=_ac ._dg .Peek (2);if _bda !=nil {return nil ,_bda ;};_cea .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_aae ));if _aae [0]=='/'{_cgd ,_dga :=_ac .parseName ();
_cea .Log .Trace ("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027",_cgd );return &_cgd ,_dga ;}else if _aae [0]=='('{_cea .Log .Trace ("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!");return _ac .parseString ();}else if _aae [0]=='['{_cea .Log .Trace ("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021");
return _ac .parseArray ();}else if (_aae [0]=='<')&&(_aae [1]=='<'){_cea .Log .Trace ("\u002d>\u0044\u0069\u0063\u0074\u0021");return _ac .parseDict ();}else if _aae [0]=='<'{_cea .Log .Trace ("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021");
return _ac .parseHexString ();}else if _aae [0]=='%'{_ac .readComment ();_ac .skipSpaces ();}else {_cea .Log .Trace ("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f");_aae ,_ =_ac ._dg .Peek (15);_gafa :=string (_aae );
_cea .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_gafa );if (len (_gafa )> 3)&&(_gafa [:4]=="\u006e\u0075\u006c\u006c"){_acf ,_edfc :=_ac .parseNull ();return &_acf ,_edfc ;}else if (len (_gafa )> 4)&&(_gafa [:5]=="\u0066\u0061\u006cs\u0065"){_gag ,_cded :=_ac .parseBool ();
return &_gag ,_cded ;}else if (len (_gafa )> 3)&&(_gafa [:4]=="\u0074\u0072\u0075\u0065"){_fee ,_faf :=_ac .parseBool ();return &_fee ,_faf ;};_fgdd :=_bd .FindStringSubmatch (_gafa );if len (_fgdd )> 1{_aae ,_ =_ac ._dg .ReadBytes ('R');_cea .Log .Trace ("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027",string (_aae [:]));
_dea ,_feg :=_bge (string (_aae ));return &_dea ,_feg ;};_ecg :=_ff .FindStringSubmatch (_gafa );if len (_ecg )> 1{_cea .Log .Trace ("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021");return _ac .parseNumber ();};_ecg =_bef .FindStringSubmatch (_gafa );
if len (_ecg )> 1{_cea .Log .Trace ("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021");_cea .Log .Trace ("\u0025\u0020\u0073",_ecg );return _ac .parseNumber ();};_cea .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)",_gafa );
return nil ,_e .New ("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e");
};};};func (_gg *fdfParser )parse ()error {_gg ._dbc .Seek (0,_a .SeekStart );_gg ._dg =_dd .NewReader (_gg ._dbc );for {_gg .skipComments ();_fdge ,_ddg :=_gg ._dg .Peek (20);if _ddg !=nil {_cea .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");
return _ddg ;};if _c .HasPrefix (string (_fdge ),"\u0074r\u0061\u0069\u006c\u0065\u0072"){_gg ._dg .Discard (7);_gg .skipSpaces ();_gg .skipComments ();_ccbf ,_ :=_gg .parseDict ();_gg ._daf =_ccbf ;break ;};_bcd :=_eg .FindStringSubmatchIndex (string (_fdge ));
if len (_bcd )< 6{_cea .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_fdge ));
return _e .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_ebe ,_ddg :=_gg .parseIndirectObject ();if _ddg !=nil {return _ddg ;};switch _gcc :=_ebe .(type ){case *_bg .PdfIndirectObject :_gg ._db [_gcc .ObjectNumber ]=_gcc ;case *_bg .PdfObjectStream :_gg ._db [_gcc .ObjectNumber ]=_gcc ;default:return _e .New ("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};};return nil ;};func (_ega *fdfParser )parseString ()(*_bg .PdfObjectString ,error ){_ega ._dg .ReadByte ();var _eae _cc .Buffer ;_gbf :=1;for {_ge ,_dbg :=_ega ._dg .Peek (1);if _dbg !=nil {return _bg .MakeString (_eae .String ()),_dbg ;};if _ge [0]=='\\'{_ega ._dg .ReadByte ();
_fe ,_fdeg :=_ega ._dg .ReadByte ();if _fdeg !=nil {return _bg .MakeString (_eae .String ()),_fdeg ;};if _bg .IsOctalDigit (_fe ){_ada ,_afab :=_ega ._dg .Peek (2);if _afab !=nil {return _bg .MakeString (_eae .String ()),_afab ;};var _gf []byte ;_gf =append (_gf ,_fe );
for _ ,_cefe :=range _ada {if _bg .IsOctalDigit (_cefe ){_gf =append (_gf ,_cefe );}else {break ;};};_ega ._dg .Discard (len (_gf )-1);_cea .Log .Trace ("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022",_gf );
_cff ,_afab :=_ce .ParseUint (string (_gf ),8,32);if _afab !=nil {return _bg .MakeString (_eae .String ()),_afab ;};_eae .WriteByte (byte (_cff ));continue ;};switch _fe {case 'n':_eae .WriteRune ('\n');case 'r':_eae .WriteRune ('\r');case 't':_eae .WriteRune ('\t');
case 'b':_eae .WriteRune ('\b');case 'f':_eae .WriteRune ('\f');case '(':_eae .WriteRune ('(');case ')':_eae .WriteRune (')');case '\\':_eae .WriteRune ('\\');};continue ;}else if _ge [0]=='('{_gbf ++;}else if _ge [0]==')'{_gbf --;if _gbf ==0{_ega ._dg .ReadByte ();
break ;};};_bdc ,_ :=_ega ._dg .ReadByte ();_eae .WriteByte (_bdc );};return _bg .MakeString (_eae .String ()),nil ;};func (_gb *fdfParser )readAtLeast (_ec []byte ,_abe int )(int ,error ){_ba :=_abe ;_gd :=0;_dab :=0;for _ba > 0{_ed ,_fd :=_gb ._dg .Read (_ec [_gd :]);
if _fd !=nil {_cea .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073",_ed ,_dab ,_fd .Error ());return _gd ,_e .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067");
};_dab ++;_gd +=_ed ;_ba -=_ed ;};return _gd ,nil ;};

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath (fdfPath string )(*Data ,error ){_cf ,_eb :=_dc .Open (fdfPath );if _eb !=nil {return nil ,_eb ;};defer _cf .Close ();return Load (_cf );};func (_cee *fdfParser )seekToEOFMarker (_fag int64 )error {_aec :=int64 (0);_dbgc :=int64 (1000);
for _aec < _fag {if _fag <=(_dbgc +_aec ){_dbgc =_fag -_aec ;};_ ,_fdg :=_cee ._dbc .Seek (-_aec -_dbgc ,_a .SeekEnd );if _fdg !=nil {return _fdg ;};_dfc :=make ([]byte ,_dbgc );_cee ._dbc .Read (_dfc );_cea .Log .Trace ("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022",string (_dfc ));
_gbd :=_afc .FindAllStringIndex (string (_dfc ),-1);if _gbd !=nil {_abcg :=_gbd [len (_gbd )-1];_cea .Log .Trace ("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064",_gbd );_cee ._dbc .Seek (-_aec -_dbgc +int64 (_abcg [0]),_a .SeekEnd );return nil ;};_cea .Log .Debug ("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067");
_aec +=_dbgc ;};_cea .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e");return _e .New ("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");
};func (_cdeb *fdfParser )trace (_abdd _bg .PdfObject )_bg .PdfObject {switch _cgc :=_abdd .(type ){case *_bg .PdfObjectReference :_dgab ,_bde :=_cdeb ._db [_cgc .ObjectNumber ].(*_bg .PdfIndirectObject );if _bde {return _dgab .PdfObject ;};_cea .Log .Debug ("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
return nil ;case *_bg .PdfIndirectObject :return _cgc .PdfObject ;};return _abdd ;};func (_ecd *fdfParser )parseFdfVersion ()(int ,int ,error ){_ecd ._dbc .Seek (0,_a .SeekStart );_beec :=20;_abgf :=make ([]byte ,_beec );_ecd ._dbc .Read (_abgf );_cdf :=_afa .FindStringSubmatch (string (_abgf ));
if len (_cdf )< 3{_eed ,_bbf ,_fed :=_ecd .seekFdfVersionTopDown ();if _fed !=nil {_cea .Log .Debug ("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e");
return 0,0,_fed ;};return _eed ,_bbf ,nil ;};_eff ,_ceae :=_ce .Atoi (_cdf [1]);if _ceae !=nil {return 0,0,_ceae ;};_gfg ,_ceae :=_ce .Atoi (_cdf [2]);if _ceae !=nil {return 0,0,_ceae ;};_cea .Log .Debug ("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064",_eff ,_gfg );
return _eff ,_gfg ,nil ;};func (_caa *fdfParser )getFileOffset ()int64 {_gc ,_ :=_caa ._dbc .Seek (0,_a .SeekCurrent );_gc -=int64 (_caa ._dg .Buffered ());return _gc ;};func (_abd *fdfParser )parseIndirectObject ()(_bg .PdfObject ,error ){_aab :=_bg .PdfIndirectObject {};
_cea .Log .Trace ("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a");_gbc ,_efc :=_abd ._dg .Peek (20);if _efc !=nil {_cea .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");
return &_aab ,_efc ;};_cea .Log .Trace ("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022",string (_gbc ));_ddbc :=_eg .FindStringSubmatchIndex (string (_gbc ));if len (_ddbc )< 6{_cea .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_gbc ));
return &_aab ,_e .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_abd ._dg .Discard (_ddbc [0]);_cea .Log .Trace ("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064",_ddbc );_eeb :=_ddbc [1]-_ddbc [0];_fccb :=make ([]byte ,_eeb );_ ,_efc =_abd .readAtLeast (_fccb ,_eeb );if _efc !=nil {_cea .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073",_efc );
return nil ,_efc ;};_cea .Log .Trace ("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073",_fccb );_dbgd :=_eg .FindStringSubmatch (string (_fccb ));if len (_dbgd )< 3{_cea .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_fccb ));
return &_aab ,_e .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_dabb ,_ :=_ce .Atoi (_dbgd [1]);_def ,_ :=_ce .Atoi (_dbgd [2]);_aab .ObjectNumber =int64 (_dabb );_aab .GenerationNumber =int64 (_def );for {_ege ,_fae :=_abd ._dg .Peek (2);if _fae !=nil {return &_aab ,_fae ;};_cea .Log .Trace ("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_ege ),string (_ege ));
if _bg .IsWhiteSpace (_ege [0]){_abd .skipSpaces ();}else if _ege [0]=='%'{_abd .skipComments ();}else if (_ege [0]=='<')&&(_ege [1]=='<'){_cea .Log .Trace ("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074");_aab .PdfObject ,_fae =_abd .parseDict ();
_cea .Log .Trace ("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076",_fae );if _fae !=nil {return &_aab ,_fae ;};_cea .Log .Trace ("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e");
}else if (_ege [0]=='/')||(_ege [0]=='(')||(_ege [0]=='[')||(_ege [0]=='<'){_aab .PdfObject ,_fae =_abd .parseObject ();if _fae !=nil {return &_aab ,_fae ;};_cea .Log .Trace ("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e");
}else {if _ege [0]=='e'{_adf ,_bcg :=_abd .readTextLine ();if _bcg !=nil {return nil ,_bcg ;};if len (_adf )>=6&&_adf [0:6]=="\u0065\u006e\u0064\u006f\u0062\u006a"{break ;};}else if _ege [0]=='s'{_ege ,_ =_abd ._dg .Peek (10);if string (_ege [:6])=="\u0073\u0074\u0072\u0065\u0061\u006d"{_fdbg :=6;
if len (_ege )> 6{if _bg .IsWhiteSpace (_ege [_fdbg ])&&_ege [_fdbg ]!='\r'&&_ege [_fdbg ]!='\n'{_cea .Log .Debug ("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072");
_fdbg ++;};if _ege [_fdbg ]=='\r'{_fdbg ++;if _ege [_fdbg ]=='\n'{_fdbg ++;};}else if _ege [_fdbg ]=='\n'{_fdbg ++;};};_abd ._dg .Discard (_fdbg );_cdg ,_bgg :=_aab .PdfObject .(*_bg .PdfObjectDictionary );if !_bgg {return nil ,_e .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};_cea .Log .Trace ("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073",_cdg );_cafe ,_gdaf :=_cdg .Get ("\u004c\u0065\u006e\u0067\u0074\u0068").(*_bg .PdfObjectInteger );if !_gdaf {return nil ,_e .New ("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072");
};_fga :=*_cafe ;if _fga < 0{return nil ,_e .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030");};if int64 (_fga )> _abd ._ae {_cea .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065");
return nil ,_e .New ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065");};_cac :=make ([]byte ,_fga );
_ ,_fae =_abd .readAtLeast (_cac ,int (_fga ));if _fae !=nil {_cea .Log .Debug ("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058",len (_cac ),_cac );_cea .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_fae );
return nil ,_fae ;};_abeb :=_bg .PdfObjectStream {};_abeb .Stream =_cac ;_abeb .PdfObjectDictionary =_aab .PdfObject .(*_bg .PdfObjectDictionary );_abeb .ObjectNumber =_aab .ObjectNumber ;_abeb .GenerationNumber =_aab .GenerationNumber ;_abd .skipSpaces ();
_abd ._dg .Discard (9);_abd .skipSpaces ();return &_abeb ,nil ;};};_aab .PdfObject ,_fae =_abd .parseObject ();return &_aab ,_fae ;};};_cea .Log .Trace ("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021");
return &_aab ,nil ;};

// Data represents forms data format (FDF) file data.
type Data struct{_cd *_bg .PdfObjectDictionary ;_cdb *_bg .PdfObjectArray ;};func (_aaf *fdfParser )parseDict ()(*_bg .PdfObjectDictionary ,error ){_cea .Log .Trace ("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021");
_fdb :=_bg .MakeDict ();_bdb ,_ :=_aaf ._dg .ReadByte ();if _bdb !='<'{return nil ,_e .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};_bdb ,_ =_aaf ._dg .ReadByte ();if _bdb !='<'{return nil ,_e .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");
};for {_aaf .skipSpaces ();_aaf .skipComments ();_eaf ,_fca :=_aaf ._dg .Peek (2);if _fca !=nil {return nil ,_fca ;};_cea .Log .Trace ("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_eaf ),string (_eaf ));
if (_eaf [0]=='>')&&(_eaf [1]=='>'){_cea .Log .Trace ("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");_aaf ._dg .ReadByte ();_aaf ._dg .ReadByte ();break ;};_cea .Log .Trace ("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021");
_eefg ,_fca :=_aaf .parseName ();_cea .Log .Trace ("\u004be\u0079\u003a\u0020\u0025\u0073",_eefg );if _fca !=nil {_cea .Log .Debug ("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073",_fca );
return nil ,_fca ;};if len (_eefg )> 4&&_eefg [len (_eefg )-4:]=="\u006e\u0075\u006c\u006c"{_ffd :=_eefg [0:len (_eefg )-4];_cea .Log .Debug ("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029",_eefg );
_cea .Log .Debug ("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c",_ffd );_aaf .skipSpaces ();_ddb ,_ :=_aaf ._dg .Peek (1);if _ddb [0]=='/'{_fdb .Set (_ffd ,_bg .MakeNull ());continue ;};};_aaf .skipSpaces ();
_fbe ,_fca :=_aaf .parseObject ();if _fca !=nil {return nil ,_fca ;};_fdb .Set (_eefg ,_fbe );_cea .Log .Trace ("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073",_eefg ,_fbe .String ());};_cea .Log .Trace ("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021");
return _fdb ,nil ;};