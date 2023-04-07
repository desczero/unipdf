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
package sighandler ;import (_ed "bytes";_f "crypto";_g "crypto/rand";_eb "crypto/rsa";_fb "crypto/x509";_dd "crypto/x509/pkix";_df "encoding/asn1";_c "errors";_ca "fmt";_cf "github.com/unidoc/pkcs7";_fg "github.com/unidoc/timestamp";_b "github.com/unidoc/unipdf/v3/core";
_bgf "github.com/unidoc/unipdf/v3/model";_gg "github.com/unidoc/unipdf/v3/model/mdp";_bg "github.com/unidoc/unipdf/v3/model/sigutil";_ec "hash";_e "math/big";_a "time";);

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts (timestampServerURL string ,hashAlgorithm _f .Hash ,opts *DocTimeStampOpts )(_bgf .SignatureHandler ,error ){if opts ==nil {opts =&DocTimeStampOpts {};};if opts .SignatureSize <=0{opts .SignatureSize =4192;};return &docTimeStamp {_abf :timestampServerURL ,_acf :hashAlgorithm ,_cgf :opts .SignatureSize ,_eagg :opts .Client },nil ;
};

// Sign adds a new reference to signature's references array.
func (_fbb *DocMDPHandler )Sign (sig *_bgf .PdfSignature ,digest _bgf .Hasher )error {return _fbb ._ebd .Sign (sig ,digest );};

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_eb .PrivateKey ,certificate *_fb .Certificate )(_bgf .SignatureHandler ,error ){return &adobeX509RSASHA1 {_gbgf :certificate ,_de :privateKey },nil ;};

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts (certificate *_fb .Certificate ,signFunc SignFunc ,opts *AdobeX509RSASHA1Opts )(_bgf .SignatureHandler ,error ){if opts ==nil {opts =&AdobeX509RSASHA1Opts {};};return &adobeX509RSASHA1 {_gbgf :certificate ,_eag :signFunc ,_acg :opts .EstimateSize ,_cab :opts .Algorithm },nil ;
};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_edc *adobeX509RSASHA1 )IsApplicable (sig *_bgf .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";
};

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
Algorithm _f .Hash ;};type adobeX509RSASHA1 struct{_de *_eb .PrivateKey ;_gbgf *_fb .Certificate ;_eag SignFunc ;_acg bool ;_cab _f .Hash ;};

// NewDigest creates a new digest.
func (_geb *DocMDPHandler )NewDigest (sig *_bgf .PdfSignature )(_bgf .Hasher ,error ){return _geb ._ebd .NewDigest (sig );};

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct{_ebd _bgf .SignatureHandler ;Permission _gg .DocMDPPermission ;};

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
Client *_bg .TimestampClient ;};

// Validate validates PdfSignature.
func (_fee *adobePKCS7Detached )Validate (sig *_bgf .PdfSignature ,digest _bgf .Hasher )(_bgf .SignatureValidationResult ,error ){_da :=sig .Contents .Bytes ();_ae ,_fea :=_cf .Parse (_da );if _fea !=nil {return _bgf .SignatureValidationResult {},_fea ;
};_af :=digest .(*_ed .Buffer );_ae .Content =_af .Bytes ();if _fea =_ae .Verify ();_fea !=nil {return _bgf .SignatureValidationResult {},_fea ;};return _bgf .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_aba *_bgf .PdfSignature ,_cg _bgf .Hasher )([]byte ,error );

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_eb .PrivateKey ,certificate *_fb .Certificate )(_bgf .SignatureHandler ,error ){return &adobePKCS7Detached {_eda :certificate ,_be :privateKey },nil ;};

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler (handler _bgf .SignatureHandler ,permission _gg .DocMDPPermission )(_bgf .SignatureHandler ,error ){return &DocMDPHandler {_ebd :handler ,Permission :permission },nil ;};

// Validate validates PdfSignature.
func (_dbb *docTimeStamp )Validate (sig *_bgf .PdfSignature ,digest _bgf .Hasher )(_bgf .SignatureValidationResult ,error ){_bfc :=sig .Contents .Bytes ();_gbe ,_egcc :=_cf .Parse (_bfc );if _egcc !=nil {return _bgf .SignatureValidationResult {},_egcc ;
};if _egcc =_gbe .Verify ();_egcc !=nil {return _bgf .SignatureValidationResult {},_egcc ;};var _gae timestampInfo ;_ ,_egcc =_df .Unmarshal (_gbe .Content ,&_gae );if _egcc !=nil {return _bgf .SignatureValidationResult {},_egcc ;};_fedc ,_egcc :=_dbe (_gae .MessageImprint .HashAlgorithm .Algorithm );
if _egcc !=nil {return _bgf .SignatureValidationResult {},_egcc ;};_afd :=_fedc .New ();_gcfb :=digest .(*_ed .Buffer );_afd .Write (_gcfb .Bytes ());_efdb :=_afd .Sum (nil );_dcbg :=_bgf .SignatureValidationResult {IsSigned :true ,IsVerified :_ed .Equal (_efdb ,_gae .MessageImprint .HashedMessage ),GeneralizedTime :_gae .GeneralizedTime };
return _dcbg ,nil ;};

// Sign sets the Contents fields for the PdfSignature.
func (_bda *adobeX509RSASHA1 )Sign (sig *_bgf .PdfSignature ,digest _bgf .Hasher )error {var _fed []byte ;var _cef error ;if _bda ._eag !=nil {_fed ,_cef =_bda ._eag (sig ,digest );if _cef !=nil {return _cef ;};}else {_dee ,_gbc :=digest .(_ec .Hash );
if !_gbc {return _c .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_cfd :=_dfd ;if _bda ._cab !=0{_cfd =_bda ._cab ;};_fed ,_cef =_eb .SignPKCS1v15 (_g .Reader ,_bda ._de ,_cfd ,_dee .Sum (nil ));if _cef !=nil {return _cef ;
};};_fed ,_cef =_df .Marshal (_fed );if _cef !=nil {return _cef ;};sig .Contents =_b .MakeHexString (string (_fed ));return nil ;};func (_bcd *adobeX509RSASHA1 )getHashAlgorithm (_dcc *_bgf .PdfSignature )(_f .Hash ,error ){_adc ,_gbb :=_bcd .getCertificate (_dcc );
if _gbb !=nil {if _bcd ._cab !=0{return _bcd ._cab ,nil ;};return _dfd ,_gbb ;};if _dcc .Contents !=nil {_gbgg :=_dcc .Contents .Bytes ();var _gcc []byte ;if _ ,_efe :=_df .Unmarshal (_gbgg ,&_gcc );_efe ==nil {_dcb :=_age (_adc .PublicKey .(*_eb .PublicKey ),_gcc );
if _dcb > 0{return _dcb ,nil ;};};};if _bcd ._cab !=0{return _bcd ._cab ,nil ;};return _dfd ,nil ;};

// InitSignature initialises the PdfSignature.
func (_cfg *adobeX509RSASHA1 )InitSignature (sig *_bgf .PdfSignature )error {if _cfg ._gbgf ==nil {return _c .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _cfg ._de ==nil &&_cfg ._eag ==nil {return _c .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");
};_bc :=*_cfg ;sig .Handler =&_bc ;sig .Filter =_b .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_b .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");
sig .Cert =_b .MakeString (string (_bc ._gbgf .Raw ));sig .Reference =nil ;_ggd ,_eagc :=_bc .NewDigest (sig );if _eagc !=nil {return _eagc ;};_ggd .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _bc .sign (sig ,_ggd ,_cfg ._acg );};func (_agf *adobePKCS7Detached )getCertificate (_ffd *_bgf .PdfSignature )(*_fb .Certificate ,error ){if _agf ._eda !=nil {return _agf ._eda ,nil ;};_edf ,_caa :=_ffd .GetCerts ();if _caa !=nil {return nil ,_caa ;
};return _edf [0],nil ;};

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_ce *DocMDPHandler )ValidateWithOpts (sig *_bgf .PdfSignature ,digest _bgf .Hasher ,params _bgf .SignatureHandlerDocMDPParams )(_bgf .SignatureValidationResult ,error ){_fge ,_ba :=_ce ._ebd .Validate (sig ,digest );if _ba !=nil {return _fge ,_ba ;
};_cff :=params .Parser ;if _cff ==nil {return _bgf .SignatureValidationResult {},_c .New ("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c");};if !_fge .IsVerified {return _fge ,nil ;};_edd :=params .DiffPolicy ;
if _edd ==nil {_edd =_gg .NewDefaultDiffPolicy ();};for _gc :=0;_gc <=_cff .GetRevisionNumber ();_gc ++{_ag ,_gca :=_cff .GetRevision (_gc );if _gca !=nil {return _bgf .SignatureValidationResult {},_gca ;};_eg :=_ag .GetTrailer ();if _eg ==nil {return _bgf .SignatureValidationResult {},_c .New ("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074");
};_feb ,_gdf :=_b .GetDict (_eg .Get ("\u0052\u006f\u006f\u0074"));if !_gdf {return _bgf .SignatureValidationResult {},_c .New ("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t");
};_dfg ,_gdf :=_b .GetDict (_feb .Get ("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"));if !_gdf {continue ;};_gbf ,_gdf :=_b .GetArray (_dfg .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_gdf {continue ;};for _ ,_ff :=range _gbf .Elements (){_fa ,_cdb :=_b .GetDict (_ff );
if !_cdb {continue ;};_agc ,_cdb :=_b .GetDict (_fa .Get ("\u0056"));if !_cdb {continue ;};if _b .EqualObjects (_agc .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"),sig .Contents ){_fge .DiffResults ,_gca =_edd .ReviewFile (_ag ,_cff ,&_gg .MDPParameters {DocMDPLevel :_ce .Permission });
if _gca !=nil {return _bgf .SignatureValidationResult {},_gca ;};_fge .IsVerified =_fge .DiffResults .IsPermitted ();return _fge ,nil ;};};};return _bgf .SignatureValidationResult {},_c .New ("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073");
};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_bgf .SignatureHandler ,error ){return &adobePKCS7Detached {_ee :true ,_bd :signatureLen },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_fedd *docTimeStamp )IsApplicable (sig *_bgf .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";
};func _age (_bf *_eb .PublicKey ,_acc []byte )_f .Hash {_ecd :=_bf .Size ();if _ecd !=len (_acc ){return 0;};_bgc :=func (_fbe *_e .Int ,_gba *_eb .PublicKey ,_efef *_e .Int )*_e .Int {_cde :=_e .NewInt (int64 (_gba .E ));_fbe .Exp (_efef ,_cde ,_gba .N );
return _fbe ;};_fbba :=new (_e .Int ).SetBytes (_acc );_gdc :=_bgc (new (_e .Int ),_bf ,_fbba );_dfdd :=_ebc (_gdc .Bytes (),_ecd );if _dfdd [0]!=0||_dfdd [1]!=1{return 0;};_aae :=[]struct{Hash _f .Hash ;Prefix []byte ;}{{Hash :_f .SHA1 ,Prefix :[]byte {0x30,0x21,0x30,0x09,0x06,0x05,0x2b,0x0e,0x03,0x02,0x1a,0x05,0x00,0x04,0x14}},{Hash :_f .SHA256 ,Prefix :[]byte {0x30,0x31,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x01,0x05,0x00,0x04,0x20}},{Hash :_f .SHA384 ,Prefix :[]byte {0x30,0x41,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x02,0x05,0x00,0x04,0x30}},{Hash :_f .SHA512 ,Prefix :[]byte {0x30,0x51,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x03,0x05,0x00,0x04,0x40}},{Hash :_f .RIPEMD160 ,Prefix :[]byte {0x30,0x20,0x30,0x08,0x06,0x06,0x28,0xcf,0x06,0x03,0x00,0x31,0x04,0x14}}};
for _ ,_abae :=range _aae {_bde :=_abae .Hash .Size ();_edcf :=len (_abae .Prefix )+_bde ;if _ed .Equal (_dfdd [_ecd -_edcf :_ecd -_bde ],_abae .Prefix ){return _abae .Hash ;};};return 0;};

// Validate validates PdfSignature.
func (_gbge *adobeX509RSASHA1 )Validate (sig *_bgf .PdfSignature ,digest _bgf .Hasher )(_bgf .SignatureValidationResult ,error ){_gcf ,_afe :=_gbge .getCertificate (sig );if _afe !=nil {return _bgf .SignatureValidationResult {},_afe ;};_aab :=sig .Contents .Bytes ();
var _fafd []byte ;if _ ,_eea :=_df .Unmarshal (_aab ,&_fafd );_eea !=nil {return _bgf .SignatureValidationResult {},_eea ;};_gbfe ,_aaa :=digest .(_ec .Hash );if !_aaa {return _bgf .SignatureValidationResult {},_c .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};_bbe ,_ :=_gbge .getHashAlgorithm (sig );if _bbe ==0{_bbe =_dfd ;};if _ga :=_eb .VerifyPKCS1v15 (_gcf .PublicKey .(*_eb .PublicKey ),_bbe ,_gbfe .Sum (nil ),_fafd );_ga !=nil {return _bgf .SignatureValidationResult {},_ga ;};return _bgf .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;
};

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _f .Hash )(_bgf .SignatureHandler ,error ){return &docTimeStamp {_abf :timestampServerURL ,_acf :hashAlgorithm },nil ;};

// InitSignature initialises the PdfSignature.
func (_bac *adobePKCS7Detached )InitSignature (sig *_bgf .PdfSignature )error {if !_bac ._ee {if _bac ._eda ==nil {return _c .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _bac ._be ==nil {return _c .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_ab :=*_bac ;sig .Handler =&_ab ;sig .Filter =_b .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");
sig .SubFilter =_b .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_aa ,_fbc :=_ab .NewDigest (sig );if _fbc !=nil {return _fbc ;};_aa .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _ab .Sign (sig ,_aa );};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_dfc *adobePKCS7Detached )IsApplicable (sig *_bgf .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";
};

// NewDigest creates a new digest.
func (_egc *adobeX509RSASHA1 )NewDigest (sig *_bgf .PdfSignature )(_bgf .Hasher ,error ){if _cda ,_faf :=_egc .getHashAlgorithm (sig );_cda !=0&&_faf ==nil {return _cda .New (),nil ;};return _dfd .New (),nil ;};func (_febg *adobeX509RSASHA1 )getCertificate (_gf *_bgf .PdfSignature )(*_fb .Certificate ,error ){if _febg ._gbgf !=nil {return _febg ._gbgf ,nil ;
};_bbg ,_dg :=_gf .GetCerts ();if _dg !=nil {return nil ,_dg ;};return _bbg [0],nil ;};type adobePKCS7Detached struct{_be *_eb .PrivateKey ;_eda *_fb .Certificate ;_ee bool ;_bd int ;};type docTimeStamp struct{_abf string ;_acf _f .Hash ;_cgf int ;_eagg *_bg .TimestampClient ;
};

// Sign sets the Contents fields for the PdfSignature.
func (_cb *docTimeStamp )Sign (sig *_bgf .PdfSignature ,digest _bgf .Hasher )error {_fgc ,_gdd :=_bg .NewTimestampRequest (digest .(*_ed .Buffer ),&_fg .RequestOptions {Hash :_cb ._acf ,Certificates :true });if _gdd !=nil {return _gdd ;};_fgf :=_cb ._eagg ;
if _fgf ==nil {_fgf =_bg .NewTimestampClient ();};_dcf ,_gdd :=_fgf .GetEncodedToken (_cb ._abf ,_fgc );if _gdd !=nil {return _gdd ;};_cgd :=len (_dcf );if _cb ._cgf > 0&&_cgd > _cb ._cgf {return _bgf .ErrSignNotEnoughSpace ;};if _cgd > 0{_cb ._cgf =_cgd +128;
};if sig .Contents !=nil {_abb :=sig .Contents .Bytes ();copy (_abb ,_dcf );_dcf =_abb ;};sig .Contents =_b .MakeHexString (string (_dcf ));return nil ;};

// NewDigest creates a new digest.
func (_cc *adobePKCS7Detached )NewDigest (sig *_bgf .PdfSignature )(_bgf .Hasher ,error ){return _ed .NewBuffer (nil ),nil ;};func _ebc (_gab []byte ,_egce int )(_aga []byte ){_afa :=len (_gab );if _afa > _egce {_afa =_egce ;};_aga =make ([]byte ,_egce );
copy (_aga [len (_aga )-_afa :],_gab );return ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom (certificate *_fb .Certificate ,signFunc SignFunc )(_bgf .SignatureHandler ,error ){return &adobeX509RSASHA1 {_gbgf :certificate ,_eag :signFunc },nil ;};func (_gfg *docTimeStamp )getCertificate (_bdd *_bgf .PdfSignature )(*_fb .Certificate ,error ){_cae ,_fgb :=_bdd .GetCerts ();
if _fgb !=nil {return nil ,_fgb ;};return _cae [0],nil ;};

// InitSignature initialization of the DocMDP signature.
func (_ad *DocMDPHandler )InitSignature (sig *_bgf .PdfSignature )error {_cfb :=_ad ._ebd .InitSignature (sig );if _cfb !=nil {return _cfb ;};sig .Handler =_ad ;if sig .Reference ==nil {sig .Reference =_b .MakeArray ();};sig .Reference .Append (_bgf .NewPdfSignatureReferenceDocMDP (_bgf .NewPdfTransformParamsDocMDP (_ad .Permission )).ToPdfObject ());
return nil ;};

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_ac *DocMDPHandler )Validate (sig *_bgf .PdfSignature ,digest _bgf .Hasher )(_bgf .SignatureValidationResult ,error ){return _bgf .SignatureValidationResult {},_c .New ("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065");
};func _dbe (_ceb _df .ObjectIdentifier )(_f .Hash ,error ){switch {case _ceb .Equal (_cf .OIDDigestAlgorithmSHA1 ),_ceb .Equal (_cf .OIDDigestAlgorithmECDSASHA1 ),_ceb .Equal (_cf .OIDDigestAlgorithmDSA ),_ceb .Equal (_cf .OIDDigestAlgorithmDSASHA1 ),_ceb .Equal (_cf .OIDEncryptionAlgorithmRSA ):return _f .SHA1 ,nil ;
case _ceb .Equal (_cf .OIDDigestAlgorithmSHA256 ),_ceb .Equal (_cf .OIDDigestAlgorithmECDSASHA256 ):return _f .SHA256 ,nil ;case _ceb .Equal (_cf .OIDDigestAlgorithmSHA384 ),_ceb .Equal (_cf .OIDDigestAlgorithmECDSASHA384 ):return _f .SHA384 ,nil ;case _ceb .Equal (_cf .OIDDigestAlgorithmSHA512 ),_ceb .Equal (_cf .OIDDigestAlgorithmECDSASHA512 ):return _f .SHA512 ,nil ;
};return _f .Hash (0),_cf .ErrUnsupportedAlgorithm ;};type timestampInfo struct{Version int ;Policy _df .RawValue ;MessageImprint struct{HashAlgorithm _dd .AlgorithmIdentifier ;HashedMessage []byte ;};SerialNumber _df .RawValue ;GeneralizedTime _a .Time ;
};

// Sign sets the Contents fields.
func (_efd *adobePKCS7Detached )Sign (sig *_bgf .PdfSignature ,digest _bgf .Hasher )error {if _efd ._ee {_db :=_efd ._bd ;if _db <=0{_db =8192;};sig .Contents =_b .MakeHexString (string (make ([]byte ,_db )));return nil ;};_efde :=digest .(*_ed .Buffer );
_aff ,_ecb :=_cf .NewSignedData (_efde .Bytes ());if _ecb !=nil {return _ecb ;};if _efg :=_aff .AddSigner (_efd ._eda ,_efd ._be ,_cf .SignerInfoConfig {});_efg !=nil {return _efg ;};_aff .Detach ();_eab ,_ecb :=_aff .Finish ();if _ecb !=nil {return _ecb ;
};_abg :=make ([]byte ,8192);copy (_abg ,_eab );sig .Contents =_b .MakeHexString (string (_abg ));return nil ;};

// NewDigest creates a new digest.
func (_fdb *docTimeStamp )NewDigest (sig *_bgf .PdfSignature )(_bgf .Hasher ,error ){return _ed .NewBuffer (nil ),nil ;};

// InitSignature initialises the PdfSignature.
func (_eddg *docTimeStamp )InitSignature (sig *_bgf .PdfSignature )error {_deeb :=*_eddg ;sig .Handler =&_deeb ;sig .Filter =_b .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_b .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");
sig .Reference =nil ;if _eddg ._cgf > 0{sig .Contents =_b .MakeHexString (string (make ([]byte ,_eddg ._cgf )));}else {_gbbb ,_edcg :=_eddg .NewDigest (sig );if _edcg !=nil {return _edcg ;};_gbbb .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
if _edcg =_deeb .Sign (sig ,_gbbb );_edcg !=nil {return _edcg ;};_eddg ._cgf =_deeb ._cgf ;};return nil ;};func (_cfa *adobeX509RSASHA1 )sign (_fgeg *_bgf .PdfSignature ,_gcb _bgf .Hasher ,_fd bool )error {if !_fd {return _cfa .Sign (_fgeg ,_gcb );};_feag ,_ccb :=_cfa ._gbgf .PublicKey .(*_eb .PublicKey );
if !_ccb {return _ca .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",_feag );};_fac ,_bce :=_df .Marshal (make ([]byte ,_feag .Size ()));if _bce !=nil {return _bce ;
};_fgeg .Contents =_b .MakeHexString (string (_fac ));return nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ef *DocMDPHandler )IsApplicable (sig *_bgf .PdfSignature )bool {_gd :=false ;for _ ,_gb :=range sig .Reference .Elements (){if _ea ,_fba :=_b .GetDict (_gb );_fba {if _dc ,_ge :=_b .GetNameVal (_ea .Get ("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"));
_ge {if _dc !="\u0044\u006f\u0063\u004d\u0044\u0050"{return false ;};if _gbg ,_fe :=_b .GetDict (_ea .Get ("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"));_fe {_ ,_cd :=_b .GetNumberAsInt64 (_gbg .Get ("\u0050"));if _cd !=nil {return false ;
};_gd =true ;break ;};};};};return _gd &&_ef ._ebd .IsApplicable (sig );};const _dfd =_f .SHA1 ;