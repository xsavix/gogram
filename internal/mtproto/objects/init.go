package objects

import (
	"github.com/xsavix/gogram/internal/encoding/tl"
)

func init() {
	tl.RegisterObjects(
		&ReqPQParams{},
		&ReqDHParamsParams{},
		&SetClientDHParamsParams{},
		&PingParams{},
		&PingDelayDisconnectParams{},
		&ResPQ{},
		&PQInnerData{},
		&PQInnerDataTempDc{},
		&ServerDHParamsFail{},
		&ServerDHParamsOk{},
		&ServerDHInnerData{},
		&ClientDHInnerData{},
		&DHGenOk{},
		&DHGenRetry{},
		&DHGenFail{},
		&RpcResult{},
		&RpcError{},
		&RpcAnswerUnknown{},
		&RpcAnswerDroppedRunning{},
		&RpcAnswerDropped{},
		&FutureSalt{},
		&FutureSalts{},
		&Pong{},
		&NewSessionCreated{},
		&MessageContainer{},
		&MsgCopy{},
		&GzipPacked{},
		&MsgsAck{},
		&BadMsgNotification{},
		&BadServerSalt{},
		&MsgResendReq{},
		&MsgsStateReq{},
		&MsgsStateInfo{},
		&MsgsAllInfo{},
		&MsgsDetailedInfo{},
		&MsgsNewDetailedInfo{},
	)
}
