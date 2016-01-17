package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qgraphicstransform.h
// dst-file: /src/widgets/qgraphicstransform.go
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
  // proto:  QVector3D QGraphicsRotation::origin();
extern void _ZNK17QGraphicsRotation6originEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsRotation::metaObject();
extern void _ZNK17QGraphicsRotation10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsRotation::applyTo(QMatrix4x4 * matrix);
extern void _ZNK17QGraphicsRotation7applyToEP10QMatrix4x4(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRotation::setOrigin(const QVector3D & point);
extern void _ZN17QGraphicsRotation9setOriginERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsRotation::QGraphicsRotation(QObject * parent);
extern void _ZN17QGraphicsRotationC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QGraphicsRotation::setAngle(qreal );
extern void _ZN17QGraphicsRotation8setAngleEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsRotation::~QGraphicsRotation();
extern void _ZN17QGraphicsRotationD2Ev(void* qthis); // 4
  // proto:  qreal QGraphicsRotation::angle();
extern void _ZNK17QGraphicsRotation5angleEv(void* qthis); // 4
  // proto:  QVector3D QGraphicsRotation::axis();
extern void _ZNK17QGraphicsRotation4axisEv(void* qthis); // 4
  // proto:  void QGraphicsRotation::setAxis(const QVector3D & axis);
extern void _ZN17QGraphicsRotation7setAxisERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  QVector3D QGraphicsScale::origin();
extern void _ZNK14QGraphicsScale6originEv(void* qthis); // 4
  // proto:  qreal QGraphicsScale::zScale();
extern void _ZNK14QGraphicsScale6zScaleEv(void* qthis); // 4
  // proto:  void QGraphicsScale::QGraphicsScale(QObject * parent);
extern void _ZN14QGraphicsScaleC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  qreal QGraphicsScale::yScale();
extern void _ZNK14QGraphicsScale6yScaleEv(void* qthis); // 4
  // proto:  void QGraphicsScale::~QGraphicsScale();
extern void _ZN14QGraphicsScaleD2Ev(void* qthis); // 4
  // proto:  void QGraphicsScale::setZScale(qreal );
extern void _ZN14QGraphicsScale9setZScaleEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsScale::setYScale(qreal );
extern void _ZN14QGraphicsScale9setYScaleEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsScale::xScale();
extern void _ZNK14QGraphicsScale6xScaleEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsScale::metaObject();
extern void _ZNK14QGraphicsScale10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsScale::applyTo(QMatrix4x4 * matrix);
extern void _ZNK14QGraphicsScale7applyToEP10QMatrix4x4(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsScale::setOrigin(const QVector3D & point);
extern void _ZN14QGraphicsScale9setOriginERK9QVector3D(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsScale::setXScale(qreal );
extern void _ZN14QGraphicsScale9setXScaleEd(void* qthis, double arg0); // 4
  // proto:  const QMetaObject * QGraphicsTransform::metaObject();
extern void _ZNK18QGraphicsTransform10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsTransform::QGraphicsTransform(QObject * parent);
extern void _ZN18QGraphicsTransformC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QGraphicsTransform::~QGraphicsTransform();
extern void _ZN18QGraphicsTransformD2Ev(void* qthis); // 4
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

// class sizeof(QGraphicsRotation)=1
type QGraphicsRotation struct {
  /*qbase*/ QGraphicsTransform;
  qclsinst unsafe.Pointer /* *C.void */;
//  _originChanged QGraphicsRotation_originChanged_signal;
//  _axisChanged QGraphicsRotation_axisChanged_signal;
//  _angleChanged QGraphicsRotation_angleChanged_signal;
}

// class sizeof(QGraphicsScale)=1
type QGraphicsScale struct {
  /*qbase*/ QGraphicsTransform;
  qclsinst unsafe.Pointer /* *C.void */;
//  _yScaleChanged QGraphicsScale_yScaleChanged_signal;
//  _xScaleChanged QGraphicsScale_xScaleChanged_signal;
//  _zScaleChanged QGraphicsScale_zScaleChanged_signal;
//  _scaleChanged QGraphicsScale_scaleChanged_signal;
//  _originChanged QGraphicsScale_originChanged_signal;
}

// class sizeof(QGraphicsTransform)=1
type QGraphicsTransform struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// origin()
func (this *QGraphicsRotation) origin(args ...interface{}) () {
  // origin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation6originEv
    // invoke: QVector3D origin()
    C._ZNK17QGraphicsRotation6originEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "origin", args)
  }

}

// metaObject()
func (this *QGraphicsRotation) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK17QGraphicsRotation10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "metaObject", args)
  }

}

// applyTo(class QMatrix4x4 *)
func (this *QGraphicsRotation) applyTo(args ...interface{}) () {
  // applyTo(class QMatrix4x4 *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "QMatrix4x4 *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation7applyToEP10QMatrix4x4
    // invoke: void applyTo(class QMatrix4x4 *)
    var arg0 = args[0].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK17QGraphicsRotation7applyToEP10QMatrix4x4(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "applyTo", args)
  }

}

// setOrigin(const class QVector3D &)
func (this *QGraphicsRotation) setOrigin(args ...interface{}) () {
  // setOrigin(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotation9setOriginERK9QVector3D
    // invoke: void setOrigin(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QGraphicsRotation9setOriginERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setOrigin", args)
  }

}

// QGraphicsRotation(class QObject *)
func NewQGraphicsRotation(args ...interface{}) QGraphicsRotation {
  // QGraphicsRotation(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotationC1EP7QObject
    // invoke: void QGraphicsRotation(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN17QGraphicsRotationC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "QGraphicsRotation", args)
  }

  return QGraphicsRotation{}
}

// setAngle(qreal)
func (this *QGraphicsRotation) setAngle(args ...interface{}) () {
  // setAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotation8setAngleEd
    // invoke: void setAngle(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN17QGraphicsRotation8setAngleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setAngle", args)
  }

}

// ~QGraphicsRotation()
func (this *QGraphicsRotation) FreeQGraphicsRotation(args ...interface{}) () {
  // ~QGraphicsRotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotationD0Ev
    // invoke: void ~QGraphicsRotation()
    C._ZN17QGraphicsRotationD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "~QGraphicsRotation", args)
  }

}

// angle()
func (this *QGraphicsRotation) angle(args ...interface{}) () {
  // angle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation5angleEv
    // invoke: qreal angle()
    C._ZNK17QGraphicsRotation5angleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "angle", args)
  }

}

// axis()
func (this *QGraphicsRotation) axis(args ...interface{}) () {
  // axis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QGraphicsRotation4axisEv
    // invoke: QVector3D axis()
    C._ZNK17QGraphicsRotation4axisEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "axis", args)
  }

}

// setAxis(const class QVector3D &)
func (this *QGraphicsRotation) setAxis(args ...interface{}) () {
  // setAxis(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QGraphicsRotation7setAxisERK9QVector3D
    // invoke: void setAxis(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN17QGraphicsRotation7setAxisERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsRotation", "setAxis", args)
  }

}

// origin()
func (this *QGraphicsScale) origin(args ...interface{}) () {
  // origin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale6originEv
    // invoke: QVector3D origin()
    C._ZNK14QGraphicsScale6originEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "origin", args)
  }

}

// zScale()
func (this *QGraphicsScale) zScale(args ...interface{}) () {
  // zScale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale6zScaleEv
    // invoke: qreal zScale()
    C._ZNK14QGraphicsScale6zScaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "zScale", args)
  }

}

// QGraphicsScale(class QObject *)
func NewQGraphicsScale(args ...interface{}) QGraphicsScale {
  // QGraphicsScale(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScaleC1EP7QObject
    // invoke: void QGraphicsScale(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QGraphicsScaleC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "QGraphicsScale", args)
  }

  return QGraphicsScale{}
}

// yScale()
func (this *QGraphicsScale) yScale(args ...interface{}) () {
  // yScale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale6yScaleEv
    // invoke: qreal yScale()
    C._ZNK14QGraphicsScale6yScaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "yScale", args)
  }

}

// ~QGraphicsScale()
func (this *QGraphicsScale) FreeQGraphicsScale(args ...interface{}) () {
  // ~QGraphicsScale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScaleD0Ev
    // invoke: void ~QGraphicsScale()
    C._ZN14QGraphicsScaleD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "~QGraphicsScale", args)
  }

}

// setZScale(qreal)
func (this *QGraphicsScale) setZScale(args ...interface{}) () {
  // setZScale(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setZScaleEd
    // invoke: void setZScale(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScale9setZScaleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setZScale", args)
  }

}

// setYScale(qreal)
func (this *QGraphicsScale) setYScale(args ...interface{}) () {
  // setYScale(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setYScaleEd
    // invoke: void setYScale(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScale9setYScaleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setYScale", args)
  }

}

// xScale()
func (this *QGraphicsScale) xScale(args ...interface{}) () {
  // xScale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale6xScaleEv
    // invoke: qreal xScale()
    C._ZNK14QGraphicsScale6xScaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "xScale", args)
  }

}

// metaObject()
func (this *QGraphicsScale) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK14QGraphicsScale10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "metaObject", args)
  }

}

// applyTo(class QMatrix4x4 *)
func (this *QGraphicsScale) applyTo(args ...interface{}) () {
  // applyTo(class QMatrix4x4 *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix4x4{}) // "QMatrix4x4 *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScale7applyToEP10QMatrix4x4
    // invoke: void applyTo(class QMatrix4x4 *)
    var arg0 = args[0].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QGraphicsScale7applyToEP10QMatrix4x4(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "applyTo", args)
  }

}

// setOrigin(const class QVector3D &)
func (this *QGraphicsScale) setOrigin(args ...interface{}) () {
  // setOrigin(const class QVector3D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setOriginERK9QVector3D
    // invoke: void setOrigin(const class QVector3D &)
    var arg0 = args[0].(QVector3D).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScale9setOriginERK9QVector3D(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setOrigin", args)
  }

}

// setXScale(qreal)
func (this *QGraphicsScale) setXScale(args ...interface{}) () {
  // setXScale(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScale9setXScaleEd
    // invoke: void setXScale(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScale9setXScaleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScale", "setXScale", args)
  }

}

// metaObject()
func (this *QGraphicsTransform) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QGraphicsTransform10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK18QGraphicsTransform10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "metaObject", args)
  }

}

// QGraphicsTransform(class QObject *)
func NewQGraphicsTransform(args ...interface{}) QGraphicsTransform {
  // QGraphicsTransform(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsTransformC1EP7QObject
    // invoke: void QGraphicsTransform(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN18QGraphicsTransformC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "QGraphicsTransform", args)
  }

  return QGraphicsTransform{}
}

// ~QGraphicsTransform()
func (this *QGraphicsTransform) FreeQGraphicsTransform(args ...interface{}) () {
  // ~QGraphicsTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGraphicsTransformD0Ev
    // invoke: void ~QGraphicsTransform()
    C._ZN18QGraphicsTransformD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsTransform", "~QGraphicsTransform", args)
  }

}

// <= body block end

