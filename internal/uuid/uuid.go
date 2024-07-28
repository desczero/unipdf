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

package uuid ;import (_b "crypto/rand";_ba "encoding/hex";_bb "io";);func _dc (_ga []byte ,_cf UUID ){_ba .Encode (_ga ,_cf [:4]);_ga [8]='-';_ba .Encode (_ga [9:13],_cf [4:6]);_ga [13]='-';_ba .Encode (_ga [14:18],_cf [6:8]);_ga [18]='-';_ba .Encode (_ga [19:23],_cf [8:10]);
_ga [23]='-';_ba .Encode (_ga [24:],_cf [10:]);};func (_bf UUID )String ()string {var _d [36]byte ;_dc (_d [:],_bf );return string (_d [:])};type UUID [16]byte ;func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_c :=_bb .ReadFull (_e ,uuid [:]);if _c !=nil {return _gc ,_c ;
};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;return uuid ,nil ;};var _e =_b .Reader ;var Nil =_gc ;var _gc UUID ;func MustUUID ()UUID {uuid ,_bg :=NewUUID ();if _bg !=nil {panic (_bg );};return uuid ;};