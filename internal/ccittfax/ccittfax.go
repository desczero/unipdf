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

package ccittfax ;import (_f "errors";_a "github.com/unidoc/unipdf/v3/internal/bitwise";_c "io";_bb "math";);type Decoder struct{_ecc int ;_acc int ;_ge int ;_ee []byte ;_bge int ;_ccf bool ;_ecda bool ;_dg bool ;_be bool ;_cca bool ;_aee bool ;_fcd bool ;
_cd int ;_bec int ;_gf []int ;_bfg []int ;_egae int ;_cbb int ;_fge int ;_cfb int ;_dca *_a .Reader ;_cef tiffType ;_bdb error ;};func init (){_e =&treeNode {_gde :true ,_ace :_fe };_g =&treeNode {_ace :_gb ,_fdc :_e };_g ._ecfe =_g ;_eg =&tree {_fgeb :&treeNode {}};
if _gbb :=_eg .fillWithNode (12,0,_g );_gbb !=nil {panic (_gbb .Error ());};if _cb :=_eg .fillWithNode (12,1,_e );_cb !=nil {panic (_cb .Error ());};_ff =&tree {_fgeb :&treeNode {}};for _ca :=0;_ca < len (_afc );_ca ++{for _af :=0;_af < len (_afc [_ca ]);
_af ++{if _db :=_ff .fill (_ca +2,int (_afc [_ca ][_af ]),int (_egb [_ca ][_af ]));_db !=nil {panic (_db .Error ());};};};if _ba :=_ff .fillWithNode (12,0,_g );_ba !=nil {panic (_ba .Error ());};if _gc :=_ff .fillWithNode (12,1,_e );_gc !=nil {panic (_gc .Error ());
};_fg =&tree {_fgeb :&treeNode {}};for _ac :=0;_ac < len (_cc );_ac ++{for _fga :=0;_fga < len (_cc [_ac ]);_fga ++{if _caa :=_fg .fill (_ac +4,int (_cc [_ac ][_fga ]),int (_fdf [_ac ][_fga ]));_caa !=nil {panic (_caa .Error ());};};};if _fa :=_fg .fillWithNode (12,0,_g );
_fa !=nil {panic (_fa .Error ());};if _gbe :=_fg .fillWithNode (12,1,_e );_gbe !=nil {panic (_gbe .Error ());};_d =&tree {_fgeb :&treeNode {}};if _gd :=_d .fill (4,1,_bf );_gd !=nil {panic (_gd .Error ());};if _bbf :=_d .fill (3,1,_aa );_bbf !=nil {panic (_bbf .Error ());
};if _fab :=_d .fill (1,1,0);_fab !=nil {panic (_fab .Error ());};if _cbe :=_d .fill (3,3,1);_cbe !=nil {panic (_cbe .Error ());};if _bd :=_d .fill (6,3,2);_bd !=nil {panic (_bd .Error ());};if _gdg :=_d .fill (7,3,3);_gdg !=nil {panic (_gdg .Error ());
};if _ad :=_d .fill (3,2,-1);_ad !=nil {panic (_ad .Error ());};if _fd :=_d .fill (6,2,-2);_fd !=nil {panic (_fd .Error ());};if _cg :=_d .fill (7,2,-3);_cg !=nil {panic (_cg .Error ());};};var _cc =[...][]uint16 {{0x7,0x8,0xb,0xc,0xe,0xf},{0x12,0x13,0x14,0x1b,0x7,0x8},{0x17,0x18,0x2a,0x2b,0x3,0x34,0x35,0x7,0x8},{0x13,0x17,0x18,0x24,0x27,0x28,0x2b,0x3,0x37,0x4,0x8,0xc},{0x12,0x13,0x14,0x15,0x16,0x17,0x1a,0x1b,0x2,0x24,0x25,0x28,0x29,0x2a,0x2b,0x2c,0x2d,0x3,0x32,0x33,0x34,0x35,0x36,0x37,0x4,0x4a,0x4b,0x5,0x52,0x53,0x54,0x55,0x58,0x59,0x5a,0x5b,0x64,0x65,0x67,0x68,0xa,0xb},{0x98,0x99,0x9a,0x9b,0xcc,0xcd,0xd2,0xd3,0xd4,0xd5,0xd6,0xd7,0xd8,0xd9,0xda,0xdb},{},{0x8,0xc,0xd},{0x12,0x13,0x14,0x15,0x16,0x17,0x1c,0x1d,0x1e,0x1f}};
func _dddg (_aaaa []byte ,_degb int ,_feea code )([]byte ,int ){_dbfa :=0;for _dbfa < _feea .BitsWritten {_beaa :=_degb /8;_cdef :=_degb %8;if _beaa >=len (_aaaa ){_aaaa =append (_aaaa ,0);};_edg :=8-_cdef ;_cfc :=_feea .BitsWritten -_dbfa ;if _edg > _cfc {_edg =_cfc ;
};if _dbfa < 8{_aaaa [_beaa ]=_aaaa [_beaa ]|byte (_feea .Code >>uint (8+_cdef -_dbfa ))&_bfc [8-_edg -_cdef ];}else {_aaaa [_beaa ]=_aaaa [_beaa ]|(byte (_feea .Code <<uint (_dbfa -8))&_bfc [8-_edg ])>>uint (_cdef );};_degb +=_edg ;_dbfa +=_edg ;};return _aaaa ,_degb ;
};type code struct{Code uint16 ;BitsWritten int ;};func (_bcdb *treeNode )walk (_facb bool )*treeNode {if _facb {return _bcdb ._fdc ;};return _bcdb ._ecfe ;};func (_abd *Decoder )tryFetchEOL1 ()(bool ,error ){_bdd ,_ade :=_abd ._dca .ReadBits (13);if _ade !=nil {return false ,_ade ;
};return _bdd ==0x3,nil ;};const (_ tiffType =iota ;_bc ;_ffe ;_bg ;);func init (){_aca =make (map[int ]code );_aca [0]=code {Code :13<<8|3<<6,BitsWritten :10};_aca [1]=code {Code :2<<(5+8),BitsWritten :3};_aca [2]=code {Code :3<<(6+8),BitsWritten :2};
_aca [3]=code {Code :2<<(6+8),BitsWritten :2};_aca [4]=code {Code :3<<(5+8),BitsWritten :3};_aca [5]=code {Code :3<<(4+8),BitsWritten :4};_aca [6]=code {Code :2<<(4+8),BitsWritten :4};_aca [7]=code {Code :3<<(3+8),BitsWritten :5};_aca [8]=code {Code :5<<(2+8),BitsWritten :6};
_aca [9]=code {Code :4<<(2+8),BitsWritten :6};_aca [10]=code {Code :4<<(1+8),BitsWritten :7};_aca [11]=code {Code :5<<(1+8),BitsWritten :7};_aca [12]=code {Code :7<<(1+8),BitsWritten :7};_aca [13]=code {Code :4<<8,BitsWritten :8};_aca [14]=code {Code :7<<8,BitsWritten :8};
_aca [15]=code {Code :12<<8,BitsWritten :9};_aca [16]=code {Code :5<<8|3<<6,BitsWritten :10};_aca [17]=code {Code :6<<8,BitsWritten :10};_aca [18]=code {Code :2<<8,BitsWritten :10};_aca [19]=code {Code :12<<8|7<<5,BitsWritten :11};_aca [20]=code {Code :13<<8,BitsWritten :11};
_aca [21]=code {Code :13<<8|4<<5,BitsWritten :11};_aca [22]=code {Code :6<<8|7<<5,BitsWritten :11};_aca [23]=code {Code :5<<8,BitsWritten :11};_aca [24]=code {Code :2<<8|7<<5,BitsWritten :11};_aca [25]=code {Code :3<<8,BitsWritten :11};_aca [26]=code {Code :12<<8|10<<4,BitsWritten :12};
_aca [27]=code {Code :12<<8|11<<4,BitsWritten :12};_aca [28]=code {Code :12<<8|12<<4,BitsWritten :12};_aca [29]=code {Code :12<<8|13<<4,BitsWritten :12};_aca [30]=code {Code :6<<8|8<<4,BitsWritten :12};_aca [31]=code {Code :6<<8|9<<4,BitsWritten :12};_aca [32]=code {Code :6<<8|10<<4,BitsWritten :12};
_aca [33]=code {Code :6<<8|11<<4,BitsWritten :12};_aca [34]=code {Code :13<<8|2<<4,BitsWritten :12};_aca [35]=code {Code :13<<8|3<<4,BitsWritten :12};_aca [36]=code {Code :13<<8|4<<4,BitsWritten :12};_aca [37]=code {Code :13<<8|5<<4,BitsWritten :12};_aca [38]=code {Code :13<<8|6<<4,BitsWritten :12};
_aca [39]=code {Code :13<<8|7<<4,BitsWritten :12};_aca [40]=code {Code :6<<8|12<<4,BitsWritten :12};_aca [41]=code {Code :6<<8|13<<4,BitsWritten :12};_aca [42]=code {Code :13<<8|10<<4,BitsWritten :12};_aca [43]=code {Code :13<<8|11<<4,BitsWritten :12};
_aca [44]=code {Code :5<<8|4<<4,BitsWritten :12};_aca [45]=code {Code :5<<8|5<<4,BitsWritten :12};_aca [46]=code {Code :5<<8|6<<4,BitsWritten :12};_aca [47]=code {Code :5<<8|7<<4,BitsWritten :12};_aca [48]=code {Code :6<<8|4<<4,BitsWritten :12};_aca [49]=code {Code :6<<8|5<<4,BitsWritten :12};
_aca [50]=code {Code :5<<8|2<<4,BitsWritten :12};_aca [51]=code {Code :5<<8|3<<4,BitsWritten :12};_aca [52]=code {Code :2<<8|4<<4,BitsWritten :12};_aca [53]=code {Code :3<<8|7<<4,BitsWritten :12};_aca [54]=code {Code :3<<8|8<<4,BitsWritten :12};_aca [55]=code {Code :2<<8|7<<4,BitsWritten :12};
_aca [56]=code {Code :2<<8|8<<4,BitsWritten :12};_aca [57]=code {Code :5<<8|8<<4,BitsWritten :12};_aca [58]=code {Code :5<<8|9<<4,BitsWritten :12};_aca [59]=code {Code :2<<8|11<<4,BitsWritten :12};_aca [60]=code {Code :2<<8|12<<4,BitsWritten :12};_aca [61]=code {Code :5<<8|10<<4,BitsWritten :12};
_aca [62]=code {Code :6<<8|6<<4,BitsWritten :12};_aca [63]=code {Code :6<<8|7<<4,BitsWritten :12};_cf =make (map[int ]code );_cf [0]=code {Code :53<<8,BitsWritten :8};_cf [1]=code {Code :7<<(2+8),BitsWritten :6};_cf [2]=code {Code :7<<(4+8),BitsWritten :4};
_cf [3]=code {Code :8<<(4+8),BitsWritten :4};_cf [4]=code {Code :11<<(4+8),BitsWritten :4};_cf [5]=code {Code :12<<(4+8),BitsWritten :4};_cf [6]=code {Code :14<<(4+8),BitsWritten :4};_cf [7]=code {Code :15<<(4+8),BitsWritten :4};_cf [8]=code {Code :19<<(3+8),BitsWritten :5};
_cf [9]=code {Code :20<<(3+8),BitsWritten :5};_cf [10]=code {Code :7<<(3+8),BitsWritten :5};_cf [11]=code {Code :8<<(3+8),BitsWritten :5};_cf [12]=code {Code :8<<(2+8),BitsWritten :6};_cf [13]=code {Code :3<<(2+8),BitsWritten :6};_cf [14]=code {Code :52<<(2+8),BitsWritten :6};
_cf [15]=code {Code :53<<(2+8),BitsWritten :6};_cf [16]=code {Code :42<<(2+8),BitsWritten :6};_cf [17]=code {Code :43<<(2+8),BitsWritten :6};_cf [18]=code {Code :39<<(1+8),BitsWritten :7};_cf [19]=code {Code :12<<(1+8),BitsWritten :7};_cf [20]=code {Code :8<<(1+8),BitsWritten :7};
_cf [21]=code {Code :23<<(1+8),BitsWritten :7};_cf [22]=code {Code :3<<(1+8),BitsWritten :7};_cf [23]=code {Code :4<<(1+8),BitsWritten :7};_cf [24]=code {Code :40<<(1+8),BitsWritten :7};_cf [25]=code {Code :43<<(1+8),BitsWritten :7};_cf [26]=code {Code :19<<(1+8),BitsWritten :7};
_cf [27]=code {Code :36<<(1+8),BitsWritten :7};_cf [28]=code {Code :24<<(1+8),BitsWritten :7};_cf [29]=code {Code :2<<8,BitsWritten :8};_cf [30]=code {Code :3<<8,BitsWritten :8};_cf [31]=code {Code :26<<8,BitsWritten :8};_cf [32]=code {Code :27<<8,BitsWritten :8};
_cf [33]=code {Code :18<<8,BitsWritten :8};_cf [34]=code {Code :19<<8,BitsWritten :8};_cf [35]=code {Code :20<<8,BitsWritten :8};_cf [36]=code {Code :21<<8,BitsWritten :8};_cf [37]=code {Code :22<<8,BitsWritten :8};_cf [38]=code {Code :23<<8,BitsWritten :8};
_cf [39]=code {Code :40<<8,BitsWritten :8};_cf [40]=code {Code :41<<8,BitsWritten :8};_cf [41]=code {Code :42<<8,BitsWritten :8};_cf [42]=code {Code :43<<8,BitsWritten :8};_cf [43]=code {Code :44<<8,BitsWritten :8};_cf [44]=code {Code :45<<8,BitsWritten :8};
_cf [45]=code {Code :4<<8,BitsWritten :8};_cf [46]=code {Code :5<<8,BitsWritten :8};_cf [47]=code {Code :10<<8,BitsWritten :8};_cf [48]=code {Code :11<<8,BitsWritten :8};_cf [49]=code {Code :82<<8,BitsWritten :8};_cf [50]=code {Code :83<<8,BitsWritten :8};
_cf [51]=code {Code :84<<8,BitsWritten :8};_cf [52]=code {Code :85<<8,BitsWritten :8};_cf [53]=code {Code :36<<8,BitsWritten :8};_cf [54]=code {Code :37<<8,BitsWritten :8};_cf [55]=code {Code :88<<8,BitsWritten :8};_cf [56]=code {Code :89<<8,BitsWritten :8};
_cf [57]=code {Code :90<<8,BitsWritten :8};_cf [58]=code {Code :91<<8,BitsWritten :8};_cf [59]=code {Code :74<<8,BitsWritten :8};_cf [60]=code {Code :75<<8,BitsWritten :8};_cf [61]=code {Code :50<<8,BitsWritten :8};_cf [62]=code {Code :51<<8,BitsWritten :8};
_cf [63]=code {Code :52<<8,BitsWritten :8};_ef =make (map[int ]code );_ef [64]=code {Code :3<<8|3<<6,BitsWritten :10};_ef [128]=code {Code :12<<8|8<<4,BitsWritten :12};_ef [192]=code {Code :12<<8|9<<4,BitsWritten :12};_ef [256]=code {Code :5<<8|11<<4,BitsWritten :12};
_ef [320]=code {Code :3<<8|3<<4,BitsWritten :12};_ef [384]=code {Code :3<<8|4<<4,BitsWritten :12};_ef [448]=code {Code :3<<8|5<<4,BitsWritten :12};_ef [512]=code {Code :3<<8|12<<3,BitsWritten :13};_ef [576]=code {Code :3<<8|13<<3,BitsWritten :13};_ef [640]=code {Code :2<<8|10<<3,BitsWritten :13};
_ef [704]=code {Code :2<<8|11<<3,BitsWritten :13};_ef [768]=code {Code :2<<8|12<<3,BitsWritten :13};_ef [832]=code {Code :2<<8|13<<3,BitsWritten :13};_ef [896]=code {Code :3<<8|18<<3,BitsWritten :13};_ef [960]=code {Code :3<<8|19<<3,BitsWritten :13};_ef [1024]=code {Code :3<<8|20<<3,BitsWritten :13};
_ef [1088]=code {Code :3<<8|21<<3,BitsWritten :13};_ef [1152]=code {Code :3<<8|22<<3,BitsWritten :13};_ef [1216]=code {Code :119<<3,BitsWritten :13};_ef [1280]=code {Code :2<<8|18<<3,BitsWritten :13};_ef [1344]=code {Code :2<<8|19<<3,BitsWritten :13};_ef [1408]=code {Code :2<<8|20<<3,BitsWritten :13};
_ef [1472]=code {Code :2<<8|21<<3,BitsWritten :13};_ef [1536]=code {Code :2<<8|26<<3,BitsWritten :13};_ef [1600]=code {Code :2<<8|27<<3,BitsWritten :13};_ef [1664]=code {Code :3<<8|4<<3,BitsWritten :13};_ef [1728]=code {Code :3<<8|5<<3,BitsWritten :13};
_ec =make (map[int ]code );_ec [64]=code {Code :27<<(3+8),BitsWritten :5};_ec [128]=code {Code :18<<(3+8),BitsWritten :5};_ec [192]=code {Code :23<<(2+8),BitsWritten :6};_ec [256]=code {Code :55<<(1+8),BitsWritten :7};_ec [320]=code {Code :54<<8,BitsWritten :8};
_ec [384]=code {Code :55<<8,BitsWritten :8};_ec [448]=code {Code :100<<8,BitsWritten :8};_ec [512]=code {Code :101<<8,BitsWritten :8};_ec [576]=code {Code :104<<8,BitsWritten :8};_ec [640]=code {Code :103<<8,BitsWritten :8};_ec [704]=code {Code :102<<8,BitsWritten :9};
_ec [768]=code {Code :102<<8|1<<7,BitsWritten :9};_ec [832]=code {Code :105<<8,BitsWritten :9};_ec [896]=code {Code :105<<8|1<<7,BitsWritten :9};_ec [960]=code {Code :106<<8,BitsWritten :9};_ec [1024]=code {Code :106<<8|1<<7,BitsWritten :9};_ec [1088]=code {Code :107<<8,BitsWritten :9};
_ec [1152]=code {Code :107<<8|1<<7,BitsWritten :9};_ec [1216]=code {Code :108<<8,BitsWritten :9};_ec [1280]=code {Code :108<<8|1<<7,BitsWritten :9};_ec [1344]=code {Code :109<<8,BitsWritten :9};_ec [1408]=code {Code :109<<8|1<<7,BitsWritten :9};_ec [1472]=code {Code :76<<8,BitsWritten :9};
_ec [1536]=code {Code :76<<8|1<<7,BitsWritten :9};_ec [1600]=code {Code :77<<8,BitsWritten :9};_ec [1664]=code {Code :24<<(2+8),BitsWritten :6};_ec [1728]=code {Code :77<<8|1<<7,BitsWritten :9};_fad =make (map[int ]code );_fad [1792]=code {Code :1<<8,BitsWritten :11};
_fad [1856]=code {Code :1<<8|4<<5,BitsWritten :11};_fad [1920]=code {Code :1<<8|5<<5,BitsWritten :11};_fad [1984]=code {Code :1<<8|2<<4,BitsWritten :12};_fad [2048]=code {Code :1<<8|3<<4,BitsWritten :12};_fad [2112]=code {Code :1<<8|4<<4,BitsWritten :12};
_fad [2176]=code {Code :1<<8|5<<4,BitsWritten :12};_fad [2240]=code {Code :1<<8|6<<4,BitsWritten :12};_fad [2304]=code {Code :1<<8|7<<4,BitsWritten :12};_fad [2368]=code {Code :1<<8|12<<4,BitsWritten :12};_fad [2432]=code {Code :1<<8|13<<4,BitsWritten :12};
_fad [2496]=code {Code :1<<8|14<<4,BitsWritten :12};_fad [2560]=code {Code :1<<8|15<<4,BitsWritten :12};_bfc =make (map[int ]byte );_bfc [0]=0xFF;_bfc [1]=0xFE;_bfc [2]=0xFC;_bfc [3]=0xF8;_bfc [4]=0xF0;_bfc [5]=0xE0;_bfc [6]=0xC0;_bfc [7]=0x80;_bfc [8]=0x00;
};var _egb =[...][]uint16 {{3,2},{1,4},{6,5},{7},{9,8},{10,11,12},{13,14},{15},{16,17,0,18,64},{24,25,23,22,19,20,21,1792,1856,1920},{1984,2048,2112,2176,2240,2304,2368,2432,2496,2560,52,55,56,59,60,320,384,448,53,54,50,51,44,45,46,47,57,58,61,256,48,49,62,63,30,31,32,33,40,41,128,192,26,27,28,29,34,35,36,37,38,39,42,43},{640,704,768,832,1280,1344,1408,1472,1536,1600,1664,1728,512,576,896,960,1024,1088,1152,1216}};
func _dbg (_febd []byte ,_ccc int ,_gaa int ,_afcf bool )([]byte ,int ){var (_afb code ;_gac bool ;);for !_gac {_afb ,_gaa ,_gac =_beb (_gaa ,_afcf );_febd ,_ccc =_dddg (_febd ,_ccc ,_afb );};return _febd ,_ccc ;};func NewDecoder (data []byte ,options DecodeOptions )(*Decoder ,error ){_cfe :=&Decoder {_dca :_a .NewReader (data ),_ecc :options .Columns ,_acc :options .Rows ,_bge :options .DamagedRowsBeforeError ,_ee :make ([]byte ,(options .Columns +7)/8),_gf :make ([]int ,options .Columns +2),_bfg :make ([]int ,options .Columns +2),_fcd :options .EncodedByteAligned ,_be :options .BlackIsOne ,_cca :options .EndOfLine ,_aee :options .EndOfBlock };
switch {case options .K ==0:_cfe ._cef =_ffe ;if len (data )< 20{return nil ,_f .New ("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m");};_dd :=data [:20];if _dd [0]!=0||(_dd [1]>>4!=1&&_dd [1]!=1){_cfe ._cef =_bc ;
_bcg :=(uint16 (_dd [0])<<8+uint16 (_dd [1]&0xff))>>4;for _cdd :=12;_cdd < 160;_cdd ++{_bcg =(_bcg <<1)+uint16 ((_dd [_cdd /8]>>uint16 (7-(_cdd %8)))&0x01);if _bcg &0xfff==1{_cfe ._cef =_ffe ;break ;};};};case options .K < 0:_cfe ._cef =_bg ;case options .K > 0:_cfe ._cef =_ffe ;
_cfe ._ccf =true ;};switch _cfe ._cef {case _bc ,_ffe ,_bg :default:return nil ,_f .New ("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065");
};return _cfe ,nil ;};func _bda (_ebf []byte ,_fcfc ,_eea ,_gef int )([]byte ,int ){_cdbc :=_cfa (_eea ,_gef );_ebf ,_fcfc =_dddg (_ebf ,_fcfc ,_cdbc );return _ebf ,_fcfc ;};func (_daf *Decoder )decodeRow ()(_fcce error ){if !_daf ._aee &&_daf ._acc > 0&&_daf ._acc ==_daf ._ge {return _c .EOF ;
};switch _daf ._cef {case _bc :_fcce =_daf .decodeRowType2 ();case _ffe :_fcce =_daf .decodeRowType4 ();case _bg :_fcce =_daf .decodeRowType6 ();};if _fcce !=nil {return _fcce ;};_dgc :=0;_gfa :=true ;_daf ._cfb =0;for _feg :=0;_feg < _daf ._cbb ;_feg ++{_bdc :=_daf ._ecc ;
if _feg !=_daf ._cbb {_bdc =_daf ._bfg [_feg ];};if _bdc > _daf ._ecc {_bdc =_daf ._ecc ;};_ceg :=_dgc /8;for _dgc %8!=0&&_bdc -_dgc > 0{var _eca byte ;if !_gfa {_eca =1<<uint (7-(_dgc %8));};_daf ._ee [_ceg ]|=_eca ;_dgc ++;};if _dgc %8==0{_ceg =_dgc /8;
var _bdf byte ;if !_gfa {_bdf =0xff;};for _bdc -_dgc > 7{_daf ._ee [_ceg ]=_bdf ;_dgc +=8;_ceg ++;};};for _bdc -_dgc > 0{if _dgc %8==0{_daf ._ee [_ceg ]=0;};var _fea byte ;if !_gfa {_fea =1<<uint (7-(_dgc %8));};_daf ._ee [_ceg ]|=_fea ;_dgc ++;};_gfa =!_gfa ;
};if _dgc !=_daf ._ecc {return _f .New ("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068");
};_daf ._cd =(_dgc +7)/8;_daf ._ge ++;return nil ;};func _dcbf (_gfac ,_gaf []byte ,_ccfd ,_gcgg ,_gge int )([]byte ,int ,int ){_cecf :=_adf (_gfac ,_gge );_cbd :=_gcgg >=0&&_gfac [_gcgg ]==_deg ||_gcgg ==-1;_gaf ,_ccfd =_dddg (_gaf ,_ccfd ,_da );var _acad int ;
if _gcgg > -1{_acad =_gge -_gcgg ;}else {_acad =_gge -_gcgg -1;};_gaf ,_ccfd =_dbg (_gaf ,_ccfd ,_acad ,_cbd );_cbd =!_cbd ;_dbca :=_cecf -_gge ;_gaf ,_ccfd =_dbg (_gaf ,_ccfd ,_dbca ,_cbd );_gcgg =_cecf ;return _gaf ,_ccfd ,_gcgg ;};var (_aca map[int ]code ;
_cf map[int ]code ;_ef map[int ]code ;_ec map[int ]code ;_fad map[int ]code ;_bfc map[int ]byte ;_aaa =code {Code :1<<4,BitsWritten :12};_ecd =code {Code :3<<3,BitsWritten :13};_cce =code {Code :2<<3,BitsWritten :13};_dc =code {Code :1<<12,BitsWritten :4};
_da =code {Code :1<<13,BitsWritten :3};_bac =code {Code :1<<15,BitsWritten :1};_fc =code {Code :3<<13,BitsWritten :3};_dcb =code {Code :3<<10,BitsWritten :6};_adg =code {Code :3<<9,BitsWritten :7};_ecb =code {Code :2<<13,BitsWritten :3};_ega =code {Code :2<<10,BitsWritten :6};
_adb =code {Code :2<<9,BitsWritten :7};);func (_gfda *Decoder )decodeG32D ()error {_gfda ._egae =_gfda ._cbb ;_gfda ._bfg ,_gfda ._gf =_gfda ._gf ,_gfda ._bfg ;_bbe :=true ;var (_beg bool ;_cbg int ;_dad error ;);_gfda ._cbb =0;_bdcf :for _cbg < _gfda ._ecc {_ecbg :=_d ._fgeb ;
for {_beg ,_dad =_gfda ._dca .ReadBool ();if _dad !=nil {return _dad ;};_ecbg =_ecbg .walk (_beg );if _ecbg ==nil {continue _bdcf ;};if !_ecbg ._gde {continue ;};switch _ecbg ._ace {case _aa :var _gda int ;if _bbe {_gda ,_dad =_gfda .decodeRun (_fg );}else {_gda ,_dad =_gfda .decodeRun (_ff );
};if _dad !=nil {return _dad ;};_cbg +=_gda ;_gfda ._bfg [_gfda ._cbb ]=_cbg ;_gfda ._cbb ++;if _bbe {_gda ,_dad =_gfda .decodeRun (_ff );}else {_gda ,_dad =_gfda .decodeRun (_fg );};if _dad !=nil {return _dad ;};_cbg +=_gda ;_gfda ._bfg [_gfda ._cbb ]=_cbg ;
_gfda ._cbb ++;case _bf :_gdc :=_gfda .getNextChangingElement (_cbg ,_bbe )+1;if _gdc >=_gfda ._egae {_cbg =_gfda ._ecc ;}else {_cbg =_gfda ._gf [_gdc ];};default:_ecf :=_gfda .getNextChangingElement (_cbg ,_bbe );if _ecf >=_gfda ._egae ||_ecf ==-1{_cbg =_gfda ._ecc +_ecbg ._ace ;
}else {_cbg =_gfda ._gf [_ecf ]+_ecbg ._ace ;};_gfda ._bfg [_gfda ._cbb ]=_cbg ;_gfda ._cbb ++;_bbe =!_bbe ;};continue _bdcf ;};};return nil ;};func _adf (_gecb []byte ,_cfd int )int {if _cfd >=len (_gecb ){return _cfd ;};if _cfd < -1{_cfd =-1;};var _efe byte ;
if _cfd > -1{_efe =_gecb [_cfd ];}else {_efe =_deg ;};_fdfc :=_cfd +1;for _fdfc < len (_gecb ){if _gecb [_fdfc ]!=_efe {break ;};_fdfc ++;};return _fdfc ;};func (_bgc *Decoder )Read (in []byte )(int ,error ){if _bgc ._bdb !=nil {return 0,_bgc ._bdb ;};
_dcac :=len (in );var (_cad int ;_ecdd int ;);for _dcac !=0{if _bgc ._bec >=_bgc ._cd {if _bfd :=_bgc .fetch ();_bfd !=nil {_bgc ._bdb =_bfd ;return 0,_bfd ;};};if _bgc ._cd ==-1{return _cad ,_c .EOF ;};switch {case _dcac <=_bgc ._cd -_bgc ._bec :_dcbg :=_bgc ._ee [_bgc ._bec :_bgc ._bec +_dcac ];
for _ ,_aeg :=range _dcbg {if !_bgc ._be {_aeg =^_aeg ;};in [_ecdd ]=_aeg ;_ecdd ++;};_cad +=len (_dcbg );_bgc ._bec +=len (_dcbg );return _cad ,nil ;default:_dgd :=_bgc ._ee [_bgc ._bec :];for _ ,_fcc :=range _dgd {if !_bgc ._be {_fcc =^_fcc ;};in [_ecdd ]=_fcc ;
_ecdd ++;};_cad +=len (_dgd );_bgc ._bec +=len (_dgd );_dcac -=len (_dgd );};};return _cad ,nil ;};var _afc =[...][]uint16 {{0x2,0x3},{0x2,0x3},{0x2,0x3},{0x3},{0x4,0x5},{0x4,0x5,0x7},{0x4,0x7},{0x18},{0x17,0x18,0x37,0x8,0xf},{0x17,0x18,0x28,0x37,0x67,0x68,0x6c,0x8,0xc,0xd},{0x12,0x13,0x14,0x15,0x16,0x17,0x1c,0x1d,0x1e,0x1f,0x24,0x27,0x28,0x2b,0x2c,0x33,0x34,0x35,0x37,0x38,0x52,0x53,0x54,0x55,0x56,0x57,0x58,0x59,0x5a,0x5b,0x64,0x65,0x66,0x67,0x68,0x69,0x6a,0x6b,0x6c,0x6d,0xc8,0xc9,0xca,0xcb,0xcc,0xcd,0xd2,0xd3,0xd4,0xd5,0xd6,0xd7,0xda,0xdb},{0x4a,0x4b,0x4c,0x4d,0x52,0x53,0x54,0x55,0x5a,0x5b,0x64,0x65,0x6c,0x6d,0x72,0x73,0x74,0x75,0x76,0x77}};
func _cbbf (_ged int )([]byte ,int ){var _abc []byte ;for _bde :=0;_bde < 6;_bde ++{_abc ,_ged =_dddg (_abc ,_ged ,_ecd );};return _abc ,_ged %8;};func (_fgfg *Decoder )tryFetchRTC2D ()(_fee error ){_fgfg ._dca .Mark ();var _adba bool ;for _egbc :=0;_egbc < 5;
_egbc ++{_adba ,_fee =_fgfg .tryFetchEOL1 ();if _fee !=nil {if _f .Is (_fee ,_c .EOF ){if _egbc ==0{break ;};return _ea ;};};if _adba {continue ;};if _egbc > 0{return _ea ;};break ;};if _adba {return _c .EOF ;};_fgfg ._dca .Reset ();return _fee ;};type tiffType int ;
func _aec (_bega int )([]byte ,int ){var _edfg []byte ;for _cgb :=0;_cgb < 2;_cgb ++{_edfg ,_bega =_dddg (_edfg ,_bega ,_aaa );};return _edfg ,_bega %8;};func (_bgb *tree )fillWithNode (_ebg ,_eefg int ,_cbda *treeNode )error {_cdc :=_bgb ._fgeb ;for _dceb :=0;
_dceb < _ebg ;_dceb ++{_gegc :=uint (_ebg -1-_dceb );_ggfe :=((_eefg >>_gegc )&1)!=0;_ccd :=_cdc .walk (_ggfe );if _ccd !=nil {if _ccd ._gde {return _f .New ("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067");
};_cdc =_ccd ;continue ;};if _dceb ==_ebg -1{_ccd =_cbda ;}else {_ccd =&treeNode {};};if _eefg ==0{_ccd ._dga =true ;};_cdc .set (_ggfe ,_ccd );_cdc =_ccd ;};return nil ;};func (_fcg tiffType )String ()string {switch _fcg {case _bc :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065";
case _ffe :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034";case _bg :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036";default:return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d";};};func (_ab *Decoder )looseFetchEOL ()(bool ,error ){_dgb ,_edf :=_ab ._dca .ReadBits (12);
if _edf !=nil {return false ,_edf ;};switch _dgb {case 0x1:return true ,nil ;case 0x0:for {_eafde ,_aef :=_ab ._dca .ReadBool ();if _aef !=nil {return false ,_aef ;};if _eafde {return true ,nil ;};};default:return false ,nil ;};};func (_aaf *Encoder )Encode (pixels [][]byte )[]byte {if _aaf .BlackIs1 {_deg =0;
_feb =1;}else {_deg =1;_feb =0;};if _aaf .K ==0{return _aaf .encodeG31D (pixels );};if _aaf .K > 0{return _aaf .encodeG32D (pixels );};if _aaf .K < 0{return _aaf .encodeG4 (pixels );};return nil ;};func (_aad *Encoder )encodeG32D (_aegd [][]byte )[]byte {var _bgce []byte ;
var _geg int ;for _gecc :=0;_gecc < len (_aegd );_gecc +=_aad .K {if _aad .Rows > 0&&!_aad .EndOfBlock &&_gecc ==_aad .Rows {break ;};_gfe ,_dfb :=_bea (_aegd [_gecc ],_geg ,_ecd );_bgce =_aad .appendEncodedRow (_bgce ,_gfe ,_geg );if _aad .EncodedByteAlign {_dfb =0;
};_geg =_dfb ;for _ffg :=_gecc +1;_ffg < (_gecc +_aad .K )&&_ffg < len (_aegd );_ffg ++{if _aad .Rows > 0&&!_aad .EndOfBlock &&_ffg ==_aad .Rows {break ;};_deb ,_fbf :=_dddg (nil ,_geg ,_cce );var _geb ,_bgf ,_ede int ;_abde :=-1;for _abde < len (_aegd [_ffg ]){_geb =_adf (_aegd [_ffg ],_abde );
_bgf =_cfcg (_aegd [_ffg ],_aegd [_ffg -1],_abde );_ede =_adf (_aegd [_ffg -1],_bgf );if _ede < _geb {_deb ,_fbf =_acca (_deb ,_fbf );_abde =_ede ;}else {if _bb .Abs (float64 (_bgf -_geb ))> 3{_deb ,_fbf ,_abde =_dcbf (_aegd [_ffg ],_deb ,_fbf ,_abde ,_geb );
}else {_deb ,_fbf =_bda (_deb ,_fbf ,_geb ,_bgf );_abde =_geb ;};};};_bgce =_aad .appendEncodedRow (_bgce ,_deb ,_geg );if _aad .EncodedByteAlign {_fbf =0;};_geg =_fbf %8;};};if _aad .EndOfBlock {_fgee ,_ :=_cbbf (_geg );_bgce =_aad .appendEncodedRow (_bgce ,_fgee ,_geg );
};return _bgce ;};func (_dbf *Decoder )decode2D ()error {_dbf ._egae =_dbf ._cbb ;_dbf ._bfg ,_dbf ._gf =_dbf ._gf ,_dbf ._bfg ;_efb :=true ;var (_de bool ;_gcb int ;_ggf error ;);_dbf ._cbb =0;_cgg :for _gcb < _dbf ._ecc {_afcc :=_d ._fgeb ;for {_de ,_ggf =_dbf ._dca .ReadBool ();
if _ggf !=nil {return _ggf ;};_afcc =_afcc .walk (_de );if _afcc ==nil {continue _cgg ;};if !_afcc ._gde {continue ;};switch _afcc ._ace {case _aa :var _cee int ;if _efb {_cee ,_ggf =_dbf .decodeRun (_fg );}else {_cee ,_ggf =_dbf .decodeRun (_ff );};if _ggf !=nil {return _ggf ;
};_gcb +=_cee ;_dbf ._bfg [_dbf ._cbb ]=_gcb ;_dbf ._cbb ++;if _efb {_cee ,_ggf =_dbf .decodeRun (_ff );}else {_cee ,_ggf =_dbf .decodeRun (_fg );};if _ggf !=nil {return _ggf ;};_gcb +=_cee ;_dbf ._bfg [_dbf ._cbb ]=_gcb ;_dbf ._cbb ++;case _bf :_cbad :=_dbf .getNextChangingElement (_gcb ,_efb )+1;
if _cbad >=_dbf ._egae {_gcb =_dbf ._ecc ;}else {_gcb =_dbf ._gf [_cbad ];};default:_dcd :=_dbf .getNextChangingElement (_gcb ,_efb );if _dcd >=_dbf ._egae ||_dcd ==-1{_gcb =_dbf ._ecc +_afcc ._ace ;}else {_gcb =_dbf ._gf [_dcd ]+_afcc ._ace ;};_dbf ._bfg [_dbf ._cbb ]=_gcb ;
_dbf ._cbb ++;_efb =!_efb ;};continue _cgg ;};};return nil ;};func (_bfa *tree )fill (_abce ,_agfd ,_aded int )error {_befd :=_bfa ._fgeb ;for _dac :=0;_dac < _abce ;_dac ++{_acf :=_abce -1-_dac ;_eee :=((_agfd >>uint (_acf ))&1)!=0;_dgf :=_befd .walk (_eee );
if _dgf !=nil {if _dgf ._gde {return _f .New ("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067");};_befd =_dgf ;continue ;
};_dgf =&treeNode {};if _dac ==_abce -1{_dgf ._ace =_aded ;_dgf ._gde =true ;};if _agfd ==0{_dgf ._dga =true ;};_befd .set (_eee ,_dgf );_befd =_dgf ;};return nil ;};func (_ebef *Decoder )tryFetchEOL ()(bool ,error ){_eafd ,_fcf :=_ebef ._dca .ReadBits (12);
if _fcf !=nil {return false ,_fcf ;};return _eafd ==0x1,nil ;};var _fdf =[...][]uint16 {{2,3,4,5,6,7},{128,8,9,64,10,11},{192,1664,16,17,13,14,15,1,12},{26,21,28,27,18,24,25,22,256,23,20,19},{33,34,35,36,37,38,31,32,29,53,54,39,40,41,42,43,44,30,61,62,63,0,320,384,45,59,60,46,49,50,51,52,55,56,57,58,448,512,640,576,47,48},{1472,1536,1600,1728,704,768,832,896,960,1024,1088,1152,1216,1280,1344,1408},{},{1792,1856,1920},{1984,2048,2112,2176,2240,2304,2368,2432,2496,2560}};
func _cfcg (_dbe ,_cdb []byte ,_ecccg int )int {_cab :=_adf (_cdb ,_ecccg );if _cab < len (_cdb )&&(_ecccg ==-1&&_cdb [_cab ]==_deg ||_ecccg >=0&&_ecccg < len (_dbe )&&_dbe [_ecccg ]==_cdb [_cab ]||_ecccg >=len (_dbe )&&_dbe [_ecccg -1]!=_cdb [_cab ]){_cab =_adf (_cdb ,_cab );
};return _cab ;};func (_aced *treeNode )set (_fbb bool ,_baca *treeNode ){if !_fbb {_aced ._ecfe =_baca ;}else {_aced ._fdc =_baca ;};};func _abg (_dbc []byte ,_cefe bool ,_ddc int )(int ,int ){_dafg :=0;for _ddc < len (_dbc ){if _cefe {if _dbc [_ddc ]!=_deg {break ;
};}else {if _dbc [_ddc ]!=_feb {break ;};};_dafg ++;_ddc ++;};return _dafg ,_ddc ;};func _acca (_dea []byte ,_feaf int )([]byte ,int ){return _dddg (_dea ,_feaf ,_dc )};func (_cba *Decoder )fetch ()error {if _cba ._cd ==-1{return nil ;};if _cba ._bec < _cba ._cd {return nil ;
};_cba ._cd =0;_gfd :=_cba .decodeRow ();if _gfd !=nil {if !_f .Is (_gfd ,_c .EOF ){return _gfd ;};if _cba ._cd !=0{return _gfd ;};_cba ._cd =-1;};_cba ._bec =0;return nil ;};type DecodeOptions struct{Columns int ;Rows int ;K int ;EncodedByteAligned bool ;
BlackIsOne bool ;EndOfBlock bool ;EndOfLine bool ;DamagedRowsBeforeError int ;};func (_bag *Decoder )decodeRun (_bcf *tree )(int ,error ){var _cggc int ;_eccc :=_bcf ._fgeb ;for {_fgf ,_bege :=_bag ._dca .ReadBool ();if _bege !=nil {return 0,_bege ;};_eccc =_eccc .walk (_fgf );
if _eccc ==nil {return 0,_f .New ("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d");};if _eccc ._gde {_cggc +=_eccc ._ace ;switch {case _eccc ._ace >=64:_eccc =_bcf ._fgeb ;
case _eccc ._ace >=0:return _cggc ,nil ;default:return _bag ._ecc ,nil ;};};};};type tree struct{_fgeb *treeNode };func _fgb (_eff ,_egg []byte ,_becb int ,_fccf bool )int {_dcbge :=_adf (_egg ,_becb );if _dcbge < len (_egg )&&(_becb ==-1&&_egg [_dcbge ]==_deg ||_becb >=0&&_becb < len (_eff )&&_eff [_becb ]==_egg [_dcbge ]||_becb >=len (_eff )&&_fccf &&_egg [_dcbge ]==_deg ||_becb >=len (_eff )&&!_fccf &&_egg [_dcbge ]==_feb ){_dcbge =_adf (_egg ,_dcbge );
};return _dcbge ;};func _bea (_age []byte ,_ga int ,_fdd code )([]byte ,int ){_ebb :=true ;var _fef []byte ;_fef ,_ga =_dddg (nil ,_ga ,_fdd );_ecba :=0;var _gcd int ;for _ecba < len (_age ){_gcd ,_ecba =_abg (_age ,_ebb ,_ecba );_fef ,_ga =_dbg (_fef ,_ga ,_gcd ,_ebb );
_ebb =!_ebb ;};return _fef ,_ga %8;};func (_eae *Decoder )decodeRowType2 ()error {if _eae ._fcd {_eae ._dca .Align ();};if _eef :=_eae .decode1D ();_eef !=nil {return _eef ;};return nil ;};type treeNode struct{_ecfe *treeNode ;_fdc *treeNode ;_ace int ;
_dga bool ;_gde bool ;};func (_bae *Decoder )decode1D ()error {var (_fb int ;_bbfb error ;);_ddb :=true ;_bae ._cbb =0;for {var _gcg int ;if _ddb {_gcg ,_bbfb =_bae .decodeRun (_fg );}else {_gcg ,_bbfb =_bae .decodeRun (_ff );};if _bbfb !=nil {return _bbfb ;
};_fb +=_gcg ;_bae ._bfg [_bae ._cbb ]=_fb ;_bae ._cbb ++;_ddb =!_ddb ;if _fb >=_bae ._ecc {break ;};};return nil ;};var (_ea =_f .New ("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043");
_ce =_f .New ("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064"););func _beb (_cgbc int ,_beag bool )(code ,int ,bool ){if _cgbc < 64{if _beag {return _cf [_cgbc ],0,true ;};return _aca [_cgbc ],0,true ;
};_aab :=_cgbc /64;if _aab > 40{return _fad [2560],_cgbc -2560,false ;};if _aab > 27{return _fad [_aab *64],_cgbc -_aab *64,false ;};if _beag {return _ec [_aab *64],_cgbc -_aab *64,false ;};return _ef [_aab *64],_cgbc -_aab *64,false ;};func (_ddg *Encoder )encodeG4 (_bgg [][]byte )[]byte {_feba :=make ([][]byte ,len (_bgg ));
copy (_feba ,_bgg );_feba =_ecga (_feba );var _fgc []byte ;var _gba int ;for _dce :=1;_dce < len (_feba );_dce ++{if _ddg .Rows > 0&&!_ddg .EndOfBlock &&_dce ==(_ddg .Rows +1){break ;};var _cec []byte ;var _efd ,_cae ,_edc int ;_aac :=_gba ;_gdf :=-1;for _gdf < len (_feba [_dce ]){_efd =_adf (_feba [_dce ],_gdf );
_cae =_cfcg (_feba [_dce ],_feba [_dce -1],_gdf );_edc =_adf (_feba [_dce -1],_cae );if _edc < _efd {_cec ,_aac =_dddg (_cec ,_aac ,_dc );_gdf =_edc ;}else {if _bb .Abs (float64 (_cae -_efd ))> 3{_cec ,_aac ,_gdf =_dcbf (_feba [_dce ],_cec ,_aac ,_gdf ,_efd );
}else {_cec ,_aac =_bda (_cec ,_aac ,_efd ,_cae );_gdf =_efd ;};};};_fgc =_ddg .appendEncodedRow (_fgc ,_cec ,_gba );if _ddg .EncodedByteAlign {_aac =0;};_gba =_aac %8;};if _ddg .EndOfBlock {_fde ,_ :=_aec (_gba );_fgc =_ddg .appendEncodedRow (_fgc ,_fde ,_gba );
};return _fgc ;};func _ecbc (_fac int )([]byte ,int ){var _abe []byte ;for _eag :=0;_eag < 6;_eag ++{_abe ,_fac =_dddg (_abe ,_fac ,_aaa );};return _abe ,_fac %8;};type Encoder struct{K int ;EndOfLine bool ;EncodedByteAlign bool ;Columns int ;Rows int ;
EndOfBlock bool ;BlackIs1 bool ;DamagedRowsBeforeError int ;};func (_eeg *Decoder )decodeRowType4 ()error {if !_eeg ._ccf {return _eeg .decoderRowType41D ();};if _eeg ._fcd {_eeg ._dca .Align ();};_eeg ._dca .Mark ();_eb ,_bcd :=_eeg .tryFetchEOL ();if _bcd !=nil {return _bcd ;
};if !_eb &&_eeg ._cca {_eeg ._fge ++;if _eeg ._fge > _eeg ._bge {return _ce ;};_eeg ._dca .Reset ();};if !_eb {_eeg ._dca .Reset ();};_ebe ,_bcd :=_eeg ._dca .ReadBool ();if _bcd !=nil {return _bcd ;};if _ebe {if _eb &&_eeg ._aee {if _bcd =_eeg .tryFetchRTC2D ();
_bcd !=nil {return _bcd ;};};_bcd =_eeg .decode1D ();}else {_bcd =_eeg .decode2D ();};if _bcd !=nil {return _bcd ;};return nil ;};func (_cefg *Decoder )decodeRowType6 ()error {if _cefg ._fcd {_cefg ._dca .Align ();};if _cefg ._aee {_cefg ._dca .Mark ();
_ag ,_cde :=_cefg .tryFetchEOL ();if _cde !=nil {return _cde ;};if _ag {_ag ,_cde =_cefg .tryFetchEOL ();if _cde !=nil {return _cde ;};if _ag {return _c .EOF ;};};_cefg ._dca .Reset ();};return _cefg .decode2D ();};func _ecga (_aagb [][]byte )[][]byte {_ebbf :=make ([]byte ,len (_aagb [0]));
for _abed :=range _ebbf {_ebbf [_abed ]=_deg ;};_aagb =append (_aagb ,[]byte {});for _efa :=len (_aagb )-1;_efa > 0;_efa --{_aagb [_efa ]=_aagb [_efa -1];};_aagb [0]=_ebbf ;return _aagb ;};func (_bce *Encoder )appendEncodedRow (_dddb ,_ddgd []byte ,_cac int )[]byte {if len (_dddb )> 0&&_cac !=0&&!_bce .EncodedByteAlign {_dddb [len (_dddb )-1]=_dddb [len (_dddb )-1]|_ddgd [0];
_dddb =append (_dddb ,_ddgd [1:]...);}else {_dddb =append (_dddb ,_ddgd ...);};return _dddb ;};func (_bef *Encoder )encodeG31D (_cge [][]byte )[]byte {var _gec []byte ;_def :=0;for _cceg :=range _cge {if _bef .Rows > 0&&!_bef .EndOfBlock &&_cceg ==_bef .Rows {break ;
};_df ,_aae :=_bea (_cge [_cceg ],_def ,_aaa );_gec =_bef .appendEncodedRow (_gec ,_df ,_def );if _bef .EncodedByteAlign {_aae =0;};_def =_aae ;};if _bef .EndOfBlock {_dgbb ,_ :=_ecbc (_def );_gec =_bef .appendEncodedRow (_gec ,_dgbb ,_def );};return _gec ;
};func _cfa (_agf ,_gdad int )code {var _fcb code ;switch _gdad -_agf {case -1:_fcb =_fc ;case -2:_fcb =_dcb ;case -3:_fcb =_adg ;case 0:_fcb =_bac ;case 1:_fcb =_ecb ;case 2:_fcb =_ega ;case 3:_fcb =_adb ;};return _fcb ;};var (_e *treeNode ;_g *treeNode ;
_ff *tree ;_fg *tree ;_eg *tree ;_d *tree ;_fe =-2000;_gb =-1000;_bf =-3000;_aa =-4000;);var (_deg byte =1;_feb byte =0;);func (_gg *Decoder )decoderRowType41D ()error {if _gg ._fcd {_gg ._dca .Align ();};_gg ._dca .Mark ();var (_ddd bool ;_fda error ;
);if _gg ._cca {_ddd ,_fda =_gg .tryFetchEOL ();if _fda !=nil {return _fda ;};if !_ddd {return _ce ;};}else {_ddd ,_fda =_gg .looseFetchEOL ();if _fda !=nil {return _fda ;};};if !_ddd {_gg ._dca .Reset ();};if _ddd &&_gg ._aee {_gg ._dca .Mark ();for _ccg :=0;
_ccg < 5;_ccg ++{_ddd ,_fda =_gg .tryFetchEOL ();if _fda !=nil {if _f .Is (_fda ,_c .EOF ){if _ccg ==0{break ;};return _ea ;};};if _ddd {continue ;};if _ccg > 0{return _ea ;};break ;};if _ddd {return _c .EOF ;};_gg ._dca .Reset ();};if _fda =_gg .decode1D ();
_fda !=nil {return _fda ;};return nil ;};func (_ed *Decoder )getNextChangingElement (_cbbb int ,_aag bool )int {_bcb :=0;if !_aag {_bcb =1;};_dda :=int (uint32 (_ed ._cfb )&0xFFFFFFFE)+_bcb ;if _dda > 2{_dda -=2;};if _cbbb ==0{return _dda ;};for _dcba :=_dda ;
_dcba < _ed ._egae ;_dcba +=2{if _cbbb < _ed ._gf [_dcba ]{_ed ._cfb =_dcba ;return _dcba ;};};return -1;};