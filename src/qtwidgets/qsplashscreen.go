//  header block begin
// /usr/include/qt/QtWidgets/qsplashscreen.h
// #include <qsplashscreen.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSplashScreen struct {
	*QWidget
}

func (this *QSplashScreen) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQSplashScreenFromPointer(cthis unsafe.Pointer) *QSplashScreen {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSplashScreen{bcthis0}
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSplashScreen) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSplashScreen10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:59
// index:0
// Public virtual
// void ~QSplashScreen()
func DeleteQSplashScreen(*QSplashScreen) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreenD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:61
// index:0
// Public
// void setPixmap(const class QPixmap &)
func (this *QSplashScreen) SetPixmap(pixmap *qtgui.QPixmap) {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen9setPixmapERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:62
// index:0
// Public
// const QPixmap pixmap()
func (this *QSplashScreen) Pixmap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSplashScreen6pixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:63
// index:0
// Public
// void finish(class QWidget *)
func (this *QSplashScreen) Finish(w unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen6finishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:64
// index:0
// Public
// void repaint()
func (this *QSplashScreen) Repaint() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen7repaintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:65
// index:0
// Public
// QString message()
func (this *QSplashScreen) Message() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSplashScreen7messageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:68
// index:0
// Public
// void showMessage(const class QString &, int, const class QColor &)
func (this *QSplashScreen) ShowMessage(message *qtcore.QString, alignment int, color *qtgui.QColor) {
	var convArg0 = message.GetCthis()
	var convArg2 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &alignment, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:70
// index:0
// Public
// void clearMessage()
func (this *QSplashScreen) ClearMessage() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen12clearMessageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:73
// index:0
// Public
// void messageChanged(const class QString &)
func (this *QSplashScreen) MessageChanged(message *qtcore.QString) {
	var convArg0 = message.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen14messageChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:76
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSplashScreen) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:77
// index:0
// Protected virtual
// void drawContents(class QPainter *)
func (this *QSplashScreen) DrawContents(painter unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen12drawContentsEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplashscreen.h:78
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QSplashScreen) MousePressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSplashScreen15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end