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

package mdp ;import (_d "errors";_ge "fmt";_dc "github.com/unidoc/unipdf/v3/core";);func (_gd *defaultDiffPolicy )compareAnnots (_dgb int ,_egf ,_dfb []_dc .PdfObject )error {_fbc :=make (map[int64 ]*_dc .PdfObjectDictionary );for _ ,_cba :=range _egf {_ffe ,_deb :=_dc .GetIndirect (_cba );
if !_deb {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_cgcd ,_deb :=_dc .GetDict (_ffe .PdfObject );if !_deb {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");
};_fbc [_ffe .ObjectNumber ]=_cgcd ;};for _ ,_dee :=range _dfb {_ed ,_gff :=_dc .GetIndirect (_dee );if !_gff {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");
};_ddb ,_gff :=_dc .GetDict (_ed .PdfObject );if !_gff {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_cee ,_ :=_dc .GetStringVal (_ddb .Get ("\u0054"));
_ecb ,_ :=_dc .GetNameVal (_ddb .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));if _ ,_cdg :=_fbc [_ed .ObjectNumber ];!_cdg {switch _gd ._c {case NoRestrictions ,FillFormsAndAnnots :_gd ._ab .addWarningWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064",_ecb ,_cee ));
default:_ebg ,_ggg :=_dc .GetDict (_ed .PdfObject );if !_ggg {return _d .New ("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e");};_dfeg ,_ggg :=_dc .GetNameVal (_ebg .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
if !_ggg {return _d .New ("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065");};if _dfeg =="\u0057\u0069\u0064\u0067\u0065\u0074"{switch _gd ._c {case NoRestrictions ,FillFormsAndAnnots ,FillForms :_gd ._ab .addWarningWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064",_ecb ,_cee ));
default:_gd ._ab .addErrorWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064",_ecb ,_cee ));};}else {_gd ._ab .addErrorWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064",_ecb ,_cee ));
};};}else {delete (_fbc ,_ed .ObjectNumber );if _ceg ,_fab :=_gd ._a [_ed .ObjectNumber ];_fab {switch _gd ._c {case NoRestrictions ,FillFormsAndAnnots :_gd ._ab .addWarningWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064",_ecb ,_cee ));
default:_dfed ,_dfc :=_dc .GetIndirect (_ceg );if !_dfc {return _d .New ("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e");};_dgbb ,_dfc :=_dc .GetDict (_dfed .PdfObject );if !_dfc {return _d .New ("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e");
};_ca ,_dfc :=_dc .GetNameVal (_dgbb .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));if !_dfc {return _d .New ("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065");
};if _ca =="\u0057\u0069\u0064\u0067\u0065\u0074"{switch _gd ._c {case NoRestrictions ,FillFormsAndAnnots ,FillForms :_gd ._ab .addWarningWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064",_ecb ,_cee ));
default:_gd ._ab .addErrorWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064",_ecb ,_cee ));};}else {_gd ._ab .addErrorWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064",_ecb ,_cee ));
};};};};};for _ ,_gbaa :=range _fbc {_be ,_ :=_dc .GetStringVal (_gbaa .Get ("\u0054"));_ddf ,_ :=_dc .GetNameVal (_gbaa .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));switch _gd ._c {case NoRestrictions ,FillFormsAndAnnots :_gd ._ab .addWarningWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064",_ddf ,_be ));
default:_gd ._ab .addErrorWithDescription (_dgb ,_ge .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064",_ddf ,_be ));};};return nil ;};func _acbd (_fbgb _dc .PdfObject )([]_dc .PdfObject ,error ){_gaec :=make ([]_dc .PdfObject ,0);
if _fbgb !=nil {_fd :=_fbgb ;if _ecg ,_cge :=_dc .GetIndirect (_fbgb );_cge {_fd =_ecg .PdfObject ;};if _dfa ,_gagc :=_dc .GetArray (_fd );_gagc {_gaec =_dfa .Elements ();}else {return nil ,_d .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0061n\u006eo\u0074s\u0027\u0020\u006f\u0062\u006a\u0065\u0063t");
};};return _gaec ,nil ;};

// DiffPolicy interface for comparing two revisions of the Pdf document.
type DiffPolicy interface{

// ReviewFile should check the revisions of the old and new parsers
// and evaluate the differences between the revisions.
// Each implementation of this interface must decide
// how to handle cases where there are multiple revisions between the old and new revisions.
ReviewFile (_bga *_dc .PdfParser ,_dff *_dc .PdfParser ,_agf *MDPParameters )(*DiffResults ,error );};func (_afgd *DiffResults )addWarning (_dce *DiffResult ){if _afgd .Warnings ==nil {_afgd .Warnings =make ([]*DiffResult ,0);};_afgd .Warnings =append (_afgd .Warnings ,_dce );
};func NewDefaultDiffPolicy ()DiffPolicy {return &defaultDiffPolicy {_a :nil ,_ab :&DiffResults {},_c :0};};

// String returns the state of the warning.
func (_ceb *DiffResult )String ()string {return _ge .Sprintf ("\u0025\u0073\u0020\u0069n \u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0023\u0025\u0064",_ceb .Description ,_ceb .Revision );};func (_abd *defaultDiffPolicy )compareFields (_cfe int ,_abb ,_ec []_dc .PdfObject )error {_cgc :=make (map[int64 ]*_dc .PdfObjectDictionary );
for _ ,_ba :=range _abb {_abe ,_egg :=_dc .GetIndirect (_ba );if !_egg {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_ef ,_egg :=_dc .GetDict (_abe .PdfObject );
if !_egg {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_cgc [_abe .ObjectNumber ]=_ef ;};for _ ,_dd :=range _ec {_cd ,_aba :=_dc .GetIndirect (_dd );
if !_aba {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_dad ,_aba :=_dc .GetDict (_cd .PdfObject );if !_aba {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");
};T :=_dad .Get ("\u0054");if _ ,_ggd :=_abd ._a [_cd .ObjectNumber ];_ggd {switch _abd ._c {case NoRestrictions ,FillForms ,FillFormsAndAnnots :_abd ._ab .addWarningWithDescription (_cfe ,_ge .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064",T ));
default:_abd ._ab .addErrorWithDescription (_cfe ,_ge .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064",T ));};};if _ ,_cb :=_cgc [_cd .ObjectNumber ];!_cb {switch _abd ._c {case NoRestrictions ,FillForms ,FillFormsAndAnnots :_abd ._ab .addWarningWithDescription (_cfe ,_ge .Sprintf ("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064",_dad .Get ("\u0054")));
default:_abd ._ab .addErrorWithDescription (_cfe ,_ge .Sprintf ("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064",_dad .Get ("\u0054")));};}else {delete (_cgc ,_cd .ObjectNumber );if _ ,_eef :=_abd ._a [_cd .ObjectNumber ];
_eef {switch _abd ._c {case NoRestrictions ,FillForms ,FillFormsAndAnnots :_abd ._ab .addWarningWithDescription (_cfe ,_ge .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064",_dad .Get ("\u0054")));
default:_abd ._ab .addErrorWithDescription (_cfe ,_ge .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064",_dad .Get ("\u0054")));};};};if FT ,_egd :=_dc .GetNameVal (_dad .Get ("\u0046\u0054"));_egd {if FT =="\u0053\u0069\u0067"{if _fcc ,_fbg :=_dc .GetIndirect (_dad .Get ("\u0056"));
_fbg {if _ ,_gc :=_abd ._a [_fcc .ObjectNumber ];_gc {switch _abd ._c {case NoRestrictions ,FillForms ,FillFormsAndAnnots :_abd ._ab .addWarningWithDescription (_cfe ,_ge .Sprintf ("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064",T ));
default:_abd ._ab .addErrorWithDescription (_cfe ,_ge .Sprintf ("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064",T ));};};};
};};};for _ ,_bb :=range _cgc {switch _abd ._c {case NoRestrictions :_abd ._ab .addWarningWithDescription (_cfe ,_ge .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064",_bb .Get ("\u0054")));default:_abd ._ab .addErrorWithDescription (_cfe ,_ge .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064",_bb .Get ("\u0054")));
};};return nil ;};

// DiffResult describes the warning or the error for the DiffPolicy results.
type DiffResult struct{Revision int ;Description string ;};

// MDPParameters describes parameters for the MDP checks (now only DocMDP).
type MDPParameters struct{DocMDPLevel DocMDPPermission ;};const (NoRestrictions DocMDPPermission =0;NoChanges DocMDPPermission =1;FillForms DocMDPPermission =2;FillFormsAndAnnots DocMDPPermission =3;);func (_aef *DiffResults )addErrorWithDescription (_gfd int ,_fcf string ){if _aef .Errors ==nil {_aef .Errors =make ([]*DiffResult ,0);
};_aef .Errors =append (_aef .Errors ,&DiffResult {Revision :_gfd ,Description :_fcf });};func (_fc *defaultDiffPolicy )compareRevisions (_abc *_dc .PdfParser ,_bg *_dc .PdfParser )(*DiffResults ,error ){var _cf error ;_fc ._a ,_cf =_bg .GetUpdatedObjects (_abc );
if _cf !=nil {return &DiffResults {},_cf ;};if len (_fc ._a )==0{return &DiffResults {},nil ;};_ff :=_bg .GetRevisionNumber ();_ged ,_cg :=_dc .GetIndirect (_dc .ResolveReference (_abc .GetTrailer ().Get ("\u0052\u006f\u006f\u0074")));_gg ,_ee :=_dc .GetIndirect (_dc .ResolveReference (_bg .GetTrailer ().Get ("\u0052\u006f\u006f\u0074")));
if !_cg ||!_ee {return &DiffResults {},_d .New ("\u0065\u0072\u0072o\u0072\u0020\u0077\u0068i\u006c\u0065\u0020\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0072\u006f\u006f\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074");};_ac ,_cg :=_dc .GetDict (_dc .ResolveReference (_ged .PdfObject ));
_ae ,_ee :=_dc .GetDict (_dc .ResolveReference (_gg .PdfObject ));if !_cg ||!_ee {return &DiffResults {},_d .New ("\u0065\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020a\u0020\u0072\u006f\u006f\u0074\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");
};if _bde ,_ga :=_dc .GetIndirect (_ae .Get ("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"));_ga {_gae ,_gag :=_dc .GetDict (_bde );if !_gag {return &DiffResults {},_d .New ("\u0065\u0072\u0072\u006f\u0072 \u0077\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};_cfb :=make ([]_dc .PdfObject ,0);if _gf ,_gba :=_dc .GetIndirect (_ac .Get ("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"));_gba {if _eeb ,_db :=_dc .GetDict (_gf );_db {if _bc ,_eg :=_dc .GetArray (_eeb .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));
_eg {_cfb =_bc .Elements ();};};};_da ,_gag :=_dc .GetArray (_gae .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_gag {return &DiffResults {},_d .New ("\u0065\u0072r\u006f\u0072\u0020\u0077h\u0069\u006ce\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0066i\u0065\u006c\u0064\u0073");
};if _fb :=_fc .compareFields (_ff ,_cfb ,_da .Elements ());_fb !=nil {return &DiffResults {},_fb ;};};_ea ,_geda :=_dc .GetIndirect (_ae .Get ("\u0050\u0061\u0067e\u0073"));if !_geda {return &DiffResults {},_d .New ("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074");
};_cfg ,_geda :=_dc .GetIndirect (_ac .Get ("\u0050\u0061\u0067e\u0073"));if !_geda {return &DiffResults {},_d .New ("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074");
};if _acb :=_fc .comparePages (_ff ,_cfg ,_ea );_acb !=nil {return &DiffResults {},_acb ;};return _fc ._ab ,nil ;};

// DocMDPPermission is values for set up access permissions for DocMDP.
// (Section 12.8.2.2, Table 254 - Entries in a signature dictionary p. 471 in PDF32000_2008).
type DocMDPPermission int64 ;

// ReviewFile implementation of DiffPolicy interface
// The default policy only checks the next types of objects:
// Page, Pages (container for page objects), Annot, Annots (container for annotation objects), Field.
// It checks adding, removing and modifying objects of these types.
func (_f *defaultDiffPolicy )ReviewFile (oldParser *_dc .PdfParser ,newParser *_dc .PdfParser ,params *MDPParameters )(*DiffResults ,error ){if oldParser .GetRevisionNumber ()> newParser .GetRevisionNumber (){return nil ,_d .New ("\u006f\u006c\u0064\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u006e\u0065\u0077\u0020r\u0065\u0076i\u0073\u0069\u006f\u006e");
};if oldParser .GetRevisionNumber ()==newParser .GetRevisionNumber (){if oldParser !=newParser {return nil ,_d .New ("\u0073\u0061m\u0065\u0020\u0072\u0065v\u0069\u0073i\u006f\u006e\u0073\u002c\u0020\u0062\u0075\u0074 \u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0070\u0061r\u0073\u0065\u0072\u0073");
};return &DiffResults {},nil ;};if params ==nil {_f ._c =NoRestrictions ;}else {_f ._c =params .DocMDPLevel ;};_df :=&DiffResults {};for _gb :=oldParser .GetRevisionNumber ()+1;_gb <=newParser .GetRevisionNumber ();_gb ++{_b ,_bd :=newParser .GetRevision (_gb -1);
if _bd !=nil {return nil ,_bd ;};_dfe ,_bd :=newParser .GetRevision (_gb );if _bd !=nil {return nil ,_bd ;};_e ,_bd :=_f .compareRevisions (_b ,_dfe );if _bd !=nil {return nil ,_bd ;};_df .Warnings =append (_df .Warnings ,_e .Warnings ...);_df .Errors =append (_df .Errors ,_e .Errors ...);
};return _df ,nil ;};func (_ag *DiffResults )addWarningWithDescription (_dbg int ,_ega string ){if _ag .Warnings ==nil {_ag .Warnings =make ([]*DiffResult ,0);};_ag .Warnings =append (_ag .Warnings ,&DiffResult {Revision :_dbg ,Description :_ega });};

// IsPermitted returns true if changes permitted.
func (_dab *DiffResults )IsPermitted ()bool {return len (_dab .Errors )==0};func (_bf *defaultDiffPolicy )comparePages (_fbd int ,_efe ,_ce *_dc .PdfIndirectObject )error {if _ ,_eb :=_bf ._a [_ce .ObjectNumber ];_eb {_bf ._ab .addErrorWithDescription (_fbd ,"\u0050a\u0067e\u0073\u0020\u0077\u0065\u0072e\u0020\u0063h\u0061\u006e\u0067\u0065\u0064");
};_af ,_fa :=_dc .GetDict (_ce .PdfObject );_ddd ,_bgc :=_dc .GetDict (_efe .PdfObject );if !_fa ||!_bgc {return _d .New ("\u0075n\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061g\u0065\u0073\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074");
};_fg ,_fa :=_dc .GetArray (_af .Get ("\u004b\u0069\u0064\u0073"));_faf ,_bgc :=_dc .GetArray (_ddd .Get ("\u004b\u0069\u0064\u0073"));if !_fa ||!_bgc {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061\u0067\u0065s\u0027 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079");
};_dcd :=_fg .Len ();if _dcd > _faf .Len (){_dcd =_faf .Len ();};for _de :=0;_de < _dcd ;_de ++{_gec ,_gcg :=_dc .GetIndirect (_dc .ResolveReference (_faf .Get (_de )));_ced ,_bgg :=_dc .GetIndirect (_dc .ResolveReference (_fg .Get (_de )));if !_gcg ||!_bgg {return _d .New ("\u0075\u006e\u0065\u0078pe\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065c\u0074");
};if _gec .ObjectNumber !=_ced .ObjectNumber {_bf ._ab .addErrorWithDescription (_fbd ,_ge .Sprintf ("p\u0061\u0067\u0065\u0020#%\u0064 \u0077\u0061\u0073\u0020\u0072e\u0070\u006c\u0061\u0063\u0065\u0064",_de ));};_dcb ,_gcg :=_dc .GetDict (_ced );_ebf ,_bgg :=_dc .GetDict (_gec );
if !_gcg ||!_bgg {return _d .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065'\u0073 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079");};_afa ,_fae :=_acbd (_dcb .Get ("\u0041\u006e\u006e\u006f\u0074\u0073"));
if _fae !=nil {return _fae ;};_aa ,_fae :=_acbd (_ebf .Get ("\u0041\u006e\u006e\u006f\u0074\u0073"));if _fae !=nil {return _fae ;};if _bce :=_bf .compareAnnots (_fbd ,_aa ,_afa );_bce !=nil {return _bce ;};};for _bfg :=_dcd +1;_bfg <=_fg .Len ();_bfg ++{_bf ._ab .addErrorWithDescription (_fbd ,_ge .Sprintf ("\u0070a\u0067e\u0020\u0023\u0025\u0064\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064",_bfg ));
};for _cgg :=_dcd +1;_cgg <=_faf .Len ();_cgg ++{_bf ._ab .addErrorWithDescription (_fbd ,_ge .Sprintf ("p\u0061g\u0065\u0020\u0023\u0025\u0064\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064",_cgg ));};return nil ;};type defaultDiffPolicy struct{_a map[int64 ]_dc .PdfObject ;
_ab *DiffResults ;_c DocMDPPermission ;};

// DiffResults describes the results of the DiffPolicy.
type DiffResults struct{Warnings []*DiffResult ;Errors []*DiffResult ;};func (_afg *DiffResults )addError (_fge *DiffResult ){if _afg .Errors ==nil {_afg .Errors =make ([]*DiffResult ,0);};_afg .Errors =append (_afg .Errors ,_fge );};