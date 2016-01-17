package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qabstracteventdispatcher.h
// dst-file: /src/core/qabstracteventdispatcher.go
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
  // proto:  void QAbstractEventDispatcher::installNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void _ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0); // 4
  // proto:  void QAbstractEventDispatcher::~QAbstractEventDispatcher();
extern void _ZN24QAbstractEventDispatcherD2Ev(void* qthis); // 4
  // proto:  void QAbstractEventDispatcher::closingDown();
extern void _ZN24QAbstractEventDispatcher11closingDownEv(void* qthis); // 4
  // proto:  void QAbstractEventDispatcher::startingUp();
extern void _ZN24QAbstractEventDispatcher10startingUpEv(void* qthis); // 4
  // proto: static QAbstractEventDispatcher * QAbstractEventDispatcher::instance(QThread * thread);
extern void _ZN24QAbstractEventDispatcher8instanceEP7QThread(void* arg0); // 4
  // proto:  void QAbstractEventDispatcher::removeNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void _ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractEventDispatcher::filterNativeEvent(const QByteArray & eventType, void * message, long * result);
extern void _ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl(void* qthis, void* arg0, void* arg1, int32_t* arg2); // 4
  // proto:  void QAbstractEventDispatcher::QAbstractEventDispatcher(QObject * parent);
extern void _ZN24QAbstractEventDispatcherC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  const QMetaObject * QAbstractEventDispatcher::metaObject();
extern void _ZNK24QAbstractEventDispatcher10metaObjectEv(void* qthis); // 4
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

// class sizeof(QAbstractEventDispatcher)=1
type QAbstractEventDispatcher struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToBlock QAbstractEventDispatcher_aboutToBlock_signal;
//  _awake QAbstractEventDispatcher_awake_signal;
}

// installNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) installNativeEventFilter(args ...interface{}) () {
  // installNativeEventFilter(class QAbstractNativeEventFilter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractNativeEventFilter{}) // "QAbstractNativeEventFilter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter
    // invoke: void installNativeEventFilter(class QAbstractNativeEventFilter *)
    var arg0 = args[0].(QAbstractNativeEventFilter).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "installNativeEventFilter", args)
  }

}

// ~QAbstractEventDispatcher()
func (this *QAbstractEventDispatcher) FreeQAbstractEventDispatcher(args ...interface{}) () {
  // ~QAbstractEventDispatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcherD0Ev
    // invoke: void ~QAbstractEventDispatcher()
    C._ZN24QAbstractEventDispatcherD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "~QAbstractEventDispatcher", args)
  }

}

// closingDown()
func (this *QAbstractEventDispatcher) closingDown(args ...interface{}) () {
  // closingDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher11closingDownEv
    // invoke: void closingDown()
    C._ZN24QAbstractEventDispatcher11closingDownEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "closingDown", args)
  }

}

// startingUp()
func (this *QAbstractEventDispatcher) startingUp(args ...interface{}) () {
  // startingUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher10startingUpEv
    // invoke: void startingUp()
    C._ZN24QAbstractEventDispatcher10startingUpEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "startingUp", args)
  }

}

// instance(class QThread *)
func (this *QAbstractEventDispatcher) instance_s(args ...interface{}) () {
  // instance(class QThread *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QThread{}) // "QThread *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher8instanceEP7QThread
    // invoke: QAbstractEventDispatcher * instance(class QThread *)
    var arg0 = args[0].(QThread).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher8instanceEP7QThread(arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "instance", args)
  }

}

// removeNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) removeNativeEventFilter(args ...interface{}) () {
  // removeNativeEventFilter(class QAbstractNativeEventFilter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractNativeEventFilter{}) // "QAbstractNativeEventFilter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter
    // invoke: void removeNativeEventFilter(class QAbstractNativeEventFilter *)
    var arg0 = args[0].(QAbstractNativeEventFilter).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "removeNativeEventFilter", args)
  }

}

// filterNativeEvent(const class QByteArray &, void *, long *)
func (this *QAbstractEventDispatcher) filterNativeEvent(args ...interface{}) () {
  // filterNativeEvent(const class QByteArray &, void *, long *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  vtys[0][2] = qtrt.Int32Ty(true) // "long *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl
    // invoke: bool filterNativeEvent(const class QByteArray &, void *, long *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C._ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "filterNativeEvent", args)
  }

}

// QAbstractEventDispatcher(class QObject *)
func NewQAbstractEventDispatcher(args ...interface{}) QAbstractEventDispatcher {
  // QAbstractEventDispatcher(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAbstractEventDispatcherC1EP7QObject
    // invoke: void QAbstractEventDispatcher(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN24QAbstractEventDispatcherC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "QAbstractEventDispatcher", args)
  }

  return QAbstractEventDispatcher{}
}

// metaObject()
func (this *QAbstractEventDispatcher) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAbstractEventDispatcher10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK24QAbstractEventDispatcher10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractEventDispatcher", "metaObject", args)
  }

}

// <= body block end

