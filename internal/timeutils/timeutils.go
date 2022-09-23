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

package timeutils ;import (_d "errors";_c "fmt";_be "regexp";_g "strconv";_bc "time";);func ParsePdfTime (pdfTime string )(_bc .Time ,error ){_af :=_ab .FindAllStringSubmatch (pdfTime ,1);if len (_af )< 1{return _bc .Time {},_c .Errorf ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029",pdfTime );
};if len (_af [0])!=10{return _bc .Time {},_d .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030");
};_gg ,_ :=_g .ParseInt (_af [0][1],10,32);_ed ,_ :=_g .ParseInt (_af [0][2],10,32);_f ,_ :=_g .ParseInt (_af [0][3],10,32);_cb ,_ :=_g .ParseInt (_af [0][4],10,32);_ad ,_ :=_g .ParseInt (_af [0][5],10,32);_edg ,_ :=_g .ParseInt (_af [0][6],10,32);var (_cd byte ;
_gb int64 ;_bcf int64 ;);if len (_af [0][7])> 0{_cd =_af [0][7][0];}else {_cd ='+';};if len (_af [0][8])> 0{_gb ,_ =_g .ParseInt (_af [0][8],10,32);}else {_gb =0;};if len (_af [0][9])> 0{_bcf ,_ =_g .ParseInt (_af [0][9],10,32);}else {_bcf =0;};_gd :=int (_gb *60*60+_bcf *60);
switch _cd {case '-':_gd =-_gd ;case 'Z':_gd =0;};_gga :=_c .Sprintf ("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064",_cd ,_gb ,_bcf );_eba :=_bc .FixedZone (_gga ,_gd );return _bc .Date (int (_gg ),_bc .Month (_ed ),int (_f ),int (_cb ),int (_ad ),int (_edg ),0,_eba ),nil ;
};var _ab =_be .MustCompile ("\u005c\u0073\u002a\u0044\u005c\u0073\u002a:\u005c\u0073\u002a\u0028\u005c\u0064\u007b\u0034\u007d\u0029\u0028\u005c\u0064\u007b2\u007d)\u0028\u005c\u0064\u007b\u0032\u007d)\u0028\u005c\u0064\u007b\u0032\u007d\u0029(\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005b\u002b\u002d\u005a\u005d\u0029\u003f\u0028\u005cd\u007b\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032\u007d)\u003f");
func FormatPdfTime (in _bc .Time )string {_bb :=in .Format ("\u002d\u0030\u0037\u003a\u0030\u0030");_ga ,_ :=_g .ParseInt (_bb [1:3],10,32);_a ,_ :=_g .ParseInt (_bb [4:6],10,32);_dd :=int64 (in .Year ());_bg :=int64 (in .Month ());_cf :=int64 (in .Day ());
_cfg :=int64 (in .Hour ());_ddb :=int64 (in .Minute ());_dg :=int64 (in .Second ());_e :=_bb [0];return _c .Sprintf ("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027",_dd ,_bg ,_cf ,_cfg ,_ddb ,_dg ,_e ,_ga ,_a );
};