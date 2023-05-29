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

package arithmetic ;import (_g "bytes";_a "github.com/unidoc/unipdf/v3/common";_ag "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_de "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_f "io";);func (_bef *Encoder )renormalize (){for {_bef ._dce <<=1;
_bef ._ea <<=1;_bef ._gc --;if _bef ._gc ==0{_bef .byteOut ();};if (_bef ._dce &0x8000)!=0{break ;};};};var _b =[]intEncRangeS {{0,3,0,2,0,2},{-1,-1,9,4,0,0},{-3,-2,5,3,2,1},{4,19,2,3,4,4},{-19,-4,3,3,4,4},{20,83,6,4,20,6},{-83,-20,7,4,20,6},{84,339,14,5,84,8},{-339,-84,15,5,84,8},{340,4435,30,6,340,12},{-4435,-340,31,6,340,12},{4436,2000000000,62,6,4436,32},{-2000000000,-4436,63,6,4436,32}};
func (_aea *Encoder )codeLPS (_ebb *codingContext ,_gga uint32 ,_bbg uint16 ,_eee byte ){_aea ._dce -=_bbg ;if _aea ._dce < _bbg {_aea ._ea +=uint32 (_bbg );}else {_aea ._dce =_bbg ;};if _fed [_eee ]._fab ==1{_ebb .flipMps (_gga );};_ebb ._e [_gga ]=_fed [_eee ]._fcg ;
_aea .renormalize ();};func (_eg *Encoder )encodeBit (_dag *codingContext ,_eaa uint32 ,_cdc uint8 )error {const _ebf ="\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074";_eg ._bf ++;if _eaa >=uint32 (len (_dag ._e )){return _de .Errorf (_ebf ,"\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027",_eaa );
};_aacb :=_dag ._e [_eaa ];_cda :=_dag .mps (_eaa );_gfg :=_fed [_aacb ]._aacbg ;_a .Log .Trace ("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064",_eg ._bf ,_cdc ,_aacb ,_cda ,_gfg ,_eg ._dce ,_eg ._ea ,_eg ._gc ,_eg ._aga ,_eg ._cac );
if _cdc ==0{_eg .code0 (_dag ,_eaa ,_gfg ,_aacb );}else {_eg .code1 (_dag ,_eaa ,_gfg ,_aacb );};return nil ;};const (_aec =65536;_bcgc =20*1024;);func (_da *Encoder )Reset (){_da ._dce =0x8000;_da ._ea =0;_da ._gc =12;_da ._cac =-1;_da ._aga =0;_da ._agae =nil ;
_da ._ae =_be (_aec );};func (_bcg *Encoder )codeMPS (_fee *codingContext ,_dde uint32 ,_bea uint16 ,_fag byte ){_bcg ._dce -=_bea ;if _bcg ._dce &0x8000!=0{_bcg ._ea +=uint32 (_bea );return ;};if _bcg ._dce < _bea {_bcg ._dce =_bea ;}else {_bcg ._ea +=uint32 (_bea );
};_fee ._e [_dde ]=_fed [_fag ]._edb ;_bcg .renormalize ();};func (_cf *Encoder )encodeIAID (_bdd ,_gfab int )error {if _cf ._agae ==nil {_cf ._agae =_be (1<<uint (_bdd ));};_gde :=uint32 (1<<uint32 (_bdd +1))-1;_gfab <<=uint (32-_bdd );_faf :=uint32 (1);
for _cfg :=0;_cfg < _bdd ;_cfg ++{_dfc :=_faf &_gde ;_ggd :=uint8 ((uint32 (_gfab )&0x80000000)>>31);if _fcd :=_cf .encodeBit (_cf ._agae ,_dfc ,_ggd );_fcd !=nil {return _fcd ;};_faf =(_faf <<1)|uint32 (_ggd );_gfab <<=1;};return nil ;};type codingContext struct{_e []byte ;
_ca []byte ;};func (_aee *Encoder )Init (){_aee ._ae =_be (_aec );_aee ._dce =0x8000;_aee ._ea =0;_aee ._gc =12;_aee ._cac =-1;_aee ._aga =0;_aee ._cb =0;_aee ._gb =make ([]byte ,_bcgc );for _fd :=0;_fd < len (_aee ._adf );_fd ++{_aee ._adf [_fd ]=_be (512);
};_aee ._agae =nil ;};const (IAAI Class =iota ;IADH ;IADS ;IADT ;IADW ;IAEX ;IAFS ;IAIT ;IARDH ;IARDW ;IARDX ;IARDY ;IARI ;);func (_afb *Encoder )flush (){_afb .setBits ();_afb ._ea <<=_afb ._gc ;_afb .byteOut ();_afb ._ea <<=_afb ._gc ;_afb .byteOut ();
_afb .emit ();if _afb ._aga !=0xff{_afb ._cac ++;_afb ._aga =0xff;_afb .emit ();};_afb ._cac ++;_afb ._aga =0xac;_afb ._cac ++;_afb .emit ();};func (_ff *codingContext )mps (_gg uint32 )int {return int (_ff ._ca [_gg ])};func (_aaf *Encoder )rBlock (){if _aaf ._cac >=0{_aaf .emit ();
};_aaf ._cac ++;_aaf ._aga =uint8 (_aaf ._ea >>20);_aaf ._ea &=0xfffff;_aaf ._gc =7;};func (_ebbd *Encoder )emit (){if _ebbd ._cb ==_bcgc {_ebbd ._gf =append (_ebbd ._gf ,_ebbd ._gb );_ebbd ._gb =make ([]byte ,_bcgc );_ebbd ._cb =0;};_ebbd ._gb [_ebbd ._cb ]=_ebbd ._aga ;
_ebbd ._cb ++;};func (_c Class )String ()string {switch _c {case IAAI :return "\u0049\u0041\u0041\u0049";case IADH :return "\u0049\u0041\u0044\u0048";case IADS :return "\u0049\u0041\u0044\u0053";case IADT :return "\u0049\u0041\u0044\u0054";case IADW :return "\u0049\u0041\u0044\u0057";
case IAEX :return "\u0049\u0041\u0045\u0058";case IAFS :return "\u0049\u0041\u0046\u0053";case IAIT :return "\u0049\u0041\u0049\u0054";case IARDH :return "\u0049\u0041\u0052D\u0048";case IARDW :return "\u0049\u0041\u0052D\u0057";case IARDX :return "\u0049\u0041\u0052D\u0058";
case IARDY :return "\u0049\u0041\u0052D\u0059";case IARI :return "\u0049\u0041\u0052\u0049";default:return "\u0055N\u004b\u004e\u004f\u0057\u004e";};};func (_fa *Encoder )code0 (_aacf *codingContext ,_fbb uint32 ,_dd uint16 ,_bed byte ){if _aacf .mps (_fbb )==0{_fa .codeMPS (_aacf ,_fbb ,_dd ,_bed );
}else {_fa .codeLPS (_aacf ,_fbb ,_dd ,_bed );};};func (_dfd *Encoder )Final (){_dfd .flush ()};func (_eab *Encoder )encodeOOB (_eda Class )error {_gebb :=_eab ._adf [_eda ];_bcgf :=_eab .encodeBit (_gebb ,1,1);if _bcgf !=nil {return _bcgf ;};_bcgf =_eab .encodeBit (_gebb ,3,0);
if _bcgf !=nil {return _bcgf ;};_bcgf =_eab .encodeBit (_gebb ,6,0);if _bcgf !=nil {return _bcgf ;};_bcgf =_eab .encodeBit (_gebb ,12,0);if _bcgf !=nil {return _bcgf ;};return nil ;};func (_fce *Encoder )EncodeInteger (proc Class ,value int )(_geb error ){_a .Log .Trace ("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027",value ,proc );
if _geb =_fce .encodeInteger (proc ,value );_geb !=nil {return _de .Wrap (_geb ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072","");};return nil ;};func (_ed *Encoder )byteOut (){if _ed ._aga ==0xff{_ed .rBlock ();return ;
};if _ed ._ea < 0x8000000{_ed .lBlock ();return ;};_ed ._aga ++;if _ed ._aga !=0xff{_ed .lBlock ();return ;};_ed ._ea &=0x7ffffff;_ed .rBlock ();};func (_dgf *Encoder )Refine (iTemp ,iTarget *_ag .Bitmap ,ox ,oy int )error {for _cgg :=0;_cgg < iTarget .Height ;
_cgg ++{var _ec int ;_cbb :=_cgg +oy ;var (_dec ,_af ,_efd ,_fg ,_agg uint16 ;_gba ,_bfd ,_aega ,_aa ,_gea byte ;);if _cbb >=1&&(_cbb -1)< iTemp .Height {_gba =iTemp .Data [(_cbb -1)*iTemp .RowStride ];};if _cbb >=0&&_cbb < iTemp .Height {_bfd =iTemp .Data [_cbb *iTemp .RowStride ];
};if _cbb >=-1&&_cbb +1< iTemp .Height {_aega =iTemp .Data [(_cbb +1)*iTemp .RowStride ];};if _cgg >=1{_aa =iTarget .Data [(_cgg -1)*iTarget .RowStride ];};_gea =iTarget .Data [_cgg *iTarget .RowStride ];_gfbg :=uint (6+ox );_dec =uint16 (_gba >>_gfbg );
_af =uint16 (_bfd >>_gfbg );_efd =uint16 (_aega >>_gfbg );_fg =uint16 (_aa >>6);_bd :=uint (2-ox );_gba <<=_bd ;_bfd <<=_bd ;_aega <<=_bd ;_aa <<=2;for _ec =0;_ec < iTarget .Width ;_ec ++{_fb :=(_dec <<10)|(_af <<7)|(_efd <<4)|(_fg <<1)|_agg ;_ebd :=_gea >>7;
_eff :=_dgf .encodeBit (_dgf ._ae ,uint32 (_fb ),_ebd );if _eff !=nil {return _eff ;};_dec <<=1;_af <<=1;_efd <<=1;_fg <<=1;_dec |=uint16 (_gba >>7);_af |=uint16 (_bfd >>7);_efd |=uint16 (_aega >>7);_fg |=uint16 (_aa >>7);_agg =uint16 (_ebd );_afd :=_ec %8;
_ffc :=_ec /8+1;if _afd ==5+ox {_gba ,_bfd ,_aega =0,0,0;if _ffc < iTemp .RowStride &&_cbb >=1&&(_cbb -1)< iTemp .Height {_gba =iTemp .Data [(_cbb -1)*iTemp .RowStride +_ffc ];};if _ffc < iTemp .RowStride &&_cbb >=0&&_cbb < iTemp .Height {_bfd =iTemp .Data [_cbb *iTemp .RowStride +_ffc ];
};if _ffc < iTemp .RowStride &&_cbb >=-1&&(_cbb +1)< iTemp .Height {_aega =iTemp .Data [(_cbb +1)*iTemp .RowStride +_ffc ];};}else {_gba <<=1;_bfd <<=1;_aega <<=1;};if _afd ==5&&_cgg >=1{_aa =0;if _ffc < iTarget .RowStride {_aa =iTarget .Data [(_cgg -1)*iTarget .RowStride +_ffc ];
};}else {_aa <<=1;};if _afd ==7{_gea =0;if _ffc < iTarget .RowStride {_gea =iTarget .Data [_cgg *iTarget .RowStride +_ffc ];};}else {_gea <<=1;};_dec &=7;_af &=7;_efd &=7;_fg &=7;};};return nil ;};var _ _f .WriterTo =&Encoder {};const _dg =0x9b25;var _fed =[]state {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
func (_gebd *Encoder )EncodeOOB (proc Class )(_efb error ){_a .Log .Trace ("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027",proc );if _efb =_gebd .encodeOOB (proc );
_efb !=nil {return _de .Wrap (_efb ,"\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB","");};return nil ;};func (_db *Encoder )encodeInteger (_caf Class ,_aeag int )error {const _acb ="E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072";
if _aeag > 2000000000||_aeag < -2000000000{return _de .Errorf (_acb ,"\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",_aeag );
};_efg :=_db ._adf [_caf ];_gbe :=uint32 (1);var _dbe int ;for ;;_dbe ++{if _b [_dbe ]._ge <=_aeag &&_b [_dbe ]._dc >=_aeag {break ;};};if _aeag < 0{_aeag =-_aeag ;};_aeag -=int (_b [_dbe ]._cg );_add :=_b [_dbe ]._ce ;for _afdd :=uint8 (0);_afdd < _b [_dbe ]._ad ;
_afdd ++{_bdb :=_add &1;if _fda :=_db .encodeBit (_efg ,_gbe ,_bdb );_fda !=nil {return _de .Wrap (_fda ,_acb ,"");};_add >>=1;if _gbe &0x100> 0{_gbe =(((_gbe <<1)|uint32 (_bdb ))&0x1ff)|0x100;}else {_gbe =(_gbe <<1)|uint32 (_bdb );};};_aeag <<=32-_b [_dbe ]._fc ;
for _ccd :=uint8 (0);_ccd < _b [_dbe ]._fc ;_ccd ++{_dad :=uint8 ((uint32 (_aeag )&0x80000000)>>31);if _ecf :=_db .encodeBit (_efg ,_gbe ,_dad );_ecf !=nil {return _de .Wrap (_ecf ,_acb ,"\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064");
};_aeag <<=1;if _gbe &0x100!=0{_gbe =(((_gbe <<1)|uint32 (_dad ))&0x1ff)|0x100;}else {_gbe =(_gbe <<1)|uint32 (_dad );};};return nil ;};func (_egg *Encoder )setBits (){_ace :=_egg ._ea +uint32 (_egg ._dce );_egg ._ea |=0xffff;if _egg ._ea >=_ace {_egg ._ea -=0x8000;
};};func _be (_cc int )*codingContext {return &codingContext {_e :make ([]byte ,_cc ),_ca :make ([]byte ,_cc )};};func (_gdc *Encoder )dataSize ()int {return _bcgc *len (_gdc ._gf )+_gdc ._cb };func (_cag *Encoder )EncodeBitmap (bm *_ag .Bitmap ,duplicateLineRemoval bool )error {_a .Log .Trace ("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s",bm .Width ,bm .Height ,bm );
var (_ccf ,_fff uint8 ;_cef ,_bb ,_gd uint16 ;_ac ,_ef ,_fdd byte ;_aeec ,_cd ,_geg int ;_gfb ,_eb []byte ;);for _aeg :=0;_aeg < bm .Height ;_aeg ++{_ac ,_ef =0,0;if _aeg >=2{_ac =bm .Data [(_aeg -2)*bm .RowStride ];};if _aeg >=1{_ef =bm .Data [(_aeg -1)*bm .RowStride ];
if duplicateLineRemoval {_cd =_aeg *bm .RowStride ;_gfb =bm .Data [_cd :_cd +bm .RowStride ];_geg =(_aeg -1)*bm .RowStride ;_eb =bm .Data [_geg :_geg +bm .RowStride ];if _g .Equal (_gfb ,_eb ){_fff =_ccf ^1;_ccf =1;}else {_fff =_ccf ;_ccf =0;};};};if duplicateLineRemoval {if _bbd :=_cag .encodeBit (_cag ._ae ,_dg ,_fff );
_bbd !=nil {return _bbd ;};if _ccf !=0{continue ;};};_fdd =bm .Data [_aeg *bm .RowStride ];_cef =uint16 (_ac >>5);_bb =uint16 (_ef >>4);_ac <<=3;_ef <<=4;_gd =0;for _aeec =0;_aeec < bm .Width ;_aeec ++{_aef :=uint32 (_cef <<11|_bb <<4|_gd );_fe :=(_fdd &0x80)>>7;
_df :=_cag .encodeBit (_cag ._ae ,_aef ,_fe );if _df !=nil {return _df ;};_cef <<=1;_bb <<=1;_gd <<=1;_cef |=uint16 ((_ac &0x80)>>7);_bb |=uint16 ((_ef &0x80)>>7);_gd |=uint16 (_fe );_gab :=_aeec %8;_bec :=_aeec /8+1;if _gab ==4&&_aeg >=2{_ac =0;if _bec < bm .RowStride {_ac =bm .Data [(_aeg -2)*bm .RowStride +_bec ];
};}else {_ac <<=1;};if _gab ==3&&_aeg >=1{_ef =0;if _bec < bm .RowStride {_ef =bm .Data [(_aeg -1)*bm .RowStride +_bec ];};}else {_ef <<=1;};if _gab ==7{_fdd =0;if _bec < bm .RowStride {_fdd =bm .Data [_aeg *bm .RowStride +_bec ];};}else {_fdd <<=1;};_cef &=31;
_bb &=127;_gd &=15;};};return nil ;};func New ()*Encoder {_ceb :=&Encoder {};_ceb .Init ();return _ceb };func (_ga *codingContext )flipMps (_cab uint32 ){_ga ._ca [_cab ]=1-_ga ._ca [_cab ]};func (_gbc *Encoder )DataSize ()int {return _gbc .dataSize ()};
type intEncRangeS struct{_ge ,_dc int ;_ce ,_ad uint8 ;_cg uint16 ;_fc uint8 ;};type Encoder struct{_ea uint32 ;_dce uint16 ;_gc ,_aga uint8 ;_cac int ;_bf int ;_gf [][]byte ;_gb []byte ;_cb int ;_ae *codingContext ;_adf [13]*codingContext ;_agae *codingContext ;
};func (_gda *Encoder )WriteTo (w _f .Writer )(int64 ,error ){const _bg ="\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f";var _bc int64 ;for _cggf ,_cdg :=range _gda ._gf {_def ,_dfde :=w .Write (_cdg );if _dfde !=nil {return 0,_de .Wrapf (_dfde ,_bg ,"\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b",_cggf );
};_bc +=int64 (_def );};_gda ._gb =_gda ._gb [:_gda ._cb ];_cdf ,_aac :=w .Write (_gda ._gb );if _aac !=nil {return 0,_de .Wrap (_aac ,_bg ,"\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073");};_bc +=int64 (_cdf );return _bc ,nil ;
};func (_gfa *Encoder )EncodeIAID (symbolCodeLength ,value int )(_ee error ){_a .Log .Trace ("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",symbolCodeLength ,value );
if _ee =_gfa .encodeIAID (symbolCodeLength ,value );_ee !=nil {return _de .Wrap (_ee ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044","");};return nil ;};type Class int ;func (_dcg *Encoder )code1 (_edf *codingContext ,_ab uint32 ,_bfa uint16 ,_acd byte ){if _edf .mps (_ab )==1{_dcg .codeMPS (_edf ,_ab ,_bfa ,_acd );
}else {_dcg .codeLPS (_edf ,_ab ,_bfa ,_acd );};};type state struct{_aacbg uint16 ;_edb ,_fcg uint8 ;_fab uint8 ;};func (_gge *Encoder )lBlock (){if _gge ._cac >=0{_gge .emit ();};_gge ._cac ++;_gge ._aga =uint8 (_gge ._ea >>19);_gge ._ea &=0x7ffff;_gge ._gc =8;
};func (_agd *Encoder )Flush (){_agd ._cb =0;_agd ._gf =nil ;_agd ._cac =-1};