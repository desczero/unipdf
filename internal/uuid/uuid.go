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

package uuid ;import (_e "crypto/rand";_ec "encoding/hex";_g "io";);func MustUUID ()UUID {uuid ,_fc :=NewUUID ();if _fc !=nil {panic (_fc );};return uuid ;};var _aa UUID ;func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_ga :=_g .ReadFull (_d ,uuid [:]);
if _ga !=nil {return _aa ,_ga ;};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;return uuid ,nil ;};func (_da UUID )String ()string {var _ee [36]byte ;_gc (_ee [:],_da );return string (_ee [:])};func _gc (_fe []byte ,_ac UUID ){_ec .Encode (_fe ,_ac [:4]);
_fe [8]='-';_ec .Encode (_fe [9:13],_ac [4:6]);_fe [13]='-';_ec .Encode (_fe [14:18],_ac [6:8]);_fe [18]='-';_ec .Encode (_fe [19:23],_ac [8:10]);_fe [23]='-';_ec .Encode (_fe [24:],_ac [10:]);};type UUID [16]byte ;var Nil =_aa ;var _d =_e .Reader ;