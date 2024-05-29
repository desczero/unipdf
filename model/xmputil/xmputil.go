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

// Package xmputil provides abstraction used by the pdf document XMP Metadata.
package xmputil ;import (_a "errors";_cd "fmt";_b "github.com/trimmer-io/go-xmp/models/pdf";_bf "github.com/trimmer-io/go-xmp/models/xmp_base";_c "github.com/trimmer-io/go-xmp/models/xmp_mm";_ef "github.com/trimmer-io/go-xmp/xmp";_fb "github.com/unidoc/unipdf/v3/core";
_g "github.com/unidoc/unipdf/v3/internal/timeutils";_aa "github.com/unidoc/unipdf/v3/internal/uuid";_be "github.com/unidoc/unipdf/v3/model/xmputil/pdfaextension";_eg "github.com/unidoc/unipdf/v3/model/xmputil/pdfaid";_fc "strconv";_f "time";);

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument (stream []byte )(*Document ,error ){_ed :=_ef .NewDocument ();if _fbb :=_ef .Unmarshal (stream ,_ed );_fbb !=nil {return nil ,_fbb ;};return &Document {_cc :_ed },nil ;};

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct{OriginalDocumentID GUID ;DocumentID GUID ;InstanceID GUID ;VersionID string ;};

// MediaManagementOptions are the options for the Media management xmp metadata.
type MediaManagementOptions struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
// By default, this value is generated.
OriginalDocumentID string ;

// NewDocumentID is a flag which generates a new Document identifier while setting media management.
// This value should be set to true only if the document is stored and saved as new document.
// Otherwise, if the document is modified and overwrites previous file, it should be set to false.
NewDocumentID bool ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
// By default, this value is generated if NewDocumentID is true or previous doesn't exist.
DocumentID string ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
// By default, this value is generated.
InstanceID string ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
// By default, the derived from structure is filled from previous XMP metadata (if exists).
DerivedFrom string ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
// By default, this values is incremented or set to the next version number.
VersionID string ;

// ModifyComment is a comment to given modification
ModifyComment string ;

// ModifyDate is a custom modification date for the versions.
// By default, this would be set to time.Now().
ModifyDate _f .Time ;

// Modifier is a person who did the modification.
Modifier string ;};

// GetPdfAID gets the pdfaid xmp metadata model.
func (_eef *Document )GetPdfAID ()(*PdfAID ,bool ){_cdb ,_gg :=_eef ._cc .FindModel (_eg .Namespace ).(*_eg .Model );if !_gg {return nil ,false ;};return &PdfAID {Part :_cdb .Part ,Conformance :_cdb .Conformance },true ;};

// SetPdfInfo sets the pdf info into selected document.
func (_bfd *Document )SetPdfInfo (options *PdfInfoOptions )error {if options ==nil {return _a .New ("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};_fbbe ,_ce :=_b .MakeModel (_bfd ._cc );
if _ce !=nil {return _ce ;};if options .Overwrite {*_fbbe =_b .PDFInfo {};};if options .InfoDict !=nil {_df ,_fa :=_fb .GetDict (options .InfoDict );if !_fa {return _cd .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",options .InfoDict );
};var _da *_fb .PdfObjectString ;for _ ,_ae :=range _df .Keys (){switch _ae {case "\u0054\u0069\u0074l\u0065":_da ,_fa =_fb .GetString (_df .Get ("\u0054\u0069\u0074l\u0065"));if _fa {_fbbe .Title =_ef .NewAltString (_da );};case "\u0041\u0075\u0074\u0068\u006f\u0072":_da ,_fa =_fb .GetString (_df .Get ("\u0041\u0075\u0074\u0068\u006f\u0072"));
if _fa {_fbbe .Author =_ef .NewStringList (_da .String ());};case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":_da ,_fa =_fb .GetString (_df .Get ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"));if _fa {_fbbe .Keywords =_da .String ();};case "\u0043r\u0065\u0061\u0074\u006f\u0072":_da ,_fa =_fb .GetString (_df .Get ("\u0043r\u0065\u0061\u0074\u006f\u0072"));
if _fa {_fbbe .Creator =_ef .AgentName (_da .String ());};case "\u0053u\u0062\u006a\u0065\u0063\u0074":_da ,_fa =_fb .GetString (_df .Get ("\u0053u\u0062\u006a\u0065\u0063\u0074"));if _fa {_fbbe .Subject =_ef .NewAltString (_da .String ());};case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":_da ,_fa =_fb .GetString (_df .Get ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"));
if _fa {_fbbe .Producer =_ef .AgentName (_da .String ());};case "\u0054r\u0061\u0070\u0070\u0065\u0064":_bea ,_fce :=_fb .GetName (_df .Get ("\u0054r\u0061\u0070\u0070\u0065\u0064"));if _fce {switch _bea .String (){case "\u0054\u0072\u0075\u0065":_fbbe .Trapped =true ;
case "\u0046\u0061\u006cs\u0065":_fbbe .Trapped =false ;default:_fbbe .Trapped =true ;};};case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":if _ge ,_dge :=_fb .GetString (_df .Get ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"));
_dge &&_ge .String ()!=""{_gf ,_fe :=_g .ParsePdfTime (_ge .String ());if _fe !=nil {return _cd .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077",_fe );
};_fbbe .CreationDate =_ef .NewDate (_gf );};case "\u004do\u0064\u0044\u0061\u0074\u0065":if _feg ,_ag :=_fb .GetString (_df .Get ("\u004do\u0064\u0044\u0061\u0074\u0065"));_ag &&_feg .String ()!=""{_dfc ,_gc :=_g .ParsePdfTime (_feg .String ());if _gc !=nil {return _cd .Errorf ("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w",_gc );
};_fbbe .ModifyDate =_ef .NewDate (_dfc );};};};};if options .PdfVersion !=""{_fbbe .PDFVersion =options .PdfVersion ;};if options .Marked {_fbbe .Marked =_ef .Bool (options .Marked );};if options .Copyright !=""{_fbbe .Copyright =options .Copyright ;};
if _ce =_fbbe .SyncToXMP (_bfd ._cc );_ce !=nil {return _ce ;};return nil ;};

// PdfInfo is the xmp document pdf info.
type PdfInfo struct{InfoDict _fb .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;};

// NewDocument creates a new document without any previous xmp information.
func NewDocument ()*Document {_ab :=_ef .NewDocument ();return &Document {_cc :_ab }};

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_fd *Document )GetGoXmpDocument ()*_ef .Document {return _fd ._cc };

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_gce *Document )GetMediaManagement ()(*MediaManagement ,bool ){_fg :=_c .FindModel (_gce ._cc );if _fg ==nil {return nil ,false ;};_cde :=make ([]MediaManagementVersion ,len (_fg .Versions ));for _egd ,_eag :=range _fg .Versions {_cde [_egd ]=MediaManagementVersion {VersionID :_eag .Version ,ModifyDate :_eag .ModifyDate .Value (),Comments :_eag .Comments ,Modifier :_eag .Modifier };
};_ecd :=&MediaManagement {OriginalDocumentID :GUID (_fg .OriginalDocumentID .Value ()),DocumentID :GUID (_fg .DocumentID .Value ()),InstanceID :GUID (_fg .InstanceID .Value ()),VersionID :_fg .VersionID ,Versions :_cde };if _fg .DerivedFrom !=nil {_ecd .DerivedFrom =&MediaManagementDerivedFrom {OriginalDocumentID :GUID (_fg .DerivedFrom .OriginalDocumentID ),DocumentID :GUID (_fg .DerivedFrom .DocumentID ),InstanceID :GUID (_fg .DerivedFrom .InstanceID ),VersionID :_fg .DerivedFrom .VersionID };
};return _ecd ,true ;};

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{_cc *_ef .Document };

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_ba *Document )MarshalIndent (prefix ,indent string )([]byte ,error ){if _ba ._cc .IsDirty (){if _aad :=_ba ._cc .SyncModels ();_aad !=nil {return nil ,_aad ;};};return _ef .MarshalIndent (_ba ._cc ,prefix ,indent );};

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_cgc *Document )SetPdfAID (part int ,conformance string )error {_gb ,_efa :=_eg .MakeModel (_cgc ._cc );if _efa !=nil {return _efa ;};_gb .Part =part ;_gb .Conformance =conformance ;if _gd :=_gb .SyncToXMP (_cgc ._cc );_gd !=nil {return _gd ;};return nil ;
};

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_ee *Document )GetPdfaExtensionSchemas ()([]_be .Schema ,error ){_dc :=_ee ._cc .FindModel (_be .Namespace );if _dc ==nil {return nil ,nil ;};_cce ,_db :=_dc .(*_be .Model );if !_db {return nil ,_cd .Errorf ("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054",_dc );
};return _cce .Schemas ,nil ;};

// GetPdfInfo gets the document pdf info.
func (_dbb *Document )GetPdfInfo ()(*PdfInfo ,bool ){_dcb :=PdfInfo {};var _cdc *_fb .PdfObjectDictionary ;_ec :=func (_ea string ,_de _fb .PdfObject ){if _cdc ==nil {_cdc =_fb .MakeDict ();};_cdc .Set (_fb .PdfObjectName (_ea ),_de );};_dfe ,_eac :=_dbb ._cc .FindModel (_b .NsPDF ).(*_b .PDFInfo );
if !_eac {_fdd ,_bd :=_dbb ._cc .FindModel (_bf .NsXmp ).(*_bf .XmpBase );if !_bd {return nil ,false ;};if _fdd .CreatorTool !=""{_ec ("\u0043r\u0065\u0061\u0074\u006f\u0072",_fb .MakeString (string (_fdd .CreatorTool )));};if !_fdd .CreateDate .IsZero (){_ec ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_fb .MakeString (_g .FormatPdfTime (_fdd .CreateDate .Value ())));
};if !_fdd .ModifyDate .IsZero (){_ec ("\u004do\u0064\u0044\u0061\u0074\u0065",_fb .MakeString (_g .FormatPdfTime (_fdd .ModifyDate .Value ())));};_dcb .InfoDict =_cdc ;return &_dcb ,true ;};_dcb .Copyright =_dfe .Copyright ;_dcb .PdfVersion =_dfe .PDFVersion ;
_dcb .Marked =bool (_dfe .Marked );if len (_dfe .Title )> 0{_ec ("\u0054\u0069\u0074l\u0065",_fb .MakeString (_dfe .Title .Default ()));};if len (_dfe .Author )> 0{_ec ("\u0041\u0075\u0074\u0068\u006f\u0072",_fb .MakeString (_dfe .Author [0]));};if _dfe .Keywords !=""{_ec ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",_fb .MakeString (_dfe .Keywords ));
};if len (_dfe .Subject )> 0{_ec ("\u0053u\u0062\u006a\u0065\u0063\u0074",_fb .MakeString (_dfe .Subject .Default ()));};if _dfe .Creator !=""{_ec ("\u0043r\u0065\u0061\u0074\u006f\u0072",_fb .MakeString (string (_dfe .Creator )));};if _dfe .Producer !=""{_ec ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",_fb .MakeString (string (_dfe .Producer )));
};if _dfe .Trapped {_ec ("\u0054r\u0061\u0070\u0070\u0065\u0064",_fb .MakeName ("\u0054\u0072\u0075\u0065"));};if !_dfe .CreationDate .IsZero (){_ec ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_fb .MakeString (_g .FormatPdfTime (_dfe .CreationDate .Value ())));
};if !_dfe .ModifyDate .IsZero (){_ec ("\u004do\u0064\u0044\u0061\u0074\u0065",_fb .MakeString (_g .FormatPdfTime (_dfe .ModifyDate .Value ())));};_dcb .InfoDict =_cdc ;return &_dcb ,true ;};

// Marshal the document into xml byte stream.
func (_cdg *Document )Marshal ()([]byte ,error ){if _cdg ._cc .IsDirty (){if _d :=_cdg ._cc .SyncModels ();_d !=nil {return nil ,_d ;};};return _ef .Marshal (_cdg ._cc );};

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct{InfoDict _fb .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;

// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
Overwrite bool ;};

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct{VersionID string ;ModifyDate _f .Time ;Comments string ;Modifier string ;};

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_eb *Document )SetMediaManagement (options *MediaManagementOptions )error {_ad ,_fad :=_c .MakeModel (_eb ._cc );if _fad !=nil {return _fad ;};if options ==nil {options =new (MediaManagementOptions );};_ac :=_c .ResourceRef {};if _ad .OriginalDocumentID .IsZero (){if options .OriginalDocumentID !=""{_ad .OriginalDocumentID =_ef .GUID (options .OriginalDocumentID );
}else {_aeg ,_baf :=_aa .NewUUID ();if _baf !=nil {return _baf ;};_ad .OriginalDocumentID =_ef .GUID (_aeg .String ());};}else {_ac .OriginalDocumentID =_ad .OriginalDocumentID ;};switch {case options .DocumentID !="":_ad .DocumentID =_ef .GUID (options .DocumentID );
case options .NewDocumentID ||_ad .DocumentID .IsZero ():if !_ad .DocumentID .IsZero (){_ac .DocumentID =_ad .DocumentID ;};_cg ,_fcg :=_aa .NewUUID ();if _fcg !=nil {return _fcg ;};_ad .DocumentID =_ef .GUID (_cg .String ());};if !_ad .InstanceID .IsZero (){_ac .InstanceID =_ad .InstanceID ;
};_ad .InstanceID =_ef .GUID (options .InstanceID );if _ad .InstanceID ==""{_bfc ,_fde :=_aa .NewUUID ();if _fde !=nil {return _fde ;};_ad .InstanceID =_ef .GUID (_bfc .String ());};if !_ac .IsZero (){_ad .DerivedFrom =&_ac ;};_ebg :=options .VersionID ;
if _ad .VersionID !=""{_faa ,_agd :=_fc .Atoi (_ad .VersionID );if _agd !=nil {_ebg =_fc .Itoa (len (_ad .Versions )+1);}else {_ebg =_fc .Itoa (_faa +1);};};if _ebg ==""{_ebg ="\u0031";};_ad .VersionID =_ebg ;if _fad =_ad .SyncToXMP (_eb ._cc );_fad !=nil {return _fad ;
};return nil ;};

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_fcb *Document )SetPdfAExtension ()error {_dg ,_dgb :=_be .MakeModel (_fcb ._cc );if _dgb !=nil {return _dgb ;};if _dgb =_be .FillModel (_fcb ._cc ,_dg );_dgb !=nil {return _dgb ;};if _dgb =_dg .SyncToXMP (_fcb ._cc );_dgb !=nil {return _dgb ;};return nil ;
};

// GUID is a string representing a globally unique identifier.
type GUID string ;

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct{Part int ;Conformance string ;};

// MediaManagement are the values from the document media management metadata.
type MediaManagement struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
OriginalDocumentID GUID ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
DocumentID GUID ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
InstanceID GUID ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
DerivedFrom *MediaManagementDerivedFrom ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
VersionID string ;

// Versions is the history of the document versions along with the comments, timestamps and issuers.
Versions []MediaManagementVersion ;};