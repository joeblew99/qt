// +build !minimal

#define protected public
#define private public

#include "nfc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QList>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNdefFilter>
#include <QNdefMessage>
#include <QNdefNfcSmartPosterRecord>
#include <QNdefNfcTextRecord>
#include <QNdefNfcUriRecord>
#include <QNdefRecord>
#include <QNearFieldManager>
#include <QNearFieldShareManager>
#include <QNearFieldShareTarget>
#include <QNearFieldTarget>
#include <QObject>
#include <QQmlNdefRecord>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>

void* QNdefFilter_NewQNdefFilter()
{
	return new QNdefFilter();
}

void* QNdefFilter_NewQNdefFilter2(void* other)
{
	return new QNdefFilter(*static_cast<QNdefFilter*>(other));
}

void QNdefFilter_AppendRecord2(void* ptr, long long typeNameFormat, void* ty, unsigned int min, unsigned int max)
{
	static_cast<QNdefFilter*>(ptr)->appendRecord(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat), *static_cast<QByteArray*>(ty), min, max);
}

void QNdefFilter_Clear(void* ptr)
{
	static_cast<QNdefFilter*>(ptr)->clear();
}

char QNdefFilter_OrderMatch(void* ptr)
{
	return static_cast<QNdefFilter*>(ptr)->orderMatch();
}

int QNdefFilter_RecordCount(void* ptr)
{
	return static_cast<QNdefFilter*>(ptr)->recordCount();
}

void QNdefFilter_SetOrderMatch(void* ptr, char on)
{
	static_cast<QNdefFilter*>(ptr)->setOrderMatch(on != 0);
}

void QNdefFilter_DestroyQNdefFilter(void* ptr)
{
	static_cast<QNdefFilter*>(ptr)->~QNdefFilter();
}

void* QNdefMessage_NewQNdefMessage()
{
	return new QNdefMessage();
}

void* QNdefMessage_NewQNdefMessage3(void* message)
{
	return new QNdefMessage(*static_cast<QNdefMessage*>(message));
}

void* QNdefMessage_NewQNdefMessage2(void* record)
{
	return new QNdefMessage(*static_cast<QNdefRecord*>(record));
}

void* QNdefMessage_QNdefMessage_FromByteArray(void* message)
{
	return new QNdefMessage(QNdefMessage::fromByteArray(*static_cast<QByteArray*>(message)));
}

void* QNdefMessage_ToByteArray(void* ptr)
{
	return new QByteArray(static_cast<QNdefMessage*>(ptr)->toByteArray());
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord()
{
	return new QNdefNfcSmartPosterRecord();
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(void* other)
{
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefNfcSmartPosterRecord*>(other));
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(void* other)
{
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefRecord*>(other));
}

long long QNdefNfcSmartPosterRecord_Action(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->action();
}

void QNdefNfcSmartPosterRecord_AddIcon2(void* ptr, void* ty, void* data)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addIcon(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(data));
}

char QNdefNfcSmartPosterRecord_AddTitle(void* ptr, void* text)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

char QNdefNfcSmartPosterRecord_AddTitle2(void* ptr, char* text, char* locale, long long encoding)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(QString(text), QString(locale), static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

char QNdefNfcSmartPosterRecord_HasAction(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasAction();
}

char QNdefNfcSmartPosterRecord_HasIcon(void* ptr, void* mimetype)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasIcon(*static_cast<QByteArray*>(mimetype));
}

char QNdefNfcSmartPosterRecord_HasSize(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasSize();
}

char QNdefNfcSmartPosterRecord_HasTitle(void* ptr, char* locale)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTitle(QString(locale));
}

char QNdefNfcSmartPosterRecord_HasTypeInfo(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTypeInfo();
}

void* QNdefNfcSmartPosterRecord_Icon(void* ptr, void* mimetype)
{
	return new QByteArray(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->icon(*static_cast<QByteArray*>(mimetype)));
}

int QNdefNfcSmartPosterRecord_IconCount(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->iconCount();
}

char QNdefNfcSmartPosterRecord_RemoveIcon2(void* ptr, void* ty)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeIcon(*static_cast<QByteArray*>(ty));
}

char QNdefNfcSmartPosterRecord_RemoveTitle(void* ptr, void* text)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

char QNdefNfcSmartPosterRecord_RemoveTitle2(void* ptr, char* locale)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(QString(locale));
}

void QNdefNfcSmartPosterRecord_SetAction(void* ptr, long long act)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setAction(static_cast<QNdefNfcSmartPosterRecord::Action>(act));
}

void QNdefNfcSmartPosterRecord_SetSize(void* ptr, unsigned int size)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setSize(size);
}

void QNdefNfcSmartPosterRecord_SetTypeInfo(void* ptr, void* ty)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setTypeInfo(*static_cast<QByteArray*>(ty));
}

void QNdefNfcSmartPosterRecord_SetUri(void* ptr, void* url)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QNdefNfcUriRecord*>(url));
}

void QNdefNfcSmartPosterRecord_SetUri2(void* ptr, void* url)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QUrl*>(url));
}

unsigned int QNdefNfcSmartPosterRecord_Size(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->size();
}

struct QtNfc_PackedString QNdefNfcSmartPosterRecord_Title(void* ptr, char* locale)
{
	return ({ QByteArray t827ac1 = static_cast<QNdefNfcSmartPosterRecord*>(ptr)->title(QString(locale)).toUtf8(); QtNfc_PackedString { const_cast<char*>(t827ac1.prepend("WHITESPACE").constData()+10), t827ac1.size()-10 }; });
}

int QNdefNfcSmartPosterRecord_TitleCount(void* ptr)
{
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->titleCount();
}

void* QNdefNfcSmartPosterRecord_TitleRecord(void* ptr, int index)
{
	return new QNdefNfcTextRecord(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->titleRecord(index));
}

void* QNdefNfcSmartPosterRecord_TypeInfo(void* ptr)
{
	return new QByteArray(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->typeInfo());
}

void* QNdefNfcSmartPosterRecord_Uri(void* ptr)
{
	return new QUrl(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->uri());
}

void* QNdefNfcSmartPosterRecord_UriRecord(void* ptr)
{
	return new QNdefNfcUriRecord(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->uriRecord());
}

void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(void* ptr)
{
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->~QNdefNfcSmartPosterRecord();
}

void* QNdefNfcSmartPosterRecord_titleRecords_atList(void* ptr, int i)
{
	return new QNdefNfcTextRecord(static_cast<QList<QNdefNfcTextRecord>*>(ptr)->at(i));
}

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord()
{
	return new QNdefNfcTextRecord();
}

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord2(void* other)
{
	return new QNdefNfcTextRecord(*static_cast<QNdefRecord*>(other));
}

long long QNdefNfcTextRecord_Encoding(void* ptr)
{
	return static_cast<QNdefNfcTextRecord*>(ptr)->encoding();
}

struct QtNfc_PackedString QNdefNfcTextRecord_Locale(void* ptr)
{
	return ({ QByteArray tb8fb0c = static_cast<QNdefNfcTextRecord*>(ptr)->locale().toUtf8(); QtNfc_PackedString { const_cast<char*>(tb8fb0c.prepend("WHITESPACE").constData()+10), tb8fb0c.size()-10 }; });
}

void QNdefNfcTextRecord_SetEncoding(void* ptr, long long encoding)
{
	static_cast<QNdefNfcTextRecord*>(ptr)->setEncoding(static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

void QNdefNfcTextRecord_SetLocale(void* ptr, char* locale)
{
	static_cast<QNdefNfcTextRecord*>(ptr)->setLocale(QString(locale));
}

void QNdefNfcTextRecord_SetText(void* ptr, char* text)
{
	static_cast<QNdefNfcTextRecord*>(ptr)->setText(QString(text));
}

struct QtNfc_PackedString QNdefNfcTextRecord_Text(void* ptr)
{
	return ({ QByteArray tbe9455 = static_cast<QNdefNfcTextRecord*>(ptr)->text().toUtf8(); QtNfc_PackedString { const_cast<char*>(tbe9455.prepend("WHITESPACE").constData()+10), tbe9455.size()-10 }; });
}

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord()
{
	return new QNdefNfcUriRecord();
}

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord2(void* other)
{
	return new QNdefNfcUriRecord(*static_cast<QNdefRecord*>(other));
}

void QNdefNfcUriRecord_SetUri(void* ptr, void* uri)
{
	static_cast<QNdefNfcUriRecord*>(ptr)->setUri(*static_cast<QUrl*>(uri));
}

void* QNdefNfcUriRecord_Uri(void* ptr)
{
	return new QUrl(static_cast<QNdefNfcUriRecord*>(ptr)->uri());
}

void* QNdefRecord_NewQNdefRecord()
{
	return new QNdefRecord();
}

void* QNdefRecord_NewQNdefRecord2(void* other)
{
	return new QNdefRecord(*static_cast<QNdefRecord*>(other));
}

void* QNdefRecord_Id(void* ptr)
{
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->id());
}

char QNdefRecord_IsEmpty(void* ptr)
{
	return static_cast<QNdefRecord*>(ptr)->isEmpty();
}

void* QNdefRecord_Payload(void* ptr)
{
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->payload());
}

void QNdefRecord_SetId(void* ptr, void* id)
{
	static_cast<QNdefRecord*>(ptr)->setId(*static_cast<QByteArray*>(id));
}

void QNdefRecord_SetPayload(void* ptr, void* payload)
{
	static_cast<QNdefRecord*>(ptr)->setPayload(*static_cast<QByteArray*>(payload));
}

void QNdefRecord_SetType(void* ptr, void* ty)
{
	static_cast<QNdefRecord*>(ptr)->setType(*static_cast<QByteArray*>(ty));
}

void QNdefRecord_SetTypeNameFormat(void* ptr, long long typeNameFormat)
{
	static_cast<QNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat));
}

void* QNdefRecord_Type(void* ptr)
{
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->type());
}

long long QNdefRecord_TypeNameFormat(void* ptr)
{
	return static_cast<QNdefRecord*>(ptr)->typeNameFormat();
}

void QNdefRecord_DestroyQNdefRecord(void* ptr)
{
	static_cast<QNdefRecord*>(ptr)->~QNdefRecord();
}

class MyQNearFieldManager: public QNearFieldManager
{
public:
	MyQNearFieldManager(QObject *parent) : QNearFieldManager(parent) {};
	void Signal_TargetDetected(QNearFieldTarget * target) { callbackQNearFieldManager_TargetDetected(this, target); };
	void Signal_TargetLost(QNearFieldTarget * target) { callbackQNearFieldManager_TargetLost(this, target); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldManager_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNearFieldManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNearFieldManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNearFieldManager_CustomEvent(this, event); };
	void deleteLater() { callbackQNearFieldManager_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNearFieldManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNearFieldManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNearFieldManager_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNearFieldManager_MetaObject(const_cast<MyQNearFieldManager*>(this))); };
};

int QNearFieldManager_RegisterNdefMessageHandler(void* ptr, void* object, char* method)
{
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QObject*>(object), const_cast<const char*>(method));
}

char QNearFieldManager_StartTargetDetection(void* ptr)
{
	return static_cast<QNearFieldManager*>(ptr)->startTargetDetection();
}

void* QNearFieldManager_NewQNearFieldManager(void* parent)
{
	return new MyQNearFieldManager(static_cast<QObject*>(parent));
}

char QNearFieldManager_IsAvailable(void* ptr)
{
	return static_cast<QNearFieldManager*>(ptr)->isAvailable();
}

int QNearFieldManager_RegisterNdefMessageHandler2(void* ptr, long long typeNameFormat, void* ty, void* object, char* method)
{
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat), *static_cast<QByteArray*>(ty), static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_RegisterNdefMessageHandler3(void* ptr, void* filter, void* object, char* method)
{
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(*static_cast<QNdefFilter*>(filter), static_cast<QObject*>(object), const_cast<const char*>(method));
}

void QNearFieldManager_SetTargetAccessModes(void* ptr, long long accessModes)
{
	static_cast<QNearFieldManager*>(ptr)->setTargetAccessModes(static_cast<QNearFieldManager::TargetAccessMode>(accessModes));
}

void QNearFieldManager_StopTargetDetection(void* ptr)
{
	static_cast<QNearFieldManager*>(ptr)->stopTargetDetection();
}

long long QNearFieldManager_TargetAccessModes(void* ptr)
{
	return static_cast<QNearFieldManager*>(ptr)->targetAccessModes();
}

void QNearFieldManager_ConnectTargetDetected(void* ptr)
{
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));
}

void QNearFieldManager_DisconnectTargetDetected(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));
}

void QNearFieldManager_TargetDetected(void* ptr, void* target)
{
	static_cast<QNearFieldManager*>(ptr)->targetDetected(static_cast<QNearFieldTarget*>(target));
}

void QNearFieldManager_ConnectTargetLost(void* ptr)
{
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));
}

void QNearFieldManager_DisconnectTargetLost(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));
}

void QNearFieldManager_TargetLost(void* ptr, void* target)
{
	static_cast<QNearFieldManager*>(ptr)->targetLost(static_cast<QNearFieldTarget*>(target));
}

char QNearFieldManager_UnregisterNdefMessageHandler(void* ptr, int handlerId)
{
	return static_cast<QNearFieldManager*>(ptr)->unregisterNdefMessageHandler(handlerId);
}

void QNearFieldManager_DestroyQNearFieldManager(void* ptr)
{
	static_cast<QNearFieldManager*>(ptr)->~QNearFieldManager();
}

void QNearFieldManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QNearFieldManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QNearFieldManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNearFieldManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QNearFieldManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNearFieldManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::customEvent(static_cast<QEvent*>(event));
}

void QNearFieldManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNearFieldManager*>(ptr), "deleteLater");
}

void QNearFieldManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::deleteLater();
}

void QNearFieldManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNearFieldManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QNearFieldManager_Event(void* ptr, void* e)
{
	return static_cast<QNearFieldManager*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNearFieldManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::event(static_cast<QEvent*>(e));
}

char QNearFieldManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNearFieldManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNearFieldManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNearFieldManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNearFieldManager*>(ptr)->metaObject());
}

void* QNearFieldManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::metaObject());
}

class MyQNearFieldShareManager: public QNearFieldShareManager
{
public:
	MyQNearFieldShareManager(QObject *parent) : QNearFieldShareManager(parent) {};
	void Signal_Error(QNearFieldShareManager::ShareError error) { callbackQNearFieldShareManager_Error(this, error); };
	void Signal_ShareModesChanged(QNearFieldShareManager::ShareModes modes) { callbackQNearFieldShareManager_ShareModesChanged(this, modes); };
	void Signal_TargetDetected(QNearFieldShareTarget * shareTarget) { callbackQNearFieldShareManager_TargetDetected(this, shareTarget); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldShareManager_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNearFieldShareManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNearFieldShareManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNearFieldShareManager_CustomEvent(this, event); };
	void deleteLater() { callbackQNearFieldShareManager_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNearFieldShareManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNearFieldShareManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNearFieldShareManager_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNearFieldShareManager_MetaObject(const_cast<MyQNearFieldShareManager*>(this))); };
};

void* QNearFieldShareManager_NewQNearFieldShareManager(void* parent)
{
	return new MyQNearFieldShareManager(static_cast<QObject*>(parent));
}

void QNearFieldShareManager_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));
}

void QNearFieldShareManager_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));
}

void QNearFieldShareManager_Error(void* ptr, long long error)
{
	static_cast<QNearFieldShareManager*>(ptr)->error(static_cast<QNearFieldShareManager::ShareError>(error));
}

void QNearFieldShareManager_SetShareModes(void* ptr, long long mode)
{
	static_cast<QNearFieldShareManager*>(ptr)->setShareModes(static_cast<QNearFieldShareManager::ShareMode>(mode));
}

long long QNearFieldShareManager_ShareError(void* ptr)
{
	return static_cast<QNearFieldShareManager*>(ptr)->shareError();
}

long long QNearFieldShareManager_ShareModes(void* ptr)
{
	return static_cast<QNearFieldShareManager*>(ptr)->shareModes();
}

void QNearFieldShareManager_ConnectShareModesChanged(void* ptr)
{
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));
}

void QNearFieldShareManager_DisconnectShareModesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));
}

void QNearFieldShareManager_ShareModesChanged(void* ptr, long long modes)
{
	static_cast<QNearFieldShareManager*>(ptr)->shareModesChanged(static_cast<QNearFieldShareManager::ShareMode>(modes));
}

long long QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes()
{
	return QNearFieldShareManager::supportedShareModes();
}

void QNearFieldShareManager_ConnectTargetDetected(void* ptr)
{
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));
}

void QNearFieldShareManager_DisconnectTargetDetected(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));
}

void QNearFieldShareManager_TargetDetected(void* ptr, void* shareTarget)
{
	static_cast<QNearFieldShareManager*>(ptr)->targetDetected(static_cast<QNearFieldShareTarget*>(shareTarget));
}

void QNearFieldShareManager_DestroyQNearFieldShareManager(void* ptr)
{
	static_cast<QNearFieldShareManager*>(ptr)->~QNearFieldShareManager();
}

void QNearFieldShareManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QNearFieldShareManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldShareManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldShareManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QNearFieldShareManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNearFieldShareManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QNearFieldShareManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNearFieldShareManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::customEvent(static_cast<QEvent*>(event));
}

void QNearFieldShareManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNearFieldShareManager*>(ptr), "deleteLater");
}

void QNearFieldShareManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::deleteLater();
}

void QNearFieldShareManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNearFieldShareManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QNearFieldShareManager_Event(void* ptr, void* e)
{
	return static_cast<QNearFieldShareManager*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNearFieldShareManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::event(static_cast<QEvent*>(e));
}

char QNearFieldShareManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNearFieldShareManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNearFieldShareManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNearFieldShareManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNearFieldShareManager*>(ptr)->metaObject());
}

void* QNearFieldShareManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::metaObject());
}

class MyQNearFieldShareTarget: public QNearFieldShareTarget
{
public:
	void Signal_Error(QNearFieldShareManager::ShareError error) { callbackQNearFieldShareTarget_Error(this, error); };
	void Signal_ShareFinished() { callbackQNearFieldShareTarget_ShareFinished(this); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldShareTarget_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNearFieldShareTarget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNearFieldShareTarget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNearFieldShareTarget_CustomEvent(this, event); };
	void deleteLater() { callbackQNearFieldShareTarget_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNearFieldShareTarget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNearFieldShareTarget_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNearFieldShareTarget_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNearFieldShareTarget_MetaObject(const_cast<MyQNearFieldShareTarget*>(this))); };
};

void QNearFieldShareTarget_Cancel(void* ptr)
{
	static_cast<QNearFieldShareTarget*>(ptr)->cancel();
}

void QNearFieldShareTarget_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareTarget::error), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareTarget::Signal_Error));
}

void QNearFieldShareTarget_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareTarget::error), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareTarget::Signal_Error));
}

void QNearFieldShareTarget_Error(void* ptr, long long error)
{
	static_cast<QNearFieldShareTarget*>(ptr)->error(static_cast<QNearFieldShareManager::ShareError>(error));
}

char QNearFieldShareTarget_IsShareInProgress(void* ptr)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->isShareInProgress();
}

char QNearFieldShareTarget_Share(void* ptr, void* message)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->share(*static_cast<QNdefMessage*>(message));
}

long long QNearFieldShareTarget_ShareError(void* ptr)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->shareError();
}

void QNearFieldShareTarget_ConnectShareFinished(void* ptr)
{
	QObject::connect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)()>(&QNearFieldShareTarget::shareFinished), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)()>(&MyQNearFieldShareTarget::Signal_ShareFinished));
}

void QNearFieldShareTarget_DisconnectShareFinished(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)()>(&QNearFieldShareTarget::shareFinished), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)()>(&MyQNearFieldShareTarget::Signal_ShareFinished));
}

void QNearFieldShareTarget_ShareFinished(void* ptr)
{
	static_cast<QNearFieldShareTarget*>(ptr)->shareFinished();
}

long long QNearFieldShareTarget_ShareModes(void* ptr)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->shareModes();
}

void QNearFieldShareTarget_DestroyQNearFieldShareTarget(void* ptr)
{
	static_cast<QNearFieldShareTarget*>(ptr)->~QNearFieldShareTarget();
}

void QNearFieldShareTarget_TimerEvent(void* ptr, void* event)
{
	static_cast<QNearFieldShareTarget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldShareTarget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldShareTarget_ChildEvent(void* ptr, void* event)
{
	static_cast<QNearFieldShareTarget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareTarget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareTarget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNearFieldShareTarget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareTarget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareTarget_CustomEvent(void* ptr, void* event)
{
	static_cast<QNearFieldShareTarget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNearFieldShareTarget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::customEvent(static_cast<QEvent*>(event));
}

void QNearFieldShareTarget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNearFieldShareTarget*>(ptr), "deleteLater");
}

void QNearFieldShareTarget_DeleteLaterDefault(void* ptr)
{
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::deleteLater();
}

void QNearFieldShareTarget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNearFieldShareTarget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldShareTarget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QNearFieldShareTarget_Event(void* ptr, void* e)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNearFieldShareTarget_EventDefault(void* ptr, void* e)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::event(static_cast<QEvent*>(e));
}

char QNearFieldShareTarget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNearFieldShareTarget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNearFieldShareTarget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNearFieldShareTarget*>(ptr)->metaObject());
}

void* QNearFieldShareTarget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::metaObject());
}

class MyQNearFieldTarget: public QNearFieldTarget
{
public:
	MyQNearFieldTarget(QObject *parent) : QNearFieldTarget(parent) {};
	AccessMethods accessMethods() const { return static_cast<QNearFieldTarget::AccessMethod>(callbackQNearFieldTarget_AccessMethods(const_cast<MyQNearFieldTarget*>(this))); };
	void Signal_Disconnected() { callbackQNearFieldTarget_Disconnected(this); };
	bool hasNdefMessage() { return callbackQNearFieldTarget_HasNdefMessage(this) != 0; };
	void Signal_NdefMessageRead(const QNdefMessage & message) { callbackQNearFieldTarget_NdefMessageRead(this, const_cast<QNdefMessage*>(&message)); };
	void Signal_NdefMessagesWritten() { callbackQNearFieldTarget_NdefMessagesWritten(this); };
	Type type() const { return static_cast<QNearFieldTarget::Type>(callbackQNearFieldTarget_Type(const_cast<MyQNearFieldTarget*>(this))); };
	QByteArray uid() const { return *static_cast<QByteArray*>(callbackQNearFieldTarget_Uid(const_cast<MyQNearFieldTarget*>(this))); };
	QUrl url() const { return *static_cast<QUrl*>(callbackQNearFieldTarget_Url(const_cast<MyQNearFieldTarget*>(this))); };
	 ~MyQNearFieldTarget() { callbackQNearFieldTarget_DestroyQNearFieldTarget(this); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldTarget_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNearFieldTarget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNearFieldTarget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNearFieldTarget_CustomEvent(this, event); };
	void deleteLater() { callbackQNearFieldTarget_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNearFieldTarget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNearFieldTarget_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNearFieldTarget_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNearFieldTarget_MetaObject(const_cast<MyQNearFieldTarget*>(this))); };
};

void* QNearFieldTarget_NewQNearFieldTarget(void* parent)
{
	return new MyQNearFieldTarget(static_cast<QObject*>(parent));
}

long long QNearFieldTarget_AccessMethods(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->accessMethods();
}

void QNearFieldTarget_ConnectDisconnected(void* ptr)
{
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));
}

void QNearFieldTarget_DisconnectDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));
}

void QNearFieldTarget_Disconnected(void* ptr)
{
	static_cast<QNearFieldTarget*>(ptr)->disconnected();
}

char QNearFieldTarget_HasNdefMessage(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->hasNdefMessage();
}

char QNearFieldTarget_HasNdefMessageDefault(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::hasNdefMessage();
}

char QNearFieldTarget_IsProcessingCommand(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->isProcessingCommand();
}

void QNearFieldTarget_ConnectNdefMessageRead(void* ptr)
{
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)(const QNdefMessage &)>(&QNearFieldTarget::ndefMessageRead), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)(const QNdefMessage &)>(&MyQNearFieldTarget::Signal_NdefMessageRead));
}

void QNearFieldTarget_DisconnectNdefMessageRead(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)(const QNdefMessage &)>(&QNearFieldTarget::ndefMessageRead), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)(const QNdefMessage &)>(&MyQNearFieldTarget::Signal_NdefMessageRead));
}

void QNearFieldTarget_NdefMessageRead(void* ptr, void* message)
{
	static_cast<QNearFieldTarget*>(ptr)->ndefMessageRead(*static_cast<QNdefMessage*>(message));
}

void QNearFieldTarget_ConnectNdefMessagesWritten(void* ptr)
{
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));
}

void QNearFieldTarget_DisconnectNdefMessagesWritten(void* ptr)
{
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));
}

void QNearFieldTarget_NdefMessagesWritten(void* ptr)
{
	static_cast<QNearFieldTarget*>(ptr)->ndefMessagesWritten();
}

long long QNearFieldTarget_Type(void* ptr)
{
	return static_cast<QNearFieldTarget*>(ptr)->type();
}

void* QNearFieldTarget_Uid(void* ptr)
{
	return new QByteArray(static_cast<QNearFieldTarget*>(ptr)->uid());
}

void* QNearFieldTarget_Url(void* ptr)
{
	return new QUrl(static_cast<QNearFieldTarget*>(ptr)->url());
}

void* QNearFieldTarget_UrlDefault(void* ptr)
{
	return new QUrl(static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::url());
}

void QNearFieldTarget_DestroyQNearFieldTarget(void* ptr)
{
	static_cast<QNearFieldTarget*>(ptr)->~QNearFieldTarget();
}

void QNearFieldTarget_DestroyQNearFieldTargetDefault(void* ptr)
{

}

void QNearFieldTarget_TimerEvent(void* ptr, void* event)
{
	static_cast<QNearFieldTarget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldTarget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldTarget_ChildEvent(void* ptr, void* event)
{
	static_cast<QNearFieldTarget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldTarget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldTarget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNearFieldTarget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldTarget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldTarget_CustomEvent(void* ptr, void* event)
{
	static_cast<QNearFieldTarget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNearFieldTarget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::customEvent(static_cast<QEvent*>(event));
}

void QNearFieldTarget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNearFieldTarget*>(ptr), "deleteLater");
}

void QNearFieldTarget_DeleteLaterDefault(void* ptr)
{
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::deleteLater();
}

void QNearFieldTarget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNearFieldTarget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNearFieldTarget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QNearFieldTarget_Event(void* ptr, void* e)
{
	return static_cast<QNearFieldTarget*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNearFieldTarget_EventDefault(void* ptr, void* e)
{
	return static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::event(static_cast<QEvent*>(e));
}

char QNearFieldTarget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNearFieldTarget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNearFieldTarget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNearFieldTarget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNearFieldTarget*>(ptr)->metaObject());
}

void* QNearFieldTarget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::metaObject());
}

class MyQQmlNdefRecord: public QQmlNdefRecord
{
public:
	MyQQmlNdefRecord(QObject *parent) : QQmlNdefRecord(parent) {};
	MyQQmlNdefRecord(const QNdefRecord &record, QObject *parent) : QQmlNdefRecord(record, parent) {};
	void Signal_RecordChanged() { callbackQQmlNdefRecord_RecordChanged(this); };
	void Signal_TypeChanged() { callbackQQmlNdefRecord_TypeChanged(this); };
	void Signal_TypeNameFormatChanged() { callbackQQmlNdefRecord_TypeNameFormatChanged(this); };
	void timerEvent(QTimerEvent * event) { callbackQQmlNdefRecord_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQQmlNdefRecord_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlNdefRecord_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQmlNdefRecord_CustomEvent(this, event); };
	void deleteLater() { callbackQQmlNdefRecord_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlNdefRecord_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQmlNdefRecord_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlNdefRecord_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlNdefRecord_MetaObject(const_cast<MyQQmlNdefRecord*>(this))); };
};

long long QQmlNdefRecord_TypeNameFormat(void* ptr)
{
	return static_cast<QQmlNdefRecord*>(ptr)->typeNameFormat();
}

void* QQmlNdefRecord_NewQQmlNdefRecord(void* parent)
{
	return new MyQQmlNdefRecord(static_cast<QObject*>(parent));
}

void* QQmlNdefRecord_NewQQmlNdefRecord2(void* record, void* parent)
{
	return new MyQQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QObject*>(parent));
}

void* QQmlNdefRecord_Record(void* ptr)
{
	return new QNdefRecord(static_cast<QQmlNdefRecord*>(ptr)->record());
}

void QQmlNdefRecord_ConnectRecordChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));
}

void QQmlNdefRecord_DisconnectRecordChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));
}

void QQmlNdefRecord_RecordChanged(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->recordChanged();
}

void QQmlNdefRecord_SetRecord(void* ptr, void* record)
{
	static_cast<QQmlNdefRecord*>(ptr)->setRecord(*static_cast<QNdefRecord*>(record));
}

void QQmlNdefRecord_SetType(void* ptr, char* newtype)
{
	static_cast<QQmlNdefRecord*>(ptr)->setType(QString(newtype));
}

void QQmlNdefRecord_SetTypeNameFormat(void* ptr, long long newTypeNameFormat)
{
	static_cast<QQmlNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QQmlNdefRecord::TypeNameFormat>(newTypeNameFormat));
}

struct QtNfc_PackedString QQmlNdefRecord_Type(void* ptr)
{
	return ({ QByteArray t9a31a1 = static_cast<QQmlNdefRecord*>(ptr)->type().toUtf8(); QtNfc_PackedString { const_cast<char*>(t9a31a1.prepend("WHITESPACE").constData()+10), t9a31a1.size()-10 }; });
}

void QQmlNdefRecord_ConnectTypeChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));
}

void QQmlNdefRecord_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));
}

void QQmlNdefRecord_TypeChanged(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->typeChanged();
}

void QQmlNdefRecord_ConnectTypeNameFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));
}

void QQmlNdefRecord_DisconnectTypeNameFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));
}

void QQmlNdefRecord_TypeNameFormatChanged(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->typeNameFormatChanged();
}

void QQmlNdefRecord_DestroyQQmlNdefRecord(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->~QQmlNdefRecord();
}

void QQmlNdefRecord_TimerEvent(void* ptr, void* event)
{
	static_cast<QQmlNdefRecord*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlNdefRecord_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlNdefRecord_ChildEvent(void* ptr, void* event)
{
	static_cast<QQmlNdefRecord*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlNdefRecord_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlNdefRecord_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlNdefRecord*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlNdefRecord_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlNdefRecord_CustomEvent(void* ptr, void* event)
{
	static_cast<QQmlNdefRecord*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlNdefRecord_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::customEvent(static_cast<QEvent*>(event));
}

void QQmlNdefRecord_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlNdefRecord*>(ptr), "deleteLater");
}

void QQmlNdefRecord_DeleteLaterDefault(void* ptr)
{
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::deleteLater();
}

void QQmlNdefRecord_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQmlNdefRecord*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlNdefRecord_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QQmlNdefRecord_Event(void* ptr, void* e)
{
	return static_cast<QQmlNdefRecord*>(ptr)->event(static_cast<QEvent*>(e));
}

char QQmlNdefRecord_EventDefault(void* ptr, void* e)
{
	return static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::event(static_cast<QEvent*>(e));
}

char QQmlNdefRecord_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlNdefRecord*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QQmlNdefRecord_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlNdefRecord_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlNdefRecord*>(ptr)->metaObject());
}

void* QQmlNdefRecord_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::metaObject());
}

