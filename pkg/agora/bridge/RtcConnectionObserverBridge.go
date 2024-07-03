package bridge

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../../third_party/agora_rtc_sdk_c/agora_rtc_sdk-prefix/src/agora_rtc_sdk/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../../third_party/agora_rtc_sdk_c/agora_rtc_sdk-prefix/src/agora_rtc_sdk/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "bridge/C_RtcConnectionObserverBridge.h"

void cgo_RtcConnectionObserverBridge_onConnected(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_TConnectionInfo *connectionInfo, enum C_CONNECTION_CHANGED_REASON_TYPE reason);

void cgo_RtcConnectionObserverBridge_onDisconnected(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_TConnectionInfo *connectionInfo, enum C_CONNECTION_CHANGED_REASON_TYPE reason);

void cgo_RtcConnectionObserverBridge_onConnecting(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_TConnectionInfo *connectionInfo, enum C_CONNECTION_CHANGED_REASON_TYPE reason);

void cgo_RtcConnectionObserverBridge_onReconnecting(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_TConnectionInfo *connectionInfo, enum C_CONNECTION_CHANGED_REASON_TYPE reason);

void cgo_RtcConnectionObserverBridge_onReconnected(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_TConnectionInfo *connectionInfo, enum C_CONNECTION_CHANGED_REASON_TYPE reason);

void cgo_RtcConnectionObserverBridge_onConnectionLost(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_TConnectionInfo *connectionInfo);

void cgo_RtcConnectionObserverBridge_onLastmileQuality(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_QUALITY_TYPE quality);

void cgo_RtcConnectionObserverBridge_onLastmileProbeResult(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_LastmileProbeResult *result);

void cgo_RtcConnectionObserverBridge_onTokenPrivilegeWillExpire(C_RtcConnectionObserverBridge *this_, void *userData,
	char *token);

void cgo_RtcConnectionObserverBridge_onTokenPrivilegeDidExpire(C_RtcConnectionObserverBridge *this_, void *userData);

void cgo_RtcConnectionObserverBridge_onConnectionFailure(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_TConnectionInfo *connectionInfo,
	enum C_CONNECTION_CHANGED_REASON_TYPE reason);

void cgo_RtcConnectionObserverBridge_onUserJoined(C_RtcConnectionObserverBridge *this_, void *userData,
	C_user_id_t userId);

void cgo_RtcConnectionObserverBridge_onUserLeft(C_RtcConnectionObserverBridge *this_, void *userData,
	C_user_id_t userId, enum C_USER_OFFLINE_REASON_TYPE reason);

void cgo_RtcConnectionObserverBridge_onTransportStats(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_RtcStats *stats);

void cgo_RtcConnectionObserverBridge_onChangeRoleSuccess(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_CLIENT_ROLE_TYPE oldRole, enum C_CLIENT_ROLE_TYPE newRole, struct C_ClientRoleOptions *newRoleOptions);

void cgo_RtcConnectionObserverBridge_onChangeRoleFailure(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_CLIENT_ROLE_CHANGE_FAILED_REASON reason, enum C_CLIENT_ROLE_TYPE currentRole);

void cgo_RtcConnectionObserverBridge_onLicenseValidationFailure(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_LICENSE_ERROR_TYPE error);

void cgo_RtcConnectionObserverBridge_onUserNetworkQuality(C_RtcConnectionObserverBridge *this_, void *userData,
	C_user_id_t userId, enum C_QUALITY_TYPE txQuality,
	enum C_QUALITY_TYPE rxQuality);

void cgo_RtcConnectionObserverBridge_onNetworkTypeChanged(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_NETWORK_TYPE type);

void cgo_RtcConnectionObserverBridge_onContentInspectResult(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_CONTENT_INSPECT_RESULT result);

void cgo_RtcConnectionObserverBridge_onSnapshotTaken(C_RtcConnectionObserverBridge *this_, void *userData,
	C_uid_t uid, char *filePath, int width, int height, int errCode);

void cgo_RtcConnectionObserverBridge_onError(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_ERROR_CODE_TYPE error, char *msg);

void cgo_RtcConnectionObserverBridge_onChannelMediaRelayStateChanged(C_RtcConnectionObserverBridge *this_, void *userData,
	int state, int code);

void cgo_RtcConnectionObserverBridge_onLocalUserRegistered(C_RtcConnectionObserverBridge *this_, void *userData,
	C_uid_t uid, char *userAccount);

void cgo_RtcConnectionObserverBridge_onUserAccountUpdated(C_RtcConnectionObserverBridge *this_, void *userData,
	C_uid_t uid, char *userAccount);

void cgo_RtcConnectionObserverBridge_onStreamMessageError(C_RtcConnectionObserverBridge *this_, void *userData,
	C_user_id_t userId, int streamId, int code, int missed,
	int cached);

void cgo_RtcConnectionObserverBridge_onEncryptionError(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_ENCRYPTION_ERROR_TYPE errorType);

void cgo_RtcConnectionObserverBridge_onUploadLogResult(C_RtcConnectionObserverBridge *this_, void *userData,
	char *requestId, bool success, enum C_UPLOAD_ERROR_REASON reason);

void cgo_RtcConnectionObserverBridge_onWlAccMessage(C_RtcConnectionObserverBridge *this_, void *userData,
	enum C_WLACC_MESSAGE_REASON reason, enum C_WLACC_SUGGEST_ACTION action, char *wlAccMsg);

void cgo_RtcConnectionObserverBridge_onWlAccStats(C_RtcConnectionObserverBridge *this_, void *userData,
	struct C_WlAccStats currentStats, struct C_WlAccStats averageStats);

*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/agora_rtc_sdk_cgo/pkg/agora"
)

type IRtcConnectionObserverHandler interface {

	/**
	 * Occurs when the connection state between the SDK and the Agora channel changes to `CONNECTION_STATE_CONNECTED(3)`.
	 *
	 * @param connectionInfo The information of the connection. See \ref agora::rtc::TConnectionInfo "TConnectionInfo".
	 * @param reason The reason of the connection state change. See #CONNECTION_CHANGED_REASON_TYPE.
	 */
	OnConnected(connectionInfo *agora.TConnectionInfo, reason agora.CONNECTION_CHANGED_REASON_TYPE)

	/**
	 * Occurs when the connection state between the SDK and the Agora channel changes to `CONNECTION_STATE_DISCONNECTED(1)`.
	 *
	 * @param connectionInfo The information of the connection. See \ref agora::rtc::TConnectionInfo "TConnectionInfo".
	 * @param reason The reason of the connection state change. See #CONNECTION_CHANGED_REASON_TYPE.
	 */
	OnDisconnected(connectionInfo *agora.TConnectionInfo, reason agora.CONNECTION_CHANGED_REASON_TYPE)

	/**
	 * Occurs when the connection state between the SDK and the Agora channel changes to `CONNECTION_STATE_CONNECTING(2)`.
	 *
	 * @param connectionInfo The information of the connection. See \ref agora::rtc::TConnectionInfo "TConnectionInfo".
	 * @param reason The reason of the connection state change. See #CONNECTION_CHANGED_REASON_TYPE.
	 */
	OnConnecting(connectionInfo *agora.TConnectionInfo, reason agora.CONNECTION_CHANGED_REASON_TYPE)

	/**
	 * Occurs when the connection state between the SDK and the Agora channel changes to `CONNECTION_STATE_RECONNECTING(4)`.
	 *
	 * @param connectionInfo The information of the connection. See \ref agora::rtc::TConnectionInfo "TConnectionInfo".
	 * @param reason The reason of the connection state change. See #CONNECTION_CHANGED_REASON_TYPE.
	 */
	OnReconnecting(connectionInfo *agora.TConnectionInfo, reason agora.CONNECTION_CHANGED_REASON_TYPE)

	// This should be deleted. onConnected is enough.
	OnReconnected(connectionInfo *agora.TConnectionInfo, reason agora.CONNECTION_CHANGED_REASON_TYPE)

	/**
	 * Occurs when the SDK loses connection with the Agora channel.
	 *
	 * @param connectionInfo The information of the connection. See \ref agora::rtc::TConnectionInfo "TConnectionInfo".
	 */
	OnConnectionLost(connectionInfo *agora.TConnectionInfo)

	/**
	 * Reports the quality of the last-mile network.
	 *
	 * The SDK triggers this callback within two seconds after the app calls \ref IRtcConnection::startLastmileProbeTest "startLastmileProbeTest".
	 *
	 * @param quality Quality of the last-mile network: #QUALITY_TYPE.
	 */
	OnLastmileQuality(quality agora.QUALITY_TYPE)

	/**
	 * Reports the result of the last-mile network probe test.
	 *
	 * The SDK triggers this callback within 30 seconds after the app calls \ref IRtcConnection::startLastmileProbeTest "startLastmileProbeTest".
	 *
	 * @param result The result of the last-mile network probe test: \ref agora::rtc::LastmileProbeResult "LastmileProbeResult".
	 */
	OnLastmileProbeResult(result *agora.LastmileProbeResult)

	/**
	 * Occurs when the token expires in 30 seconds.
	 *
	 * The SDK triggers this callback to remind the app to get a new token before the token privilege expires.
	 *
	 * Upon receiving this callback, you must generate a new token on your server and call \ref IRtcConnection::renewToken
	 * "renewToken" to pass the new token to the SDK.
	 *
	 * @param token The pointer to the token that expires in 30 seconds.
	 */
	OnTokenPrivilegeWillExpire(token string)

	/**
	 * Occurs when the token has expired.
	 *
	 * Upon receiving this callback, you must generate a new token on your server and call \ref IRtcConnection::renewToken
	 * "renewToken" to pass the new token to the SDK.
	 */
	OnTokenPrivilegeDidExpire()

	/**
	 * Occurs when the connection state between the SDK and the Agora channel changes to `CONNECTION_STATE_FAILED(5)`.
	 *
	 * @param connectionInfo The connection information: TConnectionInfo.
	 * @param reason The reason of the connection state change: #CONNECTION_CHANGED_REASON_TYPE.
	 */
	OnConnectionFailure(connectionInfo *agora.TConnectionInfo,
		reason agora.CONNECTION_CHANGED_REASON_TYPE)

	/**
	 * Occurs when a remote user joins the channel.
	 *
	 * You can get the ID of the remote user in this callback.
	 *
	 * @param userId The ID of the remote user who joins the channel.
	 */
	OnUserJoined(userId string)

	/**
	 * Occurs when a remote user leaves the channel.
	 *
	 * You can know why the user leaves the channel through the `reason` parameter.
	 *
	 * @param userId The ID of the user who leaves the channel.
	 * @param reason The reason why the remote user leaves the channel: #USER_OFFLINE_REASON_TYPE.
	 */
	OnUserLeft(userId string, reason agora.USER_OFFLINE_REASON_TYPE)

	/**
	 * Reports the transport statistics of the connection.
	 *
	 * The SDK triggers this callback once every two seconds when the connection state is `CONNECTION_STATE_CONNECTED`.
	 *
	 * @param stats The pointer to \ref rtc::RtcStats "RtcStats".
	 */
	OnTransportStats(stats *agora.RtcStats)

	/**
	 * Occurs when the role of the local user changes.
	 *
	 * @param oldRole The previous role of the local user: \ref rtc::CLIENT_ROLE_TYPE "CLIENT_ROLE_TYPE".
	 * @param newRole The current role of the local user: \ref rtc::CLIENT_ROLE_TYPE "CLIENT_ROLE_TYPE".
	 * @param newRoleOptions The client role options of the current role of the local user: \ref rtc::ClientRoleOptions "ClientRoleOptions".
	 */
	OnChangeRoleSuccess(oldRole agora.CLIENT_ROLE_TYPE, newRolem agora.CLIENT_ROLE_TYPE, newRoleOptions *agora.ClientRoleOptions)

	/**
	 * Occurs when the local user fails to change the user role.
	 */
	OnChangeRoleFailure(reason agora.CLIENT_ROLE_CHANGE_FAILED_REASON, currentRole agora.CLIENT_ROLE_TYPE)

	/**
	 * Occurs when connection license verification fails
	 *
	 * You can know the reason accordding to error code
	 * @param error verify fail reason
	 */
	OnLicenseValidationFailure(_error agora.LICENSE_ERROR_TYPE)

	/**
	 * Reports the network quality of each user.
	 *
	 * The SDK triggers this callback once every two seconds to report the uplink and downlink network conditions
	 * of each user in the channel, including the local user.
	 *
	 * @param userId The ID of the user. If `userId` is empty, this callback reports the network quality of the local user.
	 * @param txQuality The uplink network quality: #QUALITY_TYPE.
	 * @param rxQuality The downlink network quality: #QUALITY_TYPE.
	 */
	OnUserNetworkQuality(userId string, txQuality agora.QUALITY_TYPE,
		rxQuality agora.QUALITY_TYPE)

	/** Occurs when the network type is changed.
	 * @param type The current network type. See #NETWORK_TYPE.
	 */
	OnNetworkTypeChanged(_type agora.NETWORK_TYPE)

	/** Reports result of Content Inspect*/
	OnContentInspectResult(result agora.CONTENT_INSPECT_RESULT)
	/** Occurs when takeSnapshot API result is obtained
	 *
	 *
	 * @brief snapshot taken callback
	 *
	 * @param channel channel name
	 * @param uid user id
	 * @param filePath image is saveed file path
	 * @param width image width
	 * @param height image height
	 * @param errCode 0 is ok negative is error
	 */
	OnSnapshotTaken(uid agora.Uid_t, filePath string, width int, height int, errCode int)

	/**
	 * Reports the error code and error message.
	 * @param error The error code: #ERROR_CODE_TYPE.
	 * @param msg The error message.
	 */
	OnError(_error agora.ERROR_CODE_TYPE, msg string)

	/**
	 * Occurs when the state of the channel media relay changes.
	 *
	 *
	 * @param state The state code:
	 * - `RELAY_STATE_IDLE(0)`: The SDK is initializing.
	 * - `RELAY_STATE_CONNECTING(1)`: The SDK tries to relay the media stream to the destination
	 * channel.
	 * - `RELAY_STATE_RUNNING(2)`: The SDK successfully relays the media stream to the destination
	 * channel.
	 * - `RELAY_STATE_FAILURE(3)`: A failure occurs. See the details in `code`.
	 * @param code The error code:
	 * - `RELAY_OK(0)`: The state is normal.
	 * - `RELAY_ERROR_SERVER_ERROR_RESPONSE(1)`: An error occurs in the server response.
	 * - `RELAY_ERROR_SERVER_NO_RESPONSE(2)`: No server response. You can call the leaveChannel method
	 * to leave the channel.
	 * - `RELAY_ERROR_NO_RESOURCE_AVAILABLE(3)`: The SDK fails to access the service, probably due to
	 * limited resources of the server.
	 * - `RELAY_ERROR_FAILED_JOIN_SRC(4)`: Fails to send the relay request.
	 * - `RELAY_ERROR_FAILED_JOIN_DEST(5)`: Fails to accept the relay request.
	 * - `RELAY_ERROR_FAILED_PACKET_RECEIVED_FROM_SRC(6)`: The server fails to receive the media
	 * stream.
	 * - `RELAY_ERROR_FAILED_PACKET_SENT_TO_DEST(7)`: The server fails to send the media stream.
	 * - `RELAY_ERROR_SERVER_CONNECTION_LOST(8)`: The SDK disconnects from the server due to poor
	 * network connections. You can call the leaveChannel method to leave the channel.
	 * - `RELAY_ERROR_INTERNAL_ERROR(9)`: An internal error occurs in the server.
	 * - `RELAY_ERROR_SRagora.TOKEN_EXPIRED(10)`: The token of the source channel has expired.
	 * - `RELAY_ERROR_DEST_TOKEN_EXPIRED(11)`: The token of the destination channel has expired.
	 */
	OnChannelMediaRelayStateChanged(state int, code int)

	/** Occurs when the local user successfully registers a user account by calling the \ref IRtcEngine::joinChannelWithUserAccount "joinChannelWithUserAccount" method.This callback reports the user ID and user account of the local user.
	 *
	 * @param uid The ID of the local user.
	 * @param userAccount The user account of the local user.
	 */
	OnLocalUserRegistered(uid agora.Uid_t, userAccount string)

	/** Technical Preview, please do not depend on this event. */
	OnUserAccountUpdated(uid agora.Uid_t, userAccount string)

	/**
	 * Reports the error that occurs when receiving data stream messages.
	 *
	 * @param userId The ID of the user sending the data stream.
	 * @param streamId  the ID of the sent data stream, returned in the \ref agora::rtc::IRtcConnection::createDataStream "createDataStream" method.
	 * @param code The error code.
	 * @param missed The number of lost messages.
	 * @param cached The number of incoming cached messages when the data stream is interrupted.
	 */
	OnStreamMessageError(userId string, streamId int, code int, missed int,
		cached int)

	/**
	 * Reports the error type of encryption.
	 * @param type See #ENCRYPTION_ERROR_TYPE.
	 */
	OnEncryptionError(errorType agora.ENCRYPTION_ERROR_TYPE)
	/**
	 * Reports the user log upload result
	 * @param requestId RequestId of the upload
	 * @param success Is upload success
	 * @param reason Reason of the upload, 0: OK, 1 Network Error, 2 Server Error.
	 */
	OnUploadLogResult(requestId string, success bool, reason agora.UPLOAD_ERROR_REASON)

	/** Occurs when the WIFI message need be sent to the user.
	 *
	 * @param reason The reason of notifying the user of a message.
	 * @param action Suggest an action for the user.
	 * @param wlAccMsg The message content of notifying the user.
	 */
	OnWlAccMessage(reason agora.WLACC_MESSAGE_REASON, action agora.WLACC_SUGGEST_ACTION, wlAccMsg string)

	/** Occurs when SDK statistics wifi acceleration optimization effect.
	 *
	 * @param currentStats Instantaneous value of optimization effect.
	 * @param averageStats Average value of cumulative optimization effect.
	 */
	OnWlAccStats(currentStats agora.WlAccStats, averageStats agora.WlAccStats)
}

type RtcConnectionObserverBridge struct {
	handler IRtcConnectionObserverHandler
	cBridge *C.C_RtcConnectionObserverBridge
}

func (b *RtcConnectionObserverBridge) ToAgoraEventHandler() *agora.IRtcConnectionObserver {
	return (*agora.IRtcConnectionObserver)(b.cBridge)
}

func (b *RtcConnectionObserverBridge) Delete() {
	C.C_RtcConnectionObserverBridge_Delete(unsafe.Pointer(b.cBridge))
	b.handler = nil
	b.cBridge = nil
}

func NewRtcConnectionObserverBridge(handler IRtcConnectionObserverHandler) *RtcConnectionObserverBridge {
	b := RtcConnectionObserverBridge{}
	userData := unsafe.Pointer(&b)
	b.cBridge = (*C.C_RtcConnectionObserverBridge)(C.C_RtcConnectionObserverBridge_New(
		C.C_RtcConnectionObserverBridge_Callbacks{
			onConnected:                     C.C_RtcConnectionObserverBridge_onConnected(C.cgo_RtcConnectionObserverBridge_onConnected),
			onDisconnected:                  C.C_RtcConnectionObserverBridge_onDisconnected(C.cgo_RtcConnectionObserverBridge_onDisconnected),
			onConnecting:                    C.C_RtcConnectionObserverBridge_onConnecting(C.cgo_RtcConnectionObserverBridge_onConnecting),
			onReconnecting:                  C.C_RtcConnectionObserverBridge_onReconnecting(C.cgo_RtcConnectionObserverBridge_onReconnecting),
			onReconnected:                   C.C_RtcConnectionObserverBridge_onReconnected(C.cgo_RtcConnectionObserverBridge_onReconnected),
			onConnectionLost:                C.C_RtcConnectionObserverBridge_onConnectionLost(C.cgo_RtcConnectionObserverBridge_onConnectionLost),
			onLastmileQuality:               C.C_RtcConnectionObserverBridge_onLastmileQuality(C.cgo_RtcConnectionObserverBridge_onLastmileQuality),
			onLastmileProbeResult:           C.C_RtcConnectionObserverBridge_onLastmileProbeResult(C.cgo_RtcConnectionObserverBridge_onLastmileProbeResult),
			onTokenPrivilegeWillExpire:      C.C_RtcConnectionObserverBridge_onTokenPrivilegeWillExpire(C.cgo_RtcConnectionObserverBridge_onTokenPrivilegeWillExpire),
			onTokenPrivilegeDidExpire:       C.C_RtcConnectionObserverBridge_onTokenPrivilegeDidExpire(C.cgo_RtcConnectionObserverBridge_onTokenPrivilegeDidExpire),
			onConnectionFailure:             C.C_RtcConnectionObserverBridge_onConnectionFailure(C.cgo_RtcConnectionObserverBridge_onConnectionFailure),
			onUserJoined:                    C.C_RtcConnectionObserverBridge_onUserJoined(C.cgo_RtcConnectionObserverBridge_onUserJoined),
			onUserLeft:                      C.C_RtcConnectionObserverBridge_onUserLeft(C.cgo_RtcConnectionObserverBridge_onUserLeft),
			onTransportStats:                C.C_RtcConnectionObserverBridge_onTransportStats(C.cgo_RtcConnectionObserverBridge_onTransportStats),
			onChangeRoleSuccess:             C.C_RtcConnectionObserverBridge_onChangeRoleSuccess(C.cgo_RtcConnectionObserverBridge_onChangeRoleSuccess),
			onChangeRoleFailure:             C.C_RtcConnectionObserverBridge_onChangeRoleFailure(C.cgo_RtcConnectionObserverBridge_onChangeRoleFailure),
			onLicenseValidationFailure:      C.C_RtcConnectionObserverBridge_onLicenseValidationFailure(C.cgo_RtcConnectionObserverBridge_onLicenseValidationFailure),
			onUserNetworkQuality:            C.C_RtcConnectionObserverBridge_onUserNetworkQuality(C.cgo_RtcConnectionObserverBridge_onUserNetworkQuality),
			onNetworkTypeChanged:            C.C_RtcConnectionObserverBridge_onNetworkTypeChanged(C.cgo_RtcConnectionObserverBridge_onNetworkTypeChanged),
			onContentInspectResult:          C.C_RtcConnectionObserverBridge_onContentInspectResult(C.cgo_RtcConnectionObserverBridge_onContentInspectResult),
			onSnapshotTaken:                 C.C_RtcConnectionObserverBridge_onSnapshotTaken(C.cgo_RtcConnectionObserverBridge_onSnapshotTaken),
			onError:                         C.C_RtcConnectionObserverBridge_onError(C.cgo_RtcConnectionObserverBridge_onError),
			onChannelMediaRelayStateChanged: C.C_RtcConnectionObserverBridge_onChannelMediaRelayStateChanged(C.cgo_RtcConnectionObserverBridge_onChannelMediaRelayStateChanged),
			onLocalUserRegistered:           C.C_RtcConnectionObserverBridge_onLocalUserRegistered(C.cgo_RtcConnectionObserverBridge_onLocalUserRegistered),
			onUserAccountUpdated:            C.C_RtcConnectionObserverBridge_onUserAccountUpdated(C.cgo_RtcConnectionObserverBridge_onUserAccountUpdated),
			onStreamMessageError:            C.C_RtcConnectionObserverBridge_onStreamMessageError(C.cgo_RtcConnectionObserverBridge_onStreamMessageError),
			onEncryptionError:               C.C_RtcConnectionObserverBridge_onEncryptionError(C.cgo_RtcConnectionObserverBridge_onEncryptionError),
			onUploadLogResult:               C.C_RtcConnectionObserverBridge_onUploadLogResult(C.cgo_RtcConnectionObserverBridge_onUploadLogResult),
			onWlAccMessage:                  C.C_RtcConnectionObserverBridge_onWlAccMessage(C.cgo_RtcConnectionObserverBridge_onWlAccMessage),
			onWlAccStats:                    C.C_RtcConnectionObserverBridge_onWlAccStats(C.cgo_RtcConnectionObserverBridge_onWlAccStats),
		},
		userData,
	))
	b.handler = handler
	return &b
}

//export cgo_RtcConnectionObserverBridge_onConnected
func cgo_RtcConnectionObserverBridge_onConnected(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	connectionInfo *C.struct_C_TConnectionInfo, reason C.enum_C_CONNECTION_CHANGED_REASON_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnConnected(
		(*agora.TConnectionInfo)(unsafe.Pointer(connectionInfo)),
		agora.CONNECTION_CHANGED_REASON_TYPE(reason),
	)
}

//export cgo_RtcConnectionObserverBridge_onDisconnected
func cgo_RtcConnectionObserverBridge_onDisconnected(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	connectionInfo *C.struct_C_TConnectionInfo, reason C.enum_C_CONNECTION_CHANGED_REASON_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnDisconnected(
		(*agora.TConnectionInfo)(unsafe.Pointer(connectionInfo)),
		agora.CONNECTION_CHANGED_REASON_TYPE(reason),
	)
}

//export cgo_RtcConnectionObserverBridge_onConnecting
func cgo_RtcConnectionObserverBridge_onConnecting(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	connectionInfo *C.struct_C_TConnectionInfo, reason C.enum_C_CONNECTION_CHANGED_REASON_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnConnecting(
		(*agora.TConnectionInfo)(unsafe.Pointer(connectionInfo)),
		agora.CONNECTION_CHANGED_REASON_TYPE(reason),
	)
}

//export cgo_RtcConnectionObserverBridge_onReconnecting
func cgo_RtcConnectionObserverBridge_onReconnecting(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	connectionInfo *C.struct_C_TConnectionInfo, reason C.enum_C_CONNECTION_CHANGED_REASON_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnReconnecting(
		(*agora.TConnectionInfo)(unsafe.Pointer(connectionInfo)),
		agora.CONNECTION_CHANGED_REASON_TYPE(reason),
	)
}

//export cgo_RtcConnectionObserverBridge_onReconnected
func cgo_RtcConnectionObserverBridge_onReconnected(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	connectionInfo *C.struct_C_TConnectionInfo, reason C.enum_C_CONNECTION_CHANGED_REASON_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnReconnected(
		(*agora.TConnectionInfo)(unsafe.Pointer(connectionInfo)),
		agora.CONNECTION_CHANGED_REASON_TYPE(reason),
	)
}

//export cgo_RtcConnectionObserverBridge_onConnectionLost
func cgo_RtcConnectionObserverBridge_onConnectionLost(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	connectionInfo *C.struct_C_TConnectionInfo) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnConnectionLost(
		(*agora.TConnectionInfo)(unsafe.Pointer(connectionInfo)),
	)
}

//export cgo_RtcConnectionObserverBridge_onLastmileQuality
func cgo_RtcConnectionObserverBridge_onLastmileQuality(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	quality C.enum_C_QUALITY_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnLastmileQuality(
		agora.QUALITY_TYPE(quality),
	)
}

//export cgo_RtcConnectionObserverBridge_onLastmileProbeResult
func cgo_RtcConnectionObserverBridge_onLastmileProbeResult(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	result *C.struct_C_LastmileProbeResult) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnLastmileProbeResult(
		(*agora.LastmileProbeResult)(unsafe.Pointer(result)),
	)
}

//export cgo_RtcConnectionObserverBridge_onTokenPrivilegeWillExpire
func cgo_RtcConnectionObserverBridge_onTokenPrivilegeWillExpire(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	token *C.char) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnTokenPrivilegeWillExpire(
		C.GoString(token),
	)
}

//export cgo_RtcConnectionObserverBridge_onTokenPrivilegeDidExpire
func cgo_RtcConnectionObserverBridge_onTokenPrivilegeDidExpire(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnTokenPrivilegeDidExpire()
}

//export cgo_RtcConnectionObserverBridge_onConnectionFailure
func cgo_RtcConnectionObserverBridge_onConnectionFailure(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	connectionInfo *C.struct_C_TConnectionInfo,
	reason C.enum_C_CONNECTION_CHANGED_REASON_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnConnectionFailure(
		(*agora.TConnectionInfo)(unsafe.Pointer(connectionInfo)),
		agora.CONNECTION_CHANGED_REASON_TYPE(reason),
	)
}

//export cgo_RtcConnectionObserverBridge_onUserJoined
func cgo_RtcConnectionObserverBridge_onUserJoined(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnUserJoined(
		C.GoString(userId),
	)
}

//export cgo_RtcConnectionObserverBridge_onUserLeft
func cgo_RtcConnectionObserverBridge_onUserLeft(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, reason C.enum_C_USER_OFFLINE_REASON_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnUserLeft(
		C.GoString(userId),
		agora.USER_OFFLINE_REASON_TYPE(reason),
	)
}

//export cgo_RtcConnectionObserverBridge_onTransportStats
func cgo_RtcConnectionObserverBridge_onTransportStats(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	stats *C.struct_C_RtcStats) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnTransportStats(
		(*agora.RtcStats)(unsafe.Pointer(stats)),
	)
}

//export cgo_RtcConnectionObserverBridge_onChangeRoleSuccess
func cgo_RtcConnectionObserverBridge_onChangeRoleSuccess(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	oldRole C.enum_C_CLIENT_ROLE_TYPE, newRole C.enum_C_CLIENT_ROLE_TYPE, newRoleOptions *C.struct_C_ClientRoleOptions) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnChangeRoleSuccess(
		agora.CLIENT_ROLE_TYPE(oldRole),
		agora.CLIENT_ROLE_TYPE(newRole),
		(*agora.ClientRoleOptions)(unsafe.Pointer(newRoleOptions)),
	)
}

//export cgo_RtcConnectionObserverBridge_onChangeRoleFailure
func cgo_RtcConnectionObserverBridge_onChangeRoleFailure(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	reason C.enum_C_CLIENT_ROLE_CHANGE_FAILED_REASON, currentRole C.enum_C_CLIENT_ROLE_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnChangeRoleFailure(
		agora.CLIENT_ROLE_CHANGE_FAILED_REASON(reason),
		agora.CLIENT_ROLE_TYPE(currentRole),
	)
}

//export cgo_RtcConnectionObserverBridge_onLicenseValidationFailure
func cgo_RtcConnectionObserverBridge_onLicenseValidationFailure(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	_error C.enum_C_LICENSE_ERROR_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnLicenseValidationFailure(
		agora.LICENSE_ERROR_TYPE(_error),
	)
}

//export cgo_RtcConnectionObserverBridge_onUserNetworkQuality
func cgo_RtcConnectionObserverBridge_onUserNetworkQuality(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, txQuality C.enum_C_QUALITY_TYPE,
	rxQuality C.enum_C_QUALITY_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnUserNetworkQuality(
		C.GoString(userId),
		agora.QUALITY_TYPE(txQuality),
		agora.QUALITY_TYPE(rxQuality),
	)
}

//export cgo_RtcConnectionObserverBridge_onNetworkTypeChanged
func cgo_RtcConnectionObserverBridge_onNetworkTypeChanged(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	_type C.enum_C_NETWORK_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnNetworkTypeChanged(
		agora.NETWORK_TYPE(_type),
	)
}

//export cgo_RtcConnectionObserverBridge_onContentInspectResult
func cgo_RtcConnectionObserverBridge_onContentInspectResult(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	result C.enum_C_CONTENT_INSPECT_RESULT) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnContentInspectResult(
		agora.CONTENT_INSPECT_RESULT(result),
	)
}

//export cgo_RtcConnectionObserverBridge_onSnapshotTaken
func cgo_RtcConnectionObserverBridge_onSnapshotTaken(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	uid C.C_uid_t, filePath *C.char, width C.int, height C.int, errCode C.int) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnSnapshotTaken(
		agora.Uid_t(uid),
		C.GoString(filePath),
		int(width),
		int(height),
		int(errCode),
	)
}

//export cgo_RtcConnectionObserverBridge_onError
func cgo_RtcConnectionObserverBridge_onError(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	_error C.enum_C_ERROR_CODE_TYPE, msg *C.char) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnError(
		agora.ERROR_CODE_TYPE(_error),
		C.GoString(msg),
	)
}

//export cgo_RtcConnectionObserverBridge_onChannelMediaRelayStateChanged
func cgo_RtcConnectionObserverBridge_onChannelMediaRelayStateChanged(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	state C.int, code C.int) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnChannelMediaRelayStateChanged(
		int(state),
		int(code),
	)
}

//export cgo_RtcConnectionObserverBridge_onLocalUserRegistered
func cgo_RtcConnectionObserverBridge_onLocalUserRegistered(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	uid C.C_uid_t, userAccount *C.char) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnLocalUserRegistered(
		agora.Uid_t(uid),
		C.GoString(userAccount),
	)
}

//export cgo_RtcConnectionObserverBridge_onUserAccountUpdated
func cgo_RtcConnectionObserverBridge_onUserAccountUpdated(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	uid C.C_uid_t, userAccount *C.char) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnUserAccountUpdated(
		agora.Uid_t(uid),
		C.GoString(userAccount),
	)
}

//export cgo_RtcConnectionObserverBridge_onStreamMessageError
func cgo_RtcConnectionObserverBridge_onStreamMessageError(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	userId C.C_user_id_t, streamId C.int, code C.int, missed C.int,
	cached C.int) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnStreamMessageError(
		C.GoString(userId),
		int(streamId),
		int(code),
		int(missed),
		int(cached),
	)
}

//export cgo_RtcConnectionObserverBridge_onEncryptionError
func cgo_RtcConnectionObserverBridge_onEncryptionError(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	errorType C.enum_C_ENCRYPTION_ERROR_TYPE) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnEncryptionError(
		agora.ENCRYPTION_ERROR_TYPE(errorType),
	)
}

//export cgo_RtcConnectionObserverBridge_onUploadLogResult
func cgo_RtcConnectionObserverBridge_onUploadLogResult(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	requestId *C.char, success C.bool, reason C.enum_C_UPLOAD_ERROR_REASON) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnUploadLogResult(
		C.GoString(requestId),
		bool(success),
		agora.UPLOAD_ERROR_REASON(reason),
	)
}

//export cgo_RtcConnectionObserverBridge_onWlAccMessage
func cgo_RtcConnectionObserverBridge_onWlAccMessage(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	reason C.enum_C_WLACC_MESSAGE_REASON, action C.enum_C_WLACC_SUGGEST_ACTION, wlAccMsg *C.char) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnWlAccMessage(
		agora.WLACC_MESSAGE_REASON(reason),
		agora.WLACC_SUGGEST_ACTION(action),
		C.GoString(wlAccMsg),
	)
}

//export cgo_RtcConnectionObserverBridge_onWlAccStats
func cgo_RtcConnectionObserverBridge_onWlAccStats(_ *C.C_RtcConnectionObserverBridge, userData unsafe.Pointer,
	currentStats C.struct_C_WlAccStats, averageStats C.struct_C_WlAccStats) {
	if userData == nil {
		return
	}

	bridge := (*RtcConnectionObserverBridge)(userData)
	bridge.handler.OnWlAccStats(
		*(*agora.WlAccStats)(unsafe.Pointer(&currentStats)),
		*(*agora.WlAccStats)(unsafe.Pointer(&averageStats)),
	)
}
