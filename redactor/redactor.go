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

package redactor ;import (_e "errors";_g "fmt";_c "github.com/unidoc/unipdf/v3/common";_b "github.com/unidoc/unipdf/v3/contentstream";_cg "github.com/unidoc/unipdf/v3/core";_feg "github.com/unidoc/unipdf/v3/creator";_gdd "github.com/unidoc/unipdf/v3/extractor";
_fef "github.com/unidoc/unipdf/v3/model";_fe "io";_ge "regexp";_gd "sort";_a "strings";);func _fcf (_gddc string ,_dbf *_fef .PdfFont )[]byte {_ebad ,_cc :=_dbf .StringToCharcodeBytes (_gddc );if _cc !=0{_c .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0073\u006fm\u0065\u0020\u0072un\u0065\u0073\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006e\u0063\u006f\u0064\u0065d\u002e\u000a\u0009\u0025\u0073\u0020\u002d\u003e \u0025\u0076",_gddc ,_ebad );
};return _ebad ;};func _edgc (_gbec *_b .ContentStreamOperations ,PdfObj _cg .PdfObject )(*_b .ContentStreamOperation ,int ,bool ){for _afgf ,_gggc :=range *_gbec {_eee :=_gggc .Operand ;if _eee =="\u0054\u006a"{_fbf :=_cg .TraceToDirectObject (_gggc .Params [0]);
if _fbf ==PdfObj {return _gggc ,_afgf ,true ;};}else if _eee =="\u0054\u004a"{_abga ,_caea :=_cg .GetArray (_gggc .Params [0]);if !_caea {return nil ,_afgf ,_caea ;};for _ ,_abbb :=range _abga .Elements (){if _abbb ==PdfObj {return _gggc ,_afgf ,true ;
};};}else if _eee =="\u0022"{_cfd :=_cg .TraceToDirectObject (_gggc .Params [2]);if _cfd ==PdfObj {return _gggc ,_afgf ,true ;};}else if _eee =="\u0027"{_bdef :=_cg .TraceToDirectObject (_gggc .Params [0]);if _bdef ==PdfObj {return _gggc ,_afgf ,true ;
};};};return nil ,-1,false ;};func _bgf (_bac *_b .ContentStreamOperation ,_fdc _cg .PdfObject ,_bag []localSpanMarks )error {var _fefe *_cg .PdfObjectArray ;_cgg ,_abb :=_gcf (_bag );if len (_abb )==1{_ebe :=_abb [0];_dbd :=_cgg [_ebe ];if len (_dbd )==1{_feec :=_dbd [0];
_bc :=_feec ._aedc ;_ega :=_aae (_bc );_bfe ,_gee :=_egc (_fdc ,_ega );if _gee !=nil {return _gee ;};_gac ,_gee :=_eff (_feec ,_bc ,_ega ,_bfe ,_ebe );if _gee !=nil {return _gee ;};_fefe =_cg .MakeArray (_gac ...);}else {_cgb :=_dbd [0]._aedc ;_ecf :=_aae (_cgb );
_ccf ,_dcb :=_egc (_fdc ,_ecf );if _dcb !=nil {return _dcb ;};_faa ,_dcb :=_cgd (_ccf ,_dbd );if _dcb !=nil {return _dcb ;};_ddab :=_bgb (_faa );_fed :=_agde (_ccf ,_ddab ,_ecf );_fefe =_cg .MakeArray (_fed ...);};}else if len (_abb )> 1{_ac :=_bag [0];
_fb :=_ac ._aedc ;_ ,_dbe :=_beg (_fb );_bfg :=_fb .Elements ()[_dbe ];_bafb :=_bfg .Font ;_cge ,_ceda :=_egc (_fdc ,_bafb );if _ceda !=nil {return _ceda ;};_bge ,_ceda :=_cgd (_cge ,_bag );if _ceda !=nil {return _ceda ;};_gec :=_bgb (_bge );_dffe :=_agde (_cge ,_gec ,_bafb );
_fefe =_cg .MakeArray (_dffe ...);};_bac .Params [0]=_fefe ;_bac .Operand ="\u0054\u004a";return nil ;};func _cgd (_bcc string ,_gba []localSpanMarks )([]placeHolders ,error ){_gad :="";_bgfb :=[]placeHolders {};for _dgfb ,_gcba :=range _gba {_bcd :=_gcba ._aedc ;
_dgcf :=_gcba ._gdda ;_acb :=_gefc (_bcd );_edg ,_cae :=_edd (_bcd );if _cae !=nil {return nil ,_cae ;};if _acb !=_gad {var _dfb []int ;if _dgfb ==0&&_dgcf !=_acb {_efd :=_a .Index (_bcc ,_acb );_dfb =[]int {_efd };}else if _dgfb ==len (_gba )-1{_gacf :=_a .LastIndex (_bcc ,_acb );
_dfb =[]int {_gacf };}else {_dfb =_aabg (_bcc ,_acb );};_bcb :=placeHolders {_ag :_dfb ,_cd :_acb ,_fa :_edg };_bgfb =append (_bgfb ,_bcb );};_gad =_acb ;};return _bgfb ,nil ;};func _aabg (_faaa ,_fge string )[]int {if len (_fge )==0{return nil ;};var _acg []int ;
for _ggbd :=0;_ggbd < len (_faaa );{_eag :=_a .Index (_faaa [_ggbd :],_fge );if _eag < 0{return _acg ;};_acg =append (_acg ,_ggbd +_eag );_ggbd +=_eag +len (_fge );};return _acg ;};func _cf (_ba *_b .ContentStreamOperations ,_ga map[_cg .PdfObject ][]localSpanMarks )error {for _gg ,_cb :=range _ga {if _gg ==nil {continue ;
};_ed ,_bf ,_gf :=_edgc (_ba ,_gg );if !_gf {_c .Log .Debug ("Pd\u0066\u004fb\u006a\u0065\u0063\u0074\u0020\u0025\u0073\u006e\u006ft\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0069\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073\u0074r\u0065a\u006d\u0020\u006f\u0070\u0065\u0072\u0061\u0074i\u006fn\u0020\u0025s",_gg ,_ba );
return nil ;};if _ed .Operand =="\u0054\u006a"{_af :=_bgf (_ed ,_gg ,_cb );if _af !=nil {return _af ;};}else if _ed .Operand =="\u0054\u004a"{_fee :=_ced (_ed ,_gg ,_cb );if _fee !=nil {return _fee ;};}else if _ed .Operand =="\u0027"||_ed .Operand =="\u0022"{_fc :=_ca (_ba ,_ed .Operand ,_bf );
if _fc !=nil {return _fc ;};_fc =_bgf (_ed ,_gg ,_cb );if _fc !=nil {return _fc ;};};};return nil ;};type placeHolders struct{_ag []int ;_cd string ;_fa float64 ;};func _bgb (_fegef []placeHolders )[]replacement {_cdg :=[]replacement {};for _ ,_bcba :=range _fegef {_abd :=_bcba ._ag ;
_fec :=_bcba ._cd ;_ffff :=_bcba ._fa ;for _ ,_bce :=range _abd {_fbg :=replacement {_ce :_fec ,_d :_ffff ,_bg :_bce };_cdg =append (_cdg ,_fbg );};};_gd .Slice (_cdg ,func (_gegb ,_abg int )bool {return _cdg [_gegb ]._bg < _cdg [_abg ]._bg });return _cdg ;
};func _edec (_adeg *_gdd .TextMarkArray ,_dfe int ,_effc int )int {_bcedg :=_adeg .Elements ();_gbbf :=_dfe -1;_dce :=_dfe +1;_eace :=-1;if _gbbf >=0{_ccbg :=_bcedg [_gbbf ];_aeg :=_ccbg .ObjString ;_ecfb :=len (_aeg );_gbc :=_ccbg .Index ;if _gbc +1< _ecfb {_eace =_gbbf ;
return _eace ;};};if _dce < len (_bcedg ){_cgef :=_bcedg [_dce ];_bgcbc :=_cgef .ObjString ;if _bgcbc [0]!=_cgef .Text {_eace =_dce ;return _eace ;};};return _eace ;};func _dbbc (_fgd *_gdd .TextMarkArray )[]*_gdd .TextMarkArray {_cdc :=_fgd .Elements ();
_bbae :=len (_cdc );var _adaf _cg .PdfObject ;_abbc :=[]*_gdd .TextMarkArray {};_gge :=&_gdd .TextMarkArray {};_gaab :=-1;for _dage ,_ece :=range _cdc {_adfc :=_ece .DirectObject ;_gaab =_ece .Index ;if _adfc ==nil {_bfee :=_edec (_fgd ,_dage ,_gaab );
if _adaf !=nil {if _bfee ==-1||_bfee > _dage {_abbc =append (_abbc ,_gge );_gge =&_gdd .TextMarkArray {};};};}else if _adfc !=nil &&_adaf ==nil {if _gaab ==0&&_dage > 0{_abbc =append (_abbc ,_gge );_gge =&_gdd .TextMarkArray {};};}else if _adfc !=nil &&_adaf !=nil {if _adfc !=_adaf {_abbc =append (_abbc ,_gge );
_gge =&_gdd .TextMarkArray {};};};_adaf =_adfc ;_gge .Append (_ece );if _dage ==(_bbae -1){_abbc =append (_abbc ,_gge );};};return _abbc ;};type matchedBBox struct{_ceaa _fef .PdfRectangle ;_cba string ;};

// RedactRectanglePropsNew return a new pointer to a default RectangleProps object.
func RedactRectanglePropsNew ()*RectangleProps {return &RectangleProps {FillColor :_feg .ColorBlack ,BorderWidth :0.0,FillOpacity :1.0};};func _agde (_dfc string ,_ddc []replacement ,_gfg *_fef .PdfFont )[]_cg .PdfObject {_bgc :=[]_cg .PdfObject {};_ggg :=0;
_fgg :=_dfc ;for _bfb ,_fba :=range _ddc {_gfa :=_fba ._bg ;_fgbc :=_fba ._d ;_ea :=_fba ._ce ;_cgeg :=_cg .MakeFloat (_fgbc );_gce :=_dfc [_ggg :_gfa ];_cffg :=_fcf (_gce ,_gfg );_cbd :=_cg .MakeStringFromBytes (_cffg );_bgc =append (_bgc ,_cbd );_bgc =append (_bgc ,_cgeg );
_fcc :=_gfa +len (_ea );_fgg =_dfc [_fcc :];_ggg =_fcc ;if _bfb ==len (_ddc )-1{_cffg =_fcf (_fgg ,_gfg );_cbd =_cg .MakeStringFromBytes (_cffg );_bgc =append (_bgc ,_cbd );};};return _bgc ;};func _gefc (_ccb *_gdd .TextMarkArray )string {_cea :="";for _ ,_gaca :=range _ccb .Elements (){_cea +=_gaca .Text ;
};return _cea ;};func (_dcge *regexMatcher )match (_aef string )([]*matchedIndex ,error ){_fgf :=_dcge ._cdcd .Pattern ;if _fgf ==nil {return nil ,_e .New ("\u006e\u006f\u0020\u0070at\u0074\u0065\u0072\u006e\u0020\u0063\u006f\u006d\u0070\u0069\u006c\u0065\u0064");
};var (_gcab =_fgf .FindAllStringIndex (_aef ,-1);_aagb []*matchedIndex ;);for _ ,_acdb :=range _gcab {_aagb =append (_aagb ,&matchedIndex {_dbc :_acdb [0],_ada :_acdb [1],_fffa :_aef [_acdb [0]:_acdb [1]]});};return _aagb ,nil ;};

// New instantiates a Redactor object with given PdfReader and `regex` pattern.
func New (reader *_fef .PdfReader ,opts *RedactionOptions ,rectProps *RectangleProps )*Redactor {if rectProps ==nil {rectProps =RedactRectanglePropsNew ();};return &Redactor {_ggc :reader ,_fdb :opts ,_cbaf :_feg .New (),_efcd :rectProps };};type regexMatcher struct{_cdcd RedactionTerm };
type localSpanMarks struct{_aedc *_gdd .TextMarkArray ;_cbda int ;_gdda string ;};func _geeb (_bdc []int ,_eagd *_gdd .TextMarkArray ,_cag string )(*_gdd .TextMarkArray ,matchedBBox ,error ){_bbc :=matchedBBox {};_fcb :=_bdc [0];_age :=_bdc [1];_cadf :=len (_cag )-len (_a .TrimLeft (_cag ,"\u0020"));
_dfg :=len (_cag )-len (_a .TrimRight (_cag ,"\u0020\u000a"));_fcb =_fcb +_cadf ;_age =_age -_dfg ;_cag =_a .Trim (_cag ,"\u0020\u000a");_bba ,_gaf :=_eagd .RangeOffset (_fcb ,_age );if _gaf !=nil {return nil ,_bbc ,_gaf ;};_ccc ,_cdga :=_bba .BBox ();
if !_cdga {return nil ,_bbc ,_g .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_bba );
};_bbc =matchedBBox {_cba :_cag ,_ceaa :_ccc };return _bba ,_bbc ,nil ;};func (_afb *Redactor )redactPage (_gbad *_b .ContentStreamOperations ,_gcbg *_fef .PdfPageResources )([]matchedBBox ,*_b .ContentStreamOperations ,error ){_gca ,_dggf :=_gdd .NewFromContents (_gbad .String (),_gcbg );
if _dggf !=nil {return nil ,nil ,_dggf ;};_cfg ,_ ,_ ,_dggf :=_gca .ExtractPageText ();_gbad =_cfg .GetContentStreamOps ();if _dggf !=nil {return nil ,nil ,_dggf ;};_bced :=_cfg .Marks ();_cfa :=_cfg .Text ();_baab :=[]matchedBBox {};_begd :=make (map[_cg .PdfObject ][]localSpanMarks );
for _ ,_ccfg :=range _afb ._fdb .Terms {_eae ,_cefg :=_cedd (_ccfg );if _cefg !=nil {return nil ,nil ,_cefg ;};_dgdb ,_cefg :=_eae .match (_cfa );if _cefg !=nil {return nil ,nil ,_cefg ;};_egdb :=_cbfa (_dgdb );for _cfe ,_dfbf :=range _egdb {_agcd :=[]matchedBBox {};
for _ ,_cacg :=range _dfbf {_dag ,_eac ,_fefd :=_geeb (_cacg ,_bced ,_cfe );if _fefd !=nil {return nil ,nil ,_fefd ;};_gbf :=_dbbc (_dag );for _faag ,_gedb :=range _gbf {_acba :=localSpanMarks {_aedc :_gedb ,_cbda :_faag ,_gdda :_cfe };_egcb ,_ :=_beg (_gedb );
if _bgcb ,_eef :=_begd [_egcb ];_eef {_begd [_egcb ]=append (_bgcb ,_acba );}else {_begd [_egcb ]=[]localSpanMarks {_acba };};};_agcd =append (_agcd ,_eac );};_baab =append (_baab ,_agcd ...);};};_dggf =_cf (_gbad ,_begd );if _dggf !=nil {return nil ,nil ,_dggf ;
};return _baab ,_gbad ,nil ;};

// Write writes the content of `re.creator` to writer of type io.Writer interface.
func (_eda *Redactor )Write (writer _fe .Writer )error {return _eda ._cbaf .Write (writer )};func _cedd (_ddbg RedactionTerm )(*regexMatcher ,error ){return &regexMatcher {_cdcd :_ddbg },nil };

// RedactionTerm holds the regexp pattern and the replacement string for the redaction process.
type RedactionTerm struct{Pattern *_ge .Regexp ;};func _ef (_caa *_fef .PdfFont ,_ec _gdd .TextMark )float64 {_egd :=0.001;_ebc :=_ec .Th /100;if _caa .Subtype ()=="\u0054\u0079\u0070e\u0033"{_egd =1;};_gab ,_fgbf :=_caa .GetRuneMetrics (' ');if !_fgbf {_gab ,_fgbf =_caa .GetCharMetrics (32);
};if !_fgbf {_gab ,_ =_fef .DefaultFont ().GetRuneMetrics (' ');};_dda :=_egd *((_gab .Wx *_ec .FontSize +_ec .Tc +_ec .Tw )/_ebc );return _dda ;};func _edd (_gfga *_gdd .TextMarkArray )(float64 ,error ){_eed ,_cac :=_gfga .BBox ();if !_cac {return 0.0,_g .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_gfga );
};_aga :=_baag (_gfga );_egba :=0.0;_ ,_efda :=_beg (_gfga );_cggg :=_gfga .Elements ()[_efda ];_bfba :=_cggg .Font ;if _aga > 0{_egba =_ef (_bfba ,_cggg );};_dbg :=(_eed .Urx -_eed .Llx );_dbg =_dbg +_egba *float64 (_aga );Tj :=_beff (_dbg ,_cggg .FontSize ,_cggg .Th );
return Tj ,nil ;};func _beg (_bed *_gdd .TextMarkArray )(_cg .PdfObject ,int ){var _cbf _cg .PdfObject ;_bbe :=-1;for _cbb ,_dd :=range _bed .Elements (){_cbf =_dd .DirectObject ;_bbe =_cbb ;if _cbf !=nil {break ;};};return _cbf ,_bbe ;};

// Redactor represtents a Redactor object.
type Redactor struct{_ggc *_fef .PdfReader ;_fdb *RedactionOptions ;_cbaf *_feg .Creator ;_efcd *RectangleProps ;};func _ca (_fg *_b .ContentStreamOperations ,_baa string ,_bab int )error {_afa :=_b .ContentStreamOperations {};var _ff _b .ContentStreamOperation ;
for _be ,_gb :=range *_fg {if _be ==_bab {if _baa =="\u0027"{_cda :=_b .ContentStreamOperation {Operand :"\u0054\u002a"};_afa =append (_afa ,&_cda );_ff .Params =_gb .Params ;_ff .Operand ="\u0054\u006a";_afa =append (_afa ,&_ff );}else if _baa =="\u0022"{_dg :=_gb .Params [:2];
Tc ,Tw :=_dg [0],_dg [1];_bd :=_b .ContentStreamOperation {Params :[]_cg .PdfObject {Tc },Operand :"\u0054\u0063"};_afa =append (_afa ,&_bd );_bd =_b .ContentStreamOperation {Params :[]_cg .PdfObject {Tw },Operand :"\u0054\u0077"};_afa =append (_afa ,&_bd );
_ff .Params =[]_cg .PdfObject {_gb .Params [2]};_ff .Operand ="\u0054\u006a";_afa =append (_afa ,&_ff );};};_afa =append (_afa ,_gb );};*_fg =_afa ;return nil ;};

// RedactionOptions is a collection of RedactionTerm objects.
type RedactionOptions struct{Terms []RedactionTerm ;};

// WriteToFile writes the redacted document to file specified by `outputPath`.
func (_edee *Redactor )WriteToFile (outputPath string )error {if _cab :=_edee ._cbaf .WriteToFile (outputPath );_cab !=nil {return _g .Errorf ("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0077\u0072\u0069\u0074\u0065\u0020t\u0068\u0065\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0066\u0069\u006c\u0065");
};return nil ;};func _beff (_gga ,_ffd ,_acd float64 )float64 {_acd =_acd /100;_afag :=(-1000*_gga )/(_ffd *_acd );return _afag ;};type matchedIndex struct{_dbc int ;_ada int ;_fffa string ;};func _aae (_baf *_gdd .TextMarkArray )*_fef .PdfFont {_ ,_fefa :=_beg (_baf );
_befa :=_baf .Elements ()[_fefa ];_gde :=_befa .Font ;return _gde ;};func _ced (_gc *_b .ContentStreamOperation ,_bde _cg .PdfObject ,_dgc []localSpanMarks )error {_fege ,_eg :=_cg .GetArray (_gc .Params [0]);_ad :=[]_cg .PdfObject {};if !_eg {_c .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064",_gc );
return _g .Errorf ("\u0073\u0070\u0061\u006e\u004d\u0061\u0072\u006bs\u002e\u0042\u0042ox\u0020\u0068\u0061\u0073\u0020\u006eo\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062\u006f\u0078\u002e\u0020s\u0070\u0061\u006e\u004d\u0061\u0072\u006b\u0073=\u0025\u0073",_gc );
};_dff ,_aa :=_gcf (_dgc );if len (_aa )==1{_db :=_aa [0];_agg :=_dff [_db ];if len (_agg )==1{_fgb :=_agg [0];_fca :=_fgb ._aedc ;_dc :=_aae (_fca );_ffe ,_cbc :=_egc (_bde ,_dc );if _cbc !=nil {return _cbc ;};_cbcd ,_cbc :=_eff (_fgb ,_fca ,_dc ,_ffe ,_db );
if _cbc !=nil {return _cbc ;};for _ ,_geg :=range _fege .Elements (){if _geg ==_bde {_ad =append (_ad ,_cbcd ...);}else {_ad =append (_ad ,_geg );};};}else {_gbe :=_agg [0]._aedc ;_egf :=_aae (_gbe );_cff ,_gef :=_egc (_bde ,_egf );if _gef !=nil {return _gef ;
};_bdf ,_gef :=_cgd (_cff ,_agg );if _gef !=nil {return _gef ;};_ee :=_bgb (_bdf );_ae :=_agde (_cff ,_ee ,_egf );for _ ,_eba :=range _fege .Elements (){if _eba ==_bde {_ad =append (_ad ,_ae ...);}else {_ad =append (_ad ,_eba );};};};_gc .Params [0]=_cg .MakeArray (_ad ...);
}else if len (_aa )> 1{_de :=_dgc [0];_fd :=_de ._aedc ;_ ,_bef :=_beg (_fd );_gbb :=_fd .Elements ()[_bef ];_fega :=_gbb .Font ;_ab ,_aab :=_egc (_bde ,_fega );if _aab !=nil {return _aab ;};_dgg ,_aab :=_cgd (_ab ,_dgc );if _aab !=nil {return _aab ;};
_bff :=_bgb (_dgg );_fcd :=_agde (_ab ,_bff ,_fega );for _ ,_dfa :=range _fege .Elements (){if _dfa ==_bde {_ad =append (_ad ,_fcd ...);}else {_ad =append (_ad ,_dfa );};};_gc .Params [0]=_cg .MakeArray (_ad ...);};return nil ;};

// Redact executes the redact operation on a pdf file and updates the content streams of all pages of the file.
func (_cca *Redactor )Redact ()error {_aag ,_daf :=_cca ._ggc .GetNumPages ();if _daf !=nil {return _g .Errorf ("\u0066\u0061\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0074\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020P\u0061\u0067\u0065\u0073");
};_fcdbb :=_cca ._efcd .FillColor ;_cga :=_cca ._efcd .BorderWidth ;_acdc :=_cca ._efcd .FillOpacity ;for _dca :=1;_dca <=_aag ;_dca ++{_add ,_cfff :=_cca ._ggc .GetPage (_dca );if _cfff !=nil {return _cfff ;};_dfbc ,_cfff :=_gdd .New (_add );if _cfff !=nil {return _cfff ;
};_bfbb ,_ ,_ ,_cfff :=_dfbc .ExtractPageText ();if _cfff !=nil {return _cfff ;};_aedcc :=_bfbb .GetContentStreamOps ();_gdfg ,_dadb ,_cfff :=_cca .redactPage (_aedcc ,_add .Resources );if _dadb ==nil {_c .Log .Info ("N\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u0020t\u0068\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065d \u0070\u0061\u0074t\u0061r\u006e\u002e");
_dadb =_aedcc ;};_add .SetContentStreams ([]string {_dadb .String ()},_cg .NewFlateEncoder ());if _cfff !=nil {return _cfff ;};_ddb ,_cfff :=_add .GetMediaBox ();if _cfff !=nil {return _cfff ;};if _add .MediaBox ==nil {_add .MediaBox =_ddb ;};if _cad :=_cca ._cbaf .AddPage (_add );
_cad !=nil {return _cad ;};_ded :=_ddb .Ury ;for _ ,_ffda :=range _gdfg {_ade :=_ffda ._ceaa ;_bgca :=_cca ._cbaf .NewRectangle (_ade .Llx ,_ded -_ade .Lly ,_ade .Urx -_ade .Llx ,-(_ade .Ury -_ade .Lly ));_bgca .SetFillColor (_fcdbb );_bgca .SetBorderWidth (_cga );
_bgca .SetFillOpacity (_acdc );if _bae :=_cca ._cbaf .Draw (_bgca );_bae !=nil {return nil ;};};};_cca ._cbaf .SetOutlineTree (_cca ._ggc .GetOutlineTree ());return nil ;};func _egc (_cec _cg .PdfObject ,_dgfd *_fef .PdfFont )(string ,error ){_aff ,_edc :=_cg .GetStringBytes (_cec );
if !_edc {return "",_cg .ErrTypeError ;};_dbb :=_dgfd .BytesToCharcodes (_aff );_fcdb ,_edb ,_abbd :=_dgfd .CharcodesToStrings (_dbb );if _abbd > 0{_c .Log .Debug ("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064",_edb ,_abbd );
};_dae :=_a .Join (_fcdb ,"");return _dae ,nil ;};func _cbfa (_ecfbg []*matchedIndex )map[string ][][]int {_fcaf :=make (map[string ][][]int );for _ ,_aaf :=range _ecfbg {_adad :=_aaf ._fffa ;_gafg :=[]int {_aaf ._dbc ,_aaf ._ada };if _cdcc ,_ebd :=_fcaf [_adad ];
_ebd {_fcaf [_adad ]=append (_cdcc ,_gafg );}else {_fcaf [_adad ]=[][]int {_gafg };};};return _fcaf ;};func _eff (_afd localSpanMarks ,_dfd *_gdd .TextMarkArray ,_fff *_fef .PdfFont ,_cef ,_da string )([]_cg .PdfObject ,error ){_aed :=_gefc (_dfd );Tj ,_dggc :=_edd (_dfd );
if _dggc !=nil {return nil ,_dggc ;};_bca :=len (_cef );_gcb :=len (_aed );_ged :=-1;_ceg :=_cg .MakeFloat (Tj );if _aed !=_da {_cdaa :=_afd ._cbda ;if _cdaa ==0{_ged =_a .LastIndex (_cef ,_aed );}else {_ged =_a .Index (_cef ,_aed );};}else {_ged =_a .Index (_cef ,_aed );
};_egb :=_ged +_gcb ;_afg :=[]_cg .PdfObject {};if _ged ==0&&_egb ==_bca {_afg =append (_afg ,_ceg );}else if _ged ==0&&_egb < _bca {_ecg :=_fcf (_cef [_egb :],_fff );_fedf :=_cg .MakeStringFromBytes (_ecg );_afg =append (_afg ,_ceg ,_fedf );}else if _ged > 0&&_egb >=_bca {_bee :=_fcf (_cef [:_ged ],_fff );
_gdec :=_cg .MakeStringFromBytes (_bee );_afg =append (_afg ,_gdec ,_ceg );}else if _ged > 0&&_egb < _bca {_fgbg :=_fcf (_cef [:_ged ],_fff );_cfc :=_fcf (_cef [_egb :],_fff );_gdf :=_cg .MakeStringFromBytes (_fgbg );_dad :=_cg .MakeString (string (_cfc ));
_afg =append (_afg ,_gdf ,_ceg ,_dad );};return _afg ,nil ;};

// RectangleProps defines properties of the redaction rectangle to be drawn.
type RectangleProps struct{FillColor _feg .Color ;BorderWidth float64 ;FillOpacity float64 ;};func _gcf (_fac []localSpanMarks )(map[string ][]localSpanMarks ,[]string ){_dgd :=make (map[string ][]localSpanMarks );_ggb :=[]string {};for _ ,_dcg :=range _fac {_ecb :=_dcg ._gdda ;
if _adf ,_eec :=_dgd [_ecb ];_eec {_dgd [_ecb ]=append (_adf ,_dcg );}else {_dgd [_ecb ]=[]localSpanMarks {_dcg };_ggb =append (_ggb ,_ecb );};};return _dgd ,_ggb ;};type replacement struct{_ce string ;_d float64 ;_bg int ;};func _baag (_bb *_gdd .TextMarkArray )int {_df :=0;
_dgf :=_bb .Elements ();if _dgf [0].Text =="\u0020"{_df ++;};if _dgf [_bb .Len ()-1].Text =="\u0020"{_df ++;};return _df ;};