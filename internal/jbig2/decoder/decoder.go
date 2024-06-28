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

package decoder ;import (_g "github.com/unidoc/unipdf/v3/internal/bitwise";_d "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_b "github.com/unidoc/unipdf/v3/internal/jbig2/document";_ge "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_e "image";
);func (_ed *Decoder )decodePage (_af int )([]byte ,error ){const _ae ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";if _af < 0{return nil ,_ge .Errorf (_ae ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_af );
};if _af > int (_ed ._f .NumberOfPages ){return nil ,_ge .Errorf (_ae ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_af );
};_ea ,_df :=_ed ._f .GetPage (_af );if _df !=nil {return nil ,_ge .Wrap (_df ,_ae ,"");};_dab ,_df :=_ea .GetBitmap ();if _df !=nil {return nil ,_ge .Wrap (_df ,_ae ,"");};_dab .InverseData ();if !_ed ._ef .UnpaddedData {return _dab .Data ,nil ;};return _dab .GetUnpaddedData ();
};func (_edf *Decoder )decodePageImage (_cg int )(_e .Image ,error ){const _fb ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";if _cg < 0{return nil ,_ge .Errorf (_fb ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_cg );
};if _cg > int (_edf ._f .NumberOfPages ){return nil ,_ge .Errorf (_fb ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_cg );
};_efd ,_dfg :=_edf ._f .GetPage (_cg );if _dfg !=nil {return nil ,_ge .Wrap (_dfg ,_fb ,"");};_gf ,_dfg :=_efd .GetBitmap ();if _dfg !=nil {return nil ,_ge .Wrap (_dfg ,_fb ,"");};_gf .InverseData ();return _gf .ToImage (),nil ;};type Decoder struct{_bb *_g .Reader ;
_f *_b .Document ;_gd int ;_ef Parameters ;};func (_dc *Decoder )DecodePageImage (pageNumber int )(_e .Image ,error ){const _da ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";
_bbb ,_geb :=_dc .decodePageImage (pageNumber );if _geb !=nil {return nil ,_ge .Wrap (_geb ,_da ,"");};return _bbb ,nil ;};func Decode (input []byte ,parameters Parameters ,globals *_b .Globals )(*Decoder ,error ){_fd :=_g .NewReader (input );_db ,_ab :=_b .DecodeDocument (_fd ,globals );
if _ab !=nil {return nil ,_ab ;};return &Decoder {_bb :_fd ,_f :_db ,_ef :parameters },nil ;};func (_ega *Decoder )PageNumber ()(int ,error ){const _ff ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";if _ega ._f ==nil {return 0,_ge .Error (_ff ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");
};return int (_ega ._f .NumberOfPages ),nil ;};func (_eg *Decoder )DecodeNextPage ()([]byte ,error ){_eg ._gd ++;_a :=_eg ._gd ;return _eg .decodePage (_a );};func (_bg *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _bg .decodePage (pageNumber )};
type Parameters struct{UnpaddedData bool ;Color _d .Color ;};