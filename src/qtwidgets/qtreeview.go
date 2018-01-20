//  header block begin
// /usr/include/qt/QtWidgets/qtreeview.h
// #include <qtreeview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QTreeView struct {
	*QAbstractItemView
}

func (this *QTreeView) GetCthis() unsafe.Pointer {
	return this.QAbstractItemView.GetCthis()
}
func NewQTreeViewFromPointer(cthis unsafe.Pointer) *QTreeView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QTreeView{bcthis0}
}

// /usr/include/qt/QtWidgets/qtreeview.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTreeView) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:71
// index:0
// Public
// void QTreeView(class QWidget *)
func NewQTreeView(parent unsafe.Pointer) *QTreeView {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreeview.h:72
// index:0
// Public virtual
// void ~QTreeView()
func DeleteQTreeView(*QTreeView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:74
// index:0
// Public virtual
// void setModel(class QAbstractItemModel *)
func (this *QTreeView) SetModel(model unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:75
// index:0
// Public virtual
// void setRootIndex(const class QModelIndex &)
func (this *QTreeView) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:76
// index:0
// Public virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTreeView) SetSelectionModel(selectionModel unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:78
// index:0
// Public
// QHeaderView * header()
func (this *QTreeView) Header() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView6headerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:79
// index:0
// Public
// void setHeader(class QHeaderView *)
func (this *QTreeView) SetHeader(header unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9setHeaderEP11QHeaderView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), header)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:81
// index:0
// Public
// int autoExpandDelay()
func (this *QTreeView) AutoExpandDelay() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15autoExpandDelayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:82
// index:0
// Public
// void setAutoExpandDelay(int)
func (this *QTreeView) SetAutoExpandDelay(delay int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setAutoExpandDelayEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &delay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:84
// index:0
// Public
// int indentation()
func (this *QTreeView) Indentation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11indentationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:85
// index:0
// Public
// void setIndentation(int)
func (this *QTreeView) SetIndentation(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14setIndentationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:86
// index:0
// Public
// void resetIndentation()
func (this *QTreeView) ResetIndentation() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16resetIndentationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:88
// index:0
// Public
// bool rootIsDecorated()
func (this *QTreeView) RootIsDecorated() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15rootIsDecoratedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:89
// index:0
// Public
// void setRootIsDecorated(_Bool)
func (this *QTreeView) SetRootIsDecorated(show bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setRootIsDecoratedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:91
// index:0
// Public
// bool uniformRowHeights()
func (this *QTreeView) UniformRowHeights() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView17uniformRowHeightsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:92
// index:0
// Public
// void setUniformRowHeights(_Bool)
func (this *QTreeView) SetUniformRowHeights(uniform bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView20setUniformRowHeightsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &uniform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:94
// index:0
// Public
// bool itemsExpandable()
func (this *QTreeView) ItemsExpandable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15itemsExpandableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:95
// index:0
// Public
// void setItemsExpandable(_Bool)
func (this *QTreeView) SetItemsExpandable(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setItemsExpandableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:97
// index:0
// Public
// bool expandsOnDoubleClick()
func (this *QTreeView) ExpandsOnDoubleClick() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView20expandsOnDoubleClickEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:98
// index:0
// Public
// void setExpandsOnDoubleClick(_Bool)
func (this *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView23setExpandsOnDoubleClickEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:100
// index:0
// Public
// int columnViewportPosition(int)
func (this *QTreeView) ColumnViewportPosition(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView22columnViewportPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:101
// index:0
// Public
// int columnWidth(int)
func (this *QTreeView) ColumnWidth(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11columnWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:102
// index:0
// Public
// void setColumnWidth(int, int)
func (this *QTreeView) SetColumnWidth(column int, width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14setColumnWidthEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:103
// index:0
// Public
// int columnAt(int)
func (this *QTreeView) ColumnAt(x int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8columnAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:105
// index:0
// Public
// bool isColumnHidden(int)
func (this *QTreeView) IsColumnHidden(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14isColumnHiddenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:106
// index:0
// Public
// void setColumnHidden(int, _Bool)
func (this *QTreeView) SetColumnHidden(column int, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setColumnHiddenEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:108
// index:0
// Public
// bool isHeaderHidden()
func (this *QTreeView) IsHeaderHidden() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14isHeaderHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:109
// index:0
// Public
// void setHeaderHidden(_Bool)
func (this *QTreeView) SetHeaderHidden(hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setHeaderHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:111
// index:0
// Public
// bool isRowHidden(int, const class QModelIndex &)
func (this *QTreeView) IsRowHidden(row int, parent *qtcore.QModelIndex) interface{} {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:112
// index:0
// Public
// void setRowHidden(int, const class QModelIndex &, _Bool)
func (this *QTreeView) SetRowHidden(row int, parent *qtcore.QModelIndex, hide bool) {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, convArg1, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:114
// index:0
// Public
// bool isFirstColumnSpanned(int, const class QModelIndex &)
func (this *QTreeView) IsFirstColumnSpanned(row int, parent *qtcore.QModelIndex) interface{} {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:115
// index:0
// Public
// void setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
func (this *QTreeView) SetFirstColumnSpanned(row int, parent *qtcore.QModelIndex, span bool) {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, convArg1, &span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:117
// index:0
// Public
// bool isExpanded(const class QModelIndex &)
func (this *QTreeView) IsExpanded(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10isExpandedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:118
// index:0
// Public
// void setExpanded(const class QModelIndex &, _Bool)
func (this *QTreeView) SetExpanded(index *qtcore.QModelIndex, expand bool) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setExpandedERK11QModelIndexb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:120
// index:0
// Public
// void setSortingEnabled(_Bool)
func (this *QTreeView) SetSortingEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17setSortingEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:121
// index:0
// Public
// bool isSortingEnabled()
func (this *QTreeView) IsSortingEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16isSortingEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:123
// index:0
// Public
// void setAnimated(_Bool)
func (this *QTreeView) SetAnimated(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setAnimatedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:124
// index:0
// Public
// bool isAnimated()
func (this *QTreeView) IsAnimated() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10isAnimatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:126
// index:0
// Public
// void setAllColumnsShowFocus(_Bool)
func (this *QTreeView) SetAllColumnsShowFocus(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView22setAllColumnsShowFocusEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:127
// index:0
// Public
// bool allColumnsShowFocus()
func (this *QTreeView) AllColumnsShowFocus() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView19allColumnsShowFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:129
// index:0
// Public
// void setWordWrap(_Bool)
func (this *QTreeView) SetWordWrap(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setWordWrapEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:130
// index:0
// Public
// bool wordWrap()
func (this *QTreeView) WordWrap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8wordWrapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:132
// index:0
// Public
// void setTreePosition(int)
func (this *QTreeView) SetTreePosition(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setTreePositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:133
// index:0
// Public
// int treePosition()
func (this *QTreeView) TreePosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView12treePositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:135
// index:0
// Public virtual
// void keyboardSearch(const class QString &)
func (this *QTreeView) KeyboardSearch(search *qtcore.QString) {
	var convArg0 = search.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14keyboardSearchERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:137
// index:0
// Public virtual
// QRect visualRect(const class QModelIndex &)
func (this *QTreeView) VisualRect(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:138
// index:0
// Public virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QTreeView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:139
// index:0
// Public virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QTreeView) IndexAt(p *qtcore.QPoint) interface{} {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView7indexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:140
// index:0
// Public
// QModelIndex indexAbove(const class QModelIndex &)
func (this *QTreeView) IndexAbove(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10indexAboveERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:141
// index:0
// Public
// QModelIndex indexBelow(const class QModelIndex &)
func (this *QTreeView) IndexBelow(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10indexBelowERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:143
// index:0
// Public virtual
// void doItemsLayout()
func (this *QTreeView) DoItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13doItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:144
// index:0
// Public virtual
// void reset()
func (this *QTreeView) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:146
// index:0
// Public
// void sortByColumn(int, Qt::SortOrder)
func (this *QTreeView) SortByColumn(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:161
// index:1
// Public
// void sortByColumn(int)
func (this *QTreeView) SortByColumn_1(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:149
// index:0
// Public virtual
// void selectAll()
func (this *QTreeView) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:152
// index:0
// Public
// void expanded(const class QModelIndex &)
func (this *QTreeView) Expanded(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8expandedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:153
// index:0
// Public
// void collapsed(const class QModelIndex &)
func (this *QTreeView) Collapsed(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9collapsedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:156
// index:0
// Public
// void hideColumn(int)
func (this *QTreeView) HideColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10hideColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:157
// index:0
// Public
// void showColumn(int)
func (this *QTreeView) ShowColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10showColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:158
// index:0
// Public
// void expand(const class QModelIndex &)
func (this *QTreeView) Expand(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView6expandERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:159
// index:0
// Public
// void collapse(const class QModelIndex &)
func (this *QTreeView) Collapse(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8collapseERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:160
// index:0
// Public
// void resizeColumnToContents(int)
func (this *QTreeView) ResizeColumnToContents(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView22resizeColumnToContentsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:162
// index:0
// Public
// void expandAll()
func (this *QTreeView) ExpandAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9expandAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:163
// index:0
// Public
// void collapseAll()
func (this *QTreeView) CollapseAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11collapseAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:164
// index:0
// Public
// void expandToDepth(int)
func (this *QTreeView) ExpandToDepth(depth int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13expandToDepthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:167
// index:0
// Protected
// void columnResized(int, int, int)
func (this *QTreeView) ColumnResized(column int, oldSize int, newSize int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13columnResizedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &oldSize, &newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:168
// index:0
// Protected
// void columnCountChanged(int, int)
func (this *QTreeView) ColumnCountChanged(oldCount int, newCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18columnCountChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &oldCount, &newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:169
// index:0
// Protected
// void columnMoved()
func (this *QTreeView) ColumnMoved() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11columnMovedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:170
// index:0
// Protected
// void reexpand()
func (this *QTreeView) Reexpand() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8reexpandEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:171
// index:0
// Protected
// void rowsRemoved(const class QModelIndex &, int, int)
func (this *QTreeView) RowsRemoved(parent *qtcore.QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11rowsRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:175
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QTreeView) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:176
// index:0
// Protected virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QTreeView) RowsInserted(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:177
// index:0
// Protected virtual
// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QTreeView) RowsAboutToBeRemoved(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:180
// index:0
// Protected virtual
// int horizontalOffset()
func (this *QTreeView) HorizontalOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16horizontalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:181
// index:0
// Protected virtual
// int verticalOffset()
func (this *QTreeView) VerticalOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14verticalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:184
// index:0
// Protected virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QTreeView) VisualRegionForSelection(selection *qtcore.QItemSelection) interface{} {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:185
// index:0
// Protected virtual
// QModelIndexList selectedIndexes()
func (this *QTreeView) SelectedIndexes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15selectedIndexesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:187
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QTreeView) TimerEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:188
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QTreeView) PaintEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:190
// index:0
// Protected
// void drawTree(class QPainter *, const class QRegion &)
func (this *QTreeView) DrawTree(painter unsafe.Pointer, region *qtgui.QRegion) {
	var convArg1 = region.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:191
// index:0
// Protected virtual
// void drawRow(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QTreeView) DrawRow(painter unsafe.Pointer, options *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg1 = options.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:194
// index:0
// Protected virtual
// void drawBranches(class QPainter *, const class QRect &, const class QModelIndex &)
func (this *QTreeView) DrawBranches(painter unsafe.Pointer, rect *qtcore.QRect, index *qtcore.QModelIndex) {
	var convArg1 = rect.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:198
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QTreeView) MousePressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:199
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTreeView) MouseReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:200
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QTreeView) MouseDoubleClickEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:201
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QTreeView) MouseMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:202
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QTreeView) KeyPressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:204
// index:0
// Protected virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QTreeView) DragMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:206
// index:0
// Protected virtual
// bool viewportEvent(class QEvent *)
func (this *QTreeView) ViewportEvent(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:208
// index:0
// Protected virtual
// void updateGeometries()
func (this *QTreeView) UpdateGeometries() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16updateGeometriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:210
// index:0
// Protected virtual
// QSize viewportSizeHint()
func (this *QTreeView) ViewportSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16viewportSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:212
// index:0
// Protected virtual
// int sizeHintForColumn(int)
func (this *QTreeView) SizeHintForColumn(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView17sizeHintForColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:213
// index:0
// Protected
// int indexRowSizeHint(const class QModelIndex &)
func (this *QTreeView) IndexRowSizeHint(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:214
// index:0
// Protected
// int rowHeight(const class QModelIndex &)
func (this *QTreeView) RowHeight(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView9rowHeightERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:216
// index:0
// Protected virtual
// void horizontalScrollbarAction(int)
func (this *QTreeView) HorizontalScrollbarAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView25horizontalScrollbarActionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:218
// index:0
// Protected virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QTreeView) IsIndexHidden(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreeview.h:219
// index:0
// Protected virtual
// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QTreeView) SelectionChanged(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) {
	var convArg0 = selected.GetCthis()
	var convArg1 = deselected.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:221
// index:0
// Protected virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QTreeView) CurrentChanged(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end