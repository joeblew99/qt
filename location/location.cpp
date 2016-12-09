// +build !minimal

#define protected public
#define private public

#include "location.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGeoCoordinate>
#include <QGeoManeuver>
#include <QGeoRectangle>
#include <QGeoRoute>
#include <QGeoRouteReply>
#include <QGeoRouteRequest>
#include <QGeoRouteSegment>
#include <QGeoRoutingManager>
#include <QGeoRoutingManagerEngine>
#include <QGeoServiceProvider>
#include <QGeoServiceProviderFactory>
#include <QList>
#include <QLocale>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>

void* QGeoManeuver_NewQGeoManeuver()
{
	return new QGeoManeuver();
}

void* QGeoManeuver_NewQGeoManeuver2(void* other)
{
	return new QGeoManeuver(*static_cast<QGeoManeuver*>(other));
}

long long QGeoManeuver_Direction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->direction();
}

double QGeoManeuver_DistanceToNextInstruction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->distanceToNextInstruction();
}

struct QtLocation_PackedString QGeoManeuver_InstructionText(void* ptr)
{
	return ({ QByteArray t4cc4b3 = static_cast<QGeoManeuver*>(ptr)->instructionText().toUtf8(); QtLocation_PackedString { const_cast<char*>(t4cc4b3.prepend("WHITESPACE").constData()+10), t4cc4b3.size()-10 }; });
}

char QGeoManeuver_IsValid(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->isValid();
}

void* QGeoManeuver_Position(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoManeuver*>(ptr)->position());
}

void QGeoManeuver_SetDirection(void* ptr, long long direction)
{
	static_cast<QGeoManeuver*>(ptr)->setDirection(static_cast<QGeoManeuver::InstructionDirection>(direction));
}

void QGeoManeuver_SetDistanceToNextInstruction(void* ptr, double distance)
{
	static_cast<QGeoManeuver*>(ptr)->setDistanceToNextInstruction(distance);
}

void QGeoManeuver_SetInstructionText(void* ptr, char* instructionText)
{
	static_cast<QGeoManeuver*>(ptr)->setInstructionText(QString(instructionText));
}

void QGeoManeuver_SetPosition(void* ptr, void* position)
{
	static_cast<QGeoManeuver*>(ptr)->setPosition(*static_cast<QGeoCoordinate*>(position));
}

void QGeoManeuver_SetTimeToNextInstruction(void* ptr, int secs)
{
	static_cast<QGeoManeuver*>(ptr)->setTimeToNextInstruction(secs);
}

void QGeoManeuver_SetWaypoint(void* ptr, void* coordinate)
{
	static_cast<QGeoManeuver*>(ptr)->setWaypoint(*static_cast<QGeoCoordinate*>(coordinate));
}

int QGeoManeuver_TimeToNextInstruction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->timeToNextInstruction();
}

void* QGeoManeuver_Waypoint(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoManeuver*>(ptr)->waypoint());
}

void QGeoManeuver_DestroyQGeoManeuver(void* ptr)
{
	static_cast<QGeoManeuver*>(ptr)->~QGeoManeuver();
}

void* QGeoRoute_NewQGeoRoute()
{
	return new QGeoRoute();
}

void* QGeoRoute_NewQGeoRoute2(void* other)
{
	return new QGeoRoute(*static_cast<QGeoRoute*>(other));
}

void* QGeoRoute_Bounds(void* ptr)
{
	return new QGeoRectangle(static_cast<QGeoRoute*>(ptr)->bounds());
}

double QGeoRoute_Distance(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->distance();
}

void* QGeoRoute_FirstRouteSegment(void* ptr)
{
	return new QGeoRouteSegment(static_cast<QGeoRoute*>(ptr)->firstRouteSegment());
}

struct QtLocation_PackedList QGeoRoute_Path(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue = new QList<QGeoCoordinate>(static_cast<QGeoRoute*>(ptr)->path()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

void* QGeoRoute_Request(void* ptr)
{
	return new QGeoRouteRequest(static_cast<QGeoRoute*>(ptr)->request());
}

struct QtLocation_PackedString QGeoRoute_RouteId(void* ptr)
{
	return ({ QByteArray t7492be = static_cast<QGeoRoute*>(ptr)->routeId().toUtf8(); QtLocation_PackedString { const_cast<char*>(t7492be.prepend("WHITESPACE").constData()+10), t7492be.size()-10 }; });
}

void QGeoRoute_SetBounds(void* ptr, void* bounds)
{
	static_cast<QGeoRoute*>(ptr)->setBounds(*static_cast<QGeoRectangle*>(bounds));
}

void QGeoRoute_SetDistance(void* ptr, double distance)
{
	static_cast<QGeoRoute*>(ptr)->setDistance(distance);
}

void QGeoRoute_SetFirstRouteSegment(void* ptr, void* routeSegment)
{
	static_cast<QGeoRoute*>(ptr)->setFirstRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRoute_SetRequest(void* ptr, void* request)
{
	static_cast<QGeoRoute*>(ptr)->setRequest(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoute_SetRouteId(void* ptr, char* id)
{
	static_cast<QGeoRoute*>(ptr)->setRouteId(QString(id));
}

void QGeoRoute_SetTravelMode(void* ptr, long long mode)
{
	static_cast<QGeoRoute*>(ptr)->setTravelMode(static_cast<QGeoRouteRequest::TravelMode>(mode));
}

void QGeoRoute_SetTravelTime(void* ptr, int secs)
{
	static_cast<QGeoRoute*>(ptr)->setTravelTime(secs);
}

long long QGeoRoute_TravelMode(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->travelMode();
}

int QGeoRoute_TravelTime(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->travelTime();
}

void QGeoRoute_DestroyQGeoRoute(void* ptr)
{
	static_cast<QGeoRoute*>(ptr)->~QGeoRoute();
}

void* QGeoRoute_path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

class MyQGeoRouteReply: public QGeoRouteReply
{
public:
	MyQGeoRouteReply(Error error, const QString &errorString, QObject *parent) : QGeoRouteReply(error, errorString, parent) {};
	MyQGeoRouteReply(const QGeoRouteRequest &request, QObject *parent) : QGeoRouteReply(request, parent) {};
	void abort() { callbackQGeoRouteReply_Abort(this); };
	void Signal_Error2(QGeoRouteReply::Error error, const QString & errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQGeoRouteReply_Error2(this, error, errorStringPacked); };
	void Signal_Finished() { callbackQGeoRouteReply_Finished(this); };
	 ~MyQGeoRouteReply() { callbackQGeoRouteReply_DestroyQGeoRouteReply(this); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRouteReply_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQGeoRouteReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRouteReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRouteReply_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRouteReply_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRouteReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoRouteReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRouteReply_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRouteReply_MetaObject(const_cast<MyQGeoRouteReply*>(this))); };
};

void* QGeoRouteReply_NewQGeoRouteReply(long long error, char* errorString, void* parent)
{
	return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString(errorString), static_cast<QObject*>(parent));
}

void* QGeoRouteReply_NewQGeoRouteReply2(void* request, void* parent)
{
	return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QObject*>(parent));
}

void QGeoRouteReply_Abort(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->abort();
}

void QGeoRouteReply_AbortDefault(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::abort();
}

void QGeoRouteReply_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&QGeoRouteReply::error), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&MyQGeoRouteReply::Signal_Error2));
}

void QGeoRouteReply_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&QGeoRouteReply::error), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&MyQGeoRouteReply::Signal_Error2));
}

void QGeoRouteReply_Error2(void* ptr, long long error, char* errorString)
{
	static_cast<QGeoRouteReply*>(ptr)->error(static_cast<QGeoRouteReply::Error>(error), QString(errorString));
}

long long QGeoRouteReply_Error(void* ptr)
{
	return static_cast<QGeoRouteReply*>(ptr)->error();
}

struct QtLocation_PackedString QGeoRouteReply_ErrorString(void* ptr)
{
	return ({ QByteArray t834aee = static_cast<QGeoRouteReply*>(ptr)->errorString().toUtf8(); QtLocation_PackedString { const_cast<char*>(t834aee.prepend("WHITESPACE").constData()+10), t834aee.size()-10 }; });
}

void QGeoRouteReply_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));
}

void QGeoRouteReply_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));
}

void QGeoRouteReply_Finished(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->finished();
}

char QGeoRouteReply_IsFinished(void* ptr)
{
	return static_cast<QGeoRouteReply*>(ptr)->isFinished();
}

void* QGeoRouteReply_Request(void* ptr)
{
	return new QGeoRouteRequest(static_cast<QGeoRouteReply*>(ptr)->request());
}

struct QtLocation_PackedList QGeoRouteReply_Routes(void* ptr)
{
	return ({ QList<QGeoRoute>* tmpValue = new QList<QGeoRoute>(static_cast<QGeoRouteReply*>(ptr)->routes()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

void QGeoRouteReply_SetError(void* ptr, long long error, char* errorString)
{
	static_cast<QGeoRouteReply*>(ptr)->setError(static_cast<QGeoRouteReply::Error>(error), QString(errorString));
}

void QGeoRouteReply_SetFinished(void* ptr, char finished)
{
	static_cast<QGeoRouteReply*>(ptr)->setFinished(finished != 0);
}

void QGeoRouteReply_DestroyQGeoRouteReply(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->~QGeoRouteReply();
}

void QGeoRouteReply_DestroyQGeoRouteReplyDefault(void* ptr)
{

}

void* QGeoRouteReply_routes_atList(void* ptr, int i)
{
	return new QGeoRoute(static_cast<QList<QGeoRoute>*>(ptr)->at(i));
}

void QGeoRouteReply_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoRouteReply*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoRouteReply_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoRouteReply_ChildEvent(void* ptr, void* event)
{
	static_cast<QGeoRouteReply*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRouteReply_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRouteReply_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoRouteReply*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRouteReply_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRouteReply_CustomEvent(void* ptr, void* event)
{
	static_cast<QGeoRouteReply*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoRouteReply_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::customEvent(static_cast<QEvent*>(event));
}

void QGeoRouteReply_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoRouteReply*>(ptr), "deleteLater");
}

void QGeoRouteReply_DeleteLaterDefault(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::deleteLater();
}

void QGeoRouteReply_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoRouteReply*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRouteReply_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoRouteReply_Event(void* ptr, void* e)
{
	return static_cast<QGeoRouteReply*>(ptr)->event(static_cast<QEvent*>(e));
}

char QGeoRouteReply_EventDefault(void* ptr, void* e)
{
	return static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::event(static_cast<QEvent*>(e));
}

char QGeoRouteReply_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoRouteReply*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGeoRouteReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoRouteReply_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoRouteReply*>(ptr)->metaObject());
}

void* QGeoRouteReply_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::metaObject());
}

void* QGeoRouteRequest_NewQGeoRouteRequest2(void* origin, void* destination)
{
	return new QGeoRouteRequest(*static_cast<QGeoCoordinate*>(origin), *static_cast<QGeoCoordinate*>(destination));
}

void* QGeoRouteRequest_NewQGeoRouteRequest3(void* other)
{
	return new QGeoRouteRequest(*static_cast<QGeoRouteRequest*>(other));
}

struct QtLocation_PackedList QGeoRouteRequest_ExcludeAreas(void* ptr)
{
	return ({ QList<QGeoRectangle>* tmpValue = new QList<QGeoRectangle>(static_cast<QGeoRouteRequest*>(ptr)->excludeAreas()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

long long QGeoRouteRequest_FeatureWeight(void* ptr, long long featureType)
{
	return static_cast<QGeoRouteRequest*>(ptr)->featureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType));
}

long long QGeoRouteRequest_ManeuverDetail(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->maneuverDetail();
}

int QGeoRouteRequest_NumberAlternativeRoutes(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->numberAlternativeRoutes();
}

long long QGeoRouteRequest_RouteOptimization(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->routeOptimization();
}

long long QGeoRouteRequest_SegmentDetail(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->segmentDetail();
}

void QGeoRouteRequest_SetFeatureWeight(void* ptr, long long featureType, long long featureWeight)
{
	static_cast<QGeoRouteRequest*>(ptr)->setFeatureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType), static_cast<QGeoRouteRequest::FeatureWeight>(featureWeight));
}

void QGeoRouteRequest_SetManeuverDetail(void* ptr, long long maneuverDetail)
{
	static_cast<QGeoRouteRequest*>(ptr)->setManeuverDetail(static_cast<QGeoRouteRequest::ManeuverDetail>(maneuverDetail));
}

void QGeoRouteRequest_SetNumberAlternativeRoutes(void* ptr, int alternatives)
{
	static_cast<QGeoRouteRequest*>(ptr)->setNumberAlternativeRoutes(alternatives);
}

void QGeoRouteRequest_SetRouteOptimization(void* ptr, long long optimization)
{
	static_cast<QGeoRouteRequest*>(ptr)->setRouteOptimization(static_cast<QGeoRouteRequest::RouteOptimization>(optimization));
}

void QGeoRouteRequest_SetSegmentDetail(void* ptr, long long segmentDetail)
{
	static_cast<QGeoRouteRequest*>(ptr)->setSegmentDetail(static_cast<QGeoRouteRequest::SegmentDetail>(segmentDetail));
}

void QGeoRouteRequest_SetTravelModes(void* ptr, long long travelModes)
{
	static_cast<QGeoRouteRequest*>(ptr)->setTravelModes(static_cast<QGeoRouteRequest::TravelMode>(travelModes));
}

long long QGeoRouteRequest_TravelModes(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->travelModes();
}

struct QtLocation_PackedList QGeoRouteRequest_Waypoints(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue = new QList<QGeoCoordinate>(static_cast<QGeoRouteRequest*>(ptr)->waypoints()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

void QGeoRouteRequest_DestroyQGeoRouteRequest(void* ptr)
{
	static_cast<QGeoRouteRequest*>(ptr)->~QGeoRouteRequest();
}

void* QGeoRouteRequest_excludeAreas_atList(void* ptr, int i)
{
	return new QGeoRectangle(static_cast<QList<QGeoRectangle>*>(ptr)->at(i));
}

void* QGeoRouteRequest_waypoints_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void* QGeoRouteSegment_NewQGeoRouteSegment()
{
	return new QGeoRouteSegment();
}

void* QGeoRouteSegment_NewQGeoRouteSegment2(void* other)
{
	return new QGeoRouteSegment(*static_cast<QGeoRouteSegment*>(other));
}

double QGeoRouteSegment_Distance(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->distance();
}

char QGeoRouteSegment_IsValid(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->isValid();
}

void* QGeoRouteSegment_Maneuver(void* ptr)
{
	return new QGeoManeuver(static_cast<QGeoRouteSegment*>(ptr)->maneuver());
}

void* QGeoRouteSegment_NextRouteSegment(void* ptr)
{
	return new QGeoRouteSegment(static_cast<QGeoRouteSegment*>(ptr)->nextRouteSegment());
}

struct QtLocation_PackedList QGeoRouteSegment_Path(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue = new QList<QGeoCoordinate>(static_cast<QGeoRouteSegment*>(ptr)->path()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

void QGeoRouteSegment_SetDistance(void* ptr, double distance)
{
	static_cast<QGeoRouteSegment*>(ptr)->setDistance(distance);
}

void QGeoRouteSegment_SetManeuver(void* ptr, void* maneuver)
{
	static_cast<QGeoRouteSegment*>(ptr)->setManeuver(*static_cast<QGeoManeuver*>(maneuver));
}

void QGeoRouteSegment_SetNextRouteSegment(void* ptr, void* routeSegment)
{
	static_cast<QGeoRouteSegment*>(ptr)->setNextRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRouteSegment_SetTravelTime(void* ptr, int secs)
{
	static_cast<QGeoRouteSegment*>(ptr)->setTravelTime(secs);
}

int QGeoRouteSegment_TravelTime(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->travelTime();
}

void QGeoRouteSegment_DestroyQGeoRouteSegment(void* ptr)
{
	static_cast<QGeoRouteSegment*>(ptr)->~QGeoRouteSegment();
}

void* QGeoRouteSegment_path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

class MyQGeoRoutingManager: public QGeoRoutingManager
{
public:
	void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQGeoRoutingManager_Error(this, reply, error, errorStringPacked); };
	void Signal_Finished(QGeoRouteReply * reply) { callbackQGeoRoutingManager_Finished(this, reply); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRoutingManager_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQGeoRoutingManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRoutingManager_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRoutingManager_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoRoutingManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRoutingManager_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRoutingManager_MetaObject(const_cast<MyQGeoRoutingManager*>(this))); };
};

void* QGeoRoutingManager_CalculateRoute(void* ptr, void* request)
{
	return static_cast<QGeoRoutingManager*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoutingManager_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));
}

void QGeoRoutingManager_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));
}

void QGeoRoutingManager_Error(void* ptr, void* reply, long long error, char* errorString)
{
	static_cast<QGeoRoutingManager*>(ptr)->error(static_cast<QGeoRouteReply*>(reply), static_cast<QGeoRouteReply::Error>(error), QString(errorString));
}

void QGeoRoutingManager_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));
}

void QGeoRoutingManager_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));
}

void QGeoRoutingManager_Finished(void* ptr, void* reply)
{
	static_cast<QGeoRoutingManager*>(ptr)->finished(static_cast<QGeoRouteReply*>(reply));
}

void* QGeoRoutingManager_Locale(void* ptr)
{
	return new QLocale(static_cast<QGeoRoutingManager*>(ptr)->locale());
}

struct QtLocation_PackedString QGeoRoutingManager_ManagerName(void* ptr)
{
	return ({ QByteArray t48ee82 = static_cast<QGeoRoutingManager*>(ptr)->managerName().toUtf8(); QtLocation_PackedString { const_cast<char*>(t48ee82.prepend("WHITESPACE").constData()+10), t48ee82.size()-10 }; });
}

int QGeoRoutingManager_ManagerVersion(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->managerVersion();
}

long long QGeoRoutingManager_MeasurementSystem(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->measurementSystem();
}

void QGeoRoutingManager_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoRoutingManager*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManager_SetMeasurementSystem(void* ptr, long long system)
{
	static_cast<QGeoRoutingManager*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

long long QGeoRoutingManager_SupportedFeatureTypes(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureTypes();
}

long long QGeoRoutingManager_SupportedFeatureWeights(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureWeights();
}

long long QGeoRoutingManager_SupportedManeuverDetails(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedManeuverDetails();
}

long long QGeoRoutingManager_SupportedRouteOptimizations(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedRouteOptimizations();
}

long long QGeoRoutingManager_SupportedSegmentDetails(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedSegmentDetails();
}

long long QGeoRoutingManager_SupportedTravelModes(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedTravelModes();
}

void* QGeoRoutingManager_UpdateRoute(void* ptr, void* route, void* position)
{
	return static_cast<QGeoRoutingManager*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManager_DestroyQGeoRoutingManager(void* ptr)
{
	static_cast<QGeoRoutingManager*>(ptr)->~QGeoRoutingManager();
}

void QGeoRoutingManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoRoutingManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoRoutingManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoRoutingManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QGeoRoutingManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRoutingManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRoutingManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoRoutingManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QGeoRoutingManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoRoutingManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::customEvent(static_cast<QEvent*>(event));
}

void QGeoRoutingManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoRoutingManager*>(ptr), "deleteLater");
}

void QGeoRoutingManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::deleteLater();
}

void QGeoRoutingManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoRoutingManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoRoutingManager_Event(void* ptr, void* e)
{
	return static_cast<QGeoRoutingManager*>(ptr)->event(static_cast<QEvent*>(e));
}

char QGeoRoutingManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::event(static_cast<QEvent*>(e));
}

char QGeoRoutingManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoRoutingManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGeoRoutingManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoRoutingManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoRoutingManager*>(ptr)->metaObject());
}

void* QGeoRoutingManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::metaObject());
}

class MyQGeoRoutingManagerEngine: public QGeoRoutingManagerEngine
{
public:
	QGeoRouteReply * calculateRoute(const QGeoRouteRequest & request) { return static_cast<QGeoRouteReply*>(callbackQGeoRoutingManagerEngine_CalculateRoute(this, const_cast<QGeoRouteRequest*>(&request))); };
	void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQGeoRoutingManagerEngine_Error(this, reply, error, errorStringPacked); };
	void Signal_Finished(QGeoRouteReply * reply) { callbackQGeoRoutingManagerEngine_Finished(this, reply); };
	QGeoRouteReply * updateRoute(const QGeoRoute & route, const QGeoCoordinate & position) { return static_cast<QGeoRouteReply*>(callbackQGeoRoutingManagerEngine_UpdateRoute(this, const_cast<QGeoRoute*>(&route), const_cast<QGeoCoordinate*>(&position))); };
	 ~MyQGeoRoutingManagerEngine() { callbackQGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(this); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRoutingManagerEngine_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQGeoRoutingManagerEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManagerEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRoutingManagerEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRoutingManagerEngine_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManagerEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoRoutingManagerEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRoutingManagerEngine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRoutingManagerEngine_MetaObject(const_cast<MyQGeoRoutingManagerEngine*>(this))); };
};

void* QGeoRoutingManagerEngine_CalculateRoute(void* ptr, void* request)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoutingManagerEngine_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));
}

void QGeoRoutingManagerEngine_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));
}

void QGeoRoutingManagerEngine_Error(void* ptr, void* reply, long long error, char* errorString)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->error(static_cast<QGeoRouteReply*>(reply), static_cast<QGeoRouteReply::Error>(error), QString(errorString));
}

void QGeoRoutingManagerEngine_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));
}

void QGeoRoutingManagerEngine_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));
}

void QGeoRoutingManagerEngine_Finished(void* ptr, void* reply)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->finished(static_cast<QGeoRouteReply*>(reply));
}

void* QGeoRoutingManagerEngine_Locale(void* ptr)
{
	return new QLocale(static_cast<QGeoRoutingManagerEngine*>(ptr)->locale());
}

struct QtLocation_PackedString QGeoRoutingManagerEngine_ManagerName(void* ptr)
{
	return ({ QByteArray tce87b5 = static_cast<QGeoRoutingManagerEngine*>(ptr)->managerName().toUtf8(); QtLocation_PackedString { const_cast<char*>(tce87b5.prepend("WHITESPACE").constData()+10), tce87b5.size()-10 }; });
}

int QGeoRoutingManagerEngine_ManagerVersion(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->managerVersion();
}

long long QGeoRoutingManagerEngine_MeasurementSystem(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->measurementSystem();
}

void QGeoRoutingManagerEngine_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManagerEngine_SetMeasurementSystem(void* ptr, long long system)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

void QGeoRoutingManagerEngine_SetSupportedFeatureTypes(void* ptr, long long featureTypes)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedFeatureTypes(static_cast<QGeoRouteRequest::FeatureType>(featureTypes));
}

void QGeoRoutingManagerEngine_SetSupportedFeatureWeights(void* ptr, long long featureWeights)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedFeatureWeights(static_cast<QGeoRouteRequest::FeatureWeight>(featureWeights));
}

void QGeoRoutingManagerEngine_SetSupportedManeuverDetails(void* ptr, long long maneuverDetails)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedManeuverDetails(static_cast<QGeoRouteRequest::ManeuverDetail>(maneuverDetails));
}

void QGeoRoutingManagerEngine_SetSupportedRouteOptimizations(void* ptr, long long optimizations)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedRouteOptimizations(static_cast<QGeoRouteRequest::RouteOptimization>(optimizations));
}

void QGeoRoutingManagerEngine_SetSupportedSegmentDetails(void* ptr, long long segmentDetails)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedSegmentDetails(static_cast<QGeoRouteRequest::SegmentDetail>(segmentDetails));
}

void QGeoRoutingManagerEngine_SetSupportedTravelModes(void* ptr, long long travelModes)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedTravelModes(static_cast<QGeoRouteRequest::TravelMode>(travelModes));
}

long long QGeoRoutingManagerEngine_SupportedFeatureTypes(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureTypes();
}

long long QGeoRoutingManagerEngine_SupportedFeatureWeights(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureWeights();
}

long long QGeoRoutingManagerEngine_SupportedManeuverDetails(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedManeuverDetails();
}

long long QGeoRoutingManagerEngine_SupportedRouteOptimizations(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedRouteOptimizations();
}

long long QGeoRoutingManagerEngine_SupportedSegmentDetails(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedSegmentDetails();
}

long long QGeoRoutingManagerEngine_SupportedTravelModes(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedTravelModes();
}

void* QGeoRoutingManagerEngine_UpdateRoute(void* ptr, void* route, void* position)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void* QGeoRoutingManagerEngine_UpdateRouteDefault(void* ptr, void* route, void* position)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(void* ptr)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->~QGeoRoutingManagerEngine();
}

void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngineDefault(void* ptr)
{

}

void QGeoRoutingManagerEngine_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoRoutingManagerEngine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoRoutingManagerEngine_ChildEvent(void* ptr, void* event)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRoutingManagerEngine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRoutingManagerEngine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManagerEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManagerEngine_CustomEvent(void* ptr, void* event)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoRoutingManagerEngine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::customEvent(static_cast<QEvent*>(event));
}

void QGeoRoutingManagerEngine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoRoutingManagerEngine*>(ptr), "deleteLater");
}

void QGeoRoutingManagerEngine_DeleteLaterDefault(void* ptr)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::deleteLater();
}

void QGeoRoutingManagerEngine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManagerEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoRoutingManagerEngine_Event(void* ptr, void* e)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->event(static_cast<QEvent*>(e));
}

char QGeoRoutingManagerEngine_EventDefault(void* ptr, void* e)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::event(static_cast<QEvent*>(e));
}

char QGeoRoutingManagerEngine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGeoRoutingManagerEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoRoutingManagerEngine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoRoutingManagerEngine*>(ptr)->metaObject());
}

void* QGeoRoutingManagerEngine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::metaObject());
}

int QGeoServiceProvider_OnlineGeocodingFeature_Type()
{
	return QGeoServiceProvider::OnlineGeocodingFeature;
}

int QGeoServiceProvider_OfflineGeocodingFeature_Type()
{
	return QGeoServiceProvider::OfflineGeocodingFeature;
}

int QGeoServiceProvider_ReverseGeocodingFeature_Type()
{
	return QGeoServiceProvider::ReverseGeocodingFeature;
}

int QGeoServiceProvider_LocalizedGeocodingFeature_Type()
{
	return QGeoServiceProvider::LocalizedGeocodingFeature;
}

int QGeoServiceProvider_AnyGeocodingFeatures_Type()
{
	return QGeoServiceProvider::AnyGeocodingFeatures;
}

int QGeoServiceProvider_OnlineMappingFeature_Type()
{
	return QGeoServiceProvider::OnlineMappingFeature;
}

int QGeoServiceProvider_OfflineMappingFeature_Type()
{
	return QGeoServiceProvider::OfflineMappingFeature;
}

int QGeoServiceProvider_LocalizedMappingFeature_Type()
{
	return QGeoServiceProvider::LocalizedMappingFeature;
}

int QGeoServiceProvider_AnyMappingFeatures_Type()
{
	return QGeoServiceProvider::AnyMappingFeatures;
}

int QGeoServiceProvider_OnlinePlacesFeature_Type()
{
	return QGeoServiceProvider::OnlinePlacesFeature;
}

int QGeoServiceProvider_OfflinePlacesFeature_Type()
{
	return QGeoServiceProvider::OfflinePlacesFeature;
}

int QGeoServiceProvider_SavePlaceFeature_Type()
{
	return QGeoServiceProvider::SavePlaceFeature;
}

int QGeoServiceProvider_RemovePlaceFeature_Type()
{
	return QGeoServiceProvider::RemovePlaceFeature;
}

int QGeoServiceProvider_SaveCategoryFeature_Type()
{
	return QGeoServiceProvider::SaveCategoryFeature;
}

int QGeoServiceProvider_RemoveCategoryFeature_Type()
{
	return QGeoServiceProvider::RemoveCategoryFeature;
}

int QGeoServiceProvider_PlaceRecommendationsFeature_Type()
{
	return QGeoServiceProvider::PlaceRecommendationsFeature;
}

int QGeoServiceProvider_SearchSuggestionsFeature_Type()
{
	return QGeoServiceProvider::SearchSuggestionsFeature;
}

int QGeoServiceProvider_LocalizedPlacesFeature_Type()
{
	return QGeoServiceProvider::LocalizedPlacesFeature;
}

int QGeoServiceProvider_NotificationsFeature_Type()
{
	return QGeoServiceProvider::NotificationsFeature;
}

int QGeoServiceProvider_PlaceMatchingFeature_Type()
{
	return QGeoServiceProvider::PlaceMatchingFeature;
}

int QGeoServiceProvider_AnyPlacesFeatures_Type()
{
	return QGeoServiceProvider::AnyPlacesFeatures;
}

int QGeoServiceProvider_OnlineRoutingFeature_Type()
{
	return QGeoServiceProvider::OnlineRoutingFeature;
}

int QGeoServiceProvider_OfflineRoutingFeature_Type()
{
	return QGeoServiceProvider::OfflineRoutingFeature;
}

int QGeoServiceProvider_LocalizedRoutingFeature_Type()
{
	return QGeoServiceProvider::LocalizedRoutingFeature;
}

int QGeoServiceProvider_RouteUpdatesFeature_Type()
{
	return QGeoServiceProvider::RouteUpdatesFeature;
}

int QGeoServiceProvider_AlternativeRoutesFeature_Type()
{
	return QGeoServiceProvider::AlternativeRoutesFeature;
}

int QGeoServiceProvider_ExcludeAreasRoutingFeature_Type()
{
	return QGeoServiceProvider::ExcludeAreasRoutingFeature;
}

int QGeoServiceProvider_AnyRoutingFeatures_Type()
{
	return QGeoServiceProvider::AnyRoutingFeatures;
}

struct QtLocation_PackedString QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders()
{
	return ({ QByteArray teec3ca = QGeoServiceProvider::availableServiceProviders().join("|").toUtf8(); QtLocation_PackedString { const_cast<char*>(teec3ca.prepend("WHITESPACE").constData()+10), teec3ca.size()-10 }; });
}

long long QGeoServiceProvider_Error(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->error();
}

struct QtLocation_PackedString QGeoServiceProvider_ErrorString(void* ptr)
{
	return ({ QByteArray t90b712 = static_cast<QGeoServiceProvider*>(ptr)->errorString().toUtf8(); QtLocation_PackedString { const_cast<char*>(t90b712.prepend("WHITESPACE").constData()+10), t90b712.size()-10 }; });
}

long long QGeoServiceProvider_GeocodingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingFeatures();
}

void* QGeoServiceProvider_GeocodingManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingManager();
}

long long QGeoServiceProvider_MappingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->mappingFeatures();
}

void* QGeoServiceProvider_PlaceManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->placeManager();
}

long long QGeoServiceProvider_PlacesFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->placesFeatures();
}

long long QGeoServiceProvider_RoutingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->routingFeatures();
}

void* QGeoServiceProvider_RoutingManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->routingManager();
}

void QGeoServiceProvider_SetAllowExperimental(void* ptr, char allow)
{
	static_cast<QGeoServiceProvider*>(ptr)->setAllowExperimental(allow != 0);
}

void QGeoServiceProvider_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoServiceProvider*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoServiceProvider_DestroyQGeoServiceProvider(void* ptr)
{
	static_cast<QGeoServiceProvider*>(ptr)->~QGeoServiceProvider();
}

void QGeoServiceProvider_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoServiceProvider*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoServiceProvider_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoServiceProvider_ChildEvent(void* ptr, void* event)
{
	static_cast<QGeoServiceProvider*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoServiceProvider_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoServiceProvider_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoServiceProvider*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoServiceProvider_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoServiceProvider_CustomEvent(void* ptr, void* event)
{
	static_cast<QGeoServiceProvider*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoServiceProvider_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::customEvent(static_cast<QEvent*>(event));
}

void QGeoServiceProvider_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoServiceProvider*>(ptr), "deleteLater");
}

void QGeoServiceProvider_DeleteLaterDefault(void* ptr)
{
	static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::deleteLater();
}

void QGeoServiceProvider_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoServiceProvider*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoServiceProvider_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoServiceProvider_Event(void* ptr, void* e)
{
	return static_cast<QGeoServiceProvider*>(ptr)->event(static_cast<QEvent*>(e));
}

char QGeoServiceProvider_EventDefault(void* ptr, void* e)
{
	return static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::event(static_cast<QEvent*>(e));
}

char QGeoServiceProvider_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoServiceProvider*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGeoServiceProvider_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoServiceProvider_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoServiceProvider*>(ptr)->metaObject());
}

void* QGeoServiceProvider_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::metaObject());
}

class MyQGeoServiceProviderFactory: public QGeoServiceProviderFactory
{
public:
	 ~MyQGeoServiceProviderFactory() { callbackQGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(this); };
};

void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(void* ptr)
{
	static_cast<QGeoServiceProviderFactory*>(ptr)->~QGeoServiceProviderFactory();
}

void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactoryDefault(void* ptr)
{

}

