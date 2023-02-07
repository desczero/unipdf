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

package uuid ;import (_d "crypto/rand";_e "encoding/hex";_c "io";);func MustUUID ()UUID {uuid ,_geg :=NewUUID ();if _geg !=nil {panic (_geg );};return uuid ;};var _gg =_d .Reader ;var Nil =_da ;func _egf (_b []byte ,_cf UUID ){_e .Encode (_b ,_cf [:4]);
_b [8]='-';_e .Encode (_b [9:13],_cf [4:6]);_b [13]='-';_e .Encode (_b [14:18],_cf [6:8]);_b [18]='-';_e .Encode (_b [19:23],_cf [8:10]);_b [23]='-';_e .Encode (_b [24:],_cf [10:]);};var _da UUID ;func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_ge :=_c .ReadFull (_gg ,uuid [:]);
if _ge !=nil {return _da ,_ge ;};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;return uuid ,nil ;};type UUID [16]byte ;func (_eef UUID )String ()string {var _a [36]byte ;_egf (_a [:],_eef );return string (_a [:])};