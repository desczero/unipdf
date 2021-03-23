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

package pdfutil ;import (_f "github.com/unidoc/unipdf/v3/common";_b "github.com/unidoc/unipdf/v3/contentstream";_ee "github.com/unidoc/unipdf/v3/contentstream/draw";_g "github.com/unidoc/unipdf/v3/core";_c "github.com/unidoc/unipdf/v3/model";);

// NormalizePage performs the following operations on the passed in page:
// - Normalize the page rotation.
//   Rotates the contents of the page according to the Rotate entry, thus
//   flattening the rotation. The Rotate entry of the page is set to nil.
// - Normalize the media box.
//   If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//   the contents of the page are translated to (-Llx, -Lly). After
//   normalization, the media box is updated (Llx and Lly are set to 0 and
//   Urx and Ury are updated accordingly).
// - Normalize the crop box.
//   The crop box of the page is updated based on the previous operations.
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage (page *_c .PdfPage )error {_cg ,_bd :=page .GetMediaBox ();if _bd !=nil {return _bd ;};_ga :=page .Rotate ;_a :=_ga !=nil &&*_ga %360!=0&&*_ga %90==0;_cg .Normalize ();_d ,_fg ,_cgd ,_eg :=_cg .Llx ,_cg .Lly ,_cg .Width (),_cg .Height ();
_gaa :=_d !=0||_fg !=0;if !_a &&!_gaa {return nil ;};_db :=func (_fd ,_ad ,_aa float64 )_ee .BoundingBox {return _ee .Path {Points :[]_ee .Point {_ee .NewPoint (0,0).Rotate (_aa ),_ee .NewPoint (_fd ,0).Rotate (_aa ),_ee .NewPoint (0,_ad ).Rotate (_aa ),_ee .NewPoint (_fd ,_ad ).Rotate (_aa )}}.GetBoundingBox ();
};_ed :=_b .NewContentCreator ();var _ae float64 ;if _a {_ae =-float64 (*page .Rotate );_ag :=_db (_cgd ,_eg ,_ae );_ed .Translate ((_ag .Width -_cgd )/2+_cgd /2,(_ag .Height -_eg )/2+_eg /2);_ed .RotateDeg (_ae );_ed .Translate (-_cgd /2,-_eg /2);_cgd ,_eg =_ag .Width ,_ag .Height ;
};if _gaa {_ed .Translate (-_d ,-_fg );};_bdf :=_ed .Operations ();_ba ,_bd :=_g .MakeStream (_bdf .Bytes (),_g .NewFlateEncoder ());if _bd !=nil {return _bd ;};_eb :=_g .MakeArray (_ba );_eb .Append (page .GetContentStreamObjs ()...);*_cg =_c .PdfRectangle {Urx :_cgd ,Ury :_eg };
if _add :=page .CropBox ;_add !=nil {_add .Normalize ();_gf ,_fe ,_gc ,_ce :=_add .Llx -_d ,_add .Lly -_fg ,_add .Width (),_add .Height ();if _a {_cb :=_db (_gc ,_ce ,_ae );_gc ,_ce =_cb .Width ,_cb .Height ;};*_add =_c .PdfRectangle {Llx :_gf ,Lly :_fe ,Urx :_gf +_gc ,Ury :_fe +_ce };
};_f .Log .Debug ("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_ae ,_bdf ,_cg );page .Contents =_eb ;page .Rotate =nil ;
return nil ;};