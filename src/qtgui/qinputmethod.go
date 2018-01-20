//  header block begin
// /usr/include/qt/QtGui/qinputmethod.h
// #include <qinputmethod.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QInputMethod struct {
	*qtcore.QObject
}

func (this *QInputMethod) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQInputMethodFromPointer(cthis unsafe.Pointer) *QInputMethod {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QInputMethod{bcthis0}
}

// /usr/include/qt/QtGui/qinputmethod.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QInputMethod) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:68
// index:0
// Public
// QTransform inputItemTransform()
func (this *QInputMethod) InputItemTransform() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod18inputItemTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:69
// index:0
// Public
// void setInputItemTransform(const class QTransform &)
func (this *QInputMethod) SetInputItemTransform(transform *QTransform) {
	var convArg0 = transform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21setInputItemTransformERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:71
// index:0
// Public
// QRectF inputItemRectangle()
func (this *QInputMethod) InputItemRectangle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod18inputItemRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:72
// index:0
// Public
// void setInputItemRectangle(const class QRectF &)
func (this *QInputMethod) SetInputItemRectangle(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21setInputItemRectangleERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:75
// index:0
// Public
// QRectF cursorRectangle()
func (this *QInputMethod) CursorRectangle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod15cursorRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:76
// index:0
// Public
// QRectF anchorRectangle()
func (this *QInputMethod) AnchorRectangle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod15anchorRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:79
// index:0
// Public
// QRectF keyboardRectangle()
func (this *QInputMethod) KeyboardRectangle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod17keyboardRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:81
// index:0
// Public
// QRectF inputItemClipRectangle()
func (this *QInputMethod) InputItemClipRectangle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod22inputItemClipRectangleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:89
// index:0
// Public
// bool isVisible()
func (this *QInputMethod) IsVisible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:90
// index:0
// Public
// void setVisible(_Bool)
func (this *QInputMethod) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:92
// index:0
// Public
// bool isAnimating()
func (this *QInputMethod) IsAnimating() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod11isAnimatingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:94
// index:0
// Public
// QLocale locale()
func (this *QInputMethod) Locale() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod6localeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:95
// index:0
// Public
// Qt::LayoutDirection inputDirection()
func (this *QInputMethod) InputDirection() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QInputMethod14inputDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qinputmethod.h:97
// index:0
// Public static
// QVariant queryFocusObject(Qt::InputMethodQuery, class QVariant)
func (this *QInputMethod) QueryFocusObject(query int, argument *qtcore.QVariant) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod16queryFocusObjectEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_POINTER, query, argument)
	gopp.ErrPrint(err, rv)
	return rv
}
func QInputMethod_QueryFocusObject(query int, argument *qtcore.QVariant) {
	var nilthis *QInputMethod
	nilthis.QueryFocusObject(query, argument)
}

// /usr/include/qt/QtGui/qinputmethod.h:100
// index:0
// Public
// void show()
func (this *QInputMethod) Show() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod4showEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:101
// index:0
// Public
// void hide()
func (this *QInputMethod) Hide() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod4hideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:104
// index:0
// Public
// void reset()
func (this *QInputMethod) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:105
// index:0
// Public
// void commit()
func (this *QInputMethod) Commit() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod6commitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:107
// index:0
// Public
// void invokeAction(enum QInputMethod::Action, int)
func (this *QInputMethod) InvokeAction(a int, cursorPosition int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod12invokeActionENS_6ActionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &a, &cursorPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:110
// index:0
// Public
// void cursorRectangleChanged()
func (this *QInputMethod) CursorRectangleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod22cursorRectangleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:111
// index:0
// Public
// void anchorRectangleChanged()
func (this *QInputMethod) AnchorRectangleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod22anchorRectangleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:112
// index:0
// Public
// void keyboardRectangleChanged()
func (this *QInputMethod) KeyboardRectangleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod24keyboardRectangleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:113
// index:0
// Public
// void inputItemClipRectangleChanged()
func (this *QInputMethod) InputItemClipRectangleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod29inputItemClipRectangleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:114
// index:0
// Public
// void visibleChanged()
func (this *QInputMethod) VisibleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod14visibleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:115
// index:0
// Public
// void animatingChanged()
func (this *QInputMethod) AnimatingChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod16animatingChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:116
// index:0
// Public
// void localeChanged()
func (this *QInputMethod) LocaleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod13localeChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qinputmethod.h:117
// index:0
// Public
// void inputDirectionChanged(Qt::LayoutDirection)
func (this *QInputMethod) InputDirectionChanged(newDirection int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QInputMethod21inputDirectionChangedEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &newDirection)
	gopp.ErrPrint(err, rv)
}

//  body block end