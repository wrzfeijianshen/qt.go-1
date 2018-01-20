//  header block begin
// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h
// #include <qgraphicslinearlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QGraphicsLinearLayout struct {
	*QGraphicsLayout
}

func (this *QGraphicsLinearLayout) GetCthis() unsafe.Pointer {
	return this.QGraphicsLayout.GetCthis()
}
func NewQGraphicsLinearLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsLinearLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsLinearLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:56
// index:0
// Public
// void QGraphicsLinearLayout(class QGraphicsLayoutItem *)
func NewQGraphicsLinearLayout(parent unsafe.Pointer) *QGraphicsLinearLayout {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:57
// index:1
// Public
// void QGraphicsLinearLayout(Qt::Orientation, class QGraphicsLayoutItem *)
func NewQGraphicsLinearLayout_1(orientation int, parent unsafe.Pointer) *QGraphicsLinearLayout {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EN2Qt11OrientationEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:58
// index:0
// Public virtual
// void ~QGraphicsLinearLayout()
func DeleteQGraphicsLinearLayout(*QGraphicsLinearLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:60
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QGraphicsLinearLayout) SetOrientation(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:61
// index:0
// Public
// Qt::Orientation orientation()
func (this *QGraphicsLinearLayout) Orientation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:63
// index:0
// Public inline
// void addItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) AddItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:64
// index:0
// Public inline
// void addStretch(int)
func (this *QGraphicsLinearLayout) AddStretch(stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10addStretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:66
// index:0
// Public
// void insertItem(int, class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) InsertItem(index int, item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:67
// index:0
// Public
// void insertStretch(int, int)
func (this *QGraphicsLinearLayout) InsertStretch(index int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout13insertStretchEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:69
// index:0
// Public
// void removeItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) RemoveItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:70
// index:0
// Public virtual
// void removeAt(int)
func (this *QGraphicsLinearLayout) RemoveAt(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout8removeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:72
// index:0
// Public
// void setSpacing(qreal)
func (this *QGraphicsLinearLayout) SetSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10setSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:73
// index:0
// Public
// qreal spacing()
func (this *QGraphicsLinearLayout) Spacing() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout7spacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:74
// index:0
// Public
// void setItemSpacing(int, qreal)
func (this *QGraphicsLinearLayout) SetItemSpacing(index int, spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout14setItemSpacingEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:75
// index:0
// Public
// qreal itemSpacing(int)
func (this *QGraphicsLinearLayout) ItemSpacing(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout11itemSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:77
// index:0
// Public
// void setStretchFactor(class QGraphicsLayoutItem *, int)
func (this *QGraphicsLinearLayout) SetStretchFactor(item unsafe.Pointer, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:78
// index:0
// Public
// int stretchFactor(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) StretchFactor(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:81
// index:0
// Public
// Qt::Alignment alignment(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) Alignment(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout9alignmentEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:83
// index:0
// Public virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsLinearLayout) SetGeometry(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout11setGeometryERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:85
// index:0
// Public virtual
// int count()
func (this *QGraphicsLinearLayout) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:86
// index:0
// Public virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsLinearLayout) ItemAt(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:88
// index:0
// Public virtual
// void invalidate()
func (this *QGraphicsLinearLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:89
// index:0
// Public virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsLinearLayout) SizeHint(which int, constraint *qtcore.QSizeF) interface{} {
	var convArg1 = constraint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:95
// index:0
// Public
// void dump(int)
func (this *QGraphicsLinearLayout) Dump(indent int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout4dumpEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &indent)
	gopp.ErrPrint(err, rv)
}

//  body block end