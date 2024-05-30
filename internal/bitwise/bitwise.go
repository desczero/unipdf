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

package bitwise ;import (_c "encoding/binary";_gc "errors";_d "fmt";_e "github.com/unidoc/unipdf/v3/common";_ag "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_g "io";);func (_eda *Reader )ReadBits (n byte )(_abe uint64 ,_dca error ){if n < _eda ._caf {_egg :=_eda ._caf -n ;
_abe =uint64 (_eda ._ecf >>_egg );_eda ._ecf &=1<<_egg -1;_eda ._caf =_egg ;return _abe ,nil ;};if n > _eda ._caf {if _eda ._caf > 0{_abe =uint64 (_eda ._ecf );n -=_eda ._caf ;};for n >=8{_bac ,_add :=_eda .readBufferByte ();if _add !=nil {return 0,_add ;
};_abe =_abe <<8+uint64 (_bac );n -=8;};if n > 0{if _eda ._ecf ,_dca =_eda .readBufferByte ();_dca !=nil {return 0,_dca ;};_adg :=8-n ;_abe =_abe <<n +uint64 (_eda ._ecf >>_adg );_eda ._ecf &=1<<_adg -1;_eda ._caf =_adg ;}else {_eda ._caf =0;};return _abe ,nil ;
};_eda ._caf =0;return uint64 (_eda ._ecf ),nil ;};type BufferedWriter struct{_ea []byte ;_ac uint8 ;_af int ;_aa bool ;};func (_cf *Reader )AbsoluteLength ()uint64 {return uint64 (len (_cf ._ed ._afa ))};func (_ge *BufferedWriter )Write (d []byte )(int ,error ){_ge .expandIfNeeded (len (d ));
if _ge ._ac ==0{return _ge .writeFullBytes (d ),nil ;};return _ge .writeShiftedBytes (d ),nil ;};func (_bdc *BufferedWriter )writeShiftedBytes (_abg []byte )int {for _ ,_bag :=range _abg {_bdc .writeByte (_bag );};return len (_abg );};func (_ba *BufferedWriter )WriteBits (bits uint64 ,number int )(_cea int ,_age error ){const _cd ="\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s";
if number < 0||number > 64{return 0,_ag .Errorf (_cd ,"\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027",number );
};_agg :=number /8;if _agg > 0{_da :=number -_agg *8;for _ead :=_agg -1;_ead >=0;_ead --{_fd :=byte ((bits >>uint (_ead *8+_da ))&0xff);if _age =_ba .WriteByte (_fd );_age !=nil {return _cea ,_ag .Wrapf (_age ,_cd ,"\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027",_agg -_ead +1);
};};number -=_agg *8;if number ==0{return _agg ,nil ;};};var _dg int ;for _cdd :=0;_cdd < number ;_cdd ++{if _ba ._aa {_dg =int ((bits >>uint (number -1-_cdd ))&0x1);}else {_dg =int (bits &0x1);bits >>=1;};if _age =_ba .WriteBit (_dg );_age !=nil {return _cea ,_ag .Wrapf (_age ,_cd ,"\u0062i\u0074\u003a\u0020\u0025\u0064",_cdd );
};};return _agg ,nil ;};func (_ff *Reader )BitPosition ()int {return int (_ff ._caf )};func (_fa *Writer )Write (p []byte )(int ,error ){if len (p )> _fa .byteCapacity (){return 0,_g .EOF ;};for _ ,_ccb :=range p {if _cffd :=_fa .writeByte (_ccb );_cffd !=nil {return 0,_cffd ;
};};return len (p ),nil ;};func (_egf *Reader )ReadBit ()(_cdc int ,_eb error ){_fba ,_eb :=_egf .readBool ();if _eb !=nil {return 0,_eb ;};if _fba {_cdc =1;};return _cdc ,nil ;};func (_bdd *Reader )ReadUint32 ()(uint32 ,error ){_afac :=make ([]byte ,4);
_ ,_dgd :=_bdd .Read (_afac );if _dgd !=nil {return 0,_dgd ;};return _c .BigEndian .Uint32 (_afac ),nil ;};func (_gaaf *Reader )RelativePosition ()int64 {return _gaaf ._gba };func (_bfc *Reader )Seek (offset int64 ,whence int )(int64 ,error ){_bfc ._ged =-1;
_bfc ._caf =0;_bfc ._ecf =0;_bfc ._fb =0;var _afg int64 ;switch whence {case _g .SeekStart :_afg =offset ;case _g .SeekCurrent :_afg =_bfc ._gba +offset ;case _g .SeekEnd :_afg =int64 (_bfc ._ed ._cee )+offset ;default:return 0,_gc .New ("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065");
};if _afg < 0{return 0,_gc .New ("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e");};_bfc ._gba =_afg ;
_bfc ._caf =0;return _afg ,nil ;};func (_bgc *Reader )readBool ()(_acb bool ,_ceec error ){if _bgc ._caf ==0{_bgc ._ecf ,_ceec =_bgc .readBufferByte ();if _ceec !=nil {return false ,_ceec ;};_acb =(_bgc ._ecf &0x80)!=0;_bgc ._ecf ,_bgc ._caf =_bgc ._ecf &0x7f,7;
return _acb ,nil ;};_bgc ._caf --;_acb =(_bgc ._ecf &(1<<_bgc ._caf ))!=0;_bgc ._ecf &=1<<_bgc ._caf -1;return _acb ,nil ;};func (_gfgd *Writer )SkipBits (skip int )error {const _aaf ="\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073";
if skip ==0{return nil ;};_gfa :=int (_gfgd ._bb )+skip ;if _gfa >=0&&_gfa < 8{_gfgd ._bb =uint8 (_gfa );return nil ;};_gfa =int (_gfgd ._bb )+_gfgd ._ceeg *8+skip ;if _gfa < 0{return _ag .Errorf (_aaf ,"\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_bbg :=_gfa /8;_bae :=_gfa %8;_e .Log .Trace ("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073");_e .Log .Trace ("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027",_gfgd ._bb ,_gfgd ._ceeg ,int (_gfgd ._bb )+(_gfgd ._ceeg )*8,len (_gfgd ._fef ),cap (_gfgd ._fef ));
_e .Log .Trace ("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027",skip ,_gfa ,_bae );_gfgd ._bb =uint8 (_bae );if _fcc :=_bbg -_gfgd ._ceeg ;
_fcc > 0&&len (_gfgd ._fef )-1< _bbg {_e .Log .Trace ("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064",_fcc );return _ag .Errorf (_aaf ,"\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_gfgd ._ceeg =_bbg ;
_e .Log .Trace ("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027",_gfgd ._bb ,_gfgd ._ceeg );return nil ;};func (_cg *BufferedWriter )writeFullBytes (_adb []byte )int {_ca :=copy (_cg ._ea [_cg .fullOffset ():],_adb );
_cg ._af +=_ca ;return _ca ;};func (_ab *BufferedWriter )Len ()int {return _ab .byteCapacity ()};func (_ce *BufferedWriter )Data ()[]byte {return _ce ._ea };func (_gg *Reader )Read (p []byte )(_ggg int ,_dgf error ){if _gg ._caf ==0{return _gg .read (p );
};for ;_ggg < len (p );_ggg ++{if p [_ggg ],_dgf =_gg .readUnalignedByte ();_dgf !=nil {return 0,_dgf ;};};return _ggg ,nil ;};func NewWriter (data []byte )*Writer {return &Writer {_fef :data }};var _ _g .ByteWriter =&BufferedWriter {};func (_gbac *Reader )NewPartialReader (offset ,length int ,relative bool )(*Reader ,error ){if offset < 0{return nil ,_gc .New ("p\u0061\u0072\u0074\u0069\u0061\u006c\u0020\u0072\u0065\u0061\u0064\u0065\u0072\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062e \u006e\u0065\u0067a\u0074i\u0076\u0065");
};if relative {offset =_gbac ._ed ._ace +offset ;};if length > 0{_de :=len (_gbac ._ed ._afa );if relative {_de =_gbac ._ed ._cee ;};if offset +length > _de {return nil ,_d .Errorf ("\u0070\u0061r\u0074\u0069\u0061l\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u006f\u0066\u0066se\u0074\u0028\u0025\u0064\u0029\u002b\u006c\u0065\u006e\u0067\u0074\u0068\u0028\u0025\u0064\u0029\u003d\u0025d\u0020i\u0073\u0020\u0067\u0072\u0065\u0061ter\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u006f\u0072ig\u0069n\u0061\u006c\u0020\u0072e\u0061d\u0065r\u0020\u006ce\u006e\u0067th\u003a\u0020\u0025\u0064",offset ,length ,offset +length ,_gbac ._ed ._cee );
};};if length < 0{_gee :=len (_gbac ._ed ._afa );if relative {_gee =_gbac ._ed ._cee ;};length =_gee -offset ;};return &Reader {_ed :readerSource {_afa :_gbac ._ed ._afa ,_cee :length ,_ace :offset }},nil ;};func (_edb *Reader )ConsumeRemainingBits ()(uint64 ,error ){if _edb ._caf !=0{return _edb .ReadBits (_edb ._caf );
};return 0,nil ;};func (_afc *BufferedWriter )Reset (){_afc ._ea =_afc ._ea [:0];_afc ._af =0;_afc ._ac =0};const (_ae =64;_b =int (^uint (0)>>1););type BinaryWriter interface{BitWriter ;_g .Writer ;_g .ByteWriter ;Data ()[]byte ;};func (_cc *BufferedWriter )ResetBitIndex (){_cc ._ac =0};
var _ BinaryWriter =&Writer {};func (_afe *Reader )readUnalignedByte ()(_aad byte ,_cff error ){_aac :=_afe ._caf ;_aad =_afe ._ecf <<(8-_aac );_afe ._ecf ,_cff =_afe .readBufferByte ();if _cff !=nil {return 0,_cff ;};_aad |=_afe ._ecf >>_aac ;_afe ._ecf &=1<<_aac -1;
return _aad ,nil ;};func NewWriterMSB (data []byte )*Writer {return &Writer {_fef :data ,_afga :true }};func (_dee *Reader )ReadBool ()(bool ,error ){return _dee .readBool ()};func (_gdf *Writer )writeByte (_abee byte )error {if _gdf ._ceeg > len (_gdf ._fef )-1{return _g .EOF ;
};if _gdf ._ceeg ==len (_gdf ._fef )-1&&_gdf ._bb !=0{return _g .EOF ;};if _gdf ._bb ==0{_gdf ._fef [_gdf ._ceeg ]=_abee ;_gdf ._ceeg ++;return nil ;};if _gdf ._afga {_gdf ._fef [_gdf ._ceeg ]|=_abee >>_gdf ._bb ;_gdf ._ceeg ++;_gdf ._fef [_gdf ._ceeg ]=byte (uint16 (_abee )<<(8-_gdf ._bb )&0xff);
}else {_gdf ._fef [_gdf ._ceeg ]|=byte (uint16 (_abee )<<_gdf ._bb &0xff);_gdf ._ceeg ++;_gdf ._fef [_gdf ._ceeg ]=_abee >>(8-_gdf ._bb );};return nil ;};func (_fe *BufferedWriter )WriteByte (bt byte )error {if _fe ._af > len (_fe ._ea )-1||(_fe ._af ==len (_fe ._ea )-1&&_fe ._ac !=0){_fe .expandIfNeeded (1);
};_fe .writeByte (bt );return nil ;};func (_fdd *Reader )readBufferByte ()(byte ,error ){if _fdd ._gba >=int64 (_fdd ._ed ._cee ){return 0,_g .EOF ;};_fdd ._ged =-1;_agd :=_fdd ._ed ._afa [int64 (_fdd ._ed ._ace )+_fdd ._gba ];_fdd ._gba ++;_fdd ._fb =int (_agd );
return _agd ,nil ;};func (_cgc *BufferedWriter )tryGrowByReslice (_bf int )bool {if _bc :=len (_cgc ._ea );_bf <=cap (_cgc ._ea )-_bc {_cgc ._ea =_cgc ._ea [:_bc +_bf ];return true ;};return false ;};type Writer struct{_fef []byte ;_bb uint8 ;_ceeg int ;
_afga bool ;};func (_bbd *Writer )ResetBit (){_bbd ._bb =0};type Reader struct{_ed readerSource ;_ecf byte ;_caf byte ;_gba int64 ;_fb int ;_ged int ;_fc int64 ;_cda byte ;_bee byte ;_gaa int ;};func (_ceg *BufferedWriter )FinishByte (){if _ceg ._ac ==0{return ;
};_ceg ._ac =0;_ceg ._af ++;};func (_abf *Writer )UseMSB ()bool {return _abf ._afga };var _ _g .Writer =&BufferedWriter {};type BitWriter interface{WriteBit (_gf int )error ;WriteBits (_gbe uint64 ,_bdcd int )(_dd int ,_dda error );FinishByte ();SkipBits (_bdcg int )error ;
};func (_fcg *Reader )Mark (){_fcg ._fc =_fcg ._gba ;_fcg ._cda =_fcg ._caf ;_fcg ._bee =_fcg ._ecf ;_fcg ._gaa =_fcg ._fb ;};func NewReader (data []byte )*Reader {return &Reader {_ed :readerSource {_afa :data ,_cee :len (data ),_ace :0}};};func (_eggd *Reader )ReadByte ()(byte ,error ){if _eggd ._caf ==0{return _eggd .readBufferByte ();
};return _eggd .readUnalignedByte ();};func (_bad *BufferedWriter )expandIfNeeded (_eaa int ){if !_bad .tryGrowByReslice (_eaa ){_bad .grow (_eaa );};};func (_badd *Writer )WriteByte (c byte )error {return _badd .writeByte (c )};func (_acf *Writer )byteCapacity ()int {_fcd :=len (_acf ._fef )-_acf ._ceeg ;
if _acf ._bb !=0{_fcd --;};return _fcd ;};func (_gcc *Writer )WriteBits (bits uint64 ,number int )(_fcge int ,_cab error ){const _fad ="\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073";if number < 0||number > 64{return 0,_ag .Errorf (_fad ,"\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027",number );
};if number ==0{return 0,nil ;};_fadb :=number /8;if _fadb > 0{_gd :=number -_fadb *8;for _ddf :=_fadb -1;_ddf >=0;_ddf --{_bbb :=byte ((bits >>uint (_ddf *8+_gd ))&0xff);if _cab =_gcc .WriteByte (_bbb );_cab !=nil {return _fcge ,_ag .Wrapf (_cab ,_fad ,"\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027",_fadb -_ddf +1);
};};number -=_fadb *8;if number ==0{return _fadb ,nil ;};};var _cabg int ;for _db :=0;_db < number ;_db ++{if _gcc ._afga {_cabg =int ((bits >>uint (number -1-_db ))&0x1);}else {_cabg =int (bits &0x1);bits >>=1;};if _cab =_gcc .WriteBit (_cabg );_cab !=nil {return _fcge ,_ag .Wrapf (_cab ,_fad ,"\u0062i\u0074\u003a\u0020\u0025\u0064",_db );
};};return _fadb ,nil ;};func (_gca *Reader )AbsolutePosition ()int64 {return _gca ._gba +int64 (_gca ._ed ._ace )};func (_ffg *Reader )Reset (){_ffg ._gba =_ffg ._fc ;_ffg ._caf =_ffg ._cda ;_ffg ._ecf =_ffg ._bee ;_ffg ._fb =_ffg ._gaa ;};var (_ _g .Reader =&Reader {};
_ _g .ByteReader =&Reader {};_ _g .Seeker =&Reader {};_ StreamReader =&Reader {};);type StreamReader interface{_g .Reader ;_g .ByteReader ;_g .Seeker ;Align ()byte ;BitPosition ()int ;Mark ();Length ()uint64 ;ReadBit ()(int ,error );ReadBits (_fec byte )(uint64 ,error );
ReadBool ()(bool ,error );ReadUint32 ()(uint32 ,error );Reset ();AbsolutePosition ()int64 ;};func (_gfg *Reader )read (_gggb []byte )(int ,error ){if _gfg ._gba >=int64 (_gfg ._ed ._cee ){return 0,_g .EOF ;};_gfg ._ged =-1;_ccf :=copy (_gggb ,_gfg ._ed ._afa [(int64 (_gfg ._ed ._ace )+_gfg ._gba ):(_gfg ._ed ._ace +_gfg ._ed ._cee )]);
_gfg ._gba +=int64 (_ccf );return _ccf ,nil ;};func (_be *BufferedWriter )WriteBit (bit int )error {if bit !=1&&bit !=0{return _ag .Errorf ("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074","\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064",bit );
};if len (_be ._ea )-1< _be ._af {_be .expandIfNeeded (1);};_geb :=_be ._ac ;if _be ._aa {_geb =7-_be ._ac ;};_be ._ea [_be ._af ]|=byte (uint16 (bit <<_geb )&0xff);_be ._ac ++;if _be ._ac ==8{_be ._af ++;_be ._ac =0;};return nil ;};type readerSource struct{_afa []byte ;
_ace int ;_cee int ;};func (_gecf *Writer )WriteBit (bit int )error {switch bit {case 0,1:return _gecf .writeBit (uint8 (bit ));};return _ag .Error ("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074","\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065");
};func (_df *BufferedWriter )fullOffset ()int {_gbb :=_df ._af ;if _df ._ac !=0{_gbb ++;};return _gbb ;};func (_gec *BufferedWriter )writeByte (_aab byte ){switch {case _gec ._ac ==0:_gec ._ea [_gec ._af ]=_aab ;_gec ._af ++;case _gec ._aa :_gec ._ea [_gec ._af ]|=_aab >>_gec ._ac ;
_gec ._af ++;_gec ._ea [_gec ._af ]=byte (uint16 (_aab )<<(8-_gec ._ac )&0xff);default:_gec ._ea [_gec ._af ]|=byte (uint16 (_aab )<<_gec ._ac &0xff);_gec ._af ++;_gec ._ea [_gec ._af ]=_aab >>(8-_gec ._ac );};};var _ BinaryWriter =&BufferedWriter {};func (_ad *BufferedWriter )byteCapacity ()int {_dc :=len (_ad ._ea )-_ad ._af ;
if _ad ._ac !=0{_dc --;};return _dc ;};func (_daf *BufferedWriter )grow (_aba int ){if _daf ._ea ==nil &&_aba < _ae {_daf ._ea =make ([]byte ,_aba ,_ae );return ;};_eg :=len (_daf ._ea );if _daf ._ac !=0{_eg ++;};_ec :=cap (_daf ._ea );switch {case _aba <=_ec /2-_eg :_e .Log .Trace ("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027",len (_daf ._ea ),cap (_daf ._ea ),_aba );
_e .Log .Trace ("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027",_ec ,_eg );copy (_daf ._ea ,_daf ._ea [_daf .fullOffset ():]);
case _ec > _b -_ec -_aba :_e .Log .Error ("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065");return ;default:_bge :=make ([]byte ,2*_ec +_aba );copy (_bge ,_daf ._ea );_daf ._ea =_bge ;};_daf ._ea =_daf ._ea [:_eg +_aba ];
};func (_ggc *Writer )writeBit (_dfed uint8 )error {if len (_ggc ._fef )-1< _ggc ._ceeg {return _g .EOF ;};_dfeb :=_ggc ._bb ;if _ggc ._afga {_dfeb =7-_ggc ._bb ;};_ggc ._fef [_ggc ._ceeg ]|=byte (uint16 (_dfed <<_dfeb )&0xff);_ggc ._bb ++;if _ggc ._bb ==8{_ggc ._ceeg ++;
_ggc ._bb =0;};return nil ;};func (_ffc *Writer )Data ()[]byte {return _ffc ._fef };func BufferedMSB ()*BufferedWriter {return &BufferedWriter {_aa :true }};func (_cac *Writer )FinishByte (){if _cac ._bb ==0{return ;};_cac ._bb =0;_cac ._ceeg ++;};func (_dfe *Reader )Align ()(_gfe byte ){_gfe =_dfe ._caf ;
_dfe ._caf =0;return _gfe };func (_cddg *Reader )Length ()uint64 {return uint64 (_cddg ._ed ._cee )};func (_gb *BufferedWriter )SkipBits (skip int )error {if skip ==0{return nil ;};_aef :=int (_gb ._ac )+skip ;if _aef >=0&&_aef < 8{_gb ._ac =uint8 (_aef );
return nil ;};_aef =int (_gb ._ac )+_gb ._af *8+skip ;if _aef < 0{return _ag .Errorf ("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073","\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_bg :=_aef /8;_f :=_aef %8;_gb ._ac =uint8 (_f );if _ccc :=_bg -_gb ._af ;_ccc > 0&&len (_gb ._ea )-1< _bg {if _gb ._ac !=0{_ccc ++;};_gb .expandIfNeeded (_ccc );};_gb ._af =_bg ;return nil ;};