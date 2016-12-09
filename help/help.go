// +build !minimal

package help

//#include <stdint.h>
//#include <stdlib.h>
//#include "help.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtHelp_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QHelpContentItem struct {
	ptr unsafe.Pointer
}

type QHelpContentItem_ITF interface {
	QHelpContentItem_PTR() *QHelpContentItem
}

func (p *QHelpContentItem) QHelpContentItem_PTR() *QHelpContentItem {
	return p
}

func (p *QHelpContentItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QHelpContentItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQHelpContentItem(ptr QHelpContentItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentItem_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentItemFromPointer(ptr unsafe.Pointer) *QHelpContentItem {
	var n = new(QHelpContentItem)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpContentItem) Child(row int) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Child(ptr.Pointer(), C.int(int32(row))))
	}
	return nil
}

func (ptr *QHelpContentItem) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_ChildCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentItem) ChildPosition(child QHelpContentItem_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_ChildPosition(ptr.Pointer(), PointerFromQHelpContentItem(child))))
	}
	return 0
}

func (ptr *QHelpContentItem) Parent() *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentItem) Row() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentItem_Row(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentItem) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpContentItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpContentItem) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QHelpContentItem_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentItem) DestroyQHelpContentItem() {
	if ptr.Pointer() != nil {
		C.QHelpContentItem_DestroyQHelpContentItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QHelpContentModel struct {
	core.QAbstractItemModel
}

type QHelpContentModel_ITF interface {
	core.QAbstractItemModel_ITF
	QHelpContentModel_PTR() *QHelpContentModel
}

func (p *QHelpContentModel) QHelpContentModel_PTR() *QHelpContentModel {
	return p
}

func (p *QHelpContentModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractItemModel_PTR().Pointer()
	}
	return nil
}

func (p *QHelpContentModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractItemModel_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpContentModel(ptr QHelpContentModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentModel_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentModelFromPointer(ptr unsafe.Pointer) *QHelpContentModel {
	var n = new(QHelpContentModel)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpContentModel) ColumnCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndex_ITF) *QHelpContentItem {
	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentModel_ContentItemAt(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

//export callbackQHelpContentModel_ContentsCreated
func callbackQHelpContentModel_ContentsCreated(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::contentsCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::contentsCreated", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreated() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::contentsCreated")
	}
}

func (ptr *QHelpContentModel) ContentsCreated() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreated(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_ContentsCreationStarted
func callbackQHelpContentModel_ContentsCreationStarted(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::contentsCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreationStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::contentsCreationStarted", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::contentsCreationStarted")
	}
}

func (ptr *QHelpContentModel) ContentsCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) CreateContents(customFilterName string) {
	if ptr.Pointer() != nil {
		var customFilterNameC = C.CString(customFilterName)
		defer C.free(unsafe.Pointer(customFilterNameC))
		C.QHelpContentModel_CreateContents(ptr.Pointer(), customFilterNameC)
	}
}

func (ptr *QHelpContentModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) IsCreatingContents() bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_IsCreatingContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) RowCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpContentModel) DestroyQHelpContentModel() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DestroyQHelpContentModel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpContentModel_Sibling
func callbackQHelpContentModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, index unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectSibling(f func(row int, column int, index *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::sibling", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSibling() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::sibling")
	}
}

func (ptr *QHelpContentModel) Sibling(row int, column int, index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_Sibling(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) SiblingDefault(row int, column int, index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Buddy
func callbackQHelpContentModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectBuddy(f func(index *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::buddy", f)
	}
}

func (ptr *QHelpContentModel) DisconnectBuddy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::buddy")
	}
}

func (ptr *QHelpContentModel) Buddy(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_Buddy(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_CanDropMimeData
func callbackQHelpContentModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectCanDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::canDropMimeData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectCanDropMimeData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::canDropMimeData")
	}
}

func (ptr *QHelpContentModel) CanDropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_CanDropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_CanFetchMore
func callbackQHelpContentModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectCanFetchMore(f func(parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::canFetchMore", f)
	}
}

func (ptr *QHelpContentModel) DisconnectCanFetchMore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::canFetchMore")
	}
}

func (ptr *QHelpContentModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_DropMimeData
func callbackQHelpContentModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::dropMimeData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectDropMimeData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::dropMimeData")
	}
}

func (ptr *QHelpContentModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_FetchMore
func callbackQHelpContentModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpContentModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpContentModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::fetchMore", f)
	}
}

func (ptr *QHelpContentModel) DisconnectFetchMore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::fetchMore")
	}
}

func (ptr *QHelpContentModel) FetchMore(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QHelpContentModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQHelpContentModel_Flags
func callbackQHelpContentModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::flags"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex) core.Qt__ItemFlag)(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectFlags(f func(index *core.QModelIndex) core.Qt__ItemFlag) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::flags", f)
	}
}

func (ptr *QHelpContentModel) DisconnectFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::flags")
	}
}

func (ptr *QHelpContentModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpContentModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QHelpContentModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpContentModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQHelpContentModel_HasChildren
func callbackQHelpContentModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectHasChildren(f func(parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::hasChildren", f)
	}
}

func (ptr *QHelpContentModel) DisconnectHasChildren() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::hasChildren")
	}
}

func (ptr *QHelpContentModel) HasChildren(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_HeaderData
func callbackQHelpContentModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::headerData"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int, core.Qt__Orientation, int) *core.QVariant)(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpContentModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QHelpContentModel) ConnectHeaderData(f func(section int, orientation core.Qt__Orientation, role int) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::headerData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectHeaderData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::headerData")
	}
}

func (ptr *QHelpContentModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentModel_HeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_InsertColumns
func callbackQHelpContentModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectInsertColumns(f func(column int, count int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::insertColumns", f)
	}
}

func (ptr *QHelpContentModel) DisconnectInsertColumns() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::insertColumns")
	}
}

func (ptr *QHelpContentModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_InsertColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_InsertRows
func callbackQHelpContentModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectInsertRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::insertRows", f)
	}
}

func (ptr *QHelpContentModel) DisconnectInsertRows() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::insertRows")
	}
}

func (ptr *QHelpContentModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_InsertRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_MimeTypes
func callbackQHelpContentModel_MimeTypes(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::mimeTypes"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQHelpContentModelFromPointer(ptr).MimeTypesDefault(), "|"))
}

func (ptr *QHelpContentModel) ConnectMimeTypes(f func() []string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::mimeTypes", f)
	}
}

func (ptr *QHelpContentModel) DisconnectMimeTypes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::mimeTypes")
	}
}

func (ptr *QHelpContentModel) MimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpContentModel_MimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpContentModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpContentModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQHelpContentModel_MoveColumns
func callbackQHelpContentModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpContentModel) ConnectMoveColumns(f func(sourceParent *core.QModelIndex, sourceColumn int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::moveColumns", f)
	}
}

func (ptr *QHelpContentModel) DisconnectMoveColumns() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::moveColumns")
	}
}

func (ptr *QHelpContentModel) MoveColumns(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_MoveColumns(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QHelpContentModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQHelpContentModel_MoveRows
func callbackQHelpContentModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpContentModel) ConnectMoveRows(f func(sourceParent *core.QModelIndex, sourceRow int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::moveRows", f)
	}
}

func (ptr *QHelpContentModel) DisconnectMoveRows() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::moveRows")
	}
}

func (ptr *QHelpContentModel) MoveRows(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_MoveRows(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QHelpContentModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQHelpContentModel_RemoveColumns
func callbackQHelpContentModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectRemoveColumns(f func(column int, count int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::removeColumns", f)
	}
}

func (ptr *QHelpContentModel) DisconnectRemoveColumns() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::removeColumns")
	}
}

func (ptr *QHelpContentModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_RemoveRows
func callbackQHelpContentModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpContentModel) ConnectRemoveRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::removeRows", f)
	}
}

func (ptr *QHelpContentModel) DisconnectRemoveRows() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::removeRows")
	}
}

func (ptr *QHelpContentModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_RemoveRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpContentModel_ResetInternalData
func callbackQHelpContentModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QHelpContentModel) ConnectResetInternalData(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::resetInternalData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectResetInternalData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::resetInternalData")
	}
}

func (ptr *QHelpContentModel) ResetInternalData() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ResetInternalData(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_Revert
func callbackQHelpContentModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::revert"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QHelpContentModel) ConnectRevert(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::revert", f)
	}
}

func (ptr *QHelpContentModel) DisconnectRevert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::revert")
	}
}

func (ptr *QHelpContentModel) Revert() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_Revert(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentModel_SetData
func callbackQHelpContentModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, *core.QVariant, int) bool)(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpContentModel) ConnectSetData(f func(index *core.QModelIndex, value *core.QVariant, role int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::setData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::setData")
	}
}

func (ptr *QHelpContentModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QHelpContentModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQHelpContentModel_SetHeaderData
func callbackQHelpContentModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, core.Qt__Orientation, *core.QVariant, int) bool)(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpContentModel) ConnectSetHeaderData(f func(section int, orientation core.Qt__Orientation, value *core.QVariant, role int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::setHeaderData", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSetHeaderData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::setHeaderData")
	}
}

func (ptr *QHelpContentModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SetHeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QHelpContentModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQHelpContentModel_Sort
func callbackQHelpContentModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpContentModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpContentModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::sort", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::sort")
	}
}

func (ptr *QHelpContentModel) Sort(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_Sort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

func (ptr *QHelpContentModel) SortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpContentModel_Span
func callbackQHelpContentModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQHelpContentModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentModel) ConnectSpan(f func(index *core.QModelIndex) *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::span", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSpan() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::span")
	}
}

func (ptr *QHelpContentModel) Span(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentModel_Span(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentModel_Submit
func callbackQHelpContentModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QHelpContentModel) ConnectSubmit(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::submit", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSubmit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::submit")
	}
}

func (ptr *QHelpContentModel) Submit() bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpContentModel_SupportedDragActions
func callbackQHelpContentModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QHelpContentModel) ConnectSupportedDragActions(f func() core.Qt__DropAction) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::supportedDragActions", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSupportedDragActions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::supportedDragActions")
	}
}

func (ptr *QHelpContentModel) SupportedDragActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpContentModel) SupportedDragActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpContentModel_SupportedDropActions
func callbackQHelpContentModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpContentModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QHelpContentModel) ConnectSupportedDropActions(f func() core.Qt__DropAction) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::supportedDropActions", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSupportedDropActions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::supportedDropActions")
	}
}

func (ptr *QHelpContentModel) SupportedDropActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpContentModel) SupportedDropActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpContentModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpContentModel_TimerEvent
func callbackQHelpContentModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::timerEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::timerEvent")
	}
}

func (ptr *QHelpContentModel) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpContentModel_ChildEvent
func callbackQHelpContentModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::childEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::childEvent")
	}
}

func (ptr *QHelpContentModel) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpContentModel_ConnectNotify
func callbackQHelpContentModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::connectNotify", f)
	}
}

func (ptr *QHelpContentModel) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::connectNotify")
	}
}

func (ptr *QHelpContentModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpContentModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentModel_CustomEvent
func callbackQHelpContentModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::customEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::customEvent")
	}
}

func (ptr *QHelpContentModel) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentModel_DeleteLater
func callbackQHelpContentModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpContentModel) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::deleteLater", f)
	}
}

func (ptr *QHelpContentModel) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::deleteLater")
	}
}

func (ptr *QHelpContentModel) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpContentModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpContentModel_DisconnectNotify
func callbackQHelpContentModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::disconnectNotify", f)
	}
}

func (ptr *QHelpContentModel) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::disconnectNotify")
	}
}

func (ptr *QHelpContentModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpContentModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentModel_Event
func callbackQHelpContentModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpContentModel) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::event", f)
	}
}

func (ptr *QHelpContentModel) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::event")
	}
}

func (ptr *QHelpContentModel) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpContentModel_EventFilter
func callbackQHelpContentModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::eventFilter", f)
	}
}

func (ptr *QHelpContentModel) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::eventFilter")
	}
}

func (ptr *QHelpContentModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpContentModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpContentModel_MetaObject
func callbackQHelpContentModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpContentModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpContentModel) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::metaObject", f)
	}
}

func (ptr *QHelpContentModel) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentModel::metaObject")
	}
}

func (ptr *QHelpContentModel) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpContentWidget struct {
	widgets.QTreeView
}

type QHelpContentWidget_ITF interface {
	widgets.QTreeView_ITF
	QHelpContentWidget_PTR() *QHelpContentWidget
}

func (p *QHelpContentWidget) QHelpContentWidget_PTR() *QHelpContentWidget {
	return p
}

func (p *QHelpContentWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QTreeView_PTR().Pointer()
	}
	return nil
}

func (p *QHelpContentWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QTreeView_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpContentWidget(ptr QHelpContentWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentWidgetFromPointer(ptr unsafe.Pointer) *QHelpContentWidget {
	var n = new(QHelpContentWidget)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpContentWidget) DestroyQHelpContentWidget() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QHelpContentWidget) IndexOf(link core.QUrl_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexOf(ptr.Pointer(), core.PointerFromQUrl(link)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_LinkActivated
func callbackQHelpContentWidget_LinkActivated(ptr unsafe.Pointer, link unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::linkActivated"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpContentWidget) ConnectLinkActivated(f func(link *core.QUrl)) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::linkActivated", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::linkActivated")
	}
}

func (ptr *QHelpContentWidget) LinkActivated(link core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

//export callbackQHelpContentWidget_Collapse
func callbackQHelpContentWidget_Collapse(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::collapse"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CollapseDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectCollapse(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::collapse", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCollapse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::collapse")
	}
}

func (ptr *QHelpContentWidget) Collapse(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Collapse(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) CollapseDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CollapseDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_Expand
func callbackQHelpContentWidget_Expand(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::expand"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectExpand(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::expand", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectExpand() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::expand")
	}
}

func (ptr *QHelpContentWidget) Expand(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Expand(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) ExpandDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_CollapseAll
func callbackQHelpContentWidget_CollapseAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::collapseAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CollapseAllDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectCollapseAll(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::collapseAll", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCollapseAll() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::collapseAll")
	}
}

func (ptr *QHelpContentWidget) CollapseAll() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CollapseAll(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) CollapseAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CollapseAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ColumnCountChanged
func callbackQHelpContentWidget_ColumnCountChanged(ptr unsafe.Pointer, oldCount C.int, newCount C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::columnCountChanged"); signal != nil {
		signal.(func(int, int))(int(int32(oldCount)), int(int32(newCount)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnCountChangedDefault(int(int32(oldCount)), int(int32(newCount)))
	}
}

func (ptr *QHelpContentWidget) ConnectColumnCountChanged(f func(oldCount int, newCount int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::columnCountChanged", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectColumnCountChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::columnCountChanged")
	}
}

func (ptr *QHelpContentWidget) ColumnCountChanged(oldCount int, newCount int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnCountChanged(ptr.Pointer(), C.int(int32(oldCount)), C.int(int32(newCount)))
	}
}

func (ptr *QHelpContentWidget) ColumnCountChangedDefault(oldCount int, newCount int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnCountChangedDefault(ptr.Pointer(), C.int(int32(oldCount)), C.int(int32(newCount)))
	}
}

//export callbackQHelpContentWidget_ColumnMoved
func callbackQHelpContentWidget_ColumnMoved(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::columnMoved"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnMovedDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectColumnMoved(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::columnMoved", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectColumnMoved() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::columnMoved")
	}
}

func (ptr *QHelpContentWidget) ColumnMoved() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnMoved(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ColumnMovedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnMovedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ColumnResized
func callbackQHelpContentWidget_ColumnResized(ptr unsafe.Pointer, column C.int, oldSize C.int, newSize C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::columnResized"); signal != nil {
		signal.(func(int, int, int))(int(int32(column)), int(int32(oldSize)), int(int32(newSize)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ColumnResizedDefault(int(int32(column)), int(int32(oldSize)), int(int32(newSize)))
	}
}

func (ptr *QHelpContentWidget) ConnectColumnResized(f func(column int, oldSize int, newSize int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::columnResized", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectColumnResized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::columnResized")
	}
}

func (ptr *QHelpContentWidget) ColumnResized(column int, oldSize int, newSize int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnResized(ptr.Pointer(), C.int(int32(column)), C.int(int32(oldSize)), C.int(int32(newSize)))
	}
}

func (ptr *QHelpContentWidget) ColumnResizedDefault(column int, oldSize int, newSize int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ColumnResizedDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(oldSize)), C.int(int32(newSize)))
	}
}

//export callbackQHelpContentWidget_CurrentChanged
func callbackQHelpContentWidget_CurrentChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QHelpContentWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::currentChanged", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::currentChanged")
	}
}

func (ptr *QHelpContentWidget) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QHelpContentWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

//export callbackQHelpContentWidget_DragMoveEvent
func callbackQHelpContentWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::dragMoveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::dragMoveEvent")
	}
}

func (ptr *QHelpContentWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpContentWidget_DrawBranches
func callbackQHelpContentWidget_DrawBranches(ptr unsafe.Pointer, painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::drawBranches"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRect, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DrawBranchesDefault(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectDrawBranches(f func(painter *gui.QPainter, rect *core.QRect, index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::drawBranches", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDrawBranches() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::drawBranches")
	}
}

func (ptr *QHelpContentWidget) DrawBranches(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawBranches(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) DrawBranchesDefault(painter gui.QPainter_ITF, rect core.QRect_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawBranchesDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_DrawRow
func callbackQHelpContentWidget_DrawRow(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::drawRow"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DrawRowDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectDrawRow(f func(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::drawRow", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDrawRow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::drawRow")
	}
}

func (ptr *QHelpContentWidget) DrawRow(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawRow(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) DrawRowDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DrawRowDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_ExpandAll
func callbackQHelpContentWidget_ExpandAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::expandAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandAllDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectExpandAll(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::expandAll", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectExpandAll() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::expandAll")
	}
}

func (ptr *QHelpContentWidget) ExpandAll() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandAll(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ExpandAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ExpandToDepth
func callbackQHelpContentWidget_ExpandToDepth(ptr unsafe.Pointer, depth C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::expandToDepth"); signal != nil {
		signal.(func(int))(int(int32(depth)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ExpandToDepthDefault(int(int32(depth)))
	}
}

func (ptr *QHelpContentWidget) ConnectExpandToDepth(f func(depth int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::expandToDepth", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectExpandToDepth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::expandToDepth")
	}
}

func (ptr *QHelpContentWidget) ExpandToDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandToDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

func (ptr *QHelpContentWidget) ExpandToDepthDefault(depth int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ExpandToDepthDefault(ptr.Pointer(), C.int(int32(depth)))
	}
}

//export callbackQHelpContentWidget_HideColumn
func callbackQHelpContentWidget_HideColumn(ptr unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::hideColumn"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideColumnDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ConnectHideColumn(f func(column int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::hideColumn", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHideColumn() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::hideColumn")
	}
}

func (ptr *QHelpContentWidget) HideColumn(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideColumn(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) HideColumnDefault(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideColumnDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_HorizontalOffset
func callbackQHelpContentWidget_HorizontalOffset(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::horizontalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHelpContentWidget) ConnectHorizontalOffset(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::horizontalOffset", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHorizontalOffset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::horizontalOffset")
	}
}

func (ptr *QHelpContentWidget) HorizontalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HorizontalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentWidget) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpContentWidget_IndexAt
func callbackQHelpContentWidget_IndexAt(ptr unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::indexAt"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QPoint) *core.QModelIndex)(core.NewQPointFromPointer(point)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentWidgetFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(point)))
}

func (ptr *QHelpContentWidget) ConnectIndexAt(f func(point *core.QPoint) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::indexAt", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectIndexAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::indexAt")
	}
}

func (ptr *QHelpContentWidget) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) IndexAtDefault(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(point)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_IsIndexHidden
func callbackQHelpContentWidget_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHelpContentWidget) ConnectIsIndexHidden(f func(index *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::isIndexHidden", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectIsIndexHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::isIndexHidden")
	}
}

func (ptr *QHelpContentWidget) IsIndexHidden(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_KeyPressEvent
func callbackQHelpContentWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::keyPressEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::keyPressEvent")
	}
}

func (ptr *QHelpContentWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpContentWidget_KeyboardSearch
func callbackQHelpContentWidget_KeyboardSearch(ptr unsafe.Pointer, search C.struct_QtHelp_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::keyboardSearch"); signal != nil {
		signal.(func(string))(cGoUnpackString(search))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyboardSearchDefault(cGoUnpackString(search))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyboardSearch(f func(search string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::keyboardSearch", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyboardSearch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::keyboardSearch")
	}
}

func (ptr *QHelpContentWidget) KeyboardSearch(search string) {
	if ptr.Pointer() != nil {
		var searchC = C.CString(search)
		defer C.free(unsafe.Pointer(searchC))
		C.QHelpContentWidget_KeyboardSearch(ptr.Pointer(), searchC)
	}
}

func (ptr *QHelpContentWidget) KeyboardSearchDefault(search string) {
	if ptr.Pointer() != nil {
		var searchC = C.CString(search)
		defer C.free(unsafe.Pointer(searchC))
		C.QHelpContentWidget_KeyboardSearchDefault(ptr.Pointer(), searchC)
	}
}

//export callbackQHelpContentWidget_MouseDoubleClickEvent
func callbackQHelpContentWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::mouseDoubleClickEvent")
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MouseMoveEvent
func callbackQHelpContentWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::mouseMoveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::mouseMoveEvent")
	}
}

func (ptr *QHelpContentWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MousePressEvent
func callbackQHelpContentWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::mousePressEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::mousePressEvent")
	}
}

func (ptr *QHelpContentWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MouseReleaseEvent
func callbackQHelpContentWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::mouseReleaseEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::mouseReleaseEvent")
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpContentWidget_MoveCursor
func callbackQHelpContentWidget_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::moveCursor"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(widgets.QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHelpContentWidgetFromPointer(ptr).MoveCursorDefault(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHelpContentWidget) ConnectMoveCursor(f func(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::moveCursor", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMoveCursor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::moveCursor")
	}
}

func (ptr *QHelpContentWidget) MoveCursor(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_MoveCursor(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpContentWidget_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_PaintEvent
func callbackQHelpContentWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::paintEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectPaintEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::paintEvent")
	}
}

func (ptr *QHelpContentWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpContentWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpContentWidget_Reset
func callbackQHelpContentWidget_Reset(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::reset"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectReset(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::reset", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::reset")
	}
}

func (ptr *QHelpContentWidget) Reset() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ResetDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResetDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ResizeColumnToContents
func callbackQHelpContentWidget_ResizeColumnToContents(ptr unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::resizeColumnToContents"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResizeColumnToContentsDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ConnectResizeColumnToContents(f func(column int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::resizeColumnToContents", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectResizeColumnToContents() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::resizeColumnToContents")
	}
}

func (ptr *QHelpContentWidget) ResizeColumnToContents(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeColumnToContents(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ResizeColumnToContentsDefault(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeColumnToContentsDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_RowsAboutToBeRemoved
func callbackQHelpContentWidget_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::rowsAboutToBeRemoved", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRowsAboutToBeRemoved() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::rowsAboutToBeRemoved")
	}
}

func (ptr *QHelpContentWidget) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_RowsInserted
func callbackQHelpContentWidget_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::rowsInserted", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRowsInserted() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::rowsInserted")
	}
}

func (ptr *QHelpContentWidget) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_RowsRemoved
func callbackQHelpContentWidget_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::rowsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RowsRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) ConnectRowsRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::rowsRemoved", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRowsRemoved() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::rowsRemoved")
	}
}

func (ptr *QHelpContentWidget) RowsRemoved(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpContentWidget) RowsRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RowsRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpContentWidget_ScrollContentsBy
func callbackQHelpContentWidget_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(int32(dx)), int(int32(dy)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *QHelpContentWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::scrollContentsBy", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollContentsBy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::scrollContentsBy")
	}
}

func (ptr *QHelpContentWidget) ScrollContentsBy(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsBy(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

func (ptr *QHelpContentWidget) ScrollContentsByDefault(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackQHelpContentWidget_ScrollTo
func callbackQHelpContentWidget_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpContentWidget) ConnectScrollTo(f func(index *core.QModelIndex, hint widgets.QAbstractItemView__ScrollHint)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::scrollTo", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollTo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::scrollTo")
	}
}

func (ptr *QHelpContentWidget) ScrollTo(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

func (ptr *QHelpContentWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQHelpContentWidget_SelectAll
func callbackQHelpContentWidget_SelectAll(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectSelectAll(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::selectAll", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSelectAll() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::selectAll")
	}
}

func (ptr *QHelpContentWidget) SelectAll() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) SelectAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SelectionChanged
func callbackQHelpContentWidget_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::selectionChanged"); signal != nil {
		signal.(func(*core.QItemSelection, *core.QItemSelection))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SelectionChangedDefault(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *QHelpContentWidget) ConnectSelectionChanged(f func(selected *core.QItemSelection, deselected *core.QItemSelection)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::selectionChanged", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::selectionChanged")
	}
}

func (ptr *QHelpContentWidget) SelectionChanged(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectionChanged(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

func (ptr *QHelpContentWidget) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectionChangedDefault(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

//export callbackQHelpContentWidget_SetModel
func callbackQHelpContentWidget_SetModel(ptr unsafe.Pointer, model unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpContentWidget) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setModel", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setModel")
	}
}

func (ptr *QHelpContentWidget) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpContentWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

//export callbackQHelpContentWidget_SetRootIndex
func callbackQHelpContentWidget_SetRootIndex(ptr unsafe.Pointer, index unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetRootIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setRootIndex", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetRootIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setRootIndex")
	}
}

func (ptr *QHelpContentWidget) SetRootIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_SetSelection
func callbackQHelpContentWidget_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, command C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpContentWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setSelection", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setSelection")
	}
}

func (ptr *QHelpContentWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

func (ptr *QHelpContentWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

//export callbackQHelpContentWidget_SetSelectionModel
func callbackQHelpContentWidget_SetSelectionModel(ptr unsafe.Pointer, selectionModel unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpContentWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setSelectionModel", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetSelectionModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setSelectionModel")
	}
}

func (ptr *QHelpContentWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpContentWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

//export callbackQHelpContentWidget_ShowColumn
func callbackQHelpContentWidget_ShowColumn(ptr unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::showColumn"); signal != nil {
		signal.(func(int))(int(int32(column)))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowColumnDefault(int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ConnectShowColumn(f func(column int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showColumn", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowColumn() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showColumn")
	}
}

func (ptr *QHelpContentWidget) ShowColumn(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowColumn(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QHelpContentWidget) ShowColumnDefault(column int) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowColumnDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQHelpContentWidget_SizeHintForColumn
func callbackQHelpContentWidget_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::sizeHintForColumn"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(column)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).SizeHintForColumnDefault(int(int32(column)))))
}

func (ptr *QHelpContentWidget) ConnectSizeHintForColumn(f func(column int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::sizeHintForColumn", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSizeHintForColumn() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::sizeHintForColumn")
	}
}

func (ptr *QHelpContentWidget) SizeHintForColumn(column int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForColumn(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

func (ptr *QHelpContentWidget) SizeHintForColumnDefault(column int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForColumnDefault(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

//export callbackQHelpContentWidget_UpdateGeometries
func callbackQHelpContentWidget_UpdateGeometries(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectUpdateGeometries(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::updateGeometries", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectUpdateGeometries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::updateGeometries")
	}
}

func (ptr *QHelpContentWidget) UpdateGeometries() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) UpdateGeometriesDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_VerticalOffset
func callbackQHelpContentWidget_VerticalOffset(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::verticalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHelpContentWidget) ConnectVerticalOffset(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::verticalOffset", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectVerticalOffset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::verticalOffset")
	}
}

func (ptr *QHelpContentWidget) VerticalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_VerticalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpContentWidget) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpContentWidget_ViewportEvent
func callbackQHelpContentWidget_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) ConnectViewportEvent(f func(event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::viewportEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectViewportEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::viewportEvent")
	}
}

func (ptr *QHelpContentWidget) ViewportEvent(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) ViewportEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_ViewportSizeHint
func callbackQHelpContentWidget_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::viewportSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QHelpContentWidget) ConnectViewportSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::viewportSizeHint", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectViewportSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::viewportSizeHint")
	}
}

func (ptr *QHelpContentWidget) ViewportSizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_ViewportSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) ViewportSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_ViewportSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_VisualRect
func callbackQHelpContentWidget_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::visualRect"); signal != nil {
		return core.PointerFromQRect(signal.(func(*core.QModelIndex) *core.QRect)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHelpContentWidgetFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpContentWidget) ConnectVisualRect(f func(index *core.QModelIndex) *core.QRect) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::visualRect", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectVisualRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::visualRect")
	}
}

func (ptr *QHelpContentWidget) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QHelpContentWidget_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QHelpContentWidget_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_VisualRegionForSelection
func callbackQHelpContentWidget_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion(signal.(func(*core.QItemSelection) *gui.QRegion)(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHelpContentWidgetFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHelpContentWidget) ConnectVisualRegionForSelection(f func(selection *core.QItemSelection) *gui.QRegion) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::visualRegionForSelection", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectVisualRegionForSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::visualRegionForSelection")
	}
}

func (ptr *QHelpContentWidget) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQRegionFromPointer(C.QHelpContentWidget_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQRegionFromPointer(C.QHelpContentWidget_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_DragLeaveEvent
func callbackQHelpContentWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::dragLeaveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::dragLeaveEvent")
	}
}

func (ptr *QHelpContentWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpContentWidget_ClearSelection
func callbackQHelpContentWidget_ClearSelection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::clearSelection"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ClearSelectionDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectClearSelection(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::clearSelection", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectClearSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::clearSelection")
	}
}

func (ptr *QHelpContentWidget) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ClearSelectionDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ClearSelectionDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_CloseEditor
func callbackQHelpContentWidget_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CloseEditorDefault(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	}
}

func (ptr *QHelpContentWidget) ConnectCloseEditor(f func(editor *widgets.QWidget, hint widgets.QAbstractItemDelegate__EndEditHint)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::closeEditor", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCloseEditor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::closeEditor")
	}
}

func (ptr *QHelpContentWidget) CloseEditor(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

func (ptr *QHelpContentWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

//export callbackQHelpContentWidget_CommitData
func callbackQHelpContentWidget_CommitData(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CommitDataDefault(widgets.NewQWidgetFromPointer(editor))
	}
}

func (ptr *QHelpContentWidget) ConnectCommitData(f func(editor *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::commitData", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCommitData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::commitData")
	}
}

func (ptr *QHelpContentWidget) CommitData(editor widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitData(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpContentWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

//export callbackQHelpContentWidget_DragEnterEvent
func callbackQHelpContentWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::dragEnterEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::dragEnterEvent")
	}
}

func (ptr *QHelpContentWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpContentWidget_DropEvent
func callbackQHelpContentWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::dropEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::dropEvent")
	}
}

func (ptr *QHelpContentWidget) DropEvent(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpContentWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpContentWidget_Edit
func callbackQHelpContentWidget_Edit(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::edit"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectEdit(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::edit", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEdit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::edit")
	}
}

func (ptr *QHelpContentWidget) Edit(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Edit(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) EditDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_Edit2
func callbackQHelpContentWidget_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, widgets.QAbstractItemView__EditTrigger, *core.QEvent) bool)(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) ConnectEdit2(f func(index *core.QModelIndex, trigger widgets.QAbstractItemView__EditTrigger, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::edit2", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEdit2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::edit2")
	}
}

func (ptr *QHelpContentWidget) Edit2(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_Edit2(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_EditorDestroyed
func callbackQHelpContentWidget_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EditorDestroyedDefault(core.NewQObjectFromPointer(editor))
	}
}

func (ptr *QHelpContentWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::editorDestroyed", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEditorDestroyed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::editorDestroyed")
	}
}

func (ptr *QHelpContentWidget) EditorDestroyed(editor core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpContentWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

//export callbackQHelpContentWidget_FocusInEvent
func callbackQHelpContentWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::focusInEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::focusInEvent")
	}
}

func (ptr *QHelpContentWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpContentWidget_FocusNextPrevChild
func callbackQHelpContentWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpContentWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::focusNextPrevChild", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusNextPrevChild() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::focusNextPrevChild")
	}
}

func (ptr *QHelpContentWidget) FocusNextPrevChild(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_FocusOutEvent
func callbackQHelpContentWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::focusOutEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::focusOutEvent")
	}
}

func (ptr *QHelpContentWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpContentWidget_InputMethodEvent
func callbackQHelpContentWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::inputMethodEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::inputMethodEvent")
	}
}

func (ptr *QHelpContentWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpContentWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpContentWidget_InputMethodQuery
func callbackQHelpContentWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpContentWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpContentWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::inputMethodQuery", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::inputMethodQuery")
	}
}

func (ptr *QHelpContentWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpContentWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_ResizeEvent
func callbackQHelpContentWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::resizeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::resizeEvent")
	}
}

func (ptr *QHelpContentWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpContentWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpContentWidget_ScrollToBottom
func callbackQHelpContentWidget_ScrollToBottom(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::scrollToBottom"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToBottomDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectScrollToBottom(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::scrollToBottom", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollToBottom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::scrollToBottom")
	}
}

func (ptr *QHelpContentWidget) ScrollToBottom() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToBottom(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ScrollToBottomDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToBottomDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ScrollToTop
func callbackQHelpContentWidget_ScrollToTop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::scrollToTop"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollToTopDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectScrollToTop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::scrollToTop", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollToTop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::scrollToTop")
	}
}

func (ptr *QHelpContentWidget) ScrollToTop() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToTop(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ScrollToTopDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollToTopDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SelectionCommand
func callbackQHelpContentWidget_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::selectionCommand"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex, *core.QEvent) core.QItemSelectionModel__SelectionFlag)(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewQHelpContentWidgetFromPointer(ptr).SelectionCommandDefault(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
}

func (ptr *QHelpContentWidget) ConnectSelectionCommand(f func(index *core.QModelIndex, event *core.QEvent) core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::selectionCommand", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSelectionCommand() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::selectionCommand")
	}
}

func (ptr *QHelpContentWidget) SelectionCommand(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpContentWidget_SelectionCommand(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

func (ptr *QHelpContentWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpContentWidget_SelectionCommandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackQHelpContentWidget_SetCurrentIndex
func callbackQHelpContentWidget_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setCurrentIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetCurrentIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectSetCurrentIndex(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setCurrentIndex", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetCurrentIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setCurrentIndex")
	}
}

func (ptr *QHelpContentWidget) SetCurrentIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetCurrentIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetCurrentIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_SizeHintForRow
func callbackQHelpContentWidget_SizeHintForRow(ptr unsafe.Pointer, row C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::sizeHintForRow"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(row)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *QHelpContentWidget) ConnectSizeHintForRow(f func(row int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::sizeHintForRow", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSizeHintForRow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::sizeHintForRow")
	}
}

func (ptr *QHelpContentWidget) SizeHintForRow(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForRow(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

func (ptr *QHelpContentWidget) SizeHintForRowDefault(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackQHelpContentWidget_StartDrag
func callbackQHelpContentWidget_StartDrag(ptr unsafe.Pointer, supportedActions C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpContentWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::startDrag", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectStartDrag() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::startDrag")
	}
}

func (ptr *QHelpContentWidget) StartDrag(supportedActions core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDrag(ptr.Pointer(), C.longlong(supportedActions))
	}
}

func (ptr *QHelpContentWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDragDefault(ptr.Pointer(), C.longlong(supportedActions))
	}
}

//export callbackQHelpContentWidget_Update
func callbackQHelpContentWidget_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::update"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpContentWidget) ConnectUpdate(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::update", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::update")
	}
}

func (ptr *QHelpContentWidget) Update(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Update(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpContentWidget) UpdateDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpContentWidget_ViewOptions
func callbackQHelpContentWidget_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::viewOptions"); signal != nil {
		return widgets.PointerFromQStyleOptionViewItem(signal.(func() *widgets.QStyleOptionViewItem)())
	}

	return widgets.PointerFromQStyleOptionViewItem(NewQHelpContentWidgetFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *QHelpContentWidget) ConnectViewOptions(f func() *widgets.QStyleOptionViewItem) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::viewOptions", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectViewOptions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::viewOptions")
	}
}

func (ptr *QHelpContentWidget) ViewOptions() *widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQStyleOptionViewItemFromPointer(C.QHelpContentWidget_ViewOptions(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQStyleOptionViewItemFromPointer(C.QHelpContentWidget_ViewOptionsDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_ContextMenuEvent
func callbackQHelpContentWidget_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::contextMenuEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::contextMenuEvent")
	}
}

func (ptr *QHelpContentWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpContentWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQHelpContentWidget_MinimumSizeHint
func callbackQHelpContentWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpContentWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::minimumSizeHint", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMinimumSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::minimumSizeHint")
	}
}

func (ptr *QHelpContentWidget) MinimumSizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_SetupViewport
func callbackQHelpContentWidget_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpContentWidget) ConnectSetupViewport(f func(viewport *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setupViewport", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetupViewport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setupViewport")
	}
}

func (ptr *QHelpContentWidget) SetupViewport(viewport widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewport(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpContentWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

//export callbackQHelpContentWidget_SizeHint
func callbackQHelpContentWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpContentWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpContentWidget) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::sizeHint", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::sizeHint")
	}
}

func (ptr *QHelpContentWidget) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpContentWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpContentWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpContentWidget_WheelEvent
func callbackQHelpContentWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::wheelEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::wheelEvent")
	}
}

func (ptr *QHelpContentWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpContentWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQHelpContentWidget_ChangeEvent
func callbackQHelpContentWidget_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpContentWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::changeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectChangeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::changeEvent")
	}
}

func (ptr *QHelpContentWidget) ChangeEvent(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpContentWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQHelpContentWidget_ActionEvent
func callbackQHelpContentWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::actionEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectActionEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::actionEvent")
	}
}

func (ptr *QHelpContentWidget) ActionEvent(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpContentWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpContentWidget_EnterEvent
func callbackQHelpContentWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::enterEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::enterEvent")
	}
}

func (ptr *QHelpContentWidget) EnterEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_HideEvent
func callbackQHelpContentWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::hideEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::hideEvent")
	}
}

func (ptr *QHelpContentWidget) HideEvent(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpContentWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpContentWidget_LeaveEvent
func callbackQHelpContentWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::leaveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::leaveEvent")
	}
}

func (ptr *QHelpContentWidget) LeaveEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_MoveEvent
func callbackQHelpContentWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::moveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::moveEvent")
	}
}

func (ptr *QHelpContentWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpContentWidget_SetEnabled
func callbackQHelpContentWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setEnabled", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setEnabled")
	}
}

func (ptr *QHelpContentWidget) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpContentWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpContentWidget_SetStyleSheet
func callbackQHelpContentWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QHelpContentWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setStyleSheet", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetStyleSheet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setStyleSheet")
	}
}

func (ptr *QHelpContentWidget) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpContentWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QHelpContentWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpContentWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQHelpContentWidget_SetVisible
func callbackQHelpContentWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setVisible", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setVisible")
	}
}

func (ptr *QHelpContentWidget) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QHelpContentWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpContentWidget_SetWindowModified
func callbackQHelpContentWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetWindowModified(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setWindowModified", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetWindowModified() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setWindowModified")
	}
}

func (ptr *QHelpContentWidget) SetWindowModified(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpContentWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpContentWidget_SetWindowTitle
func callbackQHelpContentWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QHelpContentWidget) ConnectSetWindowTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setWindowTitle", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetWindowTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setWindowTitle")
	}
}

func (ptr *QHelpContentWidget) SetWindowTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpContentWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QHelpContentWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpContentWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQHelpContentWidget_ShowEvent
func callbackQHelpContentWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showEvent")
	}
}

func (ptr *QHelpContentWidget) ShowEvent(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpContentWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpContentWidget_Close
func callbackQHelpContentWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpContentWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::close", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::close")
	}
}

func (ptr *QHelpContentWidget) Close() bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpContentWidget_CloseEvent
func callbackQHelpContentWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::closeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCloseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::closeEvent")
	}
}

func (ptr *QHelpContentWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpContentWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpContentWidget_HasHeightForWidth
func callbackQHelpContentWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpContentWidget) ConnectHasHeightForWidth(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::hasHeightForWidth", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHasHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::hasHeightForWidth")
	}
}

func (ptr *QHelpContentWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpContentWidget_HeightForWidth
func callbackQHelpContentWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpContentWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpContentWidget) ConnectHeightForWidth(f func(w int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::heightForWidth", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::heightForWidth")
	}
}

func (ptr *QHelpContentWidget) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QHelpContentWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpContentWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpContentWidget_Hide
func callbackQHelpContentWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::hide", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::hide")
	}
}

func (ptr *QHelpContentWidget) Hide() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_KeyReleaseEvent
func callbackQHelpContentWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::keyReleaseEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::keyReleaseEvent")
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpContentWidget_Lower
func callbackQHelpContentWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::lower", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::lower")
	}
}

func (ptr *QHelpContentWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_NativeEvent
func callbackQHelpContentWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QHelpContentWidget) ConnectNativeEvent(f func(eventType *core.QByteArray, message unsafe.Pointer, result int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::nativeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectNativeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::nativeEvent")
	}
}

func (ptr *QHelpContentWidget) NativeEvent(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_NativeEvent(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQHelpContentWidget_Raise
func callbackQHelpContentWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectRaise(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::raise", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRaise() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::raise")
	}
}

func (ptr *QHelpContentWidget) Raise() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_Repaint
func callbackQHelpContentWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectRepaint(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::repaint", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectRepaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::repaint")
	}
}

func (ptr *QHelpContentWidget) Repaint() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SetDisabled
func callbackQHelpContentWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetDisabled(f func(disable bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setDisabled", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetDisabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setDisabled")
	}
}

func (ptr *QHelpContentWidget) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QHelpContentWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpContentWidget_SetFocus2
func callbackQHelpContentWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpContentWidget) ConnectSetFocus2(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setFocus2", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetFocus2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setFocus2")
	}
}

func (ptr *QHelpContentWidget) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_SetHidden
func callbackQHelpContentWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpContentWidget) ConnectSetHidden(f func(hidden bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setHidden", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::setHidden")
	}
}

func (ptr *QHelpContentWidget) SetHidden(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QHelpContentWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpContentWidget_Show
func callbackQHelpContentWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::show", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::show")
	}
}

func (ptr *QHelpContentWidget) Show() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Show(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowFullScreen
func callbackQHelpContentWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShowFullScreen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showFullScreen", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowFullScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showFullScreen")
	}
}

func (ptr *QHelpContentWidget) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowMaximized
func callbackQHelpContentWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showMaximized", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showMaximized")
	}
}

func (ptr *QHelpContentWidget) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowMinimized
func callbackQHelpContentWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShowMinimized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showMinimized", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowMinimized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showMinimized")
	}
}

func (ptr *QHelpContentWidget) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ShowNormal
func callbackQHelpContentWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectShowNormal(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showNormal", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowNormal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::showNormal")
	}
}

func (ptr *QHelpContentWidget) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_TabletEvent
func callbackQHelpContentWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::tabletEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectTabletEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::tabletEvent")
	}
}

func (ptr *QHelpContentWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpContentWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpContentWidget_UpdateMicroFocus
func callbackQHelpContentWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectUpdateMicroFocus(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::updateMicroFocus", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectUpdateMicroFocus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::updateMicroFocus")
	}
}

func (ptr *QHelpContentWidget) UpdateMicroFocus() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpContentWidget_ChildEvent
func callbackQHelpContentWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::childEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::childEvent")
	}
}

func (ptr *QHelpContentWidget) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpContentWidget_ConnectNotify
func callbackQHelpContentWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::connectNotify", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::connectNotify")
	}
}

func (ptr *QHelpContentWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpContentWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentWidget_CustomEvent
func callbackQHelpContentWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::customEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::customEvent")
	}
}

func (ptr *QHelpContentWidget) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpContentWidget_DeleteLater
func callbackQHelpContentWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpContentWidget) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::deleteLater", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::deleteLater")
	}
}

func (ptr *QHelpContentWidget) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpContentWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpContentWidget_DisconnectNotify
func callbackQHelpContentWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpContentWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::disconnectNotify", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::disconnectNotify")
	}
}

func (ptr *QHelpContentWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpContentWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpContentWidget_EventFilter
func callbackQHelpContentWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpContentWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpContentWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::eventFilter", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::eventFilter")
	}
}

func (ptr *QHelpContentWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpContentWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpContentWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpContentWidget_MetaObject
func callbackQHelpContentWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpContentWidget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpContentWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpContentWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::metaObject", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpContentWidget::metaObject")
	}
}

func (ptr *QHelpContentWidget) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpContentWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpEngine struct {
	QHelpEngineCore
}

type QHelpEngine_ITF interface {
	QHelpEngineCore_ITF
	QHelpEngine_PTR() *QHelpEngine
}

func (p *QHelpEngine) QHelpEngine_PTR() *QHelpEngine {
	return p
}

func (p *QHelpEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QHelpEngineCore_PTR().Pointer()
	}
	return nil
}

func (p *QHelpEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QHelpEngineCore_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpEngine(ptr QHelpEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineFromPointer(ptr unsafe.Pointer) *QHelpEngine {
	var n = new(QHelpEngine)
	n.SetPointer(ptr)
	return n
}
func NewQHelpEngine(collectionFile string, parent core.QObject_ITF) *QHelpEngine {
	var collectionFileC = C.CString(collectionFile)
	defer C.free(unsafe.Pointer(collectionFileC))
	var tmpValue = NewQHelpEngineFromPointer(C.QHelpEngine_NewQHelpEngine(collectionFileC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHelpEngine) ContentModel() *QHelpContentModel {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpContentModelFromPointer(C.QHelpEngine_ContentModel(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngine) ContentWidget() *QHelpContentWidget {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpContentWidgetFromPointer(C.QHelpEngine_ContentWidget(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngine) IndexModel() *QHelpIndexModel {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpIndexModelFromPointer(C.QHelpEngine_IndexModel(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngine) IndexWidget() *QHelpIndexWidget {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpIndexWidgetFromPointer(C.QHelpEngine_IndexWidget(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngine) SearchEngine() *QHelpSearchEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpSearchEngineFromPointer(C.QHelpEngine_SearchEngine(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngine) DestroyQHelpEngine() {
	if ptr.Pointer() != nil {
		C.QHelpEngine_DestroyQHelpEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpEngine_TimerEvent
func callbackQHelpEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::timerEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::timerEvent")
	}
}

func (ptr *QHelpEngine) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpEngine_ChildEvent
func callbackQHelpEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::childEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::childEvent")
	}
}

func (ptr *QHelpEngine) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpEngine_ConnectNotify
func callbackQHelpEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::connectNotify", f)
	}
}

func (ptr *QHelpEngine) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::connectNotify")
	}
}

func (ptr *QHelpEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngine_CustomEvent
func callbackQHelpEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::customEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::customEvent")
	}
}

func (ptr *QHelpEngine) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpEngine_DeleteLater
func callbackQHelpEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpEngine) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::deleteLater", f)
	}
}

func (ptr *QHelpEngine) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::deleteLater")
	}
}

func (ptr *QHelpEngine) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpEngine_DisconnectNotify
func callbackQHelpEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::disconnectNotify", f)
	}
}

func (ptr *QHelpEngine) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::disconnectNotify")
	}
}

func (ptr *QHelpEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngine_Event
func callbackQHelpEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::event", f)
	}
}

func (ptr *QHelpEngine) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::event")
	}
}

func (ptr *QHelpEngine) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpEngine_EventFilter
func callbackQHelpEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::eventFilter", f)
	}
}

func (ptr *QHelpEngine) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::eventFilter")
	}
}

func (ptr *QHelpEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpEngine_MetaObject
func callbackQHelpEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::metaObject", f)
	}
}

func (ptr *QHelpEngine) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngine::metaObject")
	}
}

func (ptr *QHelpEngine) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpEngineCore struct {
	core.QObject
}

type QHelpEngineCore_ITF interface {
	core.QObject_ITF
	QHelpEngineCore_PTR() *QHelpEngineCore
}

func (p *QHelpEngineCore) QHelpEngineCore_PTR() *QHelpEngineCore {
	return p
}

func (p *QHelpEngineCore) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QHelpEngineCore) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpEngineCore(ptr QHelpEngineCore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngineCore_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineCoreFromPointer(ptr unsafe.Pointer) *QHelpEngineCore {
	var n = new(QHelpEngineCore)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpEngineCore) AutoSaveFilter() bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AutoSaveFilter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CollectionFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_CollectionFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) CurrentFilter() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_CurrentFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) SetAutoSaveFilter(save bool) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetAutoSaveFilter(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(save))))
	}
}

func (ptr *QHelpEngineCore) SetCollectionFile(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QHelpEngineCore_SetCollectionFile(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QHelpEngineCore) SetCurrentFilter(filterName string) {
	if ptr.Pointer() != nil {
		var filterNameC = C.CString(filterName)
		defer C.free(unsafe.Pointer(filterNameC))
		C.QHelpEngineCore_SetCurrentFilter(ptr.Pointer(), filterNameC)
	}
}

func NewQHelpEngineCore(collectionFile string, parent core.QObject_ITF) *QHelpEngineCore {
	var collectionFileC = C.CString(collectionFile)
	defer C.free(unsafe.Pointer(collectionFileC))
	var tmpValue = NewQHelpEngineCoreFromPointer(C.QHelpEngineCore_NewQHelpEngineCore(collectionFileC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHelpEngineCore) AddCustomFilter(filterName string, attributes []string) bool {
	if ptr.Pointer() != nil {
		var filterNameC = C.CString(filterName)
		defer C.free(unsafe.Pointer(filterNameC))
		var attributesC = C.CString(strings.Join(attributes, "|"))
		defer C.free(unsafe.Pointer(attributesC))
		return C.QHelpEngineCore_AddCustomFilter(ptr.Pointer(), filterNameC, attributesC) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CopyCollectionFile(fileName string) bool {
	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		return C.QHelpEngineCore_CopyCollectionFile(ptr.Pointer(), fileNameC) != 0
	}
	return false
}

//export callbackQHelpEngineCore_CurrentFilterChanged
func callbackQHelpEngineCore_CurrentFilterChanged(ptr unsafe.Pointer, newFilter C.struct_QtHelp_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::currentFilterChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(newFilter))
	}

}

func (ptr *QHelpEngineCore) ConnectCurrentFilterChanged(f func(newFilter string)) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectCurrentFilterChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::currentFilterChanged", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCurrentFilterChanged() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectCurrentFilterChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::currentFilterChanged")
	}
}

func (ptr *QHelpEngineCore) CurrentFilterChanged(newFilter string) {
	if ptr.Pointer() != nil {
		var newFilterC = C.CString(newFilter)
		defer C.free(unsafe.Pointer(newFilterC))
		C.QHelpEngineCore_CurrentFilterChanged(ptr.Pointer(), newFilterC)
	}
}

func (ptr *QHelpEngineCore) CustomFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_CustomFilters(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		var tmpValue = core.NewQVariantFromPointer(C.QHelpEngineCore_CustomValue(ptr.Pointer(), keyC, core.PointerFromQVariant(defaultValue)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) DocumentationFileName(namespaceName string) string {
	if ptr.Pointer() != nil {
		var namespaceNameC = C.CString(namespaceName)
		defer C.free(unsafe.Pointer(namespaceNameC))
		return cGoUnpackString(C.QHelpEngineCore_DocumentationFileName(ptr.Pointer(), namespaceNameC))
	}
	return ""
}

func (ptr *QHelpEngineCore) Error() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QHelpEngineCore_Error(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) FileData(url core.QUrl_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QHelpEngineCore_FileData(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) Files(namespaceName string, filterAttributes []string, extensionFilter string) []*core.QUrl {
	if ptr.Pointer() != nil {
		var namespaceNameC = C.CString(namespaceName)
		defer C.free(unsafe.Pointer(namespaceNameC))
		var filterAttributesC = C.CString(strings.Join(filterAttributes, "|"))
		defer C.free(unsafe.Pointer(filterAttributesC))
		var extensionFilterC = C.CString(extensionFilter)
		defer C.free(unsafe.Pointer(extensionFilterC))
		return func(l C.struct_QtHelp_PackedList) []*core.QUrl {
			var out = make([]*core.QUrl, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQHelpEngineCoreFromPointer(l.data).files_atList(i)
			}
			return out
		}(C.QHelpEngineCore_Files(ptr.Pointer(), namespaceNameC, filterAttributesC, extensionFilterC))
	}
	return nil
}

func (ptr *QHelpEngineCore) FilterAttributes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_FilterAttributes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FilterAttributes2(filterName string) []string {
	if ptr.Pointer() != nil {
		var filterNameC = C.CString(filterName)
		defer C.free(unsafe.Pointer(filterNameC))
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_FilterAttributes2(ptr.Pointer(), filterNameC)), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FindFile(url core.QUrl_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QHelpEngineCore_FindFile(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func QHelpEngineCore_MetaData(documentationFileName string, name string) *core.QVariant {
	var documentationFileNameC = C.CString(documentationFileName)
	defer C.free(unsafe.Pointer(documentationFileNameC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(documentationFileNameC, nameC))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QHelpEngineCore) MetaData(documentationFileName string, name string) *core.QVariant {
	var documentationFileNameC = C.CString(documentationFileName)
	defer C.free(unsafe.Pointer(documentationFileNameC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(documentationFileNameC, nameC))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QHelpEngineCore_NamespaceName(documentationFileName string) string {
	var documentationFileNameC = C.CString(documentationFileName)
	defer C.free(unsafe.Pointer(documentationFileNameC))
	return cGoUnpackString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(documentationFileNameC))
}

func (ptr *QHelpEngineCore) NamespaceName(documentationFileName string) string {
	var documentationFileNameC = C.CString(documentationFileName)
	defer C.free(unsafe.Pointer(documentationFileNameC))
	return cGoUnpackString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(documentationFileNameC))
}

//export callbackQHelpEngineCore_ReadersAboutToBeInvalidated
func callbackQHelpEngineCore_ReadersAboutToBeInvalidated(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::readersAboutToBeInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectReadersAboutToBeInvalidated(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::readersAboutToBeInvalidated", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectReadersAboutToBeInvalidated() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::readersAboutToBeInvalidated")
	}
}

func (ptr *QHelpEngineCore) ReadersAboutToBeInvalidated() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ReadersAboutToBeInvalidated(ptr.Pointer())
	}
}

func (ptr *QHelpEngineCore) RegisterDocumentation(documentationFileName string) bool {
	if ptr.Pointer() != nil {
		var documentationFileNameC = C.CString(documentationFileName)
		defer C.free(unsafe.Pointer(documentationFileNameC))
		return C.QHelpEngineCore_RegisterDocumentation(ptr.Pointer(), documentationFileNameC) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_RegisteredDocumentations(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) RemoveCustomFilter(filterName string) bool {
	if ptr.Pointer() != nil {
		var filterNameC = C.CString(filterName)
		defer C.free(unsafe.Pointer(filterNameC))
		return C.QHelpEngineCore_RemoveCustomFilter(ptr.Pointer(), filterNameC) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RemoveCustomValue(key string) bool {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		return C.QHelpEngineCore_RemoveCustomValue(ptr.Pointer(), keyC) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetCustomValue(key string, value core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		return C.QHelpEngineCore_SetCustomValue(ptr.Pointer(), keyC, core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetupData() bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetupData(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpEngineCore_SetupFinished
func callbackQHelpEngineCore_SetupFinished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::setupFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::setupFinished", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupFinished() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::setupFinished")
	}
}

func (ptr *QHelpEngineCore) SetupFinished() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetupFinished(ptr.Pointer())
	}
}

//export callbackQHelpEngineCore_SetupStarted
func callbackQHelpEngineCore_SetupStarted(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::setupStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::setupStarted", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupStarted() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::setupStarted")
	}
}

func (ptr *QHelpEngineCore) SetupStarted() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetupStarted(ptr.Pointer())
	}
}

func (ptr *QHelpEngineCore) UnregisterDocumentation(namespaceName string) bool {
	if ptr.Pointer() != nil {
		var namespaceNameC = C.CString(namespaceName)
		defer C.free(unsafe.Pointer(namespaceNameC))
		return C.QHelpEngineCore_UnregisterDocumentation(ptr.Pointer(), namespaceNameC) != 0
	}
	return false
}

//export callbackQHelpEngineCore_Warning
func callbackQHelpEngineCore_Warning(ptr unsafe.Pointer, msg C.struct_QtHelp_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::warning"); signal != nil {
		signal.(func(string))(cGoUnpackString(msg))
	}

}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectWarning(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::warning", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectWarning() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectWarning(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::warning")
	}
}

func (ptr *QHelpEngineCore) Warning(msg string) {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QHelpEngineCore_Warning(ptr.Pointer(), msgC)
	}
}

//export callbackQHelpEngineCore_DestroyQHelpEngineCore
func callbackQHelpEngineCore_DestroyQHelpEngineCore(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::~QHelpEngineCore"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DestroyQHelpEngineCoreDefault()
	}
}

func (ptr *QHelpEngineCore) ConnectDestroyQHelpEngineCore(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::~QHelpEngineCore", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectDestroyQHelpEngineCore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::~QHelpEngineCore")
	}
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCore() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DestroyQHelpEngineCore(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCoreDefault() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DestroyQHelpEngineCoreDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngineCore) files_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QHelpEngineCore_files_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpEngineCore) filterAttributeSets_atList(i int) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpEngineCore_filterAttributeSets_atList(ptr.Pointer(), C.int(int32(i)))), "|")
	}
	return make([]string, 0)
}

//export callbackQHelpEngineCore_TimerEvent
func callbackQHelpEngineCore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::timerEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::timerEvent")
	}
}

func (ptr *QHelpEngineCore) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngineCore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpEngineCore_ChildEvent
func callbackQHelpEngineCore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::childEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::childEvent")
	}
}

func (ptr *QHelpEngineCore) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngineCore) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpEngineCore_ConnectNotify
func callbackQHelpEngineCore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngineCore) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::connectNotify", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::connectNotify")
	}
}

func (ptr *QHelpEngineCore) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpEngineCore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngineCore_CustomEvent
func callbackQHelpEngineCore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::customEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::customEvent")
	}
}

func (ptr *QHelpEngineCore) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpEngineCore) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpEngineCore_DeleteLater
func callbackQHelpEngineCore_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpEngineCore) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::deleteLater", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::deleteLater")
	}
}

func (ptr *QHelpEngineCore) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngineCore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpEngineCore_DisconnectNotify
func callbackQHelpEngineCore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpEngineCore) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::disconnectNotify", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::disconnectNotify")
	}
}

func (ptr *QHelpEngineCore) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpEngineCore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpEngineCore_Event
func callbackQHelpEngineCore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineCoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpEngineCore) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::event", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::event")
	}
}

func (ptr *QHelpEngineCore) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpEngineCore_EventFilter
func callbackQHelpEngineCore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpEngineCoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpEngineCore) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::eventFilter", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::eventFilter")
	}
}

func (ptr *QHelpEngineCore) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpEngineCore_MetaObject
func callbackQHelpEngineCore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpEngineCore::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpEngineCoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpEngineCore) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::metaObject", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpEngineCore::metaObject")
	}
}

func (ptr *QHelpEngineCore) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngineCore_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngineCore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpEngineCore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpIndexModel struct {
	core.QStringListModel
}

type QHelpIndexModel_ITF interface {
	core.QStringListModel_ITF
	QHelpIndexModel_PTR() *QHelpIndexModel
}

func (p *QHelpIndexModel) QHelpIndexModel_PTR() *QHelpIndexModel {
	return p
}

func (p *QHelpIndexModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QStringListModel_PTR().Pointer()
	}
	return nil
}

func (p *QHelpIndexModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QStringListModel_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpIndexModel(ptr QHelpIndexModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexModel_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexModelFromPointer(ptr unsafe.Pointer) *QHelpIndexModel {
	var n = new(QHelpIndexModel)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpIndexModel) DestroyQHelpIndexModel() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func (ptr *QHelpIndexModel) CreateIndex(customFilterName string) {
	if ptr.Pointer() != nil {
		var customFilterNameC = C.CString(customFilterName)
		defer C.free(unsafe.Pointer(customFilterNameC))
		C.QHelpIndexModel_CreateIndex(ptr.Pointer(), customFilterNameC)
	}
}

func (ptr *QHelpIndexModel) Filter(filter string, wildcard string) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
		var wildcardC = C.CString(wildcard)
		defer C.free(unsafe.Pointer(wildcardC))
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Filter(ptr.Pointer(), filterC, wildcardC))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_IndexCreated
func callbackQHelpIndexModel_IndexCreated(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::indexCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::indexCreated", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::indexCreated")
	}
}

func (ptr *QHelpIndexModel) IndexCreated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreated(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_IndexCreationStarted
func callbackQHelpIndexModel_IndexCreationStarted(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::indexCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreationStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::indexCreationStarted", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::indexCreationStarted")
	}
}

func (ptr *QHelpIndexModel) IndexCreationStarted() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_IsCreatingIndex(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Data
func callbackQHelpIndexModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::data"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*core.QModelIndex, int) *core.QVariant)(core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpIndexModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *QHelpIndexModel) ConnectData(f func(index *core.QModelIndex, role int) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::data", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::data")
	}
}

func (ptr *QHelpIndexModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Flags
func callbackQHelpIndexModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::flags"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex) core.Qt__ItemFlag)(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ConnectFlags(f func(index *core.QModelIndex) core.Qt__ItemFlag) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::flags", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::flags")
	}
}

func (ptr *QHelpIndexModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpIndexModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QHelpIndexModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QHelpIndexModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQHelpIndexModel_InsertRows
func callbackQHelpIndexModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectInsertRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::insertRows", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectInsertRows() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::insertRows")
	}
}

func (ptr *QHelpIndexModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_InsertRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_RemoveRows
func callbackQHelpIndexModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectRemoveRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::removeRows", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRemoveRows() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::removeRows")
	}
}

func (ptr *QHelpIndexModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_RemoveRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_RowCount
func callbackQHelpIndexModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::rowCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpIndexModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpIndexModel) ConnectRowCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::rowCount", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRowCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::rowCount")
	}
}

func (ptr *QHelpIndexModel) RowCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpIndexModel) RowCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_RowCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpIndexModel_SetData
func callbackQHelpIndexModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, *core.QVariant, int) bool)(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpIndexModel) ConnectSetData(f func(index *core.QModelIndex, value *core.QVariant, role int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::setData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::setData")
	}
}

func (ptr *QHelpIndexModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Sibling
func callbackQHelpIndexModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QHelpIndexModel) ConnectSibling(f func(row int, column int, idx *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::sibling", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSibling() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::sibling")
	}
}

func (ptr *QHelpIndexModel) Sibling(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Sibling(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Sort
func callbackQHelpIndexModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQHelpIndexModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpIndexModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::sort", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::sort")
	}
}

func (ptr *QHelpIndexModel) Sort(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_Sort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

func (ptr *QHelpIndexModel) SortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQHelpIndexModel_SupportedDropActions
func callbackQHelpIndexModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QHelpIndexModel) ConnectSupportedDropActions(f func() core.Qt__DropAction) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::supportedDropActions", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSupportedDropActions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::supportedDropActions")
	}
}

func (ptr *QHelpIndexModel) SupportedDropActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpIndexModel) SupportedDropActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpIndexModel_Index
func callbackQHelpIndexModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::index"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QHelpIndexModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::index", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::index")
	}
}

func (ptr *QHelpIndexModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_DropMimeData
func callbackQHelpIndexModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::dropMimeData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectDropMimeData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::dropMimeData")
	}
}

func (ptr *QHelpIndexModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Buddy
func callbackQHelpIndexModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ConnectBuddy(f func(index *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::buddy", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectBuddy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::buddy")
	}
}

func (ptr *QHelpIndexModel) Buddy(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Buddy(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_CanDropMimeData
func callbackQHelpIndexModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectCanDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::canDropMimeData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectCanDropMimeData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::canDropMimeData")
	}
}

func (ptr *QHelpIndexModel) CanDropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_CanDropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_CanFetchMore
func callbackQHelpIndexModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectCanFetchMore(f func(parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::canFetchMore", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectCanFetchMore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::canFetchMore")
	}
}

func (ptr *QHelpIndexModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ColumnCount
func callbackQHelpIndexModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::columnCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQHelpIndexModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QHelpIndexModel) ConnectColumnCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::columnCount", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectColumnCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::columnCount")
	}
}

func (ptr *QHelpIndexModel) ColumnCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QHelpIndexModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexModel_ColumnCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQHelpIndexModel_FetchMore
func callbackQHelpIndexModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQHelpIndexModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QHelpIndexModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::fetchMore", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectFetchMore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::fetchMore")
	}
}

func (ptr *QHelpIndexModel) FetchMore(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QHelpIndexModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQHelpIndexModel_HasChildren
func callbackQHelpIndexModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectHasChildren(f func(parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::hasChildren", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectHasChildren() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::hasChildren")
	}
}

func (ptr *QHelpIndexModel) HasChildren(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_HeaderData
func callbackQHelpIndexModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::headerData"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int, core.Qt__Orientation, int) *core.QVariant)(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQHelpIndexModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QHelpIndexModel) ConnectHeaderData(f func(section int, orientation core.Qt__Orientation, role int) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::headerData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectHeaderData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::headerData")
	}
}

func (ptr *QHelpIndexModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexModel_HeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_InsertColumns
func callbackQHelpIndexModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectInsertColumns(f func(column int, count int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::insertColumns", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectInsertColumns() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::insertColumns")
	}
}

func (ptr *QHelpIndexModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_InsertColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MimeTypes
func callbackQHelpIndexModel_MimeTypes(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::mimeTypes"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQHelpIndexModelFromPointer(ptr).MimeTypesDefault(), "|"))
}

func (ptr *QHelpIndexModel) ConnectMimeTypes(f func() []string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::mimeTypes", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectMimeTypes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::mimeTypes")
	}
}

func (ptr *QHelpIndexModel) MimeTypes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpIndexModel_MimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpIndexModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpIndexModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQHelpIndexModel_MoveColumns
func callbackQHelpIndexModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpIndexModel) ConnectMoveColumns(f func(sourceParent *core.QModelIndex, sourceColumn int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::moveColumns", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectMoveColumns() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::moveColumns")
	}
}

func (ptr *QHelpIndexModel) MoveColumns(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_MoveColumns(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MoveRows
func callbackQHelpIndexModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QHelpIndexModel) ConnectMoveRows(f func(sourceParent *core.QModelIndex, sourceRow int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::moveRows", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectMoveRows() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::moveRows")
	}
}

func (ptr *QHelpIndexModel) MoveRows(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_MoveRows(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Parent
func callbackQHelpIndexModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::parent"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ConnectParent(f func(index *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::parent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectParent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::parent")
	}
}

func (ptr *QHelpIndexModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_RemoveColumns
func callbackQHelpIndexModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QHelpIndexModel) ConnectRemoveColumns(f func(column int, count int, parent *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::removeColumns", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRemoveColumns() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::removeColumns")
	}
}

func (ptr *QHelpIndexModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_ResetInternalData
func callbackQHelpIndexModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QHelpIndexModel) ConnectResetInternalData(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::resetInternalData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectResetInternalData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::resetInternalData")
	}
}

func (ptr *QHelpIndexModel) ResetInternalData() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ResetInternalData(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_Revert
func callbackQHelpIndexModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::revert"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QHelpIndexModel) ConnectRevert(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::revert", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRevert() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::revert")
	}
}

func (ptr *QHelpIndexModel) Revert() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_Revert(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexModel_SetHeaderData
func callbackQHelpIndexModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, core.Qt__Orientation, *core.QVariant, int) bool)(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QHelpIndexModel) ConnectSetHeaderData(f func(section int, orientation core.Qt__Orientation, value *core.QVariant, role int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::setHeaderData", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSetHeaderData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::setHeaderData")
	}
}

func (ptr *QHelpIndexModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SetHeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQHelpIndexModel_Span
func callbackQHelpIndexModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQHelpIndexModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexModel) ConnectSpan(f func(index *core.QModelIndex) *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::span", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSpan() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::span")
	}
}

func (ptr *QHelpIndexModel) Span(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexModel_Span(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexModel_Submit
func callbackQHelpIndexModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QHelpIndexModel) ConnectSubmit(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::submit", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSubmit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::submit")
	}
}

func (ptr *QHelpIndexModel) Submit() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpIndexModel_SupportedDragActions
func callbackQHelpIndexModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQHelpIndexModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QHelpIndexModel) ConnectSupportedDragActions(f func() core.Qt__DropAction) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::supportedDragActions", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSupportedDragActions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::supportedDragActions")
	}
}

func (ptr *QHelpIndexModel) SupportedDragActions() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpIndexModel) SupportedDragActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QHelpIndexModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQHelpIndexModel_TimerEvent
func callbackQHelpIndexModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::timerEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::timerEvent")
	}
}

func (ptr *QHelpIndexModel) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpIndexModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpIndexModel_ChildEvent
func callbackQHelpIndexModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::childEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::childEvent")
	}
}

func (ptr *QHelpIndexModel) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpIndexModel_ConnectNotify
func callbackQHelpIndexModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::connectNotify", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::connectNotify")
	}
}

func (ptr *QHelpIndexModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpIndexModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexModel_CustomEvent
func callbackQHelpIndexModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::customEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::customEvent")
	}
}

func (ptr *QHelpIndexModel) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexModel_DeleteLater
func callbackQHelpIndexModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpIndexModel) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::deleteLater", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::deleteLater")
	}
}

func (ptr *QHelpIndexModel) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpIndexModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpIndexModel_DisconnectNotify
func callbackQHelpIndexModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::disconnectNotify", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::disconnectNotify")
	}
}

func (ptr *QHelpIndexModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpIndexModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexModel_Event
func callbackQHelpIndexModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpIndexModel) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::event", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::event")
	}
}

func (ptr *QHelpIndexModel) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_EventFilter
func callbackQHelpIndexModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::eventFilter", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::eventFilter")
	}
}

func (ptr *QHelpIndexModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpIndexModel_MetaObject
func callbackQHelpIndexModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpIndexModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpIndexModel) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::metaObject", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexModel::metaObject")
	}
}

func (ptr *QHelpIndexModel) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpIndexModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpIndexWidget struct {
	widgets.QListView
}

type QHelpIndexWidget_ITF interface {
	widgets.QListView_ITF
	QHelpIndexWidget_PTR() *QHelpIndexWidget
}

func (p *QHelpIndexWidget) QHelpIndexWidget_PTR() *QHelpIndexWidget {
	return p
}

func (p *QHelpIndexWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QListView_PTR().Pointer()
	}
	return nil
}

func (p *QHelpIndexWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QListView_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpIndexWidget(ptr QHelpIndexWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexWidgetFromPointer(ptr unsafe.Pointer) *QHelpIndexWidget {
	var n = new(QHelpIndexWidget)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpIndexWidget) DestroyQHelpIndexWidget() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQHelpIndexWidget_ActivateCurrentItem
func callbackQHelpIndexWidget_ActivateCurrentItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::activateCurrentItem"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexWidget) ConnectActivateCurrentItem(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::activateCurrentItem", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectActivateCurrentItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::activateCurrentItem")
	}
}

func (ptr *QHelpIndexWidget) ActivateCurrentItem() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActivateCurrentItem(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_FilterIndices
func callbackQHelpIndexWidget_FilterIndices(ptr unsafe.Pointer, filter C.struct_QtHelp_PackedString, wildcard C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::filterIndices"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(filter), cGoUnpackString(wildcard))
	}

}

func (ptr *QHelpIndexWidget) ConnectFilterIndices(f func(filter string, wildcard string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::filterIndices", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFilterIndices(filter string, wildcard string) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::filterIndices")
	}
}

func (ptr *QHelpIndexWidget) FilterIndices(filter string, wildcard string) {
	if ptr.Pointer() != nil {
		var filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
		var wildcardC = C.CString(wildcard)
		defer C.free(unsafe.Pointer(wildcardC))
		C.QHelpIndexWidget_FilterIndices(ptr.Pointer(), filterC, wildcardC)
	}
}

//export callbackQHelpIndexWidget_LinkActivated
func callbackQHelpIndexWidget_LinkActivated(ptr unsafe.Pointer, link unsafe.Pointer, keyword C.struct_QtHelp_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::linkActivated"); signal != nil {
		signal.(func(*core.QUrl, string))(core.NewQUrlFromPointer(link), cGoUnpackString(keyword))
	}

}

func (ptr *QHelpIndexWidget) ConnectLinkActivated(f func(link *core.QUrl, keyword string)) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::linkActivated", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::linkActivated")
	}
}

func (ptr *QHelpIndexWidget) LinkActivated(link core.QUrl_ITF, keyword string) {
	if ptr.Pointer() != nil {
		var keywordC = C.CString(keyword)
		defer C.free(unsafe.Pointer(keywordC))
		C.QHelpIndexWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link), keywordC)
	}
}

//export callbackQHelpIndexWidget_CurrentChanged
func callbackQHelpIndexWidget_CurrentChanged(ptr unsafe.Pointer, current unsafe.Pointer, previous unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CurrentChangedDefault(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
	}
}

func (ptr *QHelpIndexWidget) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::currentChanged", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::currentChanged")
	}
}

func (ptr *QHelpIndexWidget) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QHelpIndexWidget) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

//export callbackQHelpIndexWidget_DragLeaveEvent
func callbackQHelpIndexWidget_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::dragLeaveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::dragLeaveEvent")
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackQHelpIndexWidget_DragMoveEvent
func callbackQHelpIndexWidget_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::dragMoveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::dragMoveEvent")
	}
}

func (ptr *QHelpIndexWidget) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackQHelpIndexWidget_DropEvent
func callbackQHelpIndexWidget_DropEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::dropEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::dropEvent")
	}
}

func (ptr *QHelpIndexWidget) DropEvent(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

//export callbackQHelpIndexWidget_HorizontalOffset
func callbackQHelpIndexWidget_HorizontalOffset(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::horizontalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHelpIndexWidget) ConnectHorizontalOffset(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::horizontalOffset", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHorizontalOffset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::horizontalOffset")
	}
}

func (ptr *QHelpIndexWidget) HorizontalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HorizontalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpIndexWidget) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpIndexWidget_IndexAt
func callbackQHelpIndexWidget_IndexAt(ptr unsafe.Pointer, p unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::indexAt"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QPoint) *core.QModelIndex)(core.NewQPointFromPointer(p)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexWidgetFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(p)))
}

func (ptr *QHelpIndexWidget) ConnectIndexAt(f func(p *core.QPoint) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::indexAt", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectIndexAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::indexAt")
	}
}

func (ptr *QHelpIndexWidget) IndexAt(p core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexWidget_IndexAt(ptr.Pointer(), core.PointerFromQPoint(p)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) IndexAtDefault(p core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexWidget_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(p)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_IsIndexHidden
func callbackQHelpIndexWidget_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHelpIndexWidget) ConnectIsIndexHidden(f func(index *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::isIndexHidden", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectIsIndexHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::isIndexHidden")
	}
}

func (ptr *QHelpIndexWidget) IsIndexHidden(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_MouseMoveEvent
func callbackQHelpIndexWidget_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::mouseMoveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::mouseMoveEvent")
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQHelpIndexWidget_MouseReleaseEvent
func callbackQHelpIndexWidget_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::mouseReleaseEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::mouseReleaseEvent")
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

//export callbackQHelpIndexWidget_MoveCursor
func callbackQHelpIndexWidget_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::moveCursor"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(widgets.QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHelpIndexWidgetFromPointer(ptr).MoveCursorDefault(widgets.QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHelpIndexWidget) ConnectMoveCursor(f func(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::moveCursor", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMoveCursor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::moveCursor")
	}
}

func (ptr *QHelpIndexWidget) MoveCursor(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexWidget_MoveCursor(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) MoveCursorDefault(cursorAction widgets.QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QHelpIndexWidget_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_PaintEvent
func callbackQHelpIndexWidget_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::paintEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectPaintEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::paintEvent")
	}
}

func (ptr *QHelpIndexWidget) PaintEvent(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QHelpIndexWidget) PaintEventDefault(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

//export callbackQHelpIndexWidget_ResizeEvent
func callbackQHelpIndexWidget_ResizeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::resizeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::resizeEvent")
	}
}

func (ptr *QHelpIndexWidget) ResizeEvent(e gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

//export callbackQHelpIndexWidget_RowsAboutToBeRemoved
func callbackQHelpIndexWidget_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RowsAboutToBeRemovedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::rowsAboutToBeRemoved", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRowsAboutToBeRemoved() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::rowsAboutToBeRemoved")
	}
}

func (ptr *QHelpIndexWidget) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpIndexWidget_RowsInserted
func callbackQHelpIndexWidget_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RowsInsertedDefault(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::rowsInserted", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRowsInserted() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::rowsInserted")
	}
}

func (ptr *QHelpIndexWidget) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

func (ptr *QHelpIndexWidget) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(int32(start)), C.int(int32(end)))
	}
}

//export callbackQHelpIndexWidget_ScrollTo
func callbackQHelpIndexWidget_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, widgets.QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollTo(f func(index *core.QModelIndex, hint widgets.QAbstractItemView__ScrollHint)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::scrollTo", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollTo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::scrollTo")
	}
}

func (ptr *QHelpIndexWidget) ScrollTo(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

func (ptr *QHelpIndexWidget) ScrollToDefault(index core.QModelIndex_ITF, hint widgets.QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQHelpIndexWidget_SelectionChanged
func callbackQHelpIndexWidget_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::selectionChanged"); signal != nil {
		signal.(func(*core.QItemSelection, *core.QItemSelection))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SelectionChangedDefault(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *QHelpIndexWidget) ConnectSelectionChanged(f func(selected *core.QItemSelection, deselected *core.QItemSelection)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::selectionChanged", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::selectionChanged")
	}
}

func (ptr *QHelpIndexWidget) SelectionChanged(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectionChanged(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

func (ptr *QHelpIndexWidget) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectionChangedDefault(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

//export callbackQHelpIndexWidget_SetSelection
func callbackQHelpIndexWidget_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, command C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setSelection", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setSelection")
	}
}

func (ptr *QHelpIndexWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

//export callbackQHelpIndexWidget_StartDrag
func callbackQHelpIndexWidget_StartDrag(ptr unsafe.Pointer, supportedActions C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::startDrag", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectStartDrag() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::startDrag")
	}
}

func (ptr *QHelpIndexWidget) StartDrag(supportedActions core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDrag(ptr.Pointer(), C.longlong(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDragDefault(ptr.Pointer(), C.longlong(supportedActions))
	}
}

//export callbackQHelpIndexWidget_UpdateGeometries
func callbackQHelpIndexWidget_UpdateGeometries(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::updateGeometries"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateGeometriesDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectUpdateGeometries(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::updateGeometries", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectUpdateGeometries() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::updateGeometries")
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometries() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometriesDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_VerticalOffset
func callbackQHelpIndexWidget_VerticalOffset(ptr unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::verticalOffset"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHelpIndexWidget) ConnectVerticalOffset(f func() int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::verticalOffset", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectVerticalOffset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::verticalOffset")
	}
}

func (ptr *QHelpIndexWidget) VerticalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_VerticalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpIndexWidget) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpIndexWidget_ViewOptions
func callbackQHelpIndexWidget_ViewOptions(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::viewOptions"); signal != nil {
		return widgets.PointerFromQStyleOptionViewItem(signal.(func() *widgets.QStyleOptionViewItem)())
	}

	return widgets.PointerFromQStyleOptionViewItem(NewQHelpIndexWidgetFromPointer(ptr).ViewOptionsDefault())
}

func (ptr *QHelpIndexWidget) ConnectViewOptions(f func() *widgets.QStyleOptionViewItem) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::viewOptions", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectViewOptions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::viewOptions")
	}
}

func (ptr *QHelpIndexWidget) ViewOptions() *widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQStyleOptionViewItemFromPointer(C.QHelpIndexWidget_ViewOptions(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) ViewOptionsDefault() *widgets.QStyleOptionViewItem {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQStyleOptionViewItemFromPointer(C.QHelpIndexWidget_ViewOptionsDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*widgets.QStyleOptionViewItem).DestroyQStyleOptionViewItem)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ViewportSizeHint
func callbackQHelpIndexWidget_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::viewportSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *QHelpIndexWidget) ConnectViewportSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::viewportSizeHint", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectViewportSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::viewportSizeHint")
	}
}

func (ptr *QHelpIndexWidget) ViewportSizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_ViewportSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) ViewportSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_ViewportSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_VisualRect
func callbackQHelpIndexWidget_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::visualRect"); signal != nil {
		return core.PointerFromQRect(signal.(func(*core.QModelIndex) *core.QRect)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHelpIndexWidgetFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHelpIndexWidget) ConnectVisualRect(f func(index *core.QModelIndex) *core.QRect) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::visualRect", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectVisualRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::visualRect")
	}
}

func (ptr *QHelpIndexWidget) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QHelpIndexWidget_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QHelpIndexWidget_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_VisualRegionForSelection
func callbackQHelpIndexWidget_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion(signal.(func(*core.QItemSelection) *gui.QRegion)(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHelpIndexWidgetFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHelpIndexWidget) ConnectVisualRegionForSelection(f func(selection *core.QItemSelection) *gui.QRegion) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::visualRegionForSelection", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectVisualRegionForSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::visualRegionForSelection")
	}
}

func (ptr *QHelpIndexWidget) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQRegionFromPointer(C.QHelpIndexWidget_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQRegionFromPointer(C.QHelpIndexWidget_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		runtime.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_WheelEvent
func callbackQHelpIndexWidget_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::wheelEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::wheelEvent")
	}
}

func (ptr *QHelpIndexWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpIndexWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

//export callbackQHelpIndexWidget_ViewportEvent
func callbackQHelpIndexWidget_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).ViewportEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) ConnectViewportEvent(f func(event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::viewportEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectViewportEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::viewportEvent")
	}
}

func (ptr *QHelpIndexWidget) ViewportEvent(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) ViewportEventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_ViewportEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_ClearSelection
func callbackQHelpIndexWidget_ClearSelection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::clearSelection"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ClearSelectionDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectClearSelection(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::clearSelection", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectClearSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::clearSelection")
	}
}

func (ptr *QHelpIndexWidget) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ClearSelectionDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ClearSelectionDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_CloseEditor
func callbackQHelpIndexWidget_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CloseEditorDefault(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	}
}

func (ptr *QHelpIndexWidget) ConnectCloseEditor(f func(editor *widgets.QWidget, hint widgets.QAbstractItemDelegate__EndEditHint)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::closeEditor", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCloseEditor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::closeEditor")
	}
}

func (ptr *QHelpIndexWidget) CloseEditor(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

func (ptr *QHelpIndexWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.longlong(hint))
	}
}

//export callbackQHelpIndexWidget_CommitData
func callbackQHelpIndexWidget_CommitData(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CommitDataDefault(widgets.NewQWidgetFromPointer(editor))
	}
}

func (ptr *QHelpIndexWidget) ConnectCommitData(f func(editor *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::commitData", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCommitData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::commitData")
	}
}

func (ptr *QHelpIndexWidget) CommitData(editor widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitData(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpIndexWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

//export callbackQHelpIndexWidget_DragEnterEvent
func callbackQHelpIndexWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::dragEnterEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::dragEnterEvent")
	}
}

func (ptr *QHelpIndexWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpIndexWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpIndexWidget_Edit
func callbackQHelpIndexWidget_Edit(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::edit"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectEdit(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::edit", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEdit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::edit")
	}
}

func (ptr *QHelpIndexWidget) Edit(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Edit(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) EditDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_Edit2
func callbackQHelpIndexWidget_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, widgets.QAbstractItemView__EditTrigger, *core.QEvent) bool)(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), widgets.QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) ConnectEdit2(f func(index *core.QModelIndex, trigger widgets.QAbstractItemView__EditTrigger, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::edit2", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEdit2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::edit2")
	}
}

func (ptr *QHelpIndexWidget) Edit2(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_Edit2(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) Edit2Default(index core.QModelIndex_ITF, trigger widgets.QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_EditorDestroyed
func callbackQHelpIndexWidget_EditorDestroyed(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EditorDestroyedDefault(core.NewQObjectFromPointer(editor))
	}
}

func (ptr *QHelpIndexWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::editorDestroyed", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEditorDestroyed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::editorDestroyed")
	}
}

func (ptr *QHelpIndexWidget) EditorDestroyed(editor core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpIndexWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

//export callbackQHelpIndexWidget_FocusInEvent
func callbackQHelpIndexWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::focusInEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::focusInEvent")
	}
}

func (ptr *QHelpIndexWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpIndexWidget_FocusNextPrevChild
func callbackQHelpIndexWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpIndexWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::focusNextPrevChild", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusNextPrevChild() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::focusNextPrevChild")
	}
}

func (ptr *QHelpIndexWidget) FocusNextPrevChild(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_FocusOutEvent
func callbackQHelpIndexWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::focusOutEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::focusOutEvent")
	}
}

func (ptr *QHelpIndexWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpIndexWidget_InputMethodEvent
func callbackQHelpIndexWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::inputMethodEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::inputMethodEvent")
	}
}

func (ptr *QHelpIndexWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpIndexWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpIndexWidget_InputMethodQuery
func callbackQHelpIndexWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpIndexWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpIndexWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::inputMethodQuery", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::inputMethodQuery")
	}
}

func (ptr *QHelpIndexWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpIndexWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_KeyPressEvent
func callbackQHelpIndexWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::keyPressEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::keyPressEvent")
	}
}

func (ptr *QHelpIndexWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpIndexWidget_KeyboardSearch
func callbackQHelpIndexWidget_KeyboardSearch(ptr unsafe.Pointer, search C.struct_QtHelp_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::keyboardSearch"); signal != nil {
		signal.(func(string))(cGoUnpackString(search))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyboardSearchDefault(cGoUnpackString(search))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyboardSearch(f func(search string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::keyboardSearch", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyboardSearch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::keyboardSearch")
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearch(search string) {
	if ptr.Pointer() != nil {
		var searchC = C.CString(search)
		defer C.free(unsafe.Pointer(searchC))
		C.QHelpIndexWidget_KeyboardSearch(ptr.Pointer(), searchC)
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearchDefault(search string) {
	if ptr.Pointer() != nil {
		var searchC = C.CString(search)
		defer C.free(unsafe.Pointer(searchC))
		C.QHelpIndexWidget_KeyboardSearchDefault(ptr.Pointer(), searchC)
	}
}

//export callbackQHelpIndexWidget_MouseDoubleClickEvent
func callbackQHelpIndexWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::mouseDoubleClickEvent")
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpIndexWidget_MousePressEvent
func callbackQHelpIndexWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::mousePressEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::mousePressEvent")
	}
}

func (ptr *QHelpIndexWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpIndexWidget_Reset
func callbackQHelpIndexWidget_Reset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::reset"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectReset(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::reset", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::reset")
	}
}

func (ptr *QHelpIndexWidget) Reset() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ResetDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResetDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ScrollToBottom
func callbackQHelpIndexWidget_ScrollToBottom(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::scrollToBottom"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToBottomDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollToBottom(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::scrollToBottom", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollToBottom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::scrollToBottom")
	}
}

func (ptr *QHelpIndexWidget) ScrollToBottom() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToBottom(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ScrollToBottomDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToBottomDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ScrollToTop
func callbackQHelpIndexWidget_ScrollToTop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::scrollToTop"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollToTopDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollToTop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::scrollToTop", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollToTop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::scrollToTop")
	}
}

func (ptr *QHelpIndexWidget) ScrollToTop() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToTop(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ScrollToTopDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollToTopDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SelectAll
func callbackQHelpIndexWidget_SelectAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::selectAll"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectSelectAll(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::selectAll", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSelectAll() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::selectAll")
	}
}

func (ptr *QHelpIndexWidget) SelectAll() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) SelectAllDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SelectionCommand
func callbackQHelpIndexWidget_SelectionCommand(ptr unsafe.Pointer, index unsafe.Pointer, event unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::selectionCommand"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex, *core.QEvent) core.QItemSelectionModel__SelectionFlag)(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
	}

	return C.longlong(NewQHelpIndexWidgetFromPointer(ptr).SelectionCommandDefault(core.NewQModelIndexFromPointer(index), core.NewQEventFromPointer(event)))
}

func (ptr *QHelpIndexWidget) ConnectSelectionCommand(f func(index *core.QModelIndex, event *core.QEvent) core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::selectionCommand", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSelectionCommand() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::selectionCommand")
	}
}

func (ptr *QHelpIndexWidget) SelectionCommand(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpIndexWidget_SelectionCommand(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

func (ptr *QHelpIndexWidget) SelectionCommandDefault(index core.QModelIndex_ITF, event core.QEvent_ITF) core.QItemSelectionModel__SelectionFlag {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModel__SelectionFlag(C.QHelpIndexWidget_SelectionCommandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQEvent(event)))
	}
	return 0
}

//export callbackQHelpIndexWidget_SetCurrentIndex
func callbackQHelpIndexWidget_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setCurrentIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetCurrentIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetCurrentIndex(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setCurrentIndex", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetCurrentIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setCurrentIndex")
	}
}

func (ptr *QHelpIndexWidget) SetCurrentIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetCurrentIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) SetCurrentIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetCurrentIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_SetModel
func callbackQHelpIndexWidget_SetModel(ptr unsafe.Pointer, model unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setModel", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setModel")
	}
}

func (ptr *QHelpIndexWidget) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpIndexWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

//export callbackQHelpIndexWidget_SetRootIndex
func callbackQHelpIndexWidget_SetRootIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetRootIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setRootIndex", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetRootIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setRootIndex")
	}
}

func (ptr *QHelpIndexWidget) SetRootIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) SetRootIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_SetSelectionModel
func callbackQHelpIndexWidget_SetSelectionModel(ptr unsafe.Pointer, selectionModel unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setSelectionModel", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetSelectionModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setSelectionModel")
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

//export callbackQHelpIndexWidget_SizeHintForColumn
func callbackQHelpIndexWidget_SizeHintForColumn(ptr unsafe.Pointer, column C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::sizeHintForColumn"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(column)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).SizeHintForColumnDefault(int(int32(column)))))
}

func (ptr *QHelpIndexWidget) ConnectSizeHintForColumn(f func(column int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::sizeHintForColumn", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSizeHintForColumn() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::sizeHintForColumn")
	}
}

func (ptr *QHelpIndexWidget) SizeHintForColumn(column int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForColumn(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

func (ptr *QHelpIndexWidget) SizeHintForColumnDefault(column int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForColumnDefault(ptr.Pointer(), C.int(int32(column)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_SizeHintForRow
func callbackQHelpIndexWidget_SizeHintForRow(ptr unsafe.Pointer, row C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::sizeHintForRow"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(row)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).SizeHintForRowDefault(int(int32(row)))))
}

func (ptr *QHelpIndexWidget) ConnectSizeHintForRow(f func(row int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::sizeHintForRow", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSizeHintForRow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::sizeHintForRow")
	}
}

func (ptr *QHelpIndexWidget) SizeHintForRow(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForRow(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

func (ptr *QHelpIndexWidget) SizeHintForRowDefault(row int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_SizeHintForRowDefault(ptr.Pointer(), C.int(int32(row)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_Update
func callbackQHelpIndexWidget_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::update"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QHelpIndexWidget) ConnectUpdate(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::update", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::update")
	}
}

func (ptr *QHelpIndexWidget) Update(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Update(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHelpIndexWidget) UpdateDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQHelpIndexWidget_ContextMenuEvent
func callbackQHelpIndexWidget_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::contextMenuEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::contextMenuEvent")
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackQHelpIndexWidget_MinimumSizeHint
func callbackQHelpIndexWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpIndexWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::minimumSizeHint", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMinimumSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::minimumSizeHint")
	}
}

func (ptr *QHelpIndexWidget) MinimumSizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ScrollContentsBy
func callbackQHelpIndexWidget_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(int32(dx)), int(int32(dy)))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::scrollContentsBy", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollContentsBy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::scrollContentsBy")
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsBy(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsBy(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsByDefault(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackQHelpIndexWidget_SetupViewport
func callbackQHelpIndexWidget_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetupViewport(f func(viewport *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setupViewport", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetupViewport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setupViewport")
	}
}

func (ptr *QHelpIndexWidget) SetupViewport(viewport widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewport(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpIndexWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

//export callbackQHelpIndexWidget_SizeHint
func callbackQHelpIndexWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpIndexWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpIndexWidget) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::sizeHint", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::sizeHint")
	}
}

func (ptr *QHelpIndexWidget) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpIndexWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpIndexWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpIndexWidget_ChangeEvent
func callbackQHelpIndexWidget_ChangeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpIndexWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::changeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectChangeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::changeEvent")
	}
}

func (ptr *QHelpIndexWidget) ChangeEvent(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpIndexWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

//export callbackQHelpIndexWidget_ActionEvent
func callbackQHelpIndexWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::actionEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectActionEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::actionEvent")
	}
}

func (ptr *QHelpIndexWidget) ActionEvent(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpIndexWidget_EnterEvent
func callbackQHelpIndexWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::enterEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::enterEvent")
	}
}

func (ptr *QHelpIndexWidget) EnterEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_HideEvent
func callbackQHelpIndexWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::hideEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::hideEvent")
	}
}

func (ptr *QHelpIndexWidget) HideEvent(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpIndexWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpIndexWidget_LeaveEvent
func callbackQHelpIndexWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::leaveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::leaveEvent")
	}
}

func (ptr *QHelpIndexWidget) LeaveEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_MoveEvent
func callbackQHelpIndexWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::moveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::moveEvent")
	}
}

func (ptr *QHelpIndexWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpIndexWidget_SetEnabled
func callbackQHelpIndexWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setEnabled", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setEnabled")
	}
}

func (ptr *QHelpIndexWidget) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpIndexWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpIndexWidget_SetStyleSheet
func callbackQHelpIndexWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setStyleSheet", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetStyleSheet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setStyleSheet")
	}
}

func (ptr *QHelpIndexWidget) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpIndexWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QHelpIndexWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpIndexWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQHelpIndexWidget_SetVisible
func callbackQHelpIndexWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setVisible", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setVisible")
	}
}

func (ptr *QHelpIndexWidget) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QHelpIndexWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpIndexWidget_SetWindowModified
func callbackQHelpIndexWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetWindowModified(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setWindowModified", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetWindowModified() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setWindowModified")
	}
}

func (ptr *QHelpIndexWidget) SetWindowModified(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpIndexWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpIndexWidget_SetWindowTitle
func callbackQHelpIndexWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetWindowTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setWindowTitle", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetWindowTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setWindowTitle")
	}
}

func (ptr *QHelpIndexWidget) SetWindowTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpIndexWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QHelpIndexWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpIndexWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQHelpIndexWidget_ShowEvent
func callbackQHelpIndexWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showEvent")
	}
}

func (ptr *QHelpIndexWidget) ShowEvent(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpIndexWidget_Close
func callbackQHelpIndexWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpIndexWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::close", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::close")
	}
}

func (ptr *QHelpIndexWidget) Close() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_CloseEvent
func callbackQHelpIndexWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::closeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCloseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::closeEvent")
	}
}

func (ptr *QHelpIndexWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpIndexWidget_HasHeightForWidth
func callbackQHelpIndexWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpIndexWidget) ConnectHasHeightForWidth(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::hasHeightForWidth", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHasHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::hasHeightForWidth")
	}
}

func (ptr *QHelpIndexWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_HeightForWidth
func callbackQHelpIndexWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpIndexWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpIndexWidget) ConnectHeightForWidth(f func(w int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::heightForWidth", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::heightForWidth")
	}
}

func (ptr *QHelpIndexWidget) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QHelpIndexWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpIndexWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpIndexWidget_Hide
func callbackQHelpIndexWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::hide", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::hide")
	}
}

func (ptr *QHelpIndexWidget) Hide() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_KeyReleaseEvent
func callbackQHelpIndexWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::keyReleaseEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::keyReleaseEvent")
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpIndexWidget_Lower
func callbackQHelpIndexWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::lower", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::lower")
	}
}

func (ptr *QHelpIndexWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_NativeEvent
func callbackQHelpIndexWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QHelpIndexWidget) ConnectNativeEvent(f func(eventType *core.QByteArray, message unsafe.Pointer, result int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::nativeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectNativeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::nativeEvent")
	}
}

func (ptr *QHelpIndexWidget) NativeEvent(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_NativeEvent(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_Raise
func callbackQHelpIndexWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectRaise(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::raise", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRaise() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::raise")
	}
}

func (ptr *QHelpIndexWidget) Raise() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_Repaint
func callbackQHelpIndexWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectRepaint(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::repaint", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectRepaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::repaint")
	}
}

func (ptr *QHelpIndexWidget) Repaint() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SetDisabled
func callbackQHelpIndexWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetDisabled(f func(disable bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setDisabled", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetDisabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setDisabled")
	}
}

func (ptr *QHelpIndexWidget) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QHelpIndexWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpIndexWidget_SetFocus2
func callbackQHelpIndexWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpIndexWidget) ConnectSetFocus2(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setFocus2", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetFocus2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setFocus2")
	}
}

func (ptr *QHelpIndexWidget) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_SetHidden
func callbackQHelpIndexWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpIndexWidget) ConnectSetHidden(f func(hidden bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setHidden", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::setHidden")
	}
}

func (ptr *QHelpIndexWidget) SetHidden(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QHelpIndexWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpIndexWidget_Show
func callbackQHelpIndexWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::show", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::show")
	}
}

func (ptr *QHelpIndexWidget) Show() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Show(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowFullScreen
func callbackQHelpIndexWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShowFullScreen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showFullScreen", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowFullScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showFullScreen")
	}
}

func (ptr *QHelpIndexWidget) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowMaximized
func callbackQHelpIndexWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showMaximized", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showMaximized")
	}
}

func (ptr *QHelpIndexWidget) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowMinimized
func callbackQHelpIndexWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShowMinimized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showMinimized", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowMinimized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showMinimized")
	}
}

func (ptr *QHelpIndexWidget) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ShowNormal
func callbackQHelpIndexWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectShowNormal(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showNormal", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowNormal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::showNormal")
	}
}

func (ptr *QHelpIndexWidget) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_TabletEvent
func callbackQHelpIndexWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::tabletEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectTabletEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::tabletEvent")
	}
}

func (ptr *QHelpIndexWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpIndexWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpIndexWidget_UpdateMicroFocus
func callbackQHelpIndexWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectUpdateMicroFocus(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::updateMicroFocus", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectUpdateMicroFocus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::updateMicroFocus")
	}
}

func (ptr *QHelpIndexWidget) UpdateMicroFocus() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpIndexWidget_ChildEvent
func callbackQHelpIndexWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::childEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::childEvent")
	}
}

func (ptr *QHelpIndexWidget) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpIndexWidget_ConnectNotify
func callbackQHelpIndexWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::connectNotify", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::connectNotify")
	}
}

func (ptr *QHelpIndexWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpIndexWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexWidget_CustomEvent
func callbackQHelpIndexWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::customEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::customEvent")
	}
}

func (ptr *QHelpIndexWidget) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpIndexWidget_DeleteLater
func callbackQHelpIndexWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpIndexWidget) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::deleteLater", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::deleteLater")
	}
}

func (ptr *QHelpIndexWidget) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpIndexWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpIndexWidget_DisconnectNotify
func callbackQHelpIndexWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpIndexWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::disconnectNotify", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::disconnectNotify")
	}
}

func (ptr *QHelpIndexWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpIndexWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpIndexWidget_EventFilter
func callbackQHelpIndexWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpIndexWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpIndexWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::eventFilter", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::eventFilter")
	}
}

func (ptr *QHelpIndexWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpIndexWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpIndexWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpIndexWidget_MetaObject
func callbackQHelpIndexWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpIndexWidget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpIndexWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpIndexWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::metaObject", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpIndexWidget::metaObject")
	}
}

func (ptr *QHelpIndexWidget) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpIndexWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpIndexWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpSearchEngine struct {
	core.QObject
}

type QHelpSearchEngine_ITF interface {
	core.QObject_ITF
	QHelpSearchEngine_PTR() *QHelpSearchEngine
}

func (p *QHelpSearchEngine) QHelpSearchEngine_PTR() *QHelpSearchEngine {
	return p
}

func (p *QHelpSearchEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QHelpSearchEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpSearchEngine(ptr QHelpSearchEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchEngineFromPointer(ptr unsafe.Pointer) *QHelpSearchEngine {
	var n = new(QHelpSearchEngine)
	n.SetPointer(ptr)
	return n
}
func NewQHelpSearchEngine(helpEngine QHelpEngineCore_ITF, parent core.QObject_ITF) *QHelpSearchEngine {
	var tmpValue = NewQHelpSearchEngineFromPointer(C.QHelpSearchEngine_NewQHelpSearchEngine(PointerFromQHelpEngineCore(helpEngine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQHelpSearchEngine_CancelIndexing
func callbackQHelpSearchEngine_CancelIndexing(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::cancelIndexing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectCancelIndexing(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::cancelIndexing", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectCancelIndexing() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::cancelIndexing")
	}
}

func (ptr *QHelpSearchEngine) CancelIndexing() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelIndexing(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_CancelSearching
func callbackQHelpSearchEngine_CancelSearching(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::cancelSearching"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectCancelSearching(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::cancelSearching", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectCancelSearching() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::cancelSearching")
	}
}

func (ptr *QHelpSearchEngine) CancelSearching() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelSearching(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) HitCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchEngine_HitCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQHelpSearchEngine_IndexingFinished
func callbackQHelpSearchEngine_IndexingFinished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::indexingFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::indexingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingFinished() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::indexingFinished")
	}
}

func (ptr *QHelpSearchEngine) IndexingFinished() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_IndexingFinished(ptr.Pointer())
	}
}

//export callbackQHelpSearchEngine_IndexingStarted
func callbackQHelpSearchEngine_IndexingStarted(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::indexingStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::indexingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::indexingStarted")
	}
}

func (ptr *QHelpSearchEngine) IndexingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_IndexingStarted(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) QueryWidget() *QHelpSearchQueryWidget {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchEngine_QueryWidget(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchEngine_ReindexDocumentation
func callbackQHelpSearchEngine_ReindexDocumentation(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::reindexDocumentation"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectReindexDocumentation(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::reindexDocumentation", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectReindexDocumentation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::reindexDocumentation")
	}
}

func (ptr *QHelpSearchEngine) ReindexDocumentation() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ReindexDocumentation(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) ResultWidget() *QHelpSearchResultWidget {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpSearchResultWidgetFromPointer(C.QHelpSearchEngine_ResultWidget(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchEngine_SearchingFinished
func callbackQHelpSearchEngine_SearchingFinished(ptr unsafe.Pointer, hits C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::searchingFinished"); signal != nil {
		signal.(func(int))(int(int32(hits)))
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(hits int)) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::searchingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingFinished() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::searchingFinished")
	}
}

func (ptr *QHelpSearchEngine) SearchingFinished(hits int) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_SearchingFinished(ptr.Pointer(), C.int(int32(hits)))
	}
}

//export callbackQHelpSearchEngine_SearchingStarted
func callbackQHelpSearchEngine_SearchingStarted(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::searchingStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingStarted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::searchingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingStarted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::searchingStarted")
	}
}

func (ptr *QHelpSearchEngine) SearchingStarted() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_SearchingStarted(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngine() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DestroyQHelpSearchEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchEngine) query_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpSearchQueryFromPointer(C.QHelpSearchEngine_query_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchEngine_TimerEvent
func callbackQHelpSearchEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::timerEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::timerEvent")
	}
}

func (ptr *QHelpSearchEngine) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpSearchEngine_ChildEvent
func callbackQHelpSearchEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::childEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::childEvent")
	}
}

func (ptr *QHelpSearchEngine) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchEngine_ConnectNotify
func callbackQHelpSearchEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::connectNotify", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::connectNotify")
	}
}

func (ptr *QHelpSearchEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchEngine_CustomEvent
func callbackQHelpSearchEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::customEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::customEvent")
	}
}

func (ptr *QHelpSearchEngine) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchEngine_DeleteLater
func callbackQHelpSearchEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchEngine) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::deleteLater", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::deleteLater")
	}
}

func (ptr *QHelpSearchEngine) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchEngine_DisconnectNotify
func callbackQHelpSearchEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::disconnectNotify", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::disconnectNotify")
	}
}

func (ptr *QHelpSearchEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchEngine_Event
func callbackQHelpSearchEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHelpSearchEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::event", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::event")
	}
}

func (ptr *QHelpSearchEngine) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHelpSearchEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHelpSearchEngine_EventFilter
func callbackQHelpSearchEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::eventFilter", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::eventFilter")
	}
}

func (ptr *QHelpSearchEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpSearchEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpSearchEngine_MetaObject
func callbackQHelpSearchEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::metaObject", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchEngine::metaObject")
	}
}

func (ptr *QHelpSearchEngine) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QHelpSearchQuery::FieldName
type QHelpSearchQuery__FieldName int64

const (
	QHelpSearchQuery__DEFAULT = QHelpSearchQuery__FieldName(0)
	QHelpSearchQuery__FUZZY   = QHelpSearchQuery__FieldName(1)
	QHelpSearchQuery__WITHOUT = QHelpSearchQuery__FieldName(2)
	QHelpSearchQuery__PHRASE  = QHelpSearchQuery__FieldName(3)
	QHelpSearchQuery__ALL     = QHelpSearchQuery__FieldName(4)
	QHelpSearchQuery__ATLEAST = QHelpSearchQuery__FieldName(5)
)

type QHelpSearchQuery struct {
	ptr unsafe.Pointer
}

type QHelpSearchQuery_ITF interface {
	QHelpSearchQuery_PTR() *QHelpSearchQuery
}

func (p *QHelpSearchQuery) QHelpSearchQuery_PTR() *QHelpSearchQuery {
	return p
}

func (p *QHelpSearchQuery) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QHelpSearchQuery) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQHelpSearchQuery(ptr QHelpSearchQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQuery_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryFromPointer(ptr unsafe.Pointer) *QHelpSearchQuery {
	var n = new(QHelpSearchQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpSearchQuery) DestroyQHelpSearchQuery() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQHelpSearchQuery() *QHelpSearchQuery {
	var tmpValue = NewQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery())
	runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
	return tmpValue
}

func NewQHelpSearchQuery2(field QHelpSearchQuery__FieldName, wordList []string) *QHelpSearchQuery {
	var wordListC = C.CString(strings.Join(wordList, "|"))
	defer C.free(unsafe.Pointer(wordListC))
	var tmpValue = NewQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery2(C.longlong(field), wordListC))
	runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
	return tmpValue
}

func (ptr *QHelpSearchQuery) FieldName() QHelpSearchQuery__FieldName {
	if ptr.Pointer() != nil {
		return QHelpSearchQuery__FieldName(C.QHelpSearchQuery_FieldName(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpSearchQuery) SetFieldName(vfi QHelpSearchQuery__FieldName) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQuery_SetFieldName(ptr.Pointer(), C.longlong(vfi))
	}
}

func (ptr *QHelpSearchQuery) WordList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QHelpSearchQuery_WordList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpSearchQuery) SetWordList(vqs []string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(strings.Join(vqs, "|"))
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchQuery_SetWordList(ptr.Pointer(), vqsC)
	}
}

type QHelpSearchQueryWidget struct {
	widgets.QWidget
}

type QHelpSearchQueryWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget
}

func (p *QHelpSearchQueryWidget) QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget {
	return p
}

func (p *QHelpSearchQueryWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QHelpSearchQueryWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpSearchQueryWidget(ptr QHelpSearchQueryWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchQueryWidget {
	var n = new(QHelpSearchQueryWidget)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpSearchQueryWidget) IsCompactMode() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_IsCompactMode(ptr.Pointer()) != 0
	}
	return false
}

func NewQHelpSearchQueryWidget(parent widgets.QWidget_ITF) *QHelpSearchQueryWidget {
	var tmpValue = NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CollapseExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ExpandExtendedSearch(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_Search
func callbackQHelpSearchQueryWidget_Search(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::search"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectSearch(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::search", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectSearch(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::search")
	}
}

func (ptr *QHelpSearchQueryWidget) Search() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Search(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchQueryWidget) query_atList(i int) *QHelpSearchQuery {
	if ptr.Pointer() != nil {
		var tmpValue = NewQHelpSearchQueryFromPointer(C.QHelpSearchQueryWidget_query_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QHelpSearchQuery).DestroyQHelpSearchQuery)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_ActionEvent
func callbackQHelpSearchQueryWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::actionEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectActionEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::actionEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ActionEvent(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DragEnterEvent
func callbackQHelpSearchQueryWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::dragEnterEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::dragEnterEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DragLeaveEvent
func callbackQHelpSearchQueryWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::dragLeaveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::dragLeaveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DragMoveEvent
func callbackQHelpSearchQueryWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::dragMoveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::dragMoveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DropEvent
func callbackQHelpSearchQueryWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::dropEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::dropEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) DropEvent(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_EnterEvent
func callbackQHelpSearchQueryWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::enterEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::enterEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) EnterEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusInEvent
func callbackQHelpSearchQueryWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::focusInEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::focusInEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusOutEvent
func callbackQHelpSearchQueryWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::focusOutEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::focusOutEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_HideEvent
func callbackQHelpSearchQueryWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::hideEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::hideEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) HideEvent(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_LeaveEvent
func callbackQHelpSearchQueryWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::leaveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::leaveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) LeaveEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MinimumSizeHint
func callbackQHelpSearchQueryWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchQueryWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpSearchQueryWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::minimumSizeHint", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMinimumSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::minimumSizeHint")
	}
}

func (ptr *QHelpSearchQueryWidget) MinimumSizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_MoveEvent
func callbackQHelpSearchQueryWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::moveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::moveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_PaintEvent
func callbackQHelpSearchQueryWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::paintEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectPaintEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::paintEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_SetEnabled
func callbackQHelpSearchQueryWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setEnabled", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setEnabled")
	}
}

func (ptr *QHelpSearchQueryWidget) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchQueryWidget_SetStyleSheet
func callbackQHelpSearchQueryWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setStyleSheet", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetStyleSheet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setStyleSheet")
	}
}

func (ptr *QHelpSearchQueryWidget) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpSearchQueryWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QHelpSearchQueryWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpSearchQueryWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQHelpSearchQueryWidget_SetVisible
func callbackQHelpSearchQueryWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setVisible", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setVisible")
	}
}

func (ptr *QHelpSearchQueryWidget) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpSearchQueryWidget_SetWindowModified
func callbackQHelpSearchQueryWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetWindowModified(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setWindowModified", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetWindowModified() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setWindowModified")
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowModified(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchQueryWidget_SetWindowTitle
func callbackQHelpSearchQueryWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetWindowTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setWindowTitle", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetWindowTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setWindowTitle")
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchQueryWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QHelpSearchQueryWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchQueryWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQHelpSearchQueryWidget_ShowEvent
func callbackQHelpSearchQueryWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowEvent(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_SizeHint
func callbackQHelpSearchQueryWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchQueryWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpSearchQueryWidget) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::sizeHint", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::sizeHint")
	}
}

func (ptr *QHelpSearchQueryWidget) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchQueryWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_ChangeEvent
func callbackQHelpSearchQueryWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::changeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectChangeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::changeEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ChangeEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Close
func callbackQHelpSearchQueryWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpSearchQueryWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::close", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::close")
	}
}

func (ptr *QHelpSearchQueryWidget) Close() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_CloseEvent
func callbackQHelpSearchQueryWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::closeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectCloseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::closeEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ContextMenuEvent
func callbackQHelpSearchQueryWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::contextMenuEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::contextMenuEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_FocusNextPrevChild
func callbackQHelpSearchQueryWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpSearchQueryWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::focusNextPrevChild", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectFocusNextPrevChild() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::focusNextPrevChild")
	}
}

func (ptr *QHelpSearchQueryWidget) FocusNextPrevChild(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_HasHeightForWidth
func callbackQHelpSearchQueryWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpSearchQueryWidget) ConnectHasHeightForWidth(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::hasHeightForWidth", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHasHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::hasHeightForWidth")
	}
}

func (ptr *QHelpSearchQueryWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_HeightForWidth
func callbackQHelpSearchQueryWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpSearchQueryWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpSearchQueryWidget) ConnectHeightForWidth(f func(w int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::heightForWidth", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::heightForWidth")
	}
}

func (ptr *QHelpSearchQueryWidget) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchQueryWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QHelpSearchQueryWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchQueryWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpSearchQueryWidget_Hide
func callbackQHelpSearchQueryWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::hide", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::hide")
	}
}

func (ptr *QHelpSearchQueryWidget) Hide() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_InputMethodEvent
func callbackQHelpSearchQueryWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::inputMethodEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::inputMethodEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_InputMethodQuery
func callbackQHelpSearchQueryWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpSearchQueryWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::inputMethodQuery", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::inputMethodQuery")
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpSearchQueryWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpSearchQueryWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchQueryWidget_KeyPressEvent
func callbackQHelpSearchQueryWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::keyPressEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::keyPressEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_KeyReleaseEvent
func callbackQHelpSearchQueryWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::keyReleaseEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::keyReleaseEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Lower
func callbackQHelpSearchQueryWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::lower", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::lower")
	}
}

func (ptr *QHelpSearchQueryWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_MouseDoubleClickEvent
func callbackQHelpSearchQueryWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::mouseDoubleClickEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MouseMoveEvent
func callbackQHelpSearchQueryWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::mouseMoveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::mouseMoveEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MousePressEvent
func callbackQHelpSearchQueryWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::mousePressEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::mousePressEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_MouseReleaseEvent
func callbackQHelpSearchQueryWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::mouseReleaseEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::mouseReleaseEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_NativeEvent
func callbackQHelpSearchQueryWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QHelpSearchQueryWidget) ConnectNativeEvent(f func(eventType *core.QByteArray, message unsafe.Pointer, result int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::nativeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectNativeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::nativeEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) NativeEvent(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_NativeEvent(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_Raise
func callbackQHelpSearchQueryWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectRaise(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::raise", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectRaise() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::raise")
	}
}

func (ptr *QHelpSearchQueryWidget) Raise() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_Repaint
func callbackQHelpSearchQueryWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectRepaint(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::repaint", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectRepaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::repaint")
	}
}

func (ptr *QHelpSearchQueryWidget) Repaint() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ResizeEvent
func callbackQHelpSearchQueryWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::resizeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::resizeEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_SetDisabled
func callbackQHelpSearchQueryWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetDisabled(f func(disable bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setDisabled", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetDisabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setDisabled")
	}
}

func (ptr *QHelpSearchQueryWidget) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpSearchQueryWidget_SetFocus2
func callbackQHelpSearchQueryWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetFocus2(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setFocus2", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetFocus2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setFocus2")
	}
}

func (ptr *QHelpSearchQueryWidget) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_SetHidden
func callbackQHelpSearchQueryWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetHidden(f func(hidden bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setHidden", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::setHidden")
	}
}

func (ptr *QHelpSearchQueryWidget) SetHidden(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QHelpSearchQueryWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpSearchQueryWidget_Show
func callbackQHelpSearchQueryWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::show", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::show")
	}
}

func (ptr *QHelpSearchQueryWidget) Show() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Show(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowFullScreen
func callbackQHelpSearchQueryWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowFullScreen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showFullScreen", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowFullScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showFullScreen")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowMaximized
func callbackQHelpSearchQueryWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showMaximized", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showMaximized")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowMinimized
func callbackQHelpSearchQueryWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowMinimized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showMinimized", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowMinimized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showMinimized")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_ShowNormal
func callbackQHelpSearchQueryWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowNormal(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showNormal", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowNormal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::showNormal")
	}
}

func (ptr *QHelpSearchQueryWidget) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_TabletEvent
func callbackQHelpSearchQueryWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::tabletEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectTabletEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::tabletEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_Update
func callbackQHelpSearchQueryWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::update"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::update", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::update")
	}
}

func (ptr *QHelpSearchQueryWidget) Update() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Update(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_UpdateMicroFocus
func callbackQHelpSearchQueryWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectUpdateMicroFocus(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::updateMicroFocus", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectUpdateMicroFocus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::updateMicroFocus")
	}
}

func (ptr *QHelpSearchQueryWidget) UpdateMicroFocus() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchQueryWidget_WheelEvent
func callbackQHelpSearchQueryWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::wheelEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::wheelEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_TimerEvent
func callbackQHelpSearchQueryWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::timerEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::timerEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ChildEvent
func callbackQHelpSearchQueryWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::childEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::childEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_ConnectNotify
func callbackQHelpSearchQueryWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::connectNotify", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::connectNotify")
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchQueryWidget_CustomEvent
func callbackQHelpSearchQueryWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::customEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::customEvent")
	}
}

func (ptr *QHelpSearchQueryWidget) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchQueryWidget_DeleteLater
func callbackQHelpSearchQueryWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::deleteLater", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::deleteLater")
	}
}

func (ptr *QHelpSearchQueryWidget) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchQueryWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchQueryWidget_DisconnectNotify
func callbackQHelpSearchQueryWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::disconnectNotify", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::disconnectNotify")
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchQueryWidget_EventFilter
func callbackQHelpSearchQueryWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchQueryWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchQueryWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::eventFilter", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::eventFilter")
	}
}

func (ptr *QHelpSearchQueryWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpSearchQueryWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpSearchQueryWidget_MetaObject
func callbackQHelpSearchQueryWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchQueryWidget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchQueryWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchQueryWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::metaObject", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchQueryWidget::metaObject")
	}
}

func (ptr *QHelpSearchQueryWidget) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchQueryWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchQueryWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchQueryWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHelpSearchResultWidget struct {
	widgets.QWidget
}

type QHelpSearchResultWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget
}

func (p *QHelpSearchResultWidget) QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget {
	return p
}

func (p *QHelpSearchResultWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QHelpSearchResultWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQHelpSearchResultWidget(ptr QHelpSearchResultWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResultWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchResultWidget {
	var n = new(QHelpSearchResultWidget)
	n.SetPointer(ptr)
	return n
}
func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPoint_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QHelpSearchResultWidget_LinkAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_RequestShowLink
func callbackQHelpSearchResultWidget_RequestShowLink(ptr unsafe.Pointer, link unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::requestShowLink"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link *core.QUrl)) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectRequestShowLink(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::requestShowLink", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRequestShowLink() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectRequestShowLink(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::requestShowLink")
	}
}

func (ptr *QHelpSearchResultWidget) RequestShowLink(link core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RequestShowLink(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchResultWidget_ActionEvent
func callbackQHelpSearchResultWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::actionEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectActionEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::actionEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ActionEvent(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DragEnterEvent
func callbackQHelpSearchResultWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::dragEnterEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::dragEnterEvent")
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DragLeaveEvent
func callbackQHelpSearchResultWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::dragLeaveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::dragLeaveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DragMoveEvent
func callbackQHelpSearchResultWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::dragMoveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::dragMoveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DropEvent
func callbackQHelpSearchResultWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::dropEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDropEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::dropEvent")
	}
}

func (ptr *QHelpSearchResultWidget) DropEvent(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_EnterEvent
func callbackQHelpSearchResultWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::enterEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectEnterEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::enterEvent")
	}
}

func (ptr *QHelpSearchResultWidget) EnterEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusInEvent
func callbackQHelpSearchResultWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::focusInEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusInEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::focusInEvent")
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusOutEvent
func callbackQHelpSearchResultWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::focusOutEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusOutEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::focusOutEvent")
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_HideEvent
func callbackQHelpSearchResultWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::hideEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::hideEvent")
	}
}

func (ptr *QHelpSearchResultWidget) HideEvent(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_LeaveEvent
func callbackQHelpSearchResultWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::leaveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectLeaveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::leaveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MinimumSizeHint
func callbackQHelpSearchResultWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchResultWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QHelpSearchResultWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::minimumSizeHint", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMinimumSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::minimumSizeHint")
	}
}

func (ptr *QHelpSearchResultWidget) MinimumSizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchResultWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchResultWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_MoveEvent
func callbackQHelpSearchResultWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::moveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::moveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_PaintEvent
func callbackQHelpSearchResultWidget_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::paintEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectPaintEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::paintEvent")
	}
}

func (ptr *QHelpSearchResultWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_SetEnabled
func callbackQHelpSearchResultWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setEnabled", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setEnabled")
	}
}

func (ptr *QHelpSearchResultWidget) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpSearchResultWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchResultWidget_SetStyleSheet
func callbackQHelpSearchResultWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setStyleSheet", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetStyleSheet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setStyleSheet")
	}
}

func (ptr *QHelpSearchResultWidget) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpSearchResultWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QHelpSearchResultWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QHelpSearchResultWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQHelpSearchResultWidget_SetVisible
func callbackQHelpSearchResultWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setVisible", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setVisible")
	}
}

func (ptr *QHelpSearchResultWidget) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QHelpSearchResultWidget) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQHelpSearchResultWidget_SetWindowModified
func callbackQHelpSearchResultWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetWindowModified(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setWindowModified", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetWindowModified() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setWindowModified")
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowModified(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQHelpSearchResultWidget_SetWindowTitle
func callbackQHelpSearchResultWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtHelp_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetWindowTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setWindowTitle", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetWindowTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setWindowTitle")
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchResultWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QHelpSearchResultWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QHelpSearchResultWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQHelpSearchResultWidget_ShowEvent
func callbackQHelpSearchResultWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ShowEvent(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_SizeHint
func callbackQHelpSearchResultWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQHelpSearchResultWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QHelpSearchResultWidget) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::sizeHint", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::sizeHint")
	}
}

func (ptr *QHelpSearchResultWidget) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchResultWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QHelpSearchResultWidget_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_ChangeEvent
func callbackQHelpSearchResultWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::changeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectChangeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::changeEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ChangeEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Close
func callbackQHelpSearchResultWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QHelpSearchResultWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::close", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::close")
	}
}

func (ptr *QHelpSearchResultWidget) Close() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_CloseEvent
func callbackQHelpSearchResultWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::closeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectCloseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::closeEvent")
	}
}

func (ptr *QHelpSearchResultWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ContextMenuEvent
func callbackQHelpSearchResultWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::contextMenuEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::contextMenuEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_FocusNextPrevChild
func callbackQHelpSearchResultWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QHelpSearchResultWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::focusNextPrevChild", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusNextPrevChild() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::focusNextPrevChild")
	}
}

func (ptr *QHelpSearchResultWidget) FocusNextPrevChild(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_HasHeightForWidth
func callbackQHelpSearchResultWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QHelpSearchResultWidget) ConnectHasHeightForWidth(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::hasHeightForWidth", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHasHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::hasHeightForWidth")
	}
}

func (ptr *QHelpSearchResultWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_HeightForWidth
func callbackQHelpSearchResultWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQHelpSearchResultWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QHelpSearchResultWidget) ConnectHeightForWidth(f func(w int) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::heightForWidth", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::heightForWidth")
	}
}

func (ptr *QHelpSearchResultWidget) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchResultWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QHelpSearchResultWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHelpSearchResultWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQHelpSearchResultWidget_Hide
func callbackQHelpSearchResultWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::hide", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::hide")
	}
}

func (ptr *QHelpSearchResultWidget) Hide() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_InputMethodEvent
func callbackQHelpSearchResultWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::inputMethodEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectInputMethodEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::inputMethodEvent")
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_InputMethodQuery
func callbackQHelpSearchResultWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QHelpSearchResultWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::inputMethodQuery", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectInputMethodQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::inputMethodQuery")
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpSearchResultWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QHelpSearchResultWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQHelpSearchResultWidget_KeyPressEvent
func callbackQHelpSearchResultWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::keyPressEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::keyPressEvent")
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_KeyReleaseEvent
func callbackQHelpSearchResultWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::keyReleaseEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::keyReleaseEvent")
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Lower
func callbackQHelpSearchResultWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::lower", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::lower")
	}
}

func (ptr *QHelpSearchResultWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_MouseDoubleClickEvent
func callbackQHelpSearchResultWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::mouseDoubleClickEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MouseMoveEvent
func callbackQHelpSearchResultWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::mouseMoveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseMoveEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::mouseMoveEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MousePressEvent
func callbackQHelpSearchResultWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::mousePressEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::mousePressEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_MouseReleaseEvent
func callbackQHelpSearchResultWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::mouseReleaseEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::mouseReleaseEvent")
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_NativeEvent
func callbackQHelpSearchResultWidget_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QHelpSearchResultWidget) ConnectNativeEvent(f func(eventType *core.QByteArray, message unsafe.Pointer, result int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::nativeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectNativeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::nativeEvent")
	}
}

func (ptr *QHelpSearchResultWidget) NativeEvent(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_NativeEvent(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_Raise
func callbackQHelpSearchResultWidget_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectRaise(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::raise", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRaise() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::raise")
	}
}

func (ptr *QHelpSearchResultWidget) Raise() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_Repaint
func callbackQHelpSearchResultWidget_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectRepaint(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::repaint", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRepaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::repaint")
	}
}

func (ptr *QHelpSearchResultWidget) Repaint() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ResizeEvent
func callbackQHelpSearchResultWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::resizeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::resizeEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_SetDisabled
func callbackQHelpSearchResultWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetDisabled(f func(disable bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setDisabled", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetDisabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setDisabled")
	}
}

func (ptr *QHelpSearchResultWidget) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QHelpSearchResultWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQHelpSearchResultWidget_SetFocus2
func callbackQHelpSearchResultWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetFocus2(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setFocus2", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetFocus2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setFocus2")
	}
}

func (ptr *QHelpSearchResultWidget) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_SetHidden
func callbackQHelpSearchResultWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetHidden(f func(hidden bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setHidden", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::setHidden")
	}
}

func (ptr *QHelpSearchResultWidget) SetHidden(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QHelpSearchResultWidget) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQHelpSearchResultWidget_Show
func callbackQHelpSearchResultWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::show"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::show", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::show")
	}
}

func (ptr *QHelpSearchResultWidget) Show() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Show(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowFullScreen
func callbackQHelpSearchResultWidget_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowFullScreen(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showFullScreen", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowFullScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showFullScreen")
	}
}

func (ptr *QHelpSearchResultWidget) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowMaximized
func callbackQHelpSearchResultWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showMaximized", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showMaximized")
	}
}

func (ptr *QHelpSearchResultWidget) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowMinimized
func callbackQHelpSearchResultWidget_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowMinimized(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showMinimized", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowMinimized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showMinimized")
	}
}

func (ptr *QHelpSearchResultWidget) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_ShowNormal
func callbackQHelpSearchResultWidget_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowNormal(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showNormal", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowNormal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::showNormal")
	}
}

func (ptr *QHelpSearchResultWidget) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_TabletEvent
func callbackQHelpSearchResultWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::tabletEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectTabletEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::tabletEvent")
	}
}

func (ptr *QHelpSearchResultWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_Update
func callbackQHelpSearchResultWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::update"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::update", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::update")
	}
}

func (ptr *QHelpSearchResultWidget) Update() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_Update(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_UpdateMicroFocus
func callbackQHelpSearchResultWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectUpdateMicroFocus(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::updateMicroFocus", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectUpdateMicroFocus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::updateMicroFocus")
	}
}

func (ptr *QHelpSearchResultWidget) UpdateMicroFocus() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QHelpSearchResultWidget) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQHelpSearchResultWidget_WheelEvent
func callbackQHelpSearchResultWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::wheelEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectWheelEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::wheelEvent")
	}
}

func (ptr *QHelpSearchResultWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_TimerEvent
func callbackQHelpSearchResultWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::timerEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::timerEvent")
	}
}

func (ptr *QHelpSearchResultWidget) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ChildEvent
func callbackQHelpSearchResultWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::childEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::childEvent")
	}
}

func (ptr *QHelpSearchResultWidget) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_ConnectNotify
func callbackQHelpSearchResultWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::connectNotify", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::connectNotify")
	}
}

func (ptr *QHelpSearchResultWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchResultWidget_CustomEvent
func callbackQHelpSearchResultWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::customEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::customEvent")
	}
}

func (ptr *QHelpSearchResultWidget) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHelpSearchResultWidget_DeleteLater
func callbackQHelpSearchResultWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::deleteLater", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::deleteLater")
	}
}

func (ptr *QHelpSearchResultWidget) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchResultWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHelpSearchResultWidget_DisconnectNotify
func callbackQHelpSearchResultWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::disconnectNotify", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::disconnectNotify")
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHelpSearchResultWidget_EventFilter
func callbackQHelpSearchResultWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHelpSearchResultWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHelpSearchResultWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::eventFilter", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::eventFilter")
	}
}

func (ptr *QHelpSearchResultWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHelpSearchResultWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpSearchResultWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHelpSearchResultWidget_MetaObject
func callbackQHelpSearchResultWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHelpSearchResultWidget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHelpSearchResultWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHelpSearchResultWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::metaObject", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHelpSearchResultWidget::metaObject")
	}
}

func (ptr *QHelpSearchResultWidget) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchResultWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHelpSearchResultWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
