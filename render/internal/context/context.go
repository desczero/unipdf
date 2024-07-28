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

package context ;import (_gd "errors";_gee "github.com/unidoc/freetype/truetype";_eac "github.com/unidoc/unipdf/v3/core";_ge "github.com/unidoc/unipdf/v3/internal/cmap";_d "github.com/unidoc/unipdf/v3/internal/textencoding";_gb "github.com/unidoc/unipdf/v3/internal/transform";
_ea "github.com/unidoc/unipdf/v3/model";_e "golang.org/x/image/font";_fge "image";_f "image/color";_fg "strconv";_b "strings";);func (_dfaf *TextFont )charcodeToRunesSimple (_aad _d .CharCode )(_d .CharCode ,[]rune ){_fbe :=[]_d .CharCode {_aad };if _dfaf .Font .IsSimple ()&&_dfaf ._dcg !=nil {if _fbf :=_dfaf ._dcg .Index (rune (_aad ));
_fbf > 0{return _aad ,[]rune {rune (_aad )};};};if _dfaf ._dcg !=nil &&!_dfaf ._dcg .HasCmap ()&&_b .Contains (_dfaf .Font .Encoder ().String (),"\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-"){if _bbf :=_dfaf ._dcg .Index (rune (_aad ));_bbf > 0{return _aad ,[]rune {rune (_aad )};
};};return _aad ,_dfaf .Font .CharcodesToUnicode (_fbe );};type TextRenderingMode int ;type LineCap int ;func (_fegb *TextState )ProcTd (tx ,ty float64 ){_fegb .Tlm .Concat (_gb .TranslationMatrix (tx ,ty ));_fegb .Tm =_fegb .Tlm .Clone ();};func NewTextFontFromPath (filePath string ,size float64 )(*TextFont ,error ){_fde ,_gc :=_ea .NewPdfFontFromTTFFile (filePath );
if _gc !=nil {return nil ,_gc ;};return NewTextFont (_fde ,size );};type Pattern interface{ColorAt (_gf ,_ed int )_f .Color ;};func (_gac *TextState )ProcTj (data []byte ,ctx Context ){_gfb :=_gac .Tf .Size ;_cfdf :=_gac .Th /100.0;_cgf :=_gac .GlobalScale ;
_fgfd :=_gb .NewMatrix (_gfb *_cfdf ,0,0,_gfb ,0,_gac .Ts );_fdg :=ctx .Matrix ();_ccf :=_fdg .Clone ().Mult (_gac .Tm .Clone ().Mult (_fgfd )).ScalingFactorY ();_eag :=_gac .Tf .NewFace (_ccf );_bfc :=_gac .Tf .BytesToCharcodes (data );for _ ,_cd :=range _bfc {_adbf ,_cde :=_gac .Tf .CharcodeToRunes (_cd );
_ae :=string (_cde );if _ae =="\u0000"{continue ;};_bag :=_fdg .Clone ().Mult (_gac .Tm .Clone ().Mult (_fgfd ));_gad :=_bag .ScalingFactorY ();_bag =_bag .Scale (1/_gad ,-1/_gad );if _gac .Tr !=TextRenderingModeInvisible {ctx .SetMatrix (_bag );ctx .DrawString (_ae ,_eag ,0,0);
ctx .SetMatrix (_fdg );};_gdd :=0.0;if _ae =="\u0020"{_gdd =_gac .Tw ;};_eeb ,_ ,_adg :=_gac .Tf .GetCharMetrics (_adbf );if _adg {_eeb =_eeb *0.001*_gfb ;}else {_eeb ,_ =ctx .MeasureString (_ae ,_eag );_eeb =_eeb /_cgf ;};_fee :=(_eeb +_gac .Tc +_gdd )*_cfdf ;
_gac .Tm =_gac .Tm .Mult (_gb .TranslationMatrix (_fee ,0));};};func (_fdd *TextState )ProcTStar (){_fdd .ProcTd (0,-_fdd .Tl )};func (_defe *TextFont )WithSize (size float64 ,originalFont *_ea .PdfFont )*TextFont {return &TextFont {Font :_defe .Font ,Size :size ,_dcg :_defe ._dcg ,_afg :originalFont };
};type Context interface{Push ();Pop ();Matrix ()_gb .Matrix ;SetMatrix (_gec _gb .Matrix );Translate (_dg ,_fe float64 );Scale (_dc ,_eaf float64 );Rotate (_cg float64 );MoveTo (_gef ,_eafc float64 );LineTo (_ce ,_gg float64 );CubicTo (_df ,_a ,_fgf ,_fc ,_da ,_bb float64 );
QuadraticTo (_ga ,_bc ,_db ,_de float64 );NewSubPath ();ClosePath ();ClearPath ();Clip ();ClipPreserve ();ResetClip ();LineWidth ()float64 ;SetLineWidth (_ee float64 );SetLineCap (_def LineCap );SetLineJoin (_cgg LineJoin );SetDash (_ab ...float64 );SetDashOffset (_dd float64 );
Fill ();FillPreserve ();Stroke ();StrokePreserve ();SetRGBA (_ac ,_ad ,_gdf ,_dcb float64 );SetFillRGBA (_bf ,_bff ,_dfa ,_gde float64 );SetFillStyle (_fcg Pattern );SetFillRule (_ba FillRule );SetStrokeRGBA (_aa ,_ca ,_eda ,_fd float64 );SetStrokeStyle (_fed Pattern );
FillPattern ()Pattern ;StrokePattern ()Pattern ;TextState ()*TextState ;DrawString (_fce string ,_fb _e .Face ,_dcbb ,_dag float64 );MeasureString (_abd string ,_dcc _e .Face )(_eg ,_abdf float64 );DrawRectangle (_fef ,_ggf ,_caa ,_cf float64 );DrawImage (_bg _fge .Image ,_acd ,_cb int );
DrawImageAnchored (_bfe _fge .Image ,_af ,_bac int ,_eaff ,_ceg float64 );Height ()int ;Width ()int ;};type TextFont struct{Font *_ea .PdfFont ;Size float64 ;_dcg *_gee .Font ;_afg *_ea .PdfFont ;};type TextState struct{Tc float64 ;Tw float64 ;Th float64 ;
Tl float64 ;Tf *TextFont ;Ts float64 ;Tm _gb .Matrix ;Tlm _gb .Matrix ;Tr TextRenderingMode ;GlobalScale float64 ;};const (FillRuleWinding FillRule =iota ;FillRuleEvenOdd ;);func (_dcbe *TextState )ProcTD (tx ,ty float64 ){_dcbe .Tl =-ty ;_dcbe .ProcTd (tx ,ty )};
func (_cbc *TextState )ProcDQ (data []byte ,aw ,ac float64 ,ctx Context ){_cbc .Tw =aw ;_cbc .Tc =ac ;_cbc .ProcQ (data ,ctx );};func (_bbc *TextState )ProcTm (a ,b ,c ,d ,e ,f float64 ){_bbc .Tm =_gb .NewMatrix (a ,b ,c ,d ,e ,f );_bbc .Tlm =_bbc .Tm .Clone ();
};func (_cdb *TextState )Translate (tx ,ty float64 ){_cdb .Tm =_cdb .Tm .Mult (_gb .TranslationMatrix (tx ,ty ));};func (_eec *TextState )ProcTf (font *TextFont ){_eec .Tf =font };const (LineJoinRound LineJoin =iota ;LineJoinBevel ;);func NewTextFont (font *_ea .PdfFont ,size float64 )(*TextFont ,error ){_baa :=font .FontDescriptor ();
if _baa ==nil {return nil ,_gd .New ("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072");};_eb ,_fdf :=_eac .GetStream (_baa .FontFile2 );if !_fdf {return nil ,_gd .New ("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d");
};_aaf ,_egg :=_eac .DecodeStream (_eb );if _egg !=nil {return nil ,_egg ;};_cec ,_egg :=_gee .Parse (_aaf );if _egg !=nil {return nil ,_egg ;};_dce :=font .FontDescriptor ().FontName .String ();_cfd :=len (_dce )> 7&&_dce [6]=='+';if _baa .Flags !=nil {_ff ,_cfe :=_fg .Atoi (_baa .Flags .String ());
if _cfe ==nil &&_ff ==32{_cfd =false ;};};if !_cec .HasCmap ()&&(!_b .Contains (font .Encoder ().String (),"\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-")||!_cfd ){return nil ,_gd .New ("\u006e\u006f c\u006d\u0061\u0070 \u0061\u006e\u0064\u0020enc\u006fdi\u006e\u0067\u0020\u0069\u0073\u0020\u006eot\u0020\u0069\u0064\u0065\u006e\u0074\u0069t\u0079");
};return &TextFont {Font :font ,Size :size ,_dcg :_cec },nil ;};const (LineCapRound LineCap =iota ;LineCapButt ;LineCapSquare ;);type FillRule int ;func (_fbb *TextFont )CharcodeToRunes (charcode _d .CharCode )(_d .CharCode ,[]rune ){_cc :=[]_d .CharCode {charcode };
if _fbb ._afg ==nil ||_fbb ._afg ==_fbb .Font {return _fbb .charcodeToRunesSimple (charcode );};_ag :=_fbb ._afg .CharcodesToUnicode (_cc );_feg ,_ :=_fbb .Font .RunesToCharcodeBytes (_ag );_bfef :=_fbb .Font .BytesToCharcodes (_feg );_egb :=charcode ;
if len (_bfef )> 0&&_bfef [0]!=0{_egb =_bfef [0];};if string (_ag )==string (_ge .MissingCodeRune )&&_fbb ._afg .BaseFont ()==_fbb .Font .BaseFont (){return _fbb .charcodeToRunesSimple (charcode );};return _egb ,_ag ;};func (_bgd *TextFont )NewFace (size float64 )_e .Face {return _gee .NewFace (_bgd ._dcg ,&_gee .Options {Size :size });
};type LineJoin int ;func (_bcb *TextFont )GetCharMetrics (code _d .CharCode )(float64 ,float64 ,bool ){if _ddf ,_adbe :=_bcb .Font .GetCharMetrics (code );_adbe &&_ddf .Wx !=0{return _ddf .Wx ,_ddf .Wy ,_adbe ;};if _bcb ._afg ==nil {return 0,0,false ;
};_gbg ,_gaa :=_bcb ._afg .GetCharMetrics (code );return _gbg .Wx ,_gbg .Wy ,_gaa &&_gbg .Wx !=0;};type Gradient interface{Pattern ;AddColorStop (_c float64 ,_gba _f .Color );};func (_ef *TextFont )BytesToCharcodes (data []byte )[]_d .CharCode {if _ef ._afg !=nil {return _ef ._afg .BytesToCharcodes (data );
};return _ef .Font .BytesToCharcodes (data );};const (TextRenderingModeFill TextRenderingMode =iota ;TextRenderingModeStroke ;TextRenderingModeFillStroke ;TextRenderingModeInvisible ;TextRenderingModeFillClip ;TextRenderingModeStrokeClip ;TextRenderingModeFillStrokeClip ;
TextRenderingModeClip ;);func (_dcf *TextState )Reset (){_dcf .Tm =_gb .IdentityMatrix ();_dcf .Tlm =_gb .IdentityMatrix ()};func (_aec *TextState )ProcQ (data []byte ,ctx Context ){_aec .ProcTStar ();_aec .ProcTj (data ,ctx )};func NewTextState ()TextState {return TextState {Th :100,Tm :_gb .IdentityMatrix (),Tlm :_gb .IdentityMatrix ()};
};