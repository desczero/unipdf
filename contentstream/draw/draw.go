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

// Package draw has handy features for defining paths which can be used to draw content on a PDF page.  Handles
// defining paths as points, vector calculations and conversion to PDF content stream data which can be used in
// page content streams and XObject forms and thus also in annotation appearance streams.
//
// Also defines utility functions for drawing common shapes such as rectangles, lines and circles (ovals).
package draw ;import (_c "fmt";_e "github.com/unidoc/unipdf/v3/contentstream";_da "github.com/unidoc/unipdf/v3/core";_a "github.com/unidoc/unipdf/v3/internal/transform";_b "github.com/unidoc/unipdf/v3/model";_d "math";);

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct{Points [][]Point ;FillEnabled bool ;FillColor _b .PdfColor ;BorderEnabled bool ;BorderColor _b .PdfColor ;BorderWidth float64 ;};

// AppendCurve appends the specified Bezier curve to the path.
func (_afd CubicBezierPath )AppendCurve (curve CubicBezierCurve )CubicBezierPath {_afd .Curves =append (_afd .Curves ,curve );return _afd ;};

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_fde Vector )GetPolarAngle ()float64 {return _d .Atan2 (_fde .Dy ,_fde .Dx )};

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct{P0 Point ;P1 Point ;P2 Point ;P3 Point ;};

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_aab BasicLine )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_bc :=NewPath ();_bc =_bc .AppendPoint (NewPoint (_aab .X1 ,_aab .Y1 ));_bc =_bc .AppendPoint (NewPoint (_aab .X2 ,_aab .Y2 ));_cgg :=_e .NewContentCreator ();_cgg .Add_q ().Add_w (_aab .LineWidth ).SetStrokingColor (_aab .LineColor );
if _aab .LineStyle ==LineStyleDashed {if _aab .DashArray ==nil {_aab .DashArray =[]int64 {1,1};};_cgg .Add_d (_aab .DashArray ,_aab .DashPhase );};if len (gsName )> 1{_cgg .Add_gs (_da .PdfObjectName (gsName ));};DrawPathWithCreator (_bc ,_cgg );_cgg .Add_S ().Add_Q ();
return _cgg .Bytes (),_bc .GetBoundingBox ().ToPdfRectangle (),nil ;};

// AppendPoint adds the specified point to the path.
func (_ge Path )AppendPoint (point Point )Path {_ge .Points =append (_ge .Points ,point );return _ge };

// Draw draws the composite curve polygon. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array
// and the bounding box of the polygon.
func (_fcd CurvePolygon )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_egc :=_e .NewContentCreator ();_egc .Add_q ();_fcd .FillEnabled =_fcd .FillEnabled &&_fcd .FillColor !=nil ;if _fcd .FillEnabled {_egc .SetNonStrokingColor (_fcd .FillColor );
};_fcd .BorderEnabled =_fcd .BorderEnabled &&_fcd .BorderColor !=nil ;if _fcd .BorderEnabled {_egc .SetStrokingColor (_fcd .BorderColor );_egc .Add_w (_fcd .BorderWidth );};if len (gsName )> 1{_egc .Add_gs (_da .PdfObjectName (gsName ));};_fef :=NewCubicBezierPath ();
for _ ,_bag :=range _fcd .Rings {for _egd ,_cbgb :=range _bag {if _egd ==0{_egc .Add_m (_cbgb .P0 .X ,_cbgb .P0 .Y );}else {_egc .Add_l (_cbgb .P0 .X ,_cbgb .P0 .Y );};_egc .Add_c (_cbgb .P1 .X ,_cbgb .P1 .Y ,_cbgb .P2 .X ,_cbgb .P2 .Y ,_cbgb .P3 .X ,_cbgb .P3 .Y );
_fef =_fef .AppendCurve (_cbgb );};_egc .Add_h ();};if _fcd .FillEnabled &&_fcd .BorderEnabled {_egc .Add_B ();}else if _fcd .FillEnabled {_egc .Add_f ();}else if _fcd .BorderEnabled {_egc .Add_S ();};_egc .Add_Q ();return _egc .Bytes (),_fef .GetBoundingBox ().ToPdfRectangle (),nil ;
};

// FlipY flips the sign of the Dy component of the vector.
func (_cdea Vector )FlipY ()Vector {_cdea .Dy =-_cdea .Dy ;return _cdea };

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{Points []Point ;};

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_gdg Circle )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_ab :=_gdg .Width /2;_gb :=_gdg .Height /2;if _gdg .BorderEnabled {_ab -=_gdg .BorderWidth /2;_gb -=_gdg .BorderWidth /2;};_eb :=0.551784;_cb :=_ab *_eb ;_fb :=_gb *_eb ;_de :=NewCubicBezierPath ();
_de =_de .AppendCurve (NewCubicBezierCurve (-_ab ,0,-_ab ,_fb ,-_cb ,_gb ,0,_gb ));_de =_de .AppendCurve (NewCubicBezierCurve (0,_gb ,_cb ,_gb ,_ab ,_fb ,_ab ,0));_de =_de .AppendCurve (NewCubicBezierCurve (_ab ,0,_ab ,-_fb ,_cb ,-_gb ,0,-_gb ));_de =_de .AppendCurve (NewCubicBezierCurve (0,-_gb ,-_cb ,-_gb ,-_ab ,-_fb ,-_ab ,0));
_de =_de .Offset (_ab ,_gb );if _gdg .BorderEnabled {_de =_de .Offset (_gdg .BorderWidth /2,_gdg .BorderWidth /2);};if _gdg .X !=0||_gdg .Y !=0{_de =_de .Offset (_gdg .X ,_gdg .Y );};_aeb :=_e .NewContentCreator ();_aeb .Add_q ();if _gdg .FillEnabled {_aeb .SetNonStrokingColor (_gdg .FillColor );
};if _gdg .BorderEnabled {_aeb .SetStrokingColor (_gdg .BorderColor );_aeb .Add_w (_gdg .BorderWidth );};if len (gsName )> 1{_aeb .Add_gs (_da .PdfObjectName (gsName ));};DrawBezierPathWithCreator (_de ,_aeb );_aeb .Add_h ();if _gdg .FillEnabled &&_gdg .BorderEnabled {_aeb .Add_B ();
}else if _gdg .FillEnabled {_aeb .Add_f ();}else if _gdg .BorderEnabled {_aeb .Add_S ();};_aeb .Add_Q ();_deg :=_de .GetBoundingBox ();if _gdg .BorderEnabled {_deg .Height +=_gdg .BorderWidth ;_deg .Width +=_gdg .BorderWidth ;_deg .X -=_gdg .BorderWidth /2;
_deg .Y -=_gdg .BorderWidth /2;};return _aeb .Bytes (),_deg .ToPdfRectangle (),nil ;};

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor _b .PdfColor ;Opacity float64 ;LineWidth float64 ;LineEndingStyle1 LineEndingStyle ;LineEndingStyle2 LineEndingStyle ;LineStyle LineStyle ;};

// Flip changes the sign of the vector: -vector.
func (_gece Vector )Flip ()Vector {_dadf :=_gece .Magnitude ();_dfa :=_gece .GetPolarAngle ();_gece .Dx =_dadf *_d .Cos (_dfa +_d .Pi );_gece .Dy =_dadf *_d .Sin (_dfa +_d .Pi );return _gece ;};

// Copy returns a clone of the Bezier path.
func (_afg CubicBezierPath )Copy ()CubicBezierPath {_bd :=CubicBezierPath {};_bd .Curves =append (_bd .Curves ,_afg .Curves ...);return _bd ;};

// Offset shifts the Bezier path with the specified offsets.
func (_dfd CubicBezierPath )Offset (offX ,offY float64 )CubicBezierPath {for _eg ,_ff :=range _dfd .Curves {_dfd .Curves [_eg ]=_ff .AddOffsetXY (offX ,offY );};return _dfd ;};

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_bec Polyline )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){if _bec .LineColor ==nil {_bec .LineColor =_b .NewPdfColorDeviceRGB (0,0,0);};_bcc :=NewPath ();for _ ,_cde :=range _bec .Points {_bcc =_bcc .AppendPoint (_cde );};_becg :=_e .NewContentCreator ();
_becg .Add_q ().SetStrokingColor (_bec .LineColor ).Add_w (_bec .LineWidth );if len (gsName )> 1{_becg .Add_gs (_da .PdfObjectName (gsName ));};DrawPathWithCreator (_bcc ,_becg );_becg .Add_S ();_becg .Add_Q ();return _becg .Bytes (),_bcc .GetBoundingBox ().ToPdfRectangle (),nil ;
};

// NewPoint returns a new point with the coordinates x, y.
func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};

// GetBounds returns the bounding box of the Bezier curve.
func (_ba CubicBezierCurve )GetBounds ()_b .PdfRectangle {_cd :=_ba .P0 .X ;_ad :=_ba .P0 .X ;_fe :=_ba .P0 .Y ;_ce :=_ba .P0 .Y ;for _af :=0.0;_af <=1.0;_af +=0.001{Rx :=_ba .P0 .X *_d .Pow (1-_af ,3)+_ba .P1 .X *3*_af *_d .Pow (1-_af ,2)+_ba .P2 .X *3*_d .Pow (_af ,2)*(1-_af )+_ba .P3 .X *_d .Pow (_af ,3);
Ry :=_ba .P0 .Y *_d .Pow (1-_af ,3)+_ba .P1 .Y *3*_af *_d .Pow (1-_af ,2)+_ba .P2 .Y *3*_d .Pow (_af ,2)*(1-_af )+_ba .P3 .Y *_d .Pow (_af ,3);if Rx < _cd {_cd =Rx ;};if Rx > _ad {_ad =Rx ;};if Ry < _fe {_fe =Ry ;};if Ry > _ce {_ce =Ry ;};};_df :=_b .PdfRectangle {};
_df .Llx =_cd ;_df .Lly =_fe ;_df .Urx =_ad ;_df .Ury =_ce ;return _df ;};

// AddVector adds vector to a point.
func (_edc Point )AddVector (v Vector )Point {_edc .X +=v .Dx ;_edc .Y +=v .Dy ;return _edc };

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_ae Point )Add (dx ,dy float64 )Point {_ae .X +=dx ;_ae .Y +=dy ;return _ae };

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_adbg Line )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_cg ,_cac :=_adbg .X1 ,_adbg .X2 ;_dd ,_fecg :=_adbg .Y1 ,_adbg .Y2 ;_gda :=_fecg -_dd ;_gbd :=_cac -_cg ;_ec :=_d .Atan2 (_gda ,_gbd );L :=_d .Sqrt (_d .Pow (_gbd ,2.0)+_d .Pow (_gda ,2.0));
_fgb :=_adbg .LineWidth ;_eff :=_d .Pi ;_ceb :=1.0;if _gbd < 0{_ceb *=-1.0;};if _gda < 0{_ceb *=-1.0;};VsX :=_ceb *(-_fgb /2*_d .Cos (_ec +_eff /2));VsY :=_ceb *(-_fgb /2*_d .Sin (_ec +_eff /2)+_fgb *_d .Sin (_ec +_eff /2));V1X :=VsX +_fgb /2*_d .Cos (_ec +_eff /2);
V1Y :=VsY +_fgb /2*_d .Sin (_ec +_eff /2);V2X :=VsX +_fgb /2*_d .Cos (_ec +_eff /2)+L *_d .Cos (_ec );V2Y :=VsY +_fgb /2*_d .Sin (_ec +_eff /2)+L *_d .Sin (_ec );V3X :=VsX +_fgb /2*_d .Cos (_ec +_eff /2)+L *_d .Cos (_ec )+_fgb *_d .Cos (_ec -_eff /2);V3Y :=VsY +_fgb /2*_d .Sin (_ec +_eff /2)+L *_d .Sin (_ec )+_fgb *_d .Sin (_ec -_eff /2);
V4X :=VsX +_fgb /2*_d .Cos (_ec -_eff /2);V4Y :=VsY +_fgb /2*_d .Sin (_ec -_eff /2);_fbg :=NewPath ();_fbg =_fbg .AppendPoint (NewPoint (V1X ,V1Y ));_fbg =_fbg .AppendPoint (NewPoint (V2X ,V2Y ));_fbg =_fbg .AppendPoint (NewPoint (V3X ,V3Y ));_fbg =_fbg .AppendPoint (NewPoint (V4X ,V4Y ));
_baeg :=_adbg .LineEndingStyle1 ;_adg :=_adbg .LineEndingStyle2 ;_cc :=3*_fgb ;_acb :=3*_fgb ;_ddg :=(_acb -_fgb )/2;if _adg ==LineEndingStyleArrow {_cdc :=_fbg .GetPointNumber (2);_fab :=NewVectorPolar (_cc ,_ec +_eff );_cgf :=_cdc .AddVector (_fab );
_fga :=NewVectorPolar (_acb /2,_ec +_eff /2);_efc :=NewVectorPolar (_cc ,_ec );_bbaa :=NewVectorPolar (_ddg ,_ec +_eff /2);_egf :=_cgf .AddVector (_bbaa );_bdf :=_efc .Add (_fga .Flip ());_dfg :=_egf .AddVector (_bdf );_fdg :=_fga .Scale (2).Flip ().Add (_bdf .Flip ());
_gad :=_dfg .AddVector (_fdg );_bge :=_cgf .AddVector (NewVectorPolar (_fgb ,_ec -_eff /2));_cbgg :=NewPath ();_cbgg =_cbgg .AppendPoint (_fbg .GetPointNumber (1));_cbgg =_cbgg .AppendPoint (_cgf );_cbgg =_cbgg .AppendPoint (_egf );_cbgg =_cbgg .AppendPoint (_dfg );
_cbgg =_cbgg .AppendPoint (_gad );_cbgg =_cbgg .AppendPoint (_bge );_cbgg =_cbgg .AppendPoint (_fbg .GetPointNumber (4));_fbg =_cbgg ;};if _baeg ==LineEndingStyleArrow {_ea :=_fbg .GetPointNumber (1);_bac :=_fbg .GetPointNumber (_fbg .Length ());_ggge :=NewVectorPolar (_fgb /2,_ec +_eff +_eff /2);
_cbe :=_ea .AddVector (_ggge );_ggb :=NewVectorPolar (_cc ,_ec ).Add (NewVectorPolar (_acb /2,_ec +_eff /2));_gcb :=_cbe .AddVector (_ggb );_fage :=NewVectorPolar (_ddg ,_ec -_eff /2);_agdd :=_gcb .AddVector (_fage );_agb :=NewVectorPolar (_cc ,_ec );_cbb :=_bac .AddVector (_agb );
_abd :=NewVectorPolar (_ddg ,_ec +_eff +_eff /2);_dfdc :=_cbb .AddVector (_abd );_dda :=_cbe ;_ege :=NewPath ();_ege =_ege .AppendPoint (_cbe );_ege =_ege .AppendPoint (_gcb );_ege =_ege .AppendPoint (_agdd );for _ ,_eae :=range _fbg .Points [1:len (_fbg .Points )-1]{_ege =_ege .AppendPoint (_eae );
};_ege =_ege .AppendPoint (_cbb );_ege =_ege .AppendPoint (_dfdc );_ege =_ege .AppendPoint (_dda );_fbg =_ege ;};_ee :=_e .NewContentCreator ();_ee .Add_q ().SetNonStrokingColor (_adbg .LineColor );if len (gsName )> 1{_ee .Add_gs (_da .PdfObjectName (gsName ));
};_fbg =_fbg .Offset (_adbg .X1 ,_adbg .Y1 );_abg :=_fbg .GetBoundingBox ();DrawPathWithCreator (_fbg ,_ee );if _adbg .LineStyle ==LineStyleDashed {_ee .Add_d ([]int64 {1,1},0).Add_S ().Add_f ().Add_Q ();}else {_ee .Add_f ().Add_Q ();};return _ee .Bytes (),_abg .ToPdfRectangle (),nil ;
};

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct{

// Position and size properties.
X float64 ;Y float64 ;Width float64 ;Height float64 ;

// Fill properties.
FillEnabled bool ;FillColor _b .PdfColor ;

// Border properties.
BorderEnabled bool ;BorderColor _b .PdfColor ;BorderWidth float64 ;BorderRadiusTopLeft float64 ;BorderRadiusTopRight float64 ;BorderRadiusBottomLeft float64 ;BorderRadiusBottomRight float64 ;

// Shape opacity (0-1 interval).
Opacity float64 ;};

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_bdc Path )GetPointNumber (number int )Point {if number < 1||number > len (_bdc .Points ){return Point {};};return _bdc .Points [number -1];};

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct{Points []Point ;LineColor _b .PdfColor ;LineWidth float64 ;};const (LineEndingStyleNone LineEndingStyle =0;LineEndingStyleArrow LineEndingStyle =1;LineEndingStyleButt LineEndingStyle =2;);

// Offset shifts the path with the specified offsets.
func (_gec Path )Offset (offX ,offY float64 )Path {for _be ,_gc :=range _gec .Points {_gec .Points [_be ]=_gc .Add (offX ,offY );};return _gec ;};

// Magnitude returns the magnitude of the vector.
func (_dg Vector )Magnitude ()float64 {return _d .Sqrt (_d .Pow (_dg .Dx ,2.0)+_d .Pow (_dg .Dy ,2.0))};const (LineStyleSolid LineStyle =0;LineStyleDashed LineStyle =1;);

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_g CubicBezierCurve )AddOffsetXY (offX ,offY float64 )CubicBezierCurve {_g .P0 .X +=offX ;_g .P1 .X +=offX ;_g .P2 .X +=offX ;_g .P3 .X +=offX ;_g .P0 .Y +=offY ;_g .P1 .Y +=offY ;_g .P2 .Y +=offY ;_g .P3 .Y +=offY ;return _g ;};

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor _b .PdfColor ;BorderEnabled bool ;BorderWidth float64 ;BorderColor _b .PdfColor ;Opacity float64 ;};

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_abf Polygon )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_bbf :=_e .NewContentCreator ();_bbf .Add_q ();_abf .FillEnabled =_abf .FillEnabled &&_abf .FillColor !=nil ;if _abf .FillEnabled {_bbf .SetNonStrokingColor (_abf .FillColor );
};_abf .BorderEnabled =_abf .BorderEnabled &&_abf .BorderColor !=nil ;if _abf .BorderEnabled {_bbf .SetStrokingColor (_abf .BorderColor );_bbf .Add_w (_abf .BorderWidth );};if len (gsName )> 1{_bbf .Add_gs (_da .PdfObjectName (gsName ));};_dad :=NewPath ();
for _ ,_ag :=range _abf .Points {for _cf ,_gcf :=range _ag {_dad =_dad .AppendPoint (_gcf );if _cf ==0{_bbf .Add_m (_gcf .X ,_gcf .Y );}else {_bbf .Add_l (_gcf .X ,_gcf .Y );};};_bbf .Add_h ();};if _abf .FillEnabled &&_abf .BorderEnabled {_bbf .Add_B ();
}else if _abf .FillEnabled {_bbf .Add_f ();}else if _abf .BorderEnabled {_bbf .Add_S ();};_bbf .Add_Q ();return _bbf .Bytes (),_dad .GetBoundingBox ().ToPdfRectangle (),nil ;};

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator (path Path ,creator *_e .ContentCreator ){for _cee ,_ebf :=range path .Points {if _cee ==0{creator .Add_m (_ebf .X ,_ebf .Y );}else {creator .Add_l (_ebf .X ,_ebf .Y );};};};

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int ;

// GetBoundingBox returns the bounding box of the Bezier path.
func (_gd CubicBezierPath )GetBoundingBox ()Rectangle {_fd :=Rectangle {};_bg :=0.0;_gg :=0.0;_ega :=0.0;_adb :=0.0;for _bdg ,_db :=range _gd .Curves {_dbb :=_db .GetBounds ();if _bdg ==0{_bg =_dbb .Llx ;_gg =_dbb .Urx ;_ega =_dbb .Lly ;_adb =_dbb .Ury ;
continue ;};if _dbb .Llx < _bg {_bg =_dbb .Llx ;};if _dbb .Urx > _gg {_gg =_dbb .Urx ;};if _dbb .Lly < _ega {_ega =_dbb .Lly ;};if _dbb .Ury > _adb {_adb =_dbb .Ury ;};};_fd .X =_bg ;_fd .Y =_ega ;_fd .Width =_gg -_bg ;_fd .Height =_adb -_ega ;return _fd ;
};

// Copy returns a clone of the path.
func (_ca Path )Copy ()Path {_ed :=Path {};_ed .Points =append (_ed .Points ,_ca .Points ...);return _ed ;};

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_egae Path )RemovePoint (number int )Path {if number < 1||number > len (_egae .Points ){return _egae ;};_fg :=number -1;_egae .Points =append (_egae .Points [:_fg ],_egae .Points [_fg +1:]...);return _egae ;};func (_bae Point )String ()string {return _c .Sprintf ("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029",_bae .X ,_bae .Y );
};

// Vector represents a two-dimensional vector.
type Vector struct{Dx float64 ;Dy float64 ;};

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_fbe Rectangle )ToPdfRectangle ()*_b .PdfRectangle {return &_b .PdfRectangle {Llx :_fbe .X ,Lly :_fbe .Y ,Urx :_fbe .X +_fbe .Width ,Ury :_fbe .Y +_fbe .Height };};

// CurvePolygon is a multi-point shape with rings containing curves that can be
// drawn to a PDF content stream.
type CurvePolygon struct{Rings [][]CubicBezierCurve ;FillEnabled bool ;FillColor _b .PdfColor ;BorderEnabled bool ;BorderColor _b .PdfColor ;BorderWidth float64 ;};

// Length returns the number of points in the path.
func (_dc Path )Length ()int {return len (_dc .Points )};

// Point represents a two-dimensional point.
type Point struct{X float64 ;Y float64 ;};

// GetBoundingBox returns the bounding box of the path.
func (_ced Path )GetBoundingBox ()BoundingBox {_adbb :=BoundingBox {};_fc :=0.0;_fgf :=0.0;_gdc :=0.0;_aca :=0.0;for _fa ,_gea :=range _ced .Points {if _fa ==0{_fc =_gea .X ;_fgf =_gea .X ;_gdc =_gea .Y ;_aca =_gea .Y ;continue ;};if _gea .X < _fc {_fc =_gea .X ;
};if _gea .X > _fgf {_fgf =_gea .X ;};if _gea .Y < _gdc {_gdc =_gea .Y ;};if _gea .Y > _aca {_aca =_gea .Y ;};};_adbb .X =_fc ;_adbb .Y =_gdc ;_adbb .Width =_fgf -_fc ;_adbb .Height =_aca -_gdc ;return _adbb ;};

// LineStyle refers to how the line will be created.
type LineStyle int ;

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_baf Point )Rotate (theta float64 )Point {_bad :=_a .NewPoint (_baf .X ,_baf .Y ).Rotate (theta );return NewPoint (_bad .X ,_bad .Y );};

// Add adds the specified vector to the current one and returns the result.
func (_bfb Vector )Add (other Vector )Vector {_bfb .Dx +=other .Dx ;_bfb .Dy +=other .Dy ;return _bfb };

// Rotate rotates the vector by the specified angle.
func (_feg Vector )Rotate (phi float64 )Vector {_gcfg :=_feg .Magnitude ();_cfa :=_feg .GetPolarAngle ();return NewVectorPolar (_gcfg ,_cfa +phi );};

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_aa PolyBezierCurve )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){if _aa .BorderColor ==nil {_aa .BorderColor =_b .NewPdfColorDeviceRGB (0,0,0);};_bba :=NewCubicBezierPath ();for _ ,_fbb :=range _aa .Curves {_bba =_bba .AppendCurve (_fbb );
};_gf :=_e .NewContentCreator ();_gf .Add_q ();_aa .FillEnabled =_aa .FillEnabled &&_aa .FillColor !=nil ;if _aa .FillEnabled {_gf .SetNonStrokingColor (_aa .FillColor );};_gf .SetStrokingColor (_aa .BorderColor );_gf .Add_w (_aa .BorderWidth );if len (gsName )> 1{_gf .Add_gs (_da .PdfObjectName (gsName ));
};for _cbg ,_abe :=range _bba .Curves {if _cbg ==0{_gf .Add_m (_abe .P0 .X ,_abe .P0 .Y );}else {_gf .Add_l (_abe .P0 .X ,_abe .P0 .Y );};_gf .Add_c (_abe .P1 .X ,_abe .P1 .Y ,_abe .P2 .X ,_abe .P2 .Y ,_abe .P3 .X ,_abe .P3 .Y );};if _aa .FillEnabled {_gf .Add_h ();
_gf .Add_B ();}else {_gf .Add_S ();};_gf .Add_Q ();return _gf .Bytes (),_bba .GetBoundingBox ().ToPdfRectangle (),nil ;};

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve (x0 ,y0 ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 )CubicBezierCurve {_ac :=CubicBezierCurve {};_ac .P0 =NewPoint (x0 ,y0 );_ac .P1 =NewPoint (x1 ,y1 );_ac .P2 =NewPoint (x2 ,y2 );_ac .P3 =NewPoint (x3 ,y3 );return _ac ;};

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween (a Point ,b Point )Vector {_ccb :=Vector {};_ccb .Dx =b .X -a .X ;_ccb .Dy =b .Y -a .Y ;return _ccb ;};

// FlipX flips the sign of the Dx component of the vector.
func (_gag Vector )FlipX ()Vector {_gag .Dx =-_gag .Dx ;return _gag };

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;};

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor _b .PdfColor ;Opacity float64 ;LineWidth float64 ;LineStyle LineStyle ;DashArray []int64 ;DashPhase int64 ;};

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector (dx ,dy float64 )Vector {_ddag :=Vector {};_ddag .Dx =dx ;_ddag .Dy =dy ;return _ddag };

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath ()CubicBezierPath {_bb :=CubicBezierPath {};_bb .Curves =[]CubicBezierCurve {};return _bb ;};

// NewPath returns a new empty path.
func NewPath ()Path {return Path {}};

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{Curves []CubicBezierCurve ;};

// Scale scales the vector by the specified factor.
func (_geaa Vector )Scale (factor float64 )Vector {_ged :=_geaa .Magnitude ();_acf :=_geaa .GetPolarAngle ();_geaa .Dx =factor *_ged *_d .Cos (_acf );_geaa .Dy =factor *_ged *_d .Sin (_acf );return _geaa ;};

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar (length float64 ,theta float64 )Vector {_ecf :=Vector {};_ecf .Dx =length *_d .Cos (theta );_ecf .Dy =length *_d .Sin (theta );return _ecf ;};

// Draw draws the rectangle. A graphics state can be specified for
// setting additional properties (e.g. opacity). Otherwise pass an empty string
// for the `gsName` parameter. The method returns the content stream as a byte
// array and the bounding box of the shape.
func (_bf Rectangle )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_ggg :=_e .NewContentCreator ();_ggg .Add_q ();if _bf .FillEnabled {_ggg .SetNonStrokingColor (_bf .FillColor );};if _bf .BorderEnabled {_ggg .SetStrokingColor (_bf .BorderColor );
_ggg .Add_w (_bf .BorderWidth );};if len (gsName )> 1{_ggg .Add_gs (_da .PdfObjectName (gsName ));};var (_aef ,_cedc =_bf .X ,_bf .Y ;_fec ,_ffc =_bf .Width ,_bf .Height ;_agd =_d .Abs (_bf .BorderRadiusTopLeft );_bbe =_d .Abs (_bf .BorderRadiusTopRight );
_dbc =_d .Abs (_bf .BorderRadiusBottomLeft );_cff =_d .Abs (_bf .BorderRadiusBottomRight );_dba =0.4477;);_fcb :=Path {Points :[]Point {{X :_aef +_fec -_cff ,Y :_cedc },{X :_aef +_fec ,Y :_cedc +_ffc -_bbe },{X :_aef +_agd ,Y :_cedc +_ffc },{X :_aef ,Y :_cedc +_dbc }}};
_ef :=[][7]float64 {{_cff ,_aef +_fec -_cff *_dba ,_cedc ,_aef +_fec ,_cedc +_cff *_dba ,_aef +_fec ,_cedc +_cff },{_bbe ,_aef +_fec ,_cedc +_ffc -_bbe *_dba ,_aef +_fec -_bbe *_dba ,_cedc +_ffc ,_aef +_fec -_bbe ,_cedc +_ffc },{_agd ,_aef +_agd *_dba ,_cedc +_ffc ,_aef ,_cedc +_ffc -_agd *_dba ,_aef ,_cedc +_ffc -_agd },{_dbc ,_aef ,_cedc +_dbc *_dba ,_aef +_dbc *_dba ,_cedc ,_aef +_dbc ,_cedc }};
_ggg .Add_m (_aef +_dbc ,_cedc );for _fag :=0;_fag < 4;_fag ++{_ga :=_fcb .Points [_fag ];_ggg .Add_l (_ga .X ,_ga .Y );_egdg :=_ef [_fag ];if _gfd :=_egdg [0];_gfd !=0{_ggg .Add_c (_egdg [1],_egdg [2],_egdg [3],_egdg [4],_egdg [5],_egdg [6]);};};_ggg .Add_h ();
if _bf .FillEnabled &&_bf .BorderEnabled {_ggg .Add_B ();}else if _bf .FillEnabled {_ggg .Add_f ();}else if _bf .BorderEnabled {_ggg .Add_S ();};_ggg .Add_Q ();return _ggg .Bytes (),_fcb .GetBoundingBox ().ToPdfRectangle (),nil ;};

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_bef BoundingBox )ToPdfRectangle ()*_b .PdfRectangle {return &_b .PdfRectangle {Llx :_bef .X ,Lly :_bef .Y ,Urx :_bef .X +_bef .Width ,Ury :_bef .Y +_bef .Height };};

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator (bpath CubicBezierPath ,creator *_e .ContentCreator ){for _aaf ,_caf :=range bpath .Curves {if _aaf ==0{creator .Add_m (_caf .P0 .X ,_caf .P0 .Y );};creator .Add_c (_caf .P1 .X ,_caf .P1 .Y ,_caf .P2 .X ,_caf .P2 .Y ,_caf .P3 .X ,_caf .P3 .Y );
};};

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct{Curves []CubicBezierCurve ;BorderWidth float64 ;BorderColor _b .PdfColor ;FillEnabled bool ;FillColor _b .PdfColor ;};