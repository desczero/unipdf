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
package common ;import (_f "fmt";_a "io";_fb "os";_fac "path/filepath";_c "runtime";_fa "time";);

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Trace logs trace message.
func (_dg ConsoleLogger )Trace (format string ,args ...interface{}){if _dg .LogLevel >=LogLevelTrace {_da :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_dg .output (_fb .Stdout ,_da ,format ,args ...);};};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);const _gdb =30;

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_ad string ,_ce ...interface{});Warning (_ae string ,_b ...interface{});Notice (_g string ,_d ...interface{});Info (_af string ,_fd ...interface{});Debug (_gd string ,_gc ...interface{});Trace (_bb string ,_gcd ...interface{});
IsLogLevel (_cd LogLevel )bool ;};var ReleasedAt =_fa .Date (_ada ,_cc ,_cef ,_bbe ,_gdb ,0,0,_fa .UTC );

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_gdf WriterLogger )IsLogLevel (level LogLevel )bool {return _gdf .LogLevel >=level };

// Info logs info message.
func (_ced WriterLogger )Info (format string ,args ...interface{}){if _ced .LogLevel >=LogLevelInfo {_def :="\u005bI\u004e\u0046\u004f\u005d\u0020";_ced .logToWriter (_ced .Output ,_def ,format ,args ...);};};

// Error logs error message.
func (_aee ConsoleLogger )Error (format string ,args ...interface{}){if _aee .LogLevel >=LogLevelError {_fba :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_aee .output (_fb .Stdout ,_fba ,format ,args ...);};};const _bbe =15;

// Notice logs notice message.
func (_fbac ConsoleLogger )Notice (format string ,args ...interface{}){if _fbac .LogLevel >=LogLevelNotice {_fab :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_fbac .output (_fb .Stdout ,_fab ,format ,args ...);};};

// Warning logs warning message.
func (_bba ConsoleLogger )Warning (format string ,args ...interface{}){if _bba .LogLevel >=LogLevelWarning {_dd :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_bba .output (_fb .Stdout ,_dd ,format ,args ...);};};const _cef =6;

// Debug logs debug message.
func (_cec WriterLogger )Debug (format string ,args ...interface{}){if _cec .LogLevel >=LogLevelDebug {_acd :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_cec .logToWriter (_cec .Output ,_acd ,format ,args ...);};};

// Notice logs notice message.
func (_fdd WriterLogger )Notice (format string ,args ...interface{}){if _fdd .LogLevel >=LogLevelNotice {_ea :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_fdd .logToWriter (_fdd .Output ,_ea ,format ,args ...);};};

// Error logs error message.
func (_eg WriterLogger )Error (format string ,args ...interface{}){if _eg .LogLevel >=LogLevelError {_afb :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_eg .logToWriter (_eg .Output ,_afb ,format ,args ...);};};

// Trace logs trace message.
func (_fea WriterLogger )Trace (format string ,args ...interface{}){if _fea .LogLevel >=LogLevelTrace {_ff :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_fea .logToWriter (_fea .Output ,_ff ,format ,args ...);};};

// Warning logs warning message.
func (_df WriterLogger )Warning (format string ,args ...interface{}){if _df .LogLevel >=LogLevelWarning {_bc :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_df .logToWriter (_df .Output ,_bc ,format ,args ...);};};var Log Logger =DummyLogger {};


// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// LogLevel is the verbosity level for logging.
type LogLevel int ;const _cc =4;

// Info logs info message.
func (_ed ConsoleLogger )Info (format string ,args ...interface{}){if _ed .LogLevel >=LogLevelInfo {_ba :="\u005bI\u004e\u0046\u004f\u005d\u0020";_ed .output (_fb .Stdout ,_ba ,format ,args ...);};};const Version ="\u0033\u002e\u0034\u0035\u002e\u0030";


// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_gb ConsoleLogger )IsLogLevel (level LogLevel )bool {return _gb .LogLevel >=level };func (_ga WriterLogger )logToWriter (_fabd _a .Writer ,_dga string ,_gg string ,_gge ...interface{}){_fabe (_fabd ,_dga ,_gg ,_gge );};func (_de ConsoleLogger )output (_ag _a .Writer ,_ef string ,_fe string ,_ac ...interface{}){_fabe (_ag ,_ef ,_fe ,_ac ...);
};func _fabe (_cdg _a .Writer ,_fbe string ,_aca string ,_ffc ...interface{}){_ ,_age ,_fdb ,_dgd :=_c .Caller (3);if !_dgd {_age ="\u003f\u003f\u003f";_fdb =0;}else {_age =_fac .Base (_age );};_ee :=_f .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_fbe ,_age ,_fdb )+_aca +"\u000a";
_f .Fprintf (_cdg ,_ee ,_ffc ...);};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// Debug logs debug message.
func (_afd ConsoleLogger )Debug (format string ,args ...interface{}){if _afd .LogLevel >=LogLevelDebug {_aa :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_afd .output (_fb .Stdout ,_aa ,format ,args ...);};};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat (t _fa .Time )string {return t .Format (_dfe )+"\u0020\u0055\u0054\u0043"};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _a .Writer ;};

// DummyLogger does nothing.
type DummyLogger struct{};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _a .Writer )*WriterLogger {_dda :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_dda ;};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };const _dfe ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";const _ada =2023;