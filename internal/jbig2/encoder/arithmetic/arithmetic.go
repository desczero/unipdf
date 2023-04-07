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

package arithmetic ;import (_bb "bytes";_ae "github.com/unidoc/unipdf/v3/common";_c "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_cb "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_a "io";);func (_gc *Encoder )Init (){_gc ._ga =_bf (_fbb );
_gc ._eg =0x8000;_gc ._gea =0;_gc ._geb =12;_gc ._ec =-1;_gc ._gf =0;_gc ._ad =0;_gc ._eb =make ([]byte ,_ceg );for _bd :=0;_bd < len (_gc ._ac );_bd ++{_gc ._ac [_bd ]=_bf (512);};_gc ._dag =nil ;};type intEncRangeS struct{_g ,_f int ;_bbc ,_d uint8 ;
_ag uint16 ;_ef uint8 ;};func (_agc *Encoder )Reset (){_agc ._eg =0x8000;_agc ._gea =0;_agc ._geb =12;_agc ._ec =-1;_agc ._gf =0;_agc ._dag =nil ;_agc ._ga =_bf (_fbb );};func (_eff *Encoder )setBits (){_gebc :=_eff ._gea +uint32 (_eff ._eg );_eff ._gea |=0xffff;
if _eff ._gea >=_gebc {_eff ._gea -=0x8000;};};func (_dgfd *Encoder )codeLPS (_gg *codingContext ,_dba uint32 ,_ffed uint16 ,_adac byte ){_dgfd ._eg -=_ffed ;if _dgfd ._eg < _ffed {_dgfd ._gea +=uint32 (_ffed );}else {_dgfd ._eg =_ffed ;};if _bage [_adac ]._dga ==1{_gg .flipMps (_dba );
};_gg ._ff [_dba ]=_bage [_adac ]._caf ;_dgfd .renormalize ();};func (_eba *Encoder )EncodeOOB (proc Class )(_bg error ){_ae .Log .Trace ("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027",proc );
if _bg =_eba .encodeOOB (proc );_bg !=nil {return _cb .Wrap (_bg ,"\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB","");};return nil ;};func (_ab *Encoder )DataSize ()int {return _ab .dataSize ()};func (_gfaf *Encoder )flush (){_gfaf .setBits ();_gfaf ._gea <<=_gfaf ._geb ;
_gfaf .byteOut ();_gfaf ._gea <<=_gfaf ._geb ;_gfaf .byteOut ();_gfaf .emit ();if _gfaf ._gf !=0xff{_gfaf ._ec ++;_gfaf ._gf =0xff;_gfaf .emit ();};_gfaf ._ec ++;_gfaf ._gf =0xac;_gfaf ._ec ++;_gfaf .emit ();};func (_ba *codingContext )mps (_fb uint32 )int {return int (_ba ._ge [_fb ])};
func (_agg *Encoder )EncodeIAID (symbolCodeLength ,value int )(_afg error ){_ae .Log .Trace ("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",symbolCodeLength ,value );
if _afg =_agg .encodeIAID (symbolCodeLength ,value );_afg !=nil {return _cb .Wrap (_afg ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044","");};return nil ;};const (IAAI Class =iota ;IADH ;IADS ;IADT ;IADW ;IAEX ;IAFS ;IAIT ;IARDH ;IARDW ;
IARDX ;IARDY ;IARI ;);func (_abc *Encoder )EncodeBitmap (bm *_c .Bitmap ,duplicateLineRemoval bool )error {_ae .Log .Trace ("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s",bm .Width ,bm .Height ,bm );
var (_caa ,_cg uint8 ;_geae ,_dgc ,_aec uint16 ;_cge ,_daa ,_dc byte ;_cgf ,_ffb ,_bac int ;_df ,_af []byte ;);for _be :=0;_be < bm .Height ;_be ++{_cge ,_daa =0,0;if _be >=2{_cge =bm .Data [(_be -2)*bm .RowStride ];};if _be >=1{_daa =bm .Data [(_be -1)*bm .RowStride ];
if duplicateLineRemoval {_ffb =_be *bm .RowStride ;_df =bm .Data [_ffb :_ffb +bm .RowStride ];_bac =(_be -1)*bm .RowStride ;_af =bm .Data [_bac :_bac +bm .RowStride ];if _bb .Equal (_df ,_af ){_cg =_caa ^1;_caa =1;}else {_cg =_caa ;_caa =0;};};};if duplicateLineRemoval {if _gfa :=_abc .encodeBit (_abc ._ga ,_dbe ,_cg );
_gfa !=nil {return _gfa ;};if _caa !=0{continue ;};};_dc =bm .Data [_be *bm .RowStride ];_geae =uint16 (_cge >>5);_dgc =uint16 (_daa >>4);_cge <<=3;_daa <<=4;_aec =0;for _cgf =0;_cgf < bm .Width ;_cgf ++{_aeb :=uint32 (_geae <<11|_dgc <<4|_aec );_efg :=(_dc &0x80)>>7;
_de :=_abc .encodeBit (_abc ._ga ,_aeb ,_efg );if _de !=nil {return _de ;};_geae <<=1;_dgc <<=1;_aec <<=1;_geae |=uint16 ((_cge &0x80)>>7);_dgc |=uint16 ((_daa &0x80)>>7);_aec |=uint16 (_efg );_egf :=_cgf %8;_bbe :=_cgf /8+1;if _egf ==4&&_be >=2{_cge =0;
if _bbe < bm .RowStride {_cge =bm .Data [(_be -2)*bm .RowStride +_bbe ];};}else {_cge <<=1;};if _egf ==3&&_be >=1{_daa =0;if _bbe < bm .RowStride {_daa =bm .Data [(_be -1)*bm .RowStride +_bbe ];};}else {_daa <<=1;};if _egf ==7{_dc =0;if _bbe < bm .RowStride {_dc =bm .Data [_be *bm .RowStride +_bbe ];
};}else {_dc <<=1;};_geae &=31;_dgc &=127;_aec &=15;};};return nil ;};func New ()*Encoder {_fe :=&Encoder {};_fe .Init ();return _fe };func (_gae *Encoder )EncodeInteger (proc Class ,value int )(_cc error ){_ae .Log .Trace ("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027",value ,proc );
if _cc =_gae .encodeInteger (proc ,value );_cc !=nil {return _cb .Wrap (_cc ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072","");};return nil ;};func (_ebgf *Encoder )code1 (_debb *codingContext ,_aadf uint32 ,_ddg uint16 ,_cgec byte ){if _debb .mps (_aadf )==1{_ebgf .codeMPS (_debb ,_aadf ,_ddg ,_cgec );
}else {_ebgf .codeLPS (_debb ,_aadf ,_ddg ,_cgec );};};type codingContext struct{_ff []byte ;_ge []byte ;};func (_gd *codingContext )flipMps (_db uint32 ){_gd ._ge [_db ]=1-_gd ._ge [_db ]};const _dbe =0x9b25;func (_ddb *Encoder )codeMPS (_gcc *codingContext ,_gebd uint32 ,_eadg uint16 ,_dfg byte ){_ddb ._eg -=_eadg ;
if _ddb ._eg &0x8000!=0{_ddb ._gea +=uint32 (_eadg );return ;};if _ddb ._eg < _eadg {_ddb ._eg =_eadg ;}else {_ddb ._gea +=uint32 (_eadg );};_gcc ._ff [_gebd ]=_bage [_dfg ]._dbab ;_ddb .renormalize ();};func _bf (_ca int )*codingContext {return &codingContext {_ff :make ([]byte ,_ca ),_ge :make ([]byte ,_ca )};
};func (_caag *Encoder )encodeInteger (_adae Class ,_cdd int )error {const _bdee ="E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072";if _cdd > 2000000000||_cdd < -2000000000{return _cb .Errorf (_bdee ,"\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",_cdd );
};_dgg :=_caag ._ac [_adae ];_efe :=uint32 (1);var _bgf int ;for ;;_bgf ++{if _cf [_bgf ]._g <=_cdd &&_cf [_bgf ]._f >=_cdd {break ;};};if _cdd < 0{_cdd =-_cdd ;};_cdd -=int (_cf [_bgf ]._ag );_gdb :=_cf [_bgf ]._bbc ;for _dbg :=uint8 (0);_dbg < _cf [_bgf ]._d ;
_dbg ++{_dbb :=_gdb &1;if _dbf :=_caag .encodeBit (_dgg ,_efe ,_dbb );_dbf !=nil {return _cb .Wrap (_dbf ,_bdee ,"");};_gdb >>=1;if _efe &0x100> 0{_efe =(((_efe <<1)|uint32 (_dbb ))&0x1ff)|0x100;}else {_efe =(_efe <<1)|uint32 (_dbb );};};_cdd <<=32-_cf [_bgf ]._ef ;
for _cea :=uint8 (0);_cea < _cf [_bgf ]._ef ;_cea ++{_ffa :=uint8 ((uint32 (_cdd )&0x80000000)>>31);if _def :=_caag .encodeBit (_dgg ,_efe ,_ffa );_def !=nil {return _cb .Wrap (_def ,_bdee ,"\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064");
};_cdd <<=1;if _efe &0x100!=0{_efe =(((_efe <<1)|uint32 (_ffa ))&0x1ff)|0x100;}else {_efe =(_efe <<1)|uint32 (_ffa );};};return nil ;};func (_cfg *Encoder )renormalize (){for {_cfg ._eg <<=1;_cfg ._gea <<=1;_cfg ._geb --;if _cfg ._geb ==0{_cfg .byteOut ();
};if (_cfg ._eg &0x8000)!=0{break ;};};};var _cf =[]intEncRangeS {{0,3,0,2,0,2},{-1,-1,9,4,0,0},{-3,-2,5,3,2,1},{4,19,2,3,4,4},{-19,-4,3,3,4,4},{20,83,6,4,20,6},{-83,-20,7,4,20,6},{84,339,14,5,84,8},{-339,-84,15,5,84,8},{340,4435,30,6,340,12},{-4435,-340,31,6,340,12},{4436,2000000000,62,6,4436,32},{-2000000000,-4436,63,6,4436,32}};
func (_ee *Encoder )Final (){_ee .flush ()};type state struct{_ecb uint16 ;_dbab ,_caf uint8 ;_dga uint8 ;};func (_ead *Encoder )WriteTo (w _a .Writer )(int64 ,error ){const _cd ="\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f";
var _ce int64 ;for _bdd ,_deb :=range _ead ._da {_ffe ,_fg :=w .Write (_deb );if _fg !=nil {return 0,_cb .Wrapf (_fg ,_cd ,"\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b",_bdd );
};_ce +=int64 (_ffe );};_ead ._eb =_ead ._eb [:_ead ._ad ];_bbcb ,_dd :=w .Write (_ead ._eb );if _dd !=nil {return 0,_cb .Wrap (_dd ,_cd ,"\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073");};_ce +=int64 (_bbcb );return _ce ,nil ;
};func (_bec *Encoder )Flush (){_bec ._ad =0;_bec ._da =nil ;_bec ._ec =-1};func (_ffg *Encoder )rBlock (){if _ffg ._ec >=0{_ffg .emit ();};_ffg ._ec ++;_ffg ._gf =uint8 (_ffg ._gea >>20);_ffg ._gea &=0xfffff;_ffg ._geb =7;};func (_dad *Encoder )lBlock (){if _dad ._ec >=0{_dad .emit ();
};_dad ._ec ++;_dad ._gf =uint8 (_dad ._gea >>19);_dad ._gea &=0x7ffff;_dad ._geb =8;};type Encoder struct{_gea uint32 ;_eg uint16 ;_geb ,_gf uint8 ;_ec int ;_dg int ;_da [][]byte ;_eb []byte ;_ad int ;_ga *codingContext ;_ac [13]*codingContext ;_dag *codingContext ;
};func (_abf *Encoder )emit (){if _abf ._ad ==_ceg {_abf ._da =append (_abf ._da ,_abf ._eb );_abf ._eb =make ([]byte ,_ceg );_abf ._ad =0;};_abf ._eb [_abf ._ad ]=_abf ._gf ;_abf ._ad ++;};type Class int ;func (_deba *Encoder )encodeBit (_gba *codingContext ,_dge uint32 ,_cac uint8 )error {const _dca ="\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074";
_deba ._dg ++;if _dge >=uint32 (len (_gba ._ff )){return _cb .Errorf (_dca ,"\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027",_dge );
};_feg :=_gba ._ff [_dge ];_gfag :=_gba .mps (_dge );_gef :=_bage [_feg ]._ecb ;_ae .Log .Trace ("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064",_deba ._dg ,_cac ,_feg ,_gfag ,_gef ,_deba ._eg ,_deba ._gea ,_deba ._geb ,_deba ._gf ,_deba ._ec );
if _cac ==0{_deba .code0 (_gba ,_dge ,_gef ,_feg );}else {_deba .code1 (_gba ,_dge ,_gef ,_feg );};return nil ;};func (_cfb *Encoder )byteOut (){if _cfb ._gf ==0xff{_cfb .rBlock ();return ;};if _cfb ._gea < 0x8000000{_cfb .lBlock ();return ;};_cfb ._gf ++;
if _cfb ._gf !=0xff{_cfb .lBlock ();return ;};_cfb ._gea &=0x7ffffff;_cfb .rBlock ();};var _ _a .WriterTo =&Encoder {};func (_ffbd *Encoder )code0 (_gaa *codingContext ,_eef uint32 ,_bff uint16 ,_agd byte ){if _gaa .mps (_eef )==0{_ffbd .codeMPS (_gaa ,_eef ,_bff ,_agd );
}else {_ffbd .codeLPS (_gaa ,_eef ,_bff ,_agd );};};func (_cda *Encoder )encodeIAID (_fegg ,_eadb int )error {if _cda ._dag ==nil {_cda ._dag =_bf (1<<uint (_fegg ));};_cfa :=uint32 (1<<uint32 (_fegg +1))-1;_eadb <<=uint (32-_fegg );_efb :=uint32 (1);for _gge :=0;
_gge < _fegg ;_gge ++{_bded :=_efb &_cfa ;_eee :=uint8 ((uint32 (_eadb )&0x80000000)>>31);if _bc :=_cda .encodeBit (_cda ._dag ,_bded ,_eee );_bc !=nil {return _bc ;};_efb =(_efb <<1)|uint32 (_eee );_eadb <<=1;};return nil ;};func (_cbf *Encoder )Refine (iTemp ,iTarget *_c .Bitmap ,ox ,oy int )error {for _ea :=0;
_ea < iTarget .Height ;_ea ++{var _gcf int ;_bde :=_ea +oy ;var (_dgf ,_gebg ,_gb ,_aa ,_bdea uint16 ;_aad ,_gfab ,_dae ,_egfg ,_fd byte ;);if _bde >=1&&(_bde -1)< iTemp .Height {_aad =iTemp .Data [(_bde -1)*iTemp .RowStride ];};if _bde >=0&&_bde < iTemp .Height {_gfab =iTemp .Data [_bde *iTemp .RowStride ];
};if _bde >=-1&&_bde +1< iTemp .Height {_dae =iTemp .Data [(_bde +1)*iTemp .RowStride ];};if _ea >=1{_egfg =iTarget .Data [(_ea -1)*iTarget .RowStride ];};_fd =iTarget .Data [_ea *iTarget .RowStride ];_fbd :=uint (6+ox );_dgf =uint16 (_aad >>_fbd );_gebg =uint16 (_gfab >>_fbd );
_gb =uint16 (_dae >>_fbd );_aa =uint16 (_egfg >>6);_eec :=uint (2-ox );_aad <<=_eec ;_gfab <<=_eec ;_dae <<=_eec ;_egfg <<=2;for _gcf =0;_gcf < iTarget .Width ;_gcf ++{_ada :=(_dgf <<10)|(_gebg <<7)|(_gb <<4)|(_aa <<1)|_bdea ;_caaf :=_fd >>7;_afc :=_cbf .encodeBit (_cbf ._ga ,uint32 (_ada ),_caaf );
if _afc !=nil {return _afc ;};_dgf <<=1;_gebg <<=1;_gb <<=1;_aa <<=1;_dgf |=uint16 (_aad >>7);_gebg |=uint16 (_gfab >>7);_gb |=uint16 (_dae >>7);_aa |=uint16 (_egfg >>7);_bdea =uint16 (_caaf );_ebg :=_gcf %8;_ecc :=_gcf /8+1;if _ebg ==5+ox {_aad ,_gfab ,_dae =0,0,0;
if _ecc < iTemp .RowStride &&_bde >=1&&(_bde -1)< iTemp .Height {_aad =iTemp .Data [(_bde -1)*iTemp .RowStride +_ecc ];};if _ecc < iTemp .RowStride &&_bde >=0&&_bde < iTemp .Height {_gfab =iTemp .Data [_bde *iTemp .RowStride +_ecc ];};if _ecc < iTemp .RowStride &&_bde >=-1&&(_bde +1)< iTemp .Height {_dae =iTemp .Data [(_bde +1)*iTemp .RowStride +_ecc ];
};}else {_aad <<=1;_gfab <<=1;_dae <<=1;};if _ebg ==5&&_ea >=1{_egfg =0;if _ecc < iTarget .RowStride {_egfg =iTarget .Data [(_ea -1)*iTarget .RowStride +_ecc ];};}else {_egfg <<=1;};if _ebg ==7{_fd =0;if _ecc < iTarget .RowStride {_fd =iTarget .Data [_ea *iTarget .RowStride +_ecc ];
};}else {_fd <<=1;};_dgf &=7;_gebg &=7;_gb &=7;_aa &=7;};};return nil ;};var _bage =[]state {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
const (_fbb =65536;_ceg =20*1024;);func (_afcf *Encoder )dataSize ()int {return _ceg *len (_afcf ._da )+_afcf ._ad };func (_e Class )String ()string {switch _e {case IAAI :return "\u0049\u0041\u0041\u0049";case IADH :return "\u0049\u0041\u0044\u0048";case IADS :return "\u0049\u0041\u0044\u0053";
case IADT :return "\u0049\u0041\u0044\u0054";case IADW :return "\u0049\u0041\u0044\u0057";case IAEX :return "\u0049\u0041\u0045\u0058";case IAFS :return "\u0049\u0041\u0046\u0053";case IAIT :return "\u0049\u0041\u0049\u0054";case IARDH :return "\u0049\u0041\u0052D\u0048";
case IARDW :return "\u0049\u0041\u0052D\u0057";case IARDX :return "\u0049\u0041\u0052D\u0058";case IARDY :return "\u0049\u0041\u0052D\u0059";case IARI :return "\u0049\u0041\u0052\u0049";default:return "\u0055N\u004b\u004e\u004f\u0057\u004e";};};func (_cce *Encoder )encodeOOB (_cdb Class )error {_gca :=_cce ._ac [_cdb ];
_fbc :=_cce .encodeBit (_gca ,1,1);if _fbc !=nil {return _fbc ;};_fbc =_cce .encodeBit (_gca ,3,0);if _fbc !=nil {return _fbc ;};_fbc =_cce .encodeBit (_gca ,6,0);if _fbc !=nil {return _fbc ;};_fbc =_cce .encodeBit (_gca ,12,0);if _fbc !=nil {return _fbc ;
};return nil ;};