package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qtooltip.h
// dst-file: /src/widgets/qtooltip.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QPalette QToolTip::palette();
extern void C_ZN8QToolTip7paletteEv(); // 4
  // proto: static void QToolTip::hideText();
extern void C_ZN8QToolTip8hideTextEv(); // 2
  // proto: static QString QToolTip::text();
extern void C_ZN8QToolTip4textEv(); // 4
  // proto: static void QToolTip::setPalette(const QPalette & );
extern void C_ZN8QToolTip10setPaletteERK8QPalette(void* arg0); // 4
  // proto: static void QToolTip::showText(const QPoint & pos, const QString & text, QWidget * w, const QRect & rect);
extern void C_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRect(void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto: static void QToolTip::showText(const QPoint & pos, const QString & text, QWidget * w, const QRect & rect, int msecShowTime);
extern void C_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRecti(void* arg0, void* arg1, void* arg2, void* arg3, int32_t arg4); // 4
  // proto: static void QToolTip::showText(const QPoint & pos, const QString & text, QWidget * w);
extern void C_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidget(void* arg0, void* arg1, void* arg2); // 4
  // proto: static bool QToolTip::isVisible();
extern void C_ZN8QToolTip9isVisibleEv(); // 4
  // proto: static void QToolTip::setFont(const QFont & );
extern void C_ZN8QToolTip7setFontERK5QFont(void* arg0); // 4
  // proto: static QFont QToolTip::font();
extern void C_ZN8QToolTip4fontEv(); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QToolTip)=1
type QToolTip struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// palette()
func (this *QToolTip) palette_s(args ...interface{}) () {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolTip7paletteEv
    // invoke: QPalette palette()
    C.C_ZN8QToolTip7paletteEv()
  default:
    qtrt.ErrorResolve("QToolTip", "palette", args)
  }

}

// hideText()
func (this *QToolTip) hideText_s(args ...interface{}) () {
  // hideText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolTip8hideTextEv
    // invoke: void hideText()
    C.C_ZN8QToolTip8hideTextEv()
  default:
    qtrt.ErrorResolve("QToolTip", "hideText", args)
  }

}

// text()
func (this *QToolTip) text_s(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolTip4textEv
    // invoke: QString text()
    C.C_ZN8QToolTip4textEv()
  default:
    qtrt.ErrorResolve("QToolTip", "text", args)
  }

}

// setPalette(const class QPalette &)
func (this *QToolTip) setPalette_s(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolTip10setPaletteERK8QPalette
    // invoke: void setPalette(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QToolTip10setPaletteERK8QPalette(arg0)
  default:
    qtrt.ErrorResolve("QToolTip", "setPalette", args)
  }

}

// showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &)
func (this *QToolTip) showText_s(args ...interface{}) () {
  // showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &)
  // showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &, int)
  // showText(const class QPoint &, const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][3] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][3] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRect
    // invoke: void showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QRect).qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRect(arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRecti
    // invoke: void showText(const class QPoint &, const class QString &, class QWidget *, const class QRect &, int)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QRect).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRecti(arg0, arg1, arg2, arg3, arg4)
  case 2:
    // invoke: _ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidget
    // invoke: void showText(const class QPoint &, const class QString &, class QWidget *)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidget(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QToolTip", "showText", args)
  }

}

// isVisible()
func (this *QToolTip) isVisible_s(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolTip9isVisibleEv
    // invoke: bool isVisible()
    C.C_ZN8QToolTip9isVisibleEv()
  default:
    qtrt.ErrorResolve("QToolTip", "isVisible", args)
  }

}

// setFont(const class QFont &)
func (this *QToolTip) setFont_s(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolTip7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QToolTip7setFontERK5QFont(arg0)
  default:
    qtrt.ErrorResolve("QToolTip", "setFont", args)
  }

}

// font()
func (this *QToolTip) font_s(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolTip4fontEv
    // invoke: QFont font()
    C.C_ZN8QToolTip4fontEv()
  default:
    qtrt.ErrorResolve("QToolTip", "font", args)
  }

}

// <= body block end

