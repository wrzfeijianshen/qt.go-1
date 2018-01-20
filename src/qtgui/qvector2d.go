//  header block begin
// /usr/include/qt/QtGui/qvector2d.h
// #include <qvector2d.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QVector2D struct {
	*qtrt.CObject
}

func (this *QVector2D) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQVector2DFromPointer(cthis unsafe.Pointer) *QVector2D {
	return &QVector2D{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qvector2d.h:59
// index:0
// Public inline
// void QVector2D()
func NewQVector2D() *QVector2D {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:60
// index:1
// Public inline
// void QVector2D(Qt::Initialization)
func NewQVector2D_1(arg0 int) *QVector2D {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:61
// index:2
// Public inline
// void QVector2D(float, float)
func NewQVector2D_2(xpos float32, ypos float32) *QVector2D {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2Eff", ffiqt.FFI_TYPE_VOID, cthis, &xpos, &ypos)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:62
// index:3
// Public inline
// void QVector2D(const class QPoint &)
func NewQVector2D_3(point *qtcore.QPoint) *QVector2D {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2ERK6QPoint", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:63
// index:4
// Public inline
// void QVector2D(const class QPointF &)
func NewQVector2D_4(point *qtcore.QPointF) *QVector2D {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2ERK7QPointF", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:65
// index:5
// Public
// void QVector2D(const class QVector3D &)
func NewQVector2D_5(vector *QVector3D) *QVector2D {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2ERK9QVector3D", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:68
// index:6
// Public
// void QVector2D(const class QVector4D &)
func NewQVector2D_6(vector *QVector4D) *QVector2D {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2DC2ERK9QVector4D", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVector2DFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qvector2d.h:71
// index:0
// Public
// bool isNull()
func (this *QVector2D) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:73
// index:0
// Public inline
// float x()
func (this *QVector2D) X() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:74
// index:0
// Public inline
// float y()
func (this *QVector2D) Y() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:76
// index:0
// Public
// void setX(float)
func (this *QVector2D) SetX(x float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2D4setXEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:77
// index:0
// Public
// void setY(float)
func (this *QVector2D) SetY(y float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2D4setYEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:82
// index:0
// Public
// float length()
func (this *QVector2D) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:83
// index:0
// Public
// float lengthSquared()
func (this *QVector2D) LengthSquared() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D13lengthSquaredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:85
// index:0
// Public
// QVector2D normalized()
func (this *QVector2D) Normalized() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D10normalizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:86
// index:0
// Public
// void normalize()
func (this *QVector2D) Normalize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2D9normalizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qvector2d.h:88
// index:0
// Public
// float distanceToPoint(const class QVector2D &)
func (this *QVector2D) DistanceToPoint(point *QVector2D) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D15distanceToPointERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:89
// index:0
// Public
// float distanceToLine(const class QVector2D &, const class QVector2D &)
func (this *QVector2D) DistanceToLine(point *QVector2D, direction *QVector2D) interface{} {
	var convArg0 = point.GetCthis()
	var convArg1 = direction.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D14distanceToLineERKS_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:98
// index:0
// Public static
// float dotProduct(const class QVector2D &, const class QVector2D &)
func (this *QVector2D) DotProduct(v1 *QVector2D, v2 *QVector2D) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QVector2D10dotProductERKS_S1_", ffiqt.FFI_TYPE_POINTER, v1, v2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QVector2D_DotProduct(v1 *QVector2D, v2 *QVector2D) {
	var nilthis *QVector2D
	nilthis.DotProduct(v1, v2)
}

// /usr/include/qt/QtGui/qvector2d.h:114
// index:0
// Public
// QVector3D toVector3D()
func (this *QVector2D) ToVector3D() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D10toVector3DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:117
// index:0
// Public
// QVector4D toVector4D()
func (this *QVector2D) ToVector4D() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D10toVector4DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:120
// index:0
// Public inline
// QPoint toPoint()
func (this *QVector2D) ToPoint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D7toPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qvector2d.h:121
// index:0
// Public inline
// QPointF toPointF()
func (this *QVector2D) ToPointF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QVector2D8toPointFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end