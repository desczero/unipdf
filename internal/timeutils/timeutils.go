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

package timeutils ;import (_b "errors";_g "fmt";_a "regexp";_bf "strconv";_c "time";);var _fd =_a .MustCompile ("\u005cs\u002a\u0044\u005cs\u002a\u003a\u005cs\u002a(\\\u0064\u007b\u0034\u007d\u0029\u0028\u005cd\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064{2\u007d)\u003f\u0028\u005b\u002b\u002d\u005a]\u0029\u003f\u0028\u005c\u0064{\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032}\u0029\u003f");
func ParsePdfTime (pdfTime string )(_c .Time ,error ){_cf :=_fd .FindAllStringSubmatch (pdfTime ,1);if len (_cf )< 1{if len (pdfTime )> 0&&pdfTime [0]!='D'{pdfTime =_g .Sprintf ("\u0044\u003a\u0025\u0073",pdfTime );return ParsePdfTime (pdfTime );};return _c .Time {},_g .Errorf ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029",pdfTime );
};if len (_cf [0])!=10{return _c .Time {},_b .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030");
};_bac ,_ :=_bf .ParseInt (_cf [0][1],10,32);_ag ,_ :=_bf .ParseInt (_cf [0][2],10,32);_ac ,_ :=_bf .ParseInt (_cf [0][3],10,32);_dc ,_ :=_bf .ParseInt (_cf [0][4],10,32);_ab ,_ :=_bf .ParseInt (_cf [0][5],10,32);_cb ,_ :=_bf .ParseInt (_cf [0][6],10,32);
var (_cc byte ;_gg int64 ;_bad int64 ;);_cc ='+';if len (_cf [0][7])> 0{if _cf [0][7]=="\u002d"{_cc ='-';}else if _cf [0][7]=="\u005a"{_cc ='Z';};};if len (_cf [0][8])> 0{_gg ,_ =_bf .ParseInt (_cf [0][8],10,32);}else {_gg =0;};if len (_cf [0][9])> 0{_bad ,_ =_bf .ParseInt (_cf [0][9],10,32);
}else {_bad =0;};_ca :=int (_gg *60*60+_bad *60);switch _cc {case '-':_ca =-_ca ;case 'Z':_ca =0;};_bb :=_g .Sprintf ("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064",_cc ,_gg ,_bad );_gea :=_c .FixedZone (_bb ,_ca );return _c .Date (int (_bac ),_c .Month (_ag ),int (_ac ),int (_dc ),int (_ab ),int (_cb ),0,_gea ),nil ;
};func FormatPdfTime (in _c .Time )string {_ee :=in .Format ("\u002d\u0030\u0037\u003a\u0030\u0030");_ge ,_ :=_bf .ParseInt (_ee [1:3],10,32);_fg ,_ :=_bf .ParseInt (_ee [4:6],10,32);_cd :=int64 (in .Year ());_ba :=int64 (in .Month ());_fgc :=int64 (in .Day ());
_fcf :=int64 (in .Hour ());_fgb :=int64 (in .Minute ());_bg :=int64 (in .Second ());_d :=_ee [0];return _g .Sprintf ("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027",_cd ,_ba ,_fgc ,_fcf ,_fgb ,_bg ,_d ,_ge ,_fg );
};