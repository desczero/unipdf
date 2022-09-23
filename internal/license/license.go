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

package license ;import (_da "bytes";_g "compress/gzip";_d "crypto";_ab "crypto/aes";_aeb "crypto/cipher";_cg "crypto/hmac";_dd "crypto/rand";_aag "crypto/rsa";_gb "crypto/sha256";_fc "crypto/sha512";_abc "crypto/x509";_cd "encoding/base64";_ff "encoding/hex";
_gf "encoding/json";_ffg "encoding/pem";_aea "errors";_dag "fmt";_bge "github.com/unidoc/unipdf/v3/common";_a "io";_ga "io/ioutil";_aa "net";_b "net/http";_f "os";_ge "path/filepath";_ae "sort";_e "strings";_cf "sync";_bg "time";);const (_gfd ="\u002d\u002d\u002d--\u0042\u0045\u0047\u0049\u004e\u0020\u0055\u004e\u0049D\u004fC\u0020L\u0049C\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d";
_df ="\u002d\u002d\u002d\u002d\u002d\u0045\u004e\u0044\u0020\u0055\u004e\u0049\u0044\u004f\u0043 \u004cI\u0043\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d";);func (_eee *meteredClient )checkinUsage (_gefb meteredUsageCheckinForm )(meteredUsageCheckinResp ,error ){_gefb .Package ="\u0075\u006e\u0069\u0070\u0064\u0066";
_gefb .PackageVersion =_bge .Version ;var _fe meteredUsageCheckinResp ;_adf :=_eee ._fcc +"\u002f\u006d\u0065\u0074er\u0065\u0064\u002f\u0075\u0073\u0061\u0067\u0065\u005f\u0063\u0068\u0065\u0063\u006bi\u006e";_gge ,_feg :=_gf .Marshal (_gefb );if _feg !=nil {return _fe ,_feg ;
};_gfaf ,_feg :=_aed (_gge );if _feg !=nil {return _fe ,_feg ;};_eae ,_feg :=_b .NewRequest ("\u0050\u004f\u0053\u0054",_adf ,_gfaf );if _feg !=nil {return _fe ,_feg ;};_eae .Header .Add ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e");
_eae .Header .Add ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");_eae .Header .Add ("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");
_eae .Header .Add ("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y",_eee ._bd );_dcb ,_feg :=_eee ._fg .Do (_eae );if _feg !=nil {return _fe ,_feg ;};defer _dcb .Body .Close ();if _dcb .StatusCode !=200{_gad ,_edc :=_edd (_dcb );if _edc !=nil {return _fe ,_edc ;
};_edc =_gf .Unmarshal (_gad ,&_fe );if _edc !=nil {return _fe ,_edc ;};return _fe ,_dag .Errorf ("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064",_dcb .StatusCode );
};_efa :=_dcb .Header .Get ("\u0058\u002d\u0055\u0043\u002d\u0053\u0069\u0067\u006ea\u0074\u0075\u0072\u0065");_fff :=_gga (_gefb .MacAddress ,string (_gge ));if _fff !=_efa {_bge .Log .Error ("I\u006e\u0076\u0061l\u0069\u0064\u0020\u0072\u0065\u0073\u0070\u006f\u006e\u0073\u0065\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u002c\u0020\u0073\u0065t\u0020\u0074\u0068e\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0073\u0065\u0072\u0076e\u0072\u0020\u0074\u006f \u0068\u0074\u0074\u0070s\u003a\u002f\u002f\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069o\u002f\u0061\u0070\u0069");
return _fe ,_aea .New ("\u0066\u0061\u0069l\u0065\u0064\u0020\u0074\u006f\u0020\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0065\u0072\u0076\u0065\u0072 \u0072\u0065\u0073\u0070\u006f\u006e\u0073\u0065");
};_fcb ,_feg :=_edd (_dcb );if _feg !=nil {return _fe ,_feg ;};_feg =_gf .Unmarshal (_fcb ,&_fe );if _feg !=nil {return _fe ,_feg ;};return _fe ,nil ;};func _caba ()*meteredClient {_dff :=meteredClient {_fcc :"h\u0074\u0074\u0070\u0073\u003a\u002f/\u0063\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069d\u006f\u0063\u002ei\u006f/\u0061\u0070\u0069",_fg :&_b .Client {Timeout :30*_bg .Second }};
if _ee :=_f .Getenv ("\u0055N\u0049\u0044\u004f\u0043_\u004c\u0049\u0043\u0045\u004eS\u0045_\u0053E\u0052\u0056\u0045\u0052\u005f\u0055\u0052L");_e .HasPrefix (_ee ,"\u0068\u0074\u0074\u0070"){_dff ._fcc =_ee ;};return &_dff ;};func MakeUnlicensedKey ()*LicenseKey {_ddca :=LicenseKey {};
_ddca .CustomerName ="\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";_ddca .Tier =LicenseTierUnlicensed ;_ddca .CreatedAt =_bg .Now ().UTC ();_ddca .CreatedAtInt =_ddca .CreatedAt .Unix ();return &_ddca ;};func GetMeteredState ()(MeteredStatus ,error ){if _gcg ==nil {return MeteredStatus {},_aea .New ("\u006c\u0069\u0063\u0065ns\u0065\u0020\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");
};if !_gcg ._dcc ||len (_gcg ._ac )==0{return MeteredStatus {},_aea .New ("\u0061p\u0069 \u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");};_cda ,_ddg :=_fcf .loadState (_gcg ._ac );if _ddg !=nil {_bge .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_ddg );
return MeteredStatus {},_ddg ;};if _cda .Docs > 0{_deda :=_ddbd ("","",true );if _deda !=nil {return MeteredStatus {},_deda ;};};_bdd .Lock ();defer _bdd .Unlock ();_ggf :=_caba ();_ggf ._bd =_gcg ._ac ;_ega ,_ddg :=_ggf .getStatus ();if _ddg !=nil {return MeteredStatus {},_ddg ;
};if !_ega .Valid {return MeteredStatus {},_aea .New ("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064");};_cfbb :=MeteredStatus {OK :true ,Credits :_ega .OrgCredits ,Used :_ega .OrgUsed };return _cfbb ,nil ;};type meteredUsageCheckinForm struct{Instance string `json:"inst"`;
Next string `json:"next"`;UsageNumber int `json:"usage_number"`;NumFailed int64 `json:"num_failed"`;Hostname string `json:"hostname"`;LocalIP string `json:"local_ip"`;MacAddress string `json:"mac_address"`;Package string `json:"package"`;PackageVersion string `json:"package_version"`;
Usage map[string ]int `json:"u"`;IsPersistentCache bool `json:"is_persistent_cache"`;Timestamp int64 `json:"timestamp"`;};func (_cgg *LicenseKey )isExpired ()bool {return _cgg .getExpiryDateToCompare ().After (*_cgg .ExpiresAt )};func (_gc *LicenseKey )IsLicensed ()bool {return _gc .Tier !=LicenseTierUnlicensed ||_gc ._dcc };
type meteredUsageCheckinResp struct{Instance string `json:"inst"`;Next string `json:"next"`;Success bool `json:"success"`;Message string `json:"message"`;RemainingDocs int `json:"rd"`;LimitDocs bool `json:"ld"`;};type stateLoader interface{loadState (_ddf string )(reportState ,error );
updateState (_bda ,_dcg ,_cbe string ,_dgd int ,_cdd bool ,_ebg int ,_gag int ,_fbe _bg .Time ,_feb map[string ]int )error ;};func Track (docKey string ,useKey string )error {return _ddbd (docKey ,useKey ,!_gcg ._ded )};var _ce map[string ]int ;func (_gabb defaultStateHolder )updateState (_fec ,_dbaf ,_acc string ,_bfe int ,_aac bool ,_cae int ,_adg int ,_gdc _bg .Time ,_ebc map[string ]int )error {_eag :=_cbee ();
if len (_eag )==0{return _aea .New ("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");};_efd :=_ge .Join (_eag ,"\u002eu\u006e\u0069\u0064\u006f\u0063");_fcbd :=_f .MkdirAll (_efd ,0777);if _fcbd !=nil {return _fcbd ;
};if len (_fec )< 20{return _aea .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079");};_ddd :=[]byte (_fec );_dcgd :=_fc .Sum512_256 (_ddd [:20]);_dad :=_ff .EncodeToString (_dcgd [:]);_fdgc :=_ge .Join (_efd ,_dad );var _efc reportState ;
_efc .Docs =int64 (_bfe );_efc .NumErrors =int64 (_adg );_efc .LimitDocs =_aac ;_efc .RemainingDocs =int64 (_cae );_efc .LastWritten =_bg .Now ().UTC ();_efc .LastReported =_gdc ;_efc .Instance =_dbaf ;_efc .Next =_acc ;_efc .Usage =_ebc ;_cgd ,_fcbd :=_gf .Marshal (_efc );
if _fcbd !=nil {return _fcbd ;};const _bgd ="\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057";_cgd ,_fcbd =_bad ([]byte (_bgd ),_cgd );if _fcbd !=nil {return _fcbd ;};_fcbd =_ga .WriteFile (_fdgc ,_cgd ,0600);if _fcbd !=nil {return _fcbd ;
};return nil ;};const _dbe ="\u000a\u002d\u002d\u002d\u002d\u002d\u0042\u0045\u0047\u0049\u004e \u0050\u0055\u0042\u004c\u0049\u0043\u0020\u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\u000a\u004d\u0049I\u0042\u0049\u006a\u0041NB\u0067\u006b\u0071\u0068\u006b\u0069G\u0039\u0077\u0030\u0042\u0041\u0051\u0045\u0046A\u0041\u004f\u0043\u0041\u0051\u0038\u0041\u004d\u0049\u0049\u0042\u0043\u0067\u004b\u0043\u0041\u0051\u0045A\u006dF\u0055\u0069\u0079\u0064\u0037\u0062\u0035\u0058\u006a\u0070\u006b\u0050\u0035\u0052\u0061\u0070\u0034\u0077\u000a\u0044\u0063\u0031d\u0079\u007a\u0049\u0051\u0034\u004c\u0065\u006b\u0078\u0072\u0076\u0079\u0074\u006e\u0045\u004d\u0070\u004e\u0055\u0062\u006f\u0036i\u0041\u0037\u0034\u0056\u0038\u0072\u0075\u005a\u004f\u0076\u0072\u0053\u0063\u0073\u0066\u0032\u0051\u0065\u004e9\u002f\u0071r\u0055\u0047\u0038\u0071\u0045\u0062\u0055\u0057\u0064\u006f\u0045\u0059\u0071+\u000a\u006f\u0074\u0046\u004e\u0041\u0046N\u0078\u006c\u0047\u0062\u0078\u0062\u0044\u0048\u0063\u0064\u0047\u0056\u0061\u004d\u0030\u004f\u0058\u0064\u0058g\u0044y\u004c5\u0061\u0049\u0045\u0061\u0067\u004c\u0030\u0063\u0035\u0070\u0077\u006a\u0049\u0064\u0050G\u0049\u006e\u0034\u0036\u0066\u0037\u0038\u0065\u004d\u004a\u002b\u004a\u006b\u0064\u0063\u0070\u0044\n\u0044\u004a\u0061\u0071\u0059\u0058d\u0072\u007a5\u004b\u0065\u0073\u0068\u006aS\u0069\u0049\u0061\u0061\u0037\u006d\u0065\u006e\u0042\u0049\u0041\u0058\u0053\u0034\u0055\u0046\u0078N\u0066H\u0068\u004e\u0030\u0048\u0043\u0059\u005a\u0059\u0071\u0051\u0047\u0037\u0062K+\u0073\u0035\u0072R\u0048\u006f\u006e\u0079\u0064\u004eW\u0045\u0047\u000a\u0048\u0038M\u0079\u0076\u00722\u0070\u0079\u0061\u0032K\u0072\u004d\u0075m\u0066\u006d\u0041\u0078\u0055\u0042\u0036\u0066\u0065\u006e\u0043\u002f4\u004f\u0030\u0057\u00728\u0067\u0066\u0050\u004f\u0055\u0038R\u0069\u0074\u006d\u0062\u0044\u0076\u0051\u0050\u0049\u0052\u0058\u004fL\u0034\u0076\u0054B\u0072\u0042\u0064\u0062a\u0041\u000a9\u006e\u0077\u004e\u0050\u002b\u0069\u002f\u002f\u0032\u0030\u004d\u00542\u0062\u0078\u006d\u0065\u0057\u0042\u002b\u0067\u0070\u0063\u0045\u0068G\u0070\u0058\u005a7\u0033\u0033\u0061\u007a\u0051\u0078\u0072\u0043\u0033\u004a\u0034\u0076\u0033C\u005a\u006d\u0045\u004eS\u0074\u0044\u004b\u002f\u004b\u0044\u0053\u0050\u004b\u0055\u0047\u0066\u00756\u000a\u0066\u0077I\u0044\u0041\u0051\u0041\u0042\u000a\u002d\u002d\u002d\u002d\u002dE\u004e\u0044\u0020\u0050\u0055\u0042\u004c\u0049\u0043 \u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\n";
func (_ecd *LicenseKey )ToString ()string {if _ecd ._dcc {return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e";};_aec :=_dag .Sprintf ("\u004ci\u0063e\u006e\u0073\u0065\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a",_ecd .LicenseId );
_aec +=_dag .Sprintf ("\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a",_ecd .CustomerId );_aec +=_dag .Sprintf ("\u0043u\u0073t\u006f\u006d\u0065\u0072\u0020N\u0061\u006de\u003a\u0020\u0025\u0073\u000a",_ecd .CustomerName );
_aec +=_dag .Sprintf ("\u0054i\u0065\u0072\u003a\u0020\u0025\u0073\n",_ecd .Tier );_aec +=_dag .Sprintf ("\u0043r\u0065a\u0074\u0065\u0064\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a",_bge .UtcTimeFormat (_ecd .CreatedAt ));if _ecd .ExpiresAt ==nil {_aec +="\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041t\u003a\u0020N\u0065\u0076\u0065\u0072\u000a";
}else {_aec +=_dag .Sprintf ("\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a",_bge .UtcTimeFormat (*_ecd .ExpiresAt ));};_aec +=_dag .Sprintf ("\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u003a\u0020\u0025\u0073\u0020<\u0025\u0073\u003e\u000a",_ecd .CreatorName ,_ecd .CreatorEmail );
return _aec ;};func (_ea *LicenseKey )Validate ()error {if _ea ._dcc {return nil ;};if len (_ea .LicenseId )< 10{return _dag .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020L\u0069\u0063\u0065n\u0073e\u0020\u0049\u0064");
};if len (_ea .CustomerId )< 10{return _dag .Errorf ("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020C\u0075\u0073\u0074\u006f\u006d\u0065\u0072 \u0049\u0064");};if len (_ea .CustomerName )< 1{return _dag .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043u\u0073\u0074\u006f\u006d\u0065\u0072\u0020\u004e\u0061\u006d\u0065");
};if _aab .After (_ea .CreatedAt ){return _dag .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064 \u0041\u0074\u0020\u0069\u0073 \u0069\u006ev\u0061\u006c\u0069\u0064");
};if _ea .ExpiresAt ==nil {_gda :=_ea .CreatedAt .AddDate (1,0,0);if _dagb .After (_gda ){_gda =_dagb ;};_ea .ExpiresAt =&_gda ;};if _ea .CreatedAt .After (*_ea .ExpiresAt ){return _dag .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0041\u0074 \u0063a\u006e\u006e\u006f\u0074 \u0062\u0065 \u0047\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0045\u0078\u0070\u0069\u0072\u0065\u0073\u0020\u0041\u0074");
};if _ea .isExpired (){return _dag .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020l\u0069\u0063\u0065ns\u0065\u003a\u0020\u0054\u0068\u0065 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0065\u0078\u0070\u0069r\u0065\u0064");
};if len (_ea .CreatorName )< 1{return _dag .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020na\u006d\u0065");};if len (_ea .CreatorEmail )< 1{return _dag .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043r\u0065\u0061\u0074\u006f\u0072\u0020\u0065\u006d\u0061\u0069\u006c");
};if _ea .CreatedAt .After (_aebe ){if !_ea .UniPDF {return _dag .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020\u0054\u0068\u0069\u0073\u0020\u0055\u006e\u0069\u0044\u006f\u0063\u0020k\u0065\u0079\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d \u0066\u006f\u0072\u0020\u0055\u006e\u0069\u0050\u0044\u0046");
};};return nil ;};type defaultStateHolder struct{};type meteredStatusResp struct{Valid bool `json:"valid"`;OrgCredits int64 `json:"org_credits"`;OrgUsed int64 `json:"org_used"`;OrgRemaining int64 `json:"org_remaining"`;};func TrackUse (useKey string ){if _gcg ==nil {return ;
};if !_gcg ._dcc ||len (_gcg ._ac )==0{return ;};if len (useKey )==0{return ;};_bdd .Lock ();defer _bdd .Unlock ();if _ce ==nil {_ce =map[string ]int {};};_ce [useKey ]++;};func _cbbf (_bdag ,_efgf []byte )([]byte ,error ){_gbd :=make ([]byte ,_cd .URLEncoding .DecodedLen (len (_efgf )));
_aaa ,_ffd :=_cd .URLEncoding .Decode (_gbd ,_efgf );if _ffd !=nil {return nil ,_ffd ;};_gbd =_gbd [:_aaa ];_cgc ,_ffd :=_ab .NewCipher (_bdag );if _ffd !=nil {return nil ,_ffd ;};if len (_gbd )< _ab .BlockSize {return nil ,_aea .New ("c\u0069p\u0068\u0065\u0072\u0074\u0065\u0078\u0074\u0020t\u006f\u006f\u0020\u0073ho\u0072\u0074");
};_adfg :=_gbd [:_ab .BlockSize ];_gbd =_gbd [_ab .BlockSize :];_aeec :=_aeb .NewCFBDecrypter (_cgc ,_adfg );_aeec .XORKeyStream (_gbd ,_gbd );return _gbd ,nil ;};func _gga (_dagg ,_gbed string )string {_bdb :=[]byte (_dagg );_afg :=_cg .New (_gb .New ,_bdb );
_afg .Write ([]byte (_gbed ));return _cd .StdEncoding .EncodeToString (_afg .Sum (nil ));};func _ec (_bbd string )(LicenseKey ,error ){var _gbf LicenseKey ;_ccf ,_gd :=_dec (_gfd ,_df ,_bbd );if _gd !=nil {return _gbf ,_gd ;};_abf ,_gd :=_dfc (_dbe ,_ccf );
if _gd !=nil {return _gbf ,_gd ;};_gd =_gf .Unmarshal (_abf ,&_gbf );if _gd !=nil {return _gbf ,_gd ;};_gbf .CreatedAt =_bg .Unix (_gbf .CreatedAtInt ,0);if _gbf .ExpiresAtInt > 0{_bf :=_bg .Unix (_gbf .ExpiresAtInt ,0);_gbf .ExpiresAt =&_bf ;};return _gbf ,nil ;
};func _dec (_ag string ,_ggb string ,_gef string )(string ,error ){_ad :=_e .Index (_gef ,_ag );if _ad ==-1{return "",_dag .Errorf ("\u0068\u0065a\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_cb :=_e .Index (_gef ,_ggb );
if _cb ==-1{return "",_dag .Errorf ("\u0066\u006fo\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_db :=_ad +len (_ag )+1;return _gef [_db :_cb -1],nil ;};func _bad (_fba ,_cade []byte )([]byte ,error ){_fgf ,_bgc :=_ab .NewCipher (_fba );
if _bgc !=nil {return nil ,_bgc ;};_edb :=make ([]byte ,_ab .BlockSize +len (_cade ));_ccd :=_edb [:_ab .BlockSize ];if _ ,_gaf :=_a .ReadFull (_dd .Reader ,_ccd );_gaf !=nil {return nil ,_gaf ;};_fgfb :=_aeb .NewCFBEncrypter (_fgf ,_ccd );_fgfb .XORKeyStream (_edb [_ab .BlockSize :],_cade );
_bfb :=make ([]byte ,_cd .URLEncoding .EncodedLen (len (_edb )));_cd .URLEncoding .Encode (_bfb ,_edb );return _bfb ,nil ;};var _aab =_bg .Date (2010,1,1,0,0,0,0,_bg .UTC );type meteredClient struct{_fcc string ;_bd string ;_fg *_b .Client ;};func _fce ()([]string ,[]string ,error ){_fegg ,_gdg :=_aa .Interfaces ();
if _gdg !=nil {return nil ,nil ,_gdg ;};var _edce []string ;var _abgb []string ;for _ ,_eeeb :=range _fegg {if _eeeb .Flags &_aa .FlagUp ==0||_da .Equal (_eeeb .HardwareAddr ,nil ){continue ;};_cbdd ,_aagb :=_eeeb .Addrs ();if _aagb !=nil {return nil ,nil ,_aagb ;
};_ecc :=0;for _ ,_bcb :=range _cbdd {var _ecdd _aa .IP ;switch _cfd :=_bcb .(type ){case *_aa .IPNet :_ecdd =_cfd .IP ;case *_aa .IPAddr :_ecdd =_cfd .IP ;};if _ecdd .IsLoopback (){continue ;};if _ecdd .To4 ()==nil {continue ;};_abgb =append (_abgb ,_ecdd .String ());
_ecc ++;};_ggea :=_eeeb .HardwareAddr .String ();if _ggea !=""&&_ecc > 0{_edce =append (_edce ,_ggea );};};return _edce ,_abgb ,nil ;};func (_caa *LicenseKey )TypeToString ()string {if _caa ._dcc {return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e";
};if _caa .Tier ==LicenseTierUnlicensed {return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";};if _caa .Tier ==LicenseTierCommunity {return "\u0041\u0047PL\u0076\u0033\u0020O\u0070\u0065\u006e\u0020Sou\u0072ce\u0020\u0043\u006f\u006d\u006d\u0075\u006eit\u0079\u0020\u004c\u0069\u0063\u0065\u006es\u0065";
};if _caa .Tier ==LicenseTierIndividual ||_caa .Tier =="\u0069\u006e\u0064i\u0065"{return "\u0043\u006f\u006dm\u0065\u0072\u0063\u0069a\u006c\u0020\u004c\u0069\u0063\u0065\u006es\u0065\u0020\u002d\u0020\u0049\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c";
};return "\u0043\u006fm\u006d\u0065\u0072\u0063\u0069\u0061\u006c\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u002d\u0020\u0042\u0075\u0073\u0069ne\u0073\u0073";};func (_ed *LicenseKey )getExpiryDateToCompare ()_bg .Time {if _ed .Trial {return _bg .Now ().UTC ();
};return _bge .ReleasedAt ;};var _ffa map[string ]struct{};func SetMeteredKey (apiKey string )error {if len (apiKey )==0{_bge .Log .Error ("\u004d\u0065\u0074\u0065\u0072e\u0064\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0041\u0050\u0049 \u004b\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079");
_bge .Log .Error ("\u002d\u0020\u0047\u0072\u0061\u0062\u0020\u006f\u006e\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0072\u0065\u0065\u0020\u0054\u0069\u0065\u0072\u0020\u0061t\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002f\u0063\u006c\u006fud\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f");
return _dag .Errorf ("\u006de\u0074\u0065\u0072e\u0064\u0020\u006ci\u0063en\u0073\u0065\u0020\u0061\u0070\u0069\u0020k\u0065\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079\u003a\u0020\u0063\u0072\u0065\u0061\u0074\u0065 o\u006ee\u0020\u0061\u0074\u0020\u0068\u0074t\u0070\u0073\u003a\u002f\u002fc\u006c\u006f\u0075\u0064\u002e\u0075\u006e\u0069\u0064\u006f\u0063.\u0069\u006f");
};if _gcg !=nil &&(_gcg ._dcc ||_gcg .Tier !=LicenseTierUnlicensed ){_bge .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0043\u0061\u006e\u006eo\u0074 \u0073\u0065\u0074\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0074\u0077\u0069c\u0065\u0020\u002d\u0020\u0053\u0068\u006f\u0075\u006c\u0064\u0020\u006a\u0075\u0073\u0074\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069z\u0065\u0020\u006f\u006e\u0063\u0065");
return _aea .New ("\u006c\u0069\u0063en\u0073\u0065\u0020\u006b\u0065\u0079\u0020\u0061\u006c\u0072\u0065\u0061\u0064\u0079\u0020\u0073\u0065\u0074");};_eca :=_caba ();_eca ._bd =apiKey ;_gea ,_edcf :=_eca .getStatus ();if _edcf !=nil {return _edcf ;};
if !_gea .Valid {return _aea .New ("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064");};_ggc :=&LicenseKey {_dcc :true ,_ac :apiKey ,_ded :true };_gcg =_ggc ;return nil ;};const _fga ="U\u004eI\u0050\u0044\u0046\u005f\u0043\u0055\u0053\u0054O\u004d\u0045\u0052\u005fNA\u004d\u0045";
func SetMeteredKeyPersistentCache (val bool ){_gcg ._ded =val };func _cbee ()string {_gfc :=_f .Getenv ("\u0048\u004f\u004d\u0045");if len (_gfc )==0{_gfc ,_ =_f .UserHomeDir ();};return _gfc ;};type LicenseKey struct{LicenseId string `json:"license_id"`;
CustomerId string `json:"customer_id"`;CustomerName string `json:"customer_name"`;Tier string `json:"tier"`;CreatedAt _bg .Time `json:"-"`;CreatedAtInt int64 `json:"created_at"`;ExpiresAt *_bg .Time `json:"-"`;ExpiresAtInt int64 `json:"expires_at"`;CreatedBy string `json:"created_by"`;
CreatorName string `json:"creator_name"`;CreatorEmail string `json:"creator_email"`;UniPDF bool `json:"unipdf"`;UniOffice bool `json:"unioffice"`;UniHTML bool `json:"unihtml"`;Trial bool `json:"trial"`;_dcc bool ;_ac string ;_ded bool ;};func _ddbd (_adgd string ,_cce string ,_fcbf bool )error {if _gcg ==nil {return _aea .New ("\u006e\u006f\u0020\u006c\u0069\u0063\u0065\u006e\u0073e\u0020\u006b\u0065\u0079");
};if !_gcg ._dcc ||len (_gcg ._ac )==0{return nil ;};if len (_adgd )==0&&!_fcbf {return _aea .New ("\u0064\u006f\u0063\u004b\u0065\u0079\u0020\u006e\u006ft\u0020\u0073\u0065\u0074");};_bdd .Lock ();defer _bdd .Unlock ();if _ffa ==nil {_ffa =map[string ]struct{}{};
};if _ce ==nil {_ce =map[string ]int {};};_ceb :=0;_ ,_gaaa :=_ffa [_adgd ];if !_gaaa {_ffa [_adgd ]=struct{}{};_ceb ++;};if _ceb ==0{return nil ;};_ce [_cce ]++;_aef :=_bg .Now ();_beb ,_fgbc :=_fcf .loadState (_gcg ._ac );if _fgbc !=nil {_bge .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_fgbc );
return _fgbc ;};if _beb .Usage ==nil {_beb .Usage =map[string ]int {};};for _fed ,_bdc :=range _ce {_beb .Usage [_fed ]+=_bdc ;};_ce =nil ;const _gagd =24*_bg .Hour ;const _cbb =3*24*_bg .Hour ;if len (_beb .Instance )==0||_aef .Sub (_beb .LastReported )> _gagd ||(_beb .LimitDocs &&_beb .RemainingDocs <=_beb .Docs +int64 (_ceb ))||_fcbf {_bbg ,_gcc :=_f .Hostname ();
if _gcc !=nil {return _gcc ;};_abd :=_beb .Docs ;_cac ,_cfa ,_gcc :=_fce ();if _gcc !=nil {_bge .Log .Debug ("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u006c\u006fc\u0061\u006c\u0020\u0061\u0064\u0064\u0072\u0065\u0073\u0073\u003a\u0020\u0025\u0073",_gcc .Error ());
_cac =append (_cac ,"\u0069n\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020n\u006ft\u0020a\u0076\u0061\u0069\u006c\u0061\u0062\u006ce");_cfa =append (_cfa ,"\u0069n\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0020n\u006ft\u0020a\u0076\u0061\u0069\u006c\u0061\u0062\u006ce");
}else {_ae .Strings (_cfa );_ae .Strings (_cac );_dffe ,_bfeb :=_dgc ();if _bfeb !=nil {return _bfeb ;};_gac :=false ;for _ ,_fdc :=range _cfa {if _fdc ==_dffe .String (){_gac =true ;};};if !_gac {_cfa =append (_cfa ,_dffe .String ());};};_fge :=_caba ();
_fge ._bd =_gcg ._ac ;_abd +=int64 (_ceb );_cdcg :=meteredUsageCheckinForm {Instance :_beb .Instance ,Next :_beb .Next ,UsageNumber :int (_abd ),NumFailed :_beb .NumErrors ,Hostname :_bbg ,LocalIP :_e .Join (_cfa ,"\u002c\u0020"),MacAddress :_e .Join (_cac ,"\u002c\u0020"),Package :"\u0075\u006e\u0069\u0070\u0064\u0066",PackageVersion :_bge .Version ,Usage :_beb .Usage ,IsPersistentCache :_gcg ._ded ,Timestamp :_aef .Unix ()};
if len (_cac )==0{_cdcg .MacAddress ="\u006e\u006f\u006e\u0065";};_febd :=int64 (0);_eff :=_beb .NumErrors ;_ggg :=_aef ;_gbe :=0;_gfb :=_beb .LimitDocs ;_cca ,_gcc :=_fge .checkinUsage (_cdcg );if _gcc !=nil {if _aef .Sub (_beb .LastReported )> _cbb {if !_cca .Success {return _aea .New (_cca .Message );
};return _aea .New ("\u0074\u006f\u006f\u0020\u006c\u006f\u006e\u0067\u0020\u0073\u0069\u006e\u0063\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0073\u0075\u0063\u0063e\u0073\u0073\u0066\u0075\u006c \u0063\u0068e\u0063\u006b\u0069\u006e");};_febd =_abd ;
_eff ++;_ggg =_beb .LastReported ;}else {_gfb =_cca .LimitDocs ;_gbe =_cca .RemainingDocs ;_eff =0;};if len (_cca .Instance )==0{_cca .Instance =_cdcg .Instance ;};if len (_cca .Next )==0{_cca .Next =_cdcg .Next ;};_gcc =_fcf .updateState (_fge ._bd ,_cca .Instance ,_cca .Next ,int (_febd ),_gfb ,_gbe ,int (_eff ),_ggg ,nil );
if _gcc !=nil {return _gcc ;};if !_cca .Success {return _dag .Errorf ("\u0065r\u0072\u006f\u0072\u003a\u0020\u0025s",_cca .Message );};}else {_fgbc =_fcf .updateState (_gcg ._ac ,_beb .Instance ,_beb .Next ,int (_beb .Docs )+_ceb ,_beb .LimitDocs ,int (_beb .RemainingDocs ),int (_beb .NumErrors ),_beb .LastReported ,_beb .Usage );
if _fgbc !=nil {return _fgbc ;};};return nil ;};func _aed (_dfg []byte )(_a .Reader ,error ){_eac :=new (_da .Buffer );_beg :=_g .NewWriter (_eac );_beg .Write (_dfg );_cbef :=_beg .Close ();if _cbef !=nil {return nil ,_cbef ;};return _eac ,nil ;};type meteredStatusForm struct{};
func _gfa (_fa string ,_bgg []byte )(string ,error ){_aeae ,_ :=_ffg .Decode ([]byte (_fa ));if _aeae ==nil {return "",_dag .Errorf ("\u0050\u0072\u0069\u0076\u004b\u0065\u0079\u0020\u0066a\u0069\u006c\u0065\u0064");};_fd ,_fdg :=_abc .ParsePKCS1PrivateKey (_aeae .Bytes );
if _fdg !=nil {return "",_fdg ;};_cga :=_fc .New ();_cga .Write (_bgg );_ddc :=_cga .Sum (nil );_abg ,_fdg :=_aag .SignPKCS1v15 (_dd .Reader ,_fd ,_d .SHA512 ,_ddc );if _fdg !=nil {return "",_fdg ;};_cc :=_cd .StdEncoding .EncodeToString (_bgg );_cc +="\u000a\u002b\u000a";
_cc +=_cd .StdEncoding .EncodeToString (_abg );return _cc ,nil ;};func SetLicenseKey (content string ,customerName string )error {_ba ,_afa :=_ec (content );if _afa !=nil {_bge .Log .Error ("\u004c\u0069c\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u0072\u0072\u006f\u0072: \u0025\u0076",_afa );
return _afa ;};if !_e .EqualFold (_ba .CustomerName ,customerName ){_bge .Log .Error ("L\u0069ce\u006es\u0065 \u0063\u006f\u0064\u0065\u0020i\u0073\u0073\u0075e\u0020\u002d\u0020\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006da\u0074\u0063\u0068, e\u0078\u0070\u0065\u0063\u0074\u0065d\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067o\u0074 \u0027\u0025\u0073\u0027",customerName ,_ba .CustomerName );
return _dag .Errorf ("\u0063\u0075\u0073\u0074\u006fm\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006d\u0061t\u0063\u0068\u002c\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067\u006f\u0074\u0020\u0027\u0025\u0073'",customerName ,_ba .CustomerName );
};_afa =_ba .Validate ();if _afa !=nil {_bge .Log .Error ("\u004c\u0069\u0063\u0065\u006e\u0073e\u0020\u0063\u006f\u0064\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074i\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076",_afa );return _afa ;
};_gcg =&_ba ;return nil ;};func GetLicenseKey ()*LicenseKey {if _gcg ==nil {return nil ;};_efb :=*_gcg ;return &_efb ;};var _fcf stateLoader =defaultStateHolder {};func _dgc ()(_aa .IP ,error ){_agg ,_dffb :=_aa .Dial ("\u0075\u0064\u0070","\u0038\u002e\u0038\u002e\u0038\u002e\u0038\u003a\u0038\u0030");
if _dffb !=nil {return nil ,_dffb ;};defer _agg .Close ();_dcd :=_agg .LocalAddr ().(*_aa .UDPAddr );return _dcd .IP ,nil ;};type MeteredStatus struct{OK bool ;Credits int64 ;Used int64 ;};var _dagb =_bg .Date (2020,1,1,0,0,0,0,_bg .UTC );func _dgba (_aacf *_b .Response )(_a .ReadCloser ,error ){var _dae error ;
var _ebd _a .ReadCloser ;switch _e .ToLower (_aacf .Header .Get ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")){case "\u0067\u007a\u0069\u0070":_ebd ,_dae =_g .NewReader (_aacf .Body );if _dae !=nil {return _ebd ,_dae ;
};defer _ebd .Close ();default:_ebd =_aacf .Body ;};return _ebd ,nil ;};func _edd (_faa *_b .Response )([]byte ,error ){var _dab []byte ;_fdb ,_fgef :=_dgba (_faa );if _fgef !=nil {return _dab ,_fgef ;};return _ga .ReadAll (_fdb );};var _gcg =MakeUnlicensedKey ();
func init (){_eea :=_f .Getenv (_bac );_ggba :=_f .Getenv (_fga );if len (_eea )==0||len (_ggba )==0{return ;};_aga ,_abfd :=_ga .ReadFile (_eea );if _abfd !=nil {_bge .Log .Error ("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0072\u0065ad \u006cic\u0065\u006e\u0073\u0065\u0020\u0063\u006fde\u0020\u0066\u0069\u006c\u0065\u003a\u0020%\u0076",_abfd );
return ;};_abfd =SetLicenseKey (string (_aga ),_ggba );if _abfd !=nil {_bge .Log .Error ("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u006c\u006f\u0061\u0064\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0076",_abfd );
return ;};};type reportState struct{Instance string `json:"inst"`;Next string `json:"n"`;Docs int64 `json:"d"`;NumErrors int64 `json:"e"`;LimitDocs bool `json:"ld"`;RemainingDocs int64 `json:"rd"`;LastReported _bg .Time `json:"lr"`;LastWritten _bg .Time `json:"lw"`;
Usage map[string ]int `json:"u"`;};const _bac ="\u0055\u004e\u0049\u0050DF\u005f\u004c\u0049\u0043\u0045\u004e\u0053\u0045\u005f\u0050\u0041\u0054\u0048";func (_adfc defaultStateHolder )loadState (_ddfg string )(reportState ,error ){_dde :=_cbee ();if len (_dde )==0{return reportState {},_aea .New ("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");
};_fde :=_ge .Join (_dde ,"\u002eu\u006e\u0069\u0064\u006f\u0063");_efg :=_f .MkdirAll (_fde ,0777);if _efg !=nil {return reportState {},_efg ;};if len (_ddfg )< 20{return reportState {},_aea .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079");
};_fgb :=[]byte (_ddfg );_dac :=_fc .Sum512_256 (_fgb [:20]);_ddb :=_ff .EncodeToString (_dac [:]);_cbd :=_ge .Join (_fde ,_ddb );_bfa ,_efg :=_ga .ReadFile (_cbd );if _efg !=nil {if _f .IsNotExist (_efg ){return reportState {},nil ;};_bge .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_efg );
return reportState {},_aea .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061");};const _aeeb ="\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057";_bfa ,_efg =_cbbf ([]byte (_aeeb ),_bfa );if _efg !=nil {return reportState {},_efg ;
};var _bbdf reportState ;_efg =_gf .Unmarshal (_bfa ,&_bbdf );if _efg !=nil {_bge .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0076",_efg );return reportState {},_aea .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061");
};return _bbdf ,nil ;};const (LicenseTierUnlicensed ="\u0075\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";LicenseTierCommunity ="\u0063o\u006d\u006d\u0075\u006e\u0069\u0074y";LicenseTierIndividual ="\u0069\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c";
LicenseTierBusiness ="\u0062\u0075\u0073\u0069\u006e\u0065\u0073\u0073";);var _bdd =&_cf .Mutex {};func (_fb *meteredClient )getStatus ()(meteredStatusResp ,error ){var _cdb meteredStatusResp ;_gaa :=_fb ._fcc +"\u002fm\u0065t\u0065\u0072\u0065\u0064\u002f\u0073\u0074\u0061\u0074\u0075\u0073";
var _bea meteredStatusForm ;_dba ,_eec :=_gf .Marshal (_bea );if _eec !=nil {return _cdb ,_eec ;};_eb ,_eec :=_aed (_dba );if _eec !=nil {return _cdb ,_eec ;};_dcf ,_eec :=_b .NewRequest ("\u0050\u004f\u0053\u0054",_gaa ,_eb );if _eec !=nil {return _cdb ,_eec ;
};_dcf .Header .Add ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e");_dcf .Header .Add ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");
_dcf .Header .Add ("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");_dcf .Header .Add ("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y",_fb ._bd );_eg ,_eec :=_fb ._fg .Do (_dcf );if _eec !=nil {return _cdb ,_eec ;
};defer _eg .Body .Close ();if _eg .StatusCode !=200{return _cdb ,_dag .Errorf ("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064",_eg .StatusCode );
};_aeg ,_eec :=_edd (_eg );if _eec !=nil {return _cdb ,_eec ;};_eec =_gf .Unmarshal (_aeg ,&_cdb );if _eec !=nil {return _cdb ,_eec ;};return _cdb ,nil ;};func _dfc (_ca string ,_aee string )([]byte ,error ){var (_cfb int ;_de string ;);for _ ,_de =range []string {"\u000a\u002b\u000a","\u000d\u000a\u002b\r\u000a","\u0020\u002b\u0020"}{if _cfb =_e .Index (_aee ,_de );
_cfb !=-1{break ;};};if _cfb ==-1{return nil ,_dag .Errorf ("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u002c \u0073i\u0067n\u0061t\u0075\u0072\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u006f\u0072");};_be :=_aee [:_cfb ];
_bb :=_cfb +len (_de );_gg :=_aee [_bb :];if _be ==""||_gg ==""{return nil ,_dag .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0069n\u0070\u0075\u0074,\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0069\u0067\u0069n\u0061\u006c \u006f\u0072\u0020\u0073\u0069\u0067n\u0061\u0074u\u0072\u0065");
};_gae ,_dg :=_cd .StdEncoding .DecodeString (_be );if _dg !=nil {return nil ,_dag .Errorf ("\u0069\u006e\u0076\u0061li\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0072\u0069\u0067\u0069\u006ea\u006c");};_ccc ,_dg :=_cd .StdEncoding .DecodeString (_gg );
if _dg !=nil {return nil ,_dag .Errorf ("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_cab ,_ :=_ffg .Decode ([]byte (_ca ));if _cab ==nil {return nil ,_dag .Errorf ("\u0050\u0075\u0062\u004b\u0065\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064");
};_dc ,_dg :=_abc .ParsePKIXPublicKey (_cab .Bytes );if _dg !=nil {return nil ,_dg ;};_cccb :=_dc .(*_aag .PublicKey );if _cccb ==nil {return nil ,_dag .Errorf ("\u0050u\u0062\u004b\u0065\u0079\u0020\u0063\u006f\u006e\u0076\u0065\u0072s\u0069\u006f\u006e\u0020\u0066\u0061\u0069\u006c\u0065\u0064");
};_gab :=_fc .New ();_gab .Write (_gae );_bc :=_gab .Sum (nil );_dg =_aag .VerifyPKCS1v15 (_cccb ,_d .SHA512 ,_bc ,_ccc );if _dg !=nil {return nil ,_dg ;};return _gae ,nil ;};var _aebe =_bg .Date (2019,6,6,0,0,0,0,_bg .UTC );