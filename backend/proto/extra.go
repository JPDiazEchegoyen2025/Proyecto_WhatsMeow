package armadillo

import (
	"google.golang.org/protobuf/proto"

	"github.com/tuusuario/whatsmeow-backend/proto/instamadilloAddMessage"
	"github.com/tuusuario/whatsmeow-backend/proto/instamadilloDeleteMessage"
	"github.com/tuusuario/whatsmeow-backend/proto/instamadilloSupplementMessage"
	"github.com/tuusuario/whatsmeow-backend/proto/waArmadilloApplication"
	"github.com/tuusuario/whatsmeow-backend/proto/waCommon"
	"github.com/tuusuario/whatsmeow-backend/proto/waConsumerApplication"
	"github.com/tuusuario/whatsmeow-backend/proto/waMultiDevice"
)

type MessageApplicationSub interface {
	IsMessageApplicationSub()
}

type RealMessageApplicationSub interface {
	MessageApplicationSub
	proto.Message
}

type Unsupported_BusinessApplication waCommon.SubProtocol
type Unsupported_PaymentApplication waCommon.SubProtocol
type Unsupported_Voip waCommon.SubProtocol

var (
	_ MessageApplicationSub = (*waConsumerApplication.ConsumerApplication)(nil) // 2
	_ MessageApplicationSub = (*Unsupported_BusinessApplication)(nil)           // 3
	_ MessageApplicationSub = (*Unsupported_PaymentApplication)(nil)            // 4
	_ MessageApplicationSub = (*waMultiDevice.MultiDevice)(nil)                 // 5
	_ MessageApplicationSub = (*Unsupported_Voip)(nil)                          // 6
	_ MessageApplicationSub = (*waArmadilloApplication.Armadillo)(nil)          // 7

	_ MessageApplicationSub = (*instamadilloAddMessage.AddMessagePayload)(nil)
	_ MessageApplicationSub = (*instamadilloSupplementMessage.SupplementMessagePayload)(nil)
	_ MessageApplicationSub = (*instamadilloDeleteMessage.DeleteMessagePayload)(nil)
)

func (*Unsupported_BusinessApplication) IsMessageApplicationSub() {}
func (*Unsupported_PaymentApplication) IsMessageApplicationSub()  {}
func (*Unsupported_Voip) IsMessageApplicationSub()                {}
