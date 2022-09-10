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
package common ;import (_ce "fmt";_ea "io";_g "os";_f "path/filepath";_a "runtime";_e "time";);

// Debug logs debug message.
func (_da ConsoleLogger )Debug (format string ,args ...interface{}){if _da .LogLevel >=LogLevelDebug {_gd :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_da .output (_g .Stdout ,_gd ,format ,args ...);};};

// Trace logs trace message.
func (_bcc ConsoleLogger )Trace (format string ,args ...interface{}){if _bcc .LogLevel >=LogLevelTrace {_ee :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_bcc .output (_g .Stdout ,_ee ,format ,args ...);};};const _gaf ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";


// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fb WriterLogger )IsLogLevel (level LogLevel )bool {return _fb .LogLevel >=level };const _ceab =30;

// Notice logs notice message.
func (_gc WriterLogger )Notice (format string ,args ...interface{}){if _gc .LogLevel >=LogLevelNotice {_ga :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_gc .logToWriter (_gc .Output ,_ga ,format ,args ...);};};

// Notice logs notice message.
func (_bg ConsoleLogger )Notice (format string ,args ...interface{}){if _bg .LogLevel >=LogLevelNotice {_ad :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_bg .output (_g .Stdout ,_ad ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_aed ConsoleLogger )IsLogLevel (level LogLevel )bool {return _aed .LogLevel >=level };func (_eaa ConsoleLogger )output (_fad _ea .Writer ,_gbb string ,_ec string ,_ef ...interface{}){_afa (_fad ,_gbb ,_ec ,_ef ...);};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Debug logs debug message.
func (_eb WriterLogger )Debug (format string ,args ...interface{}){if _eb .LogLevel >=LogLevelDebug {_ba :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_eb .logToWriter (_eb .Output ,_ba ,format ,args ...);};};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// Warning logs warning message.
func (_df WriterLogger )Warning (format string ,args ...interface{}){if _df .LogLevel >=LogLevelWarning {_gdb :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_df .logToWriter (_df .Output ,_gdb ,format ,args ...);};};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};const _bb =15;

// Info logs info message.
func (_aeb ConsoleLogger )Info (format string ,args ...interface{}){if _aeb .LogLevel >=LogLevelInfo {_ed :="\u005bI\u004e\u0046\u004f\u005d\u0020";_aeb .output (_g .Stdout ,_ed ,format ,args ...);};};func (_fgc WriterLogger )logToWriter (_fcd _ea .Writer ,_edf string ,_cb string ,_ddf ...interface{}){_afa (_fcd ,_edf ,_cb ,_ddf );
};

// DummyLogger does nothing.
type DummyLogger struct{};

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat (t _e .Time )string {return t .Format (_gaf )+"\u0020\u0055\u0054\u0043"};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _ea .Writer )*WriterLogger {_efb :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_efb ;};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// Info logs info message.
func (_cc WriterLogger )Info (format string ,args ...interface{}){if _cc .LogLevel >=LogLevelInfo {_fdf :="\u005bI\u004e\u0046\u004f\u005d\u0020";_cc .logToWriter (_cc .Output ,_fdf ,format ,args ...);};};var Log Logger =DummyLogger {};var ReleasedAt =_e .Date (_ca ,_gbba ,_gcf ,_bb ,_ceab ,0,0,_e .UTC );


// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};const _gbba =9;

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _ea .Writer ;};const _gcf =10;const Version ="\u0033\u002e\u0033\u0038\u002e\u0030";const _ca =2022;const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;
LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// Error logs error message.
func (_ff WriterLogger )Error (format string ,args ...interface{}){if _ff .LogLevel >=LogLevelError {_fcb :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ff .logToWriter (_ff .Output ,_fcb ,format ,args ...);};};

// Warning logs warning message.
func (_bc ConsoleLogger )Warning (format string ,args ...interface{}){if _bc .LogLevel >=LogLevelWarning {_af :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_bc .output (_g .Stdout ,_af ,format ,args ...);};};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Error logs error message.
func (_bf ConsoleLogger )Error (format string ,args ...interface{}){if _bf .LogLevel >=LogLevelError {_gbg :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_bf .output (_g .Stdout ,_gbg ,format ,args ...);};};func _afa (_be _ea .Writer ,_bge string ,_aga string ,_ge ...interface{}){_ ,_cea ,_abg ,_bag :=_a .Caller (3);
if !_bag {_cea ="\u003f\u003f\u003f";_abg =0;}else {_cea =_f .Base (_cea );};_db :=_ce .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_bge ,_cea ,_abg )+_aga +"\u000a";_ce .Fprintf (_be ,_db ,_ge ...);};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// Trace logs trace message.
func (_edb WriterLogger )Trace (format string ,args ...interface{}){if _edb .LogLevel >=LogLevelTrace {_dd :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_edb .logToWriter (_edb .Output ,_dd ,format ,args ...);};};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_fg string ,_ac ...interface{});Warning (_ae string ,_gb ...interface{});Notice (_b string ,_ab ...interface{});Info (_aee string ,_gg ...interface{});Debug (_ag string ,_d ...interface{});Trace (_fc string ,_fa ...interface{});
IsLogLevel (_fd LogLevel )bool ;};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};