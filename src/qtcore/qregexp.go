//  header block begin
// /usr/include/qt/QtCore/qregexp.h
// #include <qregexp.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 86
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
type QRegExp struct {
	*qtrt.CObject
}

func (this *QRegExp) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQRegExpFromPointer(cthis unsafe.Pointer) *QRegExp {
	return &QRegExp{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qregexp.h:70
// index:0
// Public
// void QRegExp()
func NewQRegExp() *QRegExp {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExpC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:71
// index:1
// Public
// void QRegExp(const class QString &, Qt::CaseSensitivity, enum QRegExp::PatternSyntax)
func NewQRegExp_1(pattern *QString, cs int, syntax int) *QRegExp {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExpC2ERK7QStringN2Qt15CaseSensitivityENS_13PatternSyntaxE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &cs, &syntax)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegExpFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qregexp.h:74
// index:0
// Public
// void ~QRegExp()
func DeleteQRegExp(*QRegExp) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExpD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:79
// index:0
// Public inline
// void swap(class QRegExp &)
func (this *QRegExp) Swap(other *QRegExp) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:84
// index:0
// Public
// bool isEmpty()
func (this *QRegExp) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:85
// index:0
// Public
// bool isValid()
func (this *QRegExp) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:86
// index:0
// Public
// QString pattern()
func (this *QRegExp) Pattern() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp7patternEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:87
// index:0
// Public
// void setPattern(const class QString &)
func (this *QRegExp) SetPattern(pattern *QString) {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp10setPatternERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:88
// index:0
// Public
// Qt::CaseSensitivity caseSensitivity()
func (this *QRegExp) CaseSensitivity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp15caseSensitivityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:89
// index:0
// Public
// void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QRegExp) SetCaseSensitivity(cs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp18setCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:90
// index:0
// Public
// QRegExp::PatternSyntax patternSyntax()
func (this *QRegExp) PatternSyntax() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp13patternSyntaxEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:91
// index:0
// Public
// void setPatternSyntax(enum QRegExp::PatternSyntax)
func (this *QRegExp) SetPatternSyntax(syntax int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp16setPatternSyntaxENS_13PatternSyntaxE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &syntax)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:93
// index:0
// Public
// bool isMinimal()
func (this *QRegExp) IsMinimal() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp9isMinimalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:94
// index:0
// Public
// void setMinimal(_Bool)
func (this *QRegExp) SetMinimal(minimal bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp10setMinimalEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &minimal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregexp.h:96
// index:0
// Public
// bool exactMatch(const class QString &)
func (this *QRegExp) ExactMatch(str *QString) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp10exactMatchERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:98
// index:0
// Public
// int indexIn(const class QString &, int, enum QRegExp::CaretMode)
func (this *QRegExp) IndexIn(str *QString, offset int, caretMode int) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp7indexInERK7QStringiNS_9CaretModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &offset, &caretMode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:99
// index:0
// Public
// int lastIndexIn(const class QString &, int, enum QRegExp::CaretMode)
func (this *QRegExp) LastIndexIn(str *QString, offset int, caretMode int) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp11lastIndexInERK7QStringiNS_9CaretModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &offset, &caretMode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:100
// index:0
// Public
// int matchedLength()
func (this *QRegExp) MatchedLength() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp13matchedLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:102
// index:0
// Public
// int captureCount()
func (this *QRegExp) CaptureCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp12captureCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:103
// index:0
// Public
// QStringList capturedTexts()
func (this *QRegExp) CapturedTexts() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp13capturedTextsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:104
// index:1
// Public
// QStringList capturedTexts()
func (this *QRegExp) CapturedTexts_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp13capturedTextsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:105
// index:0
// Public
// QString cap(int)
func (this *QRegExp) Cap(nth int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp3capEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:106
// index:1
// Public
// QString cap(int)
func (this *QRegExp) Cap_1(nth int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp3capEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:107
// index:0
// Public
// int pos(int)
func (this *QRegExp) Pos(nth int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp3posEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:108
// index:1
// Public
// int pos(int)
func (this *QRegExp) Pos_1(nth int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp3posEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &nth)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:109
// index:0
// Public
// QString errorString()
func (this *QRegExp) ErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QRegExp11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:110
// index:1
// Public
// QString errorString()
func (this *QRegExp) ErrorString_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qregexp.h:113
// index:0
// Public static
// QString escape(const class QString &)
func (this *QRegExp) Escape(str *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QRegExp6escapeERK7QString", ffiqt.FFI_TYPE_POINTER, str)
	gopp.ErrPrint(err, rv)
	return rv
}
func QRegExp_Escape(str *QString) {
	var nilthis *QRegExp
	nilthis.Escape(str)
}

//  body block end