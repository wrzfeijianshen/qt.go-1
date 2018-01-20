//  header block begin
// /usr/include/qt/QtWidgets/qrubberband.h
// #include <qrubberband.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 65
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
type QRubberBand struct {
	*QWidget
}

func (this *QRubberBand) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQRubberBandFromPointer(cthis unsafe.Pointer) *QRubberBand {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QRubberBand{bcthis0}
}

// /usr/include/qt/QtWidgets/qrubberband.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QRubberBand) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QRubberBand10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qrubberband.h:59
// index:0
// Public
// void QRubberBand(enum QRubberBand::Shape, class QWidget *)
func NewQRubberBand(arg0 int, arg1 unsafe.Pointer) *QRubberBand {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBandC2ENS_5ShapeEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &arg0, arg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQRubberBandFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qrubberband.h:60
// index:0
// Public virtual
// void ~QRubberBand()
func DeleteQRubberBand(*QRubberBand) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBandD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:62
// index:0
// Public
// QRubberBand::Shape shape()
func (this *QRubberBand) Shape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QRubberBand5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qrubberband.h:64
// index:0
// Public
// void setGeometry(const class QRect &)
func (this *QRubberBand) SetGeometry(r *qtcore.QRect) {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:66
// index:1
// Public inline
// void setGeometry(int, int, int, int)
func (this *QRubberBand) SetGeometry_1(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand11setGeometryEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:67
// index:0
// Public inline
// void move(int, int)
func (this *QRubberBand) Move(x int, y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand4moveEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:68
// index:1
// Public inline
// void move(const class QPoint &)
func (this *QRubberBand) Move_1(p *qtcore.QPoint) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand4moveERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:70
// index:0
// Public inline
// void resize(int, int)
func (this *QRubberBand) Resize(w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand6resizeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:72
// index:1
// Public inline
// void resize(const class QSize &)
func (this *QRubberBand) Resize_1(s *qtcore.QSize) {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand6resizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:76
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QRubberBand) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qrubberband.h:77
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QRubberBand) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:78
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QRubberBand) ChangeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:79
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QRubberBand) ShowEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:80
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QRubberBand) ResizeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:81
// index:0
// Protected virtual
// void moveEvent(class QMoveEvent *)
func (this *QRubberBand) MoveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QRubberBand9moveEventEP10QMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qrubberband.h:82
// index:0
// Protected
// void initStyleOption(class QStyleOptionRubberBand *)
func (this *QRubberBand) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QRubberBand15initStyleOptionEP22QStyleOptionRubberBand", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end