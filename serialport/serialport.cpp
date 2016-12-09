// +build !minimal

#define protected public
#define private public

#include "serialport.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QIODevice>
#include <QList>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QSerialPort>
#include <QSerialPortInfo>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>

class MyQSerialPort: public QSerialPort
{
public:
	MyQSerialPort(QObject *parent) : QSerialPort(parent) {};
	MyQSerialPort(const QSerialPortInfo &serialPortInfo, QObject *parent) : QSerialPort(serialPortInfo, parent) {};
	MyQSerialPort(const QString &name, QObject *parent) : QSerialPort(name, parent) {};
	void Signal_BaudRateChanged(qint32 baudRate, QSerialPort::Directions directions) { callbackQSerialPort_BaudRateChanged(this, baudRate, directions); };
	void Signal_BreakEnabledChanged(bool set) { callbackQSerialPort_BreakEnabledChanged(this, set); };
	void Signal_DataBitsChanged(QSerialPort::DataBits dataBits) { callbackQSerialPort_DataBitsChanged(this, dataBits); };
	void Signal_DataTerminalReadyChanged(bool set) { callbackQSerialPort_DataTerminalReadyChanged(this, set); };
	void Signal_Error2(QSerialPort::SerialPortError error) { callbackQSerialPort_Error2(this, error); };
	void Signal_FlowControlChanged(QSerialPort::FlowControl flow) { callbackQSerialPort_FlowControlChanged(this, flow); };
	void Signal_ParityChanged(QSerialPort::Parity parity) { callbackQSerialPort_ParityChanged(this, parity); };
	void Signal_RequestToSendChanged(bool set) { callbackQSerialPort_RequestToSendChanged(this, set); };
	void Signal_StopBitsChanged(QSerialPort::StopBits stopBits) { callbackQSerialPort_StopBitsChanged(this, stopBits); };
	 ~MyQSerialPort() { callbackQSerialPort_DestroyQSerialPort(this); };
	qint64 pos() const { return callbackQSerialPort_Pos(const_cast<MyQSerialPort*>(this)); };
	bool reset() { return callbackQSerialPort_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQSerialPort_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQSerialPort_Size(const_cast<MyQSerialPort*>(this)); };
	void timerEvent(QTimerEvent * event) { callbackQSerialPort_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSerialPort_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSerialPort_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSerialPort_CustomEvent(this, event); };
	void deleteLater() { callbackQSerialPort_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSerialPort_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSerialPort_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSerialPort_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSerialPort_MetaObject(const_cast<MyQSerialPort*>(this))); };
};

int QSerialPort_BaudRate(void* ptr, long long directions)
{
	return static_cast<QSerialPort*>(ptr)->baudRate(static_cast<QSerialPort::Direction>(directions));
}

void QSerialPort_ClearError(void* ptr)
{
	static_cast<QSerialPort*>(ptr)->clearError();
}

long long QSerialPort_DataBits(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->dataBits();
}

long long QSerialPort_Error(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->error();
}

long long QSerialPort_FlowControl(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->flowControl();
}

char QSerialPort_IsBreakEnabled(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->isBreakEnabled();
}

char QSerialPort_IsDataTerminalReady(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->isDataTerminalReady();
}

char QSerialPort_IsRequestToSend(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->isRequestToSend();
}

long long QSerialPort_Parity(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->parity();
}

char QSerialPort_SetBaudRate(void* ptr, int baudRate, long long directions)
{
	return static_cast<QSerialPort*>(ptr)->setBaudRate(baudRate, static_cast<QSerialPort::Direction>(directions));
}

char QSerialPort_SetBreakEnabled(void* ptr, char set)
{
	return static_cast<QSerialPort*>(ptr)->setBreakEnabled(set != 0);
}

char QSerialPort_SetDataBits(void* ptr, long long dataBits)
{
	return static_cast<QSerialPort*>(ptr)->setDataBits(static_cast<QSerialPort::DataBits>(dataBits));
}

char QSerialPort_SetDataTerminalReady(void* ptr, char set)
{
	return static_cast<QSerialPort*>(ptr)->setDataTerminalReady(set != 0);
}

char QSerialPort_SetFlowControl(void* ptr, long long flowControl)
{
	return static_cast<QSerialPort*>(ptr)->setFlowControl(static_cast<QSerialPort::FlowControl>(flowControl));
}

char QSerialPort_SetParity(void* ptr, long long parity)
{
	return static_cast<QSerialPort*>(ptr)->setParity(static_cast<QSerialPort::Parity>(parity));
}

char QSerialPort_SetRequestToSend(void* ptr, char set)
{
	return static_cast<QSerialPort*>(ptr)->setRequestToSend(set != 0);
}

char QSerialPort_SetStopBits(void* ptr, long long stopBits)
{
	return static_cast<QSerialPort*>(ptr)->setStopBits(static_cast<QSerialPort::StopBits>(stopBits));
}

long long QSerialPort_StopBits(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->stopBits();
}

void* QSerialPort_NewQSerialPort(void* parent)
{
	return new MyQSerialPort(static_cast<QObject*>(parent));
}

void* QSerialPort_NewQSerialPort3(void* serialPortInfo, void* parent)
{
	return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QObject*>(parent));
}

void* QSerialPort_NewQSerialPort2(char* name, void* parent)
{
	return new MyQSerialPort(QString(name), static_cast<QObject*>(parent));
}

char QSerialPort_AtEnd(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->atEnd();
}

void QSerialPort_ConnectBaudRateChanged(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(qint32, QSerialPort::Directions)>(&QSerialPort::baudRateChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(qint32, QSerialPort::Directions)>(&MyQSerialPort::Signal_BaudRateChanged));
}

void QSerialPort_DisconnectBaudRateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(qint32, QSerialPort::Directions)>(&QSerialPort::baudRateChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(qint32, QSerialPort::Directions)>(&MyQSerialPort::Signal_BaudRateChanged));
}

void QSerialPort_BaudRateChanged(void* ptr, int baudRate, long long directions)
{
	static_cast<QSerialPort*>(ptr)->baudRateChanged(baudRate, static_cast<QSerialPort::Direction>(directions));
}

void QSerialPort_ConnectBreakEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::breakEnabledChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_BreakEnabledChanged));
}

void QSerialPort_DisconnectBreakEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::breakEnabledChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_BreakEnabledChanged));
}

void QSerialPort_BreakEnabledChanged(void* ptr, char set)
{
	static_cast<QSerialPort*>(ptr)->breakEnabledChanged(set != 0);
}

long long QSerialPort_BytesAvailable(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->bytesAvailable();
}

long long QSerialPort_BytesToWrite(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->bytesToWrite();
}

char QSerialPort_CanReadLine(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->canReadLine();
}

char QSerialPort_Clear(void* ptr, long long directions)
{
	return static_cast<QSerialPort*>(ptr)->clear(static_cast<QSerialPort::Direction>(directions));
}

void QSerialPort_Close(void* ptr)
{
	static_cast<QSerialPort*>(ptr)->close();
}

void QSerialPort_ConnectDataBitsChanged(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::DataBits)>(&QSerialPort::dataBitsChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::DataBits)>(&MyQSerialPort::Signal_DataBitsChanged));
}

void QSerialPort_DisconnectDataBitsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::DataBits)>(&QSerialPort::dataBitsChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::DataBits)>(&MyQSerialPort::Signal_DataBitsChanged));
}

void QSerialPort_DataBitsChanged(void* ptr, long long dataBits)
{
	static_cast<QSerialPort*>(ptr)->dataBitsChanged(static_cast<QSerialPort::DataBits>(dataBits));
}

void QSerialPort_ConnectDataTerminalReadyChanged(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::dataTerminalReadyChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_DataTerminalReadyChanged));
}

void QSerialPort_DisconnectDataTerminalReadyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::dataTerminalReadyChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_DataTerminalReadyChanged));
}

void QSerialPort_DataTerminalReadyChanged(void* ptr, char set)
{
	static_cast<QSerialPort*>(ptr)->dataTerminalReadyChanged(set != 0);
}

void QSerialPort_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::SerialPortError)>(&QSerialPort::error), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::SerialPortError)>(&MyQSerialPort::Signal_Error2));
}

void QSerialPort_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::SerialPortError)>(&QSerialPort::error), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::SerialPortError)>(&MyQSerialPort::Signal_Error2));
}

void QSerialPort_Error2(void* ptr, long long error)
{
	static_cast<QSerialPort*>(ptr)->error(static_cast<QSerialPort::SerialPortError>(error));
}

void QSerialPort_ConnectFlowControlChanged(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::FlowControl)>(&QSerialPort::flowControlChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::FlowControl)>(&MyQSerialPort::Signal_FlowControlChanged));
}

void QSerialPort_DisconnectFlowControlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::FlowControl)>(&QSerialPort::flowControlChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::FlowControl)>(&MyQSerialPort::Signal_FlowControlChanged));
}

void QSerialPort_FlowControlChanged(void* ptr, long long flow)
{
	static_cast<QSerialPort*>(ptr)->flowControlChanged(static_cast<QSerialPort::FlowControl>(flow));
}

char QSerialPort_Flush(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->flush();
}

char QSerialPort_IsSequential(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->isSequential();
}

char QSerialPort_Open(void* ptr, long long mode)
{
	return static_cast<QSerialPort*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

void QSerialPort_ConnectParityChanged(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::Parity)>(&QSerialPort::parityChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::Parity)>(&MyQSerialPort::Signal_ParityChanged));
}

void QSerialPort_DisconnectParityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::Parity)>(&QSerialPort::parityChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::Parity)>(&MyQSerialPort::Signal_ParityChanged));
}

void QSerialPort_ParityChanged(void* ptr, long long parity)
{
	static_cast<QSerialPort*>(ptr)->parityChanged(static_cast<QSerialPort::Parity>(parity));
}

long long QSerialPort_PinoutSignals(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->pinoutSignals();
}

struct QtSerialPort_PackedString QSerialPort_PortName(void* ptr)
{
	return ({ QByteArray t212a26 = static_cast<QSerialPort*>(ptr)->portName().toUtf8(); QtSerialPort_PackedString { const_cast<char*>(t212a26.prepend("WHITESPACE").constData()+10), t212a26.size()-10 }; });
}

long long QSerialPort_ReadBufferSize(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->readBufferSize();
}

long long QSerialPort_ReadLineData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QSerialPort*>(ptr)->readLineData(data, maxSize);
}

void QSerialPort_ConnectRequestToSendChanged(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::requestToSendChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_RequestToSendChanged));
}

void QSerialPort_DisconnectRequestToSendChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::requestToSendChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_RequestToSendChanged));
}

void QSerialPort_RequestToSendChanged(void* ptr, char set)
{
	static_cast<QSerialPort*>(ptr)->requestToSendChanged(set != 0);
}

void QSerialPort_SetPort(void* ptr, void* serialPortInfo)
{
	static_cast<QSerialPort*>(ptr)->setPort(*static_cast<QSerialPortInfo*>(serialPortInfo));
}

void QSerialPort_SetPortName(void* ptr, char* name)
{
	static_cast<QSerialPort*>(ptr)->setPortName(QString(name));
}

void QSerialPort_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QSerialPort*>(ptr)->setReadBufferSize(size);
}

void QSerialPort_ConnectStopBitsChanged(void* ptr)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::StopBits)>(&QSerialPort::stopBitsChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::StopBits)>(&MyQSerialPort::Signal_StopBitsChanged));
}

void QSerialPort_DisconnectStopBitsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::StopBits)>(&QSerialPort::stopBitsChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::StopBits)>(&MyQSerialPort::Signal_StopBitsChanged));
}

void QSerialPort_StopBitsChanged(void* ptr, long long stopBits)
{
	static_cast<QSerialPort*>(ptr)->stopBitsChanged(static_cast<QSerialPort::StopBits>(stopBits));
}

char QSerialPort_WaitForBytesWritten(void* ptr, int msecs)
{
	return static_cast<QSerialPort*>(ptr)->waitForBytesWritten(msecs);
}

char QSerialPort_WaitForReadyRead(void* ptr, int msecs)
{
	return static_cast<QSerialPort*>(ptr)->waitForReadyRead(msecs);
}

long long QSerialPort_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QSerialPort*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

void QSerialPort_DestroyQSerialPort(void* ptr)
{
	static_cast<QSerialPort*>(ptr)->~QSerialPort();
}

void QSerialPort_DestroyQSerialPortDefault(void* ptr)
{

}

long long QSerialPort_Pos(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->pos();
}

long long QSerialPort_PosDefault(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->QSerialPort::pos();
}

char QSerialPort_Reset(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->reset();
}

char QSerialPort_ResetDefault(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->QSerialPort::reset();
}

char QSerialPort_Seek(void* ptr, long long pos)
{
	return static_cast<QSerialPort*>(ptr)->seek(pos);
}

char QSerialPort_SeekDefault(void* ptr, long long pos)
{
	return static_cast<QSerialPort*>(ptr)->QSerialPort::seek(pos);
}

long long QSerialPort_Size(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->size();
}

long long QSerialPort_SizeDefault(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->QSerialPort::size();
}

void QSerialPort_TimerEvent(void* ptr, void* event)
{
	static_cast<QSerialPort*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSerialPort_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSerialPort*>(ptr)->QSerialPort::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSerialPort_ChildEvent(void* ptr, void* event)
{
	static_cast<QSerialPort*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSerialPort_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSerialPort*>(ptr)->QSerialPort::childEvent(static_cast<QChildEvent*>(event));
}

void QSerialPort_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSerialPort*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSerialPort_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSerialPort*>(ptr)->QSerialPort::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSerialPort_CustomEvent(void* ptr, void* event)
{
	static_cast<QSerialPort*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSerialPort_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSerialPort*>(ptr)->QSerialPort::customEvent(static_cast<QEvent*>(event));
}

void QSerialPort_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSerialPort*>(ptr), "deleteLater");
}

void QSerialPort_DeleteLaterDefault(void* ptr)
{
	static_cast<QSerialPort*>(ptr)->QSerialPort::deleteLater();
}

void QSerialPort_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSerialPort*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSerialPort_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSerialPort*>(ptr)->QSerialPort::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSerialPort_Event(void* ptr, void* e)
{
	return static_cast<QSerialPort*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSerialPort_EventDefault(void* ptr, void* e)
{
	return static_cast<QSerialPort*>(ptr)->QSerialPort::event(static_cast<QEvent*>(e));
}

char QSerialPort_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSerialPort*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSerialPort_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSerialPort*>(ptr)->QSerialPort::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSerialPort_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSerialPort*>(ptr)->metaObject());
}

void* QSerialPort_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSerialPort*>(ptr)->QSerialPort::metaObject());
}

void QSerialPortInfo_Swap(void* ptr, void* other)
{
	static_cast<QSerialPortInfo*>(ptr)->swap(*static_cast<QSerialPortInfo*>(other));
}

void* QSerialPortInfo_NewQSerialPortInfo()
{
	return new QSerialPortInfo();
}

void* QSerialPortInfo_NewQSerialPortInfo2(void* port)
{
	return new QSerialPortInfo(*static_cast<QSerialPort*>(port));
}

void* QSerialPortInfo_NewQSerialPortInfo4(void* other)
{
	return new QSerialPortInfo(*static_cast<QSerialPortInfo*>(other));
}

void* QSerialPortInfo_NewQSerialPortInfo3(char* name)
{
	return new QSerialPortInfo(QString(name));
}

struct QtSerialPort_PackedString QSerialPortInfo_Description(void* ptr)
{
	return ({ QByteArray t2edb91 = static_cast<QSerialPortInfo*>(ptr)->description().toUtf8(); QtSerialPort_PackedString { const_cast<char*>(t2edb91.prepend("WHITESPACE").constData()+10), t2edb91.size()-10 }; });
}

char QSerialPortInfo_HasProductIdentifier(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->hasProductIdentifier();
}

char QSerialPortInfo_HasVendorIdentifier(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->hasVendorIdentifier();
}

char QSerialPortInfo_IsNull(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->isNull();
}

struct QtSerialPort_PackedString QSerialPortInfo_Manufacturer(void* ptr)
{
	return ({ QByteArray ta0f36d = static_cast<QSerialPortInfo*>(ptr)->manufacturer().toUtf8(); QtSerialPort_PackedString { const_cast<char*>(ta0f36d.prepend("WHITESPACE").constData()+10), ta0f36d.size()-10 }; });
}

struct QtSerialPort_PackedString QSerialPortInfo_PortName(void* ptr)
{
	return ({ QByteArray tea4d96 = static_cast<QSerialPortInfo*>(ptr)->portName().toUtf8(); QtSerialPort_PackedString { const_cast<char*>(tea4d96.prepend("WHITESPACE").constData()+10), tea4d96.size()-10 }; });
}

unsigned short QSerialPortInfo_ProductIdentifier(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->productIdentifier();
}

struct QtSerialPort_PackedString QSerialPortInfo_SerialNumber(void* ptr)
{
	return ({ QByteArray tb01b4d = static_cast<QSerialPortInfo*>(ptr)->serialNumber().toUtf8(); QtSerialPort_PackedString { const_cast<char*>(tb01b4d.prepend("WHITESPACE").constData()+10), tb01b4d.size()-10 }; });
}

struct QtSerialPort_PackedString QSerialPortInfo_SystemLocation(void* ptr)
{
	return ({ QByteArray te47b34 = static_cast<QSerialPortInfo*>(ptr)->systemLocation().toUtf8(); QtSerialPort_PackedString { const_cast<char*>(te47b34.prepend("WHITESPACE").constData()+10), te47b34.size()-10 }; });
}

unsigned short QSerialPortInfo_VendorIdentifier(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->vendorIdentifier();
}

void QSerialPortInfo_DestroyQSerialPortInfo(void* ptr)
{
	static_cast<QSerialPortInfo*>(ptr)->~QSerialPortInfo();
}

struct QtSerialPort_PackedList QSerialPortInfo_QSerialPortInfo_AvailablePorts()
{
	return ({ QList<QSerialPortInfo>* tmpValue = new QList<QSerialPortInfo>(QSerialPortInfo::availablePorts()); QtSerialPort_PackedList { tmpValue, tmpValue->size() }; });
}

char QSerialPortInfo_IsBusy(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->isBusy();
}

void* QSerialPortInfo_availablePorts_atList(void* ptr, int i)
{
	return new QSerialPortInfo(static_cast<QList<QSerialPortInfo>*>(ptr)->at(i));
}

