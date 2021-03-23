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

package classer ;import (_c "github.com/unidoc/unipdf/v3/common";_be "github.com/unidoc/unipdf/v3/internal/jbig2/basic";_ab "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_d "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_a "image";_b "math";
);func _cb (_aa *_ab .Bitmap ,_aae ,_cgg ,_eba ,_gcf int ,_cdge *_ab .Bitmap )(_dd _a .Point ,_dcg error ){const _bgg ="\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg";if _aa ==nil {return _dd ,_d .Error (_bgg ,"\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");
};if _cdge ==nil {return _dd ,_d .Error (_bgg ,"t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064");};_dffg ,_bge :=_cdge .Width ,_cdge .Height ;_ddc ,_ddcd :=_aae -_eba -JbAddedPixels ,_cgg -_gcf -JbAddedPixels ;
_c .Log .Trace ("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027",_aae ,_cgg ,_dffg ,_bge ,_ddc ,_ddcd );
_gda ,_dcg :=_ab .Rect (_ddc ,_ddcd ,_dffg ,_bge );if _dcg !=nil {return _dd ,_d .Wrap (_dcg ,_bgg ,"");};_ef ,_ ,_dcg :=_aa .ClipRectangle (_gda );if _dcg !=nil {_c .Log .Error ("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076",_gda );
return _dd ,_d .Wrap (_dcg ,_bgg ,"");};_ddb :=_ab .New (_ef .Width ,_ef .Height );_af :=_b .MaxInt32 ;var _bde ,_fdf ,_ebd ,_dcc ,_bed int ;for _bde =-1;_bde <=1;_bde ++{for _fdf =-1;_fdf <=1;_fdf ++{if _ ,_dcg =_ab .Copy (_ddb ,_ef );_dcg !=nil {return _dd ,_d .Wrap (_dcg ,_bgg ,"");
};if _dcg =_ddb .RasterOperation (_fdf ,_bde ,_dffg ,_bge ,_ab .PixSrcXorDst ,_cdge ,0,0);_dcg !=nil {return _dd ,_d .Wrap (_dcg ,_bgg ,"");};_ebd =_ddb .CountPixels ();if _ebd < _af {_dcc =_fdf ;_bed =_bde ;_af =_ebd ;};};};_dd .X =_dcc ;_dd .Y =_bed ;
return _dd ,nil ;};var _fg bool ;func (_ceab *Classer )classifyRankHouseOne (_ffgd *_ab .Boxes ,_fgb ,_afg ,_ebad *_ab .Bitmaps ,_ddd *_ab .Points ,_ebag int )(_dae error ){const _fag ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_dbe ,_ga ,_bcb ,_fegc float32 ;_efdd int ;_fbb ,_adbd ,_fef ,_ddbg ,_fdd *_ab .Bitmap ;_gdd ,_bgc bool ;);for _bfg :=0;_bfg < len (_fgb .Values );_bfg ++{_adbd =_afg .Values [_bfg ];_fef =_ebad .Values [_bfg ];_dbe ,_ga ,_dae =_ddd .GetGeometry (_bfg );
if _dae !=nil {return _d .Wrapf (_dae ,_fag ,"\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079");};_ecdc :=len (_ceab .UndilatedTemplates .Values );_gdd =false ;_gfc :=_cgae (_ceab ,_adbd );for _efdd =_gfc .Next ();_efdd > -1;
{_ddbg ,_dae =_ceab .UndilatedTemplates .GetBitmap (_efdd );if _dae !=nil {return _d .Wrap (_dae ,_fag ,"\u0062\u006d\u0033");};_fdd ,_dae =_ceab .DilatedTemplates .GetBitmap (_efdd );if _dae !=nil {return _d .Wrap (_dae ,_fag ,"\u0062\u006d\u0034");};
_bcb ,_fegc ,_dae =_ceab .CentroidPointsTemplates .GetGeometry (_efdd );if _dae !=nil {return _d .Wrap (_dae ,_fag ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073");};_bgc ,_dae =_ab .HausTest (_adbd ,_fef ,_ddbg ,_fdd ,_dbe -_bcb ,_ga -_fegc ,MaxDiffWidth ,MaxDiffHeight );
if _dae !=nil {return _d .Wrap (_dae ,_fag ,"");};if _bgc {_gdd =true ;if _dae =_ceab .ClassIDs .Add (_efdd );_dae !=nil {return _d .Wrap (_dae ,_fag ,"");};if _dae =_ceab .ComponentPageNumbers .Add (_ebag );_dae !=nil {return _d .Wrap (_dae ,_fag ,"");
};if _ceab .Settings .KeepClassInstances {_gdaf ,_gab :=_ceab .ClassInstances .GetBitmaps (_efdd );if _gab !=nil {return _d .Wrap (_gab ,_fag ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_fbb ,_gab =_fgb .GetBitmap (_bfg );if _gab !=nil {return _d .Wrap (_gab ,_fag ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");
};_gdaf .AddBitmap (_fbb );_cbc ,_gab :=_ffgd .Get (_bfg );if _gab !=nil {return _d .Wrap (_gab ,_fag ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_gdaf .AddBox (_cbc );};break ;};};if !_gdd {if _dae =_ceab .ClassIDs .Add (_ecdc );_dae !=nil {return _d .Wrap (_dae ,_fag ,"");
};if _dae =_ceab .ComponentPageNumbers .Add (_ebag );_dae !=nil {return _d .Wrap (_dae ,_fag ,"");};_ffgc :=&_ab .Bitmaps {};_fbb ,_dae =_fgb .GetBitmap (_bfg );if _dae !=nil {return _d .Wrap (_dae ,_fag ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_ffgc .Values =append (_ffgc .Values ,_fbb );
_aec ,_ddg :=_fbb .Width ,_fbb .Height ;_ceab .TemplatesSize .Add (uint64 (_ddg )*uint64 (_aec ),_ecdc );_agd ,_deg :=_ffgd .Get (_bfg );if _deg !=nil {return _d .Wrap (_deg ,_fag ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_ffgc .AddBox (_agd );_ceab .ClassInstances .AddBitmaps (_ffgc );
_ceab .CentroidPointsTemplates .AddPoint (_dbe ,_ga );_ceab .UndilatedTemplates .AddBitmap (_adbd );_ceab .DilatedTemplates .AddBitmap (_fef );};};return nil ;};func (_ceb *Classer )addPageComponents (_gf *_ab .Bitmap ,_ade *_ab .Boxes ,_fc *_ab .Bitmaps ,_bd int ,_cdb Method )error {const _dfff ="\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s";
if _gf ==nil {return _d .Error (_dfff ,"\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065");};if _ade ==nil ||_fc ==nil ||len (*_ade )==0{_c .Log .Trace ("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064",_gf );
return nil ;};var _bea error ;switch _cdb {case RankHaus :_bea =_ceb .classifyRankHaus (_ade ,_fc ,_bd );case Correlation :_bea =_ceb .classifyCorrelation (_ade ,_fc ,_bd );default:_c .Log .Debug ("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027",_cdb );
return _d .Error (_dfff ,"\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064");};if _bea !=nil {return _d .Wrap (_bea ,_dfff ,"");};if _bea =_ceb .getULCorners (_gf ,_ade );_bea !=nil {return _d .Wrap (_bea ,_dfff ,"");
};_ge :=len (*_ade );_ceb .BaseIndex +=_ge ;if _bea =_ceb .ComponentsNumber .Add (_ge );_bea !=nil {return _d .Wrap (_bea ,_dfff ,"");};return nil ;};const (RankHaus Method =iota ;Correlation ;);const (MaxConnCompWidth =350;MaxCharCompWidth =350;MaxWordCompWidth =1000;
MaxCompHeight =120;);func (_g *Classer )AddPage (inputPage *_ab .Bitmap ,pageNumber int ,method Method )(_df error ){const _f ="\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065";_g .Widths [pageNumber ]=inputPage .Width ;
_g .Heights [pageNumber ]=inputPage .Height ;if _df =_g .verifyMethod (method );_df !=nil {return _d .Wrap (_df ,_f ,"");};_fd ,_bc ,_df :=inputPage .GetComponents (_g .Settings .Components ,_g .Settings .MaxCompWidth ,_g .Settings .MaxCompHeight );if _df !=nil {return _d .Wrap (_df ,_f ,"");
};_c .Log .Debug ("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076",_fd );if _df =_g .addPageComponents (inputPage ,_bc ,_fd ,pageNumber ,method );_df !=nil {return _d .Wrap (_df ,_f ,"");};return nil ;};type Classer struct{BaseIndex int ;
Settings Settings ;ComponentsNumber *_be .IntSlice ;TemplateAreas *_be .IntSlice ;Widths map[int ]int ;Heights map[int ]int ;NumberOfClasses int ;ClassInstances *_ab .BitmapsArray ;UndilatedTemplates *_ab .Bitmaps ;DilatedTemplates *_ab .Bitmaps ;TemplatesSize _be .IntsMap ;
FgTemplates *_be .NumSlice ;CentroidPoints *_ab .Points ;CentroidPointsTemplates *_ab .Points ;ClassIDs *_be .IntSlice ;ComponentPageNumbers *_be .IntSlice ;PtaUL *_ab .Points ;PtaLL *_ab .Points ;};const JbAddedPixels =6;type Method int ;func DefaultSettings ()Settings {_fda :=&Settings {};
_fda .SetDefault ();return *_fda };func (_ffd *Classer )classifyCorrelation (_cead *_ab .Boxes ,_baf *_ab .Bitmaps ,_bf int )error {const _da ="\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e";if _cead ==nil {return _d .Error (_da ,"\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064");
};if _baf ==nil {return _d .Error (_da ,"\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");};_ea :=len (_baf .Values );if _ea ==0{_c .Log .Debug ("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079");
return nil ;};var (_fe ,_cgf *_ab .Bitmap ;_aba error ;);_bda :=&_ab .Bitmaps {Values :make ([]*_ab .Bitmap ,_ea )};for _cf ,_fgc :=range _baf .Values {_cgf ,_aba =_fgc .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);if _aba !=nil {return _d .Wrap (_aba ,_da ,"");
};_bda .Values [_cf ]=_cgf ;};_gdc :=_ffd .FgTemplates ;_aac :=_ab .MakePixelSumTab8 ();_adef :=_ab .MakePixelCentroidTab8 ();_bba :=make ([]int ,_ea );_acg :=make ([][]int ,_ea );_cga :=_ab .Points (make ([]_ab .Point ,_ea ));_efc :=&_cga ;var (_cbd ,_gfa int ;
_acd ,_bfe ,_gdad int ;_fb ,_ebb int ;_bec byte ;);for _db ,_bef :=range _bda .Values {_acg [_db ]=make ([]int ,_bef .Height );_cbd =0;_gfa =0;_bfe =(_bef .Height -1)*_bef .RowStride ;_acd =0;for _ebb =_bef .Height -1;_ebb >=0;_ebb ,_bfe =_ebb -1,_bfe -_bef .RowStride {_acg [_db ][_ebb ]=_acd ;
_gdad =0;for _fb =0;_fb < _bef .RowStride ;_fb ++{_bec =_bef .Data [_bfe +_fb ];_gdad +=_aac [_bec ];_cbd +=_adef [_bec ]+_fb *8*_aac [_bec ];};_acd +=_gdad ;_gfa +=_gdad *_ebb ;};_bba [_db ]=_acd ;if _acd > 0{(*_efc )[_db ]=_ab .Point {X :float32 (_cbd )/float32 (_acd ),Y :float32 (_gfa )/float32 (_acd )};
}else {(*_efc )[_db ]=_ab .Point {X :float32 (_bef .Width )/float32 (2),Y :float32 (_bef .Height )/float32 (2)};};};if _aba =_ffd .CentroidPoints .Add (_efc );_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064");
};var (_acc ,_efb ,_dbf int ;_cce float64 ;_cbdg ,_efbd ,_feg ,_dda float32 ;_dfg ,_ebe _ab .Point ;_beg bool ;_fbc *similarTemplatesFinder ;_abd int ;_ced *_ab .Bitmap ;_efd *_a .Rectangle ;_gdfb *_ab .Bitmaps ;);for _abd ,_cgf =range _bda .Values {_efb =_bba [_abd ];
if _cbdg ,_efbd ,_aba =_efc .GetGeometry (_abd );_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0070t\u0061\u0020\u002d\u0020\u0069");};_beg =false ;_fdc :=len (_ffd .UndilatedTemplates .Values );_fbc =_cgae (_ffd ,_cgf );for _dcef :=_fbc .Next ();_dcef > -1;
{if _ced ,_aba =_ffd .UndilatedTemplates .GetBitmap (_dcef );_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032");};if _dbf ,_aba =_gdc .GetInt (_dcef );
_aba !=nil {_c .Log .Trace ("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_aba );};if _feg ,_dda ,_aba =_ffd .CentroidPointsTemplates .GetGeometry (_dcef );
_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020");
};if _ffd .Settings .WeightFactor > 0.0{if _acc ,_aba =_ffd .TemplateAreas .Get (_dcef );_aba !=nil {_c .Log .Trace ("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076",_aba );
};_cce =_ffd .Settings .Thresh +(1.0-_ffd .Settings .Thresh )*_ffd .Settings .WeightFactor *float64 (_dbf )/float64 (_acc );}else {_cce =_ffd .Settings .Thresh ;};_ebg ,_gbg :=_ab .CorrelationScoreThresholded (_cgf ,_ced ,_efb ,_dbf ,_dfg .X -_ebe .X ,_dfg .Y -_ebe .Y ,MaxDiffWidth ,MaxDiffHeight ,_aac ,_acg [_abd ],float32 (_cce ));
if _gbg !=nil {return _d .Wrap (_gbg ,_da ,"");};if _fg {var (_bca ,_bbaa float64 ;_afb ,_cgc int ;);_bca ,_gbg =_ab .CorrelationScore (_cgf ,_ced ,_efb ,_dbf ,_cbdg -_feg ,_efbd -_dda ,MaxDiffWidth ,MaxDiffHeight ,_aac );if _gbg !=nil {return _d .Wrap (_gbg ,_da ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_bbaa ,_gbg =_ab .CorrelationScoreSimple (_cgf ,_ced ,_efb ,_dbf ,_cbdg -_feg ,_efbd -_dda ,MaxDiffWidth ,MaxDiffHeight ,_aac );if _gbg !=nil {return _d .Wrap (_gbg ,_da ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_afb =int (_b .Sqrt (_bca *float64 (_efb )*float64 (_dbf )));_cgc =int (_b .Sqrt (_bbaa *float64 (_efb )*float64 (_dbf )));if (_bca >=_cce )!=(_bbaa >=_cce ){return _d .Errorf (_da ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066",_afb ,_bca ,_bca >=_cce ,_cgc ,_bbaa ,_bbaa >=_cce ,_bca -_bbaa );
};if _bca >=_cce !=_ebg {return _d .Errorf (_da ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076",_bca ,_bca *float64 (_efb )*float64 (_dbf ),_afb ,_cce ,float32 (_cce )*float32 (_efb )*float32 (_dbf ),_ebg );
};};if _ebg {_beg =true ;if _gbg =_ffd .ClassIDs .Add (_dcef );_gbg !=nil {return _d .Wrap (_gbg ,_da ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");};if _gbg =_ffd .ComponentPageNumbers .Add (_bf );_gbg !=nil {return _d .Wrap (_gbg ,_da ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");
};if _ffd .Settings .KeepClassInstances {if _fe ,_gbg =_baf .GetBitmap (_abd );_gbg !=nil {return _d .Wrap (_gbg ,_da ,"\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069");};if _gdfb ,_gbg =_ffd .ClassInstances .GetBitmaps (_dcef );
_gbg !=nil {return _d .Wrap (_gbg ,_da ,"K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073");};_gdfb .AddBitmap (_fe );if _efd ,_gbg =_cead .Get (_abd );_gbg !=nil {return _d .Wrap (_gbg ,_da ,"\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073");
};_gdfb .AddBox (_efd );};break ;};};if !_beg {if _aba =_ffd .ClassIDs .Add (_fdc );_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _aba =_ffd .ComponentPageNumbers .Add (_bf );_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_gdfb =&_ab .Bitmaps {};if _fe ,_aba =_baf .GetBitmap (_abd );_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_gdfb .AddBitmap (_fe );_fbf ,_eag :=_fe .Width ,_fe .Height ;_bfc :=uint64 (_eag )*uint64 (_fbf );_ffd .TemplatesSize .Add (_bfc ,_fdc );
if _efd ,_aba =_cead .Get (_abd );_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_gdfb .AddBox (_efd );_ffd .ClassInstances .AddBitmaps (_gdfb );_ffd .CentroidPointsTemplates .AddPoint (_cbdg ,_efbd );_ffd .FgTemplates .AddInt (_efb );
_ffd .UndilatedTemplates .AddBitmap (_fe );_acc =(_cgf .Width -2*JbAddedPixels )*(_cgf .Height -2*JbAddedPixels );if _aba =_ffd .TemplateAreas .Add (_acc );_aba !=nil {return _d .Wrap (_aba ,_da ,"\u0021\u0066\u006f\u0075\u006e\u0064");};};};_ffd .NumberOfClasses =len (_ffd .UndilatedTemplates .Values );
return nil ;};type Settings struct{MaxCompWidth int ;MaxCompHeight int ;SizeHaus int ;RankHaus float64 ;Thresh float64 ;WeightFactor float64 ;KeepClassInstances bool ;Components _ab .Component ;Method Method ;};func (_dffe *Settings )SetDefault (){if _dffe .MaxCompWidth ==0{switch _dffe .Components {case _ab .ComponentConn :_dffe .MaxCompWidth =MaxConnCompWidth ;
case _ab .ComponentCharacters :_dffe .MaxCompWidth =MaxCharCompWidth ;case _ab .ComponentWords :_dffe .MaxCompWidth =MaxWordCompWidth ;};};if _dffe .MaxCompHeight ==0{_dffe .MaxCompHeight =MaxCompHeight ;};if _dffe .Thresh ==0.0{_dffe .Thresh =0.9;};if _dffe .WeightFactor ==0.0{_dffe .WeightFactor =0.75;
};if _dffe .RankHaus ==0.0{_dffe .RankHaus =0.97;};if _dffe .SizeHaus ==0{_dffe .SizeHaus =2;};};func (_ff *Classer )verifyMethod (_ffg Method )error {if _ffg !=RankHaus &&_ffg !=Correlation {return _d .Error ("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064","\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064");
};return nil ;};func (_befe *Classer )classifyRankHaus (_ggf *_ab .Boxes ,_fae *_ab .Bitmaps ,_fga int )error {const _fcg ="\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073";if _ggf ==nil {return _d .Error (_fcg ,"\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");
};if _fae ==nil {return _d .Error (_fcg ,"\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");};_ee :=len (_fae .Values );if _ee ==0{return _d .Error (_fcg ,"e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073");
};_dca :=_fae .CountPixels ();_adae :=_befe .Settings .SizeHaus ;_edb :=_ab .SelCreateBrick (_adae ,_adae ,_adae /2,_adae /2,_ab .SelHit );_geb :=&_ab .Bitmaps {Values :make ([]*_ab .Bitmap ,_ee )};_gde :=&_ab .Bitmaps {Values :make ([]*_ab .Bitmap ,_ee )};
var (_agf ,_ca ,_gbb *_ab .Bitmap ;_de error ;);for _eg :=0;_eg < _ee ;_eg ++{_agf ,_de =_fae .GetBitmap (_eg );if _de !=nil {return _d .Wrap (_de ,_fcg ,"");};_ca ,_de =_agf .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);
if _de !=nil {return _d .Wrap (_de ,_fcg ,"");};_gbb ,_de =_ab .Dilate (nil ,_ca ,_edb );if _de !=nil {return _d .Wrap (_de ,_fcg ,"");};_geb .Values [_ee ]=_ca ;_gde .Values [_ee ]=_gbb ;};_gfe ,_de :=_ab .Centroids (_geb .Values );if _de !=nil {return _d .Wrap (_de ,_fcg ,"");
};if _de =_gfe .Add (_befe .CentroidPoints );_de !=nil {_c .Log .Trace ("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064");};if _befe .Settings .RankHaus ==1.0{_de =_befe .classifyRankHouseOne (_ggf ,_fae ,_geb ,_gde ,_gfe ,_fga );
}else {_de =_befe .classifyRankHouseNonOne (_ggf ,_fae ,_geb ,_gde ,_gfe ,_dca ,_fga );};if _de !=nil {return _d .Wrap (_de ,_fcg ,"");};return nil ;};func (_daeb *similarTemplatesFinder )Next ()int {var (_bae ,_ffe ,_bgee ,_cdgc int ;_cff bool ;_dbd *_ab .Bitmap ;
_eage error ;);for {if _daeb .Index >=25{return -1;};_ffe =_daeb .Width +TwoByTwoWalk [2*_daeb .Index ];_bae =_daeb .Height +TwoByTwoWalk [2*_daeb .Index +1];if _bae < 1||_ffe < 1{_daeb .Index ++;continue ;};if len (_daeb .CurrentNumbers )==0{_daeb .CurrentNumbers ,_cff =_daeb .Classer .TemplatesSize .GetSlice (uint64 (_ffe )*uint64 (_bae ));
if !_cff {_daeb .Index ++;continue ;};_daeb .N =0;};_bgee =len (_daeb .CurrentNumbers );for ;_daeb .N < _bgee ;_daeb .N ++{_cdgc =_daeb .CurrentNumbers [_daeb .N ];_dbd ,_eage =_daeb .Classer .DilatedTemplates .GetBitmap (_cdgc );if _eage !=nil {_c .Log .Debug ("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020");
return 0;};if _dbd .Width -2*JbAddedPixels ==_ffe &&_dbd .Height -2*JbAddedPixels ==_bae {return _cdgc ;};};_daeb .Index ++;_daeb .CurrentNumbers =nil ;};};func (_gfcb *Classer )classifyRankHouseNonOne (_fdff *_ab .Boxes ,_dec ,_bbe ,_bgb *_ab .Bitmaps ,_fcgd *_ab .Points ,_eec *_be .NumSlice ,_agfa int )(_fdg error ){const _edc ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_ded ,_dac ,_gebd ,_dad float32 ;_dbeg ,_aad ,_dfc int ;_ecf ,_dccg ,_aee ,_cbg ,_bcf *_ab .Bitmap ;_dfe ,_gabg bool ;);_dg :=_ab .MakePixelSumTab8 ();for _ebbe :=0;_ebbe < len (_dec .Values );_ebbe ++{if _dccg ,_fdg =_bbe .GetBitmap (_ebbe );_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029");
};if _dbeg ,_fdg =_eec .GetInt (_ebbe );_fdg !=nil {_c .Log .Trace ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_ebbe ,_fdg );
};if _aee ,_fdg =_bgb .GetBitmap (_ebbe );_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};if _ded ,_dac ,_fdg =_fcgd .GetGeometry (_ebbe );_fdg !=nil {return _d .Wrapf (_fdg ,_edc ,"\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079");
};_bgbf :=len (_gfcb .UndilatedTemplates .Values );_dfe =false ;_bbec :=_cgae (_gfcb ,_dccg );for _dfc =_bbec .Next ();_dfc > -1;{if _cbg ,_fdg =_gfcb .UndilatedTemplates .GetBitmap (_dfc );_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");
};if _aad ,_fdg =_gfcb .FgTemplates .GetInt (_dfc );_fdg !=nil {_c .Log .Trace ("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_dfc ,_fdg );
};if _bcf ,_fdg =_gfcb .DilatedTemplates .GetBitmap (_dfc );_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");};if _gebd ,_dad ,_fdg =_gfcb .CentroidPointsTemplates .GetGeometry (_dfc );
_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d");};_gabg ,_fdg =_ab .RankHausTest (_dccg ,_aee ,_cbg ,_bcf ,_ded -_gebd ,_dac -_dad ,MaxDiffWidth ,MaxDiffHeight ,_dbeg ,_aad ,float32 (_gfcb .Settings .RankHaus ),_dg );
if _fdg !=nil {return _d .Wrap (_fdg ,_edc ,"");};if _gabg {_dfe =true ;if _fdg =_gfcb .ClassIDs .Add (_dfc );_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"");};if _fdg =_gfcb .ComponentPageNumbers .Add (_agfa );_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"");
};if _gfcb .Settings .KeepClassInstances {_fee ,_ebc :=_gfcb .ClassInstances .GetBitmaps (_dfc );if _ebc !=nil {return _d .Wrap (_ebc ,_edc ,"\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029");
};if _ecf ,_ebc =_dec .GetBitmap (_ebbe );_ebc !=nil {return _d .Wrap (_ebc ,_edc ,"\u0070i\u0078\u0061\u005b\u0069\u005d");};_fee .Values =append (_fee .Values ,_ecf );_dea ,_ebc :=_fdff .Get (_ebbe );if _ebc !=nil {return _d .Wrap (_ebc ,_edc ,"b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029");
};_fee .Boxes =append (_fee .Boxes ,_dea );};break ;};};if !_dfe {if _fdg =_gfcb .ClassIDs .Add (_bgbf );_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _fdg =_gfcb .ComponentPageNumbers .Add (_agfa );_fdg !=nil {return _d .Wrap (_fdg ,_edc ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_cfc :=&_ab .Bitmaps {};_ecf =_dec .Values [_ebbe ];_cfc .AddBitmap (_ecf );_dgd ,_bgga :=_ecf .Width ,_ecf .Height ;_gfcb .TemplatesSize .Add (uint64 (_dgd )*uint64 (_bgga ),_bgbf );_aacc ,_ace :=_fdff .Get (_ebbe );if _ace !=nil {return _d .Wrap (_ace ,_edc ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_cfc .AddBox (_aacc );_gfcb .ClassInstances .AddBitmaps (_cfc );_gfcb .CentroidPointsTemplates .AddPoint (_ded ,_dac );_gfcb .UndilatedTemplates .AddBitmap (_dccg );_gfcb .DilatedTemplates .AddBitmap (_aee );_gfcb .FgTemplates .AddInt (_dbeg );};};_gfcb .NumberOfClasses =len (_gfcb .UndilatedTemplates .Values );
return nil ;};func (_cd *Classer )ComputeLLCorners ()(_gd error ){const _ad ="\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";if _cd .PtaUL ==nil {return _d .Error (_ad ,"\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064");
};_dc :=len (*_cd .PtaUL );_cd .PtaLL =&_ab .Points {};var (_dff ,_ce float32 ;_gc ,_cef int ;_bb *_ab .Bitmap ;);for _gg :=0;_gg < _dc ;_gg ++{_dff ,_ce ,_gd =_cd .PtaUL .GetGeometry (_gg );if _gd !=nil {_c .Log .Debug ("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_gd );
return _d .Wrap (_gd ,_ad ,"\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079");};_gc ,_gd =_cd .ClassIDs .Get (_gg );if _gd !=nil {_c .Log .Debug ("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076",_gd );
return _d .Wrap (_gd ,_ad ,"\u0043l\u0061\u0073\u0073\u0049\u0044");};_bb ,_gd =_cd .UndilatedTemplates .GetBitmap (_gc );if _gd !=nil {_c .Log .Debug ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_gd );
return _d .Wrap (_gd ,_ad ,"\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_cef =_bb .Height ;_cd .PtaLL .AddPoint (_dff ,_ce +float32 (_cef ));};return nil ;};var TwoByTwoWalk =[]int {0,0,0,1,-1,0,0,-1,1,0,-1,1,1,1,-1,-1,1,-1,0,-2,2,0,0,2,-2,0,-1,-2,1,-2,2,-1,2,1,1,2,-1,2,-2,1,-2,-1,-2,-2,2,-2,2,2,-2,2};
const (MaxDiffWidth =2;MaxDiffHeight =2;);func _cgae (_ddga *Classer ,_cbda *_ab .Bitmap )*similarTemplatesFinder {return &similarTemplatesFinder {Width :_cbda .Width ,Height :_cbda .Height ,Classer :_ddga };};type similarTemplatesFinder struct{Classer *Classer ;
Width int ;Height int ;Index int ;CurrentNumbers []int ;N int ;};func (_ba *Classer )getULCorners (_ceg *_ab .Bitmap ,_bab *_ab .Boxes )error {const _bag ="\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";if _ceg ==nil {return _d .Error (_bag ,"\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070");
};if _bab ==nil {return _d .Error (_bag ,"\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};if _ba .PtaUL ==nil {_ba .PtaUL =&_ab .Points {};};_gcb :=len (*_bab );var (_ae ,_gdf ,_ag ,_bg int ;_gb ,_dfd ,_ecd ,_cdc float32 ;_ggd error ;_adb *_a .Rectangle ;
_cdg *_ab .Bitmap ;_cea _a .Point ;);for _ada :=0;_ada < _gcb ;_ada ++{_ae =_ba .BaseIndex +_ada ;if _gb ,_dfd ,_ggd =_ba .CentroidPoints .GetGeometry (_ae );_ggd !=nil {return _d .Wrap (_ggd ,_bag ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073");
};if _gdf ,_ggd =_ba .ClassIDs .Get (_ae );_ggd !=nil {return _d .Wrap (_ggd ,_bag ,"\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074");};if _ecd ,_cdc ,_ggd =_ba .CentroidPointsTemplates .GetGeometry (_gdf );_ggd !=nil {return _d .Wrap (_ggd ,_bag ,"\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");
};_cg :=_ecd -_gb ;_ac :=_cdc -_dfd ;if _cg >=0{_ag =int (_cg +0.5);}else {_ag =int (_cg -0.5);};if _ac >=0{_bg =int (_ac +0.5);}else {_bg =int (_ac -0.5);};if _adb ,_ggd =_bab .Get (_ada );_ggd !=nil {return _d .Wrap (_ggd ,_bag ,"");};_cc ,_gef :=_adb .Min .X ,_adb .Min .Y ;
_cdg ,_ggd =_ba .UndilatedTemplates .GetBitmap (_gdf );if _ggd !=nil {return _d .Wrap (_ggd ,_bag ,"\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029");
};_cea ,_ggd =_cb (_ceg ,_cc ,_gef ,_ag ,_bg ,_cdg );if _ggd !=nil {return _d .Wrap (_ggd ,_bag ,"");};_ba .PtaUL .AddPoint (float32 (_cc -_ag +_cea .X ),float32 (_gef -_bg +_cea .Y ));};return nil ;};func (_cae Settings )Validate ()error {const _gbc ="\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065";
if _cae .Thresh < 0.4||_cae .Thresh > 0.98{return _d .Error (_gbc ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d");
};if _cae .WeightFactor < 0.0||_cae .WeightFactor > 1.0{return _d .Error (_gbc ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _cae .RankHaus < 0.5||_cae .RankHaus > 1.0{return _d .Error (_gbc ,"\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _cae .SizeHaus < 1||_cae .SizeHaus > 10{return _d .Error (_gbc ,"\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]");
};switch _cae .Components {case _ab .ComponentConn ,_ab .ComponentCharacters ,_ab .ComponentWords :default:return _d .Error (_gbc ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et");
};return nil ;};func Init (settings Settings )(*Classer ,error ){const _ed ="\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074";_ec :=&Classer {Settings :settings ,Widths :map[int ]int {},Heights :map[int ]int {},TemplatesSize :_be .IntsMap {},TemplateAreas :&_be .IntSlice {},ComponentPageNumbers :&_be .IntSlice {},ClassIDs :&_be .IntSlice {},ComponentsNumber :&_be .IntSlice {},CentroidPoints :&_ab .Points {},CentroidPointsTemplates :&_ab .Points {},UndilatedTemplates :&_ab .Bitmaps {},DilatedTemplates :&_ab .Bitmaps {},ClassInstances :&_ab .BitmapsArray {},FgTemplates :&_be .NumSlice {}};
if _eb :=_ec .Settings .Validate ();_eb !=nil {return nil ,_d .Wrap (_eb ,_ed ,"");};return _ec ,nil ;};