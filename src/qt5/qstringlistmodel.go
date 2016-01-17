package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qstringlistmodel.h
// dst-file: /src/core/qstringlistmodel.go
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
  // proto:  const QMetaObject * QStringListModel::metaObject();
extern void _ZNK16QStringListModel10metaObjectEv(void* qthis); // 4
  // proto:  void QStringListModel::QStringListModel(const QStringList & strings, QObject * parent);
extern void _ZN16QStringListModelC2ERK11QStringListP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QStringListModel::QStringListModel(QObject * parent);
extern void _ZN16QStringListModelC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  QModelIndex QStringListModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK16QStringListModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  int QStringListModel::rowCount(const QModelIndex & parent);
extern void _ZNK16QStringListModel8rowCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QStringListModel::insertRows(int row, int count, const QModelIndex & parent);
extern void _ZN16QStringListModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  Qt::ItemFlags QStringListModel::flags(const QModelIndex & index);
extern void _ZNK16QStringListModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QStringListModel::removeRows(int row, int count, const QModelIndex & parent);
extern void _ZN16QStringListModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QStringListModel::setStringList(const QStringList & strings);
extern void _ZN16QStringListModel13setStringListERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QVariant QStringListModel::data(const QModelIndex & index, int role);
extern void _ZNK16QStringListModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QStringList QStringListModel::stringList();
extern void _ZNK16QStringListModel10stringListEv(void* qthis); // 4
  // proto:  Qt::DropActions QStringListModel::supportedDropActions();
extern void _ZNK16QStringListModel20supportedDropActionsEv(void* qthis); // 4
  // proto:  bool QStringListModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
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

// class sizeof(QStringListModel)=1
type QStringListModel struct {
  /*qbase*/ QAbstractListModel;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QStringListModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK16QStringListModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringListModel", "metaObject", args)
  }

}

// QStringListModel(const class QStringList &, class QObject *)
func NewQStringListModel(args ...interface{}) QStringListModel {
  // QStringListModel(const class QStringList &, class QObject *)
  // QStringListModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QStringListModelC1ERK11QStringListP7QObject
    // invoke: void QStringListModel(const class QStringList &, class QObject *)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN16QStringListModelC2ERK11QStringListP7QObject(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN16QStringListModelC1EP7QObject
    // invoke: void QStringListModel(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN16QStringListModelC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QStringListModel", "QStringListModel", args)
  }

  return QStringListModel{}
}

// sibling(int, int, const class QModelIndex &)
func (this *QStringListModel) sibling(args ...interface{}) () {
  // sibling(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel7siblingEiiRK11QModelIndex
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK16QStringListModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringListModel", "sibling", args)
  }

}

// rowCount(const class QModelIndex &)
func (this *QStringListModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel8rowCountERK11QModelIndex
    // invoke: int rowCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK16QStringListModel8rowCountERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringListModel", "rowCount", args)
  }

}

// insertRows(int, int, const class QModelIndex &)
func (this *QStringListModel) insertRows(args ...interface{}) () {
  // insertRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QStringListModel10insertRowsEiiRK11QModelIndex
    // invoke: bool insertRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN16QStringListModel10insertRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringListModel", "insertRows", args)
  }

}

// flags(const class QModelIndex &)
func (this *QStringListModel) flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK16QStringListModel5flagsERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringListModel", "flags", args)
  }

}

// removeRows(int, int, const class QModelIndex &)
func (this *QStringListModel) removeRows(args ...interface{}) () {
  // removeRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QStringListModel10removeRowsEiiRK11QModelIndex
    // invoke: bool removeRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN16QStringListModel10removeRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringListModel", "removeRows", args)
  }

}

// setStringList(const class QStringList &)
func (this *QStringListModel) setStringList(args ...interface{}) () {
  // setStringList(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QStringListModel13setStringListERK11QStringList
    // invoke: void setStringList(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN16QStringListModel13setStringListERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringListModel", "setStringList", args)
  }

}

// data(const class QModelIndex &, int)
func (this *QStringListModel) data(args ...interface{}) () {
  // data(const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel4dataERK11QModelIndexi
    // invoke: QVariant data(const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK16QStringListModel4dataERK11QModelIndexi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringListModel", "data", args)
  }

}

// stringList()
func (this *QStringListModel) stringList(args ...interface{}) () {
  // stringList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel10stringListEv
    // invoke: QStringList stringList()
    C._ZNK16QStringListModel10stringListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringListModel", "stringList", args)
  }

}

// supportedDropActions()
func (this *QStringListModel) supportedDropActions(args ...interface{}) () {
  // supportedDropActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QStringListModel20supportedDropActionsEv
    // invoke: Qt::DropActions supportedDropActions()
    C._ZNK16QStringListModel20supportedDropActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringListModel", "supportedDropActions", args)
  }

}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QStringListModel) setData(args ...interface{}) () {
  // setData(const class QModelIndex &, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringListModel", "setData", args)
  }

}

// <= body block end

