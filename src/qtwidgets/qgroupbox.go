//  header block begin
// /usr/include/qt/QtWidgets/qgroupbox.h
// #include <qgroupbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 106
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
type QGroupBox struct {
	*QWidget
}

func (this *QGroupBox) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQGroupBoxFromPointer(cthis unsafe.Pointer) *QGroupBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QGroupBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qgroupbox.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGroupBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// Public
// void QGroupBox(class QWidget *)
func NewQGroupBox(parent unsafe.Pointer) *QGroupBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGroupBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// Public
// void QGroupBox(const class QString &, class QWidget *)
func NewQGroupBox_1(title *qtcore.QString, parent unsafe.Pointer) *QGroupBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGroupBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:64
// index:0
// Public virtual
// void ~QGroupBox()
func DeleteQGroupBox(*QGroupBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:66
// index:0
// Public
// QString title()
func (this *QGroupBox) Title() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox5titleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgroupbox.h:67
// index:0
// Public
// void setTitle(const class QString &)
func (this *QGroupBox) SetTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox8setTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:69
// index:0
// Public
// Qt::Alignment alignment()
func (this *QGroupBox) Alignment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgroupbox.h:70
// index:0
// Public
// void setAlignment(int)
func (this *QGroupBox) SetAlignment(alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox12setAlignmentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:72
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QGroupBox) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgroupbox.h:74
// index:0
// Public
// bool isFlat()
func (this *QGroupBox) IsFlat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox6isFlatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgroupbox.h:75
// index:0
// Public
// void setFlat(_Bool)
func (this *QGroupBox) SetFlat(flat bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7setFlatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &flat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:76
// index:0
// Public
// bool isCheckable()
func (this *QGroupBox) IsCheckable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox11isCheckableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgroupbox.h:77
// index:0
// Public
// void setCheckable(_Bool)
func (this *QGroupBox) SetCheckable(checkable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox12setCheckableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &checkable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:78
// index:0
// Public
// bool isChecked()
func (this *QGroupBox) IsChecked() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox9isCheckedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgroupbox.h:81
// index:0
// Public
// void setChecked(_Bool)
func (this *QGroupBox) SetChecked(checked bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox10setCheckedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:84
// index:0
// Public
// void clicked(_Bool)
func (this *QGroupBox) Clicked(checked bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7clickedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:85
// index:0
// Public
// void toggled(_Bool)
func (this *QGroupBox) Toggled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7toggledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:88
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QGroupBox) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgroupbox.h:89
// index:0
// Protected virtual
// void childEvent(class QChildEvent *)
func (this *QGroupBox) ChildEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox10childEventEP11QChildEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:90
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QGroupBox) ResizeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:91
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QGroupBox) PaintEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:92
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGroupBox) FocusInEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:93
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QGroupBox) ChangeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:94
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QGroupBox) MousePressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:95
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QGroupBox) MouseMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:96
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QGroupBox) MouseReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:97
// index:0
// Protected
// void initStyleOption(class QStyleOptionGroupBox *)
func (this *QGroupBox) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox15initStyleOptionEP20QStyleOptionGroupBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end