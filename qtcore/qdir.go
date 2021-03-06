package qtcore

// /usr/include/qt/QtCore/qdir.h
// #include <qdir.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 60
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QDir struct {
	*qtrt.CObject
}
type QDir_ITF interface {
	QDir_PTR() *QDir
}

func (ptr *QDir) QDir_PTR() *QDir { return ptr }

func (this *QDir) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDir) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDirFromPointer(cthis unsafe.Pointer) *QDir {
	return &QDir{&qtrt.CObject{cthis}}
}
func (*QDir) NewFromPointer(cthis unsafe.Pointer) *QDir {
	return NewQDirFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdir.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDir(const QString &)

/*
Constructs a QDir object that is a copy of the QDir object for directory dir.

See also operator=().
*/
func (*QDir) NewForInherit(path string) *QDir {
	return NewQDir(path)
}
func NewQDir(path string) *QDir {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDir(const QString &)

/*
Constructs a QDir object that is a copy of the QDir object for directory dir.

See also operator=().
*/
func (*QDir) NewForInheritp() *QDir {
	return NewQDirp()
}
func NewQDirp() *QDir {
	// arg: 0, const QString &=LValueReference, QString=Record, , Invalid
	var convArg0 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDir(const QString &, const QString &, QDir::SortFlags, QDir::Filters)

/*
Constructs a QDir object that is a copy of the QDir object for directory dir.

See also operator=().
*/
func (*QDir) NewForInherit1(path string, nameFilter string, sort int, filter int) *QDir {
	return NewQDir1(path, nameFilter, sort, filter)
}
func NewQDir1(path string, nameFilter string, sort int, filter int) *QDir {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(nameFilter)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirC2ERK7QStringS2_6QFlagsINS_8SortFlagEES3_INS_6FilterEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, sort, filter)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDir(const QString &, const QString &, QDir::SortFlags, QDir::Filters)

/*
Constructs a QDir object that is a copy of the QDir object for directory dir.

See also operator=().
*/
func (*QDir) NewForInherit1p(path string, nameFilter string) *QDir {
	return NewQDir1p(path, nameFilter)
}
func NewQDir1p(path string, nameFilter string) *QDir {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(nameFilter)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	// arg: 3, QDir::Filters=Typedef, QDir::Filters=Typedef, QFlags<QDir::Filter>, Unexposed
	filter := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirC2ERK7QStringS2_6QFlagsINS_8SortFlagEES3_INS_6FilterEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, sort, filter)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDir(const QString &, const QString &, QDir::SortFlags, QDir::Filters)

/*
Constructs a QDir object that is a copy of the QDir object for directory dir.

See also operator=().
*/
func (*QDir) NewForInherit1p1(path string, nameFilter string, sort int) *QDir {
	return NewQDir1p1(path, nameFilter, sort)
}
func NewQDir1p1(path string, nameFilter string, sort int) *QDir {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(nameFilter)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, QDir::Filters=Typedef, QDir::Filters=Typedef, QFlags<QDir::Filter>, Unexposed
	filter := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirC2ERK7QStringS2_6QFlagsINS_8SortFlagEES3_INS_6FilterEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, sort, filter)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDir)
	return gothis
}

// /usr/include/qt/QtCore/qdir.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDir()

/*

 */
func DeleteQDir(this *QDir) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDirD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdir.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QDir & operator=(const QDir &)

/*

 */
func (this *QDir) Operator_equal(arg0 QDir_ITF) *QDir {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDir_PTR() != nil {
		convArg0 = arg0.QDir_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDiraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:108
// index:1
// Public Visibility=Default Availability=Available
// [8] QDir & operator=(const QString &)

/*

 */
func (this *QDir) Operator_equal1(path string) *QDir {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDiraSERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:110
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QDir & operator=(QDir &&)

/*

 */
func (this *QDir) Operator_equal2(other unsafe.Pointer /*333*/) *QDir {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDiraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDir &)

/*
Swaps this QDir instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QDir) Swap(other QDir_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDir_PTR() != nil {
		convArg0 = other.QDir_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &)

/*
Sets the path of the directory to path. The path is cleaned of redundant ".", ".." and of multiple separators. No check is made to see whether a directory with this path actually exists; but you can check for yourself using exists().

The path can be either absolute or relative. Absolute paths begin with the directory separator "/" (optionally preceded by a drive specification under Windows). Relative file names begin with a directory name or a file name and specify a path relative to the current directory. An example of an absolute path is the string "/tmp/quartz", a relative path might look like "src/fatlib".

See also path(), absolutePath(), exists(), cleanPath(), dirName(), absoluteFilePath(), isRelative(), and makeAbsolute().
*/
func (this *QDir) SetPath(path string) {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir7setPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path() const

/*
Returns the path. This may contain symbolic links, but never contains redundant ".", ".." or multiple separators.

The returned path can be either absolute or relative (see setPath()).

See also setPath(), absolutePath(), exists(), cleanPath(), dirName(), absoluteFilePath(), toNativeSeparators(), and makeAbsolute().
*/
func (this *QDir) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absolutePath() const

/*
Returns the absolute path (a path that starts with "/" or with a drive specification), which may contain symbolic links, but never contains redundant ".", ".." or multiple separators.

See also setPath(), canonicalPath(), exists(), cleanPath(), dirName(), and absoluteFilePath().
*/
func (this *QDir) AbsolutePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir12absolutePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QString canonicalPath() const

/*
Returns the canonical path, i.e. a path without symbolic links or redundant "." or ".." elements.

On systems that do not have symbolic links this function will always return the same string that absolutePath() returns. If the canonical path does not exist (normally due to dangling symbolic links) canonicalPath() returns an empty string.

Example:


  QString bin = "/local/bin";         // where /local/bin is a symlink to /usr/bin
  QDir binDir(bin);
  QString canonicalBin = binDir.canonicalPath();
  // canonicalBin now equals "/usr/bin"

  QString ls = "/local/bin/ls";       // where ls is the executable "ls"
  QDir lsDir(ls);
  QString canonicalLs = lsDir.canonicalPath();
  // canonicalLS now equals "/usr/bin/ls".



See also path(), absolutePath(), exists(), cleanPath(), dirName(), and absoluteFilePath().
*/
func (this *QDir) CanonicalPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13canonicalPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addResourceSearchPath(const QString &)

/*

 */
func (this *QDir) AddResourceSearchPath(path string) {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir21addResourceSearchPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QDir_AddResourceSearchPath(path string) {
	var nilthis *QDir
	nilthis.AddResourceSearchPath(path)
}

// /usr/include/qt/QtCore/qdir.h:123
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setSearchPaths(const QString &, const QStringList &)

/*
Sets or replaces Qt's search paths for file names with the prefix prefix to searchPaths.

To specify a prefix for a file name, prepend the prefix followed by a single colon (e.g., "images:undo.png", "xmldocs:books.xml"). prefix can only contain letters or numbers (e.g., it cannot contain a colon, nor a slash).

Qt uses this search path to locate files with a known prefix. The search path entries are tested in order, starting with the first entry.


  QDir::setSearchPaths("icons", QStringList(QDir::homePath() + "/images"));
  QDir::setSearchPaths("docs", QStringList(":/embeddedDocuments"));
  ...
  QPixmap pixmap("icons:undo.png"); // will look for undo.png in QDir::homePath() + "/images"
  QFile file("docs:design.odf"); // will look in the :/embeddedDocuments resource path



File name prefix must be at least 2 characters long to avoid conflicts with Windows drive letters.

Search paths may contain paths to The Qt Resource System.

This function was introduced in  Qt 4.3.

See also searchPaths().
*/
func (this *QDir) SetSearchPaths(prefix string, searchPaths QStringList_ITF) {
	var tmpArg0 = NewQString5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if searchPaths != nil && searchPaths.QStringList_PTR() != nil {
		convArg1 = searchPaths.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir14setSearchPathsERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QDir_SetSearchPaths(prefix string, searchPaths QStringList_ITF) {
	var nilthis *QDir
	nilthis.SetSearchPaths(prefix, searchPaths)
}

// /usr/include/qt/QtCore/qdir.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addSearchPath(const QString &, const QString &)

/*
Adds path to the search path for prefix.

This function was introduced in  Qt 4.3.

See also setSearchPaths().
*/
func (this *QDir) AddSearchPath(prefix string, path string) {
	var tmpArg0 = NewQString5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(path)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir13addSearchPathERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QDir_AddSearchPath(prefix string, path string) {
	var nilthis *QDir
	nilthis.AddSearchPath(prefix, path)
}

// /usr/include/qt/QtCore/qdir.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList searchPaths(const QString &)

/*
Returns the search paths for prefix.

This function was introduced in  Qt 4.3.

See also setSearchPaths() and addSearchPath().
*/
func (this *QDir) SearchPaths(prefix string) *QStringList /*123*/ {
	var tmpArg0 = NewQString5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir11searchPathsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QDir_SearchPaths(prefix string) *QStringList /*123*/ {
	var nilthis *QDir
	rv := nilthis.SearchPaths(prefix)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dirName() const

/*
Returns the name of the directory; this is not the same as the path, e.g. a directory with the name "mail", might have the path "/var/spool/mail". If the directory has no name (e.g. it is the root directory) an empty string is returned.

No check is made to ensure that a directory with this name actually exists; but see exists().

See also path(), filePath(), absolutePath(), and absoluteFilePath().
*/
func (this *QDir) DirName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7dirNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath(const QString &) const

/*
Returns the path name of a file in the directory. Does not check if the file actually exists in the directory; but see exists(). If the QDir is relative the returned path name will also be relative. Redundant multiple separators or "." and ".." directories in fileName are not removed (see cleanPath()).

See also dirName(), absoluteFilePath(), isRelative(), and canonicalPath().
*/
func (this *QDir) FilePath(fileName string) string {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir8filePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absoluteFilePath(const QString &) const

/*
Returns the absolute path name of a file in the directory. Does not check if the file actually exists in the directory; but see exists(). Redundant multiple separators or "." and ".." directories in fileName are not removed (see cleanPath()).

See also relativeFilePath(), filePath(), and canonicalPath().
*/
func (this *QDir) AbsoluteFilePath(fileName string) string {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir16absoluteFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QString relativeFilePath(const QString &) const

/*
Returns the path to fileName relative to the directory.


  QDir dir("/home/bob");
  QString s;

  s = dir.relativeFilePath("images/file.jpg");     // s is "images/file.jpg"
  s = dir.relativeFilePath("/home/mary/file.txt"); // s is "../mary/file.txt"



See also absoluteFilePath(), filePath(), and canonicalPath().
*/
func (this *QDir) RelativeFilePath(fileName string) string {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir16relativeFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString toNativeSeparators(const QString &)

/*
Returns pathName with the '/' separators converted to separators that are appropriate for the underlying operating system.

On Windows, toNativeSeparators("c:/winnt/system32") returns "c:\winnt\system32".

The returned string may be the same as the argument on some operating systems, for example on Unix.

This function was introduced in  Qt 4.2.

See also fromNativeSeparators() and separator().
*/
func (this *QDir) ToNativeSeparators(pathName string) string {
	var tmpArg0 = NewQString5(pathName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir18toNativeSeparatorsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_ToNativeSeparators(pathName string) string {
	var nilthis *QDir
	rv := nilthis.ToNativeSeparators(pathName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:133
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromNativeSeparators(const QString &)

/*
Returns pathName using '/' as file separator. On Windows, for instance, fromNativeSeparators("c:\\winnt\\system32") returns "c:/winnt/system32".

The returned string may be the same as the argument on some operating systems, for example on Unix.

This function was introduced in  Qt 4.2.

See also toNativeSeparators() and separator().
*/
func (this *QDir) FromNativeSeparators(pathName string) string {
	var tmpArg0 = NewQString5(pathName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir20fromNativeSeparatorsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_FromNativeSeparators(pathName string) string {
	var nilthis *QDir
	rv := nilthis.FromNativeSeparators(pathName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cd(const QString &)

/*
Changes the QDir's directory to dirName.

Returns true if the new directory exists; otherwise returns false. Note that the logical cd() operation is not performed if the new directory does not exist.

Calling cd("..") is equivalent to calling cdUp().

See also cdUp(), isReadable(), exists(), and path().
*/
func (this *QDir) Cd(dirName string) bool {
	var tmpArg0 = NewQString5(dirName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir2cdERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cdUp()

/*
Changes directory by moving one directory up from the QDir's current directory.

Returns true if the new directory exists; otherwise returns false. Note that the logical cdUp() operation is not performed if the new directory does not exist.

See also cd(), isReadable(), exists(), and path().
*/
func (this *QDir) CdUp() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4cdUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList nameFilters() const

/*
Returns the string list set by setNameFilters()

See also setNameFilters().
*/
func (this *QDir) NameFilters() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir11nameFiltersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilters(const QStringList &)

/*
Sets the name filters used by entryList() and entryInfoList() to the list of filters specified by nameFilters.

Each name filter is a wildcard (globbing) filter that understands * and ? wildcards. (See QRegExp wildcard matching.)

For example, the following code sets three name filters on a QDir to ensure that only files with extensions typically used for C++ source files are listed:


      QStringList filters;
      filters << "*.cpp" << "*.cxx" << "*.cc";
      dir.setNameFilters(filters);



See also nameFilters() and setFilter().
*/
func (this *QDir) SetNameFilters(nameFilters QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir14setNameFiltersERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::Filters filter() const

/*
Returns the value set by setFilter()

See also setFilter().
*/
func (this *QDir) Filter() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6filterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdir.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilter(QDir::Filters)

/*
Sets the filter used by entryList() and entryInfoList() to filters. The filter is used to specify the kind of files that should be returned by entryList() and entryInfoList(). See QDir::Filter.

See also filter() and setNameFilters().
*/
func (this *QDir) SetFilter(filter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir9setFilterE6QFlagsINS_6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::SortFlags sorting() const

/*
Returns the value set by setSorting()

See also setSorting() and SortFlag.
*/
func (this *QDir) Sorting() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7sortingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdir.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSorting(QDir::SortFlags)

/*
Sets the sort order used by entryList() and entryInfoList().

The sort is specified by OR-ing values from the enum QDir::SortFlag.

See also sorting() and SortFlag.
*/
func (this *QDir) SetSorting(sort int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir10setSortingE6QFlagsINS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sort)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdir.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] uint count() const

/*
Returns the total number of directories and files in the directory.

Equivalent to entryList().count().

See also operator[]() and entryList().
*/
func (this *QDir) Count() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdir.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty(QDir::Filters) const

/*
Returns whether the directory is empty.

Equivalent to count() == 0 with filters QDir::AllEntries | QDir::NoDotAndDotDot, but faster as it just checks whether the directory contains at least one entry.

Note: Unless you set the filters flags to include QDir::NoDotAndDotDot (as the default value does), no directory is empty.

This function was introduced in  Qt 5.9.

See also count(), entryList(), and setFilter().
*/
func (this *QDir) IsEmpty(filters int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7isEmptyE6QFlagsINS_6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty(QDir::Filters) const

/*
Returns whether the directory is empty.

Equivalent to count() == 0 with filters QDir::AllEntries | QDir::NoDotAndDotDot, but faster as it just checks whether the directory contains at least one entry.

Note: Unless you set the filters flags to include QDir::NoDotAndDotDot (as the default value does), no directory is empty.

This function was introduced in  Qt 5.9.

See also count(), entryList(), and setFilter().
*/
func (this *QDir) IsEmptyp() bool {
	// arg: 0, QDir::Filters=Typedef, QDir::Filters=Typedef, QFlags<QDir::Filter>, Unexposed
	filters := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7isEmptyE6QFlagsINS_6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] QString operator[](int) const

/*

 */
func (this *QDir) Operator_get_index(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDirixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdir.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList nameFiltersFromString(const QString &)

/*

 */
func (this *QDir) NameFiltersFromString(nameFilter string) *QStringList /*123*/ {
	var tmpArg0 = NewQString5(nameFilter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir21nameFiltersFromStringERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QDir_NameFiltersFromString(nameFilter string) *QStringList /*123*/ {
	var nilthis *QDir
	rv := nilthis.NameFiltersFromString(nameFilter)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList entryList(QDir::Filters, QDir::SortFlags) const

/*
Returns a list of the names of all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryInfoList(), setNameFilters(), setSorting(), and setFilter().
*/
func (this *QDir) EntryList(filters int, sort int) *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir9entryListE6QFlagsINS_6FilterEES0_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList entryList(QDir::Filters, QDir::SortFlags) const

/*
Returns a list of the names of all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryInfoList(), setNameFilters(), setSorting(), and setFilter().
*/
func (this *QDir) EntryListp() *QStringList /*123*/ {
	// arg: 0, QDir::Filters=Typedef, QDir::Filters=Typedef, QFlags<QDir::Filter>, Unexposed
	filters := 0
	// arg: 1, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir9entryListE6QFlagsINS_6FilterEES0_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList entryList(QDir::Filters, QDir::SortFlags) const

/*
Returns a list of the names of all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryInfoList(), setNameFilters(), setSorting(), and setFilter().
*/
func (this *QDir) EntryListp1(filters int) *QStringList /*123*/ {
	// arg: 1, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir9entryListE6QFlagsINS_6FilterEES0_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:154
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList entryList(const QStringList &, QDir::Filters, QDir::SortFlags) const

/*
Returns a list of the names of all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryInfoList(), setNameFilters(), setSorting(), and setFilter().
*/
func (this *QDir) EntryList1(nameFilters QStringList_ITF, filters int, sort int) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir9entryListERK11QStringList6QFlagsINS_6FilterEES3_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:154
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList entryList(const QStringList &, QDir::Filters, QDir::SortFlags) const

/*
Returns a list of the names of all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryInfoList(), setNameFilters(), setSorting(), and setFilter().
*/
func (this *QDir) EntryList1p(nameFilters QStringList_ITF) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	// arg: 1, QDir::Filters=Typedef, QDir::Filters=Typedef, QFlags<QDir::Filter>, Unexposed
	filters := 0
	// arg: 2, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir9entryListERK11QStringList6QFlagsINS_6FilterEES3_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:154
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList entryList(const QStringList &, QDir::Filters, QDir::SortFlags) const

/*
Returns a list of the names of all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryInfoList(), setNameFilters(), setSorting(), and setFilter().
*/
func (this *QDir) EntryList1p1(nameFilters QStringList_ITF, filters int) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	// arg: 2, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir9entryListERK11QStringList6QFlagsINS_6FilterEES3_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] QFileInfoList entryInfoList(QDir::Filters, QDir::SortFlags) const

/*
Returns a list of QFileInfo objects for all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryList(), setNameFilters(), setSorting(), setFilter(), isReadable(), and exists().
*/
func (this *QDir) EntryInfoList(filters int, sort int) *QFileInfoList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13entryInfoListE6QFlagsINS_6FilterEES0_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] QFileInfoList entryInfoList(QDir::Filters, QDir::SortFlags) const

/*
Returns a list of QFileInfo objects for all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryList(), setNameFilters(), setSorting(), setFilter(), isReadable(), and exists().
*/
func (this *QDir) EntryInfoListp() *QFileInfoList /*667*/ {
	// arg: 0, QDir::Filters=Typedef, QDir::Filters=Typedef, QFlags<QDir::Filter>, Unexposed
	filters := 0
	// arg: 1, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13entryInfoListE6QFlagsINS_6FilterEES0_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] QFileInfoList entryInfoList(QDir::Filters, QDir::SortFlags) const

/*
Returns a list of QFileInfo objects for all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryList(), setNameFilters(), setSorting(), setFilter(), isReadable(), and exists().
*/
func (this *QDir) EntryInfoListp1(filters int) *QFileInfoList /*667*/ {
	// arg: 1, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13entryInfoListE6QFlagsINS_6FilterEES0_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:158
// index:1
// Public Visibility=Default Availability=Available
// [-2] QFileInfoList entryInfoList(const QStringList &, QDir::Filters, QDir::SortFlags) const

/*
Returns a list of QFileInfo objects for all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryList(), setNameFilters(), setSorting(), setFilter(), isReadable(), and exists().
*/
func (this *QDir) EntryInfoList1(nameFilters QStringList_ITF, filters int, sort int) *QFileInfoList /*667*/ {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13entryInfoListERK11QStringList6QFlagsINS_6FilterEES3_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:158
// index:1
// Public Visibility=Default Availability=Available
// [-2] QFileInfoList entryInfoList(const QStringList &, QDir::Filters, QDir::SortFlags) const

/*
Returns a list of QFileInfo objects for all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryList(), setNameFilters(), setSorting(), setFilter(), isReadable(), and exists().
*/
func (this *QDir) EntryInfoList1p(nameFilters QStringList_ITF) *QFileInfoList /*667*/ {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	// arg: 1, QDir::Filters=Typedef, QDir::Filters=Typedef, QFlags<QDir::Filter>, Unexposed
	filters := 0
	// arg: 2, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13entryInfoListERK11QStringList6QFlagsINS_6FilterEES3_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:158
// index:1
// Public Visibility=Default Availability=Available
// [-2] QFileInfoList entryInfoList(const QStringList &, QDir::Filters, QDir::SortFlags) const

/*
Returns a list of QFileInfo objects for all the files and directories in the directory, ordered according to the name and attribute filters previously set with setNameFilters() and setFilter(), and sorted according to the flags set with setSorting().

The name filter, file attribute filter, and sorting specification can be overridden using the nameFilters, filters, and sort arguments.

Returns an empty list if the directory is unreadable, does not exist, or if nothing matches the specification.

See also entryList(), setNameFilters(), setSorting(), setFilter(), isReadable(), and exists().
*/
func (this *QDir) EntryInfoList1p1(nameFilters QStringList_ITF, filters int) *QFileInfoList /*667*/ {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	// arg: 2, QDir::SortFlags=Typedef, QDir::SortFlags=Typedef, QFlags<QDir::SortFlag>, Unexposed
	sort := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir13entryInfoListERK11QStringList6QFlagsINS_6FilterEES3_INS_8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, filters, sort)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qdir.h:161
// index:0
// Public Visibility=Default Availability=Available
// [1] bool mkdir(const QString &) const

/*
Creates a sub-directory called dirName.

Returns true on success; otherwise returns false.

If the directory already exists when this function is called, it will return false.

See also rmdir().
*/
func (this *QDir) Mkdir(dirName string) bool {
	var tmpArg0 = NewQString5(dirName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir5mkdirERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:162
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rmdir(const QString &) const

/*
Removes the directory specified by dirName.

The directory must be empty for rmdir() to succeed.

Returns true if successful; otherwise returns false.

See also mkdir().
*/
func (this *QDir) Rmdir(dirName string) bool {
	var tmpArg0 = NewQString5(dirName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir5rmdirERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool mkpath(const QString &) const

/*
Creates the directory path dirPath.

The function will create all parent directories necessary to create the directory.

Returns true if successful; otherwise returns false.

If the path already exists when this function is called, it will return true.

See also rmpath().
*/
func (this *QDir) Mkpath(dirPath string) bool {
	var tmpArg0 = NewQString5(dirPath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6mkpathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rmpath(const QString &) const

/*
Removes the directory path dirPath.

The function will remove all parent directories in dirPath, provided that they are empty. This is the opposite of mkpath(dirPath).

Returns true if successful; otherwise returns false.

See also mkpath().
*/
func (this *QDir) Rmpath(dirPath string) bool {
	var tmpArg0 = NewQString5(dirPath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6rmpathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool removeRecursively()

/*
Removes the directory, including all its contents.

Returns true if successful, otherwise false.

If a file or directory cannot be removed, removeRecursively() keeps going and attempts to delete as many files and sub-directories as possible, then returns false.

If the directory was already removed, the method returns true (expected result already reached).

Note: this function is meant for removing a small application-internal directory (such as a temporary directory), but not user-visible directories. For user-visible operations, it is rather recommended to report errors more precisely to the user, to offer solutions in case of errors, to show progress during the deletion since it could take several minutes, etc.

This function was introduced in  Qt 5.0.
*/
func (this *QDir) RemoveRecursively() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir17removeRecursivelyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:168
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable() const

/*
Returns true if the directory is readable and we can open files by name; otherwise returns false.

Warning: A false value from this function is not a guarantee that files in the directory are not accessible.

See also QFileInfo::isReadable().
*/
func (this *QDir) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:169
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exists() const

/*
Returns true if the file called name exists; otherwise returns false.

Unless name contains an absolute file path, the file name is assumed to be relative to the directory itself, so this function is typically used to check for the presence of files within a directory.

See also QFileInfo::exists() and QFile::exists().
*/
func (this *QDir) Exists() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6existsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:183
// index:1
// Public Visibility=Default Availability=Available
// [1] bool exists(const QString &) const

/*
Returns true if the file called name exists; otherwise returns false.

Unless name contains an absolute file path, the file name is assumed to be relative to the directory itself, so this function is typically used to check for the presence of files within a directory.

See also QFileInfo::exists() and QFile::exists().
*/
func (this *QDir) Exists1(name string) bool {
	var tmpArg0 = NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6existsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRoot() const

/*
Returns true if the directory is the root directory; otherwise returns false.

Note: If the directory is a symbolic link to the root directory this function returns false. If you want to test for this use canonicalPath(), e.g.


  QDir dir("/tmp/root_link");
  dir = dir.canonicalPath();
  if (dir.isRoot())
      qWarning("It is a root link");



See also root() and rootPath().
*/
func (this *QDir) IsRoot() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir6isRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:172
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isRelativePath(const QString &)

/*
Returns true if path is relative; returns false if it is absolute.

See also isRelative(), isAbsolutePath(), and makeAbsolute().
*/
func (this *QDir) IsRelativePath(path string) bool {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir14isRelativePathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_IsRelativePath(path string) bool {
	var nilthis *QDir
	rv := nilthis.IsRelativePath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:173
// index:0
// Public static inline Visibility=Default Availability=Available
// [1] bool isAbsolutePath(const QString &)

/*
Returns true if path is absolute; returns false if it is relative.

See also isAbsolute(), isRelativePath(), makeAbsolute(), and cleanPath().
*/
func (this *QDir) IsAbsolutePath(path string) bool {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir14isAbsolutePathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_IsAbsolutePath(path string) bool {
	var nilthis *QDir
	rv := nilthis.IsAbsolutePath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:174
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRelative() const

/*
Returns true if the directory path is relative; otherwise returns false. (Under Unix a path is relative if it does not start with a "/").

See also makeAbsolute(), isAbsolute(), isAbsolutePath(), and cleanPath().
*/
func (this *QDir) IsRelative() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir10isRelativeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:175
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAbsolute() const

/*
Returns true if the directory's path is absolute; otherwise returns false. See isAbsolutePath().

See also isRelative(), makeAbsolute(), and cleanPath().
*/
func (this *QDir) IsAbsolute() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir10isAbsoluteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool makeAbsolute()

/*
Converts the directory path to an absolute path. If it is already absolute nothing happens. Returns true if the conversion succeeded; otherwise returns false.

See also isAbsolute(), isAbsolutePath(), isRelative(), and cleanPath().
*/
func (this *QDir) MakeAbsolute() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir12makeAbsoluteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:178
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QDir &) const

/*

 */
func (this *QDir) Operator_equal_equal(dir QDir_ITF) bool {
	var convArg0 unsafe.Pointer
	if dir != nil && dir.QDir_PTR() != nil {
		convArg0 = dir.QDir_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDireqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:179
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QDir &) const

/*

 */
func (this *QDir) Operator_not_equal(dir QDir_ITF) bool {
	var convArg0 unsafe.Pointer
	if dir != nil && dir.QDir_PTR() != nil {
		convArg0 = dir.QDir_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDirneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove(const QString &)

/*
Removes the file, fileName.

Returns true if the file is removed successfully; otherwise returns false.
*/
func (this *QDir) Remove(fileName string) bool {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir6removeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rename(const QString &, const QString &)

/*
Renames a file or directory from oldName to newName, and returns true if successful; otherwise returns false.

On most file systems, rename() fails only if oldName does not exist, or if a file with the new name already exists. However, there are also other reasons why rename() can fail. For example, on at least one file system rename() fails if newName points to an open file.

If oldName is a file (not a directory) that can't be renamed right away, Qt will try to copy oldName to newName and remove oldName.

See also QFile::rename().
*/
func (this *QDir) Rename(oldName string, newName string) bool {
	var tmpArg0 = NewQString5(oldName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(newName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir6renameERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdir.h:185
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QFileInfoList drives()

/*
Returns a list of the root directories on this system.

On Windows this returns a list of QFileInfo objects containing "C:/", "D:/", etc. On other operating systems, it returns a list containing just one root directory (i.e. "/").

See also root() and rootPath().
*/
func (this *QDir) Drives() *QFileInfoList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir6drivesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}
func QDir_Drives() *QFileInfoList /*667*/ {
	var nilthis *QDir
	rv := nilthis.Drives()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:187
// index:0
// Public static inline Visibility=Default Availability=Available
// [2] QChar listSeparator()

/*
Returns the native path list separator: ':' under Unix and ';' under Windows.

This function was introduced in  Qt 5.6.

See also separator().
*/
func (this *QDir) ListSeparator() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir13listSeparatorEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}
func QDir_ListSeparator() *QChar /*123*/ {
	var nilthis *QDir
	rv := nilthis.ListSeparator()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:196
// index:0
// Public static Visibility=Default Availability=Available
// [2] QChar separator()

/*
Returns the native directory separator: "/" under Unix and "\" under Windows.

You do not need to use this function to build file paths. If you always use "/", Qt will translate your paths to conform to the underlying operating system. If you want to display paths to the user using their operating system's separator use toNativeSeparators().

See also listSeparator().
*/
func (this *QDir) Separator() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir9separatorEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}
func QDir_Separator() *QChar /*123*/ {
	var nilthis *QDir
	rv := nilthis.Separator()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:198
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool setCurrent(const QString &)

/*
Sets the application's current working directory to path. Returns true if the directory was successfully changed; otherwise returns false.

See also current(), currentPath(), home(), root(), and temp().
*/
func (this *QDir) SetCurrent(path string) bool {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir10setCurrentERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_SetCurrent(path string) bool {
	var nilthis *QDir
	rv := nilthis.SetCurrent(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:199
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir current()

/*
Returns the application's current directory.

The directory is constructed using the absolute path of the current directory, ensuring that its path() will be the same as its absolutePath().

See also currentPath(), setCurrent(), home(), root(), and temp().
*/
func (this *QDir) Current() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir7currentEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}
func QDir_Current() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Current()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:200
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString currentPath()

/*
Returns the absolute path of the application's current directory. The current directory is the last directory set with QDir::setCurrent() or, if that was never called, the directory at which this application was started at by the parent process.

See also current(), setCurrent(), homePath(), rootPath(), tempPath(), and QCoreApplication::applicationDirPath().
*/
func (this *QDir) CurrentPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir11currentPathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_CurrentPath() string {
	var nilthis *QDir
	rv := nilthis.CurrentPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:202
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir home()

/*
Returns the user's home directory.

The directory is constructed using the absolute path of the home directory, ensuring that its path() will be the same as its absolutePath().

See homePath() for details.

See also drives(), current(), root(), and temp().
*/
func (this *QDir) Home() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4homeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}
func QDir_Home() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Home()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:203
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString homePath()

/*
Returns the absolute path of the user's home directory.

Under Windows this function will return the directory of the current user's profile. Typically, this is:


  C:/Documents and Settings/Username



Use the toNativeSeparators() function to convert the separators to the ones that are appropriate for the underlying operating system.

If the directory of the current user's profile does not exist or cannot be retrieved, the following alternatives will be checked (in the given order) until an existing and available path is found:
*/
func (this *QDir) HomePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir8homePathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_HomePath() string {
	var nilthis *QDir
	rv := nilthis.HomePath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:204
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir root()

/*
Returns the root directory.

The directory is constructed using the absolute path of the root directory, ensuring that its path() will be the same as its absolutePath().

See rootPath() for details.

See also drives(), current(), home(), and temp().
*/
func (this *QDir) Root() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4rootEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}
func QDir_Root() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Root()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:205
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString rootPath()

/*
Returns the absolute path of the root directory.

For Unix operating systems this returns "/". For Windows file systems this normally returns "c:/".

See also root(), drives(), currentPath(), homePath(), and tempPath().
*/
func (this *QDir) RootPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir8rootPathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_RootPath() string {
	var nilthis *QDir
	rv := nilthis.RootPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:206
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QDir temp()

/*
Returns the system's temporary directory.

The directory is constructed using the absolute path of the temporary directory, ensuring that its path() will be the same as its absolutePath().

See tempPath() for details.

See also drives(), current(), home(), and root().
*/
func (this *QDir) Temp() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir4tempEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}
func QDir_Temp() *QDir /*123*/ {
	var nilthis *QDir
	rv := nilthis.Temp()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:207
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString tempPath()

/*
Returns the absolute path of the system's temporary directory.

On Unix/Linux systems this is the path in the TMPDIR environment variable or /tmp if TMPDIR is not defined. On Windows this is usually the path in the TEMP or TMP environment variable. The path returned by this method doesn't end with a directory separator unless it is the root directory (of a drive).

See also temp(), currentPath(), homePath(), and rootPath().
*/
func (this *QDir) TempPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir8tempPathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_TempPath() string {
	var nilthis *QDir
	rv := nilthis.TempPath()
	return rv
}

// /usr/include/qt/QtCore/qdir.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool match(const QStringList &, const QString &)

/*
Returns true if the fileName matches the wildcard (glob) pattern filter; otherwise returns false. The filter may contain multiple patterns separated by spaces or semicolons. The matching is case insensitive.

See also QRegExp wildcard matching, QRegExp::exactMatch(), entryList(), and entryInfoList().
*/
func (this *QDir) Match(filters QStringList_ITF, fileName string) bool {
	var convArg0 unsafe.Pointer
	if filters != nil && filters.QStringList_PTR() != nil {
		convArg0 = filters.QStringList_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir5matchERK11QStringListRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_Match(filters QStringList_ITF, fileName string) bool {
	var nilthis *QDir
	rv := nilthis.Match(filters, fileName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:211
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool match(const QString &, const QString &)

/*
Returns true if the fileName matches the wildcard (glob) pattern filter; otherwise returns false. The filter may contain multiple patterns separated by spaces or semicolons. The matching is case insensitive.

See also QRegExp wildcard matching, QRegExp::exactMatch(), entryList(), and entryInfoList().
*/
func (this *QDir) Match1(filter string, fileName string) bool {
	var tmpArg0 = NewQString5(filter)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir5matchERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QDir_Match1(filter string, fileName string) bool {
	var nilthis *QDir
	rv := nilthis.Match1(filter, fileName)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:214
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString cleanPath(const QString &)

/*
Returns path with directory separators normalized (converted to "/") and redundant ones removed, and "."s and ".."s resolved (as far as possible).

Symbolic links are kept. This function does not return the canonical path, but rather the simplest version of the input. For example, "./local" becomes "local", "local/../bin" becomes "bin" and "/local/usr/../bin" becomes "/local/bin".

See also absolutePath() and canonicalPath().
*/
func (this *QDir) CleanPath(path string) string {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QDir9cleanPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QDir_CleanPath(path string) string {
	var nilthis *QDir
	rv := nilthis.CleanPath(path)
	return rv
}

// /usr/include/qt/QtCore/qdir.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh() const

/*
Refreshes the directory information.
*/
func (this *QDir) Refresh() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QDir7refreshEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QDir__Filter = int

//
const QDir__Dirs QDir__Filter = 1

//
const QDir__Files QDir__Filter = 2

//
const QDir__Drives QDir__Filter = 4

//
const QDir__NoSymLinks QDir__Filter = 8

//
const QDir__AllEntries QDir__Filter = 7

//
const QDir__TypeMask QDir__Filter = 15

//
const QDir__Readable QDir__Filter = 16

//
const QDir__Writable QDir__Filter = 32

//
const QDir__Executable QDir__Filter = 64

//
const QDir__PermissionMask QDir__Filter = 112

//
const QDir__Modified QDir__Filter = 128

//
const QDir__Hidden QDir__Filter = 256

//
const QDir__System QDir__Filter = 512

//
const QDir__AccessMask QDir__Filter = 1008

//
const QDir__AllDirs QDir__Filter = 1024

//
const QDir__CaseSensitive QDir__Filter = 2048

//
const QDir__NoDot QDir__Filter = 8192

//
const QDir__NoDotDot QDir__Filter = 16384

//
const QDir__NoDotAndDotDot QDir__Filter = 24576

//
const QDir__NoFilter QDir__Filter = -1

func (this *QDir) FilterItemName(val int) string {
	switch val {
	case QDir__Dirs: // 1
		return "Dirs"
	case QDir__Files: // 2
		return "Files"
	case QDir__Drives: // 4
		return "Drives"
	case QDir__NoSymLinks: // 8
		return "NoSymLinks"
	case QDir__AllEntries: // 7
		return "AllEntries"
	case QDir__TypeMask: // 15
		return "TypeMask"
	case QDir__Readable: // 16
		return "Readable"
	case QDir__Writable: // 32
		return "Writable"
	case QDir__Executable: // 64
		return "Executable"
	case QDir__PermissionMask: // 112
		return "PermissionMask"
	case QDir__Modified: // 128
		return "Modified"
	case QDir__Hidden: // 256
		return "Hidden"
	case QDir__System: // 512
		return "System"
	case QDir__AccessMask: // 1008
		return "AccessMask"
	case QDir__AllDirs: // 1024
		return "AllDirs"
	case QDir__CaseSensitive: // 2048
		return "CaseSensitive"
	case QDir__NoDot: // 8192
		return "NoDot"
	case QDir__NoDotDot: // 16384
		return "NoDotDot"
	case QDir__NoDotAndDotDot: // 24576
		return "NoDotAndDotDot"
	case QDir__NoFilter: // -1
		return "NoFilter"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QDir_FilterItemName(val int) string {
	var nilthis *QDir
	return nilthis.FilterItemName(val)
}

/*


 */
type QDir__SortFlag = int

//
const QDir__Name QDir__SortFlag = 0

//
const QDir__Time QDir__SortFlag = 1

//
const QDir__Size QDir__SortFlag = 2

//
const QDir__Unsorted QDir__SortFlag = 3

//
const QDir__SortByMask QDir__SortFlag = 3

//
const QDir__DirsFirst QDir__SortFlag = 4

//
const QDir__Reversed QDir__SortFlag = 8

//
const QDir__IgnoreCase QDir__SortFlag = 16

//
const QDir__DirsLast QDir__SortFlag = 32

//
const QDir__LocaleAware QDir__SortFlag = 64

//
const QDir__Type QDir__SortFlag = 128

//
const QDir__NoSort QDir__SortFlag = -1

func (this *QDir) SortFlagItemName(val int) string {
	switch val {
	case QDir__Name: // 0
		return "Name"
	case QDir__Time: // 1
		return "Time"
	case QDir__Size: // 2
		return "Size"
	case QDir__Unsorted: // 3
		return "Unsorted,SortByMask"
		// case QDir__SortByMask: // 3
		// return ""
	case QDir__DirsFirst: // 4
		return "DirsFirst"
	case QDir__Reversed: // 8
		return "Reversed"
	case QDir__IgnoreCase: // 16
		return "IgnoreCase"
	case QDir__DirsLast: // 32
		return "DirsLast"
	case QDir__LocaleAware: // 64
		return "LocaleAware"
	case QDir__Type: // 128
		return "Type"
	case QDir__NoSort: // -1
		return "NoSort"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QDir_SortFlagItemName(val int) string {
	var nilthis *QDir
	return nilthis.SortFlagItemName(val)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
