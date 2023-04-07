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

package textshaping ;import (_c "github.com/unidoc/garabic";_g "golang.org/x/text/unicode/bidi";_b "strings";);

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape (text string )(string ,error ){_f :=_g .Paragraph {};_f .SetString (text );_bg ,_cc :=_f .Order ();if _cc !=nil {return "",_cc ;};for _fd :=0;_fd < _bg .NumRuns ();_fd ++{_a :=_bg .Run (_fd );_d :=_a .String ();if _a .Direction ()==_g .RightToLeft {var (_de =_c .Shape (_d );
_ea =[]rune (_de );_ge =make ([]rune ,len (_ea )););_cg :=0;for _ef :=len (_ea )-1;_ef >=0;_ef --{_ge [_cg ]=_ea [_ef ];_cg ++;};_d =string (_ge );text =_b .Replace (text ,_b .TrimSpace (_a .String ()),_d ,1);};};return text ,nil ;};