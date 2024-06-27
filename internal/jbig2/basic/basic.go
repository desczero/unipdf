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

package basic ;import _db "github.com/unidoc/unipdf/v3/internal/jbig2/errors";func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};func (_b IntsMap )Add (key uint64 ,value int ){_b [key ]=append (_b [key ],value )};
func Abs (v int )int {if v > 0{return v ;};return -v ;};func (_bg IntsMap )Get (key uint64 )(int ,bool ){_a ,_ab :=_bg [key ];if !_ab {return 0,false ;};if len (_a )==0{return 0,false ;};return _a [0],true ;};func (_ba *Stack )peek ()(interface{},bool ){_ccf :=_ba .top ();
if _ccf ==-1{return nil ,false ;};return _ba .Data [_ccf ],true ;};func (_gf *Stack )Push (v interface{}){_gf .Data =append (_gf .Data ,v )};func (_c *IntSlice )Copy ()*IntSlice {_fa :=IntSlice (make ([]int ,len (*_c )));copy (_fa ,*_c );return &_fa ;};
func (_ee *Stack )Peek ()(_dc interface{},_bge bool ){return _ee .peek ()};func Min (x ,y int )int {if x < y {return x ;};return y ;};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};type NumSlice []float32 ;func (_cf *NumSlice )AddInt (v int ){*_cf =append (*_cf ,float32 (v ))};
func Max (x ,y int )int {if x > y {return x ;};return y ;};type IntsMap map[uint64 ][]int ;type Stack struct{Data []interface{};Aux *Stack ;};func (_bga IntSlice )Size ()int {return len (_bga )};func (_fdb NumSlice )GetIntSlice ()[]int {_bgf :=make ([]int ,len (_fdb ));
for _ef ,_dg :=range _fdb {_bgf [_ef ]=int (_dg );};return _bgf ;};func (_fe *Stack )Pop ()(_ge interface{},_ea bool ){_ge ,_ea =_fe .peek ();if !_ea {return nil ,_ea ;};_fe .Data =_fe .Data [:_fe .top ()];return _ge ,true ;};func (_ad IntsMap )Delete (key uint64 ){delete (_ad ,key )};
func (_cc *NumSlice )Add (v float32 ){*_cc =append (*_cc ,v )};func NewIntSlice (i int )*IntSlice {_dbc :=IntSlice (make ([]int ,i ));return &_dbc };func (_af *Stack )Len ()int {return len (_af .Data )};func (_ce IntSlice )Get (index int )(int ,error ){if index > len (_ce )-1{return 0,_db .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );
};return _ce [index ],nil ;};func (_ae *IntSlice )Add (v int )error {if _ae ==nil {return _db .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");
};*_ae =append (*_ae ,v );return nil ;};func (_fd NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_fd )-1{return 0,_db .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};return _fd [i ],nil ;};func NewNumSlice (i int )*NumSlice {_da :=NumSlice (make ([]float32 ,i ));return &_da };type IntSlice []int ;func (_aef *Stack )top ()int {return len (_aef .Data )-1};func (_e NumSlice )GetInt (i int )(int ,error ){const _ec ="\u0047\u0065\u0074\u0049\u006e\u0074";
if i < 0||i > len (_e )-1{return 0,_db .Errorf (_ec ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};_bc :=_e [i ];return int (_bc +Sign (_bc )*0.5),nil ;};func (_g IntsMap )GetSlice (key uint64 )([]int ,bool ){_gd ,_f :=_g [key ];
if !_f {return nil ,false ;};return _gd ,true ;};