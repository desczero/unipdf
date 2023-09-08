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

package jbig2 ;import (_f "github.com/unidoc/unipdf/v3/internal/bitwise";_ce "github.com/unidoc/unipdf/v3/internal/jbig2/decoder";_e "github.com/unidoc/unipdf/v3/internal/jbig2/document";_b "github.com/unidoc/unipdf/v3/internal/jbig2/document/segments";
_cc "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_c "sort";);type Globals map[int ]*_b .Header ;func DecodeBytes (encoded []byte ,parameters _ce .Parameters ,globals ...Globals )([]byte ,error ){var _db Globals ;if len (globals )> 0{_db =globals [0];
};_g ,_fb :=_ce .Decode (encoded ,parameters ,_db .ToDocumentGlobals ());if _fb !=nil {return nil ,_fb ;};return _g .DecodeNextPage ();};func (_dc Globals )ToDocumentGlobals ()*_e .Globals {if _dc ==nil {return nil ;};_gc :=[]*_b .Header {};for _ ,_cec :=range _dc {_gc =append (_gc ,_cec );
};_c .Slice (_gc ,func (_bg ,_a int )bool {return _gc [_bg ].SegmentNumber < _gc [_a ].SegmentNumber });return &_e .Globals {Segments :_gc };};func DecodeGlobals (encoded []byte )(Globals ,error ){const _ge ="\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073";
_gd :=_f .NewReader (encoded );_cf ,_bf :=_e .DecodeDocument (_gd ,nil );if _bf !=nil {return nil ,_cc .Wrap (_bf ,_ge ,"");};if _cf .GlobalSegments ==nil ||(_cf .GlobalSegments .Segments ==nil ){return nil ,_cc .Error (_ge ,"\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064");
};_eb :=Globals {};for _ ,_ef :=range _cf .GlobalSegments .Segments {_eb [int (_ef .SegmentNumber )]=_ef ;};return _eb ,nil ;};