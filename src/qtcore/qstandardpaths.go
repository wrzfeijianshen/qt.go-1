//  header block begin
// /usr/include/qt/QtCore/qstandardpaths.h
// #include <qstandardpaths.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 61
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QStandardPaths struct {
	*qtrt.CObject
}

func (this *QStandardPaths) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQStandardPathsFromPointer(cthis unsafe.Pointer) *QStandardPaths {
	return &QStandardPaths{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstandardpaths.h:81
// index:0
// Public static
// QString writableLocation(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) WritableLocation(type_ int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths16writableLocationENS_16StandardLocationE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStandardPaths_WritableLocation(type_ int) {
	var nilthis *QStandardPaths
	nilthis.WritableLocation(type_)
}

// /usr/include/qt/QtCore/qstandardpaths.h:82
// index:0
// Public static
// QStringList standardLocations(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) StandardLocations(type_ int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths17standardLocationsENS_16StandardLocationE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStandardPaths_StandardLocations(type_ int) {
	var nilthis *QStandardPaths
	nilthis.StandardLocations(type_)
}

// /usr/include/qt/QtCore/qstandardpaths.h:94
// index:0
// Public static
// QString displayName(enum QStandardPaths::StandardLocation)
func (this *QStandardPaths) DisplayName(type_ int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths11displayNameENS_16StandardLocationE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStandardPaths_DisplayName(type_ int) {
	var nilthis *QStandardPaths
	nilthis.DisplayName(type_)
}

// /usr/include/qt/QtCore/qstandardpaths.h:97
// index:0
// Public static
// QString findExecutable(const class QString &, const class QStringList &)
func (this *QStandardPaths) FindExecutable(executableName *QString, paths *QStringList) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths14findExecutableERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, executableName, paths)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStandardPaths_FindExecutable(executableName *QString, paths *QStringList) {
	var nilthis *QStandardPaths
	nilthis.FindExecutable(executableName, paths)
}

// /usr/include/qt/QtCore/qstandardpaths.h:100
// index:0
// Public static
// void enableTestMode(_Bool)
func (this *QStandardPaths) EnableTestMode(testMode bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths14enableTestModeEb", ffiqt.FFI_TYPE_POINTER, testMode)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_EnableTestMode(testMode bool) {
	var nilthis *QStandardPaths
	nilthis.EnableTestMode(testMode)
}

// /usr/include/qt/QtCore/qstandardpaths.h:102
// index:0
// Public static
// void setTestModeEnabled(_Bool)
func (this *QStandardPaths) SetTestModeEnabled(testMode bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths18setTestModeEnabledEb", ffiqt.FFI_TYPE_POINTER, testMode)
	gopp.ErrPrint(err, rv)
}
func QStandardPaths_SetTestModeEnabled(testMode bool) {
	var nilthis *QStandardPaths
	nilthis.SetTestModeEnabled(testMode)
}

// /usr/include/qt/QtCore/qstandardpaths.h:103
// index:0
// Public static
// bool isTestModeEnabled()
func (this *QStandardPaths) IsTestModeEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStandardPaths17isTestModeEnabledEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStandardPaths_IsTestModeEnabled() {
	var nilthis *QStandardPaths
	nilthis.IsTestModeEnabled()
}

//  body block end