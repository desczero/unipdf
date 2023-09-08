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

package graphic2d ;import (_d "image/color";_g "math";);func _fb (_dee ,_bfd float64 )bool {return _g .Abs (_dee -_bfd )<=_ffe };func (_dfb Point )Add (q Point )Point {return Point {_dfb .X +q .X ,_dfb .Y +q .Y }};type Point struct{X ,Y float64 ;};func _bc (_cg ,_ece ,_df ,_bcg ,_gb ,_bdf float64 )(float64 ,float64 ){_ea ,_dc :=_g .Sincos (_bdf );
_dfa ,_ca :=_g .Sincos (_df );_cac :=_bcg +_cg *_dc *_ca -_ece *_ea *_dfa ;_ee :=_gb +_cg *_dc *_dfa +_ece *_ea *_ca ;return _cac ,_ee ;};func QuadraticToCubicBezier (startX ,startY ,x1 ,y1 ,x ,y float64 )(Point ,Point ){_cfe :=Point {X :startX ,Y :startY };
_bad :=Point {X :x1 ,Y :y1 };_aa :=Point {X :x ,Y :y };_af :=_cfe .Interpolate (_bad ,2.0/3.0);_gbg :=_aa .Interpolate (_bad ,2.0/3.0);return _af ,_gbg ;};func _gdf (_cb ,_dd ,_fdg ,_dff ,_cf float64 ,_bfb ,_eaa bool ,_ff ,_eaaa float64 )(float64 ,float64 ,float64 ,float64 ){if _fb (_cb ,_ff )&&_fb (_dd ,_eaaa ){return _cb ,_dd ,0.0,0.0;
};_cfg ,_fgb :=_g .Sincos (_cf );_cfga :=_fgb *(_cb -_ff )/2.0+_cfg *(_dd -_eaaa )/2.0;_cbc :=-_cfg *(_cb -_ff )/2.0+_fgb *(_dd -_eaaa )/2.0;_gbb :=_cfga *_cfga /_fdg /_fdg +_cbc *_cbc /_dff /_dff ;if _gbb > 1.0{_fdg *=_g .Sqrt (_gbb );_dff *=_g .Sqrt (_gbb );
};_ffb :=(_fdg *_fdg *_dff *_dff -_fdg *_fdg *_cbc *_cbc -_dff *_dff *_cfga *_cfga )/(_fdg *_fdg *_cbc *_cbc +_dff *_dff *_cfga *_cfga );if _ffb < 0.0{_ffb =0.0;};_ed :=_g .Sqrt (_ffb );if _bfb ==_eaa {_ed =-_ed ;};_dda :=_ed *_fdg *_cbc /_dff ;_be :=_ed *-_dff *_cfga /_fdg ;
_dcc :=_fgb *_dda -_cfg *_be +(_cb +_ff )/2.0;_cgg :=_cfg *_dda +_fgb *_be +(_dd +_eaaa )/2.0;_bfbg :=(_cfga -_dda )/_fdg ;_ge :=(_cbc -_be )/_dff ;_ba :=-(_cfga +_dda )/_fdg ;_gf :=-(_cbc +_be )/_dff ;_gde :=_g .Acos (_bfbg /_g .Sqrt (_bfbg *_bfbg +_ge *_ge ));
if _ge < 0.0{_gde =-_gde ;};_gde =_ced (_gde );_egc :=(_bfbg *_ba +_ge *_gf )/_g .Sqrt ((_bfbg *_bfbg +_ge *_ge )*(_ba *_ba +_gf *_gf ));_egc =_g .Min (1.0,_g .Max (-1.0,_egc ));_bee :=_g .Acos (_egc );if _bfbg *_gf -_ge *_ba < 0.0{_bee =-_bee ;};if !_eaa &&_bee > 0.0{_bee -=2.0*_g .Pi ;
}else if _eaa &&_bee < 0.0{_bee +=2.0*_g .Pi ;};return _dcc ,_cgg ,_gde ,_gde +_bee ;};func (_cca Point )Interpolate (q Point ,t float64 )Point {return Point {(1-t )*_cca .X +t *q .X ,(1-t )*_cca .Y +t *q .Y };};const _ffe =1e-10;var (Aliceblue =_d .RGBA {0xf0,0xf8,0xff,0xff};
Antiquewhite =_d .RGBA {0xfa,0xeb,0xd7,0xff};Aqua =_d .RGBA {0x00,0xff,0xff,0xff};Aquamarine =_d .RGBA {0x7f,0xff,0xd4,0xff};Azure =_d .RGBA {0xf0,0xff,0xff,0xff};Beige =_d .RGBA {0xf5,0xf5,0xdc,0xff};Bisque =_d .RGBA {0xff,0xe4,0xc4,0xff};Black =_d .RGBA {0x00,0x00,0x00,0xff};
Blanchedalmond =_d .RGBA {0xff,0xeb,0xcd,0xff};Blue =_d .RGBA {0x00,0x00,0xff,0xff};Blueviolet =_d .RGBA {0x8a,0x2b,0xe2,0xff};Brown =_d .RGBA {0xa5,0x2a,0x2a,0xff};Burlywood =_d .RGBA {0xde,0xb8,0x87,0xff};Cadetblue =_d .RGBA {0x5f,0x9e,0xa0,0xff};Chartreuse =_d .RGBA {0x7f,0xff,0x00,0xff};
Chocolate =_d .RGBA {0xd2,0x69,0x1e,0xff};Coral =_d .RGBA {0xff,0x7f,0x50,0xff};Cornflowerblue =_d .RGBA {0x64,0x95,0xed,0xff};Cornsilk =_d .RGBA {0xff,0xf8,0xdc,0xff};Crimson =_d .RGBA {0xdc,0x14,0x3c,0xff};Cyan =_d .RGBA {0x00,0xff,0xff,0xff};Darkblue =_d .RGBA {0x00,0x00,0x8b,0xff};
Darkcyan =_d .RGBA {0x00,0x8b,0x8b,0xff};Darkgoldenrod =_d .RGBA {0xb8,0x86,0x0b,0xff};Darkgray =_d .RGBA {0xa9,0xa9,0xa9,0xff};Darkgreen =_d .RGBA {0x00,0x64,0x00,0xff};Darkgrey =_d .RGBA {0xa9,0xa9,0xa9,0xff};Darkkhaki =_d .RGBA {0xbd,0xb7,0x6b,0xff};
Darkmagenta =_d .RGBA {0x8b,0x00,0x8b,0xff};Darkolivegreen =_d .RGBA {0x55,0x6b,0x2f,0xff};Darkorange =_d .RGBA {0xff,0x8c,0x00,0xff};Darkorchid =_d .RGBA {0x99,0x32,0xcc,0xff};Darkred =_d .RGBA {0x8b,0x00,0x00,0xff};Darksalmon =_d .RGBA {0xe9,0x96,0x7a,0xff};
Darkseagreen =_d .RGBA {0x8f,0xbc,0x8f,0xff};Darkslateblue =_d .RGBA {0x48,0x3d,0x8b,0xff};Darkslategray =_d .RGBA {0x2f,0x4f,0x4f,0xff};Darkslategrey =_d .RGBA {0x2f,0x4f,0x4f,0xff};Darkturquoise =_d .RGBA {0x00,0xce,0xd1,0xff};Darkviolet =_d .RGBA {0x94,0x00,0xd3,0xff};
Deeppink =_d .RGBA {0xff,0x14,0x93,0xff};Deepskyblue =_d .RGBA {0x00,0xbf,0xff,0xff};Dimgray =_d .RGBA {0x69,0x69,0x69,0xff};Dimgrey =_d .RGBA {0x69,0x69,0x69,0xff};Dodgerblue =_d .RGBA {0x1e,0x90,0xff,0xff};Firebrick =_d .RGBA {0xb2,0x22,0x22,0xff};Floralwhite =_d .RGBA {0xff,0xfa,0xf0,0xff};
Forestgreen =_d .RGBA {0x22,0x8b,0x22,0xff};Fuchsia =_d .RGBA {0xff,0x00,0xff,0xff};Gainsboro =_d .RGBA {0xdc,0xdc,0xdc,0xff};Ghostwhite =_d .RGBA {0xf8,0xf8,0xff,0xff};Gold =_d .RGBA {0xff,0xd7,0x00,0xff};Goldenrod =_d .RGBA {0xda,0xa5,0x20,0xff};Gray =_d .RGBA {0x80,0x80,0x80,0xff};
Green =_d .RGBA {0x00,0x80,0x00,0xff};Greenyellow =_d .RGBA {0xad,0xff,0x2f,0xff};Grey =_d .RGBA {0x80,0x80,0x80,0xff};Honeydew =_d .RGBA {0xf0,0xff,0xf0,0xff};Hotpink =_d .RGBA {0xff,0x69,0xb4,0xff};Indianred =_d .RGBA {0xcd,0x5c,0x5c,0xff};Indigo =_d .RGBA {0x4b,0x00,0x82,0xff};
Ivory =_d .RGBA {0xff,0xff,0xf0,0xff};Khaki =_d .RGBA {0xf0,0xe6,0x8c,0xff};Lavender =_d .RGBA {0xe6,0xe6,0xfa,0xff};Lavenderblush =_d .RGBA {0xff,0xf0,0xf5,0xff};Lawngreen =_d .RGBA {0x7c,0xfc,0x00,0xff};Lemonchiffon =_d .RGBA {0xff,0xfa,0xcd,0xff};Lightblue =_d .RGBA {0xad,0xd8,0xe6,0xff};
Lightcoral =_d .RGBA {0xf0,0x80,0x80,0xff};Lightcyan =_d .RGBA {0xe0,0xff,0xff,0xff};Lightgoldenrodyellow =_d .RGBA {0xfa,0xfa,0xd2,0xff};Lightgray =_d .RGBA {0xd3,0xd3,0xd3,0xff};Lightgreen =_d .RGBA {0x90,0xee,0x90,0xff};Lightgrey =_d .RGBA {0xd3,0xd3,0xd3,0xff};
Lightpink =_d .RGBA {0xff,0xb6,0xc1,0xff};Lightsalmon =_d .RGBA {0xff,0xa0,0x7a,0xff};Lightseagreen =_d .RGBA {0x20,0xb2,0xaa,0xff};Lightskyblue =_d .RGBA {0x87,0xce,0xfa,0xff};Lightslategray =_d .RGBA {0x77,0x88,0x99,0xff};Lightslategrey =_d .RGBA {0x77,0x88,0x99,0xff};
Lightsteelblue =_d .RGBA {0xb0,0xc4,0xde,0xff};Lightyellow =_d .RGBA {0xff,0xff,0xe0,0xff};Lime =_d .RGBA {0x00,0xff,0x00,0xff};Limegreen =_d .RGBA {0x32,0xcd,0x32,0xff};Linen =_d .RGBA {0xfa,0xf0,0xe6,0xff};Magenta =_d .RGBA {0xff,0x00,0xff,0xff};Maroon =_d .RGBA {0x80,0x00,0x00,0xff};
Mediumaquamarine =_d .RGBA {0x66,0xcd,0xaa,0xff};Mediumblue =_d .RGBA {0x00,0x00,0xcd,0xff};Mediumorchid =_d .RGBA {0xba,0x55,0xd3,0xff};Mediumpurple =_d .RGBA {0x93,0x70,0xdb,0xff};Mediumseagreen =_d .RGBA {0x3c,0xb3,0x71,0xff};Mediumslateblue =_d .RGBA {0x7b,0x68,0xee,0xff};
Mediumspringgreen =_d .RGBA {0x00,0xfa,0x9a,0xff};Mediumturquoise =_d .RGBA {0x48,0xd1,0xcc,0xff};Mediumvioletred =_d .RGBA {0xc7,0x15,0x85,0xff};Midnightblue =_d .RGBA {0x19,0x19,0x70,0xff};Mintcream =_d .RGBA {0xf5,0xff,0xfa,0xff};Mistyrose =_d .RGBA {0xff,0xe4,0xe1,0xff};
Moccasin =_d .RGBA {0xff,0xe4,0xb5,0xff};Navajowhite =_d .RGBA {0xff,0xde,0xad,0xff};Navy =_d .RGBA {0x00,0x00,0x80,0xff};Oldlace =_d .RGBA {0xfd,0xf5,0xe6,0xff};Olive =_d .RGBA {0x80,0x80,0x00,0xff};Olivedrab =_d .RGBA {0x6b,0x8e,0x23,0xff};Orange =_d .RGBA {0xff,0xa5,0x00,0xff};
Orangered =_d .RGBA {0xff,0x45,0x00,0xff};Orchid =_d .RGBA {0xda,0x70,0xd6,0xff};Palegoldenrod =_d .RGBA {0xee,0xe8,0xaa,0xff};Palegreen =_d .RGBA {0x98,0xfb,0x98,0xff};Paleturquoise =_d .RGBA {0xaf,0xee,0xee,0xff};Palevioletred =_d .RGBA {0xdb,0x70,0x93,0xff};
Papayawhip =_d .RGBA {0xff,0xef,0xd5,0xff};Peachpuff =_d .RGBA {0xff,0xda,0xb9,0xff};Peru =_d .RGBA {0xcd,0x85,0x3f,0xff};Pink =_d .RGBA {0xff,0xc0,0xcb,0xff};Plum =_d .RGBA {0xdd,0xa0,0xdd,0xff};Powderblue =_d .RGBA {0xb0,0xe0,0xe6,0xff};Purple =_d .RGBA {0x80,0x00,0x80,0xff};
Red =_d .RGBA {0xff,0x00,0x00,0xff};Rosybrown =_d .RGBA {0xbc,0x8f,0x8f,0xff};Royalblue =_d .RGBA {0x41,0x69,0xe1,0xff};Saddlebrown =_d .RGBA {0x8b,0x45,0x13,0xff};Salmon =_d .RGBA {0xfa,0x80,0x72,0xff};Sandybrown =_d .RGBA {0xf4,0xa4,0x60,0xff};Seagreen =_d .RGBA {0x2e,0x8b,0x57,0xff};
Seashell =_d .RGBA {0xff,0xf5,0xee,0xff};Sienna =_d .RGBA {0xa0,0x52,0x2d,0xff};Silver =_d .RGBA {0xc0,0xc0,0xc0,0xff};Skyblue =_d .RGBA {0x87,0xce,0xeb,0xff};Slateblue =_d .RGBA {0x6a,0x5a,0xcd,0xff};Slategray =_d .RGBA {0x70,0x80,0x90,0xff};Slategrey =_d .RGBA {0x70,0x80,0x90,0xff};
Snow =_d .RGBA {0xff,0xfa,0xfa,0xff};Springgreen =_d .RGBA {0x00,0xff,0x7f,0xff};Steelblue =_d .RGBA {0x46,0x82,0xb4,0xff};Tan =_d .RGBA {0xd2,0xb4,0x8c,0xff};Teal =_d .RGBA {0x00,0x80,0x80,0xff};Thistle =_d .RGBA {0xd8,0xbf,0xd8,0xff};Tomato =_d .RGBA {0xff,0x63,0x47,0xff};
Turquoise =_d .RGBA {0x40,0xe0,0xd0,0xff};Violet =_d .RGBA {0xee,0x82,0xee,0xff};Wheat =_d .RGBA {0xf5,0xde,0xb3,0xff};White =_d .RGBA {0xff,0xff,0xff,0xff};Whitesmoke =_d .RGBA {0xf5,0xf5,0xf5,0xff};Yellow =_d .RGBA {0xff,0xff,0x00,0xff};Yellowgreen =_d .RGBA {0x9a,0xcd,0x32,0xff};
);var Names =[]string {"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e","\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065","\u0061\u0071\u0075\u0061","\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065","\u0061\u007a\u0075r\u0065","\u0062\u0065\u0069g\u0065","\u0062\u0069\u0073\u0071\u0075\u0065","\u0062\u006c\u0061c\u006b","\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064","\u0062\u006c\u0075\u0065","\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074","\u0062\u0072\u006fw\u006e","\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd","\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e","\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065","\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e","\u0063\u006f\u0072a\u006c","\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065","\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b","\u0063r\u0069\u006d\u0073\u006f\u006e","\u0063\u0079\u0061\u006e","\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065","\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e","\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064","\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079","\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n","\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079","\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi","d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061","\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e","\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065","\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064","\u0064a\u0072\u006b\u0072\u0065\u0064","\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e","\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e","\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065","\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079","\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079","\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065","\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074","\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b","d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065","\u0064i\u006d\u0067\u0072\u0061\u0079","\u0064i\u006d\u0067\u0072\u0065\u0079","\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065","\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k","f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065","f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e","\u0066u\u0063\u0068\u0073\u0069\u0061","\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o","\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065","\u0067\u006f\u006c\u0064","\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd","\u0067\u0072\u0061\u0079","\u0067\u0072\u0065e\u006e","g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077","\u0067\u0072\u0065\u0079","\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077","\u0068o\u0074\u0070\u0069\u006e\u006b","\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d","\u0069\u006e\u0064\u0069\u0067\u006f","\u0069\u0076\u006fr\u0079","\u006b\u0068\u0061k\u0069","\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072","\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068","\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n","\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e","\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e","\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c","\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n","l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077","\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y","\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e","\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y","\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek","l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e","\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e","\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065","\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079","\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079","\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065","l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077","\u006c\u0069\u006d\u0065","\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n","\u006c\u0069\u006ee\u006e","\u006da\u0067\u0065\u006e\u0074\u0061","\u006d\u0061\u0072\u006f\u006f\u006e","\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065","\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065","\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064","\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065","\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e","\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065","\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e","\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065","\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064","\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065","\u006di\u006e\u0074\u0063\u0072\u0065\u0061m","\u006di\u0073\u0074\u0079\u0072\u006f\u0073e","\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e","n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065","\u006e\u0061\u0076\u0079","\u006fl\u0064\u006c\u0061\u0063\u0065","\u006f\u006c\u0069v\u0065","\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b","\u006f\u0072\u0061\u006e\u0067\u0065","\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d","\u006f\u0072\u0063\u0068\u0069\u0064","\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064","\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n","\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065","\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064","\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070","\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f","\u0070\u0065\u0072\u0075","\u0070\u0069\u006e\u006b","\u0070\u006c\u0075\u006d","\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065","\u0070\u0075\u0072\u0070\u006c\u0065","\u0072\u0065\u0064","\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n","\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e","s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e","\u0073\u0061\u006c\u006d\u006f\u006e","\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e","\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e","\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c","\u0073\u0069\u0065\u006e\u006e\u0061","\u0073\u0069\u006c\u0076\u0065\u0072","\u0073k\u0079\u0062\u006c\u0075\u0065","\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e","\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y","\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y","\u0073\u006e\u006f\u0077","s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e","\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e","\u0074\u0061\u006e","\u0074\u0065\u0061\u006c","\u0074h\u0069\u0073\u0074\u006c\u0065","\u0074\u006f\u006d\u0061\u0074\u006f","\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e","\u0076\u0069\u006f\u006c\u0065\u0074","\u0077\u0068\u0065a\u0074","\u0077\u0068\u0069t\u0065","\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065","\u0079\u0065\u006c\u006c\u006f\u0077","y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e"};
func EllipseToCubicBeziers (startX ,startY ,rx ,ry ,rot float64 ,large ,sweep bool ,endX ,endY float64 )[][4]Point {rx =_g .Abs (rx );ry =_g .Abs (ry );if rx < ry {rx ,ry =ry ,rx ;rot +=90.0;};_e :=_ced (rot *_g .Pi /180.0);if _g .Pi <=_e {_e -=_g .Pi ;
};_a ,_de ,_eg ,_fg :=_gdf (startX ,startY ,rx ,ry ,_e ,large ,sweep ,endX ,endY );_eb :=_g .Pi /2.0;_b :=int (_g .Ceil (_g .Abs (_fg -_eg )/_eb ));_eb =_g .Abs (_fg -_eg )/float64 (_b );_gd :=_g .Sin (_eb )*(_g .Sqrt (4.0+3.0*_g .Pow (_g .Tan (_eb /2.0),2.0))-1.0)/3.0;
if !sweep {_eb =-_eb ;};_ab :=Point {X :startX ,Y :startY };_bf ,_fe :=_abb (rx ,ry ,_e ,sweep ,_eg );_ad :=Point {X :_bf ,Y :_fe };_ac :=[][4]Point {};for _da :=1;_da < _b +1;_da ++{_c :=_eg +float64 (_da )*_eb ;_acd ,_bd :=_bc (rx ,ry ,_e ,_a ,_de ,_c );
_fa :=Point {X :_acd ,Y :_bd };_fd ,_cd :=_abb (rx ,ry ,_e ,sweep ,_c );_ec :=Point {X :_fd ,Y :_cd };_ae :=_ab .Add (_ad .Mul (_gd ));_bg :=_fa .Sub (_ec .Mul (_gd ));_ac =append (_ac ,[4]Point {_ab ,_ae ,_bg ,_fa });_ad =_ec ;_ab =_fa ;};return _ac ;
};var ColorMap =map[string ]_d .RGBA {"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e":_d .RGBA {0xf0,0xf8,0xff,0xff},"\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065":_d .RGBA {0xfa,0xeb,0xd7,0xff},"\u0061\u0071\u0075\u0061":_d .RGBA {0x00,0xff,0xff,0xff},"\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065":_d .RGBA {0x7f,0xff,0xd4,0xff},"\u0061\u007a\u0075r\u0065":_d .RGBA {0xf0,0xff,0xff,0xff},"\u0062\u0065\u0069g\u0065":_d .RGBA {0xf5,0xf5,0xdc,0xff},"\u0062\u0069\u0073\u0071\u0075\u0065":_d .RGBA {0xff,0xe4,0xc4,0xff},"\u0062\u006c\u0061c\u006b":_d .RGBA {0x00,0x00,0x00,0xff},"\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064":_d .RGBA {0xff,0xeb,0xcd,0xff},"\u0062\u006c\u0075\u0065":_d .RGBA {0x00,0x00,0xff,0xff},"\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074":_d .RGBA {0x8a,0x2b,0xe2,0xff},"\u0062\u0072\u006fw\u006e":_d .RGBA {0xa5,0x2a,0x2a,0xff},"\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd":_d .RGBA {0xde,0xb8,0x87,0xff},"\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e":_d .RGBA {0x5f,0x9e,0xa0,0xff},"\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065":_d .RGBA {0x7f,0xff,0x00,0xff},"\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e":_d .RGBA {0xd2,0x69,0x1e,0xff},"\u0063\u006f\u0072a\u006c":_d .RGBA {0xff,0x7f,0x50,0xff},"\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065":_d .RGBA {0x64,0x95,0xed,0xff},"\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b":_d .RGBA {0xff,0xf8,0xdc,0xff},"\u0063r\u0069\u006d\u0073\u006f\u006e":_d .RGBA {0xdc,0x14,0x3c,0xff},"\u0063\u0079\u0061\u006e":_d .RGBA {0x00,0xff,0xff,0xff},"\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065":_d .RGBA {0x00,0x00,0x8b,0xff},"\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e":_d .RGBA {0x00,0x8b,0x8b,0xff},"\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064":_d .RGBA {0xb8,0x86,0x0b,0xff},"\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079":_d .RGBA {0xa9,0xa9,0xa9,0xff},"\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n":_d .RGBA {0x00,0x64,0x00,0xff},"\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079":_d .RGBA {0xa9,0xa9,0xa9,0xff},"\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi":_d .RGBA {0xbd,0xb7,0x6b,0xff},"d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061":_d .RGBA {0x8b,0x00,0x8b,0xff},"\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e":_d .RGBA {0x55,0x6b,0x2f,0xff},"\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065":_d .RGBA {0xff,0x8c,0x00,0xff},"\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064":_d .RGBA {0x99,0x32,0xcc,0xff},"\u0064a\u0072\u006b\u0072\u0065\u0064":_d .RGBA {0x8b,0x00,0x00,0xff},"\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e":_d .RGBA {0xe9,0x96,0x7a,0xff},"\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e":_d .RGBA {0x8f,0xbc,0x8f,0xff},"\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065":_d .RGBA {0x48,0x3d,0x8b,0xff},"\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079":_d .RGBA {0x2f,0x4f,0x4f,0xff},"\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079":_d .RGBA {0x2f,0x4f,0x4f,0xff},"\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065":_d .RGBA {0x00,0xce,0xd1,0xff},"\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074":_d .RGBA {0x94,0x00,0xd3,0xff},"\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b":_d .RGBA {0xff,0x14,0x93,0xff},"d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065":_d .RGBA {0x00,0xbf,0xff,0xff},"\u0064i\u006d\u0067\u0072\u0061\u0079":_d .RGBA {0x69,0x69,0x69,0xff},"\u0064i\u006d\u0067\u0072\u0065\u0079":_d .RGBA {0x69,0x69,0x69,0xff},"\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065":_d .RGBA {0x1e,0x90,0xff,0xff},"\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k":_d .RGBA {0xb2,0x22,0x22,0xff},"f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065":_d .RGBA {0xff,0xfa,0xf0,0xff},"f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e":_d .RGBA {0x22,0x8b,0x22,0xff},"\u0066u\u0063\u0068\u0073\u0069\u0061":_d .RGBA {0xff,0x00,0xff,0xff},"\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o":_d .RGBA {0xdc,0xdc,0xdc,0xff},"\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065":_d .RGBA {0xf8,0xf8,0xff,0xff},"\u0067\u006f\u006c\u0064":_d .RGBA {0xff,0xd7,0x00,0xff},"\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd":_d .RGBA {0xda,0xa5,0x20,0xff},"\u0067\u0072\u0061\u0079":_d .RGBA {0x80,0x80,0x80,0xff},"\u0067\u0072\u0065e\u006e":_d .RGBA {0x00,0x80,0x00,0xff},"g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077":_d .RGBA {0xad,0xff,0x2f,0xff},"\u0067\u0072\u0065\u0079":_d .RGBA {0x80,0x80,0x80,0xff},"\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077":_d .RGBA {0xf0,0xff,0xf0,0xff},"\u0068o\u0074\u0070\u0069\u006e\u006b":_d .RGBA {0xff,0x69,0xb4,0xff},"\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d":_d .RGBA {0xcd,0x5c,0x5c,0xff},"\u0069\u006e\u0064\u0069\u0067\u006f":_d .RGBA {0x4b,0x00,0x82,0xff},"\u0069\u0076\u006fr\u0079":_d .RGBA {0xff,0xff,0xf0,0xff},"\u006b\u0068\u0061k\u0069":_d .RGBA {0xf0,0xe6,0x8c,0xff},"\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072":_d .RGBA {0xe6,0xe6,0xfa,0xff},"\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068":_d .RGBA {0xff,0xf0,0xf5,0xff},"\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n":_d .RGBA {0x7c,0xfc,0x00,0xff},"\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e":_d .RGBA {0xff,0xfa,0xcd,0xff},"\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e":_d .RGBA {0xad,0xd8,0xe6,0xff},"\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c":_d .RGBA {0xf0,0x80,0x80,0xff},"\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n":_d .RGBA {0xe0,0xff,0xff,0xff},"l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077":_d .RGBA {0xfa,0xfa,0xd2,0xff},"\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y":_d .RGBA {0xd3,0xd3,0xd3,0xff},"\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e":_d .RGBA {0x90,0xee,0x90,0xff},"\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y":_d .RGBA {0xd3,0xd3,0xd3,0xff},"\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek":_d .RGBA {0xff,0xb6,0xc1,0xff},"l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e":_d .RGBA {0xff,0xa0,0x7a,0xff},"\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e":_d .RGBA {0x20,0xb2,0xaa,0xff},"\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065":_d .RGBA {0x87,0xce,0xfa,0xff},"\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079":_d .RGBA {0x77,0x88,0x99,0xff},"\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079":_d .RGBA {0x77,0x88,0x99,0xff},"\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065":_d .RGBA {0xb0,0xc4,0xde,0xff},"l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077":_d .RGBA {0xff,0xff,0xe0,0xff},"\u006c\u0069\u006d\u0065":_d .RGBA {0x00,0xff,0x00,0xff},"\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n":_d .RGBA {0x32,0xcd,0x32,0xff},"\u006c\u0069\u006ee\u006e":_d .RGBA {0xfa,0xf0,0xe6,0xff},"\u006da\u0067\u0065\u006e\u0074\u0061":_d .RGBA {0xff,0x00,0xff,0xff},"\u006d\u0061\u0072\u006f\u006f\u006e":_d .RGBA {0x80,0x00,0x00,0xff},"\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065":_d .RGBA {0x66,0xcd,0xaa,0xff},"\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065":_d .RGBA {0x00,0x00,0xcd,0xff},"\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064":_d .RGBA {0xba,0x55,0xd3,0xff},"\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065":_d .RGBA {0x93,0x70,0xdb,0xff},"\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e":_d .RGBA {0x3c,0xb3,0x71,0xff},"\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065":_d .RGBA {0x7b,0x68,0xee,0xff},"\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e":_d .RGBA {0x00,0xfa,0x9a,0xff},"\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065":_d .RGBA {0x48,0xd1,0xcc,0xff},"\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064":_d .RGBA {0xc7,0x15,0x85,0xff},"\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065":_d .RGBA {0x19,0x19,0x70,0xff},"\u006di\u006e\u0074\u0063\u0072\u0065\u0061m":_d .RGBA {0xf5,0xff,0xfa,0xff},"\u006di\u0073\u0074\u0079\u0072\u006f\u0073e":_d .RGBA {0xff,0xe4,0xe1,0xff},"\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e":_d .RGBA {0xff,0xe4,0xb5,0xff},"n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065":_d .RGBA {0xff,0xde,0xad,0xff},"\u006e\u0061\u0076\u0079":_d .RGBA {0x00,0x00,0x80,0xff},"\u006fl\u0064\u006c\u0061\u0063\u0065":_d .RGBA {0xfd,0xf5,0xe6,0xff},"\u006f\u006c\u0069v\u0065":_d .RGBA {0x80,0x80,0x00,0xff},"\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b":_d .RGBA {0x6b,0x8e,0x23,0xff},"\u006f\u0072\u0061\u006e\u0067\u0065":_d .RGBA {0xff,0xa5,0x00,0xff},"\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d":_d .RGBA {0xff,0x45,0x00,0xff},"\u006f\u0072\u0063\u0068\u0069\u0064":_d .RGBA {0xda,0x70,0xd6,0xff},"\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064":_d .RGBA {0xee,0xe8,0xaa,0xff},"\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n":_d .RGBA {0x98,0xfb,0x98,0xff},"\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065":_d .RGBA {0xaf,0xee,0xee,0xff},"\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064":_d .RGBA {0xdb,0x70,0x93,0xff},"\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070":_d .RGBA {0xff,0xef,0xd5,0xff},"\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f":_d .RGBA {0xff,0xda,0xb9,0xff},"\u0070\u0065\u0072\u0075":_d .RGBA {0xcd,0x85,0x3f,0xff},"\u0070\u0069\u006e\u006b":_d .RGBA {0xff,0xc0,0xcb,0xff},"\u0070\u006c\u0075\u006d":_d .RGBA {0xdd,0xa0,0xdd,0xff},"\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065":_d .RGBA {0xb0,0xe0,0xe6,0xff},"\u0070\u0075\u0072\u0070\u006c\u0065":_d .RGBA {0x80,0x00,0x80,0xff},"\u0072\u0065\u0064":_d .RGBA {0xff,0x00,0x00,0xff},"\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n":_d .RGBA {0xbc,0x8f,0x8f,0xff},"\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e":_d .RGBA {0x41,0x69,0xe1,0xff},"s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e":_d .RGBA {0x8b,0x45,0x13,0xff},"\u0073\u0061\u006c\u006d\u006f\u006e":_d .RGBA {0xfa,0x80,0x72,0xff},"\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e":_d .RGBA {0xf4,0xa4,0x60,0xff},"\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e":_d .RGBA {0x2e,0x8b,0x57,0xff},"\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c":_d .RGBA {0xff,0xf5,0xee,0xff},"\u0073\u0069\u0065\u006e\u006e\u0061":_d .RGBA {0xa0,0x52,0x2d,0xff},"\u0073\u0069\u006c\u0076\u0065\u0072":_d .RGBA {0xc0,0xc0,0xc0,0xff},"\u0073k\u0079\u0062\u006c\u0075\u0065":_d .RGBA {0x87,0xce,0xeb,0xff},"\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e":_d .RGBA {0x6a,0x5a,0xcd,0xff},"\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y":_d .RGBA {0x70,0x80,0x90,0xff},"\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y":_d .RGBA {0x70,0x80,0x90,0xff},"\u0073\u006e\u006f\u0077":_d .RGBA {0xff,0xfa,0xfa,0xff},"s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e":_d .RGBA {0x00,0xff,0x7f,0xff},"\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e":_d .RGBA {0x46,0x82,0xb4,0xff},"\u0074\u0061\u006e":_d .RGBA {0xd2,0xb4,0x8c,0xff},"\u0074\u0065\u0061\u006c":_d .RGBA {0x00,0x80,0x80,0xff},"\u0074h\u0069\u0073\u0074\u006c\u0065":_d .RGBA {0xd8,0xbf,0xd8,0xff},"\u0074\u006f\u006d\u0061\u0074\u006f":_d .RGBA {0xff,0x63,0x47,0xff},"\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e":_d .RGBA {0x40,0xe0,0xd0,0xff},"\u0076\u0069\u006f\u006c\u0065\u0074":_d .RGBA {0xee,0x82,0xee,0xff},"\u0077\u0068\u0065a\u0074":_d .RGBA {0xf5,0xde,0xb3,0xff},"\u0077\u0068\u0069t\u0065":_d .RGBA {0xff,0xff,0xff,0xff},"\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065":_d .RGBA {0xf5,0xf5,0xf5,0xff},"\u0079\u0065\u006c\u006c\u006f\u0077":_d .RGBA {0xff,0xff,0x00,0xff},"y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e":_d .RGBA {0x9a,0xcd,0x32,0xff}};
func _ced (_fbf float64 )float64 {_fbf =_g .Mod (_fbf ,2.0*_g .Pi );if _fbf < 0.0{_fbf +=2.0*_g .Pi ;};return _fbf ;};func (_deg Point )Mul (f float64 )Point {return Point {f *_deg .X ,f *_deg .Y }};func (_acg Point )Sub (q Point )Point {return Point {_acg .X -q .X ,_acg .Y -q .Y }};
func _abb (_abe ,_ce ,_dg float64 ,_fea bool ,_daa float64 )(float64 ,float64 ){_aea ,_cdf :=_g .Sincos (_daa );_gded ,_ecc :=_g .Sincos (_dg );_cc :=-_abe *_aea *_ecc -_ce *_cdf *_gded ;_ceb :=-_abe *_aea *_gded +_ce *_cdf *_ecc ;if !_fea {return -_cc ,-_ceb ;
};return _cc ,_ceb ;};