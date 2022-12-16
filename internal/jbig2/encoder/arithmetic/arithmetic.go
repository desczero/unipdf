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

package arithmetic ;import (_c "bytes";_d "github.com/unidoc/unipdf/v3/common";_b "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_bb "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_e "io";);func (_dbd *Encoder )DataSize ()int {return _dbd .dataSize ()};
func (_dea *Encoder )Init (){_dea ._cc =_ad (_dac );_dea ._ecc =0x8000;_dea ._cf =0;_dea ._f =12;_dea ._ee =-1;_dea ._bbe =0;_dea ._gdg =0;_dea ._cfg =make ([]byte ,_aff );for _ef :=0;_ef < len (_dea ._agb );_ef ++{_dea ._agb [_ef ]=_ad (512);};_dea ._ge =nil ;
};func (_bc *codingContext )mps (_ba uint32 )int {return int (_bc ._ca [_ba ])};func (_dcg *Encoder )EncodeOOB (proc Class )(_fbe error ){_d .Log .Trace ("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027",proc );
if _fbe =_dcg .encodeOOB (proc );_fbe !=nil {return _bb .Wrap (_fbe ,"\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB","");};return nil ;};func (_fc *Encoder )Final (){_fc .flush ()};const (_dac =65536;_aff =20*1024;);func (_bec *Encoder )codeLPS (_eac *codingContext ,_afg uint32 ,_da uint16 ,_aec byte ){_bec ._ecc -=_da ;
if _bec ._ecc < _da {_bec ._cf +=uint32 (_da );}else {_bec ._ecc =_da ;};if _gab [_aec ]._faa ==1{_eac .flipMps (_afg );};_eac ._be [_afg ]=_gab [_aec ]._fbg ;_bec .renormalize ();};func New ()*Encoder {_gdb :=&Encoder {};_gdb .Init ();return _gdb };const (IAAI Class =iota ;
IADH ;IADS ;IADT ;IADW ;IAEX ;IAFS ;IAIT ;IARDH ;IARDW ;IARDX ;IARDY ;IARI ;);func (_ecf *Encoder )EncodeIAID (symbolCodeLength ,value int )(_ddd error ){_d .Log .Trace ("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",symbolCodeLength ,value );
if _ddd =_ecf .encodeIAID (symbolCodeLength ,value );_ddd !=nil {return _bb .Wrap (_ddd ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044","");};return nil ;};func (_cae *Encoder )encodeOOB (_bbcb Class )error {_dec :=_cae ._agb [_bbcb ];_ccd :=_cae .encodeBit (_dec ,1,1);
if _ccd !=nil {return _ccd ;};_ccd =_cae .encodeBit (_dec ,3,0);if _ccd !=nil {return _ccd ;};_ccd =_cae .encodeBit (_dec ,6,0);if _ccd !=nil {return _ccd ;};_ccd =_cae .encodeBit (_dec ,12,0);if _ccd !=nil {return _ccd ;};return nil ;};func _ad (_ag int )*codingContext {return &codingContext {_be :make ([]byte ,_ag ),_ca :make ([]byte ,_ag )};
};func (_bba *Encoder )codeMPS (_bfe *codingContext ,_dfd uint32 ,_bbc uint16 ,_geb byte ){_bba ._ecc -=_bbc ;if _bba ._ecc &0x8000!=0{_bba ._cf +=uint32 (_bbc );return ;};if _bba ._ecc < _bbc {_bba ._ecc =_bbc ;}else {_bba ._cf +=uint32 (_bbc );};_bfe ._be [_dfd ]=_gab [_geb ]._bfb ;
_bba .renormalize ();};func (_fa *Encoder )Flush (){_fa ._gdg =0;_fa ._dg =nil ;_fa ._ee =-1};func (_eg *Encoder )code1 (_ebf *codingContext ,_gee uint32 ,_ecd uint16 ,_gce byte ){if _ebf .mps (_gee )==1{_eg .codeMPS (_ebf ,_gee ,_ecd ,_gce );}else {_eg .codeLPS (_ebf ,_gee ,_ecd ,_gce );
};};func (_fgc *Encoder )dataSize ()int {return _aff *len (_fgc ._dg )+_fgc ._gdg };func (_gdf *Encoder )EncodeInteger (proc Class ,value int )(_gccd error ){_d .Log .Trace ("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027",value ,proc );
if _gccd =_gdf .encodeInteger (proc ,value );_gccd !=nil {return _bb .Wrap (_gccd ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072","");};return nil ;};type Class int ;func (_dae *Encoder )renormalize (){for {_dae ._ecc <<=1;
_dae ._cf <<=1;_dae ._f --;if _dae ._f ==0{_dae .byteOut ();};if (_dae ._ecc &0x8000)!=0{break ;};};};func (_gbgg *Encoder )code0 (_adce *codingContext ,_cd uint32 ,_gcb uint16 ,_beg byte ){if _adce .mps (_cd )==0{_gbgg .codeMPS (_adce ,_cd ,_gcb ,_beg );
}else {_gbgg .codeLPS (_adce ,_cd ,_gcb ,_beg );};};type state struct{_fcb uint16 ;_bfb ,_fbg uint8 ;_faa uint8 ;};var _ _e .WriterTo =&Encoder {};func (_cdf *Encoder )lBlock (){if _cdf ._ee >=0{_cdf .emit ();};_cdf ._ee ++;_cdf ._bbe =uint8 (_cdf ._cf >>19);
_cdf ._cf &=0x7ffff;_cdf ._f =8;};type codingContext struct{_be []byte ;_ca []byte ;};type Encoder struct{_cf uint32 ;_ecc uint16 ;_f ,_bbe uint8 ;_ee int ;_dbf int ;_dg [][]byte ;_cfg []byte ;_gdg int ;_cc *codingContext ;_agb [13]*codingContext ;_ge *codingContext ;
};func (_ae *Encoder )Reset (){_ae ._ecc =0x8000;_ae ._cf =0;_ae ._f =12;_ae ._ee =-1;_ae ._bbe =0;_ae ._ge =nil ;_ae ._cc =_ad (_dac );};func (_fdc *Encoder )encodeInteger (_aae Class ,_cab int )error {const _fcg ="E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072";
if _cab > 2000000000||_cab < -2000000000{return _bb .Errorf (_fcg ,"\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",_cab );
};_dgc :=_fdc ._agb [_aae ];_cca :=uint32 (1);var _cag int ;for ;;_cag ++{if _gd [_cag ]._bg <=_cab &&_gd [_cag ]._db >=_cab {break ;};};if _cab < 0{_cab =-_cab ;};_cab -=int (_gd [_cag ]._ec );_fge :=_gd [_cag ]._de ;for _cagg :=uint8 (0);_cagg < _gd [_cag ]._a ;
_cagg ++{_ff :=_fge &1;if _bbfa :=_fdc .encodeBit (_dgc ,_cca ,_ff );_bbfa !=nil {return _bb .Wrap (_bbfa ,_fcg ,"");};_fge >>=1;if _cca &0x100> 0{_cca =(((_cca <<1)|uint32 (_ff ))&0x1ff)|0x100;}else {_cca =(_cca <<1)|uint32 (_ff );};};_cab <<=32-_gd [_cag ]._dba ;
for _ddc :=uint8 (0);_ddc < _gd [_cag ]._dba ;_ddc ++{_dgb :=uint8 ((uint32 (_cab )&0x80000000)>>31);if _cg :=_fdc .encodeBit (_dgc ,_cca ,_dgb );_cg !=nil {return _bb .Wrap (_cg ,_fcg ,"\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064");
};_cab <<=1;if _cca &0x100!=0{_cca =(((_cca <<1)|uint32 (_dgb ))&0x1ff)|0x100;}else {_cca =(_cca <<1)|uint32 (_dgb );};};return nil ;};const _fd =0x9b25;func (_bag *Encoder )WriteTo (w _e .Writer )(int64 ,error ){const _ea ="\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f";
var _ab int64 ;for _cbd ,_cbe :=range _bag ._dg {_ac ,_dbe :=w .Write (_cbe );if _dbe !=nil {return 0,_bb .Wrapf (_dbe ,_ea ,"\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b",_cbd );
};_ab +=int64 (_ac );};_bag ._cfg =_bag ._cfg [:_bag ._gdg ];_af ,_dfg :=w .Write (_bag ._cfg );if _dfg !=nil {return 0,_bb .Wrap (_dfg ,_ea ,"\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073");};_ab +=int64 (_af );return _ab ,nil ;
};func (_gcc *Encoder )EncodeBitmap (bm *_b .Bitmap ,duplicateLineRemoval bool )error {_d .Log .Trace ("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s",bm .Width ,bm .Height ,bm );
var (_fdg ,_dd uint8 ;_ece ,_cb ,_efg uint16 ;_cac ,_cfe ,_cad byte ;_fg ,_dc ,_agg int ;_fb ,_bf []byte ;);for _gcf :=0;_gcf < bm .Height ;_gcf ++{_cac ,_cfe =0,0;if _gcf >=2{_cac =bm .Data [(_gcf -2)*bm .RowStride ];};if _gcf >=1{_cfe =bm .Data [(_gcf -1)*bm .RowStride ];
if duplicateLineRemoval {_dc =_gcf *bm .RowStride ;_fb =bm .Data [_dc :_dc +bm .RowStride ];_agg =(_gcf -1)*bm .RowStride ;_bf =bm .Data [_agg :_agg +bm .RowStride ];if _c .Equal (_fb ,_bf ){_dd =_fdg ^1;_fdg =1;}else {_dd =_fdg ;_fdg =0;};};};if duplicateLineRemoval {if _bcb :=_gcc .encodeBit (_gcc ._cc ,_fd ,_dd );
_bcb !=nil {return _bcb ;};if _fdg !=0{continue ;};};_cad =bm .Data [_gcf *bm .RowStride ];_ece =uint16 (_cac >>5);_cb =uint16 (_cfe >>4);_cac <<=3;_cfe <<=4;_efg =0;for _fg =0;_fg < bm .Width ;_fg ++{_agbd :=uint32 (_ece <<11|_cb <<4|_efg );_gge :=(_cad &0x80)>>7;
_ga :=_gcc .encodeBit (_gcc ._cc ,_agbd ,_gge );if _ga !=nil {return _ga ;};_ece <<=1;_cb <<=1;_efg <<=1;_ece |=uint16 ((_cac &0x80)>>7);_cb |=uint16 ((_cfe &0x80)>>7);_efg |=uint16 (_gge );_gaa :=_fg %8;_adb :=_fg /8+1;if _gaa ==4&&_gcf >=2{_cac =0;if _adb < bm .RowStride {_cac =bm .Data [(_gcf -2)*bm .RowStride +_adb ];
};}else {_cac <<=1;};if _gaa ==3&&_gcf >=1{_cfe =0;if _adb < bm .RowStride {_cfe =bm .Data [(_gcf -1)*bm .RowStride +_adb ];};}else {_cfe <<=1;};if _gaa ==7{_cad =0;if _adb < bm .RowStride {_cad =bm .Data [_gcf *bm .RowStride +_adb ];};}else {_cad <<=1;
};_ece &=31;_cb &=127;_efg &=15;};};return nil ;};var _gd =[]intEncRangeS {{0,3,0,2,0,2},{-1,-1,9,4,0,0},{-3,-2,5,3,2,1},{4,19,2,3,4,4},{-19,-4,3,3,4,4},{20,83,6,4,20,6},{-83,-20,7,4,20,6},{84,339,14,5,84,8},{-339,-84,15,5,84,8},{340,4435,30,6,340,12},{-4435,-340,31,6,340,12},{4436,2000000000,62,6,4436,32},{-2000000000,-4436,63,6,4436,32}};
func (_gg *codingContext )flipMps (_gb uint32 ){_gg ._ca [_gb ]=1-_gg ._ca [_gb ]};func (_fbf *Encoder )encodeIAID (_fdf ,_abd int )error {if _fbf ._ge ==nil {_fbf ._ge =_ad (1<<uint (_fdf ));};_ead :=uint32 (1<<uint32 (_fdf +1))-1;_abd <<=uint (32-_fdf );
_dee :=uint32 (1);for _eaf :=0;_eaf < _fdf ;_eaf ++{_agc :=_dee &_ead ;_dfge :=uint8 ((uint32 (_abd )&0x80000000)>>31);if _gced :=_fbf .encodeBit (_fbf ._ge ,_agc ,_dfge );_gced !=nil {return _gced ;};_dee =(_dee <<1)|uint32 (_dfge );_abd <<=1;};return nil ;
};func (_gc Class )String ()string {switch _gc {case IAAI :return "\u0049\u0041\u0041\u0049";case IADH :return "\u0049\u0041\u0044\u0048";case IADS :return "\u0049\u0041\u0044\u0053";case IADT :return "\u0049\u0041\u0044\u0054";case IADW :return "\u0049\u0041\u0044\u0057";
case IAEX :return "\u0049\u0041\u0045\u0058";case IAFS :return "\u0049\u0041\u0046\u0053";case IAIT :return "\u0049\u0041\u0049\u0054";case IARDH :return "\u0049\u0041\u0052D\u0048";case IARDW :return "\u0049\u0041\u0052D\u0057";case IARDX :return "\u0049\u0041\u0052D\u0058";
case IARDY :return "\u0049\u0041\u0052D\u0059";case IARI :return "\u0049\u0041\u0052\u0049";default:return "\u0055N\u004b\u004e\u004f\u0057\u004e";};};func (_aad *Encoder )encodeBit (_gaag *codingContext ,_dgf uint32 ,_bcd uint8 )error {const _cba ="\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074";
_aad ._dbf ++;if _dgf >=uint32 (len (_gaag ._be )){return _bb .Errorf (_cba ,"\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027",_dgf );
};_afd :=_gaag ._be [_dgf ];_adf :=_gaag .mps (_dgf );_dbdc :=_gab [_afd ]._fcb ;_d .Log .Trace ("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064",_aad ._dbf ,_bcd ,_afd ,_adf ,_dbdc ,_aad ._ecc ,_aad ._cf ,_aad ._f ,_aad ._bbe ,_aad ._ee );
if _bcd ==0{_aad .code0 (_gaag ,_dgf ,_dbdc ,_afd );}else {_aad .code1 (_gaag ,_dgf ,_dbdc ,_afd );};return nil ;};func (_fcd *Encoder )emit (){if _fcd ._gdg ==_aff {_fcd ._dg =append (_fcd ._dg ,_fcd ._cfg );_fcd ._cfg =make ([]byte ,_aff );_fcd ._gdg =0;
};_fcd ._cfg [_fcd ._gdg ]=_fcd ._bbe ;_fcd ._gdg ++;};func (_eccb *Encoder )flush (){_eccb .setBits ();_eccb ._cf <<=_eccb ._f ;_eccb .byteOut ();_eccb ._cf <<=_eccb ._f ;_eccb .byteOut ();_eccb .emit ();if _eccb ._bbe !=0xff{_eccb ._ee ++;_eccb ._bbe =0xff;
_eccb .emit ();};_eccb ._ee ++;_eccb ._bbe =0xac;_eccb ._ee ++;_eccb .emit ();};func (_bed *Encoder )Refine (iTemp ,iTarget *_b .Bitmap ,ox ,oy int )error {for _bd :=0;_bd < iTarget .Height ;_bd ++{var _gbf int ;_ged :=_bd +oy ;var (_aa ,_ccf ,_fde ,_cce ,_bbf uint16 ;
_eb ,_eba ,_aag ,_bfg ,_ebg byte ;);if _ged >=1&&(_ged -1)< iTemp .Height {_eb =iTemp .Data [(_ged -1)*iTemp .RowStride ];};if _ged >=0&&_ged < iTemp .Height {_eba =iTemp .Data [_ged *iTemp .RowStride ];};if _ged >=-1&&_ged +1< iTemp .Height {_aag =iTemp .Data [(_ged +1)*iTemp .RowStride ];
};if _bd >=1{_bfg =iTarget .Data [(_bd -1)*iTarget .RowStride ];};_ebg =iTarget .Data [_bd *iTarget .RowStride ];_cadb :=uint (6+ox );_aa =uint16 (_eb >>_cadb );_ccf =uint16 (_eba >>_cadb );_fde =uint16 (_aag >>_cadb );_cce =uint16 (_bfg >>6);_cacc :=uint (2-ox );
_eb <<=_cacc ;_eba <<=_cacc ;_aag <<=_cacc ;_bfg <<=2;for _gbf =0;_gbf < iTarget .Width ;_gbf ++{_fe :=(_aa <<10)|(_ccf <<7)|(_fde <<4)|(_cce <<1)|_bbf ;_bab :=_ebg >>7;_gbg :=_bed .encodeBit (_bed ._cc ,uint32 (_fe ),_bab );if _gbg !=nil {return _gbg ;
};_aa <<=1;_ccf <<=1;_fde <<=1;_cce <<=1;_aa |=uint16 (_eb >>7);_ccf |=uint16 (_eba >>7);_fde |=uint16 (_aag >>7);_cce |=uint16 (_bfg >>7);_bbf =uint16 (_bab );_dda :=_gbf %8;_df :=_gbf /8+1;if _dda ==5+ox {_eb ,_eba ,_aag =0,0,0;if _df < iTemp .RowStride &&_ged >=1&&(_ged -1)< iTemp .Height {_eb =iTemp .Data [(_ged -1)*iTemp .RowStride +_df ];
};if _df < iTemp .RowStride &&_ged >=0&&_ged < iTemp .Height {_eba =iTemp .Data [_ged *iTemp .RowStride +_df ];};if _df < iTemp .RowStride &&_ged >=-1&&(_ged +1)< iTemp .Height {_aag =iTemp .Data [(_ged +1)*iTemp .RowStride +_df ];};}else {_eb <<=1;_eba <<=1;
_aag <<=1;};if _dda ==5&&_bd >=1{_bfg =0;if _df < iTarget .RowStride {_bfg =iTarget .Data [(_bd -1)*iTarget .RowStride +_df ];};}else {_bfg <<=1;};if _dda ==7{_ebg =0;if _df < iTarget .RowStride {_ebg =iTarget .Data [_bd *iTarget .RowStride +_df ];};}else {_ebg <<=1;
};_aa &=7;_ccf &=7;_fde &=7;_cce &=7;};};return nil ;};func (_baf *Encoder )rBlock (){if _baf ._ee >=0{_baf .emit ();};_baf ._ee ++;_baf ._bbe =uint8 (_baf ._cf >>20);_baf ._cf &=0xfffff;_baf ._f =7;};type intEncRangeS struct{_bg ,_db int ;_de ,_a uint8 ;
_ec uint16 ;_dba uint8 ;};func (_aef *Encoder )byteOut (){if _aef ._bbe ==0xff{_aef .rBlock ();return ;};if _aef ._cf < 0x8000000{_aef .lBlock ();return ;};_aef ._bbe ++;if _aef ._bbe !=0xff{_aef .lBlock ();return ;};_aef ._cf &=0x7ffffff;_aef .rBlock ();
};var _gab =[]state {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
func (_fdfe *Encoder )setBits (){_cgg :=_fdfe ._cf +uint32 (_fdfe ._ecc );_fdfe ._cf |=0xffff;if _fdfe ._cf >=_cgg {_fdfe ._cf -=0x8000;};};