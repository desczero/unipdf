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

package huffman ;import (_df "errors";_b "fmt";_d "github.com/unidoc/unipdf/v3/internal/bitwise";_de "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_fc "math";_f "strings";);func (_ba *EncodedTable )Decode (r _d .StreamReader )(int64 ,error ){return _ba ._c .Decode (r )};
func (_ga *FixedSizeTable )Decode (r _d .StreamReader )(int64 ,error ){return _ga ._ab .Decode (r )};func (_adg *InternalNode )Decode (r _d .StreamReader )(int64 ,error ){_fcd ,_fee :=r .ReadBit ();if _fee !=nil {return 0,_fee ;};if _fcd ==0{return _adg ._abg .Decode (r );
};return _adg ._ag .Decode (r );};func (_ed *ValueNode )String ()string {return _b .Sprintf ("\u0025\u0064\u002f%\u0064",_ed ._bff ,_ed ._bgab );};func GetStandardTable (number int )(Tabler ,error ){if number <=0||number > len (_agf ){return nil ,_df .New ("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_adda :=_agf [number -1];if _adda ==nil {var _bcb error ;_adda ,_bcb =_bde (_dgde [number -1]);if _bcb !=nil {return nil ,_bcb ;};_agf [number -1]=_adda ;};return _adda ,nil ;};func (_edf *StandardTable )InitTree (codeTable []*Code )error {_ccc (codeTable );
for _ ,_ffg :=range codeTable {if _bac :=_edf ._gbc .append (_ffg );_bac !=nil {return _bac ;};};return nil ;};func (_cfc *OutOfBandNode )String ()string {return _b .Sprintf ("\u0025\u0030\u00364\u0062",int64 (_fc .MaxInt64 ));};func (_cde *OutOfBandNode )Decode (r _d .StreamReader )(int64 ,error ){return 0,_de .ErrOOB };
func (_bf *EncodedTable )String ()string {return _bf ._c .String ()+"\u000a"};func _ec (_afb int32 )*InternalNode {return &InternalNode {_gd :_afb }};func (_af *ValueNode )Decode (r _d .StreamReader )(int64 ,error ){_cdd ,_fb :=r .ReadBits (byte (_af ._bff ));
if _fb !=nil {return 0,_fb ;};if _af ._add {_cdd =-_cdd ;};return int64 (_af ._bgab )+int64 (_cdd ),nil ;};var _ Node =&ValueNode {};type InternalNode struct{_gd int32 ;_abg Node ;_ag Node ;};var _ Tabler =&EncodedTable {};var _ Node =&OutOfBandNode {};
type FixedSizeTable struct{_ab *InternalNode };var _agf =make ([]Tabler ,len (_dgde ));func (_ff *FixedSizeTable )InitTree (codeTable []*Code )error {_ccc (codeTable );for _ ,_bga :=range codeTable {_ef :=_ff ._ab .append (_bga );if _ef !=nil {return _ef ;
};};return nil ;};type EncodedTable struct{BasicTabler ;_c *InternalNode ;};func (_faf *InternalNode )append (_ac *Code )(_dgd error ){if _ac ._gcd ==0{return nil ;};_gbb :=_ac ._gcd -1-_faf ._gd ;if _gbb < 0{return _df .New ("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064");
};_gc :=(_ac ._cfb >>uint (_gbb ))&0x1;if _gbb ==0{if _ac ._beff ==-1{if _gc ==1{if _faf ._ag !=nil {return _b .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_ac );
};_faf ._ag =_cfd (_ac );}else {if _faf ._abg !=nil {return _b .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_ac );};_faf ._abg =_cfd (_ac );};}else {if _gc ==1{if _faf ._ag !=nil {return _b .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_ac );
};_faf ._ag =_aa (_ac );}else {if _faf ._abg !=nil {return _b .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_ac );
};_faf ._abg =_aa (_ac );};};}else {if _gc ==1{if _faf ._ag ==nil {_faf ._ag =_ec (_faf ._gd +1);};if _dgd =_faf ._ag .(*InternalNode ).append (_ac );_dgd !=nil {return _dgd ;};}else {if _faf ._abg ==nil {_faf ._abg =_ec (_faf ._gd +1);};if _dgd =_faf ._abg .(*InternalNode ).append (_ac );
_dgd !=nil {return _dgd ;};};};return nil ;};func (_a *EncodedTable )InitTree (codeTable []*Code )error {_ccc (codeTable );for _ ,_ad :=range codeTable {if _dg :=_a ._c .append (_ad );_dg !=nil {return _dg ;};};return nil ;};func (_cf *EncodedTable )parseTable ()error {var (_cdg []*Code ;
_fab ,_ce ,_ca int32 ;_bd uint64 ;_bg error ;);_ae :=_cf .StreamReader ();_fe :=_cf .HtLow ();for _fe < _cf .HtHigh (){_bd ,_bg =_ae .ReadBits (byte (_cf .HtPS ()));if _bg !=nil {return _bg ;};_fab =int32 (_bd );_bd ,_bg =_ae .ReadBits (byte (_cf .HtRS ()));
if _bg !=nil {return _bg ;};_ce =int32 (_bd );_cdg =append (_cdg ,NewCode (_fab ,_ce ,_ca ,false ));_fe +=1<<uint (_ce );};_bd ,_bg =_ae .ReadBits (byte (_cf .HtPS ()));if _bg !=nil {return _bg ;};_fab =int32 (_bd );_ce =32;_ca =_cf .HtLow ()-1;_cdg =append (_cdg ,NewCode (_fab ,_ce ,_ca ,true ));
_bd ,_bg =_ae .ReadBits (byte (_cf .HtPS ()));if _bg !=nil {return _bg ;};_fab =int32 (_bd );_ce =32;_ca =_cf .HtHigh ();_cdg =append (_cdg ,NewCode (_fab ,_ce ,_ca ,false ));if _cf .HtOOB ()==1{_bd ,_bg =_ae .ReadBits (byte (_cf .HtPS ()));if _bg !=nil {return _bg ;
};_fab =int32 (_bd );_cdg =append (_cdg ,NewCode (_fab ,-1,-1,false ));};if _bg =_cf .InitTree (_cdg );_bg !=nil {return _bg ;};return nil ;};func (_eb *StandardTable )String ()string {return _eb ._gbc .String ()+"\u000a"};var _ Node =&InternalNode {};
type ValueNode struct{_bff int32 ;_bgab int32 ;_add bool ;};func (_fba *Code )String ()string {var _eg string ;if _fba ._cfb !=-1{_eg =_eca (_fba ._cfb ,_fba ._gcd );}else {_eg ="\u003f";};return _b .Sprintf ("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_eg ,_fba ._gcd ,_fba ._beff ,_fba ._efd );
};type StandardTable struct{_gbc *InternalNode };func NewFixedSizeTable (codeTable []*Code )(*FixedSizeTable ,error ){_bc :=&FixedSizeTable {_ab :&InternalNode {}};if _gb :=_bc .InitTree (codeTable );_gb !=nil {return nil ,_gb ;};return _bc ,nil ;};func (_e *EncodedTable )RootNode ()*InternalNode {return _e ._c };
type Tabler interface{Decode (_dd _d .StreamReader )(int64 ,error );InitTree (_dgg []*Code )error ;String ()string ;RootNode ()*InternalNode ;};func NewEncodedTable (table BasicTabler )(*EncodedTable ,error ){_gf :=&EncodedTable {_c :&InternalNode {},BasicTabler :table };
if _cd :=_gf .parseTable ();_cd !=nil {return nil ,_cd ;};return _gf ,nil ;};func _eca (_cc ,_ea int32 )string {var _aba int32 ;_fef :=make ([]rune ,_ea );for _gef :=int32 (1);_gef <=_ea ;_gef ++{_aba =_cc >>uint (_ea -_gef )&1;if _aba !=0{_fef [_gef -1]='1';
}else {_fef [_gef -1]='0';};};return string (_fef );};func _cfd (_ge *Code )*OutOfBandNode {return &OutOfBandNode {}};func (_bce *InternalNode )String ()string {_dc :=&_f .Builder {};_dc .WriteString ("\u000a");_bce .pad (_dc );_dc .WriteString ("\u0030\u003a\u0020");
_dc .WriteString (_bce ._abg .String ()+"\u000a");_bce .pad (_dc );_dc .WriteString ("\u0031\u003a\u0020");_dc .WriteString (_bce ._ag .String ()+"\u000a");return _dc .String ();};func _gge (_ccd ,_egg int32 )int32 {if _ccd > _egg {return _ccd ;};return _egg ;
};func (_fag *StandardTable )RootNode ()*InternalNode {return _fag ._gbc };func NewCode (prefixLength ,rangeLength ,rangeLow int32 ,isLowerRange bool )*Code {return &Code {_gcd :prefixLength ,_beff :rangeLength ,_efd :rangeLow ,_fd :isLowerRange ,_cfb :-1};
};func (_bef *FixedSizeTable )String ()string {return _bef ._ab .String ()+"\u000a"};var _dgde =[][][]int32 {{{1,4,0},{2,8,16},{3,16,272},{3,32,65808}},{{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{6,32,75},{6,-1,0}},{{8,8,-256},{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{8,32,-257,999},{7,32,75},{6,-1,0}},{{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{5,32,76}},{{7,8,-255},{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{7,32,-256,999},{6,32,76}},{{5,10,-2048},{4,9,-1024},{4,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{2,7,0},{3,7,128},{3,8,256},{4,9,512},{4,10,1024},{6,32,-2049,999},{6,32,2048}},{{4,9,-1024},{3,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{4,5,0},{5,5,32},{5,6,64},{4,7,128},{3,8,256},{3,9,512},{3,10,1024},{5,32,-1025,999},{5,32,2048}},{{8,3,-15},{9,1,-7},{8,1,-5},{9,0,-3},{7,0,-2},{4,0,-1},{2,1,0},{5,0,2},{6,0,3},{3,4,4},{6,1,20},{4,4,22},{4,5,38},{5,6,70},{5,7,134},{6,7,262},{7,8,390},{6,10,646},{9,32,-16,999},{9,32,1670},{2,-1,0}},{{8,4,-31},{9,2,-15},{8,2,-11},{9,1,-7},{7,1,-5},{4,1,-3},{3,1,-1},{3,1,1},{5,1,3},{6,1,5},{3,5,7},{6,2,39},{4,5,43},{4,6,75},{5,7,139},{5,8,267},{6,8,523},{7,9,779},{6,11,1291},{9,32,-32,999},{9,32,3339},{2,-1,0}},{{7,4,-21},{8,0,-5},{7,0,-4},{5,0,-3},{2,2,-2},{5,0,2},{6,0,3},{7,0,4},{8,0,5},{2,6,6},{5,5,70},{6,5,102},{6,6,134},{6,7,198},{6,8,326},{6,9,582},{6,10,1094},{7,11,2118},{8,32,-22,999},{8,32,4166},{2,-1,0}},{{1,0,1},{2,1,2},{4,0,4},{4,1,5},{5,1,7},{5,2,9},{6,2,13},{7,2,17},{7,3,21},{7,4,29},{7,5,45},{7,6,77},{7,32,141}},{{1,0,1},{2,0,2},{3,1,3},{5,0,5},{5,1,6},{6,1,8},{7,0,10},{7,1,11},{7,2,13},{7,3,17},{7,4,25},{8,5,41},{8,32,73}},{{1,0,1},{3,0,2},{4,0,3},{5,0,4},{4,1,5},{3,3,7},{6,1,15},{6,2,17},{6,3,21},{6,4,29},{6,5,45},{7,6,77},{7,32,141}},{{3,0,-2},{3,0,-1},{1,0,0},{3,0,1},{3,0,2}},{{7,4,-24},{6,2,-8},{5,1,-4},{4,0,-2},{3,0,-1},{1,0,0},{3,0,1},{4,0,2},{5,1,3},{6,2,5},{7,4,9},{7,32,-25,999},{7,32,25}}};
type OutOfBandNode struct{};func (_cef *FixedSizeTable )RootNode ()*InternalNode {return _cef ._ab };func _ccc (_afd []*Code ){var _aad int32 ;for _ ,_gga :=range _afd {_aad =_gge (_aad ,_gga ._gcd );};_eaf :=make ([]int32 ,_aad +1);for _ ,_cefg :=range _afd {_eaf [_cefg ._gcd ]++;
};var _dbf int32 ;_abd :=make ([]int32 ,len (_eaf )+1);_eaf [0]=0;for _egb :=int32 (1);_egb <=int32 (len (_eaf ));_egb ++{_abd [_egb ]=(_abd [_egb -1]+(_eaf [_egb -1]))<<1;_dbf =_abd [_egb ];for _ ,_fcb :=range _afd {if _fcb ._gcd ==_egb {_fcb ._cfb =_dbf ;
_dbf ++;};};};};type Node interface{Decode (_db _d .StreamReader )(int64 ,error );String ()string ;};func _aa (_cac *Code )*ValueNode {return &ValueNode {_bff :_cac ._beff ,_bgab :_cac ._efd ,_add :_cac ._fd };};func (_bge *InternalNode )pad (_fg *_f .Builder ){for _faa :=int32 (0);
_faa < _bge ._gd ;_faa ++{_fg .WriteString ("\u0020\u0020\u0020");};};func _bde (_cg [][]int32 )(*StandardTable ,error ){var _fafa []*Code ;for _gg :=0;_gg < len (_cg );_gg ++{_bffa :=_cg [_gg ][0];_fca :=_cg [_gg ][1];_aff :=_cg [_gg ][2];var _eff bool ;
if len (_cg [_gg ])> 3{_eff =true ;};_fafa =append (_fafa ,NewCode (_bffa ,_fca ,_aff ,_eff ));};_fgc :=&StandardTable {_gbc :_ec (0)};if _ee :=_fgc .InitTree (_fafa );_ee !=nil {return nil ,_ee ;};return _fgc ,nil ;};type BasicTabler interface{HtHigh ()int32 ;
HtLow ()int32 ;StreamReader ()_d .StreamReader ;HtPS ()int32 ;HtRS ()int32 ;HtOOB ()int32 ;};func (_fge *StandardTable )Decode (r _d .StreamReader )(int64 ,error ){return _fge ._gbc .Decode (r )};type Code struct{_gcd int32 ;_beff int32 ;_efd int32 ;_fd bool ;
_cfb int32 ;};