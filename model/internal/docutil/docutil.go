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

package docutil ;import (_e "errors";_b "fmt";_f "github.com/unidoc/unipdf/v3/common";_g "github.com/unidoc/unipdf/v3/core";);func (_aa *Catalog )SetMarkInfo (mi _g .PdfObject ){_aa .Object .Set ("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f",mi );
_aa ._c .Objects =append (_aa ._c .Objects ,mi );};type OutputIntent struct{Object *_g .PdfObjectDictionary ;};type Page struct{_gbf int ;Object *_g .PdfObjectDictionary ;_fece *Document ;};func (_fd *Catalog )HasMetadata ()bool {_ged :=_fd .Object .Get ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061");
return _ged !=nil ;};func (_gee *Catalog )SetOutputIntents (outputIntents *OutputIntents ){if _age :=_gee .Object .Get ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073");_age !=nil {for _eg ,_fec :=range _gee ._c .Objects {if _fec ==_age {if outputIntents ._cb ==_age {return ;
};_gee ._c .Objects =append (_gee ._c .Objects [:_eg ],_gee ._c .Objects [_eg +1:]...);break ;};};};_fdb :=outputIntents ._cb ;if _fdb ==nil {_fdb =_g .MakeIndirectObject (outputIntents ._cef );};_gee .Object .Set ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073",_fdb );
_gee ._c .Objects =append (_gee ._c .Objects ,_fdb );};func (_ae *OutputIntents )Len ()int {return _ae ._cef .Len ()};func _eggg (_baf _g .PdfObject )(_g .PdfObjectName ,error ){var _bgd *_g .PdfObjectName ;var _edc *_g .PdfObjectArray ;if _eggc ,_faf :=_baf .(*_g .PdfIndirectObject );
_faf {if _ddg ,_acb :=_eggc .PdfObject .(*_g .PdfObjectArray );_acb {_edc =_ddg ;}else if _fea ,_gdf :=_eggc .PdfObject .(*_g .PdfObjectName );_gdf {_bgd =_fea ;};}else if _ff ,_cc :=_baf .(*_g .PdfObjectArray );_cc {_edc =_ff ;}else if _be ,_aee :=_baf .(*_g .PdfObjectName );
_aee {_bgd =_be ;};if _bgd !=nil {switch *_bgd {case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079","\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B","\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":return *_bgd ,nil ;case "\u0050a\u0074\u0074\u0065\u0072\u006e":return *_bgd ,nil ;
};};if _edc !=nil &&_edc .Len ()> 0{if _gedd ,_cbe :=_edc .Get (0).(*_g .PdfObjectName );_cbe {switch *_gedd {case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079","\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B","\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":if _edc .Len ()==1{return *_gedd ,nil ;
};case "\u0043a\u006c\u0047\u0072\u0061\u0079","\u0043\u0061\u006c\u0052\u0047\u0042","\u004c\u0061\u0062":return *_gedd ,nil ;case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064","\u0050a\u0074\u0074\u0065\u0072\u006e","\u0049n\u0064\u0065\u0078\u0065\u0064":return *_gedd ,nil ;
case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e","\u0044e\u0076\u0069\u0063\u0065\u004e":return *_gedd ,nil ;};};};return "",nil ;};func (_dc *Document )GetPages ()([]Page ,bool ){_cdd ,_bdb :=_dc .FindCatalog ();if !_bdb {return nil ,false ;
};return _cdd .GetPages ();};func (_ggdc Page )GetResourcesXObject ()(*_g .PdfObjectDictionary ,bool ){_cfc ,_gcg :=_ggdc .GetResources ();if !_gcg {return nil ,false ;};return _g .GetDict (_cfc .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));};type OutputIntents struct{_cef *_g .PdfObjectArray ;
_cf *Document ;_cb *_g .PdfIndirectObject ;};func (_ca *Document )FindCatalog ()(*Catalog ,bool ){var _dged *_g .PdfObjectDictionary ;for _ ,_ab :=range _ca .Objects {_cff ,_gff :=_g .GetDict (_ab );if !_gff {continue ;};if _caa ,_fb :=_g .GetName (_cff .Get ("\u0054\u0079\u0070\u0065"));
_fb &&*_caa =="\u0043a\u0074\u0061\u006c\u006f\u0067"{_dged =_cff ;break ;};};if _dged ==nil {return nil ,false ;};return &Catalog {Object :_dged ,_c :_ca },true ;};func (_ggde *Catalog )NewOutputIntents ()*OutputIntents {return &OutputIntents {_cf :_ggde ._c }};
func (_aece *Document )AddIndirectObject (indirect *_g .PdfIndirectObject ){for _ ,_fa :=range _aece .Objects {if _fa ==indirect {return ;};};_aece .Objects =append (_aece .Objects ,indirect );};func (_fbc Page )GetResources ()(*_g .PdfObjectDictionary ,bool ){return _g .GetDict (_fbc .Object .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));
};type Content struct{Stream *_g .PdfObjectStream ;_aff int ;_aaf Page ;};func (_cca Content )GetData ()([]byte ,error ){_eea ,_eed :=_g .NewEncoderFromStream (_cca .Stream );if _eed !=nil {return nil ,_eed ;};_gda ,_eed :=_eea .DecodeStream (_cca .Stream );
if _eed !=nil {return nil ,_eed ;};return _gda ,nil ;};func (_bdd *OutputIntents )Get (i int )(OutputIntent ,bool ){if _bdd ._cef ==nil {return OutputIntent {},false ;};if i >=_bdd ._cef .Len (){return OutputIntent {},false ;};_ba :=_bdd ._cef .Get (i );
_ac ,_gc :=_g .GetIndirect (_ba );if !_gc {_gfb ,_aec :=_g .GetDict (_ba );return OutputIntent {Object :_gfb },_aec ;};_bda ,_ed :=_g .GetDict (_ac .PdfObject );return OutputIntent {Object :_bda },_ed ;};func (_eb *Catalog )GetPages ()([]Page ,bool ){_d ,_ce :=_g .GetDict (_eb .Object .Get ("\u0050\u0061\u0067e\u0073"));
if !_ce {return nil ,false ;};_fe ,_gg :=_g .GetArray (_d .Get ("\u004b\u0069\u0064\u0073"));if !_gg {return nil ,false ;};_dg :=make ([]Page ,_fe .Len ());for _bf ,_ge :=range _fe .Elements (){_gef ,_ee :=_g .GetDict (_ge );if !_ee {continue ;};_dg [_bf ]=Page {Object :_gef ,_gbf :_bf +1,_fece :_eb ._c };
};return _dg ,true ;};type ImageSMask struct{Image *Image ;Stream *_g .PdfObjectStream ;};func (_ag *Catalog )GetMetadata ()(*_g .PdfObjectStream ,bool ){_ea ,_bd :=_g .GetStream (_ag .Object .Get ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"));return _ea ,_bd ;
};func (_egd Page )FindXObjectForms ()[]*_g .PdfObjectStream {_eaa ,_dfce :=_egd .GetResourcesXObject ();if !_dfce {return nil ;};_agec :=map[*_g .PdfObjectStream ]struct{}{};for _ ,_db :=range _eaa .Keys (){_feg ,_abg :=_g .GetStream (_eaa .Get (_db ));
if !_abg {continue ;};if _ ,_geb :=_agec [_feg ];_geb {continue ;};_cdf ,_ec :=_g .GetName (_feg .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));if !_ec ||_cdf .String ()!="\u0046\u006f\u0072\u006d"{continue ;};_agec [_feg ]=struct{}{};};var _bcg []*_g .PdfObjectStream ;
for _bcf :=range _agec {_bcg =append (_bcg ,_bcf );};return _bcg ;};func (_aae Page )FindXObjectImages ()([]*Image ,error ){_de ,_adf :=_aae .GetResourcesXObject ();if !_adf {return nil ,nil ;};var _cdc []*Image ;var _bgb error ;_gde :=map[*_g .PdfObjectStream ]int {};
_gba :=map[*_g .PdfObjectStream ]struct{}{};var _ef int ;for _ ,_dcd :=range _de .Keys (){_df ,_gfbe :=_g .GetStream (_de .Get (_dcd ));if !_gfbe {continue ;};if _ ,_ddgf :=_gde [_df ];_ddgf {continue ;};_afd ,_dab :=_g .GetName (_df .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
if !_dab ||_afd .String ()!="\u0049\u006d\u0061g\u0065"{continue ;};_bc :=Image {BitsPerComponent :8,Stream :_df ,Name :string (_dcd )};if _bc .Colorspace ,_bgb =_eggg (_df .PdfObjectDictionary .Get ("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"));
_bgb !=nil {_f .Log .Error ("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073",_bgb );continue ;};if _fg ,_bgc :=_g .GetIntVal (_df .PdfObjectDictionary .Get ("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"));
_bgc {_bc .BitsPerComponent =_fg ;};if _def ,_dgee :=_g .GetIntVal (_df .PdfObjectDictionary .Get ("\u0057\u0069\u0064t\u0068"));_dgee {_bc .Width =_def ;};if _agf ,_afa :=_g .GetIntVal (_df .PdfObjectDictionary .Get ("\u0048\u0065\u0069\u0067\u0068\u0074"));
_afa {_bc .Height =_agf ;};if _gcd ,_ede :=_g .GetStream (_df .Get ("\u0053\u004d\u0061s\u006b"));_ede {_bc .SMask =&ImageSMask {Image :&_bc ,Stream :_gcd };_gba [_gcd ]=struct{}{};};switch _bc .Colorspace {case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":_bc .ColorComponents =3;
case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":_bc .ColorComponents =1;case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":_bc .ColorComponents =4;default:_bc .ColorComponents =-1;};_gde [_df ]=_ef ;_cdc =append (_cdc ,&_bc );
_ef ++;};var _gcc []int ;for _ ,_gcb :=range _cdc {if _gcb .SMask !=nil {_bca ,_ace :=_gde [_gcb .SMask .Stream ];if _ace {_gcc =append (_gcc ,_bca );};};};_gbg :=make ([]*Image ,len (_cdc )-len (_gcc ));_ef =0;_dfg :for _fdbb ,_dfc :=range _cdc {for _ ,_dee :=range _gcc {if _fdbb ==_dee {continue _dfg ;
};};_gbg [_ef ]=_dfc ;_ef ++;};return _cdc ,nil ;};func (_dge *Catalog )GetOutputIntents ()(*OutputIntents ,bool ){_ad :=_dge .Object .Get ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073");if _ad ==nil {return nil ,false ;
};_egg ,_dac :=_g .GetIndirect (_ad );if !_dac {return nil ,false ;};_geg ,_bg :=_g .GetArray (_egg .PdfObject );if !_bg {return nil ,false ;};return &OutputIntents {_cb :_egg ,_cef :_geg ,_cf :_dge ._c },true ;};type Catalog struct{Object *_g .PdfObjectDictionary ;
_c *Document ;};func (_agg Page )GetContents ()([]Content ,bool ){_ebb ,_adg :=_g .GetArray (_agg .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if !_adg {return nil ,false ;};_bga :=make ([]Content ,_ebb .Len ());for _gbb ,_edcb :=range _ebb .Elements (){_ggg ,_dgf :=_g .GetStream (_edcb );
if !_dgf {continue ;};_bga [_gbb ]=Content {Stream :_ggg ,_aaf :_agg ,_aff :_gbb };};return _bga ,true ;};func (_gbe *OutputIntents )Add (oi _g .PdfObject )error {_ggf ,_ebe :=oi .(*_g .PdfObjectDictionary );if !_ebe {return _e .New ("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};if _geed ,_ebc :=_g .GetStream (_ggf .Get ("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065"));_ebc {_gbe ._cf .Objects =append (_gbe ._cf .Objects ,_geed );};_ga ,_bdg :=oi .(*_g .PdfIndirectObject );
if !_bdg {_ga =_g .MakeIndirectObject (oi );};if _gbe ._cef ==nil {_gbe ._cef =_g .MakeArray (_ga );}else {_gbe ._cef .Append (_ga );};_gbe ._cf .Objects =append (_gbe ._cf .Objects ,_ga );return nil ;};func (_bed *Page )Number ()int {return _bed ._gbf };
func (_af *Catalog )GetMarkInfo ()(*_g .PdfObjectDictionary ,bool ){_gd ,_da :=_g .GetDict (_af .Object .Get ("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"));return _gd ,_da ;};func (_afae *Content )SetData (data []byte )error {_abb ,_bcfc :=_g .MakeStream (data ,_g .NewFlateEncoder ());
if _bcfc !=nil {return _bcfc ;};_bafg ,_ :=_g .GetArray (_afae ._aaf .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if _bcfc =_bafg .Set (_afae ._aff ,_abb );_bcfc !=nil {return _bcfc ;};_afae ._aaf ._fece .Objects =append (_afae ._aaf ._fece .Objects ,_abb );
return nil ;};type Document struct{ID [2]string ;Version _g .Version ;Objects []_g .PdfObject ;Info _g .PdfObject ;Crypt *_g .PdfCrypt ;UseHashBasedID bool ;};func (_ggd *Catalog )SetMetadata (data []byte )error {_gb ,_gea :=_g .MakeStream (data ,nil );
if _gea !=nil {return _gea ;};_gb .Set ("\u0054\u0079\u0070\u0065",_g .MakeName ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"));_gb .Set ("\u0053u\u0062\u0074\u0079\u0070\u0065",_g .MakeName ("\u0058\u004d\u004c"));_ggd .Object .Set ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061",_gb );
_ggd ._c .Objects =append (_ggd ._c .Objects ,_gb );return nil ;};func (_cac *Document )AddStream (stream *_g .PdfObjectStream ){for _ ,_dd :=range _cac .Objects {if _dd ==stream {return ;};};_cac .Objects =append (_cac .Objects ,stream );};type Image struct{Name string ;
Width int ;Height int ;Colorspace _g .PdfObjectName ;ColorComponents int ;BitsPerComponent int ;SMask *ImageSMask ;Stream *_g .PdfObjectStream ;};func (_gf *Catalog )SetVersion (){_gf .Object .Set ("\u0056e\u0072\u0073\u0069\u006f\u006e",_g .MakeName (_b .Sprintf ("\u0025\u0064\u002e%\u0064",_gf ._c .Version .Major ,_gf ._c .Version .Minor )));
};