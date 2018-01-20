//  header block begin
// /usr/include/qt/QtWidgets/qproxystyle.h
// #include <qproxystyle.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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
type QProxyStyle struct {
	*QCommonStyle
}

func (this *QProxyStyle) GetCthis() unsafe.Pointer {
	return this.QCommonStyle.GetCthis()
}
func NewQProxyStyleFromPointer(cthis unsafe.Pointer) *QProxyStyle {
	bcthis0 := NewQCommonStyleFromPointer(cthis)
	return &QProxyStyle{bcthis0}
}

// /usr/include/qt/QtWidgets/qproxystyle.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QProxyStyle) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:57
// index:0
// Public
// void QProxyStyle(class QStyle *)
func NewQProxyStyle(style unsafe.Pointer) *QProxyStyle {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyleC2EP6QStyle", ffiqt.FFI_TYPE_VOID, cthis, style)
	gopp.ErrPrint(err, rv)
	gothis := NewQProxyStyleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qproxystyle.h:58
// index:1
// Public
// void QProxyStyle(const class QString &)
func NewQProxyStyle_1(key *qtcore.QString) *QProxyStyle {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyleC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQProxyStyleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qproxystyle.h:59
// index:0
// Public virtual
// void ~QProxyStyle()
func DeleteQProxyStyle(*QProxyStyle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:61
// index:0
// Public
// QStyle * baseStyle()
func (this *QProxyStyle) BaseStyle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle9baseStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:62
// index:0
// Public
// void setBaseStyle(class QStyle *)
func (this *QProxyStyle) SetBaseStyle(style unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle12setBaseStyleEP6QStyle", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:64
// index:0
// Public virtual
// void drawPrimitive(enum QStyle::PrimitiveElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QProxyStyle) DrawPrimitive(element int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &element, option, painter, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:65
// index:0
// Public virtual
// void drawControl(enum QStyle::ControlElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QProxyStyle) DrawControl(element int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &element, option, painter, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:66
// index:0
// Public virtual
// void drawComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, class QPainter *, const class QWidget *)
func (this *QProxyStyle) DrawComplexControl(control int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &control, option, painter, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:67
// index:0
// Public virtual
// void drawItemText(class QPainter *, const class QRect &, int, const class QPalette &, _Bool, const class QString &, class QPalette::ColorRole)
func (this *QProxyStyle) DrawItemText(painter unsafe.Pointer, rect *qtcore.QRect, flags int, pal *qtgui.QPalette, enabled bool, text *qtcore.QString, textRole int) {
	var convArg1 = rect.GetCthis()
	var convArg3 = pal.GetCthis()
	var convArg5 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle12drawItemTextEP8QPainterRK5QRectiRK8QPalettebRK7QStringNS5_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, &flags, convArg3, &enabled, convArg5, &textRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:69
// index:0
// Public virtual
// void drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
func (this *QProxyStyle) DrawItemPixmap(painter unsafe.Pointer, rect *qtcore.QRect, alignment int, pixmap *qtgui.QPixmap) {
	var convArg1 = rect.GetCthis()
	var convArg3 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, &alignment, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:71
// index:0
// Public virtual
// QSize sizeFromContents(enum QStyle::ContentsType, const class QStyleOption *, const class QSize &, const class QWidget *)
func (this *QProxyStyle) SizeFromContents(type_ int, option unsafe.Pointer, size *qtcore.QSize, widget unsafe.Pointer) interface{} {
	var convArg2 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_, option, convArg2, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:73
// index:0
// Public virtual
// QRect subElementRect(enum QStyle::SubElement, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) SubElementRect(element int, option unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &element, option, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:74
// index:0
// Public virtual
// QRect subControlRect(enum QStyle::ComplexControl, const class QStyleOptionComplex *, enum QStyle::SubControl, const class QWidget *)
func (this *QProxyStyle) SubControlRect(cc int, opt unsafe.Pointer, sc int, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cc, opt, &sc, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:75
// index:0
// Public virtual
// QRect itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
func (this *QProxyStyle) ItemTextRect(fm *qtgui.QFontMetrics, r *qtcore.QRect, flags int, enabled bool, text *qtcore.QString) interface{} {
	var convArg0 = fm.GetCthis()
	var convArg1 = r.GetCthis()
	var convArg4 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &flags, &enabled, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:76
// index:0
// Public virtual
// QRect itemPixmapRect(const class QRect &, int, const class QPixmap &)
func (this *QProxyStyle) ItemPixmapRect(r *qtcore.QRect, flags int, pixmap *qtgui.QPixmap) interface{} {
	var convArg0 = r.GetCthis()
	var convArg2 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &flags, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:78
// index:0
// Public virtual
// QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, const class QPoint &, const class QWidget *)
func (this *QProxyStyle) HitTestComplexControl(control int, option unsafe.Pointer, pos *qtcore.QPoint, widget unsafe.Pointer) interface{} {
	var convArg2 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &control, option, convArg2, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:79
// index:0
// Public virtual
// int styleHint(enum QStyle::StyleHint, const class QStyleOption *, const class QWidget *, class QStyleHintReturn *)
func (this *QProxyStyle) StyleHint(hint int, option unsafe.Pointer, widget unsafe.Pointer, returnData unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &hint, option, widget, returnData)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:80
// index:0
// Public virtual
// int pixelMetric(enum QStyle::PixelMetric, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) PixelMetric(metric int, option unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &metric, option, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:81
// index:0
// Public virtual
// int layoutSpacing(class QSizePolicy::ControlType, class QSizePolicy::ControlType, Qt::Orientation, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) LayoutSpacing(control1 int, control2 int, orientation int, option unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &control1, &control2, &orientation, option, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:84
// index:0
// Public virtual
// QIcon standardIcon(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) StandardIcon(standardIcon int, option unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &standardIcon, option, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:85
// index:0
// Public virtual
// QPixmap standardPixmap(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) StandardPixmap(standardPixmap int, opt unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &standardPixmap, opt, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:86
// index:0
// Public virtual
// QPixmap generatedIconPixmap(class QIcon::Mode, const class QPixmap &, const class QStyleOption *)
func (this *QProxyStyle) GeneratedIconPixmap(iconMode int, pixmap *qtgui.QPixmap, opt unsafe.Pointer) interface{} {
	var convArg1 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &iconMode, convArg1, opt)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:87
// index:0
// Public virtual
// QPalette standardPalette()
func (this *QProxyStyle) StandardPalette() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle15standardPaletteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qproxystyle.h:89
// index:0
// Public virtual
// void polish(class QWidget *)
func (this *QProxyStyle) Polish(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle6polishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:90
// index:1
// Public virtual
// void polish(class QPalette &)
func (this *QProxyStyle) Polish_1(pal *qtgui.QPalette) {
	var convArg0 = pal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle6polishER8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:91
// index:2
// Public virtual
// void polish(class QApplication *)
func (this *QProxyStyle) Polish_2(app unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle6polishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), app)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:93
// index:0
// Public virtual
// void unpolish(class QWidget *)
func (this *QProxyStyle) Unpolish(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:94
// index:1
// Public virtual
// void unpolish(class QApplication *)
func (this *QProxyStyle) Unpolish_1(app unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), app)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:97
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QProxyStyle) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end