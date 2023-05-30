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

package timeutils ;import (_f "errors";_ff "fmt";_d "regexp";_be "strconv";_bc "time";);var _bf =_d .MustCompile ("\u005c\u0073\u002a\u0044\u005c\u0073\u002a:\u005c\u0073\u002a\u0028\u005c\u0064\u007b\u0034\u007d\u0029\u0028\u005c\u0064\u007b2\u007d)\u0028\u005c\u0064\u007b\u0032\u007d)\u0028\u005c\u0064\u007b\u0032\u007d\u0029(\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005b\u002b\u002d\u005a\u005d\u0029\u003f\u0028\u005cd\u007b\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032\u007d)\u003f");
func FormatPdfTime (in _bc .Time )string {_c :=in .Format ("\u002d\u0030\u0037\u003a\u0030\u0030");_e ,_ :=_be .ParseInt (_c [1:3],10,32);_bea ,_ :=_be .ParseInt (_c [4:6],10,32);_beaf :=int64 (in .Year ());_a :=int64 (in .Month ());_bb :=int64 (in .Day ());
_fff :=int64 (in .Hour ());_fg :=int64 (in .Minute ());_bba :=int64 (in .Second ());_dc :=_c [0];return _ff .Sprintf ("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027",_beaf ,_a ,_bb ,_fff ,_fg ,_bba ,_dc ,_e ,_bea );
};func ParsePdfTime (pdfTime string )(_bc .Time ,error ){_ba :=_bf .FindAllStringSubmatch (pdfTime ,1);if len (_ba )< 1{if len (pdfTime )> 0&&pdfTime [0]!='D'{pdfTime =_ff .Sprintf ("\u0044\u003a\u0025\u0073",pdfTime );return ParsePdfTime (pdfTime );};
return _bc .Time {},_ff .Errorf ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029",pdfTime );};if len (_ba [0])!=10{return _bc .Time {},_f .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030");
};_g ,_ :=_be .ParseInt (_ba [0][1],10,32);_cf ,_ :=_be .ParseInt (_ba [0][2],10,32);_beb ,_ :=_be .ParseInt (_ba [0][3],10,32);_ad ,_ :=_be .ParseInt (_ba [0][4],10,32);_gd ,_ :=_be .ParseInt (_ba [0][5],10,32);_ge ,_ :=_be .ParseInt (_ba [0][6],10,32);
var (_aa byte ;_bbe int64 ;_df int64 ;);if len (_ba [0][7])> 0{_aa =_ba [0][7][0];}else {_aa ='+';};if len (_ba [0][8])> 0{_bbe ,_ =_be .ParseInt (_ba [0][8],10,32);}else {_bbe =0;};if len (_ba [0][9])> 0{_df ,_ =_be .ParseInt (_ba [0][9],10,32);}else {_df =0;
};_ag :=int (_bbe *60*60+_df *60);switch _aa {case '-':_ag =-_ag ;case 'Z':_ag =0;};_ffff :=_ff .Sprintf ("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064",_aa ,_bbe ,_df );_dd :=_bc .FixedZone (_ffff ,_ag );return _bc .Date (int (_g ),_bc .Month (_cf ),int (_beb ),int (_ad ),int (_gd ),int (_ge ),0,_dd ),nil ;
};