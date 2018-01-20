//  header block begin
// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
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
type QString struct {
	*qtrt.CObject
}

func (this *QString) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQStringFromPointer(cthis unsafe.Pointer) *QString {
	return &QString{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstring.h:217
// index:0
// Public inline
// void QString()
func NewQString() *QString {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:218
// index:1
// Public
// void QString(const class QChar *, int)
func NewQString_1(unicode unsafe.Pointer, size int) *QString {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringC2EPK5QChari", ffiqt.FFI_TYPE_VOID, cthis, unicode, &size)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:219
// index:2
// Public
// void QString(class QChar)
func NewQString_2(c *QChar) *QString {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringC2E5QChar", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:220
// index:3
// Public
// void QString(int, class QChar)
func NewQString_3(size int, c *QChar) *QString {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg1 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringC2Ei5QChar", ffiqt.FFI_TYPE_VOID, cthis, &size, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:221
// index:4
// Public inline
// void QString(class QLatin1String)
func NewQString_4(latin1 *QLatin1String) *QString {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = latin1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringC2E13QLatin1String", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:682
// index:5
// Public inline
// void QString(const char *)
func NewQString_5(ch string) *QString {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = qtrt.CString(ch)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:685
// index:6
// Public inline
// void QString(const class QByteArray &)
func NewQString_6(a *QByteArray) *QString {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:811
// index:7
// Public
// void QString(int, Qt::Initialization)
func NewQString_7(size int, arg1 int) *QString {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringC2EiN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &size, &arg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:223
// index:0
// Public inline
// void ~QString()
func DeleteQString(*QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QStringD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:232
// index:0
// Public inline
// void swap(class QString &)
func (this *QString) Swap(other *QString) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:233
// index:0
// Public inline
// int size()
func (this *QString) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:234
// index:0
// Public inline
// int count()
func (this *QString) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:335
// index:1
// Public
// int count(class QChar, Qt::CaseSensitivity)
func (this *QString) Count_1(c *QChar, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5countE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:336
// index:2
// Public
// int count(const class QString &, Qt::CaseSensitivity)
func (this *QString) Count_2(s *QString, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5countERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:337
// index:3
// Public
// int count(const class QStringRef &, Qt::CaseSensitivity)
func (this *QString) Count_3(s *QStringRef, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5countERK10QStringRefN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:343
// index:4
// Public
// int count(const class QRegExp &)
func (this *QString) Count_4(arg0 *QRegExp) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5countERK7QRegExp", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:357
// index:5
// Public
// int count(const class QRegularExpression &)
func (this *QString) Count_5(re *QRegularExpression) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5countERK18QRegularExpression", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:235
// index:0
// Public inline
// int length()
func (this *QString) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:236
// index:0
// Public inline
// bool isEmpty()
func (this *QString) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:237
// index:0
// Public
// void resize(int)
func (this *QString) Resize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6resizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:238
// index:1
// Public
// void resize(int, class QChar)
func (this *QString) Resize_1(size int, fillChar *QChar) {
	var convArg1 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6resizeEi5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:240
// index:0
// Public
// QString & fill(class QChar, int)
func (this *QString) Fill(c *QChar, size int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString4fillE5QChari", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &size)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:241
// index:0
// Public
// void truncate(int)
func (this *QString) Truncate(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString8truncateEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:242
// index:0
// Public
// void chop(int)
func (this *QString) Chop(n int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString4chopEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:244
// index:0
// Public
// int capacity()
func (this *QString) Capacity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8capacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:245
// index:0
// Public inline
// void reserve(int)
func (this *QString) Reserve(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7reserveEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:246
// index:0
// Public inline
// void squeeze()
func (this *QString) Squeeze() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7squeezeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:248
// index:0
// Public inline
// const QChar * unicode()
func (this *QString) Unicode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:249
// index:0
// Public inline
// QChar * data()
func (this *QString) Data() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:250
// index:1
// Public inline
// const QChar * data()
func (this *QString) Data_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:251
// index:0
// Public inline
// const QChar * constData()
func (this *QString) ConstData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString9constDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:253
// index:0
// Public inline
// void detach()
func (this *QString) Detach() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:254
// index:0
// Public inline
// bool isDetached()
func (this *QString) IsDetached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:255
// index:0
// Public inline
// bool isSharedWith(const class QString &)
func (this *QString) IsSharedWith(other *QString) interface{} {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString12isSharedWithERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:256
// index:0
// Public
// void clear()
func (this *QString) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:258
// index:0
// Public inline
// const QChar at(int)
func (this *QString) At(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString2atEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:264
// index:0
// Public inline
// QChar front()
func (this *QString) Front() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5frontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:265
// index:1
// Public inline
// QCharRef front()
func (this *QString) Front_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString5frontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:266
// index:0
// Public inline
// QChar back()
func (this *QString) Back() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString4backEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:267
// index:1
// Public inline
// QCharRef back()
func (this *QString) Back_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString4backEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:269
// index:0
// Public
// QString arg(qlonglong, int, int, class QChar)
func (this *QString) Arg(a int64, fieldwidth int, base int, fillChar *QChar) interface{} {
	var convArg3 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argExii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldwidth, &base, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:271
// index:1
// Public
// QString arg(qulonglong, int, int, class QChar)
func (this *QString) Arg_1(a uint64, fieldwidth int, base int, fillChar *QChar) interface{} {
	var convArg3 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argEyii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldwidth, &base, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:273
// index:2
// Public
// QString arg(long, int, int, class QChar)
func (this *QString) Arg_2(a int, fieldwidth int, base int, fillChar *QChar) interface{} {
	var convArg3 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argElii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldwidth, &base, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:275
// index:3
// Public
// QString arg(ulong, int, int, class QChar)
func (this *QString) Arg_3(a uint, fieldwidth int, base int, fillChar *QChar) interface{} {
	var convArg3 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argEmii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldwidth, &base, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:277
// index:4
// Public
// QString arg(int, int, int, class QChar)
func (this *QString) Arg_4(a int, fieldWidth int, base int, fillChar *QChar) interface{} {
	var convArg3 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argEiii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldWidth, &base, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:279
// index:5
// Public
// QString arg(uint, int, int, class QChar)
func (this *QString) Arg_5(a uint, fieldWidth int, base int, fillChar *QChar) interface{} {
	var convArg3 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argEjii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldWidth, &base, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:281
// index:6
// Public
// QString arg(short, int, int, class QChar)
func (this *QString) Arg_6(a int16, fieldWidth int, base int, fillChar *QChar) interface{} {
	var convArg3 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argEsii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldWidth, &base, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:283
// index:7
// Public
// QString arg(ushort, int, int, class QChar)
func (this *QString) Arg_7(a uint16, fieldWidth int, base int, fillChar *QChar) interface{} {
	var convArg3 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argEtii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldWidth, &base, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:285
// index:8
// Public
// QString arg(double, int, char, int, class QChar)
func (this *QString) Arg_8(a float64, fieldWidth int, fmt byte, prec int, fillChar *QChar) interface{} {
	var convArg4 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argEdici5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldWidth, &fmt, &prec, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:287
// index:9
// Public
// QString arg(char, int, class QChar)
func (this *QString) Arg_9(a byte, fieldWidth int, fillChar *QChar) interface{} {
	var convArg2 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argEci5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:289
// index:10
// Public
// QString arg(class QChar, int, class QChar)
func (this *QString) Arg_10(a *QChar, fieldWidth int, fillChar *QChar) interface{} {
	var convArg0 = a.GetCthis()
	var convArg2 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argE5QChariS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:292
// index:11
// Public
// QString arg(const class QString &, int, class QChar)
func (this *QString) Arg_11(a *QString, fieldWidth int, fillChar *QChar) interface{} {
	var convArg0 = a.GetCthis()
	var convArg2 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_i5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:295
// index:12
// Public
// QString arg(class QStringView, int, class QChar)
func (this *QString) Arg_12(a *QStringView, fieldWidth int, fillChar *QChar) interface{} {
	var convArg0 = a.GetCthis()
	var convArg2 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argE11QStringViewi5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:297
// index:13
// Public
// QString arg(class QLatin1String, int, class QChar)
func (this *QString) Arg_13(a *QLatin1String, fieldWidth int, fillChar *QChar) interface{} {
	var convArg0 = a.GetCthis()
	var convArg2 = fillChar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argE13QLatin1Stringi5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:299
// index:14
// Public
// QString arg(const class QString &, const class QString &)
func (this *QString) Arg_14(a1 *QString, a2 *QString) interface{} {
	var convArg0 = a1.GetCthis()
	var convArg1 = a2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:300
// index:15
// Public
// QString arg(const class QString &, const class QString &, const class QString &)
func (this *QString) Arg_15(a1 *QString, a2 *QString, a3 *QString) interface{} {
	var convArg0 = a1.GetCthis()
	var convArg1 = a2.GetCthis()
	var convArg2 = a3.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:301
// index:16
// Public
// QString arg(const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QString) Arg_16(a1 *QString, a2 *QString, a3 *QString, a4 *QString) interface{} {
	var convArg0 = a1.GetCthis()
	var convArg1 = a2.GetCthis()
	var convArg2 = a3.GetCthis()
	var convArg3 = a4.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:303
// index:17
// Public
// QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QString) Arg_17(a1 *QString, a2 *QString, a3 *QString, a4 *QString, a5 *QString) interface{} {
	var convArg0 = a1.GetCthis()
	var convArg1 = a2.GetCthis()
	var convArg2 = a3.GetCthis()
	var convArg3 = a4.GetCthis()
	var convArg4 = a5.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:305
// index:18
// Public
// QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QString) Arg_18(a1 *QString, a2 *QString, a3 *QString, a4 *QString, a5 *QString, a6 *QString) interface{} {
	var convArg0 = a1.GetCthis()
	var convArg1 = a2.GetCthis()
	var convArg2 = a3.GetCthis()
	var convArg3 = a4.GetCthis()
	var convArg4 = a5.GetCthis()
	var convArg5 = a6.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:307
// index:19
// Public
// QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QString) Arg_19(a1 *QString, a2 *QString, a3 *QString, a4 *QString, a5 *QString, a6 *QString, a7 *QString) interface{} {
	var convArg0 = a1.GetCthis()
	var convArg1 = a2.GetCthis()
	var convArg2 = a3.GetCthis()
	var convArg3 = a4.GetCthis()
	var convArg4 = a5.GetCthis()
	var convArg5 = a6.GetCthis()
	var convArg6 = a7.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:310
// index:20
// Public
// QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QString) Arg_20(a1 *QString, a2 *QString, a3 *QString, a4 *QString, a5 *QString, a6 *QString, a7 *QString, a8 *QString) interface{} {
	var convArg0 = a1.GetCthis()
	var convArg1 = a2.GetCthis()
	var convArg2 = a3.GetCthis()
	var convArg3 = a4.GetCthis()
	var convArg4 = a5.GetCthis()
	var convArg5 = a6.GetCthis()
	var convArg6 = a7.GetCthis()
	var convArg7 = a8.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:313
// index:21
// Public
// QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QString) Arg_21(a1 *QString, a2 *QString, a3 *QString, a4 *QString, a5 *QString, a6 *QString, a7 *QString, a8 *QString, a9 *QString) interface{} {
	var convArg0 = a1.GetCthis()
	var convArg1 = a2.GetCthis()
	var convArg2 = a3.GetCthis()
	var convArg3 = a4.GetCthis()
	var convArg4 = a5.GetCthis()
	var convArg5 = a6.GetCthis()
	var convArg6 = a7.GetCthis()
	var convArg7 = a8.GetCthis()
	var convArg8 = a9.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:322
// index:0
// Public
// int indexOf(class QChar, int, Qt::CaseSensitivity)
func (this *QString) IndexOf(c *QChar, from int, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7indexOfE5QChariN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:323
// index:1
// Public
// int indexOf(const class QString &, int, Qt::CaseSensitivity)
func (this *QString) IndexOf_1(s *QString, from int, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7indexOfERKS_iN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:324
// index:2
// Public
// int indexOf(class QLatin1String, int, Qt::CaseSensitivity)
func (this *QString) IndexOf_2(s *QLatin1String, from int, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:325
// index:3
// Public
// int indexOf(const class QStringRef &, int, Qt::CaseSensitivity)
func (this *QString) IndexOf_3(s *QStringRef, from int, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7indexOfERK10QStringRefiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:340
// index:4
// Public
// int indexOf(const class QRegExp &, int)
func (this *QString) IndexOf_4(arg0 *QRegExp, from int) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7indexOfERK7QRegExpi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:345
// index:5
// Public
// int indexOf(class QRegExp &, int)
func (this *QString) IndexOf_5(arg0 *QRegExp, from int) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7indexOfER7QRegExpi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:351
// index:6
// Public
// int indexOf(const class QRegularExpression &, int)
func (this *QString) IndexOf_6(re *QRegularExpression, from int) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7indexOfERK18QRegularExpressioni", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:352
// index:7
// Public
// int indexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
func (this *QString) IndexOf_7(re *QRegularExpression, from int, rmatch unsafe.Pointer) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, rmatch)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:326
// index:0
// Public
// int lastIndexOf(class QChar, int, Qt::CaseSensitivity)
func (this *QString) LastIndexOf(c *QChar, from int, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11lastIndexOfE5QChariN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:327
// index:1
// Public
// int lastIndexOf(const class QString &, int, Qt::CaseSensitivity)
func (this *QString) LastIndexOf_1(s *QString, from int, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERKS_iN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:328
// index:2
// Public
// int lastIndexOf(class QLatin1String, int, Qt::CaseSensitivity)
func (this *QString) LastIndexOf_2(s *QLatin1String, from int, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:329
// index:3
// Public
// int lastIndexOf(const class QStringRef &, int, Qt::CaseSensitivity)
func (this *QString) LastIndexOf_3(s *QStringRef, from int, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK10QStringRefiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:341
// index:4
// Public
// int lastIndexOf(const class QRegExp &, int)
func (this *QString) LastIndexOf_4(arg0 *QRegExp, from int) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK7QRegExpi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:346
// index:5
// Public
// int lastIndexOf(class QRegExp &, int)
func (this *QString) LastIndexOf_5(arg0 *QRegExp, from int) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11lastIndexOfER7QRegExpi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:353
// index:6
// Public
// int lastIndexOf(const class QRegularExpression &, int)
func (this *QString) LastIndexOf_6(re *QRegularExpression, from int) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK18QRegularExpressioni", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:354
// index:7
// Public
// int lastIndexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
func (this *QString) LastIndexOf_7(re *QRegularExpression, from int, rmatch unsafe.Pointer) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from, rmatch)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:331
// index:0
// Public inline
// bool contains(class QChar, Qt::CaseSensitivity)
func (this *QString) Contains(c *QChar, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8containsE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:332
// index:1
// Public inline
// bool contains(const class QString &, Qt::CaseSensitivity)
func (this *QString) Contains_1(s *QString, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8containsERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:333
// index:2
// Public inline
// bool contains(class QLatin1String, Qt::CaseSensitivity)
func (this *QString) Contains_2(s *QLatin1String, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8containsE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:334
// index:3
// Public inline
// bool contains(const class QStringRef &, Qt::CaseSensitivity)
func (this *QString) Contains_3(s *QStringRef, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8containsERK10QStringRefN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:342
// index:4
// Public inline
// bool contains(const class QRegExp &)
func (this *QString) Contains_4(rx *QRegExp) interface{} {
	var convArg0 = rx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8containsERK7QRegExp", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:347
// index:5
// Public inline
// bool contains(class QRegExp &)
func (this *QString) Contains_5(rx *QRegExp) interface{} {
	var convArg0 = rx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8containsER7QRegExp", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:355
// index:6
// Public
// bool contains(const class QRegularExpression &)
func (this *QString) Contains_6(re *QRegularExpression) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8containsERK18QRegularExpression", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:356
// index:7
// Public
// bool contains(const class QRegularExpression &, class QRegularExpressionMatch *)
func (this *QString) Contains_7(re *QRegularExpression, match_ unsafe.Pointer) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, match_)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:377
// index:0
// Public
// QString left(int)
func (this *QString) Left(n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString4leftEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:378
// index:0
// Public
// QString right(int)
func (this *QString) Right(n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5rightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:379
// index:0
// Public
// QString mid(int, int)
func (this *QString) Mid(position int, n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3midEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position, &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:380
// index:0
// Public inline
// QString chopped(int)
func (this *QString) Chopped(n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7choppedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:384
// index:0
// Public
// QStringRef leftRef(int)
func (this *QString) LeftRef(n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7leftRefEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:385
// index:0
// Public
// QStringRef rightRef(int)
func (this *QString) RightRef(n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8rightRefEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:386
// index:0
// Public
// QStringRef midRef(int, int)
func (this *QString) MidRef(position int, n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString6midRefEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position, &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:389
// index:0
// Public
// bool startsWith(const class QString &, Qt::CaseSensitivity)
func (this *QString) StartsWith(s *QString, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10startsWithERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:390
// index:1
// Public
// bool startsWith(const class QStringRef &, Qt::CaseSensitivity)
func (this *QString) StartsWith_1(s *QStringRef, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10startsWithERK10QStringRefN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:392
// index:2
// Public inline
// bool startsWith(class QStringView, Qt::CaseSensitivity)
func (this *QString) StartsWith_2(s *QStringView, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10startsWithE11QStringViewN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:394
// index:3
// Public
// bool startsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QString) StartsWith_3(s *QLatin1String, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:395
// index:4
// Public
// bool startsWith(class QChar, Qt::CaseSensitivity)
func (this *QString) StartsWith_4(c *QChar, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10startsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:398
// index:0
// Public
// bool endsWith(const class QString &, Qt::CaseSensitivity)
func (this *QString) EndsWith(s *QString, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8endsWithERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:399
// index:1
// Public
// bool endsWith(const class QStringRef &, Qt::CaseSensitivity)
func (this *QString) EndsWith_1(s *QStringRef, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8endsWithERK10QStringRefN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:401
// index:2
// Public inline
// bool endsWith(class QStringView, Qt::CaseSensitivity)
func (this *QString) EndsWith_2(s *QStringView, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8endsWithE11QStringViewN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:403
// index:3
// Public
// bool endsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QString) EndsWith_3(s *QLatin1String, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:404
// index:4
// Public
// bool endsWith(class QChar, Qt::CaseSensitivity)
func (this *QString) EndsWith_4(c *QChar, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8endsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:406
// index:0
// Public
// QString leftJustified(int, class QChar, _Bool)
func (this *QString) LeftJustified(width int, fill *QChar, trunc bool) interface{} {
	var convArg1 = fill.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString13leftJustifiedEi5QCharb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width, convArg1, &trunc)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:407
// index:0
// Public
// QString rightJustified(int, class QChar, _Bool)
func (this *QString) RightJustified(width int, fill *QChar, trunc bool) interface{} {
	var convArg1 = fill.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString14rightJustifiedEi5QCharb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width, convArg1, &trunc)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:417
// index:0
// Public inline
// QString toLower()
func (this *QString) ToLower() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR7QString7toLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:419
// index:1
// Public inline
// QString toLower()
func (this *QString) ToLower_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO7QString7toLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:421
// index:0
// Public inline
// QString toUpper()
func (this *QString) ToUpper() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR7QString7toUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:423
// index:1
// Public inline
// QString toUpper()
func (this *QString) ToUpper_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO7QString7toUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:425
// index:0
// Public inline
// QString toCaseFolded()
func (this *QString) ToCaseFolded() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR7QString12toCaseFoldedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:427
// index:1
// Public inline
// QString toCaseFolded()
func (this *QString) ToCaseFolded_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO7QString12toCaseFoldedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:429
// index:0
// Public inline
// QString trimmed()
func (this *QString) Trimmed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR7QString7trimmedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:431
// index:1
// Public inline
// QString trimmed()
func (this *QString) Trimmed_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO7QString7trimmedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:433
// index:0
// Public inline
// QString simplified()
func (this *QString) Simplified() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR7QString10simplifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:435
// index:1
// Public inline
// QString simplified()
func (this *QString) Simplified_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO7QString10simplifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:447
// index:0
// Public
// QString toHtmlEscaped()
func (this *QString) ToHtmlEscaped() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString13toHtmlEscapedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:478
// index:0
// Public
// QString & remove(int, int)
func (this *QString) Remove(i int, len int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6removeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:479
// index:1
// Public
// QString & remove(class QChar, Qt::CaseSensitivity)
func (this *QString) Remove_1(c *QChar, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6removeE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:480
// index:2
// Public
// QString & remove(const class QString &, Qt::CaseSensitivity)
func (this *QString) Remove_2(s *QString, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6removeERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:495
// index:3
// Public inline
// QString & remove(const class QRegExp &)
func (this *QString) Remove_3(rx *QRegExp) interface{} {
	var convArg0 = rx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6removeERK7QRegExp", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:500
// index:4
// Public inline
// QString & remove(const class QRegularExpression &)
func (this *QString) Remove_4(re *QRegularExpression) interface{} {
	var convArg0 = re.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6removeERK18QRegularExpression", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:481
// index:0
// Public
// QString & replace(int, int, class QChar)
func (this *QString) Replace(i int, len int, after *QChar) interface{} {
	var convArg2 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceEii5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, &len, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:482
// index:1
// Public
// QString & replace(int, int, const class QChar *, int)
func (this *QString) Replace_1(i int, len int, s unsafe.Pointer, slen int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceEiiPK5QChari", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, &len, s, &slen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:483
// index:2
// Public
// QString & replace(int, int, const class QString &)
func (this *QString) Replace_2(i int, len int, after *QString) interface{} {
	var convArg2 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceEiiRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, &len, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:484
// index:3
// Public
// QString & replace(class QChar, class QChar, Qt::CaseSensitivity)
func (this *QString) Replace_3(before *QChar, after *QChar, cs int) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceE5QCharS0_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:485
// index:4
// Public
// QString & replace(const class QChar *, int, const class QChar *, int, Qt::CaseSensitivity)
func (this *QString) Replace_4(before unsafe.Pointer, blen int, after unsafe.Pointer, alen int, cs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceEPK5QChariS2_iN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), before, &blen, after, &alen, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:486
// index:5
// Public
// QString & replace(class QLatin1String, class QLatin1String, Qt::CaseSensitivity)
func (this *QString) Replace_5(before *QLatin1String, after *QLatin1String, cs int) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceE13QLatin1StringS0_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:487
// index:6
// Public
// QString & replace(class QLatin1String, const class QString &, Qt::CaseSensitivity)
func (this *QString) Replace_6(before *QLatin1String, after *QString, cs int) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceE13QLatin1StringRKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:488
// index:7
// Public
// QString & replace(const class QString &, class QLatin1String, Qt::CaseSensitivity)
func (this *QString) Replace_7(before *QString, after *QLatin1String, cs int) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceERKS_13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:489
// index:8
// Public
// QString & replace(const class QString &, const class QString &, Qt::CaseSensitivity)
func (this *QString) Replace_8(before *QString, after *QString, cs int) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceERKS_S1_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:491
// index:9
// Public
// QString & replace(class QChar, const class QString &, Qt::CaseSensitivity)
func (this *QString) Replace_9(c *QChar, after *QString, cs int) interface{} {
	var convArg0 = c.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceE5QCharRKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:492
// index:10
// Public
// QString & replace(class QChar, class QLatin1String, Qt::CaseSensitivity)
func (this *QString) Replace_10(c *QChar, after *QLatin1String, cs int) interface{} {
	var convArg0 = c.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceE5QChar13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:494
// index:11
// Public
// QString & replace(const class QRegExp &, const class QString &)
func (this *QString) Replace_11(rx *QRegExp, after *QString) interface{} {
	var convArg0 = rx.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceERK7QRegExpRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:499
// index:12
// Public
// QString & replace(const class QRegularExpression &, const class QString &)
func (this *QString) Replace_12(re *QRegularExpression, after *QString) interface{} {
	var convArg0 = re.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7replaceERK18QRegularExpressionRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:506
// index:0
// Public
// QStringList split(const class QString &, enum QString::SplitBehavior, Qt::CaseSensitivity)
func (this *QString) Split(sep *QString, behavior int, cs int) interface{} {
	var convArg0 = sep.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5splitERKS_NS_13SplitBehaviorEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:510
// index:1
// Public
// QStringList split(class QChar, enum QString::SplitBehavior, Qt::CaseSensitivity)
func (this *QString) Split_1(sep *QChar, behavior int, cs int) interface{} {
	var convArg0 = sep.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5splitE5QCharNS_13SplitBehaviorEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:515
// index:2
// Public
// QStringList split(const class QRegExp &, enum QString::SplitBehavior)
func (this *QString) Split_2(sep *QRegExp, behavior int) interface{} {
	var convArg0 = sep.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5splitERK7QRegExpNS_13SplitBehaviorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:519
// index:3
// Public
// QStringList split(const class QRegularExpression &, enum QString::SplitBehavior)
func (this *QString) Split_3(sep *QRegularExpression, behavior int) interface{} {
	var convArg0 = sep.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5splitERK18QRegularExpressionNS_13SplitBehaviorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:508
// index:0
// Public
// QVector<QStringRef> splitRef(const class QString &, enum QString::SplitBehavior, Qt::CaseSensitivity)
func (this *QString) SplitRef(sep *QString, behavior int, cs int) interface{} {
	var convArg0 = sep.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8splitRefERKS_NS_13SplitBehaviorEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:512
// index:1
// Public
// QVector<QStringRef> splitRef(class QChar, enum QString::SplitBehavior, Qt::CaseSensitivity)
func (this *QString) SplitRef_1(sep *QChar, behavior int, cs int) interface{} {
	var convArg0 = sep.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8splitRefE5QCharNS_13SplitBehaviorEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:516
// index:2
// Public
// QVector<QStringRef> splitRef(const class QRegExp &, enum QString::SplitBehavior)
func (this *QString) SplitRef_2(sep *QRegExp, behavior int) interface{} {
	var convArg0 = sep.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8splitRefERK7QRegExpNS_13SplitBehaviorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:520
// index:3
// Public
// QVector<QStringRef> splitRef(const class QRegularExpression &, enum QString::SplitBehavior)
func (this *QString) SplitRef_3(sep *QRegularExpression, behavior int) interface{} {
	var convArg0 = sep.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8splitRefERK18QRegularExpressionNS_13SplitBehaviorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:528
// index:0
// Public
// QString normalized(enum QString::NormalizationForm, class QChar::UnicodeVersion)
func (this *QString) Normalized(mode int, version int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10normalizedENS_17NormalizationFormEN5QChar14UnicodeVersionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode, &version)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:530
// index:0
// Public
// QString repeated(int)
func (this *QString) Repeated(times int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8repeatedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &times)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:532
// index:0
// Public
// const ushort * utf16()
func (this *QString) Utf16() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5utf16Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:535
// index:0
// Public inline
// QByteArray toLatin1()
func (this *QString) ToLatin1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR7QString8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:537
// index:1
// Public inline
// QByteArray toLatin1()
func (this *QString) ToLatin1_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO7QString8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:539
// index:0
// Public inline
// QByteArray toUtf8()
func (this *QString) ToUtf8() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR7QString6toUtf8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:541
// index:1
// Public inline
// QByteArray toUtf8()
func (this *QString) ToUtf8_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO7QString6toUtf8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:543
// index:0
// Public inline
// QByteArray toLocal8Bit()
func (this *QString) ToLocal8Bit() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR7QString11toLocal8BitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:545
// index:1
// Public inline
// QByteArray toLocal8Bit()
func (this *QString) ToLocal8Bit_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO7QString11toLocal8BitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:552
// index:0
// Public
// QVector<uint> toUcs4()
func (this *QString) ToUcs4() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString6toUcs4Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:555
// index:0
// Public static inline
// QString fromLatin1(const char *, int)
func (this *QString) FromLatin1(str string, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString10fromLatin1EPKci", ffiqt.FFI_TYPE_POINTER, str, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromLatin1(str string, size int) {
	var nilthis *QString
	nilthis.FromLatin1(str, size)
}

// /usr/include/qt/QtCore/qstring.h:568
// index:1
// Public static inline
// QString fromLatin1(const class QByteArray &)
func (this *QString) FromLatin1_1(str *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString10fromLatin1ERK10QByteArray", ffiqt.FFI_TYPE_POINTER, str)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromLatin1_1(str *QByteArray) {
	var nilthis *QString
	nilthis.FromLatin1_1(str)
}

// /usr/include/qt/QtCore/qstring.h:560
// index:0
// Public static inline
// QString fromUtf8(const char *, int)
func (this *QString) FromUtf8(str string, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString8fromUtf8EPKci", ffiqt.FFI_TYPE_POINTER, str, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromUtf8(str string, size int) {
	var nilthis *QString
	nilthis.FromUtf8(str, size)
}

// /usr/include/qt/QtCore/qstring.h:570
// index:1
// Public static inline
// QString fromUtf8(const class QByteArray &)
func (this *QString) FromUtf8_1(str *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString8fromUtf8ERK10QByteArray", ffiqt.FFI_TYPE_POINTER, str)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromUtf8_1(str *QByteArray) {
	var nilthis *QString
	nilthis.FromUtf8_1(str)
}

// /usr/include/qt/QtCore/qstring.h:564
// index:0
// Public static inline
// QString fromLocal8Bit(const char *, int)
func (this *QString) FromLocal8Bit(str string, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString13fromLocal8BitEPKci", ffiqt.FFI_TYPE_POINTER, str, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromLocal8Bit(str string, size int) {
	var nilthis *QString
	nilthis.FromLocal8Bit(str, size)
}

// /usr/include/qt/QtCore/qstring.h:572
// index:1
// Public static inline
// QString fromLocal8Bit(const class QByteArray &)
func (this *QString) FromLocal8Bit_1(str *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString13fromLocal8BitERK10QByteArray", ffiqt.FFI_TYPE_POINTER, str)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromLocal8Bit_1(str *QByteArray) {
	var nilthis *QString
	nilthis.FromLocal8Bit_1(str)
}

// /usr/include/qt/QtCore/qstring.h:574
// index:0
// Public static
// QString fromUtf16(const ushort *, int)
func (this *QString) FromUtf16(arg0 unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString9fromUtf16EPKti", ffiqt.FFI_TYPE_POINTER, arg0, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromUtf16(arg0 unsafe.Pointer, size int) {
	var nilthis *QString
	nilthis.FromUtf16(arg0, size)
}

// /usr/include/qt/QtCore/qstring.h:579
// index:1
// Public static inline
// QString fromUtf16(const char16_t *, int)
func (this *QString) FromUtf16_1(str unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString9fromUtf16EPKDsi", ffiqt.FFI_TYPE_POINTER, str, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromUtf16_1(str unsafe.Pointer, size int) {
	var nilthis *QString
	nilthis.FromUtf16_1(str, size)
}

// /usr/include/qt/QtCore/qstring.h:575
// index:0
// Public static
// QString fromUcs4(const uint *, int)
func (this *QString) FromUcs4(arg0 unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString8fromUcs4EPKji", ffiqt.FFI_TYPE_POINTER, arg0, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromUcs4(arg0 unsafe.Pointer, size int) {
	var nilthis *QString
	nilthis.FromUcs4(arg0, size)
}

// /usr/include/qt/QtCore/qstring.h:581
// index:1
// Public static inline
// QString fromUcs4(const char32_t *, int)
func (this *QString) FromUcs4_1(str unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString8fromUcs4EPKDii", ffiqt.FFI_TYPE_POINTER, str, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromUcs4_1(str unsafe.Pointer, size int) {
	var nilthis *QString
	nilthis.FromUcs4_1(str, size)
}

// /usr/include/qt/QtCore/qstring.h:576
// index:0
// Public static
// QString fromRawData(const class QChar *, int)
func (this *QString) FromRawData(arg0 unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString11fromRawDataEPK5QChari", ffiqt.FFI_TYPE_POINTER, arg0, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromRawData(arg0 unsafe.Pointer, size int) {
	var nilthis *QString
	nilthis.FromRawData(arg0, size)
}

// /usr/include/qt/QtCore/qstring.h:594
// index:0
// Public inline
// int toWCharArray(wchar_t *)
func (this *QString) ToWCharArray(array unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString12toWCharArrayEPw", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), array)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:595
// index:0
// Public static inline
// QString fromWCharArray(const wchar_t *, int)
func (this *QString) FromWCharArray(string unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString14fromWCharArrayEPKwi", ffiqt.FFI_TYPE_POINTER, string, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_FromWCharArray(string unsafe.Pointer, size int) {
	var nilthis *QString
	nilthis.FromWCharArray(string, size)
}

// /usr/include/qt/QtCore/qstring.h:597
// index:0
// Public
// QString & setRawData(const class QChar *, int)
func (this *QString) SetRawData(unicode unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString10setRawDataEPK5QChari", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), unicode, &size)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:598
// index:0
// Public
// QString & setUnicode(const class QChar *, int)
func (this *QString) SetUnicode(unicode unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString10setUnicodeEPK5QChari", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), unicode, &size)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:599
// index:0
// Public inline
// QString & setUtf16(const ushort *, int)
func (this *QString) SetUtf16(utf16 unsafe.Pointer, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString8setUtf16EPKti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), utf16, &size)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:601
// index:0
// Public
// int compare(const class QString &, Qt::CaseSensitivity)
func (this *QString) Compare(s *QString, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7compareERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:602
// index:1
// Public
// int compare(class QLatin1String, Qt::CaseSensitivity)
func (this *QString) Compare_1(other *QLatin1String, cs int) interface{} {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7compareE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:604
// index:2
// Public static inline
// int compare(const class QString &, const class QString &, Qt::CaseSensitivity)
func (this *QString) Compare_2(s1 *QString, s2 *QString, cs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7compareERKS_S1_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, s1, s2, cs)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Compare_2(s1 *QString, s2 *QString, cs int) {
	var nilthis *QString
	nilthis.Compare_2(s1, s2, cs)
}

// /usr/include/qt/QtCore/qstring.h:608
// index:3
// Public static inline
// int compare(const class QString &, class QLatin1String, Qt::CaseSensitivity)
func (this *QString) Compare_3(s1 *QString, s2 *QLatin1String, cs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, s1, s2, cs)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Compare_3(s1 *QString, s2 *QLatin1String, cs int) {
	var nilthis *QString
	nilthis.Compare_3(s1, s2, cs)
}

// /usr/include/qt/QtCore/qstring.h:611
// index:4
// Public static inline
// int compare(class QLatin1String, const class QString &, Qt::CaseSensitivity)
func (this *QString) Compare_4(s1 *QLatin1String, s2 *QString, cs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7compareE13QLatin1StringRKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, s1, s2, cs)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Compare_4(s1 *QLatin1String, s2 *QString, cs int) {
	var nilthis *QString
	nilthis.Compare_4(s1, s2, cs)
}

// /usr/include/qt/QtCore/qstring.h:615
// index:5
// Public inline
// int compare(const class QStringRef &, Qt::CaseSensitivity)
func (this *QString) Compare_5(s *QStringRef, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7compareERK10QStringRefN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:616
// index:6
// Public static
// int compare(const class QString &, const class QStringRef &, Qt::CaseSensitivity)
func (this *QString) Compare_6(s1 *QString, s2 *QStringRef, arg2 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString7compareERKS_RK10QStringRefN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, s1, s2, arg2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Compare_6(s1 *QString, s2 *QStringRef, arg2 int) {
	var nilthis *QString
	nilthis.Compare_6(s1, s2, arg2)
}

// /usr/include/qt/QtCore/qstring.h:619
// index:0
// Public
// int localeAwareCompare(const class QString &)
func (this *QString) LocaleAwareCompare(s *QString) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString18localeAwareCompareERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:620
// index:1
// Public static inline
// int localeAwareCompare(const class QString &, const class QString &)
func (this *QString) LocaleAwareCompare_1(s1 *QString, s2 *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString18localeAwareCompareERKS_S1_", ffiqt.FFI_TYPE_POINTER, s1, s2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_LocaleAwareCompare_1(s1 *QString, s2 *QString) {
	var nilthis *QString
	nilthis.LocaleAwareCompare_1(s1, s2)
}

// /usr/include/qt/QtCore/qstring.h:623
// index:2
// Public
// int localeAwareCompare(const class QStringRef &)
func (this *QString) LocaleAwareCompare_2(s *QStringRef) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString18localeAwareCompareERK10QStringRef", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:624
// index:3
// Public static
// int localeAwareCompare(const class QString &, const class QStringRef &)
func (this *QString) LocaleAwareCompare_3(s1 *QString, s2 *QStringRef) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString18localeAwareCompareERKS_RK10QStringRef", ffiqt.FFI_TYPE_POINTER, s1, s2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_LocaleAwareCompare_3(s1 *QString, s2 *QStringRef) {
	var nilthis *QString
	nilthis.LocaleAwareCompare_3(s1, s2)
}

// /usr/include/qt/QtCore/qstring.h:627
// index:0
// Public
// short toShort(_Bool *, int)
func (this *QString) ToShort(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7toShortEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:628
// index:0
// Public
// ushort toUShort(_Bool *, int)
func (this *QString) ToUShort(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8toUShortEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:629
// index:0
// Public
// int toInt(_Bool *, int)
func (this *QString) ToInt(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5toIntEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:630
// index:0
// Public
// uint toUInt(_Bool *, int)
func (this *QString) ToUInt(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString6toUIntEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:631
// index:0
// Public
// long toLong(_Bool *, int)
func (this *QString) ToLong(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString6toLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:632
// index:0
// Public
// ulong toULong(_Bool *, int)
func (this *QString) ToULong(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7toULongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:633
// index:0
// Public
// qlonglong toLongLong(_Bool *, int)
func (this *QString) ToLongLong(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10toLongLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:634
// index:0
// Public
// qulonglong toULongLong(_Bool *, int)
func (this *QString) ToULongLong(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11toULongLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:635
// index:0
// Public
// float toFloat(_Bool *)
func (this *QString) ToFloat(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString7toFloatEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:636
// index:0
// Public
// double toDouble(_Bool *)
func (this *QString) ToDouble(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8toDoubleEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:638
// index:0
// Public
// QString & setNum(short, int)
func (this *QString) SetNum(arg0 int16, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEsi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:639
// index:1
// Public
// QString & setNum(ushort, int)
func (this *QString) SetNum_1(arg0 uint16, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:640
// index:2
// Public
// QString & setNum(int, int)
func (this *QString) SetNum_2(arg0 int, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:641
// index:3
// Public
// QString & setNum(uint, int)
func (this *QString) SetNum_3(arg0 uint, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEji", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:642
// index:4
// Public
// QString & setNum(long, int)
func (this *QString) SetNum_4(arg0 int, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEli", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:643
// index:5
// Public
// QString & setNum(ulong, int)
func (this *QString) SetNum_5(arg0 uint, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEmi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:644
// index:6
// Public
// QString & setNum(qlonglong, int)
func (this *QString) SetNum_6(arg0 int64, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumExi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:645
// index:7
// Public
// QString & setNum(qulonglong, int)
func (this *QString) SetNum_7(arg0 uint64, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEyi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:646
// index:8
// Public
// QString & setNum(float, char, int)
func (this *QString) SetNum_8(arg0 float32, f byte, prec int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEfci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &f, &prec)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:647
// index:9
// Public
// QString & setNum(double, char, int)
func (this *QString) SetNum_9(arg0 float64, f byte, prec int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6setNumEdci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &f, &prec)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:649
// index:0
// Public static
// QString number(int, int)
func (this *QString) Number(arg0 int, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6numberEii", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Number(arg0 int, base int) {
	var nilthis *QString
	nilthis.Number(arg0, base)
}

// /usr/include/qt/QtCore/qstring.h:650
// index:1
// Public static
// QString number(uint, int)
func (this *QString) Number_1(arg0 uint, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6numberEji", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Number_1(arg0 uint, base int) {
	var nilthis *QString
	nilthis.Number_1(arg0, base)
}

// /usr/include/qt/QtCore/qstring.h:651
// index:2
// Public static
// QString number(long, int)
func (this *QString) Number_2(arg0 int, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6numberEli", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Number_2(arg0 int, base int) {
	var nilthis *QString
	nilthis.Number_2(arg0, base)
}

// /usr/include/qt/QtCore/qstring.h:652
// index:3
// Public static
// QString number(ulong, int)
func (this *QString) Number_3(arg0 uint, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6numberEmi", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Number_3(arg0 uint, base int) {
	var nilthis *QString
	nilthis.Number_3(arg0, base)
}

// /usr/include/qt/QtCore/qstring.h:653
// index:4
// Public static
// QString number(qlonglong, int)
func (this *QString) Number_4(arg0 int64, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6numberExi", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Number_4(arg0 int64, base int) {
	var nilthis *QString
	nilthis.Number_4(arg0, base)
}

// /usr/include/qt/QtCore/qstring.h:654
// index:5
// Public static
// QString number(qulonglong, int)
func (this *QString) Number_5(arg0 uint64, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6numberEyi", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Number_5(arg0 uint64, base int) {
	var nilthis *QString
	nilthis.Number_5(arg0, base)
}

// /usr/include/qt/QtCore/qstring.h:655
// index:6
// Public static
// QString number(double, char, int)
func (this *QString) Number_6(arg0 float64, f byte, prec int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString6numberEdci", ffiqt.FFI_TYPE_POINTER, arg0, f, prec)
	gopp.ErrPrint(err, rv)
	return rv
}
func QString_Number_6(arg0 float64, f byte, prec int) {
	var nilthis *QString
	nilthis.Number_6(arg0, f, prec)
}

// /usr/include/qt/QtCore/qstring.h:750
// index:0
// Public inline
// QString::iterator begin()
func (this *QString) Begin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:751
// index:1
// Public inline
// QString::const_iterator begin()
func (this *QString) Begin_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:752
// index:0
// Public inline
// QString::const_iterator cbegin()
func (this *QString) Cbegin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString6cbeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:753
// index:0
// Public inline
// QString::const_iterator constBegin()
func (this *QString) ConstBegin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString10constBeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:754
// index:0
// Public inline
// QString::iterator end()
func (this *QString) End() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:755
// index:1
// Public inline
// QString::const_iterator end()
func (this *QString) End_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:756
// index:0
// Public inline
// QString::const_iterator cend()
func (this *QString) Cend() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString4cendEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:757
// index:0
// Public inline
// QString::const_iterator constEnd()
func (this *QString) ConstEnd() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString8constEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:773
// index:0
// Public inline
// void push_back(class QChar)
func (this *QString) Push_back(c *QChar) {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString9push_backE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:774
// index:1
// Public inline
// void push_back(const class QString &)
func (this *QString) Push_back_1(s *QString) {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString9push_backERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:775
// index:0
// Public inline
// void push_front(class QChar)
func (this *QString) Push_front(c *QChar) {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString10push_frontE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:776
// index:1
// Public inline
// void push_front(const class QString &)
func (this *QString) Push_front_1(s *QString) {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString10push_frontERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:777
// index:0
// Public inline
// void shrink_to_fit()
func (this *QString) Shrink_to_fit() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QString13shrink_to_fitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:780
// index:0
// Public inline
// std::string toStdString()
func (this *QString) ToStdString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString11toStdStringB5cxx11Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:782
// index:0
// Public inline
// std::wstring toStdWString()
func (this *QString) ToStdWString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString12toStdWStringB5cxx11Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:786
// index:0
// Public inline
// std::u16string toStdU16String()
func (this *QString) ToStdU16String() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString14toStdU16StringB5cxx11Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:788
// index:0
// Public inline
// std::u32string toStdU32String()
func (this *QString) ToStdU32String() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString14toStdU32StringB5cxx11Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:805
// index:0
// Public inline
// bool isNull()
func (this *QString) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:808
// index:0
// Public
// bool isSimpleText()
func (this *QString) IsSimpleText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString12isSimpleTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:809
// index:0
// Public
// bool isRightToLeft()
func (this *QString) IsRightToLeft() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QString13isRightToLeftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end