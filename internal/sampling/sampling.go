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

package sampling ;import (_cc "github.com/unidoc/unipdf/v3/internal/bitwise";_a "github.com/unidoc/unipdf/v3/internal/imageutil";_g "io";);func (_eb *Reader )ReadSample ()(uint32 ,error ){if _eb ._dd ==_eb ._cf .Height {return 0,_g .EOF ;};_ac ,_dg :=_eb ._f .ReadBits (byte (_eb ._cf .BitsPerComponent ));
if _dg !=nil {return 0,_dg ;};_eb ._ff --;if _eb ._ff ==0{_eb ._ff =_eb ._cf .ColorComponents ;_eb ._d ++;};if _eb ._d ==_eb ._cf .Width {if _eb ._e {_eb ._f .ConsumeRemainingBits ();};_eb ._d =0;_eb ._dd ++;};return uint32 (_ac ),nil ;};func NewWriter (img _a .ImageBase )*Writer {return &Writer {_dda :_cc .NewWriterMSB (img .Data ),_bf :img ,_afe :img .ColorComponents ,_cab :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};func (_eee *Writer )WriteSample (sample uint32 )error {if _ ,_ddd :=_eee ._dda .WriteBits (uint64 (sample ),_eee ._bf .BitsPerComponent );_ddd !=nil {return _ddd ;};_eee ._afe --;if _eee ._afe ==0{_eee ._afe =_eee ._bf .ColorComponents ;_eee ._ee ++;
};if _eee ._ee ==_eee ._bf .Width {if _eee ._cab {_eee ._dda .FinishByte ();};_eee ._ee =0;};return nil ;};func (_ae *Writer )WriteSamples (samples []uint32 )error {for _bc :=0;_bc < len (samples );_bc ++{if _ada :=_ae .WriteSample (samples [_bc ]);_ada !=nil {return _ada ;
};};return nil ;};type Reader struct{_cf _a .ImageBase ;_f *_cc .Reader ;_d ,_dd ,_ff int ;_e bool ;};func NewReader (img _a .ImageBase )*Reader {return &Reader {_f :_cc .NewReader (img .Data ),_cf :img ,_ff :img .ColorComponents ,_e :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_af []uint32 )error ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _gf []uint32 ;_eg :=bitsPerSample ;var _ag uint32 ;var _b byte ;_da :=0;_ce :=0;_agd :=0;for _agd < len (data ){if _da > 0{_adc :=_da ;
if _eg < _adc {_adc =_eg ;};_ag =(_ag <<uint (_adc ))|uint32 (_b >>uint (8-_adc ));_da -=_adc ;if _da > 0{_b =_b <<uint (_adc );}else {_b =0;};_eg -=_adc ;if _eg ==0{_gf =append (_gf ,_ag );_eg =bitsPerSample ;_ag =0;_ce ++;};}else {_bb :=data [_agd ];
_agd ++;_afc :=8;if _eg < _afc {_afc =_eg ;};_da =8-_afc ;_ag =(_ag <<uint (_afc ))|uint32 (_bb >>uint (_da ));if _afc < 8{_b =_bb <<uint (_afc );};_eg -=_afc ;if _eg ==0{_gf =append (_gf ,_ag );_eg =bitsPerSample ;_ag =0;_ce ++;};};};for _da >=bitsPerSample {_bd :=_da ;
if _eg < _bd {_bd =_eg ;};_ag =(_ag <<uint (_bd ))|uint32 (_b >>uint (8-_bd ));_da -=_bd ;if _da > 0{_b =_b <<uint (_bd );}else {_b =0;};_eg -=_bd ;if _eg ==0{_gf =append (_gf ,_ag );_eg =bitsPerSample ;_ag =0;_ce ++;};};return _gf ;};type SampleWriter interface{WriteSample (_ca uint32 )error ;
WriteSamples (_aa []uint32 )error ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _fc []uint32 ;_fd :=bitsPerOutputSample ;var _ecf uint32 ;var _fg uint32 ;_bdg :=0;_gg :=0;_fca :=0;for _fca < len (data ){if _bdg > 0{_ef :=_bdg ;
if _fd < _ef {_ef =_fd ;};_ecf =(_ecf <<uint (_ef ))|(_fg >>uint (bitsPerInputSample -_ef ));_bdg -=_ef ;if _bdg > 0{_fg =_fg <<uint (_ef );}else {_fg =0;};_fd -=_ef ;if _fd ==0{_fc =append (_fc ,_ecf );_fd =bitsPerOutputSample ;_ecf =0;_gg ++;};}else {_fgb :=data [_fca ];
_fca ++;_dc :=bitsPerInputSample ;if _fd < _dc {_dc =_fd ;};_bdg =bitsPerInputSample -_dc ;_ecf =(_ecf <<uint (_dc ))|(_fgb >>uint (_bdg ));if _dc < bitsPerInputSample {_fg =_fgb <<uint (_dc );};_fd -=_dc ;if _fd ==0{_fc =append (_fc ,_ecf );_fd =bitsPerOutputSample ;
_ecf =0;_gg ++;};};};for _bdg >=bitsPerOutputSample {_agg :=_bdg ;if _fd < _agg {_agg =_fd ;};_ecf =(_ecf <<uint (_agg ))|(_fg >>uint (bitsPerInputSample -_agg ));_bdg -=_agg ;if _bdg > 0{_fg =_fg <<uint (_agg );}else {_fg =0;};_fd -=_agg ;if _fd ==0{_fc =append (_fc ,_ecf );
_fd =bitsPerOutputSample ;_ecf =0;_gg ++;};};if _fd > 0&&_fd < bitsPerOutputSample {_ecf <<=uint (_fd );_fc =append (_fc ,_ecf );};return _fc ;};type Writer struct{_bf _a .ImageBase ;_dda *_cc .Writer ;_ee ,_afe int ;_cab bool ;};func (_ad *Reader )ReadSamples (samples []uint32 )(_ec error ){for _fb :=0;
_fb < len (samples );_fb ++{samples [_fb ],_ec =_ad .ReadSample ();if _ec !=nil {return _ec ;};};return nil ;};