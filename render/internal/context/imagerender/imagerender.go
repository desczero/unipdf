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

package imagerender ;import (_dc "errors";_d "fmt";_ef "github.com/unidoc/freetype/raster";_de "github.com/unidoc/unipdf/v3/common";_g "github.com/unidoc/unipdf/v3/internal/transform";_ae "github.com/unidoc/unipdf/v3/render/internal/context";_ga "golang.org/x/image/draw";
_f "golang.org/x/image/font";_ge "golang.org/x/image/math/f64";_gc "golang.org/x/image/math/fixed";_e "image";_dd "image/color";_af "image/draw";_a "math";_bd "sort";_c "strings";);func (_fgb *Context )CubicTo (x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 ){if !_fgb ._gae {_fgb .MoveTo (x1 ,y1 );
};_aab ,_acbc :=_fgb ._cggb .X ,_fgb ._cggb .Y ;x1 ,y1 =_fgb .Transform (x1 ,y1 );x2 ,y2 =_fgb .Transform (x2 ,y2 );x3 ,y3 =_fgb .Transform (x3 ,y3 );_gfb :=_bgg (_aab ,_acbc ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 );_aec :=_cdf (_fgb ._cggb );for _ ,_fad :=range _gfb [1:]{_dbeb :=_cdf (_fad );
if _dbeb ==_aec {continue ;};_aec =_dbeb ;_fgb ._bdc .Add1 (_dbeb );_fgb ._cd .Add1 (_dbeb );_fgb ._cggb =_fad ;};};func (_ebg *Context )Fill (){_ebg .FillPreserve ();_ebg .ClearPath ()};func NewContextForRGBA (im *_e .RGBA )*Context {_fgda :=im .Bounds ().Size ().X ;
_eee :=im .Bounds ().Size ().Y ;return &Context {_bef :_fgda ,_aaf :_eee ,_dfc :_ef .NewRasterizer (_fgda ,_eee ),_gce :im ,_edf :_dd .Transparent ,_ded :_eg ,_fd :_cbd ,_cbg :1,_cgb :_ae .FillRuleWinding ,_cc :_g .IdentityMatrix (),_bdg :_ae .NewTextState ()};
};func (_dgc *Context )RotateAbout (angle ,x ,y float64 ){_dgc .Translate (x ,y );_dgc .Rotate (angle );_dgc .Translate (-x ,-y );};func (_cce *Context )SetFillStyle (pattern _ae .Pattern ){if _ca ,_ceg :=pattern .(*solidPattern );_ceg {_cce ._edf =_ca ._fdc ;
};_cce ._ded =pattern ;};func (_feb *Context )SetLineWidth (lineWidth float64 ){_feb ._cbg =lineWidth };func (_fb *Context )SetRGBA (r ,g ,b ,a float64 ){_fb ._edf =_dd .NRGBA {uint8 (r *255),uint8 (g *255),uint8 (b *255),uint8 (a *255)};_fb .setFillAndStrokeColor (_fb ._edf );
};func (_eb *Context )NewSubPath (){if _eb ._gae {_eb ._cd .Add1 (_cdf (_eb ._bc ));};_eb ._gae =false ;};func (_eagd *Context )Scale (x ,y float64 ){_eagd ._cc =_eagd ._cc .Scale (x ,y )};func (_fca stops )Len ()int {return len (_fca )};func (_abbf *Context )Height ()int {return _abbf ._aaf };
func (_bac *Context )Clear (){_adbag :=_e .NewUniform (_bac ._edf );_ga .Draw (_bac ._gce ,_bac ._gce .Bounds (),_adbag ,_e .Point {},_ga .Src );};func (_dfg *radialGradient )ColorAt (x ,y int )_dd .Color {if len (_dfg ._ebd )==0{return _dd .Transparent ;
};_cbb ,_cegdb :=float64 (x )+0.5-_dfg ._bcg ._aef ,float64 (y )+0.5-_dfg ._bcg ._bccg ;_abc :=_gacf (_cbb ,_cegdb ,_dfg ._bcg ._bgc ,_dfg ._dga ._aef ,_dfg ._dga ._bccg ,_dfg ._dga ._bgc );_gfd :=_gacf (_cbb ,_cegdb ,-_dfg ._bcg ._bgc ,_cbb ,_cegdb ,_dfg ._bcg ._bgc );
if _dfg ._dgaa ==0{if _abc ==0{return _dd .Transparent ;};_ecc :=0.5*_gfd /_abc ;if _ecc *_dfg ._dga ._bgc >=_dfg ._fegb {return _agb (_ecc ,_dfg ._ebd );};return _dd .Transparent ;};_eab :=_gacf (_abc ,_dfg ._dgaa ,0,_abc ,-_gfd ,0);if _eab >=0{_fggc :=_a .Sqrt (_eab );
_cgfc :=(_abc +_fggc )*_dfg ._bfce ;_degg :=(_abc -_fggc )*_dfg ._bfce ;if _cgfc *_dfg ._dga ._bgc >=_dfg ._fegb {return _agb (_cgfc ,_dfg ._ebd );}else if _degg *_dfg ._dga ._bgc >=_dfg ._fegb {return _agb (_degg ,_dfg ._ebd );};};return _dd .Transparent ;
};func (_ddgd *Context )SetFillRGBA (r ,g ,b ,a float64 ){_bb :=_dd .NRGBA {uint8 (r *255),uint8 (g *255),uint8 (b *255),uint8 (a *255)};_ddgd ._edf =_bb ;_ddgd ._ded =_aadbc (_bb );};func (_cgec *Context )ShearAbout (sx ,sy ,x ,y float64 ){_cgec .Translate (x ,y );
_cgec .Shear (sx ,sy );_cgec .Translate (-x ,-y );};func (_eaed *Context )Identity (){_eaed ._cc =_g .IdentityMatrix ()};func (_bcc *Context )SetStrokeStyle (pattern _ae .Pattern ){_bcc ._fd =pattern };func (_afb *Context )QuadraticTo (x1 ,y1 ,x2 ,y2 float64 ){if !_afb ._gae {_afb .MoveTo (x1 ,y1 );
};x1 ,y1 =_afb .Transform (x1 ,y1 );x2 ,y2 =_afb .Transform (x2 ,y2 );_dba :=_g .NewPoint (x1 ,y1 );_ead :=_g .NewPoint (x2 ,y2 );_eea :=_cdf (_dba );_ffcc :=_cdf (_ead );_afb ._bdc .Add2 (_eea ,_ffcc );_afb ._cd .Add2 (_eea ,_ffcc );_afb ._cggb =_ead ;
};func (_efa *Context )Matrix ()_g .Matrix {return _efa ._cc };func (_ecg *Context )SetPixel (x ,y int ){_ecg ._gce .Set (x ,y ,_ecg ._edf )};func _bgg (_dad ,_gcd ,_ce ,_afg ,_cgg ,_df ,_eae ,_aedcf float64 )[]_g .Point {_adba :=(_a .Hypot (_ce -_dad ,_afg -_gcd )+_a .Hypot (_cgg -_ce ,_df -_afg )+_a .Hypot (_eae -_cgg ,_aedcf -_df ));
_acb :=int (_adba +0.5);if _acb < 4{_acb =4;};_eag :=float64 (_acb )-1;_gea :=make ([]_g .Point ,_acb );for _aa :=0;_aa < _acb ;_aa ++{_aad :=float64 (_aa )/_eag ;_ddf ,_ed :=_fg (_dad ,_gcd ,_ce ,_afg ,_cgg ,_df ,_eae ,_aedcf ,_aad );_gea [_aa ]=_g .NewPoint (_ddf ,_ed );
};return _gea ;};func (_deg *Context )Clip (){_deg .ClipPreserve ();_deg .ClearPath ()};func NewContextForImage (im _e .Image )*Context {return NewContextForRGBA (_dgcf (im ))};func (_bfc *Context )SetMask (mask *_e .Alpha )error {if mask .Bounds ().Size ()!=_bfc ._gce .Bounds ().Size (){return _dc .New ("\u006d\u0061\u0073\u006b\u0020\u0073i\u007a\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068 \u0063\u006f\u006e\u0074\u0065\u0078\u0074 \u0073\u0069\u007a\u0065");
};_bfc ._abe =mask ;return nil ;};func NewContext (width ,height int )*Context {return NewContextForRGBA (_e .NewRGBA (_e .Rect (0,0,width ,height )));};func (_efba *Context )DrawPoint (x ,y ,r float64 ){_efba .Push ();_aacg ,_gdfb :=_efba .Transform (x ,y );
_efba .Identity ();_efba .DrawCircle (_aacg ,_gdfb ,r );_efba .Pop ();};func (_ba *Context )Width ()int {return _ba ._bef };func (_deca *Context )FillPreserve (){var _aeg _ef .Painter ;if _deca ._abe ==nil {if _daf ,_cgd :=_deca ._ded .(*solidPattern );
_cgd {_aff :=_ef .NewRGBAPainter (_deca ._gce );_aff .SetColor (_daf ._fdc );_aeg =_aff ;};};if _aeg ==nil {_aeg =_bgfg (_deca ._gce ,_deca ._abe ,_deca ._ded );};_deca .fill (_aeg );};func (_bed *Context )FillPattern ()_ae .Pattern {return _bed ._ded };
func (_ecd *Context )stroke (_eba _ef .Painter ){_eec :=_ecd ._bdc ;if len (_ecd ._aade )> 0{_eec =_fadb (_eec ,_ecd ._aade ,_ecd ._dee );}else {_eec =_gfe (_fdge (_eec ));};_abbe :=_ecd ._dfc ;_abbe .UseNonZeroWinding =true ;_abbe .Clear ();_caaa :=(_ecd ._cc .ScalingFactorX ()+_ecd ._cc .ScalingFactorY ())/2;
_abbe .AddStroke (_eec ,_fba (_ecd ._cbg *_caaa ),_ecd .capper (),_ecd .joiner ());_abbe .Rasterize (_eba );};func (_agf *Context )SetDash (dashes ...float64 ){_agf ._aade =dashes };func (_cff *Context )DrawStringAnchored (s string ,face _f .Face ,x ,y ,ax ,ay float64 ){_bdeb ,_eac :=_cff .MeasureString (s ,face );
_cff .drawString (s ,face ,x -ax *_bdeb ,y +ay *_eac );};func (_dgca *solidPattern )ColorAt (x ,y int )_dd .Color {return _dgca ._fdc };func (_gddb *Context )Stroke (){_gddb .StrokePreserve ();_gddb .ClearPath ()};func _fg (_ee ,_gge ,_efb ,_fgd ,_cf ,_adg ,_acgc ,_gbe ,_faf float64 )(_dce ,_dac float64 ){_aedc :=1-_faf ;
_fafg :=_aedc *_aedc *_aedc ;_dbe :=3*_aedc *_aedc *_faf ;_gee :=3*_aedc *_faf *_faf ;_dceb :=_faf *_faf *_faf ;_dce =_fafg *_ee +_dbe *_efb +_gee *_cf +_dceb *_acgc ;_dac =_fafg *_gge +_dbe *_fgd +_gee *_adg +_dceb *_gbe ;return ;};func (_dbb *Context )capper ()_ef .Capper {switch _dbb ._ag {case _ae .LineCapButt :return _ef .ButtCapper ;
case _ae .LineCapRound :return _ef .RoundCapper ;case _ae .LineCapSquare :return _ef .SquareCapper ;};return nil ;};func (_cge *Context )LineWidth ()float64 {return _cge ._cbg };func (_dec *Context )SetDashOffset (offset float64 ){_dec ._dee =offset };
func (_afbg *Context )Translate (x ,y float64 ){_afbg ._cc =_afbg ._cc .Translate (x ,y )};type radialGradient struct{_bcg ,_ddaf ,_dga circle ;_dgaa ,_bfce float64 ;_fegb float64 ;_ebd stops ;};func (_eecf stops )Swap (i ,j int ){_eecf [i ],_eecf [j ]=_eecf [j ],_eecf [i ]};
func (_cfe stops )Less (i ,j int )bool {return _cfe [i ]._dedf < _cfe [j ]._dedf };func (_ccab *Context )InvertMask (){if _ccab ._abe ==nil {_ccab ._abe =_e .NewAlpha (_ccab ._gce .Bounds ());}else {for _fda ,_ccad :=range _ccab ._abe .Pix {_ccab ._abe .Pix [_fda ]=255-_ccad ;
};};};func _aadbc (_daff _dd .Color )_ae .Pattern {return &solidPattern {_fdc :_daff }};type stops []stop ;func _gacf (_aae ,_eaa ,_ecgb ,_dcfa ,_agdd ,_efc float64 )float64 {return _aae *_dcfa +_eaa *_agdd +_ecgb *_efc ;};func _aeb (_bafd _e .Image ,_bfg repeatOp )_ae .Pattern {return &surfacePattern {_bdcbe :_bafd ,_cagg :_bfg };
};func (_ccg *radialGradient )AddColorStop (offset float64 ,color _dd .Color ){_ccg ._ebd =append (_ccg ._ebd ,stop {_dedf :offset ,_fcf :color });_bd .Sort (_ccg ._ebd );};func (_aba *Context )DrawImage (im _e .Image ,x ,y int ){_aba .DrawImageAnchored (im ,x ,y ,0,0)};
func _cdf (_faaa _g .Point )_gc .Point26_6 {return _gc .Point26_6 {X :_fba (_faaa .X ),Y :_fba (_faaa .Y )}};func (_baa *Context )DrawRectangle (x ,y ,w ,h float64 ){_baa .NewSubPath ();_baa .MoveTo (x ,y );_baa .LineTo (x +w ,y );_baa .LineTo (x +w ,y +h );
_baa .LineTo (x ,y +h );_baa .ClosePath ();};func (_dg *Context )SetLineJoin (lineJoin _ae .LineJoin ){_dg ._agc =lineJoin };func _bgfg (_gdfd *_e .RGBA ,_gca *_e .Alpha ,_gaeac _ae .Pattern )*patternPainter {return &patternPainter {_gdfd ,_gca ,_gaeac };
};type circle struct{_aef ,_bccg ,_bgc float64 };func (_ebab *Context )AsMask ()*_e .Alpha {_gba :=_e .NewAlpha (_ebab ._gce .Bounds ());_ga .Draw (_gba ,_ebab ._gce .Bounds (),_ebab ._gce ,_e .Point {},_ga .Src );return _gba ;};func (_bbe *Context )SetMatrix (m _g .Matrix ){_bbe ._cc =m };
type linearGradient struct{_gfg ,_bag ,_dgd ,_ddde float64 ;_geg stops ;};func _agb (_fcc float64 ,_egcf stops )_dd .Color {if _fcc <=0.0||len (_egcf )==1{return _egcf [0]._fcf ;};_bcgf :=_egcf [len (_egcf )-1];if _fcc >=_bcgf ._dedf {return _bcgf ._fcf ;
};for _cggf ,_agdg :=range _egcf [1:]{if _fcc < _agdg ._dedf {_fcc =(_fcc -_egcf [_cggf ]._dedf )/(_agdg ._dedf -_egcf [_cggf ]._dedf );return _bee (_egcf [_cggf ]._fcf ,_agdg ._fcf ,_fcc );};};return _bcgf ._fcf ;};type Context struct{_bef int ;_aaf int ;
_dfc *_ef .Rasterizer ;_gce *_e .RGBA ;_abe *_e .Alpha ;_edf _dd .Color ;_ded _ae .Pattern ;_fd _ae .Pattern ;_bdc _ef .Path ;_cd _ef .Path ;_bc _g .Point ;_cggb _g .Point ;_gae bool ;_aade []float64 ;_dee float64 ;_cbg float64 ;_ag _ae .LineCap ;_agc _ae .LineJoin ;
_cgb _ae .FillRule ;_cc _g .Matrix ;_bdg _ae .TextState ;_gggf []*Context ;};type solidPattern struct{_fdc _dd .Color };func (_dfd *Context )drawRegularPolygon (_fbg int ,_aee ,_dgec ,_cefb ,_gbgg float64 ){_fbga :=2*_a .Pi /float64 (_fbg );_gbgg -=_a .Pi /2;
if _fbg %2==0{_gbgg +=_fbga /2;};_dfd .NewSubPath ();for _fbd :=0;_fbd < _fbg ;_fbd ++{_baac :=_gbgg +_fbga *float64 (_fbd );_dfd .LineTo (_aee +_cefb *_a .Cos (_baac ),_dgec +_cefb *_a .Sin (_baac ));};_dfd .ClosePath ();};func (_cegd *Context )ScaleAbout (sx ,sy ,x ,y float64 ){_cegd .Translate (x ,y );
_cegd .Scale (sx ,sy );_cegd .Translate (-x ,-y );};func (_cad *Context )ResetClip (){_cad ._abe =nil };func _fba (_ccea float64 )_gc .Int26_6 {return _gc .Int26_6 (_ccea *64)};func (_cec *Context )ClosePath (){if _cec ._gae {_cecg :=_cdf (_cec ._bc );
_cec ._bdc .Add1 (_cecg );_cec ._cd .Add1 (_cecg );_cec ._cggb =_cec ._bc ;};};func (_cbeb *Context )SetStrokeRGBA (r ,g ,b ,a float64 ){_eeb :=_dd .NRGBA {uint8 (r *255),uint8 (g *255),uint8 (b *255),uint8 (a *255)};_cbeb ._fd =_aadbc (_eeb );};func (_ffd *Context )DrawRoundedRectangle (x ,y ,w ,h ,r float64 ){_bfa ,_bfd ,_fgg ,_gac :=x ,x +r ,x +w -r ,x +w ;
_ccef ,_ebf ,_ebc ,_fce :=y ,y +r ,y +h -r ,y +h ;_ffd .NewSubPath ();_ffd .MoveTo (_bfd ,_ccef );_ffd .LineTo (_fgg ,_ccef );_ffd .DrawArc (_fgg ,_ebf ,r ,_bcf (270),_bcf (360));_ffd .LineTo (_gac ,_ebc );_ffd .DrawArc (_fgg ,_ebc ,r ,_bcf (0),_bcf (90));
_ffd .LineTo (_bfd ,_fce );_ffd .DrawArc (_bfd ,_ebc ,r ,_bcf (90),_bcf (180));_ffd .LineTo (_bfa ,_ebf );_ffd .DrawArc (_bfd ,_ebf ,r ,_bcf (180),_bcf (270));_ffd .ClosePath ();};func NewLinearGradient (x0 ,y0 ,x1 ,y1 float64 )_ae .Gradient {_feg :=&linearGradient {_gfg :x0 ,_bag :y0 ,_dgd :x1 ,_ddde :y1 };
return _feg ;};func (_gcc *Context )DrawArc (x ,y ,r ,angle1 ,angle2 float64 ){_gcc .DrawEllipticalArc (x ,y ,r ,r ,angle1 ,angle2 );};func (_dcf *Context )Push (){_age :=*_dcf ;_dcf ._gggf =append (_dcf ._gggf ,&_age )};func _be (_fe ,_abb ,_gcb ,_adb ,_cb ,_ggg float64 )[]_g .Point {_ddd :=(_a .Hypot (_gcb -_fe ,_adb -_abb )+_a .Hypot (_cb -_gcb ,_ggg -_adb ));
_cg :=int (_ddd +0.5);if _cg < 4{_cg =4;};_db :=float64 (_cg )-1;_fa :=make ([]_g .Point ,_cg );for _ea :=0;_ea < _cg ;_ea ++{_ec :=float64 (_ea )/_db ;_ddg ,_ddc :=_ab (_fe ,_abb ,_gcb ,_adb ,_cb ,_ggg ,_ec );_fa [_ea ]=_g .NewPoint (_ddg ,_ddc );};return _fa ;
};var (_eg =_aadbc (_dd .White );_cbd =_aadbc (_dd .Black ););func (_dbg *Context )SetRGB (r ,g ,b float64 ){_dbg .SetRGBA (r ,g ,b ,1)};func (_cedg *Context )Shear (x ,y float64 ){_cedg ._cc .Shear (x ,y )};func _bcf (_aedg float64 )float64 {return _aedg *_a .Pi /180};
func (_ff *Context )setFillAndStrokeColor (_ffc _dd .Color ){_ff ._edf =_ffc ;_ff ._ded =_aadbc (_ffc );_ff ._fd =_aadbc (_ffc );};func (_gaeg *Context )Pop (){_abec :=*_gaeg ;_cdb :=_gaeg ._gggf ;_ccd :=_cdb [len (_cdb )-1];*_gaeg =*_ccd ;_gaeg ._bdc =_abec ._bdc ;
_gaeg ._cd =_abec ._cd ;_gaeg ._bc =_abec ._bc ;_gaeg ._cggb =_abec ._cggb ;_gaeg ._gae =_abec ._gae ;};func _bee (_fced ,_befe _dd .Color ,_gddf float64 )_dd .Color {_egg ,_gfbf ,_fafa ,_affd :=_fced .RGBA ();_ecf ,_dddg ,_gced ,_eadd :=_befe .RGBA ();
return _dd .RGBA {_gcca (_egg ,_ecf ,_gddf ),_gcca (_gfbf ,_dddg ,_gddf ),_gcca (_fafa ,_gced ,_gddf ),_gcca (_affd ,_eadd ,_gddf )};};func _gcbb (_egf string )(_bfef ,_fgae ,_ada ,_beae int ){_egf =_c .TrimPrefix (_egf ,"\u0023");_beae =255;if len (_egf )==3{_dbbd :="\u00251\u0078\u0025\u0031\u0078\u0025\u0031x";
_d .Sscanf (_egf ,_dbbd ,&_bfef ,&_fgae ,&_ada );_bfef |=_bfef <<4;_fgae |=_fgae <<4;_ada |=_ada <<4;};if len (_egf )==6{_gaf :="\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078";_d .Sscanf (_egf ,_gaf ,&_bfef ,&_fgae ,&_ada );};if len (_egf )==8{_dbfa :="\u0025\u00302\u0078\u0025\u00302\u0078\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078";
_d .Sscanf (_egf ,_dbfa ,&_bfef ,&_fgae ,&_ada ,&_beae );};return ;};type repeatOp int ;func (_gbeg *Context )DrawEllipse (x ,y ,rx ,ry float64 ){_gbeg .NewSubPath ();_gbeg .DrawEllipticalArc (x ,y ,rx ,ry ,0,2*_a .Pi );_gbeg .ClosePath ();};func (_gdeg *Context )joiner ()_ef .Joiner {switch _gdeg ._agc {case _ae .LineJoinBevel :return _ef .BevelJoiner ;
case _ae .LineJoinRound :return _ef .RoundJoiner ;};return nil ;};type surfacePattern struct{_bdcbe _e .Image ;_cagg repeatOp ;};func (_fee *Context )fill (_abbef _ef .Painter ){_bgeb :=_fee ._cd ;if _fee ._gae {_bgeb =make (_ef .Path ,len (_fee ._cd ));
copy (_bgeb ,_fee ._cd );_bgeb .Add1 (_cdf (_fee ._bc ));};_edc :=_fee ._dfc ;_edc .UseNonZeroWinding =_fee ._cgb ==_ae .FillRuleWinding ;_edc .Clear ();_edc .AddPath (_bgeb );_edc .Rasterize (_abbef );};func (_dggf *patternPainter )Paint (ss []_ef .Span ,done bool ){_cgab :=_dggf ._bga .Bounds ();
for _ ,_fga :=range ss {if _fga .Y < _cgab .Min .Y {continue ;};if _fga .Y >=_cgab .Max .Y {return ;};if _fga .X0 < _cgab .Min .X {_fga .X0 =_cgab .Min .X ;};if _fga .X1 > _cgab .Max .X {_fga .X1 =_cgab .Max .X ;};if _fga .X0 >=_fga .X1 {continue ;};const _fbc =1<<16-1;
_bdgd :=_fga .Y -_dggf ._bga .Rect .Min .Y ;_efe :=_fga .X0 -_dggf ._bga .Rect .Min .X ;_cded :=(_fga .Y -_dggf ._bga .Rect .Min .Y )*_dggf ._bga .Stride +(_fga .X0 -_dggf ._bga .Rect .Min .X )*4;_gcbf :=_cded +(_fga .X1 -_fga .X0 )*4;for _abaa ,_dfb :=_cded ,_efe ;
_abaa < _gcbf ;_abaa ,_dfb =_abaa +4,_dfb +1{_edg :=_fga .Alpha ;if _dggf ._abdg !=nil {_edg =_edg *uint32 (_dggf ._abdg .AlphaAt (_dfb ,_bdgd ).A )/255;if _edg ==0{continue ;};};_gdfc :=_dggf ._dgac .ColorAt (_dfb ,_bdgd );_bea ,_bda ,_fdf ,_bae :=_gdfc .RGBA ();
_eed :=uint32 (_dggf ._bga .Pix [_abaa +0]);_fgf :=uint32 (_dggf ._bga .Pix [_abaa +1]);_fcg :=uint32 (_dggf ._bga .Pix [_abaa +2]);_bbfb :=uint32 (_dggf ._bga .Pix [_abaa +3]);_aacc :=(_fbc -(_bae *_edg /_fbc ))*0x101;_dggf ._bga .Pix [_abaa +0]=uint8 ((_eed *_aacc +_bea *_edg )/_fbc >>8);
_dggf ._bga .Pix [_abaa +1]=uint8 ((_fgf *_aacc +_bda *_edg )/_fbc >>8);_dggf ._bga .Pix [_abaa +2]=uint8 ((_fcg *_aacc +_fdf *_edg )/_fbc >>8);_dggf ._bga .Pix [_abaa +3]=uint8 ((_bbfb *_aacc +_bae *_edg )/_fbc >>8);};};};func (_gaea *surfacePattern )ColorAt (x ,y int )_dd .Color {_aeea :=_gaea ._bdcbe .Bounds ();
switch _gaea ._cagg {case _bfdd :if y >=_aeea .Dy (){return _dd .Transparent ;};case _fgbf :if x >=_aeea .Dx (){return _dd .Transparent ;};case _aadb :if x >=_aeea .Dx ()||y >=_aeea .Dy (){return _dd .Transparent ;};};x =x %_aeea .Dx ()+_aeea .Min .X ;
y =y %_aeea .Dy ()+_aeea .Min .Y ;return _gaea ._bdcbe .At (x ,y );};func _gcca (_caac ,_gag uint32 ,_fefb float64 )uint8 {return uint8 (int32 (float64 (_caac )*(1.0-_fefb )+float64 (_gag )*_fefb )>>8);};const (_bdf repeatOp =iota ;_bfdd ;_fgbf ;_aadb ;
);func (_gde *Context )Image ()_e .Image {return _gde ._gce };func (_eda *Context )DrawImageAnchored (im _e .Image ,x ,y int ,ax ,ay float64 ){_agd :=im .Bounds ().Size ();x -=int (ax *float64 (_agd .X ));y -=int (ay *float64 (_agd .Y ));_ace :=_ga .BiLinear ;
_ece :=_eda ._cc .Clone ().Translate (float64 (x ),float64 (y ));_fec :=_ge .Aff3 {_ece [0],_ece [3],_ece [6],_ece [1],_ece [4],_ece [7]};if _eda ._abe ==nil {_ace .Transform (_eda ._gce ,_fec ,im ,im .Bounds (),_ga .Over ,nil );}else {_ace .Transform (_eda ._gce ,_fec ,im ,im .Bounds (),_ga .Over ,&_ga .Options {DstMask :_eda ._abe ,DstMaskP :_e .Point {}});
};};func (_cgeg *Context )MeasureString (s string ,face _f .Face )(_gff ,_gdc float64 ){_cdd :=&_f .Drawer {Face :face };_dff :=_cdd .MeasureString (s );return float64 (_dff >>6),_cgeg ._bdg .Tf .Size ;};func _dgcf (_cfa _e .Image )*_e .RGBA {_degd :=_cfa .Bounds ();
_fdd :=_e .NewRGBA (_degd );_af .Draw (_fdd ,_degd ,_cfa ,_degd .Min ,_af .Src );return _fdd ;};func (_ccc *Context )Rotate (angle float64 ){_ccc ._cc =_ccc ._cc .Rotate (angle )};type stop struct{_dedf float64 ;_fcf _dd .Color ;};type patternPainter struct{_bga *_e .RGBA ;
_abdg *_e .Alpha ;_dgac _ae .Pattern ;};func (_afa *Context )ClipPreserve (){_afbd :=_e .NewAlpha (_e .Rect (0,0,_afa ._bef ,_afa ._aaf ));_bf :=_ef .NewAlphaOverPainter (_afbd );_afa .fill (_bf );if _afa ._abe ==nil {_afa ._abe =_afbd ;}else {_agce :=_e .NewAlpha (_e .Rect (0,0,_afa ._bef ,_afa ._aaf ));
_ga .DrawMask (_agce ,_agce .Bounds (),_afbd ,_e .Point {},_afa ._abe ,_e .Point {},_ga .Over );_afa ._abe =_agce ;};};func (_faa *Context )TextState ()*_ae .TextState {return &_faa ._bdg };func (_abd *linearGradient )AddColorStop (offset float64 ,color _dd .Color ){_abd ._geg =append (_abd ._geg ,stop {_dedf :offset ,_fcf :color });
_bd .Sort (_abd ._geg );};func _cced (_gcede [][]_g .Point ,_cfd []float64 ,_fdb float64 )[][]_g .Point {var _abcg [][]_g .Point ;if len (_cfd )==0{return _gcede ;};if len (_cfd )==1{_cfd =append (_cfd ,_cfd [0]);};for _ ,_cbc :=range _gcede {if len (_cbc )< 2{continue ;
};_bbcd :=_cbc [0];_dada :=1;_edcg :=0;_ffa :=0.0;if _fdb !=0{var _decae float64 ;for _ ,_dde :=range _cfd {_decae +=_dde ;};_fdb =_a .Mod (_fdb ,_decae );if _fdb < 0{_fdb +=_decae ;};for _dgg ,_baged :=range _cfd {_fdb -=_baged ;if _fdb < 0{_edcg =_dgg ;
_ffa =_baged +_fdb ;break ;};};};var _dcfc []_g .Point ;_dcfc =append (_dcfc ,_bbcd );for _dada < len (_cbc ){_dfcc :=_cfd [_edcg ];_acbb :=_cbc [_dada ];_eccc :=_bbcd .Distance (_acbb );_abbb :=_dfcc -_ffa ;if _eccc > _abbb {_egb :=_abbb /_eccc ;_fbf :=_bbcd .Interpolate (_acbb ,_egb );
_dcfc =append (_dcfc ,_fbf );if _edcg %2==0&&len (_dcfc )> 1{_abcg =append (_abcg ,_dcfc );};_dcfc =nil ;_dcfc =append (_dcfc ,_fbf );_ffa =0;_bbcd =_fbf ;_edcg =(_edcg +1)%len (_cfd );}else {_dcfc =append (_dcfc ,_acbb );_bbcd =_acbb ;_ffa +=_eccc ;_dada ++;
};};if _edcg %2==0&&len (_dcfc )> 1{_abcg =append (_abcg ,_dcfc );};};return _abcg ;};func (_bfb *Context )DrawEllipticalArc (x ,y ,rx ,ry ,angle1 ,angle2 float64 ){const _fcd =16;for _bcd :=0;_bcd < _fcd ;_bcd ++{_ede :=float64 (_bcd +0)/_fcd ;_dbc :=float64 (_bcd +1)/_fcd ;
_bad :=angle1 +(angle2 -angle1 )*_ede ;_cef :=angle1 +(angle2 -angle1 )*_dbc ;_dddc :=x +rx *_a .Cos (_bad );_cga :=y +ry *_a .Sin (_bad );_gcf :=x +rx *_a .Cos ((_bad +_cef )/2);_cba :=y +ry *_a .Sin ((_bad +_cef )/2);_ggd :=x +rx *_a .Cos (_cef );_fac :=y +ry *_a .Sin (_cef );
_bgb :=2*_gcf -_dddc /2-_ggd /2;_gcfg :=2*_cba -_cga /2-_fac /2;if _bcd ==0{if _bfb ._gae {_bfb .LineTo (_dddc ,_cga );}else {_bfb .MoveTo (_dddc ,_cga );};};_bfb .QuadraticTo (_bgb ,_gcfg ,_ggd ,_fac );};};func (_aga *Context )DrawString (s string ,face _f .Face ,x ,y float64 ){_aga .DrawStringAnchored (s ,face ,x ,y ,0,0);
};func (_gcbg *Context )SetHexColor (x string ){_caa ,_bgf ,_bde ,_geed :=_gcbb (x );_gcbg .SetRGBA255 (_caa ,_bgf ,_bde ,_geed );};func (_bfab *Context )Transform (x ,y float64 )(_dbbe ,_afc float64 ){return _bfab ._cc .Transform (x ,y )};func _fadb (_dbd _ef .Path ,_efbg []float64 ,_bdbe float64 )_ef .Path {return _gfe (_cced (_fdge (_dbd ),_efbg ,_bdbe ));
};func (_bbf *Context )SetRGBA255 (r ,g ,b ,a int ){_bbf ._edf =_dd .NRGBA {uint8 (r ),uint8 (g ),uint8 (b ),uint8 (a )};_bbf .setFillAndStrokeColor (_bbf ._edf );};func (_cca *Context )ClearPath (){_cca ._bdc .Clear ();_cca ._cd .Clear ();_cca ._gae =false };
func (_bdcb *linearGradient )ColorAt (x ,y int )_dd .Color {if len (_bdcb ._geg )==0{return _dd .Transparent ;};_gaec ,_gdde :=float64 (x ),float64 (y );_dgf ,_afd ,_edfb ,_fff :=_bdcb ._gfg ,_bdcb ._bag ,_bdcb ._dgd ,_bdcb ._ddde ;_aacb ,_gbae :=_edfb -_dgf ,_fff -_afd ;
if _gbae ==0&&_aacb !=0{return _agb ((_gaec -_dgf )/_aacb ,_bdcb ._geg );};if _aacb ==0&&_gbae !=0{return _agb ((_gdde -_afd )/_gbae ,_bdcb ._geg );};_fef :=_aacb *(_gaec -_dgf )+_gbae *(_gdde -_afd );if _fef < 0{return _bdcb ._geg [0]._fcf ;};_aaa :=_a .Hypot (_aacb ,_gbae );
_aea :=((_gaec -_dgf )*-_gbae +(_gdde -_afd )*_aacb )/(_aaa *_aaa );_dbf ,_acc :=_dgf +_aea *-_gbae ,_afd +_aea *_aacb ;_ffg :=_a .Hypot (_gaec -_dbf ,_gdde -_acc )/_aaa ;return _agb (_ffg ,_bdcb ._geg );};func (_abg *Context )drawString (_fceg string ,_ebcb _f .Face ,_cde ,_bfe float64 ){_cgf :=&_f .Drawer {Src :_e .NewUniform (_abg ._edf ),Face :_ebcb ,Dot :_cdf (_g .NewPoint (_cde ,_bfe ))};
_ebae :=rune (-1);for _ ,_gdec :=range _fceg {if _ebae >=0{_cgf .Dot .X +=_cgf .Face .Kern (_ebae ,_gdec );};_ddb ,_bba ,_bcb ,_gccg ,_fag :=_cgf .Face .Glyph (_cgf .Dot ,_gdec );if !_fag {continue ;};_feca :=_ddb .Sub (_ddb .Min );_cgac :=_e .NewRGBA (_feca );
_ga .DrawMask (_cgac ,_feca ,_cgf .Src ,_e .Point {},_bba ,_bcb ,_ga .Over );var _dace *_ga .Options ;if _abg ._abe !=nil {_dace =&_ga .Options {DstMask :_abg ._abe ,DstMaskP :_e .Point {}};};_cgfe :=_abg ._cc .Clone ().Translate (float64 (_ddb .Min .X ),float64 (_ddb .Min .Y ));
_fafd :=_ge .Aff3 {_cgfe [0],_cgfe [3],_cgfe [6],_cgfe [1],_cgfe [4],_cgfe [7]};_ga .BiLinear .Transform (_abg ._gce ,_fafd ,_cgac ,_feca ,_ga .Over ,_dace );_cgf .Dot .X +=_gccg ;_ebae =_gdec ;};};func (_gdf *Context )StrokePattern ()_ae .Pattern {return _gdf ._fd };
func (_dbag *Context )StrokePreserve (){var _ged _ef .Painter ;if _dbag ._abe ==nil {if _gbg ,_cee :=_dbag ._fd .(*solidPattern );_cee {_fgc :=_ef .NewRGBAPainter (_dbag ._gce );_fgc .SetColor (_gbg ._fdc );_ged =_fgc ;};};if _ged ==nil {_ged =_bgfg (_dbag ._gce ,_dbag ._abe ,_dbag ._fd );
};_dbag .stroke (_ged );};func (_dda *Context )DrawCircle (x ,y ,r float64 ){_dda .NewSubPath ();_dda .DrawEllipticalArc (x ,y ,r ,r ,0,2*_a .Pi );_dda .ClosePath ();};func (_bdb *Context )SetColor (c _dd .Color ){_bdb .setFillAndStrokeColor (c )};func _gfe (_gacd [][]_g .Point )_ef .Path {var _aca _ef .Path ;
for _ ,_fcce :=range _gacd {var _gdb _gc .Point26_6 ;for _dgfb ,_cgad :=range _fcce {_aaab :=_cdf (_cgad );if _dgfb ==0{_aca .Start (_aaab );}else {_dfca :=_aaab .X -_gdb .X ;_efg :=_aaab .Y -_gdb .Y ;if _dfca < 0{_dfca =-_dfca ;};if _efg < 0{_efg =-_efg ;
};if _dfca +_efg > 8{_aca .Add1 (_aaab );};};_gdb =_aaab ;};};return _aca ;};func (_gdd *Context )MoveTo (x ,y float64 ){if _gdd ._gae {_gdd ._cd .Add1 (_cdf (_gdd ._bc ));};x ,y =_gdd .Transform (x ,y );_fdg :=_g .NewPoint (x ,y );_bge :=_cdf (_fdg );
_gdd ._bdc .Start (_bge );_gdd ._cd .Start (_bge );_gdd ._bc =_fdg ;_gdd ._cggb =_fdg ;_gdd ._gae =true ;};func (_aac *Context )SetLineCap (lineCap _ae .LineCap ){_aac ._ag =lineCap };func _ab (_dcc ,_geb ,_gd ,_ac ,_gf ,_gb ,_gg float64 )(_da ,_ad float64 ){_fc :=1-_gg ;
_bg :=_fc *_fc ;_aed :=2*_fc *_gg ;_acg :=_gg *_gg ;_da =_bg *_dcc +_aed *_gd +_acg *_gf ;_ad =_bg *_geb +_aed *_ac +_acg *_gb ;return ;};func _fbgab (_fegd _gc .Int26_6 )float64 {const _acbd ,_fgad =6,1<<6-1;if _fegd >=0{return float64 (_fegd >>_acbd )+float64 (_fegd &_fgad )/64;
};_fegd =-_fegd ;if _fegd >=0{return -(float64 (_fegd >>_acbd )+float64 (_fegd &_fgad )/64);};return 0;};func (_cbe *Context )SetFillRule (fillRule _ae .FillRule ){_cbe ._cgb =fillRule };func (_cag *Context )DrawLine (x1 ,y1 ,x2 ,y2 float64 ){_cag .MoveTo (x1 ,y1 );
_cag .LineTo (x2 ,y2 )};func _fdge (_ffb _ef .Path )[][]_g .Point {var _beg [][]_g .Point ;var _dea []_g .Point ;var _gef ,_gcdg float64 ;for _dacc :=0;_dacc < len (_ffb );{switch _ffb [_dacc ]{case 0:if len (_dea )> 0{_beg =append (_beg ,_dea );_dea =nil ;
};_dgab :=_fbgab (_ffb [_dacc +1]);_dgaf :=_fbgab (_ffb [_dacc +2]);_dea =append (_dea ,_g .NewPoint (_dgab ,_dgaf ));_gef ,_gcdg =_dgab ,_dgaf ;_dacc +=4;case 1:_edaa :=_fbgab (_ffb [_dacc +1]);_cddc :=_fbgab (_ffb [_dacc +2]);_dea =append (_dea ,_g .NewPoint (_edaa ,_cddc ));
_gef ,_gcdg =_edaa ,_cddc ;_dacc +=4;case 2:_cadd :=_fbgab (_ffb [_dacc +1]);_fde :=_fbgab (_ffb [_dacc +2]);_cac :=_fbgab (_ffb [_dacc +3]);_ccf :=_fbgab (_ffb [_dacc +4]);_ceda :=_be (_gef ,_gcdg ,_cadd ,_fde ,_cac ,_ccf );_dea =append (_dea ,_ceda ...);
_gef ,_gcdg =_cac ,_ccf ;_dacc +=6;case 3:_baf :=_fbgab (_ffb [_dacc +1]);_abeb :=_fbgab (_ffb [_dacc +2]);_deb :=_fbgab (_ffb [_dacc +3]);_ega :=_fbgab (_ffb [_dacc +4]);_bage :=_fbgab (_ffb [_dacc +5]);_eaf :=_fbgab (_ffb [_dacc +6]);_eff :=_bgg (_gef ,_gcdg ,_baf ,_abeb ,_deb ,_ega ,_bage ,_eaf );
_dea =append (_dea ,_eff ...);_gef ,_gcdg =_bage ,_eaf ;_dacc +=8;default:_de .Log .Debug ("\u0057\u0041\u0052\u004e: \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0074\u0068\u003a\u0020%\u0076",_ffb );return _beg ;};};if len (_dea )> 0{_beg =append (_beg ,_dea );
};return _beg ;};func (_febf *Context )SetRGB255 (r ,g ,b int ){_febf .SetRGBA255 (r ,g ,b ,255)};func NewRadialGradient (x0 ,y0 ,r0 ,x1 ,y1 ,r1 float64 )_ae .Gradient {_ffcf :=circle {x0 ,y0 ,r0 };_bbc :=circle {x1 ,y1 ,r1 };_gfa :=circle {x1 -x0 ,y1 -y0 ,r1 -r0 };
_dca :=_gacf (_gfa ._aef ,_gfa ._bccg ,-_gfa ._bgc ,_gfa ._aef ,_gfa ._bccg ,_gfa ._bgc );var _cagc float64 ;if _dca !=0{_cagc =1.0/_dca ;};_eead :=-_ffcf ._bgc ;_egc :=&radialGradient {_bcg :_ffcf ,_ddaf :_bbc ,_dga :_gfa ,_dgaa :_dca ,_bfce :_cagc ,_fegb :_eead };
return _egc ;};func (_ced *Context )LineTo (x ,y float64 ){if !_ced ._gae {_ced .MoveTo (x ,y );}else {x ,y =_ced .Transform (x ,y );_eage :=_g .NewPoint (x ,y );_gbea :=_cdf (_eage );_ced ._bdc .Add1 (_gbea );_ced ._cd .Add1 (_gbea );_ced ._cggb =_eage ;
};};