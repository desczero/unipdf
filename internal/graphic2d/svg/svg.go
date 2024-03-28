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

package svg ;import (_db "encoding/xml";_cf "fmt";_de "github.com/unidoc/unipdf/v3/common";_bg "github.com/unidoc/unipdf/v3/contentstream";_fa "github.com/unidoc/unipdf/v3/contentstream/draw";_fg "github.com/unidoc/unipdf/v3/internal/graphic2d";_df "golang.org/x/net/html/charset";
_a "io";_bd "math";_f "os";_d "strconv";_e "strings";_b "unicode";);type pathParserError struct{_bb string };func _cab (_edd *GraphicSVG ,_ffd *_bg .ContentCreator ){_ffd .Add_q ();_edd .Style .toContentStream (_ffd );_dga ,_cba :=_gge (_edd .Attributes ["\u0063\u0078"],64);
if _cba !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cba .Error ());};_ac ,_cba :=_gge (_edd .Attributes ["\u0063\u0079"],64);
if _cba !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cba .Error ());};_fee ,_cba :=_gge (_edd .Attributes ["\u0072"],64);
if _cba !=nil {_de .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0072\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_cba .Error ());};_be :=_fee *_edd ._bfc ;
_adg :=_fee *_edd ._bfc ;_dgd :=_be *_cg ;_caf :=_adg *_cg ;_abe :=_fa .NewCubicBezierPath ();_abe =_abe .AppendCurve (_fa .NewCubicBezierCurve (-_be ,0,-_be ,_caf ,-_dgd ,_adg ,0,_adg ));_abe =_abe .AppendCurve (_fa .NewCubicBezierCurve (0,_adg ,_dgd ,_adg ,_be ,_caf ,_be ,0));
_abe =_abe .AppendCurve (_fa .NewCubicBezierCurve (_be ,0,_be ,-_caf ,_dgd ,-_adg ,0,-_adg ));_abe =_abe .AppendCurve (_fa .NewCubicBezierCurve (0,-_adg ,-_dgd ,-_adg ,-_be ,-_caf ,-_be ,0));_abe =_abe .Offset (_dga *_edd ._bfc ,_ac *_edd ._bfc );if _edd .Style .StrokeWidth > 0{_abe =_abe .Offset (_edd .Style .StrokeWidth /2,_edd .Style .StrokeWidth /2);
};_fa .DrawBezierPathWithCreator (_abe ,_ffd );if _edd .Style .FillColor !=""&&_edd .Style .StrokeColor !=""{_ffd .Add_B ();}else if _edd .Style .FillColor !=""{_ffd .Add_f ();}else if _edd .Style .StrokeColor !=""{_ffd .Add_S ();};_ffd .Add_h ();_ffd .Add_Q ();
};var _efa commands ;type GraphicSVGStyle struct{FillColor string ;StrokeColor string ;StrokeWidth float64 ;};func _accb (_gbcc string )(_efe []float64 ,_fggb error ){var _cgb float64 ;_cbae :=0;_dgbg :=true ;for _cgfed ,_ffab :=range _gbcc {if _ffab =='.'{if _dgbg {_dgbg =false ;
continue ;};_cgb ,_fggb =_gge (_gbcc [_cbae :_cgfed ],64);if _fggb !=nil {return ;};_efe =append (_efe ,_cgb );_cbae =_cgfed ;};};_cgb ,_fggb =_gge (_gbcc [_cbae :],64);if _fggb !=nil {return ;};_efe =append (_efe ,_cgb );return ;};func _cebb (_cbgf map[string ]string ,_aceb float64 )(*GraphicSVGStyle ,error ){_baae :=_bdbg ();
_dcfd ,_dcff :=_cbgf ["\u0066\u0069\u006c\u006c"];if _dcff {_baae .FillColor =_dcfd ;if _dcfd =="\u006e\u006f\u006e\u0065"{_baae .FillColor ="";};};_ggg ,_dcag :=_cbgf ["\u0073\u0074\u0072\u006f\u006b\u0065"];if _dcag {_baae .StrokeColor =_ggg ;if _ggg =="\u006e\u006f\u006e\u0065"{_baae .StrokeColor ="";
};};_eff ,_fbd :=_cbgf ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];if _fbd {_dbbg ,_fgf :=_gge (_eff ,64);if _fgf !=nil {return nil ,_fgf ;};_baae .StrokeWidth =_dbbg *_aceb ;};return _baae ,nil ;};func ParseFromFile (path string )(*GraphicSVG ,error ){_ace ,_bdee :=_f .Open (path );
if _bdee !=nil {return nil ,_bdee ;};defer _ace .Close ();return ParseFromStream (_ace );};func (_cga *GraphicSVG )ToContentCreator (cc *_bg .ContentCreator ,scaleX ,scaleY ,translateX ,translateY float64 )*_bg .ContentCreator {if _cga .Name =="\u0073\u0076\u0067"{_cga .SetScaling (scaleX ,scaleY );
cc .Add_cm (1,0,0,1,translateX ,translateY );_cga .setDefaultScaling (_cga ._bfc );cc .Add_q ();_dggf :=_bd .Max (scaleX ,scaleY );cc .Add_re (_cga .ViewBox .X *_dggf ,_cga .ViewBox .Y *_dggf ,_cga .ViewBox .W *_dggf ,_cga .ViewBox .H *_dggf );cc .Add_W ();
cc .Add_n ();for _ ,_acdb :=range _cga .Children {_acdb .ViewBox =_cga .ViewBox ;_acdb .toContentStream (cc );};cc .Add_Q ();return cc ;};return nil ;};func _caab (_ded []*Command )*Path {_gbe :=&Path {};var _cbbc []*Command ;for _aeff ,_geb :=range _ded {switch _e .ToLower (_geb .Symbol ){case _efa ._bgfd :if len (_cbbc )> 0{_gbe .Subpaths =append (_gbe .Subpaths ,&Subpath {_cbbc });
};_cbbc =[]*Command {_geb };case _efa ._fgfg :_cbbc =append (_cbbc ,_geb );_gbe .Subpaths =append (_gbe .Subpaths ,&Subpath {_cbbc });_cbbc =[]*Command {};default:_cbbc =append (_cbbc ,_geb );if len (_ded )==_aeff +1{_gbe .Subpaths =append (_gbe .Subpaths ,&Subpath {_cbbc });
};};};return _gbe ;};func _dcfc (_gdf string )[]token {var (_debf []token ;_cgf string ;);for _ ,_efb :=range _gdf {_cdc :=string (_efb );switch {case _efa .isCommand (_cdc ):_debf ,_cgf =_dbgf (_debf ,_cgf );_debf =append (_debf ,token {_cdc ,true });
case _cdc =="\u002e":if _cgf ==""{_cgf ="\u0030";};if _e .Contains (_cgf ,_cdc ){_debf =append (_debf ,token {_cgf ,false });_cgf ="\u0030";};fallthrough;case _cdc >="\u0030"&&_cdc <="\u0039"||_cdc =="\u0065":_cgf +=_cdc ;case _cdc =="\u002d":if _e .HasSuffix (_cgf ,"\u0065"){_cgf +=_cdc ;
}else {_debf ,_ =_dbgf (_debf ,_cgf );_cgf =_cdc ;};default:_debf ,_cgf =_dbgf (_debf ,_cgf );};};_debf ,_ =_dbgf (_debf ,_cgf );return _debf ;};func _bce (_ege *GraphicSVG ,_ccc *_bg .ContentCreator ){_ccc .Add_q ();_ege .Style .toContentStream (_ccc );
_cfe ,_eef :=_adf (_ege .Attributes ["\u0070\u006f\u0069\u006e\u0074\u0073"]);if _eef !=nil {_de .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076",_eef );
return ;};if len (_cfe )%2> 0{_de .Log .Debug ("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068");return ;
};for _febg :=0;_febg < len (_cfe );{if _febg ==0{_ccc .Add_m (_cfe [_febg ]*_ege ._bfc ,_cfe [_febg +1]*_ege ._bfc );}else {_ccc .Add_l (_cfe [_febg ]*_ege ._bfc ,_cfe [_febg +1]*_ege ._bfc );};_febg +=2;};if _ege .Style .FillColor !=""&&_ege .Style .StrokeColor !=""{_ccc .Add_B ();
}else if _ege .Style .FillColor !=""{_ccc .Add_f ();}else if _ege .Style .StrokeColor !=""{_ccc .Add_S ();};_ccc .Add_h ();_ccc .Add_Q ();};func (_bcd *GraphicSVG )SetScaling (xFactor ,yFactor float64 ){_fad :=_bcd .Width /_bcd .ViewBox .W ;_ffa :=_bcd .Height /_bcd .ViewBox .H ;
_bcd .setDefaultScaling (_bd .Max (_fad ,_ffa ));for _ ,_ceg :=range _bcd .Children {_ceg .SetScaling (xFactor ,yFactor );};};func (_ccaf *Path )compare (_caa *Path )bool {if len (_ccaf .Subpaths )!=len (_caa .Subpaths ){return false ;};for _feg ,_cdb :=range _ccaf .Subpaths {if !_cdb .compare (_caa .Subpaths [_feg ]){return false ;
};};return true ;};func (_bdaf *GraphicSVG )setDefaultScaling (_cedd float64 ){_bdaf ._bfc =_cedd ;if _bdaf .Style !=nil &&_bdaf .Style .StrokeWidth > 0{_bdaf .Style .StrokeWidth =_bdaf .Style .StrokeWidth *_bdaf ._bfc ;};for _ ,_dca :=range _bdaf .Children {_dca .setDefaultScaling (_cedd );
};};func _bcf ()commands {var _ebc =map[string ]int {"\u006d":2,"\u007a":0,"\u006c":2,"\u0068":1,"\u0076":1,"\u0063":6,"\u0073":4,"\u0071":4,"\u0074":2,"\u0061":7};var _aea []string ;for _faff :=range _ebc {_aea =append (_aea ,_faff );};return commands {_aea ,_ebc ,"\u006d","\u007a"};
};func (_cfa *GraphicSVG )toContentStream (_egd *_bg .ContentCreator ){_efc ,_ccg :=_cebb (_cfa .Attributes ,_cfa ._bfc );if _ccg !=nil {_de .Log .Debug ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_ccg );
};_cfa .Style =_efc ;switch _cfa .Name {case "\u0070\u0061\u0074\u0068":_cge (_cfa ,_egd );for _ ,_ebd :=range _cfa .Children {_ebd .toContentStream (_egd );};case "\u0072\u0065\u0063\u0074":_fdg (_cfa ,_egd );for _ ,_baa :=range _cfa .Children {_baa .toContentStream (_egd );
};case "\u0063\u0069\u0072\u0063\u006c\u0065":_cab (_cfa ,_egd );for _ ,_bceb :=range _cfa .Children {_bceb .toContentStream (_egd );};case "\u0065l\u006c\u0069\u0070\u0073\u0065":_fc (_cfa ,_egd );for _ ,_feed :=range _cfa .Children {_feed .toContentStream (_egd );
};case "\u0070\u006f\u006c\u0079\u006c\u0069\u006e\u0065":_bce (_cfa ,_egd );for _ ,_bdd :=range _cfa .Children {_bdd .toContentStream (_egd );};case "\u0070o\u006c\u0079\u0067\u006f\u006e":_eab (_cfa ,_egd );for _ ,_gbc :=range _cfa .Children {_gbc .toContentStream (_egd );
};case "\u006c\u0069\u006e\u0065":_dbb (_cfa ,_egd );for _ ,_eee :=range _cfa .Children {_eee .toContentStream (_egd );};case "\u0067":_cce ,_fafd :=_cfa .Attributes ["\u0066\u0069\u006c\u006c"];_ceaa ,_bgca :=_cfa .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"];
_fag ,_acc :=_cfa .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];for _ ,_efdd :=range _cfa .Children {if _ ,_dcc :=_efdd .Attributes ["\u0066\u0069\u006c\u006c"];!_dcc &&_fafd {_efdd .Attributes ["\u0066\u0069\u006c\u006c"]=_cce ;
};if _ ,_gcd :=_efdd .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"];!_gcd &&_bgca {_efdd .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"]=_ceaa ;};if _ ,_gab :=_efdd .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];
!_gab &&_acc {_efdd .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]=_fag ;};_efdd .toContentStream (_egd );};};};func (_gd *commands )isCommand (_age string )bool {for _ ,_fea :=range _gd ._bdc {if _e .ToLower (_age )==_fea {return true ;
};};return false ;};func (_gdc *Command )isAbsolute ()bool {return _gdc .Symbol ==_e .ToUpper (_gdc .Symbol )};func _gef (_cde string )(*Path ,error ){_efa =_bcf ();_cbe ,_fff :=_abd (_dcfc (_cde ));if _fff !=nil {return nil ,_fff ;};return _caab (_cbe ),nil ;
};func (_ffe *GraphicSVG )Decode (decoder *_db .Decoder )error {for {_dea ,_daf :=decoder .Token ();if _dea ==nil &&_daf ==_a .EOF {break ;};if _daf !=nil {return _daf ;};switch _cfb :=_dea .(type ){case _db .StartElement :_gaf :=_ebe (_cfb );_fdf :=_gaf .Decode (decoder );
if _fdf !=nil {return _fdf ;};_ffe .Children =append (_ffe .Children ,_gaf );case _db .CharData :_def :=_e .TrimSpace (string (_cfb ));if _def !=""{_ffe .Content =string (_cfb );};case _db .EndElement :if _cfb .Name .Local ==_ffe .Name {return nil ;};};
};return nil ;};type Subpath struct{Commands []*Command ;};func _cge (_dfe *GraphicSVG ,_g *_bg .ContentCreator ){_g .Add_q ();_dfe .Style .toContentStream (_g );_bde ,_aa :=_gef (_dfe .Attributes ["\u0064"]);if _aa !=nil {_de .Log .Error ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025s",_aa .Error ());
};var (_fd ,_dbg =0.0,0.0;_gb ,_deb =0.0,0.0;_dg *Command ;);for _ ,_dge :=range _bde .Subpaths {for _ ,_ed :=range _dge .Commands {switch _e .ToLower (_ed .Symbol ){case "\u006d":_gb ,_deb =_ed .Params [0]*_dfe ._bfc ,_ed .Params [1]*_dfe ._bfc ;if !_ed .isAbsolute (){_gb ,_deb =_fd +_gb -_dfe .ViewBox .X ,_dbg +_deb -_dfe .ViewBox .Y ;
};_g .Add_m (_dfbb (_gb ,3),_dfbb (_deb ,3));_fd ,_dbg =_gb ,_deb ;case "\u0063":_dgg ,_af ,_dd ,_cb ,_bc ,_cfc :=_ed .Params [0]*_dfe ._bfc ,_ed .Params [1]*_dfe ._bfc ,_ed .Params [2]*_dfe ._bfc ,_ed .Params [3]*_dfe ._bfc ,_ed .Params [4]*_dfe ._bfc ,_ed .Params [5]*_dfe ._bfc ;
if !_ed .isAbsolute (){_dgg ,_af ,_dd ,_cb ,_bc ,_cfc =_fd +_dgg ,_dbg +_af ,_fd +_dd ,_dbg +_cb ,_fd +_bc ,_dbg +_cfc ;};_g .Add_c (_dfbb (_dgg ,3),_dfbb (_af ,3),_dfbb (_dd ,3),_dfbb (_cb ,3),_dfbb (_bc ,3),_dfbb (_cfc ,3));_fd ,_dbg =_bc ,_cfc ;case "\u0073":_eb ,_faa ,_dbd ,_ga :=_ed .Params [0]*_dfe ._bfc ,_ed .Params [1]*_dfe ._bfc ,_ed .Params [2]*_dfe ._bfc ,_ed .Params [3]*_dfe ._bfc ;
if !_ed .isAbsolute (){_eb ,_faa ,_dbd ,_ga =_fd +_eb ,_dbg +_faa ,_fd +_dbd ,_dbg +_ga ;};_g .Add_c (_dfbb (_fd ,3),_dfbb (_dbg ,3),_dfbb (_eb ,3),_dfbb (_faa ,3),_dfbb (_dbd ,3),_dfbb (_ga ,3));_fd ,_dbg =_dbd ,_ga ;case "\u006c":_bf ,_afa :=_ed .Params [0]*_dfe ._bfc ,_ed .Params [1]*_dfe ._bfc ;
if !_ed .isAbsolute (){_bf ,_afa =_fd +_bf ,_dbg +_afa ;};_g .Add_l (_dfbb (_bf ,3),_dfbb (_afa ,3));_fd ,_dbg =_bf ,_afa ;case "\u0068":_cc :=_ed .Params [0]*_dfe ._bfc ;if !_ed .isAbsolute (){_cc =_fd +_cc ;};_g .Add_l (_dfbb (_cc ,3),_dfbb (_dbg ,3));
_fd =_cc ;case "\u0076":_cbd :=_ed .Params [0]*_dfe ._bfc ;if !_ed .isAbsolute (){_cbd =_dbg +_cbd ;};_g .Add_l (_dfbb (_fd ,3),_dfbb (_cbd ,3));_dbg =_cbd ;case "\u0071":_bda ,_afd ,_cfcc ,_ab :=_ed .Params [0]*_dfe ._bfc ,_ed .Params [1]*_dfe ._bfc ,_ed .Params [2]*_dfe ._bfc ,_ed .Params [3]*_dfe ._bfc ;
if !_ed .isAbsolute (){_bda ,_afd ,_cfcc ,_ab =_fd +_bda ,_dbg +_afd ,_fd +_cfcc ,_dbg +_ab ;};_fba ,_ad :=_fg .QuadraticToCubicBezier (_fd ,_dbg ,_bda ,_afd ,_cfcc ,_ab );_g .Add_c (_dfbb (_fba .X ,3),_dfbb (_fba .Y ,3),_dfbb (_ad .X ,3),_dfbb (_ad .Y ,3),_dfbb (_cfcc ,3),_dfbb (_ab ,3));
_fd ,_dbg =_cfcc ,_ab ;case "\u0074":var _gc ,_da _fg .Point ;_dde ,_ag :=_ed .Params [0]*_dfe ._bfc ,_ed .Params [1]*_dfe ._bfc ;if !_ed .isAbsolute (){_dde ,_ag =_fd +_dde ,_dbg +_ag ;};if _dg !=nil &&_e .ToLower (_dg .Symbol )=="\u0071"{_ee :=_fg .Point {X :_dg .Params [0]*_dfe ._bfc ,Y :_dg .Params [1]*_dfe ._bfc };
_afdd :=_fg .Point {X :_dg .Params [2]*_dfe ._bfc ,Y :_dg .Params [3]*_dfe ._bfc };_ccf :=_afdd .Mul (2.0).Sub (_ee );_gc ,_da =_fg .QuadraticToCubicBezier (_fd ,_dbg ,_ccf .X ,_ccf .Y ,_dde ,_ag );};_g .Add_c (_dfbb (_gc .X ,3),_dfbb (_gc .Y ,3),_dfbb (_da .X ,3),_dfbb (_da .Y ,3),_dfbb (_dde ,3),_dfbb (_ag ,3));
_fd ,_dbg =_dde ,_ag ;case "\u0061":_ca ,_fdc :=_ed .Params [0]*_dfe ._bfc ,_ed .Params [1]*_dfe ._bfc ;_dab :=_ed .Params [2];_gac :=_ed .Params [3]> 0;_faf :=_ed .Params [4]> 0;_edc ,_aaf :=_ed .Params [5]*_dfe ._bfc ,_ed .Params [6]*_dfe ._bfc ;if !_ed .isAbsolute (){_edc ,_aaf =_fd +_edc ,_dbg +_aaf ;
};_bdb :=_fg .EllipseToCubicBeziers (_fd ,_dbg ,_ca ,_fdc ,_dab ,_gac ,_faf ,_edc ,_aaf );for _ ,_add :=range _bdb {_g .Add_c (_dfbb (_add [1].X ,3),_dfbb ((_add [1].Y ),3),_dfbb ((_add [2].X ),3),_dfbb ((_add [2].Y ),3),_dfbb ((_add [3].X ),3),_dfbb ((_add [3].Y ),3));
};_fd ,_dbg =_edc ,_aaf ;case "\u007a":_g .Add_h ();};_dg =_ed ;};};if _dfe .Style .FillColor !=""&&_dfe .Style .StrokeColor !=""{_g .Add_B ();}else if _dfe .Style .FillColor !=""{_g .Add_f ();}else if _dfe .Style .StrokeColor !=""{_g .Add_S ();};_g .Add_h ();
_g .Add_Q ();};func _dfbb (_affd float64 ,_cacb int )float64 {_egef :=_bd .Pow (10,float64 (_cacb ));return float64 (_dcaf (_affd *_egef ))/_egef ;};type Path struct{Subpaths []*Subpath ;};type token struct{_gcf string ;_cca bool ;};func (_agf *Subpath )compare (_dgac *Subpath )bool {if len (_agf .Commands )!=len (_dgac .Commands ){return false ;
};for _dafa ,_bad :=range _agf .Commands {if !_bad .compare (_dgac .Commands [_dafa ]){return false ;};};return true ;};func ParseFromStream (source _a .Reader )(*GraphicSVG ,error ){_ceb :=_db .NewDecoder (source );_ceb .CharsetReader =_df .NewReaderLabel ;
_fgc ,_aec :=_dff (_ceb );if _aec !=nil {return nil ,_aec ;};if _eeef :=_fgc .Decode (_ceb );_eeef !=nil &&_eeef !=_a .EOF {return nil ,_eeef ;};return _fgc ,nil ;};func _bae (_dgdde string )(_fdggg ,_adgc ,_ebcg float64 ){if (len (_dgdde )!=4&&len (_dgdde )!=7)||_dgdde [0]!='#'{_de .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_dgdde );
return _fdggg ,_adgc ,_ebcg ;};var _egefc ,_eddde ,_bee int ;if len (_dgdde )==4{var _aga ,_babd ,_cbc int ;_eccf ,_aad :=_cf .Sscanf (_dgdde ,"\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078",&_aga ,&_babd ,&_cbc );if _aad !=nil {_de .Log .Debug ("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076",_dgdde ,_aad );
return _fdggg ,_adgc ,_ebcg ;};if _eccf !=3{_de .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_dgdde );return _fdggg ,_adgc ,_ebcg ;};_egefc =_aga *16+_aga ;_eddde =_babd *16+_babd ;_bee =_cbc *16+_cbc ;
}else {_cda ,_bcfg :=_cf .Sscanf (_dgdde ,"\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078",&_egefc ,&_eddde ,&_bee );if _bcfg !=nil {_de .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_dgdde );
return _fdggg ,_adgc ,_ebcg ;};if _cda !=3{_de .Log .Debug ("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029",_dgdde ,_cda );
return _fdggg ,_adgc ,_ebcg ;};};_dbgff :=float64 (_egefc )/255.0;_gfd :=float64 (_eddde )/255.0;_ead :=float64 (_bee )/255.0;return _dbgff ,_gfd ,_ead ;};func _fdg (_eg *GraphicSVG ,_bgc *_bg .ContentCreator ){_bgc .Add_q ();_eg .Style .toContentStream (_bgc );
_fed ,_ega :=_gge (_eg .Attributes ["\u0078"],64);if _ega !=nil {_de .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0078\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_ega .Error ());
};_fac ,_ega :=_gge (_eg .Attributes ["\u0079"],64);if _ega !=nil {_de .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0079\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_ega .Error ());
};_eeb ,_ega :=_gge (_eg .Attributes ["\u0077\u0069\u0064t\u0068"],64);if _ega !=nil {_de .Log .Debug ("\u0045\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_ega .Error ());
};_dc ,_ega :=_gge (_eg .Attributes ["\u0068\u0065\u0069\u0067\u0068\u0074"],64);if _ega !=nil {_de .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0077h\u0069\u006c\u0065 \u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_ega .Error ());
};_bgc .Add_re (_fed *_eg ._bfc ,_fac *_eg ._bfc ,_eeb *_eg ._bfc ,_dc *_eg ._bfc );if _eg .Style .FillColor !=""&&_eg .Style .StrokeColor !=""{_bgc .Add_B ();}else if _eg .Style .FillColor !=""{_bgc .Add_f ();}else if _eg .Style .StrokeColor !=""{_bgc .Add_S ();
};_bgc .Add_Q ();};func (_dgdg pathParserError )Error ()string {return _dgdg ._bb };type GraphicSVG struct{ViewBox struct{X ,Y ,W ,H float64 ;};Name string ;Attributes map[string ]string ;Children []*GraphicSVG ;Content string ;Style *GraphicSVGStyle ;
Width float64 ;Height float64 ;_bfc float64 ;};func _dff (_bcgg *_db .Decoder )(*GraphicSVG ,error ){for {_ece ,_fae :=_bcgg .Token ();if _ece ==nil &&_fae ==_a .EOF {break ;};if _fae !=nil {return nil ,_fae ;};switch _ffdd :=_ece .(type ){case _db .StartElement :return _ebe (_ffdd ),nil ;
};};return &GraphicSVG {},nil ;};type commands struct{_bdc []string ;_afc map[string ]int ;_bgfd string ;_fgfg string ;};func ParseFromString (svgStr string )(*GraphicSVG ,error ){return ParseFromStream (_e .NewReader (svgStr ));};func _fced (_ggea string )(_fdga ,_ebdc string ){if _ggea ==""||(_ggea [len (_ggea )-1]>='0'&&_ggea [len (_ggea )-1]<='9'){return _ggea ,"";
};_fdga =_ggea ;for _ ,_cfg :=range _fe {if _e .Contains (_fdga ,_cfg ){_ebdc =_cfg ;};_fdga =_e .TrimSuffix (_fdga ,_cfg );};return ;};func _abd (_afg []token )([]*Command ,error ){var (_agbd []*Command ;_deg []float64 ;);for _bgce :=len (_afg )-1;_bgce >=0;
_bgce --{_bddg :=_afg [_bgce ];if _bddg ._cca {_dbgc :=_efa ._afc [_e .ToLower (_bddg ._gcf )];_abg :=len (_deg );if _dbgc ==0&&_abg ==0{_fdgee :=&Command {Symbol :_bddg ._gcf };_agbd =append ([]*Command {_fdgee },_agbd ...);}else if _dbgc !=0&&_abg %_dbgc ==0{_cgfe :=_abg /_dbgc ;
for _dgaf :=0;_dgaf < _cgfe ;_dgaf ++{_eag :=_bddg ._gcf ;if _eag =="\u006d"&&_dgaf < _cgfe -1{_eag ="\u006c";};if _eag =="\u004d"&&_dgaf < _cgfe -1{_eag ="\u004c";};_fdgg :=&Command {_eag ,_dgc (_deg [:_dbgc ])};_agbd =append ([]*Command {_fdgg },_agbd ...);
_deg =_deg [_dbgc :];};}else {_acebf :=pathParserError {"I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006fr\u0020"+_bddg ._gcf };
return nil ,_acebf ;};}else {_ged ,_dbe :=_gge (_bddg ._gcf ,64);if _dbe !=nil {return nil ,_dbe ;};_deg =append (_deg ,_ged );};};return _agbd ,nil ;};func _dgc (_efae []float64 )[]float64 {for _acee ,_cac :=0,len (_efae )-1;_acee < _cac ;_acee ,_cac =_acee +1,_cac -1{_efae [_acee ],_efae [_cac ]=_efae [_cac ],_efae [_acee ];
};return _efae ;};func _adf (_daa string )([]float64 ,error ){_bea :=-1;var _bcgb []float64 ;_gabd :=' ';for _defg ,_gbca :=range _daa {if !_b .IsNumber (_gbca )&&_gbca !='.'&&!(_gbca =='-'&&_gabd =='e')&&_gbca !='e'{if _bea !=-1{_fgcd ,_caff :=_accb (_daa [_bea :_defg ]);
if _caff !=nil {return _bcgb ,_caff ;};_bcgb =append (_bcgb ,_fgcd ...);};if _gbca =='-'{_bea =_defg ;}else {_bea =-1;};}else if _bea ==-1{_bea =_defg ;};_gabd =_gbca ;};if _bea !=-1&&_bea !=len (_daa ){_ada ,_ebg :=_accb (_daa [_bea :]);if _ebg !=nil {return _bcgb ,_ebg ;
};_bcgb =append (_bcgb ,_ada ...);};return _bcgb ,nil ;};func _dcaf (_ccag float64 )int {return int (_ccag +_bd .Copysign (0.5,_ccag ))};func (_fdff *GraphicSVGStyle )toContentStream (_ccd *_bg .ContentCreator ){if _fdff ==nil {return ;};if _fdff .FillColor !=""{var _edb ,_gag ,_gaca float64 ;
if _fbad ,_fagc :=_fg .ColorMap [_fdff .FillColor ];_fagc {_cgcd ,_bcg ,_gae ,_ :=_fbad .RGBA ();_edb ,_gag ,_gaca =float64 (_cgcd ),float64 (_bcg ),float64 (_gae );}else {_edb ,_gag ,_gaca =_bae (_fdff .FillColor );};_ccd .Add_rg (_edb ,_gag ,_gaca );
};if _fdff .StrokeColor !=""{var _fcc ,_gacf ,_ffag float64 ;if _cbb ,_fggd :=_fg .ColorMap [_fdff .StrokeColor ];_fggd {_bab ,_bga ,_bcc ,_ :=_cbb .RGBA ();_fcc ,_gacf ,_ffag =float64 (_bab )/255.0,float64 (_bga )/255.0,float64 (_bcc )/255.0;}else {_fcc ,_gacf ,_ffag =_bae (_fdff .StrokeColor );
};_ccd .Add_RG (_fcc ,_gacf ,_ffag );};if _fdff .StrokeWidth > 0{_ccd .Add_w (_fdff .StrokeWidth );};};func _eab (_acd *GraphicSVG ,_gf *_bg .ContentCreator ){_gf .Add_q ();_acd .Style .toContentStream (_gf );_fdge ,_eaf :=_adf (_acd .Attributes ["\u0070\u006f\u0069\u006e\u0074\u0073"]);
if _eaf !=nil {_de .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076",_eaf );
return ;};if len (_fdge )%2> 0{_de .Log .Debug ("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068");return ;
};for _fbg :=0;_fbg < len (_fdge );{if _fbg ==0{_gf .Add_m (_fdge [_fbg ]*_acd ._bfc ,_fdge [_fbg +1]*_acd ._bfc );}else {_gf .Add_l (_fdge [_fbg ]*_acd ._bfc ,_fdge [_fbg +1]*_acd ._bfc );};_fbg +=2;};_gf .Add_l (_fdge [0]*_acd ._bfc ,_fdge [1]*_acd ._bfc );
if _acd .Style .FillColor !=""&&_acd .Style .StrokeColor !=""{_gf .Add_B ();}else if _acd .Style .FillColor !=""{_gf .Add_f ();}else if _acd .Style .StrokeColor !=""{_gf .Add_S ();};_gf .Add_h ();_gf .Add_Q ();};func (_dad *Command )compare (_cfd *Command )bool {if _dad .Symbol !=_cfd .Symbol {return false ;
};for _dgeg ,_eeeg :=range _dad .Params {if _eeeg !=_cfd .Params [_dgeg ]{return false ;};};return true ;};func _bdbg ()*GraphicSVGStyle {return &GraphicSVGStyle {FillColor :"\u00230\u0030\u0030\u0030\u0030\u0030",StrokeColor :"",StrokeWidth :0};};type Command struct{Symbol string ;
Params []float64 ;};const (_ef =0.72;_cff =28.3464;_ce =_cff /10;_cg =0.551784;);func _dbb (_fgg *GraphicSVG ,_abea *_bg .ContentCreator ){_abea .Add_q ();_fgg .Style .toContentStream (_abea );_agg ,_cgg :=_gge (_fgg .Attributes ["\u0078\u0031"],64);if _cgg !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cgg .Error ());
};_eddd ,_cgg :=_gge (_fgg .Attributes ["\u0079\u0031"],64);if _cgg !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cgg .Error ());
};_ede ,_cgg :=_gge (_fgg .Attributes ["\u0078\u0032"],64);if _cgg !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cgg .Error ());
};_dee ,_cgg :=_gge (_fgg .Attributes ["\u0079\u0032"],64);if _cgg !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cgg .Error ());
};_abea .Add_m (_agg *_fgg ._bfc ,_eddd *_fgg ._bfc );_abea .Add_l (_ede *_fgg ._bfc ,_dee *_fgg ._bfc );if _fgg .Style .FillColor !=""&&_fgg .Style .StrokeColor !=""{_abea .Add_B ();}else if _fgg .Style .FillColor !=""{_abea .Add_f ();}else if _fgg .Style .StrokeColor !=""{_abea .Add_S ();
};_abea .Add_h ();_abea .Add_Q ();};func _gge (_beaf string ,_eeegg int )(float64 ,error ){_cgcdf ,_fcf :=_fced (_beaf );_dec ,_dafg :=_d .ParseFloat (_cgcdf ,_eeegg );if _dafg !=nil {return 0,_dafg ;};if _ecc ,_feab :=_bdf [_fcf ];_feab {_dec =_dec *_ecc ;
}else {_dec =_dec *_ef ;};return _dec ,nil ;};func _dbgf (_agbb []token ,_ccgb string )([]token ,string ){if _ccgb !=""{_agbb =append (_agbb ,token {_ccgb ,false });_ccgb ="";};return _agbb ,_ccgb ;};var (_fe =[]string {"\u0063\u006d","\u006d\u006d","\u0070\u0078","\u0070\u0074"};
_bdf =map[string ]float64 {"\u0063\u006d":_cff ,"\u006d\u006d":_ce ,"\u0070\u0078":_ef ,"\u0070\u0074":1};);func _ebe (_gcg _db .StartElement )*GraphicSVG {_agb :=&GraphicSVG {};_gabc :=make (map[string ]string );for _ ,_cd :=range _gcg .Attr {_gabc [_cd .Name .Local ]=_cd .Value ;
};_agb .Name =_gcg .Name .Local ;_agb .Attributes =_gabc ;_agb ._bfc =1;if _agb .Name =="\u0073\u0076\u0067"{_aeg ,_fce :=_adf (_gabc ["\u0076i\u0065\u0077\u0042\u006f\u0078"]);if _fce !=nil {_de .Log .Debug ("\u0055\u006ea\u0062\u006c\u0065\u0020t\u006f\u0020p\u0061\u0072\u0073\u0065\u0020\u0076\u0069\u0065w\u0042\u006f\u0078\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074e\u003a\u0020\u0025\u0076",_fce );
return nil ;};_agb .ViewBox .X =_aeg [0];_agb .ViewBox .Y =_aeg [1];_agb .ViewBox .W =_aeg [2];_agb .ViewBox .H =_aeg [3];_agb .Width =_agb .ViewBox .W ;_agb .Height =_agb .ViewBox .H ;if _bge ,_ge :=_gabc ["\u0077\u0069\u0064t\u0068"];_ge {_bff ,_gcb :=_gge (_bge ,64);
if _gcb !=nil {_de .Log .Debug ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0077\u0069\u0064\u0074\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_gcb );return nil ;};_agb .Width =_bff ;
};if _ceaf ,_cbg :=_gabc ["\u0068\u0065\u0069\u0067\u0068\u0074"];_cbg {_adgg ,_aff :=_gge (_ceaf ,64);if _aff !=nil {_de .Log .Debug ("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0070\u0061rse\u0020he\u0069\u0067\u0068\u0074\u0020\u0061\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_aff );
return nil ;};_agb .Height =_adgg ;};if _agb .Width > 0&&_agb .Height > 0{_agb ._bfc =_agb .Width /_agb .ViewBox .W ;};};return _agb ;};func _fc (_feb *GraphicSVG ,_cgc *_bg .ContentCreator ){_cgc .Add_q ();_feb .Style .toContentStream (_cgc );_gacg ,_gbf :=_gge (_feb .Attributes ["\u0063\u0078"],64);
if _gbf !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gbf .Error ());};_cffd ,_gbf :=_gge (_feb .Attributes ["\u0063\u0079"],64);
if _gbf !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gbf .Error ());};_bgcd ,_gbf :=_gge (_feb .Attributes ["\u0072\u0078"],64);
if _gbf !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gbf .Error ());};_acf ,_gbf :=_gge (_feb .Attributes ["\u0072\u0079"],64);
if _gbf !=nil {_de .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gbf .Error ());};_ea :=_bgcd *_feb ._bfc ;
_gad :=_acf *_feb ._bfc ;_gg :=_gacg *_feb ._bfc ;_cffe :=_cffd *_feb ._bfc ;_dfb :=_ea *_cg ;_ec :=_gad *_cg ;_ced :=_fa .NewCubicBezierPath ();_ced =_ced .AppendCurve (_fa .NewCubicBezierCurve (-_ea ,0,-_ea ,_ec ,-_dfb ,_gad ,0,_gad ));_ced =_ced .AppendCurve (_fa .NewCubicBezierCurve (0,_gad ,_dfb ,_gad ,_ea ,_ec ,_ea ,0));
_ced =_ced .AppendCurve (_fa .NewCubicBezierCurve (_ea ,0,_ea ,-_ec ,_dfb ,-_gad ,0,-_gad ));_ced =_ced .AppendCurve (_fa .NewCubicBezierCurve (0,-_gad ,-_dfb ,-_gad ,-_ea ,-_ec ,-_ea ,0));_ced =_ced .Offset (_gg ,_cffe );if _feb .Style .StrokeWidth > 0{_ced =_ced .Offset (_feb .Style .StrokeWidth /2,_feb .Style .StrokeWidth /2);
};_fa .DrawBezierPathWithCreator (_ced ,_cgc );if _feb .Style .FillColor !=""&&_feb .Style .StrokeColor !=""{_cgc .Add_B ();}else if _feb .Style .FillColor !=""{_cgc .Add_f ();}else if _feb .Style .StrokeColor !=""{_cgc .Add_S ();};_cgc .Add_h ();_cgc .Add_Q ();
};