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

package svg ;import (_bbf "encoding/xml";_dc "fmt";_bd "github.com/unidoc/unipdf/v3/common";_ag "github.com/unidoc/unipdf/v3/contentstream";_dg "github.com/unidoc/unipdf/v3/contentstream/draw";_bc "github.com/unidoc/unipdf/v3/internal/graphic2d";_fg "golang.org/x/net/html/charset";
_bb "io";_bf "math";_a "os";_c "strconv";_d "strings";_b "unicode";);func _cde (_eabb []*Command )*Path {_bfb :=&Path {};var _dfdf []*Command ;for _edb ,_dgf :=range _eabb {switch _d .ToLower (_dgf .Symbol ){case _edd ._dagg :if len (_dfdf )> 0{_bfb .Subpaths =append (_bfb .Subpaths ,&Subpath {_dfdf });
};_dfdf =[]*Command {_dgf };case _edd ._ddg :_dfdf =append (_dfdf ,_dgf );_bfb .Subpaths =append (_bfb .Subpaths ,&Subpath {_dfdf });_dfdf =[]*Command {};default:_dfdf =append (_dfdf ,_dgf );if len (_eabb )==_edb +1{_bfb .Subpaths =append (_bfb .Subpaths ,&Subpath {_dfdf });
};};};return _bfb ;};func ParseFromStream (source _bb .Reader )(*GraphicSVG ,error ){_fabg :=_bbf .NewDecoder (source );_fabg .CharsetReader =_fg .NewReaderLabel ;_feee ,_ee :=_cgda (_fabg );if _ee !=nil {return nil ,_ee ;};if _dgad :=_feee .Decode (_fabg );
_dgad !=nil &&_dgad !=_bb .EOF {return nil ,_dgad ;};return _feee ,nil ;};func (_efa *Command )isAbsolute ()bool {return _efa .Symbol ==_d .ToUpper (_efa .Symbol )};func (_aaee *GraphicSVGStyle )toContentStream (_fcg *_ag .ContentCreator ){if _aaee ==nil {return ;
};if _aaee .FillColor !=""{var _gca ,_bce ,_bed float64 ;if _agcb ,_gcfg :=_bc .ColorMap [_aaee .FillColor ];_gcfg {_fgd ,_cfab ,_egc ,_ :=_agcb .RGBA ();_gca ,_bce ,_bed =float64 (_fgd ),float64 (_cfab ),float64 (_egc );}else {_gca ,_bce ,_bed =_eeb (_aaee .FillColor );
};_fcg .Add_rg (_gca ,_bce ,_bed );};if _aaee .StrokeColor !=""{var _ae ,_fcf ,_agge float64 ;if _eca ,_ege :=_bc .ColorMap [_aaee .StrokeColor ];_ege {_gac ,_ffe ,_acdd ,_ :=_eca .RGBA ();_ae ,_fcf ,_agge =float64 (_gac )/255.0,float64 (_ffe )/255.0,float64 (_acdd )/255.0;
}else {_ae ,_fcf ,_agge =_eeb (_aaee .StrokeColor );};_fcg .Add_RG (_ae ,_fcf ,_agge );};if _aaee .StrokeWidth > 0{_fcg .Add_w (_aaee .StrokeWidth );};};func _gae (_ccaf string ,_dgcc int )(float64 ,error ){_abd ,_ecd :=_fgg (_ccaf );_egeb ,_ccbg :=_c .ParseFloat (_abd ,_dgcc );
if _ccbg !=nil {return 0,_ccbg ;};if _cgdae ,_gbgb :=_e [_ecd ];_gbgb {_egeb =_egeb *_cgdae ;}else {_egeb =_egeb *_dcb ;};return _egeb ,nil ;};type GraphicSVG struct{ViewBox struct{X ,Y ,W ,H float64 ;};Name string ;Attributes map[string ]string ;Children []*GraphicSVG ;
Content string ;Style *GraphicSVGStyle ;Width float64 ;Height float64 ;_daa float64 ;};func _adb (_gbc map[string ]string ,_abee float64 )(*GraphicSVGStyle ,error ){_cbbf :=_cbb ();_fbf ,_dge :=_gbc ["\u0066\u0069\u006c\u006c"];if _dge {_cbbf .FillColor =_fbf ;
if _fbf =="\u006e\u006f\u006e\u0065"{_cbbf .FillColor ="";};};_feca ,_gadd :=_gbc ["\u0073\u0074\u0072\u006f\u006b\u0065"];if _gadd {_cbbf .StrokeColor =_feca ;if _feca =="\u006e\u006f\u006e\u0065"{_cbbf .StrokeColor ="";};};_dcbg ,_gcbb :=_gbc ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];
if _gcbb {_ced ,_bdfa :=_gae (_dcbg ,64);if _bdfa !=nil {return nil ,_bdfa ;};_cbbf .StrokeWidth =_ced *_abee ;};return _cbbf ,nil ;};var (_fc =[]string {"\u0063\u006d","\u006d\u006d","\u0070\u0078","\u0070\u0074"};_e =map[string ]float64 {"\u0063\u006d":_aa ,"\u006d\u006d":_de ,"\u0070\u0078":_dcb ,"\u0070\u0074":1};
);func _bbgb (_be *GraphicSVG ,_gfg *_ag .ContentCreator ){_gfg .Add_q ();_be .Style .toContentStream (_gfg );_ccc ,_agb :=_gae (_be .Attributes ["\u0078\u0031"],64);if _agb !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_agb .Error ());
};_ggc ,_agb :=_gae (_be .Attributes ["\u0079\u0031"],64);if _agb !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_agb .Error ());
};_ebf ,_agb :=_gae (_be .Attributes ["\u0078\u0032"],64);if _agb !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_agb .Error ());
};_ba ,_agb :=_gae (_be .Attributes ["\u0079\u0032"],64);if _agb !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_agb .Error ());
};_gfg .Add_m (_ccc *_be ._daa ,_ggc *_be ._daa );_gfg .Add_l (_ebf *_be ._daa ,_ba *_be ._daa );if _be .Style .FillColor !=""&&_be .Style .StrokeColor !=""{_gfg .Add_B ();}else if _be .Style .FillColor !=""{_gfg .Add_f ();}else if _be .Style .StrokeColor !=""{_gfg .Add_S ();
};_gfg .Add_h ();_gfg .Add_Q ();};func (_ff *GraphicSVG )SetScaling (xFactor ,yFactor float64 ){_bfa :=_ff .Width /_ff .ViewBox .W ;_efe :=_ff .Height /_ff .ViewBox .H ;_ff .setDefaultScaling (_bf .Max (_bfa ,_efe ));for _ ,_fff :=range _ff .Children {_fff .SetScaling (xFactor ,yFactor );
};};func (_aeg *Command )compare (_baa *Command )bool {if _aeg .Symbol !=_baa .Symbol {return false ;};for _geg ,_dfa :=range _aeg .Params {if _dfa !=_baa .Params [_geg ]{return false ;};};return true ;};func _eeb (_fbg string )(_cag ,_ffec ,_ebc float64 ){if (len (_fbg )!=4&&len (_fbg )!=7)||_fbg [0]!='#'{_bd .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_fbg );
return _cag ,_ffec ,_ebc ;};var _cdbb ,_aga ,_dfe int ;if len (_fbg )==4{var _cfd ,_gdbg ,_aagf int ;_dfag ,_bea :=_dc .Sscanf (_fbg ,"\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078",&_cfd ,&_gdbg ,&_aagf );if _bea !=nil {_bd .Log .Debug ("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076",_fbg ,_bea );
return _cag ,_ffec ,_ebc ;};if _dfag !=3{_bd .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_fbg );return _cag ,_ffec ,_ebc ;};_cdbb =_cfd *16+_cfd ;_aga =_gdbg *16+_gdbg ;_dfe =_aagf *16+_aagf ;
}else {_dfc ,_debcd :=_dc .Sscanf (_fbg ,"\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078",&_cdbb ,&_aga ,&_dfe );if _debcd !=nil {_bd .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_fbg );
return _cag ,_ffec ,_ebc ;};if _dfc !=3{_bd .Log .Debug ("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029",_fbg ,_dfc );
return _cag ,_ffec ,_ebc ;};};_edbc :=float64 (_cdbb )/255.0;_dabe :=float64 (_aga )/255.0;_cfdf :=float64 (_dfe )/255.0;return _edbc ,_dabe ,_cfdf ;};type Path struct{Subpaths []*Subpath ;};func (_fda *GraphicSVG )toContentStream (_ge *_ag .ContentCreator ){_fab ,_bdd :=_adb (_fda .Attributes ,_fda ._daa );
if _bdd !=nil {_bd .Log .Debug ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_bdd );};_fda .Style =_fab ;switch _fda .Name {case "\u0070\u0061\u0074\u0068":_ab (_fda ,_ge );
for _ ,_aafg :=range _fda .Children {_aafg .toContentStream (_ge );};case "\u0072\u0065\u0063\u0074":_dag (_fda ,_ge );for _ ,_fgb :=range _fda .Children {_fgb .toContentStream (_ge );};case "\u0063\u0069\u0072\u0063\u006c\u0065":_bbb (_fda ,_ge );for _ ,_eba :=range _fda .Children {_eba .toContentStream (_ge );
};case "\u0065l\u006c\u0069\u0070\u0073\u0065":_fgc (_fda ,_ge );for _ ,_caa :=range _fda .Children {_caa .toContentStream (_ge );};case "\u0070\u006f\u006c\u0079\u006c\u0069\u006e\u0065":_gbb (_fda ,_ge );for _ ,_dec :=range _fda .Children {_dec .toContentStream (_ge );
};case "\u0070o\u006c\u0079\u0067\u006f\u006e":_bdg (_fda ,_ge );for _ ,_ccd :=range _fda .Children {_ccd .toContentStream (_ge );};case "\u006c\u0069\u006e\u0065":_bbgb (_fda ,_ge );for _ ,_bcd :=range _fda .Children {_bcd .toContentStream (_ge );};case "\u0067":_cdd ,_efbg :=_fda .Attributes ["\u0066\u0069\u006c\u006c"];
_bcc ,_agd :=_fda .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"];_cdg ,_dfg :=_fda .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];for _ ,_cgd :=range _fda .Children {if _ ,_fcb :=_cgd .Attributes ["\u0066\u0069\u006c\u006c"];
!_fcb &&_efbg {_cgd .Attributes ["\u0066\u0069\u006c\u006c"]=_cdd ;};if _ ,_ccf :=_cgd .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"];!_ccf &&_agd {_cgd .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"]=_bcc ;};if _ ,_gcfe :=_cgd .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];
!_gcfe &&_dfg {_cgd .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]=_cdg ;};_cgd .toContentStream (_ge );};};};func _ffc (_cgdaea float64 ,_acbd int )float64 {_dfda :=_bf .Pow (10,float64 (_acbd ));return float64 (_dfb (_cgdaea *_dfda ))/_dfda ;
};func _cdb ()commands {var _ccb =map[string ]int {"\u006d":2,"\u007a":0,"\u006c":2,"\u0068":1,"\u0076":1,"\u0063":6,"\u0073":4,"\u0071":4,"\u0074":2,"\u0061":7};var _edg []string ;for _bff :=range _ccb {_edg =append (_edg ,_bff );};return commands {_edg ,_ccb ,"\u006d","\u007a"};
};func _bdde (_ade []float64 )[]float64 {for _gaaf ,_ebfc :=0,len (_ade )-1;_gaaf < _ebfc ;_gaaf ,_ebfc =_gaaf +1,_ebfc -1{_ade [_gaaf ],_ade [_ebfc ]=_ade [_ebfc ],_ade [_gaaf ];};return _ade ;};type Subpath struct{Commands []*Command ;};func _ab (_ea *GraphicSVG ,_deg *_ag .ContentCreator ){_deg .Add_q ();
_ea .Style .toContentStream (_deg );_fe ,_ca :=_cgacg (_ea .Attributes ["\u0064"]);if _ca !=nil {_bd .Log .Error ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025s",_ca .Error ());};var (_bcf ,_cad =0.0,0.0;_g ,_da =0.0,0.0;_gf *Command ;);for _ ,_gd :=range _fe .Subpaths {for _ ,_bca :=range _gd .Commands {switch _d .ToLower (_bca .Symbol ){case "\u006d":_g ,_da =_bca .Params [0]*_ea ._daa ,_bca .Params [1]*_ea ._daa ;
if !_bca .isAbsolute (){_g ,_da =_bcf +_g -_ea .ViewBox .X ,_cad +_da -_ea .ViewBox .Y ;};_deg .Add_m (_ffc (_g ,3),_ffc (_da ,3));_bcf ,_cad =_g ,_da ;case "\u0063":_ed ,_ac ,_fa ,_fee ,_aag ,_cg :=_bca .Params [0]*_ea ._daa ,_bca .Params [1]*_ea ._daa ,_bca .Params [2]*_ea ._daa ,_bca .Params [3]*_ea ._daa ,_bca .Params [4]*_ea ._daa ,_bca .Params [5]*_ea ._daa ;
if !_bca .isAbsolute (){_ed ,_ac ,_fa ,_fee ,_aag ,_cg =_bcf +_ed ,_cad +_ac ,_bcf +_fa ,_cad +_fee ,_bcf +_aag ,_cad +_cg ;};_deg .Add_c (_ffc (_ed ,3),_ffc (_ac ,3),_ffc (_fa ,3),_ffc (_fee ,3),_ffc (_aag ,3),_ffc (_cg ,3));_bcf ,_cad =_aag ,_cg ;case "\u0073":_eg ,_bbg ,_cc ,_abe :=_bca .Params [0]*_ea ._daa ,_bca .Params [1]*_ea ._daa ,_bca .Params [2]*_ea ._daa ,_bca .Params [3]*_ea ._daa ;
if !_bca .isAbsolute (){_eg ,_bbg ,_cc ,_abe =_bcf +_eg ,_cad +_bbg ,_bcf +_cc ,_cad +_abe ;};_deg .Add_c (_ffc (_bcf ,3),_ffc (_cad ,3),_ffc (_eg ,3),_ffc (_bbg ,3),_ffc (_cc ,3),_ffc (_abe ,3));_bcf ,_cad =_cc ,_abe ;case "\u006c":_gfb ,_fge :=_bca .Params [0]*_ea ._daa ,_bca .Params [1]*_ea ._daa ;
if !_bca .isAbsolute (){_gfb ,_fge =_bcf +_gfb ,_cad +_fge ;};_deg .Add_l (_ffc (_gfb ,3),_ffc (_fge ,3));_bcf ,_cad =_gfb ,_fge ;case "\u0068":_gfe :=_bca .Params [0]*_ea ._daa ;if !_bca .isAbsolute (){_gfe =_bcf +_gfe ;};_deg .Add_l (_ffc (_gfe ,3),_ffc (_cad ,3));
_bcf =_gfe ;case "\u0076":_dad :=_bca .Params [0]*_ea ._daa ;if !_bca .isAbsolute (){_dad =_cad +_dad ;};_deg .Add_l (_ffc (_bcf ,3),_ffc (_dad ,3));_cad =_dad ;case "\u0071":_fec ,_abec ,_acd ,_agg :=_bca .Params [0]*_ea ._daa ,_bca .Params [1]*_ea ._daa ,_bca .Params [2]*_ea ._daa ,_bca .Params [3]*_ea ._daa ;
if !_bca .isAbsolute (){_fec ,_abec ,_acd ,_agg =_bcf +_fec ,_cad +_abec ,_bcf +_acd ,_cad +_agg ;};_ec ,_gdc :=_bc .QuadraticToCubicBezier (_bcf ,_cad ,_fec ,_abec ,_acd ,_agg );_deg .Add_c (_ffc (_ec .X ,3),_ffc (_ec .Y ,3),_ffc (_gdc .X ,3),_ffc (_gdc .Y ,3),_ffc (_acd ,3),_ffc (_agg ,3));
_bcf ,_cad =_acd ,_agg ;case "\u0074":var _edf ,_aac _bc .Point ;_ef ,_fd :=_bca .Params [0]*_ea ._daa ,_bca .Params [1]*_ea ._daa ;if !_bca .isAbsolute (){_ef ,_fd =_bcf +_ef ,_cad +_fd ;};if _gf !=nil &&_d .ToLower (_gf .Symbol )=="\u0071"{_ce :=_bc .Point {X :_gf .Params [0]*_ea ._daa ,Y :_gf .Params [1]*_ea ._daa };
_bcae :=_bc .Point {X :_gf .Params [2]*_ea ._daa ,Y :_gf .Params [3]*_ea ._daa };_degd :=_bcae .Mul (2.0).Sub (_ce );_edf ,_aac =_bc .QuadraticToCubicBezier (_bcf ,_cad ,_degd .X ,_degd .Y ,_ef ,_fd );};_deg .Add_c (_ffc (_edf .X ,3),_ffc (_edf .Y ,3),_ffc (_aac .X ,3),_ffc (_aac .Y ,3),_ffc (_ef ,3),_ffc (_fd ,3));
_bcf ,_cad =_ef ,_fd ;case "\u0061":_efb ,_db :=_bca .Params [0]*_ea ._daa ,_bca .Params [1]*_ea ._daa ;_adc :=_bca .Params [2];_cf :=_bca .Params [3]> 0;_eda :=_bca .Params [4]> 0;_aae ,_eaf :=_bca .Params [5]*_ea ._daa ,_bca .Params [6]*_ea ._daa ;if !_bca .isAbsolute (){_aae ,_eaf =_bcf +_aae ,_cad +_eaf ;
};_ga :=_bc .EllipseToCubicBeziers (_bcf ,_cad ,_efb ,_db ,_adc ,_cf ,_eda ,_aae ,_eaf );for _ ,_faa :=range _ga {_deg .Add_c (_ffc (_faa [1].X ,3),_ffc ((_faa [1].Y ),3),_ffc ((_faa [2].X ),3),_ffc ((_faa [2].Y ),3),_ffc ((_faa [3].X ),3),_ffc ((_faa [3].Y ),3));
};_bcf ,_cad =_aae ,_eaf ;case "\u007a":_deg .Add_h ();};_gf =_bca ;};};if _ea .Style .FillColor !=""&&_ea .Style .StrokeColor !=""{_deg .Add_B ();}else if _ea .Style .FillColor !=""{_deg .Add_f ();}else if _ea .Style .StrokeColor !=""{_deg .Add_S ();};
_deg .Add_h ();_deg .Add_Q ();};func _aafa (_baad string )(_bfae []float64 ,_ecba error ){var _eef float64 ;_fdd :=0;_gdf :=true ;for _dafg ,_bag :=range _baad {if _bag =='.'{if _gdf {_gdf =false ;continue ;};_eef ,_ecba =_gae (_baad [_fdd :_dafg ],64);
if _ecba !=nil {return ;};_bfae =append (_bfae ,_eef );_fdd =_dafg ;};};_eef ,_ecba =_gae (_baad [_fdd :],64);if _ecba !=nil {return ;};_bfae =append (_bfae ,_eef );return ;};type token struct{_ecg string ;_aea bool ;};func _cbb ()*GraphicSVGStyle {return &GraphicSVGStyle {FillColor :"\u00230\u0030\u0030\u0030\u0030\u0030",StrokeColor :"",StrokeWidth :0};
};func _cgacg (_fga string )(*Path ,error ){_edd =_cdb ();_cae ,_acb :=_fdc (_ebfa (_fga ));if _acb !=nil {return nil ,_acb ;};return _cde (_cae ),nil ;};func _fgg (_aafag string )(_egbc ,_cdbe string ){if _aafag ==""||(_aafag [len (_aafag )-1]>='0'&&_aafag [len (_aafag )-1]<='9'){return _aafag ,"";
};_egbc =_aafag ;for _ ,_dee :=range _fc {if _d .Contains (_egbc ,_dee ){_cdbe =_dee ;};_egbc =_d .TrimSuffix (_egbc ,_dee );};return ;};func (_dbfg *Subpath )compare (_cfe *Subpath )bool {if len (_dbfg .Commands )!=len (_cfe .Commands ){return false ;
};for _bad ,_cgeb :=range _dbfg .Commands {if !_cgeb .compare (_cfe .Commands [_bad ]){return false ;};};return true ;};func _cgda (_ggb *_bbf .Decoder )(*GraphicSVG ,error ){for {_gga ,_dcf :=_ggb .Token ();if _gga ==nil &&_dcf ==_bb .EOF {break ;};if _dcf !=nil {return nil ,_dcf ;
};switch _bac :=_gga .(type ){case _bbf .StartElement :return _dff (_bac ),nil ;};};return &GraphicSVG {},nil ;};const (_dcb =0.72;_aa =28.3464;_de =_aa /10;_df =0.551784;);func ParseFromString (svgStr string )(*GraphicSVG ,error ){return ParseFromStream (_d .NewReader (svgStr ));
};func _gbcg (_gbe string )([]float64 ,error ){_fgac :=-1;var _ccda []float64 ;_ebe :=' ';for _efbf ,_bfdb :=range _gbe {if !_b .IsNumber (_bfdb )&&_bfdb !='.'&&!(_bfdb =='-'&&_ebe =='e')&&_bfdb !='e'{if _fgac !=-1{_aee ,_debc :=_aafa (_gbe [_fgac :_efbf ]);
if _debc !=nil {return _ccda ,_debc ;};_ccda =append (_ccda ,_aee ...);};if _bfdb =='-'{_fgac =_efbf ;}else {_fgac =-1;};}else if _fgac ==-1{_fgac =_efbf ;};_ebe =_bfdb ;};if _fgac !=-1&&_fgac !=len (_gbe ){_accf ,_afef :=_aafa (_gbe [_fgac :]);if _afef !=nil {return _ccda ,_afef ;
};_ccda =append (_ccda ,_accf ...);};return _ccda ,nil ;};func (_bfd *commands )isCommand (_gag string )bool {for _ ,_faee :=range _bfd ._aff {if _d .ToLower (_gag )==_faee {return true ;};};return false ;};func _cff (_gdcf []token ,_dcec string )([]token ,string ){if _dcec !=""{_gdcf =append (_gdcf ,token {_dcec ,false });
_dcec ="";};return _gdcf ,_dcec ;};func _gbb (_gcf *GraphicSVG ,_dbf *_ag .ContentCreator ){_dbf .Add_q ();_gcf .Style .toContentStream (_dbf );_adg ,_fea :=_gbcg (_gcf .Attributes ["\u0070\u006f\u0069\u006e\u0074\u0073"]);if _fea !=nil {_bd .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076",_fea );
return ;};if len (_adg )%2> 0{_bd .Log .Debug ("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068");return ;
};for _dgc :=0;_dgc < len (_adg );{if _dgc ==0{_dbf .Add_m (_adg [_dgc ]*_gcf ._daa ,_adg [_dgc +1]*_gcf ._daa );}else {_dbf .Add_l (_adg [_dgc ]*_gcf ._daa ,_adg [_dgc +1]*_gcf ._daa );};_dgc +=2;};if _gcf .Style .FillColor !=""&&_gcf .Style .StrokeColor !=""{_dbf .Add_B ();
}else if _gcf .Style .FillColor !=""{_dbf .Add_f ();}else if _gcf .Style .StrokeColor !=""{_dbf .Add_S ();};_dbf .Add_h ();_dbf .Add_Q ();};func _ebfa (_cfec string )[]token {var (_bedg []token ;_cea string ;);for _ ,_eaac :=range _cfec {_cdge :=string (_eaac );
switch {case _edd .isCommand (_cdge ):_bedg ,_cea =_cff (_bedg ,_cea );_bedg =append (_bedg ,token {_cdge ,true });case _cdge =="\u002e":if _cea ==""{_cea ="\u0030";};if _d .Contains (_cea ,_cdge ){_bedg =append (_bedg ,token {_cea ,false });_cea ="\u0030";
};fallthrough;case _cdge >="\u0030"&&_cdge <="\u0039"||_cdge =="\u0065":_cea +=_cdge ;case _cdge =="\u002d":if _d .HasSuffix (_cea ,"\u0065"){_cea +=_cdge ;}else {_bedg ,_ =_cff (_bedg ,_cea );_cea =_cdge ;};default:_bedg ,_cea =_cff (_bedg ,_cea );};};
_bedg ,_ =_cff (_bedg ,_cea );return _bedg ;};type pathParserError struct{_fbb string };func _bdg (_gcfc *GraphicSVG ,_fbd *_ag .ContentCreator ){_fbd .Add_q ();_gcfc .Style .toContentStream (_fbd );_bg ,_ada :=_gbcg (_gcfc .Attributes ["\u0070\u006f\u0069\u006e\u0074\u0073"]);
if _ada !=nil {_bd .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076",_ada );
return ;};if len (_bg )%2> 0{_bd .Log .Debug ("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068");return ;
};for _fae :=0;_fae < len (_bg );{if _fae ==0{_fbd .Add_m (_bg [_fae ]*_gcfc ._daa ,_bg [_fae +1]*_gcfc ._daa );}else {_fbd .Add_l (_bg [_fae ]*_gcfc ._daa ,_bg [_fae +1]*_gcfc ._daa );};_fae +=2;};_fbd .Add_l (_bg [0]*_gcfc ._daa ,_bg [1]*_gcfc ._daa );
if _gcfc .Style .FillColor !=""&&_gcfc .Style .StrokeColor !=""{_fbd .Add_B ();}else if _gcfc .Style .FillColor !=""{_fbd .Add_f ();}else if _gcfc .Style .StrokeColor !=""{_fbd .Add_S ();};_fbd .Add_h ();_fbd .Add_Q ();};type GraphicSVGStyle struct{FillColor string ;
StrokeColor string ;StrokeWidth float64 ;};func _dfb (_ged float64 )int {return int (_ged +_bf .Copysign (0.5,_ged ))};type Command struct{Symbol string ;Params []float64 ;};var _edd commands ;func _bbb (_dce *GraphicSVG ,_eff *_ag .ContentCreator ){_eff .Add_q ();
_dce .Style .toContentStream (_eff );_bfc ,_eb :=_gae (_dce .Attributes ["\u0063\u0078"],64);if _eb !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_eb .Error ());
};_gcc ,_eb :=_gae (_dce .Attributes ["\u0063\u0079"],64);if _eb !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_eb .Error ());
};_aaf ,_eb :=_gae (_dce .Attributes ["\u0072"],64);if _eb !=nil {_bd .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0072\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_eb .Error ());
};_gab :=_aaf *_dce ._daa ;_fb :=_aaf *_dce ._daa ;_gad :=_gab *_df ;_gbf :=_fb *_df ;_bdf :=_dg .NewCubicBezierPath ();_bdf =_bdf .AppendCurve (_dg .NewCubicBezierCurve (-_gab ,0,-_gab ,_gbf ,-_gad ,_fb ,0,_fb ));_bdf =_bdf .AppendCurve (_dg .NewCubicBezierCurve (0,_fb ,_gad ,_fb ,_gab ,_gbf ,_gab ,0));
_bdf =_bdf .AppendCurve (_dg .NewCubicBezierCurve (_gab ,0,_gab ,-_gbf ,_gad ,-_fb ,0,-_fb ));_bdf =_bdf .AppendCurve (_dg .NewCubicBezierCurve (0,-_fb ,-_gad ,-_fb ,-_gab ,-_gbf ,-_gab ,0));_bdf =_bdf .Offset (_bfc *_dce ._daa ,_gcc *_dce ._daa );if _dce .Style .StrokeWidth > 0{_bdf =_bdf .Offset (_dce .Style .StrokeWidth /2,_dce .Style .StrokeWidth /2);
};_dg .DrawBezierPathWithCreator (_bdf ,_eff );if _dce .Style .FillColor !=""&&_dce .Style .StrokeColor !=""{_eff .Add_B ();}else if _dce .Style .FillColor !=""{_eff .Add_f ();}else if _dce .Style .StrokeColor !=""{_eff .Add_S ();};_eff .Add_h ();_eff .Add_Q ();
};func _fgc (_egd *GraphicSVG ,_gbg *_ag .ContentCreator ){_gbg .Add_q ();_egd .Style .toContentStream (_gbg );_fdb ,_fag :=_gae (_egd .Attributes ["\u0063\u0078"],64);if _fag !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_fag .Error ());
};_cb ,_fag :=_gae (_egd .Attributes ["\u0063\u0079"],64);if _fag !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_fag .Error ());
};_dagc ,_fag :=_gae (_egd .Attributes ["\u0072\u0078"],64);if _fag !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_fag .Error ());
};_cbf ,_fag :=_gae (_egd .Attributes ["\u0072\u0079"],64);if _fag !=nil {_bd .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_fag .Error ());
};_dd :=_dagc *_egd ._daa ;_dcc :=_cbf *_egd ._daa ;_gg :=_fdb *_egd ._daa ;_gfbe :=_cb *_egd ._daa ;_dced :=_dd *_df ;_egb :=_dcc *_df ;_egg :=_dg .NewCubicBezierPath ();_egg =_egg .AppendCurve (_dg .NewCubicBezierCurve (-_dd ,0,-_dd ,_egb ,-_dced ,_dcc ,0,_dcc ));
_egg =_egg .AppendCurve (_dg .NewCubicBezierCurve (0,_dcc ,_dced ,_dcc ,_dd ,_egb ,_dd ,0));_egg =_egg .AppendCurve (_dg .NewCubicBezierCurve (_dd ,0,_dd ,-_egb ,_dced ,-_dcc ,0,-_dcc ));_egg =_egg .AppendCurve (_dg .NewCubicBezierCurve (0,-_dcc ,-_dced ,-_dcc ,-_dd ,-_egb ,-_dd ,0));
_egg =_egg .Offset (_gg ,_gfbe );if _egd .Style .StrokeWidth > 0{_egg =_egg .Offset (_egd .Style .StrokeWidth /2,_egd .Style .StrokeWidth /2);};_dg .DrawBezierPathWithCreator (_egg ,_gbg );if _egd .Style .FillColor !=""&&_egd .Style .StrokeColor !=""{_gbg .Add_B ();
}else if _egd .Style .FillColor !=""{_gbg .Add_f ();}else if _egd .Style .StrokeColor !=""{_gbg .Add_S ();};_gbg .Add_h ();_gbg .Add_Q ();};func _dag (_daf *GraphicSVG ,_abf *_ag .ContentCreator ){_abf .Add_q ();_daf .Style .toContentStream (_abf );_ead ,_gc :=_gae (_daf .Attributes ["\u0078"],64);
if _gc !=nil {_bd .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0078\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_gc .Error ());};_dbd ,_gc :=_gae (_daf .Attributes ["\u0079"],64);
if _gc !=nil {_bd .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0079\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_gc .Error ());};_gb ,_gc :=_gae (_daf .Attributes ["\u0077\u0069\u0064t\u0068"],64);
if _gc !=nil {_bd .Log .Debug ("\u0045\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gc .Error ());
};_cef ,_gc :=_gae (_daf .Attributes ["\u0068\u0065\u0069\u0067\u0068\u0074"],64);if _gc !=nil {_bd .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0077h\u0069\u006c\u0065 \u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gc .Error ());
};_abf .Add_re (_ead *_daf ._daa ,_dbd *_daf ._daa ,_gb *_daf ._daa ,_cef *_daf ._daa );if _daf .Style .FillColor !=""&&_daf .Style .StrokeColor !=""{_abf .Add_B ();}else if _daf .Style .FillColor !=""{_abf .Add_f ();}else if _daf .Style .StrokeColor !=""{_abf .Add_S ();
};_abf .Add_Q ();};func (_ccdg *Path )compare (_cac *Path )bool {if len (_ccdg .Subpaths )!=len (_cac .Subpaths ){return false ;};for _ebff ,_afe :=range _ccdg .Subpaths {if !_afe .compare (_cac .Subpaths [_ebff ]){return false ;};};return true ;};type commands struct{_aff []string ;
_gdaa map[string ]int ;_dagg string ;_ddg string ;};func ParseFromFile (path string )(*GraphicSVG ,error ){_egde ,_fabf :=_a .Open (path );if _fabf !=nil {return nil ,_fabf ;};defer _egde .Close ();return ParseFromStream (_egde );};func (_fcec pathParserError )Error ()string {return _fcec ._fbb };
func (_dga *GraphicSVG )Decode (decoder *_bbf .Decoder )error {for {_deb ,_fdg :=decoder .Token ();if _deb ==nil &&_fdg ==_bb .EOF {break ;};if _fdg !=nil {return _fdg ;};switch _fce :=_deb .(type ){case _bbf .StartElement :_aba :=_dff (_fce );_ddd :=_aba .Decode (decoder );
if _ddd !=nil {return _ddd ;};_dga .Children =append (_dga .Children ,_aba );case _bbf .CharData :_cd :=_d .TrimSpace (string (_fce ));if _cd !=""{_dga .Content =string (_fce );};case _bbf .EndElement :if _fce .Name .Local ==_dga .Name {return nil ;};};
};return nil ;};func (_effb *GraphicSVG )setDefaultScaling (_af float64 ){_effb ._daa =_af ;if _effb .Style !=nil &&_effb .Style .StrokeWidth > 0{_effb .Style .StrokeWidth =_effb .Style .StrokeWidth *_effb ._daa ;};for _ ,_bee :=range _effb .Children {_bee .setDefaultScaling (_af );
};};func _fdc (_ddda []token )([]*Command ,error ){var (_gdb []*Command ;_ggg []float64 ;);for _aggd :=len (_ddda )-1;_aggd >=0;_aggd --{_aca :=_ddda [_aggd ];if _aca ._aea {_bae :=_edd ._gdaa [_d .ToLower (_aca ._ecg )];_cdda :=len (_ggg );if _bae ==0&&_cdda ==0{_dfd :=&Command {Symbol :_aca ._ecg };
_gdb =append ([]*Command {_dfd },_gdb ...);}else if _bae !=0&&_cdda %_bae ==0{_gfba :=_cdda /_bae ;for _dgab :=0;_dgab < _gfba ;_dgab ++{_ceb :=_aca ._ecg ;if _ceb =="\u006d"&&_dgab < _gfba -1{_ceb ="\u006c";};if _ceb =="\u004d"&&_dgab < _gfba -1{_ceb ="\u004c";
};_acc :=&Command {_ceb ,_bdde (_ggg [:_bae ])};_gdb =append ([]*Command {_acc },_gdb ...);_ggg =_ggg [_bae :];};}else {_acee :=pathParserError {"I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006fr\u0020"+_aca ._ecg };
return nil ,_acee ;};}else {_aceb ,_fdbb :=_gae (_aca ._ecg ,64);if _fdbb !=nil {return nil ,_fdbb ;};_ggg =append (_ggg ,_aceb );};};return _gdb ,nil ;};func (_dbg *GraphicSVG )ToContentCreator (cc *_ag .ContentCreator ,scaleX ,scaleY ,translateX ,translateY float64 )*_ag .ContentCreator {if _dbg .Name =="\u0073\u0076\u0067"{_dbg .SetScaling (scaleX ,scaleY );
cc .Add_cm (1,0,0,1,translateX ,translateY );_dbg .setDefaultScaling (_dbg ._daa );cc .Add_q ();_efc :=_bf .Max (scaleX ,scaleY );cc .Add_re (_dbg .ViewBox .X *_efc ,_dbg .ViewBox .Y *_efc ,_dbg .ViewBox .W *_efc ,_dbg .ViewBox .H *_efc );cc .Add_W ();
cc .Add_n ();for _ ,_bfg :=range _dbg .Children {_bfg .ViewBox =_dbg .ViewBox ;_bfg .toContentStream (cc );};cc .Add_Q ();return cc ;};return nil ;};func _dff (_cgac _bbf .StartElement )*GraphicSVG {_dbe :=&GraphicSVG {};_fdf :=make (map[string ]string );
for _ ,_gcb :=range _cgac .Attr {_fdf [_gcb .Name .Local ]=_gcb .Value ;};_dbe .Name =_cgac .Name .Local ;_dbe .Attributes =_fdf ;_dbe ._daa =1;if _dbe .Name =="\u0073\u0076\u0067"{_feb ,_cfa :=_gbcg (_fdf ["\u0076i\u0065\u0077\u0042\u006f\u0078"]);if _cfa !=nil {_bd .Log .Debug ("\u0055\u006ea\u0062\u006c\u0065\u0020t\u006f\u0020p\u0061\u0072\u0073\u0065\u0020\u0076\u0069\u0065w\u0042\u006f\u0078\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074e\u003a\u0020\u0025\u0076",_cfa );
return nil ;};_dbe .ViewBox .X =_feb [0];_dbe .ViewBox .Y =_feb [1];_dbe .ViewBox .W =_feb [2];_dbe .ViewBox .H =_feb [3];_dbe .Width =_dbe .ViewBox .W ;_dbe .Height =_dbe .ViewBox .H ;if _cged ,_eaa :=_fdf ["\u0077\u0069\u0064t\u0068"];_eaa {_gfa ,_cee :=_gae (_cged ,64);
if _cee !=nil {_bd .Log .Debug ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0077\u0069\u0064\u0074\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_cee );return nil ;};_dbe .Width =_gfa ;
};if _dde ,_dcd :=_fdf ["\u0068\u0065\u0069\u0067\u0068\u0074"];_dcd {_gda ,_dcg :=_gae (_dde ,64);if _dcg !=nil {_bd .Log .Debug ("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0070\u0061rse\u0020he\u0069\u0067\u0068\u0074\u0020\u0061\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_dcg );
return nil ;};_dbe .Height =_gda ;};if _dbe .Width > 0&&_dbe .Height > 0{_dbe ._daa =_dbe .Width /_dbe .ViewBox .W ;};};return _dbe ;};