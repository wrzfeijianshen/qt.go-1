package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qarraydata.h
// dst-file: /src/core/qarraydata.go
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
  // proto:  AllocationOptions QArrayData::detachFlags();
extern void _ZNK10QArrayData11detachFlagsEv(void* qthis); // 2
  // proto:  AllocationOptions QArrayData::cloneFlags();
extern void _ZNK10QArrayData10cloneFlagsEv(void* qthis); // 2
  // proto:  void * QArrayData::data();
extern void _ZN10QArrayData4dataEv(void* qthis); // 2
  // proto: static QArrayData * QArrayData::sharedNull();
extern void _ZN10QArrayData10sharedNullEv(); // 2
  // proto:  bool QArrayData::isMutable();
extern void _ZNK10QArrayData9isMutableEv(void* qthis); // 2
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

// class sizeof(QArrayData)=1
type QArrayData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// detachFlags()
func (this *QArrayData) detachFlags(args ...interface{}) () {
  // detachFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData11detachFlagsEv
    // invoke: AllocationOptions detachFlags()
    C._ZNK10QArrayData11detachFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QArrayData", "detachFlags", args)
  }

}

// cloneFlags()
func (this *QArrayData) cloneFlags(args ...interface{}) () {
  // cloneFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData10cloneFlagsEv
    // invoke: AllocationOptions cloneFlags()
    C._ZNK10QArrayData10cloneFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QArrayData", "cloneFlags", args)
  }

}

// data()
func (this *QArrayData) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QArrayData4dataEv
    // invoke: void * data()
    C._ZN10QArrayData4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QArrayData", "data", args)
  }

}

// sharedNull()
func (this *QArrayData) sharedNull_s(args ...interface{}) () {
  // sharedNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QArrayData10sharedNullEv
    // invoke: QArrayData * sharedNull()
    C._ZN10QArrayData10sharedNullEv()
  default:
    qtrt.ErrorResolve("QArrayData", "sharedNull", args)
  }

}

// isMutable()
func (this *QArrayData) isMutable(args ...interface{}) () {
  // isMutable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QArrayData9isMutableEv
    // invoke: bool isMutable()
    C._ZNK10QArrayData9isMutableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QArrayData", "isMutable", args)
  }

}

// <= body block end

