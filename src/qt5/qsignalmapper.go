package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qsignalmapper.h
// dst-file: /src/core/qsignalmapper.go
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
  // proto:  void QSignalMapper::setMapping(QObject * sender, QObject * object);
extern void _ZN13QSignalMapper10setMappingEP7QObjectS1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QSignalMapper::setMapping(QObject * sender, const QString & text);
extern void _ZN13QSignalMapper10setMappingEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QSignalMapper::setMapping(QObject * sender, int id);
extern void _ZN13QSignalMapper10setMappingEP7QObjecti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QSignalMapper::setMapping(QObject * sender, QWidget * widget);
extern void _ZN13QSignalMapper10setMappingEP7QObjectP7QWidget(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QSignalMapper::map();
extern void _ZN13QSignalMapper3mapEv(void* qthis); // 4
  // proto:  void QSignalMapper::map(QObject * sender);
extern void _ZN13QSignalMapper3mapEP7QObject(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QSignalMapper::metaObject();
extern void _ZNK13QSignalMapper10metaObjectEv(void* qthis); // 4
  // proto:  void QSignalMapper::~QSignalMapper();
extern void _ZN13QSignalMapperD2Ev(void* qthis); // 4
  // proto:  void QSignalMapper::removeMappings(QObject * sender);
extern void _ZN13QSignalMapper14removeMappingsEP7QObject(void* qthis, void* arg0); // 4
  // proto:  QObject * QSignalMapper::mapping(const QString & text);
extern void _ZNK13QSignalMapper7mappingERK7QString(void* qthis, void* arg0); // 4
  // proto:  QObject * QSignalMapper::mapping(int id);
extern void _ZNK13QSignalMapper7mappingEi(void* qthis, int32_t arg0); // 4
  // proto:  QObject * QSignalMapper::mapping(QObject * object);
extern void _ZNK13QSignalMapper7mappingEP7QObject(void* qthis, void* arg0); // 4
  // proto:  QObject * QSignalMapper::mapping(QWidget * widget);
extern void _ZNK13QSignalMapper7mappingEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QSignalMapper::QSignalMapper(QObject * parent);
extern void _ZN13QSignalMapperC2EP7QObject(void* qthis, void* arg0); // 3
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

// class sizeof(QSignalMapper)=1
type QSignalMapper struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _mapped QSignalMapper_mapped_signal;
}

// setMapping(class QObject *, class QObject *)
func (this *QSignalMapper) setMapping(args ...interface{}) () {
  // setMapping(class QObject *, class QObject *)
  // setMapping(class QObject *, const class QString &)
  // setMapping(class QObject *, int)
  // setMapping(class QObject *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectS1_
    // invoke: void setMapping(class QObject *, class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QSignalMapper10setMappingEP7QObjectS1_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectRK7QString
    // invoke: void setMapping(class QObject *, const class QString &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QSignalMapper10setMappingEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjecti
    // invoke: void setMapping(class QObject *, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QSignalMapper10setMappingEP7QObjecti(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectP7QWidget
    // invoke: void setMapping(class QObject *, class QWidget *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QSignalMapper10setMappingEP7QObjectP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSignalMapper", "setMapping", args)
  }

}

// map()
func (this *QSignalMapper) map_(args ...interface{}) () {
  // map()
  // map(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapper3mapEv
    // invoke: void map()
    C._ZN13QSignalMapper3mapEv(this.qclsinst)
  case 1:
    // invoke: _ZN13QSignalMapper3mapEP7QObject
    // invoke: void map(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QSignalMapper3mapEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSignalMapper", "map", args)
  }

}

// metaObject()
func (this *QSignalMapper) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSignalMapper10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QSignalMapper10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalMapper", "metaObject", args)
  }

}

// ~QSignalMapper()
func (this *QSignalMapper) FreeQSignalMapper(args ...interface{}) () {
  // ~QSignalMapper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapperD0Ev
    // invoke: void ~QSignalMapper()
    C._ZN13QSignalMapperD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalMapper", "~QSignalMapper", args)
  }

}

// removeMappings(class QObject *)
func (this *QSignalMapper) removeMappings(args ...interface{}) () {
  // removeMappings(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapper14removeMappingsEP7QObject
    // invoke: void removeMappings(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QSignalMapper14removeMappingsEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSignalMapper", "removeMappings", args)
  }

}

// mapping(const class QString &)
func (this *QSignalMapper) mapping(args ...interface{}) () {
  // mapping(const class QString &)
  // mapping(int)
  // mapping(class QObject *)
  // mapping(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSignalMapper7mappingERK7QString
    // invoke: QObject * mapping(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QSignalMapper7mappingERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK13QSignalMapper7mappingEi
    // invoke: QObject * mapping(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QSignalMapper7mappingEi(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK13QSignalMapper7mappingEP7QObject
    // invoke: QObject * mapping(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QSignalMapper7mappingEP7QObject(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK13QSignalMapper7mappingEP7QWidget
    // invoke: QObject * mapping(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QSignalMapper7mappingEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSignalMapper", "mapping", args)
  }

}

// QSignalMapper(class QObject *)
func NewQSignalMapper(args ...interface{}) QSignalMapper {
  // QSignalMapper(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapperC1EP7QObject
    // invoke: void QSignalMapper(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QSignalMapperC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QSignalMapper", "QSignalMapper", args)
  }

  return QSignalMapper{}
}

// <= body block end

