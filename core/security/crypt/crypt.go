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

package crypt ;import (_c "crypto/aes";_fc "crypto/cipher";_cg "crypto/md5";_b "crypto/rand";_fe "crypto/rc4";_f "fmt";_a "github.com/unidoc/unipdf/v3/common";_bb "github.com/unidoc/unipdf/v3/core/security";_d "io";);func init (){_gfa ("\u0041\u0045\u0053V\u0032",_ae )};
var _ Filter =filterAESV3 {};func _ae (_cdd FilterDict )(Filter ,error ){if _cdd .Length ==128{_a .Log .Debug ("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_cdd .Length );
_cdd .Length /=8;};if _cdd .Length !=0&&_cdd .Length !=16{return nil ,_f .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_cdd .Length );
};return filterAESV2 {},nil ;};

// PDFVersion implements Filter interface.
func (filterAESV2 )PDFVersion ()[2]int {return [2]int {1,5}};func (filterIdentity )PDFVersion ()[2]int {return [2]int {}};func (filterAES )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ac ,_ea :=_c .NewCipher (okey );if _ea !=nil {return nil ,_ea ;
};_a .Log .Trace ("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );const _ag =_c .BlockSize ;_bdc :=_ag -len (buf )%_ag ;for _bf :=0;_bf < _bdc ;_bf ++{buf =append (buf ,byte (_bdc ));
};_a .Log .Trace ("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073",len (buf ));_bc :=make ([]byte ,_ag +len (buf ));_cb :=_bc [:_ag ];if _ ,_de :=_d .ReadFull (_b .Reader ,_cb );_de !=nil {return nil ,_de ;};_gd :=_fc .NewCBCEncrypter (_ac ,_cb );
_gd .CryptBlocks (_bc [_ag :],buf );buf =_bc ;_a .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,nil ;};type filterV2 struct{_gf int };

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity ()Filter {return filterIdentity {}};func (filterIdentity )MakeKey (objNum ,genNum uint32 ,fkey []byte )([]byte ,error ){return fkey ,nil };type filterAESV2 struct{filterAES };func init (){_gfa ("\u0041\u0045\u0053V\u0033",_bd )};

// MakeKey implements Filter interface.
func (filterAESV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _fg (objNum ,genNum ,ekey ,true );};var _ Filter =filterV2 {};func (filterIdentity )DecryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };

// KeyLength implements Filter interface.
func (filterAESV3 )KeyLength ()int {return 256/8};

// HandlerVersion implements Filter interface.
func (_acg filterV2 )HandlerVersion ()(V ,R int ){V ,R =2,3;return ;};func _gfa (_egc string ,_ada filterFunc ){if _ ,_bbb :=_ad [_egc ];_bbb {panic ("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064");};_ad [_egc ]=_ada ;
};func (filterIdentity )HandlerVersion ()(V ,R int ){return ;};type filterAES struct{};type filterIdentity struct{};

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
MakeKey (_ec ,_ge uint32 ,_ebc []byte )([]byte ,error );

// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and encrypt data in-place.
EncryptBytes (_fac []byte ,_aa []byte )([]byte ,error );

// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and decrypt data in-place.
DecryptBytes (_gb []byte ,_aaa []byte )([]byte ,error );};

// Name implements Filter interface.
func (filterAESV3 )Name ()string {return "\u0041\u0045\u0053V\u0033"};

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct{CFM string ;AuthEvent _bb .AuthEvent ;Length int ;};func init (){_gfa ("\u0056\u0032",_fcd )};

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3 ()Filter {_cdc ,_cc :=_bd (FilterDict {});if _cc !=nil {_a .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_cc );
return filterAESV3 {};};return _cdc ;};

// Name implements Filter interface.
func (filterV2 )Name ()string {return "\u0056\u0032"};type filterAESV3 struct{filterAES };

// KeyLength implements Filter interface.
func (filterAESV2 )KeyLength ()int {return 128/8};func (filterAES )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_cda ,_cdg :=_c .NewCipher (okey );if _cdg !=nil {return nil ,_cdg ;};if len (buf )< 16{_a .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073",buf );
return buf ,_f .Errorf ("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029",len (buf ));};_cea :=buf [:16];buf =buf [16:];if len (buf )%16!=0{_a .Log .Debug ("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (_cea ),_cea );
_a .Log .Debug ("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,_f .Errorf ("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029",len (buf ));
};_eb :=_fc .NewCBCDecrypter (_cda ,_cea );_a .Log .Trace ("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );_a .Log .Trace ("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078",len (buf ),buf );
_eb .CryptBlocks (buf ,buf );_a .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );if len (buf )==0{_a .Log .Trace ("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067");
return buf ,nil ;};_cdf :=int (buf [len (buf )-1]);if _cdf > len (buf ){_a .Log .Debug ("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029",_cdf ,len (buf ));
return buf ,_f .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068");};buf =buf [:len (buf )-_cdf ];return buf ,nil ;};

// MakeKey implements Filter interface.
func (_ab filterV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _fg (objNum ,genNum ,ekey ,false );};

// PDFVersion implements Filter interface.
func (_dg filterV2 )PDFVersion ()[2]int {return [2]int {}};func _fg (_ff ,_cge uint32 ,_ga []byte ,_fa bool )([]byte ,error ){_ffb :=make ([]byte ,len (_ga )+5);copy (_ffb ,_ga );for _bff :=0;_bff < 3;_bff ++{_bg :=byte ((_ff >>uint32 (8*_bff ))&0xff);
_ffb [_bff +len (_ga )]=_bg ;};for _fbb :=0;_fbb < 2;_fbb ++{_bgd :=byte ((_cge >>uint32 (8*_fbb ))&0xff);_ffb [_fbb +len (_ga )+3]=_bgd ;};if _fa {_ffb =append (_ffb ,0x73);_ffb =append (_ffb ,0x41);_ffb =append (_ffb ,0x6C);_ffb =append (_ffb ,0x54);
};_fee :=_cg .New ();_fee .Write (_ffb );_agb :=_fee .Sum (nil );if len (_ga )+5< 16{return _agb [0:len (_ga )+5],nil ;};return _agb ,nil ;};func (filterIdentity )KeyLength ()int {return 0};var _ Filter =filterAESV2 {};

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter (d FilterDict )(Filter ,error ){_abc ,_caa :=_cca (d .CFM );if _caa !=nil {return nil ,_caa ;};_edc ,_caa :=_abc (d );if _caa !=nil {return nil ,_caa ;};return _edc ,nil ;};

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2 (length int )Filter {_gc ,_fef :=_fcd (FilterDict {Length :length });if _fef !=nil {_a .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_fef );
return filterV2 {_gf :length };};return _gc ;};func _fcd (_af FilterDict )(Filter ,error ){if _af .Length %8!=0{return nil ,_f .Errorf ("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029",_af .Length );
};if _af .Length < 5||_af .Length > 16{if _af .Length ==40||_af .Length ==64||_af .Length ==128{_a .Log .Debug ("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029",_af .Length );
_af .Length /=8;}else {return nil ,_f .Errorf ("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029",_af .Length );
};};return filterV2 {_gf :_af .Length },nil ;};

// KeyLength implements Filter interface.
func (_ca filterV2 )KeyLength ()int {return _ca ._gf };func (filterIdentity )EncryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };func (filterIdentity )Name ()string {return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"};

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2 ()Filter {_g ,_fd :=_ae (FilterDict {});if _fd !=nil {_a .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_fd );
return filterAESV2 {};};return _g ;};func _cca (_gfd string )(filterFunc ,error ){_da :=_ad [_gfd ];if _da ==nil {return nil ,_f .Errorf ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071",_gfd );
};return _da ,nil ;};

// PDFVersion implements Filter interface.
func (filterAESV3 )PDFVersion ()[2]int {return [2]int {2,0}};

// EncryptBytes implements Filter interface.
func (filterV2 )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_bba ,_gfb :=_fe .NewCipher (okey );if _gfb !=nil {return nil ,_gfb ;};_a .Log .Trace ("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );
_bba .XORKeyStream (buf ,buf );_a .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};

// MakeKey implements Filter interface.
func (filterAESV3 )MakeKey (_ ,_ uint32 ,ekey []byte )([]byte ,error ){return ekey ,nil };

// DecryptBytes implements Filter interface.
func (filterV2 )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_dd ,_cdgg :=_fe .NewCipher (okey );if _cdgg !=nil {return nil ,_cdgg ;};_a .Log .Trace ("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );
_dd .XORKeyStream (buf ,buf );_a .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};type filterFunc func (_ed FilterDict )(Filter ,error );

// HandlerVersion implements Filter interface.
func (filterAESV2 )HandlerVersion ()(V ,R int ){V ,R =4,4;return ;};var (_ad =make (map[string ]filterFunc ););func _bd (_fcc FilterDict )(Filter ,error ){if _fcc .Length ==256{_a .Log .Debug ("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_fcc .Length );
_fcc .Length /=8;};if _fcc .Length !=0&&_fcc .Length !=32{return nil ,_f .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_fcc .Length );
};return filterAESV3 {},nil ;};

// HandlerVersion implements Filter interface.
func (filterAESV3 )HandlerVersion ()(V ,R int ){V ,R =5,6;return ;};

// Name implements Filter interface.
func (filterAESV2 )Name ()string {return "\u0041\u0045\u0053V\u0032"};