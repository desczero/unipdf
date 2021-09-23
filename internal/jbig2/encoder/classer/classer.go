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

package classer ;import (_e "github.com/unidoc/unipdf/v3/common";_gd "github.com/unidoc/unipdf/v3/internal/jbig2/basic";_cf "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_a "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_gb "image";_c "math";
);func (_bg *Classer )addPageComponents (_bf *_cf .Bitmap ,_cfb *_cf .Boxes ,_ge *_cf .Bitmaps ,_ec int ,_eg Method )error {const _cfde ="\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s";
if _bf ==nil {return _a .Error (_cfde ,"\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065");};if _cfb ==nil ||_ge ==nil ||len (*_cfb )==0{_e .Log .Trace ("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064",_bf );
return nil ;};var _gee error ;switch _eg {case RankHaus :_gee =_bg .classifyRankHaus (_cfb ,_ge ,_ec );case Correlation :_gee =_bg .classifyCorrelation (_cfb ,_ge ,_ec );default:_e .Log .Debug ("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027",_eg );
return _a .Error (_cfde ,"\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064");};if _gee !=nil {return _a .Wrap (_gee ,_cfde ,"");};if _gee =_bg .getULCorners (_bf ,_cfb );_gee !=nil {return _a .Wrap (_gee ,_cfde ,"");
};_gec :=len (*_cfb );_bg .BaseIndex +=_gec ;if _gee =_bg .ComponentsNumber .Add (_gec );_gee !=nil {return _a .Wrap (_gee ,_cfde ,"");};return nil ;};type Classer struct{BaseIndex int ;Settings Settings ;ComponentsNumber *_gd .IntSlice ;TemplateAreas *_gd .IntSlice ;
Widths map[int ]int ;Heights map[int ]int ;NumberOfClasses int ;ClassInstances *_cf .BitmapsArray ;UndilatedTemplates *_cf .Bitmaps ;DilatedTemplates *_cf .Bitmaps ;TemplatesSize _gd .IntsMap ;FgTemplates *_gd .NumSlice ;CentroidPoints *_cf .Points ;CentroidPointsTemplates *_cf .Points ;
ClassIDs *_gd .IntSlice ;ComponentPageNumbers *_gd .IntSlice ;PtaUL *_cf .Points ;PtaLL *_cf .Points ;};const (MaxDiffWidth =2;MaxDiffHeight =2;);func (_ag *Classer )AddPage (inputPage *_cf .Bitmap ,pageNumber int ,method Method )(_b error ){const _bc ="\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065";
_ag .Widths [pageNumber ]=inputPage .Width ;_ag .Heights [pageNumber ]=inputPage .Height ;if _b =_ag .verifyMethod (method );_b !=nil {return _a .Wrap (_b ,_bc ,"");};_ab ,_eb ,_b :=inputPage .GetComponents (_ag .Settings .Components ,_ag .Settings .MaxCompWidth ,_ag .Settings .MaxCompHeight );
if _b !=nil {return _a .Wrap (_b ,_bc ,"");};_e .Log .Debug ("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076",_ab );if _b =_ag .addPageComponents (inputPage ,_eb ,_ab ,pageNumber ,method );_b !=nil {return _a .Wrap (_b ,_bc ,"");
};return nil ;};var TwoByTwoWalk =[]int {0,0,0,1,-1,0,0,-1,1,0,-1,1,1,1,-1,-1,1,-1,0,-2,2,0,0,2,-2,0,-1,-2,1,-2,2,-1,2,1,1,2,-1,2,-2,1,-2,-1,-2,-2,2,-2,2,2,-2,2};type Settings struct{MaxCompWidth int ;MaxCompHeight int ;SizeHaus int ;RankHaus float64 ;
Thresh float64 ;WeightFactor float64 ;KeepClassInstances bool ;Components _cf .Component ;Method Method ;};const JbAddedPixels =6;func (_bde *Classer )verifyMethod (_bdf Method )error {if _bdf !=RankHaus &&_bdf !=Correlation {return _a .Error ("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064","\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064");
};return nil ;};func (_fed *Classer )classifyRankHaus (_df *_cf .Boxes ,_dba *_cf .Bitmaps ,_dbc int )error {const _agd ="\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073";if _df ==nil {return _a .Error (_agd ,"\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");
};if _dba ==nil {return _a .Error (_agd ,"\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");};_eff :=len (_dba .Values );if _eff ==0{return _a .Error (_agd ,"e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073");
};_bfc :=_dba .CountPixels ();_gfc :=_fed .Settings .SizeHaus ;_aeg :=_cf .SelCreateBrick (_gfc ,_gfc ,_gfc /2,_gfc /2,_cf .SelHit );_dce :=&_cf .Bitmaps {Values :make ([]*_cf .Bitmap ,_eff )};_ceaf :=&_cf .Bitmaps {Values :make ([]*_cf .Bitmap ,_eff )};
var (_ceac ,_decb ,_bcd *_cf .Bitmap ;_fdg error ;);for _acf :=0;_acf < _eff ;_acf ++{_ceac ,_fdg =_dba .GetBitmap (_acf );if _fdg !=nil {return _a .Wrap (_fdg ,_agd ,"");};_decb ,_fdg =_ceac .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);
if _fdg !=nil {return _a .Wrap (_fdg ,_agd ,"");};_bcd ,_fdg =_cf .Dilate (nil ,_decb ,_aeg );if _fdg !=nil {return _a .Wrap (_fdg ,_agd ,"");};_dce .Values [_eff ]=_decb ;_ceaf .Values [_eff ]=_bcd ;};_dfa ,_fdg :=_cf .Centroids (_dce .Values );if _fdg !=nil {return _a .Wrap (_fdg ,_agd ,"");
};if _fdg =_dfa .Add (_fed .CentroidPoints );_fdg !=nil {_e .Log .Trace ("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064");};if _fed .Settings .RankHaus ==1.0{_fdg =_fed .classifyRankHouseOne (_df ,_dba ,_dce ,_ceaf ,_dfa ,_dbc );
}else {_fdg =_fed .classifyRankHouseNonOne (_df ,_dba ,_dce ,_ceaf ,_dfa ,_bfc ,_dbc );};if _fdg !=nil {return _a .Wrap (_fdg ,_agd ,"");};return nil ;};func (_gcg *Classer )classifyCorrelation (_bcb *_cf .Boxes ,_dec *_cf .Bitmaps ,_fcf int )error {const _cbd ="\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e";
if _bcb ==nil {return _a .Error (_cbd ,"\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064");
};if _dec ==nil {return _a .Error (_cbd ,"\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");};_ade :=len (_dec .Values );if _ade ==0{_e .Log .Debug ("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079");
return nil ;};var (_ae ,_gaf *_cf .Bitmap ;_fdaa error ;);_da :=&_cf .Bitmaps {Values :make ([]*_cf .Bitmap ,_ade )};for _ebe ,_gga :=range _dec .Values {_gaf ,_fdaa =_gga .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);
if _fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"");};_da .Values [_ebe ]=_gaf ;};_bca :=_gcg .FgTemplates ;_fa :=_cf .MakePixelSumTab8 ();_adec :=_cf .MakePixelCentroidTab8 ();_cdb :=make ([]int ,_ade );_fbb :=make ([][]int ,_ade );_cge :=_cf .Points (make ([]_cf .Point ,_ade ));
_fcb :=&_cge ;var (_gcb ,_ea int ;_aad ,_aac ,_eed int ;_cea ,_ffe int ;_eae byte ;);for _fae ,_ffc :=range _da .Values {_fbb [_fae ]=make ([]int ,_ffc .Height );_gcb =0;_ea =0;_aac =(_ffc .Height -1)*_ffc .RowStride ;_aad =0;for _ffe =_ffc .Height -1;
_ffe >=0;_ffe ,_aac =_ffe -1,_aac -_ffc .RowStride {_fbb [_fae ][_ffe ]=_aad ;_eed =0;for _cea =0;_cea < _ffc .RowStride ;_cea ++{_eae =_ffc .Data [_aac +_cea ];_eed +=_fa [_eae ];_gcb +=_adec [_eae ]+_cea *8*_fa [_eae ];};_aad +=_eed ;_ea +=_eed *_ffe ;
};_cdb [_fae ]=_aad ;if _aad > 0{(*_fcb )[_fae ]=_cf .Point {X :float32 (_gcb )/float32 (_aad ),Y :float32 (_ea )/float32 (_aad )};}else {(*_fcb )[_fae ]=_cf .Point {X :float32 (_ffc .Width )/float32 (2),Y :float32 (_ffc .Height )/float32 (2)};};};if _fdaa =_gcg .CentroidPoints .Add (_fcb );
_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064");};var (_fe ,_fdad ,_bce int ;_dd float64 ;_ecad ,_fee ,_efc ,_eeg float32 ;_eea ,_aacg _cf .Point ;_afb bool ;_fbf *similarTemplatesFinder ;
_adg int ;_daf *_cf .Bitmap ;_cfbc *_gb .Rectangle ;_ace *_cf .Bitmaps ;);for _adg ,_gaf =range _da .Values {_fdad =_cdb [_adg ];if _ecad ,_fee ,_fdaa =_fcb .GetGeometry (_adg );_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0070t\u0061\u0020\u002d\u0020\u0069");
};_afb =false ;_feb :=len (_gcg .UndilatedTemplates .Values );_fbf =_bfb (_gcg ,_gaf );for _gef :=_fbf .Next ();_gef > -1;{if _daf ,_fdaa =_gcg .UndilatedTemplates .GetBitmap (_gef );_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032");
};if _bce ,_fdaa =_bca .GetInt (_gef );_fdaa !=nil {_e .Log .Trace ("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_fdaa );};
if _efc ,_eeg ,_fdaa =_gcg .CentroidPointsTemplates .GetGeometry (_gef );_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020");
};if _gcg .Settings .WeightFactor > 0.0{if _fe ,_fdaa =_gcg .TemplateAreas .Get (_gef );_fdaa !=nil {_e .Log .Trace ("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076",_fdaa );
};_dd =_gcg .Settings .Thresh +(1.0-_gcg .Settings .Thresh )*_gcg .Settings .WeightFactor *float64 (_bce )/float64 (_fe );}else {_dd =_gcg .Settings .Thresh ;};_cdbe ,_eaf :=_cf .CorrelationScoreThresholded (_gaf ,_daf ,_fdad ,_bce ,_eea .X -_aacg .X ,_eea .Y -_aacg .Y ,MaxDiffWidth ,MaxDiffHeight ,_fa ,_fbb [_adg ],float32 (_dd ));
if _eaf !=nil {return _a .Wrap (_eaf ,_cbd ,"");};if _cbf {var (_bb ,_dga float64 ;_cdd ,_cbdg int ;);_bb ,_eaf =_cf .CorrelationScore (_gaf ,_daf ,_fdad ,_bce ,_ecad -_efc ,_fee -_eeg ,MaxDiffWidth ,MaxDiffHeight ,_fa );if _eaf !=nil {return _a .Wrap (_eaf ,_cbd ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_dga ,_eaf =_cf .CorrelationScoreSimple (_gaf ,_daf ,_fdad ,_bce ,_ecad -_efc ,_fee -_eeg ,MaxDiffWidth ,MaxDiffHeight ,_fa );if _eaf !=nil {return _a .Wrap (_eaf ,_cbd ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_cdd =int (_c .Sqrt (_bb *float64 (_fdad )*float64 (_bce )));_cbdg =int (_c .Sqrt (_dga *float64 (_fdad )*float64 (_bce )));if (_bb >=_dd )!=(_dga >=_dd ){return _a .Errorf (_cbd ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066",_cdd ,_bb ,_bb >=_dd ,_cbdg ,_dga ,_dga >=_dd ,_bb -_dga );
};if _bb >=_dd !=_cdbe {return _a .Errorf (_cbd ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076",_bb ,_bb *float64 (_fdad )*float64 (_bce ),_cdd ,_dd ,float32 (_dd )*float32 (_fdad )*float32 (_bce ),_cdbe );
};};if _cdbe {_afb =true ;if _eaf =_gcg .ClassIDs .Add (_gef );_eaf !=nil {return _a .Wrap (_eaf ,_cbd ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");};if _eaf =_gcg .ComponentPageNumbers .Add (_fcf );_eaf !=nil {return _a .Wrap (_eaf ,_cbd ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");
};if _gcg .Settings .KeepClassInstances {if _ae ,_eaf =_dec .GetBitmap (_adg );_eaf !=nil {return _a .Wrap (_eaf ,_cbd ,"\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069");};if _ace ,_eaf =_gcg .ClassInstances .GetBitmaps (_gef );
_eaf !=nil {return _a .Wrap (_eaf ,_cbd ,"K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073");};_ace .AddBitmap (_ae );if _cfbc ,_eaf =_bcb .Get (_adg );_eaf !=nil {return _a .Wrap (_eaf ,_cbd ,"\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073");
};_ace .AddBox (_cfbc );};break ;};};if !_afb {if _fdaa =_gcg .ClassIDs .Add (_feb );_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _fdaa =_gcg .ComponentPageNumbers .Add (_fcf );_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_ace =&_cf .Bitmaps {};if _ae ,_fdaa =_dec .GetBitmap (_adg );_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_ace .AddBitmap (_ae );_gbeb ,_fgg :=_ae .Width ,_ae .Height ;_bdc :=uint64 (_fgg )*uint64 (_gbeb );_gcg .TemplatesSize .Add (_bdc ,_feb );
if _cfbc ,_fdaa =_bcb .Get (_adg );_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_ace .AddBox (_cfbc );_gcg .ClassInstances .AddBitmaps (_ace );_gcg .CentroidPointsTemplates .AddPoint (_ecad ,_fee );_gcg .FgTemplates .AddInt (_fdad );
_gcg .UndilatedTemplates .AddBitmap (_ae );_fe =(_gaf .Width -2*JbAddedPixels )*(_gaf .Height -2*JbAddedPixels );if _fdaa =_gcg .TemplateAreas .Add (_fe );_fdaa !=nil {return _a .Wrap (_fdaa ,_cbd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};};};_gcg .NumberOfClasses =len (_gcg .UndilatedTemplates .Values );
return nil ;};func _bfb (_fce *Classer ,_gedf *_cf .Bitmap )*similarTemplatesFinder {return &similarTemplatesFinder {Width :_gedf .Width ,Height :_gedf .Height ,Classer :_fce };};func Init (settings Settings )(*Classer ,error ){const _ca ="\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074";
_gf :=&Classer {Settings :settings ,Widths :map[int ]int {},Heights :map[int ]int {},TemplatesSize :_gd .IntsMap {},TemplateAreas :&_gd .IntSlice {},ComponentPageNumbers :&_gd .IntSlice {},ClassIDs :&_gd .IntSlice {},ComponentsNumber :&_gd .IntSlice {},CentroidPoints :&_cf .Points {},CentroidPointsTemplates :&_cf .Points {},UndilatedTemplates :&_cf .Bitmaps {},DilatedTemplates :&_cf .Bitmaps {},ClassInstances :&_cf .BitmapsArray {},FgTemplates :&_gd .NumSlice {}};
if _cb :=_gf .Settings .Validate ();_cb !=nil {return nil ,_a .Wrap (_cb ,_ca ,"");};return _gf ,nil ;};var _cbf bool ;const (MaxConnCompWidth =350;MaxCharCompWidth =350;MaxWordCompWidth =1000;MaxCompHeight =120;);type similarTemplatesFinder struct{Classer *Classer ;
Width int ;Height int ;Index int ;CurrentNumbers []int ;N int ;};const (RankHaus Method =iota ;Correlation ;);func (_cbg *similarTemplatesFinder )Next ()int {var (_ddc ,_cgfb ,_aca ,_cafb int ;_bcab bool ;_cae *_cf .Bitmap ;_age error ;);for {if _cbg .Index >=25{return -1;
};_cgfb =_cbg .Width +TwoByTwoWalk [2*_cbg .Index ];_ddc =_cbg .Height +TwoByTwoWalk [2*_cbg .Index +1];if _ddc < 1||_cgfb < 1{_cbg .Index ++;continue ;};if len (_cbg .CurrentNumbers )==0{_cbg .CurrentNumbers ,_bcab =_cbg .Classer .TemplatesSize .GetSlice (uint64 (_cgfb )*uint64 (_ddc ));
if !_bcab {_cbg .Index ++;continue ;};_cbg .N =0;};_aca =len (_cbg .CurrentNumbers );for ;_cbg .N < _aca ;_cbg .N ++{_cafb =_cbg .CurrentNumbers [_cbg .N ];_cae ,_age =_cbg .Classer .DilatedTemplates .GetBitmap (_cafb );if _age !=nil {_e .Log .Debug ("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020");
return 0;};if _cae .Width -2*JbAddedPixels ==_cgfb &&_cae .Height -2*JbAddedPixels ==_ddc {return _cafb ;};};_cbg .Index ++;_cbg .CurrentNumbers =nil ;};};func (_bag *Classer )classifyRankHouseNonOne (_daff *_cf .Boxes ,_cee ,_acc ,_ggad *_cf .Bitmaps ,_cdda *_cf .Points ,_fadd *_gd .NumSlice ,_baf int )(_gge error ){const _egf ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_gbd ,_fec ,_dfe ,_gafg float32 ;_fecg ,_dbf ,_dae int ;_febg ,_fca ,_cgf ,_gac ,_edg *_cf .Bitmap ;_gbf ,_agfa bool ;);_cgff :=_cf .MakePixelSumTab8 ();for _aeb :=0;_aeb < len (_cee .Values );_aeb ++{if _fca ,_gge =_acc .GetBitmap (_aeb );_gge !=nil {return _a .Wrap (_gge ,_egf ,"b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029");
};if _fecg ,_gge =_fadd .GetInt (_aeb );_gge !=nil {_e .Log .Trace ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_aeb ,_gge );
};if _cgf ,_gge =_ggad .GetBitmap (_aeb );_gge !=nil {return _a .Wrap (_gge ,_egf ,"b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};if _gbd ,_fec ,_gge =_cdda .GetGeometry (_aeb );_gge !=nil {return _a .Wrapf (_gge ,_egf ,"\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079");
};_bec :=len (_bag .UndilatedTemplates .Values );_gbf =false ;_fdb :=_bfb (_bag ,_fca );for _dae =_fdb .Next ();_dae > -1;{if _gac ,_gge =_bag .UndilatedTemplates .GetBitmap (_dae );_gge !=nil {return _a .Wrap (_gge ,_egf ,"\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");
};if _dbf ,_gge =_bag .FgTemplates .GetInt (_dae );_gge !=nil {_e .Log .Trace ("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_dae ,_gge );
};if _edg ,_gge =_bag .DilatedTemplates .GetBitmap (_dae );_gge !=nil {return _a .Wrap (_gge ,_egf ,"\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");};if _dfe ,_gafg ,_gge =_bag .CentroidPointsTemplates .GetGeometry (_dae );
_gge !=nil {return _a .Wrap (_gge ,_egf ,"\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d");};_agfa ,_gge =_cf .RankHausTest (_fca ,_cgf ,_gac ,_edg ,_gbd -_dfe ,_fec -_gafg ,MaxDiffWidth ,MaxDiffHeight ,_fecg ,_dbf ,float32 (_bag .Settings .RankHaus ),_cgff );
if _gge !=nil {return _a .Wrap (_gge ,_egf ,"");};if _agfa {_gbf =true ;if _gge =_bag .ClassIDs .Add (_dae );_gge !=nil {return _a .Wrap (_gge ,_egf ,"");};if _gge =_bag .ComponentPageNumbers .Add (_baf );_gge !=nil {return _a .Wrap (_gge ,_egf ,"");};
if _bag .Settings .KeepClassInstances {_dac ,_gba :=_bag .ClassInstances .GetBitmaps (_dae );if _gba !=nil {return _a .Wrap (_gba ,_egf ,"\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029");
};if _febg ,_gba =_cee .GetBitmap (_aeb );_gba !=nil {return _a .Wrap (_gba ,_egf ,"\u0070i\u0078\u0061\u005b\u0069\u005d");};_dac .Values =append (_dac .Values ,_febg );_bcc ,_gba :=_daff .Get (_aeb );if _gba !=nil {return _a .Wrap (_gba ,_egf ,"b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029");
};_dac .Boxes =append (_dac .Boxes ,_bcc );};break ;};};if !_gbf {if _gge =_bag .ClassIDs .Add (_bec );_gge !=nil {return _a .Wrap (_gge ,_egf ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _gge =_bag .ComponentPageNumbers .Add (_baf );_gge !=nil {return _a .Wrap (_gge ,_egf ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_edd :=&_cf .Bitmaps {};_febg =_cee .Values [_aeb ];_edd .AddBitmap (_febg );_gfb ,_afd :=_febg .Width ,_febg .Height ;_bag .TemplatesSize .Add (uint64 (_gfb )*uint64 (_afd ),_bec );_gdg ,_gdd :=_daff .Get (_aeb );if _gdd !=nil {return _a .Wrap (_gdd ,_egf ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_edd .AddBox (_gdg );_bag .ClassInstances .AddBitmaps (_edd );_bag .CentroidPointsTemplates .AddPoint (_gbd ,_fec );_bag .UndilatedTemplates .AddBitmap (_fca );_bag .DilatedTemplates .AddBitmap (_cgf );_bag .FgTemplates .AddInt (_fecg );};};_bag .NumberOfClasses =len (_bag .UndilatedTemplates .Values );
return nil ;};func DefaultSettings ()Settings {_ebg :=&Settings {};_ebg .SetDefault ();return *_ebg };func _fc (_dg *_cf .Bitmap ,_cd ,_caa ,_dca ,_cfdc int ,_cce *_cf .Bitmap )(_fb _gb .Point ,_de error ){const _cab ="\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg";
if _dg ==nil {return _fb ,_a .Error (_cab ,"\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};if _cce ==nil {return _fb ,_a .Error (_cab ,"t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064");
};_ef ,_efe :=_cce .Width ,_cce .Height ;_ee ,_aab :=_cd -_dca -JbAddedPixels ,_caa -_cfdc -JbAddedPixels ;_e .Log .Trace ("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027",_cd ,_caa ,_ef ,_efe ,_ee ,_aab );
_dbg ,_de :=_cf .Rect (_ee ,_aab ,_ef ,_efe );if _de !=nil {return _fb ,_a .Wrap (_de ,_cab ,"");};_ce ,_ ,_de :=_dg .ClipRectangle (_dbg );if _de !=nil {_e .Log .Error ("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076",_dbg );
return _fb ,_a .Wrap (_de ,_cab ,"");};_gbe :=_cf .New (_ce .Width ,_ce .Height );_bdb :=_c .MaxInt32 ;var _fcc ,_bdec ,_cgc ,_eca ,_fge int ;for _fcc =-1;_fcc <=1;_fcc ++{for _bdec =-1;_bdec <=1;_bdec ++{if _ ,_de =_cf .Copy (_gbe ,_ce );_de !=nil {return _fb ,_a .Wrap (_de ,_cab ,"");
};if _de =_gbe .RasterOperation (_bdec ,_fcc ,_ef ,_efe ,_cf .PixSrcXorDst ,_cce ,0,0);_de !=nil {return _fb ,_a .Wrap (_de ,_cab ,"");};_cgc =_gbe .CountPixels ();if _cgc < _bdb {_eca =_bdec ;_fge =_fcc ;_bdb =_cgc ;};};};_fb .X =_eca ;_fb .Y =_fge ;return _fb ,nil ;
};func (_efd *Classer )classifyRankHouseOne (_fea *_cf .Boxes ,_caag ,_bfe ,_cdbf *_cf .Bitmaps ,_dcd *_cf .Points ,_dcc int )(_ceae error ){const _deb ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_abc ,_dbaa ,_aag ,_bgbe float32 ;_cfdb int ;_egb ,_eeee ,_fad ,_fede ,_bbd *_cf .Bitmap ;_cag ,_cgb bool ;);for _gece :=0;_gece < len (_caag .Values );_gece ++{_eeee =_bfe .Values [_gece ];_fad =_cdbf .Values [_gece ];_abc ,_dbaa ,_ceae =_dcd .GetGeometry (_gece );
if _ceae !=nil {return _a .Wrapf (_ceae ,_deb ,"\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079");};_afbc :=len (_efd .UndilatedTemplates .Values );_cag =false ;_aacd :=_bfb (_efd ,_eeee );for _cfdb =_aacd .Next ();_cfdb > -1;
{_fede ,_ceae =_efd .UndilatedTemplates .GetBitmap (_cfdb );if _ceae !=nil {return _a .Wrap (_ceae ,_deb ,"\u0062\u006d\u0033");};_bbd ,_ceae =_efd .DilatedTemplates .GetBitmap (_cfdb );if _ceae !=nil {return _a .Wrap (_ceae ,_deb ,"\u0062\u006d\u0034");
};_aag ,_bgbe ,_ceae =_efd .CentroidPointsTemplates .GetGeometry (_cfdb );if _ceae !=nil {return _a .Wrap (_ceae ,_deb ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073");};_cgb ,_ceae =_cf .HausTest (_eeee ,_fad ,_fede ,_bbd ,_abc -_aag ,_dbaa -_bgbe ,MaxDiffWidth ,MaxDiffHeight );
if _ceae !=nil {return _a .Wrap (_ceae ,_deb ,"");};if _cgb {_cag =true ;if _ceae =_efd .ClassIDs .Add (_cfdb );_ceae !=nil {return _a .Wrap (_ceae ,_deb ,"");};if _ceae =_efd .ComponentPageNumbers .Add (_dcc );_ceae !=nil {return _a .Wrap (_ceae ,_deb ,"");
};if _efd .Settings .KeepClassInstances {_ged ,_be :=_efd .ClassInstances .GetBitmaps (_cfdb );if _be !=nil {return _a .Wrap (_be ,_deb ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_egb ,_be =_caag .GetBitmap (_gece );if _be !=nil {return _a .Wrap (_be ,_deb ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");
};_ged .AddBitmap (_egb );_eda ,_be :=_fea .Get (_gece );if _be !=nil {return _a .Wrap (_be ,_deb ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_ged .AddBox (_eda );};break ;};};if !_cag {if _ceae =_efd .ClassIDs .Add (_afbc );_ceae !=nil {return _a .Wrap (_ceae ,_deb ,"");
};if _ceae =_efd .ComponentPageNumbers .Add (_dcc );_ceae !=nil {return _a .Wrap (_ceae ,_deb ,"");};_ggb :=&_cf .Bitmaps {};_egb ,_ceae =_caag .GetBitmap (_gece );if _ceae !=nil {return _a .Wrap (_ceae ,_deb ,"\u0021\u0066\u006f\u0075\u006e\u0064");};
_ggb .Values =append (_ggb .Values ,_egb );_fbfc ,_fbe :=_egb .Width ,_egb .Height ;_efd .TemplatesSize .Add (uint64 (_fbe )*uint64 (_fbfc ),_afbc );_cead ,_acea :=_fea .Get (_gece );if _acea !=nil {return _a .Wrap (_acea ,_deb ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_ggb .AddBox (_cead );_efd .ClassInstances .AddBitmaps (_ggb );_efd .CentroidPointsTemplates .AddPoint (_abc ,_dbaa );_efd .UndilatedTemplates .AddBitmap (_eeee );_efd .DilatedTemplates .AddBitmap (_fad );};};return nil ;};type Method int ;func (_gg *Classer )ComputeLLCorners ()(_cfc error ){const _bd ="\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";
if _gg .PtaUL ==nil {return _a .Error (_bd ,"\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064");};_cg :=len (*_gg .PtaUL );_gg .PtaLL =&_cf .Points {};var (_cfd ,_d float32 ;_cbc ,_f int ;
_ggf *_cf .Bitmap ;);for _dc :=0;_dc < _cg ;_dc ++{_cfd ,_d ,_cfc =_gg .PtaUL .GetGeometry (_dc );if _cfc !=nil {_e .Log .Debug ("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_cfc );
return _a .Wrap (_cfc ,_bd ,"\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079");};_cbc ,_cfc =_gg .ClassIDs .Get (_dc );if _cfc !=nil {_e .Log .Debug ("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076",_cfc );
return _a .Wrap (_cfc ,_bd ,"\u0043l\u0061\u0073\u0073\u0049\u0044");};_ggf ,_cfc =_gg .UndilatedTemplates .GetBitmap (_cbc );if _cfc !=nil {_e .Log .Debug ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_cfc );
return _a .Wrap (_cfc ,_bd ,"\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_f =_ggf .Height ;_gg .PtaLL .AddPoint (_cfd ,_d +float32 (_f ));};return nil ;};func (_ded Settings )Validate ()error {const _fef ="\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065";
if _ded .Thresh < 0.4||_ded .Thresh > 0.98{return _a .Error (_fef ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d");
};if _ded .WeightFactor < 0.0||_ded .WeightFactor > 1.0{return _a .Error (_fef ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _ded .RankHaus < 0.5||_ded .RankHaus > 1.0{return _a .Error (_fef ,"\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _ded .SizeHaus < 1||_ded .SizeHaus > 10{return _a .Error (_fef ,"\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]");
};switch _ded .Components {case _cf .ComponentConn ,_cf .ComponentCharacters ,_cf .ComponentWords :default:return _a .Error (_fef ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et");
};return nil ;};func (_bbb *Settings )SetDefault (){if _bbb .MaxCompWidth ==0{switch _bbb .Components {case _cf .ComponentConn :_bbb .MaxCompWidth =MaxConnCompWidth ;case _cf .ComponentCharacters :_bbb .MaxCompWidth =MaxCharCompWidth ;case _cf .ComponentWords :_bbb .MaxCompWidth =MaxWordCompWidth ;
};};if _bbb .MaxCompHeight ==0{_bbb .MaxCompHeight =MaxCompHeight ;};if _bbb .Thresh ==0.0{_bbb .Thresh =0.9;};if _bbb .WeightFactor ==0.0{_bbb .WeightFactor =0.75;};if _bbb .RankHaus ==0.0{_bbb .RankHaus =0.97;};if _bbb .SizeHaus ==0{_bbb .SizeHaus =2;
};};func (_ggfg *Classer )getULCorners (_caf *_cf .Bitmap ,_aa *_cf .Boxes )error {const _ba ="\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";if _caf ==nil {return _a .Error (_ba ,"\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070");
};if _aa ==nil {return _a .Error (_ba ,"\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};if _ggfg .PtaUL ==nil {_ggfg .PtaUL =&_cf .Points {};};_ed :=len (*_aa );var (_gda ,_ad ,_ff ,_cc int ;_ga ,_gc ,_agf ,_bge float32 ;_fg error ;_af *_gb .Rectangle ;
_ffa *_cf .Bitmap ;_ac _gb .Point ;);for _db :=0;_db < _ed ;_db ++{_gda =_ggfg .BaseIndex +_db ;if _ga ,_gc ,_fg =_ggfg .CentroidPoints .GetGeometry (_gda );_fg !=nil {return _a .Wrap (_fg ,_ba ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073");
};if _ad ,_fg =_ggfg .ClassIDs .Get (_gda );_fg !=nil {return _a .Wrap (_fg ,_ba ,"\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074");};if _agf ,_bge ,_fg =_ggfg .CentroidPointsTemplates .GetGeometry (_ad );_fg !=nil {return _a .Wrap (_fg ,_ba ,"\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");
};_fd :=_agf -_ga ;_fda :=_bge -_gc ;if _fd >=0{_ff =int (_fd +0.5);}else {_ff =int (_fd -0.5);};if _fda >=0{_cc =int (_fda +0.5);}else {_cc =int (_fda -0.5);};if _af ,_fg =_aa .Get (_db );_fg !=nil {return _a .Wrap (_fg ,_ba ,"");};_bgb ,_bff :=_af .Min .X ,_af .Min .Y ;
_ffa ,_fg =_ggfg .UndilatedTemplates .GetBitmap (_ad );if _fg !=nil {return _a .Wrap (_fg ,_ba ,"\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029");
};_ac ,_fg =_fc (_caf ,_bgb ,_bff ,_ff ,_cc ,_ffa );if _fg !=nil {return _a .Wrap (_fg ,_ba ,"");};_ggfg .PtaUL .AddPoint (float32 (_bgb -_ff +_ac .X ),float32 (_bff -_cc +_ac .Y ));};return nil ;};