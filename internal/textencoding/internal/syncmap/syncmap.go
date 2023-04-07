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

package syncmap ;import _b "sync";type RuneUint16Map struct{_fee map[rune ]uint16 ;_cgge _b .RWMutex ;};func (_fc *RuneSet )Length ()int {_fc ._ca .RLock ();defer _fc ._ca .RUnlock ();return len (_fc ._bb )};func (_ed *RuneByteMap )Write (r rune ,b byte ){_ed ._fe .Lock ();
defer _ed ._fe .Unlock ();_ed ._bc [r ]=b };type RuneStringMap struct{_dd map[rune ]string ;_cg _b .RWMutex ;};func MakeRuneUint16Map (length int )*RuneUint16Map {return &RuneUint16Map {_fee :make (map[rune ]uint16 ,length )};};func (_bg *RuneStringMap )Read (r rune )(string ,bool ){_bg ._cg .RLock ();
defer _bg ._cg .RUnlock ();_cgg ,_fdc :=_bg ._dd [r ];return _cgg ,_fdc ;};func (_df *RuneStringMap )Length ()int {_df ._cg .RLock ();defer _df ._cg .RUnlock ();return len (_df ._dd )};func NewStringRuneMap (m map[string ]rune )*StringRuneMap {return &StringRuneMap {_eda :m }};
func NewByteRuneMap (m map[byte ]rune )*ByteRuneMap {return &ByteRuneMap {_d :m }};func (_fac *StringsMap )Read (g string )(string ,bool ){_fac ._ccd .RLock ();defer _fac ._ccd .RUnlock ();_bdf ,_dac :=_fac ._afd [g ];return _bdf ,_dac ;};func (_bed *RuneUint16Map )Write (r rune ,g uint16 ){_bed ._cgge .Lock ();
defer _bed ._cgge .Unlock ();_bed ._fee [r ]=g ;};func (_a *ByteRuneMap )Read (b byte )(rune ,bool ){_a ._f .RLock ();defer _a ._f .RUnlock ();_fb ,_af :=_a ._d [b ];return _fb ,_af ;};type RuneSet struct{_bb map[rune ]struct{};_ca _b .RWMutex ;};func (_eeb *StringRuneMap )Read (g string )(rune ,bool ){_eeb ._ccgf .RLock ();
defer _eeb ._ccgf .RUnlock ();_acg ,_egcd :=_eeb ._eda [g ];return _acg ,_egcd ;};func (_ab *ByteRuneMap )Length ()int {_ab ._f .RLock ();defer _ab ._f .RUnlock ();return len (_ab ._d )};func (_e *ByteRuneMap )Range (f func (_ac byte ,_dc rune )(_eg bool )){_e ._f .RLock ();
defer _e ._f .RUnlock ();for _ea ,_aa :=range _e ._d {if f (_ea ,_aa ){break ;};};};func (_g *RuneByteMap )Read (r rune )(byte ,bool ){_g ._fe .RLock ();defer _g ._fe .RUnlock ();_fba ,_ead :=_g ._bc [r ];return _fba ,_ead ;};func MakeByteRuneMap (length int )*ByteRuneMap {return &ByteRuneMap {_d :make (map[byte ]rune ,length )}};
func (_ebd *StringsMap )Range (f func (_gf ,_bgg string )(_bae bool )){_ebd ._ccd .RLock ();defer _ebd ._ccd .RUnlock ();for _gde ,_bea :=range _ebd ._afd {if f (_gde ,_bea ){break ;};};};type StringsMap struct{_afd map[string ]string ;_ccd _b .RWMutex ;
};type StringsTuple struct{Key ,Value string ;};func (_aad *RuneByteMap )Range (f func (_be rune ,_da byte )(_cf bool )){_aad ._fe .RLock ();defer _aad ._fe .RUnlock ();for _ae ,_fd :=range _aad ._bc {if f (_ae ,_fd ){break ;};};};func (_agb *StringRuneMap )Write (g string ,r rune ){_agb ._ccgf .Lock ();
defer _agb ._ccgf .Unlock ();_agb ._eda [g ]=r ;};func (_ec *RuneSet )Write (r rune ){_ec ._ca .Lock ();defer _ec ._ca .Unlock ();_ec ._bb [r ]=struct{}{}};func (_cb *ByteRuneMap )Write (b byte ,r rune ){_cb ._f .Lock ();defer _cb ._f .Unlock ();_cb ._d [b ]=r };
func (_fbag *RuneUint16Map )Length ()int {_fbag ._cgge .RLock ();defer _fbag ._cgge .RUnlock ();return len (_fbag ._fee );};func (_aga *StringRuneMap )Range (f func (_dgd string ,_aab rune )(_gea bool )){_aga ._ccgf .RLock ();defer _aga ._ccgf .RUnlock ();
for _fcb ,_dcb :=range _aga ._eda {if f (_fcb ,_dcb ){break ;};};};func MakeRuneSet (length int )*RuneSet {return &RuneSet {_bb :make (map[rune ]struct{},length )}};func (_egc *RuneUint16Map )Read (r rune )(uint16 ,bool ){_egc ._cgge .RLock ();defer _egc ._cgge .RUnlock ();
_ccg ,_bd :=_egc ._fee [r ];return _ccg ,_bd ;};func MakeRuneByteMap (length int )*RuneByteMap {_bf :=make (map[rune ]byte ,length );return &RuneByteMap {_bc :_bf };};func NewRuneStringMap (m map[rune ]string )*RuneStringMap {return &RuneStringMap {_dd :m }};
func (_bcd *StringRuneMap )Length ()int {_bcd ._ccgf .RLock ();defer _bcd ._ccgf .RUnlock ();return len (_bcd ._eda );};type ByteRuneMap struct{_d map[byte ]rune ;_f _b .RWMutex ;};type StringRuneMap struct{_eda map[string ]rune ;_ccgf _b .RWMutex ;};func (_ee *RuneStringMap )Range (f func (_fa rune ,_ccc string )(_ge bool )){_ee ._cg .RLock ();
defer _ee ._cg .RUnlock ();for _acf ,_fag :=range _ee ._dd {if f (_acf ,_fag ){break ;};};};func (_afc *RuneUint16Map )Delete (r rune ){_afc ._cgge .Lock ();defer _afc ._cgge .Unlock ();delete (_afc ._fee ,r );};func (_gd *RuneSet )Range (f func (_gcc rune )(_cc bool )){_gd ._ca .RLock ();
defer _gd ._ca .RUnlock ();for _cd :=range _gd ._bb {if f (_cd ){break ;};};};func (_de *RuneUint16Map )RangeDelete (f func (_gee rune ,_faa uint16 )(_ff bool ,_bfc bool )){_de ._cgge .Lock ();defer _de ._cgge .Unlock ();for _ef ,_cba :=range _de ._fee {_fdb ,_fcc :=f (_ef ,_cba );
if _fdb {delete (_de ._fee ,_ef );};if _fcc {break ;};};};func (_ag *RuneSet )Exists (r rune )bool {_ag ._ca .RLock ();defer _ag ._ca .RUnlock ();_ ,_fbb :=_ag ._bb [r ];return _fbb ;};type RuneByteMap struct{_bc map[rune ]byte ;_fe _b .RWMutex ;};func NewStringsMap (tuples []StringsTuple )*StringsMap {_dgg :=map[string ]string {};
for _ ,_dgbb :=range tuples {_dgg [_dgbb .Key ]=_dgbb .Value ;};return &StringsMap {_afd :_dgg };};func (_dg *RuneByteMap )Length ()int {_dg ._fe .RLock ();defer _dg ._fe .RUnlock ();return len (_dg ._bc )};func (_ba *StringsMap )Write (g1 ,g2 string ){_ba ._ccd .Lock ();
defer _ba ._ccd .Unlock ();_ba ._afd [g1 ]=g2 ;};func (_bfg *RuneUint16Map )Range (f func (_eb rune ,_dde uint16 )(_ecb bool )){_bfg ._cgge .RLock ();defer _bfg ._cgge .RUnlock ();for _ga ,_eed :=range _bfg ._fee {if f (_ga ,_eed ){break ;};};};func (_ecg *StringsMap )Copy ()*StringsMap {_ecg ._ccd .RLock ();
defer _ecg ._ccd .RUnlock ();_edb :=map[string ]string {};for _fbc ,_eee :=range _ecg ._afd {_edb [_fbc ]=_eee ;};return &StringsMap {_afd :_edb };};func (_dgb *RuneStringMap )Write (r rune ,s string ){_dgb ._cg .Lock ();defer _dgb ._cg .Unlock ();_dgb ._dd [r ]=s ;
};