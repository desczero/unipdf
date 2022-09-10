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

// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
//
// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
package ps ;import (_ga "bufio";_eg "bytes";_e "errors";_gc "fmt";_b "github.com/unidoc/unipdf/v3/common";_f "github.com/unidoc/unipdf/v3/core";_eb "io";_d "math";);func (_fdc *PSOperand )ln (_def *PSStack )error {_dgbc ,_gfgd :=_def .PopNumberAsFloat64 ();
if _gfgd !=nil {return _gfgd ;};_eda :=_d .Log (_dgbc );_gfgd =_def .Push (MakeReal (_eda ));return _gfgd ;};

// Parse parses the postscript and store as a program that can be executed.
func (_bce *PSParser )Parse ()(*PSProgram ,error ){_bce .skipSpaces ();_gaac ,_fcff :=_bce ._bdaf .Peek (2);if _fcff !=nil {return nil ,_fcff ;};if _gaac [0]!='{'{return nil ,_e .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b");
};_acaf ,_fcff :=_bce .parseFunction ();if _fcff !=nil &&_fcff !=_eb .EOF {return nil ,_fcff ;};return _acaf ,_fcff ;};func (_ef *PSProgram )Duplicate ()PSObject {_faa :=&PSProgram {};for _ ,_gfa :=range *_ef {_faa .Append (_gfa .Duplicate ());};return _faa ;
};

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool (val bool )*PSBoolean {_ecfb :=PSBoolean {};_ecfb .Val =val ;return &_ecfb };var ErrUnsupportedOperand =_e .New ("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064");func (_gaca *PSOperand )roll (_eaf *PSStack )error {_dacg ,_efe :=_eaf .Pop ();
if _efe !=nil {return _efe ;};_fcee ,_efe :=_eaf .Pop ();if _efe !=nil {return _efe ;};_gebg ,_bgd :=_dacg .(*PSInteger );if !_bgd {return ErrTypeCheck ;};_dbcg ,_bgd :=_fcee .(*PSInteger );if !_bgd {return ErrTypeCheck ;};if _dbcg .Val < 0{return ErrRangeCheck ;
};if _dbcg .Val ==0||_dbcg .Val ==1{return nil ;};if _dbcg .Val > len (*_eaf ){return ErrStackUnderflow ;};for _fad :=0;_fad < _fcdd (_gebg .Val );_fad ++{var _adbca []PSObject ;_adbca =(*_eaf )[len (*_eaf )-(_dbcg .Val ):len (*_eaf )];if _gebg .Val > 0{_gfgg :=_adbca [len (_adbca )-1];
_adbca =append ([]PSObject {_gfgg },_adbca [0:len (_adbca )-1]...);}else {_fdbg :=_adbca [len (_adbca )-_dbcg .Val ];_adbca =append (_adbca [1:],_fdbg );};_cbfg :=append ((*_eaf )[0:len (*_eaf )-_dbcg .Val ],_adbca ...);_eaf =&_cbfg ;};return nil ;};

// PSBoolean represents a boolean value.
type PSBoolean struct{Val bool ;};func (_fge *PSOperand )le (_bdff *PSStack )error {_bb ,_aadc :=_bdff .PopNumberAsFloat64 ();if _aadc !=nil {return _aadc ;};_fef ,_aadc :=_bdff .PopNumberAsFloat64 ();if _aadc !=nil {return _aadc ;};if _d .Abs (_fef -_bb )< _a {_eegc :=_bdff .Push (MakeBool (true ));
return _eegc ;}else if _fef < _bb {_gae :=_bdff .Push (MakeBool (true ));return _gae ;}else {_gab :=_bdff .Push (MakeBool (false ));return _gab ;};};func (_fbe *PSOperand )ceiling (_afbb *PSStack )error {_ca ,_dcg :=_afbb .Pop ();if _dcg !=nil {return _dcg ;
};if _ffg ,_cac :=_ca .(*PSReal );_cac {_dcg =_afbb .Push (MakeReal (_d .Ceil (_ffg .Val )));}else if _bdf ,_cc :=_ca .(*PSInteger );_cc {_dcg =_afbb .Push (MakeInteger (_bdf .Val ));}else {_dcg =ErrTypeCheck ;};return _dcg ;};

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal (val float64 )*PSReal {_acgc :=PSReal {};_acgc .Val =val ;return &_acgc };func (_da *PSOperand )abs (_cd *PSStack )error {_fag ,_ed :=_cd .Pop ();if _ed !=nil {return _ed ;};if _ae ,_cg :=_fag .(*PSReal );_cg {_ac :=_ae .Val ;if _ac < 0{_ed =_cd .Push (MakeReal (-_ac ));
}else {_ed =_cd .Push (MakeReal (_ac ));};}else if _fea ,_adf :=_fag .(*PSInteger );_adf {_bge :=_fea .Val ;if _bge < 0{_ed =_cd .Push (MakeInteger (-_bge ));}else {_ed =_cd .Push (MakeInteger (_bge ));};}else {return ErrTypeCheck ;};return _ed ;};func (_eeb *PSProgram )String ()string {_cbe :="\u007b\u0020";
for _ ,_fgga :=range *_eeb {_cbe +=_fgga .String ();_cbe +="\u0020";};_cbe +="\u007d";return _cbe ;};

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct{Stack *PSStack ;_fa *PSProgram ;};

// Append appends an object to the PSProgram.
func (_dce *PSProgram )Append (obj PSObject ){*_dce =append (*_dce ,obj )};

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand (val string )*PSOperand {_dda :=PSOperand (val );return &_dda };func (_gdd *PSOperand )ge (_ffe *PSStack )error {_fecg ,_adc :=_ffe .PopNumberAsFloat64 ();if _adc !=nil {return _adc ;};_ffgc ,_adc :=_ffe .PopNumberAsFloat64 ();if _adc !=nil {return _adc ;
};if _d .Abs (_ffgc -_fecg )< _a {_efc :=_ffe .Push (MakeBool (true ));return _efc ;}else if _ffgc > _fecg {_feb :=_ffe .Push (MakeBool (true ));return _feb ;}else {_cad :=_ffe .Push (MakeBool (false ));return _cad ;};};func (_aa *PSReal )DebugString ()string {return _gc .Sprintf ("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f",_aa .Val );
};func (_bffg *PSOperand )truncate (_ccf *PSStack )error {_aaga ,_cae :=_ccf .Pop ();if _cae !=nil {return _cae ;};if _ecdb ,_fece :=_aaga .(*PSReal );_fece {_bdg :=int (_ecdb .Val );_cae =_ccf .Push (MakeReal (float64 (_bdg )));}else if _ebgd ,_geg :=_aaga .(*PSInteger );
_geg {_cae =_ccf .Push (MakeInteger (_ebgd .Val ));}else {return ErrTypeCheck ;};return _cae ;};func (_gbcd *PSOperand )round (_fbcd *PSStack )error {_cce ,_fcbd :=_fbcd .Pop ();if _fcbd !=nil {return _fcbd ;};if _bbb ,_edg :=_cce .(*PSReal );_edg {_fcbd =_fbcd .Push (MakeReal (_d .Floor (_bbb .Val +0.5)));
}else if _dgg ,_eddf :=_cce .(*PSInteger );_eddf {_fcbd =_fbcd .Push (MakeInteger (_dgg .Val ));}else {return ErrTypeCheck ;};return _fcbd ;};func (_caag *PSOperand )sub (_bed *PSStack )error {_efbb ,_aae :=_bed .Pop ();if _aae !=nil {return _aae ;};_fbf ,_aae :=_bed .Pop ();
if _aae !=nil {return _aae ;};_eba ,_bdde :=_efbb .(*PSReal );_gagc ,_dae :=_efbb .(*PSInteger );if !_bdde &&!_dae {return ErrTypeCheck ;};_cadb ,_faf :=_fbf .(*PSReal );_fefa ,_dgf :=_fbf .(*PSInteger );if !_faf &&!_dgf {return ErrTypeCheck ;};if _dae &&_dgf {_aacc :=_fefa .Val -_gagc .Val ;
_ebcc :=_bed .Push (MakeInteger (_aacc ));return _ebcc ;};var _gace float64 =0;if _faf {_gace =_cadb .Val ;}else {_gace =float64 (_fefa .Val );};if _bdde {_gace -=_eba .Val ;}else {_gace -=float64 (_gagc .Val );};_aae =_bed .Push (MakeReal (_gace ));return _aae ;
};

// PSParser is a basic Postscript parser.
type PSParser struct{_bdaf *_ga .Reader };func (_bda *PSOperand )gt (_caa *PSStack )error {_beba ,_gbge :=_caa .PopNumberAsFloat64 ();if _gbge !=nil {return _gbge ;};_ecfd ,_gbge :=_caa .PopNumberAsFloat64 ();if _gbge !=nil {return _gbge ;};if _d .Abs (_ecfd -_beba )< _a {_agg :=_caa .Push (MakeBool (false ));
return _agg ;}else if _ecfd > _beba {_cfa :=_caa .Push (MakeBool (true ));return _cfa ;}else {_fgaa :=_caa .Push (MakeBool (false ));return _fgaa ;};};func (_gcc *PSBoolean )Duplicate ()PSObject {_be :=PSBoolean {};_be .Val =_gcc .Val ;return &_be };

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_dc *PSExecutor )Execute (objects []PSObject )([]PSObject ,error ){for _ ,_ff :=range objects {_de :=_dc .Stack .Push (_ff );if _de !=nil {return nil ,_de ;};};_fgg :=_dc ._fa .Exec (_dc .Stack );if _fgg !=nil {_b .Log .Debug ("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_fgg );
return nil ,_fgg ;};_gba :=[]PSObject (*_dc .Stack );_dc .Stack .Empty ();return _gba ,nil ;};func (_ee *PSInteger )Duplicate ()PSObject {_fe :=PSInteger {};_fe .Val =_ee .Val ;return &_fe };func (_ggg *PSOperand )log (_dffd *PSStack )error {_adg ,_add :=_dffd .PopNumberAsFloat64 ();
if _add !=nil {return _add ;};_cgf :=_d .Log10 (_adg );_add =_dffd .Push (MakeReal (_cgf ));return _add ;};func (_dac *PSOperand )copy (_fd *PSStack )error {_aca ,_bfd :=_fd .PopInteger ();if _bfd !=nil {return _bfd ;};if _aca < 0{return ErrRangeCheck ;
};if _aca > len (*_fd ){return ErrRangeCheck ;};*_fd =append (*_fd ,(*_fd )[len (*_fd )-_aca :]...);return nil ;};func (_eegg *PSOperand )String ()string {return string (*_eegg )};func (_caad *PSParser )skipSpaces ()(int ,error ){_efef :=0;for {_cdg ,_afbd :=_caad ._bdaf .Peek (1);
if _afbd !=nil {return 0,_afbd ;};if _f .IsWhiteSpace (_cdg [0]){_caad ._bdaf .ReadByte ();_efef ++;}else {break ;};};return _efef ,nil ;};func (_bfec *PSOperand )ifCondition (_cbed *PSStack )error {_cgea ,_aaac :=_cbed .Pop ();if _aaac !=nil {return _aaac ;
};_edc ,_aaac :=_cbed .Pop ();if _aaac !=nil {return _aaac ;};_eab ,_afbbc :=_cgea .(*PSProgram );if !_afbbc {return ErrTypeCheck ;};_cbf ,_afbbc :=_edc .(*PSBoolean );if !_afbbc {return ErrTypeCheck ;};if _cbf .Val {_adbc :=_eab .Exec (_cbed );return _adbc ;
};return nil ;};

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_bgda *PSStack )DebugString ()string {_ddb :="\u005b\u0020";for _ ,_faga :=range *_bgda {_ddb +=_faga .DebugString ();_ddb +="\u0020";};_ddb +="\u005d";return _ddb ;};func (_gagf *PSOperand )cos (_aab *PSStack )error {_fga ,_gbd :=_aab .PopNumberAsFloat64 ();
if _gbd !=nil {return _gbd ;};_ded :=_d .Cos (_fga *_d .Pi /180.0);_gbd =_aab .Push (MakeReal (_ded ));return _gbd ;};var ErrUndefinedResult =_e .New ("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072");
func (_gcaf *PSOperand )neg (_daga *PSStack )error {_cgfc ,_cec :=_daga .Pop ();if _cec !=nil {return _cec ;};if _ccg ,_egf :=_cgfc .(*PSReal );_egf {_cec =_daga .Push (MakeReal (-_ccg .Val ));return _cec ;}else if _baf ,_bbca :=_cgfc .(*PSInteger );_bbca {_cec =_daga .Push (MakeInteger (-_baf .Val ));
return _cec ;}else {return ErrTypeCheck ;};};

// Pop pops an object from the top of the stack.
func (_aacb *PSStack )Pop ()(PSObject ,error ){if len (*_aacb )< 1{return nil ,ErrStackUnderflow ;};_fcbe :=(*_aacb )[len (*_aacb )-1];*_aacb =(*_aacb )[0:len (*_aacb )-1];return _fcbe ,nil ;};func (_db *PSProgram )DebugString ()string {_eef :="\u007b\u0020";
for _ ,_dbf :=range *_db {_eef +=_dbf .DebugString ();_eef +="\u0020";};_eef +="\u007d";return _eef ;};

// PSOperand represents a Postscript operand (text string).
type PSOperand string ;func (_gcge *PSOperand )idiv (_dbg *PSStack )error {_cggg ,_fcf :=_dbg .Pop ();if _fcf !=nil {return _fcf ;};_gcgd ,_fcf :=_dbg .Pop ();if _fcf !=nil {return _fcf ;};_eefg ,_gdg :=_cggg .(*PSInteger );if !_gdg {return ErrTypeCheck ;
};if _eefg .Val ==0{return ErrUndefinedResult ;};_ffee ,_gdg :=_gcgd .(*PSInteger );if !_gdg {return ErrTypeCheck ;};_eadf :=_ffee .Val /_eefg .Val ;_fcf =_dbg .Push (MakeInteger (_eadf ));return _fcf ;};func (_fcb *PSOperand )pop (_fefe *PSStack )error {_ ,_edbc :=_fefe .Pop ();
if _edbc !=nil {return _edbc ;};return nil ;};func _fcdd (_edbb int )int {if _edbb < 0{return -_edbb ;};return _edbb ;};

// Empty empties the stack.
func (_abdc *PSStack )Empty (){*_abdc =[]PSObject {}};

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_bbde *PSStack )PopNumberAsFloat64 ()(float64 ,error ){_dcf ,_ebf :=_bbde .Pop ();if _ebf !=nil {return 0,_ebf ;};if _fff ,_bcd :=_dcf .(*PSReal );_bcd {return _fff .Val ,nil ;}else if _ddc ,_cage :=_dcf .(*PSInteger );_cage {return float64 (_ddc .Val ),nil ;
}else {return 0,ErrTypeCheck ;};};func (_af *PSReal )Duplicate ()PSObject {_gd :=PSReal {};_gd .Val =_af .Val ;return &_gd };func (_cge *PSOperand )bitshift (_gff *PSStack )error {_ace ,_ce :=_gff .PopInteger ();if _ce !=nil {return _ce ;};_gga ,_ce :=_gff .PopInteger ();
if _ce !=nil {return _ce ;};var _agb int ;if _ace >=0{_agb =_gga <<uint (_ace );}else {_agb =_gga >>uint (-_ace );};_ce =_gff .Push (MakeInteger (_agb ));return _ce ;};func (_gbe *PSOperand )lt (_bgb *PSStack )error {_eae ,_cab :=_bgb .PopNumberAsFloat64 ();
if _cab !=nil {return _cab ;};_acg ,_cab :=_bgb .PopNumberAsFloat64 ();if _cab !=nil {return _cab ;};if _d .Abs (_acg -_eae )< _a {_aef :=_bgb .Push (MakeBool (false ));return _aef ;}else if _acg < _eae {_efbg :=_bgb .Push (MakeBool (true ));return _efbg ;
}else {_bddf :=_bgb .Push (MakeBool (false ));return _bddf ;};};

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_ebcb *PSStack )PopInteger ()(int ,error ){_aee ,_bbdc :=_ebcb .Pop ();if _bbdc !=nil {return 0,_bbdc ;};if _bggd ,_gddb :=_aee .(*PSInteger );_gddb {return _bggd .Val ,nil ;};return 0,ErrTypeCheck ;};

// NewPSStack returns an initialized PSStack.
func NewPSStack ()*PSStack {return &PSStack {}};func (_adfc *PSOperand )cvr (_gcg *PSStack )error {_dcbd ,_afgb :=_gcg .Pop ();if _afgb !=nil {return _afgb ;};if _ead ,_ffb :=_dcbd .(*PSReal );_ffb {_afgb =_gcg .Push (MakeReal (_ead .Val ));}else if _gec ,_dcba :=_dcbd .(*PSInteger );
_dcba {_afgb =_gcg .Push (MakeReal (float64 (_gec .Val )));}else {return ErrTypeCheck ;};return _afgb ;};

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array (objects []PSObject )([]float64 ,error ){var _c []float64 ;for _ ,_fg :=range objects {if _ge ,_bca :=_fg .(*PSInteger );_bca {_c =append (_c ,float64 (_ge .Val ));}else if _gf ,_cb :=_fg .(*PSReal );_cb {_c =append (_c ,_gf .Val );
}else {return nil ,ErrTypeCheck ;};};return _c ,nil ;};func (_ggd *PSOperand )ifelse (_bdd *PSStack )error {_gcbg ,_eac :=_bdd .Pop ();if _eac !=nil {return _eac ;};_ebg ,_eac :=_bdd .Pop ();if _eac !=nil {return _eac ;};_acc ,_eac :=_bdd .Pop ();if _eac !=nil {return _eac ;
};_dgb ,_agab :=_gcbg .(*PSProgram );if !_agab {return ErrTypeCheck ;};_cbg ,_agab :=_ebg .(*PSProgram );if !_agab {return ErrTypeCheck ;};_abb ,_agab :=_acc .(*PSBoolean );if !_agab {return ErrTypeCheck ;};if _abb .Val {_cff :=_cbg .Exec (_bdd );return _cff ;
};_eac =_dgb .Exec (_bdd );return _eac ;};func (_acb *PSParser )parseNumber ()(PSObject ,error ){_feaf ,_edec :=_f .ParseNumber (_acb ._bdaf );if _edec !=nil {return nil ,_edec ;};switch _bbce :=_feaf .(type ){case *_f .PdfObjectFloat :return MakeReal (float64 (*_bbce )),nil ;
case *_f .PdfObjectInteger :return MakeInteger (int (*_bbce )),nil ;};return nil ,_gc .Errorf ("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054",_feaf );};func (_ecb *PSOperand )and (_cf *PSStack )error {_bab ,_gac :=_cf .Pop ();
if _gac !=nil {return _gac ;};_ea ,_gac :=_cf .Pop ();if _gac !=nil {return _gac ;};if _gfe ,_dff :=_bab .(*PSBoolean );_dff {_bde ,_gbg :=_ea .(*PSBoolean );if !_gbg {return ErrTypeCheck ;};_gac =_cf .Push (MakeBool (_gfe .Val &&_bde .Val ));return _gac ;
};if _gcd ,_agd :=_bab .(*PSInteger );_agd {_aed ,_bee :=_ea .(*PSInteger );if !_bee {return ErrTypeCheck ;};_gac =_cf .Push (MakeInteger (_gcd .Val &_aed .Val ));return _gac ;};return ErrTypeCheck ;};const _a =0.000001;

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject ;func (_ebb *PSOperand )atan (_ebbe *PSStack )error {_ab ,_bfe :=_ebbe .PopNumberAsFloat64 ();if _bfe !=nil {return _bfe ;};_dea ,_bfe :=_ebbe .PopNumberAsFloat64 ();if _bfe !=nil {return _bfe ;};if _ab ==0{var _bad error ;if _dea < 0{_bad =_ebbe .Push (MakeReal (270));
}else {_bad =_ebbe .Push (MakeReal (90));};return _bad ;};_afb :=_dea /_ab ;_ebc :=_d .Atan (_afb )*180/_d .Pi ;_bfe =_ebbe .Push (MakeReal (_ebc ));return _bfe ;};func (_agbb *PSOperand )sin (_cfag *PSStack )error {_aggg ,_dbfb :=_cfag .PopNumberAsFloat64 ();
if _dbfb !=nil {return _dbfb ;};_ecd :=_d .Sin (_aggg *_d .Pi /180.0);_dbfb =_cfag .Push (MakeReal (_ecd ));return _dbfb ;};

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser (content []byte )*PSParser {_cecg :=PSParser {};_gcgc :=_eg .NewBuffer (content );_cecg ._bdaf =_ga .NewReader (_gcgc );return &_cecg ;};func (_dg *PSOperand )div (_gbb *PSStack )error {_fege ,_fdb :=_gbb .Pop ();if _fdb !=nil {return _fdb ;
};_bgg ,_fdb :=_gbb .Pop ();if _fdb !=nil {return _fdb ;};_ged ,_feag :=_fege .(*PSReal );_aad ,_dbc :=_fege .(*PSInteger );if !_feag &&!_dbc {return ErrTypeCheck ;};if _feag &&_ged .Val ==0{return ErrUndefinedResult ;};if _dbc &&_aad .Val ==0{return ErrUndefinedResult ;
};_cde ,_bdb :=_bgg .(*PSReal );_acf ,_cdd :=_bgg .(*PSInteger );if !_bdb &&!_cdd {return ErrTypeCheck ;};var _fc float64 ;if _bdb {_fc =_cde .Val ;}else {_fc =float64 (_acf .Val );};if _feag {_fc /=_ged .Val ;}else {_fc /=float64 (_aad .Val );};_fdb =_gbb .Push (MakeReal (_fc ));
return _fdb ;};func (_bcc *PSOperand )cvi (_ceg *PSStack )error {_cbee ,_beee :=_ceg .Pop ();if _beee !=nil {return _beee ;};if _ffdd ,_bcf :=_cbee .(*PSReal );_bcf {_bfc :=int (_ffdd .Val );_beee =_ceg .Push (MakeInteger (_bfc ));}else if _gbgd ,_ecbd :=_cbee .(*PSInteger );
_ecbd {_bgc :=_gbgd .Val ;_beee =_ceg .Push (MakeInteger (_bgc ));}else {return ErrTypeCheck ;};return _beee ;};

// Exec executes the program, typically leaving output values on the stack.
func (_aac *PSProgram )Exec (stack *PSStack )error {for _ ,_bg :=range *_aac {var _ec error ;switch _gfd :=_bg .(type ){case *PSInteger :_gg :=_gfd ;_ec =stack .Push (_gg );case *PSReal :_ag :=_gfd ;_ec =stack .Push (_ag );case *PSBoolean :_ecf :=_gfd ;
_ec =stack .Push (_ecf );case *PSProgram :_eeg :=_gfd ;_ec =stack .Push (_eeg );case *PSOperand :_ad :=_gfd ;_ec =_ad .Exec (stack );default:return ErrTypeCheck ;};if _ec !=nil {return _ec ;};};return nil ;};

// String returns a string representation of the stack.
func (_eeggd *PSStack )String ()string {_cbga :="\u005b\u0020";for _ ,_aagga :=range *_eeggd {_cbga +=_aagga .String ();_cbga +="\u0020";};_cbga +="\u005d";return _cbga ;};

// Push pushes an object on top of the stack.
func (_abda *PSStack )Push (obj PSObject )error {if len (*_abda )> 100{return ErrStackOverflow ;};*_abda =append (*_abda ,obj );return nil ;};func (_dfd *PSOperand )xor (_edced *PSStack )error {_fagf ,_gcce :=_edced .Pop ();if _gcce !=nil {return _gcce ;
};_aceb ,_gcce :=_edced .Pop ();if _gcce !=nil {return _gcce ;};if _acfa ,_dgeg :=_fagf .(*PSBoolean );_dgeg {_feae ,_ffge :=_aceb .(*PSBoolean );if !_ffge {return ErrTypeCheck ;};_gcce =_edced .Push (MakeBool (_acfa .Val !=_feae .Val ));return _gcce ;
};if _ccc ,_dbb :=_fagf .(*PSInteger );_dbb {_cbc ,_ccgb :=_aceb .(*PSInteger );if !_ccgb {return ErrTypeCheck ;};_gcce =_edced .Push (MakeInteger (_ccc .Val ^_cbc .Val ));return _gcce ;};return ErrTypeCheck ;};func (_fbbd *PSOperand )eq (_gbdf *PSStack )error {_egg ,_fec :=_gbdf .Pop ();
if _fec !=nil {return _fec ;};_cga ,_fec :=_gbdf .Pop ();if _fec !=nil {return _fec ;};_adfe ,_ede :=_egg .(*PSBoolean );_cag ,_dfg :=_cga .(*PSBoolean );if _ede ||_dfg {var _adb error ;if _ede &&_dfg {_adb =_gbdf .Push (MakeBool (_adfe .Val ==_cag .Val ));
}else {_adb =_gbdf .Push (MakeBool (false ));};return _adb ;};var _dfb float64 ;var _ccd float64 ;if _bff ,_dbe :=_egg .(*PSInteger );_dbe {_dfb =float64 (_bff .Val );}else if _adbf ,_dedb :=_egg .(*PSReal );_dedb {_dfb =_adbf .Val ;}else {return ErrTypeCheck ;
};if _ffgd ,_ggb :=_cga .(*PSInteger );_ggb {_ccd =float64 (_ffgd .Val );}else if _cgef ,_eada :=_cga .(*PSReal );_eada {_ccd =_cgef .Val ;}else {return ErrTypeCheck ;};if _d .Abs (_ccd -_dfb )< _a {_fec =_gbdf .Push (MakeBool (true ));}else {_fec =_gbdf .Push (MakeBool (false ));
};return _fec ;};func (_ddg *PSOperand )not (_fgaf *PSStack )error {_fecc ,_edd :=_fgaf .Pop ();if _edd !=nil {return _edd ;};if _cca ,_eec :=_fecc .(*PSBoolean );_eec {_edd =_fgaf .Push (MakeBool (!_cca .Val ));return _edd ;}else if _bgf ,_egc :=_fecc .(*PSInteger );
_egc {_edd =_fgaf .Push (MakeInteger (^_bgf .Val ));return _edd ;}else {return ErrTypeCheck ;};};func (_fae *PSParser )parseOperand ()(*PSOperand ,error ){var _efg []byte ;for {_abf ,_aadd :=_fae ._bdaf .Peek (1);if _aadd !=nil {if _aadd ==_eb .EOF {break ;
};return nil ,_aadd ;};if _f .IsDelimiter (_abf [0]){break ;};if _f .IsWhiteSpace (_abf [0]){break ;};_edeb ,_ :=_fae ._bdaf .ReadByte ();_efg =append (_efg ,_edeb );};if len (_efg )==0{return nil ,_e .New ("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029");
};return MakeOperand (string (_efg )),nil ;};func (_fb *PSOperand )DebugString ()string {return _gc .Sprintf ("\u006fp\u003a\u0027\u0025\u0073\u0027",*_fb );};func (_aag *PSReal )String ()string {return _gc .Sprintf ("\u0025\u002e\u0035\u0066",_aag .Val )};
func (_aaf *PSOperand )mul (_bbc *PSStack )error {_ece ,_eggd :=_bbc .Pop ();if _eggd !=nil {return _eggd ;};_bbf ,_eggd :=_bbc .Pop ();if _eggd !=nil {return _eggd ;};_deb ,_cfe :=_ece .(*PSReal );_gcf ,_cgag :=_ece .(*PSInteger );if !_cfe &&!_cgag {return ErrTypeCheck ;
};_cgeg ,_gbf :=_bbf .(*PSReal );_edf ,_bfg :=_bbf .(*PSInteger );if !_gbf &&!_bfg {return ErrTypeCheck ;};if _cgag &&_bfg {_ced :=_gcf .Val *_edf .Val ;_bgbd :=_bbc .Push (MakeInteger (_ced ));return _bgbd ;};var _cef float64 ;if _cfe {_cef =_deb .Val ;
}else {_cef =float64 (_gcf .Val );};if _gbf {_cef *=_cgeg .Val ;}else {_cef *=float64 (_edf .Val );};_eggd =_bbc .Push (MakeReal (_cef ));return _eggd ;};func (_dge *PSOperand )index (_cadd *PSStack )error {_ggf ,_eff :=_cadd .Pop ();if _eff !=nil {return _eff ;
};_gdc ,_fdd :=_ggf .(*PSInteger );if !_fdd {return ErrTypeCheck ;};if _gdc .Val < 0{return ErrRangeCheck ;};if _gdc .Val > len (*_cadd )-1{return ErrStackUnderflow ;};_ebca :=(*_cadd )[len (*_cadd )-1-_gdc .Val ];_eff =_cadd .Push (_ebca .Duplicate ());
return _eff ;};

// PSObject represents a postscript object.
type PSObject interface{

// Duplicate makes a fresh copy of the PSObject.
Duplicate ()PSObject ;

// DebugString returns a descriptive representation of the PSObject with more information than String()
// for debugging purposes.
DebugString ()string ;

// String returns a string representation of the PSObject.
String ()string ;};var ErrRangeCheck =_e .New ("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072");func (_dag *PSOperand )exch (_fce *PSStack )error {_cbec ,_fcd :=_fce .Pop ();if _fcd !=nil {return _fcd ;
};_dgc ,_fcd :=_fce .Pop ();if _fcd !=nil {return _fcd ;};_fcd =_fce .Push (_cbec );if _fcd !=nil {return _fcd ;};_fcd =_fce .Push (_dgc );return _fcd ;};func (_beb *PSOperand )floor (_dbfe *PSStack )error {_agc ,_eed :=_dbfe .Pop ();if _eed !=nil {return _eed ;
};if _cea ,_bfde :=_agc .(*PSReal );_bfde {_eed =_dbfe .Push (MakeReal (_d .Floor (_cea .Val )));}else if _faaa ,_edbd :=_agc .(*PSInteger );_edbd {_eed =_dbfe .Push (MakeInteger (_faaa .Val ));}else {return ErrTypeCheck ;};return _eed ;};var ErrStackOverflow =_e .New ("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077");
func (_fbb *PSOperand )dup (_bccb *PSStack )error {_gbc ,_bded :=_bccb .Pop ();if _bded !=nil {return _bded ;};_bded =_bccb .Push (_gbc );if _bded !=nil {return _bded ;};_bded =_bccb .Push (_gbc .Duplicate ());return _bded ;};

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram ()*PSProgram {return &PSProgram {}};

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger (val int )*PSInteger {_baaf :=PSInteger {};_baaf .Val =val ;return &_baaf };func (_ccca *PSParser )parseFunction ()(*PSProgram ,error ){_cegc ,_ :=_ccca ._bdaf .ReadByte ();if _cegc !='{'{return nil ,_e .New ("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");
};_eedd :=NewPSProgram ();for {_ccca .skipSpaces ();_fgeb ,_fcde :=_ccca ._bdaf .Peek (2);if _fcde !=nil {if _fcde ==_eb .EOF {break ;};return nil ,_fcde ;};_b .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_fgeb ));
if _fgeb [0]=='}'{_b .Log .Trace ("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");_ccca ._bdaf .ReadByte ();break ;}else if _fgeb [0]=='{'{_b .Log .Trace ("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!");_bcg ,_gdae :=_ccca .parseFunction ();
if _gdae !=nil {return nil ,_gdae ;};_eedd .Append (_bcg );}else if _f .IsDecimalDigit (_fgeb [0])||(_fgeb [0]=='-'&&_f .IsDecimalDigit (_fgeb [1])){_b .Log .Trace ("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!");_ccfb ,_afcb :=_ccca .parseNumber ();if _afcb !=nil {return nil ,_afcb ;
};_eedd .Append (_ccfb );}else {_b .Log .Trace ("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f");_fgeb ,_ =_ccca ._bdaf .Peek (5);_gbdc :=string (_fgeb );_b .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_gbdc );
if (len (_gbdc )> 4)&&(_gbdc [:5]=="\u0066\u0061\u006cs\u0065"){_agff ,_fggc :=_ccca .parseBool ();if _fggc !=nil {return nil ,_fggc ;};_eedd .Append (_agff );}else if (len (_gbdc )> 3)&&(_gbdc [:4]=="\u0074\u0072\u0075\u0065"){_ebgb ,_efeb :=_ccca .parseBool ();
if _efeb !=nil {return nil ,_efeb ;};_eedd .Append (_ebgb );}else {_fegc ,_abc :=_ccca .parseOperand ();if _abc !=nil {return nil ,_abc ;};_eedd .Append (_fegc );};};};return _eedd ,nil ;};func (_cgaf *PSOperand )mod (_adbg *PSStack )error {_fbc ,_bbd :=_adbg .Pop ();
if _bbd !=nil {return _bbd ;};_eede ,_bbd :=_adbg .Pop ();if _bbd !=nil {return _bbd ;};_edce ,_gfac :=_fbc .(*PSInteger );if !_gfac {return ErrTypeCheck ;};if _edce .Val ==0{return ErrUndefinedResult ;};_ceab ,_gfac :=_eede .(*PSInteger );if !_gfac {return ErrTypeCheck ;
};_dca :=_ceab .Val %_edce .Val ;_bbd =_adbg .Push (MakeInteger (_dca ));return _bbd ;};func (_ega *PSInteger )DebugString ()string {return _gc .Sprintf ("\u0069\u006e\u0074\u003a\u0025\u0064",_ega .Val );};var ErrStackUnderflow =_e .New ("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077");
func (_aagd *PSOperand )Duplicate ()PSObject {_dcb :=*_aagd ;return &_dcb };func (_cgg *PSOperand )add (_ba *PSStack )error {_aga ,_efb :=_ba .Pop ();if _efb !=nil {return _efb ;};_agf ,_efb :=_ba .Pop ();if _efb !=nil {return _efb ;};_aaa ,_bd :=_aga .(*PSReal );
_afg ,_cbd :=_aga .(*PSInteger );if !_bd &&!_cbd {return ErrTypeCheck ;};_geb ,_aagg :=_agf .(*PSReal );_ffd ,_baa :=_agf .(*PSInteger );if !_aagg &&!_baa {return ErrTypeCheck ;};if _cbd &&_baa {_edb :=_afg .Val +_ffd .Val ;_deg :=_ba .Push (MakeInteger (_edb ));
return _deg ;};var _acd float64 ;if _bd {_acd =_aaa .Val ;}else {_acd =float64 (_afg .Val );};if _aagg {_acd +=_geb .Val ;}else {_acd +=float64 (_ffd .Val );};_efb =_ba .Push (MakeReal (_acd ));return _efb ;};func (_fafc *PSParser )parseBool ()(*PSBoolean ,error ){_fddc ,_gde :=_fafc ._bdaf .Peek (4);
if _gde !=nil {return MakeBool (false ),_gde ;};if (len (_fddc )>=4)&&(string (_fddc [:4])=="\u0074\u0072\u0075\u0065"){_fafc ._bdaf .Discard (4);return MakeBool (true ),nil ;};_fddc ,_gde =_fafc ._bdaf .Peek (5);if _gde !=nil {return MakeBool (false ),_gde ;
};if (len (_fddc )>=5)&&(string (_fddc [:5])=="\u0066\u0061\u006cs\u0065"){_fafc ._bdaf .Discard (5);return MakeBool (false ),nil ;};return MakeBool (false ),_e .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");
};func (_afc *PSBoolean )DebugString ()string {return _gc .Sprintf ("\u0062o\u006f\u006c\u003a\u0025\u0076",_afc .Val );};var ErrTypeCheck =_e .New ("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072");func (_eacg *PSOperand )sqrt (_cba *PSStack )error {_fgb ,_fgba :=_cba .PopNumberAsFloat64 ();
if _fgba !=nil {return _fgba ;};if _fgb < 0{return ErrRangeCheck ;};_aec :=_d .Sqrt (_fgb );_fgba =_cba .Push (MakeReal (_aec ));return _fgba ;};

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor (program *PSProgram )*PSExecutor {_bc :=&PSExecutor {};_bc .Stack =NewPSStack ();_bc ._fa =program ;return _bc ;};func (_gaa *PSInteger )String ()string {return _gc .Sprintf ("\u0025\u0064",_gaa .Val )};

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject ;

// PSInteger represents an integer.
type PSInteger struct{Val int ;};

// Exec executes the operand `op` in the state specified by `stack`.
func (_df *PSOperand )Exec (stack *PSStack )error {_gef :=ErrUnsupportedOperand ;switch *_df {case "\u0061\u0062\u0073":_gef =_df .abs (stack );case "\u0061\u0064\u0064":_gef =_df .add (stack );case "\u0061\u006e\u0064":_gef =_df .and (stack );case "\u0061\u0074\u0061\u006e":_gef =_df .atan (stack );
case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":_gef =_df .bitshift (stack );case "\u0063e\u0069\u006c\u0069\u006e\u0067":_gef =_df .ceiling (stack );case "\u0063\u006f\u0070\u0079":_gef =_df .copy (stack );case "\u0063\u006f\u0073":_gef =_df .cos (stack );
case "\u0063\u0076\u0069":_gef =_df .cvi (stack );case "\u0063\u0076\u0072":_gef =_df .cvr (stack );case "\u0064\u0069\u0076":_gef =_df .div (stack );case "\u0064\u0075\u0070":_gef =_df .dup (stack );case "\u0065\u0071":_gef =_df .eq (stack );case "\u0065\u0078\u0063\u0068":_gef =_df .exch (stack );
case "\u0065\u0078\u0070":_gef =_df .exp (stack );case "\u0066\u006c\u006fo\u0072":_gef =_df .floor (stack );case "\u0067\u0065":_gef =_df .ge (stack );case "\u0067\u0074":_gef =_df .gt (stack );case "\u0069\u0064\u0069\u0076":_gef =_df .idiv (stack );
case "\u0069\u0066":_gef =_df .ifCondition (stack );case "\u0069\u0066\u0065\u006c\u0073\u0065":_gef =_df .ifelse (stack );case "\u0069\u006e\u0064e\u0078":_gef =_df .index (stack );case "\u006c\u0065":_gef =_df .le (stack );case "\u006c\u006f\u0067":_gef =_df .log (stack );
case "\u006c\u006e":_gef =_df .ln (stack );case "\u006c\u0074":_gef =_df .lt (stack );case "\u006d\u006f\u0064":_gef =_df .mod (stack );case "\u006d\u0075\u006c":_gef =_df .mul (stack );case "\u006e\u0065":_gef =_df .ne (stack );case "\u006e\u0065\u0067":_gef =_df .neg (stack );
case "\u006e\u006f\u0074":_gef =_df .not (stack );case "\u006f\u0072":_gef =_df .or (stack );case "\u0070\u006f\u0070":_gef =_df .pop (stack );case "\u0072\u006f\u0075n\u0064":_gef =_df .round (stack );case "\u0072\u006f\u006c\u006c":_gef =_df .roll (stack );
case "\u0073\u0069\u006e":_gef =_df .sin (stack );case "\u0073\u0071\u0072\u0074":_gef =_df .sqrt (stack );case "\u0073\u0075\u0062":_gef =_df .sub (stack );case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":_gef =_df .truncate (stack );case "\u0078\u006f\u0072":_gef =_df .xor (stack );
};return _gef ;};func (_dec *PSOperand )ne (_gda *PSStack )error {_dad :=_dec .eq (_gda );if _dad !=nil {return _dad ;};_dad =_dec .not (_gda );return _dad ;};func (_gdb *PSOperand )or (_dceb *PSStack )error {_accb ,_deag :=_dceb .Pop ();if _deag !=nil {return _deag ;
};_cgd ,_deag :=_dceb .Pop ();if _deag !=nil {return _deag ;};if _acag ,_beeb :=_accb .(*PSBoolean );_beeb {_ebge ,_afe :=_cgd .(*PSBoolean );if !_afe {return ErrTypeCheck ;};_deag =_dceb .Push (MakeBool (_acag .Val ||_ebge .Val ));return _deag ;};if _bdag ,_cfae :=_accb .(*PSInteger );
_cfae {_gad ,_fbeb :=_cgd .(*PSInteger );if !_fbeb {return ErrTypeCheck ;};_deag =_dceb .Push (MakeInteger (_bdag .Val |_gad .Val ));return _deag ;};return ErrTypeCheck ;};

// PSReal represents a real number.
type PSReal struct{Val float64 ;};func (_feg *PSBoolean )String ()string {return _gc .Sprintf ("\u0025\u0076",_feg .Val )};func (_cbb *PSOperand )exp (_gfg *PSStack )error {_abd ,_aced :=_gfg .PopNumberAsFloat64 ();if _aced !=nil {return _aced ;};_gbaa ,_aced :=_gfg .PopNumberAsFloat64 ();
if _aced !=nil {return _aced ;};if _d .Abs (_abd )< 1&&_gbaa < 0{return ErrUndefinedResult ;};_cega :=_d .Pow (_gbaa ,_abd );_aced =_gfg .Push (MakeReal (_cega ));return _aced ;};