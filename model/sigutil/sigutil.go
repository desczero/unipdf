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

package sigutil ;import (_a "bytes";_e "crypto";_gg "crypto/x509";_gb "encoding/asn1";_ade "encoding/pem";_eb "errors";_fc "fmt";_ee "github.com/unidoc/timestamp";_d "github.com/unidoc/unipdf/v3/common";_cb "golang.org/x/crypto/ocsp";_g "io";_b "io/ioutil";
_f "net/http";_ad "time";);

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_cf *CertClient )IsCA (cert *_gg .Certificate )bool {return cert .IsCA &&_a .Equal (cert .RawIssuer ,cert .RawSubject );};func _fed ()*_f .Client {return &_f .Client {Timeout :5*_ad .Second }};

// NewCRLClient returns a new CRL client.
func NewCRLClient ()*CRLClient {return &CRLClient {HTTPClient :_fed ()}};

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct{

// HTTPClient is the HTTP client used to make timestamp requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_f .Client ;

// Callbacks.
BeforeHTTPRequest func (_ae *_f .Request )error ;};

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct{

// HTTPClient is the HTTP client used to make CRL requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_f .Client ;};

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient ()*TimestampClient {return &TimestampClient {HTTPClient :_fed ()}};

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_aeg *TimestampClient )GetEncodedToken (serverURL string ,req *_ee .Request )([]byte ,error ){if serverURL ==""{return nil ,_fc .Errorf ("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c");
};if req ==nil {return nil ,_fc .Errorf ("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c");};_bbf ,_gc :=req .Marshal ();if _gc !=nil {return nil ,_gc ;
};_ded ,_gc :=_f .NewRequest ("\u0050\u004f\u0053\u0054",serverURL ,_a .NewBuffer (_bbf ));if _gc !=nil {return nil ,_gc ;};_ded .Header .Set ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079");
if _aeg .BeforeHTTPRequest !=nil {if _gbf :=_aeg .BeforeHTTPRequest (_ded );_gbf !=nil {return nil ,_gbf ;};};_gcf :=_aeg .HTTPClient ;if _gcf ==nil {_gcf =_fed ();};_gge ,_gc :=_gcf .Do (_ded );if _gc !=nil {return nil ,_gc ;};defer _gge .Body .Close ();
_faa ,_gc :=_b .ReadAll (_gge .Body );if _gc !=nil {return nil ,_gc ;};if _gge .StatusCode !=_f .StatusOK {return nil ,_fc .Errorf ("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064",_gge .StatusCode );
};var _ag struct{Version _gb .RawValue ;Content _gb .RawValue ;};if _ ,_gc =_gb .Unmarshal (_faa ,&_ag );_gc !=nil {return nil ,_gc ;};return _ag .Content .FullBytes ,nil ;};

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_fg *CRLClient )MakeRequest (serverURL string ,cert *_gg .Certificate )([]byte ,error ){if _fg .HTTPClient ==nil {_fg .HTTPClient =_fed ();};if serverURL ==""{if len (cert .CRLDistributionPoints )==0{return nil ,_eb .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073");
};serverURL =cert .CRLDistributionPoints [0];};_ffc ,_faf :=_fg .HTTPClient .Get (serverURL );if _faf !=nil {return nil ,_faf ;};defer _ffc .Body .Close ();_ace ,_faf :=_b .ReadAll (_ffc .Body );if _faf !=nil {return nil ,_faf ;};if _fb ,_ :=_ade .Decode (_ace );
_fb !=nil {_ace =_fb .Bytes ;};return _ace ,nil ;};

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_ga *OCSPClient )MakeRequest (serverURL string ,cert ,issuer *_gg .Certificate )(*_cb .Response ,[]byte ,error ){if _ga .HTTPClient ==nil {_ga .HTTPClient =_fed ();};if serverURL ==""{if len (cert .OCSPServer )==0{return nil ,nil ,_eb .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073");
};serverURL =cert .OCSPServer [0];};_fbe ,_bcg :=_cb .CreateRequest (cert ,issuer ,&_cb .RequestOptions {Hash :_ga .Hash });if _bcg !=nil {return nil ,nil ,_bcg ;};_abc ,_bcg :=_ga .HTTPClient .Post (serverURL ,"\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074",_a .NewReader (_fbe ));
if _bcg !=nil {return nil ,nil ,_bcg ;};defer _abc .Body .Close ();_gba ,_bcg :=_b .ReadAll (_abc .Body );if _bcg !=nil {return nil ,nil ,_bcg ;};if _bb ,_ :=_ade .Decode (_gba );_bb !=nil {_gba =_bb .Bytes ;};_eee ,_bcg :=_cb .ParseResponseForCert (_gba ,cert ,issuer );
if _bcg !=nil {return nil ,nil ,_bcg ;};return _eee ,_gba ,nil ;};

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest (body _g .Reader ,opts *_ee .RequestOptions )(*_ee .Request ,error ){if opts ==nil {opts =&_ee .RequestOptions {};};if opts .Hash ==0{opts .Hash =_e .SHA256 ;};if !opts .Hash .Available (){return nil ,_gg .ErrUnsupportedAlgorithm ;
};_cd :=opts .Hash .New ();if _ ,_fae :=_g .Copy (_cd ,body );_fae !=nil {return nil ,_fae ;};return &_ee .Request {HashAlgorithm :opts .Hash ,HashedMessage :_cd .Sum (nil ),Certificates :opts .Certificates ,TSAPolicyOID :opts .TSAPolicyOID ,Nonce :opts .Nonce },nil ;
};

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient ()*OCSPClient {return &OCSPClient {HTTPClient :_fed (),Hash :_e .SHA1 }};

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct{

// HTTPClient is the HTTP client used to make OCSP requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_f .Client ;

// Hash is the hash function  used when constructing the OCSP
// requests. If zero, SHA-1 will be used.
Hash _e .Hash ;};

// NewCertClient returns a new certificate client.
func NewCertClient ()*CertClient {return &CertClient {HTTPClient :_fed ()}};

// GetIssuer retrieves the issuer of the provided certificate.
func (_fa *CertClient )GetIssuer (cert *_gg .Certificate )(*_gg .Certificate ,error ){for _ ,_eed :=range cert .IssuingCertificateURL {_de ,_bc :=_fa .Get (_eed );if _bc !=nil {_d .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076",cert .Subject .CommonName ,_bc );
continue ;};return _de ,nil ;};return nil ,_fc .Errorf ("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064");};

// Get retrieves the certificate at the specified URL.
func (_ca *CertClient )Get (url string )(*_gg .Certificate ,error ){if _ca .HTTPClient ==nil {_ca .HTTPClient =_fed ();};_ac ,_ff :=_ca .HTTPClient .Get (url );if _ff !=nil {return nil ,_ff ;};defer _ac .Body .Close ();_ef ,_ff :=_b .ReadAll (_ac .Body );
if _ff !=nil {return nil ,_ff ;};if _fd ,_ :=_ade .Decode (_ef );_fd !=nil {_ef =_fd .Bytes ;};_ggg ,_ff :=_gg .ParseCertificate (_ef );if _ff !=nil {return nil ,_ff ;};return _ggg ,nil ;};

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct{

// HTTPClient is the HTTP client used to make certificate requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_f .Client ;};