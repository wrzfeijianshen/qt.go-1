//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QTextFrameFormat struct {
	*QTextFormat
}

func (this *QTextFrameFormat) GetCthis() unsafe.Pointer {
	return this.QTextFormat.GetCthis()
}
func NewQTextFrameFormatFromPointer(cthis unsafe.Pointer) *QTextFrameFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextFrameFormat{bcthis0}
}

// /usr/include/qt/QtGui/qtextformat.h:770
// index:0
// Public
// void QTextFrameFormat()
func NewQTextFrameFormat() *QTextFrameFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFrameFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:852
// index:1
// Protected
// void QTextFrameFormat(const class QTextFormat &)
func NewQTextFrameFormat_1(fmt *QTextFormat) *QTextFrameFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = fmt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFrameFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:772
// index:0
// Public inline
// bool isValid()
func (this *QTextFrameFormat) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:796
// index:0
// Public inline
// void setPosition(enum QTextFrameFormat::Position)
func (this *QTextFrameFormat) SetPosition(f int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat11setPositionENS_8PositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:798
// index:0
// Public inline
// QTextFrameFormat::Position position()
func (this *QTextFrameFormat) Position() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:801
// index:0
// Public inline
// void setBorder(qreal)
func (this *QTextFrameFormat) SetBorder(border float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat9setBorderEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &border)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:802
// index:0
// Public inline
// qreal border()
func (this *QTextFrameFormat) Border() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat6borderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:805
// index:0
// Public inline
// void setBorderBrush(const class QBrush &)
func (this *QTextFrameFormat) SetBorderBrush(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat14setBorderBrushERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:807
// index:0
// Public inline
// QBrush borderBrush()
func (this *QTextFrameFormat) BorderBrush() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat11borderBrushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:810
// index:0
// Public inline
// void setBorderStyle(enum QTextFrameFormat::BorderStyle)
func (this *QTextFrameFormat) SetBorderStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat14setBorderStyleENS_11BorderStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:812
// index:0
// Public inline
// QTextFrameFormat::BorderStyle borderStyle()
func (this *QTextFrameFormat) BorderStyle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat11borderStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:815
// index:0
// Public
// void setMargin(qreal)
func (this *QTextFrameFormat) SetMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat9setMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:816
// index:0
// Public inline
// qreal margin()
func (this *QTextFrameFormat) Margin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat6marginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:819
// index:0
// Public inline
// void setTopMargin(qreal)
func (this *QTextFrameFormat) SetTopMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat12setTopMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:820
// index:0
// Public
// qreal topMargin()
func (this *QTextFrameFormat) TopMargin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat9topMarginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:822
// index:0
// Public inline
// void setBottomMargin(qreal)
func (this *QTextFrameFormat) SetBottomMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat15setBottomMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:823
// index:0
// Public
// qreal bottomMargin()
func (this *QTextFrameFormat) BottomMargin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat12bottomMarginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:825
// index:0
// Public inline
// void setLeftMargin(qreal)
func (this *QTextFrameFormat) SetLeftMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat13setLeftMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:826
// index:0
// Public
// qreal leftMargin()
func (this *QTextFrameFormat) LeftMargin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat10leftMarginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:828
// index:0
// Public inline
// void setRightMargin(qreal)
func (this *QTextFrameFormat) SetRightMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat14setRightMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:829
// index:0
// Public
// qreal rightMargin()
func (this *QTextFrameFormat) RightMargin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat11rightMarginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:831
// index:0
// Public inline
// void setPadding(qreal)
func (this *QTextFrameFormat) SetPadding(padding float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat10setPaddingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:832
// index:0
// Public inline
// qreal padding()
func (this *QTextFrameFormat) Padding() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat7paddingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:835
// index:0
// Public inline
// void setWidth(qreal)
func (this *QTextFrameFormat) SetWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat8setWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:836
// index:1
// Public inline
// void setWidth(const class QTextLength &)
func (this *QTextFrameFormat) SetWidth_1(length *QTextLength) {
	var convArg0 = length.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat8setWidthERK11QTextLength", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:838
// index:0
// Public inline
// QTextLength width()
func (this *QTextFrameFormat) Width() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:841
// index:0
// Public inline
// void setHeight(qreal)
func (this *QTextFrameFormat) SetHeight(height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat9setHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:842
// index:1
// Public inline
// void setHeight(const class QTextLength &)
func (this *QTextFrameFormat) SetHeight_1(height *QTextLength) {
	var convArg0 = height.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat9setHeightERK11QTextLength", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:843
// index:0
// Public inline
// QTextLength height()
func (this *QTextFrameFormat) Height() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextformat.h:848
// index:0
// Public inline
// QTextFormat::PageBreakFlags pageBreakPolicy()
func (this *QTextFrameFormat) PageBreakPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat15pageBreakPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end