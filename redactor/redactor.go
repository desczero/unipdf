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

package redactor ;import (_a "errors";_f "fmt";_fa "github.com/unidoc/unipdf/v3/common";_ec "github.com/unidoc/unipdf/v3/contentstream";_cc "github.com/unidoc/unipdf/v3/core";_ae "github.com/unidoc/unipdf/v3/creator";_ef "github.com/unidoc/unipdf/v3/extractor";
_g "github.com/unidoc/unipdf/v3/model";_cg "io";_fg "regexp";_fgd "sort";_e "strings";);func _gfbc (_ede localSpanMarks ,_abc *_ef .TextMarkArray ,_cgd *_g .PdfFont ,_ade ,_bde string )([]_cc .PdfObject ,error ){_dee :=_bfced (_abc );Tj ,_bec :=_ggbg (_abc );
if _bec !=nil {return nil ,_bec ;};_bbf :=len (_ade );_beg :=len (_dee );_dcg :=-1;_ggfa :=_cc .MakeFloat (Tj );if _dee !=_bde {_fb :=_ede ._ggfc ;if _fb ==0{_dcg =_e .LastIndex (_ade ,_dee );}else {_dcg =_e .Index (_ade ,_dee );};}else {_dcg =_e .Index (_ade ,_dee );
};_bbfb :=_dcg +_beg ;_bdef :=[]_cc .PdfObject {};if _dcg ==0&&_bbfb ==_bbf {_bdef =append (_bdef ,_ggfa );}else if _dcg ==0&&_bbfb < _bbf {_bgd :=_caa (_ade [_bbfb :],_cgd );_dfd :=_cc .MakeStringFromBytes (_bgd );_bdef =append (_bdef ,_ggfa ,_dfd );}else if _dcg > 0&&_bbfb >=_bbf {_eac :=_caa (_ade [:_dcg ],_cgd );
_gac :=_cc .MakeStringFromBytes (_eac );_bdef =append (_bdef ,_gac ,_ggfa );}else if _dcg > 0&&_bbfb < _bbf {_aaf :=_caa (_ade [:_dcg ],_cgd );_ddg :=_caa (_ade [_bbfb :],_cgd );_fbg :=_cc .MakeStringFromBytes (_aaf );_bcg :=_cc .MakeString (string (_ddg ));
_bdef =append (_bdef ,_fbg ,_ggfa ,_bcg );};return _bdef ,nil ;};func _cege (_gade string )(string ,[][]int ){_gafd :=_fg .MustCompile ("\u005c\u006e");_efce :=_gafd .FindAllStringIndex (_gade ,-1);_edc :=_gafd .ReplaceAllString (_gade ,"\u0020");return _edc ,_efce ;
};func _gfb (_dc *_ef .TextMarkArray )int {_cgb :=0;_eea :=_dc .Elements ();if _eea [0].Text =="\u0020"{_cgb ++;};if _eea [_dc .Len ()-1].Text =="\u0020"{_cgb ++;};return _cgb ;};

// Write writes the content of `re.creator` to writer of type io.Writer interface.
func (_fea *Redactor )Write (writer _cg .Writer )error {return _fea ._cfgd .Write (writer )};

// RedactRectanglePropsNew return a new pointer to a default RectangleProps object.
func RedactRectanglePropsNew ()*RectangleProps {return &RectangleProps {FillColor :_ae .ColorBlack ,BorderWidth :0.0,FillOpacity :1.0};};

// Redact executes the redact operation on a pdf file and updates the content streams of all pages of the file.
func (_dceg *Redactor )Redact ()error {_bef ,_ced :=_dceg ._adc .GetNumPages ();if _ced !=nil {return _f .Errorf ("\u0066\u0061\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0074\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020P\u0061\u0067\u0065\u0073");
};_fgg :=_dceg ._deca .FillColor ;_befa :=_dceg ._deca .BorderWidth ;_cceg :=_dceg ._deca .FillOpacity ;for _eae :=1;_eae <=_bef ;_eae ++{_aegc ,_gcfa :=_dceg ._adc .GetPage (_eae );if _gcfa !=nil {return _gcfa ;};_bece ,_gcfa :=_ef .New (_aegc );if _gcfa !=nil {return _gcfa ;
};_acac ,_ ,_ ,_gcfa :=_bece .ExtractPageText ();if _gcfa !=nil {return _gcfa ;};_aeca :=_acac .GetContentStreamOps ();_edg ,_aaaf ,_gcfa :=_dceg .redactPage (_aeca ,_aegc .Resources );if _aaaf ==nil {_fa .Log .Info ("N\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u0020t\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065d \u0070\u0061\u0074t\u0061r\u006e\u002e");
_aaaf =_aeca ;};_gcgdb :=_ec .ContentStreamOperation {Operand :"\u006e"};*_aaaf =append (*_aaaf ,&_gcgdb );_aegc .SetContentStreams ([]string {_aaaf .String ()},_cc .NewFlateEncoder ());if _gcfa !=nil {return _gcfa ;};_ffaa ,_gcfa :=_aegc .GetMediaBox ();
if _gcfa !=nil {return _gcfa ;};if _aegc .MediaBox ==nil {_aegc .MediaBox =_ffaa ;};if _eeeg :=_dceg ._cfgd .AddPage (_aegc );_eeeg !=nil {return _eeeg ;};_fgd .Slice (_edg ,func (_faaa ,_ebfb int )bool {return _edg [_faaa ]._gcfb < _edg [_ebfb ]._gcfb });
_cedd :=_ffaa .Ury ;for _ ,_gde :=range _edg {_bga :=_gde ._bgb ;_adbc :=_dceg ._cfgd .NewRectangle (_bga .Llx ,_cedd -_bga .Lly ,_bga .Urx -_bga .Llx ,-(_bga .Ury -_bga .Lly ));_adbc .SetFillColor (_fgg );_adbc .SetBorderWidth (_befa );_adbc .SetFillOpacity (_cceg );
if _abba :=_dceg ._cfgd .Draw (_adbc );_abba !=nil {return nil ;};};};_dceg ._cfgd .SetOutlineTree (_dceg ._adc .GetOutlineTree ());return nil ;};func _gddc (_egb ,_gfbca targetMap )(bool ,[]int ){_caab :=_e .Contains (_egb ._geea ,_gfbca ._geea );var _ebfg []int ;
for _ ,_bcgf :=range _egb ._cfdd {for _bcd ,_bgg :=range _gfbca ._cfdd {if _bgg [0]>=_bcgf [0]&&_bgg [1]<=_bcgf [1]{_ebfg =append (_ebfg ,_bcd );};};};return _caab ,_ebfg ;};func _eb (_ded []localSpanMarks )(map[string ][]localSpanMarks ,[]string ){_daf :=make (map[string ][]localSpanMarks );
_acd :=[]string {};for _ ,_eafc :=range _ded {_acg :=_eafc ._ccb ;if _gfd ,_cbf :=_daf [_acg ];_cbf {_daf [_acg ]=append (_gfd ,_eafc );}else {_daf [_acg ]=[]localSpanMarks {_eafc };_acd =append (_acd ,_acg );};};return _daf ,_acd ;};func _fbc (_efe []placeHolders )[]replacement {_cbe :=[]replacement {};
for _ ,_bea :=range _efe {_gaea :=_bea ._b ;_cfgg :=_bea ._gc ;_aaa :=_bea ._cf ;for _ ,_bfcd :=range _gaea {_bff :=replacement {_d :_cfgg ,_cfb :_aaa ,_ge :_bfcd };_cbe =append (_cbe ,_bff );};};_fgd .Slice (_cbe ,func (_fgcf ,_ead int )bool {return _cbe [_fgcf ]._ge < _cbe [_ead ]._ge });
return _cbe ;};type localSpanMarks struct{_dbc *_ef .TextMarkArray ;_ggfc int ;_ccb string ;};func _bfdf (_fgf *_ef .TextMarkArray )[]*_ef .TextMarkArray {_efc :=_fgf .Elements ();_acdf :=len (_efc );var _ecf _cc .PdfObject ;_caf :=[]*_ef .TextMarkArray {};
_eabc :=&_ef .TextMarkArray {};_gcc :=-1;for _dgfg ,_gccc :=range _efc {_dgg :=_gccc .DirectObject ;_gcc =_gccc .Index ;if _dgg ==nil {_eaa :=_gab (_fgf ,_dgfg ,_gcc );if _ecf !=nil {if _eaa ==-1||_eaa > _dgfg {_caf =append (_caf ,_eabc );_eabc =&_ef .TextMarkArray {};
};};}else if _dgg !=nil &&_ecf ==nil {if _gcc ==0&&_dgfg > 0{_caf =append (_caf ,_eabc );_eabc =&_ef .TextMarkArray {};};}else if _dgg !=nil &&_ecf !=nil {if _dgg !=_ecf {_caf =append (_caf ,_eabc );_eabc =&_ef .TextMarkArray {};};};_ecf =_dgg ;_eabc .Append (_gccc );
if _dgfg ==(_acdf -1){_caf =append (_caf ,_eabc );};};return _caf ;};func _cfbed (_bbb string ,_ggd []localSpanMarks )([]placeHolders ,error ){_bgf :="";_dce :=[]placeHolders {};for _gba ,_ag :=range _ggd {_fdaa :=_ag ._dbc ;_ddc :=_ag ._ccb ;_abb :=_bfced (_fdaa );
_bcb ,_aafa :=_ggbg (_fdaa );if _aafa !=nil {return nil ,_aafa ;};if _abb !=_bgf {var _eafg []int ;if _gba ==0&&_ddc !=_abb {_eec :=_e .Index (_bbb ,_abb );_eafg =[]int {_eec };}else if _gba ==len (_ggd )-1{_bca :=_e .LastIndex (_bbb ,_abb );_eafg =[]int {_bca };
}else {_eafg =_ggfe (_bbb ,_abb );};_db :=placeHolders {_b :_eafg ,_gc :_abb ,_cf :_bcb };_dce =append (_dce ,_db );};_bgf =_abb ;};return _dce ,nil ;};func _dbcdg (_acc []int ,_dgd *_ef .TextMarkArray ,_ceg string )(*_ef .TextMarkArray ,matchedBBox ,error ){_gaef :=matchedBBox {};
_dbb :=_acc [0];_dfaa :=_acc [1];_gcbd :=len (_ceg )-len (_e .TrimLeft (_ceg ,"\u0020"));_ccdc :=len (_ceg )-len (_e .TrimRight (_ceg ,"\u0020\u000a"));_dbb =_dbb +_gcbd ;_dfaa =_dfaa -_ccdc ;_ceg =_e .Trim (_ceg ,"\u0020\u000a");_gebg ,_faad :=_dgd .RangeOffset (_dbb ,_dfaa );
if _faad !=nil {return nil ,_gaef ,_faad ;};_fedb ,_dcdf :=_gebg .BBox ();if !_dcdf {return nil ,_gaef ,_f .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_gebg );
};_gaef =matchedBBox {_gcfb :_ceg ,_bgb :_fedb };return _gebg ,_gaef ,nil ;};

// New instantiates a Redactor object with given PdfReader and `regex` pattern.
func New (reader *_g .PdfReader ,opts *RedactionOptions ,rectProps *RectangleProps )*Redactor {if rectProps ==nil {rectProps =RedactRectanglePropsNew ();};return &Redactor {_adc :reader ,_fgdf :opts ,_cfgd :_ae .New (),_deca :rectProps };};func (_fafa *regexMatcher )match (_cea string )([]*matchedIndex ,error ){_gec :=_fafa ._fbgc .Pattern ;
if _gec ==nil {return nil ,_a .New ("\u006e\u006f\u0020\u0070at\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006d\u0070\u0069\u006c\u0065\u0064");};var (_fgcb =_gec .FindAllStringIndex (_cea ,-1);_egac []*matchedIndex ;);for _ ,_afc :=range _fgcb {_egac =append (_egac ,&matchedIndex {_affd :_afc [0],_acdg :_afc [1],_afb :_cea [_afc [0]:_afc [1]]});
};return _egac ,nil ;};func _ggeb (_ebdb []*matchedIndex )[]*targetMap {_cfe :=make (map[string ][][]int );_eeca :=[]*targetMap {};for _ ,_abbc :=range _ebdb {_fce :=_abbc ._afb ;_bcbb :=[]int {_abbc ._affd ,_abbc ._acdg };if _dedb ,_bdfd :=_cfe [_fce ];
_bdfd {_cfe [_fce ]=append (_dedb ,_bcbb );}else {_cfe [_fce ]=[][]int {_bcbb };};};for _ecg ,_aac :=range _cfe {_aed :=&targetMap {_geea :_ecg ,_cfdd :_aac };_eeca =append (_eeca ,_aed );};return _eeca ;};type targetMap struct{_geea string ;_cfdd [][]int ;
};

// Redactor represents a Redactor object.
type Redactor struct{_adc *_g .PdfReader ;_fgdf *RedactionOptions ;_cfgd *_ae .Creator ;_deca *RectangleProps ;};func _bcf (_fdb *_ec .ContentStreamOperation ,_efb _cc .PdfObject ,_gbf []localSpanMarks )error {var _ab *_cc .PdfObjectArray ;_gea ,_cgf :=_eb (_gbf );
if len (_cgf )==1{_feb :=_cgf [0];_gga :=_gea [_feb ];if len (_gga )==1{_gd :=_gga [0];_dfa :=_gd ._dbc ;_bg :=_dcf (_dfa );_ffg ,_adba :=_eacb (_efb ,_bg );if _adba !=nil {return _adba ;};_aec ,_adba :=_gfbc (_gd ,_dfa ,_bg ,_ffg ,_feb );if _adba !=nil {return _adba ;
};_ab =_cc .MakeArray (_aec ...);}else {_gcf :=_gga [0]._dbc ;_cfbe :=_dcf (_gcf );_egf ,_gge :=_eacb (_efb ,_cfbe );if _gge !=nil {return _gge ;};_eccc ,_gge :=_cfbed (_egf ,_gga );if _gge !=nil {return _gge ;};_cac :=_fbc (_eccc );_gdd :=_gad (_egf ,_cac ,_cfbe );
_ab =_cc .MakeArray (_gdd ...);};}else if len (_cgf )> 1{_fec :=_gbf [0];_bad :=_fec ._dbc ;_ ,_bfce :=_eeg (_bad );_cfg :=_bad .Elements ()[_bfce ];_ffd :=_cfg .Font ;_gcdf ,_gae :=_eacb (_efb ,_ffd );if _gae !=nil {return _gae ;};_edd ,_gae :=_cfbed (_gcdf ,_gbf );
if _gae !=nil {return _gae ;};_ebf :=_fbc (_edd );_gce :=_gad (_gcdf ,_ebf ,_ffd );_ab =_cc .MakeArray (_gce ...);};_fdb .Params [0]=_ab ;_fdb .Operand ="\u0054\u004a";return nil ;};func _decc (_bdc []*targetMap ){for _fcfd ,_febd :=range _bdc {for _adf ,_dbfe :=range _bdc {if _fcfd !=_adf {_gfbg ,_ffgd :=_gddc (*_febd ,*_dbfe );
if _gfbg {_caef (_dbfe ,_ffgd );};};};};};func _ba (_eead *_ec .ContentStreamOperation ,_dddb _cc .PdfObject ,_fe []localSpanMarks )error {_gee ,_gcb :=_cc .GetArray (_eead .Params [0]);_ea :=[]_cc .PdfObject {};if !_gcb {_fa .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064",_eead );
return _f .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_eead );
};_aeg ,_ed :=_eb (_fe );if len (_ed )==1{_aa :=_ed [0];_bf :=_aeg [_aa ];if len (_bf )==1{_bc :=_bf [0];_be :=_bc ._dbc ;_gb :=_dcf (_be );_fag ,_faa :=_eacb (_dddb ,_gb );if _faa !=nil {return _faa ;};_aff ,_faa :=_gfbc (_bc ,_be ,_gb ,_fag ,_aa );if _faa !=nil {return _faa ;
};for _ ,_aeb :=range _gee .Elements (){if _aeb ==_dddb {_ea =append (_ea ,_aff ...);}else {_ea =append (_ea ,_aeb );};};}else {_ecd :=_bf [0]._dbc ;_egg :=_dcf (_ecd );_eebe ,_de :=_eacb (_dddb ,_egg );if _de !=nil {return _de ;};_cfd ,_de :=_cfbed (_eebe ,_bf );
if _de !=nil {return _de ;};_eab :=_fbc (_cfd );_fda :=_gad (_eebe ,_eab ,_egg );for _ ,_bfg :=range _gee .Elements (){if _bfg ==_dddb {_ea =append (_ea ,_fda ...);}else {_ea =append (_ea ,_bfg );};};};_eead .Params [0]=_cc .MakeArray (_ea ...);}else if len (_ed )> 1{_ggf :=_fe [0];
_fdd :=_ggf ._dbc ;_ ,_ga :=_eeg (_fdd );_cb :=_fdd .Elements ()[_ga ];_bee :=_cb .Font ;_afe ,_aab :=_eacb (_dddb ,_bee );if _aab !=nil {return _aab ;};_adb ,_aab :=_cfbed (_afe ,_fe );if _aab !=nil {return _aab ;};_dg :=_fbc (_adb );_gfff :=_gad (_afe ,_dg ,_bee );
for _ ,_eaf :=range _gee .Elements (){if _eaf ==_dddb {_ea =append (_ea ,_gfff ...);}else {_ea =append (_ea ,_eaf );};};_eead .Params [0]=_cc .MakeArray (_ea ...);};return nil ;};

// RedactionOptions is a collection of RedactionTerm objects.
type RedactionOptions struct{Terms []RedactionTerm ;};type placeHolders struct{_b []int ;_gc string ;_cf float64 ;};func _da (_gg *_ec .ContentStreamOperations ,_gf map[_cc .PdfObject ][]localSpanMarks )error {for _dd ,_ca :=range _gf {if _dd ==nil {continue ;
};_cae ,_fgc ,_fd :=_aee (_gg ,_dd );if !_fd {_fa .Log .Debug ("Pd\u0066\u004fb\u006a\u0065\u0063\u0074\u0020\u0025\u0073\u006e\u006ft\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0069\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073\u0074r\u0065a\u006d\u0020\u006f\u0070\u0065\u0072\u0061\u0074i\u006fn\u0020\u0025s",_dd ,_gg );
return nil ;};if _cae .Operand =="\u0054\u006a"{_geb :=_bcf (_cae ,_dd ,_ca );if _geb !=nil {return _geb ;};}else if _cae .Operand =="\u0054\u004a"{_ee :=_ba (_cae ,_dd ,_ca );if _ee !=nil {return _ee ;};}else if _cae .Operand =="\u0027"||_cae .Operand =="\u0022"{_eeb :=_ad (_gg ,_cae .Operand ,_fgc );
if _eeb !=nil {return _eeb ;};_eeb =_bcf (_cae ,_dd ,_ca );if _eeb !=nil {return _eeb ;};};};return nil ;};func _afee (_aaac *matchedIndex ,_agfb [][]int )(bool ,[][]int ){_daga :=[][]int {};for _ ,_eag :=range _agfb {if _aaac ._affd < _eag [0]&&_aaac ._acdg > _eag [1]{_daga =append (_daga ,_eag );
};};return len (_daga )> 0,_daga ;};func _caef (_bcfe *targetMap ,_begf []int ){var _cgc [][]int ;for _bdge ,_gbb :=range _bcfe ._cfdd {if _fgb (_bdge ,_begf ){continue ;};_cgc =append (_cgc ,_gbb );};_bcfe ._cfdd =_cgc ;};func _eeg (_af *_ef .TextMarkArray )(_cc .PdfObject ,int ){var _dcd _cc .PdfObject ;
_gfe :=-1;for _dda ,_ff :=range _af .Elements (){_dcd =_ff .DirectObject ;_gfe =_dda ;if _dcd !=nil {break ;};};return _dcd ,_gfe ;};type replacement struct{_d string ;_cfb float64 ;_ge int ;};type matchedBBox struct{_bgb _g .PdfRectangle ;_gcfb string ;
};func _df (_ffa *_g .PdfFont ,_ecc _ef .TextMark )float64 {_ffc :=0.001;_fc :=_ecc .Th /100;if _ffa .Subtype ()=="\u0054\u0079\u0070e\u0033"{_ffc =1;};_dca ,_cd :=_ffa .GetRuneMetrics (' ');if !_cd {_dca ,_cd =_ffa .GetCharMetrics (32);};if !_cd {_dca ,_ =_g .DefaultFont ().GetRuneMetrics (' ');
};_fca :=_ffc *((_dca .Wx *_ecc .FontSize +_ecc .Tc +_ecc .Tw )/_fc );return _fca ;};func _dcf (_fff *_ef .TextMarkArray )*_g .PdfFont {_ ,_cdc :=_eeg (_fff );_ccg :=_fff .Elements ()[_cdc ];_fge :=_ccg .Font ;return _fge ;};func _gab (_gcbcb *_ef .TextMarkArray ,_cegd int ,_cee int )int {_eddc :=_gcbcb .Elements ();
_bbg :=_cegd -1;_deeb :=_cegd +1;_bgc :=-1;if _bbg >=0{_efae :=_eddc [_bbg ];_aeed :=_efae .ObjString ;_agf :=len (_aeed );_ffb :=_efae .Index ;if _ffb +1< _agf {_bgc =_bbg ;return _bgc ;};};if _deeb < len (_eddc ){_fbec :=_eddc [_deeb ];_acgg :=_fbec .ObjString ;
if _acgg [0]!=_fbec .Text {_bgc =_deeb ;return _bgc ;};};return _bgc ;};func _geg (_aafgg RedactionTerm )(*regexMatcher ,error ){return &regexMatcher {_fbgc :_aafgg },nil };func _aee (_acb *_ec .ContentStreamOperations ,PdfObj _cc .PdfObject )(*_ec .ContentStreamOperation ,int ,bool ){for _bbbe ,_dedf :=range *_acb {_gaa :=_dedf .Operand ;
if _gaa =="\u0054\u006a"{_fdfa :=_cc .TraceToDirectObject (_dedf .Params [0]);if _fdfa ==PdfObj {return _dedf ,_bbbe ,true ;};}else if _gaa =="\u0054\u004a"{_ega ,_dece :=_cc .GetArray (_dedf .Params [0]);if !_dece {return nil ,_bbbe ,_dece ;};for _ ,_aad :=range _ega .Elements (){if _aad ==PdfObj {return _dedf ,_bbbe ,true ;
};};}else if _gaa =="\u0022"{_gafc :=_cc .TraceToDirectObject (_dedf .Params [2]);if _gafc ==PdfObj {return _dedf ,_bbbe ,true ;};}else if _gaa =="\u0027"{_gbae :=_cc .TraceToDirectObject (_dedf .Params [0]);if _gbae ==PdfObj {return _dedf ,_bbbe ,true ;
};};};return nil ,-1,false ;};type regexMatcher struct{_fbgc RedactionTerm };func (_bgbc *Redactor )redactPage (_eeba *_ec .ContentStreamOperations ,_geeb *_g .PdfPageResources )([]matchedBBox ,*_ec .ContentStreamOperations ,error ){_abd ,_cec :=_ef .NewFromContents (_eeba .String (),_geeb );
if _cec !=nil {return nil ,nil ,_cec ;};_fdc ,_ ,_ ,_cec :=_abd .ExtractPageText ();if _cec !=nil {return nil ,nil ,_cec ;};_eeba =_fdc .GetContentStreamOps ();_abde :=_fdc .Marks ();_dbcd :=_fdc .Text ();_dbcd ,_ebc :=_cege (_dbcd );_adg :=[]matchedBBox {};
_gaaa :=make (map[_cc .PdfObject ][]localSpanMarks );_age :=[]*targetMap {};for _ ,_ggab :=range _bgbc ._fgdf .Terms {_acad ,_cag :=_geg (_ggab );if _cag !=nil {return nil ,nil ,_cag ;};_cbfa ,_cag :=_acad .match (_dbcd );if _cag !=nil {return nil ,nil ,_cag ;
};_cbfa =_afge (_cbfa ,_ebc );_dgfc :=_ggeb (_cbfa );_age =append (_age ,_dgfc ...);};_decc (_age );for _ ,_gace :=range _age {_dcc :=_gace ._geea ;_aae :=_gace ._cfdd ;_agb :=[]matchedBBox {};for _ ,_gcbc :=range _aae {_adcb ,_gca ,_dag :=_dbcdg (_gcbc ,_abde ,_dcc );
if _dag !=nil {return nil ,nil ,_dag ;};_bae :=_bfdf (_adcb );for _fad ,_bac :=range _bae {_gebb :=localSpanMarks {_dbc :_bac ,_ggfc :_fad ,_ccb :_dcc };_acdb ,_ :=_eeg (_bac );if _gdea ,_dac :=_gaaa [_acdb ];_dac {_gaaa [_acdb ]=append (_gdea ,_gebb );
}else {_gaaa [_acdb ]=[]localSpanMarks {_gebb };};};_agb =append (_agb ,_gca );};_adg =append (_adg ,_agb ...);};_cec =_da (_eeba ,_gaaa );if _cec !=nil {return nil ,nil ,_cec ;};return _adg ,_eeba ,nil ;};func _feg (_cgde *matchedIndex ,_aafg [][]int )[]*matchedIndex {_dde :=[]*matchedIndex {};
_efeg :=_cgde ._affd ;_eegb :=_efeg ;_bbfg :=_cgde ._afb ;_cfgb :=0;for _ ,_bacd :=range _aafg {_gcfbe :=_bacd [0]-_efeg ;if _cfgb >=_gcfbe {continue ;};_ecgb :=_bbfg [_cfgb :_gcfbe ];_fcac :=&matchedIndex {_afb :_ecgb ,_affd :_eegb ,_acdg :_bacd [0]};
if len (_e .TrimSpace (_ecgb ))!=0{_dde =append (_dde ,_fcac );};_cfgb =_bacd [1]-_efeg ;_eegb =_efeg +_cfgb ;};_dff :=_bbfg [_cfgb :];_aafb :=&matchedIndex {_afb :_dff ,_affd :_eegb ,_acdg :_cgde ._acdg };if len (_e .TrimSpace (_dff ))!=0{_dde =append (_dde ,_aafb );
};return _dde ;};func _fcf (_ebd ,_efbd ,_fac float64 )float64 {_fac =_fac /100;_bfee :=(-1000*_ebd )/(_efbd *_fac );return _bfee ;};

// WriteToFile writes the redacted document to file specified by `outputPath`.
func (_gdec *Redactor )WriteToFile (outputPath string )error {if _dccg :=_gdec ._cfgd .WriteToFile (outputPath );_dccg !=nil {return _f .Errorf ("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0077\u0072\u0069\u0074\u0065\u0020t\u0068\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0066\u0069\u006c\u0065");
};return nil ;};func _ad (_ac *_ec .ContentStreamOperations ,_ccd string ,_ddd int )error {_gff :=_ec .ContentStreamOperations {};var _cfc _ec .ContentStreamOperation ;for _eg ,_gcd :=range *_ac {if _eg ==_ddd {if _ccd =="\u0027"{_cce :=_ec .ContentStreamOperation {Operand :"\u0054\u002a"};
_gff =append (_gff ,&_cce );_cfc .Params =_gcd .Params ;_cfc .Operand ="\u0054\u006a";_gff =append (_gff ,&_cfc );}else if _ccd =="\u0022"{_dae :=_gcd .Params [:2];Tc ,Tw :=_dae [0],_dae [1];_ged :=_ec .ContentStreamOperation {Params :[]_cc .PdfObject {Tc },Operand :"\u0054\u0063"};
_gff =append (_gff ,&_ged );_ged =_ec .ContentStreamOperation {Params :[]_cc .PdfObject {Tw },Operand :"\u0054\u0077"};_gff =append (_gff ,&_ged );_cfc .Params =[]_cc .PdfObject {_gcd .Params [2]};_cfc .Operand ="\u0054\u006a";_gff =append (_gff ,&_cfc );
};};_gff =append (_gff ,_gcd );};*_ac =_gff ;return nil ;};func _bfced (_gcgd *_ef .TextMarkArray )string {_afff :="";for _ ,_bbff :=range _gcgd .Elements (){_afff +=_bbff .Text ;};return _afff ;};

// RedactionTerm holds the regexp pattern and the replacement string for the redaction process.
type RedactionTerm struct{Pattern *_fg .Regexp ;};func _fgb (_dafg int ,_gacg []int )bool {for _ ,_fcd :=range _gacg {if _fcd ==_dafg {return true ;};};return false ;};func _eacb (_dge _cc .PdfObject ,_dfc *_g .PdfFont )(string ,error ){_ccdf ,_eace :=_cc .GetStringBytes (_dge );
if !_eace {return "",_cc .ErrTypeError ;};_def :=_dfc .BytesToCharcodes (_ccdf );_ege ,_fbd ,_egeb :=_dfc .CharcodesToStrings (_def ,"");if _egeb > 0{_fa .Log .Debug ("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064",_fbd ,_egeb );
};_eece :=_e .Join (_ege ,"");return _eece ,nil ;};type matchedIndex struct{_affd int ;_acdg int ;_afb string ;};func _ggfe (_faag ,_dab string )[]int {if len (_dab )==0{return nil ;};var _bdg []int ;for _gcdd :=0;_gcdd < len (_faag );{_dgfd :=_e .Index (_faag [_gcdd :],_dab );
if _dgfd < 0{return _bdg ;};_bdg =append (_bdg ,_gcdd +_dgfd );_gcdd +=_dgfd +len (_dab );};return _bdg ;};func _caa (_bd string ,_bb *_g .PdfFont )[]byte {_ce ,_cdg :=_bb .StringToCharcodeBytes (_bd );if _cdg !=0{_fa .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0073\u006fm\u0065\u0020\u0072un\u0065\u0073\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065d\u002e\u000a\u0009\u0025\u0073\u0020\u002d\u003e \u0025\u0076",_bd ,_ce );
};return _ce ;};

// RectangleProps defines properties of the redaction rectangle to be drawn.
type RectangleProps struct{FillColor _ae .Color ;BorderWidth float64 ;FillOpacity float64 ;};func _gad (_gdg string ,_fbe []replacement ,_aca *_g .PdfFont )[]_cc .PdfObject {_dec :=[]_cc .PdfObject {};_efg :=0;_fed :=_gdg ;for _daa ,_cbb :=range _fbe {_gdc :=_cbb ._ge ;
_fdf :=_cbb ._cfb ;_gddd :=_cbb ._d ;_bbc :=_cc .MakeFloat (_fdf );if _efg > _gdc ||_gdc ==-1{continue ;};_eaca :=_gdg [_efg :_gdc ];_ggb :=_caa (_eaca ,_aca );_ggc :=_cc .MakeStringFromBytes (_ggb );_dec =append (_dec ,_ggc );_dec =append (_dec ,_bbc );
_ccc :=_gdc +len (_gddd );_fed =_gdg [_ccc :];_efg =_ccc ;if _daa ==len (_fbe )-1{_ggb =_caa (_fed ,_aca );_ggc =_cc .MakeStringFromBytes (_ggb );_dec =append (_dec ,_ggc );};};return _dec ;};func _ggbg (_efge *_ef .TextMarkArray )(float64 ,error ){_gcg ,_gaf :=_efge .BBox ();
if !_gaf {return 0.0,_f .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_efge );
};_becb :=_gfb (_efge );_fdg :=0.0;_ ,_dbf :=_eeg (_efge );_bcab :=_efge .Elements ()[_dbf ];_fef :=_bcab .Font ;if _becb > 0{_fdg =_df (_fef ,_bcab );};_bfd :=(_gcg .Urx -_gcg .Llx );_bfd =_bfd +_fdg *float64 (_becb );Tj :=_fcf (_bfd ,_bcab .FontSize ,_bcab .Th );
return Tj ,nil ;};func _afge (_bge []*matchedIndex ,_gfdd [][]int )[]*matchedIndex {_fbca :=[]*matchedIndex {};for _ ,_cde :=range _bge {_efbf ,_fgef :=_afee (_cde ,_gfdd );if _efbf {_cagc :=_feg (_cde ,_fgef );_fbca =append (_fbca ,_cagc ...);}else {_fbca =append (_fbca ,_cde );
};};return _fbca ;};