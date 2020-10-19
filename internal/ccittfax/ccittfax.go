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

package ccittfax ;import (_f "errors";_a "math";);func _ddfg (_fg *decodingTreeNode ,_ebca uint16 ,_dcba int )(*int ,*code ){if _fg ==nil {return nil ,nil ;};if _dcba ==16{return _fg .RunLen ,_fg .Code ;};_cga :=_eeb (_ebca ,_dcba );_dcba ++;var _gbea *int ;var _ggab *code ;if _cga ==1{_gbea ,_ggab =_ddfg (_fg .Right ,_ebca ,_dcba );}else {_gbea ,_ggab =_ddfg (_fg .Left ,_ebca ,_dcba );};if _gbea ==nil {_gbea =_fg .RunLen ;_ggab =_fg .Code ;};return _gbea ,_ggab ;};func (_aae *Encoder )appendEncodedRow (_adge ,_accc []byte ,_eacg int )[]byte {if len (_adge )> 0&&_eacg !=0&&!_aae .EncodedByteAlign {_adge [len (_adge )-1]=_adge [len (_adge )-1]|_accc [0];_adge =append (_adge ,_accc [1:]...);}else {_adge =append (_adge ,_accc ...);};return _adge ;};func init (){_d =make (map[int ]code );_d [0]=code {Code :13<<8|3<<6,BitsWritten :10};_d [1]=code {Code :2<<(5+8),BitsWritten :3};_d [2]=code {Code :3<<(6+8),BitsWritten :2};_d [3]=code {Code :2<<(6+8),BitsWritten :2};_d [4]=code {Code :3<<(5+8),BitsWritten :3};_d [5]=code {Code :3<<(4+8),BitsWritten :4};_d [6]=code {Code :2<<(4+8),BitsWritten :4};_d [7]=code {Code :3<<(3+8),BitsWritten :5};_d [8]=code {Code :5<<(2+8),BitsWritten :6};_d [9]=code {Code :4<<(2+8),BitsWritten :6};_d [10]=code {Code :4<<(1+8),BitsWritten :7};_d [11]=code {Code :5<<(1+8),BitsWritten :7};_d [12]=code {Code :7<<(1+8),BitsWritten :7};_d [13]=code {Code :4<<8,BitsWritten :8};_d [14]=code {Code :7<<8,BitsWritten :8};_d [15]=code {Code :12<<8,BitsWritten :9};_d [16]=code {Code :5<<8|3<<6,BitsWritten :10};_d [17]=code {Code :6<<8,BitsWritten :10};_d [18]=code {Code :2<<8,BitsWritten :10};_d [19]=code {Code :12<<8|7<<5,BitsWritten :11};_d [20]=code {Code :13<<8,BitsWritten :11};_d [21]=code {Code :13<<8|4<<5,BitsWritten :11};_d [22]=code {Code :6<<8|7<<5,BitsWritten :11};_d [23]=code {Code :5<<8,BitsWritten :11};_d [24]=code {Code :2<<8|7<<5,BitsWritten :11};_d [25]=code {Code :3<<8,BitsWritten :11};_d [26]=code {Code :12<<8|10<<4,BitsWritten :12};_d [27]=code {Code :12<<8|11<<4,BitsWritten :12};_d [28]=code {Code :12<<8|12<<4,BitsWritten :12};_d [29]=code {Code :12<<8|13<<4,BitsWritten :12};_d [30]=code {Code :6<<8|8<<4,BitsWritten :12};_d [31]=code {Code :6<<8|9<<4,BitsWritten :12};_d [32]=code {Code :6<<8|10<<4,BitsWritten :12};_d [33]=code {Code :6<<8|11<<4,BitsWritten :12};_d [34]=code {Code :13<<8|2<<4,BitsWritten :12};_d [35]=code {Code :13<<8|3<<4,BitsWritten :12};_d [36]=code {Code :13<<8|4<<4,BitsWritten :12};_d [37]=code {Code :13<<8|5<<4,BitsWritten :12};_d [38]=code {Code :13<<8|6<<4,BitsWritten :12};_d [39]=code {Code :13<<8|7<<4,BitsWritten :12};_d [40]=code {Code :6<<8|12<<4,BitsWritten :12};_d [41]=code {Code :6<<8|13<<4,BitsWritten :12};_d [42]=code {Code :13<<8|10<<4,BitsWritten :12};_d [43]=code {Code :13<<8|11<<4,BitsWritten :12};_d [44]=code {Code :5<<8|4<<4,BitsWritten :12};_d [45]=code {Code :5<<8|5<<4,BitsWritten :12};_d [46]=code {Code :5<<8|6<<4,BitsWritten :12};_d [47]=code {Code :5<<8|7<<4,BitsWritten :12};_d [48]=code {Code :6<<8|4<<4,BitsWritten :12};_d [49]=code {Code :6<<8|5<<4,BitsWritten :12};_d [50]=code {Code :5<<8|2<<4,BitsWritten :12};_d [51]=code {Code :5<<8|3<<4,BitsWritten :12};_d [52]=code {Code :2<<8|4<<4,BitsWritten :12};_d [53]=code {Code :3<<8|7<<4,BitsWritten :12};_d [54]=code {Code :3<<8|8<<4,BitsWritten :12};_d [55]=code {Code :2<<8|7<<4,BitsWritten :12};_d [56]=code {Code :2<<8|8<<4,BitsWritten :12};_d [57]=code {Code :5<<8|8<<4,BitsWritten :12};_d [58]=code {Code :5<<8|9<<4,BitsWritten :12};_d [59]=code {Code :2<<8|11<<4,BitsWritten :12};_d [60]=code {Code :2<<8|12<<4,BitsWritten :12};_d [61]=code {Code :5<<8|10<<4,BitsWritten :12};_d [62]=code {Code :6<<8|6<<4,BitsWritten :12};_d [63]=code {Code :6<<8|7<<4,BitsWritten :12};_cc =make (map[int ]code );_cc [0]=code {Code :53<<8,BitsWritten :8};_cc [1]=code {Code :7<<(2+8),BitsWritten :6};_cc [2]=code {Code :7<<(4+8),BitsWritten :4};_cc [3]=code {Code :8<<(4+8),BitsWritten :4};_cc [4]=code {Code :11<<(4+8),BitsWritten :4};_cc [5]=code {Code :12<<(4+8),BitsWritten :4};_cc [6]=code {Code :14<<(4+8),BitsWritten :4};_cc [7]=code {Code :15<<(4+8),BitsWritten :4};_cc [8]=code {Code :19<<(3+8),BitsWritten :5};_cc [9]=code {Code :20<<(3+8),BitsWritten :5};_cc [10]=code {Code :7<<(3+8),BitsWritten :5};_cc [11]=code {Code :8<<(3+8),BitsWritten :5};_cc [12]=code {Code :8<<(2+8),BitsWritten :6};_cc [13]=code {Code :3<<(2+8),BitsWritten :6};_cc [14]=code {Code :52<<(2+8),BitsWritten :6};_cc [15]=code {Code :53<<(2+8),BitsWritten :6};_cc [16]=code {Code :42<<(2+8),BitsWritten :6};_cc [17]=code {Code :43<<(2+8),BitsWritten :6};_cc [18]=code {Code :39<<(1+8),BitsWritten :7};_cc [19]=code {Code :12<<(1+8),BitsWritten :7};_cc [20]=code {Code :8<<(1+8),BitsWritten :7};_cc [21]=code {Code :23<<(1+8),BitsWritten :7};_cc [22]=code {Code :3<<(1+8),BitsWritten :7};_cc [23]=code {Code :4<<(1+8),BitsWritten :7};_cc [24]=code {Code :40<<(1+8),BitsWritten :7};_cc [25]=code {Code :43<<(1+8),BitsWritten :7};_cc [26]=code {Code :19<<(1+8),BitsWritten :7};_cc [27]=code {Code :36<<(1+8),BitsWritten :7};_cc [28]=code {Code :24<<(1+8),BitsWritten :7};_cc [29]=code {Code :2<<8,BitsWritten :8};_cc [30]=code {Code :3<<8,BitsWritten :8};_cc [31]=code {Code :26<<8,BitsWritten :8};_cc [32]=code {Code :27<<8,BitsWritten :8};_cc [33]=code {Code :18<<8,BitsWritten :8};_cc [34]=code {Code :19<<8,BitsWritten :8};_cc [35]=code {Code :20<<8,BitsWritten :8};_cc [36]=code {Code :21<<8,BitsWritten :8};_cc [37]=code {Code :22<<8,BitsWritten :8};_cc [38]=code {Code :23<<8,BitsWritten :8};_cc [39]=code {Code :40<<8,BitsWritten :8};_cc [40]=code {Code :41<<8,BitsWritten :8};_cc [41]=code {Code :42<<8,BitsWritten :8};_cc [42]=code {Code :43<<8,BitsWritten :8};_cc [43]=code {Code :44<<8,BitsWritten :8};_cc [44]=code {Code :45<<8,BitsWritten :8};_cc [45]=code {Code :4<<8,BitsWritten :8};_cc [46]=code {Code :5<<8,BitsWritten :8};_cc [47]=code {Code :10<<8,BitsWritten :8};_cc [48]=code {Code :11<<8,BitsWritten :8};_cc [49]=code {Code :82<<8,BitsWritten :8};_cc [50]=code {Code :83<<8,BitsWritten :8};_cc [51]=code {Code :84<<8,BitsWritten :8};_cc [52]=code {Code :85<<8,BitsWritten :8};_cc [53]=code {Code :36<<8,BitsWritten :8};_cc [54]=code {Code :37<<8,BitsWritten :8};_cc [55]=code {Code :88<<8,BitsWritten :8};_cc [56]=code {Code :89<<8,BitsWritten :8};_cc [57]=code {Code :90<<8,BitsWritten :8};_cc [58]=code {Code :91<<8,BitsWritten :8};_cc [59]=code {Code :74<<8,BitsWritten :8};_cc [60]=code {Code :75<<8,BitsWritten :8};_cc [61]=code {Code :50<<8,BitsWritten :8};_cc [62]=code {Code :51<<8,BitsWritten :8};_cc [63]=code {Code :52<<8,BitsWritten :8};_ad =make (map[int ]code );_ad [64]=code {Code :3<<8|3<<6,BitsWritten :10};_ad [128]=code {Code :12<<8|8<<4,BitsWritten :12};_ad [192]=code {Code :12<<8|9<<4,BitsWritten :12};_ad [256]=code {Code :5<<8|11<<4,BitsWritten :12};_ad [320]=code {Code :3<<8|3<<4,BitsWritten :12};_ad [384]=code {Code :3<<8|4<<4,BitsWritten :12};_ad [448]=code {Code :3<<8|5<<4,BitsWritten :12};_ad [512]=code {Code :3<<8|12<<3,BitsWritten :13};_ad [576]=code {Code :3<<8|13<<3,BitsWritten :13};_ad [640]=code {Code :2<<8|10<<3,BitsWritten :13};_ad [704]=code {Code :2<<8|11<<3,BitsWritten :13};_ad [768]=code {Code :2<<8|12<<3,BitsWritten :13};_ad [832]=code {Code :2<<8|13<<3,BitsWritten :13};_ad [896]=code {Code :3<<8|18<<3,BitsWritten :13};_ad [960]=code {Code :3<<8|19<<3,BitsWritten :13};_ad [1024]=code {Code :3<<8|20<<3,BitsWritten :13};_ad [1088]=code {Code :3<<8|21<<3,BitsWritten :13};_ad [1152]=code {Code :3<<8|22<<3,BitsWritten :13};_ad [1216]=code {Code :119<<3,BitsWritten :13};_ad [1280]=code {Code :2<<8|18<<3,BitsWritten :13};_ad [1344]=code {Code :2<<8|19<<3,BitsWritten :13};_ad [1408]=code {Code :2<<8|20<<3,BitsWritten :13};_ad [1472]=code {Code :2<<8|21<<3,BitsWritten :13};_ad [1536]=code {Code :2<<8|26<<3,BitsWritten :13};_ad [1600]=code {Code :2<<8|27<<3,BitsWritten :13};_ad [1664]=code {Code :3<<8|4<<3,BitsWritten :13};_ad [1728]=code {Code :3<<8|5<<3,BitsWritten :13};_af =make (map[int ]code );_af [64]=code {Code :27<<(3+8),BitsWritten :5};_af [128]=code {Code :18<<(3+8),BitsWritten :5};_af [192]=code {Code :23<<(2+8),BitsWritten :6};_af [256]=code {Code :55<<(1+8),BitsWritten :7};_af [320]=code {Code :54<<8,BitsWritten :8};_af [384]=code {Code :55<<8,BitsWritten :8};_af [448]=code {Code :100<<8,BitsWritten :8};_af [512]=code {Code :101<<8,BitsWritten :8};_af [576]=code {Code :104<<8,BitsWritten :8};_af [640]=code {Code :103<<8,BitsWritten :8};_af [704]=code {Code :102<<8,BitsWritten :9};_af [768]=code {Code :102<<8|1<<7,BitsWritten :9};_af [832]=code {Code :105<<8,BitsWritten :9};_af [896]=code {Code :105<<8|1<<7,BitsWritten :9};_af [960]=code {Code :106<<8,BitsWritten :9};_af [1024]=code {Code :106<<8|1<<7,BitsWritten :9};_af [1088]=code {Code :107<<8,BitsWritten :9};_af [1152]=code {Code :107<<8|1<<7,BitsWritten :9};_af [1216]=code {Code :108<<8,BitsWritten :9};_af [1280]=code {Code :108<<8|1<<7,BitsWritten :9};_af [1344]=code {Code :109<<8,BitsWritten :9};_af [1408]=code {Code :109<<8|1<<7,BitsWritten :9};_af [1472]=code {Code :76<<8,BitsWritten :9};_af [1536]=code {Code :76<<8|1<<7,BitsWritten :9};_af [1600]=code {Code :77<<8,BitsWritten :9};_af [1664]=code {Code :24<<(2+8),BitsWritten :6};_af [1728]=code {Code :77<<8|1<<7,BitsWritten :9};_ff =make (map[int ]code );_ff [1792]=code {Code :1<<8,BitsWritten :11};_ff [1856]=code {Code :1<<8|4<<5,BitsWritten :11};_ff [1920]=code {Code :1<<8|5<<5,BitsWritten :11};_ff [1984]=code {Code :1<<8|2<<4,BitsWritten :12};_ff [2048]=code {Code :1<<8|3<<4,BitsWritten :12};_ff [2112]=code {Code :1<<8|4<<4,BitsWritten :12};_ff [2176]=code {Code :1<<8|5<<4,BitsWritten :12};_ff [2240]=code {Code :1<<8|6<<4,BitsWritten :12};_ff [2304]=code {Code :1<<8|7<<4,BitsWritten :12};_ff [2368]=code {Code :1<<8|12<<4,BitsWritten :12};_ff [2432]=code {Code :1<<8|13<<4,BitsWritten :12};_ff [2496]=code {Code :1<<8|14<<4,BitsWritten :12};_ff [2560]=code {Code :1<<8|15<<4,BitsWritten :12};_dc =make (map[int ]byte );_dc [0]=0xFF;_dc [1]=0xFE;_dc [2]=0xFC;_dc [3]=0xF8;_dc [4]=0xF0;_dc [5]=0xE0;_dc [6]=0xC0;_dc [7]=0x80;_dc [8]=0x00;};func _cae (_bdaf []byte ,_abd int )(bool ,int ){return _aag (_bdaf ,_abd ,_ac )};func _aac (_acb [][]byte ,_ab []byte ,_fca bool ,_dbf ,_bb int )([]byte ,int ){_aeb :=_dgg (_ab ,_acb [len (_acb )-1],_dbf ,_fca );_bc :=_aeb +_bb ;if _dbf ==-1{_ab =_fefc (_ab ,_fca ,_bc -_dbf -1);}else {_ab =_fefc (_ab ,_fca ,_bc -_dbf );};_dbf =_bc ;return _ab ,_dbf ;};func init (){for _cac ,_efb :=range _cc {_cdg (_bab ,_efb ,0,_cac );};for _ce ,_bf :=range _af {_cdg (_bab ,_bf ,0,_ce );};for _gd ,_dg :=range _d {_cdg (_ca ,_dg ,0,_gd );};for _ae ,_gde :=range _ad {_cdg (_ca ,_gde ,0,_ae );};for _ga ,_cdb :=range _ff {_cdg (_bab ,_cdb ,0,_ga );_cdg (_ca ,_cdb ,0,_ga );};_cdg (_eb ,_b ,0,0);_cdg (_eb ,_dcb ,0,0);_cdg (_eb ,_e ,0,0);_cdg (_eb ,_cd ,0,0);_cdg (_eb ,_eg ,0,0);_cdg (_eb ,_ba ,0,0);_cdg (_eb ,_dd ,0,0);_cdg (_eb ,_ef ,0,0);_cdg (_eb ,_fe ,0,0);};func _ffd (_gdf []byte ,_dde int )(bool ,int ,error ){_dda :=_dde ;var _ged =false ;for _ed :=0;_ed < 6;_ed ++{_ged ,_dde =_cae (_gdf ,_dde );if !_ged {if _ed > 1{return false ,_dda ,_bd ;}else {_dde =_dda ;break ;};};};return _ged ,_dde ,nil ;};func _bfe (_dec ,_fcb []byte ,_bef int ,_dea bool )([]byte ,int ,error ){_dbb :=_bef ;var _fac int ;for _fac ,_bef =_gbe (_dec ,_bef ,_dea );_fac !=-1;_fac ,_bef =_gbe (_dec ,_bef ,_dea ){_fcb =_fefc (_fcb ,_dea ,_fac );if _fac < 64{break ;};};if _fac ==-1{return _fcb ,_dbb ,_ec ;};return _fcb ,_bef ,nil ;};func _fecb (_cbfd ,_fbc []byte ,_egeb ,_faf ,_eedcc int )([]byte ,int ,int ){_begg :=_gfed (_cbfd ,_eedcc );_gcda :=_faf >=0&&_cbfd [_faf ]==_aebe ||_faf ==-1;_fbc ,_egeb =_ffdd (_fbc ,_egeb ,_dcb );var _aed int ;if _faf > -1{_aed =_eedcc -_faf ;}else {_aed =_eedcc -_faf -1;};_fbc ,_egeb =_gfdb (_fbc ,_egeb ,_aed ,_gcda );_gcda =!_gcda ;_edae :=_begg -_eedcc ;_fbc ,_egeb =_gfdb (_fbc ,_egeb ,_edae ,_gcda );_faf =_begg ;return _fbc ,_egeb ,_faf ;};func (_cef *Encoder )encodeG31D (_bfa [][]byte )[]byte {var _gfa []byte ;_bdeac :=0;for _ggb :=range _bfa {if _cef .Rows > 0&&!_cef .EndOfBlock &&_ggb ==_cef .Rows {break ;};_gdc ,_dgd :=_dge (_bfa [_ggb ],_bdeac ,_g );_gfa =_cef .appendEncodedRow (_gfa ,_gdc ,_bdeac );if _cef .EncodedByteAlign {_dgd =0;};_bdeac =_dgd ;};if _cef .EndOfBlock {_afd ,_ :=_cdfb (_bdeac );_gfa =_cef .appendEncodedRow (_gfa ,_afd ,_bdeac );};return _gfa ;};func (_cea *Encoder )decodeG31D (_fd []byte )([][]byte ,error ){var _gb [][]byte ;var _adg int ;for (_adg /8)< len (_fd ){var _ggg bool ;_ggg ,_adg =_bcf (_fd ,_adg );if !_ggg {if _cea .EndOfLine {return nil ,_fc ;};}else {for _be :=0;_be < 5;_be ++{_ggg ,_adg =_bcf (_fd ,_adg );if !_ggg {if _be ==0{break ;};return nil ,_fa ;};};if _ggg {break ;};};var _ge []byte ;_ge ,_adg =_cea .decodeRow1D (_fd ,_adg );if _cea .EncodedByteAlign &&_adg %8!=0{_adg +=8-_adg %8;};_gb =append (_gb ,_ge );if _cea .Rows > 0&&!_cea .EndOfBlock &&len (_gb )>=_cea .Rows {break ;};};return _gb ,nil ;};func _eac (_deae []byte ,_fefd int )(bool ,int ){return _aag (_deae ,_fefd ,_gg )};type code struct{Code uint16 ;BitsWritten int ;};func _ecf (_bdd [][]byte ,_dbff []byte ,_bee bool ,_cba int )([]byte ,int ){_efc :=_dgg (_dbff ,_bdd [len (_bdd )-1],_cba ,_bee );_ced :=_gfed (_bdd [len (_bdd )-1],_efc );if _cba ==-1{_dbff =_fefc (_dbff ,_bee ,_ced -_cba -1);}else {_dbff =_fefc (_dbff ,_bee ,_ced -_cba );};_cba =_ced ;return _dbff ,_cba ;};var (_d map[int ]code ;_cc map[int ]code ;_ad map[int ]code ;_af map[int ]code ;_ff map[int ]code ;_dc map[int ]byte ;_g =code {Code :1<<4,BitsWritten :12};_ac =code {Code :3<<3,BitsWritten :13};_gg =code {Code :2<<3,BitsWritten :13};_b =code {Code :1<<12,BitsWritten :4};_dcb =code {Code :1<<13,BitsWritten :3};_e =code {Code :1<<15,BitsWritten :1};_cd =code {Code :3<<13,BitsWritten :3};_eg =code {Code :3<<10,BitsWritten :6};_ba =code {Code :3<<9,BitsWritten :7};_dd =code {Code :2<<13,BitsWritten :3};_ef =code {Code :2<<10,BitsWritten :6};_fe =code {Code :2<<9,BitsWritten :7};);type decodingTreeNode struct{Val byte ;RunLen *int ;Code *code ;Left *decodingTreeNode ;Right *decodingTreeNode ;};func (_cf *Encoder )Decode (encoded []byte )([][]byte ,error ){if _cf .BlackIs1 {_aebe =0;_bdea =1;}else {_aebe =1;_bdea =0;};if _cf .K ==0{return _cf .decodeG31D (encoded );};if _cf .K > 0{return _cf .decodeG32D (encoded );};if _cf .K < 4{return _cf .decodeG4 (encoded );};return nil ,nil ;};func _fae (_cace uint16 ,_bgf int ,_fcag bool )(int ,code ){var _eag *int ;var _bga *code ;if _fcag {_eag ,_bga =_ddfg (_bab ,_cace ,_bgf );}else {_eag ,_bga =_ddfg (_ca ,_cace ,_bgf );};if _eag ==nil {return -1,code {};};return *_eag ,*_bga ;};func (_deafa *Encoder )encodeG4 (_dbbe [][]byte )[]byte {_cbag :=make ([][]byte ,len (_dbbe ));copy (_cbag ,_dbbe );_cbag =_cdcbc (_cbag );var _fee []byte ;var _cdd int ;for _cbga :=1;_cbga < len (_cbag );_cbga ++{if _deafa .Rows > 0&&!_deafa .EndOfBlock &&_cbga ==(_deafa .Rows +1){break ;};var _eae []byte ;var _efbc ,_agg ,_bfg int ;_ddbf :=_cdd ;_bge :=-1;for _bge < len (_cbag [_cbga ]){_efbc =_gfed (_cbag [_cbga ],_bge );_agg =_gfebg (_cbag [_cbga ],_cbag [_cbga -1],_bge );_bfg =_gfed (_cbag [_cbga -1],_agg );if _bfg < _efbc {_eae ,_ddbf =_ffdd (_eae ,_ddbf ,_b );_bge =_bfg ;}else {if _a .Abs (float64 (_agg -_efbc ))> 3{_eae ,_ddbf ,_bge =_fecb (_cbag [_cbga ],_eae ,_ddbf ,_bge ,_efbc );}else {_eae ,_ddbf =_ecd (_eae ,_ddbf ,_efbc ,_agg );_bge =_efbc ;};};};_fee =_deafa .appendEncodedRow (_fee ,_eae ,_cdd );if _deafa .EncodedByteAlign {_ddbf =0;};_cdd =_ddbf %8;};if _deafa .EndOfBlock {_bad ,_ :=_fde (_cdd );_fee =_deafa .appendEncodedRow (_fee ,_bad ,_cdd );};return _fee ;};func _gfdb (_gcfd []byte ,_fgd int ,_dgeb int ,_bgd bool )([]byte ,int ){var (_dcda code ;_beb bool ;);for !_beb {_dcda ,_dgeb ,_beb =_edf (_dgeb ,_bgd );_gcfd ,_fgd =_ffdd (_gcfd ,_fgd ,_dcda );};return _gcfd ,_fgd ;};func _aag (_fbd []byte ,_daf int ,_cfd code )(bool ,int ){_gda :=_daf ;var (_aeee uint16 ;_baf int ;);_aeee ,_baf ,_daf =_cdbe (_fbd ,_daf );if _baf > 3{return false ,_gda ;};_aeee >>=uint (3-_baf );_aeee <<=3;if _aeee !=_cfd .Code {return false ,_gda ;}else {return true ,_daf -3+_baf ;};};func _gbe (_dbffb []byte ,_ccb int ,_geb bool )(int ,int ){var (_bde uint16 ;_eda int ;_fef int ;);_fef =_ccb ;_bde ,_eda ,_ccb =_cdbe (_dbffb ,_ccb );_ecb ,_aee :=_fae (_bde ,_eda ,_geb );if _ecb ==-1{return -1,_fef ;};return _ecb ,_fef +_aee .BitsWritten ;};func _ccg (_bbd int )([]byte ,int ){var _bbg []byte ;for _bdg :=0;_bdg < 6;_bdg ++{_bbg ,_bbd =_ffdd (_bbg ,_bbd ,_ac );};return _bbg ,_bbd %8;};func _cdcbc (_ccde [][]byte )[][]byte {_gaca :=make ([]byte ,len (_ccde [0]));for _abbc :=range _gaca {_gaca [_abbc ]=_aebe ;};_ccde =append (_ccde ,[]byte {});for _egc :=len (_ccde )-1;_egc > 0;_egc --{_ccde [_egc ]=_ccde [_egc -1];};_ccde [0]=_gaca ;return _ccde ;};func (_acbc *Encoder )decodeRow1D (_ecg []byte ,_fcc int )([]byte ,int ){var _befc []byte ;_fccc :=true ;var _cdf int ;_cdf ,_fcc =_gbe (_ecg ,_fcc ,_fccc );for _cdf !=-1{_befc =_fefc (_befc ,_fccc ,_cdf );if _cdf < 64{if len (_befc )>=_acbc .Columns {break ;};_fccc =!_fccc ;};_cdf ,_fcc =_gbe (_ecg ,_fcc ,_fccc );};return _befc ,_fcc ;};func _cdbe (_edd []byte ,_ege int )(uint16 ,int ,int ){_gbc :=_ege ;_daa :=_ege /8;_ege %=8;if _daa >=len (_edd ){return 0,16,_gbc ;};_gcd :=byte (0xFF>>uint (_ege ));_gfeb :=uint16 ((_edd [_daa ]&_gcd )<<uint (_ege ))<<8;_dac :=8-_ege ;_daa ++;_ege =0;if _daa >=len (_edd ){return _gfeb >>(16-uint (_dac )),16-_dac ,_gbc +_dac ;};_gfeb |=uint16 (_edd [_daa ])<<(8-uint (_dac ));_dac +=8;_daa ++;_ege =0;if _daa >=len (_edd ){return _gfeb >>(16-uint (_dac )),16-_dac ,_gbc +_dac ;};if _dac ==16{return _gfeb ,0,_gbc +_dac ;};_dbfd :=16-_dac ;_gfeb |=uint16 (_edd [_daa ]>>(8-uint (_dbfd )));return _gfeb ,0,_gbc +16;};var (_aebe byte =1;_bdea byte =0;);func _fde (_feb int )([]byte ,int ){var _gdfe []byte ;for _gca :=0;_gca < 2;_gca ++{_gdfe ,_feb =_ffdd (_gdfe ,_feb ,_g );};return _gdfe ,_feb %8;};func _fba (_dgbf []byte ,_ebe int )([]byte ,int ){return _ffdd (_dgbf ,_ebe ,_b )};func _beg (_gfb []byte ,_bbc int )(code ,int ,bool ){var (_ggge uint16 ;_abc int ;_bfeb int ;);_bfeb =_bbc ;_ggge ,_abc ,_bbc =_cdbe (_gfb ,_bbc );_abb ,_cfb :=_befa (_ggge ,_abc );if !_cfb {return code {},_bfeb ,false ;};return _abb ,_bfeb +_abb .BitsWritten ,true ;};func _fefc (_eed []byte ,_dbd bool ,_abca int )[]byte {if _abca < 0{return _eed ;};_eedc :=make ([]byte ,_abca );if _dbd {for _cdc :=0;_cdc < len (_eedc );_cdc ++{_eedc [_cdc ]=_aebe ;};}else {for _gae :=0;_gae < len (_eedc );_gae ++{_eedc [_gae ]=_bdea ;};};_eed =append (_eed ,_eedc ...);return _eed ;};func _eec (_bdf ,_bda []byte ,_ebc int ,_ggga bool ,_adgd int )([]byte ,int ,int ,error ){_ded :=_ebc ;var _bba error ;_bda ,_ebc ,_bba =_bfe (_bdf ,_bda ,_ebc ,_ggga );if _bba !=nil {return _bda ,_ded ,_adgd ,_bba ;};_ggga =!_ggga ;_bda ,_ebc ,_bba =_bfe (_bdf ,_bda ,_ebc ,_ggga );if _bba !=nil {return _bda ,_ded ,_adgd ,_bba ;};_adgd =len (_bda );return _bda ,_ebc ,_adgd ,nil ;};var (_acc =_f .New ("\u0045\u004f\u0046\u0042 c\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074e\u0064");_bd =_f .New ("R\u0054\u0043\u0020\u0063od\u0065 \u0069\u0073\u0020\u0063\u006fr\u0072\u0075\u0070\u0074\u0065\u0064");_ec =_f .New ("\u0077\u0072o\u006e\u0067\u0020\u0063\u006f\u0064\u0065\u0020\u0069\u006e\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020mo\u0064\u0065");_fc =_f .New ("\u006e\u006f\u0020\u0045\u004f\u004c\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0074\u0068\u0065\u0020\u0045\u006e\u0064O\u0066\u004c\u0069\u006e\u0065\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0065");_fa =_f .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u004f\u004c");_gf =_f .New ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0032\u0044\u0020\u0063\u006f\u0064\u0065"););func (_aa *Encoder )decodeG32D (_cg []byte )([][]byte ,error ){var (_dgb [][]byte ;_ebg int ;_fec error ;);_aef :for (_ebg /8)< len (_cg ){var _ee bool ;_ee ,_ebg ,_fec =_ffd (_cg ,_ebg );if _fec !=nil {return nil ,_fec ;};if _ee {break ;};_ee ,_ebg =_cae (_cg ,_ebg );if !_ee {if _aa .EndOfLine {return nil ,_fc ;};};var _gga []byte ;_gga ,_ebg =_aa .decodeRow1D (_cg ,_ebg );if _aa .EncodedByteAlign &&_ebg %8!=0{_ebg +=8-_ebg %8;};if _gga !=nil {_dgb =append (_dgb ,_gga );};if _aa .Rows > 0&&!_aa .EndOfBlock &&len (_dgb )>=_aa .Rows {break ;};for _gdg :=1;_gdg < _aa .K &&(_ebg /8)< len (_cg );_gdg ++{_ee ,_ebg =_eac (_cg ,_ebg );if !_ee {_ee ,_ebg ,_fec =_ffd (_cg ,_ebg );if _fec !=nil {return nil ,_fec ;};if _ee {break _aef ;}else {if _aa .EndOfLine {return nil ,_fc ;};};};var (_ccc code ;_da bool ;);_eeg :=true ;var _ccd []byte ;_fb :=-1;for _ccc ,_ebg ,_da =_beg (_cg ,_ebg );_da ;_ccc ,_ebg ,_da =_beg (_cg ,_ebg ){switch _ccc {case _b :_ccd ,_fb =_ecf (_dgb ,_ccd ,_eeg ,_fb );case _dcb :_ccd ,_ebg ,_fb ,_fec =_eec (_cg ,_ccd ,_ebg ,_eeg ,_fb );if _fec !=nil {return nil ,_fec ;};case _e :_ccd ,_fb =_aac (_dgb ,_ccd ,_eeg ,_fb ,0);_eeg =!_eeg ;case _cd :_ccd ,_fb =_aac (_dgb ,_ccd ,_eeg ,_fb ,1);_eeg =!_eeg ;case _eg :_ccd ,_fb =_aac (_dgb ,_ccd ,_eeg ,_fb ,2);_eeg =!_eeg ;case _ba :_ccd ,_fb =_aac (_dgb ,_ccd ,_eeg ,_fb ,3);_eeg =!_eeg ;case _dd :_ccd ,_fb =_aac (_dgb ,_ccd ,_eeg ,_fb ,-1);_eeg =!_eeg ;case _ef :_ccd ,_fb =_aac (_dgb ,_ccd ,_eeg ,_fb ,-2);_eeg =!_eeg ;case _fe :_ccd ,_fb =_aac (_dgb ,_ccd ,_eeg ,_fb ,-3);_eeg =!_eeg ;};if len (_ccd )>=_aa .Columns {break ;};};if _aa .EncodedByteAlign &&_ebg %8!=0{_ebg +=8-_ebg %8;};if _ccd !=nil {_dgb =append (_dgb ,_ccd );};if _aa .Rows > 0&&!_aa .EndOfBlock &&len (_dgb )>=_aa .Rows {break _aef ;};};};return _dgb ,nil ;};func _gge (_cdae ,_fdfe int )code {var _fbcc code ;switch _fdfe -_cdae {case -1:_fbcc =_cd ;case -2:_fbcc =_eg ;case -3:_fbcc =_ba ;case 0:_fbcc =_e ;case 1:_fbcc =_dd ;case 2:_fbcc =_ef ;case 3:_fbcc =_fe ;};return _fbcc ;};func _cdfb (_dfd int )([]byte ,int ){var _dcca []byte ;for _eece :=0;_eece < 6;_eece ++{_dcca ,_dfd =_ffdd (_dcca ,_dfd ,_g );};return _dcca ,_dfd %8;};func _dge (_gbeac []byte ,_fab int ,_agc code )([]byte ,int ){_bdeg :=true ;var _gac []byte ;_gac ,_fab =_ffdd (nil ,_fab ,_agc );_eef :=0;var _dcd int ;for _eef < len (_gbeac ){_dcd ,_eef =_ecge (_gbeac ,_bdeg ,_eef );_gac ,_fab =_gfdb (_gac ,_fab ,_dcd ,_bdeg );_bdeg =!_bdeg ;};return _gac ,_fab %8;};func (_cb *Encoder )decodeG4 (_fdb []byte )([][]byte ,error ){_ddf :=make ([]byte ,_cb .Columns );for _caf :=range _ddf {_ddf [_caf ]=_aebe ;};_ebb :=make ([][]byte ,1);_ebb [0]=_ddf ;var (_cbg bool ;_agd error ;_acf int ;);for (_acf /8)< len (_fdb ){_cbg ,_acf ,_agd =_fbb (_fdb ,_acf );if _agd !=nil {return nil ,_agd ;};if _cbg {break ;};var (_cbf code ;_caa bool ;);_de :=true ;var _gee []byte ;_ea :=-1;for _ea < _cb .Columns {_cbf ,_acf ,_caa =_beg (_fdb ,_acf );if !_caa {return nil ,_gf ;};switch _cbf {case _b :_gee ,_ea =_ecf (_ebb ,_gee ,_de ,_ea );case _dcb :_gee ,_acf ,_ea ,_agd =_eec (_fdb ,_gee ,_acf ,_de ,_ea );if _agd !=nil {return nil ,_agd ;};case _e :_gee ,_ea =_aac (_ebb ,_gee ,_de ,_ea ,0);_de =!_de ;case _cd :_gee ,_ea =_aac (_ebb ,_gee ,_de ,_ea ,1);_de =!_de ;case _eg :_gee ,_ea =_aac (_ebb ,_gee ,_de ,_ea ,2);_de =!_de ;case _ba :_gee ,_ea =_aac (_ebb ,_gee ,_de ,_ea ,3);_de =!_de ;case _dd :_gee ,_ea =_aac (_ebb ,_gee ,_de ,_ea ,-1);_de =!_de ;case _ef :_gee ,_ea =_aac (_ebb ,_gee ,_de ,_ea ,-2);_de =!_de ;case _fe :_gee ,_ea =_aac (_ebb ,_gee ,_de ,_ea ,-3);_de =!_de ;};if len (_gee )>=_cb .Columns {break ;};};if _cb .EncodedByteAlign &&_acf %8!=0{_acf +=8-_acf %8;};_ebb =append (_ebb ,_gee );if _cb .Rows > 0&&!_cb .EndOfBlock &&len (_ebb )>=(_cb .Rows +1){break ;};};_ebb =_ebb [1:];return _ebb ,nil ;};func _ecd (_edfd []byte ,_fea ,_agdg ,_bbe int )([]byte ,int ){_cda :=_gge (_agdg ,_bbe );_edfd ,_fea =_ffdd (_edfd ,_fea ,_cda );return _edfd ,_fea ;};func _cdg (_fbbb *decodingTreeNode ,_gfd code ,_dcc int ,_bag int ){_df :=_eeb (_gfd .Code ,_dcc );_dcc ++;if _df ==1{if _fbbb .Right ==nil {_fbbb .Right =&decodingTreeNode {Val :_df };};if _dcc ==_gfd .BitsWritten {_fbbb .Right .RunLen =&_bag ;_fbbb .Right .Code =&_gfd ;}else {_cdg (_fbbb .Right ,_gfd ,_dcc ,_bag );};}else {if _fbbb .Left ==nil {_fbbb .Left =&decodingTreeNode {Val :_df };};if _dcc ==_gfd .BitsWritten {_fbbb .Left .RunLen =&_bag ;_fbbb .Left .Code =&_gfd ;}else {_cdg (_fbbb .Left ,_gfd ,_dcc ,_bag );};};};func _bcf (_fed []byte ,_faa int )(bool ,int ){_cdcb :=_faa ;var (_gebd uint16 ;_gbd int ;);_gebd ,_gbd ,_faa =_cdbe (_fed ,_faa );if _gbd > 4{return false ,_cdcb ;};_gebd >>=uint (4-_gbd );_gebd <<=4;if _gebd !=_g .Code {return false ,_cdcb ;}else {return true ,_faa -4+_gbd ;};};func _ecge (_fbf []byte ,_bdegd bool ,_acd int )(int ,int ){_bebf :=0;for _acd < len (_fbf ){if _bdegd {if _fbf [_acd ]!=_aebe {break ;};}else {if _fbf [_acd ]!=_bdea {break ;};};_bebf ++;_acd ++;};return _bebf ,_acd ;};var (_bab =&decodingTreeNode {Val :255};_ca =&decodingTreeNode {Val :255};_eb =&decodingTreeNode {Val :255};);type Encoder struct{K int ;EndOfLine bool ;EncodedByteAlign bool ;Columns int ;Rows int ;EndOfBlock bool ;BlackIs1 bool ;DamagedRowsBeforeError int ;};func _befa (_fdbb uint16 ,_bcd int )(code ,bool ){_ ,_egf :=_ddfg (_eb ,_fdbb ,_bcd );if _egf ==nil {return code {},false ;};return *_egf ,true ;};func _fbb (_efe []byte ,_cccb int )(bool ,int ,error ){_efg :=_cccb ;var _cbfc bool ;_cbfc ,_cccb =_bcf (_efe ,_cccb );if _cbfc {_cbfc ,_cccb =_bcf (_efe ,_cccb );if _cbfc {return true ,_cccb ,nil ;}else {return false ,_efg ,_acc ;};};return false ,_efg ,nil ;};func _ffdd (_eegc []byte ,_dafc int ,_ccf code )([]byte ,int ){_def :=0;for _def < _ccf .BitsWritten {_fdf :=_dafc /8;_cefg :=_dafc %8;if _fdf >=len (_eegc ){_eegc =append (_eegc ,0);};_begc :=8-_cefg ;_ddaa :=_ccf .BitsWritten -_def ;if _begc > _ddaa {_begc =_ddaa ;};if _def < 8{_eegc [_fdf ]=_eegc [_fdf ]|byte (_ccf .Code >>uint (8+_cefg -_def ))&_dc [8-_begc -_cefg ];}else {_eegc [_fdf ]=_eegc [_fdf ]|(byte (_ccf .Code <<uint (_def -8))&_dc [8-_begc ])>>uint (_cefg );};_dafc +=_begc ;_def +=_begc ;};return _eegc ,_dafc ;};func _gfed (_fgc []byte ,_agf int )int {if _agf >=len (_fgc ){return _agf ;};if _agf < -1{_agf =-1;};var _bgdd byte ;if _agf > -1{_bgdd =_fgc [_agf ];}else {_bgdd =_aebe ;};_cefb :=_agf +1;for _cefb < len (_fgc ){if _fgc [_cefb ]!=_bgdd {break ;};_cefb ++;};return _cefb ;};func _dgg (_bfge ,_dfe []byte ,_abg int ,_efgc bool )int {_egd :=_gfed (_dfe ,_abg );if _egd < len (_dfe )&&(_abg ==-1&&_dfe [_egd ]==_aebe ||_abg >=0&&_abg < len (_bfge )&&_bfge [_abg ]==_dfe [_egd ]||_abg >=len (_bfge )&&_efgc &&_dfe [_egd ]==_aebe ||_abg >=len (_bfge )&&!_efgc &&_dfe [_egd ]==_bdea ){_egd =_gfed (_dfe ,_egd );};return _egd ;};func _eeb (_bg uint16 ,_gc int )byte {if _gc < 8{_bg >>=8;};_gc %=8;_gbb :=byte (0x01<<(7-uint (_gc )));return (byte (_bg )&_gbb )>>(7-uint (_gc ));};func _edf (_fdc int ,_bcc bool )(code ,int ,bool ){if _fdc < 64{if _bcc {return _cc [_fdc ],0,true ;}else {return _d [_fdc ],0,true ;};}else {_gcc :=_fdc /64;if _gcc > 40{return _ff [2560],_fdc -2560,false ;};if _gcc > 27{return _ff [_gcc *64],_fdc -_gcc *64,false ;};if _bcc {return _af [_gcc *64],_fdc -_gcc *64,false ;}else {return _ad [_gcc *64],_fdc -_gcc *64,false ;};};};func (_gfc *Encoder )Encode (pixels [][]byte )[]byte {if _gfc .BlackIs1 {_aebe =0;_bdea =1;}else {_aebe =1;_bdea =0;};if _gfc .K ==0{return _gfc .encodeG31D (pixels );};if _gfc .K > 0{return _gfc .encodeG32D (pixels );};if _gfc .K < 0{return _gfc .encodeG4 (pixels );};return nil ;};func _gfebg (_aec ,_ddc []byte ,_ddbg int )int {_cde :=_gfed (_ddc ,_ddbg );if _cde < len (_ddc )&&(_ddbg ==-1&&_ddc [_cde ]==_aebe ||_ddbg >=0&&_ddbg < len (_aec )&&_aec [_ddbg ]==_ddc [_cde ]||_ddbg >=len (_aec )&&_aec [_ddbg -1]!=_ddc [_cde ]){_cde =_gfed (_ddc ,_cde );};return _cde ;};func (_gcf *Encoder )encodeG32D (_ddb [][]byte )[]byte {var _fbe []byte ;var _ceg int ;for _dgdc :=0;_dgdc < len (_ddb );_dgdc +=_gcf .K {if _gcf .Rows > 0&&!_gcf .EndOfBlock &&_dgdc ==_gcf .Rows {break ;};_gbdb ,_eab :=_dge (_ddb [_dgdc ],_ceg ,_ac );_fbe =_gcf .appendEncodedRow (_fbe ,_gbdb ,_ceg );if _gcf .EncodedByteAlign {_eab =0;};_ceg =_eab ;for _dccc :=_dgdc +1;_dccc < (_dgdc +_gcf .K )&&_dccc < len (_ddb );_dccc ++{if _gcf .Rows > 0&&!_gcf .EndOfBlock &&_dccc ==_gcf .Rows {break ;};_babg ,_deg :=_ffdd (nil ,_ceg ,_gg );var _bcfd ,_gce ,_edc int ;_deaf :=-1;for _deaf < len (_ddb [_dccc ]){_bcfd =_gfed (_ddb [_dccc ],_deaf );_gce =_gfebg (_ddb [_dccc ],_ddb [_dccc -1],_deaf );_edc =_gfed (_ddb [_dccc -1],_gce );if _edc < _bcfd {_babg ,_deg =_fba (_babg ,_deg );_deaf =_edc ;}else {if _a .Abs (float64 (_gce -_bcfd ))> 3{_babg ,_deg ,_deaf =_fecb (_ddb [_dccc ],_babg ,_deg ,_deaf ,_bcfd );}else {_babg ,_deg =_ecd (_babg ,_deg ,_bcfd ,_gce );_deaf =_bcfd ;};};};_fbe =_gcf .appendEncodedRow (_fbe ,_babg ,_ceg );if _gcf .EncodedByteAlign {_deg =0;};_ceg =_deg %8;};};if _gcf .EndOfBlock {_facd ,_ :=_ccg (_ceg );_fbe =_gcf .appendEncodedRow (_fbe ,_facd ,_ceg );};return _fbe ;};