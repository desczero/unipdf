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

package jbig2 ;import (_cc "github.com/unidoc/unipdf/v3/internal/bitwise";_a "github.com/unidoc/unipdf/v3/internal/jbig2/decoder";_ae "github.com/unidoc/unipdf/v3/internal/jbig2/document";_e "github.com/unidoc/unipdf/v3/internal/jbig2/document/segments";
_cb "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_f "sort";);func DecodeGlobals (encoded []byte )(Globals ,error ){const _g ="\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073";_d :=_cc .NewReader (encoded );_gg ,_bf :=_ae .DecodeDocument (_d ,nil );
if _bf !=nil {return nil ,_cb .Wrap (_bf ,_g ,"");};if _gg .GlobalSegments ==nil ||(_gg .GlobalSegments .Segments ==nil ){return nil ,_cb .Error (_g ,"\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064");
};_gd :=Globals {};for _ ,_ce :=range _gg .GlobalSegments .Segments {_gd [int (_ce .SegmentNumber )]=_ce ;};return _gd ,nil ;};type Globals map[int ]*_e .Header ;func (_cg Globals )ToDocumentGlobals ()*_ae .Globals {if _cg ==nil {return nil ;};_ea :=[]*_e .Header {};
for _ ,_ad :=range _cg {_ea =append (_ea ,_ad );};_f .Slice (_ea ,func (_fa ,_ceg int )bool {return _ea [_fa ].SegmentNumber < _ea [_ceg ].SegmentNumber });return &_ae .Globals {Segments :_ea };};func DecodeBytes (encoded []byte ,parameters _a .Parameters ,globals ...Globals )([]byte ,error ){var _fb Globals ;
if len (globals )> 0{_fb =globals [0];};_b ,_ec :=_a .Decode (encoded ,parameters ,_fb .ToDocumentGlobals ());if _ec !=nil {return nil ,_ec ;};return _b .DecodeNextPage ();};