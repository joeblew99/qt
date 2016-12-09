// +build !minimal

package macextras

//#include <stdint.h>
//#include <stdlib.h>
//#include "macextras.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtMacExtras_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QMacPasteboardMime struct {
	ptr unsafe.Pointer
}

type QMacPasteboardMime_ITF interface {
	QMacPasteboardMime_PTR() *QMacPasteboardMime
}

func (p *QMacPasteboardMime) QMacPasteboardMime_PTR() *QMacPasteboardMime {
	return p
}

func (p *QMacPasteboardMime) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QMacPasteboardMime) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQMacPasteboardMime(ptr QMacPasteboardMime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacPasteboardMime_PTR().Pointer()
	}
	return nil
}

func NewQMacPasteboardMimeFromPointer(ptr unsafe.Pointer) *QMacPasteboardMime {
	var n = new(QMacPasteboardMime)
	n.SetPointer(ptr)
	return n
}

//export callbackQMacPasteboardMime_CanConvert
func callbackQMacPasteboardMime_CanConvert(ptr unsafe.Pointer, mime C.struct_QtMacExtras_PackedString, flav C.struct_QtMacExtras_PackedString) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::canConvert"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string) bool)(cGoUnpackString(mime), cGoUnpackString(flav)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QMacPasteboardMime) ConnectCanConvert(f func(mime string, flav string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::canConvert", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectCanConvert(mime string, flav string) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::canConvert")
	}
}

func (ptr *QMacPasteboardMime) CanConvert(mime string, flav string) bool {
	if ptr.Pointer() != nil {
		var mimeC = C.CString(mime)
		defer C.free(unsafe.Pointer(mimeC))
		var flavC = C.CString(flav)
		defer C.free(unsafe.Pointer(flavC))
		return C.QMacPasteboardMime_CanConvert(ptr.Pointer(), mimeC, flavC) != 0
	}
	return false
}

//export callbackQMacPasteboardMime_ConvertorName
func callbackQMacPasteboardMime_ConvertorName(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::convertorName"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QMacPasteboardMime) ConnectConvertorName(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::convertorName", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectConvertorName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::convertorName")
	}
}

func (ptr *QMacPasteboardMime) ConvertorName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QMacPasteboardMime_ConvertorName(ptr.Pointer()))
	}
	return ""
}

//export callbackQMacPasteboardMime_Count
func callbackQMacPasteboardMime_Count(ptr unsafe.Pointer, mimeData unsafe.Pointer) C.int {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::count"); signal != nil {
		return C.int(int32(signal.(func(*core.QMimeData) int)(core.NewQMimeDataFromPointer(mimeData))))
	}

	return C.int(int32(NewQMacPasteboardMimeFromPointer(ptr).CountDefault(core.NewQMimeDataFromPointer(mimeData))))
}

func (ptr *QMacPasteboardMime) ConnectCount(f func(mimeData *core.QMimeData) int) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::count", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::count")
	}
}

func (ptr *QMacPasteboardMime) Count(mimeData core.QMimeData_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QMacPasteboardMime_Count(ptr.Pointer(), core.PointerFromQMimeData(mimeData))))
	}
	return 0
}

func (ptr *QMacPasteboardMime) CountDefault(mimeData core.QMimeData_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QMacPasteboardMime_CountDefault(ptr.Pointer(), core.PointerFromQMimeData(mimeData))))
	}
	return 0
}

//export callbackQMacPasteboardMime_FlavorFor
func callbackQMacPasteboardMime_FlavorFor(ptr unsafe.Pointer, mime C.struct_QtMacExtras_PackedString) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::flavorFor"); signal != nil {
		return C.CString(signal.(func(string) string)(cGoUnpackString(mime)))
	}

	return C.CString("")
}

func (ptr *QMacPasteboardMime) ConnectFlavorFor(f func(mime string) string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::flavorFor", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectFlavorFor(mime string) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::flavorFor")
	}
}

func (ptr *QMacPasteboardMime) FlavorFor(mime string) string {
	if ptr.Pointer() != nil {
		var mimeC = C.CString(mime)
		defer C.free(unsafe.Pointer(mimeC))
		return cGoUnpackString(C.QMacPasteboardMime_FlavorFor(ptr.Pointer(), mimeC))
	}
	return ""
}

//export callbackQMacPasteboardMime_MimeFor
func callbackQMacPasteboardMime_MimeFor(ptr unsafe.Pointer, flav C.struct_QtMacExtras_PackedString) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::mimeFor"); signal != nil {
		return C.CString(signal.(func(string) string)(cGoUnpackString(flav)))
	}

	return C.CString("")
}

func (ptr *QMacPasteboardMime) ConnectMimeFor(f func(flav string) string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::mimeFor", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectMimeFor(flav string) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::mimeFor")
	}
}

func (ptr *QMacPasteboardMime) MimeFor(flav string) string {
	if ptr.Pointer() != nil {
		var flavC = C.CString(flav)
		defer C.free(unsafe.Pointer(flavC))
		return cGoUnpackString(C.QMacPasteboardMime_MimeFor(ptr.Pointer(), flavC))
	}
	return ""
}

//export callbackQMacPasteboardMime_DestroyQMacPasteboardMime
func callbackQMacPasteboardMime_DestroyQMacPasteboardMime(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacPasteboardMime::~QMacPasteboardMime"); signal != nil {
		signal.(func())()
	} else {
		NewQMacPasteboardMimeFromPointer(ptr).DestroyQMacPasteboardMimeDefault()
	}
}

func (ptr *QMacPasteboardMime) ConnectDestroyQMacPasteboardMime(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::~QMacPasteboardMime", f)
	}
}

func (ptr *QMacPasteboardMime) DisconnectDestroyQMacPasteboardMime() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacPasteboardMime::~QMacPasteboardMime")
	}
}

func (ptr *QMacPasteboardMime) DestroyQMacPasteboardMime() {
	if ptr.Pointer() != nil {
		C.QMacPasteboardMime_DestroyQMacPasteboardMime(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacPasteboardMime) DestroyQMacPasteboardMimeDefault() {
	if ptr.Pointer() != nil {
		C.QMacPasteboardMime_DestroyQMacPasteboardMimeDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacPasteboardMime) convertFromMime_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QMacPasteboardMime_convertFromMime_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

type QMacToolBar struct {
	core.QObject
}

type QMacToolBar_ITF interface {
	core.QObject_ITF
	QMacToolBar_PTR() *QMacToolBar
}

func (p *QMacToolBar) QMacToolBar_PTR() *QMacToolBar {
	return p
}

func (p *QMacToolBar) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QMacToolBar) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQMacToolBar(ptr QMacToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBar_PTR().Pointer()
	}
	return nil
}

func NewQMacToolBarFromPointer(ptr unsafe.Pointer) *QMacToolBar {
	var n = new(QMacToolBar)
	n.SetPointer(ptr)
	return n
}
func NewQMacToolBar(parent core.QObject_ITF) *QMacToolBar {
	var tmpValue = NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQMacToolBar2(identifier string, parent core.QObject_ITF) *QMacToolBar {
	var identifierC = C.CString(identifier)
	defer C.free(unsafe.Pointer(identifierC))
	var tmpValue = NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar2(identifierC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMacToolBar) AddAllowedItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		var tmpValue = NewQMacToolBarItemFromPointer(C.QMacToolBar_AddAllowedItem(ptr.Pointer(), gui.PointerFromQIcon(icon), textC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) AddItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		var tmpValue = NewQMacToolBarItemFromPointer(C.QMacToolBar_AddItem(ptr.Pointer(), gui.PointerFromQIcon(icon), textC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) AddSeparator() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_AddSeparator(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) AllowedItems() []*QMacToolBarItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtMacExtras_PackedList) []*QMacToolBarItem {
			var out = make([]*QMacToolBarItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQMacToolBarFromPointer(l.data).allowedItems_atList(i)
			}
			return out
		}(C.QMacToolBar_AllowedItems(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBar) AttachToWindow(window gui.QWindow_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_AttachToWindow(ptr.Pointer(), gui.PointerFromQWindow(window))
	}
}

func (ptr *QMacToolBar) DetachFromWindow() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DetachFromWindow(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) Items() []*QMacToolBarItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtMacExtras_PackedList) []*QMacToolBarItem {
			var out = make([]*QMacToolBarItem, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQMacToolBarFromPointer(l.data).items_atList(i)
			}
			return out
		}(C.QMacToolBar_Items(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBar) DestroyQMacToolBar() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DestroyQMacToolBar(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBar) allowedItems_atList(i int) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQMacToolBarItemFromPointer(C.QMacToolBar_allowedItems_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBar) items_atList(i int) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		var tmpValue = NewQMacToolBarItemFromPointer(C.QMacToolBar_items_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQMacToolBar_TimerEvent
func callbackQMacToolBar_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::timerEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::timerEvent")
	}
}

func (ptr *QMacToolBar) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacToolBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMacToolBar_ChildEvent
func callbackQMacToolBar_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::childEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::childEvent")
	}
}

func (ptr *QMacToolBar) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacToolBar) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMacToolBar_ConnectNotify
func callbackQMacToolBar_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBar) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::connectNotify", f)
	}
}

func (ptr *QMacToolBar) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::connectNotify")
	}
}

func (ptr *QMacToolBar) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMacToolBar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBar_CustomEvent
func callbackQMacToolBar_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacToolBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacToolBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::customEvent", f)
	}
}

func (ptr *QMacToolBar) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::customEvent")
	}
}

func (ptr *QMacToolBar) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacToolBar) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMacToolBar_DeleteLater
func callbackQMacToolBar_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMacToolBar) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::deleteLater", f)
	}
}

func (ptr *QMacToolBar) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::deleteLater")
	}
}

func (ptr *QMacToolBar) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBar) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQMacToolBar_DisconnectNotify
func callbackQMacToolBar_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBar) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::disconnectNotify", f)
	}
}

func (ptr *QMacToolBar) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::disconnectNotify")
	}
}

func (ptr *QMacToolBar) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMacToolBar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBar_Event
func callbackQMacToolBar_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMacToolBar) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::event", f)
	}
}

func (ptr *QMacToolBar) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::event")
	}
}

func (ptr *QMacToolBar) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBar_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMacToolBar) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBar_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMacToolBar_EventFilter
func callbackQMacToolBar_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMacToolBar) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::eventFilter", f)
	}
}

func (ptr *QMacToolBar) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::eventFilter")
	}
}

func (ptr *QMacToolBar) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBar_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMacToolBar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBar_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMacToolBar_MetaObject
func callbackQMacToolBar_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBar::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMacToolBarFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMacToolBar) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::metaObject", f)
	}
}

func (ptr *QMacToolBar) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBar::metaObject")
	}
}

func (ptr *QMacToolBar) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBar_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBar) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBar_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QMacToolBarItem::StandardItem
type QMacToolBarItem__StandardItem int64

const (
	QMacToolBarItem__NoStandardItem = QMacToolBarItem__StandardItem(0)
	QMacToolBarItem__Space          = QMacToolBarItem__StandardItem(1)
	QMacToolBarItem__FlexibleSpace  = QMacToolBarItem__StandardItem(2)
)

type QMacToolBarItem struct {
	core.QObject
}

type QMacToolBarItem_ITF interface {
	core.QObject_ITF
	QMacToolBarItem_PTR() *QMacToolBarItem
}

func (p *QMacToolBarItem) QMacToolBarItem_PTR() *QMacToolBarItem {
	return p
}

func (p *QMacToolBarItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QMacToolBarItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQMacToolBarItem(ptr QMacToolBarItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBarItem_PTR().Pointer()
	}
	return nil
}

func NewQMacToolBarItemFromPointer(ptr unsafe.Pointer) *QMacToolBarItem {
	var n = new(QMacToolBarItem)
	n.SetPointer(ptr)
	return n
}
func NewQMacToolBarItem(parent core.QObject_ITF) *QMacToolBarItem {
	var tmpValue = NewQMacToolBarItemFromPointer(C.QMacToolBarItem_NewQMacToolBarItem(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQMacToolBarItem_Activated
func callbackQMacToolBarItem_Activated(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::activated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMacToolBarItem) ConnectActivated(f func()) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::activated", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::activated")
	}
}

func (ptr *QMacToolBarItem) Activated() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_Activated(ptr.Pointer())
	}
}

//export callbackQMacToolBarItem_DestroyQMacToolBarItem
func callbackQMacToolBarItem_DestroyQMacToolBarItem(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::~QMacToolBarItem"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarItemFromPointer(ptr).DestroyQMacToolBarItemDefault()
	}
}

func (ptr *QMacToolBarItem) ConnectDestroyQMacToolBarItem(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::~QMacToolBarItem", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectDestroyQMacToolBarItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::~QMacToolBarItem")
	}
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItem() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DestroyQMacToolBarItem(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItemDefault() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DestroyQMacToolBarItemDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBarItem) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QMacToolBarItem_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QMacToolBarItem) Selectable() bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_Selectable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMacToolBarItem) SetSelectable(selectable bool) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetSelectable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(selectable))))
	}
}

func (ptr *QMacToolBarItem) SetStandardItem(standardItem QMacToolBarItem__StandardItem) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetStandardItem(ptr.Pointer(), C.longlong(standardItem))
	}
}

func (ptr *QMacToolBarItem) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QMacToolBarItem_SetText(ptr.Pointer(), textC)
	}
}

func (ptr *QMacToolBarItem) StandardItem() QMacToolBarItem__StandardItem {
	if ptr.Pointer() != nil {
		return QMacToolBarItem__StandardItem(C.QMacToolBarItem_StandardItem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMacToolBarItem) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QMacToolBarItem_Text(ptr.Pointer()))
	}
	return ""
}

//export callbackQMacToolBarItem_TimerEvent
func callbackQMacToolBarItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::timerEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::timerEvent")
	}
}

func (ptr *QMacToolBarItem) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacToolBarItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMacToolBarItem_ChildEvent
func callbackQMacToolBarItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::childEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::childEvent")
	}
}

func (ptr *QMacToolBarItem) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacToolBarItem) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMacToolBarItem_ConnectNotify
func callbackQMacToolBarItem_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBarItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::connectNotify", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::connectNotify")
	}
}

func (ptr *QMacToolBarItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMacToolBarItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBarItem_CustomEvent
func callbackQMacToolBarItem_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::customEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::customEvent")
	}
}

func (ptr *QMacToolBarItem) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacToolBarItem) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMacToolBarItem_DeleteLater
func callbackQMacToolBarItem_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMacToolBarItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMacToolBarItem) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::deleteLater", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::deleteLater")
	}
}

func (ptr *QMacToolBarItem) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBarItem) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQMacToolBarItem_DisconnectNotify
func callbackQMacToolBarItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMacToolBarItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMacToolBarItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::disconnectNotify", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::disconnectNotify")
	}
}

func (ptr *QMacToolBarItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMacToolBarItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMacToolBarItem_Event
func callbackQMacToolBarItem_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarItemFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMacToolBarItem) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::event", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::event")
	}
}

func (ptr *QMacToolBarItem) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMacToolBarItem_EventFilter
func callbackQMacToolBarItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMacToolBarItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMacToolBarItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::eventFilter", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::eventFilter")
	}
}

func (ptr *QMacToolBarItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMacToolBarItem_MetaObject
func callbackQMacToolBarItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMacToolBarItem::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMacToolBarItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMacToolBarItem) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::metaObject", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMacToolBarItem::metaObject")
	}
}

func (ptr *QMacToolBarItem) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBarItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBarItem) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMacToolBarItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
