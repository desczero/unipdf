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

package arithmetic ;import (_d "bytes";_de "github.com/unidoc/unipdf/v3/common";_bb "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_e "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_c "io";);func (_bac *Encoder )code0 (_cbb *codingContext ,_ae uint32 ,_aab uint16 ,_be byte ){if _cbb .mps (_ae )==0{_bac .codeMPS (_cbb ,_ae ,_aab ,_be );
}else {_bac .codeLPS (_cbb ,_ae ,_aab ,_be );};};func (_cga *Encoder )emit (){if _cga ._ec ==_efa {_cga ._bbc =append (_cga ._bbc ,_cga ._gea );_cga ._gea =make ([]byte ,_efa );_cga ._ec =0;};_cga ._gea [_cga ._ec ]=_cga ._aa ;_cga ._ec ++;};func (_aac *Encoder )Final (){_aac .flush ()};
type intEncRangeS struct{_fc ,_da int ;_g ,_fe uint8 ;_a uint16 ;_ag uint8 ;};func (_fb *Encoder )Flush (){_fb ._ec =0;_fb ._bbc =nil ;_fb ._bae =-1};func New ()*Encoder {_dee :=&Encoder {};_dee .Init ();return _dee };type codingContext struct{_ba []byte ;
_ca []byte ;};type Encoder struct{_fed uint32 ;_db uint16 ;_ded ,_aa uint8 ;_bae int ;_cacg int ;_bbc [][]byte ;_gea []byte ;_ec int ;_eg *codingContext ;_bf [13]*codingContext ;_dd *codingContext ;};func (_efc *Encoder )Reset (){_efc ._db =0x8000;_efc ._fed =0;
_efc ._ded =12;_efc ._bae =-1;_efc ._aa =0;_efc ._dd =nil ;_efc ._eg =_cac (_bge );};type state struct{_bdb uint16 ;_dgd ,_fdc uint8 ;_fga uint8 ;};func (_agc *Encoder )DataSize ()int {return _agc .dataSize ()};func (_fff *Encoder )EncodeIAID (symbolCodeLength ,value int )(_ea error ){_de .Log .Trace ("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",symbolCodeLength ,value );
if _ea =_fff .encodeIAID (symbolCodeLength ,value );_ea !=nil {return _e .Wrap (_ea ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044","");};return nil ;};func (_ga *codingContext )flipMps (_bg uint32 ){_ga ._ca [_bg ]=1-_ga ._ca [_bg ]};func (_f Class )String ()string {switch _f {case IAAI :return "\u0049\u0041\u0041\u0049";
case IADH :return "\u0049\u0041\u0044\u0048";case IADS :return "\u0049\u0041\u0044\u0053";case IADT :return "\u0049\u0041\u0044\u0054";case IADW :return "\u0049\u0041\u0044\u0057";case IAEX :return "\u0049\u0041\u0045\u0058";case IAFS :return "\u0049\u0041\u0046\u0053";
case IAIT :return "\u0049\u0041\u0049\u0054";case IARDH :return "\u0049\u0041\u0052D\u0048";case IARDW :return "\u0049\u0041\u0052D\u0057";case IARDX :return "\u0049\u0041\u0052D\u0058";case IARDY :return "\u0049\u0041\u0052D\u0059";case IARI :return "\u0049\u0041\u0052\u0049";
default:return "\u0055N\u004b\u004e\u004f\u0057\u004e";};};func (_efg *Encoder )lBlock (){if _efg ._bae >=0{_efg .emit ();};_efg ._bae ++;_efg ._aa =uint8 (_efg ._fed >>19);_efg ._fed &=0x7ffff;_efg ._ded =8;};func (_ee *Encoder )byteOut (){if _ee ._aa ==0xff{_ee .rBlock ();
return ;};if _ee ._fed < 0x8000000{_ee .lBlock ();return ;};_ee ._aa ++;if _ee ._aa !=0xff{_ee .lBlock ();return ;};_ee ._fed &=0x7ffffff;_ee .rBlock ();};func (_daa *Encoder )Refine (iTemp ,iTarget *_bb .Bitmap ,ox ,oy int )error {for _fa :=0;_fa < iTarget .Height ;
_fa ++{var _aad int ;_bafe :=_fa +oy ;var (_fedc ,_fgd ,_bc ,_eb ,_faa uint16 ;_eac ,_gag ,_dea ,_cf ,_eba byte ;);if _bafe >=1&&(_bafe -1)< iTemp .Height {_eac =iTemp .Data [(_bafe -1)*iTemp .RowStride ];};if _bafe >=0&&_bafe < iTemp .Height {_gag =iTemp .Data [_bafe *iTemp .RowStride ];
};if _bafe >=-1&&_bafe +1< iTemp .Height {_dea =iTemp .Data [(_bafe +1)*iTemp .RowStride ];};if _fa >=1{_cf =iTarget .Data [(_fa -1)*iTarget .RowStride ];};_eba =iTarget .Data [_fa *iTarget .RowStride ];_cdf :=uint (6+ox );_fedc =uint16 (_eac >>_cdf );
_fgd =uint16 (_gag >>_cdf );_bc =uint16 (_dea >>_cdf );_eb =uint16 (_cf >>6);_faae :=uint (2-ox );_eac <<=_faae ;_gag <<=_faae ;_dea <<=_faae ;_cf <<=2;for _aad =0;_aad < iTarget .Width ;_aad ++{_gad :=(_fedc <<10)|(_fgd <<7)|(_bc <<4)|(_eb <<1)|_faa ;
_ce :=_eba >>7;_egg :=_daa .encodeBit (_daa ._eg ,uint32 (_gad ),_ce );if _egg !=nil {return _egg ;};_fedc <<=1;_fgd <<=1;_bc <<=1;_eb <<=1;_fedc |=uint16 (_eac >>7);_fgd |=uint16 (_gag >>7);_bc |=uint16 (_dea >>7);_eb |=uint16 (_cf >>7);_faa =uint16 (_ce );
_dde :=_aad %8;_ffc :=_aad /8+1;if _dde ==5+ox {_eac ,_gag ,_dea =0,0,0;if _ffc < iTemp .RowStride &&_bafe >=1&&(_bafe -1)< iTemp .Height {_eac =iTemp .Data [(_bafe -1)*iTemp .RowStride +_ffc ];};if _ffc < iTemp .RowStride &&_bafe >=0&&_bafe < iTemp .Height {_gag =iTemp .Data [_bafe *iTemp .RowStride +_ffc ];
};if _ffc < iTemp .RowStride &&_bafe >=-1&&(_bafe +1)< iTemp .Height {_dea =iTemp .Data [(_bafe +1)*iTemp .RowStride +_ffc ];};}else {_eac <<=1;_gag <<=1;_dea <<=1;};if _dde ==5&&_fa >=1{_cf =0;if _ffc < iTarget .RowStride {_cf =iTarget .Data [(_fa -1)*iTarget .RowStride +_ffc ];
};}else {_cf <<=1;};if _dde ==7{_eba =0;if _ffc < iTarget .RowStride {_eba =iTarget .Data [_fa *iTarget .RowStride +_ffc ];};}else {_eba <<=1;};_fedc &=7;_fgd &=7;_bc &=7;_eb &=7;};};return nil ;};func (_bgd *Encoder )flush (){_bgd .setBits ();_bgd ._fed <<=_bgd ._ded ;
_bgd .byteOut ();_bgd ._fed <<=_bgd ._ded ;_bgd .byteOut ();_bgd .emit ();if _bgd ._aa !=0xff{_bgd ._bae ++;_bgd ._aa =0xff;_bgd .emit ();};_bgd ._bae ++;_bgd ._aa =0xac;_bgd ._bae ++;_bgd .emit ();};func (_bad *Encoder )WriteTo (w _c .Writer )(int64 ,error ){const _gec ="\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f";
var _bcb int64 ;for _ged ,_cgb :=range _bad ._bbc {_eacc ,_bfb :=w .Write (_cgb );if _bfb !=nil {return 0,_e .Wrapf (_bfb ,_gec ,"\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b",_ged );
};_bcb +=int64 (_eacc );};_bad ._gea =_bad ._gea [:_bad ._ec ];_ddg ,_eae :=w .Write (_bad ._gea );if _eae !=nil {return 0,_e .Wrap (_eae ,_gec ,"\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073");};_bcb +=int64 (_ddg );
return _bcb ,nil ;};const _ac =0x9b25;func (_aca *Encoder )EncodeOOB (proc Class )(_fee error ){_de .Log .Trace ("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027",proc );
if _fee =_aca .encodeOOB (proc );_fee !=nil {return _e .Wrap (_fee ,"\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB","");};return nil ;};func (_deea *Encoder )encodeInteger (_edb Class ,_cgd int )error {const _bd ="E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072";
if _cgd > 2000000000||_cgd < -2000000000{return _e .Errorf (_bd ,"\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",_cgd );
};_edf :=_deea ._bf [_edb ];_gg :=uint32 (1);var _bcbe int ;for ;;_bcbe ++{if _fd [_bcbe ]._fc <=_cgd &&_fd [_bcbe ]._da >=_cgd {break ;};};if _cgd < 0{_cgd =-_cgd ;};_cgd -=int (_fd [_bcbe ]._a );_fab :=_fd [_bcbe ]._g ;for _bacd :=uint8 (0);_bacd < _fd [_bcbe ]._fe ;
_bacd ++{_cbg :=_fab &1;if _abc :=_deea .encodeBit (_edf ,_gg ,_cbg );_abc !=nil {return _e .Wrap (_abc ,_bd ,"");};_fab >>=1;if _gg &0x100> 0{_gg =(((_gg <<1)|uint32 (_cbg ))&0x1ff)|0x100;}else {_gg =(_gg <<1)|uint32 (_cbg );};};_cgd <<=32-_fd [_bcbe ]._ag ;
for _dff :=uint8 (0);_dff < _fd [_bcbe ]._ag ;_dff ++{_edef :=uint8 ((uint32 (_cgd )&0x80000000)>>31);if _fdg :=_deea .encodeBit (_edf ,_gg ,_edef );_fdg !=nil {return _e .Wrap (_fdg ,_bd ,"\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064");
};_cgd <<=1;if _gg &0x100!=0{_gg =(((_gg <<1)|uint32 (_edef ))&0x1ff)|0x100;}else {_gg =(_gg <<1)|uint32 (_edef );};};return nil ;};func (_eee *Encoder )code1 (_gd *codingContext ,_df uint32 ,_egd uint16 ,_cbf byte ){if _gd .mps (_df )==1{_eee .codeMPS (_gd ,_df ,_egd ,_cbf );
}else {_eee .codeLPS (_gd ,_df ,_egd ,_cbf );};};func (_afcg *Encoder )encodeIAID (_acb ,_gab int )error {if _afcg ._dd ==nil {_afcg ._dd =_cac (1<<uint (_acb ));};_fbe :=uint32 (1<<uint32 (_acb +1))-1;_gab <<=uint (32-_acb );_cdfc :=uint32 (1);for _gge :=0;
_gge < _acb ;_gge ++{_fdga :=_cdfc &_fbe ;_bag :=uint8 ((uint32 (_gab )&0x80000000)>>31);if _afb :=_afcg .encodeBit (_afcg ._dd ,_fdga ,_bag );_afb !=nil {return _afb ;};_cdfc =(_cdfc <<1)|uint32 (_bag );_gab <<=1;};return nil ;};func (_fgdf *Encoder )codeMPS (_ffe *codingContext ,_ede uint32 ,_fdeg uint16 ,_dgb byte ){_fgdf ._db -=_fdeg ;
if _fgdf ._db &0x8000!=0{_fgdf ._fed +=uint32 (_fdeg );return ;};if _fgdf ._db < _fdeg {_fgdf ._db =_fdeg ;}else {_fgdf ._fed +=uint32 (_fdeg );};_ffe ._ba [_ede ]=_ccf [_dgb ]._dgd ;_fgdf .renormalize ();};const (_bge =65536;_efa =20*1024;);func (_cd *Encoder )EncodeBitmap (bm *_bb .Bitmap ,duplicateLineRemoval bool )error {_de .Log .Trace ("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s",bm .Width ,bm .Height ,bm );
var (_ab ,_ef uint8 ;_bbd ,_aga ,_cc uint16 ;_af ,_dc ,_ed byte ;_feda ,_afc ,_cb int ;_fdf ,_fde []byte ;);for _aaa :=0;_aaa < bm .Height ;_aaa ++{_af ,_dc =0,0;if _aaa >=2{_af =bm .Data [(_aaa -2)*bm .RowStride ];};if _aaa >=1{_dc =bm .Data [(_aaa -1)*bm .RowStride ];
if duplicateLineRemoval {_afc =_aaa *bm .RowStride ;_fdf =bm .Data [_afc :_afc +bm .RowStride ];_cb =(_aaa -1)*bm .RowStride ;_fde =bm .Data [_cb :_cb +bm .RowStride ];if _d .Equal (_fdf ,_fde ){_ef =_ab ^1;_ab =1;}else {_ef =_ab ;_ab =0;};};};if duplicateLineRemoval {if _baf :=_cd .encodeBit (_cd ._eg ,_ac ,_ef );
_baf !=nil {return _baf ;};if _ab !=0{continue ;};};_ed =bm .Data [_aaa *bm .RowStride ];_bbd =uint16 (_af >>5);_aga =uint16 (_dc >>4);_af <<=3;_dc <<=4;_cc =0;for _feda =0;_feda < bm .Width ;_feda ++{_cg :=uint32 (_bbd <<11|_aga <<4|_cc );_fg :=(_ed &0x80)>>7;
_fgg :=_cd .encodeBit (_cd ._eg ,_cg ,_fg );if _fgg !=nil {return _fgg ;};_bbd <<=1;_aga <<=1;_cc <<=1;_bbd |=uint16 ((_af &0x80)>>7);_aga |=uint16 ((_dc &0x80)>>7);_cc |=uint16 (_fg );_aag :=_feda %8;_gee :=_feda /8+1;if _aag ==4&&_aaa >=2{_af =0;if _gee < bm .RowStride {_af =bm .Data [(_aaa -2)*bm .RowStride +_gee ];
};}else {_af <<=1;};if _aag ==3&&_aaa >=1{_dc =0;if _gee < bm .RowStride {_dc =bm .Data [(_aaa -1)*bm .RowStride +_gee ];};}else {_dc <<=1;};if _aag ==7{_ed =0;if _gee < bm .RowStride {_ed =bm .Data [_aaa *bm .RowStride +_gee ];};}else {_ed <<=1;};_bbd &=31;
_aga &=127;_cc &=15;};};return nil ;};func (_ge *codingContext )mps (_cab uint32 )int {return int (_ge ._ca [_cab ])};func (_dcb *Encoder )encodeOOB (_cfd Class )error {_aade :=_dcb ._bf [_cfd ];_dfd :=_dcb .encodeBit (_aade ,1,1);if _dfd !=nil {return _dfd ;
};_dfd =_dcb .encodeBit (_aade ,3,0);if _dfd !=nil {return _dfd ;};_dfd =_dcb .encodeBit (_aade ,6,0);if _dfd !=nil {return _dfd ;};_dfd =_dcb .encodeBit (_aade ,12,0);if _dfd !=nil {return _dfd ;};return nil ;};const (IAAI Class =iota ;IADH ;IADS ;IADT ;
IADW ;IAEX ;IAFS ;IAIT ;IARDH ;IARDW ;IARDX ;IARDY ;IARI ;);func _cac (_ff int )*codingContext {return &codingContext {_ba :make ([]byte ,_ff ),_ca :make ([]byte ,_ff )};};func (_baa *Encoder )rBlock (){if _baa ._bae >=0{_baa .emit ();};_baa ._bae ++;_baa ._aa =uint8 (_baa ._fed >>20);
_baa ._fed &=0xfffff;_baa ._ded =7;};func (_agaf *Encoder )encodeBit (_fae *codingContext ,_gdc uint32 ,_def uint8 )error {const _fcc ="\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074";_agaf ._cacg ++;if _gdc >=uint32 (len (_fae ._ba )){return _e .Errorf (_fcc ,"\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027",_gdc );
};_gba :=_fae ._ba [_gdc ];_bfg :=_fae .mps (_gdc );_gcc :=_ccf [_gba ]._bdb ;_de .Log .Trace ("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064",_agaf ._cacg ,_def ,_gba ,_bfg ,_gcc ,_agaf ._db ,_agaf ._fed ,_agaf ._ded ,_agaf ._aa ,_agaf ._bae );
if _def ==0{_agaf .code0 (_fae ,_gdc ,_gcc ,_gba );}else {_agaf .code1 (_fae ,_gdc ,_gcc ,_gba );};return nil ;};func (_fccg *Encoder )renormalize (){for {_fccg ._db <<=1;_fccg ._fed <<=1;_fccg ._ded --;if _fccg ._ded ==0{_fccg .byteOut ();};if (_fccg ._db &0x8000)!=0{break ;
};};};var _ccf =[]state {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
func (_gb *Encoder )Init (){_gb ._eg =_cac (_bge );_gb ._db =0x8000;_gb ._fed =0;_gb ._ded =12;_gb ._bae =-1;_gb ._aa =0;_gb ._ec =0;_gb ._gea =make ([]byte ,_efa );for _gc :=0;_gc < len (_gb ._bf );_gc ++{_gb ._bf [_gc ]=_cac (512);};_gb ._dd =nil ;};
func (_fef *Encoder )EncodeInteger (proc Class ,value int )(_dg error ){_de .Log .Trace ("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027",value ,proc );
if _dg =_fef .encodeInteger (proc ,value );_dg !=nil {return _e .Wrap (_dg ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072","");};return nil ;};func (_fad *Encoder )codeLPS (_afe *codingContext ,_dfa uint32 ,_acf uint16 ,_fec byte ){_fad ._db -=_acf ;
if _fad ._db < _acf {_fad ._fed +=uint32 (_acf );}else {_fad ._db =_acf ;};if _ccf [_fec ]._fga ==1{_afe .flipMps (_dfa );};_afe ._ba [_dfa ]=_ccf [_fec ]._fdc ;_fad .renormalize ();};type Class int ;var _fd =[]intEncRangeS {{0,3,0,2,0,2},{-1,-1,9,4,0,0},{-3,-2,5,3,2,1},{4,19,2,3,4,4},{-19,-4,3,3,4,4},{20,83,6,4,20,6},{-83,-20,7,4,20,6},{84,339,14,5,84,8},{-339,-84,15,5,84,8},{340,4435,30,6,340,12},{-4435,-340,31,6,340,12},{4436,2000000000,62,6,4436,32},{-2000000000,-4436,63,6,4436,32}};
func (_acab *Encoder )dataSize ()int {return _efa *len (_acab ._bbc )+_acab ._ec };var _ _c .WriterTo =&Encoder {};func (_gcd *Encoder )setBits (){_cbd :=_gcd ._fed +uint32 (_gcd ._db );_gcd ._fed |=0xffff;if _gcd ._fed >=_cbd {_gcd ._fed -=0x8000;};};
