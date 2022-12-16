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

package ccittfax ;import (_f "errors";_b "github.com/unidoc/unipdf/v3/internal/bitwise";_ea "io";_a "math";);func _bcee (_dadd ,_eeg []byte ,_dge int )int {_abca :=_fbfd (_eeg ,_dge );if _abca < len (_eeg )&&(_dge ==-1&&_eeg [_abca ]==_fbf ||_dge >=0&&_dge < len (_dadd )&&_dadd [_dge ]==_eeg [_abca ]||_dge >=len (_dadd )&&_dadd [_dge -1]!=_eeg [_abca ]){_abca =_fbfd (_eeg ,_abca );
};return _abca ;};func init (){_ad =&treeNode {_fca :true ,_fagcb :_ge };_g =&treeNode {_fagcb :_gf ,_gag :_ad };_g ._eag =_g ;_fa =&tree {_cab :&treeNode {}};if _gd :=_fa .fillWithNode (12,0,_g );_gd !=nil {panic (_gd .Error ());};if _bef :=_fa .fillWithNode (12,1,_ad );
_bef !=nil {panic (_bef .Error ());};_ef =&tree {_cab :&treeNode {}};for _fce :=0;_fce < len (_dc );_fce ++{for _eg :=0;_eg < len (_dc [_fce ]);_eg ++{if _ag :=_ef .fill (_fce +2,int (_dc [_fce ][_eg ]),int (_ac [_fce ][_eg ]));_ag !=nil {panic (_ag .Error ());
};};};if _bd :=_ef .fillWithNode (12,0,_g );_bd !=nil {panic (_bd .Error ());};if _fag :=_ef .fillWithNode (12,1,_ad );_fag !=nil {panic (_fag .Error ());};_fc =&tree {_cab :&treeNode {}};for _eb :=0;_eb < len (_gec );_eb ++{for _cb :=0;_cb < len (_gec [_eb ]);
_cb ++{if _fg :=_fc .fill (_eb +4,int (_gec [_eb ][_cb ]),int (_cgg [_eb ][_cb ]));_fg !=nil {panic (_fg .Error ());};};};if _gea :=_fc .fillWithNode (12,0,_g );_gea !=nil {panic (_gea .Error ());};if _cf :=_fc .fillWithNode (12,1,_ad );_cf !=nil {panic (_cf .Error ());
};_c =&tree {_cab :&treeNode {}};if _aa :=_c .fill (4,1,_ab );_aa !=nil {panic (_aa .Error ());};if _d :=_c .fill (3,1,_be );_d !=nil {panic (_d .Error ());};if _ed :=_c .fill (1,1,0);_ed !=nil {panic (_ed .Error ());};if _eaa :=_c .fill (3,3,1);_eaa !=nil {panic (_eaa .Error ());
};if _cbc :=_c .fill (6,3,2);_cbc !=nil {panic (_cbc .Error ());};if _ae :=_c .fill (7,3,3);_ae !=nil {panic (_ae .Error ());};if _aab :=_c .fill (3,2,-1);_aab !=nil {panic (_aab .Error ());};if _cg :=_c .fill (6,2,-2);_cg !=nil {panic (_cg .Error ());
};if _ca :=_c .fill (7,2,-3);_ca !=nil {panic (_ca .Error ());};};func (_eaf *Decoder )Read (in []byte )(int ,error ){if _eaf ._bda !=nil {return 0,_eaf ._bda ;};_efa :=len (in );var (_eba int ;_dfc int ;);for _efa !=0{if _eaf ._beb >=_eaf ._cdba {if _abc :=_eaf .fetch ();
_abc !=nil {_eaf ._bda =_abc ;return 0,_abc ;};};if _eaf ._cdba ==-1{return _eba ,_ea .EOF ;};switch {case _efa <=_eaf ._cdba -_eaf ._beb :_bdb :=_eaf ._gdg [_eaf ._beb :_eaf ._beb +_efa ];for _ ,_dfb :=range _bdb {if !_eaf ._fbd {_dfb =^_dfb ;};in [_dfc ]=_dfb ;
_dfc ++;};_eba +=len (_bdb );_eaf ._beb +=len (_bdb );return _eba ,nil ;default:_abe :=_eaf ._gdg [_eaf ._beb :];for _ ,_ebac :=range _abe {if !_eaf ._fbd {_ebac =^_ebac ;};in [_dfc ]=_ebac ;_dfc ++;};_eba +=len (_abe );_eaf ._beb +=len (_abe );_efa -=len (_abe );
};};return _eba ,nil ;};type tree struct{_cab *treeNode };func (_gee *Decoder )decodeG32D ()error {_gee ._aaf =_gee ._df ;_gee ._gecfa ,_gee ._cge =_gee ._cge ,_gee ._gecfa ;_abf :=true ;var (_bdf bool ;_agc int ;_bg error ;);_gee ._df =0;_bcf :for _agc < _gee ._ec {_ebg :=_c ._cab ;
for {_bdf ,_bg =_gee ._fcg .ReadBool ();if _bg !=nil {return _bg ;};_ebg =_ebg .walk (_bdf );if _ebg ==nil {continue _bcf ;};if !_ebg ._fca {continue ;};switch _ebg ._fagcb {case _be :var _cbb int ;if _abf {_cbb ,_bg =_gee .decodeRun (_fc );}else {_cbb ,_bg =_gee .decodeRun (_ef );
};if _bg !=nil {return _bg ;};_agc +=_cbb ;_gee ._gecfa [_gee ._df ]=_agc ;_gee ._df ++;if _abf {_cbb ,_bg =_gee .decodeRun (_ef );}else {_cbb ,_bg =_gee .decodeRun (_fc );};if _bg !=nil {return _bg ;};_agc +=_cbb ;_gee ._gecfa [_gee ._df ]=_agc ;_gee ._df ++;
case _ab :_daa :=_gee .getNextChangingElement (_agc ,_abf )+1;if _daa >=_gee ._aaf {_agc =_gee ._ec ;}else {_agc =_gee ._cge [_daa ];};default:_ce :=_gee .getNextChangingElement (_agc ,_abf );if _ce >=_gee ._aaf ||_ce ==-1{_agc =_gee ._ec +_ebg ._fagcb ;
}else {_agc =_gee ._cge [_ce ]+_ebg ._fagcb ;};_gee ._gecfa [_gee ._df ]=_agc ;_gee ._df ++;_abf =!_abf ;};continue _bcf ;};};return nil ;};func (_gced *Encoder )encodeG4 (_faf [][]byte )[]byte {_cde :=make ([][]byte ,len (_faf ));copy (_cde ,_faf );_cde =_ggb (_cde );
var _cga []byte ;var _baed int ;for _egd :=1;_egd < len (_cde );_egd ++{if _gced .Rows > 0&&!_gced .EndOfBlock &&_egd ==(_gced .Rows +1){break ;};var _fgdc []byte ;var _cfcd ,_cdg ,_ga int ;_aeb :=_baed ;_bag :=-1;for _bag < len (_cde [_egd ]){_cfcd =_fbfd (_cde [_egd ],_bag );
_cdg =_bcee (_cde [_egd ],_cde [_egd -1],_bag );_ga =_fbfd (_cde [_egd -1],_cdg );if _ga < _cfcd {_fgdc ,_aeb =_dcfa (_fgdc ,_aeb ,_aga );_bag =_ga ;}else {if _a .Abs (float64 (_cdg -_cfcd ))> 3{_fgdc ,_aeb ,_bag =_bgde (_cde [_egd ],_fgdc ,_aeb ,_bag ,_cfcd );
}else {_fgdc ,_aeb =_cdf (_fgdc ,_aeb ,_cfcd ,_cdg );_bag =_cfcd ;};};};_cga =_gced .appendEncodedRow (_cga ,_fgdc ,_baed );if _gced .EncodedByteAlign {_aeb =0;};_baed =_aeb %8;};if _gced .EndOfBlock {_adag ,_ :=_dgb (_baed );_cga =_gced .appendEncodedRow (_cga ,_adag ,_baed );
};return _cga ;};func _efb (_bafd []byte ,_dbe int )([]byte ,int ){return _dcfa (_bafd ,_dbe ,_aga )};func (_ade *Encoder )encodeG31D (_ecg [][]byte )[]byte {var _bgbb []byte ;_gdaf :=0;for _egc :=range _ecg {if _ade .Rows > 0&&!_ade .EndOfBlock &&_egc ==_ade .Rows {break ;
};_fagd ,_gecc :=_dgg (_ecg [_egc ],_gdaf ,_cad );_bgbb =_ade .appendEncodedRow (_bgbb ,_fagd ,_gdaf );if _ade .EncodedByteAlign {_gecc =0;};_gdaf =_gecc ;};if _ade .EndOfBlock {_ecgb ,_ :=_bfc (_gdaf );_bgbb =_ade .appendEncodedRow (_bgbb ,_ecgb ,_gdaf );
};return _bgbb ;};func init (){_af =make (map[int ]code );_af [0]=code {Code :13<<8|3<<6,BitsWritten :10};_af [1]=code {Code :2<<(5+8),BitsWritten :3};_af [2]=code {Code :3<<(6+8),BitsWritten :2};_af [3]=code {Code :2<<(6+8),BitsWritten :2};_af [4]=code {Code :3<<(5+8),BitsWritten :3};
_af [5]=code {Code :3<<(4+8),BitsWritten :4};_af [6]=code {Code :2<<(4+8),BitsWritten :4};_af [7]=code {Code :3<<(3+8),BitsWritten :5};_af [8]=code {Code :5<<(2+8),BitsWritten :6};_af [9]=code {Code :4<<(2+8),BitsWritten :6};_af [10]=code {Code :4<<(1+8),BitsWritten :7};
_af [11]=code {Code :5<<(1+8),BitsWritten :7};_af [12]=code {Code :7<<(1+8),BitsWritten :7};_af [13]=code {Code :4<<8,BitsWritten :8};_af [14]=code {Code :7<<8,BitsWritten :8};_af [15]=code {Code :12<<8,BitsWritten :9};_af [16]=code {Code :5<<8|3<<6,BitsWritten :10};
_af [17]=code {Code :6<<8,BitsWritten :10};_af [18]=code {Code :2<<8,BitsWritten :10};_af [19]=code {Code :12<<8|7<<5,BitsWritten :11};_af [20]=code {Code :13<<8,BitsWritten :11};_af [21]=code {Code :13<<8|4<<5,BitsWritten :11};_af [22]=code {Code :6<<8|7<<5,BitsWritten :11};
_af [23]=code {Code :5<<8,BitsWritten :11};_af [24]=code {Code :2<<8|7<<5,BitsWritten :11};_af [25]=code {Code :3<<8,BitsWritten :11};_af [26]=code {Code :12<<8|10<<4,BitsWritten :12};_af [27]=code {Code :12<<8|11<<4,BitsWritten :12};_af [28]=code {Code :12<<8|12<<4,BitsWritten :12};
_af [29]=code {Code :12<<8|13<<4,BitsWritten :12};_af [30]=code {Code :6<<8|8<<4,BitsWritten :12};_af [31]=code {Code :6<<8|9<<4,BitsWritten :12};_af [32]=code {Code :6<<8|10<<4,BitsWritten :12};_af [33]=code {Code :6<<8|11<<4,BitsWritten :12};_af [34]=code {Code :13<<8|2<<4,BitsWritten :12};
_af [35]=code {Code :13<<8|3<<4,BitsWritten :12};_af [36]=code {Code :13<<8|4<<4,BitsWritten :12};_af [37]=code {Code :13<<8|5<<4,BitsWritten :12};_af [38]=code {Code :13<<8|6<<4,BitsWritten :12};_af [39]=code {Code :13<<8|7<<4,BitsWritten :12};_af [40]=code {Code :6<<8|12<<4,BitsWritten :12};
_af [41]=code {Code :6<<8|13<<4,BitsWritten :12};_af [42]=code {Code :13<<8|10<<4,BitsWritten :12};_af [43]=code {Code :13<<8|11<<4,BitsWritten :12};_af [44]=code {Code :5<<8|4<<4,BitsWritten :12};_af [45]=code {Code :5<<8|5<<4,BitsWritten :12};_af [46]=code {Code :5<<8|6<<4,BitsWritten :12};
_af [47]=code {Code :5<<8|7<<4,BitsWritten :12};_af [48]=code {Code :6<<8|4<<4,BitsWritten :12};_af [49]=code {Code :6<<8|5<<4,BitsWritten :12};_af [50]=code {Code :5<<8|2<<4,BitsWritten :12};_af [51]=code {Code :5<<8|3<<4,BitsWritten :12};_af [52]=code {Code :2<<8|4<<4,BitsWritten :12};
_af [53]=code {Code :3<<8|7<<4,BitsWritten :12};_af [54]=code {Code :3<<8|8<<4,BitsWritten :12};_af [55]=code {Code :2<<8|7<<4,BitsWritten :12};_af [56]=code {Code :2<<8|8<<4,BitsWritten :12};_af [57]=code {Code :5<<8|8<<4,BitsWritten :12};_af [58]=code {Code :5<<8|9<<4,BitsWritten :12};
_af [59]=code {Code :2<<8|11<<4,BitsWritten :12};_af [60]=code {Code :2<<8|12<<4,BitsWritten :12};_af [61]=code {Code :5<<8|10<<4,BitsWritten :12};_af [62]=code {Code :6<<8|6<<4,BitsWritten :12};_af [63]=code {Code :6<<8|7<<4,BitsWritten :12};_cd =make (map[int ]code );
_cd [0]=code {Code :53<<8,BitsWritten :8};_cd [1]=code {Code :7<<(2+8),BitsWritten :6};_cd [2]=code {Code :7<<(4+8),BitsWritten :4};_cd [3]=code {Code :8<<(4+8),BitsWritten :4};_cd [4]=code {Code :11<<(4+8),BitsWritten :4};_cd [5]=code {Code :12<<(4+8),BitsWritten :4};
_cd [6]=code {Code :14<<(4+8),BitsWritten :4};_cd [7]=code {Code :15<<(4+8),BitsWritten :4};_cd [8]=code {Code :19<<(3+8),BitsWritten :5};_cd [9]=code {Code :20<<(3+8),BitsWritten :5};_cd [10]=code {Code :7<<(3+8),BitsWritten :5};_cd [11]=code {Code :8<<(3+8),BitsWritten :5};
_cd [12]=code {Code :8<<(2+8),BitsWritten :6};_cd [13]=code {Code :3<<(2+8),BitsWritten :6};_cd [14]=code {Code :52<<(2+8),BitsWritten :6};_cd [15]=code {Code :53<<(2+8),BitsWritten :6};_cd [16]=code {Code :42<<(2+8),BitsWritten :6};_cd [17]=code {Code :43<<(2+8),BitsWritten :6};
_cd [18]=code {Code :39<<(1+8),BitsWritten :7};_cd [19]=code {Code :12<<(1+8),BitsWritten :7};_cd [20]=code {Code :8<<(1+8),BitsWritten :7};_cd [21]=code {Code :23<<(1+8),BitsWritten :7};_cd [22]=code {Code :3<<(1+8),BitsWritten :7};_cd [23]=code {Code :4<<(1+8),BitsWritten :7};
_cd [24]=code {Code :40<<(1+8),BitsWritten :7};_cd [25]=code {Code :43<<(1+8),BitsWritten :7};_cd [26]=code {Code :19<<(1+8),BitsWritten :7};_cd [27]=code {Code :36<<(1+8),BitsWritten :7};_cd [28]=code {Code :24<<(1+8),BitsWritten :7};_cd [29]=code {Code :2<<8,BitsWritten :8};
_cd [30]=code {Code :3<<8,BitsWritten :8};_cd [31]=code {Code :26<<8,BitsWritten :8};_cd [32]=code {Code :27<<8,BitsWritten :8};_cd [33]=code {Code :18<<8,BitsWritten :8};_cd [34]=code {Code :19<<8,BitsWritten :8};_cd [35]=code {Code :20<<8,BitsWritten :8};
_cd [36]=code {Code :21<<8,BitsWritten :8};_cd [37]=code {Code :22<<8,BitsWritten :8};_cd [38]=code {Code :23<<8,BitsWritten :8};_cd [39]=code {Code :40<<8,BitsWritten :8};_cd [40]=code {Code :41<<8,BitsWritten :8};_cd [41]=code {Code :42<<8,BitsWritten :8};
_cd [42]=code {Code :43<<8,BitsWritten :8};_cd [43]=code {Code :44<<8,BitsWritten :8};_cd [44]=code {Code :45<<8,BitsWritten :8};_cd [45]=code {Code :4<<8,BitsWritten :8};_cd [46]=code {Code :5<<8,BitsWritten :8};_cd [47]=code {Code :10<<8,BitsWritten :8};
_cd [48]=code {Code :11<<8,BitsWritten :8};_cd [49]=code {Code :82<<8,BitsWritten :8};_cd [50]=code {Code :83<<8,BitsWritten :8};_cd [51]=code {Code :84<<8,BitsWritten :8};_cd [52]=code {Code :85<<8,BitsWritten :8};_cd [53]=code {Code :36<<8,BitsWritten :8};
_cd [54]=code {Code :37<<8,BitsWritten :8};_cd [55]=code {Code :88<<8,BitsWritten :8};_cd [56]=code {Code :89<<8,BitsWritten :8};_cd [57]=code {Code :90<<8,BitsWritten :8};_cd [58]=code {Code :91<<8,BitsWritten :8};_cd [59]=code {Code :74<<8,BitsWritten :8};
_cd [60]=code {Code :75<<8,BitsWritten :8};_cd [61]=code {Code :50<<8,BitsWritten :8};_cd [62]=code {Code :51<<8,BitsWritten :8};_cd [63]=code {Code :52<<8,BitsWritten :8};_fb =make (map[int ]code );_fb [64]=code {Code :3<<8|3<<6,BitsWritten :10};_fb [128]=code {Code :12<<8|8<<4,BitsWritten :12};
_fb [192]=code {Code :12<<8|9<<4,BitsWritten :12};_fb [256]=code {Code :5<<8|11<<4,BitsWritten :12};_fb [320]=code {Code :3<<8|3<<4,BitsWritten :12};_fb [384]=code {Code :3<<8|4<<4,BitsWritten :12};_fb [448]=code {Code :3<<8|5<<4,BitsWritten :12};_fb [512]=code {Code :3<<8|12<<3,BitsWritten :13};
_fb [576]=code {Code :3<<8|13<<3,BitsWritten :13};_fb [640]=code {Code :2<<8|10<<3,BitsWritten :13};_fb [704]=code {Code :2<<8|11<<3,BitsWritten :13};_fb [768]=code {Code :2<<8|12<<3,BitsWritten :13};_fb [832]=code {Code :2<<8|13<<3,BitsWritten :13};_fb [896]=code {Code :3<<8|18<<3,BitsWritten :13};
_fb [960]=code {Code :3<<8|19<<3,BitsWritten :13};_fb [1024]=code {Code :3<<8|20<<3,BitsWritten :13};_fb [1088]=code {Code :3<<8|21<<3,BitsWritten :13};_fb [1152]=code {Code :3<<8|22<<3,BitsWritten :13};_fb [1216]=code {Code :119<<3,BitsWritten :13};_fb [1280]=code {Code :2<<8|18<<3,BitsWritten :13};
_fb [1344]=code {Code :2<<8|19<<3,BitsWritten :13};_fb [1408]=code {Code :2<<8|20<<3,BitsWritten :13};_fb [1472]=code {Code :2<<8|21<<3,BitsWritten :13};_fb [1536]=code {Code :2<<8|26<<3,BitsWritten :13};_fb [1600]=code {Code :2<<8|27<<3,BitsWritten :13};
_fb [1664]=code {Code :3<<8|4<<3,BitsWritten :13};_fb [1728]=code {Code :3<<8|5<<3,BitsWritten :13};_dd =make (map[int ]code );_dd [64]=code {Code :27<<(3+8),BitsWritten :5};_dd [128]=code {Code :18<<(3+8),BitsWritten :5};_dd [192]=code {Code :23<<(2+8),BitsWritten :6};
_dd [256]=code {Code :55<<(1+8),BitsWritten :7};_dd [320]=code {Code :54<<8,BitsWritten :8};_dd [384]=code {Code :55<<8,BitsWritten :8};_dd [448]=code {Code :100<<8,BitsWritten :8};_dd [512]=code {Code :101<<8,BitsWritten :8};_dd [576]=code {Code :104<<8,BitsWritten :8};
_dd [640]=code {Code :103<<8,BitsWritten :8};_dd [704]=code {Code :102<<8,BitsWritten :9};_dd [768]=code {Code :102<<8|1<<7,BitsWritten :9};_dd [832]=code {Code :105<<8,BitsWritten :9};_dd [896]=code {Code :105<<8|1<<7,BitsWritten :9};_dd [960]=code {Code :106<<8,BitsWritten :9};
_dd [1024]=code {Code :106<<8|1<<7,BitsWritten :9};_dd [1088]=code {Code :107<<8,BitsWritten :9};_dd [1152]=code {Code :107<<8|1<<7,BitsWritten :9};_dd [1216]=code {Code :108<<8,BitsWritten :9};_dd [1280]=code {Code :108<<8|1<<7,BitsWritten :9};_dd [1344]=code {Code :109<<8,BitsWritten :9};
_dd [1408]=code {Code :109<<8|1<<7,BitsWritten :9};_dd [1472]=code {Code :76<<8,BitsWritten :9};_dd [1536]=code {Code :76<<8|1<<7,BitsWritten :9};_dd [1600]=code {Code :77<<8,BitsWritten :9};_dd [1664]=code {Code :24<<(2+8),BitsWritten :6};_dd [1728]=code {Code :77<<8|1<<7,BitsWritten :9};
_fgf =make (map[int ]code );_fgf [1792]=code {Code :1<<8,BitsWritten :11};_fgf [1856]=code {Code :1<<8|4<<5,BitsWritten :11};_fgf [1920]=code {Code :1<<8|5<<5,BitsWritten :11};_fgf [1984]=code {Code :1<<8|2<<4,BitsWritten :12};_fgf [2048]=code {Code :1<<8|3<<4,BitsWritten :12};
_fgf [2112]=code {Code :1<<8|4<<4,BitsWritten :12};_fgf [2176]=code {Code :1<<8|5<<4,BitsWritten :12};_fgf [2240]=code {Code :1<<8|6<<4,BitsWritten :12};_fgf [2304]=code {Code :1<<8|7<<4,BitsWritten :12};_fgf [2368]=code {Code :1<<8|12<<4,BitsWritten :12};
_fgf [2432]=code {Code :1<<8|13<<4,BitsWritten :12};_fgf [2496]=code {Code :1<<8|14<<4,BitsWritten :12};_fgf [2560]=code {Code :1<<8|15<<4,BitsWritten :12};_gc =make (map[int ]byte );_gc [0]=0xFF;_gc [1]=0xFE;_gc [2]=0xFC;_gc [3]=0xF8;_gc [4]=0xF0;_gc [5]=0xE0;
_gc [6]=0xC0;_gc [7]=0x80;_gc [8]=0x00;};func _cbf (_cdcg int ,_gdge bool )(code ,int ,bool ){if _cdcg < 64{if _gdge {return _cd [_cdcg ],0,true ;};return _af [_cdcg ],0,true ;};_bce :=_cdcg /64;if _bce > 40{return _fgf [2560],_cdcg -2560,false ;};if _bce > 27{return _fgf [_bce *64],_cdcg -_bce *64,false ;
};if _gdge {return _dd [_bce *64],_cdcg -_bce *64,false ;};return _fb [_bce *64],_cdcg -_bce *64,false ;};func (_bdgg *Decoder )tryFetchEOL1 ()(bool ,error ){_bdcd ,_cec :=_bdgg ._fcg .ReadBits (13);if _cec !=nil {return false ,_cec ;};return _bdcd ==0x3,nil ;
};type DecodeOptions struct{Columns int ;Rows int ;K int ;EncodedByteAligned bool ;BlackIsOne bool ;EndOfBlock bool ;EndOfLine bool ;DamagedRowsBeforeError int ;};func _ddf (_fced []byte ,_gff int ,_efc int ,_fbe bool )([]byte ,int ){var (_ded code ;_ceb bool ;
);for !_ceb {_ded ,_efc ,_ceb =_cbf (_efc ,_fbe );_fced ,_gff =_dcfa (_fced ,_gff ,_ded );};return _fced ,_gff ;};func _dgg (_gbda []byte ,_bfa int ,_efagf code )([]byte ,int ){_fcd :=true ;var _feag []byte ;_feag ,_bfa =_dcfa (nil ,_bfa ,_efagf );_faa :=0;
var _bgd int ;for _faa < len (_gbda ){_bgd ,_faa =_egde (_gbda ,_fcd ,_faa );_feag ,_bfa =_ddf (_feag ,_bfa ,_bgd ,_fcd );_fcd =!_fcd ;};return _feag ,_bfa %8;};func (_ecgg *Encoder )encodeG32D (_fba [][]byte )[]byte {var _cfg []byte ;var _dcd int ;for _adbc :=0;
_adbc < len (_fba );_adbc +=_ecgg .K {if _ecgg .Rows > 0&&!_ecgg .EndOfBlock &&_adbc ==_ecgg .Rows {break ;};_dcfb ,_ggg :=_dgg (_fba [_adbc ],_dcd ,_gda );_cfg =_ecgg .appendEncodedRow (_cfg ,_dcfb ,_dcd );if _ecgg .EncodedByteAlign {_ggg =0;};_dcd =_ggg ;
for _ddga :=_adbc +1;_ddga < (_adbc +_ecgg .K )&&_ddga < len (_fba );_ddga ++{if _ecgg .Rows > 0&&!_ecgg .EndOfBlock &&_ddga ==_ecgg .Rows {break ;};_gcgd ,_fagc :=_dcfa (nil ,_dcd ,_gb );var _daaa ,_eea ,_dbd int ;_eda :=-1;for _eda < len (_fba [_ddga ]){_daaa =_fbfd (_fba [_ddga ],_eda );
_eea =_bcee (_fba [_ddga ],_fba [_ddga -1],_eda );_dbd =_fbfd (_fba [_ddga -1],_eea );if _dbd < _daaa {_gcgd ,_fagc =_efb (_gcgd ,_fagc );_eda =_dbd ;}else {if _a .Abs (float64 (_eea -_daaa ))> 3{_gcgd ,_fagc ,_eda =_bgde (_fba [_ddga ],_gcgd ,_fagc ,_eda ,_daaa );
}else {_gcgd ,_fagc =_cdf (_gcgd ,_fagc ,_daaa ,_eea );_eda =_daaa ;};};};_cfg =_ecgg .appendEncodedRow (_cfg ,_gcgd ,_dcd );if _ecgg .EncodedByteAlign {_fagc =0;};_dcd =_fagc %8;};};if _ecgg .EndOfBlock {_afe ,_ :=_bdagf (_dcd );_cfg =_ecgg .appendEncodedRow (_cfg ,_afe ,_dcd );
};return _cfg ;};func _cea (_dcb ,_dedc int )code {var _fef code ;switch _dedc -_dcb {case -1:_fef =_dg ;case -2:_fef =_bdg ;case -3:_fef =_ee ;case 0:_fef =_gecf ;case 1:_fef =_cdb ;case 2:_fef =_fad ;case 3:_fef =_geb ;};return _fef ;};var (_fbf byte =1;
_cac byte =0;);type Encoder struct{K int ;EndOfLine bool ;EncodedByteAlign bool ;Columns int ;Rows int ;EndOfBlock bool ;BlackIs1 bool ;DamagedRowsBeforeError int ;};func (_agd *treeNode )walk (_gafd bool )*treeNode {if _gafd {return _agd ._gag ;};return _agd ._eag ;
};type Decoder struct{_ec int ;_bdc int ;_fd int ;_gdg []byte ;_afa int ;_agec bool ;_gde bool ;_ff bool ;_fbd bool ;_ebe bool ;_gcb bool ;_dag bool ;_cdba int ;_beb int ;_cge []int ;_gecfa []int ;_aaf int ;_df int ;_eac int ;_dde int ;_fcg *_b .Reader ;
_ecb tiffType ;_bda error ;};func (_de *Decoder )decoderRowType41D ()error {if _de ._dag {_de ._fcg .Align ();};_de ._fcg .Mark ();var (_adb bool ;_eca error ;);if _de ._ebe {_adb ,_eca =_de .tryFetchEOL ();if _eca !=nil {return _eca ;};if !_adb {return _age ;
};}else {_adb ,_eca =_de .looseFetchEOL ();if _eca !=nil {return _eca ;};};if !_adb {_de ._fcg .Reset ();};if _adb &&_de ._gcb {_de ._fcg .Mark ();for _gbb :=0;_gbb < 5;_gbb ++{_adb ,_eca =_de .tryFetchEOL ();if _eca !=nil {if _f .Is (_eca ,_ea .EOF ){if _gbb ==0{break ;
};return _fe ;};};if _adb {continue ;};if _gbb > 0{return _fe ;};break ;};if _adb {return _ea .EOF ;};_de ._fcg .Reset ();};if _eca =_de .decode1D ();_eca !=nil {return _eca ;};return nil ;};func (_dgc *Decoder )decodeRowType4 ()error {if !_dgc ._agec {return _dgc .decoderRowType41D ();
};if _dgc ._dag {_dgc ._fcg .Align ();};_dgc ._fcg .Mark ();_cc ,_cdc :=_dgc .tryFetchEOL ();if _cdc !=nil {return _cdc ;};if !_cc &&_dgc ._ebe {_dgc ._eac ++;if _dgc ._eac > _dgc ._afa {return _age ;};_dgc ._fcg .Reset ();};if !_cc {_dgc ._fcg .Reset ();
};_fadg ,_cdc :=_dgc ._fcg .ReadBool ();if _cdc !=nil {return _cdc ;};if _fadg {if _cc &&_dgc ._gcb {if _cdc =_dgc .tryFetchRTC2D ();_cdc !=nil {return _cdc ;};};_cdc =_dgc .decode1D ();}else {_cdc =_dgc .decode2D ();};if _cdc !=nil {return _cdc ;};return nil ;
};func (_ecab *Decoder )decode2D ()error {_ecab ._aaf =_ecab ._df ;_ecab ._gecfa ,_ecab ._cge =_ecab ._cge ,_ecab ._gecfa ;_gg :=true ;var (_fea bool ;_ddg int ;_gbc error ;);_ecab ._df =0;_gcg :for _ddg < _ecab ._ec {_edg :=_c ._cab ;for {_fea ,_gbc =_ecab ._fcg .ReadBool ();
if _gbc !=nil {return _gbc ;};_edg =_edg .walk (_fea );if _edg ==nil {continue _gcg ;};if !_edg ._fca {continue ;};switch _edg ._fagcb {case _be :var _dfe int ;if _gg {_dfe ,_gbc =_ecab .decodeRun (_fc );}else {_dfe ,_gbc =_ecab .decodeRun (_ef );};if _gbc !=nil {return _gbc ;
};_ddg +=_dfe ;_ecab ._gecfa [_ecab ._df ]=_ddg ;_ecab ._df ++;if _gg {_dfe ,_gbc =_ecab .decodeRun (_ef );}else {_dfe ,_gbc =_ecab .decodeRun (_fc );};if _gbc !=nil {return _gbc ;};_ddg +=_dfe ;_ecab ._gecfa [_ecab ._df ]=_ddg ;_ecab ._df ++;case _ab :_edc :=_ecab .getNextChangingElement (_ddg ,_gg )+1;
if _edc >=_ecab ._aaf {_ddg =_ecab ._ec ;}else {_ddg =_ecab ._cge [_edc ];};default:_bf :=_ecab .getNextChangingElement (_ddg ,_gg );if _bf >=_ecab ._aaf ||_bf ==-1{_ddg =_ecab ._ec +_edg ._fagcb ;}else {_ddg =_ecab ._cge [_bf ]+_edg ._fagcb ;};_ecab ._gecfa [_ecab ._df ]=_ddg ;
_ecab ._df ++;_gg =!_gg ;};continue _gcg ;};};return nil ;};func (_fgde *tree )fillWithNode (_bdgb ,_cbe int ,_eegg *treeNode )error {_bab :=_fgde ._cab ;for _fcb :=0;_fcb < _bdgb ;_fcb ++{_dcfc :=uint (_bdgb -1-_fcb );_dca :=((_cbe >>_dcfc )&1)!=0;_eafa :=_bab .walk (_dca );
if _eafa !=nil {if _eafa ._fca {return _f .New ("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067");};_bab =_eafa ;continue ;
};if _fcb ==_bdgb -1{_eafa =_eegg ;}else {_eafa =&treeNode {};};if _cbe ==0{_eafa ._feaa =true ;};_bab .set (_dca ,_eafa );_bab =_eafa ;};return nil ;};var (_af map[int ]code ;_cd map[int ]code ;_fb map[int ]code ;_dd map[int ]code ;_fgf map[int ]code ;
_gc map[int ]byte ;_cad =code {Code :1<<4,BitsWritten :12};_gda =code {Code :3<<3,BitsWritten :13};_gb =code {Code :2<<3,BitsWritten :13};_aga =code {Code :1<<12,BitsWritten :4};_bc =code {Code :1<<13,BitsWritten :3};_gecf =code {Code :1<<15,BitsWritten :1};
_dg =code {Code :3<<13,BitsWritten :3};_bdg =code {Code :3<<10,BitsWritten :6};_ee =code {Code :3<<9,BitsWritten :7};_cdb =code {Code :2<<13,BitsWritten :3};_fad =code {Code :2<<10,BitsWritten :6};_geb =code {Code :2<<9,BitsWritten :7};);func _dgb (_eggb int )([]byte ,int ){var _dadg []byte ;
for _dfge :=0;_dfge < 2;_dfge ++{_dadg ,_eggb =_dcfa (_dadg ,_eggb ,_cad );};return _dadg ,_eggb %8;};var _cgg =[...][]uint16 {{2,3,4,5,6,7},{128,8,9,64,10,11},{192,1664,16,17,13,14,15,1,12},{26,21,28,27,18,24,25,22,256,23,20,19},{33,34,35,36,37,38,31,32,29,53,54,39,40,41,42,43,44,30,61,62,63,0,320,384,45,59,60,46,49,50,51,52,55,56,57,58,448,512,640,576,47,48},{1472,1536,1600,1728,704,768,832,896,960,1024,1088,1152,1216,1280,1344,1408},{},{1792,1856,1920},{1984,2048,2112,2176,2240,2304,2368,2432,2496,2560}};
func (_ccb *Decoder )tryFetchRTC2D ()(_bac error ){_ccb ._fcg .Mark ();var _cee bool ;for _ada :=0;_ada < 5;_ada ++{_cee ,_bac =_ccb .tryFetchEOL1 ();if _bac !=nil {if _f .Is (_bac ,_ea .EOF ){if _ada ==0{break ;};return _fe ;};};if _cee {continue ;};if _ada > 0{return _fe ;
};break ;};if _cee {return _ea .EOF ;};_ccb ._fcg .Reset ();return _bac ;};func (_dadb *tree )fill (_ede ,_fefe ,_dcgc int )error {_acg :=_dadb ._cab ;for _gcd :=0;_gcd < _ede ;_gcd ++{_fdg :=_ede -1-_gcd ;_afc :=((_fefe >>uint (_fdg ))&1)!=0;_ddad :=_acg .walk (_afc );
if _ddad !=nil {if _ddad ._fca {return _f .New ("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067");};_acg =_ddad ;continue ;
};_ddad =&treeNode {};if _gcd ==_ede -1{_ddad ._fagcb =_dcgc ;_ddad ._fca =true ;};if _fefe ==0{_ddad ._feaa =true ;};_acg .set (_afc ,_ddad );_acg =_ddad ;};return nil ;};var _gec =[...][]uint16 {{0x7,0x8,0xb,0xc,0xe,0xf},{0x12,0x13,0x14,0x1b,0x7,0x8},{0x17,0x18,0x2a,0x2b,0x3,0x34,0x35,0x7,0x8},{0x13,0x17,0x18,0x24,0x27,0x28,0x2b,0x3,0x37,0x4,0x8,0xc},{0x12,0x13,0x14,0x15,0x16,0x17,0x1a,0x1b,0x2,0x24,0x25,0x28,0x29,0x2a,0x2b,0x2c,0x2d,0x3,0x32,0x33,0x34,0x35,0x36,0x37,0x4,0x4a,0x4b,0x5,0x52,0x53,0x54,0x55,0x58,0x59,0x5a,0x5b,0x64,0x65,0x67,0x68,0xa,0xb},{0x98,0x99,0x9a,0x9b,0xcc,0xcd,0xd2,0xd3,0xd4,0xd5,0xd6,0xd7,0xd8,0xd9,0xda,0xdb},{},{0x8,0xc,0xd},{0x12,0x13,0x14,0x15,0x16,0x17,0x1c,0x1d,0x1e,0x1f}};
func (_acde *Encoder )appendEncodedRow (_gbf ,_ecd []byte ,_bfb int )[]byte {if len (_gbf )> 0&&_bfb !=0&&!_acde .EncodedByteAlign {_gbf [len (_gbf )-1]=_gbf [len (_gbf )-1]|_ecd [0];_gbf =append (_gbf ,_ecd [1:]...);}else {_gbf =append (_gbf ,_ecd ...);
};return _gbf ;};func _bdagf (_dade int )([]byte ,int ){var _bdcc []byte ;for _fede :=0;_fede < 6;_fede ++{_bdcc ,_dade =_dcfa (_bdcc ,_dade ,_gda );};return _bdcc ,_dade %8;};func (_fed *Decoder )looseFetchEOL ()(bool ,error ){_bge ,_ffd :=_fed ._fcg .ReadBits (12);
if _ffd !=nil {return false ,_ffd ;};switch _bge {case 0x1:return true ,nil ;case 0x0:for {_bae ,_egb :=_fed ._fcg .ReadBool ();if _egb !=nil {return false ,_egb ;};if _bae {return true ,nil ;};};default:return false ,nil ;};};func (_cff *treeNode )set (_ddge bool ,_dfbf *treeNode ){if !_ddge {_cff ._eag =_dfbf ;
}else {_cff ._gag =_dfbf ;};};func _egde (_beg []byte ,_gdae bool ,_bdca int )(int ,int ){_aca :=0;for _bdca < len (_beg ){if _gdae {if _beg [_bdca ]!=_fbf {break ;};}else {if _beg [_bdca ]!=_cac {break ;};};_aca ++;_bdca ++;};return _aca ,_bdca ;};func (_agf *Decoder )fetch ()error {if _agf ._cdba ==-1{return nil ;
};if _agf ._beb < _agf ._cdba {return nil ;};_agf ._cdba =0;_ba :=_agf .decodeRow ();if _ba !=nil {if !_f .Is (_ba ,_ea .EOF ){return _ba ;};if _agf ._cdba !=0{return _ba ;};_agf ._cdba =-1;};_agf ._beb =0;return nil ;};func (_bee *Decoder )decodeRun (_fgg *tree )(int ,error ){var _agb int ;
_bgb :=_fgg ._cab ;for {_ecac ,_adc :=_bee ._fcg .ReadBool ();if _adc !=nil {return 0,_adc ;};_bgb =_bgb .walk (_ecac );if _bgb ==nil {return 0,_f .New ("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d");
};if _bgb ._fca {_agb +=_bgb ._fagcb ;switch {case _bgb ._fagcb >=64:_bgb =_fgg ._cab ;case _bgb ._fagcb >=0:return _agb ,nil ;default:return _bee ._ec ,nil ;};};};};func (_bdag *Decoder )decodeRowType2 ()error {if _bdag ._dag {_bdag ._fcg .Align ();};
if _bbd :=_bdag .decode1D ();_bbd !=nil {return _bbd ;};return nil ;};func (_ggc *Decoder )getNextChangingElement (_cfc int ,_dcf bool )int {_fadb :=0;if !_dcf {_fadb =1;};_gfe :=int (uint32 (_ggc ._dde )&0xFFFFFFFE)+_fadb ;if _gfe > 2{_gfe -=2;};if _cfc ==0{return _gfe ;
};for _ecf :=_gfe ;_ecf < _ggc ._aaf ;_ecf +=2{if _cfc < _ggc ._cge [_ecf ]{_ggc ._dde =_ecf ;return _ecf ;};};return -1;};func (_acc *Decoder )decodeRowType6 ()error {if _acc ._dag {_acc ._fcg .Align ();};if _acc ._gcb {_acc ._fcg .Mark ();_abb ,_fee :=_acc .tryFetchEOL ();
if _fee !=nil {return _fee ;};if _abb {_abb ,_fee =_acc .tryFetchEOL ();if _fee !=nil {return _fee ;};if _abb {return _ea .EOF ;};};_acc ._fcg .Reset ();};return _acc .decode2D ();};func _dcfa (_gaf []byte ,_acaa int ,_ecae code )([]byte ,int ){_dae :=0;
for _dae < _ecae .BitsWritten {_gbcb :=_acaa /8;_fafa :=_acaa %8;if _gbcb >=len (_gaf ){_gaf =append (_gaf ,0);};_bcc :=8-_fafa ;_cegg :=_ecae .BitsWritten -_dae ;if _bcc > _cegg {_bcc =_cegg ;};if _dae < 8{_gaf [_gbcb ]=_gaf [_gbcb ]|byte (_ecae .Code >>uint (8+_fafa -_dae ))&_gc [8-_bcc -_fafa ];
}else {_gaf [_gbcb ]=_gaf [_gbcb ]|(byte (_ecae .Code <<uint (_dae -8))&_gc [8-_bcc ])>>uint (_fafa );};_acaa +=_bcc ;_dae +=_bcc ;};return _gaf ,_acaa ;};func NewDecoder (data []byte ,options DecodeOptions )(*Decoder ,error ){_aea :=&Decoder {_fcg :_b .NewReader (data ),_ec :options .Columns ,_bdc :options .Rows ,_afa :options .DamagedRowsBeforeError ,_gdg :make ([]byte ,(options .Columns +7)/8),_cge :make ([]int ,options .Columns +2),_gecfa :make ([]int ,options .Columns +2),_dag :options .EncodedByteAligned ,_fbd :options .BlackIsOne ,_ebe :options .EndOfLine ,_gcb :options .EndOfBlock };
switch {case options .K ==0:_aea ._ecb =_da ;if len (data )< 20{return nil ,_f .New ("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m");};_fgd :=data [:20];if _fgd [0]!=0||(_fgd [1]>>4!=1&&_fgd [1]!=1){_aea ._ecb =_cda ;
_bb :=(uint16 (_fgd [0])<<8+uint16 (_fgd [1]&0xff))>>4;for _cfe :=12;_cfe < 160;_cfe ++{_bb =(_bb <<1)+uint16 ((_fgd [_cfe /8]>>uint16 (7-(_cfe %8)))&0x01);if _bb &0xfff==1{_aea ._ecb =_da ;break ;};};};case options .K < 0:_aea ._ecb =_egg ;case options .K > 0:_aea ._ecb =_da ;
_aea ._agec =true ;};switch _aea ._ecb {case _cda ,_da ,_egg :default:return nil ,_f .New ("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065");
};return _aea ,nil ;};func _cdf (_ffg []byte ,_feef ,_fde ,_dbg int )([]byte ,int ){_bbe :=_cea (_fde ,_dbg );_ffg ,_feef =_dcfa (_ffg ,_feef ,_bbe );return _ffg ,_feef ;};func _ggb (_adcf [][]byte )[][]byte {_gfa :=make ([]byte ,len (_adcf [0]));for _gbab :=range _gfa {_gfa [_gbab ]=_fbf ;
};_adcf =append (_adcf ,[]byte {});for _egdd :=len (_adcf )-1;_egdd > 0;_egdd --{_adcf [_egdd ]=_adcf [_egdd -1];};_adcf [0]=_gfa ;return _adcf ;};type treeNode struct{_eag *treeNode ;_gag *treeNode ;_fagcb int ;_feaa bool ;_fca bool ;};func _bgde (_dbf ,_cdcge []byte ,_bde ,_bff ,_daee int )([]byte ,int ,int ){_aff :=_fbfd (_dbf ,_daee );
_cbg :=_bff >=0&&_dbf [_bff ]==_fbf ||_bff ==-1;_cdcge ,_bde =_dcfa (_cdcge ,_bde ,_bc );var _bed int ;if _bff > -1{_bed =_daee -_bff ;}else {_bed =_daee -_bff -1;};_cdcge ,_bde =_ddf (_cdcge ,_bde ,_bed ,_cbg );_cbg =!_cbg ;_dggd :=_aff -_daee ;_cdcge ,_bde =_ddf (_cdcge ,_bde ,_dggd ,_cbg );
_bff =_aff ;return _cdcge ,_bde ,_bff ;};func (_db *Decoder )decodeRow ()(_gdad error ){if !_db ._gcb &&_db ._bdc > 0&&_db ._bdc ==_db ._fd {return _ea .EOF ;};switch _db ._ecb {case _cda :_gdad =_db .decodeRowType2 ();case _da :_gdad =_db .decodeRowType4 ();
case _egg :_gdad =_db .decodeRowType6 ();};if _gdad !=nil {return _gdad ;};_dfbd :=0;_bcd :=true ;_db ._dde =0;for _agfg :=0;_agfg < _db ._df ;_agfg ++{_gef :=_db ._ec ;if _agfg !=_db ._df {_gef =_db ._gecfa [_agfg ];};if _gef > _db ._ec {_gef =_db ._ec ;
};_gce :=_dfbd /8;for _dfbd %8!=0&&_gef -_dfbd > 0{var _afg byte ;if !_bcd {_afg =1<<uint (7-(_dfbd %8));};_db ._gdg [_gce ]|=_afg ;_dfbd ++;};if _dfbd %8==0{_gce =_dfbd /8;var _dcg byte ;if !_bcd {_dcg =0xff;};for _gef -_dfbd > 7{_db ._gdg [_gce ]=_dcg ;
_dfbd +=8;_gce ++;};};for _gef -_dfbd > 0{if _dfbd %8==0{_db ._gdg [_gce ]=0;};var _dfg byte ;if !_bcd {_dfg =1<<uint (7-(_dfbd %8));};_db ._gdg [_gce ]|=_dfg ;_dfbd ++;};_bcd =!_bcd ;};if _dfbd !=_db ._ec {return _f .New ("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068");
};_db ._cdba =(_dfbd +7)/8;_db ._fd ++;return nil ;};func _fbfd (_cag []byte ,_gab int )int {if _gab >=len (_cag ){return _gab ;};if _gab < -1{_gab =-1;};var _fbea byte ;if _gab > -1{_fbea =_cag [_gab ];}else {_fbea =_fbf ;};_fdf :=_gab +1;for _fdf < len (_cag ){if _cag [_fdf ]!=_fbea {break ;
};_fdf ++;};return _fdf ;};type tiffType int ;const (_ tiffType =iota ;_cda ;_da ;_egg ;);var _dc =[...][]uint16 {{0x2,0x3},{0x2,0x3},{0x2,0x3},{0x3},{0x4,0x5},{0x4,0x5,0x7},{0x4,0x7},{0x18},{0x17,0x18,0x37,0x8,0xf},{0x17,0x18,0x28,0x37,0x67,0x68,0x6c,0x8,0xc,0xd},{0x12,0x13,0x14,0x15,0x16,0x17,0x1c,0x1d,0x1e,0x1f,0x24,0x27,0x28,0x2b,0x2c,0x33,0x34,0x35,0x37,0x38,0x52,0x53,0x54,0x55,0x56,0x57,0x58,0x59,0x5a,0x5b,0x64,0x65,0x66,0x67,0x68,0x69,0x6a,0x6b,0x6c,0x6d,0xc8,0xc9,0xca,0xcb,0xcc,0xcd,0xd2,0xd3,0xd4,0xd5,0xd6,0xd7,0xda,0xdb},{0x4a,0x4b,0x4c,0x4d,0x52,0x53,0x54,0x55,0x5a,0x5b,0x64,0x65,0x6c,0x6d,0x72,0x73,0x74,0x75,0x76,0x77}};
func (_gbd *Encoder )Encode (pixels [][]byte )[]byte {if _gbd .BlackIs1 {_fbf =0;_cac =1;}else {_fbf =1;_cac =0;};if _gbd .K ==0{return _gbd .encodeG31D (pixels );};if _gbd .K > 0{return _gbd .encodeG32D (pixels );};if _gbd .K < 0{return _gbd .encodeG4 (pixels );
};return nil ;};var _ac =[...][]uint16 {{3,2},{1,4},{6,5},{7},{9,8},{10,11,12},{13,14},{15},{16,17,0,18,64},{24,25,23,22,19,20,21,1792,1856,1920},{1984,2048,2112,2176,2240,2304,2368,2432,2496,2560,52,55,56,59,60,320,384,448,53,54,50,51,44,45,46,47,57,58,61,256,48,49,62,63,30,31,32,33,40,41,128,192,26,27,28,29,34,35,36,37,38,39,42,43},{640,704,768,832,1280,1344,1408,1472,1536,1600,1664,1728,512,576,896,960,1024,1088,1152,1216}};
func (_fgb tiffType )String ()string {switch _fgb {case _cda :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065";case _da :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034";
case _egg :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036";default:return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d";};};type code struct{Code uint16 ;BitsWritten int ;};func _bfc (_cbce int )([]byte ,int ){var _cfcf []byte ;for _baf :=0;
_baf < 6;_baf ++{_cfcf ,_cbce =_dcfa (_cfcf ,_cbce ,_cad );};return _cfcf ,_cbce %8;};var (_ad *treeNode ;_g *treeNode ;_ef *tree ;_fc *tree ;_fa *tree ;_c *tree ;_ge =-2000;_gf =-1000;_ab =-3000;_be =-4000;);func (_bfd *Decoder )tryFetchEOL ()(bool ,error ){_efag ,_ffc :=_bfd ._fcg .ReadBits (12);
if _ffc !=nil {return false ,_ffc ;};return _efag ==0x1,nil ;};var (_fe =_f .New ("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043");_age =_f .New ("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");
);func _gedd (_dgcc ,_dbb []byte ,_ddgf int ,_afgg bool )int {_dda :=_fbfd (_dbb ,_ddgf );if _dda < len (_dbb )&&(_ddgf ==-1&&_dbb [_dda ]==_fbf ||_ddgf >=0&&_ddgf < len (_dgcc )&&_dgcc [_ddgf ]==_dbb [_dda ]||_ddgf >=len (_dgcc )&&_afgg &&_dbb [_dda ]==_fbf ||_ddgf >=len (_dgcc )&&!_afgg &&_dbb [_dda ]==_cac ){_dda =_fbfd (_dbb ,_dda );
};return _dda ;};func (_dba *Decoder )decode1D ()error {var (_bebd int ;_gdgf error ;);_dad :=true ;_dba ._df =0;for {var _ceg int ;if _dad {_ceg ,_gdgf =_dba .decodeRun (_fc );}else {_ceg ,_gdgf =_dba .decodeRun (_ef );};if _gdgf !=nil {return _gdgf ;
};_bebd +=_ceg ;_dba ._gecfa [_dba ._df ]=_bebd ;_dba ._df ++;_dad =!_dad ;if _bebd >=_dba ._ec {break ;};};return nil ;};