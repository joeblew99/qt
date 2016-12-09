// +build !minimal

#define protected public
#define private public

#include "bluetooth.h"
#include "_cgo_export.h"

#include <QBluetoothAddress>
#include <QBluetoothDeviceDiscoveryAgent>
#include <QBluetoothDeviceInfo>
#include <QBluetoothHostInfo>
#include <QBluetoothLocalDevice>
#include <QBluetoothServer>
#include <QBluetoothServiceDiscoveryAgent>
#include <QBluetoothServiceInfo>
#include <QBluetoothSocket>
#include <QBluetoothTransferManager>
#include <QBluetoothTransferReply>
#include <QBluetoothTransferRequest>
#include <QBluetoothUuid>
#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QIODevice>
#include <QList>
#include <QLowEnergyAdvertisingData>
#include <QLowEnergyAdvertisingParameters>
#include <QLowEnergyCharacteristic>
#include <QLowEnergyCharacteristicData>
#include <QLowEnergyConnectionParameters>
#include <QLowEnergyController>
#include <QLowEnergyDescriptor>
#include <QLowEnergyDescriptorData>
#include <QLowEnergyService>
#include <QLowEnergyServiceData>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUuid>
#include <QVariant>

void* QBluetoothAddress_NewQBluetoothAddress()
{
	return new QBluetoothAddress();
}

void* QBluetoothAddress_NewQBluetoothAddress4(void* other)
{
	return new QBluetoothAddress(*static_cast<QBluetoothAddress*>(other));
}

void* QBluetoothAddress_NewQBluetoothAddress3(char* address)
{
	return new QBluetoothAddress(QString(address));
}

void* QBluetoothAddress_NewQBluetoothAddress2(unsigned long long address)
{
	return new QBluetoothAddress(address);
}

void QBluetoothAddress_Clear(void* ptr)
{
	static_cast<QBluetoothAddress*>(ptr)->clear();
}

char QBluetoothAddress_IsNull(void* ptr)
{
	return static_cast<QBluetoothAddress*>(ptr)->isNull();
}

struct QtBluetooth_PackedString QBluetoothAddress_ToString(void* ptr)
{
	return ({ QByteArray t00210e = static_cast<QBluetoothAddress*>(ptr)->toString().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t00210e.prepend("WHITESPACE").constData()+10), t00210e.size()-10 }; });
}

unsigned long long QBluetoothAddress_ToUInt64(void* ptr)
{
	return static_cast<QBluetoothAddress*>(ptr)->toUInt64();
}

void QBluetoothAddress_DestroyQBluetoothAddress(void* ptr)
{
	static_cast<QBluetoothAddress*>(ptr)->~QBluetoothAddress();
}

class MyQBluetoothDeviceDiscoveryAgent: public QBluetoothDeviceDiscoveryAgent
{
public:
	MyQBluetoothDeviceDiscoveryAgent(QObject *parent) : QBluetoothDeviceDiscoveryAgent(parent) {};
	MyQBluetoothDeviceDiscoveryAgent(const QBluetoothAddress &deviceAdapter, QObject *parent) : QBluetoothDeviceDiscoveryAgent(deviceAdapter, parent) {};
	void Signal_Canceled() { callbackQBluetoothDeviceDiscoveryAgent_Canceled(this); };
	void Signal_DeviceDiscovered(const QBluetoothDeviceInfo & info) { callbackQBluetoothDeviceDiscoveryAgent_DeviceDiscovered(this, const_cast<QBluetoothDeviceInfo*>(&info)); };
	void Signal_Error2(QBluetoothDeviceDiscoveryAgent::Error error) { callbackQBluetoothDeviceDiscoveryAgent_Error2(this, error); };
	void Signal_Finished() { callbackQBluetoothDeviceDiscoveryAgent_Finished(this); };
	void start() { callbackQBluetoothDeviceDiscoveryAgent_Start(this); };
	void stop() { callbackQBluetoothDeviceDiscoveryAgent_Stop(this); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothDeviceDiscoveryAgent_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothDeviceDiscoveryAgent_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBluetoothDeviceDiscoveryAgent_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBluetoothDeviceDiscoveryAgent_CustomEvent(this, event); };
	void deleteLater() { callbackQBluetoothDeviceDiscoveryAgent_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQBluetoothDeviceDiscoveryAgent_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBluetoothDeviceDiscoveryAgent_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBluetoothDeviceDiscoveryAgent_MetaObject(const_cast<MyQBluetoothDeviceDiscoveryAgent*>(this))); };
};

void QBluetoothDeviceDiscoveryAgent_ConnectCanceled(void* ptr)
{
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));
}

void QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::canceled), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Canceled));
}

void QBluetoothDeviceDiscoveryAgent_Canceled(void* ptr)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->canceled();
}

void QBluetoothDeviceDiscoveryAgent_ConnectDeviceDiscovered(void* ptr)
{
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)(const QBluetoothDeviceInfo &)>(&QBluetoothDeviceDiscoveryAgent::deviceDiscovered), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)(const QBluetoothDeviceInfo &)>(&MyQBluetoothDeviceDiscoveryAgent::Signal_DeviceDiscovered));
}

void QBluetoothDeviceDiscoveryAgent_DisconnectDeviceDiscovered(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)(const QBluetoothDeviceInfo &)>(&QBluetoothDeviceDiscoveryAgent::deviceDiscovered), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)(const QBluetoothDeviceInfo &)>(&MyQBluetoothDeviceDiscoveryAgent::Signal_DeviceDiscovered));
}

void QBluetoothDeviceDiscoveryAgent_DeviceDiscovered(void* ptr, void* info)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->deviceDiscovered(*static_cast<QBluetoothDeviceInfo*>(info));
}

void QBluetoothDeviceDiscoveryAgent_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)(QBluetoothDeviceDiscoveryAgent::Error)>(&QBluetoothDeviceDiscoveryAgent::error), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)(QBluetoothDeviceDiscoveryAgent::Error)>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Error2));
}

void QBluetoothDeviceDiscoveryAgent_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)(QBluetoothDeviceDiscoveryAgent::Error)>(&QBluetoothDeviceDiscoveryAgent::error), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)(QBluetoothDeviceDiscoveryAgent::Error)>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Error2));
}

void QBluetoothDeviceDiscoveryAgent_Error2(void* ptr, long long error)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->error(static_cast<QBluetoothDeviceDiscoveryAgent::Error>(error));
}

void QBluetoothDeviceDiscoveryAgent_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));
}

void QBluetoothDeviceDiscoveryAgent_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothDeviceDiscoveryAgent::*)()>(&QBluetoothDeviceDiscoveryAgent::finished), static_cast<MyQBluetoothDeviceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothDeviceDiscoveryAgent::*)()>(&MyQBluetoothDeviceDiscoveryAgent::Signal_Finished));
}

void QBluetoothDeviceDiscoveryAgent_Finished(void* ptr)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->finished();
}

void* QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(void* parent)
{
	return new MyQBluetoothDeviceDiscoveryAgent(static_cast<QObject*>(parent));
}

void* QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(void* deviceAdapter, void* parent)
{
	return new MyQBluetoothDeviceDiscoveryAgent(*static_cast<QBluetoothAddress*>(deviceAdapter), static_cast<QObject*>(parent));
}

struct QtBluetooth_PackedList QBluetoothDeviceDiscoveryAgent_DiscoveredDevices(void* ptr)
{
	return ({ QList<QBluetoothDeviceInfo>* tmpValue = new QList<QBluetoothDeviceInfo>(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->discoveredDevices()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

long long QBluetoothDeviceDiscoveryAgent_Error(void* ptr)
{
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->error();
}

struct QtBluetooth_PackedString QBluetoothDeviceDiscoveryAgent_ErrorString(void* ptr)
{
	return ({ QByteArray ta5ce9a = static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->errorString().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(ta5ce9a.prepend("WHITESPACE").constData()+10), ta5ce9a.size()-10 }; });
}

long long QBluetoothDeviceDiscoveryAgent_InquiryType(void* ptr)
{
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->inquiryType();
}

char QBluetoothDeviceDiscoveryAgent_IsActive(void* ptr)
{
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->isActive();
}

void QBluetoothDeviceDiscoveryAgent_SetInquiryType(void* ptr, long long ty)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->setInquiryType(static_cast<QBluetoothDeviceDiscoveryAgent::InquiryType>(ty));
}

void QBluetoothDeviceDiscoveryAgent_Start(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "start");
}

void QBluetoothDeviceDiscoveryAgent_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "stop");
}

void QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(void* ptr)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->~QBluetoothDeviceDiscoveryAgent();
}

void* QBluetoothDeviceDiscoveryAgent_discoveredDevices_atList(void* ptr, int i)
{
	return new QBluetoothDeviceInfo(static_cast<QList<QBluetoothDeviceInfo>*>(ptr)->at(i));
}

void QBluetoothDeviceDiscoveryAgent_TimerEvent(void* ptr, void* event)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_ChildEvent(void* ptr, void* event)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothDeviceDiscoveryAgent_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothDeviceDiscoveryAgent_CustomEvent(void* ptr, void* event)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::customEvent(static_cast<QEvent*>(event));
}

void QBluetoothDeviceDiscoveryAgent_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr), "deleteLater");
}

void QBluetoothDeviceDiscoveryAgent_DeleteLaterDefault(void* ptr)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::deleteLater();
}

void QBluetoothDeviceDiscoveryAgent_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothDeviceDiscoveryAgent_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QBluetoothDeviceDiscoveryAgent_Event(void* ptr, void* e)
{
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->event(static_cast<QEvent*>(e));
}

char QBluetoothDeviceDiscoveryAgent_EventDefault(void* ptr, void* e)
{
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::event(static_cast<QEvent*>(e));
}

char QBluetoothDeviceDiscoveryAgent_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QBluetoothDeviceDiscoveryAgent_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QBluetoothDeviceDiscoveryAgent_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->metaObject());
}

void* QBluetoothDeviceDiscoveryAgent_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothDeviceDiscoveryAgent*>(ptr)->QBluetoothDeviceDiscoveryAgent::metaObject());
}

void* QBluetoothDeviceInfo_NewQBluetoothDeviceInfo()
{
	return new QBluetoothDeviceInfo();
}

void* QBluetoothDeviceInfo_NewQBluetoothDeviceInfo2(void* address, char* name, unsigned int classOfDevice)
{
	return new QBluetoothDeviceInfo(*static_cast<QBluetoothAddress*>(address), QString(name), classOfDevice);
}

void* QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(void* other)
{
	return new QBluetoothDeviceInfo(*static_cast<QBluetoothDeviceInfo*>(other));
}

void* QBluetoothDeviceInfo_NewQBluetoothDeviceInfo3(void* uuid, char* name, unsigned int classOfDevice)
{
	return new QBluetoothDeviceInfo(*static_cast<QBluetoothUuid*>(uuid), QString(name), classOfDevice);
}

void* QBluetoothDeviceInfo_Address(void* ptr)
{
	return new QBluetoothAddress(static_cast<QBluetoothDeviceInfo*>(ptr)->address());
}

long long QBluetoothDeviceInfo_CoreConfigurations(void* ptr)
{
	return static_cast<QBluetoothDeviceInfo*>(ptr)->coreConfigurations();
}

void* QBluetoothDeviceInfo_DeviceUuid(void* ptr)
{
	return new QBluetoothUuid(static_cast<QBluetoothDeviceInfo*>(ptr)->deviceUuid());
}

char QBluetoothDeviceInfo_IsCached(void* ptr)
{
	return static_cast<QBluetoothDeviceInfo*>(ptr)->isCached();
}

char QBluetoothDeviceInfo_IsValid(void* ptr)
{
	return static_cast<QBluetoothDeviceInfo*>(ptr)->isValid();
}

long long QBluetoothDeviceInfo_MajorDeviceClass(void* ptr)
{
	return static_cast<QBluetoothDeviceInfo*>(ptr)->majorDeviceClass();
}

struct QtBluetooth_PackedString QBluetoothDeviceInfo_MinorDeviceClass(void* ptr)
{
	return ({ quint8 pret789f3c = static_cast<QBluetoothDeviceInfo*>(ptr)->minorDeviceClass(); char* t789f3c = static_cast<char*>(static_cast<void*>(&pret789f3c)); QtBluetooth_PackedString { t789f3c, -1 }; });
}

struct QtBluetooth_PackedString QBluetoothDeviceInfo_Name(void* ptr)
{
	return ({ QByteArray t605955 = static_cast<QBluetoothDeviceInfo*>(ptr)->name().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t605955.prepend("WHITESPACE").constData()+10), t605955.size()-10 }; });
}

short QBluetoothDeviceInfo_Rssi(void* ptr)
{
	return static_cast<QBluetoothDeviceInfo*>(ptr)->rssi();
}

long long QBluetoothDeviceInfo_ServiceClasses(void* ptr)
{
	return static_cast<QBluetoothDeviceInfo*>(ptr)->serviceClasses();
}

long long QBluetoothDeviceInfo_ServiceUuidsCompleteness(void* ptr)
{
	return static_cast<QBluetoothDeviceInfo*>(ptr)->serviceUuidsCompleteness();
}

void QBluetoothDeviceInfo_SetCached(void* ptr, char cached)
{
	static_cast<QBluetoothDeviceInfo*>(ptr)->setCached(cached != 0);
}

void QBluetoothDeviceInfo_SetCoreConfigurations(void* ptr, long long coreConfigs)
{
	static_cast<QBluetoothDeviceInfo*>(ptr)->setCoreConfigurations(static_cast<QBluetoothDeviceInfo::CoreConfiguration>(coreConfigs));
}

void QBluetoothDeviceInfo_SetDeviceUuid(void* ptr, void* uuid)
{
	static_cast<QBluetoothDeviceInfo*>(ptr)->setDeviceUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void QBluetoothDeviceInfo_SetRssi(void* ptr, short sign)
{
	static_cast<QBluetoothDeviceInfo*>(ptr)->setRssi(sign);
}

void QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(void* ptr)
{
	static_cast<QBluetoothDeviceInfo*>(ptr)->~QBluetoothDeviceInfo();
}

void* QBluetoothDeviceInfo_serviceUuids_atList(void* ptr, int i)
{
	return new QBluetoothUuid(static_cast<QList<QBluetoothUuid>*>(ptr)->at(i));
}

void* QBluetoothHostInfo_NewQBluetoothHostInfo()
{
	return new QBluetoothHostInfo();
}

void* QBluetoothHostInfo_NewQBluetoothHostInfo2(void* other)
{
	return new QBluetoothHostInfo(*static_cast<QBluetoothHostInfo*>(other));
}

void* QBluetoothHostInfo_Address(void* ptr)
{
	return new QBluetoothAddress(static_cast<QBluetoothHostInfo*>(ptr)->address());
}

struct QtBluetooth_PackedString QBluetoothHostInfo_Name(void* ptr)
{
	return ({ QByteArray tabf4bd = static_cast<QBluetoothHostInfo*>(ptr)->name().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(tabf4bd.prepend("WHITESPACE").constData()+10), tabf4bd.size()-10 }; });
}

void QBluetoothHostInfo_SetAddress(void* ptr, void* address)
{
	static_cast<QBluetoothHostInfo*>(ptr)->setAddress(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothHostInfo_SetName(void* ptr, char* name)
{
	static_cast<QBluetoothHostInfo*>(ptr)->setName(QString(name));
}

void QBluetoothHostInfo_DestroyQBluetoothHostInfo(void* ptr)
{
	static_cast<QBluetoothHostInfo*>(ptr)->~QBluetoothHostInfo();
}

class MyQBluetoothLocalDevice: public QBluetoothLocalDevice
{
public:
	MyQBluetoothLocalDevice(QObject *parent) : QBluetoothLocalDevice(parent) {};
	MyQBluetoothLocalDevice(const QBluetoothAddress &address, QObject *parent) : QBluetoothLocalDevice(address, parent) {};
	void Signal_DeviceConnected(const QBluetoothAddress & address) { callbackQBluetoothLocalDevice_DeviceConnected(this, const_cast<QBluetoothAddress*>(&address)); };
	void Signal_DeviceDisconnected(const QBluetoothAddress & address) { callbackQBluetoothLocalDevice_DeviceDisconnected(this, const_cast<QBluetoothAddress*>(&address)); };
	void Signal_Error(QBluetoothLocalDevice::Error error) { callbackQBluetoothLocalDevice_Error(this, error); };
	void Signal_HostModeStateChanged(QBluetoothLocalDevice::HostMode state) { callbackQBluetoothLocalDevice_HostModeStateChanged(this, state); };
	void Signal_PairingDisplayConfirmation(const QBluetoothAddress & address, QString pin) { QByteArray td145a2 = pin.toUtf8(); QtBluetooth_PackedString pinPacked = { const_cast<char*>(td145a2.prepend("WHITESPACE").constData()+10), td145a2.size()-10 };callbackQBluetoothLocalDevice_PairingDisplayConfirmation(this, const_cast<QBluetoothAddress*>(&address), pinPacked); };
	void Signal_PairingDisplayPinCode(const QBluetoothAddress & address, QString pin) { QByteArray td145a2 = pin.toUtf8(); QtBluetooth_PackedString pinPacked = { const_cast<char*>(td145a2.prepend("WHITESPACE").constData()+10), td145a2.size()-10 };callbackQBluetoothLocalDevice_PairingDisplayPinCode(this, const_cast<QBluetoothAddress*>(&address), pinPacked); };
	void Signal_PairingFinished(const QBluetoothAddress & address, QBluetoothLocalDevice::Pairing pairing) { callbackQBluetoothLocalDevice_PairingFinished(this, const_cast<QBluetoothAddress*>(&address), pairing); };
	 ~MyQBluetoothLocalDevice() { callbackQBluetoothLocalDevice_DestroyQBluetoothLocalDevice(this); };
	void pairingConfirmation(bool accept) { callbackQBluetoothLocalDevice_PairingConfirmation(this, accept); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothLocalDevice_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothLocalDevice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBluetoothLocalDevice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBluetoothLocalDevice_CustomEvent(this, event); };
	void deleteLater() { callbackQBluetoothLocalDevice_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBluetoothLocalDevice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQBluetoothLocalDevice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBluetoothLocalDevice_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBluetoothLocalDevice_MetaObject(const_cast<MyQBluetoothLocalDevice*>(this))); };
};

void QBluetoothLocalDevice_ConnectDeviceConnected(void* ptr)
{
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &)>(&QBluetoothLocalDevice::deviceConnected), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &)>(&MyQBluetoothLocalDevice::Signal_DeviceConnected));
}

void QBluetoothLocalDevice_DisconnectDeviceConnected(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &)>(&QBluetoothLocalDevice::deviceConnected), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &)>(&MyQBluetoothLocalDevice::Signal_DeviceConnected));
}

void QBluetoothLocalDevice_DeviceConnected(void* ptr, void* address)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->deviceConnected(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothLocalDevice_ConnectDeviceDisconnected(void* ptr)
{
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &)>(&QBluetoothLocalDevice::deviceDisconnected), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &)>(&MyQBluetoothLocalDevice::Signal_DeviceDisconnected));
}

void QBluetoothLocalDevice_DisconnectDeviceDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &)>(&QBluetoothLocalDevice::deviceDisconnected), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &)>(&MyQBluetoothLocalDevice::Signal_DeviceDisconnected));
}

void QBluetoothLocalDevice_DeviceDisconnected(void* ptr, void* address)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->deviceDisconnected(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothLocalDevice_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&QBluetoothLocalDevice::error), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&MyQBluetoothLocalDevice::Signal_Error));
}

void QBluetoothLocalDevice_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&QBluetoothLocalDevice::error), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::Error)>(&MyQBluetoothLocalDevice::Signal_Error));
}

void QBluetoothLocalDevice_Error(void* ptr, long long error)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->error(static_cast<QBluetoothLocalDevice::Error>(error));
}

void QBluetoothLocalDevice_ConnectHostModeStateChanged(void* ptr)
{
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&QBluetoothLocalDevice::hostModeStateChanged), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&MyQBluetoothLocalDevice::Signal_HostModeStateChanged));
}

void QBluetoothLocalDevice_DisconnectHostModeStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&QBluetoothLocalDevice::hostModeStateChanged), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(QBluetoothLocalDevice::HostMode)>(&MyQBluetoothLocalDevice::Signal_HostModeStateChanged));
}

void QBluetoothLocalDevice_HostModeStateChanged(void* ptr, long long state)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->hostModeStateChanged(static_cast<QBluetoothLocalDevice::HostMode>(state));
}

void QBluetoothLocalDevice_ConnectPairingDisplayConfirmation(void* ptr)
{
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &, QString)>(&QBluetoothLocalDevice::pairingDisplayConfirmation), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &, QString)>(&MyQBluetoothLocalDevice::Signal_PairingDisplayConfirmation));
}

void QBluetoothLocalDevice_DisconnectPairingDisplayConfirmation(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &, QString)>(&QBluetoothLocalDevice::pairingDisplayConfirmation), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &, QString)>(&MyQBluetoothLocalDevice::Signal_PairingDisplayConfirmation));
}

void QBluetoothLocalDevice_PairingDisplayConfirmation(void* ptr, void* address, char* pin)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->pairingDisplayConfirmation(*static_cast<QBluetoothAddress*>(address), QString(pin));
}

void QBluetoothLocalDevice_ConnectPairingDisplayPinCode(void* ptr)
{
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &, QString)>(&QBluetoothLocalDevice::pairingDisplayPinCode), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &, QString)>(&MyQBluetoothLocalDevice::Signal_PairingDisplayPinCode));
}

void QBluetoothLocalDevice_DisconnectPairingDisplayPinCode(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &, QString)>(&QBluetoothLocalDevice::pairingDisplayPinCode), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &, QString)>(&MyQBluetoothLocalDevice::Signal_PairingDisplayPinCode));
}

void QBluetoothLocalDevice_PairingDisplayPinCode(void* ptr, void* address, char* pin)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->pairingDisplayPinCode(*static_cast<QBluetoothAddress*>(address), QString(pin));
}

void QBluetoothLocalDevice_ConnectPairingFinished(void* ptr)
{
	QObject::connect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &, QBluetoothLocalDevice::Pairing)>(&QBluetoothLocalDevice::pairingFinished), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &, QBluetoothLocalDevice::Pairing)>(&MyQBluetoothLocalDevice::Signal_PairingFinished));
}

void QBluetoothLocalDevice_DisconnectPairingFinished(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothLocalDevice*>(ptr), static_cast<void (QBluetoothLocalDevice::*)(const QBluetoothAddress &, QBluetoothLocalDevice::Pairing)>(&QBluetoothLocalDevice::pairingFinished), static_cast<MyQBluetoothLocalDevice*>(ptr), static_cast<void (MyQBluetoothLocalDevice::*)(const QBluetoothAddress &, QBluetoothLocalDevice::Pairing)>(&MyQBluetoothLocalDevice::Signal_PairingFinished));
}

void QBluetoothLocalDevice_PairingFinished(void* ptr, void* address, long long pairing)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->pairingFinished(*static_cast<QBluetoothAddress*>(address), static_cast<QBluetoothLocalDevice::Pairing>(pairing));
}

char QBluetoothLocalDevice_IsValid(void* ptr)
{
	return static_cast<QBluetoothLocalDevice*>(ptr)->isValid();
}

void QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(void* ptr)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->~QBluetoothLocalDevice();
}

void QBluetoothLocalDevice_DestroyQBluetoothLocalDeviceDefault(void* ptr)
{

}

void* QBluetoothLocalDevice_NewQBluetoothLocalDevice(void* parent)
{
	return new MyQBluetoothLocalDevice(static_cast<QObject*>(parent));
}

void* QBluetoothLocalDevice_NewQBluetoothLocalDevice2(void* address, void* parent)
{
	return new MyQBluetoothLocalDevice(*static_cast<QBluetoothAddress*>(address), static_cast<QObject*>(parent));
}

void* QBluetoothLocalDevice_Address(void* ptr)
{
	return new QBluetoothAddress(static_cast<QBluetoothLocalDevice*>(ptr)->address());
}

struct QtBluetooth_PackedList QBluetoothLocalDevice_QBluetoothLocalDevice_AllDevices()
{
	return ({ QList<QBluetoothHostInfo>* tmpValue = new QList<QBluetoothHostInfo>(QBluetoothLocalDevice::allDevices()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtBluetooth_PackedList QBluetoothLocalDevice_ConnectedDevices(void* ptr)
{
	return ({ QList<QBluetoothAddress>* tmpValue = new QList<QBluetoothAddress>(static_cast<QBluetoothLocalDevice*>(ptr)->connectedDevices()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

long long QBluetoothLocalDevice_HostMode(void* ptr)
{
	return static_cast<QBluetoothLocalDevice*>(ptr)->hostMode();
}

struct QtBluetooth_PackedString QBluetoothLocalDevice_Name(void* ptr)
{
	return ({ QByteArray td05b9d = static_cast<QBluetoothLocalDevice*>(ptr)->name().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(td05b9d.prepend("WHITESPACE").constData()+10), td05b9d.size()-10 }; });
}

void QBluetoothLocalDevice_PairingConfirmation(void* ptr, char accept)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothLocalDevice*>(ptr), "pairingConfirmation", Q_ARG(bool, accept != 0));
}

long long QBluetoothLocalDevice_PairingStatus(void* ptr, void* address)
{
	return static_cast<QBluetoothLocalDevice*>(ptr)->pairingStatus(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothLocalDevice_PowerOn(void* ptr)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->powerOn();
}

void QBluetoothLocalDevice_RequestPairing(void* ptr, void* address, long long pairing)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->requestPairing(*static_cast<QBluetoothAddress*>(address), static_cast<QBluetoothLocalDevice::Pairing>(pairing));
}

void QBluetoothLocalDevice_SetHostMode(void* ptr, long long mode)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->setHostMode(static_cast<QBluetoothLocalDevice::HostMode>(mode));
}

void* QBluetoothLocalDevice_allDevices_atList(void* ptr, int i)
{
	return new QBluetoothHostInfo(static_cast<QList<QBluetoothHostInfo>*>(ptr)->at(i));
}

void* QBluetoothLocalDevice_connectedDevices_atList(void* ptr, int i)
{
	return new QBluetoothAddress(static_cast<QList<QBluetoothAddress>*>(ptr)->at(i));
}

void QBluetoothLocalDevice_TimerEvent(void* ptr, void* event)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothLocalDevice_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothLocalDevice_ChildEvent(void* ptr, void* event)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothLocalDevice_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothLocalDevice_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothLocalDevice_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothLocalDevice_CustomEvent(void* ptr, void* event)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothLocalDevice_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::customEvent(static_cast<QEvent*>(event));
}

void QBluetoothLocalDevice_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothLocalDevice*>(ptr), "deleteLater");
}

void QBluetoothLocalDevice_DeleteLaterDefault(void* ptr)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::deleteLater();
}

void QBluetoothLocalDevice_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothLocalDevice_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QBluetoothLocalDevice_Event(void* ptr, void* e)
{
	return static_cast<QBluetoothLocalDevice*>(ptr)->event(static_cast<QEvent*>(e));
}

char QBluetoothLocalDevice_EventDefault(void* ptr, void* e)
{
	return static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::event(static_cast<QEvent*>(e));
}

char QBluetoothLocalDevice_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothLocalDevice*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QBluetoothLocalDevice_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QBluetoothLocalDevice_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothLocalDevice*>(ptr)->metaObject());
}

void* QBluetoothLocalDevice_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothLocalDevice*>(ptr)->QBluetoothLocalDevice::metaObject());
}

class MyQBluetoothServer: public QBluetoothServer
{
public:
	MyQBluetoothServer(QBluetoothServiceInfo::Protocol serverType, QObject *parent) : QBluetoothServer(serverType, parent) {};
	void Signal_Error2(QBluetoothServer::Error error) { callbackQBluetoothServer_Error2(this, error); };
	void Signal_NewConnection() { callbackQBluetoothServer_NewConnection(this); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothServer_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBluetoothServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBluetoothServer_CustomEvent(this, event); };
	void deleteLater() { callbackQBluetoothServer_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBluetoothServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQBluetoothServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBluetoothServer_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBluetoothServer_MetaObject(const_cast<MyQBluetoothServer*>(this))); };
};

void* QBluetoothServer_NewQBluetoothServer(long long serverType, void* parent)
{
	return new MyQBluetoothServer(static_cast<QBluetoothServiceInfo::Protocol>(serverType), static_cast<QObject*>(parent));
}

void QBluetoothServer_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)(QBluetoothServer::Error)>(&QBluetoothServer::error), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)(QBluetoothServer::Error)>(&MyQBluetoothServer::Signal_Error2));
}

void QBluetoothServer_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)(QBluetoothServer::Error)>(&QBluetoothServer::error), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)(QBluetoothServer::Error)>(&MyQBluetoothServer::Signal_Error2));
}

void QBluetoothServer_Error2(void* ptr, long long error)
{
	static_cast<QBluetoothServer*>(ptr)->error(static_cast<QBluetoothServer::Error>(error));
}

void QBluetoothServer_ConnectNewConnection(void* ptr)
{
	QObject::connect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));
}

void QBluetoothServer_DisconnectNewConnection(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothServer*>(ptr), static_cast<void (QBluetoothServer::*)()>(&QBluetoothServer::newConnection), static_cast<MyQBluetoothServer*>(ptr), static_cast<void (MyQBluetoothServer::*)()>(&MyQBluetoothServer::Signal_NewConnection));
}

void QBluetoothServer_NewConnection(void* ptr)
{
	static_cast<QBluetoothServer*>(ptr)->newConnection();
}

long long QBluetoothServer_Error(void* ptr)
{
	return static_cast<QBluetoothServer*>(ptr)->error();
}

char QBluetoothServer_IsListening(void* ptr)
{
	return static_cast<QBluetoothServer*>(ptr)->isListening();
}

void* QBluetoothServer_Listen2(void* ptr, void* uuid, char* serviceName)
{
	return new QBluetoothServiceInfo(static_cast<QBluetoothServer*>(ptr)->listen(*static_cast<QBluetoothUuid*>(uuid), QString(serviceName)));
}

int QBluetoothServer_MaxPendingConnections(void* ptr)
{
	return static_cast<QBluetoothServer*>(ptr)->maxPendingConnections();
}

long long QBluetoothServer_ServerType(void* ptr)
{
	return static_cast<QBluetoothServer*>(ptr)->serverType();
}

void QBluetoothServer_DestroyQBluetoothServer(void* ptr)
{
	static_cast<QBluetoothServer*>(ptr)->~QBluetoothServer();
}

void QBluetoothServer_Close(void* ptr)
{
	static_cast<QBluetoothServer*>(ptr)->close();
}

char QBluetoothServer_HasPendingConnections(void* ptr)
{
	return static_cast<QBluetoothServer*>(ptr)->hasPendingConnections();
}

char QBluetoothServer_Listen(void* ptr, void* address, unsigned short port)
{
	return static_cast<QBluetoothServer*>(ptr)->listen(*static_cast<QBluetoothAddress*>(address), port);
}

void* QBluetoothServer_NextPendingConnection(void* ptr)
{
	return static_cast<QBluetoothServer*>(ptr)->nextPendingConnection();
}

void* QBluetoothServer_ServerAddress(void* ptr)
{
	return new QBluetoothAddress(static_cast<QBluetoothServer*>(ptr)->serverAddress());
}

unsigned short QBluetoothServer_ServerPort(void* ptr)
{
	return static_cast<QBluetoothServer*>(ptr)->serverPort();
}

void QBluetoothServer_SetMaxPendingConnections(void* ptr, int numConnections)
{
	static_cast<QBluetoothServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QBluetoothServer_TimerEvent(void* ptr, void* event)
{
	static_cast<QBluetoothServer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothServer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothServer_ChildEvent(void* ptr, void* event)
{
	static_cast<QBluetoothServer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothServer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothServer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothServer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothServer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothServer_CustomEvent(void* ptr, void* event)
{
	static_cast<QBluetoothServer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothServer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::customEvent(static_cast<QEvent*>(event));
}

void QBluetoothServer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothServer*>(ptr), "deleteLater");
}

void QBluetoothServer_DeleteLaterDefault(void* ptr)
{
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::deleteLater();
}

void QBluetoothServer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothServer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QBluetoothServer_Event(void* ptr, void* e)
{
	return static_cast<QBluetoothServer*>(ptr)->event(static_cast<QEvent*>(e));
}

char QBluetoothServer_EventDefault(void* ptr, void* e)
{
	return static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::event(static_cast<QEvent*>(e));
}

char QBluetoothServer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothServer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QBluetoothServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QBluetoothServer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothServer*>(ptr)->metaObject());
}

void* QBluetoothServer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothServer*>(ptr)->QBluetoothServer::metaObject());
}

class MyQBluetoothServiceDiscoveryAgent: public QBluetoothServiceDiscoveryAgent
{
public:
	MyQBluetoothServiceDiscoveryAgent(QObject *parent) : QBluetoothServiceDiscoveryAgent(parent) {};
	MyQBluetoothServiceDiscoveryAgent(const QBluetoothAddress &deviceAdapter, QObject *parent) : QBluetoothServiceDiscoveryAgent(deviceAdapter, parent) {};
	void Signal_Canceled() { callbackQBluetoothServiceDiscoveryAgent_Canceled(this); };
	void Signal_Error2(QBluetoothServiceDiscoveryAgent::Error error) { callbackQBluetoothServiceDiscoveryAgent_Error2(this, error); };
	void Signal_Finished() { callbackQBluetoothServiceDiscoveryAgent_Finished(this); };
	void Signal_ServiceDiscovered(const QBluetoothServiceInfo & info) { callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered(this, const_cast<QBluetoothServiceInfo*>(&info)); };
	void clear() { callbackQBluetoothServiceDiscoveryAgent_Clear(this); };
	void start(QBluetoothServiceDiscoveryAgent::DiscoveryMode mode) { callbackQBluetoothServiceDiscoveryAgent_Start(this, mode); };
	void stop() { callbackQBluetoothServiceDiscoveryAgent_Stop(this); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothServiceDiscoveryAgent_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothServiceDiscoveryAgent_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBluetoothServiceDiscoveryAgent_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBluetoothServiceDiscoveryAgent_CustomEvent(this, event); };
	void deleteLater() { callbackQBluetoothServiceDiscoveryAgent_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQBluetoothServiceDiscoveryAgent_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBluetoothServiceDiscoveryAgent_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBluetoothServiceDiscoveryAgent_MetaObject(const_cast<MyQBluetoothServiceDiscoveryAgent*>(this))); };
};

void QBluetoothServiceDiscoveryAgent_ConnectCanceled(void* ptr)
{
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));
}

void QBluetoothServiceDiscoveryAgent_DisconnectCanceled(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::canceled), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Canceled));
}

void QBluetoothServiceDiscoveryAgent_Canceled(void* ptr)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->canceled();
}

void QBluetoothServiceDiscoveryAgent_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)(QBluetoothServiceDiscoveryAgent::Error)>(&QBluetoothServiceDiscoveryAgent::error), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)(QBluetoothServiceDiscoveryAgent::Error)>(&MyQBluetoothServiceDiscoveryAgent::Signal_Error2));
}

void QBluetoothServiceDiscoveryAgent_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)(QBluetoothServiceDiscoveryAgent::Error)>(&QBluetoothServiceDiscoveryAgent::error), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)(QBluetoothServiceDiscoveryAgent::Error)>(&MyQBluetoothServiceDiscoveryAgent::Signal_Error2));
}

void QBluetoothServiceDiscoveryAgent_Error2(void* ptr, long long error)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->error(static_cast<QBluetoothServiceDiscoveryAgent::Error>(error));
}

void QBluetoothServiceDiscoveryAgent_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));
}

void QBluetoothServiceDiscoveryAgent_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)()>(&QBluetoothServiceDiscoveryAgent::finished), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)()>(&MyQBluetoothServiceDiscoveryAgent::Signal_Finished));
}

void QBluetoothServiceDiscoveryAgent_Finished(void* ptr)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->finished();
}

void QBluetoothServiceDiscoveryAgent_ConnectServiceDiscovered(void* ptr)
{
	QObject::connect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)(const QBluetoothServiceInfo &)>(&QBluetoothServiceDiscoveryAgent::serviceDiscovered), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)(const QBluetoothServiceInfo &)>(&MyQBluetoothServiceDiscoveryAgent::Signal_ServiceDiscovered));
}

void QBluetoothServiceDiscoveryAgent_DisconnectServiceDiscovered(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (QBluetoothServiceDiscoveryAgent::*)(const QBluetoothServiceInfo &)>(&QBluetoothServiceDiscoveryAgent::serviceDiscovered), static_cast<MyQBluetoothServiceDiscoveryAgent*>(ptr), static_cast<void (MyQBluetoothServiceDiscoveryAgent::*)(const QBluetoothServiceInfo &)>(&MyQBluetoothServiceDiscoveryAgent::Signal_ServiceDiscovered));
}

void QBluetoothServiceDiscoveryAgent_ServiceDiscovered(void* ptr, void* info)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->serviceDiscovered(*static_cast<QBluetoothServiceInfo*>(info));
}

void* QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(void* parent)
{
	return new MyQBluetoothServiceDiscoveryAgent(static_cast<QObject*>(parent));
}

void* QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(void* deviceAdapter, void* parent)
{
	return new MyQBluetoothServiceDiscoveryAgent(*static_cast<QBluetoothAddress*>(deviceAdapter), static_cast<QObject*>(parent));
}

void QBluetoothServiceDiscoveryAgent_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "clear");
}

struct QtBluetooth_PackedList QBluetoothServiceDiscoveryAgent_DiscoveredServices(void* ptr)
{
	return ({ QList<QBluetoothServiceInfo>* tmpValue = new QList<QBluetoothServiceInfo>(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->discoveredServices()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

long long QBluetoothServiceDiscoveryAgent_Error(void* ptr)
{
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->error();
}

struct QtBluetooth_PackedString QBluetoothServiceDiscoveryAgent_ErrorString(void* ptr)
{
	return ({ QByteArray t581ebe = static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->errorString().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t581ebe.prepend("WHITESPACE").constData()+10), t581ebe.size()-10 }; });
}

char QBluetoothServiceDiscoveryAgent_IsActive(void* ptr)
{
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->isActive();
}

void* QBluetoothServiceDiscoveryAgent_RemoteAddress(void* ptr)
{
	return new QBluetoothAddress(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->remoteAddress());
}

char QBluetoothServiceDiscoveryAgent_SetRemoteAddress(void* ptr, void* address)
{
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->setRemoteAddress(*static_cast<QBluetoothAddress*>(address));
}

void QBluetoothServiceDiscoveryAgent_SetUuidFilter2(void* ptr, void* uuid)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->setUuidFilter(*static_cast<QBluetoothUuid*>(uuid));
}

void QBluetoothServiceDiscoveryAgent_Start(void* ptr, long long mode)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "start", Q_ARG(QBluetoothServiceDiscoveryAgent::DiscoveryMode, static_cast<QBluetoothServiceDiscoveryAgent::DiscoveryMode>(mode)));
}

void QBluetoothServiceDiscoveryAgent_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "stop");
}

struct QtBluetooth_PackedList QBluetoothServiceDiscoveryAgent_UuidFilter(void* ptr)
{
	return ({ QList<QBluetoothUuid>* tmpValue = new QList<QBluetoothUuid>(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->uuidFilter()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

void QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(void* ptr)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->~QBluetoothServiceDiscoveryAgent();
}

void* QBluetoothServiceDiscoveryAgent_discoveredServices_atList(void* ptr, int i)
{
	return new QBluetoothServiceInfo(static_cast<QList<QBluetoothServiceInfo>*>(ptr)->at(i));
}

void* QBluetoothServiceDiscoveryAgent_uuidFilter_atList(void* ptr, int i)
{
	return new QBluetoothUuid(static_cast<QList<QBluetoothUuid>*>(ptr)->at(i));
}

void QBluetoothServiceDiscoveryAgent_TimerEvent(void* ptr, void* event)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_ChildEvent(void* ptr, void* event)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothServiceDiscoveryAgent_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothServiceDiscoveryAgent_CustomEvent(void* ptr, void* event)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::customEvent(static_cast<QEvent*>(event));
}

void QBluetoothServiceDiscoveryAgent_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr), "deleteLater");
}

void QBluetoothServiceDiscoveryAgent_DeleteLaterDefault(void* ptr)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::deleteLater();
}

void QBluetoothServiceDiscoveryAgent_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothServiceDiscoveryAgent_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QBluetoothServiceDiscoveryAgent_Event(void* ptr, void* e)
{
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->event(static_cast<QEvent*>(e));
}

char QBluetoothServiceDiscoveryAgent_EventDefault(void* ptr, void* e)
{
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::event(static_cast<QEvent*>(e));
}

char QBluetoothServiceDiscoveryAgent_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QBluetoothServiceDiscoveryAgent_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QBluetoothServiceDiscoveryAgent_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->metaObject());
}

void* QBluetoothServiceDiscoveryAgent_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothServiceDiscoveryAgent*>(ptr)->QBluetoothServiceDiscoveryAgent::metaObject());
}

int QBluetoothServiceInfo_ServiceName_Type()
{
	return QBluetoothServiceInfo::ServiceName;
}

int QBluetoothServiceInfo_ServiceDescription_Type()
{
	return QBluetoothServiceInfo::ServiceDescription;
}

int QBluetoothServiceInfo_ServiceProvider_Type()
{
	return QBluetoothServiceInfo::ServiceProvider;
}

struct QtBluetooth_PackedString QBluetoothServiceInfo_ServiceAvailability(void* ptr)
{
	return ({ quint8 pret18d8cb = static_cast<QBluetoothServiceInfo*>(ptr)->serviceAvailability(); char* t18d8cb = static_cast<char*>(static_cast<void*>(&pret18d8cb)); QtBluetooth_PackedString { t18d8cb, -1 }; });
}

struct QtBluetooth_PackedString QBluetoothServiceInfo_ServiceDescription(void* ptr)
{
	return ({ QByteArray t928c16 = static_cast<QBluetoothServiceInfo*>(ptr)->serviceDescription().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t928c16.prepend("WHITESPACE").constData()+10), t928c16.size()-10 }; });
}

struct QtBluetooth_PackedString QBluetoothServiceInfo_ServiceName(void* ptr)
{
	return ({ QByteArray tf38ab3 = static_cast<QBluetoothServiceInfo*>(ptr)->serviceName().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(tf38ab3.prepend("WHITESPACE").constData()+10), tf38ab3.size()-10 }; });
}

struct QtBluetooth_PackedString QBluetoothServiceInfo_ServiceProvider(void* ptr)
{
	return ({ QByteArray tae7568 = static_cast<QBluetoothServiceInfo*>(ptr)->serviceProvider().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(tae7568.prepend("WHITESPACE").constData()+10), tae7568.size()-10 }; });
}

void* QBluetoothServiceInfo_ServiceUuid(void* ptr)
{
	return new QBluetoothUuid(static_cast<QBluetoothServiceInfo*>(ptr)->serviceUuid());
}

void QBluetoothServiceInfo_SetAttribute2(void* ptr, unsigned short attributeId, void* value)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->setAttribute(attributeId, *static_cast<QBluetoothUuid*>(value));
}

void QBluetoothServiceInfo_SetServiceAvailability(void* ptr, char* availability)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceAvailability(*static_cast<quint8*>(static_cast<void*>(availability)));
}

void QBluetoothServiceInfo_SetServiceDescription(void* ptr, char* description)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceDescription(QString(description));
}

void QBluetoothServiceInfo_SetServiceName(void* ptr, char* name)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceName(QString(name));
}

void QBluetoothServiceInfo_SetServiceProvider(void* ptr, char* provider)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceProvider(QString(provider));
}

void QBluetoothServiceInfo_SetServiceUuid(void* ptr, void* uuid)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->setServiceUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void* QBluetoothServiceInfo_NewQBluetoothServiceInfo()
{
	return new QBluetoothServiceInfo();
}

void* QBluetoothServiceInfo_NewQBluetoothServiceInfo2(void* other)
{
	return new QBluetoothServiceInfo(*static_cast<QBluetoothServiceInfo*>(other));
}

void* QBluetoothServiceInfo_Attribute(void* ptr, unsigned short attributeId)
{
	return new QVariant(static_cast<QBluetoothServiceInfo*>(ptr)->attribute(attributeId));
}

char QBluetoothServiceInfo_Contains(void* ptr, unsigned short attributeId)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->contains(attributeId);
}

void* QBluetoothServiceInfo_Device(void* ptr)
{
	return new QBluetoothDeviceInfo(static_cast<QBluetoothServiceInfo*>(ptr)->device());
}

char QBluetoothServiceInfo_IsComplete(void* ptr)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->isComplete();
}

char QBluetoothServiceInfo_IsRegistered(void* ptr)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->isRegistered();
}

char QBluetoothServiceInfo_IsValid(void* ptr)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->isValid();
}

int QBluetoothServiceInfo_ProtocolServiceMultiplexer(void* ptr)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->protocolServiceMultiplexer();
}

char QBluetoothServiceInfo_RegisterService(void* ptr, void* localAdapter)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->registerService(*static_cast<QBluetoothAddress*>(localAdapter));
}

void QBluetoothServiceInfo_RemoveAttribute(void* ptr, unsigned short attributeId)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->removeAttribute(attributeId);
}

int QBluetoothServiceInfo_ServerChannel(void* ptr)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->serverChannel();
}

struct QtBluetooth_PackedList QBluetoothServiceInfo_ServiceClassUuids(void* ptr)
{
	return ({ QList<QBluetoothUuid>* tmpValue = new QList<QBluetoothUuid>(static_cast<QBluetoothServiceInfo*>(ptr)->serviceClassUuids()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

void QBluetoothServiceInfo_SetAttribute(void* ptr, unsigned short attributeId, void* value)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->setAttribute(attributeId, *static_cast<QVariant*>(value));
}

void QBluetoothServiceInfo_SetDevice(void* ptr, void* device)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->setDevice(*static_cast<QBluetoothDeviceInfo*>(device));
}

long long QBluetoothServiceInfo_SocketProtocol(void* ptr)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->socketProtocol();
}

char QBluetoothServiceInfo_UnregisterService(void* ptr)
{
	return static_cast<QBluetoothServiceInfo*>(ptr)->unregisterService();
}

void QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(void* ptr)
{
	static_cast<QBluetoothServiceInfo*>(ptr)->~QBluetoothServiceInfo();
}

void* QBluetoothServiceInfo_serviceClassUuids_atList(void* ptr, int i)
{
	return new QBluetoothUuid(static_cast<QList<QBluetoothUuid>*>(ptr)->at(i));
}

class MyQBluetoothSocket: public QBluetoothSocket
{
public:
	MyQBluetoothSocket(QBluetoothServiceInfo::Protocol socketType, QObject *parent) : QBluetoothSocket(socketType, parent) {};
	MyQBluetoothSocket(QObject *parent) : QBluetoothSocket(parent) {};
	void Signal_Connected() { callbackQBluetoothSocket_Connected(this); };
	void Signal_Disconnected() { callbackQBluetoothSocket_Disconnected(this); };
	void Signal_Error2(QBluetoothSocket::SocketError error) { callbackQBluetoothSocket_Error2(this, error); };
	void Signal_StateChanged(QBluetoothSocket::SocketState state) { callbackQBluetoothSocket_StateChanged(this, state); };
	qint64 bytesAvailable() const { return callbackQBluetoothSocket_BytesAvailable(const_cast<MyQBluetoothSocket*>(this)); };
	qint64 bytesToWrite() const { return callbackQBluetoothSocket_BytesToWrite(const_cast<MyQBluetoothSocket*>(this)); };
	bool canReadLine() const { return callbackQBluetoothSocket_CanReadLine(const_cast<MyQBluetoothSocket*>(this)) != 0; };
	void close() { callbackQBluetoothSocket_Close(this); };
	qint64 writeData(const char * data, qint64 maxSize) { QtBluetooth_PackedString dataPacked = { const_cast<char*>(data), maxSize };return callbackQBluetoothSocket_WriteData(this, dataPacked, maxSize); };
	 ~MyQBluetoothSocket() { callbackQBluetoothSocket_DestroyQBluetoothSocket(this); };
	bool atEnd() const { return callbackQBluetoothSocket_AtEnd(const_cast<MyQBluetoothSocket*>(this)) != 0; };
	bool open(QIODevice::OpenMode mode) { return callbackQBluetoothSocket_Open(this, mode) != 0; };
	qint64 pos() const { return callbackQBluetoothSocket_Pos(const_cast<MyQBluetoothSocket*>(this)); };
	qint64 readLineData(char * data, qint64 maxSize) { QtBluetooth_PackedString dataPacked = { data, maxSize };return callbackQBluetoothSocket_ReadLineData(this, dataPacked, maxSize); };
	bool reset() { return callbackQBluetoothSocket_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQBluetoothSocket_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQBluetoothSocket_Size(const_cast<MyQBluetoothSocket*>(this)); };
	bool waitForBytesWritten(int msecs) { return callbackQBluetoothSocket_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQBluetoothSocket_WaitForReadyRead(this, msecs) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothSocket_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothSocket_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBluetoothSocket_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBluetoothSocket_CustomEvent(this, event); };
	void deleteLater() { callbackQBluetoothSocket_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBluetoothSocket_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQBluetoothSocket_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBluetoothSocket_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBluetoothSocket_MetaObject(const_cast<MyQBluetoothSocket*>(this))); };
};

void QBluetoothSocket_ConnectConnected(void* ptr)
{
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));
}

void QBluetoothSocket_DisconnectConnected(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::connected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Connected));
}

void QBluetoothSocket_Connected(void* ptr)
{
	static_cast<QBluetoothSocket*>(ptr)->connected();
}

void QBluetoothSocket_ConnectDisconnected(void* ptr)
{
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));
}

void QBluetoothSocket_DisconnectDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)()>(&QBluetoothSocket::disconnected), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)()>(&MyQBluetoothSocket::Signal_Disconnected));
}

void QBluetoothSocket_Disconnected(void* ptr)
{
	static_cast<QBluetoothSocket*>(ptr)->disconnected();
}

void QBluetoothSocket_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketError)>(&QBluetoothSocket::error), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketError)>(&MyQBluetoothSocket::Signal_Error2));
}

void QBluetoothSocket_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketError)>(&QBluetoothSocket::error), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketError)>(&MyQBluetoothSocket::Signal_Error2));
}

void QBluetoothSocket_Error2(void* ptr, long long error)
{
	static_cast<QBluetoothSocket*>(ptr)->error(static_cast<QBluetoothSocket::SocketError>(error));
}

void QBluetoothSocket_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));
}

void QBluetoothSocket_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothSocket*>(ptr), static_cast<void (QBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&QBluetoothSocket::stateChanged), static_cast<MyQBluetoothSocket*>(ptr), static_cast<void (MyQBluetoothSocket::*)(QBluetoothSocket::SocketState)>(&MyQBluetoothSocket::Signal_StateChanged));
}

void QBluetoothSocket_StateChanged(void* ptr, long long state)
{
	static_cast<QBluetoothSocket*>(ptr)->stateChanged(static_cast<QBluetoothSocket::SocketState>(state));
}

void* QBluetoothSocket_NewQBluetoothSocket(long long socketType, void* parent)
{
	return new MyQBluetoothSocket(static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QObject*>(parent));
}

void* QBluetoothSocket_NewQBluetoothSocket2(void* parent)
{
	return new MyQBluetoothSocket(static_cast<QObject*>(parent));
}

void QBluetoothSocket_Abort(void* ptr)
{
	static_cast<QBluetoothSocket*>(ptr)->abort();
}

long long QBluetoothSocket_BytesAvailable(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->bytesAvailable();
}

long long QBluetoothSocket_BytesAvailableDefault(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::bytesAvailable();
}

long long QBluetoothSocket_BytesToWrite(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->bytesToWrite();
}

long long QBluetoothSocket_BytesToWriteDefault(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::bytesToWrite();
}

char QBluetoothSocket_CanReadLine(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->canReadLine();
}

char QBluetoothSocket_CanReadLineDefault(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::canReadLine();
}

void QBluetoothSocket_Close(void* ptr)
{
	static_cast<QBluetoothSocket*>(ptr)->close();
}

void QBluetoothSocket_CloseDefault(void* ptr)
{
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::close();
}

void QBluetoothSocket_ConnectToService2(void* ptr, void* address, void* uuid, long long openMode)
{
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothAddress*>(address), *static_cast<QBluetoothUuid*>(uuid), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_ConnectToService3(void* ptr, void* address, unsigned short port, long long openMode)
{
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothAddress*>(address), port, static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_ConnectToService(void* ptr, void* service, long long openMode)
{
	static_cast<QBluetoothSocket*>(ptr)->connectToService(*static_cast<QBluetoothServiceInfo*>(service), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_DisconnectFromService(void* ptr)
{
	static_cast<QBluetoothSocket*>(ptr)->disconnectFromService();
}

void QBluetoothSocket_DoDeviceDiscovery(void* ptr, void* service, long long openMode)
{
	static_cast<QBluetoothSocket*>(ptr)->doDeviceDiscovery(*static_cast<QBluetoothServiceInfo*>(service), static_cast<QIODevice::OpenModeFlag>(openMode));
}

long long QBluetoothSocket_Error(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->error();
}

struct QtBluetooth_PackedString QBluetoothSocket_ErrorString(void* ptr)
{
	return ({ QByteArray t6109ad = static_cast<QBluetoothSocket*>(ptr)->errorString().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t6109ad.prepend("WHITESPACE").constData()+10), t6109ad.size()-10 }; });
}

char QBluetoothSocket_IsSequential(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->isSequential();
}

void* QBluetoothSocket_LocalAddress(void* ptr)
{
	return new QBluetoothAddress(static_cast<QBluetoothSocket*>(ptr)->localAddress());
}

struct QtBluetooth_PackedString QBluetoothSocket_LocalName(void* ptr)
{
	return ({ QByteArray t705949 = static_cast<QBluetoothSocket*>(ptr)->localName().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t705949.prepend("WHITESPACE").constData()+10), t705949.size()-10 }; });
}

unsigned short QBluetoothSocket_LocalPort(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->localPort();
}

void* QBluetoothSocket_PeerAddress(void* ptr)
{
	return new QBluetoothAddress(static_cast<QBluetoothSocket*>(ptr)->peerAddress());
}

struct QtBluetooth_PackedString QBluetoothSocket_PeerName(void* ptr)
{
	return ({ QByteArray te958b4 = static_cast<QBluetoothSocket*>(ptr)->peerName().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(te958b4.prepend("WHITESPACE").constData()+10), te958b4.size()-10 }; });
}

unsigned short QBluetoothSocket_PeerPort(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->peerPort();
}

char QBluetoothSocket_SetSocketDescriptor(void* ptr, int socketDescriptor, long long socketType, long long socketState, long long openMode)
{
	return static_cast<QBluetoothSocket*>(ptr)->setSocketDescriptor(socketDescriptor, static_cast<QBluetoothServiceInfo::Protocol>(socketType), static_cast<QBluetoothSocket::SocketState>(socketState), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void QBluetoothSocket_SetSocketError(void* ptr, long long error_)
{
	static_cast<QBluetoothSocket*>(ptr)->setSocketError(static_cast<QBluetoothSocket::SocketError>(error_));
}

void QBluetoothSocket_SetSocketState(void* ptr, long long state)
{
	static_cast<QBluetoothSocket*>(ptr)->setSocketState(static_cast<QBluetoothSocket::SocketState>(state));
}

int QBluetoothSocket_SocketDescriptor(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->socketDescriptor();
}

long long QBluetoothSocket_SocketType(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->socketType();
}

long long QBluetoothSocket_State(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->state();
}

long long QBluetoothSocket_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QBluetoothSocket*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

long long QBluetoothSocket_WriteDataDefault(void* ptr, char* data, long long maxSize)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::writeData(const_cast<const char*>(data), maxSize);
}

void QBluetoothSocket_DestroyQBluetoothSocket(void* ptr)
{
	static_cast<QBluetoothSocket*>(ptr)->~QBluetoothSocket();
}

void QBluetoothSocket_DestroyQBluetoothSocketDefault(void* ptr)
{

}

char QBluetoothSocket_AtEnd(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->atEnd();
}

char QBluetoothSocket_AtEndDefault(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::atEnd();
}

char QBluetoothSocket_Open(void* ptr, long long mode)
{
	return static_cast<QBluetoothSocket*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QBluetoothSocket_OpenDefault(void* ptr, long long mode)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QBluetoothSocket_Pos(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->pos();
}

long long QBluetoothSocket_PosDefault(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::pos();
}

long long QBluetoothSocket_ReadLineData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QBluetoothSocket*>(ptr)->readLineData(data, maxSize);
}

long long QBluetoothSocket_ReadLineDataDefault(void* ptr, char* data, long long maxSize)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::readLineData(data, maxSize);
}

char QBluetoothSocket_Reset(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->reset();
}

char QBluetoothSocket_ResetDefault(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::reset();
}

char QBluetoothSocket_Seek(void* ptr, long long pos)
{
	return static_cast<QBluetoothSocket*>(ptr)->seek(pos);
}

char QBluetoothSocket_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::seek(pos);
}

long long QBluetoothSocket_Size(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->size();
}

long long QBluetoothSocket_SizeDefault(void* ptr)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::size();
}

char QBluetoothSocket_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QBluetoothSocket*>(ptr)->waitForBytesWritten(msecs);
}

char QBluetoothSocket_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::waitForBytesWritten(msecs);
}

char QBluetoothSocket_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QBluetoothSocket*>(ptr)->waitForReadyRead(msecs);
}

char QBluetoothSocket_WaitForReadyReadDefault(void* ptr, int msecs)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::waitForReadyRead(msecs);
}

void QBluetoothSocket_TimerEvent(void* ptr, void* event)
{
	static_cast<QBluetoothSocket*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothSocket_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothSocket_ChildEvent(void* ptr, void* event)
{
	static_cast<QBluetoothSocket*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothSocket_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothSocket_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothSocket*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothSocket_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothSocket_CustomEvent(void* ptr, void* event)
{
	static_cast<QBluetoothSocket*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothSocket_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::customEvent(static_cast<QEvent*>(event));
}

void QBluetoothSocket_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothSocket*>(ptr), "deleteLater");
}

void QBluetoothSocket_DeleteLaterDefault(void* ptr)
{
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::deleteLater();
}

void QBluetoothSocket_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothSocket*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothSocket_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QBluetoothSocket_Event(void* ptr, void* e)
{
	return static_cast<QBluetoothSocket*>(ptr)->event(static_cast<QEvent*>(e));
}

char QBluetoothSocket_EventDefault(void* ptr, void* e)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::event(static_cast<QEvent*>(e));
}

char QBluetoothSocket_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothSocket*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QBluetoothSocket_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QBluetoothSocket_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothSocket*>(ptr)->metaObject());
}

void* QBluetoothSocket_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothSocket*>(ptr)->QBluetoothSocket::metaObject());
}

class MyQBluetoothTransferManager: public QBluetoothTransferManager
{
public:
	MyQBluetoothTransferManager(QObject *parent) : QBluetoothTransferManager(parent) {};
	void Signal_Finished(QBluetoothTransferReply * reply) { callbackQBluetoothTransferManager_Finished(this, reply); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothTransferManager_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothTransferManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBluetoothTransferManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBluetoothTransferManager_CustomEvent(this, event); };
	void deleteLater() { callbackQBluetoothTransferManager_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBluetoothTransferManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQBluetoothTransferManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBluetoothTransferManager_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBluetoothTransferManager_MetaObject(const_cast<MyQBluetoothTransferManager*>(this))); };
};

void* QBluetoothTransferManager_Put(void* ptr, void* request, void* data)
{
	return static_cast<QBluetoothTransferManager*>(ptr)->put(*static_cast<QBluetoothTransferRequest*>(request), static_cast<QIODevice*>(data));
}

void* QBluetoothTransferManager_NewQBluetoothTransferManager(void* parent)
{
	return new MyQBluetoothTransferManager(static_cast<QObject*>(parent));
}

void QBluetoothTransferManager_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QBluetoothTransferManager*>(ptr), static_cast<void (QBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&QBluetoothTransferManager::finished), static_cast<MyQBluetoothTransferManager*>(ptr), static_cast<void (MyQBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferManager::Signal_Finished));
}

void QBluetoothTransferManager_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothTransferManager*>(ptr), static_cast<void (QBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&QBluetoothTransferManager::finished), static_cast<MyQBluetoothTransferManager*>(ptr), static_cast<void (MyQBluetoothTransferManager::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferManager::Signal_Finished));
}

void QBluetoothTransferManager_Finished(void* ptr, void* reply)
{
	static_cast<QBluetoothTransferManager*>(ptr)->finished(static_cast<QBluetoothTransferReply*>(reply));
}

void QBluetoothTransferManager_DestroyQBluetoothTransferManager(void* ptr)
{
	static_cast<QBluetoothTransferManager*>(ptr)->~QBluetoothTransferManager();
}

void QBluetoothTransferManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QBluetoothTransferManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothTransferManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothTransferManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QBluetoothTransferManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothTransferManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothTransferManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothTransferManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothTransferManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothTransferManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QBluetoothTransferManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothTransferManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::customEvent(static_cast<QEvent*>(event));
}

void QBluetoothTransferManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothTransferManager*>(ptr), "deleteLater");
}

void QBluetoothTransferManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::deleteLater();
}

void QBluetoothTransferManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothTransferManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothTransferManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QBluetoothTransferManager_Event(void* ptr, void* e)
{
	return static_cast<QBluetoothTransferManager*>(ptr)->event(static_cast<QEvent*>(e));
}

char QBluetoothTransferManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::event(static_cast<QEvent*>(e));
}

char QBluetoothTransferManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothTransferManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QBluetoothTransferManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QBluetoothTransferManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothTransferManager*>(ptr)->metaObject());
}

void* QBluetoothTransferManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothTransferManager*>(ptr)->QBluetoothTransferManager::metaObject());
}

class MyQBluetoothTransferReply: public QBluetoothTransferReply
{
public:
	MyQBluetoothTransferReply(QObject *parent) : QBluetoothTransferReply(parent) {};
	void abort() { callbackQBluetoothTransferReply_Abort(this); };
	void Signal_Error2(QBluetoothTransferReply::TransferError errorType) { callbackQBluetoothTransferReply_Error2(this, errorType); };
	TransferError error() const { return static_cast<QBluetoothTransferReply::TransferError>(callbackQBluetoothTransferReply_Error(const_cast<MyQBluetoothTransferReply*>(this))); };
	QString errorString() const { return QString(callbackQBluetoothTransferReply_ErrorString(const_cast<MyQBluetoothTransferReply*>(this))); };
	void Signal_Finished(QBluetoothTransferReply * reply) { callbackQBluetoothTransferReply_Finished(this, reply); };
	bool isFinished() const { return callbackQBluetoothTransferReply_IsFinished(const_cast<MyQBluetoothTransferReply*>(this)) != 0; };
	bool isRunning() const { return callbackQBluetoothTransferReply_IsRunning(const_cast<MyQBluetoothTransferReply*>(this)) != 0; };
	void Signal_TransferProgress(qint64 bytesTransferred, qint64 bytesTotal) { callbackQBluetoothTransferReply_TransferProgress(this, bytesTransferred, bytesTotal); };
	void timerEvent(QTimerEvent * event) { callbackQBluetoothTransferReply_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQBluetoothTransferReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBluetoothTransferReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBluetoothTransferReply_CustomEvent(this, event); };
	void deleteLater() { callbackQBluetoothTransferReply_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBluetoothTransferReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQBluetoothTransferReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBluetoothTransferReply_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBluetoothTransferReply_MetaObject(const_cast<MyQBluetoothTransferReply*>(this))); };
};

void QBluetoothTransferReply_Abort(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothTransferReply*>(ptr), "abort");
}

void QBluetoothTransferReply_SetManager(void* ptr, void* manager)
{
	static_cast<QBluetoothTransferReply*>(ptr)->setManager(static_cast<QBluetoothTransferManager*>(manager));
}

void QBluetoothTransferReply_SetRequest(void* ptr, void* request)
{
	static_cast<QBluetoothTransferReply*>(ptr)->setRequest(*static_cast<QBluetoothTransferRequest*>(request));
}

void* QBluetoothTransferReply_NewQBluetoothTransferReply(void* parent)
{
	return new MyQBluetoothTransferReply(static_cast<QObject*>(parent));
}

void QBluetoothTransferReply_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply::TransferError)>(&QBluetoothTransferReply::error), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply::TransferError)>(&MyQBluetoothTransferReply::Signal_Error2));
}

void QBluetoothTransferReply_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply::TransferError)>(&QBluetoothTransferReply::error), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply::TransferError)>(&MyQBluetoothTransferReply::Signal_Error2));
}

void QBluetoothTransferReply_Error2(void* ptr, long long errorType)
{
	static_cast<QBluetoothTransferReply*>(ptr)->error(static_cast<QBluetoothTransferReply::TransferError>(errorType));
}

long long QBluetoothTransferReply_Error(void* ptr)
{
	return static_cast<QBluetoothTransferReply*>(ptr)->error();
}

struct QtBluetooth_PackedString QBluetoothTransferReply_ErrorString(void* ptr)
{
	return ({ QByteArray te5ca71 = static_cast<QBluetoothTransferReply*>(ptr)->errorString().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(te5ca71.prepend("WHITESPACE").constData()+10), te5ca71.size()-10 }; });
}

void QBluetoothTransferReply_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&QBluetoothTransferReply::finished), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferReply::Signal_Finished));
}

void QBluetoothTransferReply_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&QBluetoothTransferReply::finished), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(QBluetoothTransferReply *)>(&MyQBluetoothTransferReply::Signal_Finished));
}

void QBluetoothTransferReply_Finished(void* ptr, void* reply)
{
	static_cast<QBluetoothTransferReply*>(ptr)->finished(static_cast<QBluetoothTransferReply*>(reply));
}

char QBluetoothTransferReply_IsFinished(void* ptr)
{
	return static_cast<QBluetoothTransferReply*>(ptr)->isFinished();
}

char QBluetoothTransferReply_IsRunning(void* ptr)
{
	return static_cast<QBluetoothTransferReply*>(ptr)->isRunning();
}

void* QBluetoothTransferReply_Manager(void* ptr)
{
	return static_cast<QBluetoothTransferReply*>(ptr)->manager();
}

void* QBluetoothTransferReply_Request(void* ptr)
{
	return new QBluetoothTransferRequest(static_cast<QBluetoothTransferReply*>(ptr)->request());
}

void QBluetoothTransferReply_ConnectTransferProgress(void* ptr)
{
	QObject::connect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(qint64, qint64)>(&QBluetoothTransferReply::transferProgress), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(qint64, qint64)>(&MyQBluetoothTransferReply::Signal_TransferProgress));
}

void QBluetoothTransferReply_DisconnectTransferProgress(void* ptr)
{
	QObject::disconnect(static_cast<QBluetoothTransferReply*>(ptr), static_cast<void (QBluetoothTransferReply::*)(qint64, qint64)>(&QBluetoothTransferReply::transferProgress), static_cast<MyQBluetoothTransferReply*>(ptr), static_cast<void (MyQBluetoothTransferReply::*)(qint64, qint64)>(&MyQBluetoothTransferReply::Signal_TransferProgress));
}

void QBluetoothTransferReply_TransferProgress(void* ptr, long long bytesTransferred, long long bytesTotal)
{
	static_cast<QBluetoothTransferReply*>(ptr)->transferProgress(bytesTransferred, bytesTotal);
}

void QBluetoothTransferReply_DestroyQBluetoothTransferReply(void* ptr)
{
	static_cast<QBluetoothTransferReply*>(ptr)->~QBluetoothTransferReply();
}

void QBluetoothTransferReply_TimerEvent(void* ptr, void* event)
{
	static_cast<QBluetoothTransferReply*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothTransferReply_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void QBluetoothTransferReply_ChildEvent(void* ptr, void* event)
{
	static_cast<QBluetoothTransferReply*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothTransferReply_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::childEvent(static_cast<QChildEvent*>(event));
}

void QBluetoothTransferReply_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothTransferReply*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothTransferReply_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothTransferReply_CustomEvent(void* ptr, void* event)
{
	static_cast<QBluetoothTransferReply*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QBluetoothTransferReply_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::customEvent(static_cast<QEvent*>(event));
}

void QBluetoothTransferReply_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QBluetoothTransferReply*>(ptr), "deleteLater");
}

void QBluetoothTransferReply_DeleteLaterDefault(void* ptr)
{
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::deleteLater();
}

void QBluetoothTransferReply_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QBluetoothTransferReply*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBluetoothTransferReply_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QBluetoothTransferReply_Event(void* ptr, void* e)
{
	return static_cast<QBluetoothTransferReply*>(ptr)->event(static_cast<QEvent*>(e));
}

char QBluetoothTransferReply_EventDefault(void* ptr, void* e)
{
	return static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::event(static_cast<QEvent*>(e));
}

char QBluetoothTransferReply_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothTransferReply*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QBluetoothTransferReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QBluetoothTransferReply_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothTransferReply*>(ptr)->metaObject());
}

void* QBluetoothTransferReply_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QBluetoothTransferReply*>(ptr)->QBluetoothTransferReply::metaObject());
}

void* QBluetoothTransferRequest_NewQBluetoothTransferRequest(void* address)
{
	return new QBluetoothTransferRequest(*static_cast<QBluetoothAddress*>(address));
}

void* QBluetoothTransferRequest_NewQBluetoothTransferRequest2(void* other)
{
	return new QBluetoothTransferRequest(*static_cast<QBluetoothTransferRequest*>(other));
}

void* QBluetoothTransferRequest_Address(void* ptr)
{
	return new QBluetoothAddress(static_cast<QBluetoothTransferRequest*>(ptr)->address());
}

void* QBluetoothTransferRequest_Attribute(void* ptr, long long code, void* defaultValue)
{
	return new QVariant(static_cast<QBluetoothTransferRequest*>(ptr)->attribute(static_cast<QBluetoothTransferRequest::Attribute>(code), *static_cast<QVariant*>(defaultValue)));
}

void QBluetoothTransferRequest_SetAttribute(void* ptr, long long code, void* value)
{
	static_cast<QBluetoothTransferRequest*>(ptr)->setAttribute(static_cast<QBluetoothTransferRequest::Attribute>(code), *static_cast<QVariant*>(value));
}

void QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(void* ptr)
{
	static_cast<QBluetoothTransferRequest*>(ptr)->~QBluetoothTransferRequest();
}

void* QBluetoothUuid_NewQBluetoothUuid()
{
	return new QBluetoothUuid();
}

void* QBluetoothUuid_NewQBluetoothUuid4(long long uuid)
{
	return new QBluetoothUuid(static_cast<QBluetoothUuid::CharacteristicType>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid5(long long uuid)
{
	return new QBluetoothUuid(static_cast<QBluetoothUuid::DescriptorType>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid2(long long uuid)
{
	return new QBluetoothUuid(static_cast<QBluetoothUuid::ProtocolUuid>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid3(long long uuid)
{
	return new QBluetoothUuid(static_cast<QBluetoothUuid::ServiceClassUuid>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid10(void* uuid)
{
	return new QBluetoothUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid9(char* uuid)
{
	return new QBluetoothUuid(QString(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid11(void* uuid)
{
	return new QBluetoothUuid(*static_cast<QUuid*>(uuid));
}

void* QBluetoothUuid_NewQBluetoothUuid6(unsigned short uuid)
{
	return new QBluetoothUuid(uuid);
}

void* QBluetoothUuid_NewQBluetoothUuid7(unsigned int uuid)
{
	return new QBluetoothUuid(uuid);
}

struct QtBluetooth_PackedString QBluetoothUuid_QBluetoothUuid_CharacteristicToString(long long uuid)
{
	return ({ QByteArray t44a871 = QBluetoothUuid::characteristicToString(static_cast<QBluetoothUuid::CharacteristicType>(uuid)).toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t44a871.prepend("WHITESPACE").constData()+10), t44a871.size()-10 }; });
}

struct QtBluetooth_PackedString QBluetoothUuid_QBluetoothUuid_DescriptorToString(long long uuid)
{
	return ({ QByteArray t3c5c65 = QBluetoothUuid::descriptorToString(static_cast<QBluetoothUuid::DescriptorType>(uuid)).toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t3c5c65.prepend("WHITESPACE").constData()+10), t3c5c65.size()-10 }; });
}

int QBluetoothUuid_MinimumSize(void* ptr)
{
	return static_cast<QBluetoothUuid*>(ptr)->minimumSize();
}

struct QtBluetooth_PackedString QBluetoothUuid_QBluetoothUuid_ProtocolToString(long long uuid)
{
	return ({ QByteArray tc75cac = QBluetoothUuid::protocolToString(static_cast<QBluetoothUuid::ProtocolUuid>(uuid)).toUtf8(); QtBluetooth_PackedString { const_cast<char*>(tc75cac.prepend("WHITESPACE").constData()+10), tc75cac.size()-10 }; });
}

struct QtBluetooth_PackedString QBluetoothUuid_QBluetoothUuid_ServiceClassToString(long long uuid)
{
	return ({ QByteArray t3b0299 = QBluetoothUuid::serviceClassToString(static_cast<QBluetoothUuid::ServiceClassUuid>(uuid)).toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t3b0299.prepend("WHITESPACE").constData()+10), t3b0299.size()-10 }; });
}

unsigned short QBluetoothUuid_ToUInt16(void* ptr, char ok)
{
	return static_cast<QBluetoothUuid*>(ptr)->toUInt16(NULL);
}

unsigned int QBluetoothUuid_ToUInt32(void* ptr, char ok)
{
	return static_cast<QBluetoothUuid*>(ptr)->toUInt32(NULL);
}

void QBluetoothUuid_DestroyQBluetoothUuid(void* ptr)
{
	static_cast<QBluetoothUuid*>(ptr)->~QBluetoothUuid();
}

void* QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData()
{
	return new QLowEnergyAdvertisingData();
}

void* QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData2(void* other)
{
	return new QLowEnergyAdvertisingData(*static_cast<QLowEnergyAdvertisingData*>(other));
}

long long QLowEnergyAdvertisingData_Discoverability(void* ptr)
{
	return static_cast<QLowEnergyAdvertisingData*>(ptr)->discoverability();
}

char QLowEnergyAdvertisingData_IncludePowerLevel(void* ptr)
{
	return static_cast<QLowEnergyAdvertisingData*>(ptr)->includePowerLevel();
}

unsigned short QLowEnergyAdvertisingData_QLowEnergyAdvertisingData_InvalidManufacturerId()
{
	return QLowEnergyAdvertisingData::invalidManufacturerId();
}

struct QtBluetooth_PackedString QLowEnergyAdvertisingData_LocalName(void* ptr)
{
	return ({ QByteArray t61112a = static_cast<QLowEnergyAdvertisingData*>(ptr)->localName().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t61112a.prepend("WHITESPACE").constData()+10), t61112a.size()-10 }; });
}

void* QLowEnergyAdvertisingData_ManufacturerData(void* ptr)
{
	return new QByteArray(static_cast<QLowEnergyAdvertisingData*>(ptr)->manufacturerData());
}

unsigned short QLowEnergyAdvertisingData_ManufacturerId(void* ptr)
{
	return static_cast<QLowEnergyAdvertisingData*>(ptr)->manufacturerId();
}

void* QLowEnergyAdvertisingData_RawData(void* ptr)
{
	return new QByteArray(static_cast<QLowEnergyAdvertisingData*>(ptr)->rawData());
}

struct QtBluetooth_PackedList QLowEnergyAdvertisingData_Services(void* ptr)
{
	return ({ QList<QBluetoothUuid>* tmpValue = new QList<QBluetoothUuid>(static_cast<QLowEnergyAdvertisingData*>(ptr)->services()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

void QLowEnergyAdvertisingData_SetDiscoverability(void* ptr, long long mode)
{
	static_cast<QLowEnergyAdvertisingData*>(ptr)->setDiscoverability(static_cast<QLowEnergyAdvertisingData::Discoverability>(mode));
}

void QLowEnergyAdvertisingData_SetIncludePowerLevel(void* ptr, char doInclude)
{
	static_cast<QLowEnergyAdvertisingData*>(ptr)->setIncludePowerLevel(doInclude != 0);
}

void QLowEnergyAdvertisingData_SetLocalName(void* ptr, char* name)
{
	static_cast<QLowEnergyAdvertisingData*>(ptr)->setLocalName(QString(name));
}

void QLowEnergyAdvertisingData_SetManufacturerData(void* ptr, unsigned short id, void* data)
{
	static_cast<QLowEnergyAdvertisingData*>(ptr)->setManufacturerData(id, *static_cast<QByteArray*>(data));
}

void QLowEnergyAdvertisingData_SetRawData(void* ptr, void* data)
{
	static_cast<QLowEnergyAdvertisingData*>(ptr)->setRawData(*static_cast<QByteArray*>(data));
}

void QLowEnergyAdvertisingData_Swap(void* ptr, void* other)
{
	static_cast<QLowEnergyAdvertisingData*>(ptr)->swap(*static_cast<QLowEnergyAdvertisingData*>(other));
}

void QLowEnergyAdvertisingData_DestroyQLowEnergyAdvertisingData(void* ptr)
{
	static_cast<QLowEnergyAdvertisingData*>(ptr)->~QLowEnergyAdvertisingData();
}

void* QLowEnergyAdvertisingData_services_atList(void* ptr, int i)
{
	return new QBluetoothUuid(static_cast<QList<QBluetoothUuid>*>(ptr)->at(i));
}

void* QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters()
{
	return new QLowEnergyAdvertisingParameters();
}

void* QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters2(void* other)
{
	return new QLowEnergyAdvertisingParameters(*static_cast<QLowEnergyAdvertisingParameters*>(other));
}

long long QLowEnergyAdvertisingParameters_FilterPolicy(void* ptr)
{
	return static_cast<QLowEnergyAdvertisingParameters*>(ptr)->filterPolicy();
}

int QLowEnergyAdvertisingParameters_MaximumInterval(void* ptr)
{
	return static_cast<QLowEnergyAdvertisingParameters*>(ptr)->maximumInterval();
}

int QLowEnergyAdvertisingParameters_MinimumInterval(void* ptr)
{
	return static_cast<QLowEnergyAdvertisingParameters*>(ptr)->minimumInterval();
}

long long QLowEnergyAdvertisingParameters_Mode(void* ptr)
{
	return static_cast<QLowEnergyAdvertisingParameters*>(ptr)->mode();
}

void QLowEnergyAdvertisingParameters_SetInterval(void* ptr, unsigned short minimum, unsigned short maximum)
{
	static_cast<QLowEnergyAdvertisingParameters*>(ptr)->setInterval(minimum, maximum);
}

void QLowEnergyAdvertisingParameters_SetMode(void* ptr, long long mode)
{
	static_cast<QLowEnergyAdvertisingParameters*>(ptr)->setMode(static_cast<QLowEnergyAdvertisingParameters::Mode>(mode));
}

void QLowEnergyAdvertisingParameters_Swap(void* ptr, void* other)
{
	static_cast<QLowEnergyAdvertisingParameters*>(ptr)->swap(*static_cast<QLowEnergyAdvertisingParameters*>(other));
}

void QLowEnergyAdvertisingParameters_DestroyQLowEnergyAdvertisingParameters(void* ptr)
{
	static_cast<QLowEnergyAdvertisingParameters*>(ptr)->~QLowEnergyAdvertisingParameters();
}

void* QLowEnergyCharacteristic_NewQLowEnergyCharacteristic()
{
	return new QLowEnergyCharacteristic();
}

void* QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(void* other)
{
	return new QLowEnergyCharacteristic(*static_cast<QLowEnergyCharacteristic*>(other));
}

void* QLowEnergyCharacteristic_Descriptor(void* ptr, void* uuid)
{
	return new QLowEnergyDescriptor(static_cast<QLowEnergyCharacteristic*>(ptr)->descriptor(*static_cast<QBluetoothUuid*>(uuid)));
}

struct QtBluetooth_PackedList QLowEnergyCharacteristic_Descriptors(void* ptr)
{
	return ({ QList<QLowEnergyDescriptor>* tmpValue = new QList<QLowEnergyDescriptor>(static_cast<QLowEnergyCharacteristic*>(ptr)->descriptors()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

char QLowEnergyCharacteristic_IsValid(void* ptr)
{
	return static_cast<QLowEnergyCharacteristic*>(ptr)->isValid();
}

struct QtBluetooth_PackedString QLowEnergyCharacteristic_Name(void* ptr)
{
	return ({ QByteArray td530f8 = static_cast<QLowEnergyCharacteristic*>(ptr)->name().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(td530f8.prepend("WHITESPACE").constData()+10), td530f8.size()-10 }; });
}

long long QLowEnergyCharacteristic_Properties(void* ptr)
{
	return static_cast<QLowEnergyCharacteristic*>(ptr)->properties();
}

void* QLowEnergyCharacteristic_Uuid(void* ptr)
{
	return new QBluetoothUuid(static_cast<QLowEnergyCharacteristic*>(ptr)->uuid());
}

void* QLowEnergyCharacteristic_Value(void* ptr)
{
	return new QByteArray(static_cast<QLowEnergyCharacteristic*>(ptr)->value());
}

void QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(void* ptr)
{
	static_cast<QLowEnergyCharacteristic*>(ptr)->~QLowEnergyCharacteristic();
}

void* QLowEnergyCharacteristic_descriptors_atList(void* ptr, int i)
{
	return new QLowEnergyDescriptor(static_cast<QList<QLowEnergyDescriptor>*>(ptr)->at(i));
}

void* QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData()
{
	return new QLowEnergyCharacteristicData();
}

void* QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData2(void* other)
{
	return new QLowEnergyCharacteristicData(*static_cast<QLowEnergyCharacteristicData*>(other));
}

void QLowEnergyCharacteristicData_AddDescriptor(void* ptr, void* descriptor)
{
	static_cast<QLowEnergyCharacteristicData*>(ptr)->addDescriptor(*static_cast<QLowEnergyDescriptorData*>(descriptor));
}

struct QtBluetooth_PackedList QLowEnergyCharacteristicData_Descriptors(void* ptr)
{
	return ({ QList<QLowEnergyDescriptorData>* tmpValue = new QList<QLowEnergyDescriptorData>(static_cast<QLowEnergyCharacteristicData*>(ptr)->descriptors()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

char QLowEnergyCharacteristicData_IsValid(void* ptr)
{
	return static_cast<QLowEnergyCharacteristicData*>(ptr)->isValid();
}

int QLowEnergyCharacteristicData_MaximumValueLength(void* ptr)
{
	return static_cast<QLowEnergyCharacteristicData*>(ptr)->maximumValueLength();
}

int QLowEnergyCharacteristicData_MinimumValueLength(void* ptr)
{
	return static_cast<QLowEnergyCharacteristicData*>(ptr)->minimumValueLength();
}

long long QLowEnergyCharacteristicData_Properties(void* ptr)
{
	return static_cast<QLowEnergyCharacteristicData*>(ptr)->properties();
}

void QLowEnergyCharacteristicData_SetProperties(void* ptr, long long properties)
{
	static_cast<QLowEnergyCharacteristicData*>(ptr)->setProperties(static_cast<QLowEnergyCharacteristic::PropertyType>(properties));
}

void QLowEnergyCharacteristicData_SetUuid(void* ptr, void* uuid)
{
	static_cast<QLowEnergyCharacteristicData*>(ptr)->setUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void QLowEnergyCharacteristicData_SetValue(void* ptr, void* value)
{
	static_cast<QLowEnergyCharacteristicData*>(ptr)->setValue(*static_cast<QByteArray*>(value));
}

void QLowEnergyCharacteristicData_SetValueLength(void* ptr, int minimum, int maximum)
{
	static_cast<QLowEnergyCharacteristicData*>(ptr)->setValueLength(minimum, maximum);
}

void QLowEnergyCharacteristicData_Swap(void* ptr, void* other)
{
	static_cast<QLowEnergyCharacteristicData*>(ptr)->swap(*static_cast<QLowEnergyCharacteristicData*>(other));
}

void* QLowEnergyCharacteristicData_Uuid(void* ptr)
{
	return new QBluetoothUuid(static_cast<QLowEnergyCharacteristicData*>(ptr)->uuid());
}

void* QLowEnergyCharacteristicData_Value(void* ptr)
{
	return new QByteArray(static_cast<QLowEnergyCharacteristicData*>(ptr)->value());
}

void QLowEnergyCharacteristicData_DestroyQLowEnergyCharacteristicData(void* ptr)
{
	static_cast<QLowEnergyCharacteristicData*>(ptr)->~QLowEnergyCharacteristicData();
}

void* QLowEnergyCharacteristicData_descriptors_atList(void* ptr, int i)
{
	return new QLowEnergyDescriptorData(static_cast<QList<QLowEnergyDescriptorData>*>(ptr)->at(i));
}

void* QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters()
{
	return new QLowEnergyConnectionParameters();
}

void* QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters2(void* other)
{
	return new QLowEnergyConnectionParameters(*static_cast<QLowEnergyConnectionParameters*>(other));
}

int QLowEnergyConnectionParameters_Latency(void* ptr)
{
	return static_cast<QLowEnergyConnectionParameters*>(ptr)->latency();
}

double QLowEnergyConnectionParameters_MaximumInterval(void* ptr)
{
	return static_cast<QLowEnergyConnectionParameters*>(ptr)->maximumInterval();
}

double QLowEnergyConnectionParameters_MinimumInterval(void* ptr)
{
	return static_cast<QLowEnergyConnectionParameters*>(ptr)->minimumInterval();
}

void QLowEnergyConnectionParameters_SetIntervalRange(void* ptr, double minimum, double maximum)
{
	static_cast<QLowEnergyConnectionParameters*>(ptr)->setIntervalRange(minimum, maximum);
}

void QLowEnergyConnectionParameters_SetLatency(void* ptr, int latency)
{
	static_cast<QLowEnergyConnectionParameters*>(ptr)->setLatency(latency);
}

void QLowEnergyConnectionParameters_SetSupervisionTimeout(void* ptr, int timeout)
{
	static_cast<QLowEnergyConnectionParameters*>(ptr)->setSupervisionTimeout(timeout);
}

int QLowEnergyConnectionParameters_SupervisionTimeout(void* ptr)
{
	return static_cast<QLowEnergyConnectionParameters*>(ptr)->supervisionTimeout();
}

void QLowEnergyConnectionParameters_Swap(void* ptr, void* other)
{
	static_cast<QLowEnergyConnectionParameters*>(ptr)->swap(*static_cast<QLowEnergyConnectionParameters*>(other));
}

void QLowEnergyConnectionParameters_DestroyQLowEnergyConnectionParameters(void* ptr)
{
	static_cast<QLowEnergyConnectionParameters*>(ptr)->~QLowEnergyConnectionParameters();
}

class MyQLowEnergyController: public QLowEnergyController
{
public:
	void Signal_Connected() { callbackQLowEnergyController_Connected(this); };
	void Signal_ConnectionUpdated(const QLowEnergyConnectionParameters & newParameters) { callbackQLowEnergyController_ConnectionUpdated(this, const_cast<QLowEnergyConnectionParameters*>(&newParameters)); };
	void Signal_Disconnected() { callbackQLowEnergyController_Disconnected(this); };
	void Signal_DiscoveryFinished() { callbackQLowEnergyController_DiscoveryFinished(this); };
	void Signal_Error2(QLowEnergyController::Error newError) { callbackQLowEnergyController_Error2(this, newError); };
	void Signal_ServiceDiscovered(const QBluetoothUuid & newService) { callbackQLowEnergyController_ServiceDiscovered(this, const_cast<QBluetoothUuid*>(&newService)); };
	void Signal_StateChanged(QLowEnergyController::ControllerState state) { callbackQLowEnergyController_StateChanged(this, state); };
	void timerEvent(QTimerEvent * event) { callbackQLowEnergyController_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQLowEnergyController_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLowEnergyController_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLowEnergyController_CustomEvent(this, event); };
	void deleteLater() { callbackQLowEnergyController_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLowEnergyController_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLowEnergyController_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLowEnergyController_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLowEnergyController_MetaObject(const_cast<MyQLowEnergyController*>(this))); };
};

void QLowEnergyController_ConnectConnected(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::connected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Connected));
}

void QLowEnergyController_DisconnectConnected(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::connected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Connected));
}

void QLowEnergyController_Connected(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->connected();
}

void QLowEnergyController_ConnectConnectionUpdated(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(const QLowEnergyConnectionParameters &)>(&QLowEnergyController::connectionUpdated), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(const QLowEnergyConnectionParameters &)>(&MyQLowEnergyController::Signal_ConnectionUpdated));
}

void QLowEnergyController_DisconnectConnectionUpdated(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(const QLowEnergyConnectionParameters &)>(&QLowEnergyController::connectionUpdated), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(const QLowEnergyConnectionParameters &)>(&MyQLowEnergyController::Signal_ConnectionUpdated));
}

void QLowEnergyController_ConnectionUpdated(void* ptr, void* newParameters)
{
	static_cast<QLowEnergyController*>(ptr)->connectionUpdated(*static_cast<QLowEnergyConnectionParameters*>(newParameters));
}

void QLowEnergyController_ConnectDisconnected(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::disconnected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Disconnected));
}

void QLowEnergyController_DisconnectDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::disconnected), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_Disconnected));
}

void QLowEnergyController_Disconnected(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->disconnected();
}

void QLowEnergyController_ConnectDiscoveryFinished(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::discoveryFinished), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_DiscoveryFinished));
}

void QLowEnergyController_DisconnectDiscoveryFinished(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)()>(&QLowEnergyController::discoveryFinished), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)()>(&MyQLowEnergyController::Signal_DiscoveryFinished));
}

void QLowEnergyController_DiscoveryFinished(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->discoveryFinished();
}

void QLowEnergyController_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::Error)>(&QLowEnergyController::error), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::Error)>(&MyQLowEnergyController::Signal_Error2));
}

void QLowEnergyController_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::Error)>(&QLowEnergyController::error), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::Error)>(&MyQLowEnergyController::Signal_Error2));
}

void QLowEnergyController_Error2(void* ptr, long long newError)
{
	static_cast<QLowEnergyController*>(ptr)->error(static_cast<QLowEnergyController::Error>(newError));
}

void QLowEnergyController_ConnectServiceDiscovered(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(const QBluetoothUuid &)>(&QLowEnergyController::serviceDiscovered), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(const QBluetoothUuid &)>(&MyQLowEnergyController::Signal_ServiceDiscovered));
}

void QLowEnergyController_DisconnectServiceDiscovered(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(const QBluetoothUuid &)>(&QLowEnergyController::serviceDiscovered), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(const QBluetoothUuid &)>(&MyQLowEnergyController::Signal_ServiceDiscovered));
}

void QLowEnergyController_ServiceDiscovered(void* ptr, void* newService)
{
	static_cast<QLowEnergyController*>(ptr)->serviceDiscovered(*static_cast<QBluetoothUuid*>(newService));
}

void QLowEnergyController_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::ControllerState)>(&QLowEnergyController::stateChanged), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::ControllerState)>(&MyQLowEnergyController::Signal_StateChanged));
}

void QLowEnergyController_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyController*>(ptr), static_cast<void (QLowEnergyController::*)(QLowEnergyController::ControllerState)>(&QLowEnergyController::stateChanged), static_cast<MyQLowEnergyController*>(ptr), static_cast<void (MyQLowEnergyController::*)(QLowEnergyController::ControllerState)>(&MyQLowEnergyController::Signal_StateChanged));
}

void QLowEnergyController_StateChanged(void* ptr, long long state)
{
	static_cast<QLowEnergyController*>(ptr)->stateChanged(static_cast<QLowEnergyController::ControllerState>(state));
}

void* QLowEnergyController_AddService(void* ptr, void* service, void* parent)
{
	return static_cast<QLowEnergyController*>(ptr)->addService(*static_cast<QLowEnergyServiceData*>(service), static_cast<QObject*>(parent));
}

void QLowEnergyController_ConnectToDevice(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->connectToDevice();
}

void* QLowEnergyController_QLowEnergyController_CreateCentral(void* remoteDevice, void* parent)
{
	return QLowEnergyController::createCentral(*static_cast<QBluetoothDeviceInfo*>(remoteDevice), static_cast<QObject*>(parent));
}

void* QLowEnergyController_QLowEnergyController_CreatePeripheral(void* parent)
{
	return QLowEnergyController::createPeripheral(static_cast<QObject*>(parent));
}

void* QLowEnergyController_CreateServiceObject(void* ptr, void* serviceUuid, void* parent)
{
	return static_cast<QLowEnergyController*>(ptr)->createServiceObject(*static_cast<QBluetoothUuid*>(serviceUuid), static_cast<QObject*>(parent));
}

void QLowEnergyController_DisconnectFromDevice(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->disconnectFromDevice();
}

void QLowEnergyController_DiscoverServices(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->discoverServices();
}

long long QLowEnergyController_Error(void* ptr)
{
	return static_cast<QLowEnergyController*>(ptr)->error();
}

struct QtBluetooth_PackedString QLowEnergyController_ErrorString(void* ptr)
{
	return ({ QByteArray t56a9bd = static_cast<QLowEnergyController*>(ptr)->errorString().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(t56a9bd.prepend("WHITESPACE").constData()+10), t56a9bd.size()-10 }; });
}

void* QLowEnergyController_LocalAddress(void* ptr)
{
	return new QBluetoothAddress(static_cast<QLowEnergyController*>(ptr)->localAddress());
}

void* QLowEnergyController_RemoteAddress(void* ptr)
{
	return new QBluetoothAddress(static_cast<QLowEnergyController*>(ptr)->remoteAddress());
}

long long QLowEnergyController_RemoteAddressType(void* ptr)
{
	return static_cast<QLowEnergyController*>(ptr)->remoteAddressType();
}

struct QtBluetooth_PackedString QLowEnergyController_RemoteName(void* ptr)
{
	return ({ QByteArray tf2501f = static_cast<QLowEnergyController*>(ptr)->remoteName().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(tf2501f.prepend("WHITESPACE").constData()+10), tf2501f.size()-10 }; });
}

void QLowEnergyController_RequestConnectionUpdate(void* ptr, void* parameters)
{
	static_cast<QLowEnergyController*>(ptr)->requestConnectionUpdate(*static_cast<QLowEnergyConnectionParameters*>(parameters));
}

long long QLowEnergyController_Role(void* ptr)
{
	return static_cast<QLowEnergyController*>(ptr)->role();
}

struct QtBluetooth_PackedList QLowEnergyController_Services(void* ptr)
{
	return ({ QList<QBluetoothUuid>* tmpValue = new QList<QBluetoothUuid>(static_cast<QLowEnergyController*>(ptr)->services()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

void QLowEnergyController_SetRemoteAddressType(void* ptr, long long ty)
{
	static_cast<QLowEnergyController*>(ptr)->setRemoteAddressType(static_cast<QLowEnergyController::RemoteAddressType>(ty));
}

void QLowEnergyController_StartAdvertising(void* ptr, void* parameters, void* advertisingData, void* scanResponseData)
{
	static_cast<QLowEnergyController*>(ptr)->startAdvertising(*static_cast<QLowEnergyAdvertisingParameters*>(parameters), *static_cast<QLowEnergyAdvertisingData*>(advertisingData), *static_cast<QLowEnergyAdvertisingData*>(scanResponseData));
}

long long QLowEnergyController_State(void* ptr)
{
	return static_cast<QLowEnergyController*>(ptr)->state();
}

void QLowEnergyController_StopAdvertising(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->stopAdvertising();
}

void QLowEnergyController_DestroyQLowEnergyController(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->~QLowEnergyController();
}

void* QLowEnergyController_services_atList(void* ptr, int i)
{
	return new QBluetoothUuid(static_cast<QList<QBluetoothUuid>*>(ptr)->at(i));
}

void QLowEnergyController_TimerEvent(void* ptr, void* event)
{
	static_cast<QLowEnergyController*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLowEnergyController_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLowEnergyController_ChildEvent(void* ptr, void* event)
{
	static_cast<QLowEnergyController*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLowEnergyController_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::childEvent(static_cast<QChildEvent*>(event));
}

void QLowEnergyController_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QLowEnergyController*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLowEnergyController_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLowEnergyController_CustomEvent(void* ptr, void* event)
{
	static_cast<QLowEnergyController*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLowEnergyController_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::customEvent(static_cast<QEvent*>(event));
}

void QLowEnergyController_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLowEnergyController*>(ptr), "deleteLater");
}

void QLowEnergyController_DeleteLaterDefault(void* ptr)
{
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::deleteLater();
}

void QLowEnergyController_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QLowEnergyController*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLowEnergyController_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QLowEnergyController_Event(void* ptr, void* e)
{
	return static_cast<QLowEnergyController*>(ptr)->event(static_cast<QEvent*>(e));
}

char QLowEnergyController_EventDefault(void* ptr, void* e)
{
	return static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::event(static_cast<QEvent*>(e));
}

char QLowEnergyController_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QLowEnergyController*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QLowEnergyController_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QLowEnergyController_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLowEnergyController*>(ptr)->metaObject());
}

void* QLowEnergyController_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLowEnergyController*>(ptr)->QLowEnergyController::metaObject());
}

void* QLowEnergyDescriptor_NewQLowEnergyDescriptor()
{
	return new QLowEnergyDescriptor();
}

void* QLowEnergyDescriptor_NewQLowEnergyDescriptor2(void* other)
{
	return new QLowEnergyDescriptor(*static_cast<QLowEnergyDescriptor*>(other));
}

char QLowEnergyDescriptor_IsValid(void* ptr)
{
	return static_cast<QLowEnergyDescriptor*>(ptr)->isValid();
}

struct QtBluetooth_PackedString QLowEnergyDescriptor_Name(void* ptr)
{
	return ({ QByteArray tf5ebfa = static_cast<QLowEnergyDescriptor*>(ptr)->name().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(tf5ebfa.prepend("WHITESPACE").constData()+10), tf5ebfa.size()-10 }; });
}

long long QLowEnergyDescriptor_Type(void* ptr)
{
	return static_cast<QLowEnergyDescriptor*>(ptr)->type();
}

void* QLowEnergyDescriptor_Uuid(void* ptr)
{
	return new QBluetoothUuid(static_cast<QLowEnergyDescriptor*>(ptr)->uuid());
}

void* QLowEnergyDescriptor_Value(void* ptr)
{
	return new QByteArray(static_cast<QLowEnergyDescriptor*>(ptr)->value());
}

void QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(void* ptr)
{
	static_cast<QLowEnergyDescriptor*>(ptr)->~QLowEnergyDescriptor();
}

void* QLowEnergyDescriptorData_NewQLowEnergyDescriptorData()
{
	return new QLowEnergyDescriptorData();
}

void* QLowEnergyDescriptorData_NewQLowEnergyDescriptorData2(void* uuid, void* value)
{
	return new QLowEnergyDescriptorData(*static_cast<QBluetoothUuid*>(uuid), *static_cast<QByteArray*>(value));
}

void* QLowEnergyDescriptorData_NewQLowEnergyDescriptorData3(void* other)
{
	return new QLowEnergyDescriptorData(*static_cast<QLowEnergyDescriptorData*>(other));
}

char QLowEnergyDescriptorData_IsReadable(void* ptr)
{
	return static_cast<QLowEnergyDescriptorData*>(ptr)->isReadable();
}

char QLowEnergyDescriptorData_IsValid(void* ptr)
{
	return static_cast<QLowEnergyDescriptorData*>(ptr)->isValid();
}

char QLowEnergyDescriptorData_IsWritable(void* ptr)
{
	return static_cast<QLowEnergyDescriptorData*>(ptr)->isWritable();
}

void QLowEnergyDescriptorData_SetUuid(void* ptr, void* uuid)
{
	static_cast<QLowEnergyDescriptorData*>(ptr)->setUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void QLowEnergyDescriptorData_SetValue(void* ptr, void* value)
{
	static_cast<QLowEnergyDescriptorData*>(ptr)->setValue(*static_cast<QByteArray*>(value));
}

void QLowEnergyDescriptorData_Swap(void* ptr, void* other)
{
	static_cast<QLowEnergyDescriptorData*>(ptr)->swap(*static_cast<QLowEnergyDescriptorData*>(other));
}

void* QLowEnergyDescriptorData_Uuid(void* ptr)
{
	return new QBluetoothUuid(static_cast<QLowEnergyDescriptorData*>(ptr)->uuid());
}

void* QLowEnergyDescriptorData_Value(void* ptr)
{
	return new QByteArray(static_cast<QLowEnergyDescriptorData*>(ptr)->value());
}

void QLowEnergyDescriptorData_DestroyQLowEnergyDescriptorData(void* ptr)
{
	static_cast<QLowEnergyDescriptorData*>(ptr)->~QLowEnergyDescriptorData();
}

class MyQLowEnergyService: public QLowEnergyService
{
public:
	void Signal_CharacteristicChanged(const QLowEnergyCharacteristic & characteristic, const QByteArray & newValue) { callbackQLowEnergyService_CharacteristicChanged(this, const_cast<QLowEnergyCharacteristic*>(&characteristic), const_cast<QByteArray*>(&newValue)); };
	void Signal_CharacteristicRead(const QLowEnergyCharacteristic & characteristic, const QByteArray & value) { callbackQLowEnergyService_CharacteristicRead(this, const_cast<QLowEnergyCharacteristic*>(&characteristic), const_cast<QByteArray*>(&value)); };
	void Signal_CharacteristicWritten(const QLowEnergyCharacteristic & characteristic, const QByteArray & newValue) { callbackQLowEnergyService_CharacteristicWritten(this, const_cast<QLowEnergyCharacteristic*>(&characteristic), const_cast<QByteArray*>(&newValue)); };
	void Signal_DescriptorRead(const QLowEnergyDescriptor & descriptor, const QByteArray & value) { callbackQLowEnergyService_DescriptorRead(this, const_cast<QLowEnergyDescriptor*>(&descriptor), const_cast<QByteArray*>(&value)); };
	void Signal_DescriptorWritten(const QLowEnergyDescriptor & descriptor, const QByteArray & newValue) { callbackQLowEnergyService_DescriptorWritten(this, const_cast<QLowEnergyDescriptor*>(&descriptor), const_cast<QByteArray*>(&newValue)); };
	void Signal_Error2(QLowEnergyService::ServiceError newError) { callbackQLowEnergyService_Error2(this, newError); };
	void Signal_StateChanged(QLowEnergyService::ServiceState newState) { callbackQLowEnergyService_StateChanged(this, newState); };
	void timerEvent(QTimerEvent * event) { callbackQLowEnergyService_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQLowEnergyService_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLowEnergyService_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLowEnergyService_CustomEvent(this, event); };
	void deleteLater() { callbackQLowEnergyService_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLowEnergyService_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLowEnergyService_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLowEnergyService_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLowEnergyService_MetaObject(const_cast<MyQLowEnergyService*>(this))); };
};

void QLowEnergyService_ConnectCharacteristicChanged(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&QLowEnergyService::characteristicChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&MyQLowEnergyService::Signal_CharacteristicChanged));
}

void QLowEnergyService_DisconnectCharacteristicChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&QLowEnergyService::characteristicChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&MyQLowEnergyService::Signal_CharacteristicChanged));
}

void QLowEnergyService_CharacteristicChanged(void* ptr, void* characteristic, void* newValue)
{
	static_cast<QLowEnergyService*>(ptr)->characteristicChanged(*static_cast<QLowEnergyCharacteristic*>(characteristic), *static_cast<QByteArray*>(newValue));
}

void QLowEnergyService_ConnectCharacteristicRead(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&QLowEnergyService::characteristicRead), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&MyQLowEnergyService::Signal_CharacteristicRead));
}

void QLowEnergyService_DisconnectCharacteristicRead(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&QLowEnergyService::characteristicRead), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&MyQLowEnergyService::Signal_CharacteristicRead));
}

void QLowEnergyService_CharacteristicRead(void* ptr, void* characteristic, void* value)
{
	static_cast<QLowEnergyService*>(ptr)->characteristicRead(*static_cast<QLowEnergyCharacteristic*>(characteristic), *static_cast<QByteArray*>(value));
}

void QLowEnergyService_ConnectCharacteristicWritten(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&QLowEnergyService::characteristicWritten), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&MyQLowEnergyService::Signal_CharacteristicWritten));
}

void QLowEnergyService_DisconnectCharacteristicWritten(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&QLowEnergyService::characteristicWritten), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyCharacteristic &, const QByteArray &)>(&MyQLowEnergyService::Signal_CharacteristicWritten));
}

void QLowEnergyService_CharacteristicWritten(void* ptr, void* characteristic, void* newValue)
{
	static_cast<QLowEnergyService*>(ptr)->characteristicWritten(*static_cast<QLowEnergyCharacteristic*>(characteristic), *static_cast<QByteArray*>(newValue));
}

void QLowEnergyService_ConnectDescriptorRead(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyDescriptor &, const QByteArray &)>(&QLowEnergyService::descriptorRead), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyDescriptor &, const QByteArray &)>(&MyQLowEnergyService::Signal_DescriptorRead));
}

void QLowEnergyService_DisconnectDescriptorRead(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyDescriptor &, const QByteArray &)>(&QLowEnergyService::descriptorRead), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyDescriptor &, const QByteArray &)>(&MyQLowEnergyService::Signal_DescriptorRead));
}

void QLowEnergyService_DescriptorRead(void* ptr, void* descriptor, void* value)
{
	static_cast<QLowEnergyService*>(ptr)->descriptorRead(*static_cast<QLowEnergyDescriptor*>(descriptor), *static_cast<QByteArray*>(value));
}

void QLowEnergyService_ConnectDescriptorWritten(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyDescriptor &, const QByteArray &)>(&QLowEnergyService::descriptorWritten), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyDescriptor &, const QByteArray &)>(&MyQLowEnergyService::Signal_DescriptorWritten));
}

void QLowEnergyService_DisconnectDescriptorWritten(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(const QLowEnergyDescriptor &, const QByteArray &)>(&QLowEnergyService::descriptorWritten), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(const QLowEnergyDescriptor &, const QByteArray &)>(&MyQLowEnergyService::Signal_DescriptorWritten));
}

void QLowEnergyService_DescriptorWritten(void* ptr, void* descriptor, void* newValue)
{
	static_cast<QLowEnergyService*>(ptr)->descriptorWritten(*static_cast<QLowEnergyDescriptor*>(descriptor), *static_cast<QByteArray*>(newValue));
}

void QLowEnergyService_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceError)>(&QLowEnergyService::error), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceError)>(&MyQLowEnergyService::Signal_Error2));
}

void QLowEnergyService_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceError)>(&QLowEnergyService::error), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceError)>(&MyQLowEnergyService::Signal_Error2));
}

void QLowEnergyService_Error2(void* ptr, long long newError)
{
	static_cast<QLowEnergyService*>(ptr)->error(static_cast<QLowEnergyService::ServiceError>(newError));
}

void QLowEnergyService_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceState)>(&QLowEnergyService::stateChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceState)>(&MyQLowEnergyService::Signal_StateChanged));
}

void QLowEnergyService_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLowEnergyService*>(ptr), static_cast<void (QLowEnergyService::*)(QLowEnergyService::ServiceState)>(&QLowEnergyService::stateChanged), static_cast<MyQLowEnergyService*>(ptr), static_cast<void (MyQLowEnergyService::*)(QLowEnergyService::ServiceState)>(&MyQLowEnergyService::Signal_StateChanged));
}

void QLowEnergyService_StateChanged(void* ptr, long long newState)
{
	static_cast<QLowEnergyService*>(ptr)->stateChanged(static_cast<QLowEnergyService::ServiceState>(newState));
}

void* QLowEnergyService_Characteristic(void* ptr, void* uuid)
{
	return new QLowEnergyCharacteristic(static_cast<QLowEnergyService*>(ptr)->characteristic(*static_cast<QBluetoothUuid*>(uuid)));
}

struct QtBluetooth_PackedList QLowEnergyService_Characteristics(void* ptr)
{
	return ({ QList<QLowEnergyCharacteristic>* tmpValue = new QList<QLowEnergyCharacteristic>(static_cast<QLowEnergyService*>(ptr)->characteristics()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

char QLowEnergyService_Contains(void* ptr, void* characteristic)
{
	return static_cast<QLowEnergyService*>(ptr)->contains(*static_cast<QLowEnergyCharacteristic*>(characteristic));
}

char QLowEnergyService_Contains2(void* ptr, void* descriptor)
{
	return static_cast<QLowEnergyService*>(ptr)->contains(*static_cast<QLowEnergyDescriptor*>(descriptor));
}

void QLowEnergyService_DiscoverDetails(void* ptr)
{
	static_cast<QLowEnergyService*>(ptr)->discoverDetails();
}

long long QLowEnergyService_Error(void* ptr)
{
	return static_cast<QLowEnergyService*>(ptr)->error();
}

struct QtBluetooth_PackedList QLowEnergyService_IncludedServices(void* ptr)
{
	return ({ QList<QBluetoothUuid>* tmpValue = new QList<QBluetoothUuid>(static_cast<QLowEnergyService*>(ptr)->includedServices()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

void QLowEnergyService_ReadCharacteristic(void* ptr, void* characteristic)
{
	static_cast<QLowEnergyService*>(ptr)->readCharacteristic(*static_cast<QLowEnergyCharacteristic*>(characteristic));
}

void QLowEnergyService_ReadDescriptor(void* ptr, void* descriptor)
{
	static_cast<QLowEnergyService*>(ptr)->readDescriptor(*static_cast<QLowEnergyDescriptor*>(descriptor));
}

struct QtBluetooth_PackedString QLowEnergyService_ServiceName(void* ptr)
{
	return ({ QByteArray tf34544 = static_cast<QLowEnergyService*>(ptr)->serviceName().toUtf8(); QtBluetooth_PackedString { const_cast<char*>(tf34544.prepend("WHITESPACE").constData()+10), tf34544.size()-10 }; });
}

void* QLowEnergyService_ServiceUuid(void* ptr)
{
	return new QBluetoothUuid(static_cast<QLowEnergyService*>(ptr)->serviceUuid());
}

long long QLowEnergyService_State(void* ptr)
{
	return static_cast<QLowEnergyService*>(ptr)->state();
}

long long QLowEnergyService_Type(void* ptr)
{
	return static_cast<QLowEnergyService*>(ptr)->type();
}

void QLowEnergyService_WriteCharacteristic(void* ptr, void* characteristic, void* newValue, long long mode)
{
	static_cast<QLowEnergyService*>(ptr)->writeCharacteristic(*static_cast<QLowEnergyCharacteristic*>(characteristic), *static_cast<QByteArray*>(newValue), static_cast<QLowEnergyService::WriteMode>(mode));
}

void QLowEnergyService_WriteDescriptor(void* ptr, void* descriptor, void* newValue)
{
	static_cast<QLowEnergyService*>(ptr)->writeDescriptor(*static_cast<QLowEnergyDescriptor*>(descriptor), *static_cast<QByteArray*>(newValue));
}

void QLowEnergyService_DestroyQLowEnergyService(void* ptr)
{
	static_cast<QLowEnergyService*>(ptr)->~QLowEnergyService();
}

void* QLowEnergyService_characteristics_atList(void* ptr, int i)
{
	return new QLowEnergyCharacteristic(static_cast<QList<QLowEnergyCharacteristic>*>(ptr)->at(i));
}

void* QLowEnergyService_includedServices_atList(void* ptr, int i)
{
	return new QBluetoothUuid(static_cast<QList<QBluetoothUuid>*>(ptr)->at(i));
}

void QLowEnergyService_TimerEvent(void* ptr, void* event)
{
	static_cast<QLowEnergyService*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLowEnergyService_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLowEnergyService_ChildEvent(void* ptr, void* event)
{
	static_cast<QLowEnergyService*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLowEnergyService_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::childEvent(static_cast<QChildEvent*>(event));
}

void QLowEnergyService_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QLowEnergyService*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLowEnergyService_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLowEnergyService_CustomEvent(void* ptr, void* event)
{
	static_cast<QLowEnergyService*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLowEnergyService_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::customEvent(static_cast<QEvent*>(event));
}

void QLowEnergyService_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLowEnergyService*>(ptr), "deleteLater");
}

void QLowEnergyService_DeleteLaterDefault(void* ptr)
{
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::deleteLater();
}

void QLowEnergyService_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QLowEnergyService*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLowEnergyService_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QLowEnergyService_Event(void* ptr, void* e)
{
	return static_cast<QLowEnergyService*>(ptr)->event(static_cast<QEvent*>(e));
}

char QLowEnergyService_EventDefault(void* ptr, void* e)
{
	return static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::event(static_cast<QEvent*>(e));
}

char QLowEnergyService_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QLowEnergyService*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QLowEnergyService_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QLowEnergyService_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLowEnergyService*>(ptr)->metaObject());
}

void* QLowEnergyService_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLowEnergyService*>(ptr)->QLowEnergyService::metaObject());
}

void* QLowEnergyServiceData_NewQLowEnergyServiceData()
{
	return new QLowEnergyServiceData();
}

void* QLowEnergyServiceData_NewQLowEnergyServiceData2(void* other)
{
	return new QLowEnergyServiceData(*static_cast<QLowEnergyServiceData*>(other));
}

void QLowEnergyServiceData_AddCharacteristic(void* ptr, void* characteristic)
{
	static_cast<QLowEnergyServiceData*>(ptr)->addCharacteristic(*static_cast<QLowEnergyCharacteristicData*>(characteristic));
}

void QLowEnergyServiceData_AddIncludedService(void* ptr, void* service)
{
	static_cast<QLowEnergyServiceData*>(ptr)->addIncludedService(static_cast<QLowEnergyService*>(service));
}

struct QtBluetooth_PackedList QLowEnergyServiceData_Characteristics(void* ptr)
{
	return ({ QList<QLowEnergyCharacteristicData>* tmpValue = new QList<QLowEnergyCharacteristicData>(static_cast<QLowEnergyServiceData*>(ptr)->characteristics()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtBluetooth_PackedList QLowEnergyServiceData_IncludedServices(void* ptr)
{
	return ({ QList<QLowEnergyService *>* tmpValue = new QList<QLowEnergyService *>(static_cast<QLowEnergyServiceData*>(ptr)->includedServices()); QtBluetooth_PackedList { tmpValue, tmpValue->size() }; });
}

char QLowEnergyServiceData_IsValid(void* ptr)
{
	return static_cast<QLowEnergyServiceData*>(ptr)->isValid();
}

void QLowEnergyServiceData_SetType(void* ptr, long long ty)
{
	static_cast<QLowEnergyServiceData*>(ptr)->setType(static_cast<QLowEnergyServiceData::ServiceType>(ty));
}

void QLowEnergyServiceData_SetUuid(void* ptr, void* uuid)
{
	static_cast<QLowEnergyServiceData*>(ptr)->setUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void QLowEnergyServiceData_Swap(void* ptr, void* other)
{
	static_cast<QLowEnergyServiceData*>(ptr)->swap(*static_cast<QLowEnergyServiceData*>(other));
}

long long QLowEnergyServiceData_Type(void* ptr)
{
	return static_cast<QLowEnergyServiceData*>(ptr)->type();
}

void* QLowEnergyServiceData_Uuid(void* ptr)
{
	return new QBluetoothUuid(static_cast<QLowEnergyServiceData*>(ptr)->uuid());
}

void QLowEnergyServiceData_DestroyQLowEnergyServiceData(void* ptr)
{
	static_cast<QLowEnergyServiceData*>(ptr)->~QLowEnergyServiceData();
}

void* QLowEnergyServiceData_characteristics_atList(void* ptr, int i)
{
	return new QLowEnergyCharacteristicData(static_cast<QList<QLowEnergyCharacteristicData>*>(ptr)->at(i));
}

void* QLowEnergyServiceData_includedServices_atList(void* ptr, int i)
{
	return const_cast<QLowEnergyService*>(static_cast<QList<QLowEnergyService *>*>(ptr)->at(i));
}

