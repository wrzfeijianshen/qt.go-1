package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtGui/qevent.h
// dst-file: /src/gui/qevent.go
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
  // proto:  void QWhatsThisClickedEvent::QWhatsThisClickedEvent(const QString & href);
extern void _ZN22QWhatsThisClickedEventC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  void QWhatsThisClickedEvent::~QWhatsThisClickedEvent();
extern void _ZN22QWhatsThisClickedEventD2Ev(void* qthis); // 4
  // proto:  QString QWhatsThisClickedEvent::href();
extern void _ZNK22QWhatsThisClickedEvent4hrefEv(void* qthis); // 2
  // proto:  const QRegion & QExposeEvent::region();
extern void _ZNK12QExposeEvent6regionEv(void* qthis); // 2
  // proto:  void QExposeEvent::~QExposeEvent();
extern void _ZN12QExposeEventD2Ev(void* qthis); // 4
  // proto:  void QExposeEvent::QExposeEvent(const QRegion & rgn);
extern void _ZN12QExposeEventC2ERK7QRegion(void* qthis, void* arg0); // 3
  // proto:  int QInputMethodEvent::replacementStart();
extern void _ZNK17QInputMethodEvent16replacementStartEv(void* qthis); // 2
  // proto:  int QInputMethodEvent::replacementLength();
extern void _ZNK17QInputMethodEvent17replacementLengthEv(void* qthis); // 2
  // proto:  void QInputMethodEvent::QInputMethodEvent(const QInputMethodEvent & other);
extern void _ZN17QInputMethodEventC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QInputMethodEvent::QInputMethodEvent();
extern void _ZN17QInputMethodEventC2Ev(void* qthis); // 3
  // proto:  const QString & QInputMethodEvent::preeditString();
extern void _ZNK17QInputMethodEvent13preeditStringEv(void* qthis); // 2
  // proto:  const QList<QInputMethodEvent::Attribute> & QInputMethodEvent::attributes();
extern void _ZNK17QInputMethodEvent10attributesEv(void* qthis); // 2
  // proto:  const QString & QInputMethodEvent::commitString();
extern void _ZNK17QInputMethodEvent12commitStringEv(void* qthis); // 2
  // proto:  void QInputMethodEvent::setCommitString(const QString & commitString, int replaceFrom, int replaceLength);
extern void _ZN17QInputMethodEvent15setCommitStringERK7QStringii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  const QPoint & QHelpEvent::globalPos();
extern void _ZNK10QHelpEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPoint & QHelpEvent::pos();
extern void _ZNK10QHelpEvent3posEv(void* qthis); // 2
  // proto:  int QHelpEvent::y();
extern void _ZNK10QHelpEvent1yEv(void* qthis); // 2
  // proto:  int QHelpEvent::x();
extern void _ZNK10QHelpEvent1xEv(void* qthis); // 2
  // proto:  int QHelpEvent::globalX();
extern void _ZNK10QHelpEvent7globalXEv(void* qthis); // 2
  // proto:  int QHelpEvent::globalY();
extern void _ZNK10QHelpEvent7globalYEv(void* qthis); // 2
  // proto:  void QHelpEvent::~QHelpEvent();
extern void _ZN10QHelpEventD2Ev(void* qthis); // 4
  // proto:  QAction * QActionEvent::action();
extern void _ZNK12QActionEvent6actionEv(void* qthis); // 2
  // proto:  void QActionEvent::QActionEvent(int type, QAction * action, QAction * before);
extern void _ZN12QActionEventC2EiP7QActionS1_(void* qthis, int32_t arg0, void* arg1, void* arg2); // 3
  // proto:  void QActionEvent::~QActionEvent();
extern void _ZN12QActionEventD2Ev(void* qthis); // 4
  // proto:  QAction * QActionEvent::before();
extern void _ZNK12QActionEvent6beforeEv(void* qthis); // 2
  // proto:  const QPointF & QMouseEvent::localPos();
extern void _ZNK11QMouseEvent8localPosEv(void* qthis); // 2
  // proto:  const QPointF & QMouseEvent::screenPos();
extern void _ZNK11QMouseEvent9screenPosEv(void* qthis); // 2
  // proto:  QPoint QMouseEvent::globalPos();
extern void _ZNK11QMouseEvent9globalPosEv(void* qthis); // 2
  // proto:  Qt::MouseButton QMouseEvent::button();
extern void _ZNK11QMouseEvent6buttonEv(void* qthis); // 2
  // proto:  void QMouseEvent::~QMouseEvent();
extern void _ZN11QMouseEventD2Ev(void* qthis); // 4
  // proto:  QPoint QMouseEvent::pos();
extern void _ZNK11QMouseEvent3posEv(void* qthis); // 2
  // proto:  Qt::MouseButtons QMouseEvent::buttons();
extern void _ZNK11QMouseEvent7buttonsEv(void* qthis); // 2
  // proto:  Qt::MouseEventSource QMouseEvent::source();
extern void _ZNK11QMouseEvent6sourceEv(void* qthis); // 4
  // proto:  Qt::MouseEventFlags QMouseEvent::flags();
extern void _ZNK11QMouseEvent5flagsEv(void* qthis); // 4
  // proto:  int QMouseEvent::globalY();
extern void _ZNK11QMouseEvent7globalYEv(void* qthis); // 2
  // proto:  int QMouseEvent::y();
extern void _ZNK11QMouseEvent1yEv(void* qthis); // 2
  // proto:  int QMouseEvent::x();
extern void _ZNK11QMouseEvent1xEv(void* qthis); // 2
  // proto:  const QPointF & QMouseEvent::windowPos();
extern void _ZNK11QMouseEvent9windowPosEv(void* qthis); // 2
  // proto:  int QMouseEvent::globalX();
extern void _ZNK11QMouseEvent7globalXEv(void* qthis); // 2
  // proto:  QUrl QFileOpenEvent::url();
extern void _ZNK14QFileOpenEvent3urlEv(void* qthis); // 2
  // proto:  void QFileOpenEvent::~QFileOpenEvent();
extern void _ZN14QFileOpenEventD2Ev(void* qthis); // 4
  // proto:  QString QFileOpenEvent::file();
extern void _ZNK14QFileOpenEvent4fileEv(void* qthis); // 2
  // proto:  void QFileOpenEvent::QFileOpenEvent(const QString & file);
extern void _ZN14QFileOpenEventC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  void QFileOpenEvent::QFileOpenEvent(const QUrl & url);
extern void _ZN14QFileOpenEventC2ERK4QUrl(void* qthis, void* arg0); // 3
  // proto:  void QToolBarChangeEvent::~QToolBarChangeEvent();
extern void _ZN19QToolBarChangeEventD2Ev(void* qthis); // 4
  // proto:  bool QToolBarChangeEvent::toggle();
extern void _ZNK19QToolBarChangeEvent6toggleEv(void* qthis); // 2
  // proto:  void QToolBarChangeEvent::QToolBarChangeEvent(bool t);
extern void _ZN19QToolBarChangeEventC2Eb(void* qthis, bool arg0); // 3
  // proto:  int QTabletEvent::xTilt();
extern void _ZNK12QTabletEvent5xTiltEv(void* qthis); // 2
  // proto:  QPoint QTabletEvent::pos();
extern void _ZNK12QTabletEvent3posEv(void* qthis); // 2
  // proto:  QTabletEvent::PointerType QTabletEvent::pointerType();
extern void _ZNK12QTabletEvent11pointerTypeEv(void* qthis); // 2
  // proto:  Qt::MouseButtons QTabletEvent::buttons();
extern void _ZNK12QTabletEvent7buttonsEv(void* qthis); // 4
  // proto:  qreal QTabletEvent::hiResGlobalX();
extern void _ZNK12QTabletEvent12hiResGlobalXEv(void* qthis); // 2
  // proto:  qreal QTabletEvent::hiResGlobalY();
extern void _ZNK12QTabletEvent12hiResGlobalYEv(void* qthis); // 2
  // proto:  void QTabletEvent::~QTabletEvent();
extern void _ZN12QTabletEventD2Ev(void* qthis); // 4
  // proto:  qint64 QTabletEvent::uniqueId();
extern void _ZNK12QTabletEvent8uniqueIdEv(void* qthis); // 2
  // proto:  int QTabletEvent::globalX();
extern void _ZNK12QTabletEvent7globalXEv(void* qthis); // 2
  // proto:  int QTabletEvent::globalY();
extern void _ZNK12QTabletEvent7globalYEv(void* qthis); // 2
  // proto:  int QTabletEvent::yTilt();
extern void _ZNK12QTabletEvent5yTiltEv(void* qthis); // 2
  // proto:  const QPointF & QTabletEvent::globalPosF();
extern void _ZNK12QTabletEvent10globalPosFEv(void* qthis); // 2
  // proto:  qreal QTabletEvent::pressure();
extern void _ZNK12QTabletEvent8pressureEv(void* qthis); // 2
  // proto:  QTabletEvent::TabletDevice QTabletEvent::device();
extern void _ZNK12QTabletEvent6deviceEv(void* qthis); // 2
  // proto:  qreal QTabletEvent::rotation();
extern void _ZNK12QTabletEvent8rotationEv(void* qthis); // 2
  // proto:  QPoint QTabletEvent::globalPos();
extern void _ZNK12QTabletEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPointF & QTabletEvent::posF();
extern void _ZNK12QTabletEvent4posFEv(void* qthis); // 2
  // proto:  Qt::MouseButton QTabletEvent::button();
extern void _ZNK12QTabletEvent6buttonEv(void* qthis); // 4
  // proto:  int QTabletEvent::y();
extern void _ZNK12QTabletEvent1yEv(void* qthis); // 2
  // proto:  int QTabletEvent::x();
extern void _ZNK12QTabletEvent1xEv(void* qthis); // 2
  // proto:  int QTabletEvent::z();
extern void _ZNK12QTabletEvent1zEv(void* qthis); // 2
  // proto:  qreal QTabletEvent::tangentialPressure();
extern void _ZNK12QTabletEvent18tangentialPressureEv(void* qthis); // 2
  // proto:  Qt::TouchPointStates QTouchEvent::touchPointStates();
extern void _ZNK11QTouchEvent16touchPointStatesEv(void* qthis); // 2
  // proto:  QObject * QTouchEvent::target();
extern void _ZNK11QTouchEvent6targetEv(void* qthis); // 2
  // proto:  void QTouchEvent::setTarget(QObject * atarget);
extern void _ZN11QTouchEvent9setTargetEP7QObject(void* qthis, void* arg0); // 2
  // proto:  const QList<QTouchEvent::TouchPoint> & QTouchEvent::touchPoints();
extern void _ZNK11QTouchEvent11touchPointsEv(void* qthis); // 2
  // proto:  void QTouchEvent::~QTouchEvent();
extern void _ZN11QTouchEventD2Ev(void* qthis); // 4
  // proto:  QWindow * QTouchEvent::window();
extern void _ZNK11QTouchEvent6windowEv(void* qthis); // 2
  // proto:  void QTouchEvent::setDevice(QTouchDevice * adevice);
extern void _ZN11QTouchEvent9setDeviceEP12QTouchDevice(void* qthis, void* arg0); // 2
  // proto:  QTouchDevice * QTouchEvent::device();
extern void _ZNK11QTouchEvent6deviceEv(void* qthis); // 2
  // proto:  void QTouchEvent::setWindow(QWindow * awindow);
extern void _ZN11QTouchEvent9setWindowEP7QWindow(void* qthis, void* arg0); // 2
  // proto:  QScreen * QScreenOrientationChangeEvent::screen();
extern void _ZNK29QScreenOrientationChangeEvent6screenEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientation QScreenOrientationChangeEvent::orientation();
extern void _ZNK29QScreenOrientationChangeEvent11orientationEv(void* qthis); // 4
  // proto:  void QScreenOrientationChangeEvent::~QScreenOrientationChangeEvent();
extern void _ZN29QScreenOrientationChangeEventD2Ev(void* qthis); // 4
  // proto:  void QIconDragEvent::~QIconDragEvent();
extern void _ZN14QIconDragEventD2Ev(void* qthis); // 4
  // proto:  void QIconDragEvent::QIconDragEvent();
extern void _ZN14QIconDragEventC2Ev(void* qthis); // 3
  // proto:  void QCloseEvent::QCloseEvent();
extern void _ZN11QCloseEventC2Ev(void* qthis); // 3
  // proto:  void QCloseEvent::~QCloseEvent();
extern void _ZN11QCloseEventD2Ev(void* qthis); // 4
  // proto:  void QDragEnterEvent::~QDragEnterEvent();
extern void _ZN15QDragEnterEventD2Ev(void* qthis); // 4
  // proto:  QPoint QWheelEvent::pixelDelta();
extern void _ZNK11QWheelEvent10pixelDeltaEv(void* qthis); // 2
  // proto:  const QPointF & QWheelEvent::globalPosF();
extern void _ZNK11QWheelEvent10globalPosFEv(void* qthis); // 2
  // proto:  void QWheelEvent::~QWheelEvent();
extern void _ZN11QWheelEventD2Ev(void* qthis); // 4
  // proto:  Qt::Orientation QWheelEvent::orientation();
extern void _ZNK11QWheelEvent11orientationEv(void* qthis); // 2
  // proto:  const QPointF & QWheelEvent::posF();
extern void _ZNK11QWheelEvent4posFEv(void* qthis); // 2
  // proto:  QPoint QWheelEvent::angleDelta();
extern void _ZNK11QWheelEvent10angleDeltaEv(void* qthis); // 2
  // proto:  QPoint QWheelEvent::pos();
extern void _ZNK11QWheelEvent3posEv(void* qthis); // 2
  // proto:  int QWheelEvent::y();
extern void _ZNK11QWheelEvent1yEv(void* qthis); // 2
  // proto:  Qt::MouseButtons QWheelEvent::buttons();
extern void _ZNK11QWheelEvent7buttonsEv(void* qthis); // 2
  // proto:  Qt::MouseEventSource QWheelEvent::source();
extern void _ZNK11QWheelEvent6sourceEv(void* qthis); // 2
  // proto:  int QWheelEvent::delta();
extern void _ZNK11QWheelEvent5deltaEv(void* qthis); // 2
  // proto:  Qt::ScrollPhase QWheelEvent::phase();
extern void _ZNK11QWheelEvent5phaseEv(void* qthis); // 2
  // proto:  int QWheelEvent::globalY();
extern void _ZNK11QWheelEvent7globalYEv(void* qthis); // 2
  // proto:  int QWheelEvent::globalX();
extern void _ZNK11QWheelEvent7globalXEv(void* qthis); // 2
  // proto:  int QWheelEvent::x();
extern void _ZNK11QWheelEvent1xEv(void* qthis); // 2
  // proto:  QPoint QWheelEvent::globalPos();
extern void _ZNK11QWheelEvent9globalPosEv(void* qthis); // 2
  // proto:  QPointF QScrollEvent::contentPos();
extern void _ZNK12QScrollEvent10contentPosEv(void* qthis); // 4
  // proto:  QPointF QScrollEvent::overshootDistance();
extern void _ZNK12QScrollEvent17overshootDistanceEv(void* qthis); // 4
  // proto:  void QScrollEvent::~QScrollEvent();
extern void _ZN12QScrollEventD2Ev(void* qthis); // 4
  // proto:  QScrollEvent::ScrollState QScrollEvent::scrollState();
extern void _ZNK12QScrollEvent11scrollStateEv(void* qthis); // 4
  // proto:  const QPointF & QHoverEvent::posF();
extern void _ZNK11QHoverEvent4posFEv(void* qthis); // 2
  // proto:  QPoint QHoverEvent::pos();
extern void _ZNK11QHoverEvent3posEv(void* qthis); // 2
  // proto:  QPoint QHoverEvent::oldPos();
extern void _ZNK11QHoverEvent6oldPosEv(void* qthis); // 2
  // proto:  const QPointF & QHoverEvent::oldPosF();
extern void _ZNK11QHoverEvent7oldPosFEv(void* qthis); // 2
  // proto:  void QHoverEvent::~QHoverEvent();
extern void _ZN11QHoverEventD2Ev(void* qthis); // 4
  // proto:  void QDragMoveEvent::ignore();
extern void _ZN14QDragMoveEvent6ignoreEv(void* qthis); // 2
  // proto:  void QDragMoveEvent::ignore(const QRect & r);
extern void _ZN14QDragMoveEvent6ignoreERK5QRect(void* qthis, void* arg0); // 2
  // proto:  void QDragMoveEvent::~QDragMoveEvent();
extern void _ZN14QDragMoveEventD2Ev(void* qthis); // 4
  // proto:  QRect QDragMoveEvent::answerRect();
extern void _ZNK14QDragMoveEvent10answerRectEv(void* qthis); // 2
  // proto:  void QDragMoveEvent::accept(const QRect & r);
extern void _ZN14QDragMoveEvent6acceptERK5QRect(void* qthis, void* arg0); // 2
  // proto:  void QDragMoveEvent::accept();
extern void _ZN14QDragMoveEvent6acceptEv(void* qthis); // 2
  // proto:  void QShowEvent::~QShowEvent();
extern void _ZN10QShowEventD2Ev(void* qthis); // 4
  // proto:  void QShowEvent::QShowEvent();
extern void _ZN10QShowEventC2Ev(void* qthis); // 3
  // proto:  QPlatformSurfaceEvent::SurfaceEventType QPlatformSurfaceEvent::surfaceEventType();
extern void _ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv(void* qthis); // 2
  // proto:  void QPlatformSurfaceEvent::~QPlatformSurfaceEvent();
extern void _ZN21QPlatformSurfaceEventD2Ev(void* qthis); // 4
  // proto:  void QPaintEvent::QPaintEvent(const QRect & paintRect);
extern void _ZN11QPaintEventC2ERK5QRect(void* qthis, void* arg0); // 3
  // proto:  void QPaintEvent::QPaintEvent(const QRegion & paintRegion);
extern void _ZN11QPaintEventC2ERK7QRegion(void* qthis, void* arg0); // 3
  // proto:  void QPaintEvent::~QPaintEvent();
extern void _ZN11QPaintEventD2Ev(void* qthis); // 4
  // proto:  const QRegion & QPaintEvent::region();
extern void _ZNK11QPaintEvent6regionEv(void* qthis); // 2
  // proto:  const QRect & QPaintEvent::rect();
extern void _ZNK11QPaintEvent4rectEv(void* qthis); // 2
  // proto:  Qt::FocusReason QFocusEvent::reason();
extern void _ZNK11QFocusEvent6reasonEv(void* qthis); // 4
  // proto:  void QFocusEvent::~QFocusEvent();
extern void _ZN11QFocusEventD2Ev(void* qthis); // 4
  // proto:  bool QFocusEvent::gotFocus();
extern void _ZNK11QFocusEvent8gotFocusEv(void* qthis); // 2
  // proto:  bool QFocusEvent::lostFocus();
extern void _ZNK11QFocusEvent9lostFocusEv(void* qthis); // 2
  // proto:  const QPointF & QNativeGestureEvent::localPos();
extern void _ZNK19QNativeGestureEvent8localPosEv(void* qthis); // 2
  // proto:  const QPointF & QNativeGestureEvent::screenPos();
extern void _ZNK19QNativeGestureEvent9screenPosEv(void* qthis); // 2
  // proto:  const QPoint QNativeGestureEvent::globalPos();
extern void _ZNK19QNativeGestureEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPoint QNativeGestureEvent::pos();
extern void _ZNK19QNativeGestureEvent3posEv(void* qthis); // 2
  // proto:  qreal QNativeGestureEvent::value();
extern void _ZNK19QNativeGestureEvent5valueEv(void* qthis); // 2
  // proto:  Qt::NativeGestureType QNativeGestureEvent::gestureType();
extern void _ZNK19QNativeGestureEvent11gestureTypeEv(void* qthis); // 2
  // proto:  const QPointF & QNativeGestureEvent::windowPos();
extern void _ZNK19QNativeGestureEvent9windowPosEv(void* qthis); // 2
  // proto:  void QResizeEvent::~QResizeEvent();
extern void _ZN12QResizeEventD2Ev(void* qthis); // 4
  // proto:  const QSize & QResizeEvent::oldSize();
extern void _ZNK12QResizeEvent7oldSizeEv(void* qthis); // 2
  // proto:  void QResizeEvent::QResizeEvent(const QSize & size, const QSize & oldSize);
extern void _ZN12QResizeEventC2ERK5QSizeS2_(void* qthis, void* arg0, void* arg1); // 3
  // proto:  const QSize & QResizeEvent::size();
extern void _ZNK12QResizeEvent4sizeEv(void* qthis); // 2
  // proto:  QString QStatusTipEvent::tip();
extern void _ZNK15QStatusTipEvent3tipEv(void* qthis); // 2
  // proto:  void QStatusTipEvent::~QStatusTipEvent();
extern void _ZN15QStatusTipEventD2Ev(void* qthis); // 4
  // proto:  void QStatusTipEvent::QStatusTipEvent(const QString & tip);
extern void _ZN15QStatusTipEventC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  const QPointF & QEnterEvent::localPos();
extern void _ZNK11QEnterEvent8localPosEv(void* qthis); // 2
  // proto:  const QPointF & QEnterEvent::screenPos();
extern void _ZNK11QEnterEvent9screenPosEv(void* qthis); // 2
  // proto:  QPoint QEnterEvent::globalPos();
extern void _ZNK11QEnterEvent9globalPosEv(void* qthis); // 2
  // proto:  void QEnterEvent::QEnterEvent(const QPointF & localPos, const QPointF & windowPos, const QPointF & screenPos);
extern void _ZN11QEnterEventC2ERK7QPointFS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2); // 3
  // proto:  QPoint QEnterEvent::pos();
extern void _ZNK11QEnterEvent3posEv(void* qthis); // 2
  // proto:  int QEnterEvent::globalX();
extern void _ZNK11QEnterEvent7globalXEv(void* qthis); // 2
  // proto:  void QEnterEvent::~QEnterEvent();
extern void _ZN11QEnterEventD2Ev(void* qthis); // 4
  // proto:  int QEnterEvent::globalY();
extern void _ZNK11QEnterEvent7globalYEv(void* qthis); // 2
  // proto:  int QEnterEvent::y();
extern void _ZNK11QEnterEvent1yEv(void* qthis); // 2
  // proto:  int QEnterEvent::x();
extern void _ZNK11QEnterEvent1xEv(void* qthis); // 2
  // proto:  const QPointF & QEnterEvent::windowPos();
extern void _ZNK11QEnterEvent9windowPosEv(void* qthis); // 2
  // proto:  void QMoveEvent::QMoveEvent(const QPoint & pos, const QPoint & oldPos);
extern void _ZN10QMoveEventC2ERK6QPointS2_(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QMoveEvent::~QMoveEvent();
extern void _ZN10QMoveEventD2Ev(void* qthis); // 4
  // proto:  const QPoint & QMoveEvent::oldPos();
extern void _ZNK10QMoveEvent6oldPosEv(void* qthis); // 2
  // proto:  const QPoint & QMoveEvent::pos();
extern void _ZNK10QMoveEvent3posEv(void* qthis); // 2
  // proto:  void QHideEvent::QHideEvent();
extern void _ZN10QHideEventC2Ev(void* qthis); // 3
  // proto:  void QHideEvent::~QHideEvent();
extern void _ZN10QHideEventD2Ev(void* qthis); // 4
  // proto:  void QDragLeaveEvent::~QDragLeaveEvent();
extern void _ZN15QDragLeaveEventD2Ev(void* qthis); // 4
  // proto:  void QDragLeaveEvent::QDragLeaveEvent();
extern void _ZN15QDragLeaveEventC2Ev(void* qthis); // 3
  // proto:  const QMimeData * QDropEvent::mimeData();
extern void _ZNK10QDropEvent8mimeDataEv(void* qthis); // 2
  // proto:  void QDropEvent::acceptProposedAction();
extern void _ZN10QDropEvent20acceptProposedActionEv(void* qthis); // 2
  // proto:  Qt::DropActions QDropEvent::possibleActions();
extern void _ZNK10QDropEvent15possibleActionsEv(void* qthis); // 2
  // proto:  const QPointF & QDropEvent::posF();
extern void _ZNK10QDropEvent4posFEv(void* qthis); // 2
  // proto:  QPoint QDropEvent::pos();
extern void _ZNK10QDropEvent3posEv(void* qthis); // 2
  // proto:  QObject * QDropEvent::source();
extern void _ZNK10QDropEvent6sourceEv(void* qthis); // 4
  // proto:  Qt::DropAction QDropEvent::proposedAction();
extern void _ZNK10QDropEvent14proposedActionEv(void* qthis); // 2
  // proto:  Qt::DropAction QDropEvent::dropAction();
extern void _ZNK10QDropEvent10dropActionEv(void* qthis); // 2
  // proto:  Qt::KeyboardModifiers QDropEvent::keyboardModifiers();
extern void _ZNK10QDropEvent17keyboardModifiersEv(void* qthis); // 2
  // proto:  void QDropEvent::~QDropEvent();
extern void _ZN10QDropEventD2Ev(void* qthis); // 4
  // proto:  Qt::MouseButtons QDropEvent::mouseButtons();
extern void _ZNK10QDropEvent12mouseButtonsEv(void* qthis); // 2
  // proto:  Qt::KeyboardModifiers QInputEvent::modifiers();
extern void _ZNK11QInputEvent9modifiersEv(void* qthis); // 2
  // proto:  ulong QInputEvent::timestamp();
extern void _ZNK11QInputEvent9timestampEv(void* qthis); // 2
  // proto:  void QInputEvent::setTimestamp(ulong atimestamp);
extern void _ZN11QInputEvent12setTimestampEm(void* qthis, int32_t arg0); // 2
  // proto:  void QInputEvent::~QInputEvent();
extern void _ZN11QInputEventD2Ev(void* qthis); // 4
  // proto:  Qt::ApplicationState QApplicationStateChangeEvent::applicationState();
extern void _ZNK28QApplicationStateChangeEvent16applicationStateEv(void* qthis); // 4
  // proto:  int QKeyEvent::count();
extern void _ZNK9QKeyEvent5countEv(void* qthis); // 2
  // proto:  Qt::KeyboardModifiers QKeyEvent::modifiers();
extern void _ZNK9QKeyEvent9modifiersEv(void* qthis); // 4
  // proto:  quint32 QKeyEvent::nativeModifiers();
extern void _ZNK9QKeyEvent15nativeModifiersEv(void* qthis); // 2
  // proto:  QString QKeyEvent::text();
extern void _ZNK9QKeyEvent4textEv(void* qthis); // 2
  // proto:  void QKeyEvent::~QKeyEvent();
extern void _ZN9QKeyEventD2Ev(void* qthis); // 4
  // proto:  quint32 QKeyEvent::nativeScanCode();
extern void _ZNK9QKeyEvent14nativeScanCodeEv(void* qthis); // 2
  // proto:  bool QKeyEvent::isAutoRepeat();
extern void _ZNK9QKeyEvent12isAutoRepeatEv(void* qthis); // 2
  // proto:  int QKeyEvent::key();
extern void _ZNK9QKeyEvent3keyEv(void* qthis); // 2
  // proto:  quint32 QKeyEvent::nativeVirtualKey();
extern void _ZNK9QKeyEvent16nativeVirtualKeyEv(void* qthis); // 2
  // proto:  const QPoint & QContextMenuEvent::globalPos();
extern void _ZNK17QContextMenuEvent9globalPosEv(void* qthis); // 2
  // proto:  const QPoint & QContextMenuEvent::pos();
extern void _ZNK17QContextMenuEvent3posEv(void* qthis); // 2
  // proto:  int QContextMenuEvent::y();
extern void _ZNK17QContextMenuEvent1yEv(void* qthis); // 2
  // proto:  QContextMenuEvent::Reason QContextMenuEvent::reason();
extern void _ZNK17QContextMenuEvent6reasonEv(void* qthis); // 2
  // proto:  int QContextMenuEvent::x();
extern void _ZNK17QContextMenuEvent1xEv(void* qthis); // 2
  // proto:  int QContextMenuEvent::globalX();
extern void _ZNK17QContextMenuEvent7globalXEv(void* qthis); // 2
  // proto:  int QContextMenuEvent::globalY();
extern void _ZNK17QContextMenuEvent7globalYEv(void* qthis); // 2
  // proto:  void QContextMenuEvent::~QContextMenuEvent();
extern void _ZN17QContextMenuEventD2Ev(void* qthis); // 4
  // proto:  void QScrollPrepareEvent::setContentPosRange(const QRectF & rect);
extern void _ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  QPointF QScrollPrepareEvent::contentPos();
extern void _ZNK19QScrollPrepareEvent10contentPosEv(void* qthis); // 4
  // proto:  void QScrollPrepareEvent::setContentPos(const QPointF & pos);
extern void _ZN19QScrollPrepareEvent13setContentPosERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QScrollPrepareEvent::~QScrollPrepareEvent();
extern void _ZN19QScrollPrepareEventD2Ev(void* qthis); // 4
  // proto:  QSizeF QScrollPrepareEvent::viewportSize();
extern void _ZNK19QScrollPrepareEvent12viewportSizeEv(void* qthis); // 4
  // proto:  void QScrollPrepareEvent::setViewportSize(const QSizeF & size);
extern void _ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QScrollPrepareEvent::QScrollPrepareEvent(const QPointF & startPos);
extern void _ZN19QScrollPrepareEventC2ERK7QPointF(void* qthis, void* arg0); // 3
  // proto:  QRectF QScrollPrepareEvent::contentPosRange();
extern void _ZNK19QScrollPrepareEvent15contentPosRangeEv(void* qthis); // 4
  // proto:  QPointF QScrollPrepareEvent::startPos();
extern void _ZNK19QScrollPrepareEvent8startPosEv(void* qthis); // 4
  // proto:  void QShortcutEvent::~QShortcutEvent();
extern void _ZN14QShortcutEventD2Ev(void* qthis); // 4
  // proto:  bool QShortcutEvent::isAmbiguous();
extern void _ZNK14QShortcutEvent11isAmbiguousEv(void* qthis); // 2
  // proto:  int QShortcutEvent::shortcutId();
extern void _ZNK14QShortcutEvent10shortcutIdEv(void* qthis); // 2
  // proto:  const QKeySequence & QShortcutEvent::key();
extern void _ZNK14QShortcutEvent3keyEv(void* qthis); // 2
  // proto:  void QShortcutEvent::QShortcutEvent(const QKeySequence & key, int id, bool ambiguous);
extern void _ZN14QShortcutEventC2ERK12QKeySequenceib(void* qthis, void* arg0, int32_t arg1, bool arg2); // 3
  // proto:  bool QWindowStateChangeEvent::isOverride();
extern void _ZNK23QWindowStateChangeEvent10isOverrideEv(void* qthis); // 4
  // proto:  Qt::WindowStates QWindowStateChangeEvent::oldState();
extern void _ZNK23QWindowStateChangeEvent8oldStateEv(void* qthis); // 2
  // proto:  void QWindowStateChangeEvent::~QWindowStateChangeEvent();
extern void _ZN23QWindowStateChangeEventD2Ev(void* qthis); // 4
  // proto:  void QInputMethodQueryEvent::~QInputMethodQueryEvent();
extern void _ZN22QInputMethodQueryEventD2Ev(void* qthis); // 4
  // proto:  Qt::InputMethodQueries QInputMethodQueryEvent::queries();
extern void _ZNK22QInputMethodQueryEvent7queriesEv(void* qthis); // 2
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

// class sizeof(QWhatsThisClickedEvent)=32
type QWhatsThisClickedEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QExposeEvent)=32
type QExposeEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QInputMethodEvent)=1
type QInputMethodEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QHelpEvent)=40
type QHelpEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QActionEvent)=40
type QActionEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMouseEvent)=1
type QMouseEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFileOpenEvent)=40
type QFileOpenEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QToolBarChangeEvent)=24
type QToolBarChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTabletEvent)=1
type QTabletEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTouchEvent)=1
type QTouchEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QScreenOrientationChangeEvent)=40
type QScreenOrientationChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QIconDragEvent)=24
type QIconDragEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QCloseEvent)=24
type QCloseEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDragEnterEvent)=1
type QDragEnterEvent struct {
  /*qbase*/ QDragMoveEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QWheelEvent)=1
type QWheelEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QScrollEvent)=64
type QScrollEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QHoverEvent)=1
type QHoverEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDragMoveEvent)=1
type QDragMoveEvent struct {
  /*qbase*/ QDropEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QShowEvent)=24
type QShowEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPlatformSurfaceEvent)=24
type QPlatformSurfaceEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPaintEvent)=56
type QPaintEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFocusEvent)=24
type QFocusEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QNativeGestureEvent)=1
type QNativeGestureEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QResizeEvent)=40
type QResizeEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStatusTipEvent)=32
type QStatusTipEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QEnterEvent)=72
type QEnterEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMoveEvent)=40
type QMoveEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QHideEvent)=24
type QHideEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDragLeaveEvent)=24
type QDragLeaveEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDropEvent)=1
type QDropEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QInputEvent)=1
type QInputEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QApplicationStateChangeEvent)=24
type QApplicationStateChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QKeyEvent)=1
type QKeyEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QContextMenuEvent)=1
type QContextMenuEvent struct {
  /*qbase*/ QInputEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QScrollPrepareEvent)=112
type QScrollPrepareEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QShortcutEvent)=40
type QShortcutEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QWindowStateChangeEvent)=1
type QWindowStateChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QInputMethodQueryEvent)=1
type QInputMethodQueryEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QWhatsThisClickedEvent(const class QString &)
func NewQWhatsThisClickedEvent(args ...interface{}) QWhatsThisClickedEvent {
  // QWhatsThisClickedEvent(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QWhatsThisClickedEventC1ERK7QString
    // invoke: void QWhatsThisClickedEvent(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN22QWhatsThisClickedEventC2ERK7QString(qthis, arg0)
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "QWhatsThisClickedEvent", args)
  }

  return QWhatsThisClickedEvent{}
}

// ~QWhatsThisClickedEvent()
func (this *QWhatsThisClickedEvent) FreeQWhatsThisClickedEvent(args ...interface{}) () {
  // ~QWhatsThisClickedEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QWhatsThisClickedEventD0Ev
    // invoke: void ~QWhatsThisClickedEvent()
    C._ZN22QWhatsThisClickedEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "~QWhatsThisClickedEvent", args)
  }

}

// href()
func (this *QWhatsThisClickedEvent) href(args ...interface{}) () {
  // href()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QWhatsThisClickedEvent4hrefEv
    // invoke: QString href()
    C._ZNK22QWhatsThisClickedEvent4hrefEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWhatsThisClickedEvent", "href", args)
  }

}

// region()
func (this *QExposeEvent) region(args ...interface{}) () {
  // region()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QExposeEvent6regionEv
    // invoke: const QRegion & region()
    C._ZNK12QExposeEvent6regionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QExposeEvent", "region", args)
  }

}

// ~QExposeEvent()
func (this *QExposeEvent) FreeQExposeEvent(args ...interface{}) () {
  // ~QExposeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QExposeEventD0Ev
    // invoke: void ~QExposeEvent()
    C._ZN12QExposeEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QExposeEvent", "~QExposeEvent", args)
  }

}

// QExposeEvent(const class QRegion &)
func NewQExposeEvent(args ...interface{}) QExposeEvent {
  // QExposeEvent(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QExposeEventC1ERK7QRegion
    // invoke: void QExposeEvent(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QExposeEventC2ERK7QRegion(qthis, arg0)
  default:
    qtrt.ErrorResolve("QExposeEvent", "QExposeEvent", args)
  }

  return QExposeEvent{}
}

// replacementStart()
func (this *QInputMethodEvent) replacementStart(args ...interface{}) () {
  // replacementStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent16replacementStartEv
    // invoke: int replacementStart()
    C._ZNK17QInputMethodEvent16replacementStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "replacementStart", args)
  }

}

// replacementLength()
func (this *QInputMethodEvent) replacementLength(args ...interface{}) () {
  // replacementLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent17replacementLengthEv
    // invoke: int replacementLength()
    C._ZNK17QInputMethodEvent17replacementLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "replacementLength", args)
  }

}

// QInputMethodEvent(const class QInputMethodEvent &)
func NewQInputMethodEvent(args ...interface{}) QInputMethodEvent {
  // QInputMethodEvent(const class QInputMethodEvent &)
  // QInputMethodEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QInputMethodEvent{}) // "const QInputMethodEvent &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QInputMethodEventC1ERKS_
    // invoke: void QInputMethodEvent(const class QInputMethodEvent &)
    var arg0 = args[0].(QInputMethodEvent).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN17QInputMethodEventC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN17QInputMethodEventC1Ev
    // invoke: void QInputMethodEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN17QInputMethodEventC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "QInputMethodEvent", args)
  }

  return QInputMethodEvent{}
}

// preeditString()
func (this *QInputMethodEvent) preeditString(args ...interface{}) () {
  // preeditString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent13preeditStringEv
    // invoke: const QString & preeditString()
    C._ZNK17QInputMethodEvent13preeditStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "preeditString", args)
  }

}

// attributes()
func (this *QInputMethodEvent) attributes(args ...interface{}) () {
  // attributes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent10attributesEv
    // invoke: const QList<QInputMethodEvent::Attribute> & attributes()
    C._ZNK17QInputMethodEvent10attributesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "attributes", args)
  }

}

// commitString()
func (this *QInputMethodEvent) commitString(args ...interface{}) () {
  // commitString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QInputMethodEvent12commitStringEv
    // invoke: const QString & commitString()
    C._ZNK17QInputMethodEvent12commitStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "commitString", args)
  }

}

// setCommitString(const class QString &, int, int)
func (this *QInputMethodEvent) setCommitString(args ...interface{}) () {
  // setCommitString(const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QInputMethodEvent15setCommitStringERK7QStringii
    // invoke: void setCommitString(const class QString &, int, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN17QInputMethodEvent15setCommitStringERK7QStringii(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QInputMethodEvent", "setCommitString", args)
  }

}

// globalPos()
func (this *QHelpEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent9globalPosEv
    // invoke: const QPoint & globalPos()
    C._ZNK10QHelpEvent9globalPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalPos", args)
  }

}

// pos()
func (this *QHelpEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent3posEv
    // invoke: const QPoint & pos()
    C._ZNK10QHelpEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "pos", args)
  }

}

// y()
func (this *QHelpEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent1yEv
    // invoke: int y()
    C._ZNK10QHelpEvent1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "y", args)
  }

}

// x()
func (this *QHelpEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent1xEv
    // invoke: int x()
    C._ZNK10QHelpEvent1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "x", args)
  }

}

// globalX()
func (this *QHelpEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent7globalXEv
    // invoke: int globalX()
    C._ZNK10QHelpEvent7globalXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalX", args)
  }

}

// globalY()
func (this *QHelpEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QHelpEvent7globalYEv
    // invoke: int globalY()
    C._ZNK10QHelpEvent7globalYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "globalY", args)
  }

}

// ~QHelpEvent()
func (this *QHelpEvent) FreeQHelpEvent(args ...interface{}) () {
  // ~QHelpEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QHelpEventD0Ev
    // invoke: void ~QHelpEvent()
    C._ZN10QHelpEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHelpEvent", "~QHelpEvent", args)
  }

}

// action()
func (this *QActionEvent) action(args ...interface{}) () {
  // action()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionEvent6actionEv
    // invoke: QAction * action()
    C._ZNK12QActionEvent6actionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionEvent", "action", args)
  }

}

// QActionEvent(int, class QAction *, class QAction *)
func NewQActionEvent(args ...interface{}) QActionEvent {
  // QActionEvent(int, class QAction *, class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][2] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionEventC1EiP7QActionS1_
    // invoke: void QActionEvent(int, class QAction *, class QAction *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAction).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QAction).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QActionEventC2EiP7QActionS1_(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QActionEvent", "QActionEvent", args)
  }

  return QActionEvent{}
}

// ~QActionEvent()
func (this *QActionEvent) FreeQActionEvent(args ...interface{}) () {
  // ~QActionEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionEventD0Ev
    // invoke: void ~QActionEvent()
    C._ZN12QActionEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionEvent", "~QActionEvent", args)
  }

}

// before()
func (this *QActionEvent) before(args ...interface{}) () {
  // before()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionEvent6beforeEv
    // invoke: QAction * before()
    C._ZNK12QActionEvent6beforeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionEvent", "before", args)
  }

}

// localPos()
func (this *QMouseEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent8localPosEv
    // invoke: const QPointF & localPos()
    C._ZNK11QMouseEvent8localPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "localPos", args)
  }

}

// screenPos()
func (this *QMouseEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9screenPosEv
    // invoke: const QPointF & screenPos()
    C._ZNK11QMouseEvent9screenPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "screenPos", args)
  }

}

// globalPos()
func (this *QMouseEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9globalPosEv
    // invoke: QPoint globalPos()
    C._ZNK11QMouseEvent9globalPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalPos", args)
  }

}

// button()
func (this *QMouseEvent) button(args ...interface{}) () {
  // button()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent6buttonEv
    // invoke: Qt::MouseButton button()
    C._ZNK11QMouseEvent6buttonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "button", args)
  }

}

// ~QMouseEvent()
func (this *QMouseEvent) FreeQMouseEvent(args ...interface{}) () {
  // ~QMouseEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMouseEventD0Ev
    // invoke: void ~QMouseEvent()
    C._ZN11QMouseEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "~QMouseEvent", args)
  }

}

// pos()
func (this *QMouseEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent3posEv
    // invoke: QPoint pos()
    C._ZNK11QMouseEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "pos", args)
  }

}

// buttons()
func (this *QMouseEvent) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C._ZNK11QMouseEvent7buttonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "buttons", args)
  }

}

// source()
func (this *QMouseEvent) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent6sourceEv
    // invoke: Qt::MouseEventSource source()
    C._ZNK11QMouseEvent6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "source", args)
  }

}

// flags()
func (this *QMouseEvent) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent5flagsEv
    // invoke: Qt::MouseEventFlags flags()
    C._ZNK11QMouseEvent5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "flags", args)
  }

}

// globalY()
func (this *QMouseEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7globalYEv
    // invoke: int globalY()
    C._ZNK11QMouseEvent7globalYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalY", args)
  }

}

// y()
func (this *QMouseEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent1yEv
    // invoke: int y()
    C._ZNK11QMouseEvent1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "y", args)
  }

}

// x()
func (this *QMouseEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent1xEv
    // invoke: int x()
    C._ZNK11QMouseEvent1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "x", args)
  }

}

// windowPos()
func (this *QMouseEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent9windowPosEv
    // invoke: const QPointF & windowPos()
    C._ZNK11QMouseEvent9windowPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "windowPos", args)
  }

}

// globalX()
func (this *QMouseEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMouseEvent7globalXEv
    // invoke: int globalX()
    C._ZNK11QMouseEvent7globalXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMouseEvent", "globalX", args)
  }

}

// url()
func (this *QFileOpenEvent) url(args ...interface{}) () {
  // url()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFileOpenEvent3urlEv
    // invoke: QUrl url()
    C._ZNK14QFileOpenEvent3urlEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "url", args)
  }

}

// ~QFileOpenEvent()
func (this *QFileOpenEvent) FreeQFileOpenEvent(args ...interface{}) () {
  // ~QFileOpenEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QFileOpenEventD0Ev
    // invoke: void ~QFileOpenEvent()
    C._ZN14QFileOpenEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "~QFileOpenEvent", args)
  }

}

// file()
func (this *QFileOpenEvent) file(args ...interface{}) () {
  // file()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFileOpenEvent4fileEv
    // invoke: QString file()
    C._ZNK14QFileOpenEvent4fileEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "file", args)
  }

}

// QFileOpenEvent(const class QString &)
func NewQFileOpenEvent(args ...interface{}) QFileOpenEvent {
  // QFileOpenEvent(const class QString &)
  // QFileOpenEvent(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QFileOpenEventC1ERK7QString
    // invoke: void QFileOpenEvent(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QFileOpenEventC2ERK7QString(qthis, arg0)
  case 1:
    // invoke: _ZN14QFileOpenEventC1ERK4QUrl
    // invoke: void QFileOpenEvent(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QFileOpenEventC2ERK4QUrl(qthis, arg0)
  default:
    qtrt.ErrorResolve("QFileOpenEvent", "QFileOpenEvent", args)
  }

  return QFileOpenEvent{}
}

// ~QToolBarChangeEvent()
func (this *QToolBarChangeEvent) FreeQToolBarChangeEvent(args ...interface{}) () {
  // ~QToolBarChangeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QToolBarChangeEventD0Ev
    // invoke: void ~QToolBarChangeEvent()
    C._ZN19QToolBarChangeEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "~QToolBarChangeEvent", args)
  }

}

// toggle()
func (this *QToolBarChangeEvent) toggle(args ...interface{}) () {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QToolBarChangeEvent6toggleEv
    // invoke: bool toggle()
    C._ZNK19QToolBarChangeEvent6toggleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "toggle", args)
  }

}

// QToolBarChangeEvent(_Bool)
func NewQToolBarChangeEvent(args ...interface{}) QToolBarChangeEvent {
  // QToolBarChangeEvent(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QToolBarChangeEventC1Eb
    // invoke: void QToolBarChangeEvent(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QToolBarChangeEventC2Eb(qthis, arg0)
  default:
    qtrt.ErrorResolve("QToolBarChangeEvent", "QToolBarChangeEvent", args)
  }

  return QToolBarChangeEvent{}
}

// xTilt()
func (this *QTabletEvent) xTilt(args ...interface{}) () {
  // xTilt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent5xTiltEv
    // invoke: int xTilt()
    C._ZNK12QTabletEvent5xTiltEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "xTilt", args)
  }

}

// pos()
func (this *QTabletEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent3posEv
    // invoke: QPoint pos()
    C._ZNK12QTabletEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "pos", args)
  }

}

// pointerType()
func (this *QTabletEvent) pointerType(args ...interface{}) () {
  // pointerType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent11pointerTypeEv
    // invoke: QTabletEvent::PointerType pointerType()
    C._ZNK12QTabletEvent11pointerTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "pointerType", args)
  }

}

// buttons()
func (this *QTabletEvent) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C._ZNK12QTabletEvent7buttonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "buttons", args)
  }

}

// hiResGlobalX()
func (this *QTabletEvent) hiResGlobalX(args ...interface{}) () {
  // hiResGlobalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent12hiResGlobalXEv
    // invoke: qreal hiResGlobalX()
    C._ZNK12QTabletEvent12hiResGlobalXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "hiResGlobalX", args)
  }

}

// hiResGlobalY()
func (this *QTabletEvent) hiResGlobalY(args ...interface{}) () {
  // hiResGlobalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent12hiResGlobalYEv
    // invoke: qreal hiResGlobalY()
    C._ZNK12QTabletEvent12hiResGlobalYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "hiResGlobalY", args)
  }

}

// ~QTabletEvent()
func (this *QTabletEvent) FreeQTabletEvent(args ...interface{}) () {
  // ~QTabletEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTabletEventD0Ev
    // invoke: void ~QTabletEvent()
    C._ZN12QTabletEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "~QTabletEvent", args)
  }

}

// uniqueId()
func (this *QTabletEvent) uniqueId(args ...interface{}) () {
  // uniqueId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8uniqueIdEv
    // invoke: qint64 uniqueId()
    C._ZNK12QTabletEvent8uniqueIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "uniqueId", args)
  }

}

// globalX()
func (this *QTabletEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7globalXEv
    // invoke: int globalX()
    C._ZNK12QTabletEvent7globalXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalX", args)
  }

}

// globalY()
func (this *QTabletEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent7globalYEv
    // invoke: int globalY()
    C._ZNK12QTabletEvent7globalYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalY", args)
  }

}

// yTilt()
func (this *QTabletEvent) yTilt(args ...interface{}) () {
  // yTilt()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent5yTiltEv
    // invoke: int yTilt()
    C._ZNK12QTabletEvent5yTiltEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "yTilt", args)
  }

}

// globalPosF()
func (this *QTabletEvent) globalPosF(args ...interface{}) () {
  // globalPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent10globalPosFEv
    // invoke: const QPointF & globalPosF()
    C._ZNK12QTabletEvent10globalPosFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalPosF", args)
  }

}

// pressure()
func (this *QTabletEvent) pressure(args ...interface{}) () {
  // pressure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8pressureEv
    // invoke: qreal pressure()
    C._ZNK12QTabletEvent8pressureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "pressure", args)
  }

}

// device()
func (this *QTabletEvent) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent6deviceEv
    // invoke: QTabletEvent::TabletDevice device()
    C._ZNK12QTabletEvent6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "device", args)
  }

}

// rotation()
func (this *QTabletEvent) rotation(args ...interface{}) () {
  // rotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent8rotationEv
    // invoke: qreal rotation()
    C._ZNK12QTabletEvent8rotationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "rotation", args)
  }

}

// globalPos()
func (this *QTabletEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent9globalPosEv
    // invoke: QPoint globalPos()
    C._ZNK12QTabletEvent9globalPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "globalPos", args)
  }

}

// posF()
func (this *QTabletEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent4posFEv
    // invoke: const QPointF & posF()
    C._ZNK12QTabletEvent4posFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "posF", args)
  }

}

// button()
func (this *QTabletEvent) button(args ...interface{}) () {
  // button()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent6buttonEv
    // invoke: Qt::MouseButton button()
    C._ZNK12QTabletEvent6buttonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "button", args)
  }

}

// y()
func (this *QTabletEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1yEv
    // invoke: int y()
    C._ZNK12QTabletEvent1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "y", args)
  }

}

// x()
func (this *QTabletEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1xEv
    // invoke: int x()
    C._ZNK12QTabletEvent1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "x", args)
  }

}

// z()
func (this *QTabletEvent) z(args ...interface{}) () {
  // z()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent1zEv
    // invoke: int z()
    C._ZNK12QTabletEvent1zEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "z", args)
  }

}

// tangentialPressure()
func (this *QTabletEvent) tangentialPressure(args ...interface{}) () {
  // tangentialPressure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTabletEvent18tangentialPressureEv
    // invoke: qreal tangentialPressure()
    C._ZNK12QTabletEvent18tangentialPressureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabletEvent", "tangentialPressure", args)
  }

}

// touchPointStates()
func (this *QTouchEvent) touchPointStates(args ...interface{}) () {
  // touchPointStates()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent16touchPointStatesEv
    // invoke: Qt::TouchPointStates touchPointStates()
    C._ZNK11QTouchEvent16touchPointStatesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "touchPointStates", args)
  }

}

// target()
func (this *QTouchEvent) target(args ...interface{}) () {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6targetEv
    // invoke: QObject * target()
    C._ZNK11QTouchEvent6targetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "target", args)
  }

}

// setTarget(class QObject *)
func (this *QTouchEvent) setTarget(args ...interface{}) () {
  // setTarget(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setTargetEP7QObject
    // invoke: void setTarget(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QTouchEvent9setTargetEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchEvent", "setTarget", args)
  }

}

// touchPoints()
func (this *QTouchEvent) touchPoints(args ...interface{}) () {
  // touchPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent11touchPointsEv
    // invoke: const QList<QTouchEvent::TouchPoint> & touchPoints()
    C._ZNK11QTouchEvent11touchPointsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "touchPoints", args)
  }

}

// ~QTouchEvent()
func (this *QTouchEvent) FreeQTouchEvent(args ...interface{}) () {
  // ~QTouchEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEventD0Ev
    // invoke: void ~QTouchEvent()
    C._ZN11QTouchEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "~QTouchEvent", args)
  }

}

// window()
func (this *QTouchEvent) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6windowEv
    // invoke: QWindow * window()
    C._ZNK11QTouchEvent6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "window", args)
  }

}

// setDevice(class QTouchDevice *)
func (this *QTouchEvent) setDevice(args ...interface{}) () {
  // setDevice(class QTouchDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTouchDevice{}) // "QTouchDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setDeviceEP12QTouchDevice
    // invoke: void setDevice(class QTouchDevice *)
    var arg0 = args[0].(QTouchDevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QTouchEvent9setDeviceEP12QTouchDevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchEvent", "setDevice", args)
  }

}

// device()
func (this *QTouchEvent) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTouchEvent6deviceEv
    // invoke: QTouchDevice * device()
    C._ZNK11QTouchEvent6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchEvent", "device", args)
  }

}

// setWindow(class QWindow *)
func (this *QTouchEvent) setWindow(args ...interface{}) () {
  // setWindow(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTouchEvent9setWindowEP7QWindow
    // invoke: void setWindow(class QWindow *)
    var arg0 = args[0].(QWindow).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QTouchEvent9setWindowEP7QWindow(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchEvent", "setWindow", args)
  }

}

// screen()
func (this *QScreenOrientationChangeEvent) screen(args ...interface{}) () {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QScreenOrientationChangeEvent6screenEv
    // invoke: QScreen * screen()
    C._ZNK29QScreenOrientationChangeEvent6screenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "screen", args)
  }

}

// orientation()
func (this *QScreenOrientationChangeEvent) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QScreenOrientationChangeEvent11orientationEv
    // invoke: Qt::ScreenOrientation orientation()
    C._ZNK29QScreenOrientationChangeEvent11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "orientation", args)
  }

}

// ~QScreenOrientationChangeEvent()
func (this *QScreenOrientationChangeEvent) FreeQScreenOrientationChangeEvent(args ...interface{}) () {
  // ~QScreenOrientationChangeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN29QScreenOrientationChangeEventD0Ev
    // invoke: void ~QScreenOrientationChangeEvent()
    C._ZN29QScreenOrientationChangeEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreenOrientationChangeEvent", "~QScreenOrientationChangeEvent", args)
  }

}

// ~QIconDragEvent()
func (this *QIconDragEvent) FreeQIconDragEvent(args ...interface{}) () {
  // ~QIconDragEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QIconDragEventD0Ev
    // invoke: void ~QIconDragEvent()
    C._ZN14QIconDragEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIconDragEvent", "~QIconDragEvent", args)
  }

}

// QIconDragEvent()
func NewQIconDragEvent(args ...interface{}) QIconDragEvent {
  // QIconDragEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QIconDragEventC1Ev
    // invoke: void QIconDragEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QIconDragEventC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QIconDragEvent", "QIconDragEvent", args)
  }

  return QIconDragEvent{}
}

// QCloseEvent()
func NewQCloseEvent(args ...interface{}) QCloseEvent {
  // QCloseEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QCloseEventC1Ev
    // invoke: void QCloseEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QCloseEventC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QCloseEvent", "QCloseEvent", args)
  }

  return QCloseEvent{}
}

// ~QCloseEvent()
func (this *QCloseEvent) FreeQCloseEvent(args ...interface{}) () {
  // ~QCloseEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QCloseEventD0Ev
    // invoke: void ~QCloseEvent()
    C._ZN11QCloseEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCloseEvent", "~QCloseEvent", args)
  }

}

// ~QDragEnterEvent()
func (this *QDragEnterEvent) FreeQDragEnterEvent(args ...interface{}) () {
  // ~QDragEnterEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QDragEnterEventD0Ev
    // invoke: void ~QDragEnterEvent()
    C._ZN15QDragEnterEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDragEnterEvent", "~QDragEnterEvent", args)
  }

}

// pixelDelta()
func (this *QWheelEvent) pixelDelta(args ...interface{}) () {
  // pixelDelta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10pixelDeltaEv
    // invoke: QPoint pixelDelta()
    C._ZNK11QWheelEvent10pixelDeltaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "pixelDelta", args)
  }

}

// globalPosF()
func (this *QWheelEvent) globalPosF(args ...interface{}) () {
  // globalPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10globalPosFEv
    // invoke: const QPointF & globalPosF()
    C._ZNK11QWheelEvent10globalPosFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalPosF", args)
  }

}

// ~QWheelEvent()
func (this *QWheelEvent) FreeQWheelEvent(args ...interface{}) () {
  // ~QWheelEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWheelEventD0Ev
    // invoke: void ~QWheelEvent()
    C._ZN11QWheelEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "~QWheelEvent", args)
  }

}

// orientation()
func (this *QWheelEvent) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent11orientationEv
    // invoke: Qt::Orientation orientation()
    C._ZNK11QWheelEvent11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "orientation", args)
  }

}

// posF()
func (this *QWheelEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent4posFEv
    // invoke: const QPointF & posF()
    C._ZNK11QWheelEvent4posFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "posF", args)
  }

}

// angleDelta()
func (this *QWheelEvent) angleDelta(args ...interface{}) () {
  // angleDelta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent10angleDeltaEv
    // invoke: QPoint angleDelta()
    C._ZNK11QWheelEvent10angleDeltaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "angleDelta", args)
  }

}

// pos()
func (this *QWheelEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent3posEv
    // invoke: QPoint pos()
    C._ZNK11QWheelEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "pos", args)
  }

}

// y()
func (this *QWheelEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent1yEv
    // invoke: int y()
    C._ZNK11QWheelEvent1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "y", args)
  }

}

// buttons()
func (this *QWheelEvent) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7buttonsEv
    // invoke: Qt::MouseButtons buttons()
    C._ZNK11QWheelEvent7buttonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "buttons", args)
  }

}

// source()
func (this *QWheelEvent) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent6sourceEv
    // invoke: Qt::MouseEventSource source()
    C._ZNK11QWheelEvent6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "source", args)
  }

}

// delta()
func (this *QWheelEvent) delta(args ...interface{}) () {
  // delta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent5deltaEv
    // invoke: int delta()
    C._ZNK11QWheelEvent5deltaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "delta", args)
  }

}

// phase()
func (this *QWheelEvent) phase(args ...interface{}) () {
  // phase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent5phaseEv
    // invoke: Qt::ScrollPhase phase()
    C._ZNK11QWheelEvent5phaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "phase", args)
  }

}

// globalY()
func (this *QWheelEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7globalYEv
    // invoke: int globalY()
    C._ZNK11QWheelEvent7globalYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalY", args)
  }

}

// globalX()
func (this *QWheelEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent7globalXEv
    // invoke: int globalX()
    C._ZNK11QWheelEvent7globalXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalX", args)
  }

}

// x()
func (this *QWheelEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent1xEv
    // invoke: int x()
    C._ZNK11QWheelEvent1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "x", args)
  }

}

// globalPos()
func (this *QWheelEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWheelEvent9globalPosEv
    // invoke: QPoint globalPos()
    C._ZNK11QWheelEvent9globalPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWheelEvent", "globalPos", args)
  }

}

// contentPos()
func (this *QScrollEvent) contentPos(args ...interface{}) () {
  // contentPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent10contentPosEv
    // invoke: QPointF contentPos()
    C._ZNK12QScrollEvent10contentPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollEvent", "contentPos", args)
  }

}

// overshootDistance()
func (this *QScrollEvent) overshootDistance(args ...interface{}) () {
  // overshootDistance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent17overshootDistanceEv
    // invoke: QPointF overshootDistance()
    C._ZNK12QScrollEvent17overshootDistanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollEvent", "overshootDistance", args)
  }

}

// ~QScrollEvent()
func (this *QScrollEvent) FreeQScrollEvent(args ...interface{}) () {
  // ~QScrollEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QScrollEventD0Ev
    // invoke: void ~QScrollEvent()
    C._ZN12QScrollEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollEvent", "~QScrollEvent", args)
  }

}

// scrollState()
func (this *QScrollEvent) scrollState(args ...interface{}) () {
  // scrollState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QScrollEvent11scrollStateEv
    // invoke: QScrollEvent::ScrollState scrollState()
    C._ZNK12QScrollEvent11scrollStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollEvent", "scrollState", args)
  }

}

// posF()
func (this *QHoverEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent4posFEv
    // invoke: const QPointF & posF()
    C._ZNK11QHoverEvent4posFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHoverEvent", "posF", args)
  }

}

// pos()
func (this *QHoverEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent3posEv
    // invoke: QPoint pos()
    C._ZNK11QHoverEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHoverEvent", "pos", args)
  }

}

// oldPos()
func (this *QHoverEvent) oldPos(args ...interface{}) () {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent6oldPosEv
    // invoke: QPoint oldPos()
    C._ZNK11QHoverEvent6oldPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHoverEvent", "oldPos", args)
  }

}

// oldPosF()
func (this *QHoverEvent) oldPosF(args ...interface{}) () {
  // oldPosF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHoverEvent7oldPosFEv
    // invoke: const QPointF & oldPosF()
    C._ZNK11QHoverEvent7oldPosFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHoverEvent", "oldPosF", args)
  }

}

// ~QHoverEvent()
func (this *QHoverEvent) FreeQHoverEvent(args ...interface{}) () {
  // ~QHoverEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHoverEventD0Ev
    // invoke: void ~QHoverEvent()
    C._ZN11QHoverEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHoverEvent", "~QHoverEvent", args)
  }

}

// ignore()
func (this *QDragMoveEvent) ignore(args ...interface{}) () {
  // ignore()
  // ignore(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEvent6ignoreEv
    // invoke: void ignore()
    C._ZN14QDragMoveEvent6ignoreEv(this.qclsinst)
  case 1:
    // invoke: _ZN14QDragMoveEvent6ignoreERK5QRect
    // invoke: void ignore(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QDragMoveEvent6ignoreERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "ignore", args)
  }

}

// ~QDragMoveEvent()
func (this *QDragMoveEvent) FreeQDragMoveEvent(args ...interface{}) () {
  // ~QDragMoveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEventD0Ev
    // invoke: void ~QDragMoveEvent()
    C._ZN14QDragMoveEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "~QDragMoveEvent", args)
  }

}

// answerRect()
func (this *QDragMoveEvent) answerRect(args ...interface{}) () {
  // answerRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDragMoveEvent10answerRectEv
    // invoke: QRect answerRect()
    C._ZNK14QDragMoveEvent10answerRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "answerRect", args)
  }

}

// accept(const class QRect &)
func (this *QDragMoveEvent) accept(args ...interface{}) () {
  // accept(const class QRect &)
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDragMoveEvent6acceptERK5QRect
    // invoke: void accept(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QDragMoveEvent6acceptERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN14QDragMoveEvent6acceptEv
    // invoke: void accept()
    C._ZN14QDragMoveEvent6acceptEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDragMoveEvent", "accept", args)
  }

}

// ~QShowEvent()
func (this *QShowEvent) FreeQShowEvent(args ...interface{}) () {
  // ~QShowEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QShowEventD0Ev
    // invoke: void ~QShowEvent()
    C._ZN10QShowEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShowEvent", "~QShowEvent", args)
  }

}

// QShowEvent()
func NewQShowEvent(args ...interface{}) QShowEvent {
  // QShowEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QShowEventC1Ev
    // invoke: void QShowEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QShowEventC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QShowEvent", "QShowEvent", args)
  }

  return QShowEvent{}
}

// surfaceEventType()
func (this *QPlatformSurfaceEvent) surfaceEventType(args ...interface{}) () {
  // surfaceEventType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv
    // invoke: QPlatformSurfaceEvent::SurfaceEventType surfaceEventType()
    C._ZNK21QPlatformSurfaceEvent16surfaceEventTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlatformSurfaceEvent", "surfaceEventType", args)
  }

}

// ~QPlatformSurfaceEvent()
func (this *QPlatformSurfaceEvent) FreeQPlatformSurfaceEvent(args ...interface{}) () {
  // ~QPlatformSurfaceEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QPlatformSurfaceEventD0Ev
    // invoke: void ~QPlatformSurfaceEvent()
    C._ZN21QPlatformSurfaceEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlatformSurfaceEvent", "~QPlatformSurfaceEvent", args)
  }

}

// QPaintEvent(const class QRect &)
func NewQPaintEvent(args ...interface{}) QPaintEvent {
  // QPaintEvent(const class QRect &)
  // QPaintEvent(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPaintEventC1ERK5QRect
    // invoke: void QPaintEvent(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QPaintEventC2ERK5QRect(qthis, arg0)
  case 1:
    // invoke: _ZN11QPaintEventC1ERK7QRegion
    // invoke: void QPaintEvent(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QPaintEventC2ERK7QRegion(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPaintEvent", "QPaintEvent", args)
  }

  return QPaintEvent{}
}

// ~QPaintEvent()
func (this *QPaintEvent) FreeQPaintEvent(args ...interface{}) () {
  // ~QPaintEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPaintEventD0Ev
    // invoke: void ~QPaintEvent()
    C._ZN11QPaintEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEvent", "~QPaintEvent", args)
  }

}

// region()
func (this *QPaintEvent) region(args ...interface{}) () {
  // region()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPaintEvent6regionEv
    // invoke: const QRegion & region()
    C._ZNK11QPaintEvent6regionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEvent", "region", args)
  }

}

// rect()
func (this *QPaintEvent) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPaintEvent4rectEv
    // invoke: const QRect & rect()
    C._ZNK11QPaintEvent4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintEvent", "rect", args)
  }

}

// reason()
func (this *QFocusEvent) reason(args ...interface{}) () {
  // reason()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent6reasonEv
    // invoke: Qt::FocusReason reason()
    C._ZNK11QFocusEvent6reasonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFocusEvent", "reason", args)
  }

}

// ~QFocusEvent()
func (this *QFocusEvent) FreeQFocusEvent(args ...interface{}) () {
  // ~QFocusEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFocusEventD0Ev
    // invoke: void ~QFocusEvent()
    C._ZN11QFocusEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFocusEvent", "~QFocusEvent", args)
  }

}

// gotFocus()
func (this *QFocusEvent) gotFocus(args ...interface{}) () {
  // gotFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent8gotFocusEv
    // invoke: bool gotFocus()
    C._ZNK11QFocusEvent8gotFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFocusEvent", "gotFocus", args)
  }

}

// lostFocus()
func (this *QFocusEvent) lostFocus(args ...interface{}) () {
  // lostFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusEvent9lostFocusEv
    // invoke: bool lostFocus()
    C._ZNK11QFocusEvent9lostFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFocusEvent", "lostFocus", args)
  }

}

// localPos()
func (this *QNativeGestureEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent8localPosEv
    // invoke: const QPointF & localPos()
    C._ZNK19QNativeGestureEvent8localPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "localPos", args)
  }

}

// screenPos()
func (this *QNativeGestureEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9screenPosEv
    // invoke: const QPointF & screenPos()
    C._ZNK19QNativeGestureEvent9screenPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "screenPos", args)
  }

}

// globalPos()
func (this *QNativeGestureEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9globalPosEv
    // invoke: const QPoint globalPos()
    C._ZNK19QNativeGestureEvent9globalPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "globalPos", args)
  }

}

// pos()
func (this *QNativeGestureEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent3posEv
    // invoke: const QPoint pos()
    C._ZNK19QNativeGestureEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "pos", args)
  }

}

// value()
func (this *QNativeGestureEvent) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent5valueEv
    // invoke: qreal value()
    C._ZNK19QNativeGestureEvent5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "value", args)
  }

}

// gestureType()
func (this *QNativeGestureEvent) gestureType(args ...interface{}) () {
  // gestureType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent11gestureTypeEv
    // invoke: Qt::NativeGestureType gestureType()
    C._ZNK19QNativeGestureEvent11gestureTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "gestureType", args)
  }

}

// windowPos()
func (this *QNativeGestureEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QNativeGestureEvent9windowPosEv
    // invoke: const QPointF & windowPos()
    C._ZNK19QNativeGestureEvent9windowPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNativeGestureEvent", "windowPos", args)
  }

}

// ~QResizeEvent()
func (this *QResizeEvent) FreeQResizeEvent(args ...interface{}) () {
  // ~QResizeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QResizeEventD0Ev
    // invoke: void ~QResizeEvent()
    C._ZN12QResizeEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResizeEvent", "~QResizeEvent", args)
  }

}

// oldSize()
func (this *QResizeEvent) oldSize(args ...interface{}) () {
  // oldSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QResizeEvent7oldSizeEv
    // invoke: const QSize & oldSize()
    C._ZNK12QResizeEvent7oldSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResizeEvent", "oldSize", args)
  }

}

// QResizeEvent(const class QSize &, const class QSize &)
func NewQResizeEvent(args ...interface{}) QResizeEvent {
  // QResizeEvent(const class QSize &, const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[0][1] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QResizeEventC1ERK5QSizeS2_
    // invoke: void QResizeEvent(const class QSize &, const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSize).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QResizeEventC2ERK5QSizeS2_(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QResizeEvent", "QResizeEvent", args)
  }

  return QResizeEvent{}
}

// size()
func (this *QResizeEvent) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QResizeEvent4sizeEv
    // invoke: const QSize & size()
    C._ZNK12QResizeEvent4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QResizeEvent", "size", args)
  }

}

// tip()
func (this *QStatusTipEvent) tip(args ...interface{}) () {
  // tip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QStatusTipEvent3tipEv
    // invoke: QString tip()
    C._ZNK15QStatusTipEvent3tipEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "tip", args)
  }

}

// ~QStatusTipEvent()
func (this *QStatusTipEvent) FreeQStatusTipEvent(args ...interface{}) () {
  // ~QStatusTipEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QStatusTipEventD0Ev
    // invoke: void ~QStatusTipEvent()
    C._ZN15QStatusTipEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "~QStatusTipEvent", args)
  }

}

// QStatusTipEvent(const class QString &)
func NewQStatusTipEvent(args ...interface{}) QStatusTipEvent {
  // QStatusTipEvent(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QStatusTipEventC1ERK7QString
    // invoke: void QStatusTipEvent(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QStatusTipEventC2ERK7QString(qthis, arg0)
  default:
    qtrt.ErrorResolve("QStatusTipEvent", "QStatusTipEvent", args)
  }

  return QStatusTipEvent{}
}

// localPos()
func (this *QEnterEvent) localPos(args ...interface{}) () {
  // localPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent8localPosEv
    // invoke: const QPointF & localPos()
    C._ZNK11QEnterEvent8localPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "localPos", args)
  }

}

// screenPos()
func (this *QEnterEvent) screenPos(args ...interface{}) () {
  // screenPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9screenPosEv
    // invoke: const QPointF & screenPos()
    C._ZNK11QEnterEvent9screenPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "screenPos", args)
  }

}

// globalPos()
func (this *QEnterEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9globalPosEv
    // invoke: QPoint globalPos()
    C._ZNK11QEnterEvent9globalPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalPos", args)
  }

}

// QEnterEvent(const class QPointF &, const class QPointF &, const class QPointF &)
func NewQEnterEvent(args ...interface{}) QEnterEvent {
  // QEnterEvent(const class QPointF &, const class QPointF &, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QEnterEventC1ERK7QPointFS2_S2_
    // invoke: void QEnterEvent(const class QPointF &, const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QEnterEventC2ERK7QPointFS2_S2_(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QEnterEvent", "QEnterEvent", args)
  }

  return QEnterEvent{}
}

// pos()
func (this *QEnterEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent3posEv
    // invoke: QPoint pos()
    C._ZNK11QEnterEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "pos", args)
  }

}

// globalX()
func (this *QEnterEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent7globalXEv
    // invoke: int globalX()
    C._ZNK11QEnterEvent7globalXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalX", args)
  }

}

// ~QEnterEvent()
func (this *QEnterEvent) FreeQEnterEvent(args ...interface{}) () {
  // ~QEnterEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QEnterEventD0Ev
    // invoke: void ~QEnterEvent()
    C._ZN11QEnterEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "~QEnterEvent", args)
  }

}

// globalY()
func (this *QEnterEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent7globalYEv
    // invoke: int globalY()
    C._ZNK11QEnterEvent7globalYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "globalY", args)
  }

}

// y()
func (this *QEnterEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent1yEv
    // invoke: int y()
    C._ZNK11QEnterEvent1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "y", args)
  }

}

// x()
func (this *QEnterEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent1xEv
    // invoke: int x()
    C._ZNK11QEnterEvent1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "x", args)
  }

}

// windowPos()
func (this *QEnterEvent) windowPos(args ...interface{}) () {
  // windowPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QEnterEvent9windowPosEv
    // invoke: const QPointF & windowPos()
    C._ZNK11QEnterEvent9windowPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEnterEvent", "windowPos", args)
  }

}

// QMoveEvent(const class QPoint &, const class QPoint &)
func NewQMoveEvent(args ...interface{}) QMoveEvent {
  // QMoveEvent(const class QPoint &, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMoveEventC1ERK6QPointS2_
    // invoke: void QMoveEvent(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QMoveEventC2ERK6QPointS2_(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMoveEvent", "QMoveEvent", args)
  }

  return QMoveEvent{}
}

// ~QMoveEvent()
func (this *QMoveEvent) FreeQMoveEvent(args ...interface{}) () {
  // ~QMoveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QMoveEventD0Ev
    // invoke: void ~QMoveEvent()
    C._ZN10QMoveEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMoveEvent", "~QMoveEvent", args)
  }

}

// oldPos()
func (this *QMoveEvent) oldPos(args ...interface{}) () {
  // oldPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMoveEvent6oldPosEv
    // invoke: const QPoint & oldPos()
    C._ZNK10QMoveEvent6oldPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMoveEvent", "oldPos", args)
  }

}

// pos()
func (this *QMoveEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QMoveEvent3posEv
    // invoke: const QPoint & pos()
    C._ZNK10QMoveEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMoveEvent", "pos", args)
  }

}

// QHideEvent()
func NewQHideEvent(args ...interface{}) QHideEvent {
  // QHideEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QHideEventC1Ev
    // invoke: void QHideEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QHideEventC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QHideEvent", "QHideEvent", args)
  }

  return QHideEvent{}
}

// ~QHideEvent()
func (this *QHideEvent) FreeQHideEvent(args ...interface{}) () {
  // ~QHideEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QHideEventD0Ev
    // invoke: void ~QHideEvent()
    C._ZN10QHideEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QHideEvent", "~QHideEvent", args)
  }

}

// ~QDragLeaveEvent()
func (this *QDragLeaveEvent) FreeQDragLeaveEvent(args ...interface{}) () {
  // ~QDragLeaveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QDragLeaveEventD0Ev
    // invoke: void ~QDragLeaveEvent()
    C._ZN15QDragLeaveEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDragLeaveEvent", "~QDragLeaveEvent", args)
  }

}

// QDragLeaveEvent()
func NewQDragLeaveEvent(args ...interface{}) QDragLeaveEvent {
  // QDragLeaveEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QDragLeaveEventC1Ev
    // invoke: void QDragLeaveEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QDragLeaveEventC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QDragLeaveEvent", "QDragLeaveEvent", args)
  }

  return QDragLeaveEvent{}
}

// mimeData()
func (this *QDropEvent) mimeData(args ...interface{}) () {
  // mimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent8mimeDataEv
    // invoke: const QMimeData * mimeData()
    C._ZNK10QDropEvent8mimeDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "mimeData", args)
  }

}

// acceptProposedAction()
func (this *QDropEvent) acceptProposedAction(args ...interface{}) () {
  // acceptProposedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QDropEvent20acceptProposedActionEv
    // invoke: void acceptProposedAction()
    C._ZN10QDropEvent20acceptProposedActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "acceptProposedAction", args)
  }

}

// possibleActions()
func (this *QDropEvent) possibleActions(args ...interface{}) () {
  // possibleActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent15possibleActionsEv
    // invoke: Qt::DropActions possibleActions()
    C._ZNK10QDropEvent15possibleActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "possibleActions", args)
  }

}

// posF()
func (this *QDropEvent) posF(args ...interface{}) () {
  // posF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent4posFEv
    // invoke: const QPointF & posF()
    C._ZNK10QDropEvent4posFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "posF", args)
  }

}

// pos()
func (this *QDropEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent3posEv
    // invoke: QPoint pos()
    C._ZNK10QDropEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "pos", args)
  }

}

// source()
func (this *QDropEvent) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent6sourceEv
    // invoke: QObject * source()
    C._ZNK10QDropEvent6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "source", args)
  }

}

// proposedAction()
func (this *QDropEvent) proposedAction(args ...interface{}) () {
  // proposedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent14proposedActionEv
    // invoke: Qt::DropAction proposedAction()
    C._ZNK10QDropEvent14proposedActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "proposedAction", args)
  }

}

// dropAction()
func (this *QDropEvent) dropAction(args ...interface{}) () {
  // dropAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent10dropActionEv
    // invoke: Qt::DropAction dropAction()
    C._ZNK10QDropEvent10dropActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "dropAction", args)
  }

}

// keyboardModifiers()
func (this *QDropEvent) keyboardModifiers(args ...interface{}) () {
  // keyboardModifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent17keyboardModifiersEv
    // invoke: Qt::KeyboardModifiers keyboardModifiers()
    C._ZNK10QDropEvent17keyboardModifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "keyboardModifiers", args)
  }

}

// ~QDropEvent()
func (this *QDropEvent) FreeQDropEvent(args ...interface{}) () {
  // ~QDropEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QDropEventD0Ev
    // invoke: void ~QDropEvent()
    C._ZN10QDropEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "~QDropEvent", args)
  }

}

// mouseButtons()
func (this *QDropEvent) mouseButtons(args ...interface{}) () {
  // mouseButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QDropEvent12mouseButtonsEv
    // invoke: Qt::MouseButtons mouseButtons()
    C._ZNK10QDropEvent12mouseButtonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDropEvent", "mouseButtons", args)
  }

}

// modifiers()
func (this *QInputEvent) modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QInputEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C._ZNK11QInputEvent9modifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputEvent", "modifiers", args)
  }

}

// timestamp()
func (this *QInputEvent) timestamp(args ...interface{}) () {
  // timestamp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QInputEvent9timestampEv
    // invoke: ulong timestamp()
    C._ZNK11QInputEvent9timestampEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputEvent", "timestamp", args)
  }

}

// setTimestamp(ulong)
func (this *QInputEvent) setTimestamp(args ...interface{}) () {
  // setTimestamp(ulong)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "ulong"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QInputEvent12setTimestampEm
    // invoke: void setTimestamp(ulong)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QInputEvent12setTimestampEm(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputEvent", "setTimestamp", args)
  }

}

// ~QInputEvent()
func (this *QInputEvent) FreeQInputEvent(args ...interface{}) () {
  // ~QInputEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QInputEventD0Ev
    // invoke: void ~QInputEvent()
    C._ZN11QInputEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputEvent", "~QInputEvent", args)
  }

}

// applicationState()
func (this *QApplicationStateChangeEvent) applicationState(args ...interface{}) () {
  // applicationState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK28QApplicationStateChangeEvent16applicationStateEv
    // invoke: Qt::ApplicationState applicationState()
    C._ZNK28QApplicationStateChangeEvent16applicationStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QApplicationStateChangeEvent", "applicationState", args)
  }

}

// count()
func (this *QKeyEvent) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent5countEv
    // invoke: int count()
    C._ZNK9QKeyEvent5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "count", args)
  }

}

// modifiers()
func (this *QKeyEvent) modifiers(args ...interface{}) () {
  // modifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent9modifiersEv
    // invoke: Qt::KeyboardModifiers modifiers()
    C._ZNK9QKeyEvent9modifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "modifiers", args)
  }

}

// nativeModifiers()
func (this *QKeyEvent) nativeModifiers(args ...interface{}) () {
  // nativeModifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent15nativeModifiersEv
    // invoke: quint32 nativeModifiers()
    C._ZNK9QKeyEvent15nativeModifiersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeModifiers", args)
  }

}

// text()
func (this *QKeyEvent) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent4textEv
    // invoke: QString text()
    C._ZNK9QKeyEvent4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "text", args)
  }

}

// ~QKeyEvent()
func (this *QKeyEvent) FreeQKeyEvent(args ...interface{}) () {
  // ~QKeyEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QKeyEventD0Ev
    // invoke: void ~QKeyEvent()
    C._ZN9QKeyEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "~QKeyEvent", args)
  }

}

// nativeScanCode()
func (this *QKeyEvent) nativeScanCode(args ...interface{}) () {
  // nativeScanCode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent14nativeScanCodeEv
    // invoke: quint32 nativeScanCode()
    C._ZNK9QKeyEvent14nativeScanCodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeScanCode", args)
  }

}

// isAutoRepeat()
func (this *QKeyEvent) isAutoRepeat(args ...interface{}) () {
  // isAutoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent12isAutoRepeatEv
    // invoke: bool isAutoRepeat()
    C._ZNK9QKeyEvent12isAutoRepeatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "isAutoRepeat", args)
  }

}

// key()
func (this *QKeyEvent) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent3keyEv
    // invoke: int key()
    C._ZNK9QKeyEvent3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "key", args)
  }

}

// nativeVirtualKey()
func (this *QKeyEvent) nativeVirtualKey(args ...interface{}) () {
  // nativeVirtualKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QKeyEvent16nativeVirtualKeyEv
    // invoke: quint32 nativeVirtualKey()
    C._ZNK9QKeyEvent16nativeVirtualKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEvent", "nativeVirtualKey", args)
  }

}

// globalPos()
func (this *QContextMenuEvent) globalPos(args ...interface{}) () {
  // globalPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent9globalPosEv
    // invoke: const QPoint & globalPos()
    C._ZNK17QContextMenuEvent9globalPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalPos", args)
  }

}

// pos()
func (this *QContextMenuEvent) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent3posEv
    // invoke: const QPoint & pos()
    C._ZNK17QContextMenuEvent3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "pos", args)
  }

}

// y()
func (this *QContextMenuEvent) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent1yEv
    // invoke: int y()
    C._ZNK17QContextMenuEvent1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "y", args)
  }

}

// reason()
func (this *QContextMenuEvent) reason(args ...interface{}) () {
  // reason()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent6reasonEv
    // invoke: QContextMenuEvent::Reason reason()
    C._ZNK17QContextMenuEvent6reasonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "reason", args)
  }

}

// x()
func (this *QContextMenuEvent) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent1xEv
    // invoke: int x()
    C._ZNK17QContextMenuEvent1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "x", args)
  }

}

// globalX()
func (this *QContextMenuEvent) globalX(args ...interface{}) () {
  // globalX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent7globalXEv
    // invoke: int globalX()
    C._ZNK17QContextMenuEvent7globalXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalX", args)
  }

}

// globalY()
func (this *QContextMenuEvent) globalY(args ...interface{}) () {
  // globalY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QContextMenuEvent7globalYEv
    // invoke: int globalY()
    C._ZNK17QContextMenuEvent7globalYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "globalY", args)
  }

}

// ~QContextMenuEvent()
func (this *QContextMenuEvent) FreeQContextMenuEvent(args ...interface{}) () {
  // ~QContextMenuEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QContextMenuEventD0Ev
    // invoke: void ~QContextMenuEvent()
    C._ZN17QContextMenuEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QContextMenuEvent", "~QContextMenuEvent", args)
  }

}

// setContentPosRange(const class QRectF &)
func (this *QScrollPrepareEvent) setContentPosRange(args ...interface{}) () {
  // setContentPosRange(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF
    // invoke: void setContentPosRange(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QScrollPrepareEvent18setContentPosRangeERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setContentPosRange", args)
  }

}

// contentPos()
func (this *QScrollPrepareEvent) contentPos(args ...interface{}) () {
  // contentPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent10contentPosEv
    // invoke: QPointF contentPos()
    C._ZNK19QScrollPrepareEvent10contentPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "contentPos", args)
  }

}

// setContentPos(const class QPointF &)
func (this *QScrollPrepareEvent) setContentPos(args ...interface{}) () {
  // setContentPos(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent13setContentPosERK7QPointF
    // invoke: void setContentPos(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QScrollPrepareEvent13setContentPosERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setContentPos", args)
  }

}

// ~QScrollPrepareEvent()
func (this *QScrollPrepareEvent) FreeQScrollPrepareEvent(args ...interface{}) () {
  // ~QScrollPrepareEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEventD0Ev
    // invoke: void ~QScrollPrepareEvent()
    C._ZN19QScrollPrepareEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "~QScrollPrepareEvent", args)
  }

}

// viewportSize()
func (this *QScrollPrepareEvent) viewportSize(args ...interface{}) () {
  // viewportSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent12viewportSizeEv
    // invoke: QSizeF viewportSize()
    C._ZNK19QScrollPrepareEvent12viewportSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "viewportSize", args)
  }

}

// setViewportSize(const class QSizeF &)
func (this *QScrollPrepareEvent) setViewportSize(args ...interface{}) () {
  // setViewportSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF
    // invoke: void setViewportSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QScrollPrepareEvent15setViewportSizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "setViewportSize", args)
  }

}

// QScrollPrepareEvent(const class QPointF &)
func NewQScrollPrepareEvent(args ...interface{}) QScrollPrepareEvent {
  // QScrollPrepareEvent(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollPrepareEventC1ERK7QPointF
    // invoke: void QScrollPrepareEvent(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QScrollPrepareEventC2ERK7QPointF(qthis, arg0)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "QScrollPrepareEvent", args)
  }

  return QScrollPrepareEvent{}
}

// contentPosRange()
func (this *QScrollPrepareEvent) contentPosRange(args ...interface{}) () {
  // contentPosRange()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent15contentPosRangeEv
    // invoke: QRectF contentPosRange()
    C._ZNK19QScrollPrepareEvent15contentPosRangeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "contentPosRange", args)
  }

}

// startPos()
func (this *QScrollPrepareEvent) startPos(args ...interface{}) () {
  // startPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QScrollPrepareEvent8startPosEv
    // invoke: QPointF startPos()
    C._ZNK19QScrollPrepareEvent8startPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScrollPrepareEvent", "startPos", args)
  }

}

// ~QShortcutEvent()
func (this *QShortcutEvent) FreeQShortcutEvent(args ...interface{}) () {
  // ~QShortcutEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QShortcutEventD0Ev
    // invoke: void ~QShortcutEvent()
    C._ZN14QShortcutEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcutEvent", "~QShortcutEvent", args)
  }

}

// isAmbiguous()
func (this *QShortcutEvent) isAmbiguous(args ...interface{}) () {
  // isAmbiguous()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent11isAmbiguousEv
    // invoke: bool isAmbiguous()
    C._ZNK14QShortcutEvent11isAmbiguousEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcutEvent", "isAmbiguous", args)
  }

}

// shortcutId()
func (this *QShortcutEvent) shortcutId(args ...interface{}) () {
  // shortcutId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent10shortcutIdEv
    // invoke: int shortcutId()
    C._ZNK14QShortcutEvent10shortcutIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcutEvent", "shortcutId", args)
  }

}

// key()
func (this *QShortcutEvent) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QShortcutEvent3keyEv
    // invoke: const QKeySequence & key()
    C._ZNK14QShortcutEvent3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcutEvent", "key", args)
  }

}

// QShortcutEvent(const class QKeySequence &, int, _Bool)
func NewQShortcutEvent(args ...interface{}) QShortcutEvent {
  // QShortcutEvent(const class QKeySequence &, int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QShortcutEventC1ERK12QKeySequenceib
    // invoke: void QShortcutEvent(const class QKeySequence &, int, _Bool)
    var arg0 = args[0].(QKeySequence).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QShortcutEventC2ERK12QKeySequenceib(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QShortcutEvent", "QShortcutEvent", args)
  }

  return QShortcutEvent{}
}

// isOverride()
func (this *QWindowStateChangeEvent) isOverride(args ...interface{}) () {
  // isOverride()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QWindowStateChangeEvent10isOverrideEv
    // invoke: bool isOverride()
    C._ZNK23QWindowStateChangeEvent10isOverrideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "isOverride", args)
  }

}

// oldState()
func (this *QWindowStateChangeEvent) oldState(args ...interface{}) () {
  // oldState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QWindowStateChangeEvent8oldStateEv
    // invoke: Qt::WindowStates oldState()
    C._ZNK23QWindowStateChangeEvent8oldStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "oldState", args)
  }

}

// ~QWindowStateChangeEvent()
func (this *QWindowStateChangeEvent) FreeQWindowStateChangeEvent(args ...interface{}) () {
  // ~QWindowStateChangeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QWindowStateChangeEventD0Ev
    // invoke: void ~QWindowStateChangeEvent()
    C._ZN23QWindowStateChangeEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWindowStateChangeEvent", "~QWindowStateChangeEvent", args)
  }

}

// ~QInputMethodQueryEvent()
func (this *QInputMethodQueryEvent) FreeQInputMethodQueryEvent(args ...interface{}) () {
  // ~QInputMethodQueryEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QInputMethodQueryEventD0Ev
    // invoke: void ~QInputMethodQueryEvent()
    C._ZN22QInputMethodQueryEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodQueryEvent", "~QInputMethodQueryEvent", args)
  }

}

// queries()
func (this *QInputMethodQueryEvent) queries(args ...interface{}) () {
  // queries()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QInputMethodQueryEvent7queriesEv
    // invoke: Qt::InputMethodQueries queries()
    C._ZNK22QInputMethodQueryEvent7queriesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethodQueryEvent", "queries", args)
  }

}

// <= body block end

