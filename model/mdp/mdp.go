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

package mdp ;import (_b "errors";_g "fmt";_a "github.com/unidoc/unipdf/v3/core";);

// DiffPolicy interface for comparing two revisions of the Pdf document.
type DiffPolicy interface{

// ReviewFile should check the revisions of the old and new parsers
// and evaluate the differences between the revisions.
// Each implementation of this interface must decide
// how to handle cases where there are multiple revisions between the old and new revisions.
ReviewFile (_dece *_a .PdfParser ,_gf *_a .PdfParser ,_edeee *MDPParameters )(*DiffResults ,error );};func (_bcb *DiffResults )addError (_eba *DiffResult ){if _bcb .Errors ==nil {_bcb .Errors =make ([]*DiffResult ,0);};_bcb .Errors =append (_bcb .Errors ,_eba );
};

// ReviewFile implementation of DiffPolicy interface
// The default policy only checks the next types of objects:
// Page, Pages (container for page objects), Annot, Annots (container for annotation objects), Field.
// It checks adding, removing and modifying objects of these types.
func (_bc *defaultDiffPolicy )ReviewFile (oldParser *_a .PdfParser ,newParser *_a .PdfParser ,params *MDPParameters )(*DiffResults ,error ){if oldParser .GetRevisionNumber ()> newParser .GetRevisionNumber (){return nil ,_b .New ("\u006f\u006c\u0064\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u006e\u0065\u0077\u0020r\u0065\u0076i\u0073\u0069\u006f\u006e");
};if oldParser .GetRevisionNumber ()==newParser .GetRevisionNumber (){if oldParser !=newParser {return nil ,_b .New ("\u0073\u0061m\u0065\u0020\u0072\u0065v\u0069\u0073i\u006f\u006e\u0073\u002c\u0020\u0062\u0075\u0074 \u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0070\u0061r\u0073\u0065\u0072\u0073");
};return &DiffResults {},nil ;};if params ==nil {_bc ._e =NoRestrictions ;}else {_bc ._e =params .DocMDPLevel ;};_ac :=&DiffResults {};for _fe :=oldParser .GetRevisionNumber ()+1;_fe <=newParser .GetRevisionNumber ();_fe ++{_c ,_dg :=newParser .GetRevision (_fe -1);
if _dg !=nil {return nil ,_dg ;};_cg ,_dg :=newParser .GetRevision (_fe );if _dg !=nil {return nil ,_dg ;};_ea ,_dg :=_bc .compareRevisions (_c ,_cg );if _dg !=nil {return nil ,_dg ;};_ac .Warnings =append (_ac .Warnings ,_ea .Warnings ...);_ac .Errors =append (_ac .Errors ,_ea .Errors ...);
};return _ac ,nil ;};func (_eef *DiffResults )addWarning (_dce *DiffResult ){if _eef .Warnings ==nil {_eef .Warnings =make ([]*DiffResult ,0);};_eef .Warnings =append (_eef .Warnings ,_dce );};

// String returns the state of the warning.
func (_fef *DiffResult )String ()string {return _g .Sprintf ("\u0025\u0073\u0020\u0069n \u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0023\u0025\u0064",_fef .Description ,_fef .Revision );};

// DiffResults describes the results of the DiffPolicy.
type DiffResults struct{Warnings []*DiffResult ;Errors []*DiffResult ;};const (NoRestrictions DocMDPPermission =0;NoChanges DocMDPPermission =1;FillForms DocMDPPermission =2;FillFormsAndAnnots DocMDPPermission =3;);func NewDefaultDiffPolicy ()DiffPolicy {return &defaultDiffPolicy {_d :nil ,_bd :&DiffResults {},_e :0};
};

// DocMDPPermission is values for set up access permissions for DocMDP.
// (Section 12.8.2.2, Table 254 - Entries in a signature dictionary p. 471 in PDF32000_2008).
type DocMDPPermission int64 ;func (_fgg *defaultDiffPolicy )comparePages (_cbf int ,_aa ,_cbg *_a .PdfIndirectObject )error {if _ ,_fb :=_fgg ._d [_cbg .ObjectNumber ];_fb {_fgg ._bd .addErrorWithDescription (_cbf ,"\u0050a\u0067e\u0073\u0020\u0077\u0065\u0072e\u0020\u0063h\u0061\u006e\u0067\u0065\u0064");
};_dgb ,_cgb :=_a .GetDict (_cbg .PdfObject );_da ,_bdbf :=_a .GetDict (_aa .PdfObject );if !_cgb ||!_bdbf {return _b .New ("\u0075n\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061g\u0065\u0073\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074");
};_aef ,_cgb :=_a .GetArray (_dgb .Get ("\u004b\u0069\u0064\u0073"));_ede ,_bdbf :=_a .GetArray (_da .Get ("\u004b\u0069\u0064\u0073"));if !_cgb ||!_bdbf {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0050\u0061\u0067\u0065s\u0027 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079");
};_fgb :=_aef .Len ();if _fgb > _ede .Len (){_fgb =_ede .Len ();};for _fcaf :=0;_fcaf < _fgb ;_fcaf ++{_egb ,_cff :=_a .GetIndirect (_a .ResolveReference (_ede .Get (_fcaf )));_bfd ,_eee :=_a .GetIndirect (_a .ResolveReference (_aef .Get (_fcaf )));if !_cff ||!_eee {return _b .New ("\u0075\u006e\u0065\u0078pe\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065c\u0074");
};if _egb .ObjectNumber !=_bfd .ObjectNumber {_fgg ._bd .addErrorWithDescription (_cbf ,_g .Sprintf ("p\u0061\u0067\u0065\u0020#%\u0064 \u0077\u0061\u0073\u0020\u0072e\u0070\u006c\u0061\u0063\u0065\u0064",_fcaf ));};_ag ,_cff :=_a .GetDict (_bfd );_fcf ,_eee :=_a .GetDict (_egb );
if !_cff ||!_eee {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0067\u0065'\u0073 \u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079");};_edee ,_bda :=_gce (_ag .Get ("\u0041\u006e\u006e\u006f\u0074\u0073"));
if _bda !=nil {return _bda ;};_cag ,_bda :=_gce (_fcf .Get ("\u0041\u006e\u006e\u006f\u0074\u0073"));if _bda !=nil {return _bda ;};if _eb :=_fgg .compareAnnots (_cbf ,_cag ,_edee );_eb !=nil {return _eb ;};};for _fgbe :=_fgb +1;_fgbe <=_aef .Len ();_fgbe ++{_fgg ._bd .addErrorWithDescription (_cbf ,_g .Sprintf ("\u0070a\u0067e\u0020\u0023\u0025\u0064\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064",_fgbe ));
};for _egg :=_fgb +1;_egg <=_ede .Len ();_egg ++{_fgg ._bd .addErrorWithDescription (_cbf ,_g .Sprintf ("p\u0061g\u0065\u0020\u0023\u0025\u0064\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064",_egg ));};return nil ;};func (_ee *defaultDiffPolicy )compareRevisions (_cf *_a .PdfParser ,_bb *_a .PdfParser )(*DiffResults ,error ){var _ba error ;
_ee ._d ,_ba =_bb .GetUpdatedObjects (_cf );if _ba !=nil {return &DiffResults {},_ba ;};if len (_ee ._d )==0{return &DiffResults {},nil ;};_bbe :=_bb .GetRevisionNumber ();_fd ,_de :=_a .GetIndirect (_a .ResolveReference (_cf .GetTrailer ().Get ("\u0052\u006f\u006f\u0074")));
_fdf ,_fc :=_a .GetIndirect (_a .ResolveReference (_bb .GetTrailer ().Get ("\u0052\u006f\u006f\u0074")));if !_de ||!_fc {return &DiffResults {},_b .New ("\u0065\u0072\u0072o\u0072\u0020\u0077\u0068i\u006c\u0065\u0020\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0072\u006f\u006f\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074");
};_be ,_de :=_a .GetDict (_a .ResolveReference (_fd .PdfObject ));_df ,_fc :=_a .GetDict (_a .ResolveReference (_fdf .PdfObject ));if !_de ||!_fc {return &DiffResults {},_b .New ("\u0065\u0072\u0072\u006f\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0067e\u0074\u0074\u0069\u006e\u0067\u0020a\u0020\u0072\u006f\u006f\u0074\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");
};if _dga ,_bdb :=_a .GetIndirect (_df .Get ("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"));_bdb {_baf ,_dc :=_a .GetDict (_dga );if !_dc {return &DiffResults {},_b .New ("\u0065\u0072\u0072\u006f\u0072 \u0077\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};_eg :=make ([]_a .PdfObject ,0);if _db ,_feb :=_a .GetIndirect (_be .Get ("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"));_feb {if _dee ,_cc :=_a .GetDict (_db );_cc {if _fdc ,_bf :=_a .GetArray (_dee .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));
_bf {_eg =_fdc .Elements ();};};};_ef ,_dc :=_a .GetArray (_baf .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_dc {return &DiffResults {},_b .New ("\u0065\u0072r\u006f\u0072\u0020\u0077h\u0069\u006ce\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067 \u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d\u0027\u0073\u0020\u0066i\u0065\u006c\u0064\u0073");
};if _ec :=_ee .compareFields (_bbe ,_eg ,_ef .Elements ());_ec !=nil {return &DiffResults {},_ec ;};};_fg ,_acg :=_a .GetIndirect (_df .Get ("\u0050\u0061\u0067e\u0073"));if !_acg {return &DiffResults {},_b .New ("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074");
};_eae ,_acg :=_a .GetIndirect (_be .Get ("\u0050\u0061\u0067e\u0073"));if !_acg {return &DiffResults {},_b .New ("\u0065\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020p\u0061\u0067\u0065\u0073\u0027\u0020\u006fb\u006a\u0065\u0063\u0074");
};if _ab :=_ee .comparePages (_bbe ,_eae ,_fg );_ab !=nil {return &DiffResults {},_ab ;};return _ee ._bd ,nil ;};func _gce (_gg _a .PdfObject )([]_a .PdfObject ,error ){_ggd :=make ([]_a .PdfObject ,0);if _gg !=nil {_ggb :=_gg ;if _agd ,_gced :=_a .GetIndirect (_gg );
_gced {_ggb =_agd .PdfObject ;};if _edg ,_adg :=_a .GetArray (_ggb );_adg {_ggd =_edg .Elements ();}else {return nil ,_b .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0061n\u006eo\u0074s\u0027\u0020\u006f\u0062\u006a\u0065\u0063t");};
};return _ggd ,nil ;};type defaultDiffPolicy struct{_d map[int64 ]_a .PdfObject ;_bd *DiffResults ;_e DocMDPPermission ;};func (_ebab *DiffResults )addErrorWithDescription (_cae int ,_eafb string ){if _ebab .Errors ==nil {_ebab .Errors =make ([]*DiffResult ,0);
};_ebab .Errors =append (_ebab .Errors ,&DiffResult {Revision :_cae ,Description :_eafb });};func (_fca *defaultDiffPolicy )compareFields (_abg int ,_deeg ,_ccg []_a .PdfObject )error {_bbf :=make (map[int64 ]*_a .PdfObjectDictionary );for _ ,_gd :=range _deeg {_baa ,_af :=_a .GetIndirect (_gd );
if !_af {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_efc ,_af :=_a .GetDict (_baa .PdfObject );if !_af {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");
};_bbf [_baa .ObjectNumber ]=_efc ;};for _ ,_gc :=range _ccg {_cb ,_bea :=_a .GetIndirect (_gc );if !_bea {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");
};_ece ,_bea :=_a .GetDict (_cb .PdfObject );if !_bea {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0066\u0069\u0065\u006cd\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};T :=_ece .Get ("\u0054");if _ ,_efa :=_fca ._d [_cb .ObjectNumber ];
_efa {switch _fca ._e {case NoRestrictions ,FillForms ,FillFormsAndAnnots :_fca ._bd .addWarningWithDescription (_abg ,_g .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064",T ));default:_fca ._bd .addErrorWithDescription (_abg ,_g .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064",T ));
};};if _ ,_fgd :=_bbf [_cb .ObjectNumber ];!_fgd {switch _fca ._e {case NoRestrictions ,FillForms ,FillFormsAndAnnots :_fca ._bd .addWarningWithDescription (_abg ,_g .Sprintf ("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064",_ece .Get ("\u0054")));
default:_fca ._bd .addErrorWithDescription (_abg ,_g .Sprintf ("\u0046i\u0065l\u0064\u0020\u0025\u0073\u0020w\u0061\u0073 \u0061\u0064\u0064\u0065\u0064",_ece .Get ("\u0054")));};}else {delete (_bbf ,_cb .ObjectNumber );if _ ,_dcg :=_fca ._d [_cb .ObjectNumber ];
_dcg {switch _fca ._e {case NoRestrictions ,FillForms ,FillFormsAndAnnots :_fca ._bd .addWarningWithDescription (_abg ,_g .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064",_ece .Get ("\u0054")));
default:_fca ._bd .addErrorWithDescription (_abg ,_g .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0063\u0068\u0061ng\u0065\u0064",_ece .Get ("\u0054")));};};};if FT ,_ff :=_a .GetNameVal (_ece .Get ("\u0046\u0054"));_ff {if FT =="\u0053\u0069\u0067"{if _dfg ,_ca :=_a .GetIndirect (_ece .Get ("\u0056"));
_ca {if _ ,_cd :=_fca ._d [_dfg .ObjectNumber ];_cd {switch _fca ._e {case NoRestrictions ,FillForms ,FillFormsAndAnnots :_fca ._bd .addWarningWithDescription (_abg ,_g .Sprintf ("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064",T ));
default:_fca ._bd .addErrorWithDescription (_abg ,_g .Sprintf ("\u0053\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0066\u006f\u0072\u0020%\u0073 \u0066i\u0065l\u0064\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0061\u006e\u0067\u0065\u0064",T ));};};};};
};};for _ ,_bbfd :=range _bbf {switch _fca ._e {case NoRestrictions :_fca ._bd .addWarningWithDescription (_abg ,_g .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064",_bbfd .Get ("\u0054")));default:_fca ._bd .addErrorWithDescription (_abg ,_g .Sprintf ("F\u0069e\u006c\u0064\u0020\u0025\u0073\u0020\u0077\u0061s\u0020\u0072\u0065\u006dov\u0065\u0064",_bbfd .Get ("\u0054")));
};};return nil ;};func (_dbb *defaultDiffPolicy )compareAnnots (_ebe int ,_bed ,_bedb []_a .PdfObject )error {_dfe :=make (map[int64 ]*_a .PdfObjectDictionary );for _ ,_ce :=range _bed {_cga ,_dgd :=_a .GetIndirect (_ce );if !_dgd {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");
};_baac ,_dgd :=_a .GetDict (_cga .PdfObject );if !_dgd {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_dfe [_cga .ObjectNumber ]=_baac ;
};for _ ,_gee :=range _bedb {_bad ,_eaa :=_a .GetIndirect (_gee );if !_eaa {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_gdb ,_eaa :=_a .GetDict (_bad .PdfObject );
if !_eaa {return _b .New ("\u0075\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0061\u006e\u006e\u006ft\u0027s\u0020\u0073\u0074\u0072\u0075\u0063\u0074u\u0072\u0065");};_baab ,_ :=_a .GetStringVal (_gdb .Get ("\u0054"));_fa ,_ :=_a .GetNameVal (_gdb .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
if _ ,_eeb :=_dfe [_bad .ObjectNumber ];!_eeb {switch _dbb ._e {case NoRestrictions ,FillFormsAndAnnots :_dbb ._bd .addWarningWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064",_fa ,_baab ));
default:_caa ,_afc :=_a .GetDict (_bad .PdfObject );if !_afc {return _b .New ("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e");};_eeba ,_afc :=_a .GetNameVal (_caa .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
if !_afc {return _b .New ("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065");};if _eeba =="\u0057\u0069\u0064\u0067\u0065\u0074"{switch _dbb ._e {case NoRestrictions ,FillFormsAndAnnots ,FillForms :_dbb ._bd .addWarningWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064",_fa ,_baab ));
default:_dbb ._bd .addErrorWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064",_fa ,_baab ));};}else {_dbb ._bd .addErrorWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0025\u0073\u0020\u0077\u0061\u0073\u0020\u0061d\u0064\u0065\u0064",_fa ,_baab ));
};};}else {delete (_dfe ,_bad .ObjectNumber );if _dgbg ,_ced :=_dbb ._d [_bad .ObjectNumber ];_ced {switch _dbb ._e {case NoRestrictions ,FillFormsAndAnnots :_dbb ._bd .addWarningWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064",_fa ,_baab ));
default:_bba ,_eafc :=_a .GetIndirect (_dgbg );if !_eafc {return _b .New ("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e");};_dfa ,_eafc :=_a .GetDict (_bba .PdfObject );if !_eafc {return _b .New ("u\u006ed\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0061n\u006e\u006f\u0074\u0061ti\u006f\u006e");
};_cca ,_eafc :=_a .GetNameVal (_dfa .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));if !_eafc {return _b .New ("\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020a\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0027\u0073\u0020\u0073\u0075\u0062\u0074\u0079\u0070\u0065");
};if _cca =="\u0057\u0069\u0064\u0067\u0065\u0074"{switch _dbb ._e {case NoRestrictions ,FillFormsAndAnnots ,FillForms :_dbb ._bd .addWarningWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064",_fa ,_baab ));
default:_dbb ._bd .addErrorWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064",_fa ,_baab ));};}else {_dbb ._bd .addErrorWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0063\u0068\u0061\u006eg\u0065\u0064",_fa ,_baab ));
};};};};};for _ ,_cge :=range _dfe {_faf ,_ :=_a .GetStringVal (_cge .Get ("\u0054"));_fcc ,_ :=_a .GetNameVal (_cge .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));switch _dbb ._e {case NoRestrictions ,FillFormsAndAnnots :_dbb ._bd .addWarningWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064",_fcc ,_faf ));
default:_dbb ._bd .addErrorWithDescription (_ebe ,_g .Sprintf ("\u0025\u0073\u0020\u0061n\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0025s\u0020w\u0061\u0073\u0020\u0072\u0065\u006d\u006fv\u0065\u0064",_fcc ,_faf ));};};return nil ;};

// IsPermitted returns true if changes permitted.
func (_ace *DiffResults )IsPermitted ()bool {return len (_ace .Errors )==0};func (_gb *DiffResults )addWarningWithDescription (_fefb int ,_adf string ){if _gb .Warnings ==nil {_gb .Warnings =make ([]*DiffResult ,0);};_gb .Warnings =append (_gb .Warnings ,&DiffResult {Revision :_fefb ,Description :_adf });
};

// DiffResult describes the warning or the error for the DiffPolicy results.
type DiffResult struct{Revision int ;Description string ;};

// MDPParameters describes parameters for the MDP checks (now only DocMDP).
type MDPParameters struct{DocMDPLevel DocMDPPermission ;};