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

package strutils ;import (_g "bytes";_ad "github.com/unidoc/unipdf/v3/common";_f "unicode/utf16";);func UTF16ToRunes (b []byte )[]rune {if len (b )==1{return []rune {rune (b [0])};};if len (b )%2!=0{b =append (b ,0);_ad .Log .Debug ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u0054\u0046\u0031\u0036\u0054\u006f\u0052\u0075\u006e\u0065\u0073.\u0020\u0050\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0077it\u0068\u0020\u007ae\u0072o\u0073\u002e");
};_ae :=len (b )>>1;_e :=make ([]uint16 ,_ae );for _cc :=0;_cc < _ae ;_cc ++{_e [_cc ]=uint16 (b [_cc <<1])<<8+uint16 (b [_cc <<1+1]);};_ca :=_f .Decode (_e );return _ca ;};func init (){_fa =map[rune ]byte {};for _ga ,_c :=range _ee {_fa [_c ]=_ga ;};};
func StringToPDFDocEncoding (s string )[]byte {var _b _g .Buffer ;for _ ,_bf :=range s {_eb ,_gde :=_fa [_bf ];if !_gde {_ad .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0044\u0046\u0044\u006f\u0063\u0045\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0061\u0070\u0070\u0069\u006e\u0067\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0025\u0063\u002f\u0025\u0058\u0020\u002d\u0020s\u006b\u0069\u0070\u0070\u0069n\u0067",_bf ,_bf );
continue ;};_b .WriteByte (_eb );};return _b .Bytes ();};func StringToUTF16 (s string )string {_gg :=_f .Encode ([]rune (s ));var _cd _g .Buffer ;for _ ,_fad :=range _gg {_cd .WriteByte (byte ((_fad >>8)&0xff));_cd .WriteByte (byte (_fad &0xff));};return _cd .String ();
};func PDFDocEncodingToString (b []byte )string {return string (PDFDocEncodingToRunes (b ))};func UTF16ToString (b []byte )string {return string (UTF16ToRunes (b ))};var _fa map[rune ]byte ;func PDFDocEncodingToRunes (b []byte )[]rune {var _ag []rune ;
for _ ,_ada :=range b {_gb ,_eg :=_ee [_ada ];if !_eg {_ad .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u003a\u0020P\u0044\u0046\u0044o\u0063\u0045\u006ec\u006f\u0064i\u006e\u0067\u0020\u0069\u006e\u0070u\u0074 m\u0061\u0070\u0070\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0025\u0064\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067",_ada );
continue ;};_ag =append (_ag ,_gb );};return _ag ;};var _ee =map[byte ]rune {0x01:'\u0001',0x02:'\u0002',0x03:'\u0003',0x04:'\u0004',0x05:'\u0005',0x06:'\u0006',0x07:'\u0007',0x08:'\u0008',0x09:'\u0009',0x0a:'\u000a',0x0b:'\u000b',0x0c:'\u000c',0x0d:'\u000d',0x0e:'\u000e',0x0f:'\u000f',0x10:'\u0010',0x11:'\u0011',0x12:'\u0012',0x13:'\u0013',0x14:'\u0014',0x15:'\u0015',0x16:'\u0017',0x17:'\u0017',0x18:'\u02d8',0x19:'\u02c7',0x1a:'\u02c6',0x1b:'\u02d9',0x1c:'\u02dd',0x1d:'\u02db',0x1e:'\u02da',0x1f:'\u02dc',0x20:'\u0020',0x21:'\u0021',0x22:'\u0022',0x23:'\u0023',0x24:'\u0024',0x25:'\u0025',0x26:'\u0026',0x27:'\u0027',0x28:'\u0028',0x29:'\u0029',0x2a:'\u002a',0x2b:'\u002b',0x2c:'\u002c',0x2d:'\u002d',0x2e:'\u002e',0x2f:'\u002f',0x30:'\u0030',0x31:'\u0031',0x32:'\u0032',0x33:'\u0033',0x34:'\u0034',0x35:'\u0035',0x36:'\u0036',0x37:'\u0037',0x38:'\u0038',0x39:'\u0039',0x3a:'\u003a',0x3b:'\u003b',0x3c:'\u003c',0x3d:'\u003d',0x3e:'\u003e',0x3f:'\u003f',0x40:'\u0040',0x41:'\u0041',0x42:'\u0042',0x43:'\u0043',0x44:'\u0044',0x45:'\u0045',0x46:'\u0046',0x47:'\u0047',0x48:'\u0048',0x49:'\u0049',0x4a:'\u004a',0x4b:'\u004b',0x4c:'\u004c',0x4d:'\u004d',0x4e:'\u004e',0x4f:'\u004f',0x50:'\u0050',0x51:'\u0051',0x52:'\u0052',0x53:'\u0053',0x54:'\u0054',0x55:'\u0055',0x56:'\u0056',0x57:'\u0057',0x58:'\u0058',0x59:'\u0059',0x5a:'\u005a',0x5b:'\u005b',0x5c:'\u005c',0x5d:'\u005d',0x5e:'\u005e',0x5f:'\u005f',0x60:'\u0060',0x61:'\u0061',0x62:'\u0062',0x63:'\u0063',0x64:'\u0064',0x65:'\u0065',0x66:'\u0066',0x67:'\u0067',0x68:'\u0068',0x69:'\u0069',0x6a:'\u006a',0x6b:'\u006b',0x6c:'\u006c',0x6d:'\u006d',0x6e:'\u006e',0x6f:'\u006f',0x70:'\u0070',0x71:'\u0071',0x72:'\u0072',0x73:'\u0073',0x74:'\u0074',0x75:'\u0075',0x76:'\u0076',0x77:'\u0077',0x78:'\u0078',0x79:'\u0079',0x7a:'\u007a',0x7b:'\u007b',0x7c:'\u007c',0x7d:'\u007d',0x7e:'\u007e',0x80:'\u2022',0x81:'\u2020',0x82:'\u2021',0x83:'\u2026',0x84:'\u2014',0x85:'\u2013',0x86:'\u0192',0x87:'\u2044',0x88:'\u2039',0x89:'\u203a',0x8a:'\u2212',0x8b:'\u2030',0x8c:'\u201e',0x8d:'\u201c',0x8e:'\u201d',0x8f:'\u2018',0x90:'\u2019',0x91:'\u201a',0x92:'\u2122',0x93:'\ufb01',0x94:'\ufb02',0x95:'\u0141',0x96:'\u0152',0x97:'\u0160',0x98:'\u0178',0x99:'\u017d',0x9a:'\u0131',0x9b:'\u0142',0x9c:'\u0153',0x9d:'\u0161',0x9e:'\u017e',0xa0:'\u20ac',0xa1:'\u00a1',0xa2:'\u00a2',0xa3:'\u00a3',0xa4:'\u00a4',0xa5:'\u00a5',0xa6:'\u00a6',0xa7:'\u00a7',0xa8:'\u00a8',0xa9:'\u00a9',0xaa:'\u00aa',0xab:'\u00ab',0xac:'\u00ac',0xae:'\u00ae',0xaf:'\u00af',0xb0:'\u00b0',0xb1:'\u00b1',0xb2:'\u00b2',0xb3:'\u00b3',0xb4:'\u00b4',0xb5:'\u00b5',0xb6:'\u00b6',0xb7:'\u00b7',0xb8:'\u00b8',0xb9:'\u00b9',0xba:'\u00ba',0xbb:'\u00bb',0xbc:'\u00bc',0xbd:'\u00bd',0xbe:'\u00be',0xbf:'\u00bf',0xc0:'\u00c0',0xc1:'\u00c1',0xc2:'\u00c2',0xc3:'\u00c3',0xc4:'\u00c4',0xc5:'\u00c5',0xc6:'\u00c6',0xc7:'\u00c7',0xc8:'\u00c8',0xc9:'\u00c9',0xca:'\u00ca',0xcb:'\u00cb',0xcc:'\u00cc',0xcd:'\u00cd',0xce:'\u00ce',0xcf:'\u00cf',0xd0:'\u00d0',0xd1:'\u00d1',0xd2:'\u00d2',0xd3:'\u00d3',0xd4:'\u00d4',0xd5:'\u00d5',0xd6:'\u00d6',0xd7:'\u00d7',0xd8:'\u00d8',0xd9:'\u00d9',0xda:'\u00da',0xdb:'\u00db',0xdc:'\u00dc',0xdd:'\u00dd',0xde:'\u00de',0xdf:'\u00df',0xe0:'\u00e0',0xe1:'\u00e1',0xe2:'\u00e2',0xe3:'\u00e3',0xe4:'\u00e4',0xe5:'\u00e5',0xe6:'\u00e6',0xe7:'\u00e7',0xe8:'\u00e8',0xe9:'\u00e9',0xea:'\u00ea',0xeb:'\u00eb',0xec:'\u00ec',0xed:'\u00ed',0xee:'\u00ee',0xef:'\u00ef',0xf0:'\u00f0',0xf1:'\u00f1',0xf2:'\u00f2',0xf3:'\u00f3',0xf4:'\u00f4',0xf5:'\u00f5',0xf6:'\u00f6',0xf7:'\u00f7',0xf8:'\u00f8',0xf9:'\u00f9',0xfa:'\u00fa',0xfb:'\u00fb',0xfc:'\u00fc',0xfd:'\u00fd',0xfe:'\u00fe',0xff:'\u00ff'};
