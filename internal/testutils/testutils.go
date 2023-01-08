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

package testutils ;import (_c "crypto/md5";_b "encoding/hex";_gc "errors";_f "fmt";_bf "github.com/unidoc/unipdf/v3/common";_ae "github.com/unidoc/unipdf/v3/core";_a "image";_cd "image/png";_ec "io";_ge "os";_gb "os/exec";_d "path/filepath";_cf "strings";
_g "testing";);func _fe (_ega _ae .PdfObject ,_dc map[int64 ]_ae .PdfObject )error {switch _fc :=_ega .(type ){case *_ae .PdfIndirectObject :_acg :=_fc ;_fe (_acg .PdfObject ,_dc );case *_ae .PdfObjectDictionary :_afg :=_fc ;for _ ,_ed :=range _afg .Keys (){_cc :=_afg .Get (_ed );
if _bce ,_adf :=_cc .(*_ae .PdfObjectReference );_adf {_ffc ,_fb :=_dc [_bce .ObjectNumber ];if !_fb {return _gc .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_afg .Set (_ed ,_ffc );}else {_fe (_cc ,_dc );};};case *_ae .PdfObjectArray :_ddca :=_fc ;for _dec ,_gbdd :=range _ddca .Elements (){if _bec ,_bac :=_gbdd .(*_ae .PdfObjectReference );_bac {_cfag ,_acd :=_dc [_bec .ObjectNumber ];if !_acd {return _gc .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_ddca .Set (_dec ,_cfag );}else {_fe (_gbdd ,_dc );};};};return nil ;};func CompareImages (img1 ,img2 _a .Image )(bool ,error ){_gd :=img1 .Bounds ();_cdf :=0;for _gdg :=0;_gdg < _gd .Size ().X ;_gdg ++{for _ee :=0;_ee < _gd .Size ().Y ;_ee ++{_gbd ,_ded ,_db ,_ :=img1 .At (_gdg ,_ee ).RGBA ();
_ac ,_ba ,_bfc ,_ :=img2 .At (_gdg ,_ee ).RGBA ();if _gbd !=_ac ||_ded !=_ba ||_db !=_bfc {_cdf ++;};};};_eac :=float64 (_cdf )/float64 (_gd .Dx ()*_gd .Dy ());if _eac > 0.0001{_f .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_eac ,_cdf );
return false ,nil ;};return true ,nil ;};var (ErrRenderNotSupported =_gc .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m");
);func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;};if _ ,_aef :=_gb .LookPath ("\u0067\u0073");_aef !=nil {return ErrRenderNotSupported ;};return _gb .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_f .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();
};func CompareDictionariesDeep (d1 ,d2 *_ae .PdfObjectDictionary )bool {if len (d1 .Keys ())!=len (d2 .Keys ()){_bf .Log .Debug ("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",len (d1 .Keys ()),len (d2 .Keys ()));
_bf .Log .Debug ("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027",d1 .WriteString (),d2 .WriteString ());return false ;};for _ ,_bbc :=range d1 .Keys (){if _bbc =="\u0050\u0061\u0072\u0065\u006e\u0074"{continue ;
};_acc :=_ae .TraceToDirectObject (d1 .Get (_bbc ));_efg :=_ae .TraceToDirectObject (d2 .Get (_bbc ));if _acc ==nil {_bf .Log .Debug ("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l");return false ;};if _efg ==nil {_bf .Log .Debug ("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l");
return false ;};switch _gf :=_acc .(type ){case *_ae .PdfObjectDictionary :_fba ,_ag :=_efg .(*_ae .PdfObjectDictionary );if !_ag {_bf .Log .Debug ("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054",_acc ,_efg );
return false ;};if !CompareDictionariesDeep (_gf ,_fba ){return false ;};continue ;case *_ae .PdfObjectArray :_cfb ,_aa :=_efg .(*_ae .PdfObjectArray );if !_aa {_bf .Log .Debug ("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");
return false ;};if _gf .Len ()!=_cfb .Len (){_bf .Log .Debug ("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",_gf .Len (),_cfb .Len ());
return false ;};for _afd :=0;_afd < _gf .Len ();_afd ++{_gbe :=_ae .TraceToDirectObject (_gf .Get (_afd ));_egb :=_ae .TraceToDirectObject (_cfb .Get (_afd ));if _ab ,_bd :=_gbe .(*_ae .PdfObjectDictionary );_bd {_fd ,_bfd :=_egb .(*_ae .PdfObjectDictionary );
if !_bfd {return false ;};if !CompareDictionariesDeep (_ab ,_fd ){return false ;};}else {if _gbe .WriteString ()!=_egb .WriteString (){_bf .Log .Debug ("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027",_gbe .WriteString (),_egb .WriteString ());
return false ;};};};continue ;};if _acc .String ()!=_efg .String (){_bf .Log .Debug ("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027",_bbc ,_acc .String (),_efg .String ());
_bf .Log .Debug ("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027",_acc ,_efg );_bf .Log .Debug ("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027",_acc ,_efg );return false ;
};};return true ;};func ReadPNG (file string )(_a .Image ,error ){_gg ,_aec :=_ge .Open (file );if _aec !=nil {return nil ,_aec ;};defer _gg .Close ();return _cd .Decode (_gg );};func CopyFile (src ,dst string )error {_ad ,_af :=_ge .Open (src );if _af !=nil {return _af ;
};defer _ad .Close ();_de ,_af :=_ge .Create (dst );if _af !=nil {return _af ;};defer _de .Close ();_ ,_af =_ec .Copy (_de ,_ad );return _af ;};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_dedf ,_ef :=HashFile (file1 );if _ef !=nil {return false ,_ef ;
};_eb ,_ef :=HashFile (file2 );if _ef !=nil {return false ,_ef ;};if _dedf ==_eb {return true ,nil ;};_bb ,_ef :=ReadPNG (file1 );if _ef !=nil {return false ,_ef ;};_efa ,_ef :=ReadPNG (file2 );if _ef !=nil {return false ,_ef ;};if _bb .Bounds ()!=_efa .Bounds (){return false ,nil ;
};return CompareImages (_bb ,_efa );};func HashFile (file string )(string ,error ){_ce ,_ea :=_ge .Open (file );if _ea !=nil {return "",_ea ;};defer _ce .Close ();_ff :=_c .New ();if _ ,_ea =_ec .Copy (_ff ,_ce );_ea !=nil {return "",_ea ;};return _b .EncodeToString (_ff .Sum (nil )),nil ;
};func RunRenderTest (t *_g .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ){_eff :=_cf .TrimSuffix (_d .Base (pdfPath ),_d .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_ecf *_g .T ){_ga :=_d .Join (outputDir ,_eff );
_adc :=_ga +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _bc :=RenderPDFToPNGs (pdfPath ,0,_adc );_bc !=nil {_ecf .Skip (_bc );};for _cg :=1;true ;_cg ++{_ggc :=_f .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_ga ,_cg );_be :=_d .Join (baselineRenderPath ,_f .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_eff ,_cg ));
if _ ,_ace :=_ge .Stat (_ggc );_ace !=nil {break ;};_ecf .Logf ("\u0025\u0073",_be );if saveBaseline {_ecf .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_ggc ,_be );_eag :=CopyFile (_ggc ,_be );if _eag !=nil {_ecf .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_be ,_eag );
};continue ;};_ecf .Run (_f .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_cg ),func (_ffa *_g .T ){_ffa .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_ggc ,_be );_gee ,_gab :=ComparePNGFiles (_ggc ,_be );
if _ge .IsNotExist (_gab ){_ffa .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_gee {_ffa .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");
};});};});};func ParseIndirectObjects (rawpdf string )(map[int64 ]_ae .PdfObject ,error ){_cfe :=_ae .NewParserFromString (rawpdf );_dd :=map[int64 ]_ae .PdfObject {};for {_bga ,_ddc :=_cfe .ParseIndirectObject ();if _ddc !=nil {if _ddc ==_ec .EOF {break ;
};return nil ,_ddc ;};switch _ca :=_bga .(type ){case *_ae .PdfIndirectObject :_dd [_ca .ObjectNumber ]=_bga ;case *_ae .PdfObjectStream :_dd [_ca .ObjectNumber ]=_bga ;};};for _ ,_dde :=range _dd {_fe (_dde ,_dd );};return _dd ,nil ;};