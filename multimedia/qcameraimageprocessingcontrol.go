package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraImageProcessingControl struct {
	QMediaControl
}

type QCameraImageProcessingControl_ITF interface {
	QMediaControl_ITF
	QCameraImageProcessingControl_PTR() *QCameraImageProcessingControl
}

func PointerFromQCameraImageProcessingControl(ptr QCameraImageProcessingControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageProcessingControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageProcessingControlFromPointer(ptr unsafe.Pointer) *QCameraImageProcessingControl {
	var n = new(QCameraImageProcessingControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraImageProcessingControl_") {
		n.SetObjectName("QCameraImageProcessingControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraImageProcessingControl) QCameraImageProcessingControl_PTR() *QCameraImageProcessingControl {
	return ptr
}

//QCameraImageProcessingControl::ProcessingParameter
type QCameraImageProcessingControl__ProcessingParameter int64

const (
	QCameraImageProcessingControl__WhiteBalancePreset   = QCameraImageProcessingControl__ProcessingParameter(0)
	QCameraImageProcessingControl__ColorTemperature     = QCameraImageProcessingControl__ProcessingParameter(1)
	QCameraImageProcessingControl__Contrast             = QCameraImageProcessingControl__ProcessingParameter(2)
	QCameraImageProcessingControl__Saturation           = QCameraImageProcessingControl__ProcessingParameter(3)
	QCameraImageProcessingControl__Brightness           = QCameraImageProcessingControl__ProcessingParameter(4)
	QCameraImageProcessingControl__Sharpening           = QCameraImageProcessingControl__ProcessingParameter(5)
	QCameraImageProcessingControl__Denoising            = QCameraImageProcessingControl__ProcessingParameter(6)
	QCameraImageProcessingControl__ContrastAdjustment   = QCameraImageProcessingControl__ProcessingParameter(7)
	QCameraImageProcessingControl__SaturationAdjustment = QCameraImageProcessingControl__ProcessingParameter(8)
	QCameraImageProcessingControl__BrightnessAdjustment = QCameraImageProcessingControl__ProcessingParameter(9)
	QCameraImageProcessingControl__SharpeningAdjustment = QCameraImageProcessingControl__ProcessingParameter(10)
	QCameraImageProcessingControl__DenoisingAdjustment  = QCameraImageProcessingControl__ProcessingParameter(11)
	QCameraImageProcessingControl__ColorFilter          = QCameraImageProcessingControl__ProcessingParameter(12)
	QCameraImageProcessingControl__ExtendedParameter    = QCameraImageProcessingControl__ProcessingParameter(1000)
)

func (ptr *QCameraImageProcessingControl) IsParameterSupported(parameter QCameraImageProcessingControl__ProcessingParameter) bool {
	defer qt.Recovering("QCameraImageProcessingControl::isParameterSupported")

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessingControl_IsParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessingControl) IsParameterValueSupported(parameter QCameraImageProcessingControl__ProcessingParameter, value core.QVariant_ITF) bool {
	defer qt.Recovering("QCameraImageProcessingControl::isParameterValueSupported")

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessingControl_IsParameterValueSupported(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessingControl) Parameter(parameter QCameraImageProcessingControl__ProcessingParameter) *core.QVariant {
	defer qt.Recovering("QCameraImageProcessingControl::parameter")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraImageProcessingControl_Parameter(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraImageProcessingControl) SetParameter(parameter QCameraImageProcessingControl__ProcessingParameter, value core.QVariant_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::setParameter")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_SetParameter(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraImageProcessingControl) DestroyQCameraImageProcessingControl() {
	defer qt.Recovering("QCameraImageProcessingControl::~QCameraImageProcessingControl")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraImageProcessingControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraImageProcessingControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraImageProcessingControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessingControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraImageProcessingControlTimerEvent
func callbackQCameraImageProcessingControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessingControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraImageProcessingControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessingControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraImageProcessingControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraImageProcessingControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessingControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraImageProcessingControlChildEvent
func callbackQCameraImageProcessingControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessingControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraImageProcessingControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessingControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraImageProcessingControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraImageProcessingControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraImageProcessingControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraImageProcessingControlCustomEvent
func callbackQCameraImageProcessingControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageProcessingControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraImageProcessingControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraImageProcessingControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraImageProcessingControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageProcessingControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageProcessingControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}