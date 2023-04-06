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

package svg ;import (_gg "encoding/xml";_g "fmt";_fc "github.com/unidoc/unipdf/v3/common";_fd "github.com/unidoc/unipdf/v3/contentstream";_bf "github.com/unidoc/unipdf/v3/contentstream/draw";_bd "github.com/unidoc/unipdf/v3/internal/graphic2d";_eb "golang.org/x/net/html/charset";
_f "io";_a "math";_cd "os";_e "strconv";_c "strings";_bc "unicode";);func (_bcf *GraphicSVG )Decode (decoder *_gg .Decoder )error {for {_ffd ,_bgf :=decoder .Token ();if _ffd ==nil &&_bgf ==_f .EOF {break ;};if _bgf !=nil {return _bgf ;};switch _eeb :=_ffd .(type ){case _gg .StartElement :_fcb :=_cadg (_eeb );
_egf :=_fcb .Decode (decoder );if _egf !=nil {return _egf ;};_bcf .Children =append (_bcf .Children ,_fcb );case _gg .CharData :_gc :=_c .TrimSpace (string (_eeb ));if _gc !=""{_bcf .Content =string (_eeb );};case _gg .EndElement :if _eeb .Name .Local ==_bcf .Name {return nil ;
};};};return nil ;};type commands struct{_bcbb []string ;_gff map[string ]int ;_bde string ;_gcb string ;};func _cgg (_egg *GraphicSVG ,_gag *_fd .ContentCreator ){_gag .Add_q ();_egg .Style .toContentStream (_gag );_ccbb ,_cb :=_fga (_egg .Attributes ["\u0078\u0031"],64);
if _cb !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cb .Error ());};_ace ,_cb :=_fga (_egg .Attributes ["\u0079\u0031"],64);
if _cb !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cb .Error ());};_fcfe ,_cb :=_fga (_egg .Attributes ["\u0078\u0032"],64);
if _cb !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cb .Error ());};_acb ,_cb :=_fga (_egg .Attributes ["\u0079\u0032"],64);
if _cb !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_cb .Error ());};_gag .Add_m (_ccbb *_egg ._faf ,_ace *_egg ._faf );
_gag .Add_l (_fcfe *_egg ._faf ,_acb *_egg ._faf );if _egg .Style .FillColor !=""&&_egg .Style .StrokeColor !=""{_gag .Add_B ();}else if _egg .Style .FillColor !=""{_gag .Add_f ();}else if _egg .Style .StrokeColor !=""{_gag .Add_S ();};_gag .Add_h ();_gag .Add_Q ();
};func _fcfa (_gfdd string )([]float64 ,error ){_aaad :=-1;var _cgac []float64 ;_aade :=' ';for _acdc ,_ede :=range _gfdd {if !_bc .IsNumber (_ede )&&_ede !='.'&&!(_ede =='-'&&_aade =='e')&&_ede !='e'{if _aaad !=-1{_dca ,_abbb :=_gcba (_gfdd [_aaad :_acdc ]);
if _abbb !=nil {return _cgac ,_abbb ;};_cgac =append (_cgac ,_dca ...);};if _ede =='-'{_aaad =_acdc ;}else {_aaad =-1;};}else if _aaad ==-1{_aaad =_acdc ;};_aade =_ede ;};if _aaad !=-1&&_aaad !=len (_gfdd ){_caded ,_dcfe :=_gcba (_gfdd [_aaad :]);if _dcfe !=nil {return _cgac ,_dcfe ;
};_cgac =append (_cgac ,_caded ...);};return _cgac ,nil ;};var (_bg =[]string {"\u0063\u006d","\u006d\u006d","\u0070\u0078","\u0070\u0074"};_bdc =map[string ]float64 {"\u0063\u006d":_cf ,"\u006d\u006d":_baf ,"\u0070\u0078":_ba ,"\u0070\u0074":1};);func _cgaec (_beea string )(_ecd ,_ebca ,_cda float64 ){if (len (_beea )!=4&&len (_beea )!=7)||_beea [0]!='#'{_fc .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_beea );
return _ecd ,_ebca ,_cda ;};var _cbc ,_afa ,_ecbbc int ;if len (_beea )==4{var _gee ,_dbff ,_daf int ;_cabc ,_ceca :=_g .Sscanf (_beea ,"\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078",&_gee ,&_dbff ,&_daf );if _ceca !=nil {_fc .Log .Debug ("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076",_beea ,_ceca );
return _ecd ,_ebca ,_cda ;};if _cabc !=3{_fc .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_beea );return _ecd ,_ebca ,_cda ;};_cbc =_gee *16+_gee ;_afa =_dbff *16+_dbff ;_ecbbc =_daf *16+_daf ;
}else {_cgbf ,_acf :=_g .Sscanf (_beea ,"\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078",&_cbc ,&_afa ,&_ecbbc );if _acf !=nil {_fc .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_beea );
return _ecd ,_ebca ,_cda ;};if _cgbf !=3{_fc .Log .Debug ("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029",_beea ,_cgbf );
return _ecd ,_ebca ,_cda ;};};_gbfb :=float64 (_cbc )/255.0;_aeed :=float64 (_afa )/255.0;_bcbff :=float64 (_ecbbc )/255.0;return _gbfb ,_aeed ,_bcbff ;};func ParseFromStream (source _f .Reader )(*GraphicSVG ,error ){_eef :=_gg .NewDecoder (source );_eef .CharsetReader =_eb .NewReaderLabel ;
_fdeg ,_fgg :=_gfa (_eef );if _fgg !=nil {return nil ,_fgg ;};if _ege :=_fdeg .Decode (_eef );_ege !=nil &&_ege !=_f .EOF {return nil ,_ege ;};return _fdeg ,nil ;};func _gcba (_ead string )(_gffd []float64 ,_acab error ){var _beba float64 ;_dddf :=0;_fdfb :=true ;
for _eda ,_eabd :=range _ead {if _eabd =='.'{if _fdfb {_fdfb =false ;continue ;};_beba ,_acab =_fga (_ead [_dddf :_eda ],64);if _acab !=nil {return ;};_gffd =append (_gffd ,_beba );_dddf =_eda ;};};_beba ,_acab =_fga (_ead [_dddf :],64);if _acab !=nil {return ;
};_gffd =append (_gffd ,_beba );return ;};const (_ba =0.72;_cf =28.3464;_baf =_cf /10;_ggg =0.551784;);func _cadg (_fcg _gg .StartElement )*GraphicSVG {_bdg :=&GraphicSVG {};_fdd :=make (map[string ]string );for _ ,_dff :=range _fcg .Attr {_fdd [_dff .Name .Local ]=_dff .Value ;
};_bdg .Name =_fcg .Name .Local ;_bdg .Attributes =_fdd ;_bdg ._faf =1;if _bdg .Name =="\u0073\u0076\u0067"{_dgd ,_afef :=_fcfa (_fdd ["\u0076i\u0065\u0077\u0042\u006f\u0078"]);if _afef !=nil {_fc .Log .Debug ("\u0055\u006ea\u0062\u006c\u0065\u0020t\u006f\u0020p\u0061\u0072\u0073\u0065\u0020\u0076\u0069\u0065w\u0042\u006f\u0078\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074e\u003a\u0020\u0025\u0076",_afef );
return nil ;};_bdg .ViewBox .X =_dgd [0];_bdg .ViewBox .Y =_dgd [1];_bdg .ViewBox .W =_dgd [2];_bdg .ViewBox .H =_dgd [3];_bdg .Width =_bdg .ViewBox .W ;_bdg .Height =_bdg .ViewBox .H ;if _dcf ,_bed :=_fdd ["\u0077\u0069\u0064t\u0068"];_bed {_gce ,_abd :=_fga (_dcf ,64);
if _abd !=nil {_fc .Log .Debug ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0077\u0069\u0064\u0074\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_abd );return nil ;};_bdg .Width =_gce ;
};if _ddg ,_cgd :=_fdd ["\u0068\u0065\u0069\u0067\u0068\u0074"];_cgd {_cabd ,_aeeg :=_fga (_ddg ,64);if _aeeg !=nil {_fc .Log .Debug ("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0070\u0061rse\u0020he\u0069\u0067\u0068\u0074\u0020\u0061\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_aeeg );
return nil ;};_bdg .Height =_cabd ;};if _bdg .Width > 0&&_bdg .Height > 0{_bdg ._faf =_bdg .Width /_bdg .ViewBox .W ;};};return _bdg ;};func _adg (_egge string )[]token {var (_effa []token ;_ddd string ;);for _ ,_bbe :=range _egge {_agad :=string (_bbe );
switch {case _egea .isCommand (_agad ):_effa ,_ddd =_cgae (_effa ,_ddd );_effa =append (_effa ,token {_agad ,true });case _agad =="\u002e":if _ddd ==""{_ddd ="\u0030";};if _c .Contains (_ddd ,_agad ){_effa =append (_effa ,token {_ddd ,false });_ddd ="\u0030";
};fallthrough;case _agad >="\u0030"&&_agad <="\u0039"||_agad =="\u0065":_ddd +=_agad ;case _agad =="\u002d":if _c .HasSuffix (_ddd ,"\u0065"){_ddd +=_agad ;}else {_effa ,_ =_cgae (_effa ,_ddd );_ddd =_agad ;};default:_effa ,_ddd =_cgae (_effa ,_ddd );};
};_effa ,_ =_cgae (_effa ,_ddd );return _effa ;};func _dbe (_gefa []*Command )*Path {_eab :=&Path {};var _fega []*Command ;for _cfd ,_gfdb :=range _gefa {switch _c .ToLower (_gfdb .Symbol ){case _egea ._bde :if len (_fega )> 0{_eab .Subpaths =append (_eab .Subpaths ,&Subpath {_fega });
};_fega =[]*Command {_gfdb };case _egea ._gcb :_fega =append (_fega ,_gfdb );_eab .Subpaths =append (_eab .Subpaths ,&Subpath {_fega });_fega =[]*Command {};default:_fega =append (_fega ,_gfdb );if len (_gefa )==_cfd +1{_eab .Subpaths =append (_eab .Subpaths ,&Subpath {_fega });
};};};return _eab ;};func (_beb *GraphicSVG )setDefaultScaling (_gd float64 ){_beb ._faf =_gd ;if _beb .Style !=nil &&_beb .Style .StrokeWidth > 0{_beb .Style .StrokeWidth =_beb .Style .StrokeWidth *_beb ._faf ;};for _ ,_bea :=range _beb .Children {_bea .setDefaultScaling (_gd );
};};type Subpath struct{Commands []*Command ;};func (_gdcb *Subpath )compare (_cfg *Subpath )bool {if len (_gdcb .Commands )!=len (_cfg .Commands ){return false ;};for _fefd ,_dffg :=range _gdcb .Commands {if !_dffg .compare (_cfg .Commands [_fefd ]){return false ;
};};return true ;};func _fcf (_egc *GraphicSVG ,_ccb *_fd .ContentCreator ){_ccb .Add_q ();_egc .Style .toContentStream (_ccb );_gfeb ,_bee :=_fga (_egc .Attributes ["\u0078"],64);if _bee !=nil {_fc .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0078\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_bee .Error ());
};_ea ,_bee :=_fga (_egc .Attributes ["\u0079"],64);if _bee !=nil {_fc .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0079\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_bee .Error ());
};_ebc ,_bee :=_fga (_egc .Attributes ["\u0077\u0069\u0064t\u0068"],64);if _bee !=nil {_fc .Log .Debug ("\u0045\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_bee .Error ());
};_eca ,_bee :=_fga (_egc .Attributes ["\u0068\u0065\u0069\u0067\u0068\u0074"],64);if _bee !=nil {_fc .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0077h\u0069\u006c\u0065 \u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_bee .Error ());
};_ccb .Add_re (_gfeb *_egc ._faf ,_ea *_egc ._faf ,_ebc *_egc ._faf ,_eca *_egc ._faf );if _egc .Style .FillColor !=""&&_egc .Style .StrokeColor !=""{_ccb .Add_B ();}else if _egc .Style .FillColor !=""{_ccb .Add_f ();}else if _egc .Style .StrokeColor !=""{_ccb .Add_S ();
};_ccb .Add_Q ();};func _da ()*GraphicSVGStyle {return &GraphicSVGStyle {FillColor :"\u00230\u0030\u0030\u0030\u0030\u0030",StrokeColor :"",StrokeWidth :0};};func _cea (_edce float64 ,_bebb int )float64 {_fea :=_a .Pow (10,float64 (_bebb ));return float64 (_afdd (_edce *_fea ))/_fea ;
};type GraphicSVG struct{ViewBox struct{X ,Y ,W ,H float64 ;};Name string ;Attributes map[string ]string ;Children []*GraphicSVG ;Content string ;Style *GraphicSVGStyle ;Width float64 ;Height float64 ;_faf float64 ;};func (_edbb *Command )compare (_fcd *Command )bool {if _edbb .Symbol !=_fcd .Symbol {return false ;
};for _fef ,_dad :=range _edbb .Params {if _dad !=_fcd .Params [_fef ]{return false ;};};return true ;};func (_gefe *GraphicSVG )SetScaling (xFactor ,yFactor float64 ){_acda :=_gefe .Width /_gefe .ViewBox .W ;_dfd :=_gefe .Height /_gefe .ViewBox .H ;_gefe .setDefaultScaling (_a .Max (_acda ,_dfd ));
for _ ,_dc :=range _gefe .Children {_dc .SetScaling (xFactor ,yFactor );};};func _bfe (_dcg string )(_agd ,_egb string ){if _dcg ==""||(_dcg [len (_dcg )-1]>='0'&&_dcg [len (_dcg )-1]<='9'){return _dcg ,"";};_agd =_dcg ;for _ ,_dce :=range _bg {if _c .Contains (_agd ,_dce ){_egb =_dce ;
};_agd =_c .TrimSuffix (_agd ,_dce );};return ;};func ParseFromFile (path string )(*GraphicSVG ,error ){_gae ,_afga :=_cd .Open (path );if _afga !=nil {return nil ,_afga ;};defer _gae .Close ();return ParseFromStream (_gae );};func ParseFromString (svgStr string )(*GraphicSVG ,error ){return ParseFromStream (_c .NewReader (svgStr ));
};func (_gaf *commands )isCommand (_cdc string )bool {for _ ,_abb :=range _gaf ._bcbb {if _c .ToLower (_cdc )==_abb {return true ;};};return false ;};func _abac ()commands {var _dece =map[string ]int {"\u006d":2,"\u007a":0,"\u006c":2,"\u0068":1,"\u0076":1,"\u0063":6,"\u0073":4,"\u0071":4,"\u0074":2,"\u0061":7};
var _aff []string ;for _cgef :=range _dece {_aff =append (_aff ,_cgef );};return commands {_aff ,_dece ,"\u006d","\u007a"};};func (_edcc *Path )compare (_bcc *Path )bool {if len (_edcc .Subpaths )!=len (_bcc .Subpaths ){return false ;};for _ageb ,_gfeba :=range _edcc .Subpaths {if !_gfeba .compare (_bcc .Subpaths [_ageb ]){return false ;
};};return true ;};type token struct{_cabf string ;_dcfc bool ;};func _bgfe (_gbf map[string ]string ,_gbdf float64 )(*GraphicSVGStyle ,error ){_ecae :=_da ();_aeec ,_efc :=_gbf ["\u0066\u0069\u006c\u006c"];if _efc {_ecae .FillColor =_aeec ;if _aeec =="\u006e\u006f\u006e\u0065"{_ecae .FillColor ="";
};};_agga ,_daa :=_gbf ["\u0073\u0074\u0072\u006f\u006b\u0065"];if _daa {_ecae .StrokeColor =_agga ;if _agga =="\u006e\u006f\u006e\u0065"{_ecae .StrokeColor ="";};};_fdfe ,_beed :=_gbf ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];
if _beed {_fgcb ,_fddd :=_fga (_fdfe ,64);if _fddd !=nil {return nil ,_fddd ;};_ecae .StrokeWidth =_fgcb *_gbdf ;};return _ecae ,nil ;};func _cgae (_fggf []token ,_ecbb string )([]token ,string ){if _ecbb !=""{_fggf =append (_fggf ,token {_ecbb ,false });
_ecbb ="";};return _fggf ,_ecbb ;};var _egea commands ;type pathParserError struct{_cecf string };func (_ecb pathParserError )Error ()string {return _ecb ._cecf };func _gbeff (_bdbc []token )([]*Command ,error ){var (_fbg []*Command ;_cdd []float64 ;);
for _edcg :=len (_bdbc )-1;_edcg >=0;_edcg --{_ggc :=_bdbc [_edcg ];if _ggc ._dcfc {_fbf :=_egea ._gff [_c .ToLower (_ggc ._cabf )];_adb :=len (_cdd );if _fbf ==0&&_adb ==0{_edbbd :=&Command {Symbol :_ggc ._cabf };_fbg =append ([]*Command {_edbbd },_fbg ...);
}else if _fbf !=0&&_adb %_fbf ==0{_gac :=_adb /_fbf ;for _cffe :=0;_cffe < _gac ;_cffe ++{_cade :=_ggc ._cabf ;if _cade =="\u006d"&&_cffe < _gac -1{_cade ="\u006c";};if _cade =="\u004d"&&_cffe < _gac -1{_cade ="\u004c";};_ceg :=&Command {_cade ,_gfb (_cdd [:_fbf ])};
_fbg =append ([]*Command {_ceg },_fbg ...);_cdd =_cdd [_fbf :];};}else {_dfb :=pathParserError {"I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006fr\u0020"+_ggc ._cabf };
return nil ,_dfb ;};}else {_fgf ,_eee :=_fga (_ggc ._cabf ,64);if _eee !=nil {return nil ,_eee ;};_cdd =append (_cdd ,_fgf );};};return _fbg ,nil ;};func _aad (_bae *GraphicSVG ,_ade *_fd .ContentCreator ){_ade .Add_q ();_bae .Style .toContentStream (_ade );
_cg ,_gfec :=_fga (_bae .Attributes ["\u0063\u0078"],64);if _gfec !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gfec .Error ());
};_ada ,_gfec :=_fga (_bae .Attributes ["\u0063\u0079"],64);if _gfec !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gfec .Error ());
};_fe ,_gfec :=_fga (_bae .Attributes ["\u0072\u0078"],64);if _gfec !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gfec .Error ());
};_aba ,_gfec :=_fga (_bae .Attributes ["\u0072\u0079"],64);if _gfec !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gfec .Error ());
};_fbd :=_fe *_bae ._faf ;_cgb :=_aba *_bae ._faf ;_cga :=_cg *_bae ._faf ;_fa :=_ada *_bae ._faf ;_cad :=_fbd *_ggg ;_ff :=_cgb *_ggg ;_ee :=_bf .NewCubicBezierPath ();_ee =_ee .AppendCurve (_bf .NewCubicBezierCurve (-_fbd ,0,-_fbd ,_ff ,-_cad ,_cgb ,0,_cgb ));
_ee =_ee .AppendCurve (_bf .NewCubicBezierCurve (0,_cgb ,_cad ,_cgb ,_fbd ,_ff ,_fbd ,0));_ee =_ee .AppendCurve (_bf .NewCubicBezierCurve (_fbd ,0,_fbd ,-_ff ,_cad ,-_cgb ,0,-_cgb ));_ee =_ee .AppendCurve (_bf .NewCubicBezierCurve (0,-_cgb ,-_cad ,-_cgb ,-_fbd ,-_ff ,-_fbd ,0));
_ee =_ee .Offset (_cga ,_fa );if _bae .Style .StrokeWidth > 0{_ee =_ee .Offset (_bae .Style .StrokeWidth /2,_bae .Style .StrokeWidth /2);};_bf .DrawBezierPathWithCreator (_ee ,_ade );if _bae .Style .FillColor !=""&&_bae .Style .StrokeColor !=""{_ade .Add_B ();
}else if _bae .Style .FillColor !=""{_ade .Add_f ();}else if _bae .Style .StrokeColor !=""{_ade .Add_S ();};_ade .Add_h ();_ade .Add_Q ();};func _gfb (_eea []float64 )[]float64 {for _beedg ,_bag :=0,len (_eea )-1;_beedg < _bag ;_beedg ,_bag =_beedg +1,_bag -1{_eea [_beedg ],_eea [_bag ]=_eea [_bag ],_eea [_beedg ];
};return _eea ;};func _edc (_agg *GraphicSVG ,_fad *_fd .ContentCreator ){_fad .Add_q ();_agg .Style .toContentStream (_fad );_bdb ,_acd :=_fcfa (_agg .Attributes ["\u0070\u006f\u0069\u006e\u0074\u0073"]);if _acd !=nil {_fc .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076",_acd );
return ;};if len (_bdb )%2> 0{_fc .Log .Debug ("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068");return ;
};for _bfd :=0;_bfd < len (_bdb );{if _bfd ==0{_fad .Add_m (_bdb [_bfd ]*_agg ._faf ,_bdb [_bfd +1]*_agg ._faf );}else {_fad .Add_l (_bdb [_bfd ]*_agg ._faf ,_bdb [_bfd +1]*_agg ._faf );};_bfd +=2;};if _agg .Style .FillColor !=""&&_agg .Style .StrokeColor !=""{_fad .Add_B ();
}else if _agg .Style .FillColor !=""{_fad .Add_f ();}else if _agg .Style .StrokeColor !=""{_fad .Add_S ();};_fad .Add_h ();_fad .Add_Q ();};func _d (_af *GraphicSVG ,_bb *_fd .ContentCreator ){_bb .Add_q ();_af .Style .toContentStream (_bb );_afg ,_fca :=_bbg (_af .Attributes ["\u0064"]);
if _fca !=nil {_fc .Log .Error ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025s",_fca .Error ());};var (_db ,_eba =0.0,0.0;_ec ,_bcb =0.0,0.0;_fce *Command ;);for _ ,_gb :=range _afg .Subpaths {for _ ,_be :=range _gb .Commands {switch _c .ToLower (_be .Symbol ){case "\u006d":_ec ,_bcb =_be .Params [0]*_af ._faf ,_be .Params [1]*_af ._faf ;
if !_be .isAbsolute (){_ec ,_bcb =_db +_ec -_af .ViewBox .X ,_eba +_bcb -_af .ViewBox .Y ;};_bb .Add_m (_cea (_ec ,3),_cea (_bcb ,3));_db ,_eba =_ec ,_bcb ;case "\u0063":_adf ,_ca ,_ge ,_fcc ,_gf ,_cc :=_be .Params [0]*_af ._faf ,_be .Params [1]*_af ._faf ,_be .Params [2]*_af ._faf ,_be .Params [3]*_af ._faf ,_be .Params [4]*_af ._faf ,_be .Params [5]*_af ._faf ;
if !_be .isAbsolute (){_adf ,_ca ,_ge ,_fcc ,_gf ,_cc =_db +_adf ,_eba +_ca ,_db +_ge ,_eba +_fcc ,_db +_gf ,_eba +_cc ;};_bb .Add_c (_cea (_adf ,3),_cea (_ca ,3),_cea (_ge ,3),_cea (_fcc ,3),_cea (_gf ,3),_cea (_cc ,3));_db ,_eba =_gf ,_cc ;case "\u0073":_gfe ,_dd ,_ab ,_bce :=_be .Params [0]*_af ._faf ,_be .Params [1]*_af ._faf ,_be .Params [2]*_af ._faf ,_be .Params [3]*_af ._faf ;
if !_be .isAbsolute (){_gfe ,_dd ,_ab ,_bce =_db +_gfe ,_eba +_dd ,_db +_ab ,_eba +_bce ;};_bb .Add_c (_cea (_db ,3),_cea (_eba ,3),_cea (_gfe ,3),_cea (_dd ,3),_cea (_ab ,3),_cea (_bce ,3));_db ,_eba =_ab ,_bce ;case "\u006c":_fg ,_ac :=_be .Params [0]*_af ._faf ,_be .Params [1]*_af ._faf ;
if !_be .isAbsolute (){_fg ,_ac =_db +_fg ,_eba +_ac ;};_bb .Add_l (_cea (_fg ,3),_cea (_ac ,3));_db ,_eba =_fg ,_ac ;case "\u0068":_eg :=_be .Params [0]*_af ._faf ;if !_be .isAbsolute (){_eg =_db +_eg ;};_bb .Add_l (_cea (_eg ,3),_cea (_eba ,3));_db =_eg ;
case "\u0076":_dg :=_be .Params [0]*_af ._faf ;if !_be .isAbsolute (){_dg =_eba +_dg ;};_bb .Add_l (_cea (_db ,3),_cea (_dg ,3));_eba =_dg ;case "\u0071":_gfg ,_bcbf ,_ed ,_fgc :=_be .Params [0]*_af ._faf ,_be .Params [1]*_af ._faf ,_be .Params [2]*_af ._faf ,_be .Params [3]*_af ._faf ;
if !_be .isAbsolute (){_gfg ,_bcbf ,_ed ,_fgc =_db +_gfg ,_eba +_bcbf ,_db +_ed ,_eba +_fgc ;};_beg ,_gbd :=_bd .QuadraticToCubicBezier (_db ,_eba ,_gfg ,_bcbf ,_ed ,_fgc );_bb .Add_c (_cea (_beg .X ,3),_cea (_beg .Y ,3),_cea (_gbd .X ,3),_cea (_gbd .Y ,3),_cea (_ed ,3),_cea (_fgc ,3));
_db ,_eba =_ed ,_fgc ;case "\u0074":var _edd ,_de _bd .Point ;_aa ,_geb :=_be .Params [0]*_af ._faf ,_be .Params [1]*_af ._faf ;if !_be .isAbsolute (){_aa ,_geb =_db +_aa ,_eba +_geb ;};if _fce !=nil &&_c .ToLower (_fce .Symbol )=="\u0071"{_aaa :=_bd .Point {X :_fce .Params [0]*_af ._faf ,Y :_fce .Params [1]*_af ._faf };
_fb :=_bd .Point {X :_fce .Params [2]*_af ._faf ,Y :_fce .Params [3]*_af ._faf };_abf :=_fb .Mul (2.0).Sub (_aaa );_edd ,_de =_bd .QuadraticToCubicBezier (_db ,_eba ,_abf .X ,_abf .Y ,_aa ,_geb );};_bb .Add_c (_cea (_edd .X ,3),_cea (_edd .Y ,3),_cea (_de .X ,3),_cea (_de .Y ,3),_cea (_aa ,3),_cea (_geb ,3));
_db ,_eba =_aa ,_geb ;case "\u0061":_cfc ,_edb :=_be .Params [0]*_af ._faf ,_be .Params [1]*_af ._faf ;_bgg :=_be .Params [2];_ag :=_be .Params [3]> 0;_abg :=_be .Params [4]> 0;_dbg ,_fba :=_be .Params [5]*_af ._faf ,_be .Params [6]*_af ._faf ;if !_be .isAbsolute (){_dbg ,_fba =_db +_dbg ,_eba +_fba ;
};_gbe :=_bd .EllipseToCubicBeziers (_db ,_eba ,_cfc ,_edb ,_bgg ,_ag ,_abg ,_dbg ,_fba );for _ ,_ce :=range _gbe {_bb .Add_c (_cea (_ce [1].X ,3),_cea ((_ce [1].Y ),3),_cea ((_ce [2].X ),3),_cea ((_ce [2].Y ),3),_cea ((_ce [3].X ),3),_cea ((_ce [3].Y ),3));
};_db ,_eba =_dbg ,_fba ;case "\u007a":_bb .Add_h ();};_fce =_be ;};};if _af .Style .FillColor !=""&&_af .Style .StrokeColor !=""{_bb .Add_B ();}else if _af .Style .FillColor !=""{_bb .Add_f ();}else if _af .Style .StrokeColor !=""{_bb .Add_S ();};_bb .Add_h ();
_bb .Add_Q ();};func _afdd (_aae float64 )int {return int (_aae +_a .Copysign (0.5,_aae ))};func (_cece *GraphicSVG )toContentStream (_ef *_fd .ContentCreator ){_eggg ,_dbf :=_bgfe (_cece .Attributes ,_cece ._faf );if _dbf !=nil {_fc .Log .Debug ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_dbf );
};_cece .Style =_eggg ;switch _cece .Name {case "\u0070\u0061\u0074\u0068":_d (_cece ,_ef );for _ ,_gbb :=range _cece .Children {_gbb .toContentStream (_ef );};case "\u0072\u0065\u0063\u0074":_fcf (_cece ,_ef );for _ ,_gef :=range _cece .Children {_gef .toContentStream (_ef );
};case "\u0063\u0069\u0072\u0063\u006c\u0065":_bdcg (_cece ,_ef );for _ ,_cfa :=range _cece .Children {_cfa .toContentStream (_ef );};case "\u0065l\u006c\u0069\u0070\u0073\u0065":_aad (_cece ,_ef );for _ ,_cfe :=range _cece .Children {_cfe .toContentStream (_ef );
};case "\u0070\u006f\u006c\u0079\u006c\u0069\u006e\u0065":_edc (_cece ,_ef );for _ ,_faa :=range _cece .Children {_faa .toContentStream (_ef );};case "\u0070o\u006c\u0079\u0067\u006f\u006e":_cec (_cece ,_ef );for _ ,_fde :=range _cece .Children {_fde .toContentStream (_ef );
};case "\u006c\u0069\u006e\u0065":_cgg (_cece ,_ef );for _ ,_ae :=range _cece .Children {_ae .toContentStream (_ef );};case "\u0067":_gbc ,_bba :=_cece .Attributes ["\u0066\u0069\u006c\u006c"];_age ,_cffb :=_cece .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"];
_ggd ,_cce :=_cece .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];for _ ,_caf :=range _cece .Children {if _ ,_cba :=_caf .Attributes ["\u0066\u0069\u006c\u006c"];!_cba &&_bba {_caf .Attributes ["\u0066\u0069\u006c\u006c"]=_gbc ;
};if _ ,_cab :=_caf .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"];!_cab &&_cffb {_caf .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"]=_age ;};if _ ,_edg :=_caf .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];
!_edg &&_cce {_caf .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]=_ggd ;};_caf .toContentStream (_ef );};};};func _bdcg (_bfa *GraphicSVG ,_gfc *_fd .ContentCreator ){_gfc .Add_q ();_bfa .Style .toContentStream (_gfc );
_ccd ,_ga :=_fga (_bfa .Attributes ["\u0063\u0078"],64);if _ga !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_ga .Error ());
};_aca ,_ga :=_fga (_bfa .Attributes ["\u0063\u0079"],64);if _ga !=nil {_fc .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_ga .Error ());
};_afe ,_ga :=_fga (_bfa .Attributes ["\u0072"],64);if _ga !=nil {_fc .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0072\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_ga .Error ());
};_fgb :=_afe *_bfa ._faf ;_dgc :=_afe *_bfa ._faf ;_cff :=_fgb *_ggg ;_eac :=_dgc *_ggg ;_cde :=_bf .NewCubicBezierPath ();_cde =_cde .AppendCurve (_bf .NewCubicBezierCurve (-_fgb ,0,-_fgb ,_eac ,-_cff ,_dgc ,0,_dgc ));_cde =_cde .AppendCurve (_bf .NewCubicBezierCurve (0,_dgc ,_cff ,_dgc ,_fgb ,_eac ,_fgb ,0));
_cde =_cde .AppendCurve (_bf .NewCubicBezierCurve (_fgb ,0,_fgb ,-_eac ,_cff ,-_dgc ,0,-_dgc ));_cde =_cde .AppendCurve (_bf .NewCubicBezierCurve (0,-_dgc ,-_cff ,-_dgc ,-_fgb ,-_eac ,-_fgb ,0));_cde =_cde .Offset (_ccd *_bfa ._faf ,_aca *_bfa ._faf );
if _bfa .Style .StrokeWidth > 0{_cde =_cde .Offset (_bfa .Style .StrokeWidth /2,_bfa .Style .StrokeWidth /2);};_bf .DrawBezierPathWithCreator (_cde ,_gfc );if _bfa .Style .FillColor !=""&&_bfa .Style .StrokeColor !=""{_gfc .Add_B ();}else if _bfa .Style .FillColor !=""{_gfc .Add_f ();
}else if _bfa .Style .StrokeColor !=""{_gfc .Add_S ();};_gfc .Add_h ();_gfc .Add_Q ();};type Command struct{Symbol string ;Params []float64 ;};func _fga (_gcbb string ,_gba int )(float64 ,error ){_bbc ,_cgacf :=_bfe (_gcbb );_edfd ,_aac :=_e .ParseFloat (_bbc ,_gba );
if _aac !=nil {return 0,_aac ;};if _dffe ,_bafb :=_bdc [_cgacf ];_bafb {_edfd =_edfd *_dffe ;}else {_edfd =_edfd *_ba ;};return _edfd ,nil ;};func _bbg (_gebb string )(*Path ,error ){_egea =_abac ();_fdfd ,_bggc :=_gbeff (_adg (_gebb ));if _bggc !=nil {return nil ,_bggc ;
};return _dbe (_fdfd ),nil ;};func (_aafe *GraphicSVGStyle )toContentStream (_eag *_fd .ContentCreator ){if _aafe ==nil {return ;};if _aafe .FillColor !=""{var _dbgb ,_eff ,_bfg float64 ;if _eacb ,_ccbg :=_bd .ColorMap [_aafe .FillColor ];_ccbg {_ecf ,_fgbg ,_bad ,_ :=_eacb .RGBA ();
_dbgb ,_eff ,_bfg =float64 (_ecf ),float64 (_fgbg ),float64 (_bad );}else {_dbgb ,_eff ,_bfg =_cgaec (_aafe .FillColor );};_eag .Add_rg (_dbgb ,_eff ,_bfg );};if _aafe .StrokeColor !=""{var _gfd ,_fdf ,_edf float64 ;if _ega ,_fff :=_bd .ColorMap [_aafe .StrokeColor ];
_fff {_cgaf ,_effb ,_cafe ,_ :=_ega .RGBA ();_gfd ,_fdf ,_edf =float64 (_cgaf )/255.0,float64 (_effb )/255.0,float64 (_cafe )/255.0;}else {_gfd ,_fdf ,_edf =_cgaec (_aafe .StrokeColor );};_eag .Add_RG (_gfd ,_fdf ,_edf );};if _aafe .StrokeWidth > 0{_eag .Add_w (_aafe .StrokeWidth );
};};func _cec (_bbd *GraphicSVG ,_eaa *_fd .ContentCreator ){_eaa .Add_q ();_bbd .Style .toContentStream (_eaa );_fac ,_gbef :=_fcfa (_bbd .Attributes ["\u0070\u006f\u0069\u006e\u0074\u0073"]);if _gbef !=nil {_fc .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076",_gbef );
return ;};if len (_fac )%2> 0{_fc .Log .Debug ("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068");return ;
};for _cffa :=0;_cffa < len (_fac );{if _cffa ==0{_eaa .Add_m (_fac [_cffa ]*_bbd ._faf ,_fac [_cffa +1]*_bbd ._faf );}else {_eaa .Add_l (_fac [_cffa ]*_bbd ._faf ,_fac [_cffa +1]*_bbd ._faf );};_cffa +=2;};_eaa .Add_l (_fac [0]*_bbd ._faf ,_fac [1]*_bbd ._faf );
if _bbd .Style .FillColor !=""&&_bbd .Style .StrokeColor !=""{_eaa .Add_B ();}else if _bbd .Style .FillColor !=""{_eaa .Add_f ();}else if _bbd .Style .StrokeColor !=""{_eaa .Add_S ();};_eaa .Add_h ();_eaa .Add_Q ();};type Path struct{Subpaths []*Subpath ;
};type GraphicSVGStyle struct{FillColor string ;StrokeColor string ;StrokeWidth float64 ;};func (_bca *Command )isAbsolute ()bool {return _bca .Symbol ==_c .ToUpper (_bca .Symbol )};func _gfa (_bgc *_gg .Decoder )(*GraphicSVG ,error ){for {_feg ,_agb :=_bgc .Token ();
if _feg ==nil &&_agb ==_f .EOF {break ;};if _agb !=nil {return nil ,_agb ;};switch _cge :=_feg .(type ){case _gg .StartElement :return _cadg (_cge ),nil ;};};return &GraphicSVG {},nil ;};func (_fbac *GraphicSVG )ToContentCreator (cc *_fd .ContentCreator )*_fd .ContentCreator {if _fbac .Name =="\u0073\u0076\u0067"{cc .Add_cm (1,0,0,1,0,0);
_fbac .setDefaultScaling (_fbac ._faf );for _ ,_aaf :=range _fbac .Children {_aaf .ViewBox =_fbac .ViewBox ;_aaf .toContentStream (cc );};return cc ;};return nil ;};