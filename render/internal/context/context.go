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

package context ;import (_ff "errors";_af "github.com/unidoc/freetype/truetype";_cg "github.com/unidoc/unipdf/v3/core";_a "github.com/unidoc/unipdf/v3/internal/cmap";_g "github.com/unidoc/unipdf/v3/internal/textencoding";_bc "github.com/unidoc/unipdf/v3/internal/transform";
_bd "github.com/unidoc/unipdf/v3/model";_cf "golang.org/x/image/font";_cb "image";_f "image/color";_c "strings";);const (LineCapRound LineCap =iota ;LineCapButt ;LineCapSquare ;);func (_gce *TextFont )BytesToCharcodes (data []byte )[]_g .CharCode {if _gce ._ebe !=nil {return _gce ._ebe .BytesToCharcodes (data );
};return _gce .Font .BytesToCharcodes (data );};func (_gdf *TextState )Reset (){_gdf .Tm =_bc .IdentityMatrix ();_gdf .Tlm =_bc .IdentityMatrix ()};type TextRenderingMode int ;func (_fgad *TextState )ProcDQ (data []byte ,aw ,ac float64 ,ctx Context ){_fgad .Tw =aw ;
_fgad .Tc =ac ;_fgad .ProcQ (data ,ctx );};const (FillRuleWinding FillRule =iota ;FillRuleEvenOdd ;);type TextFont struct{Font *_bd .PdfFont ;Size float64 ;_ada *_af .Font ;_ebe *_bd .PdfFont ;};func (_be *TextFont )GetCharMetrics (code _g .CharCode )(float64 ,float64 ,bool ){if _bgc ,_bbc :=_be .Font .GetCharMetrics (code );
_bbc &&_bgc .Wx !=0{return _bgc .Wx ,_bgc .Wy ,_bbc ;};if _be ._ebe ==nil {return 0,0,false ;};_ab ,_age :=_be ._ebe .GetCharMetrics (code );return _ab .Wx ,_ab .Wy ,_age &&_ab .Wx !=0;};type LineJoin int ;func (_bec *TextState )ProcTj (data []byte ,ctx Context ){_abe :=_bec .Tf .Size ;
_bgg :=_bec .Th /100.0;_gbe :=_bec .GlobalScale ;_afg :=_bc .NewMatrix (_abe *_bgg ,0,0,_abe ,0,_bec .Ts );_cbe :=ctx .Matrix ();_ed :=_cbe .Clone ().Mult (_bec .Tm .Clone ().Mult (_afg )).ScalingFactorY ();_bgd :=_bec .Tf .NewFace (_ed );_cbcf :=_bec .Tf .BytesToCharcodes (data );
for _ ,_bbe :=range _cbcf {_eff ,_eag :=_bec .Tf .CharcodeToRunes (_bbe );_cfbc :=string (_eag );if _cfbc =="\u0000"{continue ;};_afd :=_cbe .Clone ().Mult (_bec .Tm .Clone ().Mult (_afg ));_cagd :=_afd .ScalingFactorY ();_afd =_afd .Scale (1/_cagd ,-1/_cagd );
if _bec .Tr !=TextRenderingModeInvisible {ctx .SetMatrix (_afd );ctx .DrawString (_cfbc ,_bgd ,0,0);ctx .SetMatrix (_cbe );};_cgcd :=0.0;if _cfbc =="\u0020"{_cgcd =_bec .Tw ;};_bgb ,_ ,_bfc :=_bec .Tf .GetCharMetrics (_eff );if _bfc {_bgb =_bgb *0.001*_abe ;
}else {_bgb ,_ =ctx .MeasureString (_cfbc ,_bgd );_bgb =_bgb /_gbe ;};_bddg :=(_bgb +_bec .Tc +_cgcd )*_bgg ;_bec .Tm =_bec .Tm .Mult (_bc .TranslationMatrix (_bddg ,0));};};func (_cde *TextState )ProcTf (font *TextFont ){_cde .Tf =font };type TextState struct{Tc float64 ;
Tw float64 ;Th float64 ;Tl float64 ;Tf *TextFont ;Ts float64 ;Tm _bc .Matrix ;Tlm _bc .Matrix ;Tr TextRenderingMode ;GlobalScale float64 ;};type Context interface{Push ();Pop ();Matrix ()_bc .Matrix ;SetMatrix (_dc _bc .Matrix );Translate (_e ,_fd float64 );
Scale (_dg ,_bdd float64 );Rotate (_ac float64 );MoveTo (_ae ,_dga float64 );LineTo (_ef ,_ga float64 );CubicTo (_ea ,_fb ,_ge ,_eg ,_ee ,_df float64 );QuadraticTo (_aee ,_cfb ,_aa ,_dfd float64 );NewSubPath ();ClosePath ();ClearPath ();Clip ();ClipPreserve ();
ResetClip ();LineWidth ()float64 ;SetLineWidth (_ce float64 );SetLineCap (_gc LineCap );SetLineJoin (_de LineJoin );SetDash (_cc ...float64 );SetDashOffset (_gae float64 );Fill ();FillPreserve ();Stroke ();StrokePreserve ();SetRGBA (_eef ,_aeeg ,_aaa ,_fa float64 );
SetFillRGBA (_ad ,_ca ,_fba ,_gb float64 );SetFillStyle (_gg Pattern );SetFillRule (_ccd FillRule );SetStrokeRGBA (_bda ,_ag ,_cd ,_eb float64 );SetStrokeStyle (_cbc Pattern );FillPattern ()Pattern ;StrokePattern ()Pattern ;TextState ()*TextState ;DrawString (_bdeg string ,_deg _cf .Face ,_db ,_bg float64 );
MeasureString (_fe string ,_ceg _cf .Face )(_gd ,_dee float64 );DrawRectangle (_bca ,_aca ,_eea ,_eeb float64 );DrawImage (_cff _cb .Image ,_aeb ,_gf int );DrawImageAnchored (_dbd _cb .Image ,_deb ,_dbe int ,_afc ,_cgc float64 );Height ()int ;Width ()int ;
};func NewTextState ()TextState {return TextState {Th :100,Tm :_bc .IdentityMatrix (),Tlm :_bc .IdentityMatrix ()};};func (_degc *TextFont )WithSize (size float64 ,originalFont *_bd .PdfFont )*TextFont {return &TextFont {Font :_degc .Font ,Size :size ,_ada :_degc ._ada ,_ebe :originalFont };
};type FillRule int ;func (_bce *TextFont )CharcodeToRunes (charcode _g .CharCode )(_g .CharCode ,[]rune ){_cge :=[]_g .CharCode {charcode };if _bce ._ebe ==nil ||_bce ._ebe ==_bce .Font {return _bce .charcodeToRunesSimple (charcode );};_dge :=_bce ._ebe .CharcodesToUnicode (_cge );
_ega ,_ :=_bce .Font .RunesToCharcodeBytes (_dge );_aaab :=_bce .Font .BytesToCharcodes (_ega );_gec :=charcode ;if len (_aaab )> 0&&_aaab [0]!=0{_gec =_aaab [0];};if string (_dge )==string (_a .MissingCodeRune )&&_bce ._ebe .BaseFont ()==_bce .Font .BaseFont (){return _bce .charcodeToRunesSimple (charcode );
};return _gec ,_dge ;};type Gradient interface{Pattern ;AddColorStop (_bde float64 ,_bb _f .Color );};const (TextRenderingModeFill TextRenderingMode =iota ;TextRenderingModeStroke ;TextRenderingModeFillStroke ;TextRenderingModeInvisible ;TextRenderingModeFillClip ;
TextRenderingModeStrokeClip ;TextRenderingModeFillStrokeClip ;TextRenderingModeClip ;);func (_cbd *TextState )Translate (tx ,ty float64 ){_cbd .Tm =_cbd .Tm .Mult (_bc .TranslationMatrix (tx ,ty ));};type Pattern interface{ColorAt (_d ,_fg int )_f .Color ;
};func (_afce *TextFont )NewFace (size float64 )_cf .Face {return _af .NewFace (_afce ._ada ,&_af .Options {Size :size });};func (_aeeb *TextState )ProcTd (tx ,ty float64 ){_aeeb .Tlm .Concat (_bc .TranslationMatrix (tx ,ty ));_aeeb .Tm =_aeeb .Tlm .Clone ();
};func (_bbcd *TextState )ProcQ (data []byte ,ctx Context ){_bbcd .ProcTStar ();_bbcd .ProcTj (data ,ctx )};type LineCap int ;func (_gfb *TextState )ProcTStar (){_gfb .ProcTd (0,-_gfb .Tl )};func (_cage *TextState )ProcTm (a ,b ,c ,d ,e ,f float64 ){_cage .Tm =_bc .NewMatrix (a ,b ,c ,d ,e ,f );
_cage .Tlm =_cage .Tm .Clone ();};func NewTextFontFromPath (filePath string ,size float64 )(*TextFont ,error ){_cffg ,_fdg :=_bd .NewPdfFontFromTTFFile (filePath );if _fdg !=nil {return nil ,_fdg ;};return NewTextFont (_cffg ,size );};const (LineJoinRound LineJoin =iota ;
LineJoinBevel ;);func (_fca *TextState )ProcTD (tx ,ty float64 ){_fca .Tl =-ty ;_fca .ProcTd (tx ,ty )};func NewTextFont (font *_bd .PdfFont ,size float64 )(*TextFont ,error ){_afb :=font .FontDescriptor ();if _afb ==nil {return nil ,_ff .New ("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072");
};_cag ,_fc :=_cg .GetStream (_afb .FontFile2 );if !_fc {return nil ,_ff .New ("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d");};_ead ,_ced :=_cg .DecodeStream (_cag );
if _ced !=nil {return nil ,_ced ;};_dcc ,_ced :=_af .Parse (_ead );if _ced !=nil {return nil ,_ced ;};_bdab :=font .FontDescriptor ().FontName .String ();_gac :=len (_bdab )> 7&&_bdab [6]=='+';if !_dcc .HasCmap ()&&(!_c .Contains (font .Encoder ().String (),"\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-")||!_gac ){return nil ,_ff .New ("\u006e\u006f c\u006d\u0061\u0070 \u0061\u006e\u0064\u0020enc\u006fdi\u006e\u0067\u0020\u0069\u0073\u0020\u006eot\u0020\u0069\u0064\u0065\u006e\u0074\u0069t\u0079");
};return &TextFont {Font :font ,Size :size ,_ada :_dcc },nil ;};func (_fga *TextFont )charcodeToRunesSimple (_bf _g .CharCode )(_g .CharCode ,[]rune ){_gad :=[]_g .CharCode {_bf };if _fga .Font .IsSimple ()&&_fga ._ada !=nil {if _bff :=_fga ._ada .Index (rune (_bf ));
_bff > 0{return _bf ,[]rune {rune (_bf )};};};if _fga ._ada !=nil &&!_fga ._ada .HasCmap ()&&_c .Contains (_fga .Font .Encoder ().String (),"\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-"){if _da :=_fga ._ada .Index (rune (_bf ));_da > 0{return _bf ,[]rune {rune (_bf )};
};};return _bf ,_fga .Font .CharcodesToUnicode (_gad );};