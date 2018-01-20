//  header block begin
// /usr/include/qt/QtWidgets/qdatetimeedit.h
// #include <qdatetimeedit.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QDateTimeEdit struct {
	*QAbstractSpinBox
}

func (this *QDateTimeEdit) GetCthis() unsafe.Pointer {
	return this.QAbstractSpinBox.GetCthis()
}
func NewQDateTimeEditFromPointer(cthis unsafe.Pointer) *QDateTimeEdit {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QDateTimeEdit{bcthis0}
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDateTimeEdit) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:95
// index:0
// Public
// void QDateTimeEdit(class QWidget *)
func NewQDateTimeEdit(parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:96
// index:1
// Public
// void QDateTimeEdit(const class QDateTime &, class QWidget *)
func NewQDateTimeEdit_1(dt *qtcore.QDateTime, parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = dt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK9QDateTimeP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:97
// index:2
// Public
// void QDateTimeEdit(const class QDate &, class QWidget *)
func NewQDateTimeEdit_2(d *qtcore.QDate, parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = d.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QDateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:98
// index:3
// Public
// void QDateTimeEdit(const class QTime &, class QWidget *)
func NewQDateTimeEdit_3(t *qtcore.QTime, parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = t.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QTimeP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:194
// index:4
// Protected
// void QDateTimeEdit(const class QVariant &, class QVariant::Type, class QWidget *)
func NewQDateTimeEdit_4(val *qtcore.QVariant, parserType int, parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = val.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK8QVariantNS0_4TypeEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &parserType, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:99
// index:0
// Public virtual
// void ~QDateTimeEdit()
func DeleteQDateTimeEdit(*QDateTimeEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:101
// index:0
// Public
// QDateTime dateTime()
func (this *QDateTimeEdit) DateTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit8dateTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:102
// index:0
// Public
// QDate date()
func (this *QDateTimeEdit) Date() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit4dateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:103
// index:0
// Public
// QTime time()
func (this *QDateTimeEdit) Time() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit4timeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:105
// index:0
// Public
// QDateTime minimumDateTime()
func (this *QDateTimeEdit) MinimumDateTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit15minimumDateTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:106
// index:0
// Public
// void clearMinimumDateTime()
func (this *QDateTimeEdit) ClearMinimumDateTime() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit20clearMinimumDateTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:107
// index:0
// Public
// void setMinimumDateTime(const class QDateTime &)
func (this *QDateTimeEdit) SetMinimumDateTime(dt *qtcore.QDateTime) {
	var convArg0 = dt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:109
// index:0
// Public
// QDateTime maximumDateTime()
func (this *QDateTimeEdit) MaximumDateTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit15maximumDateTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:110
// index:0
// Public
// void clearMaximumDateTime()
func (this *QDateTimeEdit) ClearMaximumDateTime() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit20clearMaximumDateTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:111
// index:0
// Public
// void setMaximumDateTime(const class QDateTime &)
func (this *QDateTimeEdit) SetMaximumDateTime(dt *qtcore.QDateTime) {
	var convArg0 = dt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:113
// index:0
// Public
// void setDateTimeRange(const class QDateTime &, const class QDateTime &)
func (this *QDateTimeEdit) SetDateTimeRange(min *qtcore.QDateTime, max *qtcore.QDateTime) {
	var convArg0 = min.GetCthis()
	var convArg1 = max.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:115
// index:0
// Public
// QDate minimumDate()
func (this *QDateTimeEdit) MinimumDate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11minimumDateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:116
// index:0
// Public
// void setMinimumDate(const class QDate &)
func (this *QDateTimeEdit) SetMinimumDate(min *qtcore.QDate) {
	var convArg0 = min.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMinimumDateERK5QDate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:117
// index:0
// Public
// void clearMinimumDate()
func (this *QDateTimeEdit) ClearMinimumDate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMinimumDateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:119
// index:0
// Public
// QDate maximumDate()
func (this *QDateTimeEdit) MaximumDate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11maximumDateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:120
// index:0
// Public
// void setMaximumDate(const class QDate &)
func (this *QDateTimeEdit) SetMaximumDate(max *qtcore.QDate) {
	var convArg0 = max.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMaximumDateERK5QDate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:121
// index:0
// Public
// void clearMaximumDate()
func (this *QDateTimeEdit) ClearMaximumDate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMaximumDateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:123
// index:0
// Public
// void setDateRange(const class QDate &, const class QDate &)
func (this *QDateTimeEdit) SetDateRange(min *qtcore.QDate, max *qtcore.QDate) {
	var convArg0 = min.GetCthis()
	var convArg1 = max.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit12setDateRangeERK5QDateS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:125
// index:0
// Public
// QTime minimumTime()
func (this *QDateTimeEdit) MinimumTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11minimumTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:126
// index:0
// Public
// void setMinimumTime(const class QTime &)
func (this *QDateTimeEdit) SetMinimumTime(min *qtcore.QTime) {
	var convArg0 = min.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMinimumTimeERK5QTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:127
// index:0
// Public
// void clearMinimumTime()
func (this *QDateTimeEdit) ClearMinimumTime() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMinimumTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:129
// index:0
// Public
// QTime maximumTime()
func (this *QDateTimeEdit) MaximumTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11maximumTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:130
// index:0
// Public
// void setMaximumTime(const class QTime &)
func (this *QDateTimeEdit) SetMaximumTime(max *qtcore.QTime) {
	var convArg0 = max.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMaximumTimeERK5QTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:131
// index:0
// Public
// void clearMaximumTime()
func (this *QDateTimeEdit) ClearMaximumTime() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMaximumTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:133
// index:0
// Public
// void setTimeRange(const class QTime &, const class QTime &)
func (this *QDateTimeEdit) SetTimeRange(min *qtcore.QTime, max *qtcore.QTime) {
	var convArg0 = min.GetCthis()
	var convArg1 = max.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:135
// index:0
// Public
// QDateTimeEdit::Sections displayedSections()
func (this *QDateTimeEdit) DisplayedSections() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit17displayedSectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:136
// index:0
// Public
// QDateTimeEdit::Section currentSection()
func (this *QDateTimeEdit) CurrentSection() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit14currentSectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:137
// index:0
// Public
// QDateTimeEdit::Section sectionAt(int)
func (this *QDateTimeEdit) SectionAt(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit9sectionAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:138
// index:0
// Public
// void setCurrentSection(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SetCurrentSection(section int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit17setCurrentSectionENS_7SectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:140
// index:0
// Public
// int currentSectionIndex()
func (this *QDateTimeEdit) CurrentSectionIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit19currentSectionIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:141
// index:0
// Public
// void setCurrentSectionIndex(int)
func (this *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit22setCurrentSectionIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:143
// index:0
// Public
// QCalendarWidget * calendarWidget()
func (this *QDateTimeEdit) CalendarWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit14calendarWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:144
// index:0
// Public
// void setCalendarWidget(class QCalendarWidget *)
func (this *QDateTimeEdit) SetCalendarWidget(calendarWidget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), calendarWidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:146
// index:0
// Public
// int sectionCount()
func (this *QDateTimeEdit) SectionCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit12sectionCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:148
// index:0
// Public
// void setSelectedSection(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SetSelectedSection(section int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit18setSelectedSectionENS_7SectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:150
// index:0
// Public
// QString sectionText(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SectionText(section int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11sectionTextENS_7SectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:152
// index:0
// Public
// QString displayFormat()
func (this *QDateTimeEdit) DisplayFormat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit13displayFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:153
// index:0
// Public
// void setDisplayFormat(const class QString &)
func (this *QDateTimeEdit) SetDisplayFormat(format *qtcore.QString) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16setDisplayFormatERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:155
// index:0
// Public
// bool calendarPopup()
func (this *QDateTimeEdit) CalendarPopup() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit13calendarPopupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:156
// index:0
// Public
// void setCalendarPopup(_Bool)
func (this *QDateTimeEdit) SetCalendarPopup(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16setCalendarPopupEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:158
// index:0
// Public
// Qt::TimeSpec timeSpec()
func (this *QDateTimeEdit) TimeSpec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit8timeSpecEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:159
// index:0
// Public
// void setTimeSpec(Qt::TimeSpec)
func (this *QDateTimeEdit) SetTimeSpec(spec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit11setTimeSpecEN2Qt8TimeSpecE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:161
// index:0
// Public virtual
// QSize sizeHint()
func (this *QDateTimeEdit) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:163
// index:0
// Public virtual
// void clear()
func (this *QDateTimeEdit) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:164
// index:0
// Public virtual
// void stepBy(int)
func (this *QDateTimeEdit) StepBy(steps int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit6stepByEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:166
// index:0
// Public virtual
// bool event(class QEvent *)
func (this *QDateTimeEdit) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:168
// index:0
// Public
// void dateTimeChanged(const class QDateTime &)
func (this *QDateTimeEdit) DateTimeChanged(dateTime *qtcore.QDateTime) {
	var convArg0 = dateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit15dateTimeChangedERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:169
// index:0
// Public
// void timeChanged(const class QTime &)
func (this *QDateTimeEdit) TimeChanged(time *qtcore.QTime) {
	var convArg0 = time.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit11timeChangedERK5QTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:170
// index:0
// Public
// void dateChanged(const class QDate &)
func (this *QDateTimeEdit) DateChanged(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit11dateChangedERK5QDate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:173
// index:0
// Public
// void setDateTime(const class QDateTime &)
func (this *QDateTimeEdit) SetDateTime(dateTime *qtcore.QDateTime) {
	var convArg0 = dateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit11setDateTimeERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:174
// index:0
// Public
// void setDate(const class QDate &)
func (this *QDateTimeEdit) SetDate(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit7setDateERK5QDate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:175
// index:0
// Public
// void setTime(const class QTime &)
func (this *QDateTimeEdit) SetTime(time *qtcore.QTime) {
	var convArg0 = time.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit7setTimeERK5QTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:178
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QDateTimeEdit) KeyPressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:180
// index:0
// Protected virtual
// void wheelEvent(class QWheelEvent *)
func (this *QDateTimeEdit) WheelEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:182
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QDateTimeEdit) FocusInEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:183
// index:0
// Protected virtual
// bool focusNextPrevChild(_Bool)
func (this *QDateTimeEdit) FocusNextPrevChild(next bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:184
// index:0
// Protected virtual
// QValidator::State validate(class QString &, int &)
func (this *QDateTimeEdit) Validate(input *qtcore.QString, pos int) interface{} {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:185
// index:0
// Protected virtual
// void fixup(class QString &)
func (this *QDateTimeEdit) Fixup(input *qtcore.QString) {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit5fixupER7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:187
// index:0
// Protected virtual
// QDateTime dateTimeFromText(const class QString &)
func (this *QDateTimeEdit) DateTimeFromText(text *qtcore.QString) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit16dateTimeFromTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:188
// index:0
// Protected virtual
// QString textFromDateTime(const class QDateTime &)
func (this *QDateTimeEdit) TextFromDateTime(dt *qtcore.QDateTime) interface{} {
	var convArg0 = dt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit16textFromDateTimeERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:189
// index:0
// Protected virtual
// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QDateTimeEdit) StepEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11stepEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:190
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QDateTimeEdit) MousePressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:191
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QDateTimeEdit) PaintEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:192
// index:0
// Protected
// void initStyleOption(class QStyleOptionSpinBox *)
func (this *QDateTimeEdit) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit15initStyleOptionEP19QStyleOptionSpinBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end