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

// Package common contains common properties used by the subpackages.
package common ;import (_e "fmt";_dg "io";_b "os";_g "path/filepath";_da "runtime";_c "time";);

// Debug logs debug message.
func (_fc WriterLogger )Debug (format string ,args ...interface{}){if _fc .LogLevel >=LogLevelDebug {_bb :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_fc .logToWriter (_fc .Output ,_bb ,format ,args ...);};};

// Warning logs warning message.
func (_cg WriterLogger )Warning (format string ,args ...interface{}){if _cg .LogLevel >=LogLevelWarning {_dcc :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_cg .logToWriter (_cg .Output ,_dcc ,format ,args ...);};};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};func (_cd WriterLogger )logToWriter (_cf _dg .Writer ,_aaa string ,_cb string ,_bgc ...interface{}){_fgg (_cf ,_aaa ,_cb ,_bgc );};

// Notice logs notice message.
func (_aa WriterLogger )Notice (format string ,args ...interface{}){if _aa .LogLevel >=LogLevelNotice {_df :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_aa .logToWriter (_aa .Output ,_df ,format ,args ...);};};

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat (t _c .Time )string {return t .Format (_daf )+"\u0020\u0055\u0054\u0043"};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// Error logs error message.
func (_eeb WriterLogger )Error (format string ,args ...interface{}){if _eeb .LogLevel >=LogLevelError {_ea :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_eeb .logToWriter (_eeb .Output ,_ea ,format ,args ...);};};

// Warning logs warning message.
func (_fa ConsoleLogger )Warning (format string ,args ...interface{}){if _fa .LogLevel >=LogLevelWarning {_ge :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_fa .output (_b .Stdout ,_ge ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_db ConsoleLogger )IsLogLevel (level LogLevel )bool {return _db .LogLevel >=level };

// Info logs info message.
func (_cad WriterLogger )Info (format string ,args ...interface{}){if _cad .LogLevel >=LogLevelInfo {_eb :="\u005bI\u004e\u0046\u004f\u005d\u0020";_cad .logToWriter (_cad .Output ,_eb ,format ,args ...);};};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _dg .Writer )*WriterLogger {_ba :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_ba ;};

// Debug logs debug message.
func (_add ConsoleLogger )Debug (format string ,args ...interface{}){if _add .LogLevel >=LogLevelDebug {_aga :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_add .output (_b .Stdout ,_aga ,format ,args ...);};};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};var ReleasedAt =_c .Date (_ebg ,_agb ,_aed ,_gf ,_adg ,0,0,_c .UTC );

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_dgf WriterLogger )IsLogLevel (level LogLevel )bool {return _dgf .LogLevel >=level };

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };var Log Logger =DummyLogger {};

// Trace logs trace message.
func (_dab WriterLogger )Trace (format string ,args ...interface{}){if _dab .LogLevel >=LogLevelTrace {_cc :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_dab .logToWriter (_dab .Output ,_cc ,format ,args ...);};};const Version ="\u0033\u002e\u0035\u0036\u002e\u0030";


// DummyLogger does nothing.
type DummyLogger struct{};const _aed =19;

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};const _adg =30;

// Info logs info message.
func (_dc ConsoleLogger )Info (format string ,args ...interface{}){if _dc .LogLevel >=LogLevelInfo {_dd :="\u005bI\u004e\u0046\u004f\u005d\u0020";_dc .output (_b .Stdout ,_dd ,format ,args ...);};};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_bf string ,_dag ...interface{});Warning (_a string ,_f ...interface{});Notice (_fg string ,_gb ...interface{});Info (_ce string ,_dgg ...interface{});Debug (_cee string ,_de ...interface{});Trace (_ca string ,_ee ...interface{});
IsLogLevel (_ag LogLevel )bool ;};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};const _daf ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };const _agb =3;

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Trace logs trace message.
func (_bg ConsoleLogger )Trace (format string ,args ...interface{}){if _bg .LogLevel >=LogLevelTrace {_ae :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_bg .output (_b .Stdout ,_ae ,format ,args ...);};};const _ebg =2024;

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _dg .Writer ;};

// Notice logs notice message.
func (_ac ConsoleLogger )Notice (format string ,args ...interface{}){if _ac .LogLevel >=LogLevelNotice {_ad :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_ac .output (_b .Stdout ,_ad ,format ,args ...);};};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;
LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);func (_gbg ConsoleLogger )output (_ddg _dg .Writer ,_ga string ,_cab string ,_gec ...interface{}){_fgg (_ddg ,_ga ,_cab ,_gec ...);};

// Error logs error message.
func (_ec ConsoleLogger )Error (format string ,args ...interface{}){if _ec .LogLevel >=LogLevelError {_caa :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ec .output (_b .Stdout ,_caa ,format ,args ...);};};func _fgg (_eee _dg .Writer ,_dae string ,_dcd string ,_fd ...interface{}){_ ,_gd ,_ddf ,_fb :=_da .Caller (3);
if !_fb {_gd ="\u003f\u003f\u003f";_ddf =0;}else {_gd =_g .Base (_gd );};_ef :=_e .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_dae ,_gd ,_ddf )+_dcd +"\u000a";_e .Fprintf (_eee ,_ef ,_fd ...);};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};const _gf =15;