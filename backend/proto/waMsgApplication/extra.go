package waMsgApplication

import (
	"github.com/tuusuario/whatsmeow-backend/proto/armadilloutil"
	"github.com/tuusuario/whatsmeow-backend/proto/waArmadilloApplication"
	"github.com/tuusuario/whatsmeow-backend/proto/waConsumerApplication"
	"github.com/tuusuario/whatsmeow-backend/proto/waMultiDevice"
)

const (
	ConsumerApplicationVersion    = 1
	ArmadilloApplicationVersion   = 1
	MultiDeviceApplicationVersion = 1 // TODO: check
)

func (msg *MessageApplication_SubProtocolPayload_ConsumerMessage) Decode() (*waConsumerApplication.ConsumerApplication, error) {
	return armadilloutil.Unmarshal(&waConsumerApplication.ConsumerApplication{}, msg.ConsumerMessage, ConsumerApplicationVersion)
}

func (msg *MessageApplication_SubProtocolPayload_ConsumerMessage) Set(payload *waConsumerApplication.ConsumerApplication) (err error) {
	msg.ConsumerMessage, err = armadilloutil.Marshal(payload, ConsumerApplicationVersion)
	return
}

func (msg *MessageApplication_SubProtocolPayload_Armadillo) Decode() (*waArmadilloApplication.Armadillo, error) {
	return armadilloutil.Unmarshal(&waArmadilloApplication.Armadillo{}, msg.Armadillo, ArmadilloApplicationVersion)
}

func (msg *MessageApplication_SubProtocolPayload_Armadillo) Set(payload *waArmadilloApplication.Armadillo) (err error) {
	msg.Armadillo, err = armadilloutil.Marshal(payload, ArmadilloApplicationVersion)
	return
}

func (msg *MessageApplication_SubProtocolPayload_MultiDevice) Decode() (*waMultiDevice.MultiDevice, error) {
	return armadilloutil.Unmarshal(&waMultiDevice.MultiDevice{}, msg.MultiDevice, MultiDeviceApplicationVersion)
}

func (msg *MessageApplication_SubProtocolPayload_MultiDevice) Set(payload *waMultiDevice.MultiDevice) (err error) {
	msg.MultiDevice, err = armadilloutil.Marshal(payload, MultiDeviceApplicationVersion)
	return
}
