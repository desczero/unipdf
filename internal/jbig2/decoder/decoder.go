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

package decoder ;import (_d "github.com/unidoc/unipdf/v3/internal/bitwise";_e "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_dg "github.com/unidoc/unipdf/v3/internal/jbig2/document";_b "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_f "image";
);func (_dd *Decoder )DecodeNextPage ()([]byte ,error ){_dd ._bg ++;_ddb :=_dd ._bg ;return _dd .decodePage (_ddb );};func (_fg *Decoder )DecodePageImage (pageNumber int )(_f .Image ,error ){const _fb ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";
_ag ,_ff :=_fg .decodePageImage (pageNumber );if _ff !=nil {return nil ,_b .Wrap (_ff ,_fb ,"");};return _ag ,nil ;};type Decoder struct{_bb *_d .Reader ;_eg *_dg .Document ;_bg int ;_g Parameters ;};func (_bd *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _bd .decodePage (pageNumber )};
func (_fe *Decoder )PageNumber ()(int ,error ){const _dga ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";if _fe ._eg ==nil {return 0,_b .Error (_dga ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");
};return int (_fe ._eg .NumberOfPages ),nil ;};func (_fc *Decoder )decodePage (_ee int )([]byte ,error ){const _c ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";if _ee < 0{return nil ,_b .Errorf (_c ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_ee );
};if _ee > int (_fc ._eg .NumberOfPages ){return nil ,_b .Errorf (_c ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_ee );
};_ab ,_fff :=_fc ._eg .GetPage (_ee );if _fff !=nil {return nil ,_b .Wrap (_fff ,_c ,"");};_fa ,_fff :=_ab .GetBitmap ();if _fff !=nil {return nil ,_b .Wrap (_fff ,_c ,"");};_fa .InverseData ();if !_fc ._g .UnpaddedData {return _fa .Data ,nil ;};return _fa .GetUnpaddedData ();
};func Decode (input []byte ,parameters Parameters ,globals *_dg .Globals )(*Decoder ,error ){_ec :=_d .NewReader (input );_ge ,_de :=_dg .DecodeDocument (_ec ,globals );if _de !=nil {return nil ,_de ;};return &Decoder {_bb :_ec ,_eg :_ge ,_g :parameters },nil ;
};func (_cg *Decoder )decodePageImage (_fd int )(_f .Image ,error ){const _eef ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";if _fd < 0{return nil ,_b .Errorf (_eef ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_fd );
};if _fd > int (_cg ._eg .NumberOfPages ){return nil ,_b .Errorf (_eef ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_fd );
};_eb ,_ea :=_cg ._eg .GetPage (_fd );if _ea !=nil {return nil ,_b .Wrap (_ea ,_eef ,"");};_bdb ,_ea :=_eb .GetBitmap ();if _ea !=nil {return nil ,_b .Wrap (_ea ,_eef ,"");};_bdb .InverseData ();return _bdb .ToImage (),nil ;};type Parameters struct{UnpaddedData bool ;
Color _e .Color ;};