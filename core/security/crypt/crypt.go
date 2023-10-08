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

package crypt ;import (_f "crypto/aes";_b "crypto/cipher";_c "crypto/md5";_aa "crypto/rand";_de "crypto/rc4";_e "fmt";_cf "github.com/unidoc/unipdf/v3/common";_fb "github.com/unidoc/unipdf/v3/core/security";_a "io";);func init (){_bca ("\u0041\u0045\u0053V\u0032",_be )};


// KeyLength implements Filter interface.
func (_fgg filterV2 )KeyLength ()int {return _fgg ._eaf };

// HandlerVersion implements Filter interface.
func (_geb filterV2 )HandlerVersion ()(V ,R int ){V ,R =2,3;return ;};

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3 ()Filter {_db ,_fa :=_df (FilterDict {});if _fa !=nil {_cf .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_fa );
return filterAESV3 {};};return _db ;};

// Name implements Filter interface.
func (filterV2 )Name ()string {return "\u0056\u0032"};

// MakeKey implements Filter interface.
func (filterAESV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _cd (objNum ,genNum ,ekey ,true );};func _bca (_adg string ,_gce filterFunc ){if _ ,_aga :=_fce [_adg ];_aga {panic ("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064");
};_fce [_adg ]=_gce ;};

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity ()Filter {return filterIdentity {}};

// Name implements Filter interface.
func (filterAESV3 )Name ()string {return "\u0041\u0045\u0053V\u0033"};func _be (_g FilterDict )(Filter ,error ){if _g .Length ==128{_cf .Log .Debug ("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_g .Length );
_g .Length /=8;};if _g .Length !=0&&_g .Length !=16{return nil ,_e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_g .Length );
};return filterAESV2 {},nil ;};func init (){_bca ("\u0041\u0045\u0053V\u0033",_df )};

// EncryptBytes implements Filter interface.
func (filterV2 )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_dfd ,_bde :=_de .NewCipher (okey );if _bde !=nil {return nil ,_bde ;};_cf .Log .Trace ("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );
_dfd .XORKeyStream (buf ,buf );_cf .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};func (filterIdentity )MakeKey (objNum ,genNum uint32 ,fkey []byte )([]byte ,error ){return fkey ,nil };type filterAESV2 struct{filterAES };


// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter (d FilterDict )(Filter ,error ){_aac ,_bcc :=_fbd (d .CFM );if _bcc !=nil {return nil ,_bcc ;};_ae ,_bcc :=_aac (d );if _bcc !=nil {return nil ,_bcc ;};return _ae ,nil ;};var (_fce =make (map[string ]filterFunc ););

// DecryptBytes implements Filter interface.
func (filterV2 )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ba ,_dc :=_de .NewCipher (okey );if _dc !=nil {return nil ,_dc ;};_cf .Log .Trace ("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_ba .XORKeyStream (buf ,buf );
_cf .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};

// HandlerVersion implements Filter interface.
func (filterAESV3 )HandlerVersion ()(V ,R int ){V ,R =5,6;return ;};var _ Filter =filterAESV3 {};

// KeyLength implements Filter interface.
func (filterAESV3 )KeyLength ()int {return 256/8};func (filterIdentity )EncryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };type filterIdentity struct{};

// PDFVersion implements Filter interface.
func (filterAESV2 )PDFVersion ()[2]int {return [2]int {1,5}};func init (){_bca ("\u0056\u0032",_bd )};func (filterIdentity )KeyLength ()int {return 0};

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2 ()Filter {_ce ,_dea :=_be (FilterDict {});if _dea !=nil {_cf .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_dea );
return filterAESV2 {};};return _ce ;};

// MakeKey implements Filter interface.
func (filterAESV3 )MakeKey (_ ,_ uint32 ,ekey []byte )([]byte ,error ){return ekey ,nil };func _cd (_fg ,_gfb uint32 ,_eg []byte ,_gde bool )([]byte ,error ){_fgc :=make ([]byte ,len (_eg )+5);copy (_fgc ,_eg );for _fbe :=0;_fbe < 3;_fbe ++{_cae :=byte ((_fg >>uint32 (8*_fbe ))&0xff);
_fgc [_fbe +len (_eg )]=_cae ;};for _gc :=0;_gc < 2;_gc ++{_caec :=byte ((_gfb >>uint32 (8*_gc ))&0xff);_fgc [_gc +len (_eg )+3]=_caec ;};if _gde {_fgc =append (_fgc ,0x73);_fgc =append (_fgc ,0x41);_fgc =append (_fgc ,0x6C);_fgc =append (_fgc ,0x54);};
_gdb :=_c .New ();_gdb .Write (_fgc );_ced :=_gdb .Sum (nil );if len (_eg )+5< 16{return _ced [0:len (_eg )+5],nil ;};return _ced ,nil ;};func (filterIdentity )HandlerVersion ()(V ,R int ){return ;};type filterFunc func (_adf FilterDict )(Filter ,error );
func _bd (_ea FilterDict )(Filter ,error ){if _ea .Length %8!=0{return nil ,_e .Errorf ("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029",_ea .Length );
};if _ea .Length < 5||_ea .Length > 16{if _ea .Length ==40||_ea .Length ==64||_ea .Length ==128{_cf .Log .Debug ("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029",_ea .Length );
_ea .Length /=8;}else {return nil ,_e .Errorf ("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029",_ea .Length );
};};return filterV2 {_eaf :_ea .Length },nil ;};var _ Filter =filterV2 {};

// PDFVersion implements Filter interface.
func (filterAESV3 )PDFVersion ()[2]int {return [2]int {2,0}};type filterAES struct{};type filterV2 struct{_eaf int };var _ Filter =filterAESV2 {};func (filterIdentity )Name ()string {return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"};

// Filter is a common interface for crypt filter methods.
type Filter interface{

// Name returns a name of the filter that should be used in CFM field of Encrypt dictionary.
Name ()string ;

// KeyLength returns a length of the encryption key in bytes.
KeyLength ()int ;

// PDFVersion reports the minimal version of PDF document that introduced this filter.
PDFVersion ()[2]int ;

// HandlerVersion reports V and R parameters that should be used for this filter.
HandlerVersion ()(V ,R int );

// MakeKey generates a object encryption key based on file encryption key and object numbers.
// Used only for legacy filters - AESV3 doesn't change the key for each object.
MakeKey (_bee ,_fcg uint32 ,_eb []byte )([]byte ,error );

// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and encrypt data in-place.
EncryptBytes (_efb []byte ,_ff []byte )([]byte ,error );

// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and decrypt data in-place.
DecryptBytes (_cfa []byte ,_afc []byte )([]byte ,error );};func (filterIdentity )PDFVersion ()[2]int {return [2]int {}};

// Name implements Filter interface.
func (filterAESV2 )Name ()string {return "\u0041\u0045\u0053V\u0032"};

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2 (length int )Filter {_abd ,_ge :=_bd (FilterDict {Length :length });if _ge !=nil {_cf .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_ge );
return filterV2 {_eaf :length };};return _abd ;};func (filterAES )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_cb ,_ggd :=_f .NewCipher (okey );if _ggd !=nil {return nil ,_ggd ;};if len (buf )< 16{_cf .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073",buf );
return buf ,_e .Errorf ("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029",len (buf ));};_fc :=buf [:16];buf =buf [16:];if len (buf )%16!=0{_cf .Log .Debug ("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (_fc ),_fc );
_cf .Log .Debug ("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,_e .Errorf ("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029",len (buf ));
};_fbc :=_b .NewCBCDecrypter (_cb ,_fc );_cf .Log .Trace ("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );_cf .Log .Trace ("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078",len (buf ),buf );
_fbc .CryptBlocks (buf ,buf );_cf .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );if len (buf )==0{_cf .Log .Trace ("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067");
return buf ,nil ;};_agd :=int (buf [len (buf )-1]);if _agd > len (buf ){_cf .Log .Debug ("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029",_agd ,len (buf ));
return buf ,_e .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068");};buf =buf [:len (buf )-_agd ];return buf ,nil ;};func _fbd (_ebc string )(filterFunc ,error ){_agc :=_fce [_ebc ];if _agc ==nil {return nil ,_e .Errorf ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071",_ebc );
};return _agc ,nil ;};

// MakeKey implements Filter interface.
func (_eeg filterV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _cd (objNum ,genNum ,ekey ,false );};func (filterIdentity )DecryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };type filterAESV3 struct{filterAES };


// HandlerVersion implements Filter interface.
func (filterAESV2 )HandlerVersion ()(V ,R int ){V ,R =4,4;return ;};

// KeyLength implements Filter interface.
func (filterAESV2 )KeyLength ()int {return 128/8};

// PDFVersion implements Filter interface.
func (_feb filterV2 )PDFVersion ()[2]int {return [2]int {}};func (filterAES )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_gf ,_ee :=_f .NewCipher (okey );if _ee !=nil {return nil ,_ee ;};_cf .Log .Trace ("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );
const _af =_f .BlockSize ;_ca :=_af -len (buf )%_af ;for _gd :=0;_gd < _ca ;_gd ++{buf =append (buf ,byte (_ca ));};_cf .Log .Trace ("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073",len (buf ));_ef :=make ([]byte ,_af +len (buf ));
_ag :=_ef [:_af ];if _ ,_ec :=_a .ReadFull (_aa .Reader ,_ag );_ec !=nil {return nil ,_ec ;};_bb :=_b .NewCBCEncrypter (_gf ,_ag );_bb .CryptBlocks (_ef [_af :],buf );buf =_ef ;_cf .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );
return buf ,nil ;};func _df (_ab FilterDict )(Filter ,error ){if _ab .Length ==256{_cf .Log .Debug ("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_ab .Length );
_ab .Length /=8;};if _ab .Length !=0&&_ab .Length !=32{return nil ,_e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_ab .Length );
};return filterAESV3 {},nil ;};

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct{CFM string ;AuthEvent _fb .AuthEvent ;Length int ;};