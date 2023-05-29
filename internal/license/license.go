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

package license ;import (_dc "bytes";_bd "compress/gzip";_g "crypto";_ec "crypto/aes";_cc "crypto/cipher";_c "crypto/hmac";_gf "crypto/rand";_ba "crypto/rsa";_ee "crypto/sha256";_ca "crypto/sha512";_ff "crypto/x509";_fa "encoding/base64";_eb "encoding/hex";
_dcg "encoding/json";_aa "encoding/pem";_af "errors";_afb "fmt";_ab "github.com/unidoc/unipdf/v3/common";_be "io";_e "io/ioutil";_d "net";_gbf "net/http";_gb "os";_gg "path/filepath";_gd "sort";_f "strings";_a "sync";_de "time";);type LicenseKey struct{LicenseId string `json:"license_id"`;
CustomerId string `json:"customer_id"`;CustomerName string `json:"customer_name"`;Tier string `json:"tier"`;CreatedAt _de .Time `json:"-"`;CreatedAtInt int64 `json:"created_at"`;ExpiresAt *_de .Time `json:"-"`;ExpiresAtInt int64 `json:"expires_at"`;CreatedBy string `json:"created_by"`;
CreatorName string `json:"creator_name"`;CreatorEmail string `json:"creator_email"`;UniPDF bool `json:"unipdf"`;UniOffice bool `json:"unioffice"`;UniHTML bool `json:"unihtml"`;Trial bool `json:"trial"`;_da bool ;_geeb string ;_fd bool ;};func MakeUnlicensedKey ()*LicenseKey {_eca :=LicenseKey {};
_eca .CustomerName ="\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";_eca .Tier =LicenseTierUnlicensed ;_eca .CreatedAt =_de .Now ().UTC ();_eca .CreatedAtInt =_eca .CreatedAt .Unix ();return &_eca ;};const _afac ="U\u004eI\u0050\u0044\u0046\u005f\u0043\u0055\u0053\u0054O\u004d\u0045\u0052\u005fNA\u004d\u0045";
func (_bcg defaultStateHolder )updateState (_efd ,_gc ,_dde string ,_dgb int ,_egb bool ,_agd int ,_bga int ,_gcd _de .Time ,_gcb map[string ]int )error {_cd :=_bbf ();if len (_cd )==0{return _af .New ("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");
};_bb :=_gg .Join (_cd ,"\u002eu\u006e\u0069\u0064\u006f\u0063");_df :=_gb .MkdirAll (_bb ,0777);if _df !=nil {return _df ;};if len (_efd )< 20{return _af .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079");};_fba :=[]byte (_efd );_gfb :=_ca .Sum512_256 (_fba [:20]);
_dac :=_eb .EncodeToString (_gfb [:]);_dacb :=_gg .Join (_bb ,_dac );var _gagb reportState ;_gagb .Docs =int64 (_dgb );_gagb .NumErrors =int64 (_bga );_gagb .LimitDocs =_egb ;_gagb .RemainingDocs =int64 (_agd );_gagb .LastWritten =_de .Now ().UTC ();_gagb .LastReported =_gcd ;
_gagb .Instance =_gc ;_gagb .Next =_dde ;_gagb .Usage =_gcb ;_gfa ,_df :=_dcg .Marshal (_gagb );if _df !=nil {return _df ;};const _ebde ="\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057";_gfa ,_df =_ecca ([]byte (_ebde ),_gfa );
if _df !=nil {return _df ;};_df =_e .WriteFile (_dacb ,_gfa ,0600);if _df !=nil {return _df ;};return nil ;};var _efg =MakeUnlicensedKey ();func (_aca *LicenseKey )TypeToString ()string {if _aca ._da {return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e";
};if _aca .Tier ==LicenseTierUnlicensed {return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";};if _aca .Tier ==LicenseTierCommunity {return "\u0041\u0047PL\u0076\u0033\u0020O\u0070\u0065\u006e\u0020Sou\u0072ce\u0020\u0043\u006f\u006d\u006d\u0075\u006eit\u0079\u0020\u004c\u0069\u0063\u0065\u006es\u0065";
};if _aca .Tier ==LicenseTierIndividual ||_aca .Tier =="\u0069\u006e\u0064i\u0065"{return "\u0043\u006f\u006dm\u0065\u0072\u0063\u0069a\u006c\u0020\u004c\u0069\u0063\u0065\u006es\u0065\u0020\u002d\u0020\u0049\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c";
};return "\u0043\u006fm\u006d\u0065\u0072\u0063\u0069\u0061\u006c\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u002d\u0020\u0042\u0075\u0073\u0069ne\u0073\u0073";};func TrackUse (useKey string ){if _efg ==nil {return ;};if !_efg ._da ||len (_efg ._geeb )==0{return ;
};if len (useKey )==0{return ;};_adc .Lock ();defer _adc .Unlock ();if _bed ==nil {_bed =map[string ]int {};};_bed [useKey ]++;};const (LicenseTierUnlicensed ="\u0075\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";LicenseTierCommunity ="\u0063o\u006d\u006d\u0075\u006e\u0069\u0074y";
LicenseTierIndividual ="\u0069\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c";LicenseTierBusiness ="\u0062\u0075\u0073\u0069\u006e\u0065\u0073\u0073";);func _eafd ()*meteredClient {_bfb :=meteredClient {_ed :"h\u0074\u0074\u0070\u0073\u003a\u002f/\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069d\u006f\u0063\u002ei\u006f/\u0061\u0070\u0069",_dbe :&_gbf .Client {Timeout :30*_de .Second }};
if _gbe :=_gb .Getenv ("\u0055N\u0049\u0044\u004f\u0043_\u004c\u0049\u0043\u0045\u004eS\u0045_\u0053E\u0052\u0056\u0045\u0052\u005f\u0055\u0052L");_f .HasPrefix (_gbe ,"\u0068\u0074\u0074\u0070"){_bfb ._ed =_gbe ;};return &_bfb ;};func (_ggg *meteredClient )checkinUsage (_cea meteredUsageCheckinForm )(meteredUsageCheckinResp ,error ){_cea .Package ="\u0075\u006e\u0069\u0070\u0064\u0066";
_cea .PackageVersion =_ab .Version ;var _ebd meteredUsageCheckinResp ;_beg :=_ggg ._ed +"\u002f\u006d\u0065\u0074er\u0065\u0064\u002f\u0075\u0073\u0061\u0067\u0065\u005f\u0063\u0068\u0065\u0063\u006bi\u006e";_gbec ,_ccf :=_dcg .Marshal (_cea );if _ccf !=nil {return _ebd ,_ccf ;
};_gab ,_ccf :=_faaf (_gbec );if _ccf !=nil {return _ebd ,_ccf ;};_bgc ,_ccf :=_gbf .NewRequest ("\u0050\u004f\u0053\u0054",_beg ,_gab );if _ccf !=nil {return _ebd ,_ccf ;};_bgc .Header .Add ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e");
_bgc .Header .Add ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");_bgc .Header .Add ("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");
_bgc .Header .Add ("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y",_ggg ._fbc );_fda ,_ccf :=_ggg ._dbe .Do (_bgc );if _ccf !=nil {return _ebd ,_ccf ;};defer _fda .Body .Close ();if _fda .StatusCode !=200{_fgdg ,_gfdf :=_ggd (_fda );if _gfdf !=nil {return _ebd ,_gfdf ;
};_gfdf =_dcg .Unmarshal (_fgdg ,&_ebd );if _gfdf !=nil {return _ebd ,_gfdf ;};return _ebd ,_afb .Errorf ("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064",_fda .StatusCode );
};_ecd :=_fda .Header .Get ("\u0058\u002d\u0055\u0043\u002d\u0053\u0069\u0067\u006ea\u0074\u0075\u0072\u0065");_egg :=_efda (_cea .MacAddress ,string (_gbec ));if _egg !=_ecd {_ab .Log .Error ("I\u006e\u0076\u0061l\u0069\u0064\u0020\u0072\u0065\u0073\u0070\u006f\u006e\u0073\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u002c\u0020\u0073\u0065t\u0020\u0074\u0068e\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0073\u0065\u0072\u0076e\u0072\u0020\u0074\u006f \u0068\u0074\u0074\u0070s\u003a\u002f\u002f\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069o\u002f\u0061\u0070\u0069");
return _ebd ,_af .New ("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0065\u0072\u0076\u0065\u0072 \u0072\u0065\u0073\u0070\u006f\u006e\u0073\u0065");
};_fde ,_ccf :=_ggd (_fda );if _ccf !=nil {return _ebd ,_ccf ;};_ccf =_dcg .Unmarshal (_fde ,&_ebd );if _ccf !=nil {return _ebd ,_ccf ;};return _ebd ,nil ;};func (_cbg *LicenseKey )ToString ()string {if _cbg ._da {return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e";
};_gfd :=_afb .Sprintf ("\u004ci\u0063e\u006e\u0073\u0065\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a",_cbg .LicenseId );_gfd +=_afb .Sprintf ("\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a",_cbg .CustomerId );
_gfd +=_afb .Sprintf ("\u0043u\u0073t\u006f\u006d\u0065\u0072\u0020N\u0061\u006de\u003a\u0020\u0025\u0073\u000a",_cbg .CustomerName );_gfd +=_afb .Sprintf ("\u0054i\u0065\u0072\u003a\u0020\u0025\u0073\n",_cbg .Tier );_gfd +=_afb .Sprintf ("\u0043r\u0065a\u0074\u0065\u0064\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a",_ab .UtcTimeFormat (_cbg .CreatedAt ));
if _cbg .ExpiresAt ==nil {_gfd +="\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041t\u003a\u0020N\u0065\u0076\u0065\u0072\u000a";}else {_gfd +=_afb .Sprintf ("\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a",_ab .UtcTimeFormat (*_cbg .ExpiresAt ));
};_gfd +=_afb .Sprintf ("\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u003a\u0020\u0025\u0073\u0020<\u0025\u0073\u003e\u000a",_cbg .CreatorName ,_cbg .CreatorEmail );return _gfd ;};func _bbf ()string {_caaf :=_gb .Getenv ("\u0048\u004f\u004d\u0045");if len (_caaf )==0{_caaf ,_ =_gb .UserHomeDir ();
};return _caaf ;};var _ade stateLoader =defaultStateHolder {};var _eae =_de .Date (2020,1,1,0,0,0,0,_de .UTC );type meteredUsageCheckinForm struct{Instance string `json:"inst"`;Next string `json:"next"`;UsageNumber int `json:"usage_number"`;NumFailed int64 `json:"num_failed"`;
Hostname string `json:"hostname"`;LocalIP string `json:"local_ip"`;MacAddress string `json:"mac_address"`;Package string `json:"package"`;PackageVersion string `json:"package_version"`;Usage map[string ]int `json:"u"`;IsPersistentCache bool `json:"is_persistent_cache"`;
Timestamp int64 `json:"timestamp"`;};func (_ebc *meteredClient )getStatus ()(meteredStatusResp ,error ){var _gfc meteredStatusResp ;_bcc :=_ebc ._ed +"\u002fm\u0065t\u0065\u0072\u0065\u0064\u002f\u0073\u0074\u0061\u0074\u0075\u0073";var _ae meteredStatusForm ;
_ceb ,_dg :=_dcg .Marshal (_ae );if _dg !=nil {return _gfc ,_dg ;};_dbeb ,_dg :=_faaf (_ceb );if _dg !=nil {return _gfc ,_dg ;};_deb ,_dg :=_gbf .NewRequest ("\u0050\u004f\u0053\u0054",_bcc ,_dbeb );if _dg !=nil {return _gfc ,_dg ;};_deb .Header .Add ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e");
_deb .Header .Add ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");_deb .Header .Add ("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");
_deb .Header .Add ("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y",_ebc ._fbc );_gfce ,_dg :=_ebc ._dbe .Do (_deb );if _dg !=nil {return _gfc ,_dg ;};defer _gfce .Body .Close ();if _gfce .StatusCode !=200{return _gfc ,_afb .Errorf ("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064",_gfce .StatusCode );
};_bfc ,_dg :=_ggd (_gfce );if _dg !=nil {return _gfc ,_dg ;};_dg =_dcg .Unmarshal (_bfc ,&_gfc );if _dg !=nil {return _gfc ,_dg ;};return _gfc ,nil ;};func (_ce *LicenseKey )getExpiryDateToCompare ()_de .Time {if _ce .Trial {return _de .Now ().UTC ();
};return _ab .ReleasedAt ;};type meteredUsageCheckinResp struct{Instance string `json:"inst"`;Next string `json:"next"`;Success bool `json:"success"`;Message string `json:"message"`;RemainingDocs int `json:"rd"`;LimitDocs bool `json:"ld"`;};func (_eaf *LicenseKey )Validate ()error {if _eaf ._da {return nil ;
};if len (_eaf .LicenseId )< 10{return _afb .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020L\u0069\u0063\u0065n\u0073e\u0020\u0049\u0064");};if len (_eaf .CustomerId )< 10{return _afb .Errorf ("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020C\u0075\u0073\u0074\u006f\u006d\u0065\u0072 \u0049\u0064");
};if len (_eaf .CustomerName )< 1{return _afb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043u\u0073\u0074\u006f\u006d\u0065\u0072\u0020\u004e\u0061\u006d\u0065");};if _ad .After (_eaf .CreatedAt ){return _afb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064 \u0041\u0074\u0020\u0069\u0073 \u0069\u006ev\u0061\u006c\u0069\u0064");
};if _eaf .ExpiresAt ==nil {_bdf :=_eaf .CreatedAt .AddDate (1,0,0);if _eae .After (_bdf ){_bdf =_eae ;};_eaf .ExpiresAt =&_bdf ;};if _eaf .CreatedAt .After (*_eaf .ExpiresAt ){return _afb .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0041\u0074 \u0063a\u006e\u006e\u006f\u0074 \u0062\u0065 \u0047\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0045\u0078\u0070\u0069\u0072\u0065\u0073\u0020\u0041\u0074");
};if _eaf .isExpired (){_acgg :="\u0054\u0068\u0065\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0065\u0078\u0070\u0069r\u0065\u0064\u002e\u000a"+"\u0059o\u0075\u0020\u006d\u0061y\u0020n\u0065\u0065\u0064\u0020\u0074\u006f\u0020\u0075\u0070d\u0061\u0074\u0065\u0020\u0074\u0068\u0065\u0020l\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020t\u006f\u0020\u0074\u0068\u0065\u0020\u006e\u0065\u0077\u0065s\u0074\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0072\u0020\u0079o\u0075\u0072\u0020\u006f\u0072\u0067\u0061\u006e\u0069\u007a\u0061\u0074i\u006fn\u002e\u000a"+"\u0054o\u0020\u0066\u0069\u006ed y\u006f\u0075\u0072\u0020n\u0065\u0077\u0065\u0073\u0074\u0020\u006c\u0069\u0063\u0065n\u0073\u0065\u0020\u006b\u0065\u0079\u002c\u0020\u0067\u006f\u0020\u0074\u006f\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002f\u0063l\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064oc\u002e\u0069\u006f \u0061\u006e\u0064\u0020\u0067o\u0020t\u006f\u0020\u0074\u0068\u0065\u0020\u006c\u0069\u0063e\u006e\u0073\u0065\u0020\u006d\u0065\u006e\u0075\u002e";
return _afb .Errorf ("\u0069\u006e\u0076\u0061li\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0025\u0073",_acgg );};if len (_eaf .CreatorName )< 1{return _afb .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020na\u006d\u0065");
};if len (_eaf .CreatorEmail )< 1{return _afb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043r\u0065\u0061\u0074\u006f\u0072\u0020\u0065\u006d\u0061\u0069\u006c");};if _eaf .CreatedAt .After (_bcb ){if !_eaf .UniPDF {return _afb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020\u0054\u0068\u0069\u0073\u0020\u0055\u006e\u0069\u0044\u006f\u0063\u0020k\u0065\u0079\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d \u0066\u006f\u0072\u0020\u0055\u006e\u0069\u0050\u0044\u0046");
};};return nil ;};func _daa (_gbfg *_gbf .Response )(_be .ReadCloser ,error ){var _ede error ;var _eggb _be .ReadCloser ;switch _f .ToLower (_gbfg .Header .Get ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")){case "\u0067\u007a\u0069\u0070":_eggb ,_ede =_bd .NewReader (_gbfg .Body );
if _ede !=nil {return _eggb ,_ede ;};defer _eggb .Close ();default:_eggb =_gbfg .Body ;};return _eggb ,nil ;};func SetMeteredKey (apiKey string )error {if len (apiKey )==0{_ab .Log .Error ("\u004d\u0065\u0074\u0065\u0072e\u0064\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0041\u0050\u0049 \u004b\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079");
_ab .Log .Error ("\u002d\u0020\u0047\u0072\u0061\u0062\u0020\u006f\u006e\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0072\u0065\u0065\u0020\u0054\u0069\u0065\u0072\u0020\u0061t\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002f\u0063\u006c\u006fud\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f");
return _afb .Errorf ("\u006de\u0074\u0065\u0072e\u0064\u0020\u006ci\u0063en\u0073\u0065\u0020\u0061\u0070\u0069\u0020k\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079\u003a\u0020\u0063\u0072\u0065\u0061\u0074\u0065 o\u006ee\u0020\u0061\u0074\u0020\u0068\u0074t\u0070\u0073\u003a\u002f\u002fc\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063.\u0069\u006f");
};if _efg !=nil &&(_efg ._da ||_efg .Tier !=LicenseTierUnlicensed ){_ab .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0043\u0061\u006e\u006eo\u0074 \u0073\u0065\u0074\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0074\u0077\u0069c\u0065\u0020\u002d\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u006a\u0075\u0073\u0074\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069z\u0065\u0020\u006f\u006e\u0063\u0065");
return _af .New ("\u006c\u0069\u0063en\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0061\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0073\u0065\u0074");};_baf :=_eafd ();_baf ._fbc =apiKey ;_gag ,_fce :=_baf .getStatus ();if _fce !=nil {return _fce ;};if !_gag .Valid {return _af .New ("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064");
};_dgf :=&LicenseKey {_da :true ,_geeb :apiKey ,_fd :true };_efg =_dgf ;return nil ;};type defaultStateHolder struct{};var _ebg map[string ]struct{};func _abf (_fc string )(LicenseKey ,error ){var _ge LicenseKey ;_gee ,_ffc :=_bfe (_dd ,_cb ,_fc );if _ffc !=nil {return _ge ,_ffc ;
};_gbda ,_ffc :=_gbd (_fbaf ,_gee );if _ffc !=nil {return _ge ,_ffc ;};_ffc =_dcg .Unmarshal (_gbda ,&_ge );if _ffc !=nil {return _ge ,_ffc ;};_ge .CreatedAt =_de .Unix (_ge .CreatedAtInt ,0);if _ge .ExpiresAtInt > 0{_agc :=_de .Unix (_ge .ExpiresAtInt ,0);
_ge .ExpiresAt =&_agc ;};return _ge ,nil ;};type meteredStatusResp struct{Valid bool `json:"valid"`;OrgCredits int64 `json:"org_credits"`;OrgUsed int64 `json:"org_used"`;OrgRemaining int64 `json:"org_remaining"`;};type MeteredStatus struct{OK bool ;Credits int64 ;
Used int64 ;};func _faaf (_gaga []byte )(_be .Reader ,error ){_ecab :=new (_dc .Buffer );_abe :=_bd .NewWriter (_ecab );_abe .Write (_gaga );_eff :=_abe .Close ();if _eff !=nil {return nil ,_eff ;};return _ecab ,nil ;};var _ad =_de .Date (2010,1,1,0,0,0,0,_de .UTC );
func _add (_dag string ,_cde string ,_bfgd bool )error {if _efg ==nil {return _af .New ("\u006e\u006f\u0020\u006c\u0069\u0063\u0065\u006e\u0073e\u0020\u006b\u0065\u0079");};if !_efg ._da ||len (_efg ._geeb )==0{return nil ;};if len (_dag )==0&&!_bfgd {return _af .New ("\u0064\u006f\u0063\u004b\u0065\u0079\u0020\u006e\u006ft\u0020\u0073\u0065\u0074");
};_adc .Lock ();defer _adc .Unlock ();if _ebg ==nil {_ebg =map[string ]struct{}{};};if _bed ==nil {_bed =map[string ]int {};};_ebca :=0;_ ,_dgfg :=_ebg [_dag ];if !_dgfg {_ebg [_dag ]=struct{}{};_ebca ++;};if _ebca ==0{return nil ;};_bed [_cde ]++;_eac :=_de .Now ();
_def ,_eacb :=_ade .loadState (_efg ._geeb );if _eacb !=nil {_ab .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_eacb );return _eacb ;};if _def .Usage ==nil {_def .Usage =map[string ]int {};};for _edg ,_fca :=range _bed {_def .Usage [_edg ]+=_fca ;
};_bed =nil ;const _egbg =24*_de .Hour ;const _bega =3*24*_de .Hour ;if len (_def .Instance )==0||_eac .Sub (_def .LastReported )> _egbg ||(_def .LimitDocs &&_def .RemainingDocs <=_def .Docs +int64 (_ebca ))||_bfgd {_age ,_cac :=_gb .Hostname ();if _cac !=nil {return _cac ;
};_adg :=_def .Docs ;_fdf ,_eafff ,_cac :=_ege ();if _cac !=nil {_ab .Log .Debug ("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u006c\u006fc\u0061\u006c\u0020\u0061\u0064\u0064\u0072\u0065\u0073\u0073\u003a\u0020\u0025\u0073",_cac .Error ());
_fdf =append (_fdf ,"\u0069n\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020n\u006ft\u0020a\u0076\u0061\u0069\u006c\u0061\u0062\u006ce");_eafff =append (_eafff ,"\u0069n\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020n\u006ft\u0020a\u0076\u0061\u0069\u006c\u0061\u0062\u006ce");
}else {_gd .Strings (_eafff );_gd .Strings (_fdf );_cef ,_ddga :=_afe ();if _ddga !=nil {return _ddga ;};_gaa :=false ;for _ ,_agf :=range _eafff {if _agf ==_cef .String (){_gaa =true ;};};if !_gaa {_eafff =append (_eafff ,_cef .String ());};};_fee :=_eafd ();
_fee ._fbc =_efg ._geeb ;_adg +=int64 (_ebca );_gec :=meteredUsageCheckinForm {Instance :_def .Instance ,Next :_def .Next ,UsageNumber :int (_adg ),NumFailed :_def .NumErrors ,Hostname :_age ,LocalIP :_f .Join (_eafff ,"\u002c\u0020"),MacAddress :_f .Join (_fdf ,"\u002c\u0020"),Package :"\u0075\u006e\u0069\u0070\u0064\u0066",PackageVersion :_ab .Version ,Usage :_def .Usage ,IsPersistentCache :_efg ._fd ,Timestamp :_eac .Unix ()};
if len (_fdf )==0{_gec .MacAddress ="\u006e\u006f\u006e\u0065";};_edb :=int64 (0);_abg :=_def .NumErrors ;_eef :=_eac ;_fdae :=0;_adce :=_def .LimitDocs ;_bgf ,_cac :=_fee .checkinUsage (_gec );if _cac !=nil {if _eac .Sub (_def .LastReported )> _bega {if !_bgf .Success {return _af .New (_bgf .Message );
};return _af .New ("\u0074\u006f\u006f\u0020\u006c\u006f\u006e\u0067\u0020\u0073\u0069\u006e\u0063\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0073\u0075\u0063\u0063e\u0073\u0073\u0066\u0075\u006c \u0063\u0068e\u0063\u006b\u0069\u006e");};_edb =_adg ;_abg ++;
_eef =_def .LastReported ;}else {_adce =_bgf .LimitDocs ;_fdae =_bgf .RemainingDocs ;_abg =0;};if len (_bgf .Instance )==0{_bgf .Instance =_gec .Instance ;};if len (_bgf .Next )==0{_bgf .Next =_gec .Next ;};_cac =_ade .updateState (_fee ._fbc ,_bgf .Instance ,_bgf .Next ,int (_edb ),_adce ,_fdae ,int (_abg ),_eef ,nil );
if _cac !=nil {return _cac ;};if !_bgf .Success {return _afb .Errorf ("\u0065r\u0072\u006f\u0072\u003a\u0020\u0025s",_bgf .Message );};}else {_eacb =_ade .updateState (_efg ._geeb ,_def .Instance ,_def .Next ,int (_def .Docs )+_ebca ,_def .LimitDocs ,int (_def .RemainingDocs ),int (_def .NumErrors ),_def .LastReported ,_def .Usage );
if _eacb !=nil {return _eacb ;};};return nil ;};func _ggd (_cbc *_gbf .Response )([]byte ,error ){var _ecga []byte ;_abc ,_adb :=_daa (_cbc );if _adb !=nil {return _ecga ,_adb ;};return _e .ReadAll (_abc );};func _gbd (_db string ,_bfgb string )([]byte ,error ){var (_faa int ;
_cff string ;);for _ ,_cff =range []string {"\u000a\u002b\u000a","\u000d\u000a\u002b\r\u000a","\u0020\u002b\u0020"}{if _faa =_f .Index (_bfgb ,_cff );_faa !=-1{break ;};};if _faa ==-1{return nil ,_afb .Errorf ("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u002c \u0073i\u0067n\u0061t\u0075\u0072\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u006f\u0072");
};_bae :=_bfgb [:_faa ];_bgb :=_faa +len (_cff );_fg :=_bfgb [_bgb :];if _bae ==""||_fg ==""{return nil ,_afb .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0069n\u0070\u0075\u0074,\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0069\u0067\u0069n\u0061\u006c \u006f\u0072\u0020\u0073\u0069\u0067n\u0061\u0074u\u0072\u0065");
};_gbdg ,_ef :=_fa .StdEncoding .DecodeString (_bae );if _ef !=nil {return nil ,_afb .Errorf ("\u0069\u006e\u0076\u0061li\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0072\u0069\u0067\u0069\u006ea\u006c");};_acg ,_ef :=_fa .StdEncoding .DecodeString (_fg );
if _ef !=nil {return nil ,_afb .Errorf ("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_cfg ,_ :=_aa .Decode ([]byte (_db ));if _cfg ==nil {return nil ,_afb .Errorf ("\u0050\u0075\u0062\u004b\u0065\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064");
};_cfa ,_ef :=_ff .ParsePKIXPublicKey (_cfg .Bytes );if _ef !=nil {return nil ,_ef ;};_ea :=_cfa .(*_ba .PublicKey );if _ea ==nil {return nil ,_afb .Errorf ("\u0050u\u0062\u004b\u0065\u0079\u0020\u0063\u006f\u006e\u0076\u0065\u0072s\u0069\u006f\u006e\u0020\u0066\u0061\u0069\u006c\u0065\u0064");
};_eg :=_ca .New ();_eg .Write (_gbdg );_ag :=_eg .Sum (nil );_ef =_ba .VerifyPKCS1v15 (_ea ,_g .SHA512 ,_ag ,_acg );if _ef !=nil {return nil ,_ef ;};return _gbdg ,nil ;};func _bfe (_cgg string ,_fgd string ,_acd string )(string ,error ){_ga :=_f .Index (_acd ,_cgg );
if _ga ==-1{return "",_afb .Errorf ("\u0068\u0065a\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_dbb :=_f .Index (_acd ,_fgd );if _dbb ==-1{return "",_afb .Errorf ("\u0066\u006fo\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");
};_cgge :=_ga +len (_cgg )+1;return _acd [_cgge :_dbb -1],nil ;};func _ege ()([]string ,[]string ,error ){_gabc ,_bab :=_d .Interfaces ();if _bab !=nil {return nil ,nil ,_bab ;};var _bbg []string ;var _edbb []string ;for _ ,_faag :=range _gabc {if _faag .Flags &_d .FlagUp ==0||_dc .Equal (_faag .HardwareAddr ,nil ){continue ;
};_cacf ,_cca :=_faag .Addrs ();if _cca !=nil {return nil ,nil ,_cca ;};_eggc :=0;for _ ,_bag :=range _cacf {var _adgg _d .IP ;switch _gbc :=_bag .(type ){case *_d .IPNet :_adgg =_gbc .IP ;case *_d .IPAddr :_adgg =_gbc .IP ;};if _adgg .IsLoopback (){continue ;
};if _adgg .To4 ()==nil {continue ;};_edbb =append (_edbb ,_adgg .String ());_eggc ++;};_ced :=_faag .HardwareAddr .String ();if _ced !=""&&_eggc > 0{_bbg =append (_bbg ,_ced );};};return _bbg ,_edbb ,nil ;};var _bcb =_de .Date (2019,6,6,0,0,0,0,_de .UTC );
const _fbaf ="\u000a\u002d\u002d\u002d\u002d\u002d\u0042\u0045\u0047\u0049\u004e \u0050\u0055\u0042\u004c\u0049\u0043\u0020\u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\u000a\u004d\u0049I\u0042\u0049\u006a\u0041NB\u0067\u006b\u0071\u0068\u006b\u0069G\u0039\u0077\u0030\u0042\u0041\u0051\u0045\u0046A\u0041\u004f\u0043\u0041\u0051\u0038\u0041\u004d\u0049\u0049\u0042\u0043\u0067\u004b\u0043\u0041\u0051\u0045A\u006dF\u0055\u0069\u0079\u0064\u0037\u0062\u0035\u0058\u006a\u0070\u006b\u0050\u0035\u0052\u0061\u0070\u0034\u0077\u000a\u0044\u0063\u0031d\u0079\u007a\u0049\u0051\u0034\u004c\u0065\u006b\u0078\u0072\u0076\u0079\u0074\u006e\u0045\u004d\u0070\u004e\u0055\u0062\u006f\u0036i\u0041\u0037\u0034\u0056\u0038\u0072\u0075\u005a\u004f\u0076\u0072\u0053\u0063\u0073\u0066\u0032\u0051\u0065\u004e9\u002f\u0071r\u0055\u0047\u0038\u0071\u0045\u0062\u0055\u0057\u0064\u006f\u0045\u0059\u0071+\u000a\u006f\u0074\u0046\u004e\u0041\u0046N\u0078\u006c\u0047\u0062\u0078\u0062\u0044\u0048\u0063\u0064\u0047\u0056\u0061\u004d\u0030\u004f\u0058\u0064\u0058g\u0044y\u004c5\u0061\u0049\u0045\u0061\u0067\u004c\u0030\u0063\u0035\u0070\u0077\u006a\u0049\u0064\u0050G\u0049\u006e\u0034\u0036\u0066\u0037\u0038\u0065\u004d\u004a\u002b\u004a\u006b\u0064\u0063\u0070\u0044\n\u0044\u004a\u0061\u0071\u0059\u0058d\u0072\u007a5\u004b\u0065\u0073\u0068\u006aS\u0069\u0049\u0061\u0061\u0037\u006d\u0065\u006e\u0042\u0049\u0041\u0058\u0053\u0034\u0055\u0046\u0078N\u0066H\u0068\u004e\u0030\u0048\u0043\u0059\u005a\u0059\u0071\u0051\u0047\u0037\u0062K+\u0073\u0035\u0072R\u0048\u006f\u006e\u0079\u0064\u004eW\u0045\u0047\u000a\u0048\u0038M\u0079\u0076\u00722\u0070\u0079\u0061\u0032K\u0072\u004d\u0075m\u0066\u006d\u0041\u0078\u0055\u0042\u0036\u0066\u0065\u006e\u0043\u002f4\u004f\u0030\u0057\u00728\u0067\u0066\u0050\u004f\u0055\u0038R\u0069\u0074\u006d\u0062\u0044\u0076\u0051\u0050\u0049\u0052\u0058\u004fL\u0034\u0076\u0054B\u0072\u0042\u0064\u0062a\u0041\u000a9\u006e\u0077\u004e\u0050\u002b\u0069\u002f\u002f\u0032\u0030\u004d\u00542\u0062\u0078\u006d\u0065\u0057\u0042\u002b\u0067\u0070\u0063\u0045\u0068G\u0070\u0058\u005a7\u0033\u0033\u0061\u007a\u0051\u0078\u0072\u0043\u0033\u004a\u0034\u0076\u0033C\u005a\u006d\u0045\u004eS\u0074\u0044\u004b\u002f\u004b\u0044\u0053\u0050\u004b\u0055\u0047\u0066\u00756\u000a\u0066\u0077I\u0044\u0041\u0051\u0041\u0042\u000a\u002d\u002d\u002d\u002d\u002dE\u004e\u0044\u0020\u0050\u0055\u0042\u004c\u0049\u0043 \u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\n";
const (_dd ="\u002d\u002d\u002d--\u0042\u0045\u0047\u0049\u004e\u0020\u0055\u004e\u0049D\u004fC\u0020L\u0049C\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d";_cb ="\u002d\u002d\u002d\u002d\u002d\u0045\u004e\u0044\u0020\u0055\u004e\u0049\u0044\u004f\u0043 \u004cI\u0043\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d";
);func _dfb (_effd ,_bgd []byte )([]byte ,error ){_fff :=make ([]byte ,_fa .URLEncoding .DecodedLen (len (_bgd )));_eacg ,_cfe :=_fa .URLEncoding .Decode (_fff ,_bgd );if _cfe !=nil {return nil ,_cfe ;};_fff =_fff [:_eacg ];_aae ,_cfe :=_ec .NewCipher (_effd );
if _cfe !=nil {return nil ,_cfe ;};if len (_fff )< _ec .BlockSize {return nil ,_af .New ("c\u0069p\u0068\u0065\u0072\u0074\u0065\u0078\u0074\u0020t\u006f\u006f\u0020\u0073ho\u0072\u0074");};_ddf :=_fff [:_ec .BlockSize ];_fff =_fff [_ec .BlockSize :];_fef :=_cc .NewCFBDecrypter (_aae ,_ddf );
_fef .XORKeyStream (_fff ,_fff );return _fff ,nil ;};func SetLicenseKey (content string ,customerName string )error {_gda ,_gfdfd :=_abf (content );if _gfdfd !=nil {_ab .Log .Error ("\u004c\u0069c\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u0072\u0072\u006f\u0072: \u0025\u0076",_gfdfd );
return _gfdfd ;};if !_f .EqualFold (_gda .CustomerName ,customerName ){_ab .Log .Error ("L\u0069ce\u006es\u0065 \u0063\u006f\u0064\u0065\u0020i\u0073\u0073\u0075e\u0020\u002d\u0020\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006da\u0074\u0063\u0068, e\u0078\u0070\u0065\u0063\u0074\u0065d\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067o\u0074 \u0027\u0025\u0073\u0027",_gda .CustomerName ,customerName );
return _afb .Errorf ("\u0063\u0075\u0073\u0074\u006fm\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006d\u0061t\u0063\u0068\u002c\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067\u006f\u0074\u0020\u0027\u0025\u0073'",_gda .CustomerName ,customerName );
};_gfdfd =_gda .Validate ();if _gfdfd !=nil {_ab .Log .Error ("\u004c\u0069\u0063\u0065\u006e\u0073e\u0020\u0063\u006f\u0064\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074i\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076",_gfdfd );
return _gfdfd ;};_efg =&_gda ;return nil ;};func _efda (_adca ,_feg string )string {_gff :=[]byte (_adca );_fga :=_c .New (_ee .New ,_gff );_fga .Write ([]byte (_feg ));return _fa .StdEncoding .EncodeToString (_fga .Sum (nil ));};type stateLoader interface{loadState (_ega string )(reportState ,error );
updateState (_efa ,_eeg ,_bee string ,_fdeb int ,_dab bool ,_abfg int ,_dga int ,_daf _de .Time ,_cge map[string ]int )error ;};func (_afa *LicenseKey )isExpired ()bool {return _afa .getExpiryDateToCompare ().After (*_afa .ExpiresAt )};type reportState struct{Instance string `json:"inst"`;
Next string `json:"n"`;Docs int64 `json:"d"`;NumErrors int64 `json:"e"`;LimitDocs bool `json:"ld"`;RemainingDocs int64 `json:"rd"`;LastReported _de .Time `json:"lr"`;LastWritten _de .Time `json:"lw"`;Usage map[string ]int `json:"u"`;};func GetLicenseKey ()*LicenseKey {if _efg ==nil {return nil ;
};_caa :=*_efg ;return &_caa ;};func _afe ()(_d .IP ,error ){_cag ,_bcf :=_d .Dial ("\u0075\u0064\u0070","\u0038\u002e\u0038\u002e\u0038\u002e\u0038\u003a\u0038\u0030");if _bcf !=nil {return nil ,_bcf ;};defer _cag .Close ();_fea :=_cag .LocalAddr ().(*_d .UDPAddr );
return _fea .IP ,nil ;};var _adc =&_a .Mutex {};func (_dcd *LicenseKey )IsLicensed ()bool {return _dcd .Tier !=LicenseTierUnlicensed ||_dcd ._da };type meteredStatusForm struct{};const _agec ="\u0055\u004e\u0049\u0050DF\u005f\u004c\u0049\u0043\u0045\u004e\u0053\u0045\u005f\u0050\u0041\u0054\u0048";
func init (){_eee :=_gb .Getenv (_agec );_egd :=_gb .Getenv (_afac );if len (_eee )==0||len (_egd )==0{return ;};_ffe ,_bcce :=_e .ReadFile (_eee );if _bcce !=nil {_ab .Log .Error ("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0072\u0065ad \u006cic\u0065\u006e\u0073\u0065\u0020\u0063\u006fde\u0020\u0066\u0069\u006c\u0065\u003a\u0020%\u0076",_bcce );
return ;};_bcce =SetLicenseKey (string (_ffe ),_egd );if _bcce !=nil {_ab .Log .Error ("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u006c\u006f\u0061\u0064\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0076",_bcce );
return ;};};func SetMeteredKeyPersistentCache (val bool ){_efg ._fd =val };type meteredClient struct{_ed string ;_fbc string ;_dbe *_gbf .Client ;};var _bed map[string ]int ;func _ecca (_efc ,_agag []byte )([]byte ,error ){_ecdd ,_cab :=_ec .NewCipher (_efc );
if _cab !=nil {return nil ,_cab ;};_ffaa :=make ([]byte ,_ec .BlockSize +len (_agag ));_edd :=_ffaa [:_ec .BlockSize ];if _ ,_fdfb :=_be .ReadFull (_gf .Reader ,_edd );_fdfb !=nil {return nil ,_fdfb ;};_dbg :=_cc .NewCFBEncrypter (_ecdd ,_edd );_dbg .XORKeyStream (_ffaa [_ec .BlockSize :],_agag );
_gdfe :=make ([]byte ,_fa .URLEncoding .EncodedLen (len (_ffaa )));_fa .URLEncoding .Encode (_gdfe ,_ffaa );return _gdfe ,nil ;};func _bg (_bda string ,_fe []byte )(string ,error ){_bc ,_ :=_aa .Decode ([]byte (_bda ));if _bc ==nil {return "",_afb .Errorf ("\u0050\u0072\u0069\u0076\u004b\u0065\u0079\u0020\u0066a\u0069\u006c\u0065\u0064");
};_gdf ,_ac :=_ff .ParsePKCS1PrivateKey (_bc .Bytes );if _ac !=nil {return "",_ac ;};_bf :=_ca .New ();_bf .Write (_fe );_fb :=_bf .Sum (nil );_bfg ,_ac :=_ba .SignPKCS1v15 (_gf .Reader ,_gdf ,_g .SHA512 ,_fb );if _ac !=nil {return "",_ac ;};_cf :=_fa .StdEncoding .EncodeToString (_fe );
_cf +="\u000a\u002b\u000a";_cf +=_fa .StdEncoding .EncodeToString (_bfg );return _cf ,nil ;};func (_gdb defaultStateHolder )loadState (_eaff string )(reportState ,error ){_ecg :=_bbf ();if len (_ecg )==0{return reportState {},_af .New ("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");
};_adf :=_gg .Join (_ecg ,"\u002eu\u006e\u0069\u0064\u006f\u0063");_acdg :=_gb .MkdirAll (_adf ,0777);if _acdg !=nil {return reportState {},_acdg ;};if len (_eaff )< 20{return reportState {},_af .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079");
};_ddg :=[]byte (_eaff );_aef :=_ca .Sum512_256 (_ddg [:20]);_aafc :=_eb .EncodeToString (_aef [:]);_afd :=_gg .Join (_adf ,_aafc );_dcb ,_acdg :=_e .ReadFile (_afd );if _acdg !=nil {if _gb .IsNotExist (_acdg ){return reportState {},nil ;};_ab .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_acdg );
return reportState {},_af .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061");};const _efde ="\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057";_dcb ,_acdg =_dfb ([]byte (_efde ),_dcb );if _acdg !=nil {return reportState {},_acdg ;
};var _bdfb reportState ;_acdg =_dcg .Unmarshal (_dcb ,&_bdfb );if _acdg !=nil {_ab .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0076",_acdg );return reportState {},_af .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061");
};return _bdfb ,nil ;};func GetMeteredState ()(MeteredStatus ,error ){if _efg ==nil {return MeteredStatus {},_af .New ("\u006c\u0069\u0063\u0065ns\u0065\u0020\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");};if !_efg ._da ||len (_efg ._geeb )==0{return MeteredStatus {},_af .New ("\u0061p\u0069 \u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");
};_bcbe ,_ceac :=_ade .loadState (_efg ._geeb );if _ceac !=nil {_ab .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_ceac );return MeteredStatus {},_ceac ;};if _bcbe .Docs > 0{_egf :=_add ("","",true );if _egf !=nil {return MeteredStatus {},_egf ;
};};_adc .Lock ();defer _adc .Unlock ();_aga :=_eafd ();_aga ._fbc =_efg ._geeb ;_dcf ,_ceac :=_aga .getStatus ();if _ceac !=nil {return MeteredStatus {},_ceac ;};if !_dcf .Valid {return MeteredStatus {},_af .New ("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064");
};_abb :=MeteredStatus {OK :true ,Credits :_dcf .OrgCredits ,Used :_dcf .OrgUsed };return _abb ,nil ;};func Track (docKey string ,useKey string )error {return _add (docKey ,useKey ,!_efg ._fd )};