// +build !minimal

package bluetooth

//#include <stdint.h>
//#include <stdlib.h>
//#include "bluetooth.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtBluetooth_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QBluetoothAddress struct {
	ptr unsafe.Pointer
}

type QBluetoothAddress_ITF interface {
	QBluetoothAddress_PTR() *QBluetoothAddress
}

func (p *QBluetoothAddress) QBluetoothAddress_PTR() *QBluetoothAddress {
	return p
}

func (p *QBluetoothAddress) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothAddress) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothAddress(ptr QBluetoothAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothAddress_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothAddressFromPointer(ptr unsafe.Pointer) *QBluetoothAddress {
	var n = new(QBluetoothAddress)
	n.SetPointer(ptr)
	return n
}
func NewQBluetoothAddress() *QBluetoothAddress {
	var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress())
	runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress4(other QBluetoothAddress_ITF) *QBluetoothAddress {
	var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress4(PointerFromQBluetoothAddress(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress3(address string) *QBluetoothAddress {
	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress3(addressC))
	runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress2(address uint64) *QBluetoothAddress {
	var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress2(C.ulonglong(address)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func (ptr *QBluetoothAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QBluetoothAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothAddress) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothAddress) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothAddress) ToUInt64() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QBluetoothAddress_ToUInt64(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothAddress) DestroyQBluetoothAddress() {
	if ptr.Pointer() != nil {
		C.QBluetoothAddress_DestroyQBluetoothAddress(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QBluetoothDeviceDiscoveryAgent::Error
type QBluetoothDeviceDiscoveryAgent__Error int64

const (
	QBluetoothDeviceDiscoveryAgent__NoError                      = QBluetoothDeviceDiscoveryAgent__Error(0)
	QBluetoothDeviceDiscoveryAgent__InputOutputError             = QBluetoothDeviceDiscoveryAgent__Error(1)
	QBluetoothDeviceDiscoveryAgent__PoweredOffError              = QBluetoothDeviceDiscoveryAgent__Error(2)
	QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothDeviceDiscoveryAgent__Error(3)
	QBluetoothDeviceDiscoveryAgent__UnsupportedPlatformError     = QBluetoothDeviceDiscoveryAgent__Error(4)
	QBluetoothDeviceDiscoveryAgent__UnknownError                 = QBluetoothDeviceDiscoveryAgent__Error(100)
)

//QBluetoothDeviceDiscoveryAgent::InquiryType
type QBluetoothDeviceDiscoveryAgent__InquiryType int64

const (
	QBluetoothDeviceDiscoveryAgent__GeneralUnlimitedInquiry = QBluetoothDeviceDiscoveryAgent__InquiryType(0)
	QBluetoothDeviceDiscoveryAgent__LimitedInquiry          = QBluetoothDeviceDiscoveryAgent__InquiryType(1)
)

type QBluetoothDeviceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothDeviceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent
}

func (p *QBluetoothDeviceDiscoveryAgent) QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent {
	return p
}

func (p *QBluetoothDeviceDiscoveryAgent) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothDeviceDiscoveryAgent) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothDeviceDiscoveryAgent(ptr QBluetoothDeviceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceDiscoveryAgent {
	var n = new(QBluetoothDeviceDiscoveryAgent)
	n.SetPointer(ptr)
	return n
}

//export callbackQBluetoothDeviceDiscoveryAgent_Canceled
func callbackQBluetoothDeviceDiscoveryAgent_Canceled(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCanceled(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::canceled", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCanceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::canceled")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Canceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Canceled(ptr.Pointer())
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_DeviceDiscovered
func callbackQBluetoothDeviceDiscoveryAgent_DeviceDiscovered(ptr unsafe.Pointer, info unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::deviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothDeviceInfo))(NewQBluetoothDeviceInfoFromPointer(info))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeviceDiscovered(f func(info *QBluetoothDeviceInfo)) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectDeviceDiscovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::deviceDiscovered", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDeviceDiscovered() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectDeviceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::deviceDiscovered")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeviceDiscovered(info QBluetoothDeviceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DeviceDiscovered(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(info))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Error2
func callbackQBluetoothDeviceDiscoveryAgent_Error2(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::error2"); signal != nil {
		signal.(func(QBluetoothDeviceDiscoveryAgent__Error))(QBluetoothDeviceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectError2(f func(error QBluetoothDeviceDiscoveryAgent__Error)) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::error2", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::error2")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error2(error QBluetoothDeviceDiscoveryAgent__Error) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Error2(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Finished
func callbackQBluetoothDeviceDiscoveryAgent_Finished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::finished", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::finished")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Finished() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Finished(ptr.Pointer())
	}
}

func NewQBluetoothDeviceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	var tmpValue = NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothDeviceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	var tmpValue = NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DiscoveredDevices() []*QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothDeviceInfo {
			var out = make([]*QBluetoothDeviceInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothDeviceDiscoveryAgentFromPointer(l.data).discoveredDevices_atList(i)
			}
			return out
		}(C.QBluetoothDeviceDiscoveryAgent_DiscoveredDevices(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error() QBluetoothDeviceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__Error(C.QBluetoothDeviceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceDiscoveryAgent) InquiryType() QBluetoothDeviceDiscoveryAgent__InquiryType {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__InquiryType(C.QBluetoothDeviceDiscoveryAgent_InquiryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetInquiryType(ty QBluetoothDeviceDiscoveryAgent__InquiryType) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_SetInquiryType(ptr.Pointer(), C.longlong(ty))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Start
func callbackQBluetoothDeviceDiscoveryAgent_Start(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::start"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStart(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::start", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::start")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Start(ptr.Pointer())
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Stop
func callbackQBluetoothDeviceDiscoveryAgent_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::stop", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::stop")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Stop() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgent() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) discoveredDevices_atList(i int) *QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceDiscoveryAgent_discoveredDevices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothDeviceDiscoveryAgent_TimerEvent
func callbackQBluetoothDeviceDiscoveryAgent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::timerEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::timerEvent")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_ChildEvent
func callbackQBluetoothDeviceDiscoveryAgent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::childEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::childEvent")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_ConnectNotify
func callbackQBluetoothDeviceDiscoveryAgent_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::connectNotify", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::connectNotify")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_CustomEvent
func callbackQBluetoothDeviceDiscoveryAgent_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::customEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::customEvent")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_DeleteLater
func callbackQBluetoothDeviceDiscoveryAgent_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::deleteLater", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::deleteLater")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify
func callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::disconnectNotify", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::disconnectNotify")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Event
func callbackQBluetoothDeviceDiscoveryAgent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::event", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::event")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothDeviceDiscoveryAgent_EventFilter
func callbackQBluetoothDeviceDiscoveryAgent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::eventFilter", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::eventFilter")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothDeviceDiscoveryAgent_MetaObject
func callbackQBluetoothDeviceDiscoveryAgent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothDeviceDiscoveryAgent::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::metaObject", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothDeviceDiscoveryAgent::metaObject")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothDeviceInfo::CoreConfiguration
type QBluetoothDeviceInfo__CoreConfiguration int64

const (
	QBluetoothDeviceInfo__UnknownCoreConfiguration              = QBluetoothDeviceInfo__CoreConfiguration(0x0)
	QBluetoothDeviceInfo__LowEnergyCoreConfiguration            = QBluetoothDeviceInfo__CoreConfiguration(0x01)
	QBluetoothDeviceInfo__BaseRateCoreConfiguration             = QBluetoothDeviceInfo__CoreConfiguration(0x02)
	QBluetoothDeviceInfo__BaseRateAndLowEnergyCoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x03)
)

//QBluetoothDeviceInfo::DataCompleteness
type QBluetoothDeviceInfo__DataCompleteness int64

const (
	QBluetoothDeviceInfo__DataComplete    = QBluetoothDeviceInfo__DataCompleteness(0)
	QBluetoothDeviceInfo__DataIncomplete  = QBluetoothDeviceInfo__DataCompleteness(1)
	QBluetoothDeviceInfo__DataUnavailable = QBluetoothDeviceInfo__DataCompleteness(2)
)

//QBluetoothDeviceInfo::MajorDeviceClass
type QBluetoothDeviceInfo__MajorDeviceClass int64

const (
	QBluetoothDeviceInfo__MiscellaneousDevice = QBluetoothDeviceInfo__MajorDeviceClass(0)
	QBluetoothDeviceInfo__ComputerDevice      = QBluetoothDeviceInfo__MajorDeviceClass(1)
	QBluetoothDeviceInfo__PhoneDevice         = QBluetoothDeviceInfo__MajorDeviceClass(2)
	QBluetoothDeviceInfo__LANAccessDevice     = QBluetoothDeviceInfo__MajorDeviceClass(3)
	QBluetoothDeviceInfo__AudioVideoDevice    = QBluetoothDeviceInfo__MajorDeviceClass(4)
	QBluetoothDeviceInfo__PeripheralDevice    = QBluetoothDeviceInfo__MajorDeviceClass(5)
	QBluetoothDeviceInfo__ImagingDevice       = QBluetoothDeviceInfo__MajorDeviceClass(6)
	QBluetoothDeviceInfo__WearableDevice      = QBluetoothDeviceInfo__MajorDeviceClass(7)
	QBluetoothDeviceInfo__ToyDevice           = QBluetoothDeviceInfo__MajorDeviceClass(8)
	QBluetoothDeviceInfo__HealthDevice        = QBluetoothDeviceInfo__MajorDeviceClass(9)
	QBluetoothDeviceInfo__UncategorizedDevice = QBluetoothDeviceInfo__MajorDeviceClass(31)
)

//QBluetoothDeviceInfo::MinorAudioVideoClass
type QBluetoothDeviceInfo__MinorAudioVideoClass int64

const (
	QBluetoothDeviceInfo__UncategorizedAudioVideoDevice = QBluetoothDeviceInfo__MinorAudioVideoClass(0)
	QBluetoothDeviceInfo__WearableHeadsetDevice         = QBluetoothDeviceInfo__MinorAudioVideoClass(1)
	QBluetoothDeviceInfo__HandsFreeDevice               = QBluetoothDeviceInfo__MinorAudioVideoClass(2)
	QBluetoothDeviceInfo__Microphone                    = QBluetoothDeviceInfo__MinorAudioVideoClass(4)
	QBluetoothDeviceInfo__Loudspeaker                   = QBluetoothDeviceInfo__MinorAudioVideoClass(5)
	QBluetoothDeviceInfo__Headphones                    = QBluetoothDeviceInfo__MinorAudioVideoClass(6)
	QBluetoothDeviceInfo__PortableAudioDevice           = QBluetoothDeviceInfo__MinorAudioVideoClass(7)
	QBluetoothDeviceInfo__CarAudio                      = QBluetoothDeviceInfo__MinorAudioVideoClass(8)
	QBluetoothDeviceInfo__SetTopBox                     = QBluetoothDeviceInfo__MinorAudioVideoClass(9)
	QBluetoothDeviceInfo__HiFiAudioDevice               = QBluetoothDeviceInfo__MinorAudioVideoClass(10)
	QBluetoothDeviceInfo__Vcr                           = QBluetoothDeviceInfo__MinorAudioVideoClass(11)
	QBluetoothDeviceInfo__VideoCamera                   = QBluetoothDeviceInfo__MinorAudioVideoClass(12)
	QBluetoothDeviceInfo__Camcorder                     = QBluetoothDeviceInfo__MinorAudioVideoClass(13)
	QBluetoothDeviceInfo__VideoMonitor                  = QBluetoothDeviceInfo__MinorAudioVideoClass(14)
	QBluetoothDeviceInfo__VideoDisplayAndLoudspeaker    = QBluetoothDeviceInfo__MinorAudioVideoClass(15)
	QBluetoothDeviceInfo__VideoConferencing             = QBluetoothDeviceInfo__MinorAudioVideoClass(16)
	QBluetoothDeviceInfo__GamingDevice                  = QBluetoothDeviceInfo__MinorAudioVideoClass(18)
)

//QBluetoothDeviceInfo::MinorComputerClass
type QBluetoothDeviceInfo__MinorComputerClass int64

const (
	QBluetoothDeviceInfo__UncategorizedComputer     = QBluetoothDeviceInfo__MinorComputerClass(0)
	QBluetoothDeviceInfo__DesktopComputer           = QBluetoothDeviceInfo__MinorComputerClass(1)
	QBluetoothDeviceInfo__ServerComputer            = QBluetoothDeviceInfo__MinorComputerClass(2)
	QBluetoothDeviceInfo__LaptopComputer            = QBluetoothDeviceInfo__MinorComputerClass(3)
	QBluetoothDeviceInfo__HandheldClamShellComputer = QBluetoothDeviceInfo__MinorComputerClass(4)
	QBluetoothDeviceInfo__HandheldComputer          = QBluetoothDeviceInfo__MinorComputerClass(5)
	QBluetoothDeviceInfo__WearableComputer          = QBluetoothDeviceInfo__MinorComputerClass(6)
)

//QBluetoothDeviceInfo::MinorHealthClass
type QBluetoothDeviceInfo__MinorHealthClass int64

const (
	QBluetoothDeviceInfo__UncategorizedHealthDevice  = QBluetoothDeviceInfo__MinorHealthClass(0)
	QBluetoothDeviceInfo__HealthBloodPressureMonitor = QBluetoothDeviceInfo__MinorHealthClass(0x1)
	QBluetoothDeviceInfo__HealthThermometer          = QBluetoothDeviceInfo__MinorHealthClass(0x2)
	QBluetoothDeviceInfo__HealthWeightScale          = QBluetoothDeviceInfo__MinorHealthClass(0x3)
	QBluetoothDeviceInfo__HealthGlucoseMeter         = QBluetoothDeviceInfo__MinorHealthClass(0x4)
	QBluetoothDeviceInfo__HealthPulseOximeter        = QBluetoothDeviceInfo__MinorHealthClass(0x5)
	QBluetoothDeviceInfo__HealthDataDisplay          = QBluetoothDeviceInfo__MinorHealthClass(0x7)
	QBluetoothDeviceInfo__HealthStepCounter          = QBluetoothDeviceInfo__MinorHealthClass(0x8)
)

//QBluetoothDeviceInfo::MinorImagingClass
type QBluetoothDeviceInfo__MinorImagingClass int64

const (
	QBluetoothDeviceInfo__UncategorizedImagingDevice = QBluetoothDeviceInfo__MinorImagingClass(0)
	QBluetoothDeviceInfo__ImageDisplay               = QBluetoothDeviceInfo__MinorImagingClass(0x04)
	QBluetoothDeviceInfo__ImageCamera                = QBluetoothDeviceInfo__MinorImagingClass(0x08)
	QBluetoothDeviceInfo__ImageScanner               = QBluetoothDeviceInfo__MinorImagingClass(0x10)
	QBluetoothDeviceInfo__ImagePrinter               = QBluetoothDeviceInfo__MinorImagingClass(0x20)
)

//QBluetoothDeviceInfo::MinorMiscellaneousClass
type QBluetoothDeviceInfo__MinorMiscellaneousClass int64

const (
	QBluetoothDeviceInfo__UncategorizedMiscellaneous = QBluetoothDeviceInfo__MinorMiscellaneousClass(0)
)

//QBluetoothDeviceInfo::MinorNetworkClass
type QBluetoothDeviceInfo__MinorNetworkClass int64

const (
	QBluetoothDeviceInfo__NetworkFullService     = QBluetoothDeviceInfo__MinorNetworkClass(0x00)
	QBluetoothDeviceInfo__NetworkLoadFactorOne   = QBluetoothDeviceInfo__MinorNetworkClass(0x08)
	QBluetoothDeviceInfo__NetworkLoadFactorTwo   = QBluetoothDeviceInfo__MinorNetworkClass(0x10)
	QBluetoothDeviceInfo__NetworkLoadFactorThree = QBluetoothDeviceInfo__MinorNetworkClass(0x18)
	QBluetoothDeviceInfo__NetworkLoadFactorFour  = QBluetoothDeviceInfo__MinorNetworkClass(0x20)
	QBluetoothDeviceInfo__NetworkLoadFactorFive  = QBluetoothDeviceInfo__MinorNetworkClass(0x28)
	QBluetoothDeviceInfo__NetworkLoadFactorSix   = QBluetoothDeviceInfo__MinorNetworkClass(0x30)
	QBluetoothDeviceInfo__NetworkNoService       = QBluetoothDeviceInfo__MinorNetworkClass(0x38)
)

//QBluetoothDeviceInfo::MinorPeripheralClass
type QBluetoothDeviceInfo__MinorPeripheralClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0)
	QBluetoothDeviceInfo__KeyboardPeripheral                   = QBluetoothDeviceInfo__MinorPeripheralClass(0x10)
	QBluetoothDeviceInfo__PointingDevicePeripheral             = QBluetoothDeviceInfo__MinorPeripheralClass(0x20)
	QBluetoothDeviceInfo__KeyboardWithPointingDevicePeripheral = QBluetoothDeviceInfo__MinorPeripheralClass(0x30)
	QBluetoothDeviceInfo__JoystickPeripheral                   = QBluetoothDeviceInfo__MinorPeripheralClass(0x01)
	QBluetoothDeviceInfo__GamepadPeripheral                    = QBluetoothDeviceInfo__MinorPeripheralClass(0x02)
	QBluetoothDeviceInfo__RemoteControlPeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0x03)
	QBluetoothDeviceInfo__SensingDevicePeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0x04)
	QBluetoothDeviceInfo__DigitizerTabletPeripheral            = QBluetoothDeviceInfo__MinorPeripheralClass(0x05)
	QBluetoothDeviceInfo__CardReaderPeripheral                 = QBluetoothDeviceInfo__MinorPeripheralClass(0x06)
)

//QBluetoothDeviceInfo::MinorPhoneClass
type QBluetoothDeviceInfo__MinorPhoneClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPhone            = QBluetoothDeviceInfo__MinorPhoneClass(0)
	QBluetoothDeviceInfo__CellularPhone                 = QBluetoothDeviceInfo__MinorPhoneClass(1)
	QBluetoothDeviceInfo__CordlessPhone                 = QBluetoothDeviceInfo__MinorPhoneClass(2)
	QBluetoothDeviceInfo__SmartPhone                    = QBluetoothDeviceInfo__MinorPhoneClass(3)
	QBluetoothDeviceInfo__WiredModemOrVoiceGatewayPhone = QBluetoothDeviceInfo__MinorPhoneClass(4)
	QBluetoothDeviceInfo__CommonIsdnAccessPhone         = QBluetoothDeviceInfo__MinorPhoneClass(5)
)

//QBluetoothDeviceInfo::MinorToyClass
type QBluetoothDeviceInfo__MinorToyClass int64

const (
	QBluetoothDeviceInfo__UncategorizedToy = QBluetoothDeviceInfo__MinorToyClass(0)
	QBluetoothDeviceInfo__ToyRobot         = QBluetoothDeviceInfo__MinorToyClass(1)
	QBluetoothDeviceInfo__ToyVehicle       = QBluetoothDeviceInfo__MinorToyClass(2)
	QBluetoothDeviceInfo__ToyDoll          = QBluetoothDeviceInfo__MinorToyClass(3)
	QBluetoothDeviceInfo__ToyController    = QBluetoothDeviceInfo__MinorToyClass(4)
	QBluetoothDeviceInfo__ToyGame          = QBluetoothDeviceInfo__MinorToyClass(5)
)

//QBluetoothDeviceInfo::MinorWearableClass
type QBluetoothDeviceInfo__MinorWearableClass int64

const (
	QBluetoothDeviceInfo__UncategorizedWearableDevice = QBluetoothDeviceInfo__MinorWearableClass(0)
	QBluetoothDeviceInfo__WearableWristWatch          = QBluetoothDeviceInfo__MinorWearableClass(1)
	QBluetoothDeviceInfo__WearablePager               = QBluetoothDeviceInfo__MinorWearableClass(2)
	QBluetoothDeviceInfo__WearableJacket              = QBluetoothDeviceInfo__MinorWearableClass(3)
	QBluetoothDeviceInfo__WearableHelmet              = QBluetoothDeviceInfo__MinorWearableClass(4)
	QBluetoothDeviceInfo__WearableGlasses             = QBluetoothDeviceInfo__MinorWearableClass(5)
)

//QBluetoothDeviceInfo::ServiceClass
type QBluetoothDeviceInfo__ServiceClass int64

const (
	QBluetoothDeviceInfo__NoService             = QBluetoothDeviceInfo__ServiceClass(0x0000)
	QBluetoothDeviceInfo__PositioningService    = QBluetoothDeviceInfo__ServiceClass(0x0001)
	QBluetoothDeviceInfo__NetworkingService     = QBluetoothDeviceInfo__ServiceClass(0x0002)
	QBluetoothDeviceInfo__RenderingService      = QBluetoothDeviceInfo__ServiceClass(0x0004)
	QBluetoothDeviceInfo__CapturingService      = QBluetoothDeviceInfo__ServiceClass(0x0008)
	QBluetoothDeviceInfo__ObjectTransferService = QBluetoothDeviceInfo__ServiceClass(0x0010)
	QBluetoothDeviceInfo__AudioService          = QBluetoothDeviceInfo__ServiceClass(0x0020)
	QBluetoothDeviceInfo__TelephonyService      = QBluetoothDeviceInfo__ServiceClass(0x0040)
	QBluetoothDeviceInfo__InformationService    = QBluetoothDeviceInfo__ServiceClass(0x0080)
	QBluetoothDeviceInfo__AllServices           = QBluetoothDeviceInfo__ServiceClass(0x07ff)
)

type QBluetoothDeviceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothDeviceInfo_ITF interface {
	QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo
}

func (p *QBluetoothDeviceInfo) QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo {
	return p
}

func (p *QBluetoothDeviceInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothDeviceInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothDeviceInfo(ptr QBluetoothDeviceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothDeviceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceInfo {
	var n = new(QBluetoothDeviceInfo)
	n.SetPointer(ptr)
	return n
}
func NewQBluetoothDeviceInfo() *QBluetoothDeviceInfo {
	var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo())
	runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo2(address QBluetoothAddress_ITF, name string, classOfDevice uint) *QBluetoothDeviceInfo {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo2(PointerFromQBluetoothAddress(address), nameC, C.uint(uint32(classOfDevice))))
	runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo4(other QBluetoothDeviceInfo_ITF) *QBluetoothDeviceInfo {
	var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(PointerFromQBluetoothDeviceInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo3(uuid QBluetoothUuid_ITF, name string, classOfDevice uint) *QBluetoothDeviceInfo {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo3(PointerFromQBluetoothUuid(uuid), nameC, C.uint(uint32(classOfDevice))))
	runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func (ptr *QBluetoothDeviceInfo) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothDeviceInfo_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) CoreConfigurations() QBluetoothDeviceInfo__CoreConfiguration {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__CoreConfiguration(C.QBluetoothDeviceInfo_CoreConfigurations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) DeviceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo_DeviceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) IsCached() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsCached(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) MajorDeviceClass() QBluetoothDeviceInfo__MajorDeviceClass {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__MajorDeviceClass(C.QBluetoothDeviceInfo_MajorDeviceClass(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) MinorDeviceClass() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceInfo_MinorDeviceClass(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) Rssi() int16 {
	if ptr.Pointer() != nil {
		return int16(C.QBluetoothDeviceInfo_Rssi(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) ServiceClasses() QBluetoothDeviceInfo__ServiceClass {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__ServiceClass(C.QBluetoothDeviceInfo_ServiceClasses(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) ServiceUuidsCompleteness() QBluetoothDeviceInfo__DataCompleteness {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__DataCompleteness(C.QBluetoothDeviceInfo_ServiceUuidsCompleteness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) SetCached(cached bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCached(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(cached))))
	}
}

func (ptr *QBluetoothDeviceInfo) SetCoreConfigurations(coreConfigs QBluetoothDeviceInfo__CoreConfiguration) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCoreConfigurations(ptr.Pointer(), C.longlong(coreConfigs))
	}
}

func (ptr *QBluetoothDeviceInfo) SetDeviceUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetDeviceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothDeviceInfo) SetRssi(sign int16) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetRssi(ptr.Pointer(), C.short(sign))
	}
}

func (ptr *QBluetoothDeviceInfo) DestroyQBluetoothDeviceInfo() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceInfo) serviceUuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo_serviceUuids_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

type QBluetoothHostInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothHostInfo_ITF interface {
	QBluetoothHostInfo_PTR() *QBluetoothHostInfo
}

func (p *QBluetoothHostInfo) QBluetoothHostInfo_PTR() *QBluetoothHostInfo {
	return p
}

func (p *QBluetoothHostInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothHostInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothHostInfo(ptr QBluetoothHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothHostInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothHostInfoFromPointer(ptr unsafe.Pointer) *QBluetoothHostInfo {
	var n = new(QBluetoothHostInfo)
	n.SetPointer(ptr)
	return n
}
func NewQBluetoothHostInfo() *QBluetoothHostInfo {
	var tmpValue = NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo())
	runtime.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
	return tmpValue
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {
	var tmpValue = NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(PointerFromQBluetoothHostInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
	return tmpValue
}

func (ptr *QBluetoothHostInfo) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothHostInfo_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothHostInfo) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothHostInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothHostInfo) SetAddress(address QBluetoothAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

func (ptr *QBluetoothHostInfo) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QBluetoothHostInfo_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {
	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QBluetoothLocalDevice::Error
type QBluetoothLocalDevice__Error int64

const (
	QBluetoothLocalDevice__NoError      = QBluetoothLocalDevice__Error(0)
	QBluetoothLocalDevice__PairingError = QBluetoothLocalDevice__Error(1)
	QBluetoothLocalDevice__UnknownError = QBluetoothLocalDevice__Error(100)
)

//QBluetoothLocalDevice::HostMode
type QBluetoothLocalDevice__HostMode int64

const (
	QBluetoothLocalDevice__HostPoweredOff                 = QBluetoothLocalDevice__HostMode(0)
	QBluetoothLocalDevice__HostConnectable                = QBluetoothLocalDevice__HostMode(1)
	QBluetoothLocalDevice__HostDiscoverable               = QBluetoothLocalDevice__HostMode(2)
	QBluetoothLocalDevice__HostDiscoverableLimitedInquiry = QBluetoothLocalDevice__HostMode(3)
)

//QBluetoothLocalDevice::Pairing
type QBluetoothLocalDevice__Pairing int64

const (
	QBluetoothLocalDevice__Unpaired         = QBluetoothLocalDevice__Pairing(0)
	QBluetoothLocalDevice__Paired           = QBluetoothLocalDevice__Pairing(1)
	QBluetoothLocalDevice__AuthorizedPaired = QBluetoothLocalDevice__Pairing(2)
)

type QBluetoothLocalDevice struct {
	core.QObject
}

type QBluetoothLocalDevice_ITF interface {
	core.QObject_ITF
	QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice
}

func (p *QBluetoothLocalDevice) QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice {
	return p
}

func (p *QBluetoothLocalDevice) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothLocalDevice) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothLocalDevice(ptr QBluetoothLocalDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothLocalDevice_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) *QBluetoothLocalDevice {
	var n = new(QBluetoothLocalDevice)
	n.SetPointer(ptr)
	return n
}

//export callbackQBluetoothLocalDevice_DeviceConnected
func callbackQBluetoothLocalDevice_DeviceConnected(ptr unsafe.Pointer, address unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::deviceConnected"); signal != nil {
		signal.(func(*QBluetoothAddress))(NewQBluetoothAddressFromPointer(address))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectDeviceConnected(f func(address *QBluetoothAddress)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectDeviceConnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::deviceConnected", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDeviceConnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectDeviceConnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::deviceConnected")
	}
}

func (ptr *QBluetoothLocalDevice) DeviceConnected(address QBluetoothAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeviceConnected(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

//export callbackQBluetoothLocalDevice_DeviceDisconnected
func callbackQBluetoothLocalDevice_DeviceDisconnected(ptr unsafe.Pointer, address unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::deviceDisconnected"); signal != nil {
		signal.(func(*QBluetoothAddress))(NewQBluetoothAddressFromPointer(address))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectDeviceDisconnected(f func(address *QBluetoothAddress)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectDeviceDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::deviceDisconnected", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDeviceDisconnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectDeviceDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::deviceDisconnected")
	}
}

func (ptr *QBluetoothLocalDevice) DeviceDisconnected(address QBluetoothAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeviceDisconnected(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

//export callbackQBluetoothLocalDevice_Error
func callbackQBluetoothLocalDevice_Error(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::error"); signal != nil {
		signal.(func(QBluetoothLocalDevice__Error))(QBluetoothLocalDevice__Error(error))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectError(f func(error QBluetoothLocalDevice__Error)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::error", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::error")
	}
}

func (ptr *QBluetoothLocalDevice) Error(error QBluetoothLocalDevice__Error) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_Error(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothLocalDevice_HostModeStateChanged
func callbackQBluetoothLocalDevice_HostModeStateChanged(ptr unsafe.Pointer, state C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::hostModeStateChanged"); signal != nil {
		signal.(func(QBluetoothLocalDevice__HostMode))(QBluetoothLocalDevice__HostMode(state))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectHostModeStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::hostModeStateChanged", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectHostModeStateChanged() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectHostModeStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::hostModeStateChanged")
	}
}

func (ptr *QBluetoothLocalDevice) HostModeStateChanged(state QBluetoothLocalDevice__HostMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_HostModeStateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQBluetoothLocalDevice_PairingDisplayConfirmation
func callbackQBluetoothLocalDevice_PairingDisplayConfirmation(ptr unsafe.Pointer, address unsafe.Pointer, pin C.struct_QtBluetooth_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::pairingDisplayConfirmation"); signal != nil {
		signal.(func(*QBluetoothAddress, string))(NewQBluetoothAddressFromPointer(address), cGoUnpackString(pin))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayConfirmation(f func(address *QBluetoothAddress, pin string)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectPairingDisplayConfirmation(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::pairingDisplayConfirmation", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingDisplayConfirmation() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingDisplayConfirmation(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::pairingDisplayConfirmation")
	}
}

func (ptr *QBluetoothLocalDevice) PairingDisplayConfirmation(address QBluetoothAddress_ITF, pin string) {
	if ptr.Pointer() != nil {
		var pinC = C.CString(pin)
		defer C.free(unsafe.Pointer(pinC))
		C.QBluetoothLocalDevice_PairingDisplayConfirmation(ptr.Pointer(), PointerFromQBluetoothAddress(address), pinC)
	}
}

//export callbackQBluetoothLocalDevice_PairingDisplayPinCode
func callbackQBluetoothLocalDevice_PairingDisplayPinCode(ptr unsafe.Pointer, address unsafe.Pointer, pin C.struct_QtBluetooth_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::pairingDisplayPinCode"); signal != nil {
		signal.(func(*QBluetoothAddress, string))(NewQBluetoothAddressFromPointer(address), cGoUnpackString(pin))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayPinCode(f func(address *QBluetoothAddress, pin string)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectPairingDisplayPinCode(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::pairingDisplayPinCode", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingDisplayPinCode() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingDisplayPinCode(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::pairingDisplayPinCode")
	}
}

func (ptr *QBluetoothLocalDevice) PairingDisplayPinCode(address QBluetoothAddress_ITF, pin string) {
	if ptr.Pointer() != nil {
		var pinC = C.CString(pin)
		defer C.free(unsafe.Pointer(pinC))
		C.QBluetoothLocalDevice_PairingDisplayPinCode(ptr.Pointer(), PointerFromQBluetoothAddress(address), pinC)
	}
}

//export callbackQBluetoothLocalDevice_PairingFinished
func callbackQBluetoothLocalDevice_PairingFinished(ptr unsafe.Pointer, address unsafe.Pointer, pairing C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::pairingFinished"); signal != nil {
		signal.(func(*QBluetoothAddress, QBluetoothLocalDevice__Pairing))(NewQBluetoothAddressFromPointer(address), QBluetoothLocalDevice__Pairing(pairing))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingFinished(f func(address *QBluetoothAddress, pairing QBluetoothLocalDevice__Pairing)) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectPairingFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::pairingFinished", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::pairingFinished")
	}
}

func (ptr *QBluetoothLocalDevice) PairingFinished(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingFinished(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.longlong(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_IsValid(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_DestroyQBluetoothLocalDevice
func callbackQBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::~QBluetoothLocalDevice"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DestroyQBluetoothLocalDeviceDefault()
	}
}

func (ptr *QBluetoothLocalDevice) ConnectDestroyQBluetoothLocalDevice(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::~QBluetoothLocalDevice", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDestroyQBluetoothLocalDevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::~QBluetoothLocalDevice")
	}
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDevice() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDeviceDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDeviceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {
	var tmpValue = NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {
	var tmpValue = NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(PointerFromQBluetoothAddress(address), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothLocalDevice) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothLocalDevice_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func QBluetoothLocalDevice_AllDevices() []*QBluetoothHostInfo {
	return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothHostInfo {
		var out = make([]*QBluetoothHostInfo, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQBluetoothLocalDeviceFromPointer(l.data).allDevices_atList(i)
		}
		return out
	}(C.QBluetoothLocalDevice_QBluetoothLocalDevice_AllDevices())
}

func (ptr *QBluetoothLocalDevice) AllDevices() []*QBluetoothHostInfo {
	return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothHostInfo {
		var out = make([]*QBluetoothHostInfo, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQBluetoothLocalDeviceFromPointer(l.data).allDevices_atList(i)
		}
		return out
	}(C.QBluetoothLocalDevice_QBluetoothLocalDevice_AllDevices())
}

func (ptr *QBluetoothLocalDevice) ConnectedDevices() []*QBluetoothAddress {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothAddress {
			var out = make([]*QBluetoothAddress, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothLocalDeviceFromPointer(l.data).connectedDevices_atList(i)
			}
			return out
		}(C.QBluetoothLocalDevice_ConnectedDevices(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__HostMode(C.QBluetoothLocalDevice_HostMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothLocalDevice_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQBluetoothLocalDevice_PairingConfirmation
func callbackQBluetoothLocalDevice_PairingConfirmation(ptr unsafe.Pointer, accept C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::pairingConfirmation"); signal != nil {
		signal.(func(bool))(int8(accept) != 0)
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingConfirmation(f func(accept bool)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::pairingConfirmation", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingConfirmation(accept bool) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::pairingConfirmation")
	}
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(accept bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmation(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(accept))))
	}
}

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddress_ITF) QBluetoothLocalDevice__Pairing {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__Pairing(C.QBluetoothLocalDevice_PairingStatus(ptr.Pointer(), PointerFromQBluetoothAddress(address)))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) PowerOn() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PowerOn(ptr.Pointer())
	}
}

func (ptr *QBluetoothLocalDevice) RequestPairing(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_RequestPairing(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.longlong(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) SetHostMode(mode QBluetoothLocalDevice__HostMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_SetHostMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QBluetoothLocalDevice) allDevices_atList(i int) *QBluetoothHostInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothHostInfoFromPointer(C.QBluetoothLocalDevice_allDevices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) connectedDevices_atList(i int) *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothLocalDevice_connectedDevices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothLocalDevice_TimerEvent
func callbackQBluetoothLocalDevice_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::timerEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::timerEvent")
	}
}

func (ptr *QBluetoothLocalDevice) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothLocalDevice_ChildEvent
func callbackQBluetoothLocalDevice_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::childEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::childEvent")
	}
}

func (ptr *QBluetoothLocalDevice) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothLocalDevice_ConnectNotify
func callbackQBluetoothLocalDevice_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::connectNotify", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::connectNotify")
	}
}

func (ptr *QBluetoothLocalDevice) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothLocalDevice_CustomEvent
func callbackQBluetoothLocalDevice_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::customEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::customEvent")
	}
}

func (ptr *QBluetoothLocalDevice) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothLocalDevice_DeleteLater
func callbackQBluetoothLocalDevice_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothLocalDevice) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::deleteLater", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::deleteLater")
	}
}

func (ptr *QBluetoothLocalDevice) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothLocalDevice) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothLocalDevice_DisconnectNotify
func callbackQBluetoothLocalDevice_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::disconnectNotify", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::disconnectNotify")
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothLocalDevice_Event
func callbackQBluetoothLocalDevice_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothLocalDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothLocalDevice) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::event", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::event")
	}
}

func (ptr *QBluetoothLocalDevice) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_EventFilter
func callbackQBluetoothLocalDevice_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothLocalDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothLocalDevice) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::eventFilter", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::eventFilter")
	}
}

func (ptr *QBluetoothLocalDevice) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_MetaObject
func callbackQBluetoothLocalDevice_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothLocalDevice::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothLocalDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothLocalDevice) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::metaObject", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothLocalDevice::metaObject")
	}
}

func (ptr *QBluetoothLocalDevice) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothLocalDevice_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothLocalDevice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothServer::Error
type QBluetoothServer__Error int64

const (
	QBluetoothServer__NoError                       = QBluetoothServer__Error(0)
	QBluetoothServer__UnknownError                  = QBluetoothServer__Error(1)
	QBluetoothServer__PoweredOffError               = QBluetoothServer__Error(2)
	QBluetoothServer__InputOutputError              = QBluetoothServer__Error(3)
	QBluetoothServer__ServiceAlreadyRegisteredError = QBluetoothServer__Error(4)
	QBluetoothServer__UnsupportedProtocolError      = QBluetoothServer__Error(5)
)

type QBluetoothServer struct {
	core.QObject
}

type QBluetoothServer_ITF interface {
	core.QObject_ITF
	QBluetoothServer_PTR() *QBluetoothServer
}

func (p *QBluetoothServer) QBluetoothServer_PTR() *QBluetoothServer {
	return p
}

func (p *QBluetoothServer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothServer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothServer(ptr QBluetoothServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServer_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServerFromPointer(ptr unsafe.Pointer) *QBluetoothServer {
	var n = new(QBluetoothServer)
	n.SetPointer(ptr)
	return n
}
func NewQBluetoothServer(serverType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothServer {
	var tmpValue = NewQBluetoothServerFromPointer(C.QBluetoothServer_NewQBluetoothServer(C.longlong(serverType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothServer_Error2
func callbackQBluetoothServer_Error2(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::error2"); signal != nil {
		signal.(func(QBluetoothServer__Error))(QBluetoothServer__Error(error))
	}

}

func (ptr *QBluetoothServer) ConnectError2(f func(error QBluetoothServer__Error)) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::error2", f)
	}
}

func (ptr *QBluetoothServer) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::error2")
	}
}

func (ptr *QBluetoothServer) Error2(error QBluetoothServer__Error) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_Error2(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothServer_NewConnection
func callbackQBluetoothServer_NewConnection(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::newConnection", f)
	}
}

func (ptr *QBluetoothServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::newConnection")
	}
}

func (ptr *QBluetoothServer) NewConnection() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_NewConnection(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServer__Error(C.QBluetoothServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) Listen2(uuid QBluetoothUuid_ITF, serviceName string) *QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServer_Listen2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid), serviceNameC))
		runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServer_ServerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DestroyQBluetoothServer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) Close() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) Listen(address QBluetoothAddress_ITF, port uint16) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_Listen(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.ushort(port)) != 0
	}
	return false
}

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothSocketFromPointer(C.QBluetoothServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) ServerAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothServer_ServerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) ServerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothServer_ServerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_SetMaxPendingConnections(ptr.Pointer(), C.int(int32(numConnections)))
	}
}

//export callbackQBluetoothServer_TimerEvent
func callbackQBluetoothServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::timerEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::timerEvent")
	}
}

func (ptr *QBluetoothServer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothServer_ChildEvent
func callbackQBluetoothServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::childEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::childEvent")
	}
}

func (ptr *QBluetoothServer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothServer_ConnectNotify
func callbackQBluetoothServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::connectNotify", f)
	}
}

func (ptr *QBluetoothServer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::connectNotify")
	}
}

func (ptr *QBluetoothServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServer_CustomEvent
func callbackQBluetoothServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::customEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::customEvent")
	}
}

func (ptr *QBluetoothServer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothServer_DeleteLater
func callbackQBluetoothServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothServer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::deleteLater", f)
	}
}

func (ptr *QBluetoothServer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::deleteLater")
	}
}

func (ptr *QBluetoothServer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothServer_DisconnectNotify
func callbackQBluetoothServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::disconnectNotify", f)
	}
}

func (ptr *QBluetoothServer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::disconnectNotify")
	}
}

func (ptr *QBluetoothServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServer_Event
func callbackQBluetoothServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothServer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::event", f)
	}
}

func (ptr *QBluetoothServer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::event")
	}
}

func (ptr *QBluetoothServer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothServer_EventFilter
func callbackQBluetoothServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::eventFilter", f)
	}
}

func (ptr *QBluetoothServer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::eventFilter")
	}
}

func (ptr *QBluetoothServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothServer_MetaObject
func callbackQBluetoothServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothServer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::metaObject", f)
	}
}

func (ptr *QBluetoothServer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServer::metaObject")
	}
}

func (ptr *QBluetoothServer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothServiceDiscoveryAgent::DiscoveryMode
type QBluetoothServiceDiscoveryAgent__DiscoveryMode int64

const (
	QBluetoothServiceDiscoveryAgent__MinimalDiscovery = QBluetoothServiceDiscoveryAgent__DiscoveryMode(0)
	QBluetoothServiceDiscoveryAgent__FullDiscovery    = QBluetoothServiceDiscoveryAgent__DiscoveryMode(1)
)

//QBluetoothServiceDiscoveryAgent::Error
type QBluetoothServiceDiscoveryAgent__Error int64

const (
	QBluetoothServiceDiscoveryAgent__NoError                      = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__NoError)
	QBluetoothServiceDiscoveryAgent__InputOutputError             = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InputOutputError)
	QBluetoothServiceDiscoveryAgent__PoweredOffError              = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__PoweredOffError)
	QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError)
	QBluetoothServiceDiscoveryAgent__UnknownError                 = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__UnknownError)
)

type QBluetoothServiceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothServiceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent
}

func (p *QBluetoothServiceDiscoveryAgent) QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent {
	return p
}

func (p *QBluetoothServiceDiscoveryAgent) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothServiceDiscoveryAgent) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothServiceDiscoveryAgent(ptr QBluetoothServiceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothServiceDiscoveryAgent {
	var n = new(QBluetoothServiceDiscoveryAgent)
	n.SetPointer(ptr)
	return n
}

//export callbackQBluetoothServiceDiscoveryAgent_Canceled
func callbackQBluetoothServiceDiscoveryAgent_Canceled(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCanceled(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::canceled", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::canceled")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Canceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Canceled(ptr.Pointer())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Error2
func callbackQBluetoothServiceDiscoveryAgent_Error2(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::error2"); signal != nil {
		signal.(func(QBluetoothServiceDiscoveryAgent__Error))(QBluetoothServiceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectError2(f func(error QBluetoothServiceDiscoveryAgent__Error)) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::error2", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::error2")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error2(error QBluetoothServiceDiscoveryAgent__Error) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Error2(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Finished
func callbackQBluetoothServiceDiscoveryAgent_Finished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::finished", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::finished")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Finished() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Finished(ptr.Pointer())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered
func callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered(ptr unsafe.Pointer, info unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::serviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothServiceInfo))(NewQBluetoothServiceInfoFromPointer(info))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectServiceDiscovered(f func(info *QBluetoothServiceInfo)) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectServiceDiscovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::serviceDiscovered", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectServiceDiscovered() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectServiceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::serviceDiscovered")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ServiceDiscovered(info QBluetoothServiceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ServiceDiscovered(ptr.Pointer(), PointerFromQBluetoothServiceInfo(info))
	}
}

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	var tmpValue = NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	var tmpValue = NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothServiceDiscoveryAgent_Clear
func callbackQBluetoothServiceDiscoveryAgent_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::clear", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::clear")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DiscoveredServices() []*QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothServiceInfo {
			var out = make([]*QBluetoothServiceInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothServiceDiscoveryAgentFromPointer(l.data).discoveredServices_atList(i)
			}
			return out
		}(C.QBluetoothServiceDiscoveryAgent_DiscoveredServices(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) RemoteAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothServiceDiscoveryAgent_RemoteAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Start
func callbackQBluetoothServiceDiscoveryAgent_Start(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::start"); signal != nil {
		signal.(func(QBluetoothServiceDiscoveryAgent__DiscoveryMode))(QBluetoothServiceDiscoveryAgent__DiscoveryMode(mode))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStart(f func(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::start", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectStart(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::start")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Start(ptr.Pointer(), C.longlong(mode))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Stop
func callbackQBluetoothServiceDiscoveryAgent_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::stop", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::stop")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) UuidFilter() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothServiceDiscoveryAgentFromPointer(l.data).uuidFilter_atList(i)
			}
			return out
		}(C.QBluetoothServiceDiscoveryAgent_UuidFilter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) discoveredServices_atList(i int) *QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceDiscoveryAgent_discoveredServices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) uuidFilter_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothServiceDiscoveryAgent_uuidFilter_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothServiceDiscoveryAgent_TimerEvent
func callbackQBluetoothServiceDiscoveryAgent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::timerEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::timerEvent")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ChildEvent
func callbackQBluetoothServiceDiscoveryAgent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::childEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::childEvent")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ConnectNotify
func callbackQBluetoothServiceDiscoveryAgent_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::connectNotify", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::connectNotify")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_CustomEvent
func callbackQBluetoothServiceDiscoveryAgent_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::customEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::customEvent")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_DeleteLater
func callbackQBluetoothServiceDiscoveryAgent_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::deleteLater", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::deleteLater")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify
func callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::disconnectNotify", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::disconnectNotify")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Event
func callbackQBluetoothServiceDiscoveryAgent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::event", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::event")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_EventFilter
func callbackQBluetoothServiceDiscoveryAgent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::eventFilter", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::eventFilter")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_MetaObject
func callbackQBluetoothServiceDiscoveryAgent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothServiceDiscoveryAgent::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::metaObject", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothServiceDiscoveryAgent::metaObject")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServiceDiscoveryAgent_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServiceDiscoveryAgent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothServiceInfo::AttributeId
type QBluetoothServiceInfo__AttributeId int64

var (
	QBluetoothServiceInfo__ServiceRecordHandle              = QBluetoothServiceInfo__AttributeId(0x0000)
	QBluetoothServiceInfo__ServiceClassIds                  = QBluetoothServiceInfo__AttributeId(0x0001)
	QBluetoothServiceInfo__ServiceRecordState               = QBluetoothServiceInfo__AttributeId(0x0002)
	QBluetoothServiceInfo__ServiceId                        = QBluetoothServiceInfo__AttributeId(0x0003)
	QBluetoothServiceInfo__ProtocolDescriptorList           = QBluetoothServiceInfo__AttributeId(0x0004)
	QBluetoothServiceInfo__BrowseGroupList                  = QBluetoothServiceInfo__AttributeId(0x0005)
	QBluetoothServiceInfo__LanguageBaseAttributeIdList      = QBluetoothServiceInfo__AttributeId(0x0006)
	QBluetoothServiceInfo__ServiceInfoTimeToLive            = QBluetoothServiceInfo__AttributeId(0x0007)
	QBluetoothServiceInfo__ServiceAvailability              = QBluetoothServiceInfo__AttributeId(0x0008)
	QBluetoothServiceInfo__BluetoothProfileDescriptorList   = QBluetoothServiceInfo__AttributeId(0x0009)
	QBluetoothServiceInfo__DocumentationUrl                 = QBluetoothServiceInfo__AttributeId(0x000A)
	QBluetoothServiceInfo__ClientExecutableUrl              = QBluetoothServiceInfo__AttributeId(0x000B)
	QBluetoothServiceInfo__IconUrl                          = QBluetoothServiceInfo__AttributeId(0x000C)
	QBluetoothServiceInfo__AdditionalProtocolDescriptorList = QBluetoothServiceInfo__AttributeId(0x000D)
	QBluetoothServiceInfo__PrimaryLanguageBase              = QBluetoothServiceInfo__AttributeId(0x0100)
	QBluetoothServiceInfo__ServiceName                      = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceName_Type())
	QBluetoothServiceInfo__ServiceDescription               = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceDescription_Type())
	QBluetoothServiceInfo__ServiceProvider                  = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceProvider_Type())
)

//QBluetoothServiceInfo::Protocol
type QBluetoothServiceInfo__Protocol int64

const (
	QBluetoothServiceInfo__UnknownProtocol = QBluetoothServiceInfo__Protocol(0)
	QBluetoothServiceInfo__L2capProtocol   = QBluetoothServiceInfo__Protocol(1)
	QBluetoothServiceInfo__RfcommProtocol  = QBluetoothServiceInfo__Protocol(2)
)

type QBluetoothServiceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothServiceInfo_ITF interface {
	QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo
}

func (p *QBluetoothServiceInfo) QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo {
	return p
}

func (p *QBluetoothServiceInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothServiceInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothServiceInfo(ptr QBluetoothServiceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServiceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothServiceInfo {
	var n = new(QBluetoothServiceInfo)
	n.SetPointer(ptr)
	return n
}
func (ptr *QBluetoothServiceInfo) ServiceAvailability() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceInfo_ServiceAvailability(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceDescription() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceInfo_ServiceDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceInfo_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceProvider() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceInfo_ServiceProvider(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothServiceInfo_ServiceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) SetAttribute2(attributeId uint16, value QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetAttribute2(ptr.Pointer(), C.ushort(attributeId), PointerFromQBluetoothUuid(value))
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceAvailability(availability string) {
	if ptr.Pointer() != nil {
		var availabilityC = C.CString(availability)
		defer C.free(unsafe.Pointer(availabilityC))
		C.QBluetoothServiceInfo_SetServiceAvailability(ptr.Pointer(), availabilityC)
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC = C.CString(description)
		defer C.free(unsafe.Pointer(descriptionC))
		C.QBluetoothServiceInfo_SetServiceDescription(ptr.Pointer(), descriptionC)
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QBluetoothServiceInfo_SetServiceName(ptr.Pointer(), nameC)
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceProvider(provider string) {
	if ptr.Pointer() != nil {
		var providerC = C.CString(provider)
		defer C.free(unsafe.Pointer(providerC))
		C.QBluetoothServiceInfo_SetServiceProvider(ptr.Pointer(), providerC)
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func NewQBluetoothServiceInfo() *QBluetoothServiceInfo {
	var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo())
	runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
	return tmpValue
}

func NewQBluetoothServiceInfo2(other QBluetoothServiceInfo_ITF) *QBluetoothServiceInfo {
	var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo2(PointerFromQBluetoothServiceInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
	return tmpValue
}

func (ptr *QBluetoothServiceInfo) Attribute(attributeId uint16) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QBluetoothServiceInfo_Attribute(ptr.Pointer(), C.ushort(attributeId)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) Contains(attributeId uint16) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_Contains(ptr.Pointer(), C.ushort(attributeId)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) Device() *QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothServiceInfo_Device(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) IsComplete() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsRegistered() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsRegistered(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ProtocolServiceMultiplexer() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothServiceInfo_ProtocolServiceMultiplexer(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) RegisterService(localAdapter QBluetoothAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_RegisterService(ptr.Pointer(), PointerFromQBluetoothAddress(localAdapter)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) RemoveAttribute(attributeId uint16) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_RemoveAttribute(ptr.Pointer(), C.ushort(attributeId))
	}
}

func (ptr *QBluetoothServiceInfo) ServerChannel() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothServiceInfo_ServerChannel(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) ServiceClassUuids() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothServiceInfoFromPointer(l.data).serviceClassUuids_atList(i)
			}
			return out
		}(C.QBluetoothServiceInfo_ServiceClassUuids(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) SetAttribute(attributeId uint16, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetAttribute(ptr.Pointer(), C.ushort(attributeId), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothServiceInfo) SetDevice(device QBluetoothDeviceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetDevice(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(device))
	}
}

func (ptr *QBluetoothServiceInfo) SocketProtocol() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServiceInfo_SocketProtocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) UnregisterService() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_UnregisterService(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) DestroyQBluetoothServiceInfo() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceInfo) serviceClassUuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothServiceInfo_serviceClassUuids_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

//QBluetoothSocket::SocketError
type QBluetoothSocket__SocketError int64

const (
	QBluetoothSocket__NoSocketError            = QBluetoothSocket__SocketError(-2)
	QBluetoothSocket__UnknownSocketError       = QBluetoothSocket__SocketError(network.QAbstractSocket__UnknownSocketError)
	QBluetoothSocket__HostNotFoundError        = QBluetoothSocket__SocketError(network.QAbstractSocket__HostNotFoundError)
	QBluetoothSocket__ServiceNotFoundError     = QBluetoothSocket__SocketError(network.QAbstractSocket__SocketAddressNotAvailableError)
	QBluetoothSocket__NetworkError             = QBluetoothSocket__SocketError(network.QAbstractSocket__NetworkError)
	QBluetoothSocket__UnsupportedProtocolError = QBluetoothSocket__SocketError(8)
	QBluetoothSocket__OperationError           = QBluetoothSocket__SocketError(network.QAbstractSocket__OperationError)
)

//QBluetoothSocket::SocketState
type QBluetoothSocket__SocketState int64

const (
	QBluetoothSocket__UnconnectedState   = QBluetoothSocket__SocketState(network.QAbstractSocket__UnconnectedState)
	QBluetoothSocket__ServiceLookupState = QBluetoothSocket__SocketState(network.QAbstractSocket__HostLookupState)
	QBluetoothSocket__ConnectingState    = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectingState)
	QBluetoothSocket__ConnectedState     = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectedState)
	QBluetoothSocket__BoundState         = QBluetoothSocket__SocketState(network.QAbstractSocket__BoundState)
	QBluetoothSocket__ClosingState       = QBluetoothSocket__SocketState(network.QAbstractSocket__ClosingState)
	QBluetoothSocket__ListeningState     = QBluetoothSocket__SocketState(network.QAbstractSocket__ListeningState)
)

type QBluetoothSocket struct {
	core.QIODevice
}

type QBluetoothSocket_ITF interface {
	core.QIODevice_ITF
	QBluetoothSocket_PTR() *QBluetoothSocket
}

func (p *QBluetoothSocket) QBluetoothSocket_PTR() *QBluetoothSocket {
	return p
}

func (p *QBluetoothSocket) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QIODevice_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothSocket) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QIODevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothSocket(ptr QBluetoothSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothSocket_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothSocketFromPointer(ptr unsafe.Pointer) *QBluetoothSocket {
	var n = new(QBluetoothSocket)
	n.SetPointer(ptr)
	return n
}

//export callbackQBluetoothSocket_Connected
func callbackQBluetoothSocket_Connected(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::connected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::connected")
	}
}

func (ptr *QBluetoothSocket) Connected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Connected(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Disconnected
func callbackQBluetoothSocket_Disconnected(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::disconnected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::disconnected")
	}
}

func (ptr *QBluetoothSocket) Disconnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Error2
func callbackQBluetoothSocket_Error2(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::error2"); signal != nil {
		signal.(func(QBluetoothSocket__SocketError))(QBluetoothSocket__SocketError(error))
	}

}

func (ptr *QBluetoothSocket) ConnectError2(f func(error QBluetoothSocket__SocketError)) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::error2", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::error2")
	}
}

func (ptr *QBluetoothSocket) Error2(error QBluetoothSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Error2(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothSocket_StateChanged
func callbackQBluetoothSocket_StateChanged(ptr unsafe.Pointer, state C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::stateChanged"); signal != nil {
		signal.(func(QBluetoothSocket__SocketState))(QBluetoothSocket__SocketState(state))
	}

}

func (ptr *QBluetoothSocket) ConnectStateChanged(f func(state QBluetoothSocket__SocketState)) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::stateChanged", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::stateChanged")
	}
}

func (ptr *QBluetoothSocket) StateChanged(state QBluetoothSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

func NewQBluetoothSocket(socketType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothSocket {
	var tmpValue = NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket(C.longlong(socketType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothSocket2(parent core.QObject_ITF) *QBluetoothSocket {
	var tmpValue = NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Abort(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_BytesAvailable
func callbackQBluetoothSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QBluetoothSocket) ConnectBytesAvailable(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::bytesAvailable", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectBytesAvailable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::bytesAvailable")
	}
}

func (ptr *QBluetoothSocket) BytesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_BytesToWrite
func callbackQBluetoothSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QBluetoothSocket) ConnectBytesToWrite(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::bytesToWrite", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectBytesToWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::bytesToWrite")
	}
}

func (ptr *QBluetoothSocket) BytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_CanReadLine
func callbackQBluetoothSocket_CanReadLine(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QBluetoothSocket) ConnectCanReadLine(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::canReadLine", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectCanReadLine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::canReadLine")
	}
}

func (ptr *QBluetoothSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Close
func callbackQBluetoothSocket_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::close"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QBluetoothSocket) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::close", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::close")
	}
}

func (ptr *QBluetoothSocket) Close() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) ConnectToService2(address QBluetoothAddress_ITF, uuid QBluetoothUuid_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService2(ptr.Pointer(), PointerFromQBluetoothAddress(address), PointerFromQBluetoothUuid(uuid), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) ConnectToService3(address QBluetoothAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService3(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) ConnectToService(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) DisconnectFromService() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectFromService(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) DoDeviceDiscovery(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DoDeviceDiscovery(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) Error() QBluetoothSocket__SocketError {
	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketError(C.QBluetoothSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) LocalAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothSocket_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) LocalName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothSocket_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) LocalPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothSocket_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) PeerAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothSocket_PeerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) PeerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothSocket_PeerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) SetSocketDescriptor(socketDescriptor int, socketType QBluetoothServiceInfo__Protocol, socketState QBluetoothSocket__SocketState, openMode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SetSocketDescriptor(ptr.Pointer(), C.int(int32(socketDescriptor)), C.longlong(socketType), C.longlong(socketState), C.longlong(openMode)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) SetSocketError(error_ QBluetoothSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_SetSocketError(ptr.Pointer(), C.longlong(error_))
	}
}

func (ptr *QBluetoothSocket) SetSocketState(state QBluetoothSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_SetSocketState(ptr.Pointer(), C.longlong(state))
	}
}

func (ptr *QBluetoothSocket) SocketDescriptor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothSocket_SocketDescriptor(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothSocket) SocketType() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) State() QBluetoothSocket__SocketState {
	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketState(C.QBluetoothSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_WriteData
func callbackQBluetoothSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtBluetooth_PackedString, maxSize C.longlong) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).WriteDataDefault(cGoUnpackString(data), int64(maxSize)))
}

func (ptr *QBluetoothSocket) ConnectWriteData(f func(data string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::writeData", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::writeData")
	}
}

func (ptr *QBluetoothSocket) WriteData(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QBluetoothSocket_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) WriteDataDefault(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QBluetoothSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQBluetoothSocket_DestroyQBluetoothSocket
func callbackQBluetoothSocket_DestroyQBluetoothSocket(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::~QBluetoothSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).DestroyQBluetoothSocketDefault()
	}
}

func (ptr *QBluetoothSocket) ConnectDestroyQBluetoothSocket(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::~QBluetoothSocket", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDestroyQBluetoothSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::~QBluetoothSocket")
	}
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocket() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DestroyQBluetoothSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocketDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DestroyQBluetoothSocketDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothSocket_AtEnd
func callbackQBluetoothSocket_AtEnd(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QBluetoothSocket) ConnectAtEnd(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::atEnd", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectAtEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::atEnd")
	}
}

func (ptr *QBluetoothSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Open
func callbackQBluetoothSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QBluetoothSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::open", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::open")
	}
}

func (ptr *QBluetoothSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Pos
func callbackQBluetoothSocket_Pos(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).PosDefault())
}

func (ptr *QBluetoothSocket) ConnectPos(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::pos", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectPos() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::pos")
	}
}

func (ptr *QBluetoothSocket) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_ReadLineData
func callbackQBluetoothSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtBluetooth_PackedString, maxSize C.longlong) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxSize)))
}

func (ptr *QBluetoothSocket) ConnectReadLineData(f func(data string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::readLineData", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectReadLineData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::readLineData")
	}
}

func (ptr *QBluetoothSocket) ReadLineData(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QBluetoothSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) ReadLineDataDefault(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QBluetoothSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQBluetoothSocket_Reset
func callbackQBluetoothSocket_Reset(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QBluetoothSocket) ConnectReset(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::reset", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::reset")
	}
}

func (ptr *QBluetoothSocket) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Seek
func callbackQBluetoothSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QBluetoothSocket) ConnectSeek(f func(pos int64) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::seek", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectSeek() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::seek")
	}
}

func (ptr *QBluetoothSocket) Seek(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Size
func callbackQBluetoothSocket_Size(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QBluetoothSocket) ConnectSize(f func() int64) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::size", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::size")
	}
}

func (ptr *QBluetoothSocket) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_WaitForBytesWritten
func callbackQBluetoothSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QBluetoothSocket) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::waitForBytesWritten", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectWaitForBytesWritten() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::waitForBytesWritten")
	}
}

func (ptr *QBluetoothSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_WaitForReadyRead
func callbackQBluetoothSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QBluetoothSocket) ConnectWaitForReadyRead(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::waitForReadyRead", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectWaitForReadyRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::waitForReadyRead")
	}
}

func (ptr *QBluetoothSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_TimerEvent
func callbackQBluetoothSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::timerEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::timerEvent")
	}
}

func (ptr *QBluetoothSocket) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothSocket_ChildEvent
func callbackQBluetoothSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::childEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::childEvent")
	}
}

func (ptr *QBluetoothSocket) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothSocket_ConnectNotify
func callbackQBluetoothSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::connectNotify", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::connectNotify")
	}
}

func (ptr *QBluetoothSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothSocket_CustomEvent
func callbackQBluetoothSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::customEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::customEvent")
	}
}

func (ptr *QBluetoothSocket) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothSocket_DeleteLater
func callbackQBluetoothSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothSocket) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::deleteLater", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::deleteLater")
	}
}

func (ptr *QBluetoothSocket) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothSocket_DisconnectNotify
func callbackQBluetoothSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::disconnectNotify", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::disconnectNotify")
	}
}

func (ptr *QBluetoothSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothSocket_Event
func callbackQBluetoothSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::event", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::event")
	}
}

func (ptr *QBluetoothSocket) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_EventFilter
func callbackQBluetoothSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::eventFilter", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::eventFilter")
	}
}

func (ptr *QBluetoothSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_MetaObject
func callbackQBluetoothSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::metaObject", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothSocket::metaObject")
	}
}

func (ptr *QBluetoothSocket) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QBluetoothTransferManager struct {
	core.QObject
}

type QBluetoothTransferManager_ITF interface {
	core.QObject_ITF
	QBluetoothTransferManager_PTR() *QBluetoothTransferManager
}

func (p *QBluetoothTransferManager) QBluetoothTransferManager_PTR() *QBluetoothTransferManager {
	return p
}

func (p *QBluetoothTransferManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothTransferManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothTransferManager(ptr QBluetoothTransferManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferManager_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) *QBluetoothTransferManager {
	var n = new(QBluetoothTransferManager)
	n.SetPointer(ptr)
	return n
}
func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequest_ITF, data core.QIODevice_ITF) *QBluetoothTransferReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferManager_Put(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQBluetoothTransferManager(parent core.QObject_ITF) *QBluetoothTransferManager {
	var tmpValue = NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferManager_NewQBluetoothTransferManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothTransferManager_Finished
func callbackQBluetoothTransferManager_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::finished", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::finished")
	}
}

func (ptr *QBluetoothTransferManager) Finished(reply QBluetoothTransferReply_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothTransferManager_TimerEvent
func callbackQBluetoothTransferManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::timerEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::timerEvent")
	}
}

func (ptr *QBluetoothTransferManager) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothTransferManager_ChildEvent
func callbackQBluetoothTransferManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::childEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::childEvent")
	}
}

func (ptr *QBluetoothTransferManager) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothTransferManager_ConnectNotify
func callbackQBluetoothTransferManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::connectNotify", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::connectNotify")
	}
}

func (ptr *QBluetoothTransferManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothTransferManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferManager_CustomEvent
func callbackQBluetoothTransferManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::customEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::customEvent")
	}
}

func (ptr *QBluetoothTransferManager) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothTransferManager_DeleteLater
func callbackQBluetoothTransferManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothTransferManager) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::deleteLater", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::deleteLater")
	}
}

func (ptr *QBluetoothTransferManager) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothTransferManager_DisconnectNotify
func callbackQBluetoothTransferManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::disconnectNotify", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::disconnectNotify")
	}
}

func (ptr *QBluetoothTransferManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothTransferManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferManager_Event
func callbackQBluetoothTransferManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothTransferManager) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::event", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::event")
	}
}

func (ptr *QBluetoothTransferManager) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothTransferManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothTransferManager_EventFilter
func callbackQBluetoothTransferManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothTransferManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::eventFilter", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::eventFilter")
	}
}

func (ptr *QBluetoothTransferManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothTransferManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothTransferManager_MetaObject
func callbackQBluetoothTransferManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothTransferManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothTransferManager) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::metaObject", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferManager::metaObject")
	}
}

func (ptr *QBluetoothTransferManager) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothTransferReply::TransferError
type QBluetoothTransferReply__TransferError int64

const (
	QBluetoothTransferReply__NoError                   = QBluetoothTransferReply__TransferError(0)
	QBluetoothTransferReply__UnknownError              = QBluetoothTransferReply__TransferError(1)
	QBluetoothTransferReply__FileNotFoundError         = QBluetoothTransferReply__TransferError(2)
	QBluetoothTransferReply__HostNotFoundError         = QBluetoothTransferReply__TransferError(3)
	QBluetoothTransferReply__UserCanceledTransferError = QBluetoothTransferReply__TransferError(4)
	QBluetoothTransferReply__IODeviceNotReadableError  = QBluetoothTransferReply__TransferError(5)
	QBluetoothTransferReply__ResourceBusyError         = QBluetoothTransferReply__TransferError(6)
	QBluetoothTransferReply__SessionError              = QBluetoothTransferReply__TransferError(7)
)

type QBluetoothTransferReply struct {
	core.QObject
}

type QBluetoothTransferReply_ITF interface {
	core.QObject_ITF
	QBluetoothTransferReply_PTR() *QBluetoothTransferReply
}

func (p *QBluetoothTransferReply) QBluetoothTransferReply_PTR() *QBluetoothTransferReply {
	return p
}

func (p *QBluetoothTransferReply) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothTransferReply) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothTransferReply(ptr QBluetoothTransferReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferReply_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) *QBluetoothTransferReply {
	var n = new(QBluetoothTransferReply)
	n.SetPointer(ptr)
	return n
}

//export callbackQBluetoothTransferReply_Abort
func callbackQBluetoothTransferReply_Abort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::abort"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothTransferReply) ConnectAbort(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::abort", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectAbort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::abort")
	}
}

func (ptr *QBluetoothTransferReply) Abort() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothTransferReply) SetManager(manager QBluetoothTransferManager_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_SetManager(ptr.Pointer(), PointerFromQBluetoothTransferManager(manager))
	}
}

func (ptr *QBluetoothTransferReply) SetRequest(request QBluetoothTransferRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_SetRequest(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request))
	}
}

func NewQBluetoothTransferReply(parent core.QObject_ITF) *QBluetoothTransferReply {
	var tmpValue = NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferReply_NewQBluetoothTransferReply(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothTransferReply_Error2
func callbackQBluetoothTransferReply_Error2(ptr unsafe.Pointer, errorType C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::error2"); signal != nil {
		signal.(func(QBluetoothTransferReply__TransferError))(QBluetoothTransferReply__TransferError(errorType))
	}

}

func (ptr *QBluetoothTransferReply) ConnectError2(f func(errorType QBluetoothTransferReply__TransferError)) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::error2", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::error2")
	}
}

func (ptr *QBluetoothTransferReply) Error2(errorType QBluetoothTransferReply__TransferError) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Error2(ptr.Pointer(), C.longlong(errorType))
	}
}

//export callbackQBluetoothTransferReply_Error
func callbackQBluetoothTransferReply_Error(ptr unsafe.Pointer) C.longlong {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::error"); signal != nil {
		return C.longlong(signal.(func() QBluetoothTransferReply__TransferError)())
	}

	return C.longlong(0)
}

func (ptr *QBluetoothTransferReply) ConnectError(f func() QBluetoothTransferReply__TransferError) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::error", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::error")
	}
}

func (ptr *QBluetoothTransferReply) Error() QBluetoothTransferReply__TransferError {
	if ptr.Pointer() != nil {
		return QBluetoothTransferReply__TransferError(C.QBluetoothTransferReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothTransferReply_ErrorString
func callbackQBluetoothTransferReply_ErrorString(ptr unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QBluetoothTransferReply) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::errorString", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::errorString")
	}
}

func (ptr *QBluetoothTransferReply) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothTransferReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQBluetoothTransferReply_Finished
func callbackQBluetoothTransferReply_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::finished", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::finished")
	}
}

func (ptr *QBluetoothTransferReply) Finished(reply QBluetoothTransferReply_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

//export callbackQBluetoothTransferReply_IsFinished
func callbackQBluetoothTransferReply_IsFinished(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::isFinished"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QBluetoothTransferReply) ConnectIsFinished(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::isFinished", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectIsFinished() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::isFinished")
	}
}

func (ptr *QBluetoothTransferReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_IsRunning
func callbackQBluetoothTransferReply_IsRunning(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::isRunning"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QBluetoothTransferReply) ConnectIsRunning(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::isRunning", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectIsRunning() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::isRunning")
	}
}

func (ptr *QBluetoothTransferReply) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferReply_Manager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferReply) Request() *QBluetoothTransferRequest {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferReply_Request(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothTransferReply_TransferProgress
func callbackQBluetoothTransferReply_TransferProgress(ptr unsafe.Pointer, bytesTransferred C.longlong, bytesTotal C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::transferProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesTransferred), int64(bytesTotal))
	}

}

func (ptr *QBluetoothTransferReply) ConnectTransferProgress(f func(bytesTransferred int64, bytesTotal int64)) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectTransferProgress(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::transferProgress", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTransferProgress() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectTransferProgress(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::transferProgress")
	}
}

func (ptr *QBluetoothTransferReply) TransferProgress(bytesTransferred int64, bytesTotal int64) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TransferProgress(ptr.Pointer(), C.longlong(bytesTransferred), C.longlong(bytesTotal))
	}
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothTransferReply_TimerEvent
func callbackQBluetoothTransferReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::timerEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::timerEvent")
	}
}

func (ptr *QBluetoothTransferReply) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothTransferReply_ChildEvent
func callbackQBluetoothTransferReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::childEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::childEvent")
	}
}

func (ptr *QBluetoothTransferReply) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothTransferReply_ConnectNotify
func callbackQBluetoothTransferReply_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferReply) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::connectNotify", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::connectNotify")
	}
}

func (ptr *QBluetoothTransferReply) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothTransferReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferReply_CustomEvent
func callbackQBluetoothTransferReply_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::customEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::customEvent")
	}
}

func (ptr *QBluetoothTransferReply) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothTransferReply_DeleteLater
func callbackQBluetoothTransferReply_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothTransferReply) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::deleteLater", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::deleteLater")
	}
}

func (ptr *QBluetoothTransferReply) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferReply) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothTransferReply_DisconnectNotify
func callbackQBluetoothTransferReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferReply) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::disconnectNotify", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::disconnectNotify")
	}
}

func (ptr *QBluetoothTransferReply) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothTransferReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferReply_Event
func callbackQBluetoothTransferReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothTransferReply) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::event", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::event")
	}
}

func (ptr *QBluetoothTransferReply) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_EventFilter
func callbackQBluetoothTransferReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothTransferReply) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::eventFilter", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::eventFilter")
	}
}

func (ptr *QBluetoothTransferReply) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_MetaObject
func callbackQBluetoothTransferReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QBluetoothTransferReply::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothTransferReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothTransferReply) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::metaObject", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QBluetoothTransferReply::metaObject")
	}
}

func (ptr *QBluetoothTransferReply) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferReply_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferReply) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothTransferRequest::Attribute
type QBluetoothTransferRequest__Attribute int64

const (
	QBluetoothTransferRequest__DescriptionAttribute = QBluetoothTransferRequest__Attribute(0)
	QBluetoothTransferRequest__TimeAttribute        = QBluetoothTransferRequest__Attribute(1)
	QBluetoothTransferRequest__TypeAttribute        = QBluetoothTransferRequest__Attribute(2)
	QBluetoothTransferRequest__LengthAttribute      = QBluetoothTransferRequest__Attribute(3)
	QBluetoothTransferRequest__NameAttribute        = QBluetoothTransferRequest__Attribute(4)
)

type QBluetoothTransferRequest struct {
	ptr unsafe.Pointer
}

type QBluetoothTransferRequest_ITF interface {
	QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest
}

func (p *QBluetoothTransferRequest) QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest {
	return p
}

func (p *QBluetoothTransferRequest) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothTransferRequest) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothTransferRequest(ptr QBluetoothTransferRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferRequest_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) *QBluetoothTransferRequest {
	var n = new(QBluetoothTransferRequest)
	n.SetPointer(ptr)
	return n
}
func NewQBluetoothTransferRequest(address QBluetoothAddress_ITF) *QBluetoothTransferRequest {
	var tmpValue = NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest(PointerFromQBluetoothAddress(address)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
	return tmpValue
}

func NewQBluetoothTransferRequest2(other QBluetoothTransferRequest_ITF) *QBluetoothTransferRequest {
	var tmpValue = NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest2(PointerFromQBluetoothTransferRequest(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
	return tmpValue
}

func (ptr *QBluetoothTransferRequest) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothTransferRequest_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) Attribute(code QBluetoothTransferRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QBluetoothTransferRequest_Attribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(defaultValue)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) SetAttribute(code QBluetoothTransferRequest__Attribute, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_SetAttribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothTransferRequest) DestroyQBluetoothTransferRequest() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QBluetoothUuid::CharacteristicType
type QBluetoothUuid__CharacteristicType int64

const (
	QBluetoothUuid__DeviceName                                    = QBluetoothUuid__CharacteristicType(0x2a00)
	QBluetoothUuid__Appearance                                    = QBluetoothUuid__CharacteristicType(0x2a01)
	QBluetoothUuid__PeripheralPrivacyFlag                         = QBluetoothUuid__CharacteristicType(0x2a02)
	QBluetoothUuid__ReconnectionAddress                           = QBluetoothUuid__CharacteristicType(0x2a03)
	QBluetoothUuid__PeripheralPreferredConnectionParameters       = QBluetoothUuid__CharacteristicType(0x2a04)
	QBluetoothUuid__ServiceChanged                                = QBluetoothUuid__CharacteristicType(0x2a05)
	QBluetoothUuid__AlertLevel                                    = QBluetoothUuid__CharacteristicType(0x2a06)
	QBluetoothUuid__TxPowerLevel                                  = QBluetoothUuid__CharacteristicType(0x2a07)
	QBluetoothUuid__DateTime                                      = QBluetoothUuid__CharacteristicType(0x2a08)
	QBluetoothUuid__DayOfWeek                                     = QBluetoothUuid__CharacteristicType(0x2a09)
	QBluetoothUuid__DayDateTime                                   = QBluetoothUuid__CharacteristicType(0x2a0a)
	QBluetoothUuid__ExactTime256                                  = QBluetoothUuid__CharacteristicType(0x2a0c)
	QBluetoothUuid__DSTOffset                                     = QBluetoothUuid__CharacteristicType(0x2a0d)
	QBluetoothUuid__TimeZone                                      = QBluetoothUuid__CharacteristicType(0x2a0e)
	QBluetoothUuid__LocalTimeInformation                          = QBluetoothUuid__CharacteristicType(0x2a0f)
	QBluetoothUuid__TimeWithDST                                   = QBluetoothUuid__CharacteristicType(0x2a11)
	QBluetoothUuid__TimeAccuracy                                  = QBluetoothUuid__CharacteristicType(0x2a12)
	QBluetoothUuid__TimeSource                                    = QBluetoothUuid__CharacteristicType(0x2a13)
	QBluetoothUuid__ReferenceTimeInformation                      = QBluetoothUuid__CharacteristicType(0x2a14)
	QBluetoothUuid__TimeUpdateControlPoint                        = QBluetoothUuid__CharacteristicType(0x2a16)
	QBluetoothUuid__TimeUpdateState                               = QBluetoothUuid__CharacteristicType(0x2a17)
	QBluetoothUuid__GlucoseMeasurement                            = QBluetoothUuid__CharacteristicType(0x2a18)
	QBluetoothUuid__BatteryLevel                                  = QBluetoothUuid__CharacteristicType(0x2a19)
	QBluetoothUuid__TemperatureMeasurement                        = QBluetoothUuid__CharacteristicType(0x2a1c)
	QBluetoothUuid__TemperatureType                               = QBluetoothUuid__CharacteristicType(0x2a1d)
	QBluetoothUuid__IntermediateTemperature                       = QBluetoothUuid__CharacteristicType(0x2a1e)
	QBluetoothUuid__MeasurementInterval                           = QBluetoothUuid__CharacteristicType(0x2a21)
	QBluetoothUuid__BootKeyboardInputReport                       = QBluetoothUuid__CharacteristicType(0x2a22)
	QBluetoothUuid__SystemID                                      = QBluetoothUuid__CharacteristicType(0x2a23)
	QBluetoothUuid__ModelNumberString                             = QBluetoothUuid__CharacteristicType(0x2a24)
	QBluetoothUuid__SerialNumberString                            = QBluetoothUuid__CharacteristicType(0x2a25)
	QBluetoothUuid__FirmwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a26)
	QBluetoothUuid__HardwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a27)
	QBluetoothUuid__SoftwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a28)
	QBluetoothUuid__ManufacturerNameString                        = QBluetoothUuid__CharacteristicType(0x2a29)
	QBluetoothUuid__IEEE1107320601RegulatoryCertificationDataList = QBluetoothUuid__CharacteristicType(0x2a2a)
	QBluetoothUuid__CurrentTime                                   = QBluetoothUuid__CharacteristicType(0x2a2b)
	QBluetoothUuid__MagneticDeclination                           = QBluetoothUuid__CharacteristicType(0x2a2c)
	QBluetoothUuid__ScanRefresh                                   = QBluetoothUuid__CharacteristicType(0x2a31)
	QBluetoothUuid__BootKeyboardOutputReport                      = QBluetoothUuid__CharacteristicType(0x2a32)
	QBluetoothUuid__BootMouseInputReport                          = QBluetoothUuid__CharacteristicType(0x2a33)
	QBluetoothUuid__GlucoseMeasurementContext                     = QBluetoothUuid__CharacteristicType(0x2a34)
	QBluetoothUuid__BloodPressureMeasurement                      = QBluetoothUuid__CharacteristicType(0x2a35)
	QBluetoothUuid__IntermediateCuffPressure                      = QBluetoothUuid__CharacteristicType(0x2a36)
	QBluetoothUuid__HeartRateMeasurement                          = QBluetoothUuid__CharacteristicType(0x2a37)
	QBluetoothUuid__BodySensorLocation                            = QBluetoothUuid__CharacteristicType(0x2a38)
	QBluetoothUuid__HeartRateControlPoint                         = QBluetoothUuid__CharacteristicType(0x2a39)
	QBluetoothUuid__AlertStatus                                   = QBluetoothUuid__CharacteristicType(0x2a3f)
	QBluetoothUuid__RingerControlPoint                            = QBluetoothUuid__CharacteristicType(0x2a40)
	QBluetoothUuid__RingerSetting                                 = QBluetoothUuid__CharacteristicType(0x2a41)
	QBluetoothUuid__AlertCategoryIDBitMask                        = QBluetoothUuid__CharacteristicType(0x2a42)
	QBluetoothUuid__AlertCategoryID                               = QBluetoothUuid__CharacteristicType(0x2a43)
	QBluetoothUuid__AlertNotificationControlPoint                 = QBluetoothUuid__CharacteristicType(0x2a44)
	QBluetoothUuid__UnreadAlertStatus                             = QBluetoothUuid__CharacteristicType(0x2a45)
	QBluetoothUuid__NewAlert                                      = QBluetoothUuid__CharacteristicType(0x2a46)
	QBluetoothUuid__SupportedNewAlertCategory                     = QBluetoothUuid__CharacteristicType(0x2a47)
	QBluetoothUuid__SupportedUnreadAlertCategory                  = QBluetoothUuid__CharacteristicType(0x2a48)
	QBluetoothUuid__BloodPressureFeature                          = QBluetoothUuid__CharacteristicType(0x2a49)
	QBluetoothUuid__HIDInformation                                = QBluetoothUuid__CharacteristicType(0x2a4a)
	QBluetoothUuid__ReportMap                                     = QBluetoothUuid__CharacteristicType(0x2a4b)
	QBluetoothUuid__HIDControlPoint                               = QBluetoothUuid__CharacteristicType(0x2a4c)
	QBluetoothUuid__Report                                        = QBluetoothUuid__CharacteristicType(0x2a4d)
	QBluetoothUuid__ProtocolMode                                  = QBluetoothUuid__CharacteristicType(0x2a4e)
	QBluetoothUuid__ScanIntervalWindow                            = QBluetoothUuid__CharacteristicType(0x2a4f)
	QBluetoothUuid__PnPID                                         = QBluetoothUuid__CharacteristicType(0x2a50)
	QBluetoothUuid__GlucoseFeature                                = QBluetoothUuid__CharacteristicType(0x2a51)
	QBluetoothUuid__RecordAccessControlPoint                      = QBluetoothUuid__CharacteristicType(0x2a52)
	QBluetoothUuid__RSCMeasurement                                = QBluetoothUuid__CharacteristicType(0x2a53)
	QBluetoothUuid__RSCFeature                                    = QBluetoothUuid__CharacteristicType(0x2a54)
	QBluetoothUuid__SCControlPoint                                = QBluetoothUuid__CharacteristicType(0x2a55)
	QBluetoothUuid__CSCMeasurement                                = QBluetoothUuid__CharacteristicType(0x2a5b)
	QBluetoothUuid__CSCFeature                                    = QBluetoothUuid__CharacteristicType(0x2a5c)
	QBluetoothUuid__SensorLocation                                = QBluetoothUuid__CharacteristicType(0x2a5d)
	QBluetoothUuid__CyclingPowerMeasurement                       = QBluetoothUuid__CharacteristicType(0x2a63)
	QBluetoothUuid__CyclingPowerVector                            = QBluetoothUuid__CharacteristicType(0x2a64)
	QBluetoothUuid__CyclingPowerFeature                           = QBluetoothUuid__CharacteristicType(0x2a65)
	QBluetoothUuid__CyclingPowerControlPoint                      = QBluetoothUuid__CharacteristicType(0x2a66)
	QBluetoothUuid__LocationAndSpeed                              = QBluetoothUuid__CharacteristicType(0x2a67)
	QBluetoothUuid__Navigation                                    = QBluetoothUuid__CharacteristicType(0x2a68)
	QBluetoothUuid__PositionQuality                               = QBluetoothUuid__CharacteristicType(0x2a69)
	QBluetoothUuid__LNFeature                                     = QBluetoothUuid__CharacteristicType(0x2a6a)
	QBluetoothUuid__LNControlPoint                                = QBluetoothUuid__CharacteristicType(0x2a6b)
	QBluetoothUuid__Elevation                                     = QBluetoothUuid__CharacteristicType(0x2a6c)
	QBluetoothUuid__Pressure                                      = QBluetoothUuid__CharacteristicType(0x2a6d)
	QBluetoothUuid__Temperature                                   = QBluetoothUuid__CharacteristicType(0x2a6e)
	QBluetoothUuid__Humidity                                      = QBluetoothUuid__CharacteristicType(0x2a6f)
	QBluetoothUuid__TrueWindSpeed                                 = QBluetoothUuid__CharacteristicType(0x2a70)
	QBluetoothUuid__TrueWindDirection                             = QBluetoothUuid__CharacteristicType(0x2a71)
	QBluetoothUuid__ApparentWindSpeed                             = QBluetoothUuid__CharacteristicType(0x2a72)
	QBluetoothUuid__ApparentWindDirection                         = QBluetoothUuid__CharacteristicType(0x2a73)
	QBluetoothUuid__GustFactor                                    = QBluetoothUuid__CharacteristicType(0x2a74)
	QBluetoothUuid__PollenConcentration                           = QBluetoothUuid__CharacteristicType(0x2a75)
	QBluetoothUuid__UVIndex                                       = QBluetoothUuid__CharacteristicType(0x2a76)
	QBluetoothUuid__Irradiance                                    = QBluetoothUuid__CharacteristicType(0x2a77)
	QBluetoothUuid__Rainfall                                      = QBluetoothUuid__CharacteristicType(0x2a78)
	QBluetoothUuid__WindChill                                     = QBluetoothUuid__CharacteristicType(0x2a79)
	QBluetoothUuid__HeatIndex                                     = QBluetoothUuid__CharacteristicType(0x2a7a)
	QBluetoothUuid__DewPoint                                      = QBluetoothUuid__CharacteristicType(0x2a7b)
	QBluetoothUuid__DescriptorValueChanged                        = QBluetoothUuid__CharacteristicType(0x2a7d)
	QBluetoothUuid__AerobicHeartRateLowerLimit                    = QBluetoothUuid__CharacteristicType(0x2a7e)
	QBluetoothUuid__AerobicThreshold                              = QBluetoothUuid__CharacteristicType(0x2a7f)
	QBluetoothUuid__Age                                           = QBluetoothUuid__CharacteristicType(0x2a80)
	QBluetoothUuid__AnaerobicHeartRateLowerLimit                  = QBluetoothUuid__CharacteristicType(0x2a81)
	QBluetoothUuid__AnaerobicHeartRateUpperLimit                  = QBluetoothUuid__CharacteristicType(0x2a82)
	QBluetoothUuid__AnaerobicThreshold                            = QBluetoothUuid__CharacteristicType(0x2a83)
	QBluetoothUuid__AerobicHeartRateUpperLimit                    = QBluetoothUuid__CharacteristicType(0x2a84)
	QBluetoothUuid__DateOfBirth                                   = QBluetoothUuid__CharacteristicType(0x2a85)
	QBluetoothUuid__DateOfThresholdAssessment                     = QBluetoothUuid__CharacteristicType(0x2a86)
	QBluetoothUuid__EmailAddress                                  = QBluetoothUuid__CharacteristicType(0x2a87)
	QBluetoothUuid__FatBurnHeartRateLowerLimit                    = QBluetoothUuid__CharacteristicType(0x2a88)
	QBluetoothUuid__FatBurnHeartRateUpperLimit                    = QBluetoothUuid__CharacteristicType(0x2a89)
	QBluetoothUuid__FirstName                                     = QBluetoothUuid__CharacteristicType(0x2a8a)
	QBluetoothUuid__FiveZoneHeartRateLimits                       = QBluetoothUuid__CharacteristicType(0x2a8b)
	QBluetoothUuid__Gender                                        = QBluetoothUuid__CharacteristicType(0x2a8c)
	QBluetoothUuid__HeartRateMax                                  = QBluetoothUuid__CharacteristicType(0x2a8d)
	QBluetoothUuid__Height                                        = QBluetoothUuid__CharacteristicType(0x2a8e)
	QBluetoothUuid__HipCircumference                              = QBluetoothUuid__CharacteristicType(0x2a8f)
	QBluetoothUuid__LastName                                      = QBluetoothUuid__CharacteristicType(0x2a90)
	QBluetoothUuid__MaximumRecommendedHeartRate                   = QBluetoothUuid__CharacteristicType(0x2a91)
	QBluetoothUuid__RestingHeartRate                              = QBluetoothUuid__CharacteristicType(0x2a92)
	QBluetoothUuid__SportTypeForAerobicAnaerobicThresholds        = QBluetoothUuid__CharacteristicType(0x2a93)
	QBluetoothUuid__ThreeZoneHeartRateLimits                      = QBluetoothUuid__CharacteristicType(0x2a94)
	QBluetoothUuid__TwoZoneHeartRateLimits                        = QBluetoothUuid__CharacteristicType(0x2a95)
	QBluetoothUuid__VO2Max                                        = QBluetoothUuid__CharacteristicType(0x2a96)
	QBluetoothUuid__WaistCircumference                            = QBluetoothUuid__CharacteristicType(0x2a97)
	QBluetoothUuid__Weight                                        = QBluetoothUuid__CharacteristicType(0x2a98)
	QBluetoothUuid__DatabaseChangeIncrement                       = QBluetoothUuid__CharacteristicType(0x2a99)
	QBluetoothUuid__UserIndex                                     = QBluetoothUuid__CharacteristicType(0x2a9a)
	QBluetoothUuid__BodyCompositionFeature                        = QBluetoothUuid__CharacteristicType(0x2a9b)
	QBluetoothUuid__BodyCompositionMeasurement                    = QBluetoothUuid__CharacteristicType(0x2a9c)
	QBluetoothUuid__WeightMeasurement                             = QBluetoothUuid__CharacteristicType(0x2a9d)
	QBluetoothUuid__WeightScaleFeature                            = QBluetoothUuid__CharacteristicType(0x2a9e)
	QBluetoothUuid__UserControlPoint                              = QBluetoothUuid__CharacteristicType(0x2a9f)
	QBluetoothUuid__MagneticFluxDensity2D                         = QBluetoothUuid__CharacteristicType(0x2aa0)
	QBluetoothUuid__MagneticFluxDensity3D                         = QBluetoothUuid__CharacteristicType(0x2aa1)
	QBluetoothUuid__Language                                      = QBluetoothUuid__CharacteristicType(0x2aa2)
	QBluetoothUuid__BarometricPressureTrend                       = QBluetoothUuid__CharacteristicType(0x2aa3)
)

//QBluetoothUuid::DescriptorType
type QBluetoothUuid__DescriptorType int64

const (
	QBluetoothUuid__UnknownDescriptorType              = QBluetoothUuid__DescriptorType(0x0)
	QBluetoothUuid__CharacteristicExtendedProperties   = QBluetoothUuid__DescriptorType(0x2900)
	QBluetoothUuid__CharacteristicUserDescription      = QBluetoothUuid__DescriptorType(0x2901)
	QBluetoothUuid__ClientCharacteristicConfiguration  = QBluetoothUuid__DescriptorType(0x2902)
	QBluetoothUuid__ServerCharacteristicConfiguration  = QBluetoothUuid__DescriptorType(0x2903)
	QBluetoothUuid__CharacteristicPresentationFormat   = QBluetoothUuid__DescriptorType(0x2904)
	QBluetoothUuid__CharacteristicAggregateFormat      = QBluetoothUuid__DescriptorType(0x2905)
	QBluetoothUuid__ValidRange                         = QBluetoothUuid__DescriptorType(0x2906)
	QBluetoothUuid__ExternalReportReference            = QBluetoothUuid__DescriptorType(0x2907)
	QBluetoothUuid__ReportReference                    = QBluetoothUuid__DescriptorType(0x2908)
	QBluetoothUuid__EnvironmentalSensingConfiguration  = QBluetoothUuid__DescriptorType(0x290b)
	QBluetoothUuid__EnvironmentalSensingMeasurement    = QBluetoothUuid__DescriptorType(0x290c)
	QBluetoothUuid__EnvironmentalSensingTriggerSetting = QBluetoothUuid__DescriptorType(0x290d)
)

//QBluetoothUuid::ProtocolUuid
type QBluetoothUuid__ProtocolUuid int64

const (
	QBluetoothUuid__Sdp                    = QBluetoothUuid__ProtocolUuid(0x0001)
	QBluetoothUuid__Udp                    = QBluetoothUuid__ProtocolUuid(0x0002)
	QBluetoothUuid__Rfcomm                 = QBluetoothUuid__ProtocolUuid(0x0003)
	QBluetoothUuid__Tcp                    = QBluetoothUuid__ProtocolUuid(0x0004)
	QBluetoothUuid__TcsBin                 = QBluetoothUuid__ProtocolUuid(0x0005)
	QBluetoothUuid__TcsAt                  = QBluetoothUuid__ProtocolUuid(0x0006)
	QBluetoothUuid__Att                    = QBluetoothUuid__ProtocolUuid(0x0007)
	QBluetoothUuid__Obex                   = QBluetoothUuid__ProtocolUuid(0x0008)
	QBluetoothUuid__Ip                     = QBluetoothUuid__ProtocolUuid(0x0009)
	QBluetoothUuid__Ftp                    = QBluetoothUuid__ProtocolUuid(0x000A)
	QBluetoothUuid__Http                   = QBluetoothUuid__ProtocolUuid(0x000C)
	QBluetoothUuid__Wsp                    = QBluetoothUuid__ProtocolUuid(0x000E)
	QBluetoothUuid__Bnep                   = QBluetoothUuid__ProtocolUuid(0x000F)
	QBluetoothUuid__Upnp                   = QBluetoothUuid__ProtocolUuid(0x0010)
	QBluetoothUuid__Hidp                   = QBluetoothUuid__ProtocolUuid(0x0011)
	QBluetoothUuid__HardcopyControlChannel = QBluetoothUuid__ProtocolUuid(0x0012)
	QBluetoothUuid__HardcopyDataChannel    = QBluetoothUuid__ProtocolUuid(0x0014)
	QBluetoothUuid__HardcopyNotification   = QBluetoothUuid__ProtocolUuid(0x0016)
	QBluetoothUuid__Avctp                  = QBluetoothUuid__ProtocolUuid(0x0017)
	QBluetoothUuid__Avdtp                  = QBluetoothUuid__ProtocolUuid(0x0019)
	QBluetoothUuid__Cmtp                   = QBluetoothUuid__ProtocolUuid(0x001B)
	QBluetoothUuid__UdiCPlain              = QBluetoothUuid__ProtocolUuid(0x001D)
	QBluetoothUuid__McapControlChannel     = QBluetoothUuid__ProtocolUuid(0x001E)
	QBluetoothUuid__McapDataChannel        = QBluetoothUuid__ProtocolUuid(0x001F)
	QBluetoothUuid__L2cap                  = QBluetoothUuid__ProtocolUuid(0x0100)
)

//QBluetoothUuid::ServiceClassUuid
type QBluetoothUuid__ServiceClassUuid int64

const (
	QBluetoothUuid__ServiceDiscoveryServer                = QBluetoothUuid__ServiceClassUuid(0x1000)
	QBluetoothUuid__BrowseGroupDescriptor                 = QBluetoothUuid__ServiceClassUuid(0x1001)
	QBluetoothUuid__PublicBrowseGroup                     = QBluetoothUuid__ServiceClassUuid(0x1002)
	QBluetoothUuid__SerialPort                            = QBluetoothUuid__ServiceClassUuid(0x1101)
	QBluetoothUuid__LANAccessUsingPPP                     = QBluetoothUuid__ServiceClassUuid(0x1102)
	QBluetoothUuid__DialupNetworking                      = QBluetoothUuid__ServiceClassUuid(0x1103)
	QBluetoothUuid__IrMCSync                              = QBluetoothUuid__ServiceClassUuid(0x1104)
	QBluetoothUuid__ObexObjectPush                        = QBluetoothUuid__ServiceClassUuid(0x1105)
	QBluetoothUuid__OBEXFileTransfer                      = QBluetoothUuid__ServiceClassUuid(0x1106)
	QBluetoothUuid__IrMCSyncCommand                       = QBluetoothUuid__ServiceClassUuid(0x1107)
	QBluetoothUuid__Headset                               = QBluetoothUuid__ServiceClassUuid(0x1108)
	QBluetoothUuid__AudioSource                           = QBluetoothUuid__ServiceClassUuid(0x110a)
	QBluetoothUuid__AudioSink                             = QBluetoothUuid__ServiceClassUuid(0x110b)
	QBluetoothUuid__AV_RemoteControlTarget                = QBluetoothUuid__ServiceClassUuid(0x110c)
	QBluetoothUuid__AdvancedAudioDistribution             = QBluetoothUuid__ServiceClassUuid(0x110d)
	QBluetoothUuid__AV_RemoteControl                      = QBluetoothUuid__ServiceClassUuid(0x110e)
	QBluetoothUuid__AV_RemoteControlController            = QBluetoothUuid__ServiceClassUuid(0x110f)
	QBluetoothUuid__HeadsetAG                             = QBluetoothUuid__ServiceClassUuid(0x1112)
	QBluetoothUuid__PANU                                  = QBluetoothUuid__ServiceClassUuid(0x1115)
	QBluetoothUuid__NAP                                   = QBluetoothUuid__ServiceClassUuid(0x1116)
	QBluetoothUuid__GN                                    = QBluetoothUuid__ServiceClassUuid(0x1117)
	QBluetoothUuid__DirectPrinting                        = QBluetoothUuid__ServiceClassUuid(0x1118)
	QBluetoothUuid__ReferencePrinting                     = QBluetoothUuid__ServiceClassUuid(0x1119)
	QBluetoothUuid__BasicImage                            = QBluetoothUuid__ServiceClassUuid(0x111a)
	QBluetoothUuid__ImagingResponder                      = QBluetoothUuid__ServiceClassUuid(0x111b)
	QBluetoothUuid__ImagingAutomaticArchive               = QBluetoothUuid__ServiceClassUuid(0x111c)
	QBluetoothUuid__ImagingReferenceObjects               = QBluetoothUuid__ServiceClassUuid(0x111d)
	QBluetoothUuid__Handsfree                             = QBluetoothUuid__ServiceClassUuid(0x111e)
	QBluetoothUuid__HandsfreeAudioGateway                 = QBluetoothUuid__ServiceClassUuid(0x111f)
	QBluetoothUuid__DirectPrintingReferenceObjectsService = QBluetoothUuid__ServiceClassUuid(0x1120)
	QBluetoothUuid__ReflectedUI                           = QBluetoothUuid__ServiceClassUuid(0x1121)
	QBluetoothUuid__BasicPrinting                         = QBluetoothUuid__ServiceClassUuid(0x1122)
	QBluetoothUuid__PrintingStatus                        = QBluetoothUuid__ServiceClassUuid(0x1123)
	QBluetoothUuid__HumanInterfaceDeviceService           = QBluetoothUuid__ServiceClassUuid(0x1124)
	QBluetoothUuid__HardcopyCableReplacement              = QBluetoothUuid__ServiceClassUuid(0x1125)
	QBluetoothUuid__HCRPrint                              = QBluetoothUuid__ServiceClassUuid(0x1126)
	QBluetoothUuid__HCRScan                               = QBluetoothUuid__ServiceClassUuid(0x1127)
	QBluetoothUuid__SIMAccess                             = QBluetoothUuid__ServiceClassUuid(0x112d)
	QBluetoothUuid__PhonebookAccessPCE                    = QBluetoothUuid__ServiceClassUuid(0x112e)
	QBluetoothUuid__PhonebookAccessPSE                    = QBluetoothUuid__ServiceClassUuid(0x112f)
	QBluetoothUuid__PhonebookAccess                       = QBluetoothUuid__ServiceClassUuid(0x1130)
	QBluetoothUuid__HeadsetHS                             = QBluetoothUuid__ServiceClassUuid(0x1131)
	QBluetoothUuid__MessageAccessServer                   = QBluetoothUuid__ServiceClassUuid(0x1132)
	QBluetoothUuid__MessageNotificationServer             = QBluetoothUuid__ServiceClassUuid(0x1133)
	QBluetoothUuid__MessageAccessProfile                  = QBluetoothUuid__ServiceClassUuid(0x1134)
	QBluetoothUuid__GNSS                                  = QBluetoothUuid__ServiceClassUuid(0x1135)
	QBluetoothUuid__GNSSServer                            = QBluetoothUuid__ServiceClassUuid(0x1136)
	QBluetoothUuid__Display3D                             = QBluetoothUuid__ServiceClassUuid(0x1137)
	QBluetoothUuid__Glasses3D                             = QBluetoothUuid__ServiceClassUuid(0x1138)
	QBluetoothUuid__Synchronization3D                     = QBluetoothUuid__ServiceClassUuid(0x1139)
	QBluetoothUuid__MPSProfile                            = QBluetoothUuid__ServiceClassUuid(0x113a)
	QBluetoothUuid__MPSService                            = QBluetoothUuid__ServiceClassUuid(0x113b)
	QBluetoothUuid__PnPInformation                        = QBluetoothUuid__ServiceClassUuid(0x1200)
	QBluetoothUuid__GenericNetworking                     = QBluetoothUuid__ServiceClassUuid(0x1201)
	QBluetoothUuid__GenericFileTransfer                   = QBluetoothUuid__ServiceClassUuid(0x1202)
	QBluetoothUuid__GenericAudio                          = QBluetoothUuid__ServiceClassUuid(0x1203)
	QBluetoothUuid__GenericTelephony                      = QBluetoothUuid__ServiceClassUuid(0x1204)
	QBluetoothUuid__VideoSource                           = QBluetoothUuid__ServiceClassUuid(0x1303)
	QBluetoothUuid__VideoSink                             = QBluetoothUuid__ServiceClassUuid(0x1304)
	QBluetoothUuid__VideoDistribution                     = QBluetoothUuid__ServiceClassUuid(0x1305)
	QBluetoothUuid__HDP                                   = QBluetoothUuid__ServiceClassUuid(0x1400)
	QBluetoothUuid__HDPSource                             = QBluetoothUuid__ServiceClassUuid(0x1401)
	QBluetoothUuid__HDPSink                               = QBluetoothUuid__ServiceClassUuid(0x1402)
	QBluetoothUuid__GenericAccess                         = QBluetoothUuid__ServiceClassUuid(0x1800)
	QBluetoothUuid__GenericAttribute                      = QBluetoothUuid__ServiceClassUuid(0x1801)
	QBluetoothUuid__ImmediateAlert                        = QBluetoothUuid__ServiceClassUuid(0x1802)
	QBluetoothUuid__LinkLoss                              = QBluetoothUuid__ServiceClassUuid(0x1803)
	QBluetoothUuid__TxPower                               = QBluetoothUuid__ServiceClassUuid(0x1804)
	QBluetoothUuid__CurrentTimeService                    = QBluetoothUuid__ServiceClassUuid(0x1805)
	QBluetoothUuid__ReferenceTimeUpdateService            = QBluetoothUuid__ServiceClassUuid(0x1806)
	QBluetoothUuid__NextDSTChangeService                  = QBluetoothUuid__ServiceClassUuid(0x1807)
	QBluetoothUuid__Glucose                               = QBluetoothUuid__ServiceClassUuid(0x1808)
	QBluetoothUuid__HealthThermometer                     = QBluetoothUuid__ServiceClassUuid(0x1809)
	QBluetoothUuid__DeviceInformation                     = QBluetoothUuid__ServiceClassUuid(0x180a)
	QBluetoothUuid__HeartRate                             = QBluetoothUuid__ServiceClassUuid(0x180d)
	QBluetoothUuid__PhoneAlertStatusService               = QBluetoothUuid__ServiceClassUuid(0x180e)
	QBluetoothUuid__BatteryService                        = QBluetoothUuid__ServiceClassUuid(0x180f)
	QBluetoothUuid__BloodPressure                         = QBluetoothUuid__ServiceClassUuid(0x1810)
	QBluetoothUuid__AlertNotificationService              = QBluetoothUuid__ServiceClassUuid(0x1811)
	QBluetoothUuid__HumanInterfaceDevice                  = QBluetoothUuid__ServiceClassUuid(0x1812)
	QBluetoothUuid__ScanParameters                        = QBluetoothUuid__ServiceClassUuid(0x1813)
	QBluetoothUuid__RunningSpeedAndCadence                = QBluetoothUuid__ServiceClassUuid(0x1814)
	QBluetoothUuid__CyclingSpeedAndCadence                = QBluetoothUuid__ServiceClassUuid(0x1816)
	QBluetoothUuid__CyclingPower                          = QBluetoothUuid__ServiceClassUuid(0x1818)
	QBluetoothUuid__LocationAndNavigation                 = QBluetoothUuid__ServiceClassUuid(0x1819)
	QBluetoothUuid__EnvironmentalSensing                  = QBluetoothUuid__ServiceClassUuid(0x181a)
	QBluetoothUuid__BodyComposition                       = QBluetoothUuid__ServiceClassUuid(0x181b)
	QBluetoothUuid__UserData                              = QBluetoothUuid__ServiceClassUuid(0x181c)
	QBluetoothUuid__WeightScale                           = QBluetoothUuid__ServiceClassUuid(0x181d)
	QBluetoothUuid__BondManagement                        = QBluetoothUuid__ServiceClassUuid(0x181e)
	QBluetoothUuid__ContinuousGlucoseMonitoring           = QBluetoothUuid__ServiceClassUuid(0x181f)
)

type QBluetoothUuid struct {
	core.QUuid
}

type QBluetoothUuid_ITF interface {
	core.QUuid_ITF
	QBluetoothUuid_PTR() *QBluetoothUuid
}

func (p *QBluetoothUuid) QBluetoothUuid_PTR() *QBluetoothUuid {
	return p
}

func (p *QBluetoothUuid) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QUuid_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothUuid) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QUuid_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothUuid(ptr QBluetoothUuid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothUuid_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothUuidFromPointer(ptr unsafe.Pointer) *QBluetoothUuid {
	var n = new(QBluetoothUuid)
	n.SetPointer(ptr)
	return n
}
func NewQBluetoothUuid() *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid())
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid4(uuid QBluetoothUuid__CharacteristicType) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid4(C.longlong(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid5(uuid QBluetoothUuid__DescriptorType) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid5(C.longlong(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid2(uuid QBluetoothUuid__ProtocolUuid) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid2(C.longlong(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid3(uuid QBluetoothUuid__ServiceClassUuid) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid3(C.longlong(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid10(uuid QBluetoothUuid_ITF) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid10(PointerFromQBluetoothUuid(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid9(uuid string) *QBluetoothUuid {
	var uuidC = C.CString(uuid)
	defer C.free(unsafe.Pointer(uuidC))
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid9(uuidC))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid11(uuid core.QUuid_ITF) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid11(core.PointerFromQUuid(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid6(uuid uint16) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid6(C.ushort(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid7(uuid uint) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid7(C.uint(uint32(uuid))))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func QBluetoothUuid_CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_CharacteristicToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_CharacteristicToString(C.longlong(uuid)))
}

func QBluetoothUuid_DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_DescriptorToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_DescriptorToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) MinimumSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothUuid_MinimumSize(ptr.Pointer())))
	}
	return 0
}

func QBluetoothUuid_ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_ProtocolToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_ProtocolToString(C.longlong(uuid)))
}

func QBluetoothUuid_ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_ServiceClassToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_ServiceClassToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) ToUInt16(ok bool) uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothUuid_ToUInt16(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok)))))
	}
	return 0
}

func (ptr *QBluetoothUuid) ToUInt32(ok bool) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QBluetoothUuid_ToUInt32(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))))
	}
	return 0
}

func (ptr *QBluetoothUuid) DestroyQBluetoothUuid() {
	if ptr.Pointer() != nil {
		C.QBluetoothUuid_DestroyQBluetoothUuid(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyAdvertisingData::Discoverability
type QLowEnergyAdvertisingData__Discoverability int64

const (
	QLowEnergyAdvertisingData__DiscoverabilityNone    = QLowEnergyAdvertisingData__Discoverability(0)
	QLowEnergyAdvertisingData__DiscoverabilityLimited = QLowEnergyAdvertisingData__Discoverability(1)
	QLowEnergyAdvertisingData__DiscoverabilityGeneral = QLowEnergyAdvertisingData__Discoverability(2)
)

type QLowEnergyAdvertisingData struct {
	ptr unsafe.Pointer
}

type QLowEnergyAdvertisingData_ITF interface {
	QLowEnergyAdvertisingData_PTR() *QLowEnergyAdvertisingData
}

func (p *QLowEnergyAdvertisingData) QLowEnergyAdvertisingData_PTR() *QLowEnergyAdvertisingData {
	return p
}

func (p *QLowEnergyAdvertisingData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyAdvertisingData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyAdvertisingData(ptr QLowEnergyAdvertisingData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyAdvertisingData_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyAdvertisingDataFromPointer(ptr unsafe.Pointer) *QLowEnergyAdvertisingData {
	var n = new(QLowEnergyAdvertisingData)
	n.SetPointer(ptr)
	return n
}
func NewQLowEnergyAdvertisingData() *QLowEnergyAdvertisingData {
	var tmpValue = NewQLowEnergyAdvertisingDataFromPointer(C.QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingData).DestroyQLowEnergyAdvertisingData)
	return tmpValue
}

func NewQLowEnergyAdvertisingData2(other QLowEnergyAdvertisingData_ITF) *QLowEnergyAdvertisingData {
	var tmpValue = NewQLowEnergyAdvertisingDataFromPointer(C.QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData2(PointerFromQLowEnergyAdvertisingData(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingData).DestroyQLowEnergyAdvertisingData)
	return tmpValue
}

func (ptr *QLowEnergyAdvertisingData) Discoverability() QLowEnergyAdvertisingData__Discoverability {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingData__Discoverability(C.QLowEnergyAdvertisingData_Discoverability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingData) IncludePowerLevel() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyAdvertisingData_IncludePowerLevel(ptr.Pointer()) != 0
	}
	return false
}

func QLowEnergyAdvertisingData_InvalidManufacturerId() uint16 {
	return uint16(C.QLowEnergyAdvertisingData_QLowEnergyAdvertisingData_InvalidManufacturerId())
}

func (ptr *QLowEnergyAdvertisingData) InvalidManufacturerId() uint16 {
	return uint16(C.QLowEnergyAdvertisingData_QLowEnergyAdvertisingData_InvalidManufacturerId())
}

func (ptr *QLowEnergyAdvertisingData) LocalName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyAdvertisingData_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyAdvertisingData) ManufacturerData() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyAdvertisingData_ManufacturerData(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) ManufacturerId() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QLowEnergyAdvertisingData_ManufacturerId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingData) RawData() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyAdvertisingData_RawData(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) Services() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyAdvertisingDataFromPointer(l.data).services_atList(i)
			}
			return out
		}(C.QLowEnergyAdvertisingData_Services(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) SetDiscoverability(mode QLowEnergyAdvertisingData__Discoverability) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetDiscoverability(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetIncludePowerLevel(doInclude bool) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetIncludePowerLevel(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(doInclude))))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetLocalName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QLowEnergyAdvertisingData_SetLocalName(ptr.Pointer(), nameC)
	}
}

func (ptr *QLowEnergyAdvertisingData) SetManufacturerData(id uint16, data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetManufacturerData(ptr.Pointer(), C.ushort(id), core.PointerFromQByteArray(data))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetRawData(data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetRawData(ptr.Pointer(), core.PointerFromQByteArray(data))
	}
}

func (ptr *QLowEnergyAdvertisingData) Swap(other QLowEnergyAdvertisingData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_Swap(ptr.Pointer(), PointerFromQLowEnergyAdvertisingData(other))
	}
}

func (ptr *QLowEnergyAdvertisingData) DestroyQLowEnergyAdvertisingData() {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_DestroyQLowEnergyAdvertisingData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyAdvertisingData) services_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyAdvertisingData_services_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

//QLowEnergyAdvertisingParameters::FilterPolicy
type QLowEnergyAdvertisingParameters__FilterPolicy int64

const (
	QLowEnergyAdvertisingParameters__IgnoreWhiteList                      = QLowEnergyAdvertisingParameters__FilterPolicy(0x00)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanning              = QLowEnergyAdvertisingParameters__FilterPolicy(0x01)
	QLowEnergyAdvertisingParameters__UseWhiteListForConnecting            = QLowEnergyAdvertisingParameters__FilterPolicy(0x02)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanningAndConnecting = QLowEnergyAdvertisingParameters__FilterPolicy(0x03)
)

//QLowEnergyAdvertisingParameters::Mode
type QLowEnergyAdvertisingParameters__Mode int64

const (
	QLowEnergyAdvertisingParameters__AdvInd        = QLowEnergyAdvertisingParameters__Mode(0x0)
	QLowEnergyAdvertisingParameters__AdvScanInd    = QLowEnergyAdvertisingParameters__Mode(0x2)
	QLowEnergyAdvertisingParameters__AdvNonConnInd = QLowEnergyAdvertisingParameters__Mode(0x3)
)

type QLowEnergyAdvertisingParameters struct {
	ptr unsafe.Pointer
}

type QLowEnergyAdvertisingParameters_ITF interface {
	QLowEnergyAdvertisingParameters_PTR() *QLowEnergyAdvertisingParameters
}

func (p *QLowEnergyAdvertisingParameters) QLowEnergyAdvertisingParameters_PTR() *QLowEnergyAdvertisingParameters {
	return p
}

func (p *QLowEnergyAdvertisingParameters) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyAdvertisingParameters) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyAdvertisingParameters(ptr QLowEnergyAdvertisingParameters_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyAdvertisingParameters_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyAdvertisingParametersFromPointer(ptr unsafe.Pointer) *QLowEnergyAdvertisingParameters {
	var n = new(QLowEnergyAdvertisingParameters)
	n.SetPointer(ptr)
	return n
}
func NewQLowEnergyAdvertisingParameters() *QLowEnergyAdvertisingParameters {
	var tmpValue = NewQLowEnergyAdvertisingParametersFromPointer(C.QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingParameters).DestroyQLowEnergyAdvertisingParameters)
	return tmpValue
}

func NewQLowEnergyAdvertisingParameters2(other QLowEnergyAdvertisingParameters_ITF) *QLowEnergyAdvertisingParameters {
	var tmpValue = NewQLowEnergyAdvertisingParametersFromPointer(C.QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters2(PointerFromQLowEnergyAdvertisingParameters(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingParameters).DestroyQLowEnergyAdvertisingParameters)
	return tmpValue
}

func (ptr *QLowEnergyAdvertisingParameters) FilterPolicy() QLowEnergyAdvertisingParameters__FilterPolicy {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingParameters__FilterPolicy(C.QLowEnergyAdvertisingParameters_FilterPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) MaximumInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyAdvertisingParameters_MaximumInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) MinimumInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyAdvertisingParameters_MinimumInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) Mode() QLowEnergyAdvertisingParameters__Mode {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingParameters__Mode(C.QLowEnergyAdvertisingParameters_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) SetInterval(minimum uint16, maximum uint16) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_SetInterval(ptr.Pointer(), C.ushort(minimum), C.ushort(maximum))
	}
}

func (ptr *QLowEnergyAdvertisingParameters) SetMode(mode QLowEnergyAdvertisingParameters__Mode) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_SetMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QLowEnergyAdvertisingParameters) Swap(other QLowEnergyAdvertisingParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_Swap(ptr.Pointer(), PointerFromQLowEnergyAdvertisingParameters(other))
	}
}

func (ptr *QLowEnergyAdvertisingParameters) DestroyQLowEnergyAdvertisingParameters() {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_DestroyQLowEnergyAdvertisingParameters(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyCharacteristic::PropertyType
type QLowEnergyCharacteristic__PropertyType int64

const (
	QLowEnergyCharacteristic__Unknown          = QLowEnergyCharacteristic__PropertyType(0x00)
	QLowEnergyCharacteristic__Broadcasting     = QLowEnergyCharacteristic__PropertyType(0x01)
	QLowEnergyCharacteristic__Read             = QLowEnergyCharacteristic__PropertyType(0x02)
	QLowEnergyCharacteristic__WriteNoResponse  = QLowEnergyCharacteristic__PropertyType(0x04)
	QLowEnergyCharacteristic__Write            = QLowEnergyCharacteristic__PropertyType(0x08)
	QLowEnergyCharacteristic__Notify           = QLowEnergyCharacteristic__PropertyType(0x10)
	QLowEnergyCharacteristic__Indicate         = QLowEnergyCharacteristic__PropertyType(0x20)
	QLowEnergyCharacteristic__WriteSigned      = QLowEnergyCharacteristic__PropertyType(0x40)
	QLowEnergyCharacteristic__ExtendedProperty = QLowEnergyCharacteristic__PropertyType(0x80)
)

type QLowEnergyCharacteristic struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristic_ITF interface {
	QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic
}

func (p *QLowEnergyCharacteristic) QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic {
	return p
}

func (p *QLowEnergyCharacteristic) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyCharacteristic) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyCharacteristic(ptr QLowEnergyCharacteristic_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristic_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristic {
	var n = new(QLowEnergyCharacteristic)
	n.SetPointer(ptr)
	return n
}
func NewQLowEnergyCharacteristic() *QLowEnergyCharacteristic {
	var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
	return tmpValue
}

func NewQLowEnergyCharacteristic2(other QLowEnergyCharacteristic_ITF) *QLowEnergyCharacteristic {
	var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(PointerFromQLowEnergyCharacteristic(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
	return tmpValue
}

func (ptr *QLowEnergyCharacteristic) Descriptor(uuid QBluetoothUuid_ITF) *QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyCharacteristic_Descriptor(ptr.Pointer(), PointerFromQBluetoothUuid(uuid)))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) Descriptors() []*QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyDescriptor {
			var out = make([]*QLowEnergyDescriptor, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyCharacteristicFromPointer(l.data).descriptors_atList(i)
			}
			return out
		}(C.QLowEnergyCharacteristic_Descriptors(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristic_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristic) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyCharacteristic_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristic) Properties() QLowEnergyCharacteristic__PropertyType {
	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristic_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristic) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyCharacteristic_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyCharacteristic_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) DestroyQLowEnergyCharacteristic() {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyCharacteristic) descriptors_atList(i int) *QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyCharacteristic_descriptors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
		return tmpValue
	}
	return nil
}

type QLowEnergyCharacteristicData struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristicData_ITF interface {
	QLowEnergyCharacteristicData_PTR() *QLowEnergyCharacteristicData
}

func (p *QLowEnergyCharacteristicData) QLowEnergyCharacteristicData_PTR() *QLowEnergyCharacteristicData {
	return p
}

func (p *QLowEnergyCharacteristicData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyCharacteristicData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyCharacteristicData(ptr QLowEnergyCharacteristicData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristicData_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyCharacteristicDataFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristicData {
	var n = new(QLowEnergyCharacteristicData)
	n.SetPointer(ptr)
	return n
}
func NewQLowEnergyCharacteristicData() *QLowEnergyCharacteristicData {
	var tmpValue = NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
	return tmpValue
}

func NewQLowEnergyCharacteristicData2(other QLowEnergyCharacteristicData_ITF) *QLowEnergyCharacteristicData {
	var tmpValue = NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData2(PointerFromQLowEnergyCharacteristicData(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
	return tmpValue
}

func (ptr *QLowEnergyCharacteristicData) AddDescriptor(descriptor QLowEnergyDescriptorData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_AddDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(descriptor))
	}
}

func (ptr *QLowEnergyCharacteristicData) Descriptors() []*QLowEnergyDescriptorData {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyDescriptorData {
			var out = make([]*QLowEnergyDescriptorData, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyCharacteristicDataFromPointer(l.data).descriptors_atList(i)
			}
			return out
		}(C.QLowEnergyCharacteristicData_Descriptors(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristicData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristicData) MaximumValueLength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyCharacteristicData_MaximumValueLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) MinimumValueLength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyCharacteristicData_MinimumValueLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) Properties() QLowEnergyCharacteristic__PropertyType {
	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristicData_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) SetProperties(properties QLowEnergyCharacteristic__PropertyType) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetProperties(ptr.Pointer(), C.longlong(properties))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetValue(value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetValue(ptr.Pointer(), core.PointerFromQByteArray(value))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetValueLength(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetValueLength(ptr.Pointer(), C.int(int32(minimum)), C.int(int32(maximum)))
	}
}

func (ptr *QLowEnergyCharacteristicData) Swap(other QLowEnergyCharacteristicData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_Swap(ptr.Pointer(), PointerFromQLowEnergyCharacteristicData(other))
	}
}

func (ptr *QLowEnergyCharacteristicData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyCharacteristicData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyCharacteristicData_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) DestroyQLowEnergyCharacteristicData() {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_DestroyQLowEnergyCharacteristicData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyCharacteristicData) descriptors_atList(i int) *QLowEnergyDescriptorData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyCharacteristicData_descriptors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
		return tmpValue
	}
	return nil
}

type QLowEnergyConnectionParameters struct {
	ptr unsafe.Pointer
}

type QLowEnergyConnectionParameters_ITF interface {
	QLowEnergyConnectionParameters_PTR() *QLowEnergyConnectionParameters
}

func (p *QLowEnergyConnectionParameters) QLowEnergyConnectionParameters_PTR() *QLowEnergyConnectionParameters {
	return p
}

func (p *QLowEnergyConnectionParameters) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyConnectionParameters) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyConnectionParameters(ptr QLowEnergyConnectionParameters_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyConnectionParameters_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyConnectionParametersFromPointer(ptr unsafe.Pointer) *QLowEnergyConnectionParameters {
	var n = new(QLowEnergyConnectionParameters)
	n.SetPointer(ptr)
	return n
}
func NewQLowEnergyConnectionParameters() *QLowEnergyConnectionParameters {
	var tmpValue = NewQLowEnergyConnectionParametersFromPointer(C.QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyConnectionParameters).DestroyQLowEnergyConnectionParameters)
	return tmpValue
}

func NewQLowEnergyConnectionParameters2(other QLowEnergyConnectionParameters_ITF) *QLowEnergyConnectionParameters {
	var tmpValue = NewQLowEnergyConnectionParametersFromPointer(C.QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters2(PointerFromQLowEnergyConnectionParameters(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyConnectionParameters).DestroyQLowEnergyConnectionParameters)
	return tmpValue
}

func (ptr *QLowEnergyConnectionParameters) Latency() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyConnectionParameters_Latency(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) MaximumInterval() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLowEnergyConnectionParameters_MaximumInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) MinimumInterval() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLowEnergyConnectionParameters_MinimumInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) SetIntervalRange(minimum float64, maximum float64) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_SetIntervalRange(ptr.Pointer(), C.double(minimum), C.double(maximum))
	}
}

func (ptr *QLowEnergyConnectionParameters) SetLatency(latency int) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_SetLatency(ptr.Pointer(), C.int(int32(latency)))
	}
}

func (ptr *QLowEnergyConnectionParameters) SetSupervisionTimeout(timeout int) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_SetSupervisionTimeout(ptr.Pointer(), C.int(int32(timeout)))
	}
}

func (ptr *QLowEnergyConnectionParameters) SupervisionTimeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyConnectionParameters_SupervisionTimeout(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) Swap(other QLowEnergyConnectionParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_Swap(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(other))
	}
}

func (ptr *QLowEnergyConnectionParameters) DestroyQLowEnergyConnectionParameters() {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_DestroyQLowEnergyConnectionParameters(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyController::ControllerState
type QLowEnergyController__ControllerState int64

const (
	QLowEnergyController__UnconnectedState = QLowEnergyController__ControllerState(0)
	QLowEnergyController__ConnectingState  = QLowEnergyController__ControllerState(1)
	QLowEnergyController__ConnectedState   = QLowEnergyController__ControllerState(2)
	QLowEnergyController__DiscoveringState = QLowEnergyController__ControllerState(3)
	QLowEnergyController__DiscoveredState  = QLowEnergyController__ControllerState(4)
	QLowEnergyController__ClosingState     = QLowEnergyController__ControllerState(5)
	QLowEnergyController__AdvertisingState = QLowEnergyController__ControllerState(6)
)

//QLowEnergyController::Error
type QLowEnergyController__Error int64

const (
	QLowEnergyController__NoError                      = QLowEnergyController__Error(0)
	QLowEnergyController__UnknownError                 = QLowEnergyController__Error(1)
	QLowEnergyController__UnknownRemoteDeviceError     = QLowEnergyController__Error(2)
	QLowEnergyController__NetworkError                 = QLowEnergyController__Error(3)
	QLowEnergyController__InvalidBluetoothAdapterError = QLowEnergyController__Error(4)
	QLowEnergyController__ConnectionError              = QLowEnergyController__Error(5)
	QLowEnergyController__AdvertisingError             = QLowEnergyController__Error(6)
)

//QLowEnergyController::RemoteAddressType
type QLowEnergyController__RemoteAddressType int64

const (
	QLowEnergyController__PublicAddress = QLowEnergyController__RemoteAddressType(0)
	QLowEnergyController__RandomAddress = QLowEnergyController__RemoteAddressType(1)
)

//QLowEnergyController::Role
type QLowEnergyController__Role int64

const (
	QLowEnergyController__CentralRole    = QLowEnergyController__Role(0)
	QLowEnergyController__PeripheralRole = QLowEnergyController__Role(1)
)

type QLowEnergyController struct {
	core.QObject
}

type QLowEnergyController_ITF interface {
	core.QObject_ITF
	QLowEnergyController_PTR() *QLowEnergyController
}

func (p *QLowEnergyController) QLowEnergyController_PTR() *QLowEnergyController {
	return p
}

func (p *QLowEnergyController) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QLowEnergyController) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQLowEnergyController(ptr QLowEnergyController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyController_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyControllerFromPointer(ptr unsafe.Pointer) *QLowEnergyController {
	var n = new(QLowEnergyController)
	n.SetPointer(ptr)
	return n
}

//export callbackQLowEnergyController_Connected
func callbackQLowEnergyController_Connected(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::connected", f)
	}
}

func (ptr *QLowEnergyController) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::connected")
	}
}

func (ptr *QLowEnergyController) Connected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_Connected(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_ConnectionUpdated
func callbackQLowEnergyController_ConnectionUpdated(ptr unsafe.Pointer, newParameters unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::connectionUpdated"); signal != nil {
		signal.(func(*QLowEnergyConnectionParameters))(NewQLowEnergyConnectionParametersFromPointer(newParameters))
	}

}

func (ptr *QLowEnergyController) ConnectConnectionUpdated(f func(newParameters *QLowEnergyConnectionParameters)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectConnectionUpdated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::connectionUpdated", f)
	}
}

func (ptr *QLowEnergyController) DisconnectConnectionUpdated() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectConnectionUpdated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::connectionUpdated")
	}
}

func (ptr *QLowEnergyController) ConnectionUpdated(newParameters QLowEnergyConnectionParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectionUpdated(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(newParameters))
	}
}

//export callbackQLowEnergyController_Disconnected
func callbackQLowEnergyController_Disconnected(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::disconnected", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::disconnected")
	}
}

func (ptr *QLowEnergyController) Disconnected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_Disconnected(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_DiscoveryFinished
func callbackQLowEnergyController_DiscoveryFinished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::discoveryFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectDiscoveryFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectDiscoveryFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::discoveryFinished", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDiscoveryFinished() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDiscoveryFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::discoveryFinished")
	}
}

func (ptr *QLowEnergyController) DiscoveryFinished() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoveryFinished(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_Error2
func callbackQLowEnergyController_Error2(ptr unsafe.Pointer, newError C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::error2"); signal != nil {
		signal.(func(QLowEnergyController__Error))(QLowEnergyController__Error(newError))
	}

}

func (ptr *QLowEnergyController) ConnectError2(f func(newError QLowEnergyController__Error)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::error2", f)
	}
}

func (ptr *QLowEnergyController) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::error2")
	}
}

func (ptr *QLowEnergyController) Error2(newError QLowEnergyController__Error) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_Error2(ptr.Pointer(), C.longlong(newError))
	}
}

//export callbackQLowEnergyController_ServiceDiscovered
func callbackQLowEnergyController_ServiceDiscovered(ptr unsafe.Pointer, newService unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::serviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothUuid))(NewQBluetoothUuidFromPointer(newService))
	}

}

func (ptr *QLowEnergyController) ConnectServiceDiscovered(f func(newService *QBluetoothUuid)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectServiceDiscovered(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::serviceDiscovered", f)
	}
}

func (ptr *QLowEnergyController) DisconnectServiceDiscovered() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectServiceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::serviceDiscovered")
	}
}

func (ptr *QLowEnergyController) ServiceDiscovered(newService QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ServiceDiscovered(ptr.Pointer(), PointerFromQBluetoothUuid(newService))
	}
}

//export callbackQLowEnergyController_StateChanged
func callbackQLowEnergyController_StateChanged(ptr unsafe.Pointer, state C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::stateChanged"); signal != nil {
		signal.(func(QLowEnergyController__ControllerState))(QLowEnergyController__ControllerState(state))
	}

}

func (ptr *QLowEnergyController) ConnectStateChanged(f func(state QLowEnergyController__ControllerState)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::stateChanged", f)
	}
}

func (ptr *QLowEnergyController) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::stateChanged")
	}
}

func (ptr *QLowEnergyController) StateChanged(state QLowEnergyController__ControllerState) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

func (ptr *QLowEnergyController) AddService(service QLowEnergyServiceData_ITF, parent core.QObject_ITF) *QLowEnergyService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyServiceFromPointer(C.QLowEnergyController_AddService(ptr.Pointer(), PointerFromQLowEnergyServiceData(service), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) ConnectToDevice() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectToDevice(ptr.Pointer())
	}
}

func QLowEnergyController_CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	var tmpValue = NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreateCentral(PointerFromQBluetoothDeviceInfo(remoteDevice), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	var tmpValue = NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreateCentral(PointerFromQBluetoothDeviceInfo(remoteDevice), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QLowEnergyController_CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {
	var tmpValue = NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreatePeripheral(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {
	var tmpValue = NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreatePeripheral(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) CreateServiceObject(serviceUuid QBluetoothUuid_ITF, parent core.QObject_ITF) *QLowEnergyService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyServiceFromPointer(C.QLowEnergyController_CreateServiceObject(ptr.Pointer(), PointerFromQBluetoothUuid(serviceUuid), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) DisconnectFromDevice() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectFromDevice(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) DiscoverServices() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoverServices(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) Error() QLowEnergyController__Error {
	if ptr.Pointer() != nil {
		return QLowEnergyController__Error(C.QLowEnergyController_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyController_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) LocalAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QLowEnergyController_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QLowEnergyController_RemoteAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteAddressType() QLowEnergyController__RemoteAddressType {
	if ptr.Pointer() != nil {
		return QLowEnergyController__RemoteAddressType(C.QLowEnergyController_RemoteAddressType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) RemoteName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyController_RemoteName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) RequestConnectionUpdate(parameters QLowEnergyConnectionParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_RequestConnectionUpdate(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(parameters))
	}
}

func (ptr *QLowEnergyController) Role() QLowEnergyController__Role {
	if ptr.Pointer() != nil {
		return QLowEnergyController__Role(C.QLowEnergyController_Role(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) Services() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyControllerFromPointer(l.data).services_atList(i)
			}
			return out
		}(C.QLowEnergyController_Services(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyController) SetRemoteAddressType(ty QLowEnergyController__RemoteAddressType) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_SetRemoteAddressType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QLowEnergyController) StartAdvertising(parameters QLowEnergyAdvertisingParameters_ITF, advertisingData QLowEnergyAdvertisingData_ITF, scanResponseData QLowEnergyAdvertisingData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_StartAdvertising(ptr.Pointer(), PointerFromQLowEnergyAdvertisingParameters(parameters), PointerFromQLowEnergyAdvertisingData(advertisingData), PointerFromQLowEnergyAdvertisingData(scanResponseData))
	}
}

func (ptr *QLowEnergyController) State() QLowEnergyController__ControllerState {
	if ptr.Pointer() != nil {
		return QLowEnergyController__ControllerState(C.QLowEnergyController_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) StopAdvertising() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_StopAdvertising(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) DestroyQLowEnergyController() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DestroyQLowEnergyController(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyController) services_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyController_services_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

//export callbackQLowEnergyController_TimerEvent
func callbackQLowEnergyController_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::timerEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::timerEvent")
	}
}

func (ptr *QLowEnergyController) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLowEnergyController) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLowEnergyController_ChildEvent
func callbackQLowEnergyController_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::childEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::childEvent")
	}
}

func (ptr *QLowEnergyController) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLowEnergyController) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLowEnergyController_ConnectNotify
func callbackQLowEnergyController_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyController) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::connectNotify", f)
	}
}

func (ptr *QLowEnergyController) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::connectNotify")
	}
}

func (ptr *QLowEnergyController) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLowEnergyController) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyController_CustomEvent
func callbackQLowEnergyController_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::customEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::customEvent")
	}
}

func (ptr *QLowEnergyController) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLowEnergyController) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLowEnergyController_DeleteLater
func callbackQLowEnergyController_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLowEnergyController) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::deleteLater", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::deleteLater")
	}
}

func (ptr *QLowEnergyController) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyController) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLowEnergyController_DisconnectNotify
func callbackQLowEnergyController_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyController) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::disconnectNotify", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::disconnectNotify")
	}
}

func (ptr *QLowEnergyController) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLowEnergyController) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyController_Event
func callbackQLowEnergyController_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyControllerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLowEnergyController) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::event", f)
	}
}

func (ptr *QLowEnergyController) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::event")
	}
}

func (ptr *QLowEnergyController) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyController_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLowEnergyController) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyController_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLowEnergyController_EventFilter
func callbackQLowEnergyController_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyControllerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLowEnergyController) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::eventFilter", f)
	}
}

func (ptr *QLowEnergyController) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::eventFilter")
	}
}

func (ptr *QLowEnergyController) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyController_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLowEnergyController) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyController_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLowEnergyController_MetaObject
func callbackQLowEnergyController_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyController::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLowEnergyControllerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLowEnergyController) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::metaObject", f)
	}
}

func (ptr *QLowEnergyController) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyController::metaObject")
	}
}

func (ptr *QLowEnergyController) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyController_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyController) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyController_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QLowEnergyDescriptor struct {
	ptr unsafe.Pointer
}

type QLowEnergyDescriptor_ITF interface {
	QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor
}

func (p *QLowEnergyDescriptor) QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor {
	return p
}

func (p *QLowEnergyDescriptor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyDescriptor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyDescriptor(ptr QLowEnergyDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptor_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptor {
	var n = new(QLowEnergyDescriptor)
	n.SetPointer(ptr)
	return n
}
func NewQLowEnergyDescriptor() *QLowEnergyDescriptor {
	var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
	return tmpValue
}

func NewQLowEnergyDescriptor2(other QLowEnergyDescriptor_ITF) *QLowEnergyDescriptor {
	var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor2(PointerFromQLowEnergyDescriptor(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
	return tmpValue
}

func (ptr *QLowEnergyDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptor) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyDescriptor_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyDescriptor) Type() QBluetoothUuid__DescriptorType {
	if ptr.Pointer() != nil {
		return QBluetoothUuid__DescriptorType(C.QLowEnergyDescriptor_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptor) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyDescriptor_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyDescriptor_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) DestroyQLowEnergyDescriptor() {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLowEnergyDescriptorData struct {
	ptr unsafe.Pointer
}

type QLowEnergyDescriptorData_ITF interface {
	QLowEnergyDescriptorData_PTR() *QLowEnergyDescriptorData
}

func (p *QLowEnergyDescriptorData) QLowEnergyDescriptorData_PTR() *QLowEnergyDescriptorData {
	return p
}

func (p *QLowEnergyDescriptorData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyDescriptorData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyDescriptorData(ptr QLowEnergyDescriptorData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptorData_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyDescriptorDataFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptorData {
	var n = new(QLowEnergyDescriptorData)
	n.SetPointer(ptr)
	return n
}
func NewQLowEnergyDescriptorData() *QLowEnergyDescriptorData {
	var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func NewQLowEnergyDescriptorData2(uuid QBluetoothUuid_ITF, value core.QByteArray_ITF) *QLowEnergyDescriptorData {
	var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData2(PointerFromQBluetoothUuid(uuid), core.PointerFromQByteArray(value)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func NewQLowEnergyDescriptorData3(other QLowEnergyDescriptorData_ITF) *QLowEnergyDescriptorData {
	var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData3(PointerFromQLowEnergyDescriptorData(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func (ptr *QLowEnergyDescriptorData) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) SetUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyDescriptorData) SetValue(value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_SetValue(ptr.Pointer(), core.PointerFromQByteArray(value))
	}
}

func (ptr *QLowEnergyDescriptorData) Swap(other QLowEnergyDescriptorData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_Swap(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(other))
	}
}

func (ptr *QLowEnergyDescriptorData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyDescriptorData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyDescriptorData_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) DestroyQLowEnergyDescriptorData() {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_DestroyQLowEnergyDescriptorData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyService::ServiceError
type QLowEnergyService__ServiceError int64

const (
	QLowEnergyService__NoError                  = QLowEnergyService__ServiceError(0)
	QLowEnergyService__OperationError           = QLowEnergyService__ServiceError(1)
	QLowEnergyService__CharacteristicWriteError = QLowEnergyService__ServiceError(2)
	QLowEnergyService__DescriptorWriteError     = QLowEnergyService__ServiceError(3)
	QLowEnergyService__UnknownError             = QLowEnergyService__ServiceError(4)
	QLowEnergyService__CharacteristicReadError  = QLowEnergyService__ServiceError(5)
	QLowEnergyService__DescriptorReadError      = QLowEnergyService__ServiceError(6)
)

//QLowEnergyService::ServiceState
type QLowEnergyService__ServiceState int64

const (
	QLowEnergyService__InvalidService      = QLowEnergyService__ServiceState(0)
	QLowEnergyService__DiscoveryRequired   = QLowEnergyService__ServiceState(1)
	QLowEnergyService__DiscoveringServices = QLowEnergyService__ServiceState(2)
	QLowEnergyService__ServiceDiscovered   = QLowEnergyService__ServiceState(3)
	QLowEnergyService__LocalService        = QLowEnergyService__ServiceState(4)
)

//QLowEnergyService::ServiceType
type QLowEnergyService__ServiceType int64

const (
	QLowEnergyService__PrimaryService  = QLowEnergyService__ServiceType(0x0001)
	QLowEnergyService__IncludedService = QLowEnergyService__ServiceType(0x0002)
)

//QLowEnergyService::WriteMode
type QLowEnergyService__WriteMode int64

const (
	QLowEnergyService__WriteWithResponse    = QLowEnergyService__WriteMode(0)
	QLowEnergyService__WriteWithoutResponse = QLowEnergyService__WriteMode(1)
	QLowEnergyService__WriteSigned          = QLowEnergyService__WriteMode(2)
)

type QLowEnergyService struct {
	core.QObject
}

type QLowEnergyService_ITF interface {
	core.QObject_ITF
	QLowEnergyService_PTR() *QLowEnergyService
}

func (p *QLowEnergyService) QLowEnergyService_PTR() *QLowEnergyService {
	return p
}

func (p *QLowEnergyService) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QLowEnergyService) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQLowEnergyService(ptr QLowEnergyService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyService_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyServiceFromPointer(ptr unsafe.Pointer) *QLowEnergyService {
	var n = new(QLowEnergyService)
	n.SetPointer(ptr)
	return n
}

//export callbackQLowEnergyService_CharacteristicChanged
func callbackQLowEnergyService_CharacteristicChanged(ptr unsafe.Pointer, characteristic unsafe.Pointer, newValue unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::characteristicChanged"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicChanged(f func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectCharacteristicChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::characteristicChanged", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicChanged() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::characteristicChanged")
	}
}

func (ptr *QLowEnergyService) CharacteristicChanged(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CharacteristicChanged(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(newValue))
	}
}

//export callbackQLowEnergyService_CharacteristicRead
func callbackQLowEnergyService_CharacteristicRead(ptr unsafe.Pointer, characteristic unsafe.Pointer, value unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::characteristicRead"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(value))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicRead(f func(characteristic *QLowEnergyCharacteristic, value *core.QByteArray)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectCharacteristicRead(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::characteristicRead", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicRead() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicRead(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::characteristicRead")
	}
}

func (ptr *QLowEnergyService) CharacteristicRead(characteristic QLowEnergyCharacteristic_ITF, value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CharacteristicRead(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(value))
	}
}

//export callbackQLowEnergyService_CharacteristicWritten
func callbackQLowEnergyService_CharacteristicWritten(ptr unsafe.Pointer, characteristic unsafe.Pointer, newValue unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::characteristicWritten"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicWritten(f func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectCharacteristicWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::characteristicWritten", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicWritten() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::characteristicWritten")
	}
}

func (ptr *QLowEnergyService) CharacteristicWritten(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CharacteristicWritten(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(newValue))
	}
}

//export callbackQLowEnergyService_DescriptorRead
func callbackQLowEnergyService_DescriptorRead(ptr unsafe.Pointer, descriptor unsafe.Pointer, value unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::descriptorRead"); signal != nil {
		signal.(func(*QLowEnergyDescriptor, *core.QByteArray))(NewQLowEnergyDescriptorFromPointer(descriptor), core.NewQByteArrayFromPointer(value))
	}

}

func (ptr *QLowEnergyService) ConnectDescriptorRead(f func(descriptor *QLowEnergyDescriptor, value *core.QByteArray)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectDescriptorRead(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::descriptorRead", f)
	}
}

func (ptr *QLowEnergyService) DisconnectDescriptorRead() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectDescriptorRead(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::descriptorRead")
	}
}

func (ptr *QLowEnergyService) DescriptorRead(descriptor QLowEnergyDescriptor_ITF, value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DescriptorRead(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), core.PointerFromQByteArray(value))
	}
}

//export callbackQLowEnergyService_DescriptorWritten
func callbackQLowEnergyService_DescriptorWritten(ptr unsafe.Pointer, descriptor unsafe.Pointer, newValue unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::descriptorWritten"); signal != nil {
		signal.(func(*QLowEnergyDescriptor, *core.QByteArray))(NewQLowEnergyDescriptorFromPointer(descriptor), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectDescriptorWritten(f func(descriptor *QLowEnergyDescriptor, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectDescriptorWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::descriptorWritten", f)
	}
}

func (ptr *QLowEnergyService) DisconnectDescriptorWritten() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectDescriptorWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::descriptorWritten")
	}
}

func (ptr *QLowEnergyService) DescriptorWritten(descriptor QLowEnergyDescriptor_ITF, newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DescriptorWritten(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), core.PointerFromQByteArray(newValue))
	}
}

//export callbackQLowEnergyService_Error2
func callbackQLowEnergyService_Error2(ptr unsafe.Pointer, newError C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::error2"); signal != nil {
		signal.(func(QLowEnergyService__ServiceError))(QLowEnergyService__ServiceError(newError))
	}

}

func (ptr *QLowEnergyService) ConnectError2(f func(newError QLowEnergyService__ServiceError)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::error2", f)
	}
}

func (ptr *QLowEnergyService) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::error2")
	}
}

func (ptr *QLowEnergyService) Error2(newError QLowEnergyService__ServiceError) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_Error2(ptr.Pointer(), C.longlong(newError))
	}
}

//export callbackQLowEnergyService_StateChanged
func callbackQLowEnergyService_StateChanged(ptr unsafe.Pointer, newState C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::stateChanged"); signal != nil {
		signal.(func(QLowEnergyService__ServiceState))(QLowEnergyService__ServiceState(newState))
	}

}

func (ptr *QLowEnergyService) ConnectStateChanged(f func(newState QLowEnergyService__ServiceState)) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::stateChanged", f)
	}
}

func (ptr *QLowEnergyService) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::stateChanged")
	}
}

func (ptr *QLowEnergyService) StateChanged(newState QLowEnergyService__ServiceState) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_StateChanged(ptr.Pointer(), C.longlong(newState))
	}
}

func (ptr *QLowEnergyService) Characteristic(uuid QBluetoothUuid_ITF) *QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyService_Characteristic(ptr.Pointer(), PointerFromQBluetoothUuid(uuid)))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) Characteristics() []*QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyCharacteristic {
			var out = make([]*QLowEnergyCharacteristic, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyServiceFromPointer(l.data).characteristics_atList(i)
			}
			return out
		}(C.QLowEnergyService_Characteristics(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyService) Contains(characteristic QLowEnergyCharacteristic_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) Contains2(descriptor QLowEnergyDescriptor_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains2(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) DiscoverDetails() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DiscoverDetails(ptr.Pointer())
	}
}

func (ptr *QLowEnergyService) Error() QLowEnergyService__ServiceError {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceError(C.QLowEnergyService_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) IncludedServices() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyServiceFromPointer(l.data).includedServices_atList(i)
			}
			return out
		}(C.QLowEnergyService_IncludedServices(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyService) ReadCharacteristic(characteristic QLowEnergyCharacteristic_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic))
	}
}

func (ptr *QLowEnergyService) ReadDescriptor(descriptor QLowEnergyDescriptor_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor))
	}
}

func (ptr *QLowEnergyService) ServiceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyService_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyService) ServiceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyService_ServiceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) State() QLowEnergyService__ServiceState {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceState(C.QLowEnergyService_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) Type() QLowEnergyService__ServiceType {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceType(C.QLowEnergyService_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) WriteCharacteristic(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF, mode QLowEnergyService__WriteMode) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_WriteCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(newValue), C.longlong(mode))
	}
}

func (ptr *QLowEnergyService) WriteDescriptor(descriptor QLowEnergyDescriptor_ITF, newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_WriteDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), core.PointerFromQByteArray(newValue))
	}
}

func (ptr *QLowEnergyService) DestroyQLowEnergyService() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DestroyQLowEnergyService(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyService) characteristics_atList(i int) *QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyService_characteristics_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) includedServices_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyService_includedServices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

//export callbackQLowEnergyService_TimerEvent
func callbackQLowEnergyService_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::timerEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::timerEvent")
	}
}

func (ptr *QLowEnergyService) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLowEnergyService) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLowEnergyService_ChildEvent
func callbackQLowEnergyService_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::childEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::childEvent")
	}
}

func (ptr *QLowEnergyService) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLowEnergyService) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLowEnergyService_ConnectNotify
func callbackQLowEnergyService_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyService) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::connectNotify", f)
	}
}

func (ptr *QLowEnergyService) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::connectNotify")
	}
}

func (ptr *QLowEnergyService) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLowEnergyService) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyService_CustomEvent
func callbackQLowEnergyService_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::customEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::customEvent")
	}
}

func (ptr *QLowEnergyService) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLowEnergyService) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLowEnergyService_DeleteLater
func callbackQLowEnergyService_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLowEnergyService) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::deleteLater", f)
	}
}

func (ptr *QLowEnergyService) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::deleteLater")
	}
}

func (ptr *QLowEnergyService) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyService) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLowEnergyService_DisconnectNotify
func callbackQLowEnergyService_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyService) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::disconnectNotify", f)
	}
}

func (ptr *QLowEnergyService) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::disconnectNotify")
	}
}

func (ptr *QLowEnergyService) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLowEnergyService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyService_Event
func callbackQLowEnergyService_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyServiceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLowEnergyService) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::event", f)
	}
}

func (ptr *QLowEnergyService) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::event")
	}
}

func (ptr *QLowEnergyService) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLowEnergyService_EventFilter
func callbackQLowEnergyService_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyServiceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLowEnergyService) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::eventFilter", f)
	}
}

func (ptr *QLowEnergyService) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::eventFilter")
	}
}

func (ptr *QLowEnergyService) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLowEnergyService_MetaObject
func callbackQLowEnergyService_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLowEnergyService::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLowEnergyServiceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLowEnergyService) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::metaObject", f)
	}
}

func (ptr *QLowEnergyService) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLowEnergyService::metaObject")
	}
}

func (ptr *QLowEnergyService) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyService_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyService) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyService_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QLowEnergyServiceData::ServiceType
type QLowEnergyServiceData__ServiceType int64

const (
	QLowEnergyServiceData__ServiceTypePrimary   = QLowEnergyServiceData__ServiceType(0x2800)
	QLowEnergyServiceData__ServiceTypeSecondary = QLowEnergyServiceData__ServiceType(0x2801)
)

type QLowEnergyServiceData struct {
	ptr unsafe.Pointer
}

type QLowEnergyServiceData_ITF interface {
	QLowEnergyServiceData_PTR() *QLowEnergyServiceData
}

func (p *QLowEnergyServiceData) QLowEnergyServiceData_PTR() *QLowEnergyServiceData {
	return p
}

func (p *QLowEnergyServiceData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyServiceData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyServiceData(ptr QLowEnergyServiceData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyServiceData_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyServiceDataFromPointer(ptr unsafe.Pointer) *QLowEnergyServiceData {
	var n = new(QLowEnergyServiceData)
	n.SetPointer(ptr)
	return n
}
func NewQLowEnergyServiceData() *QLowEnergyServiceData {
	var tmpValue = NewQLowEnergyServiceDataFromPointer(C.QLowEnergyServiceData_NewQLowEnergyServiceData())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyServiceData).DestroyQLowEnergyServiceData)
	return tmpValue
}

func NewQLowEnergyServiceData2(other QLowEnergyServiceData_ITF) *QLowEnergyServiceData {
	var tmpValue = NewQLowEnergyServiceDataFromPointer(C.QLowEnergyServiceData_NewQLowEnergyServiceData2(PointerFromQLowEnergyServiceData(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyServiceData).DestroyQLowEnergyServiceData)
	return tmpValue
}

func (ptr *QLowEnergyServiceData) AddCharacteristic(characteristic QLowEnergyCharacteristicData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_AddCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristicData(characteristic))
	}
}

func (ptr *QLowEnergyServiceData) AddIncludedService(service QLowEnergyService_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_AddIncludedService(ptr.Pointer(), PointerFromQLowEnergyService(service))
	}
}

func (ptr *QLowEnergyServiceData) Characteristics() []*QLowEnergyCharacteristicData {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyCharacteristicData {
			var out = make([]*QLowEnergyCharacteristicData, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyServiceDataFromPointer(l.data).characteristics_atList(i)
			}
			return out
		}(C.QLowEnergyServiceData_Characteristics(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyServiceData) IncludedServices() []*QLowEnergyService {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyService {
			var out = make([]*QLowEnergyService, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyServiceDataFromPointer(l.data).includedServices_atList(i)
			}
			return out
		}(C.QLowEnergyServiceData_IncludedServices(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyServiceData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyServiceData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyServiceData) SetType(ty QLowEnergyServiceData__ServiceType) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QLowEnergyServiceData) SetUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyServiceData) Swap(other QLowEnergyServiceData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_Swap(ptr.Pointer(), PointerFromQLowEnergyServiceData(other))
	}
}

func (ptr *QLowEnergyServiceData) Type() QLowEnergyServiceData__ServiceType {
	if ptr.Pointer() != nil {
		return QLowEnergyServiceData__ServiceType(C.QLowEnergyServiceData_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyServiceData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyServiceData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) DestroyQLowEnergyServiceData() {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_DestroyQLowEnergyServiceData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyServiceData) characteristics_atList(i int) *QLowEnergyCharacteristicData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyServiceData_characteristics_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) includedServices_atList(i int) *QLowEnergyService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyServiceFromPointer(C.QLowEnergyServiceData_includedServices_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}
