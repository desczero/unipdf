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

package bitwise ;import (_f "encoding/binary";_ae "errors";_a "fmt";_g "github.com/unidoc/unipdf/v3/common";_e "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_dc "io";);var _ _dc .ByteWriter =&BufferedWriter {};func (_gde *Reader )ReadBit ()(_agc int ,_fba error ){_feg ,_fba :=_gde .readBool ();
if _fba !=nil {return 0,_fba ;};if _feg {_agc =1;};return _agc ,nil ;};func (_eac *Reader )RelativePosition ()int64 {return _eac ._fcca };func (_dfb *Reader )NewPartialReader (offset ,length int ,relative bool )(*Reader ,error ){if offset < 0{return nil ,_ae .New ("p\u0061\u0072\u0074\u0069\u0061\u006c\u0020\u0072\u0065\u0061\u0064\u0065\u0072\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062e \u006e\u0065\u0067a\u0074i\u0076\u0065");
};if relative {offset =_dfb ._bfe ._dgc +offset ;};if length > 0{_gbf :=len (_dfb ._bfe ._abb );if relative {_gbf =_dfb ._bfe ._daa ;};if offset +length > _gbf {return nil ,_a .Errorf ("\u0070\u0061r\u0074\u0069\u0061l\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u006f\u0066\u0066se\u0074\u0028\u0025\u0064\u0029\u002b\u006c\u0065\u006e\u0067\u0074\u0068\u0028\u0025\u0064\u0029\u003d\u0025d\u0020i\u0073\u0020\u0067\u0072\u0065\u0061ter\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u006f\u0072ig\u0069n\u0061\u006c\u0020\u0072e\u0061d\u0065r\u0020\u006ce\u006e\u0067th\u003a\u0020\u0025\u0064",offset ,length ,offset +length ,_dfb ._bfe ._daa );
};};if length < 0{_fe :=len (_dfb ._bfe ._abb );if relative {_fe =_dfb ._bfe ._daa ;};length =_fe -offset ;};return &Reader {_bfe :readerSource {_abb :_dfb ._bfe ._abb ,_daa :length ,_dgc :offset }},nil ;};func (_fad *Writer )ResetBit (){_fad ._agf =0};
func NewReader (data []byte )*Reader {return &Reader {_bfe :readerSource {_abb :data ,_daa :len (data ),_dgc :0}};};type BitWriter interface{WriteBit (_eg int )error ;WriteBits (_baf uint64 ,_df int )(_cfb int ,_ebc error );FinishByte ();SkipBits (_fcc int )error ;
};func (_af *BufferedWriter )WriteByte (bt byte )error {if _af ._bg > len (_af ._dd )-1||(_af ._bg ==len (_af ._dd )-1&&_af ._eb !=0){_af .expandIfNeeded (1);};_af .writeByte (bt );return nil ;};func (_adgc *Writer )Data ()[]byte {return _adgc ._cge };
func (_bgb *Reader )Length ()uint64 {return uint64 (_bgb ._bfe ._daa )};const (_ge =64;_b =int (^uint (0)>>1););func (_afd *BufferedWriter )writeByte (_cg byte ){switch {case _afd ._eb ==0:_afd ._dd [_afd ._bg ]=_cg ;_afd ._bg ++;case _afd ._bc :_afd ._dd [_afd ._bg ]|=_cg >>_afd ._eb ;
_afd ._bg ++;_afd ._dd [_afd ._bg ]=byte (uint16 (_cg )<<(8-_afd ._eb )&0xff);default:_afd ._dd [_afd ._bg ]|=byte (uint16 (_cg )<<_afd ._eb &0xff);_afd ._bg ++;_afd ._dd [_afd ._bg ]=_cg >>(8-_afd ._eb );};};func (_fg *Reader )ReadUint32 ()(uint32 ,error ){_aaee :=make ([]byte ,4);
_ ,_ebfg :=_fg .Read (_aaee );if _ebfg !=nil {return 0,_ebfg ;};return _f .BigEndian .Uint32 (_aaee ),nil ;};var _ BinaryWriter =&BufferedWriter {};type BinaryWriter interface{BitWriter ;_dc .Writer ;_dc .ByteWriter ;Data ()[]byte ;};func (_ea *BufferedWriter )SkipBits (skip int )error {if skip ==0{return nil ;
};_ddf :=int (_ea ._eb )+skip ;if _ddf >=0&&_ddf < 8{_ea ._eb =uint8 (_ddf );return nil ;};_ddf =int (_ea ._eb )+_ea ._bg *8+skip ;if _ddf < 0{return _e .Errorf ("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073","\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_gc :=_ddf /8;_cf :=_ddf %8;_ea ._eb =uint8 (_cf );if _be :=_gc -_ea ._bg ;_be > 0&&len (_ea ._dd )-1< _gc {if _ea ._eb !=0{_be ++;};_ea .expandIfNeeded (_be );};_ea ._bg =_gc ;return nil ;};func (_ddb *Writer )writeBit (_gdc uint8 )error {if len (_ddb ._cge )-1< _ddb ._aec {return _dc .EOF ;
};_fag :=_ddb ._agf ;if _ddb ._gag {_fag =7-_ddb ._agf ;};_ddb ._cge [_ddb ._aec ]|=byte (uint16 (_gdc <<_fag )&0xff);_ddb ._agf ++;if _ddb ._agf ==8{_ddb ._aec ++;_ddb ._agf =0;};return nil ;};func (_dca *BufferedWriter )Reset (){_dca ._dd =_dca ._dd [:0];
_dca ._bg =0;_dca ._eb =0};func (_ac *BufferedWriter )Write (d []byte )(int ,error ){_ac .expandIfNeeded (len (d ));if _ac ._eb ==0{return _ac .writeFullBytes (d ),nil ;};return _ac .writeShiftedBytes (d ),nil ;};func (_gbg *BufferedWriter )tryGrowByReslice (_agd int )bool {if _cee :=len (_gbg ._dd );
_agd <=cap (_gbg ._dd )-_cee {_gbg ._dd =_gbg ._dd [:_cee +_agd ];return true ;};return false ;};func (_aa *BufferedWriter )grow (_ag int ){if _aa ._dd ==nil &&_ag < _ge {_aa ._dd =make ([]byte ,_ag ,_ge );return ;};_aae :=len (_aa ._dd );if _aa ._eb !=0{_aae ++;
};_gfg :=cap (_aa ._dd );switch {case _ag <=_gfg /2-_aae :_g .Log .Trace ("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027",len (_aa ._dd ),cap (_aa ._dd ),_ag );
_g .Log .Trace ("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027",_gfg ,_aae );copy (_aa ._dd ,_aa ._dd [_aa .fullOffset ():]);
case _gfg > _b -_gfg -_ag :_g .Log .Error ("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065");return ;default:_dg :=make ([]byte ,2*_gfg +_ag );copy (_dg ,_aa ._dd );_aa ._dd =_dg ;};_aa ._dd =_aa ._dd [:_aae +_ag ];
};type Writer struct{_cge []byte ;_agf uint8 ;_aec int ;_gag bool ;};func (_dba *Reader )ReadBool ()(bool ,error ){return _dba .readBool ()};func (_db *Reader )ConsumeRemainingBits ()(uint64 ,error ){if _db ._fdd !=0{return _db .ReadBits (_db ._fdd );};
return 0,nil ;};func (_da *BufferedWriter )fullOffset ()int {_abg :=_da ._bg ;if _da ._eb !=0{_abg ++;};return _abg ;};type readerSource struct{_abb []byte ;_dgc int ;_daa int ;};func (_bd *BufferedWriter )FinishByte (){if _bd ._eb ==0{return ;};_bd ._eb =0;
_bd ._bg ++;};func (_bcc *BufferedWriter )Len ()int {return _bcc .byteCapacity ()};func (_bcf *Reader )readBufferByte ()(byte ,error ){if _bcf ._fcca >=int64 (_bcf ._bfe ._daa ){return 0,_dc .EOF ;};_bcf ._ad =-1;_gbd :=_bcf ._bfe ._abb [int64 (_bcf ._bfe ._dgc )+_bcf ._fcca ];
_bcf ._fcca ++;_bcf ._age =int (_gbd );return _gbd ,nil ;};func (_dcc *Reader )ReadBits (n byte )(_aad uint64 ,_adg error ){if n < _dcc ._fdd {_fde :=_dcc ._fdd -n ;_aad =uint64 (_dcc ._ga >>_fde );_dcc ._ga &=1<<_fde -1;_dcc ._fdd =_fde ;return _aad ,nil ;
};if n > _dcc ._fdd {if _dcc ._fdd > 0{_aad =uint64 (_dcc ._ga );n -=_dcc ._fdd ;};for n >=8{_fed ,_cgf :=_dcc .readBufferByte ();if _cgf !=nil {return 0,_cgf ;};_aad =_aad <<8+uint64 (_fed );n -=8;};if n > 0{if _dcc ._ga ,_adg =_dcc .readBufferByte ();
_adg !=nil {return 0,_adg ;};_debd :=8-n ;_aad =_aad <<n +uint64 (_dcc ._ga >>_debd );_dcc ._ga &=1<<_debd -1;_dcc ._fdd =_debd ;}else {_dcc ._fdd =0;};return _aad ,nil ;};_dcc ._fdd =0;return uint64 (_dcc ._ga ),nil ;};func (_dfbf *Reader )Mark (){_dfbf ._geg =_dfbf ._fcca ;
_dfbf ._dcd =_dfbf ._fdd ;_dfbf ._gfgg =_dfbf ._ga ;_dfbf ._ef =_dfbf ._age ;};func (_cfbe *Reader )Align ()(_ccc byte ){_ccc =_cfbe ._fdd ;_cfbe ._fdd =0;return _ccc };func (_gb *BufferedWriter )Data ()[]byte {return _gb ._dd };func (_gd *Reader )BitPosition ()int {return int (_gd ._fdd )};
func (_dea *BufferedWriter )writeFullBytes (_ce []byte )int {_ged :=copy (_dea ._dd [_dea .fullOffset ():],_ce );_dea ._bg +=_ged ;return _ged ;};func (_ege *Reader )Read (p []byte )(_bafc int ,_fcb error ){if _ege ._fdd ==0{return _ege .read (p );};for ;
_bafc < len (p );_bafc ++{if p [_bafc ],_fcb =_ege .readUnalignedByte ();_fcb !=nil {return 0,_fcb ;};};return _bafc ,nil ;};var _ BinaryWriter =&Writer {};func (_fbf *Writer )WriteBits (bits uint64 ,number int )(_ffg int ,_ddc error ){const _cff ="\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073";
if number < 0||number > 64{return 0,_e .Errorf (_cff ,"\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027",number );
};if number ==0{return 0,nil ;};_cfg :=number /8;if _cfg > 0{_fac :=number -_cfg *8;for _ebeg :=_cfg -1;_ebeg >=0;_ebeg --{_ebd :=byte ((bits >>uint (_ebeg *8+_fac ))&0xff);if _ddc =_fbf .WriteByte (_ebd );_ddc !=nil {return _ffg ,_e .Wrapf (_ddc ,_cff ,"\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027",_cfg -_ebeg +1);
};};number -=_cfg *8;if number ==0{return _cfg ,nil ;};};var _ced int ;for _cbc :=0;_cbc < number ;_cbc ++{if _fbf ._gag {_ced =int ((bits >>uint (number -1-_cbc ))&0x1);}else {_ced =int (bits &0x1);bits >>=1;};if _ddc =_fbf .WriteBit (_ced );_ddc !=nil {return _ffg ,_e .Wrapf (_ddc ,_cff ,"\u0062i\u0074\u003a\u0020\u0025\u0064",_cbc );
};};return _cfg ,nil ;};func (_cd *BufferedWriter )writeShiftedBytes (_dee []byte )int {for _ ,_cc :=range _dee {_cd .writeByte (_cc );};return len (_dee );};func NewWriter (data []byte )*Writer {return &Writer {_cge :data }};func (_cdc *Reader )ReadByte ()(byte ,error ){if _cdc ._fdd ==0{return _cdc .readBufferByte ();
};return _cdc .readUnalignedByte ();};type StreamReader interface{_dc .Reader ;_dc .ByteReader ;_dc .Seeker ;Align ()byte ;BitPosition ()int ;Mark ();Length ()uint64 ;ReadBit ()(int ,error );ReadBits (_abe byte )(uint64 ,error );ReadBool ()(bool ,error );
ReadUint32 ()(uint32 ,error );Reset ();AbsolutePosition ()int64 ;};func (_fgf *Reader )readUnalignedByte ()(_aee byte ,_dccd error ){_dbae :=_fgf ._fdd ;_aee =_fgf ._ga <<(8-_dbae );_fgf ._ga ,_dccd =_fgf .readBufferByte ();if _dccd !=nil {return 0,_dccd ;
};_aee |=_fgf ._ga >>_dbae ;_fgf ._ga &=1<<_dbae -1;return _aee ,nil ;};func (_cdf *Writer )FinishByte (){if _cdf ._agf ==0{return ;};_cdf ._agf =0;_cdf ._aec ++;};func (_gab *Reader )Seek (offset int64 ,whence int )(int64 ,error ){_gab ._ad =-1;_gab ._fdd =0;
_gab ._ga =0;_gab ._age =0;var _adgf int64 ;switch whence {case _dc .SeekStart :_adgf =offset ;case _dc .SeekCurrent :_adgf =_gab ._fcca +offset ;case _dc .SeekEnd :_adgf =int64 (_gab ._bfe ._daa )+offset ;default:return 0,_ae .New ("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065");
};if _adgf < 0{return 0,_ae .New ("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e");};_gab ._fcca =_adgf ;
_gab ._fdd =0;return _adgf ,nil ;};func (_dac *Reader )AbsolutePosition ()int64 {return _dac ._fcca +int64 (_dac ._bfe ._dgc )};func (_fb *BufferedWriter )WriteBit (bit int )error {if bit !=1&&bit !=0{return _e .Errorf ("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074","\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064",bit );
};if len (_fb ._dd )-1< _fb ._bg {_fb .expandIfNeeded (1);};_ab :=_fb ._eb ;if _fb ._bc {_ab =7-_fb ._eb ;};_fb ._dd [_fb ._bg ]|=byte (uint16 (bit <<_ab )&0xff);_fb ._eb ++;if _fb ._eb ==8{_fb ._bg ++;_fb ._eb =0;};return nil ;};func (_c *BufferedWriter )ResetBitIndex (){_c ._eb =0};
func (_bff *Reader )Reset (){_bff ._fcca =_bff ._geg ;_bff ._fdd =_bff ._dcd ;_bff ._ga =_bff ._gfgg ;_bff ._age =_bff ._ef ;};func (_aadd *Reader )read (_dbb []byte )(int ,error ){if _aadd ._fcca >=int64 (_aadd ._bfe ._daa ){return 0,_dc .EOF ;};_aadd ._ad =-1;
_dab :=copy (_dbb ,_aadd ._bfe ._abb [(int64 (_aadd ._bfe ._dgc )+_aadd ._fcca ):(_aadd ._bfe ._dgc +_aadd ._bfe ._daa )]);_aadd ._fcca +=int64 (_dab );return _dab ,nil ;};var _ _dc .Writer =&BufferedWriter {};func NewWriterMSB (data []byte )*Writer {return &Writer {_cge :data ,_gag :true }};
func (_dae *Writer )UseMSB ()bool {return _dae ._gag };func (_ebe *BufferedWriter )WriteBits (bits uint64 ,number int )(_bcb int ,_aba error ){const _ca ="\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s";
if number < 0||number > 64{return 0,_e .Errorf (_ca ,"\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027",number );
};_ba :=number /8;if _ba > 0{_gf :=number -_ba *8;for _ff :=_ba -1;_ff >=0;_ff --{_acb :=byte ((bits >>uint (_ff *8+_gf ))&0xff);if _aba =_ebe .WriteByte (_acb );_aba !=nil {return _bcb ,_e .Wrapf (_aba ,_ca ,"\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027",_ba -_ff +1);
};};number -=_ba *8;if number ==0{return _ba ,nil ;};};var _fc int ;for _bf :=0;_bf < number ;_bf ++{if _ebe ._bc {_fc =int ((bits >>uint (number -1-_bf ))&0x1);}else {_fc =int (bits &0x1);bits >>=1;};if _aba =_ebe .WriteBit (_fc );_aba !=nil {return _bcb ,_e .Wrapf (_aba ,_ca ,"\u0062i\u0074\u003a\u0020\u0025\u0064",_bf );
};};return _ba ,nil ;};type BufferedWriter struct{_dd []byte ;_eb uint8 ;_bg int ;_bc bool ;};func (_gded *Writer )WriteByte (c byte )error {return _gded .writeByte (c )};func (_bee *BufferedWriter )byteCapacity ()int {_ec :=len (_bee ._dd )-_bee ._bg ;
if _bee ._eb !=0{_ec --;};return _ec ;};type Reader struct{_bfe readerSource ;_ga byte ;_fdd byte ;_fcca int64 ;_age int ;_ad int ;_geg int64 ;_dcd byte ;_gfgg byte ;_ef int ;};func (_ebf *Reader )AbsoluteLength ()uint64 {return uint64 (len (_ebf ._bfe ._abb ))};
func (_gcb *Writer )SkipBits (skip int )error {const _gbe ="\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073";if skip ==0{return nil ;};_bfc :=int (_gcb ._agf )+skip ;if _bfc >=0&&_bfc < 8{_gcb ._agf =uint8 (_bfc );return nil ;
};_bfc =int (_gcb ._agf )+_gcb ._aec *8+skip ;if _bfc < 0{return _e .Errorf (_gbe ,"\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_bec :=_bfc /8;_ageg :=_bfc %8;_g .Log .Trace ("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073");
_g .Log .Trace ("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027",_gcb ._agf ,_gcb ._aec ,int (_gcb ._agf )+(_gcb ._aec )*8,len (_gcb ._cge ),cap (_gcb ._cge ));
_g .Log .Trace ("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027",skip ,_bfc ,_ageg );_gcb ._agf =uint8 (_ageg );if _dgb :=_bec -_gcb ._aec ;
_dgb > 0&&len (_gcb ._cge )-1< _bec {_g .Log .Trace ("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064",_dgb );return _e .Errorf (_gbe ,"\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_gcb ._aec =_bec ;
_g .Log .Trace ("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027",_gcb ._agf ,_gcb ._aec );return nil ;};func (_dec *Writer )WriteBit (bit int )error {switch bit {case 0,1:return _dec .writeBit (uint8 (bit ));
};return _e .Error ("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074","\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065");};func (_ece *Writer )writeByte (_dccg byte )error {if _ece ._aec > len (_ece ._cge )-1{return _dc .EOF ;
};if _ece ._aec ==len (_ece ._cge )-1&&_ece ._agf !=0{return _dc .EOF ;};if _ece ._agf ==0{_ece ._cge [_ece ._aec ]=_dccg ;_ece ._aec ++;return nil ;};if _ece ._gag {_ece ._cge [_ece ._aec ]|=_dccg >>_ece ._agf ;_ece ._aec ++;_ece ._cge [_ece ._aec ]=byte (uint16 (_dccg )<<(8-_ece ._agf )&0xff);
}else {_ece ._cge [_ece ._aec ]|=byte (uint16 (_dccg )<<_ece ._agf &0xff);_ece ._aec ++;_ece ._cge [_ece ._aec ]=_dccg >>(8-_ece ._agf );};return nil ;};func (_fee *Writer )Write (p []byte )(int ,error ){if len (p )> _fee .byteCapacity (){return 0,_dc .EOF ;
};for _ ,_aadeb :=range p {if _gdf :=_fee .writeByte (_aadeb );_gdf !=nil {return 0,_gdf ;};};return len (p ),nil ;};func (_gcd *BufferedWriter )expandIfNeeded (_fa int ){if !_gcd .tryGrowByReslice (_fa ){_gcd .grow (_fa );};};func (_abd *Writer )byteCapacity ()int {_cgc :=len (_abd ._cge )-_abd ._aec ;
if _abd ._agf !=0{_cgc --;};return _cgc ;};func (_ageb *Reader )readBool ()(_bea bool ,_ee error ){if _ageb ._fdd ==0{_ageb ._ga ,_ee =_ageb .readBufferByte ();if _ee !=nil {return false ,_ee ;};_bea =(_ageb ._ga &0x80)!=0;_ageb ._ga ,_ageb ._fdd =_ageb ._ga &0x7f,7;
return _bea ,nil ;};_ageb ._fdd --;_bea =(_ageb ._ga &(1<<_ageb ._fdd ))!=0;_ageb ._ga &=1<<_ageb ._fdd -1;return _bea ,nil ;};func BufferedMSB ()*BufferedWriter {return &BufferedWriter {_bc :true }};var (_ _dc .Reader =&Reader {};_ _dc .ByteReader =&Reader {};
_ _dc .Seeker =&Reader {};_ StreamReader =&Reader {};);