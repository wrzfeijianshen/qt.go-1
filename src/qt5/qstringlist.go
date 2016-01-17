package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qstringlist.h
// dst-file: /src/core/qstringlist.go
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
  // proto:  int QStringList::indexOf(const QRegExp & rx, int from);
extern void _ZNK11QStringList7indexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QStringList::indexOf(const QRegularExpression & re, int from);
extern void _ZNK11QStringList7indexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QStringList::indexOf(QRegExp & rx, int from);
extern void _ZNK11QStringList7indexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  void QStringList::QStringList();
extern void _ZN11QStringListC2Ev(void* qthis); // 1
  // proto:  void QStringList::QStringList(const QString & i);
extern void _ZN11QStringListC2ERK7QString(void* qthis, void* arg0); // 1
  // proto:  int QStringList::lastIndexOf(const QRegExp & rx, int from);
extern void _ZNK11QStringList11lastIndexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QStringList::lastIndexOf(const QRegularExpression & re, int from);
extern void _ZNK11QStringList11lastIndexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QStringList::lastIndexOf(QRegExp & rx, int from);
extern void _ZNK11QStringList11lastIndexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  QStringList & QListSpecialMethods<QString>::replaceInStrings(const QRegularExpression & re, const QString & after);
extern void _ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK18QRegularExpressionRKS0_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QStringList & QListSpecialMethods<QString>::replaceInStrings(const QRegExp & rx, const QString & after);
extern void _ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK7QRegExpRKS0_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QString QListSpecialMethods<QString>::join(const QString & sep);
extern void _ZNK19QListSpecialMethodsI7QStringE4joinERKS0_(void* qthis, void* arg0); // 2
  // proto:  QString QListSpecialMethods<QString>::join(QChar sep);
extern void _ZNK19QListSpecialMethodsI7QStringE4joinE5QChar(void* qthis, void* arg0); // 2
  // proto:  QStringList QListSpecialMethods<QString>::filter(const QRegularExpression & re);
extern void _ZNK19QListSpecialMethodsI7QStringE6filterERK18QRegularExpression(void* qthis, void* arg0); // 2
  // proto:  QStringList QListSpecialMethods<QString>::filter(const QRegExp & rx);
extern void _ZNK19QListSpecialMethodsI7QStringE6filterERK7QRegExp(void* qthis, void* arg0); // 2
  // proto:  int QListSpecialMethods<QString>::removeDuplicates();
extern void _ZN19QListSpecialMethodsI7QStringE16removeDuplicatesEv(void* qthis); // 2
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

// class sizeof(QStringList)=1
type QStringList struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QListSpecialMethods<QString>)=1
type QListSpecialMethods_QString_ struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// indexOf(const class QRegExp &, int)
func (this *QStringList) indexOf(args ...interface{}) () {
  // indexOf(const class QRegExp &, int)
  // indexOf(const class QRegularExpression &, int)
  // indexOf(class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStringList7indexOfERK7QRegExpi
    // invoke: int indexOf(const class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK11QStringList7indexOfERK7QRegExpi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK11QStringList7indexOfERK18QRegularExpressioni
    // invoke: int indexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK11QStringList7indexOfERK18QRegularExpressioni(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK11QStringList7indexOfER7QRegExpi
    // invoke: int indexOf(class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK11QStringList7indexOfER7QRegExpi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringList", "indexOf", args)
  }

}

// QStringList()
func NewQStringList(args ...interface{}) QStringList {
  // QStringList()
  // QStringList(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStringListC1Ev
    // invoke: void QStringList()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QStringListC2Ev(qthis)
  case 1:
    // invoke: _ZN11QStringListC1ERK7QString
    // invoke: void QStringList(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QStringListC2ERK7QString(qthis, arg0)
  default:
    qtrt.ErrorResolve("QStringList", "QStringList", args)
  }

  return QStringList{}
}

// lastIndexOf(const class QRegExp &, int)
func (this *QStringList) lastIndexOf(args ...interface{}) () {
  // lastIndexOf(const class QRegExp &, int)
  // lastIndexOf(const class QRegularExpression &, int)
  // lastIndexOf(class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStringList11lastIndexOfERK7QRegExpi
    // invoke: int lastIndexOf(const class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK11QStringList11lastIndexOfERK7QRegExpi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK11QStringList11lastIndexOfERK18QRegularExpressioni
    // invoke: int lastIndexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK11QStringList11lastIndexOfERK18QRegularExpressioni(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK11QStringList11lastIndexOfER7QRegExpi
    // invoke: int lastIndexOf(class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK11QStringList11lastIndexOfER7QRegExpi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringList", "lastIndexOf", args)
  }

}

// replaceInStrings(const class QRegularExpression &, const class QString &)
func (this *QListSpecialMethods_QString_) replaceInStrings(args ...interface{}) () {
  // replaceInStrings(const class QRegularExpression &, const class QString &)
  // replaceInStrings(const class QRegExp &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK18QRegularExpressionRKS0_
    // invoke: QStringList & replaceInStrings(const class QRegularExpression &, const class QString &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK18QRegularExpressionRKS0_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK7QRegExpRKS0_
    // invoke: QStringList & replaceInStrings(const class QRegExp &, const class QString &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK7QRegExpRKS0_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QString>", "replaceInStrings", args)
  }

}

// join(const class QString &)
func (this *QListSpecialMethods_QString_) join(args ...interface{}) () {
  // join(const class QString &)
  // join(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QListSpecialMethodsI7QStringE4joinERKS0_
    // invoke: QString join(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QListSpecialMethodsI7QStringE4joinERKS0_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK19QListSpecialMethodsI7QStringE4joinE5QChar
    // invoke: QString join(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QListSpecialMethodsI7QStringE4joinE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QString>", "join", args)
  }

}

// filter(const class QRegularExpression &)
func (this *QListSpecialMethods_QString_) filter(args ...interface{}) () {
  // filter(const class QRegularExpression &)
  // filter(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QListSpecialMethodsI7QStringE6filterERK18QRegularExpression
    // invoke: QStringList filter(const class QRegularExpression &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QListSpecialMethodsI7QStringE6filterERK18QRegularExpression(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK19QListSpecialMethodsI7QStringE6filterERK7QRegExp
    // invoke: QStringList filter(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QListSpecialMethodsI7QStringE6filterERK7QRegExp(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QString>", "filter", args)
  }

}

// removeDuplicates()
func (this *QListSpecialMethods_QString_) removeDuplicates(args ...interface{}) () {
  // removeDuplicates()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QListSpecialMethodsI7QStringE16removeDuplicatesEv
    // invoke: int removeDuplicates()
    C._ZN19QListSpecialMethodsI7QStringE16removeDuplicatesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QString>", "removeDuplicates", args)
  }

}

// <= body block end

