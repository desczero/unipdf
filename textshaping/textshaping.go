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

package textshaping ;import (_a "github.com/unidoc/garabic";_c "golang.org/x/text/unicode/bidi";_bc "strings";);

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape (text string )(string ,error ){_e :=_c .Paragraph {};_e .SetString (text );_ee ,_g :=_e .Order ();if _g !=nil {return "",_g ;};for _ae :=0;_ae < _ee .NumRuns ();_ae ++{_d :=_ee .Run (_ae );_af :=_d .String ();if _d .Direction ()==_c .RightToLeft {var (_cd =_a .Shape (_af );
_bb =[]rune (_cd );_ab =make ([]rune ,len (_bb )););_cg :=0;for _f :=len (_bb )-1;_f >=0;_f --{_ab [_cg ]=_bb [_f ];_cg ++;};_af =string (_ab );text =_bc .Replace (text ,_bc .TrimSpace (_d .String ()),_af ,1);};};return text ,nil ;};