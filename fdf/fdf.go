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
package fdf ;import (_f "bufio";_ga "bytes";_ac "encoding/hex";_e "errors";_gcd "fmt";_ef "github.com/unidoc/unipdf/v3/common";_ed "github.com/unidoc/unipdf/v3/core";_dc "io";_aa "os";_d "regexp";_c "sort";_a "strconv";_gc "strings";);var _bd =_d .MustCompile ("\u0025\u0025\u0045O\u0046");


// Data represents forms data format (FDF) file data.
type Data struct{_gb *_ed .PdfObjectDictionary ;_dca *_ed .PdfObjectArray ;};var _dcc =_d .MustCompile ("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029");func _agd (_bag string )(*fdfParser ,error ){_ccbf :=fdfParser {};
_fcg :=[]byte (_bag );_eeeb :=_ga .NewReader (_fcg );_ccbf ._ecc =_eeeb ;_ccbf ._ab =map[int64 ]_ed .PdfObject {};_egg :=_f .NewReader (_eeeb );_ccbf ._ffb =_egg ;_ccbf ._ccf =int64 (len (_bag ));return &_ccbf ,_ccbf .parse ();};func (_gdf *fdfParser )parseNull ()(_ed .PdfObjectNull ,error ){_ ,_cga :=_gdf ._ffb .Discard (4);
return _ed .PdfObjectNull {},_cga ;};var _fb =_d .MustCompile ("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029");func _bec (_edd _dc .ReadSeeker )(*fdfParser ,error ){_ddd :=&fdfParser {};_ddd ._ecc =_edd ;
_ddd ._ab =map[int64 ]_ed .PdfObject {};_cgdd ,_ddb ,_cge :=_ddd .parseFdfVersion ();if _cge !=nil {_ef .Log .Error ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076",_cge );
return nil ,_cge ;};_ddd ._fd =_cgdd ;_ddd ._eb =_ddb ;_cge =_ddd .parse ();return _ddd ,_cge ;};type fdfParser struct{_fd int ;_eb int ;_ab map[int64 ]_ed .PdfObject ;_ecc _dc .ReadSeeker ;_ffb *_f .Reader ;_ccf int64 ;_gbf *_ed .PdfObjectDictionary ;
};func _bccb (_efc string )(_ed .PdfObjectReference ,error ){_gcb :=_ed .PdfObjectReference {};_dab :=_ce .FindStringSubmatch (_efc );if len (_dab )< 3{_ef .Log .Debug ("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065");
return _gcb ,_e .New ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e");};_edf ,_ea :=_a .Atoi (_dab [1]);if _ea !=nil {_ef .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030",_dab [1]);
return _gcb ,nil ;};_gcb .ObjectNumber =int64 (_edf );_fbc ,_ea :=_a .Atoi (_dab [2]);if _ea !=nil {_ef .Log .Debug ("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030",_dab [2]);
return _gcb ,nil ;};_gcb .GenerationNumber =int64 (_fbc );return _gcb ,nil ;};func (_dbd *fdfParser )readComment ()(string ,error ){var _de _ga .Buffer ;_ ,_geb :=_dbd .skipSpaces ();if _geb !=nil {return _de .String (),_geb ;};_efd :=true ;for {_dba ,_gcef :=_dbd ._ffb .Peek (1);
if _gcef !=nil {_ef .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_gcef .Error ());return _de .String (),_gcef ;};if _efd &&_dba [0]!='%'{return _de .String (),_e .New ("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025");
};_efd =false ;if (_dba [0]!='\r')&&(_dba [0]!='\n'){_gfg ,_ :=_dbd ._ffb .ReadByte ();_de .WriteByte (_gfg );}else {break ;};};return _de .String (),nil ;};func (_eddc *fdfParser )trace (_efbf _ed .PdfObject )_ed .PdfObject {switch _gdd :=_efbf .(type ){case *_ed .PdfObjectReference :_fcf ,_ecce :=_eddc ._ab [_gdd .ObjectNumber ].(*_ed .PdfIndirectObject );
if _ecce {return _fcf .PdfObject ;};_ef .Log .Debug ("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");return nil ;case *_ed .PdfIndirectObject :return _gdd .PdfObject ;};return _efbf ;};func (_cg *fdfParser )readAtLeast (_ba []byte ,_bb int )(int ,error ){_ff :=_bb ;
_eg :=0;_dbg :=0;for _ff > 0{_bg ,_fag :=_cg ._ffb .Read (_ba [_eg :]);if _fag !=nil {_ef .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073",_bg ,_dbg ,_fag .Error ());
return _eg ,_e .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067");};_dbg ++;_eg +=_bg ;_ff -=_bg ;};return _eg ,nil ;};func (_af *fdfParser )setFileOffset (_dcb int64 ){_af ._ecc .Seek (_dcb ,_dc .SeekStart );_af ._ffb =_f .NewReader (_af ._ecc );
};func (_fbee *fdfParser )parseDict ()(*_ed .PdfObjectDictionary ,error ){_ef .Log .Trace ("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021");_afe :=_ed .MakeDict ();_baa ,_ :=_fbee ._ffb .ReadByte ();if _baa !='<'{return nil ,_e .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");
};_baa ,_ =_fbee ._ffb .ReadByte ();if _baa !='<'{return nil ,_e .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};for {_fbee .skipSpaces ();_fbee .skipComments ();_faec ,_eeb :=_fbee ._ffb .Peek (2);if _eeb !=nil {return nil ,_eeb ;
};_ef .Log .Trace ("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_faec ),string (_faec ));if (_faec [0]=='>')&&(_faec [1]=='>'){_ef .Log .Trace ("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");
_fbee ._ffb .ReadByte ();_fbee ._ffb .ReadByte ();break ;};_ef .Log .Trace ("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021");_dfb ,_eeb :=_fbee .parseName ();_ef .Log .Trace ("\u004be\u0079\u003a\u0020\u0025\u0073",_dfb );
if _eeb !=nil {_ef .Log .Debug ("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073",_eeb );return nil ,_eeb ;};if len (_dfb )> 4&&_dfb [len (_dfb )-4:]=="\u006e\u0075\u006c\u006c"{_fce :=_dfb [0:len (_dfb )-4];
_ef .Log .Debug ("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029",_dfb );_ef .Log .Debug ("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c",_fce );
_fbee .skipSpaces ();_cdc ,_ :=_fbee ._ffb .Peek (1);if _cdc [0]=='/'{_afe .Set (_fce ,_ed .MakeNull ());continue ;};};_fbee .skipSpaces ();_decg ,_eeb :=_fbee .parseObject ();if _eeb !=nil {return nil ,_eeb ;};_afe .Set (_dfb ,_decg );_ef .Log .Trace ("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073",_dfb ,_decg .String ());
};_ef .Log .Trace ("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021");return _afe ,nil ;};func (_egd *fdfParser )parseHexString ()(*_ed .PdfObjectString ,error ){_egd ._ffb .ReadByte ();var _dccd _ga .Buffer ;
for {_fda ,_bac :=_egd ._ffb .Peek (1);if _bac !=nil {return _ed .MakeHexString (""),_bac ;};if _fda [0]=='>'{_egd ._ffb .ReadByte ();break ;};_bgd ,_ :=_egd ._ffb .ReadByte ();if !_ed .IsWhiteSpace (_bgd ){_dccd .WriteByte (_bgd );};};if _dccd .Len ()%2==1{_dccd .WriteRune ('0');
};_bcc ,_bbg :=_ac .DecodeString (_dccd .String ());if _bbg !=nil {_ef .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067",_dccd .String ());
return _ed .MakeHexString (""),nil ;};return _ed .MakeHexString (string (_bcc )),nil ;};func (_gbb *fdfParser )parseBool ()(_ed .PdfObjectBool ,error ){_bgdd ,_fcd :=_gbb ._ffb .Peek (4);if _fcd !=nil {return _ed .PdfObjectBool (false ),_fcd ;};if (len (_bgdd )>=4)&&(string (_bgdd [:4])=="\u0074\u0072\u0075\u0065"){_gbb ._ffb .Discard (4);
return _ed .PdfObjectBool (true ),nil ;};_bgdd ,_fcd =_gbb ._ffb .Peek (5);if _fcd !=nil {return _ed .PdfObjectBool (false ),_fcd ;};if (len (_bgdd )>=5)&&(string (_bgdd [:5])=="\u0066\u0061\u006cs\u0065"){_gbb ._ffb .Discard (5);return _ed .PdfObjectBool (false ),nil ;
};return _ed .PdfObjectBool (false ),_e .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");};

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath (fdfPath string )(*Data ,error ){_da ,_aca :=_aa .Open (fdfPath );if _aca !=nil {return nil ,_aca ;};defer _da .Close ();return Load (_da );};func (_dedf *fdfParser )parseArray ()(*_ed .PdfObjectArray ,error ){_dbgd :=_ed .MakeArray ();
_dedf ._ffb .ReadByte ();for {_dedf .skipSpaces ();_ag ,_efff :=_dedf ._ffb .Peek (1);if _efff !=nil {return _dbgd ,_efff ;};if _ag [0]==']'{_dedf ._ffb .ReadByte ();break ;};_fdc ,_efff :=_dedf .parseObject ();if _efff !=nil {return _dbgd ,_efff ;};_dbgd .Append (_fdc );
};return _dbgd ,nil ;};func (_ecf *fdfParser )parse ()error {_ecf ._ecc .Seek (0,_dc .SeekStart );_ecf ._ffb =_f .NewReader (_ecf ._ecc );for {_ecf .skipComments ();_cgee ,_beae :=_ecf ._ffb .Peek (20);if _beae !=nil {_ef .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");
return _beae ;};if _gc .HasPrefix (string (_cgee ),"\u0074r\u0061\u0069\u006c\u0065\u0072"){_ecf ._ffb .Discard (7);_ecf .skipSpaces ();_ecf .skipComments ();_faee ,_ :=_ecf .parseDict ();_ecf ._gbf =_faee ;break ;};_dbc :=_acd .FindStringSubmatchIndex (string (_cgee ));
if len (_dbc )< 6{_ef .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_cgee ));
return _e .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_caf ,_beae :=_ecf .parseIndirectObject ();if _beae !=nil {return _beae ;};switch _eggb :=_caf .(type ){case *_ed .PdfIndirectObject :_ecf ._ab [_eggb .ObjectNumber ]=_eggb ;case *_ed .PdfObjectStream :_ecf ._ab [_eggb .ObjectNumber ]=_eggb ;default:return _e .New ("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};};return nil ;};var _ce =_d .MustCompile ("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052");func (_bda *fdfParser )parseName ()(_ed .PdfObjectName ,error ){var _fdgb _ga .Buffer ;_ade :=false ;for {_fe ,_ee :=_bda ._ffb .Peek (1);
if _ee ==_dc .EOF {break ;};if _ee !=nil {return _ed .PdfObjectName (_fdgb .String ()),_ee ;};if !_ade {if _fe [0]=='/'{_ade =true ;_bda ._ffb .ReadByte ();}else if _fe [0]=='%'{_bda .readComment ();_bda .skipSpaces ();}else {_ef .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029",_fe ,_fe );
return _ed .PdfObjectName (_fdgb .String ()),_gcd .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029",_fe [0]);};}else {if _ed .IsWhiteSpace (_fe [0]){break ;}else if (_fe [0]=='/')||(_fe [0]=='[')||(_fe [0]=='(')||(_fe [0]==']')||(_fe [0]=='<')||(_fe [0]=='>'){break ;
}else if _fe [0]=='#'{_cb ,_fbe :=_bda ._ffb .Peek (3);if _fbe !=nil {return _ed .PdfObjectName (_fdgb .String ()),_fbe ;};_bda ._ffb .Discard (3);_fff ,_fbe :=_ac .DecodeString (string (_cb [1:3]));if _fbe !=nil {return _ed .PdfObjectName (_fdgb .String ()),_fbe ;
};_fdgb .Write (_fff );}else {_ebg ,_ :=_bda ._ffb .ReadByte ();_fdgb .WriteByte (_ebg );};};};return _ed .PdfObjectName (_fdgb .String ()),nil ;};func (_acb *fdfParser )getFileOffset ()int64 {_cgg ,_ :=_acb ._ecc .Seek (0,_dc .SeekCurrent );_cgg -=int64 (_acb ._ffb .Buffered ());
return _cgg ;};func (_gea *fdfParser )parseString ()(*_ed .PdfObjectString ,error ){_gea ._ffb .ReadByte ();var _fbd _ga .Buffer ;_dee :=1;for {_bdg ,_cba :=_gea ._ffb .Peek (1);if _cba !=nil {return _ed .MakeString (_fbd .String ()),_cba ;};if _bdg [0]=='\\'{_gea ._ffb .ReadByte ();
_daa ,_bdf :=_gea ._ffb .ReadByte ();if _bdf !=nil {return _ed .MakeString (_fbd .String ()),_bdf ;};if _ed .IsOctalDigit (_daa ){_aaf ,_ded :=_gea ._ffb .Peek (2);if _ded !=nil {return _ed .MakeString (_fbd .String ()),_ded ;};var _cgd []byte ;_cgd =append (_cgd ,_daa );
for _ ,_cee :=range _aaf {if _ed .IsOctalDigit (_cee ){_cgd =append (_cgd ,_cee );}else {break ;};};_gea ._ffb .Discard (len (_cgd )-1);_ef .Log .Trace ("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022",_cgd );
_ffg ,_ded :=_a .ParseUint (string (_cgd ),8,32);if _ded !=nil {return _ed .MakeString (_fbd .String ()),_ded ;};_fbd .WriteByte (byte (_ffg ));continue ;};switch _daa {case 'n':_fbd .WriteRune ('\n');case 'r':_fbd .WriteRune ('\r');case 't':_fbd .WriteRune ('\t');
case 'b':_fbd .WriteRune ('\b');case 'f':_fbd .WriteRune ('\f');case '(':_fbd .WriteRune ('(');case ')':_fbd .WriteRune (')');case '\\':_fbd .WriteRune ('\\');};continue ;}else if _bdg [0]=='('{_dee ++;}else if _bdg [0]==')'{_dee --;if _dee ==0{_gea ._ffb .ReadByte ();
break ;};};_bf ,_ :=_gea ._ffb .ReadByte ();_fbd .WriteByte (_bf );};return _ed .MakeString (_fbd .String ()),nil ;};func (_dfd *fdfParser )parseNumber ()(_ed .PdfObject ,error ){return _ed .ParseNumber (_dfd ._ffb )};func (_cdg *fdfParser )seekToEOFMarker (_bbe int64 )error {_gcfg :=int64 (0);
_dcf :=int64 (1000);for _gcfg < _bbe {if _bbe <=(_dcf +_gcfg ){_dcf =_bbe -_gcfg ;};_ ,_fagb :=_cdg ._ecc .Seek (-_gcfg -_dcf ,_dc .SeekEnd );if _fagb !=nil {return _fagb ;};_cbe :=make ([]byte ,_dcf );_cdg ._ecc .Read (_cbe );_ef .Log .Trace ("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022",string (_cbe ));
_eac :=_bd .FindAllStringIndex (string (_cbe ),-1);if _eac !=nil {_def :=_eac [len (_eac )-1];_ef .Log .Trace ("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064",_eac );_cdg ._ecc .Seek (-_gcfg -_dcf +int64 (_def [0]),_dc .SeekEnd );return nil ;};_ef .Log .Debug ("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067");
_gcfg +=_dcf ;};_ef .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e");return _e .New ("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");
};func (_aac *fdfParser )seekFdfVersionTopDown ()(int ,int ,error ){_aac ._ecc .Seek (0,_dc .SeekStart );_aac ._ffb =_f .NewReader (_aac ._ecc );_ffbf :=20;_dde :=make ([]byte ,_ffbf );for {_dac ,_abg :=_aac ._ffb .ReadByte ();if _abg !=nil {if _abg ==_dc .EOF {break ;
}else {return 0,0,_abg ;};};if _ed .IsDecimalDigit (_dac )&&_dde [_ffbf -1]=='.'&&_ed .IsDecimalDigit (_dde [_ffbf -2])&&_dde [_ffbf -3]=='-'&&_dde [_ffbf -4]=='F'&&_dde [_ffbf -5]=='D'&&_dde [_ffbf -6]=='P'{_bbcc :=int (_dde [_ffbf -2]-'0');_cfce :=int (_dac -'0');
return _bbcc ,_cfce ,nil ;};_dde =append (_dde [1:_ffbf ],_dac );};return 0,0,_e .New ("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};func (_bdc *fdfParser )skipSpaces ()(int ,error ){_bceb :=0;for {_gac ,_gad :=_bdc ._ffb .ReadByte ();
if _gad !=nil {return 0,_gad ;};if _ed .IsWhiteSpace (_gac ){_bceb ++;}else {_bdc ._ffb .UnreadByte ();break ;};};return _bceb ,nil ;};

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data )FieldDictionaries ()(map[string ]*_ed .PdfObjectDictionary ,error ){_dg :=map[string ]*_ed .PdfObjectDictionary {};for _bc :=0;_bc < fdf ._dca .Len ();_bc ++{_aaa ,_edb :=_ed .GetDict (fdf ._dca .Get (_bc ));if _edb {_gd ,_ :=_ed .GetString (_aaa .Get ("\u0054"));
if _gd !=nil {_dg [_gd .Str ()]=_aaa ;};};};return _dg ,nil ;};

// Root returns the Root of the FDF document.
func (_agb *fdfParser )Root ()(*_ed .PdfObjectDictionary ,error ){if _agb ._gbf !=nil {if _cfe ,_ggb :=_agb .trace (_agb ._gbf .Get ("\u0052\u006f\u006f\u0074")).(*_ed .PdfObjectDictionary );_ggb {if _daf ,_acc :=_agb .trace (_cfe .Get ("\u0046\u0044\u0046")).(*_ed .PdfObjectDictionary );
_acc {return _daf ,nil ;};};};var _daaa []int64 ;for _bede :=range _agb ._ab {_daaa =append (_daaa ,_bede );};_c .Slice (_daaa ,func (_agg ,_fdbc int )bool {return _daaa [_agg ]< _daaa [_fdbc ]});for _ ,_ca :=range _daaa {_dad :=_agb ._ab [_ca ];if _fbea ,_gdb :=_agb .trace (_dad ).(*_ed .PdfObjectDictionary );
_gdb {if _cfea ,_faed :=_agb .trace (_fbea .Get ("\u0046\u0044\u0046")).(*_ed .PdfObjectDictionary );_faed {return _cfea ,nil ;};};};return nil ,_e .New ("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};var _efa =_d .MustCompile ("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029");
func (_aafc *fdfParser )parseFdfVersion ()(int ,int ,error ){_aafc ._ecc .Seek (0,_dc .SeekStart );_bcd :=20;_bab :=make ([]byte ,_bcd );_aafc ._ecc .Read (_bab );_gcae :=_dcc .FindStringSubmatch (string (_bab ));if len (_gcae )< 3{_bgcd ,_aaee ,_dae :=_aafc .seekFdfVersionTopDown ();
if _dae !=nil {_ef .Log .Debug ("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e");return 0,0,_dae ;
};return _bgcd ,_aaee ,nil ;};_afa ,_cbc :=_a .Atoi (_gcae [1]);if _cbc !=nil {return 0,0,_cbc ;};_gec ,_cbc :=_a .Atoi (_gcae [2]);if _cbc !=nil {return 0,0,_cbc ;};_ef .Log .Debug ("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064",_afa ,_gec );
return _afa ,_gec ,nil ;};func (_bed *fdfParser )parseIndirectObject ()(_ed .PdfObject ,error ){_eda :=_ed .PdfIndirectObject {};_ef .Log .Trace ("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a");_ceb ,_ebc :=_bed ._ffb .Peek (20);
if _ebc !=nil {_ef .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");return &_eda ,_ebc ;};_ef .Log .Trace ("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022",string (_ceb ));
_gagc :=_acd .FindStringSubmatchIndex (string (_ceb ));if len (_gagc )< 6{_ef .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_ceb ));
return &_eda ,_e .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_bed ._ffb .Discard (_gagc [0]);_ef .Log .Trace ("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064",_gagc );_fgc :=_gagc [1]-_gagc [0];_bacg :=make ([]byte ,_fgc );_ ,_ebc =_bed .readAtLeast (_bacg ,_fgc );if _ebc !=nil {_ef .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073",_ebc );
return nil ,_ebc ;};_ef .Log .Trace ("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073",_bacg );_efb :=_acd .FindStringSubmatch (string (_bacg ));if len (_efb )< 3{_ef .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_bacg ));
return &_eda ,_e .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_ebcf ,_ :=_a .Atoi (_efb [1]);_fage ,_ :=_a .Atoi (_efb [2]);_eda .ObjectNumber =int64 (_ebcf );_eda .GenerationNumber =int64 (_fage );for {_cgc ,_egc :=_bed ._ffb .Peek (2);if _egc !=nil {return &_eda ,_egc ;};_ef .Log .Trace ("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_cgc ),string (_cgc ));
if _ed .IsWhiteSpace (_cgc [0]){_bed .skipSpaces ();}else if _cgc [0]=='%'{_bed .skipComments ();}else if (_cgc [0]=='<')&&(_cgc [1]=='<'){_ef .Log .Trace ("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074");_eda .PdfObject ,_egc =_bed .parseDict ();
_ef .Log .Trace ("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076",_egc );if _egc !=nil {return &_eda ,_egc ;};_ef .Log .Trace ("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e");
}else if (_cgc [0]=='/')||(_cgc [0]=='(')||(_cgc [0]=='[')||(_cgc [0]=='<'){_eda .PdfObject ,_egc =_bed .parseObject ();if _egc !=nil {return &_eda ,_egc ;};_ef .Log .Trace ("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e");
}else {if _cgc [0]=='e'{_bbf ,_cfc :=_bed .readTextLine ();if _cfc !=nil {return nil ,_cfc ;};if len (_bbf )>=6&&_bbf [0:6]=="\u0065\u006e\u0064\u006f\u0062\u006a"{break ;};}else if _cgc [0]=='s'{_cgc ,_ =_bed ._ffb .Peek (10);if string (_cgc [:6])=="\u0073\u0074\u0072\u0065\u0061\u006d"{_ggc :=6;
if len (_cgc )> 6{if _ed .IsWhiteSpace (_cgc [_ggc ])&&_cgc [_ggc ]!='\r'&&_cgc [_ggc ]!='\n'{_ef .Log .Debug ("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072");
_ggc ++;};if _cgc [_ggc ]=='\r'{_ggc ++;if _cgc [_ggc ]=='\n'{_ggc ++;};}else if _cgc [_ggc ]=='\n'{_ggc ++;};};_bed ._ffb .Discard (_ggc );_gab ,_ddf :=_eda .PdfObject .(*_ed .PdfObjectDictionary );if !_ddf {return nil ,_e .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};_ef .Log .Trace ("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073",_gab );_dbf ,_gbab :=_gab .Get ("\u004c\u0065\u006e\u0067\u0074\u0068").(*_ed .PdfObjectInteger );if !_gbab {return nil ,_e .New ("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072");
};_cdgd :=*_dbf ;if _cdgd < 0{return nil ,_e .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030");};if int64 (_cdgd )> _bed ._ccf {_ef .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065");
return nil ,_e .New ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065");};_aea :=make ([]byte ,_cdgd );
_ ,_egc =_bed .readAtLeast (_aea ,int (_cdgd ));if _egc !=nil {_ef .Log .Debug ("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058",len (_aea ),_aea );_ef .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_egc );
return nil ,_egc ;};_ecg :=_ed .PdfObjectStream {};_ecg .Stream =_aea ;_ecg .PdfObjectDictionary =_eda .PdfObject .(*_ed .PdfObjectDictionary );_ecg .ObjectNumber =_eda .ObjectNumber ;_ecg .GenerationNumber =_eda .GenerationNumber ;_bed .skipSpaces ();
_bed ._ffb .Discard (9);_bed .skipSpaces ();return &_ecg ,nil ;};};_eda .PdfObject ,_egc =_bed .parseObject ();return &_eda ,_egc ;};};_ef .Log .Trace ("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021");
return &_eda ,nil ;};func (_fdg *fdfParser )readTextLine ()(string ,error ){var _fc _ga .Buffer ;for {_gcf ,_adg :=_fdg ._ffb .Peek (1);if _adg !=nil {_ef .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_adg .Error ());return _fc .String (),_adg ;
};if (_gcf [0]!='\r')&&(_gcf [0]!='\n'){_adgb ,_ :=_fdg ._ffb .ReadByte ();_fc .WriteByte (_adgb );}else {break ;};};return _fc .String (),nil ;};var _acd =_d .MustCompile ("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a");
func (_fagc *fdfParser )parseObject ()(_ed .PdfObject ,error ){_ef .Log .Trace ("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074");_fagc .skipSpaces ();for {_bgc ,_fad :=_fagc ._ffb .Peek (2);if _fad !=nil {return nil ,_fad ;
};_ef .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_bgc ));if _bgc [0]=='/'{_bdd ,_ebf :=_fagc .parseName ();_ef .Log .Trace ("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027",_bdd );
return &_bdd ,_ebf ;}else if _bgc [0]=='('{_ef .Log .Trace ("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!");return _fagc .parseString ();}else if _bgc [0]=='['{_ef .Log .Trace ("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021");return _fagc .parseArray ();
}else if (_bgc [0]=='<')&&(_bgc [1]=='<'){_ef .Log .Trace ("\u002d>\u0044\u0069\u0063\u0074\u0021");return _fagc .parseDict ();}else if _bgc [0]=='<'{_ef .Log .Trace ("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021");return _fagc .parseHexString ();
}else if _bgc [0]=='%'{_fagc .readComment ();_fagc .skipSpaces ();}else {_ef .Log .Trace ("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f");_bgc ,_ =_fagc ._ffb .Peek (15);_bea :=string (_bgc );_ef .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_bea );
if (len (_bea )> 3)&&(_bea [:4]=="\u006e\u0075\u006c\u006c"){_bba ,_fae :=_fagc .parseNull ();return &_bba ,_fae ;}else if (len (_bea )> 4)&&(_bea [:5]=="\u0066\u0061\u006cs\u0065"){_bfb ,_fg :=_fagc .parseBool ();return &_bfb ,_fg ;}else if (len (_bea )> 3)&&(_bea [:4]=="\u0074\u0072\u0075\u0065"){_aae ,_beb :=_fagc .parseBool ();
return &_aae ,_beb ;};_gdfe :=_ce .FindStringSubmatch (_bea );if len (_gdfe )> 1{_bgc ,_ =_fagc ._ffb .ReadBytes ('R');_ef .Log .Trace ("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027",string (_bgc [:]));_baca ,_bgb :=_bccb (string (_bgc ));
return &_baca ,_bgb ;};_eef :=_fb .FindStringSubmatch (_bea );if len (_eef )> 1{_ef .Log .Trace ("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021");return _fagc .parseNumber ();};_eef =_efa .FindStringSubmatch (_bea );if len (_eef )> 1{_ef .Log .Trace ("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021");
_ef .Log .Trace ("\u0025\u0020\u0073",_eef );return _fagc .parseNumber ();};_ef .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)",_bea );return nil ,_e .New ("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e");
};};};func (_cd *fdfParser )skipComments ()error {if _ ,_fdb :=_cd .skipSpaces ();_fdb !=nil {return _fdb ;};_gce :=true ;for {_gg ,_ge :=_cd ._ffb .Peek (1);if _ge !=nil {_ef .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_ge .Error ());
return _ge ;};if _gce &&_gg [0]!='%'{return nil ;};_gce =false ;if (_gg [0]!='\r')&&(_gg [0]!='\n'){_cd ._ffb .ReadByte ();}else {break ;};};return _cd .skipComments ();};

// Load loads FDF form data from `r`.
func Load (r _dc .ReadSeeker )(*Data ,error ){_gba ,_eff :=_bec (r );if _eff !=nil {return nil ,_eff ;};_db ,_eff :=_gba .Root ();if _eff !=nil {return nil ,_eff ;};_cc ,_b :=_ed .GetArray (_db .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_b {return nil ,_e .New ("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067");
};return &Data {_dca :_cc ,_gb :_db },nil ;};

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data )FieldValues ()(map[string ]_ed .PdfObject ,error ){_fa ,_ad :=fdf .FieldDictionaries ();if _ad !=nil {return nil ,_ad ;};var _gbd []string ;for _bce :=range _fa {_gbd =append (_gbd ,_bce );};_c .Strings (_gbd );_be :=map[string ]_ed .PdfObject {};
for _ ,_dgd :=range _gbd {_cf :=_fa [_dgd ];_dd :=_ed .TraceToDirectObject (_cf .Get ("\u0056"));_be [_dgd ]=_dd ;};return _be ,nil ;};