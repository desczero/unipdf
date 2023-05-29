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

// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler ;import (_gc "bytes";_gg "crypto";_cc "crypto/rand";_ce "crypto/rsa";_e "crypto/x509";_cf "crypto/x509/pkix";_b "encoding/asn1";_d "errors";_dd "fmt";_ee "github.com/unidoc/pkcs7";_da "github.com/unidoc/timestamp";_af "github.com/unidoc/unipdf/v3/core";
_cd "github.com/unidoc/unipdf/v3/model";_ca "github.com/unidoc/unipdf/v3/model/mdp";_caa "github.com/unidoc/unipdf/v3/model/sigutil";_db "hash";_g "math/big";_c "time";);

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler (handler _cd .SignatureHandler ,permission _ca .DocMDPPermission )(_cd .SignatureHandler ,error ){return &DocMDPHandler {_dad :handler ,Permission :permission },nil ;};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_cd .SignatureHandler ,error ){return &adobePKCS7Detached {_cgd :true ,_eee :signatureLen },nil ;};

// InitSignature initialises the PdfSignature.
func (_bdd *adobeX509RSASHA1 )InitSignature (sig *_cd .PdfSignature )error {if _bdd ._ec ==nil {return _d .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _bdd ._gde ==nil &&_bdd ._bgf ==nil {return _d .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");
};_gcc :=*_bdd ;sig .Handler =&_gcc ;sig .Filter =_af .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_af .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");
sig .Cert =_af .MakeString (string (_gcc ._ec .Raw ));sig .Reference =nil ;_gcbd ,_dbg :=_gcc .NewDigest (sig );if _dbg !=nil {return _dbg ;};_gcbd .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _gcc .sign (sig ,_gcbd ,_bdd ._afc );};

// Validate validates PdfSignature.
func (_cdf *adobePKCS7Detached )Validate (sig *_cd .PdfSignature ,digest _cd .Hasher )(_cd .SignatureValidationResult ,error ){_cdb :=sig .Contents .Bytes ();_dg ,_bdcf :=_ee .Parse (_cdb );if _bdcf !=nil {return _cd .SignatureValidationResult {},_bdcf ;
};_cgf :=digest .(*_gc .Buffer );_dg .Content =_cgf .Bytes ();if _bdcf =_dg .Verify ();_bdcf !=nil {return _cd .SignatureValidationResult {},_bdcf ;};return _cd .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};

// NewDigest creates a new digest.
func (_aa *DocMDPHandler )NewDigest (sig *_cd .PdfSignature )(_cd .Hasher ,error ){return _aa ._dad .NewDigest (sig );};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_gce *adobeX509RSASHA1 )IsApplicable (sig *_cd .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";
};func _edf (_eef *_ce .PublicKey ,_bab []byte )_gg .Hash {_bb :=_eef .Size ();if _bb !=len (_bab ){return 0;};_aff :=func (_fff *_g .Int ,_edg *_ce .PublicKey ,_gae *_g .Int )*_g .Int {_ged :=_g .NewInt (int64 (_edg .E ));_fff .Exp (_gae ,_ged ,_edg .N );
return _fff ;};_gba :=new (_g .Int ).SetBytes (_bab );_dea :=_aff (new (_g .Int ),_eef ,_gba );_eb :=_fbag (_dea .Bytes (),_bb );if _eb [0]!=0||_eb [1]!=1{return 0;};_eefg :=[]struct{Hash _gg .Hash ;Prefix []byte ;}{{Hash :_gg .SHA1 ,Prefix :[]byte {0x30,0x21,0x30,0x09,0x06,0x05,0x2b,0x0e,0x03,0x02,0x1a,0x05,0x00,0x04,0x14}},{Hash :_gg .SHA256 ,Prefix :[]byte {0x30,0x31,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x01,0x05,0x00,0x04,0x20}},{Hash :_gg .SHA384 ,Prefix :[]byte {0x30,0x41,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x02,0x05,0x00,0x04,0x30}},{Hash :_gg .SHA512 ,Prefix :[]byte {0x30,0x51,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x03,0x05,0x00,0x04,0x40}},{Hash :_gg .RIPEMD160 ,Prefix :[]byte {0x30,0x20,0x30,0x08,0x06,0x06,0x28,0xcf,0x06,0x03,0x00,0x31,0x04,0x14}}};
for _ ,_cggf :=range _eefg {_gaae :=_cggf .Hash .Size ();_cbd :=len (_cggf .Prefix )+_gaae ;if _gc .Equal (_eb [_bb -_cbd :_bb -_gaae ],_cggf .Prefix ){return _cggf .Hash ;};};return 0;};

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _gg .Hash )(_cd .SignatureHandler ,error ){return &docTimeStamp {_adda :timestampServerURL ,_aga :hashAlgorithm },nil ;};

// AdobeX509RSASHA1Opts defines options for configuring the adbe.x509.rsa_sha1
// signature handler.
type AdobeX509RSASHA1Opts struct{

// EstimateSize specifies whether the size of the signature contents
// should be estimated based on the modulus size of the public key
// extracted from the signing certificate. If set to false, a mock Sign
// call is made in order to estimate the size of the signature contents.
EstimateSize bool ;

// Algorithm specifies the algorithm used for performing signing.
// If not specified, defaults to SHA1.
Algorithm _gg .Hash ;};func (_ae *adobeX509RSASHA1 )getCertificate (_ag *_cd .PdfSignature )(*_e .Certificate ,error ){if _ae ._ec !=nil {return _ae ._ec ,nil ;};_gcd ,_add :=_ag .GetCerts ();if _add !=nil {return nil ,_add ;};return _gcd [0],nil ;};

// InitSignature initialization of the DocMDP signature.
func (_daa *DocMDPHandler )InitSignature (sig *_cd .PdfSignature )error {_ea :=_daa ._dad .InitSignature (sig );if _ea !=nil {return _ea ;};sig .Handler =_daa ;if sig .Reference ==nil {sig .Reference =_af .MakeArray ();};sig .Reference .Append (_cd .NewPdfSignatureReferenceDocMDP (_cd .NewPdfTransformParamsDocMDP (_daa .Permission )).ToPdfObject ());
return nil ;};

// Sign sets the Contents fields.
func (_ddb *adobePKCS7Detached )Sign (sig *_cd .PdfSignature ,digest _cd .Hasher )error {if _ddb ._cgd {_aab :=_ddb ._eee ;if _aab <=0{_aab =8192;};sig .Contents =_af .MakeHexString (string (make ([]byte ,_aab )));return nil ;};_cbg :=digest .(*_gc .Buffer );
_gag ,_gaa :=_ee .NewSignedData (_cbg .Bytes ());if _gaa !=nil {return _gaa ;};if _gca :=_gag .AddSigner (_ddb ._eec ,_ddb ._gd ,_ee .SignerInfoConfig {});_gca !=nil {return _gca ;};_gag .Detach ();_bf ,_gaa :=_gag .Finish ();if _gaa !=nil {return _gaa ;
};_ed :=make ([]byte ,8192);copy (_ed ,_bf );sig .Contents =_af .MakeHexString (string (_ed ));return nil ;};

// NewDigest creates a new digest.
func (_ddf *docTimeStamp )NewDigest (sig *_cd .PdfSignature )(_cd .Hasher ,error ){return _gc .NewBuffer (nil ),nil ;};func _fbag (_ead []byte ,_bcg int )(_cec []byte ){_dc :=len (_ead );if _dc > _bcg {_dc =_bcg ;};_cec =make ([]byte ,_bcg );copy (_cec [len (_cec )-_dc :],_ead );
return ;};func (_ge *adobeX509RSASHA1 )getHashAlgorithm (_bdb *_cd .PdfSignature )(_gg .Hash ,error ){_bfc ,_fecb :=_ge .getCertificate (_bdb );if _fecb !=nil {if _ge ._dac !=0{return _ge ._dac ,nil ;};return _ggd ,_fecb ;};if _bdb .Contents !=nil {_gad :=_bdb .Contents .Bytes ();
var _ffb []byte ;if _ ,_ddbf :=_b .Unmarshal (_gad ,&_ffb );_ddbf ==nil {_cda :=_edf (_bfc .PublicKey .(*_ce .PublicKey ),_ffb );if _cda > 0{return _cda ,nil ;};};};if _ge ._dac !=0{return _ge ._dac ,nil ;};return _ggd ,nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_fab *adobePKCS7Detached )IsApplicable (sig *_cd .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";
};

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_fd *DocMDPHandler )Validate (sig *_cd .PdfSignature ,digest _cd .Hasher )(_cd .SignatureValidationResult ,error ){return _cd .SignatureValidationResult {},_d .New ("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065");
};func (_eca *docTimeStamp )getCertificate (_aed *_cd .PdfSignature )(*_e .Certificate ,error ){_afg ,_gbac :=_aed .GetCerts ();if _gbac !=nil {return nil ,_gbac ;};return _afg [0],nil ;};func (_dge *adobeX509RSASHA1 )sign (_geg *_cd .PdfSignature ,_cdg _cd .Hasher ,_gdce bool )error {if !_gdce {return _dge .Sign (_geg ,_cdg );
};_edd ,_fbf :=_dge ._ec .PublicKey .(*_ce .PublicKey );if !_fbf {return _dd .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",_edd );};_bc ,_bcf :=_b .Marshal (make ([]byte ,_edd .Size ()));
if _bcf !=nil {return _bcf ;};_geg .Contents =_af .MakeHexString (string (_bc ));return nil ;};

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct{_dad _cd .SignatureHandler ;Permission _ca .DocMDPPermission ;};type timestampInfo struct{Version int ;Policy _b .RawValue ;MessageImprint struct{HashAlgorithm _cf .AlgorithmIdentifier ;HashedMessage []byte ;};SerialNumber _b .RawValue ;
GeneralizedTime _c .Time ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom (certificate *_e .Certificate ,signFunc SignFunc )(_cd .SignatureHandler ,error ){return &adobeX509RSASHA1 {_ec :certificate ,_bgf :signFunc },nil ;};

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_ce .PrivateKey ,certificate *_e .Certificate )(_cd .SignatureHandler ,error ){return &adobePKCS7Detached {_eec :certificate ,_gd :privateKey },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_aede *docTimeStamp )IsApplicable (sig *_cd .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";
};

// Sign adds a new reference to signature's references array.
func (_gcf *DocMDPHandler )Sign (sig *_cd .PdfSignature ,digest _cd .Hasher )error {return _gcf ._dad .Sign (sig ,digest );};

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_fe *DocMDPHandler )ValidateWithOpts (sig *_cd .PdfSignature ,digest _cd .Hasher ,params _cd .SignatureHandlerDocMDPParams )(_cd .SignatureValidationResult ,error ){_cfe ,_fc :=_fe ._dad .Validate (sig ,digest );if _fc !=nil {return _cfe ,_fc ;};
_ffd :=params .Parser ;if _ffd ==nil {return _cd .SignatureValidationResult {},_d .New ("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c");};if !_cfe .IsVerified {return _cfe ,nil ;};_gcb :=params .DiffPolicy ;
if _gcb ==nil {_gcb =_ca .NewDefaultDiffPolicy ();};for _beb :=0;_beb <=_ffd .GetRevisionNumber ();_beb ++{_cfa ,_gcbb :=_ffd .GetRevision (_beb );if _gcbb !=nil {return _cd .SignatureValidationResult {},_gcbb ;};_bd :=_cfa .GetTrailer ();if _bd ==nil {return _cd .SignatureValidationResult {},_d .New ("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074");
};_ad ,_fba :=_af .GetDict (_bd .Get ("\u0052\u006f\u006f\u0074"));if !_fba {return _cd .SignatureValidationResult {},_d .New ("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t");};
_fec ,_fba :=_af .GetDict (_ad .Get ("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"));if !_fba {continue ;};_ef ,_fba :=_af .GetArray (_fec .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_fba {continue ;};for _ ,_bdc :=range _ef .Elements (){_ddd ,_de :=_af .GetDict (_bdc );
if !_de {continue ;};_ba ,_de :=_af .GetDict (_ddd .Get ("\u0056"));if !_de {continue ;};if _af .EqualObjects (_ba .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"),sig .Contents ){_cfe .DiffResults ,_gcbb =_gcb .ReviewFile (_cfa ,_ffd ,&_ca .MDPParameters {DocMDPLevel :_fe .Permission });
if _gcbb !=nil {return _cd .SignatureValidationResult {},_gcbb ;};_cfe .IsVerified =_cfe .DiffResults .IsPermitted ();return _cfe ,nil ;};};};return _cd .SignatureValidationResult {},_d .New ("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073");
};

// Sign sets the Contents fields for the PdfSignature.
func (_bad *docTimeStamp )Sign (sig *_cd .PdfSignature ,digest _cd .Hasher )error {_cce ,_abf :=_caa .NewTimestampRequest (digest .(*_gc .Buffer ),&_da .RequestOptions {Hash :_bad ._aga ,Certificates :true });if _abf !=nil {return _abf ;};_feb :=_bad ._cef ;
if _feb ==nil {_feb =_caa .NewTimestampClient ();};_bgd ,_abf :=_feb .GetEncodedToken (_bad ._adda ,_cce );if _abf !=nil {return _abf ;};_dcc :=len (_bgd );if _bad ._bdf > 0&&_dcc > _bad ._bdf {return _cd .ErrSignNotEnoughSpace ;};if _dcc > 0{_bad ._bdf =_dcc +128;
};if sig .Contents !=nil {_gge :=sig .Contents .Bytes ();copy (_gge ,_bgd );_bgd =_gge ;};sig .Contents =_af .MakeHexString (string (_bgd ));return nil ;};type adobePKCS7Detached struct{_gd *_ce .PrivateKey ;_eec *_e .Certificate ;_cgd bool ;_eee int ;
};

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts (timestampServerURL string ,hashAlgorithm _gg .Hash ,opts *DocTimeStampOpts )(_cd .SignatureHandler ,error ){if opts ==nil {opts =&DocTimeStampOpts {};};if opts .SignatureSize <=0{opts .SignatureSize =4192;};return &docTimeStamp {_adda :timestampServerURL ,_aga :hashAlgorithm ,_bdf :opts .SignatureSize ,_cef :opts .Client },nil ;
};type docTimeStamp struct{_adda string ;_aga _gg .Hash ;_bdf int ;_cef *_caa .TimestampClient ;};

// Sign sets the Contents fields for the PdfSignature.
func (_dae *adobeX509RSASHA1 )Sign (sig *_cd .PdfSignature ,digest _cd .Hasher )error {var _gdc []byte ;var _bac error ;if _dae ._bgf !=nil {_gdc ,_bac =_dae ._bgf (sig ,digest );if _bac !=nil {return _bac ;};}else {_fecbe ,_ced :=digest .(_db .Hash );
if !_ced {return _d .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_geee :=_ggd ;if _dae ._dac !=0{_geee =_dae ._dac ;};_gdc ,_bac =_ce .SignPKCS1v15 (_cc .Reader ,_dae ._gde ,_geee ,_fecbe .Sum (nil ));if _bac !=nil {return _bac ;
};};_gdc ,_bac =_b .Marshal (_gdc );if _bac !=nil {return _bac ;};sig .Contents =_af .MakeHexString (string (_gdc ));return nil ;};

// Validate validates PdfSignature.
func (_aag *docTimeStamp )Validate (sig *_cd .PdfSignature ,digest _cd .Hasher )(_cd .SignatureValidationResult ,error ){_bce :=sig .Contents .Bytes ();_fdf ,_aeda :=_ee .Parse (_bce );if _aeda !=nil {return _cd .SignatureValidationResult {},_aeda ;};if _aeda =_fdf .Verify ();
_aeda !=nil {return _cd .SignatureValidationResult {},_aeda ;};var _fdc timestampInfo ;_ ,_aeda =_b .Unmarshal (_fdf .Content ,&_fdc );if _aeda !=nil {return _cd .SignatureValidationResult {},_aeda ;};_ac ,_aeda :=_bee (_fdc .MessageImprint .HashAlgorithm .Algorithm );
if _aeda !=nil {return _cd .SignatureValidationResult {},_aeda ;};_gcab :=_ac .New ();_dcg :=digest .(*_gc .Buffer );_gcab .Write (_dcg .Bytes ());_dbf :=_gcab .Sum (nil );_deg :=_cd .SignatureValidationResult {IsSigned :true ,IsVerified :_gc .Equal (_dbf ,_fdc .MessageImprint .HashedMessage ),GeneralizedTime :_fdc .GeneralizedTime };
return _deg ,nil ;};func _bee (_fad _b .ObjectIdentifier )(_gg .Hash ,error ){switch {case _fad .Equal (_ee .OIDDigestAlgorithmSHA1 ),_fad .Equal (_ee .OIDDigestAlgorithmECDSASHA1 ),_fad .Equal (_ee .OIDDigestAlgorithmDSA ),_fad .Equal (_ee .OIDDigestAlgorithmDSASHA1 ),_fad .Equal (_ee .OIDEncryptionAlgorithmRSA ):return _gg .SHA1 ,nil ;
case _fad .Equal (_ee .OIDDigestAlgorithmSHA256 ),_fad .Equal (_ee .OIDDigestAlgorithmECDSASHA256 ):return _gg .SHA256 ,nil ;case _fad .Equal (_ee .OIDDigestAlgorithmSHA384 ),_fad .Equal (_ee .OIDDigestAlgorithmECDSASHA384 ):return _gg .SHA384 ,nil ;case _fad .Equal (_ee .OIDDigestAlgorithmSHA512 ),_fad .Equal (_ee .OIDDigestAlgorithmECDSASHA512 ):return _gg .SHA512 ,nil ;
};return _gg .Hash (0),_ee .ErrUnsupportedAlgorithm ;};

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct{

// SignatureSize is the estimated size of the signature contents in bytes.
// If not provided, a default signature size of 4192 is used.
// The signing process will report the model.ErrSignNotEnoughSpace error
// if the estimated signature size is smaller than the actual size of the
// signature.
SignatureSize int ;

// Client is the timestamp client used to make the signature request.
// If no client is provided, a default one is used.
Client *_caa .TimestampClient ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_f *DocMDPHandler )IsApplicable (sig *_cd .PdfSignature )bool {_be :=false ;for _ ,_cb :=range sig .Reference .Elements (){if _ccb ,_cg :=_af .GetDict (_cb );_cg {if _bg ,_df :=_af .GetNameVal (_ccb .Get ("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"));
_df {if _bg !="\u0044\u006f\u0063\u004d\u0044\u0050"{return false ;};if _beg ,_fb :=_af .GetDict (_ccb .Get ("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"));_fb {_ ,_fa :=_af .GetNumberAsInt64 (_beg .Get ("\u0050"));
if _fa !=nil {return false ;};_be =true ;break ;};};};};return _be &&_f ._dad .IsApplicable (sig );};

// InitSignature initialises the PdfSignature.
func (_gb *adobePKCS7Detached )InitSignature (sig *_cd .PdfSignature )error {if !_gb ._cgd {if _gb ._eec ==nil {return _d .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _gb ._gd ==nil {return _d .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_fed :=*_gb ;sig .Handler =&_fed ;sig .Filter =_af .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");
sig .SubFilter =_af .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_bea ,_beaf :=_fed .NewDigest (sig );if _beaf !=nil {return _beaf ;};_bea .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _fed .Sign (sig ,_bea );};const _ggd =_gg .SHA1 ;type adobeX509RSASHA1 struct{_gde *_ce .PrivateKey ;_ec *_e .Certificate ;_bgf SignFunc ;_afc bool ;_dac _gg .Hash ;};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_gf *_cd .PdfSignature ,_cff _cd .Hasher )([]byte ,error );

// NewDigest creates a new digest.
func (_gfb *adobeX509RSASHA1 )NewDigest (sig *_cd .PdfSignature )(_cd .Hasher ,error ){if _caag ,_baf :=_gfb .getHashAlgorithm (sig );_caag !=0&&_baf ==nil {return _caag .New (),nil ;};return _ggd .New (),nil ;};

// InitSignature initialises the PdfSignature.
func (_ebe *docTimeStamp )InitSignature (sig *_cd .PdfSignature )error {_gdcc :=*_ebe ;sig .Handler =&_gdcc ;sig .Filter =_af .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_af .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");
sig .Reference =nil ;if _ebe ._bdf > 0{sig .Contents =_af .MakeHexString (string (make ([]byte ,_ebe ._bdf )));}else {_deb ,_dga :=_ebe .NewDigest (sig );if _dga !=nil {return _dga ;};_deb .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
if _dga =_gdcc .Sign (sig ,_deb );_dga !=nil {return _dga ;};_ebe ._bdf =_gdcc ._bdf ;};return nil ;};func (_ga *adobePKCS7Detached )getCertificate (_adc *_cd .PdfSignature )(*_e .Certificate ,error ){if _ga ._eec !=nil {return _ga ._eec ,nil ;};_eecc ,_cgg :=_adc .GetCerts ();
if _cgg !=nil {return nil ,_cgg ;};return _eecc [0],nil ;};

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_ce .PrivateKey ,certificate *_e .Certificate )(_cd .SignatureHandler ,error ){return &adobeX509RSASHA1 {_ec :certificate ,_gde :privateKey },nil ;};

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts (certificate *_e .Certificate ,signFunc SignFunc ,opts *AdobeX509RSASHA1Opts )(_cd .SignatureHandler ,error ){if opts ==nil {opts =&AdobeX509RSASHA1Opts {};};return &adobeX509RSASHA1 {_ec :certificate ,_bgf :signFunc ,_afc :opts .EstimateSize ,_dac :opts .Algorithm },nil ;
};

// Validate validates PdfSignature.
func (_dab *adobeX509RSASHA1 )Validate (sig *_cd .PdfSignature ,digest _cd .Hasher )(_cd .SignatureValidationResult ,error ){_dec ,_cdbc :=_dab .getCertificate (sig );if _cdbc !=nil {return _cd .SignatureValidationResult {},_cdbc ;};_fabb :=sig .Contents .Bytes ();
var _gee []byte ;if _ ,_afd :=_b .Unmarshal (_fabb ,&_gee );_afd !=nil {return _cd .SignatureValidationResult {},_afd ;};_aea ,_fcg :=digest .(_db .Hash );if !_fcg {return _cd .SignatureValidationResult {},_d .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};_ab ,_ :=_dab .getHashAlgorithm (sig );if _ab ==0{_ab =_ggd ;};if _gda :=_ce .VerifyPKCS1v15 (_dec .PublicKey .(*_ce .PublicKey ),_ab ,_aea .Sum (nil ),_gee );_gda !=nil {return _cd .SignatureValidationResult {},_gda ;};return _cd .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;
};

// NewDigest creates a new digest.
func (_eaa *adobePKCS7Detached )NewDigest (sig *_cd .PdfSignature )(_cd .Hasher ,error ){return _gc .NewBuffer (nil ),nil ;};