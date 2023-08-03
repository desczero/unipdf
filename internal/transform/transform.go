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

package transform ;import (_f "fmt";_b "github.com/unidoc/unipdf/v3/common";_d "math";);func RotationMatrix (angle float64 )Matrix {_c :=_d .Cos (angle );_ca :=_d .Sin (angle );return NewMatrix (_c ,_ca ,-_ca ,_c ,0,0);};func (_bgcc Matrix )Angle ()float64 {_bfe :=_d .Atan2 (-_bgcc [1],_bgcc [0]);
if _bfe < 0.0{_bfe +=2*_d .Pi ;};return _bfe /_d .Pi *180.0;};func ScaleMatrix (x ,y float64 )Matrix {return NewMatrix (x ,0,0,y ,0,0)};func (_gf Matrix )Rotate (theta float64 )Matrix {return _gf .Mult (RotationMatrix (theta ))};func NewMatrix (a ,b ,c ,d ,tx ,ty float64 )Matrix {_bg :=Matrix {a ,b ,0,c ,d ,0,tx ,ty ,1};
_bg .clampRange ();return _bg ;};func (_ffb Matrix )String ()string {_ag ,_fd ,_g ,_fda ,_gc ,_dg :=_ffb [0],_ffb [1],_ffb [3],_ffb [4],_ffb [6],_ffb [7];return _f .Sprintf ("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d",_ag ,_fd ,_g ,_fda ,_gc ,_dg );
};func (_ac *Matrix )Shear (x ,y float64 ){_ac .Concat (ShearMatrix (x ,y ))};func ShearMatrix (x ,y float64 )Matrix {return NewMatrix (1,y ,x ,1,0,0)};func (_ce Matrix )Translate (tx ,ty float64 )Matrix {return _ce .Mult (TranslationMatrix (tx ,ty ))};
func (_dgcf *Point )Transform (a ,b ,c ,d ,tx ,ty float64 ){_dc :=NewMatrix (a ,b ,c ,d ,tx ,ty );_dgcf .transformByMatrix (_dc );};func (_a Matrix )Identity ()bool {return _a [0]==1&&_a [1]==0&&_a [2]==0&&_a [3]==0&&_a [4]==1&&_a [5]==0&&_a [6]==0&&_a [7]==0&&_a [8]==1;
};func TranslationMatrix (tx ,ty float64 )Matrix {return NewMatrix (1,0,0,1,tx ,ty )};func (_ff Matrix )Round (precision float64 )Matrix {for _ef :=range _ff {_ff [_ef ]=_d .Round (_ff [_ef ]/precision )*precision ;};return _ff ;};func (_ab Matrix )Singular ()bool {return _d .Abs (_ab [0]*_ab [4]-_ab [1]*_ab [3])< _ad };
func (_gcc Matrix )Scale (xScale ,yScale float64 )Matrix {return _gcc .Mult (ScaleMatrix (xScale ,yScale ));};const _efc =1e9;func (_cd *Matrix )Clone ()Matrix {return NewMatrix (_cd [0],_cd [1],_cd [3],_cd [4],_cd [6],_cd [7])};func IdentityMatrix ()Matrix {return NewMatrix (1,0,0,1,0,0)};
func (_cdf Matrix )Inverse ()(Matrix ,bool ){_cbg ,_ee :=_cdf [0],_cdf [1];_aa ,_bb :=_cdf [3],_cdf [4];_eb ,_abb :=_cdf [6],_cdf [7];_ec :=_cbg *_bb -_ee *_aa ;if _d .Abs (_ec )< _dd {return Matrix {},false ;};_dgc ,_bdd :=_bb /_ec ,-_ee /_ec ;_gcf ,_eef :=-_aa /_ec ,_cbg /_ec ;
_gcd :=-(_dgc *_eb +_gcf *_abb );_ffc :=-(_bdd *_eb +_eef *_abb );return NewMatrix (_dgc ,_bdd ,_gcf ,_eef ,_gcd ,_ffc ),true ;};func (_cg Matrix )ScalingFactorX ()float64 {return _d .Hypot (_cg [0],_cg [1])};func (_bf Matrix )Transform (x ,y float64 )(float64 ,float64 ){_df :=x *_bf [0]+y *_bf [3]+_bf [6];
_cec :=x *_bf [1]+y *_bf [4]+_bf [7];return _df ,_cec ;};func (_da Point )String ()string {return _f .Sprintf ("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029",_da .X ,_da .Y );};func (_abc Point )Rotate (theta float64 )Point {_fa :=_d .Hypot (_abc .X ,_abc .Y );
_cf :=_d .Atan2 (_abc .Y ,_abc .X );_adb ,_cgg :=_d .Sincos (_cf +theta /180.0*_d .Pi );return Point {_fa *_cgg ,_fa *_adb };};const _ad =1e-10;func (_ga *Matrix )Concat (b Matrix ){*_ga =Matrix {b [0]*_ga [0]+b [1]*_ga [3],b [0]*_ga [1]+b [1]*_ga [4],0,b [3]*_ga [0]+b [4]*_ga [3],b [3]*_ga [1]+b [4]*_ga [4],0,b [6]*_ga [0]+b [7]*_ga [3]+_ga [6],b [6]*_ga [1]+b [7]*_ga [4]+_ga [7],1};
_ga .clampRange ();};func NewMatrixFromTransforms (xScale ,yScale ,theta ,tx ,ty float64 )Matrix {return IdentityMatrix ().Scale (xScale ,yScale ).Rotate (theta ).Translate (tx ,ty );};const _dd =1.0e-6;func (_add Matrix )Unrealistic ()bool {_ae ,_abf ,_gbc ,_gcb :=_d .Abs (_add [0]),_d .Abs (_add [1]),_d .Abs (_add [3]),_d .Abs (_add [4]);
_fde :=_ae > _fdc &&_gcb > _fdc ;_acd :=_abf > _fdc &&_gbc > _fdc ;return !(_fde ||_acd );};func (_caf *Point )Set (x ,y float64 ){_caf .X ,_caf .Y =x ,y };func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};func (_aaa *Point )transformByMatrix (_fe Matrix ){_aaa .X ,_aaa .Y =_fe .Transform (_aaa .X ,_aaa .Y )};
type Point struct{X float64 ;Y float64 ;};const _fdc =1e-6;func (_aab Point )Distance (b Point )float64 {return _d .Hypot (_aab .X -b .X ,_aab .Y -b .Y )};func (_cb Matrix )Mult (b Matrix )Matrix {_cb .Concat (b );return _cb };func (_agbc Point )Interpolate (b Point ,t float64 )Point {return Point {X :(1-t )*_agbc .X +t *b .X ,Y :(1-t )*_agbc .Y +t *b .Y };
};func (_de Matrix )ScalingFactorY ()float64 {return _d .Hypot (_de [3],_de [4])};func (_bgc *Matrix )Set (a ,b ,c ,d ,tx ,ty float64 ){_bgc [0],_bgc [1]=a ,b ;_bgc [3],_bgc [4]=c ,d ;_bgc [6],_bgc [7]=tx ,ty ;_bgc .clampRange ();};func (_acg Point )Displace (delta Point )Point {return Point {_acg .X +delta .X ,_acg .Y +delta .Y }};
type Matrix [9]float64 ;func (_gb *Matrix )clampRange (){for _agb ,_bdb :=range _gb {if _bdb > _efc {_b .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_bdb ,_efc );_gb [_agb ]=_efc ;}else if _bdb < -_efc {_b .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_bdb ,-_efc );
_gb [_agb ]=-_efc ;};};};func (_bd Matrix )Translation ()(float64 ,float64 ){return _bd [6],_bd [7]};