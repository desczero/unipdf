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

package security ;import (_ec "bytes";_ag "crypto/aes";_a "crypto/cipher";_e "crypto/md5";_ce "crypto/rand";_cc "crypto/rc4";_gg "crypto/sha256";_ga "crypto/sha512";_f "encoding/binary";_c "errors";_b "fmt";_cce "github.com/unidoc/unipdf/v3/common";_ge "hash";
_gf "io";_cd "math";);func _ff (_gaf _a .Block )*ecb {return &ecb {_bd :_gaf ,_ged :_gaf .BlockSize ()}};type errInvalidField struct{Func string ;Field string ;Exp int ;Got int ;};func (_ggb *ecbEncrypter )CryptBlocks (dst ,src []byte ){if len (src )%_ggb ._ged !=0{_cce .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_cce .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_ggb ._bd .Encrypt (dst ,src [:_ggb ._ged ]);src =src [_ggb ._ged :];dst =dst [_ggb ._ged :];};};

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4 (id0 string ,length int )StdHandler {return stdHandlerR4 {ID0 :id0 ,Length :length }};func (_de *ecbDecrypter )CryptBlocks (dst ,src []byte ){if len (src )%_de ._ged !=0{_cce .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_cce .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_de ._bd .Decrypt (dst ,src [:_de ._ged ]);src =src [_de ._ged :];dst =dst [_de ._ged :];};};type ecbEncrypter ecb ;func (_db *ecbDecrypter )BlockSize ()int {return _db ._ged };func (_ecd stdHandlerR6 )alg2a (_gge *StdEncryptDict ,_fc []byte )([]byte ,Permissions ,error ){if _cfc :=_bc ("\u0061\u006c\u00672\u0061","\u004f",48,_gge .O );
_cfc !=nil {return nil ,0,_cfc ;};if _ade :=_bc ("\u0061\u006c\u00672\u0061","\u0055",48,_gge .U );_ade !=nil {return nil ,0,_ade ;};if len (_fc )> 127{_fc =_fc [:127];};_dc ,_cad :=_ecd .alg12 (_gge ,_fc );if _cad !=nil {return nil ,0,_cad ;};var (_fgg []byte ;
_dac []byte ;_fbbf []byte ;);var _bge Permissions ;if len (_dc )!=0{_bge =PermOwner ;_ada :=make ([]byte ,len (_fc )+8+48);_eac :=copy (_ada ,_fc );_eac +=copy (_ada [_eac :],_gge .O [40:48]);copy (_ada [_eac :],_gge .U [0:48]);_fgg =_ada ;_dac =_gge .OE ;
_fbbf =_gge .U [0:48];}else {_dc ,_cad =_ecd .alg11 (_gge ,_fc );if _cad ==nil &&len (_dc )==0{_dc ,_cad =_ecd .alg11 (_gge ,[]byte (""));};if _cad !=nil {return nil ,0,_cad ;}else if len (_dc )==0{return nil ,0,nil ;};_bge =_gge .P ;_ee :=make ([]byte ,len (_fc )+8);
_agd :=copy (_ee ,_fc );copy (_ee [_agd :],_gge .U [40:48]);_fgg =_ee ;_dac =_gge .UE ;_fbbf =nil ;};if _eda :=_bc ("\u0061\u006c\u00672\u0061","\u004b\u0065\u0079",32,_dac );_eda !=nil {return nil ,0,_eda ;};_dac =_dac [:32];_gea ,_cad :=_ecd .alg2b (_gge .R ,_fgg ,_fc ,_fbbf );
if _cad !=nil {return nil ,0,_cad ;};_cef ,_cad :=_ag .NewCipher (_gea [:32]);if _cad !=nil {return nil ,0,_cad ;};_cceb :=make ([]byte ,_ag .BlockSize );_dbb :=_a .NewCBCDecrypter (_cef ,_cceb );_fcc :=make ([]byte ,32);_dbb .CryptBlocks (_fcc ,_dac );
if _gge .R ==5{return _fcc ,_bge ,nil ;};_cad =_ecd .alg13 (_gge ,_fcc );if _cad !=nil {return nil ,0,_cad ;};return _fcc ,_bge ,nil ;};

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32 ;func (_dfb stdHandlerR4 )alg4 (_bae []byte ,_ebg []byte )([]byte ,error ){_aea ,_ad :=_cc .NewCipher (_bae );if _ad !=nil {return nil ,_c .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_fab :=[]byte (_be );_fe :=make ([]byte ,len (_fab ));_aea .XORKeyStream (_fe ,_fab );return _fe ,nil ;};func (_abf stdHandlerR6 )alg12 (_ccdf *StdEncryptDict ,_gad []byte )([]byte ,error ){if _cafbb :=_bc ("\u0061\u006c\u00671\u0032","\u0055",48,_ccdf .U );
_cafbb !=nil {return nil ,_cafbb ;};if _egb :=_bc ("\u0061\u006c\u00671\u0032","\u004f",48,_ccdf .O );_egb !=nil {return nil ,_egb ;};_bbd :=make ([]byte ,len (_gad )+8+48);_abg :=copy (_bbd ,_gad );_abg +=copy (_bbd [_abg :],_ccdf .O [32:40]);_abg +=copy (_bbd [_abg :],_ccdf .U [0:48]);
_gaa ,_cg :=_abf .alg2b (_ccdf .R ,_bbd ,_gad ,_ccdf .U [0:48]);if _cg !=nil {return nil ,_cg ;};_gaa =_gaa [:32];if !_ec .Equal (_gaa ,_ccdf .O [:32]){return nil ,nil ;};return _gaa ,nil ;};func _ccd (_cda _a .Block )_a .BlockMode {return (*ecbEncrypter )(_ff (_cda ))};
func (stdHandlerR4 )paddedPass (_bcd []byte )[]byte {_bfb :=make ([]byte ,32);_eb :=copy (_bfb ,_bcd );for ;_eb < 32;_eb ++{_bfb [_eb ]=_be [_eb -len (_bcd )];};return _bfb ;};type stdHandlerR6 struct{};func (_bfef stdHandlerR6 )alg2b (R int ,_adg ,_fcf ,_bbe []byte )([]byte ,error ){if R ==5{return _ebe (_adg );
};return _aeaa (_adg ,_fcf ,_bbe );};func _aeaa (_ecc ,_fgc ,_ccb []byte )([]byte ,error ){var (_cba ,_ffa ,_bfda _ge .Hash ;);_cba =_gg .New ();_fac :=make ([]byte ,64);_ddg :=_cba ;_ddg .Write (_ecc );K :=_ddg .Sum (_fac [:0]);_cdcf :=make ([]byte ,64*(127+64+48));
_gae :=func (_afa int )([]byte ,error ){_ffg :=len (_fgc )+len (K )+len (_ccb );_facf :=_cdcf [:_ffg ];_aa :=copy (_facf ,_fgc );_aa +=copy (_facf [_aa :],K [:]);_aa +=copy (_facf [_aa :],_ccb );if _aa !=_ffg {_cce .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e");
return nil ,_c .New ("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065");};K1 :=_cdcf [:_ffg *64];_ddd (K1 ,_ffg );_fd ,_bac :=_fgfc (K [0:16]);if _bac !=nil {return nil ,_bac ;};_gcd :=_a .NewCBCEncrypter (_fd ,K [16:32]);_gcd .CryptBlocks (K1 ,K1 );
E :=K1 ;_ab :=0;for _aab :=0;_aab < 16;_aab ++{_ab +=int (E [_aab ]%3);};var _eeg _ge .Hash ;switch _ab %3{case 0:_eeg =_cba ;case 1:if _ffa ==nil {_ffa =_ga .New384 ();};_eeg =_ffa ;case 2:if _bfda ==nil {_bfda =_ga .New ();};_eeg =_bfda ;};_eeg .Reset ();
_eeg .Write (E );K =_eeg .Sum (_fac [:0]);return E ,nil ;};for _baad :=0;;{E ,_bcf :=_gae (_baad );if _bcf !=nil {return nil ,_bcf ;};_efe :=E [len (E )-1];_baad ++;if _baad >=64&&_efe <=uint8 (_baad -32){break ;};};return K [:32],nil ;};func _bc (_bgf ,_cdc string ,_dg int ,_efg []byte )error {if len (_efg )< _dg {return errInvalidField {Func :_bgf ,Field :_cdc ,Exp :_dg ,Got :len (_efg )};
};return nil ;};const (PermOwner =Permissions (_cd .MaxUint32 );PermPrinting =Permissions (1<<2);PermModify =Permissions (1<<3);PermExtractGraphics =Permissions (1<<4);PermAnnotate =Permissions (1<<5);PermFillForms =Permissions (1<<8);PermDisabilityExtract =Permissions (1<<9);
PermRotateInsert =Permissions (1<<10);PermFullPrintQuality =Permissions (1<<11););func (_adf stdHandlerR6 )alg8 (_fdb *StdEncryptDict ,_bfdd []byte ,_baea []byte )error {if _gfb :=_bc ("\u0061\u006c\u0067\u0038","\u004b\u0065\u0079",32,_bfdd );_gfb !=nil {return _gfb ;
};var _fgb [16]byte ;if _ ,_aaa :=_gf .ReadFull (_ce .Reader ,_fgb [:]);_aaa !=nil {return _aaa ;};_eg :=_fgb [0:8];_gcedg :=_fgb [8:16];_cca :=make ([]byte ,len (_baea )+len (_eg ));_cfg :=copy (_cca ,_baea );copy (_cca [_cfg :],_eg );_fcfg ,_ace :=_adf .alg2b (_fdb .R ,_cca ,_baea ,nil );
if _ace !=nil {return _ace ;};U :=make ([]byte ,len (_fcfg )+len (_eg )+len (_gcedg ));_cfg =copy (U ,_fcfg [:32]);_cfg +=copy (U [_cfg :],_eg );copy (U [_cfg :],_gcedg );_fdb .U =U ;_cfg =len (_baea );copy (_cca [_cfg :],_gcedg );_fcfg ,_ace =_adf .alg2b (_fdb .R ,_cca ,_baea ,nil );
if _ace !=nil {return _ace ;};_efcc ,_ace :=_fgfc (_fcfg [:32]);if _ace !=nil {return _ace ;};_cdf :=make ([]byte ,_ag .BlockSize );_gcg :=_a .NewCBCEncrypter (_efcc ,_cdf );UE :=make ([]byte ,32);_gcg .CryptBlocks (UE ,_bfdd [:32]);_fdb .UE =UE ;return nil ;
};var _ StdHandler =stdHandlerR6 {};

// Authenticate implements StdHandler interface.
func (_ffffa stdHandlerR6 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){return _ffffa .alg2a (d ,pass );};const (EventDocOpen =AuthEvent ("\u0044o\u0063\u004f\u0070\u0065\u006e");EventEFOpen =AuthEvent ("\u0045\u0046\u004f\u0070\u0065\u006e");
);type stdHandlerR4 struct{Length int ;ID0 string ;};func (_cgg stdHandlerR6 )alg13 (_gaca *StdEncryptDict ,_efgg []byte )error {if _cee :=_bc ("\u0061\u006c\u00671\u0033","\u004b\u0065\u0079",32,_efgg );_cee !=nil {return _cee ;};if _gefd :=_bc ("\u0061\u006c\u00671\u0033","\u0050\u0065\u0072m\u0073",16,_gaca .Perms );
_gefd !=nil {return _gefd ;};_gdc :=make ([]byte ,16);copy (_gdc ,_gaca .Perms [:16]);_bgea ,_abd :=_ag .NewCipher (_efgg [:32]);if _abd !=nil {return _abd ;};_cdaa :=_ea (_bgea );_cdaa .CryptBlocks (_gdc ,_gdc );if !_ec .Equal (_gdc [9:12],[]byte ("\u0061\u0064\u0062")){return _c .New ("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};_ffff :=Permissions (_f .LittleEndian .Uint32 (_gdc [0:4]));if _ffff !=_gaca .P {return _c .New ("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064");
};var _ggc bool ;if _gdc [8]=='T'{_ggc =true ;}else if _gdc [8]=='F'{_ggc =false ;}else {return _c .New ("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};if _ggc !=_gaca .EncryptMetadata {return _c .New ("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064");
};return nil ;};func (_ggf stdHandlerR4 )alg7 (_bce *StdEncryptDict ,_ead []byte )([]byte ,error ){_gce :=_ggf .alg3Key (_bce .R ,_ead );_bfcf :=make ([]byte ,len (_bce .O ));if _bce .R ==2{_fbe ,_dfbdf :=_cc .NewCipher (_gce );if _dfbdf !=nil {return nil ,_c .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");
};_fbe .XORKeyStream (_bfcf ,_bce .O );}else if _bce .R >=3{_acc :=append ([]byte {},_bce .O ...);for _aeca :=0;_aeca < 20;_aeca ++{_fbc :=append ([]byte {},_gce ...);for _dga :=0;_dga < len (_gce );_dga ++{_fbc [_dga ]^=byte (19-_aeca );};_caf ,_dab :=_cc .NewCipher (_fbc );
if _dab !=nil {return nil ,_c .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");};_caf .XORKeyStream (_bfcf ,_acc );_acc =append ([]byte {},_bfcf ...);};}else {return nil ,_c .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");
};_adb ,_gced :=_ggf .alg6 (_bce ,_bfcf );if _gced !=nil {return nil ,nil ;};return _adb ,nil ;};func (_cfb stdHandlerR4 )alg3 (R int ,_ffe ,_bb []byte )([]byte ,error ){var _bfc []byte ;if len (_bb )> 0{_bfc =_cfb .alg3Key (R ,_bb );}else {_bfc =_cfb .alg3Key (R ,_ffe );
};_fg ,_gfg :=_cc .NewCipher (_bfc );if _gfg !=nil {return nil ,_c .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_gc :=_cfb .paddedPass (_ffe );_dd :=make ([]byte ,len (_gc ));_fg .XORKeyStream (_dd ,_gc );
if R >=3{_aec :=make ([]byte ,len (_bfc ));for _bdf :=0;_bdf < 19;_bdf ++{for _bfa :=0;_bfa < len (_bfc );_bfa ++{_aec [_bfa ]=_bfc [_bfa ]^byte (_bdf +1);};_bfd ,_ca :=_cc .NewCipher (_aec );if _ca !=nil {return nil ,_c .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_bfd .XORKeyStream (_dd ,_dd );};};return _dd ,nil ;};func (_gb stdHandlerR6 )alg9 (_cbaf *StdEncryptDict ,_bfae []byte ,_eeb []byte )error {if _abc :=_bc ("\u0061\u006c\u0067\u0039","\u004b\u0065\u0079",32,_bfae );_abc !=nil {return _abc ;};if _geb :=_bc ("\u0061\u006c\u0067\u0039","\u0055",48,_cbaf .U );
_geb !=nil {return _geb ;};var _gac [16]byte ;if _ ,_dad :=_gf .ReadFull (_ce .Reader ,_gac [:]);_dad !=nil {return _dad ;};_bgce :=_gac [0:8];_dddg :=_gac [8:16];_dgeg :=_cbaf .U [:48];_dfea :=make ([]byte ,len (_eeb )+len (_bgce )+len (_dgeg ));_ccad :=copy (_dfea ,_eeb );
_ccad +=copy (_dfea [_ccad :],_bgce );_ccad +=copy (_dfea [_ccad :],_dgeg );_dcc ,_bdc :=_gb .alg2b (_cbaf .R ,_dfea ,_eeb ,_dgeg );if _bdc !=nil {return _bdc ;};O :=make ([]byte ,len (_dcc )+len (_bgce )+len (_dddg ));_ccad =copy (O ,_dcc [:32]);_ccad +=copy (O [_ccad :],_bgce );
_ccad +=copy (O [_ccad :],_dddg );_cbaf .O =O ;_ccad =len (_eeb );_ccad +=copy (_dfea [_ccad :],_dddg );_dcc ,_bdc =_gb .alg2b (_cbaf .R ,_dfea ,_eeb ,_dgeg );if _bdc !=nil {return _bdc ;};_cefg ,_bdc :=_fgfc (_dcc [:32]);if _bdc !=nil {return _bdc ;};
_fad :=make ([]byte ,_ag .BlockSize );_dbbe :=_a .NewCBCEncrypter (_cefg ,_fad );OE :=make ([]byte ,32);_dbbe .CryptBlocks (OE ,_bfae [:32]);_cbaf .OE =OE ;return nil ;};type ecb struct{_bd _a .Block ;_ged int ;};func (_ae stdHandlerR4 )alg2 (_af *StdEncryptDict ,_fb []byte )[]byte {_cce .Log .Trace ("\u0061\u006c\u0067\u0032");
_eab :=_ae .paddedPass (_fb );_da :=_e .New ();_da .Write (_eab );_da .Write (_af .O );var _ccf [4]byte ;_f .LittleEndian .PutUint32 (_ccf [:],uint32 (_af .P ));_da .Write (_ccf [:]);_cce .Log .Trace ("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x",_ccf );
_da .Write ([]byte (_ae .ID0 ));_cce .Log .Trace ("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076",_af .R ,_af .EncryptMetadata );
if (_af .R >=4)&&!_af .EncryptMetadata {_da .Write ([]byte {0xff,0xff,0xff,0xff});};_agb :=_da .Sum (nil );if _af .R >=3{_da =_e .New ();for _fbb :=0;_fbb < 50;_fbb ++{_da .Reset ();_da .Write (_agb [0:_ae .Length /8]);_agb =_da .Sum (nil );};};if _af .R >=3{return _agb [0:_ae .Length /8];
};return _agb [0:5];};

// Allowed checks if a set of permissions can be granted.
func (_eag Permissions )Allowed (p2 Permissions )bool {return _eag &p2 ==p2 };func _ebe (_bed []byte )([]byte ,error ){_baa :=_gg .New ();_baa .Write (_bed );return _baa .Sum (nil ),nil ;};func (_bfgb stdHandlerR6 )alg10 (_adab *StdEncryptDict ,_ccaf []byte )error {if _aefg :=_bc ("\u0061\u006c\u00671\u0030","\u004b\u0065\u0079",32,_ccaf );
_aefg !=nil {return _aefg ;};_ebd :=uint64 (uint32 (_adab .P ))|(_cd .MaxUint32 <<32);Perms :=make ([]byte ,16);_f .LittleEndian .PutUint64 (Perms [:8],_ebd );if _adab .EncryptMetadata {Perms [8]='T';}else {Perms [8]='F';};copy (Perms [9:12],"\u0061\u0064\u0062");
if _ ,_egf :=_gf .ReadFull (_ce .Reader ,Perms [12:16]);_egf !=nil {return _egf ;};_ecb ,_bba :=_fgfc (_ccaf [:32]);if _bba !=nil {return _bba ;};_fbee :=_ccd (_ecb );_fbee .CryptBlocks (Perms ,Perms );_adab .Perms =Perms [:16];return nil ;};func (_agf stdHandlerR6 )alg11 (_bfcd *StdEncryptDict ,_baf []byte )([]byte ,error ){if _caa :=_bc ("\u0061\u006c\u00671\u0031","\u0055",48,_bfcd .U );
_caa !=nil {return nil ,_caa ;};_aeg :=make ([]byte ,len (_baf )+8);_age :=copy (_aeg ,_baf );_age +=copy (_aeg [_age :],_bfcd .U [32:40]);_gbc ,_gff :=_agf .alg2b (_bfcd .R ,_aeg ,_baf ,nil );if _gff !=nil {return nil ,_gff ;};_gbc =_gbc [:32];if !_ec .Equal (_gbc ,_bfcd .U [:32]){return nil ,nil ;
};return _gbc ,nil ;};const _be ="\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377"+"\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014"+"\251\xfe\x64\x53\x69\172";func _ddd (_aeb []byte ,_bfe int ){_beeb :=_bfe ;for _beeb < len (_aeb ){copy (_aeb [_beeb :],_aeb [:_beeb ]);
_beeb *=2;};};func _ea (_d _a .Block )_a .BlockMode {return (*ecbDecrypter )(_ff (_d ))};

// StdHandler is an interface for standard security handlers.
type StdHandler interface{

// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
// It assumes that R, P and EncryptMetadata are already set.
GenerateParams (_eaa *StdEncryptDict ,_ecf ,_bg []byte )([]byte ,error );

// Authenticate uses encryption dictionary parameters and the password to calculate
// the document encryption key. It also returns permissions that should be granted to a user.
// In case of failed authentication, it returns empty key and zero permissions with no error.
Authenticate (_bf *StdEncryptDict ,_fa []byte )([]byte ,Permissions ,error );};func (_gec stdHandlerR4 )alg5 (_bgd []byte ,_gd []byte )([]byte ,error ){_bcb :=_e .New ();_bcb .Write ([]byte (_be ));_bcb .Write ([]byte (_gec .ID0 ));_dfe :=_bcb .Sum (nil );
_cce .Log .Trace ("\u0061\u006c\u0067\u0035");_cce .Log .Trace ("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x",_bgd );_cce .Log .Trace ("\u0049D\u003a\u0020\u0025\u0020\u0078",_gec .ID0 );if len (_dfe )!=16{return nil ,_c .New ("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073");
};_geg ,_fgf :=_cc .NewCipher (_bgd );if _fgf !=nil {return nil ,_c .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_ffd :=make ([]byte ,16);_geg .XORKeyStream (_ffd ,_dfe );_gab :=make ([]byte ,len (_bgd ));
for _bda :=0;_bda < 19;_bda ++{for _dgf :=0;_dgf < len (_bgd );_dgf ++{_gab [_dgf ]=_bgd [_dgf ]^byte (_bda +1);};_geg ,_fgf =_cc .NewCipher (_gab );if _fgf !=nil {return nil ,_c .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_geg .XORKeyStream (_ffd ,_ffd );_cce .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078",_bda ,_gab );_cce .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078",_bda ,_ffd );
};_bcdb :=make ([]byte ,32);for _bcc :=0;_bcc < 16;_bcc ++{_bcdb [_bcc ]=_ffd [_bcc ];};_ ,_fgf =_ce .Read (_bcdb [16:32]);if _fgf !=nil {return nil ,_c .New ("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r");
};return _bcdb ,nil ;};func (_ef *ecbEncrypter )BlockSize ()int {return _ef ._ged };func (_ecg errInvalidField )Error ()string {return _b .Sprintf ("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064",_ecg .Func ,_ecg .Field ,_ecg .Exp ,_ecg .Got );
};var _ StdHandler =stdHandlerR4 {};func _fgfc (_bgc []byte )(_a .Block ,error ){_gceb ,_gdb :=_ag .NewCipher (_bgc );if _gdb !=nil {_cce .Log .Error ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076",_gdb );
return nil ,_gdb ;};return _gceb ,nil ;};func (_dfbd stdHandlerR4 )alg6 (_ccc *StdEncryptDict ,_ebb []byte )([]byte ,error ){var (_fff []byte ;_bee error ;);_aef :=_dfbd .alg2 (_ccc ,_ebb );if _ccc .R ==2{_fff ,_bee =_dfbd .alg4 (_aef ,_ebb );}else if _ccc .R >=3{_fff ,_bee =_dfbd .alg5 (_aef ,_ebb );
}else {return nil ,_c .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};if _bee !=nil {return nil ,_bee ;};_cce .Log .Trace ("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f",string (_fff ),string (_ccc .U ));
_gcb :=_fff ;_cdcg :=_ccc .U ;if _ccc .R >=3{if len (_gcb )> 16{_gcb =_gcb [0:16];};if len (_cdcg )> 16{_cdcg =_cdcg [0:16];};};if !_ec .Equal (_gcb ,_cdcg ){return nil ,nil ;};return _aef ,nil ;};

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6 ()StdHandler {return stdHandlerR6 {}};

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_cb stdHandlerR4 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){O ,_bfg :=_cb .alg3 (d .R ,upass ,opass );if _bfg !=nil {_cce .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_bfg );
return nil ,_bfg ;};d .O =O ;_cce .Log .Trace ("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078",O );_cafb :=_cb .alg2 (d ,upass );U ,_bfg :=_cb .alg5 (_cafb ,upass );if _bfg !=nil {_cce .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_bfg );
return nil ,_bfg ;};d .U =U ;_cce .Log .Trace ("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078",U );return _cafb ,nil ;};

// AuthEvent is an event type that triggers authentication.
type AuthEvent string ;

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_bea stdHandlerR6 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){_ecff :=make ([]byte ,32);if _ ,_cfe :=_gf .ReadFull (_ce .Reader ,_ecff );_cfe !=nil {return nil ,_cfe ;};d .U =nil ;d .O =nil ;d .UE =nil ;d .OE =nil ;d .Perms =nil ;
if len (upass )> 127{upass =upass [:127];};if len (opass )> 127{opass =opass [:127];};if _dgae :=_bea .alg8 (d ,_ecff ,upass );_dgae !=nil {return nil ,_dgae ;};if _acg :=_bea .alg9 (d ,_ecff ,opass );_acg !=nil {return nil ,_acg ;};if d .R ==5{return _ecff ,nil ;
};if _dba :=_bea .alg10 (d ,_ecff );_dba !=nil {return nil ,_dba ;};return _ecff ,nil ;};type ecbDecrypter ecb ;

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct{R int ;P Permissions ;EncryptMetadata bool ;O ,U []byte ;OE ,UE []byte ;Perms []byte ;};func (_cf stdHandlerR4 )alg3Key (R int ,_bgb []byte )[]byte {_fba :=_e .New ();_ac :=_cf .paddedPass (_bgb );_fba .Write (_ac );if R >=3{for _df :=0;
_df < 50;_df ++{_dbg :=_fba .Sum (nil );_fba =_e .New ();_fba .Write (_dbg );};};_bef :=_fba .Sum (nil );if R ==2{_bef =_bef [0:5];}else {_bef =_bef [0:_cf .Length /8];};return _bef ;};

// Authenticate implements StdHandler interface.
func (_efc stdHandlerR4 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){_cce .Log .Trace ("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073");
_gabb ,_gdg :=_efc .alg7 (d ,pass );if _gdg !=nil {return nil ,0,_gdg ;};if _gabb !=nil {_cce .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _gabb ,PermOwner ,nil ;
};_cce .Log .Trace ("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073");_gabb ,_gdg =_efc .alg6 (d ,pass );if _gdg !=nil {return nil ,0,_gdg ;
};if _gabb !=nil {_cce .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _gabb ,d .P ,nil ;};return nil ,0,nil ;};