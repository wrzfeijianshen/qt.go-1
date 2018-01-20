//  header block begin
// /usr/include/qt/QtGui/qicon.h
// #include <qicon.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QIcon struct {
	*qtrt.CObject
}

func (this *QIcon) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQIconFromPointer(cthis unsafe.Pointer) *QIcon {
	return &QIcon{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qicon.h:60
// index:0
// Public
// void QIcon()
func NewQIcon() *QIcon {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:61
// index:1
// Public
// void QIcon(const class QPixmap &)
func NewQIcon_1(pixmap *QPixmap) *QIcon {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2ERK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:68
// index:2
// Public
// void QIcon(const class QString &)
func NewQIcon_2(fileName *qtcore.QString) *QIcon {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:69
// index:3
// Public
// void QIcon(class QIconEngine *)
func NewQIcon_3(engine unsafe.Pointer) *QIcon {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2EP11QIconEngine", ffiqt.FFI_TYPE_VOID, cthis, engine)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:70
// index:0
// Public
// void ~QIcon()
func DeleteQIcon(*QIcon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:76
// index:0
// Public inline
// void swap(class QIcon &)
func (this *QIcon) Swap(other *QIcon) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:81
// index:0
// Public
// QPixmap pixmap(const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap(size *qtcore.QSize, mode int, state int) interface{} {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapERK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, &state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:82
// index:1
// Public inline
// QPixmap pixmap(int, int, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_1(w int, h int, mode int, state int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiiNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w, &h, &mode, &state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:84
// index:2
// Public inline
// QPixmap pixmap(int, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_2(extent int, mode int, state int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extent, &mode, &state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:86
// index:3
// Public
// QPixmap pixmap(class QWindow *, const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_3(window unsafe.Pointer, size *qtcore.QSize, mode int, state int) interface{} {
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEP7QWindowRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), window, convArg1, &mode, &state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:88
// index:0
// Public
// QSize actualSize(const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) ActualSize(size *qtcore.QSize, mode int, state int) interface{} {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10actualSizeERK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, &state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:89
// index:1
// Public
// QSize actualSize(class QWindow *, const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) ActualSize_1(window unsafe.Pointer, size *qtcore.QSize, mode int, state int) interface{} {
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10actualSizeEP7QWindowRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), window, convArg1, &mode, &state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:91
// index:0
// Public
// QString name()
func (this *QIcon) Name() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:97
// index:0
// Public
// bool isNull()
func (this *QIcon) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:98
// index:0
// Public
// bool isDetached()
func (this *QIcon) IsDetached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:99
// index:0
// Public
// void detach()
func (this *QIcon) Detach() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:104
// index:0
// Public
// qint64 cacheKey()
func (this *QIcon) CacheKey() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon8cacheKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:106
// index:0
// Public
// void addPixmap(const class QPixmap &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AddPixmap(pixmap *QPixmap, mode int, state int) {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9addPixmapERK7QPixmapNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:107
// index:0
// Public
// void addFile(const class QString &, const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AddFile(fileName *qtcore.QString, size *qtcore.QSize, mode int, state int) {
	var convArg0 = fileName.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon7addFileERK7QStringRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:109
// index:0
// Public
// QList<QSize> availableSizes(enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AvailableSizes(mode int, state int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon14availableSizesENS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode, &state)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:111
// index:0
// Public
// void setIsMask(_Bool)
func (this *QIcon) SetIsMask(isMask bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9setIsMaskEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &isMask)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:112
// index:0
// Public
// bool isMask()
func (this *QIcon) IsMask() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6isMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:114
// index:0
// Public static
// QIcon fromTheme(const class QString &)
func (this *QIcon) FromTheme(name *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QString", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
	return rv
}
func QIcon_FromTheme(name *qtcore.QString) {
	var nilthis *QIcon
	nilthis.FromTheme(name)
}

// /usr/include/qt/QtGui/qicon.h:115
// index:1
// Public static
// QIcon fromTheme(const class QString &, const class QIcon &)
func (this *QIcon) FromTheme_1(name *qtcore.QString, fallback *QIcon) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QStringRKS_", ffiqt.FFI_TYPE_POINTER, name, fallback)
	gopp.ErrPrint(err, rv)
	return rv
}
func QIcon_FromTheme_1(name *qtcore.QString, fallback *QIcon) {
	var nilthis *QIcon
	nilthis.FromTheme_1(name, fallback)
}

// /usr/include/qt/QtGui/qicon.h:116
// index:0
// Public static
// bool hasThemeIcon(const class QString &)
func (this *QIcon) HasThemeIcon(name *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon12hasThemeIconERK7QString", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
	return rv
}
func QIcon_HasThemeIcon(name *qtcore.QString) {
	var nilthis *QIcon
	nilthis.HasThemeIcon(name)
}

// /usr/include/qt/QtGui/qicon.h:118
// index:0
// Public static
// QStringList themeSearchPaths()
func (this *QIcon) ThemeSearchPaths() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon16themeSearchPathsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QIcon_ThemeSearchPaths() {
	var nilthis *QIcon
	nilthis.ThemeSearchPaths()
}

// /usr/include/qt/QtGui/qicon.h:119
// index:0
// Public static
// void setThemeSearchPaths(const class QStringList &)
func (this *QIcon) SetThemeSearchPaths(searchpath *qtcore.QStringList) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon19setThemeSearchPathsERK11QStringList", ffiqt.FFI_TYPE_POINTER, searchpath)
	gopp.ErrPrint(err, rv)
}
func QIcon_SetThemeSearchPaths(searchpath *qtcore.QStringList) {
	var nilthis *QIcon
	nilthis.SetThemeSearchPaths(searchpath)
}

// /usr/include/qt/QtGui/qicon.h:121
// index:0
// Public static
// QString themeName()
func (this *QIcon) ThemeName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9themeNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QIcon_ThemeName() {
	var nilthis *QIcon
	nilthis.ThemeName()
}

// /usr/include/qt/QtGui/qicon.h:122
// index:0
// Public static
// void setThemeName(const class QString &)
func (this *QIcon) SetThemeName(path *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon12setThemeNameERK7QString", ffiqt.FFI_TYPE_POINTER, path)
	gopp.ErrPrint(err, rv)
}
func QIcon_SetThemeName(path *qtcore.QString) {
	var nilthis *QIcon
	nilthis.SetThemeName(path)
}

//  body block end