// +build android

#define protected public
#define private public

#include "androidextras_android.h"
#include "_cgo_export.h"

#include <QAndroidActivityResultReceiver>
#include <QAndroidJniEnvironment>
#include <QAndroidJniObject>
#include <QByteArray>
#include <QString>

class MyQAndroidActivityResultReceiver: public QAndroidActivityResultReceiver
{
public:
	void handleActivityResult(int receiverRequestCode, int resultCode, const QAndroidJniObject & data) { callbackQAndroidActivityResultReceiver_HandleActivityResult(this, receiverRequestCode, resultCode, const_cast<QAndroidJniObject*>(&data)); };
};

void QAndroidActivityResultReceiver_HandleActivityResult(void* ptr, int receiverRequestCode, int resultCode, void* data)
{
	static_cast<QAndroidActivityResultReceiver*>(ptr)->handleActivityResult(receiverRequestCode, resultCode, *static_cast<QAndroidJniObject*>(data));
}

void* QAndroidJniEnvironment_NewQAndroidJniEnvironment()
{
	return new QAndroidJniEnvironment();
}

void* QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM()
{
	return QAndroidJniEnvironment::javaVM();
}

void QAndroidJniEnvironment_DestroyQAndroidJniEnvironment(void* ptr)
{
	static_cast<QAndroidJniEnvironment*>(ptr)->~QAndroidJniEnvironment();
}

void* QAndroidJniObject_NewQAndroidJniObject()
{
	return new QAndroidJniObject();
}

void* QAndroidJniObject_NewQAndroidJniObject2(char* className)
{
	return new QAndroidJniObject(const_cast<const char*>(className));
}

void* QAndroidJniObject_NewQAndroidJniObject3(char* className, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(const_cast<const char*>(className), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void* QAndroidJniObject_NewQAndroidJniObject4(void* clazz)
{
	return new QAndroidJniObject(static_cast<jclass>(clazz));
}

void* QAndroidJniObject_NewQAndroidJniObject5(void* clazz, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(static_cast<jclass>(clazz), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void* QAndroidJniObject_NewQAndroidJniObject6(void* object)
{
	return new QAndroidJniObject(static_cast<jobject>(object));
}

int QAndroidJniObject_CallMethodInt(void* ptr, char* methodName)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName));
}

char QAndroidJniObject_CallMethodBoolean(void* ptr, char* methodName)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName));
}

void QAndroidJniObject_CallMethodVoid(void* ptr, char* methodName)
{
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName));
}

int QAndroidJniObject_CallMethodInt2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jint>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_CallMethodBoolean2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return static_cast<QAndroidJniObject*>(ptr)->callMethod<jboolean>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_CallMethodVoid2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	static_cast<QAndroidJniObject*>(ptr)->callMethod<void>(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void* QAndroidJniObject_CallObjectMethod(void* ptr, char* methodName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod<jobject>(const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_CallMethodString(void* ptr, char* methodName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod<jobject>(const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_CallObjectMethod2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_CallMethodString2(void* ptr, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->callObjectMethod(const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(char* className, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jint>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(char* className, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(char* className, char* methodName)
{
	QAndroidJniObject::callStaticMethod<void>(const_cast<const char*>(className), const_cast<const char*>(methodName));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jint>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	QAndroidJniObject::callStaticMethod<void>(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(void* clazz, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jint>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(void* clazz, char* methodName)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(void* clazz, char* methodName)
{
	QAndroidJniObject::callStaticMethod<void>(static_cast<jclass>(clazz), const_cast<const char*>(methodName));
}

int QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jint>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

char QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return QAndroidJniObject::callStaticMethod<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	QAndroidJniObject::callStaticMethod<void>(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v));
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(char* className, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(const_cast<const char*>(className), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString(char* className, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(const_cast<const char*>(className), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2(char* className, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(const_cast<const char*>(className), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(void* clazz, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3(void* clazz, char* methodName)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(methodName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4(void* clazz, char* methodName, char* signature, void* v0, void* v1, void* v2, void* v3, void* v4, void* v5, void* v6, void* v7, void* v8, void* v)
{
	return new QAndroidJniObject(QAndroidJniObject::callStaticObjectMethod(static_cast<jclass>(clazz), const_cast<const char*>(methodName), const_cast<const char*>(signature), static_cast<jobject>(v0), static_cast<jobject>(v1), static_cast<jobject>(v2), static_cast<jobject>(v3), static_cast<jobject>(v4), static_cast<jobject>(v5), static_cast<jobject>(v6), static_cast<jobject>(v7), static_cast<jobject>(v8), static_cast<jobject>(v)).object());
}

void* QAndroidJniObject_QAndroidJniObject_FromLocalRef(void* localRef)
{
	return new QAndroidJniObject(QAndroidJniObject::fromLocalRef(static_cast<jobject>(localRef)).object());
}

void* QAndroidJniObject_QAndroidJniObject_FromString(char* stri)
{
	return new QAndroidJniObject(QAndroidJniObject::fromString(QString(stri)).object());
}

int QAndroidJniObject_GetFieldInt(void* ptr, char* fieldName)
{
	return static_cast<QAndroidJniObject*>(ptr)->getField<jint>(const_cast<const char*>(fieldName));
}

char QAndroidJniObject_GetFieldBoolean(void* ptr, char* fieldName)
{
	return static_cast<QAndroidJniObject*>(ptr)->getField<jboolean>(const_cast<const char*>(fieldName));
}

void* QAndroidJniObject_GetObjectField(void* ptr, char* fieldName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_GetFieldString(void* ptr, char* fieldName)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_GetObjectField2(void* ptr, char* fieldName, char* signature)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_GetFieldString2(void* ptr, char* fieldName, char* signature)
{
	return new QAndroidJniObject(static_cast<QAndroidJniObject*>(ptr)->getObjectField<jobject>(const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(char* className, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jint>(const_cast<const char*>(className), const_cast<const char*>(fieldName));
}

char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(char* className, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jboolean>(const_cast<const char*>(className), const_cast<const char*>(fieldName));
}

int QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(void* clazz, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jint>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName));
}

char QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(void* clazz, char* fieldName)
{
	return QAndroidJniObject::getStaticField<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName));
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(char* className, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString(char* className, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(char* className, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2(char* className, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(void* clazz, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3(void* clazz, char* fieldName)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(void* clazz, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

void* QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4(void* clazz, char* fieldName, char* signature)
{
	return new QAndroidJniObject(QAndroidJniObject::getStaticObjectField<jobject>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature)).object());
}

char QAndroidJniObject_QAndroidJniObject_IsClassAvailable(char* className)
{
	return QAndroidJniObject::isClassAvailable(const_cast<const char*>(className));
}

char QAndroidJniObject_IsValid(void* ptr)
{
	return static_cast<QAndroidJniObject*>(ptr)->isValid();
}

struct QtAndroidExtras_PackedString QAndroidJniObject_ToString(void* ptr)
{
	return ({ QByteArray t3150b4 = static_cast<QAndroidJniObject*>(ptr)->toString().toUtf8(); QtAndroidExtras_PackedString { const_cast<char*>(t3150b4.prepend("WHITESPACE").constData()+10), t3150b4.size()-10 }; });
}

void* QAndroidJniObject_Object(void* ptr)
{
	return static_cast<QAndroidJniObject*>(ptr)->object();
}

void QAndroidJniObject_SetField(void* ptr, char* fieldName, void* value)
{
	static_cast<QAndroidJniObject*>(ptr)->setField(const_cast<const char*>(fieldName), static_cast<jobject>(value));
}

void QAndroidJniObject_SetField2(void* ptr, char* fieldName, char* signature, void* value)
{
	static_cast<QAndroidJniObject*>(ptr)->setField(const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2(char* className, char* fieldName, int value)
{
	QAndroidJniObject::setStaticField<jint>(const_cast<const char*>(className), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2(char* className, char* fieldName, char value)
{
	QAndroidJniObject::setStaticField<jboolean>(const_cast<const char*>(className), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticField(char* className, char* fieldName, char* signature, void* value)
{
	QAndroidJniObject::setStaticField(const_cast<const char*>(className), const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4(void* clazz, char* fieldName, int value)
{
	QAndroidJniObject::setStaticField<jint>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4(void* clazz, char* fieldName, char value)
{
	QAndroidJniObject::setStaticField<jboolean>(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), value);
}

void QAndroidJniObject_QAndroidJniObject_SetStaticField3(void* clazz, char* fieldName, char* signature, void* value)
{
	QAndroidJniObject::setStaticField(static_cast<jclass>(clazz), const_cast<const char*>(fieldName), const_cast<const char*>(signature), static_cast<jobject>(value));
}

void QAndroidJniObject_DestroyQAndroidJniObject(void* ptr)
{
	static_cast<QAndroidJniObject*>(ptr)->~QAndroidJniObject();
}

