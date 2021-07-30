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

package optimize ;import (_eaf "bytes";_ge "crypto/md5";_f "errors";_b "github.com/unidoc/unipdf/v3/common";_dg "github.com/unidoc/unipdf/v3/contentstream";_g "github.com/unidoc/unipdf/v3/core";_c "github.com/unidoc/unipdf/v3/extractor";_ae "github.com/unidoc/unipdf/v3/internal/imageutil";
_ed "github.com/unidoc/unipdf/v3/internal/textencoding";_bc "github.com/unidoc/unipdf/v3/model";_ea "github.com/unidoc/unitype";_a "golang.org/x/image/draw";_e "math";);type content struct{_cdd string ;_ecb *_bc .PdfPageResources ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_gfcb *ImagePPI )Optimize (objects []_g .PdfObject )(_dedf []_g .PdfObject ,_gbg error ){if _gfcb .ImageUpperPPI <=0{return objects ,nil ;};_caee :=_ggg (objects );if len (_caee )==0{return objects ,nil ;};_eaa :=make (map[_g .PdfObject ]struct{});
for _ ,_gbgb :=range _caee {_fda :=_gbgb .Stream .PdfObjectDictionary .Get ("\u0053\u004d\u0061s\u006b");_eaa [_fda ]=struct{}{};};_gbfb :=make (map[*_g .PdfObjectStream ]*imageInfo );for _ ,_aeeb :=range _caee {_gbfb [_aeeb .Stream ]=_aeeb ;};var _ffc *_g .PdfObjectDictionary ;
for _ ,_decg :=range objects {if _gcdc ,_ffde :=_g .GetDict (_decg );_ffc ==nil &&_ffde {if _eca ,_agca :=_g .GetName (_gcdc .Get ("\u0054\u0079\u0070\u0065"));_agca &&*_eca =="\u0043a\u0074\u0061\u006c\u006f\u0067"{_ffc =_gcdc ;};};};if _ffc ==nil {return objects ,nil ;
};_ffcg ,_aafab :=_g .GetDict (_ffc .Get ("\u0050\u0061\u0067e\u0073"));if !_aafab {return objects ,nil ;};_dbba ,_dgea :=_g .GetArray (_ffcg .Get ("\u004b\u0069\u0064\u0073"));if !_dgea {return objects ,nil ;};for _ ,_egcg :=range _dbba .Elements (){_egf :=make (map[string ]*imageInfo );
_dgcg ,_cddd :=_g .GetDict (_egcg );if !_cddd {continue ;};_dee ,_ :=_eeec (_dgcg .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if len (_dee )==0{continue ;};_bega ,_faad :=_g .GetDict (_dgcg .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));
if !_faad {continue ;};_dba ,_bedc :=_bc .NewPdfPageResourcesFromDict (_bega );if _bedc !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020-\u0020\u0069\u0067\u006e\u006fr\u0069\u006eg\u003a\u0020\u0025\u0076",_bedc );
continue ;};_afc ,_fdbf :=_g .GetDict (_bega .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));if !_fdbf {continue ;};_cgcg :=_afc .Keys ();for _ ,_bdg :=range _cgcg {if _adbc ,_egda :=_g .GetStream (_afc .Get (_bdg ));_egda {if _dbe ,_dcba :=_gbfb [_adbc ];
_dcba {_egf [string (_bdg )]=_dbe ;};};};_beda :=_dg .NewContentStreamParser (_dee );_cfcf ,_bedc :=_beda .Parse ();if _bedc !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_bedc );return nil ,_bedc ;};_dag :=_dg .NewContentStreamProcessor (*_cfcf );
_dag .AddHandler (_dg .HandlerConditionEnumAllOperands ,"",func (_bcgf *_dg .ContentStreamOperation ,_dagf _dg .GraphicsState ,_ggfe *_bc .PdfPageResources )error {switch _bcgf .Operand {case "\u0044\u006f":if len (_bcgf .Params )!=1{_b .Log .Debug ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020w\u0069\u0074\u0068\u0020\u006c\u0065\u006e\u0028\u0070\u0061ra\u006d\u0073\u0029 \u0021=\u0020\u0031");
return nil ;};_cde ,_aec :=_g .GetName (_bcgf .Params [0]);if !_aec {_b .Log .Debug ("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u006e\u0020\u004e\u0061\u006d\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072");
return nil ;};if _adef ,_feg :=_egf [string (*_cde )];_feg {_cdcf :=_dagf .CTM .ScalingFactorX ();_cdfa :=_dagf .CTM .ScalingFactorY ();_gbdf ,_gcf :=_cdcf /72.0,_cdfa /72.0;_efa ,_gdaeb :=float64 (_adef .Width )/_gbdf ,float64 (_adef .Height )/_gcf ;if _gbdf ==0||_gcf ==0{_efa =72.0;
_gdaeb =72.0;};_adef .PPI =_e .Max (_adef .PPI ,_efa );_adef .PPI =_e .Max (_adef .PPI ,_gdaeb );};};return nil ;});_bedc =_dag .Process (_dba );if _bedc !=nil {_b .Log .Debug ("E\u0052\u0052\u004f\u0052 p\u0072o\u0063\u0065\u0073\u0073\u0069n\u0067\u003a\u0020\u0025\u002b\u0076",_bedc );
return nil ,_bedc ;};};for _ ,_bgab :=range _caee {if _ ,_edca :=_eaa [_bgab .Stream ];_edca {continue ;};if _bgab .PPI <=_gfcb .ImageUpperPPI {continue ;};_efag ,_gbeg :=_bc .NewXObjectImageFromStream (_bgab .Stream );if _gbeg !=nil {return nil ,_gbeg ;
};var _bggg imageModifications ;_bggg .Scale =_gfcb .ImageUpperPPI /_bgab .PPI ;if _bgab .BitsPerComponent ==1&&_bgab .ColorComponents ==1{_bcb :=_e .Round (_bgab .PPI /_gfcb .ImageUpperPPI );_dde :=_ae .NextPowerOf2 (uint (_bcb ));if _ae .InDelta (float64 (_dde ),1/_bggg .Scale ,0.3){_bggg .Scale =float64 (1)/float64 (_dde );
};if _ ,_bca :=_efag .Filter .(*_g .JBIG2Encoder );!_bca {_bggg .Encoding =_g .NewJBIG2Encoder ();};};if _gbeg =_ccfg (_efag ,_bggg );_gbeg !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_gbeg );
continue ;};_bggg .Encoding =nil ;if _fcg ,_adfa :=_g .GetStream (_bgab .Stream .PdfObjectDictionary .Get ("\u0053\u004d\u0061s\u006b"));_adfa {_fbac ,_fddb :=_bc .NewXObjectImageFromStream (_fcg );if _fddb !=nil {return nil ,_fddb ;};if _fddb =_ccfg (_fbac ,_bggg );
_fddb !=nil {return nil ,_fddb ;};};};return objects ,nil ;};func _eeec (_bcdg _g .PdfObject )(_bcbb string ,_bcff []_g .PdfObject ){var _gefb _eaf .Buffer ;switch _ecde :=_bcdg .(type ){case *_g .PdfIndirectObject :_bcff =append (_bcff ,_ecde );_bcdg =_ecde .PdfObject ;
};switch _gddab :=_bcdg .(type ){case *_g .PdfObjectStream :if _bfdg ,_cfea :=_g .DecodeStream (_gddab );_cfea ==nil {_gefb .Write (_bfdg );_bcff =append (_bcff ,_gddab );};case *_g .PdfObjectArray :for _ ,_ffb :=range _gddab .Elements (){switch _fddf :=_ffb .(type ){case *_g .PdfObjectStream :if _fag ,_eec :=_g .DecodeStream (_fddf );
_eec ==nil {_gefb .Write (_fag );_bcff =append (_bcff ,_fddf );};};};};return _gefb .String (),_bcff ;};func _cfe (_cd *_g .PdfObjectStream )error {_ac ,_cgb :=_g .DecodeStream (_cd );if _cgb !=nil {return _cgb ;};_fa :=_dg .NewContentStreamParser (string (_ac ));
_gd ,_cgb :=_fa .Parse ();if _cgb !=nil {return _cgb ;};_gd =_cf (_gd );_aeg :=_gd .Bytes ();if len (_aeg )>=len (_ac ){return nil ;};_ec ,_cgb :=_g .MakeStream (_gd .Bytes (),_g .NewFlateEncoder ());if _cgb !=nil {return _cgb ;};_cd .Stream =_ec .Stream ;
_cd .Merge (_ec .PdfObjectDictionary );return nil ;};func _edb (_bcaa []_g .PdfObject ){for _cagc ,_aaed :=range _bcaa {switch _gfeb :=_aaed .(type ){case *_g .PdfIndirectObject :_gfeb .ObjectNumber =int64 (_cagc +1);_gfeb .GenerationNumber =0;case *_g .PdfObjectStream :_gfeb .ObjectNumber =int64 (_cagc +1);
_gfeb .GenerationNumber =0;case *_g .PdfObjectStreams :_gfeb .ObjectNumber =int64 (_cagc +1);_gfeb .GenerationNumber =0;};};};func _ggg (_bafg []_g .PdfObject )[]*imageInfo {_agbe :=_g .PdfObjectName ("\u0053u\u0062\u0074\u0079\u0070\u0065");_daad :=make (map[*_g .PdfObjectStream ]struct{});
var _gdfg []*imageInfo ;for _ ,_bcfg :=range _bafg {_dfa ,_bagg :=_g .GetStream (_bcfg );if !_bagg {continue ;};if _ ,_egb :=_daad [_dfa ];_egb {continue ;};_daad [_dfa ]=struct{}{};_agc :=_dfa .PdfObjectDictionary .Get (_agbe );_cgc ,_bagg :=_g .GetName (_agc );
if !_bagg ||string (*_cgc )!="\u0049\u006d\u0061g\u0065"{continue ;};_bafa :=&imageInfo {Stream :_dfa ,BitsPerComponent :8};if _ccf ,_fce :=_g .GetIntVal (_dfa .Get ("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"));
_fce {_bafa .BitsPerComponent =_ccf ;};if _dbce ,_abfe :=_g .GetIntVal (_dfa .Get ("\u0057\u0069\u0064t\u0068"));_abfe {_bafa .Width =_dbce ;};if _bafc ,_aded :=_g .GetIntVal (_dfa .Get ("\u0048\u0065\u0069\u0067\u0068\u0074"));_aded {_bafa .Height =_bafc ;
};_bbad ,_gfe :=_bc .NewPdfColorspaceFromPdfObject (_dfa .Get ("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"));if _gfe !=nil {_b .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_gfe );continue ;};if _bbad ==nil {_aff ,_bga :=_g .GetName (_dfa .Get ("\u0046\u0069\u006c\u0074\u0065\u0072"));
if _bga {switch _aff .String (){case "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065","J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065":_bbad =_bc .NewPdfColorspaceDeviceGray ();_bafa .BitsPerComponent =1;
};};};switch _dcgf :=_bbad .(type ){case *_bc .PdfColorspaceDeviceRGB :_bafa .ColorComponents =3;case *_bc .PdfColorspaceDeviceGray :_bafa .ColorComponents =1;default:_b .Log .Debug ("\u004f\u0070\u0074\u0069\u006d\u0069\u007aa\u0074\u0069\u006fn\u0020\u0069\u0073 \u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006fr\u0074ed\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u006b\u0069\u0070",_dcgf );
continue ;};_gdfg =append (_gdfg ,_bafa );};return _gdfg ;};func _gebfd (_bfa []_g .PdfObject ,_fgde map[_g .PdfObject ]_g .PdfObject ){if len (_fgde )==0{return ;};for _fbdf ,_babd :=range _bfa {if _cce ,_fabg :=_fgde [_babd ];_fabg {_bfa [_fbdf ]=_cce ;
continue ;};_fgde [_babd ]=_babd ;switch _fade :=_babd .(type ){case *_g .PdfObjectArray :_bad :=make ([]_g .PdfObject ,_fade .Len ());copy (_bad ,_fade .Elements ());_gebfd (_bad ,_fgde );for _acb ,_gfge :=range _bad {_fade .Set (_acb ,_gfge );};case *_g .PdfObjectStreams :_gebfd (_fade .Elements (),_fgde );
case *_g .PdfObjectStream :_edge :=[]_g .PdfObject {_fade .PdfObjectDictionary };_gebfd (_edge ,_fgde );_fade .PdfObjectDictionary =_edge [0].(*_g .PdfObjectDictionary );case *_g .PdfObjectDictionary :_bedg :=_fade .Keys ();_ccg :=make ([]_g .PdfObject ,len (_bedg ));
for _cag ,_gdda :=range _bedg {_ccg [_cag ]=_fade .Get (_gdda );};_gebfd (_ccg ,_fgde );for _bgd ,_fgef :=range _bedg {_fade .Set (_fgef ,_ccg [_bgd ]);};case *_g .PdfIndirectObject :_bbcd :=[]_g .PdfObject {_fade .PdfObject };_gebfd (_bbcd ,_fgde );_fade .PdfObject =_bbcd [0];
};};};type imageModifications struct{Scale float64 ;Encoding _g .StreamEncoder ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_gcdf *CombineDuplicateStreams )Optimize (objects []_g .PdfObject )(_efgf []_g .PdfObject ,_ffad error ){_fff :=make (map[_g .PdfObject ]_g .PdfObject );_fbbe :=make (map[_g .PdfObject ]struct{});_gac :=make (map[string ][]*_g .PdfObjectStream );
for _ ,_gba :=range objects {if _cdf ,_efe :=_gba .(*_g .PdfObjectStream );_efe {_fdf :=_ge .New ();_fdf .Write (_cdf .Stream );_fdf .Write ([]byte (_cdf .PdfObjectDictionary .WriteString ()));_dcd :=string (_fdf .Sum (nil ));_gac [_dcd ]=append (_gac [_dcd ],_cdf );
};};for _ ,_gga :=range _gac {if len (_gga )< 2{continue ;};_aag :=_gga [0];for _cfdd :=1;_cfdd < len (_gga );_cfdd ++{_gbeb :=_gga [_cfdd ];_fff [_gbeb ]=_aag ;_fbbe [_gbeb ]=struct{}{};};};_efgf =make ([]_g .PdfObject ,0,len (objects )-len (_fbbe ));
for _ ,_ecbf :=range objects {if _ ,_efgfg :=_fbbe [_ecbf ];_efgfg {continue ;};_efgf =append (_efgf ,_ecbf );};_gebfd (_efgf ,_fff );return _efgf ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_dcc *CleanFonts )Optimize (objects []_g .PdfObject )(_dcb []_g .PdfObject ,_gdaf error ){var _bdd map[*_g .PdfObjectStream ]struct{};if _dcc .Subset {var _baf error ;_bdd ,_baf =_gda (objects );if _baf !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004fR\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0073u\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u003a \u0025\u0076",_baf );
return nil ,_baf ;};};for _ ,_bfe :=range objects {_dfc ,_dbc :=_g .GetStream (_bfe );if !_dbc {continue ;};if _ ,_ece :=_bdd [_dfc ];_ece {continue ;};_aaeb ,_fbbb :=_g .NewEncoderFromStream (_dfc );if _fbbb !=nil {_b .Log .Debug ("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067",_fbbb );
continue ;};_bfc ,_fbbb :=_aaeb .DecodeStream (_dfc );if _fbbb !=nil {_b .Log .Debug ("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067",_fbbb );
continue ;};if len (_bfc )< 4{continue ;};_ade :=string (_bfc [:4]);if _ade =="\u004f\u0054\u0054\u004f"{continue ;};if _ade !="\u0000\u0001\u0000\u0000"&&_ade !="\u0074\u0072\u0075\u0065"{continue ;};_bag ,_fbbb :=_ea .Parse (_eaf .NewReader (_bfc ));
if _fbbb !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_fbbb );continue ;};_fbbb =_bag .Optimize ();
if _fbbb !=nil {_b .Log .Debug ("\u0045\u0052RO\u0052\u0020\u004fp\u0074\u0069\u006d\u0069zin\u0067 f\u006f\u006e\u0074\u003a\u0020\u0025\u0076 -\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067",_fbbb );continue ;};var _afg _eaf .Buffer ;_fbbb =_bag .Write (&_afg );
if _fbbb !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_fbbb );continue ;};if _afg .Len ()> len (_bfc ){_b .Log .Debug ("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070");
continue ;};_cabf ,_fbbb :=_g .MakeStream (_afg .Bytes (),_g .NewFlateEncoder ());if _fbbb !=nil {continue ;};*_dfc =*_cabf ;_dfc .Set ("\u004ce\u006e\u0067\u0074\u0068\u0031",_g .MakeInteger (int64 (_afg .Len ())));};return objects ,nil ;};func _efbc (_bdb *_bc .Image ,_eeeg float64 )(*_bc .Image ,error ){_bbbe ,_ceg :=_bdb .ToGoImage ();
if _ceg !=nil {return nil ,_ceg ;};var _cdg _ae .Image ;_edc ,_dcgfb :=_bbbe .(*_ae .Monochrome );if _dcgfb {if _ceg =_edc .ResolveDecode ();_ceg !=nil {return nil ,_ceg ;};_cdg ,_ceg =_edc .Scale (_eeeg );if _ceg !=nil {return nil ,_ceg ;};}else {_efbb :=int (_e .RoundToEven (float64 (_bdb .Width )*_eeeg ));
_gbd :=int (_e .RoundToEven (float64 (_bdb .Height )*_eeeg ));_cdg ,_ceg =_ae .NewImage (_efbb ,_gbd ,int (_bdb .BitsPerComponent ),_bdb .ColorComponents ,nil ,nil ,nil );if _ceg !=nil {return nil ,_ceg ;};_a .CatmullRom .Scale (_cdg ,_cdg .Bounds (),_bbbe ,_bbbe .Bounds (),_a .Over ,&_a .Options {});
};_cbd :=_cdg .Base ();_bbeg :=&_bc .Image {Width :int64 (_cbd .Width ),Height :int64 (_cbd .Height ),BitsPerComponent :int64 (_cbd .BitsPerComponent ),ColorComponents :_cbd .ColorComponents ,Data :_cbd .Data };_bbeg .SetDecode (_cbd .Decode );_bbeg .SetAlpha (_cbd .Alpha );
return _bbeg ,nil ;};

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ImageQuality int ;};

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{};

// Options describes PDF optimization parameters.
type Options struct{CombineDuplicateStreams bool ;CombineDuplicateDirectObjects bool ;ImageUpperPPI float64 ;ImageQuality int ;UseObjectStreams bool ;CombineIdenticalIndirectObjects bool ;CompressStreams bool ;CleanFonts bool ;SubsetFonts bool ;CleanContentstream bool ;
};func _gda (_cab []_g .PdfObject )(_de map[*_g .PdfObjectStream ]struct{},_fggf error ){_de =map[*_g .PdfObjectStream ]struct{}{};_dge :=map[*_bc .PdfFont ]struct{}{};_fdb :=_dfd (_cab );for _ ,_abe :=range _fdb ._geg {_ddc ,_ace :=_g .GetDict (_abe .PdfObject );
if !_ace {continue ;};_bf ,_ace :=_g .GetDict (_ddc .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));if !_ace {continue ;};_ddg ,_ :=_eeec (_ddc .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));_adg ,_ga :=_bc .NewPdfPageResourcesFromDict (_bf );
if _ga !=nil {return nil ,_ga ;};_ffa :=[]content {{_cdd :_ddg ,_ecb :_adg }};_eafdg :=_ggf (_ddc .Get ("\u0041\u006e\u006e\u006f\u0074\u0073"));if _eafdg !=nil {_ffa =append (_ffa ,_eafdg ...);};for _ ,_be :=range _ffa {_bed ,_def :=_c .NewFromContents (_be ._cdd ,_be ._ecb );
if _def !=nil {return nil ,_def ;};_ebb ,_ ,_ ,_def :=_bed .ExtractPageText ();if _def !=nil {return nil ,_def ;};for _ ,_fdg :=range _ebb .Marks ().Elements (){if _fdg .Font ==nil {continue ;};if _ ,_fdd :=_dge [_fdg .Font ];!_fdd {_dge [_fdg .Font ]=struct{}{};
};};};};_ffg :=map[*_g .PdfObjectStream ][]*_bc .PdfFont {};for _cdc :=range _dge {_gef :=_cdc .FontDescriptor ();if _gef ==nil ||_gef .FontFile2 ==nil {continue ;};_fgf ,_eeg :=_g .GetStream (_gef .FontFile2 );if !_eeg {continue ;};_ffg [_fgf ]=append (_ffg [_fgf ],_cdc );
};for _beg :=range _ffg {var _add []rune ;var _eba []_ea .GlyphIndex ;for _ ,_bfd :=range _ffg [_beg ]{switch _bbe :=_bfd .Encoder ().(type ){case *_ed .IdentityEncoder :_cc :=_bbe .RegisteredRunes ();_cac :=make ([]_ea .GlyphIndex ,len (_cc ));for _fad ,_bac :=range _cc {_cac [_fad ]=_ea .GlyphIndex (_bac );
};_eba =append (_eba ,_cac ...);case *_ed .TrueTypeFontEncoder :_dgc :=_bbe .RegisteredRunes ();_add =append (_add ,_dgc ...);case _ed .SimpleEncoder :_fcc :=_bbe .Charcodes ();for _ ,_bg :=range _fcc {_fgd ,_fba :=_bbe .CharcodeToRune (_bg );if !_fba {_b .Log .Debug ("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064",_bg );
continue ;};_add =append (_add ,_fgd );};};};_fggf =_dbg (_beg ,_add ,_eba );if _fggf !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076",_fggf );
return nil ,_fggf ;};_de [_beg ]=struct{}{};};return _de ,nil ;};

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{};func _ccfg (_gbb *_bc .XObjectImage ,_beaeg imageModifications )error {_abd ,_cbef :=_gbb .ToImage ();if _cbef !=nil {return _cbef ;};if _beaeg .Scale !=0{_abd ,_cbef =_efbc (_abd ,_beaeg .Scale );if _cbef !=nil {return _cbef ;
};};if _beaeg .Encoding !=nil {_gbb .Filter =_beaeg .Encoding ;};_gbb .Decode =nil ;switch _acga :=_gbb .Filter .(type ){case *_g .FlateEncoder :if _acga .Predictor !=1&&_acga .Predictor !=11{_acga .Predictor =1;};};if _cbef =_gbb .SetImage (_abd ,nil );
_cbef !=nil {_b .Log .Debug ("\u0045\u0072\u0072or\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0076",_cbef );return _cbef ;};_gbb .ToPdfObject ();return nil ;};type imageInfo struct{BitsPerComponent int ;
ColorComponents int ;Width int ;Height int ;Stream *_g .PdfObjectStream ;PPI float64 ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_ebcd *ObjectStreams )Optimize (objects []_g .PdfObject )(_daf []_g .PdfObject ,_dfga error ){_bfcc :=&_g .PdfObjectStreams {};_gce :=make ([]_g .PdfObject ,0,len (objects ));for _ ,_cgf :=range objects {if _gbdc ,_adea :=_cgf .(*_g .PdfIndirectObject );
_adea &&_gbdc .GenerationNumber ==0{_bfcc .Append (_cgf );}else {_gce =append (_gce ,_cgf );};};if _bfcc .Len ()==0{return _gce ,nil ;};_daf =make ([]_g .PdfObject ,0,len (_gce )+_bfcc .Len ()+1);if _bfcc .Len ()> 1{_daf =append (_daf ,_bfcc );};_daf =append (_daf ,_bfcc .Elements ()...);
_daf =append (_daf ,_gce ...);return _daf ,nil ;};func _dbg (_egd *_g .PdfObjectStream ,_egdb []rune ,_bee []_ea .GlyphIndex )error {_egd ,_bce :=_g .GetStream (_egd );if !_bce {_b .Log .Debug ("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067");
return _f .New ("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_fca ,_fgdg :=_g .DecodeStream (_egd );if _fgdg !=nil {_b .Log .Debug ("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076",_fgdg );
return _fgdg ;};_aeb ,_fgdg :=_ea .Parse (_eaf .NewReader (_fca ));if _fgdg !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074",len (_egd .Stream ));
return _fgdg ;};_cae :=_bee ;if len (_egdb )> 0{_gcd :=_aeb .LookupRunes (_egdb );_cae =append (_cae ,_gcd ...);};_aeb ,_fgdg =_aeb .SubsetKeepIndices (_cae );if _fgdg !=nil {_b .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v",_fgdg );
return _fgdg ;};var _ag _eaf .Buffer ;_fgdg =_aeb .Write (&_ag );if _fgdg !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076",_fgdg );return _fgdg ;};if _ag .Len ()> len (_fca ){_b .Log .Debug ("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070");
return nil ;};_dgd ,_fgdg :=_g .MakeStream (_ag .Bytes (),_g .NewFlateEncoder ());if _fgdg !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076",_fgdg );return _fgdg ;
};*_egd =*_dgd ;_egd .Set ("\u004ce\u006e\u0067\u0074\u0068\u0031",_g .MakeInteger (int64 (_ag .Len ())));return nil ;};type objectStructure struct{_ccb *_g .PdfObjectDictionary ;_bfde *_g .PdfObjectDictionary ;_geg []*_g .PdfIndirectObject ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_fg *Chain )Optimize (objects []_g .PdfObject )(_ef []_g .PdfObject ,_cg error ){_ef =objects ;for _ ,_ca :=range _fg ._eda {_ef ,_cg =_ca .Optimize (_ef );if _cg !=nil {return _ef ,_cg ;};};return _ef ,nil ;};

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ImageUpperPPI float64 ;};func _cf (_gg *_dg .ContentStreamOperations )*_dg .ContentStreamOperations {if _gg ==nil {return nil ;};_db :=_dg .ContentStreamOperations {};for _ ,_fc :=range *_gg {switch _fc .Operand {case "\u0042\u0044\u0043","\u0042\u004d\u0043","\u0045\u004d\u0043":continue ;
case "\u0054\u006d":if len (_fc .Params )==6{if _ad ,_bcg :=_g .GetNumbersAsFloat (_fc .Params );_bcg ==nil {if _ad [0]==1&&_ad [1]==0&&_ad [2]==0&&_ad [3]==1{_fc =&_dg .ContentStreamOperation {Params :[]_g .PdfObject {_fc .Params [4],_fc .Params [5]},Operand :"\u0054\u0064"};
};};};};_db =append (_db ,_fc );};return &_db ;};

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{};func _dfd (_fgga []_g .PdfObject )objectStructure {_efaf :=objectStructure {};_caba :=false ;for _ ,_eag :=range _fgga {switch _dfb :=_eag .(type ){case *_g .PdfIndirectObject :_cffa ,_abc :=_g .GetDict (_dfb );if !_abc {continue ;
};_fgcg ,_abc :=_g .GetName (_cffa .Get ("\u0054\u0079\u0070\u0065"));if !_abc {continue ;};switch _fgcg .String (){case "\u0043a\u0074\u0061\u006c\u006f\u0067":_efaf ._ccb =_cffa ;_caba =true ;};};if _caba {break ;};};if !_caba {return _efaf ;};_ccc ,_ebff :=_g .GetDict (_efaf ._ccb .Get ("\u0050\u0061\u0067e\u0073"));
if !_ebff {return _efaf ;};_efaf ._bfde =_ccc ;_agg ,_ebff :=_g .GetArray (_ccc .Get ("\u004b\u0069\u0064\u0073"));if !_ebff {return _efaf ;};for _ ,_fdbg :=range _agg .Elements (){_bgga ,_gfad :=_g .GetIndirect (_fdbg );if !_gfad {break ;};_efaf ._geg =append (_efaf ._geg ,_bgga );
};return _efaf ;};

// New creates a optimizers chain from options.
func New (options Options )*Chain {_deeg :=new (Chain );if options .CleanFonts ||options .SubsetFonts {_deeg .Append (&CleanFonts {Subset :options .SubsetFonts });};if options .CleanContentstream {_deeg .Append (new (CleanContentstream ));};if options .ImageUpperPPI > 0{_bbge :=new (ImagePPI );
_bbge .ImageUpperPPI =options .ImageUpperPPI ;_deeg .Append (_bbge );};if options .ImageQuality > 0{_ddgc :=new (Image );_ddgc .ImageQuality =options .ImageQuality ;_deeg .Append (_ddgc );};if options .CombineDuplicateDirectObjects {_deeg .Append (new (CombineDuplicateDirectObjects ));
};if options .CombineDuplicateStreams {_deeg .Append (new (CombineDuplicateStreams ));};if options .CombineIdenticalIndirectObjects {_deeg .Append (new (CombineIdenticalIndirectObjects ));};if options .UseObjectStreams {_deeg .Append (new (ObjectStreams ));
};if options .CompressStreams {_deeg .Append (new (CompressStreams ));};return _deeg ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_cfa *CleanContentstream )Optimize (objects []_g .PdfObject )(_fe []_g .PdfObject ,_geb error ){_eae :=map[*_g .PdfObjectStream ]struct{}{};var _df []*_g .PdfObjectStream ;_gdc :=func (_cb *_g .PdfObjectStream ){if _ ,_gf :=_eae [_cb ];!_gf {_eae [_cb ]=struct{}{};
_df =append (_df ,_cb );};};_aa :=map[_g .PdfObject ]bool {};_fbd :=map[_g .PdfObject ]bool {};for _ ,_cba :=range objects {switch _fgg :=_cba .(type ){case *_g .PdfIndirectObject :switch _fd :=_fgg .PdfObject .(type ){case *_g .PdfObjectDictionary :if _fcf ,_fbb :=_g .GetName (_fd .Get ("\u0054\u0079\u0070\u0065"));
!_fbb ||_fcf .String ()!="\u0050\u0061\u0067\u0065"{continue ;};if _eg ,_dd :=_g .GetStream (_fd .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));_dd {_gdc (_eg );}else if _bb ,_ecf :=_g .GetArray (_fd .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
_ecf {var _bd []*_g .PdfObjectStream ;for _ ,_ecc :=range _bb .Elements (){if _adb ,_ab :=_g .GetStream (_ecc );_ab {_bd =append (_bd ,_adb );};};if len (_bd )> 0{var _fef _eaf .Buffer ;for _ ,_eabb :=range _bd {if _gb ,_ff :=_g .DecodeStream (_eabb );
_ff ==nil {_fef .Write (_gb );};_aa [_eabb ]=true ;};_bcd ,_acg :=_g .MakeStream (_fef .Bytes (),_g .NewFlateEncoder ());if _acg !=nil {return nil ,_acg ;};_fbd [_bcd ]=true ;_fd .Set ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073",_bcd );_gdc (_bcd );
};};};case *_g .PdfObjectStream :if _aae ,_dc :=_g .GetName (_fgg .Get ("\u0054\u0079\u0070\u0065"));!_dc ||_aae .String ()!="\u0058O\u0062\u006a\u0065\u0063\u0074"{continue ;};if _da ,_edg :=_g .GetName (_fgg .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
!_edg ||_da .String ()!="\u0046\u006f\u0072\u006d"{continue ;};_gdc (_fgg );};};for _ ,_eafd :=range _df {_geb =_cfe (_eafd );if _geb !=nil {return nil ,_geb ;};};_fe =nil ;for _ ,_eb :=range objects {if _aa [_eb ]{continue ;};_fe =append (_fe ,_eb );};
for _ba :=range _fbd {_fe =append (_fe ,_ba );};return _fe ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_abf *CompressStreams )Optimize (objects []_g .PdfObject )(_agbf []_g .PdfObject ,_adf error ){_agbf =make ([]_g .PdfObject ,len (objects ));copy (_agbf ,objects );for _ ,_efb :=range objects {_faf ,_aafa :=_g .GetStream (_efb );if !_aafa {continue ;
};if _cfg :=_faf .Get ("\u0046\u0069\u006c\u0074\u0065\u0072");_cfg !=nil {if _ ,_gcb :=_g .GetName (_cfg );_gcb {continue ;};if _fbf ,_fabd :=_g .GetArray (_cfg );_fabd &&_fbf .Len ()> 0{continue ;};};_abee :=_g .NewFlateEncoder ();var _fgc []byte ;_fgc ,_adf =_abee .EncodeBytes (_faf .Stream );
if _adf !=nil {return _agbf ,_adf ;};_eed :=_abee .MakeStreamDict ();if len (_fgc )+len (_eed .WriteString ())< len (_faf .Stream ){_faf .Stream =_fgc ;_faf .PdfObjectDictionary .Merge (_eed );_faf .PdfObjectDictionary .Set ("\u004c\u0065\u006e\u0067\u0074\u0068",_g .MakeInteger (int64 (len (_faf .Stream ))));
};};return _agbf ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_ce *CombineDuplicateDirectObjects )Optimize (objects []_g .PdfObject )(_aed []_g .PdfObject ,_dea error ){_edb (objects );_cfec :=make (map[string ][]*_g .PdfObjectDictionary );var _bea func (_gebf *_g .PdfObjectDictionary );_bea =func (_caea *_g .PdfObjectDictionary ){for _ ,_gae :=range _caea .Keys (){_cbc :=_caea .Get (_gae );
if _dcg ,_fefe :=_cbc .(*_g .PdfObjectDictionary );_fefe {_egg :=_ge .New ();_egg .Write ([]byte (_dcg .WriteString ()));_bdf :=string (_egg .Sum (nil ));_cfec [_bdf ]=append (_cfec [_bdf ],_dcg );_bea (_dcg );};};};for _ ,_efc :=range objects {_cdcg ,_deaa :=_efc .(*_g .PdfIndirectObject );
if !_deaa {continue ;};if _gfa ,_cdcgd :=_cdcg .PdfObject .(*_g .PdfObjectDictionary );_cdcgd {_bea (_gfa );};};_fge :=make ([]_g .PdfObject ,0,len (_cfec ));_baag :=make (map[_g .PdfObject ]_g .PdfObject );for _ ,_bbc :=range _cfec {if len (_bbc )< 2{continue ;
};_bcde :=_g .MakeDict ();_bcde .Merge (_bbc [0]);_eaea :=_g .MakeIndirectObject (_bcde );_fge =append (_fge ,_eaea );for _ffae :=0;_ffae < len (_bbc );_ffae ++{_gec :=_bbc [_ffae ];_baag [_gec ]=_eaea ;};};_aed =make ([]_g .PdfObject ,len (objects ));
copy (_aed ,objects );_aed =append (_fge ,_aed ...);_gebfd (_aed ,_baag );return _aed ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_cddc *CombineIdenticalIndirectObjects )Optimize (objects []_g .PdfObject )(_dbge []_g .PdfObject ,_gcc error ){_edb (objects );_bde :=make (map[_g .PdfObject ]_g .PdfObject );_afe :=make (map[_g .PdfObject ]struct{});_agb :=make (map[string ][]*_g .PdfIndirectObject );
for _ ,_beae :=range objects {_gdfc ,_daa :=_beae .(*_g .PdfIndirectObject );if !_daa {continue ;};if _bgg ,_cff :=_gdfc .PdfObject .(*_g .PdfObjectDictionary );_cff {if _eega ,_efed :=_bgg .Get ("\u0054\u0079\u0070\u0065").(*_g .PdfObjectName );_efed &&*_eega =="\u0050\u0061\u0067\u0065"{continue ;
};_aaf :=_ge .New ();_aaf .Write ([]byte (_bgg .WriteString ()));_acf :=string (_aaf .Sum (nil ));_agb [_acf ]=append (_agb [_acf ],_gdfc );};};for _ ,_cbe :=range _agb {if len (_cbe )< 2{continue ;};_beed :=_cbe [0];for _bbg :=1;_bbg < len (_cbe );_bbg ++{_gfbb :=_cbe [_bbg ];
_bde [_gfbb ]=_beed ;_afe [_gfbb ]=struct{}{};};};_dbge =make ([]_g .PdfObject ,0,len (objects )-len (_afe ));for _ ,_ecd :=range objects {if _ ,_gad :=_afe [_ecd ];_gad {continue ;};_dbge =append (_dbge ,_ecd );};_gebfd (_dbge ,_bde );return _dbge ,nil ;
};

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct{

// Subset embedded fonts if encountered (if true).
// Otherwise attempts to reduce the font program.
Subset bool ;};

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{};func _ggf (_aegg _g .PdfObject )[]content {if _aegg ==nil {return nil ;};_ddf ,_aee :=_g .GetArray (_aegg );if !_aee {_b .Log .Debug ("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");
return nil ;};var _ddgf []content ;for _ ,_ede :=range _ddf .Elements (){_gbe ,_edga :=_g .GetDict (_ede );if !_edga {_b .Log .Debug ("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073");
continue ;};_bcf ,_edga :=_g .GetDict (_gbe .Get ("\u0041\u0050"));if !_edga {_b .Log .Debug ("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067");continue ;};_cbg :=_g .TraceToDirectObject (_bcf .Get ("\u004e"));
if _cbg ==nil {_b .Log .Debug ("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067");continue ;};var _eef *_g .PdfObjectStream ;switch _agd :=_cbg .(type ){case *_g .PdfObjectDictionary :_dbb ,_bbd :=_g .GetName (_gbe .Get ("\u0041\u0053"));
if !_bbd {_b .Log .Debug ("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067");continue ;};_eef ,_bbd =_g .GetStream (_agd .Get (*_dbb ));if !_bbd {_b .Log .Debug ("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg");
continue ;};case *_g .PdfObjectStream :_eef =_agd ;};if _eef ==nil {_b .Log .Debug ("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067");
continue ;};_eafa ,_ded :=_bc .NewXObjectFormFromStream (_eef );if _ded !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_ded );
continue ;};_baa ,_ded :=_eafa .GetContentStream ();if _ded !=nil {_b .Log .Debug ("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076",_ded );continue ;};_ddgf =append (_ddgf ,content {_cdd :string (_baa ),_ecb :_eafa .Resources });
};return _ddgf ;};

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{};

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{};

// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type Chain struct{_eda []_bc .Optimizer };

// Append appends optimizers to the chain.
func (_fb *Chain )Append (optimizers ..._bc .Optimizer ){_fb ._eda =append (_fb ._eda ,optimizers ...)};

// Optimize optimizes PDF objects to decrease PDF size.
func (_deg *Image )Optimize (objects []_g .PdfObject )(_aaa []_g .PdfObject ,_dfg error ){if _deg .ImageQuality <=0{return objects ,nil ;};_ecba :=_ggg (objects );if len (_ecba )==0{return objects ,nil ;};_gfc :=make (map[_g .PdfObject ]_g .PdfObject );
_bbb :=make (map[_g .PdfObject ]struct{});for _ ,_gadf :=range _ecba {_ffdb :=_gadf .Stream .Get ("\u0053\u004d\u0061s\u006b");_bbb [_ffdb ]=struct{}{};};for _afeg ,_ecfg :=range _ecba {_eeea :=_ecfg .Stream ;if _ ,_gdd :=_bbb [_eeea ];_gdd {continue ;
};_gdae ,_aedd :=_bc .NewXObjectImageFromStream (_eeea );if _aedd !=nil {return nil ,_aedd ;};switch _gdae .Filter .(type ){case *_g .JBIG2Encoder :continue ;case *_g .CCITTFaxEncoder :continue ;};_agf ,_aedd :=_gdae .ToImage ();if _aedd !=nil {return nil ,_aedd ;
};_aaff :=_g .NewDCTEncoder ();_aaff .ColorComponents =_agf .ColorComponents ;_aaff .Quality =_deg .ImageQuality ;_aaff .BitsPerComponent =_ecfg .BitsPerComponent ;_aaff .Width =_ecfg .Width ;_aaff .Height =_ecfg .Height ;_cdcb ,_aedd :=_aaff .EncodeBytes (_agf .Data );
if _aedd !=nil {_b .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_aedd );return nil ,_aedd ;};var _adc _g .StreamEncoder ;_adc =_aaff ;{_gdab :=_g .NewFlateEncoder ();_eea :=_g .NewMultiEncoder ();_eea .AddEncoder (_gdab );_eea .AddEncoder (_aaff );
_gfec ,_dec :=_eea .EncodeBytes (_agf .Data );if _dec !=nil {return nil ,_dec ;};if len (_gfec )< len (_cdcb ){_b .Log .Trace ("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029",len (_cdcb ),len (_gfec ),len (_eeea .Stream ));
_cdcb =_gfec ;_adc =_eea ;};};_aebc :=len (_eeea .Stream );if _aebc < len (_cdcb ){continue ;};_fbg :=&_g .PdfObjectStream {Stream :_cdcb };_fbg .PdfObjectReference =_eeea .PdfObjectReference ;_fbg .PdfObjectDictionary =_g .MakeDict ();_fbg .Merge (_eeea .PdfObjectDictionary );
_fbg .Merge (_adc .MakeStreamDict ());_fbg .Set ("\u004c\u0065\u006e\u0067\u0074\u0068",_g .MakeInteger (int64 (len (_cdcb ))));_gfc [_eeea ]=_fbg ;_ecba [_afeg ].Stream =_fbg ;};_aaa =make ([]_g .PdfObject ,len (objects ));copy (_aaa ,objects );_gebfd (_aaa ,_gfc );
return _aaa ,nil ;};