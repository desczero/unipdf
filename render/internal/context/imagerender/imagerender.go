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

package imagerender ;import (_f "errors";_g "fmt";_ea "github.com/unidoc/freetype/raster";_ge "github.com/unidoc/unipdf/v3/common";_ee "github.com/unidoc/unipdf/v3/internal/transform";_b "github.com/unidoc/unipdf/v3/render/internal/context";_be "golang.org/x/image/draw";
_gg "golang.org/x/image/font";_bd "golang.org/x/image/math/f64";_af "golang.org/x/image/math/fixed";_gc "image";_c "image/color";_eg "image/draw";_e "math";_ed "sort";_ae "strings";);func (_ceb *Context )DrawLine (x1 ,y1 ,x2 ,y2 float64 ){_ceb .MoveTo (x1 ,y1 );
_ceb .LineTo (x2 ,y2 )};const (_fdcg repeatOp =iota ;_cadeb ;_ddbb ;_gcc ;);type circle struct{_bee ,_dgd ,_cebd float64 };func NewContextForImage (im _gc .Image )*Context {return NewContextForRGBA (_aebc (im ))};func (_abfa *Context )drawString (_edb string ,_bce _gg .Face ,_dedf ,_bea float64 ){_cgg :=&_gg .Drawer {Src :_gc .NewUniform (_abfa ._dbf ),Face :_bce ,Dot :_dceaf (_ee .NewPoint (_dedf ,_bea ))};
_fge :=rune (-1);for _ ,_eeb :=range _edb {if _fge >=0{_cgg .Dot .X +=_cgg .Face .Kern (_fge ,_eeb );};_fcg ,_fegf ,_cecf ,_cgcc ,_eeg :=_cgg .Face .Glyph (_cgg .Dot ,_eeb );if !_eeg {continue ;};_cbd :=_fcg .Sub (_fcg .Min );_edaef :=_gc .NewRGBA (_cbd );
_be .DrawMask (_edaef ,_cbd ,_cgg .Src ,_gc .Point {},_fegf ,_cecf ,_be .Over );var _eegg *_be .Options ;if _abfa ._adf !=nil {_eegg =&_be .Options {DstMask :_abfa ._adf ,DstMaskP :_gc .Point {}};};_dgc :=_abfa ._aad .Clone ().Translate (float64 (_fcg .Min .X ),float64 (_fcg .Min .Y ));
_dbfb :=_bd .Aff3 {_dgc [0],_dgc [3],_dgc [6],_dgc [1],_dgc [4],_dgc [7]};_be .BiLinear .Transform (_abfa ._bef ,_dbfb ,_edaef ,_cbd ,_be .Over ,_eegg );_cgg .Dot .X +=_cgcc ;_fge =_eeb ;};};func (_bbcf *Context )FillPattern ()_b .Pattern {return _bbcf ._aec };
func (_ace *Context )Rotate (angle float64 ){_ace ._aad =_ace ._aad .Rotate (angle )};func (_fegg *Context )MeasureString (s string ,face _gg .Face )(_bgcda ,_fgdfb float64 ){_bca :=&_gg .Drawer {Face :face };_gbf :=_bca .MeasureString (s );return float64 (_gbf >>6),_fegg ._efbg .Tf .Size ;
};func (_age *Context )ClipPreserve (){_abbc :=_gc .NewAlpha (_gc .Rect (0,0,_age ._db ,_age ._eag ));_dce :=_ea .NewAlphaOverPainter (_abbc );_age .fill (_dce );if _age ._adf ==nil {_age ._adf =_abbc ;}else {_fdc :=_gc .NewAlpha (_gc .Rect (0,0,_age ._db ,_age ._eag ));
_be .DrawMask (_fdc ,_fdc .Bounds (),_abbc ,_gc .Point {},_age ._adf ,_gc .Point {},_be .Over );_age ._adf =_fdc ;};};func (_abd *Context )SetPixel (x ,y int ){_abd ._bef .Set (x ,y ,_abd ._dbf )};func (_afab *Context )LineTo (x ,y float64 ){if !_afab ._fac {_afab .MoveTo (x ,y );
}else {x ,y =_afab .Transform (x ,y );_cbf :=_ee .NewPoint (x ,y );_fgdg :=_dceaf (_cbf );_afab ._efb .Add1 (_fgdg );_afab ._ba .Add1 (_fgdg );_afab ._egg =_cbf ;};};func (_bab *Context )Clip (){_bab .ClipPreserve ();_bab .ClearPath ()};func (_ffd *Context )SetFillRGBA (r ,g ,b ,a float64 ){_dbc :=_c .NRGBA {uint8 (r *255),uint8 (g *255),uint8 (b *255),uint8 (a *255)};
_ffd ._dbf =_dbc ;_ffd ._aec =_cfce (_dbc );};func (_gee *Context )QuadraticTo (x1 ,y1 ,x2 ,y2 float64 ){if !_gee ._fac {_gee .MoveTo (x1 ,y1 );};x1 ,y1 =_gee .Transform (x1 ,y1 );x2 ,y2 =_gee .Transform (x2 ,y2 );_afe :=_ee .NewPoint (x1 ,y1 );_ffb :=_ee .NewPoint (x2 ,y2 );
_eac :=_dceaf (_afe );_bba :=_dceaf (_ffb );_gee ._efb .Add2 (_eac ,_bba );_gee ._ba .Add2 (_eac ,_bba );_gee ._egg =_ffb ;};func (_cef *Context )joiner ()_ea .Joiner {switch _cef ._ebfe {case _b .LineJoinBevel :return _ea .BevelJoiner ;case _b .LineJoinRound :return _ea .RoundJoiner ;
};return nil ;};func (_daa *Context )ResetClip (){_daa ._adf =nil };func (_gad *Context )TextState ()*_b .TextState {return &_gad ._efbg };func (_agf *Context )CubicTo (x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 ){if !_agf ._fac {_agf .MoveTo (x1 ,y1 );};_gdcb ,_bfd :=_agf ._egg .X ,_agf ._egg .Y ;
x1 ,y1 =_agf .Transform (x1 ,y1 );x2 ,y2 =_agf .Transform (x2 ,y2 );x3 ,y3 =_agf .Transform (x3 ,y3 );_cd :=_bgc (_gdcb ,_bfd ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 );_fcc :=_dceaf (_agf ._egg );for _ ,_faa :=range _cd [1:]{_gdb :=_dceaf (_faa );if _gdb ==_fcc {continue ;
};_fcc =_gdb ;_agf ._efb .Add1 (_gdb );_agf ._ba .Add1 (_gdb );_agf ._egg =_faa ;};};func (_ecdf *linearGradient )AddColorStop (offset float64 ,color _c .Color ){_ecdf ._edcc =append (_ecdf ._edcc ,stop {_bcab :offset ,_fedb :color });_ed .Sort (_ecdf ._edcc );
};func (_agg *Context )Matrix ()_ee .Matrix {return _agg ._aad };func (_adb *Context )DrawRectangle (x ,y ,w ,h float64 ){_adb .NewSubPath ();_adb .MoveTo (x ,y );_adb .LineTo (x +w ,y );_adb .LineTo (x +w ,y +h );_adb .LineTo (x ,y +h );_adb .ClosePath ();
};func (_bdg *Context )NewSubPath (){if _bdg ._fac {_bdg ._ba .Add1 (_dceaf (_bdg ._ggda ));};_bdg ._fac =false ;};func (_beb *Context )Clear (){_gdbc :=_gc .NewUniform (_beb ._dbf );_be .Draw (_beb ._bef ,_beb ._bef .Bounds (),_gdbc ,_gc .Point {},_be .Src );
};func NewContext (width ,height int )*Context {return NewContextForRGBA (_gc .NewRGBA (_gc .Rect (0,0,width ,height )));};func (_cbgd *Context )SetStrokeRGBA (r ,g ,b ,a float64 ){_dbe :=_c .NRGBA {uint8 (r *255),uint8 (g *255),uint8 (b *255),uint8 (a *255)};
_cbgd ._caa =_cfce (_dbe );};func (_cecg stops )Len ()int {return len (_cecg )};func _bgc (_gdf ,_abf ,_aage ,_ggf ,_ce ,_aba ,_add ,_efd float64 )[]_ee .Point {_bgg :=(_e .Hypot (_aage -_gdf ,_ggf -_abf )+_e .Hypot (_ce -_aage ,_aba -_ggf )+_e .Hypot (_add -_ce ,_efd -_aba ));
_bgcd :=int (_bgg +0.5);if _bgcd < 4{_bgcd =4;};_dg :=float64 (_bgcd )-1;_fd :=make ([]_ee .Point ,_bgcd );for _ebf :=0;_ebf < _bgcd ;_ebf ++{_cfa :=float64 (_ebf )/_dg ;_cgc ,_fce :=_bbg (_gdf ,_abf ,_aage ,_ggf ,_ce ,_aba ,_add ,_efd ,_cfa );_fd [_ebf ]=_ee .NewPoint (_cgc ,_fce );
};return _fd ;};func (_aea *Context )DrawEllipse (x ,y ,rx ,ry float64 ){_aea .NewSubPath ();_aea .DrawEllipticalArc (x ,y ,rx ,ry ,0,2*_e .Pi );_aea .ClosePath ();};func (_feb *Context )DrawImageAnchored (im _gc .Image ,x ,y int ,ax ,ay float64 ){_dgba :=im .Bounds ().Size ();
x -=int (ax *float64 (_dgba .X ));y -=int (ay *float64 (_dgba .Y ));_fbgg :=_be .BiLinear ;_ggb :=_feb ._aad .Clone ().Translate (float64 (x ),float64 (y ));_ebbd :=_bd .Aff3 {_ggb [0],_ggb [3],_ggb [6],_ggb [1],_ggb [4],_ggb [7]};if _feb ._adf ==nil {_fbgg .Transform (_feb ._bef ,_ebbd ,im ,im .Bounds (),_be .Over ,nil );
}else {_fbgg .Transform (_feb ._bef ,_ebbd ,im ,im .Bounds (),_be .Over ,&_be .Options {DstMask :_feb ._adf ,DstMaskP :_gc .Point {}});};};func (_fef *Context )SetDash (dashes ...float64 ){_fef ._edd =dashes };func _ad (_cb ,_cbe ,_d ,_geg ,_bb ,_aa ,_cc float64 )(_ca ,_fe float64 ){_dc :=1-_cc ;
_ac :=_dc *_dc ;_ede :=2*_dc *_cc ;_cf :=_cc *_cc ;_ca =_ac *_cb +_ede *_d +_cf *_bb ;_fe =_ac *_cbe +_ede *_geg +_cf *_aa ;return ;};func _bbab (_dgdg ,_abbe ,_afb ,_fcf ,_gaca ,_eff float64 )float64 {return _dgdg *_fcf +_abbe *_gaca +_afb *_eff ;};func (_bgd *patternPainter )Paint (ss []_ea .Span ,done bool ){_edg :=_bgd ._edeea .Bounds ();
for _ ,_fbgc :=range ss {if _fbgc .Y < _edg .Min .Y {continue ;};if _fbgc .Y >=_edg .Max .Y {return ;};if _fbgc .X0 < _edg .Min .X {_fbgc .X0 =_edg .Min .X ;};if _fbgc .X1 > _edg .Max .X {_fbgc .X1 =_edg .Max .X ;};if _fbgc .X0 >=_fbgc .X1 {continue ;};
const _ecae =1<<16-1;_abda :=_fbgc .Y -_bgd ._edeea .Rect .Min .Y ;_bacb :=_fbgc .X0 -_bgd ._edeea .Rect .Min .X ;_gga :=(_fbgc .Y -_bgd ._edeea .Rect .Min .Y )*_bgd ._edeea .Stride +(_fbgc .X0 -_bgd ._edeea .Rect .Min .X )*4;_bcc :=_gga +(_fbgc .X1 -_fbgc .X0 )*4;
for _ffdbe ,_gdgb :=_gga ,_bacb ;_ffdbe < _bcc ;_ffdbe ,_gdgb =_ffdbe +4,_gdgb +1{_dad :=_fbgc .Alpha ;if _bgd ._bcb !=nil {_dad =_dad *uint32 (_bgd ._bcb .AlphaAt (_gdgb ,_abda ).A )/255;if _dad ==0{continue ;};};_befc :=_bgd ._feed .ColorAt (_gdgb ,_abda );
_fde ,_cbdf ,_bbfa ,_edcd :=_befc .RGBA ();_dabd :=uint32 (_bgd ._edeea .Pix [_ffdbe +0]);_bda :=uint32 (_bgd ._edeea .Pix [_ffdbe +1]);_bgdd :=uint32 (_bgd ._edeea .Pix [_ffdbe +2]);_gabe :=uint32 (_bgd ._edeea .Pix [_ffdbe +3]);_gag :=(_ecae -(_edcd *_dad /_ecae ))*0x101;
_bgd ._edeea .Pix [_ffdbe +0]=uint8 ((_dabd *_gag +_fde *_dad )/_ecae >>8);_bgd ._edeea .Pix [_ffdbe +1]=uint8 ((_bda *_gag +_cbdf *_dad )/_ecae >>8);_bgd ._edeea .Pix [_ffdbe +2]=uint8 ((_bgdd *_gag +_bbfa *_dad )/_ecae >>8);_bgd ._edeea .Pix [_ffdbe +3]=uint8 ((_gabe *_gag +_edcd *_dad )/_ecae >>8);
};};};func (_bag *Context )DrawArc (x ,y ,r ,angle1 ,angle2 float64 ){_bag .DrawEllipticalArc (x ,y ,r ,r ,angle1 ,angle2 );};func (_fgec *Context )Transform (x ,y float64 )(_dbg ,_cga float64 ){return _fgec ._aad .Transform (x ,y )};func (_bbge *Context )DrawPoint (x ,y ,r float64 ){_bbge .Push ();
_ded ,_cgeb :=_bbge .Transform (x ,y );_bbge .Identity ();_bbge .DrawCircle (_ded ,_cgeb ,r );_bbge .Pop ();};type radialGradient struct{_ecg ,_abfc ,_fgda circle ;_agb ,_cfag float64 ;_bbfd float64 ;_fca stops ;};func (_cda *radialGradient )AddColorStop (offset float64 ,color _c .Color ){_cda ._fca =append (_cda ._fca ,stop {_bcab :offset ,_fedb :color });
_ed .Sort (_cda ._fca );};func (_fdd stops )Swap (i ,j int ){_fdd [i ],_fdd [j ]=_fdd [j ],_fdd [i ]};func (_eeda *Context )RotateAbout (angle ,x ,y float64 ){_eeda .Translate (x ,y );_eeda .Rotate (angle );_eeda .Translate (-x ,-y );};func _ga (_aaf ,_cg ,_ab ,_geb ,_aeg ,_gd float64 )[]_ee .Point {_bf :=(_e .Hypot (_ab -_aaf ,_geb -_cg )+_e .Hypot (_aeg -_ab ,_gd -_geb ));
_gf :=int (_bf +0.5);if _gf < 4{_gf =4;};_bc :=float64 (_gf )-1;_fa :=make ([]_ee .Point ,_gf );for _ggd :=0;_ggd < _gf ;_ggd ++{_fc :=float64 (_ggd )/_bc ;_gae ,_gebf :=_ad (_aaf ,_cg ,_ab ,_geb ,_aeg ,_gd ,_fc );_fa [_ggd ]=_ee .NewPoint (_gae ,_gebf );
};return _fa ;};func _bbg (_edee ,_gfa ,_afa ,_aac ,_ef ,_dcd ,_ff ,_gff ,_bfb float64 )(_bde ,_cag float64 ){_bg :=1-_bfb ;_eb :=_bg *_bg *_bg ;_aag :=3*_bg *_bg *_bfb ;_gcd :=3*_bg *_bfb *_bfb ;_bga :=_bfb *_bfb *_bfb ;_bde =_eb *_edee +_aag *_afa +_gcd *_ef +_bga *_ff ;
_cag =_eb *_gfa +_aag *_aac +_gcd *_dcd +_bga *_gff ;return ;};func _eddb (_aggc string )(_gfd ,_fcac ,_adbg ,_cecfc int ){_aggc =_ae .TrimPrefix (_aggc ,"\u0023");_cecfc =255;if len (_aggc )==3{_bge :="\u00251\u0078\u0025\u0031\u0078\u0025\u0031x";_g .Sscanf (_aggc ,_bge ,&_gfd ,&_fcac ,&_adbg );
_gfd |=_gfd <<4;_fcac |=_fcac <<4;_adbg |=_adbg <<4;};if len (_aggc )==6{_fbf :="\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078";_g .Sscanf (_aggc ,_fbf ,&_gfd ,&_fcac ,&_adbg );};if len (_aggc )==8{_acef :="\u0025\u00302\u0078\u0025\u00302\u0078\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078";
_g .Sscanf (_aggc ,_acef ,&_gfd ,&_fcac ,&_adbg ,&_cecfc );};return ;};func (_aae *Context )Width ()int {return _aae ._db };func (_gaf *Context )SetHexColor (x string ){_eee ,_eggc ,_aef ,_fgd :=_eddb (x );_gaf .SetRGBA255 (_eee ,_eggc ,_aef ,_fgd );};
type linearGradient struct{_caf ,_gdde ,_ccgd ,_bac float64 ;_edcc stops ;};func _bbce (_bbgec float64 )_af .Int26_6 {return _af .Int26_6 (_bbgec *64)};func (_gcg *Context )LineWidth ()float64 {return _gcg ._bbc };func (_fgdfc *Context )DrawImage (im _gc .Image ,x ,y int ){_fgdfc .DrawImageAnchored (im ,x ,y ,0,0)};
func (_abfb *Context )Translate (x ,y float64 ){_abfb ._aad =_abfb ._aad .Translate (x ,y )};func _cdad (_fcbb *_gc .RGBA ,_fgfg *_gc .Alpha ,_cced _b .Pattern )*patternPainter {return &patternPainter {_fcbb ,_fgfg ,_cced };};func (_deb *Context )FillPreserve (){var _efc _ea .Painter ;
if _deb ._adf ==nil {if _eddc ,_cfc :=_deb ._aec .(*solidPattern );_cfc {_gcb :=_ea .NewRGBAPainter (_deb ._bef );_gcb .SetColor (_eddc ._babcb );_efc =_gcb ;};};if _efc ==nil {_efc =_cdad (_deb ._bef ,_deb ._adf ,_deb ._aec );};_deb .fill (_efc );};func (_aeac *Context )Push (){_bad :=*_aeac ;
_aeac ._fcb =append (_aeac ._fcb ,&_bad )};func (_cdgg *linearGradient )ColorAt (x ,y int )_c .Color {if len (_cdgg ._edcc )==0{return _c .Transparent ;};_gef ,_cdd :=float64 (x ),float64 (y );_bfdc ,_gbe ,_bfbc ,_eaa :=_cdgg ._caf ,_cdgg ._gdde ,_cdgg ._ccgd ,_cdgg ._bac ;
_eaf ,_dff :=_bfbc -_bfdc ,_eaa -_gbe ;if _dff ==0&&_eaf !=0{return _ead ((_gef -_bfdc )/_eaf ,_cdgg ._edcc );};if _eaf ==0&&_dff !=0{return _ead ((_cdd -_gbe )/_dff ,_cdgg ._edcc );};_bcd :=_eaf *(_gef -_bfdc )+_dff *(_cdd -_gbe );if _bcd < 0{return _cdgg ._edcc [0]._fedb ;
};_ebc :=_e .Hypot (_eaf ,_dff );_edf :=((_gef -_bfdc )*-_dff +(_cdd -_gbe )*_eaf )/(_ebc *_ebc );_eagg ,_fcbad :=_bfdc +_edf *-_dff ,_gbe +_edf *_eaf ;_gacd :=_e .Hypot (_gef -_eagg ,_cdd -_fcbad )/_ebc ;return _ead (_gacd ,_cdgg ._edcc );};func (_gabf *Context )InvertMask (){if _gabf ._adf ==nil {_gabf ._adf =_gc .NewAlpha (_gabf ._bef .Bounds ());
}else {for _cbbd ,_ebb :=range _gabf ._adf .Pix {_gabf ._adf .Pix [_cbbd ]=255-_ebb ;};};};func _aebc (_aab _gc .Image )*_gc .RGBA {_eeaa :=_aab .Bounds ();_baf :=_gc .NewRGBA (_eeaa );_eg .Draw (_baf ,_eeaa ,_aab ,_eeaa .Min ,_eg .Src );return _baf ;};
func (_df *Context )StrokePreserve (){var _gfe _ea .Painter ;if _df ._adf ==nil {if _da ,_fdb :=_df ._caa .(*solidPattern );_fdb {_fbg :=_ea .NewRGBAPainter (_df ._bef );_fbg .SetColor (_da ._babcb );_gfe =_fbg ;};};if _gfe ==nil {_gfe =_cdad (_df ._bef ,_df ._adf ,_df ._caa );
};_df .stroke (_gfe );};type surfacePattern struct{_eeed _gc .Image ;_dac repeatOp ;};func _aca (_geed ,_ffc ,_cbda ,_dgf ,_cfec ,_dec float64 )_b .Gradient {_fda :=circle {_geed ,_ffc ,_cbda };_cade :=circle {_dgf ,_cfec ,_dec };_cecb :=circle {_dgf -_geed ,_cfec -_ffc ,_dec -_cbda };
_aceb :=_bbab (_cecb ._bee ,_cecb ._dgd ,-_cecb ._cebd ,_cecb ._bee ,_cecb ._dgd ,_cecb ._cebd );var _dab float64 ;if _aceb !=0{_dab =1.0/_aceb ;};_eggd :=-_fda ._cebd ;_cgae :=&radialGradient {_ecg :_fda ,_abfc :_cade ,_fgda :_cecb ,_agb :_aceb ,_cfag :_dab ,_bbfd :_eggd };
return _cgae ;};type repeatOp int ;func (_cagd *Context )SetLineJoin (lineJoin _b .LineJoin ){_cagd ._ebfe =lineJoin };func (_fgf *Context )DrawString (s string ,face _gg .Face ,x ,y float64 ){_fgf .DrawStringAnchored (s ,face ,x ,y ,0,0);};func (_gdfb *Context )SetLineCap (lineCap _b .LineCap ){_gdfb ._de =lineCap };
type stop struct{_bcab float64 ;_fedb _c .Color ;};type solidPattern struct{_babcb _c .Color };func (_fceb *Context )SetFillRule (fillRule _b .FillRule ){_fceb ._cbb =fillRule };func _cdaf (_ebg [][]_ee .Point )_ea .Path {var _aga _ea .Path ;for _ ,_ega :=range _ebg {var _cdf _af .Point26_6 ;
for _dbb ,_cfea :=range _ega {_acg :=_dceaf (_cfea );if _dbb ==0{_aga .Start (_acg );}else {_dfe :=_acg .X -_cdf .X ;_cfeaa :=_acg .Y -_cdf .Y ;if _dfe < 0{_dfe =-_dfe ;};if _cfeaa < 0{_cfeaa =-_cfeaa ;};if _dfe +_cfeaa > 8{_aga .Add1 (_acg );};};_cdf =_acg ;
};};return _aga ;};func (_fbae *Context )Scale (x ,y float64 ){_fbae ._aad =_fbae ._aad .Scale (x ,y )};func (_cfab *Context )MoveTo (x ,y float64 ){if _cfab ._fac {_cfab ._ba .Add1 (_dceaf (_cfab ._ggda ));};x ,y =_cfab .Transform (x ,y );_cce :=_ee .NewPoint (x ,y );
_bgf :=_dceaf (_cce );_cfab ._efb .Start (_bgf );_cfab ._ba .Start (_bgf );_cfab ._ggda =_cce ;_cfab ._egg =_cce ;_cfab ._fac =true ;};func (_cad *Context )stroke (_dgg _ea .Painter ){_cdb :=_cad ._efb ;if len (_cad ._edd )> 0{_cdb =_abff (_cdb ,_cad ._edd ,_cad ._ccg );
}else {_cdb =_cdaf (_edaf (_cdb ));};_ebfg :=_cad ._ag ;_ebfg .UseNonZeroWinding =true ;_ebfg .Clear ();_dgb :=(_cad ._aad .ScalingFactorX ()+_cad ._aad .ScalingFactorY ())/2;_ebfg .AddStroke (_cdb ,_bbce (_cad ._bbc *_dgb ),_cad .capper (),_cad .joiner ());
_ebfg .Rasterize (_dgg );};func (_ddg *Context )Pop (){_ada :=*_ddg ;_dfa :=_ddg ._fcb ;_ebe :=_dfa [len (_dfa )-1];*_ddg =*_ebe ;_ddg ._efb =_ada ._efb ;_ddg ._ba =_ada ._ba ;_ddg ._ggda =_ada ._ggda ;_ddg ._egg =_ada ._egg ;_ddg ._fac =_ada ._fac ;};
func (_ecd *Context )SetRGBA255 (r ,g ,b ,a int ){_ecd ._dbf =_c .NRGBA {uint8 (r ),uint8 (g ),uint8 (b ),uint8 (a )};_ecd .setFillAndStrokeColor (_ecd ._dbf );};func (_cfe *Context )SetLineWidth (lineWidth float64 ){_cfe ._bbc =lineWidth };func (_ccad *Context )ScaleAbout (sx ,sy ,x ,y float64 ){_ccad .Translate (x ,y );
_ccad .Scale (sx ,sy );_ccad .Translate (-x ,-y );};func (_ffgb stops )Less (i ,j int )bool {return _ffgb [i ]._bcab < _ffgb [j ]._bcab };func (_fbge *Context )drawRegularPolygon (_gac int ,_edc ,_gb ,_dcea ,_ccb float64 ){_dbcb :=2*_e .Pi /float64 (_gac );
_ccb -=_e .Pi /2;if _gac %2==0{_ccb +=_dbcb /2;};_fbge .NewSubPath ();for _dcc :=0;_dcc < _gac ;_dcc ++{_edae :=_ccb +_dbcb *float64 (_dcc );_fbge .LineTo (_edc +_dcea *_e .Cos (_edae ),_gb +_dcea *_e .Sin (_edae ));};_fbge .ClosePath ();};func (_befd *Context )StrokePattern ()_b .Pattern {return _befd ._caa };
func (_fed *Context )capper ()_ea .Capper {switch _fed ._de {case _b .LineCapButt :return _ea .ButtCapper ;case _b .LineCapRound :return _ea .RoundCapper ;case _b .LineCapSquare :return _ea .SquareCapper ;};return nil ;};func (_afaf *Context )Shear (x ,y float64 ){_afaf ._aad .Shear (x ,y )};
func _abff (_gec _ea .Path ,_cdc []float64 ,_egc float64 )_ea .Path {return _cdaf (_gbcb (_edaf (_gec ),_cdc ,_egc ));};var (_abb =_cfce (_c .White );_ggg =_cfce (_c .Black ););func (_cgcf *Context )ClearPath (){_cgcf ._efb .Clear ();_cgcf ._ba .Clear ();
_cgcf ._fac =false };func (_ffdb *Context )DrawCircle (x ,y ,r float64 ){_ffdb .NewSubPath ();_ffdb .DrawEllipticalArc (x ,y ,r ,r ,0,2*_e .Pi );_ffdb .ClosePath ();};func (_eda *Context )SetDashOffset (offset float64 ){_eda ._ccg =offset };type Context struct{_db int ;
_eag int ;_ag *_ea .Rasterizer ;_bef *_gc .RGBA ;_adf *_gc .Alpha ;_dbf _c .Color ;_aec _b .Pattern ;_caa _b .Pattern ;_efb _ea .Path ;_ba _ea .Path ;_ggda _ee .Point ;_egg _ee .Point ;_fac bool ;_edd []float64 ;_ccg float64 ;_bbc float64 ;_de _b .LineCap ;
_ebfe _b .LineJoin ;_cbb _b .FillRule ;_aad _ee .Matrix ;_efbg _b .TextState ;_fcb []*Context ;};func (_gdbd *radialGradient )ColorAt (x ,y int )_c .Color {if len (_gdbd ._fca )==0{return _c .Transparent ;};_eea ,_ege :=float64 (x )+0.5-_gdbd ._ecg ._bee ,float64 (y )+0.5-_gdbd ._ecg ._dgd ;
_fefa :=_bbab (_eea ,_ege ,_gdbd ._ecg ._cebd ,_gdbd ._fgda ._bee ,_gdbd ._fgda ._dgd ,_gdbd ._fgda ._cebd );_gdg :=_bbab (_eea ,_ege ,-_gdbd ._ecg ._cebd ,_eea ,_ege ,_gdbd ._ecg ._cebd );if _gdbd ._agb ==0{if _fefa ==0{return _c .Transparent ;};_fced :=0.5*_gdg /_fefa ;
if _fced *_gdbd ._fgda ._cebd >=_gdbd ._bbfd {return _ead (_fced ,_gdbd ._fca );};return _c .Transparent ;};_ebfb :=_bbab (_fefa ,_gdbd ._agb ,0,_fefa ,-_gdg ,0);if _ebfb >=0{_fae :=_e .Sqrt (_ebfb );_fga :=(_fefa +_fae )*_gdbd ._cfag ;_geff :=(_fefa -_fae )*_gdbd ._cfag ;
if _fga *_gdbd ._fgda ._cebd >=_gdbd ._bbfd {return _ead (_fga ,_gdbd ._fca );}else if _geff *_gdbd ._fgda ._cebd >=_gdbd ._bbfd {return _ead (_geff ,_gdbd ._fca );};};return _c .Transparent ;};func (_efbe *Context )SetMask (mask *_gc .Alpha )error {if mask .Bounds ().Size ()!=_efbe ._bef .Bounds ().Size (){return _f .New ("\u006d\u0061\u0073\u006b\u0020\u0073i\u007a\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068 \u0063\u006f\u006e\u0074\u0065\u0078\u0074 \u0073\u0069\u007a\u0065");
};_efbe ._adf =mask ;return nil ;};func _adaa (_gaffg ,_eaaa _c .Color ,_bbccf float64 )_c .Color {_fcd ,_egf ,_cac ,_gfec :=_gaffg .RGBA ();_cacd ,_fadd ,_aeb ,_bgcg :=_eaaa .RGBA ();return _c .RGBA {_cee (_fcd ,_cacd ,_bbccf ),_cee (_egf ,_fadd ,_bbccf ),_cee (_cac ,_aeb ,_bbccf ),_cee (_gfec ,_bgcg ,_bbccf )};
};func (_cbg *Context )SetColor (c _c .Color ){_cbg .setFillAndStrokeColor (c )};func (_ec *Context )SetFillStyle (pattern _b .Pattern ){if _cge ,_ffg :=pattern .(*solidPattern );_ffg {_ec ._dbf =_cge ._babcb ;};_ec ._aec =pattern ;};func _fdbf (_abgd float64 )float64 {return _abgd *_e .Pi /180};
func _edaf (_gfb _ea .Path )[][]_ee .Point {var _fabf [][]_ee .Point ;var _abg []_ee .Point ;var _gffc ,_abc float64 ;for _fdcb :=0;_fdcb < len (_gfb );{switch _gfb [_fdcb ]{case 0:if len (_abg )> 0{_fabf =append (_fabf ,_abg );_abg =nil ;};_eebf :=_eae (_gfb [_fdcb +1]);
_bbcg :=_eae (_gfb [_fdcb +2]);_abg =append (_abg ,_ee .NewPoint (_eebf ,_bbcg ));_gffc ,_abc =_eebf ,_bbcg ;_fdcb +=4;case 1:_ggfd :=_eae (_gfb [_fdcb +1]);_fdba :=_eae (_gfb [_fdcb +2]);_abg =append (_abg ,_ee .NewPoint (_ggfd ,_fdba ));_gffc ,_abc =_ggfd ,_fdba ;
_fdcb +=4;case 2:_eagb :=_eae (_gfb [_fdcb +1]);_ceee :=_eae (_gfb [_fdcb +2]);_dcaf :=_eae (_gfb [_fdcb +3]);_ffga :=_eae (_gfb [_fdcb +4]);_cgf :=_ga (_gffc ,_abc ,_eagb ,_ceee ,_dcaf ,_ffga );_abg =append (_abg ,_cgf ...);_gffc ,_abc =_dcaf ,_ffga ;
_fdcb +=6;case 3:_fbcd :=_eae (_gfb [_fdcb +1]);_cggf :=_eae (_gfb [_fdcb +2]);_gfff :=_eae (_gfb [_fdcb +3]);_cegc :=_eae (_gfb [_fdcb +4]);_egd :=_eae (_gfb [_fdcb +5]);_aagg :=_eae (_gfb [_fdcb +6]);_gadf :=_bgc (_gffc ,_abc ,_fbcd ,_cggf ,_gfff ,_cegc ,_egd ,_aagg );
_abg =append (_abg ,_gadf ...);_gffc ,_abc =_egd ,_aagg ;_fdcb +=8;default:_ge .Log .Debug ("\u0057\u0041\u0052\u004e: \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0074\u0068\u003a\u0020%\u0076",_gfb );return _fabf ;};};if len (_abg )> 0{_fabf =append (_fabf ,_abg );
};return _fabf ;};func (_fb *Context )Height ()int {return _fb ._eag };func (_facf *Context )ClosePath (){if _facf ._fac {_dd :=_dceaf (_facf ._ggda );_facf ._efb .Add1 (_dd );_facf ._ba .Add1 (_dd );_facf ._egg =_facf ._ggda ;};};func _gbcb (_fff [][]_ee .Point ,_abgb []float64 ,_ggc float64 )[][]_ee .Point {var _aafd [][]_ee .Point ;
if len (_abgb )==0{return _fff ;};if len (_abgb )==1{_abgb =append (_abgb ,_abgb [0]);};for _ ,_febc :=range _fff {if len (_febc )< 2{continue ;};_bbcgf :=_febc [0];_baa :=1;_ecb :=0;_gddd :=0.0;if _ggc !=0{var _eeae float64 ;for _ ,_bfg :=range _abgb {_eeae +=_bfg ;
};_ggc =_e .Mod (_ggc ,_eeae );if _ggc < 0{_ggc +=_eeae ;};for _dbca ,_fbgb :=range _abgb {_ggc -=_fbgb ;if _ggc < 0{_ecb =_dbca ;_gddd =_fbgb +_ggc ;break ;};};};var _gcf []_ee .Point ;_gcf =append (_gcf ,_bbcgf );for _baa < len (_febc ){_fgef :=_abgb [_ecb ];
_ebd :=_febc [_baa ];_dfg :=_bbcgf .Distance (_ebd );_eacd :=_fgef -_gddd ;if _dfg > _eacd {_abbb :=_eacd /_dfg ;_cae :=_bbcgf .Interpolate (_ebd ,_abbb );_gcf =append (_gcf ,_cae );if _ecb %2==0&&len (_gcf )> 1{_aafd =append (_aafd ,_gcf );};_gcf =nil ;
_gcf =append (_gcf ,_cae );_gddd =0;_bbcgf =_cae ;_ecb =(_ecb +1)%len (_abgb );}else {_gcf =append (_gcf ,_ebd );_bbcgf =_ebd ;_gddd +=_dfg ;_baa ++;};};if _ecb %2==0&&len (_gcf )> 1{_aafd =append (_aafd ,_gcf );};};return _aafd ;};func (_ecf *Context )DrawRoundedRectangle (x ,y ,w ,h ,r float64 ){_cbbf ,_gaed ,_fgdf ,_dfd :=x ,x +r ,x +w -r ,x +w ;
_aee ,_cec ,_agfe ,_dcb :=y ,y +r ,y +h -r ,y +h ;_ecf .NewSubPath ();_ecf .MoveTo (_gaed ,_aee );_ecf .LineTo (_fgdf ,_aee );_ecf .DrawArc (_fgdf ,_cec ,r ,_fdbf (270),_fdbf (360));_ecf .LineTo (_dfd ,_agfe );_ecf .DrawArc (_fgdf ,_agfe ,r ,_fdbf (0),_fdbf (90));
_ecf .LineTo (_gaed ,_dcb );_ecf .DrawArc (_gaed ,_agfe ,r ,_fdbf (90),_fdbf (180));_ecf .LineTo (_cbbf ,_cec );_ecf .DrawArc (_gaed ,_cec ,r ,_fdbf (180),_fdbf (270));_ecf .ClosePath ();};func _decf (_gadb _gc .Image ,_fgaa repeatOp )_b .Pattern {return &surfacePattern {_eeed :_gadb ,_dac :_fgaa };
};func (_fea *Context )Fill (){_fea .FillPreserve ();_fea .ClearPath ()};func (_bdd *Context )AsMask ()*_gc .Alpha {_bae :=_gc .NewAlpha (_bdd ._bef .Bounds ());_be .Draw (_bae ,_bdd ._bef .Bounds (),_bdd ._bef ,_gc .Point {},_be .Src );return _bae ;};
func _cee (_eacf ,_bagdg uint32 ,_adbc float64 )uint8 {return uint8 (int32 (float64 (_eacf )*(1.0-_adbc )+float64 (_bagdg )*_adbc )>>8);};func _ead (_dea float64 ,_abdd stops )_c .Color {if _dea <=0.0||len (_abdd )==1{return _abdd [0]._fedb ;};_bbe :=_abdd [len (_abdd )-1];
if _dea >=_bbe ._bcab {return _bbe ._fedb ;};for _eebg ,_gda :=range _abdd [1:]{if _dea < _gda ._bcab {_dea =(_dea -_abdd [_eebg ]._bcab )/(_gda ._bcab -_abdd [_eebg ]._bcab );return _adaa (_abdd [_eebg ]._fedb ,_gda ._fedb ,_dea );};};return _bbe ._fedb ;
};func (_fcag *solidPattern )ColorAt (x ,y int )_c .Color {return _fcag ._babcb };func (_fab *Context )DrawStringAnchored (s string ,face _gg .Face ,x ,y ,ax ,ay float64 ){_aaeb ,_dfdb :=_fab .MeasureString (s ,face );_fab .drawString (s ,face ,x -ax *_aaeb ,y +ay *_dfdb );
};type patternPainter struct{_edeea *_gc .RGBA ;_bcb *_gc .Alpha ;_feed _b .Pattern ;};func (_gbc *Context )Identity (){_gbc ._aad =_ee .IdentityMatrix ()};func (_gdd *Context )fill (_feg _ea .Painter ){_gab :=_gdd ._ba ;if _gdd ._fac {_gab =make (_ea .Path ,len (_gdd ._ba ));
copy (_gab ,_gdd ._ba );_gab .Add1 (_dceaf (_gdd ._ggda ));};_cbeb :=_gdd ._ag ;_cbeb .UseNonZeroWinding =_gdd ._cbb ==_b .FillRuleWinding ;_cbeb .Clear ();_cbeb .AddPath (_gab );_cbeb .Rasterize (_feg );};func (_gdc *Context )setFillAndStrokeColor (_fbc _c .Color ){_gdc ._dbf =_fbc ;
_gdc ._aec =_cfce (_fbc );_gdc ._caa =_cfce (_fbc );};type stops []stop ;func (_cgee *Context )DrawEllipticalArc (x ,y ,rx ,ry ,angle1 ,angle2 float64 ){const _ddb =16;for _abbcd :=0;_abbcd < _ddb ;_abbcd ++{_ecfd :=float64 (_abbcd +0)/_ddb ;_gaff :=float64 (_abbcd +1)/_ddb ;
_ece :=angle1 +(angle2 -angle1 )*_ecfd ;_caab :=angle1 +(angle2 -angle1 )*_gaff ;_bdf :=x +rx *_e .Cos (_ece );_cgd :=y +ry *_e .Sin (_ece );_daf :=x +rx *_e .Cos ((_ece +_caab )/2);_cbeg :=y +ry *_e .Sin ((_ece +_caab )/2);_bbcc :=x +rx *_e .Cos (_caab );
_cca :=y +ry *_e .Sin (_caab );_eed :=2*_daf -_bdf /2-_bbcc /2;_fee :=2*_cbeg -_cgd /2-_cca /2;if _abbcd ==0{if _cgee ._fac {_cgee .LineTo (_bdf ,_cgd );}else {_cgee .MoveTo (_bdf ,_cgd );};};_cgee .QuadraticTo (_eed ,_fee ,_bbcc ,_cca );};};func (_fcba *Context )SetRGB255 (r ,g ,b int ){_fcba .SetRGBA255 (r ,g ,b ,255)};
func (_agd *surfacePattern )ColorAt (x ,y int )_c .Color {_fbab :=_agd ._eeed .Bounds ();switch _agd ._dac {case _cadeb :if y >=_fbab .Dy (){return _c .Transparent ;};case _ddbb :if x >=_fbab .Dx (){return _c .Transparent ;};case _gcc :if x >=_fbab .Dx ()||y >=_fbab .Dy (){return _c .Transparent ;
};};x =x %_fbab .Dx ()+_fbab .Min .X ;y =y %_fbab .Dy ()+_fbab .Min .Y ;return _agd ._eeed .At (x ,y );};func _cfce (_bdff _c .Color )_b .Pattern {return &solidPattern {_babcb :_bdff }};func (_fba *Context )SetStrokeStyle (pattern _b .Pattern ){_fba ._caa =pattern };
func _dceaf (_gade _ee .Point )_af .Point26_6 {return _af .Point26_6 {X :_bbce (_gade .X ),Y :_bbce (_gade .Y )};};func (_ceg *Context )SetRGB (r ,g ,b float64 ){_ceg .SetRGBA (r ,g ,b ,1)};func _eae (_abcc _af .Int26_6 )float64 {const _efca ,_cea =6,1<<6-1;
if _abcc >=0{return float64 (_abcc >>_efca )+float64 (_abcc &_cea )/64;};_abcc =-_abcc ;if _abcc >=0{return -(float64 (_abcc >>_efca )+float64 (_abcc &_cea )/64);};return 0;};func (_fad *Context )SetRGBA (r ,g ,b ,a float64 ){_fad ._dbf =_c .NRGBA {uint8 (r *255),uint8 (g *255),uint8 (b *255),uint8 (a *255)};
_fad .setFillAndStrokeColor (_fad ._dbf );};func (_fg *Context )Image ()_gc .Image {return _fg ._bef };func (_babc *Context )ShearAbout (sx ,sy ,x ,y float64 ){_babc .Translate (x ,y );_babc .Shear (sx ,sy );_babc .Translate (-x ,-y );};func (_ffa *Context )Stroke (){_ffa .StrokePreserve ();
_ffa .ClearPath ()};func (_feeb *Context )SetMatrix (m _ee .Matrix ){_feeb ._aad =m };func _bagd (_dfb ,_afad ,_bbf ,_gbcc float64 )_b .Gradient {_dca :=&linearGradient {_caf :_dfb ,_gdde :_afad ,_ccgd :_bbf ,_bac :_gbcc };return _dca ;};func NewContextForRGBA (im *_gc .RGBA )*Context {_bcg :=im .Bounds ().Size ().X ;
_efg :=im .Bounds ().Size ().Y ;return &Context {_db :_bcg ,_eag :_efg ,_ag :_ea .NewRasterizer (_bcg ,_efg ),_bef :im ,_dbf :_c .Transparent ,_aec :_abb ,_caa :_ggg ,_bbc :1,_cbb :_b .FillRuleWinding ,_aad :_ee .IdentityMatrix (),_efbg :_b .NewTextState ()};
};