//  header block begin
// /usr/include/qt/QtCore/qtextstream.h
// #include <qtextstream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTextStream struct {
	*qtrt.CObject
}

func (this *QTextStream) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextStreamFromPointer(cthis unsafe.Pointer) *QTextStream {
	return &QTextStream{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtextstream.h:93
// index:0
// Public
// void QTextStream()
func NewQTextStream() *QTextStream {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:94
// index:1
// Public
// void QTextStream(class QIODevice *)
func NewQTextStream_1(device unsafe.Pointer) *QTextStream {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, device)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextstream.h:99
// index:0
// Public virtual
// void ~QTextStream()
func DeleteQTextStream(*QTextStream) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStreamD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:102
// index:0
// Public
// void setCodec(class QTextCodec *)
func (this *QTextStream) SetCodec(codec unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream8setCodecEP10QTextCodec", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), codec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:103
// index:1
// Public
// void setCodec(const char *)
func (this *QTextStream) SetCodec_1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream8setCodecEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:104
// index:0
// Public
// QTextCodec * codec()
func (this *QTextStream) Codec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream5codecEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:105
// index:0
// Public
// void setAutoDetectUnicode(_Bool)
func (this *QTextStream) SetAutoDetectUnicode(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream20setAutoDetectUnicodeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:106
// index:0
// Public
// bool autoDetectUnicode()
func (this *QTextStream) AutoDetectUnicode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream17autoDetectUnicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:107
// index:0
// Public
// void setGenerateByteOrderMark(_Bool)
func (this *QTextStream) SetGenerateByteOrderMark(generate bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream24setGenerateByteOrderMarkEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &generate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:108
// index:0
// Public
// bool generateByteOrderMark()
func (this *QTextStream) GenerateByteOrderMark() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream21generateByteOrderMarkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:111
// index:0
// Public
// void setLocale(const class QLocale &)
func (this *QTextStream) SetLocale(locale *QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream9setLocaleERK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:112
// index:0
// Public
// QLocale locale()
func (this *QTextStream) Locale() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream6localeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:114
// index:0
// Public
// void setDevice(class QIODevice *)
func (this *QTextStream) SetDevice(device unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:115
// index:0
// Public
// QIODevice * device()
func (this *QTextStream) Device() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:118
// index:0
// Public
// QString * string()
func (this *QTextStream) String() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream6stringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:120
// index:0
// Public
// QTextStream::Status status()
func (this *QTextStream) Status() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:121
// index:0
// Public
// void setStatus(enum QTextStream::Status)
func (this *QTextStream) SetStatus(status int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream9setStatusENS_6StatusE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &status)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:122
// index:0
// Public
// void resetStatus()
func (this *QTextStream) ResetStatus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream11resetStatusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:124
// index:0
// Public
// bool atEnd()
func (this *QTextStream) AtEnd() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:125
// index:0
// Public
// void reset()
func (this *QTextStream) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:126
// index:0
// Public
// void flush()
func (this *QTextStream) Flush() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream5flushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:127
// index:0
// Public
// bool seek(qint64)
func (this *QTextStream) Seek(pos int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream4seekEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:128
// index:0
// Public
// qint64 pos()
func (this *QTextStream) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:130
// index:0
// Public
// void skipWhiteSpace()
func (this *QTextStream) SkipWhiteSpace() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream14skipWhiteSpaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:132
// index:0
// Public
// QString readLine(qint64)
func (this *QTextStream) ReadLine(maxlen int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream8readLineEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:133
// index:0
// Public
// bool readLineInto(class QString *, qint64)
func (this *QTextStream) ReadLineInto(line unsafe.Pointer, maxlen int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream12readLineIntoEP7QStringx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), line, &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:134
// index:0
// Public
// QString readAll()
func (this *QTextStream) ReadAll() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream7readAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:135
// index:0
// Public
// QString read(qint64)
func (this *QTextStream) Read(maxlen int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream4readEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:137
// index:0
// Public
// void setFieldAlignment(enum QTextStream::FieldAlignment)
func (this *QTextStream) SetFieldAlignment(alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream17setFieldAlignmentENS_14FieldAlignmentE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:138
// index:0
// Public
// QTextStream::FieldAlignment fieldAlignment()
func (this *QTextStream) FieldAlignment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream14fieldAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:140
// index:0
// Public
// void setPadChar(class QChar)
func (this *QTextStream) SetPadChar(ch *QChar) {
	var convArg0 = ch.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream10setPadCharE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:141
// index:0
// Public
// QChar padChar()
func (this *QTextStream) PadChar() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream7padCharEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:143
// index:0
// Public
// void setFieldWidth(int)
func (this *QTextStream) SetFieldWidth(width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream13setFieldWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:144
// index:0
// Public
// int fieldWidth()
func (this *QTextStream) FieldWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream10fieldWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:147
// index:0
// Public
// QTextStream::NumberFlags numberFlags()
func (this *QTextStream) NumberFlags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream11numberFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:149
// index:0
// Public
// void setIntegerBase(int)
func (this *QTextStream) SetIntegerBase(base int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream14setIntegerBaseEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:150
// index:0
// Public
// int integerBase()
func (this *QTextStream) IntegerBase() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream11integerBaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:152
// index:0
// Public
// void setRealNumberNotation(enum QTextStream::RealNumberNotation)
func (this *QTextStream) SetRealNumberNotation(notation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream21setRealNumberNotationENS_18RealNumberNotationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &notation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:153
// index:0
// Public
// QTextStream::RealNumberNotation realNumberNotation()
func (this *QTextStream) RealNumberNotation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream18realNumberNotationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qtextstream.h:155
// index:0
// Public
// void setRealNumberPrecision(int)
func (this *QTextStream) SetRealNumberPrecision(precision int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextStream22setRealNumberPrecisionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:156
// index:0
// Public
// int realNumberPrecision()
func (this *QTextStream) RealNumberPrecision() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextStream19realNumberPrecisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end