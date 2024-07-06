package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "C_NGIAgoraLocalUser.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

// #ifndef C_NGI_AGORA_LOCAL_USER_H
// #define C_NGI_AGORA_LOCAL_USER_H

// #include <string.h>
// // #include <vector>
// #include "C_AgoraBase.h"
// #include "C_AgoraOptional.h"

// #ifdef __cplusplus
// extern "C"
// {
// #endif // __cplusplus

// #region agora

// #region agora::media

//   // class IAudioFrameObserver;

// #endregion agora::media

//   // class ILocalDataChannel;
//   // class IDataChannelObserver;

// #region agora::rtc

//   // class IAudioEngineWrapper;

//   // class ILocalUserObserver;
//   // class ILocalVideoTrack;
//   // class IRemoteVideoTrack;
//   // class IVideoFrameObserver2;

//   // class IMediaControlPacketSender;
//   // class IMediaControlPacketReceiver;
//   struct C_AudioEncoderConfiguration;
//   // struct VideoEncoderConfiguration;

//   // class ILocalAudioTrack;
//   // class IAudioMixerSource;
//   // struct RemoteAudioTrackStats;
//   // class IRemoteAudioTrack;
//   // struct LocalVideoTrackStats;
//   // struct RemoteVideoTrackStats;
//   // class IMediaPacketReceiver;
//   // class IVideoSinkBase;

/**
 * The definition of the AudioVolumeInformation struct.
 */
type AudioVolumeInformation C.struct_C_AudioVolumeInformation

// {
//   //   /**
//   //    * User ID of the speaker.
//   //    */
//   //   user_id_t userId;

//   //   /**
//   //    * The volume of the speaker that ranges from 0 to 255.
//   //    */
//   //   unsigned int volume;

//   //   /*
//   //    * The activity status of remote users
//   //    */
//   //   unsigned int vad;

//   //   /**
//   //    * Voice pitch frequency in Hz
//   //    */
//   //   double voicePitch;

//   //   AudioVolumeInformation() : userId(NULL), volume(0), vad(0), voicePitch(0.0) {}
//   // }

/**
 * The `ILocalUserObserver` class.
 */
type ILocalUserObserver C.C_ILocalUserObserver

// #region C_ILocalUserObserver
// // class ILocalUserObserver {
// //  public:
// //   virtual ~ILocalUserObserver() {}

// //   virtual void onAudioTrackPublishStart(agora_refptr<ILocalAudioTrack> audioTrack) = 0;
// //   /**
// //    * Occurs when the first packet of the local audio track is sent, indicating that the local audio track
// //    * is successfully published.
// //    *
// //    * @param audioTrack The pointer to \ref agora::rtc::ILocalAudioTrack "ILocalAudioTrack".
// //    */
// //   virtual void onAudioTrackPublishSuccess(agora_refptr<ILocalAudioTrack> audioTrack) = 0;

// //   virtual void onAudioTrackUnpublished(agora_refptr<ILocalAudioTrack> audioTrack) = 0;

// //   /**
// //    * Occurs when a local audio track fails to be published.
// //    *
// //    * @param audioTrack The pointer to `ILocalAudioTrack`.
// //    * @param error The error information: #ERROR_CODE_TYPE.
// //    */
// //   virtual void onAudioTrackPublicationFailure(agora_refptr<ILocalAudioTrack> audioTrack,
// //                                               ERROR_CODE_TYPE error) = 0;

// //   /**
// //    * Reports the statistics of a local audio track.
// //    *
// //    * @param stats The reference to the statistics of the local audio track.
// //    */
// //   virtual void onLocalAudioTrackStatistics(const LocalAudioStats& stats) = 0;
// //   /**
// //    * Reports the statistics of a remote audio track.
// //    *
// //    * @param audioTrack The pointer to `IRemoteAudioTrack`.
// //    * @param stats The statistics of the remote audio track.
// //    */
// //   virtual void onRemoteAudioTrackStatistics(agora_refptr<rtc::IRemoteAudioTrack> audioTrack, const RemoteAudioTrackStats& stats) = 0;
// //   /**
// //    * Occurs when the first remote audio frame is received.
// //    *
// //    * This callback indicates that the local user has subscribed to a specified remote audio track, and the first
// //    * frame of this audio track has been received.
// //    *
// //    * @param userId The ID of the remote user whose audio frame is received.
// //    * @param audioTrack The pointer to `IRemoteAudioTrack`.
// //    */
// //   virtual void onUserAudioTrackSubscribed(user_id_t userId,
// //                                           agora_refptr<rtc::IRemoteAudioTrack> audioTrack) = 0;

// //   /**
// //    * Occurs when the state of a remote audio track changes.
// //    *
// //    * @param userId The ID of the remote user whose audio track state has changed.
// //    * @param audioTrack The pointer to `IRemoteAudioTrack`.
// //    * @param state The current state of the audio track.
// //    * @param reason The reason for the state change.
// //    * @param elapsed The time (ms) since the user connects to an Agora channel.
// //    */
// //   virtual void onUserAudioTrackStateChanged(user_id_t userId,
// //                                             agora_refptr<rtc::IRemoteAudioTrack> audioTrack,
// //                                             REMOTE_AUDIO_STATE state,
// //                                             REMOTE_AUDIO_STATE_REASON reason,
// //                                             int elapsed) = 0;

// //   virtual void onVideoTrackPublishStart(agora_refptr<ILocalVideoTrack> videoTrack) = 0;
// //   /**
// //    * Occurs when the first packet of a local video track is sent, indicating that the local video track
// //    * is successfully published.
// //    * @param videoTrack The pointer to `ILocalVideoTrack`.
// //    */
// //   virtual void onVideoTrackPublishSuccess(agora_refptr<ILocalVideoTrack> videoTrack) = 0;

// //   /**
// //    * Occurs when a local video track fails to be published.
// //    *
// //    * @param videoTrack The pointer to `ILocalVideoTrack`.
// //    * @param error The error information: #ERROR_CODE_TYPE.
// //    */
// //   virtual void onVideoTrackPublicationFailure(agora_refptr<ILocalVideoTrack> videoTrack,
// //                                               ERROR_CODE_TYPE error) = 0;

// //   virtual void onVideoTrackUnpublished(agora_refptr<ILocalVideoTrack> videoTrack) = 0;
// //   /**
// //    * Occurs when the state of a local video track changes.
// //    * @note
// //    * When you receive error from this callback, the corresponding track is in error state.
// //    * To make the track recover from error state, we highly recommend that you disable the track and
// //    * try re-enabling it again.
// //    *
// //    * @param videoTrack The pointer to `ILocalVideoTrack`.
// //    * @param state The state of the local video track.
// //    * @param errorCode The error information.
// //    */
// //   virtual void onLocalVideoTrackStateChanged(agora_refptr<rtc::ILocalVideoTrack> videoTrack,
// //                                              LOCAL_VIDEO_STREAM_STATE state,
// //                                              LOCAL_VIDEO_STREAM_ERROR errorCode) = 0;

// //   /**
// //    * Reports the statistics of a local video track.
// //    *
// //    * @param videoTrack The pointer to `ILocalVideoTrack`.
// //    * @param stats The statistics of the local video track.
// //    */
// //   virtual void onLocalVideoTrackStatistics(agora_refptr<rtc::ILocalVideoTrack> videoTrack,
// //                                            const LocalVideoTrackStats& stats) = 0;

// //   /**
// //    * Occurs when the first remote video frame is received.
// //    *
// //    * This callback indicates that the local user has subscribed to a specified remote video track, and
// //    * the first frame of this video track has been received.
// //    *
// //    * @param userId The ID of the remote user whose video frame is received.
// //    * @param trackInfo The information of the remote video track.
// //    * @param videoTrack The pointer to `IRemoteVideoTrack`.
// //    */
// //   virtual void onUserVideoTrackSubscribed(user_id_t userId, VideoTrackInfo trackInfo,
// //                                           agora_refptr<rtc::IRemoteVideoTrack> videoTrack) = 0;

// //   /**
// //    * Occurs when the state of a remote video track changes.
// //    *
// //    * @param userId the ID of the remote user whose video track state has changed.
// //    * @param videoTrack The pointer to `IRemoteVideoTrack`.
// //    * @param state The current state of the video track.
// //    * @param reason The reason for the state change.
// //    * @param elapsed The time (ms) since the user connects to an Agora channel.
// //    */
// //   virtual void onUserVideoTrackStateChanged(user_id_t userId,
// //                                             agora_refptr<rtc::IRemoteVideoTrack> videoTrack,
// //                                             REMOTE_VIDEO_STATE state,
// //                                             REMOTE_VIDEO_STATE_REASON reason,
// //                                             int elapsed) = 0;

// //   /**
// //    * Occurs when the first remote video frame is rendered.
// //    *
// //    * @param userId the ID of the remote user.
// //    * @param width Width (px) of the video frame.
// //    * @param height Height (px) of the video stream.
// //    * @param elapsed The time (ms) since the user connects to an Agora channel.
// //   */
// //   virtual void onFirstRemoteVideoFrameRendered(user_id_t userId, int width,
// //                                                int height, int elapsed) = 0;

// //   /**
// //    * Reports the statistics of a remote video track.
// //    *
// //    * @param videoTrack The pointer to `IRemoteVideoTrack`.
// //    * @param stats The statistics of the remote video track.
// //    */
// //   virtual void onRemoteVideoTrackStatistics(agora_refptr<rtc::IRemoteVideoTrack> videoTrack,
// //                                             const RemoteVideoTrackStats& stats) = 0;

// //   /**
// //    * Reports which users are speaking, the speakers' volumes, and whether the local user is speaking.
// //    *
// //    * This callback reports the IDs and volumes of the loudest speakers at the moment in the channel,
// //    * and whether the local user is speaking.
// //    *
// //    * You can set the time interval of this callback using \ref ILocalUser::setAudioVolumeIndicationParameters "setAudioVolumeIndicationParameters".
// //    *
// //    * The SDK triggers two `onAudioVolumeIndication` callbacks at one time, one reporting the
// //    * volume information of the local user and the other reporting the volume information of all remote users.
// //    *
// //    * @note
// //    * - To detect whether the local user is speaking, set `report_vad` in `enableAudioVolumeIndication` to `true`.
// //    *
// //    * @param speakers The pointer to \ref agora::rtc::AudioVolumeInformation "AudioVolumeInformation", which is an array containing the user ID and volume information for each speaker.
// //    * - In the local user's callback, this array contains the following members:
// //    *   - `uid`, which is always `0`
// //    *   - `volume`, which reports the sum of the voice volume and the audio-mixing volume of the local user
// //    *   - `vad`, which reports whether the local user is speaking
// //    * - In the remote users' callback, this array contains the following members:
// //    *   - `uid`, which is the UID of each remote speaker
// //    *   - `volume`, which reports the sum of the voice volume and the audio-mixing volume of each remote speaker.
// //    *   - `vad`, which is always `0`
// //    * An empty `speakers` array in the callback indicates that no remote user is speaking at the moment.
// //    * @param speakerNumber The total number of the speakers.
// //    * @param totalVolume The total volume after audio mixing. The value ranges between 0 (lowest volume) and 255 (highest volume).
// //    * - In the local user's callback, `totalVolume` is the sum of the voice volume and the audio-mixing volume of the local user.
// //    * - In the remote speakers' callback, `totalVolume` is the sum of the voice volume and the audio-mixing volume of all remote speakers.
// //    */
// //   virtual void onAudioVolumeIndication(const AudioVolumeInformation* speakers, unsigned int speakerNumber,
// //                                        int totalVolume) = 0;

// //   /**
// //    * Occurs when an active speaker is detected.
// //    *
// //    * You can add relative functions on your app, for example, the active speaker, once detected,
// //    * will have the portrait zoomed in.
// //    *
// //    * @note
// //    * - The active speaker means the user ID of the speaker who speaks at the highest volume during a
// //    * certain period of time.
// //    *
// //    * @param userId The ID of the active speaker. A `userId` of `0` means the local user.
// //    */
// //   virtual void onActiveSpeaker(user_id_t userId) = 0;

// //   /**
// //    * Occurs when the audio subscribe state changed.
// //    *
// //    * @param channel The channel name of user joined.
// //    * @param userId The remote string user ID that is subscribed to.
// //    * @param oldState The old state of the audio stream subscribe : #STREAM_SUBSCRIBE_STATE.
// //    * @param newState The new state of the audio stream subscribe : #STREAM_SUBSCRIBE_STATE.
// //    * @param elapseSinceLastState The time elapsed (ms) from the old state to the new state.
// //    */
// //   virtual void onAudioSubscribeStateChanged(const char* channel, user_id_t userId, STREAM_SUBSCRIBE_STATE oldState, STREAM_SUBSCRIBE_STATE newState, int elapseSinceLastState) = 0;

// //   /**
// //    * Occurs when the video subscribe state changed.
// //    *
// //    * @param channel The channel name of user joined.
// //    * @param userId The remote string user ID that is subscribed to.
// //    * @param oldState The old state of the video stream subscribe : #STREAM_SUBSCRIBE_STATE.
// //    * @param newState The new state of the video stream subscribe : #STREAM_SUBSCRIBE_STATE.
// //    * @param elapseSinceLastState The time elapsed (ms) from the old state to the new state.
// //    */
// //   virtual void onVideoSubscribeStateChanged(const char* channel, user_id_t userId, STREAM_SUBSCRIBE_STATE oldState, STREAM_SUBSCRIBE_STATE newState, int elapseSinceLastState) = 0;

// //   /**
// //    * Occurs when the audio publish state changed.
// //    *
// //    * @param channel The channel name of user joined.
// //    * @param oldState The old state of the audio stream publish : #STREAM_PUBLISH_STATE.
// //    * @param newState The new state of the audio stream publish : #STREAM_PUBLISH_STATE.
// //    * @param elapseSinceLastState The time elapsed (ms) from the old state to the new state.
// //    */
// //   virtual void onAudioPublishStateChanged(const char* channel, STREAM_PUBLISH_STATE oldState, STREAM_PUBLISH_STATE newState, int elapseSinceLastState) = 0;

// //   /**
// //    * Occurs when the video publish state changed.
// //    *
// //    * @param channel The channel name of user joined.
// //    * @param oldState The old state of the video stream publish : #STREAM_PUBLISH_STATE.
// //    * @param newState The new state of the video stream publish : #STREAM_PUBLISH_STATE.
// //    * @param elapseSinceLastState The time elapsed (ms) from the old state to the new state.
// //    */
// //   virtual void onVideoPublishStateChanged(const char* channel, STREAM_PUBLISH_STATE oldState, STREAM_PUBLISH_STATE newState, int elapseSinceLastState) = 0;

// //   /** Occurs when the first remote audio frame is received.
// //    *
// //    * @param userId ID of the remote user.
// //    * @param isFallbackOrRecover Whether the remotely subscribed media stream
// //    * falls back to audio-only or switches back to the video:
// //    * - true: The remotely subscribed media stream falls back to audio-only
// //    * due to poor network conditions.
// //    * - false: The remotely subscribed media stream switches back to the
// //    * video stream after the network conditions improved.
// //    **/
// //   virtual void onRemoteSubscribeFallbackToAudioOnly(user_id_t userId, bool isFallbackOrRecover) {
// //     (void)userId;
// //     (void)isFallbackOrRecover;
// //   }

// //   /** Occurs when the first remote audio frame is received.
// //    *
// //    * @param userId ID of the remote user.
// //    * @param elapsed The time (ms) since the user connects to an Agora channel.
// //    **/
// //   virtual void onFirstRemoteAudioFrame(user_id_t userId, int elapsed) = 0;

// //   /**
// //    * Occurs when the SDK decodes the first remote audio frame for playback.
// //    *
// //    * @param uid User ID of the remote user sending the audio stream.
// //    * @param elapsed The time (ms) since the user connects to an Agora channel.
// //    */
// //   virtual void onFirstRemoteAudioDecoded(user_id_t userId, int elapsed) = 0;

// //   /**
// //    * Occurs when the first remote video frame is rendered.
// //    * The SDK triggers this callback when the first frame of the remote video is displayed in the user's video window. The application can get the time elapsed from a user joining the channel until the first video frame is displayed.
// //    *
// //    * @param userId ID of the remote user.
// //    * @param width Width (px) of the video frame.
// //    * @param height Height (px) of the video stream.
// //    * @param elapsed Time elapsed (ms) from the local user calling the \ref IRtcEngine::joinChannel "joinChannel" method until the SDK triggers this callback.
// //    */
// //   virtual void onFirstRemoteVideoFrame(user_id_t userId, int width, int height, int elapsed) = 0;

// //   /**
// //    * Occurs when the SDK decodes the first remote video frame for playback.
// //    *
// //    * @param userId ID of the remote user.
// //    * @param width Width (px) of the video stream.
// //    * @param height Height (px) of the video stream.
// //    * @param elapsed The time (ms) since the user connects to an Agora channel.
// //    */
// //   virtual void onFirstRemoteVideoDecoded(user_id_t userId, int width, int height, int elapsed) = 0;

// //   /**
// //    * The local or remote video size or rotation changed.
// //    *
// //    * @param uid User ID of the user whose video size or rotation has changed.
// //    * @param width Width (pixels) of the video stream.
// //    * @param height Height (pixels) of the video stream.
// //    * @param rotation Rotation [0 to 360).
// //    */
// //   virtual void onVideoSizeChanged(user_id_t userId, int width, int height, int rotation) = 0;

/**
 * The media information of a specified user.
 */
type USER_MEDIA_INFO C.enum_C_USER_MEDIA_INFO

const (
	/**
	 * 0: The user has muted the audio.
	 */
	USER_MEDIA_INFO_MUTE_AUDIO USER_MEDIA_INFO = C.USER_MEDIA_INFO_MUTE_AUDIO
	/**
	 * 1: The user has muted the video.
	 */
	USER_MEDIA_INFO_MUTE_VIDEO USER_MEDIA_INFO = C.USER_MEDIA_INFO_MUTE_VIDEO
	/**
	 * 4: The user has enabled the video, which includes video capturing and encoding.
	 */
	USER_MEDIA_INFO_ENABLE_VIDEO USER_MEDIA_INFO = C.USER_MEDIA_INFO_ENABLE_VIDEO
	/**
	 * 8: The user has enabled the local video capturing.
	 */
	USER_MEDIA_INFO_ENABLE_LOCAL_VIDEO USER_MEDIA_INFO = C.USER_MEDIA_INFO_ENABLE_LOCAL_VIDEO
)

// //   /**
// //    * Occurs when the user media information is updated.
// //    *
// //    *
// //    * @param userId The ID of the user.
// //    * @param msg The media information of the user. See #USER_MEDIA_INFO.
// //    * @param val Whether the user is muted.
// //    */
// //   virtual void onUserInfoUpdated(user_id_t userId, USER_MEDIA_INFO msg, bool val) {
// //     (void)userId;
// //     (void)msg;
// //     (void)val;
// //   }

// //   /**
// //    * Occurs when the intra request is received from a remote user.
// //    *
// //    * The method notifies the local user to encode a key frame.
// //    *
// //    */
// //   virtual void onIntraRequestReceived() {}

// //   /**
// //    * datastream from this connection.
// //    */
// //   virtual void onStreamMessage(user_id_t userId, int streamId, const char* data, size_t length) {}

// //   /**
// //    * Occurs when the remote user state is updated.
// //    * @param uid The uid of the remote user.
// //    * @param state The remote user state.Just & #REMOTE_USER_STATE
// //    */
// //   virtual void onUserStateChanged(user_id_t userId, uint32_t state){}

// //   virtual void onVideoRenderingTracingResult(user_id_t user_id, MEDIA_TRACE_EVENT currentState, VideoRenderingTracingInfo tracingInfo) {}
// // };
// #endregion ILocalUserObserver

//   typedef void C_IVideoFrameObserver2;
// #region C_IVideoFrameObserver2

//   // class IVideoFrameObserver2 {
//   //  public:
//   //   /**
//   //    * Occurs each time the player receives a video frame.
//   //    *
//   //    * After registering the video frame observer,
//   //    * the callback occurs each time receives a video frame to report the detailed information of the video frame.
//   //    * @param channelId The channel name
//   //    * @param remoteUid ID of the remote user.
//   //    * @param frame The detailed information of the video frame. See {@link VideoFrame}.
//   //    */
//   //   virtual void onFrame(const char* channelId, user_id_t remoteUid, const media::base::VideoFrame* frame) = 0;

//   //   virtual ~IVideoFrameObserver2() {}
//   // };

// #endregion ILocalUserObserver

/**
 * The ILocalUser class defines the behavior and state of a local user.
 *
 * Each RTC connection has its own local user. Apps can get the local
 * user object by calling \ref agora::rtc::IRtcConnection::getLocalUser
 * "IRtcConnection::getLocalUser".
 *
 * Each local user has two user roles: broadcaster (publisher and subscriber) and
 * audience (subscriber only). A publisher publishes audio and video tracks, while
 * audience receive them.
 */
type ILocalUser C.C_ILocalUser

// #region ILocalUser
//   // class ILocalUser {
//   //  public:
//   //   /**
//   //    * The statistics related to audio network adaptation (ANA).
//   //    */
//   //   struct ANAStats {
//   //     /**
//   //      * The number of actions taken by the ANA bitrate controller since the start of the call.
//   //      *
//   //      * If you do not set this parameter, the ANA bitrate controller is disabled.
//   //      */
//   //     agora::Optional<uint32_t> bitrate_action_counter;
//   //     /**
//   //      * The number of actions taken by the ANA channel controller since the start of the call.
//   //      *
//   //      * If you do not set this parameter, the ANA channel controller is disabled.
//   //      */
//   //     agora::Optional<uint32_t> channel_action_counter;
//   //     /**
//   //      * The number of actions taken by the ANA DTX controller since the start of the call.
//   //      *
//   //      * If you do not set this parameter, the ANA DTX controller is disabled.
//   //      */
//   //     agora::Optional<uint32_t> dtx_action_counter;
//   //     /**
//   //      * The number of actions taken by the ANA FEC (Forward Error Correction) controller since the start of the call.
//   //      *
//   //      * If you do not set this parameter, the ANA FEC controller is disabled.
//   //      */
//   //     agora::Optional<uint32_t> fec_action_counter;
//   //     /**
//   //      * The number of times that the ANA frame length controller increases the frame length
//   //      * since the start of the call.
//   //      *
//   //      * If you do not set this parameter, the ANA frame length controller is disabled.
//   //      */
//   //     agora::Optional<uint32_t> frame_length_increase_counter;
//   //     /**
//   //      * The number of times that the ANA frame length controller decreases the frame length
//   //      * since the start of the call.
//   //      *
//   //      * If you so not set this parameter, the ANA frame length controller is disabled.
//   //      */
//   //     agora::Optional<uint32_t> frame_length_decrease_counter;
//   //     /**
//   //      * The uplink packet loss fractions set by the ANA FEC controller.
//   //      *
//   //      * If you do not set this parameter, the ANA FEC controller is not active.
//   //      */
//   //     agora::Optional<float> uplink_packet_loss_fraction;
//   //   };

//   //   /**
//   //    * The statistics related to audio processing.
//   //    */
//   //   struct AudioProcessingStats {
//   //     /**
//   //      * The echo return loss (ERL).
//   //      *
//   //      * ERL = 10log_10(P_far / P_echo).
//   //      *
//   //      * ERL measures the signal loss that comes back as an echo. A higher ratio corresponds to a smaller
//   //      * amount of echo. The higher the ERL the better.
//   //      */
//   //     agora::Optional<double> echo_return_loss;
//   //     //
//   //     /**
//   //      * The echo return loss enhancement (ERLE).
//   //      *
//   //      * ERLE = 10log_10(P_echo / P_out).
//   //      *
//   //      * The performance of an echo canceller is measured in echo return loss enhancement, which is the amount
//   //      * of additional signal loss applied by the echo canceller.
//   //      *
//   //      * The total signal loss of the echo is the sum of ERL and ERLE.
//   //      */
//   //     agora::Optional<double> echo_return_loss_enhancement;
//   //     /**
//   //      * The fraction of time that the AEC (Acoustic Echo Cancelling) linear filter is divergent, in a
//   //      * 1-second non-overlapped aggregation window.
//   //      */
//   //     agora::Optional<double> divergent_filter_fraction;

//   //     /**
//   //      * The delay metrics (ms).
//   //      *
//   //      * It consists of the delay median and standard deviation. It also consists of the
//   //      * fraction of delay estimates that can make the echo cancellation perform poorly. The values are
//   //      * aggregated until the first call of \ref agora::rtc::IRemoteAudioTrack::getStatistics "getStatistics" and afterwards aggregated and updated every
//   //      * second.
//   //      * @note
//   //      * If there are several clients pulling metrics from
//   //      * `getStatistics` during a session, the first call from any of them will change to one second
//   //      * aggregation window for all.
//   //      */
//   //     agora::Optional<int32_t> delay_median_ms;
//   //     /**
//   //      * The delay standard deviation(ms).
//   //      */
//   //     agora::Optional<int32_t> delay_standard_deviation_ms;

//   //     /**
//   //      * The residual echo detector likelihood.
//   //      */
//   //     agora::Optional<double> residual_echo_likelihood;
//   //     /**
//   //      * The maximum residual echo likelihood from the last time period.
//   //      */
//   //     agora::Optional<double> residual_echo_likelihood_recent_max;

//   //     /**
//   //      * The instantaneous delay estimate produced in the AEC (ms).
//   //      * The value is the instantaneous value at the time of calling \ref agora::rtc::IRemoteAudioTrack::getStatistics "getStatistics".
//   //      */
//   //     agora::Optional<int32_t> delay_ms;
//   //   };

//   //   /**
//   //    * The detailed statistics of the local audio.
//   //    */
//   //   struct LocalAudioDetailedStats {
//   //     /**
//   //      * The synchronization source (SSRC) to identify the stream of RTP packets from the local audio.
//   //      */
//   //     uint32_t local_ssrc;
//   //     /**
//   //      * The number of sent audio bytes.
//   //      */
//   //     int64_t bytes_sent;
//   //     /**
//   //      * The number of sent audio packets.
//   //      */
//   //     int32_t packets_sent;
//   //     /**
//   //      * The number of lost audio packets.
//   //      */
//   //     int32_t packets_lost;
//   //     /**
//   //      * The fraction packet loss reported for this SSRC.
//   //      */
//   //     float fraction_lost;
//   //     /**
//   //      * The codec name.
//   //      */
//   //     char codec_name[media::base::kMaxCodecNameLength];
//   //     /**
//   //      * The payload type as used in RTP encoding or decoding.
//   //      */
//   //     agora::Optional<int> codec_payload_type;
//   //     /**
//   //      * The ext sequence number.
//   //      */
//   //     int32_t ext_seqnum;
//   //     /**
//   //      * The jitter duration (ms).
//   //      */
//   //     int32_t jitter_ms;
//   //     /**
//   //      * The RTT (Round-Trip Time) duration (ms).
//   //      */
//   //     int64_t rtt_ms;
//   //     /**
//   //      * The audio level (dBov) of the media source.
//   //      */
//   //     int32_t audio_level;
//   //     /**
//   //      * The audio energy of the media source.
//   //      */
//   //     double total_input_energy;
//   //     /**
//   //      * The audio duration of the media source.
//   //      */
//   //     double total_input_duration;
//   //     /**
//   //      * Whether the typing noise is detected.
//   //      * - `true`: The typing noise is detected.
//   //      * - `false`: The typing noise is not detected.
//   //      */
//   //     bool typing_noise_detected;

//   //     /**
//   //      * The audio process time from record done to encode done
//   //      */

//   //     ANAStats ana_statistics;
//   //     AudioProcessingStats apm_statistics;

//   //     LocalAudioDetailedStats() : local_ssrc(0), bytes_sent(0), packets_sent(0), packets_lost(-1), fraction_lost(-1.0f),
//   //                                 ext_seqnum(-1), jitter_ms(-1), rtt_ms(-1), audio_level(-1),
//   //                                 total_input_energy(0.0), total_input_duration(0.0), typing_noise_detected(false) {
//   //       memset(codec_name, 0, sizeof(codec_name));
//   //     }
//   //   };

type NS_MODE C.enum_C_NS_MODE

const (
	ElderNsStatistical NS_MODE = C.ElderNsStatistical /* Elder Statistical Noise Suppression.*/
	NsNGStatistical    NS_MODE = C.NsNGStatistical    /* Next Generation Statistical Noise Suppression.*/
	NsNG               NS_MODE = C.NsNG               /* Next Generation Noise Suppression.*/
)

type NS_LEVEL C.enum_C_NS_LEVEL

const (
	Soft       NS_LEVEL = C.Soft       /* Soft Noise Suppression.*/
	Aggressive NS_LEVEL = C.Aggressive /* Aggressiveness Noise Suppression.*/
)

type NS_DELAY C.enum_C_NS_DELAY

const (
	HighQuality NS_DELAY = C.HighQuality /* High Audio Quality with High Delay.*/
	Balance     NS_DELAY = C.Balance     /* Balanced Audio Quality and Delay.*/
	LowDelay    NS_DELAY = C.LowDelay    /* Slight Low Audio Quality with Low Delay.*/
)

func (this_ *ILocalUser) Delete() {
	C.C_ILocalUser_Delete(unsafe.Pointer(this_))
}

/**
 * Sets the role of the user.
 *
 * You can call this method either before or after connecting to an Agora channel:
 * - Before connecting: This method sets the user role as publisher or subscriber (default).
 * - After connecting: This method allows you to switch the user role between publisher and
 * subscriber.
 * The \ref IRtcConnectionObserver::onChangeRoleSuccess "onChangeRoleSuccess" and
 * \ref IRtcConnectionObserver::onChangeRoleFailure "onChangeRoleFailure"
 * callbacks indicate the result of this method call.
 *
 * @note
 * If the token in the \ref IRtcConnection::connect "connect" method does not have the same role
 * as `role`, the connection fails with the \ref IRtcConnectionObserver::onConnectionFailure "onConnectionFailure" callback.
 * @param role The role of the user. See \ref rtc::CLIENT_ROLE_TYPE "CLIENT_ROLE_TYPE".
 */
func (this_ *ILocalUser) SetUserRole(role CLIENT_ROLE_TYPE) {
	C.C_ILocalUser_setUserRole(unsafe.Pointer(this_),
		C.enum_C_CLIENT_ROLE_TYPE(role),
	)
}

//   /**
//    * Gets the role of the user.
//    *
//    * @return The role of the user: Success.
//    */
//   enum C_CLIENT_ROLE_TYPE C_ILocalUser_getUserRole(C_ILocalUser *this_);

//   void C_ILocalUser_setAudienceLatencyLevel(C_ILocalUser *this_, enum C_AUDIENCE_LATENCY_LEVEL_TYPE level);

//   enum C_AUDIENCE_LATENCY_LEVEL_TYPE C_ILocalUser_getAudienceLatencyLevel(C_ILocalUser *this_);

//   /**
//    * Configures the audio encoder.
//    *
//    * The SDK applies the configurations to all the sending audio tracks.
//    *
//    * @param config The reference to the audio encoder configurations.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_setAudioEncoderConfiguration(C_ILocalUser *this_, const struct C_AudioEncoderConfiguration *config);

/**
 * Sets the audio parameters and application scenarios.
 *
 * @param scenario Sets the audio application scenarios: #AUDIO_SCENARIO_TYPE.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SetAudioScenario(scenario AUDIO_SCENARIO_TYPE) int {
	ret := int(C.C_ILocalUser_setAudioScenario(unsafe.Pointer(this_),
		C.enum_C_AUDIO_SCENARIO_TYPE(scenario),
	))
	return ret
}

/**
 *  You can call this method to set the expected video scenario.
 * The SDK will optimize the video experience for each scenario you set.
 *
 * @param  scenarioType The video application scenario.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SetVideoScenario(scenarioType VIDEO_APPLICATION_SCENARIO_TYPE) int {
	ret := int(C.C_ILocalUser_setVideoScenario(unsafe.Pointer(this_),
		C.enum_C_VIDEO_APPLICATION_SCENARIO_TYPE(scenarioType),
	))
	return ret
}

//   //   /**
//   //    * Gets the detailed statistics of the local audio.
//   //    *
//   //    * @param[out] stats The reference to the detailed statistics of the local audio.
//   //    * @return
//   //    * - `true`: Success.
//   //    * - `false`: Failure.
//   //    */
//   //   virtual bool C_ILocalUser_getLocalAudioStatistics(LocalAudioDetailedStats& stats) = 0;

//   //   /**
//   //    * Publishes a local audio track to the channel.
//   //    *
//   //    * By default, all published audio tracks are mixed.
//   //    *
//   //    * @param audioTrack The local audio track to be published.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    *   - -5(ERR_REFUSED), if the role of the local user is not broadcaster.
//   //    */
//   //   virtual int C_ILocalUser_publishAudio(agora_refptr<ILocalAudioTrack> audioTrack) = 0;

//   //   /**
//   //    * Stops publishing the local audio track to the channel.
//   //    *
//   //    * @param audioTrack The local audio track that you want to stop publishing.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_unpublishAudio(agora_refptr<ILocalAudioTrack> audioTrack) = 0;

//   //   /**
//   //    * Publishes a local video track to the channel.
//   //    *
//   //    * @param videoTrack The local video track to be published.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_publishVideo(agora_refptr<ILocalVideoTrack> videoTrack) = 0;

//   //   /**
//   //    * Stops publishing the local video track to the channel.
//   //    *
//   //    * @param videoTrack The local video track that you want to stop publishing.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_unpublishVideo(agora_refptr<ILocalVideoTrack> videoTrack) = 0;

/**
 * Subscribes to the audio of a specified remote user in channel.
 *
 * @param userId The ID of the remote user whose audio you want to subscribe to.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 *   - -2(ERR_INVALID_ARGUMENT), if no such user exists or `userId` is invalid.
 */
func (this_ *ILocalUser) SubscribeAudio(userId string) int {
	cUserId := C.CString(userId)
	ret := int(C.C_ILocalUser_subscribeAudio(unsafe.Pointer(this_), cUserId))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Subscribes to the audio of all remote users in the channel.
 *
 * This method also automatically subscribes to the audio of any subsequent user.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SubscribeAllAudio() int {
	ret := int(C.C_ILocalUser_subscribeAllAudio(unsafe.Pointer(this_)))
	return ret
}

/**
 * Stops subscribing to the audio of a specified remote user in the channel.
 *
 * @param userId The ID of the user whose audio you want to stop subscribing to.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 *   - -2(ERR_INVALID_ARGUMENT), if no such user exists or `userId` is invalid.
 */
func (this_ *ILocalUser) UnsubscribeAudio(userId string) int {
	cUserId := C.CString(userId)
	ret := int(C.C_ILocalUser_unsubscribeAudio(unsafe.Pointer(this_), cUserId))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Stops subscribing to the audio of all remote users in the channel.
 *
 * This method automatically stops subscribing to the audio of any subsequent user, unless you explicitly
 * call \ref subscribeAudio "subscribeAudio" or \ref subscribeAllAudio "subscribeAllAudio".
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) UnsubscribeAllAudio() int {
	ret := int(C.C_ILocalUser_unsubscribeAllAudio(unsafe.Pointer(this_)))
	return ret
}

//   /**
//    * Adjusts the playback signal volume.
//    * @param volume The playback volume. The value ranges between 0 and 400, including the following:
//    * 0: Mute.
//    * 100: (Default) Original volume.
//    * 400: Four times the original volume with signal-clipping protection.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_adjustPlaybackSignalVolume(C_ILocalUser *this_, int volume);

//   /**
//    * Gets the current playback signal volume.
//    * @param volume A pointer to the playback signal volume.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_getPlaybackSignalVolume(C_ILocalUser *this_, int *volume);

//   /*
//    * Adjust the playback volume of the user specified by uid.
//    *
//    * You can call this method to adjust the playback volume of the user specified by uid
//    * in call. If you want to adjust playback volume of the multi user, you can call this
//    * this method multi times.
//    *
//    * @note
//    * Please call this method after join channel.
//    * This method adjust the playback volume of specified user.
//    *
//    * @param userId The ID of the Remote user.
//    * @param volume The playback volume of the specified remote user. The value ranges between 0 and 400, including the following:
//    * 0: Mute.
//    * 100: (Default) Original volume.
//    * 400: Four times the original volume with signal-clipping protection.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_adjustUserPlaybackSignalVolume(C_ILocalUser *this_, C_user_id_t userId, int volume);

//   /**
//    * Gets the current playback signal volume of specified user.
//    * @param userId The ID of the Remote user.
//    * @param volume A pointer to the playback signal volume.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_getUserPlaybackSignalVolume(C_ILocalUser *this_, C_user_id_t userId, int *volume);

/*
* Enables/Disables stereo panning for remote users.

	Ensure that you call this method before joinChannel to enable stereo panning for remote users so that the local user can track the position of a remote user by calling \ref agora::rtc::IRtcEngine::setRemoteVoicePosition "setRemoteVoicePosition".

	@param enabled Sets whether or not to enable stereo panning for remote users:
	- true: enables stereo panning.
	- false: disables stereo panning.

	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *ILocalUser) EnableSoundPositionIndication(enabled bool) int {
	ret := int(C.C_ILocalUser_enableSoundPositionIndication(unsafe.Pointer(this_),
		C.bool(enabled),
	))
	return ret
}

//   /** Sets the sound position and gain of a remote user.

//    When the local user calls this method to set the sound position of a remote user, the sound difference between the left and right channels allows the local user to track the real-time position of the remote user, creating a real sense of space. This method applies to massively multiplayer online games, such as Battle Royale games.

//    @note
//    - For this method to work, enable stereo panning for remote users by calling the \ref agora::rtc::IRtcEngine::enableSoundPositionIndication "enableSoundPositionIndication" method before joining a channel.
//    - This method requires hardware support. For the best sound positioning, we recommend using a wired headset.
//    - Ensure that you call this method after joining a channel.

//    @param userId The ID of the remote user.
//    @param pan The sound position of the remote user. The value ranges from -1.0 to 1.0:
//    - 0.0: the remote sound comes from the front.
//    - -1.0: the remote sound comes from the left.
//    - 1.0: the remote sound comes from the right.
//    @param gain Gain of the remote user. The value ranges from 0.0 to 100.0. The default value is 100.0 (the original gain of the remote user). The smaller the value, the less the gain.

//    @return
//    - 0: Success.
//    - < 0: Failure.
//    */
//   int C_ILocalUser_setRemoteVoicePosition(C_ILocalUser *this_, C_user_id_t userId, double pan, double gain);

/*
* enable spatial audio

	@param enabled enable/disable spatial audio:
	- true: enable spatial audio.
	- false: disable spatial audio.
	@return
	- 0: Success.
	- < 0: Failure.
*/
func (this_ *ILocalUser) EnableSpatialAudio(enabled bool) int {
	ret := int(C.C_ILocalUser_enableSpatialAudio(unsafe.Pointer(this_),
		C.bool(enabled),
	))
	return ret
}

//   //   /** Sets remote user parameters for spatial audio

//   //    @param userId The ID of the remote user.
//   //    @param param spatial audio parameters

//   //    @return
//   //    - 0: Success.
//   //    - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_setRemoteUserSpatialAudioParams(user_id_t userId, const agora::SpatialAudioParams& param) = 0;

/**
 * Sets the audio frame parameters for the \ref agora::media::IAudioFrameObserver::onPlaybackAudioFrame
 * "onPlaybackAudioFrame" callback.
 *
 * @param numberOfChannels The number of audio channels of the audio frame in the `onPlaybackAudioFrame` callback.
 * - 1: Mono.
 * - 2: Stereo.
 * @param sampleRateHz The sample rate (Hz) of the audio frame in the `onPlaybackAudioFrame` callback. You can
 * set it as 8000, 16000, 32000, 44100, or 48000.
 * @param mode Use mode of the audio frame. See #RAW_AUDIO_FRAME_OP_MODE_TYPE.
 * @param samplesPerCall The number of samples of the audio frame.   *
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SetPlaybackAudioFrameParameters(numberOfChannels uint,
	sampleRateHz uint32,
	mode RAW_AUDIO_FRAME_OP_MODE_TYPE,
	samplesPerCall int) int {
	ret := int(C.C_ILocalUser_setPlaybackAudioFrameParameters(unsafe.Pointer(this_),
		C.size_t(numberOfChannels),
		C.uint32_t(sampleRateHz),
		C.enum_C_RAW_AUDIO_FRAME_OP_MODE_TYPE(mode),
		C.int(samplesPerCall),
	))
	return ret
}

/**
 * Sets the audio frame parameters for the \ref agora::media::IAudioFrameObserver::onRecordAudioFrame
 * "onRecordAudioFrame" callback.
 *
 * @param numberOfChannels The number of channels of the audio frame in the `onRecordAudioFrame` callback.
 * - 1: Mono.
 * - 2: Stereo.
 * @param sampleRateHz The sample rate (Hz) of the audio frame in the `onRecordAudioFrame` callback. You can
 * set it as 8000, 16000, 32000, 44100, or 48000.
 * @param mode Use mode of the audio frame. See #RAW_AUDIO_FRAME_OP_MODE_TYPE.
 * @param samplesPerCall The number of samples of the audio frame.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SetRecordingAudioFrameParameters(numberOfChannels uint,
	sampleRateHz uint32,
	mode RAW_AUDIO_FRAME_OP_MODE_TYPE,
	samplesPerCall int) int {
	ret := int(C.C_ILocalUser_setRecordingAudioFrameParameters(unsafe.Pointer(this_),
		C.size_t(numberOfChannels),
		C.uint32_t(sampleRateHz),
		C.enum_C_RAW_AUDIO_FRAME_OP_MODE_TYPE(mode),
		C.int(samplesPerCall),
	))
	return ret
}

/**
 * Sets the audio frame parameters for the \ref agora::media::IAudioFrameObserver::onMixedAudioFrame
 * "onMixedAudioFrame" callback.
 *
 * @param numberOfChannels The number of channels of the audio frame in the `onMixedAudioFrame` callback.
 * - 1: Mono.
 * - 2: Stereo.
 * @param sampleRateHz The sample rate (Hz) of the audio frame in the `onMixedAudioFrame` callback. You can
 * set it as 8000, 16000, 32000, 44100, or 48000.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SetMixedAudioFrameParameters(numberOfChannels uint,
	sampleRateHz uint32,
	samplesPerCall int) int {
	ret := int(C.C_ILocalUser_setMixedAudioFrameParameters(unsafe.Pointer(this_),
		C.size_t(numberOfChannels),
		C.uint32_t(sampleRateHz),
		C.int(samplesPerCall),
	))
	return ret
}

/**
 * Sets the audio frame parameters for the \ref agora::media::IAudioFrameObserver::onEarMonitoringAudioFrame
 * "onEarMonitoringAudioFrame" callback.
 * @param enabled Determines whether to enable ear monitoring audio frame observer.
 * - true: Enable ear monitoring audio frame observer.
 * - false: Disable ear monitoring audio frame observer.
 * @param numberOfChannels The number of audio channels of the audio frame in the `onEarMonitoringAudioFrame` callback.
 * - 1: Mono.
 * - 2: Stereo.
 * @param sampleRateHz The sample rate (Hz) of the audio frame in the `onEarMonitoringAudioFrame` callback. You can
 * set it as 8000, 16000, 32000, 44100, or 48000.
 * @param mode Use mode of the audio frame. See #RAW_AUDIO_FRAME_OP_MODE_TYPE.
 * @param samplesPerCall The number of samples of the audio frame.   *
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SetEarMonitoringAudioFrameParameters(enabled bool,
	numberOfChannels uint,
	sampleRateHz uint32,
	mode RAW_AUDIO_FRAME_OP_MODE_TYPE,
	samplesPerCall int) int {
	ret := int(C.C_ILocalUser_setEarMonitoringAudioFrameParameters(unsafe.Pointer(this_),
		C.bool(enabled),
		C.size_t(numberOfChannels),
		C.uint32_t(sampleRateHz),
		C.enum_C_RAW_AUDIO_FRAME_OP_MODE_TYPE(mode),
		C.int(samplesPerCall),
	))
	return ret
}

/**
 * Sets the audio frame parameters for the \ref agora::media::IAudioFrameObserver::onPlaybackAudioFrameBeforeMixing
 * "onPlaybackAudioFrameBeforeMixing" callback.
 *
 * @param numberOfChannels The number of channels of the audio frame in the `onPlaybackAudioFrameBeforeMixing` callback.
 * - 1: Mono.
 * - 2: Stereo.
 * @param sampleRateHz The sample rate (Hz) of the audio frame in the `onPlaybackAudioFrameBeforeMixing` callback. You can
 * set it as 8000, 16000, 32000, 44100, or 48000.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SetPlaybackAudioFrameBeforeMixingParameters(numberOfChannels uint, sampleRateHz uint32) int {
	ret := int(C.C_ILocalUser_setPlaybackAudioFrameBeforeMixingParameters(unsafe.Pointer(this_),
		C.size_t(numberOfChannels),
		C.uint32_t(sampleRateHz),
	))
	return ret
}

/**
 * Registers an audio frame observer.
 *
 * You need to implement the `IAudioFrameObserverBase` class in this method, and register the following callbacks
 * according to your scenario:
 * - \ref agora::media::IAudioFrameObserverBase::onRecordAudioFrame "onRecordAudioFrame": Occurs when the SDK receives the audio data captured by the local recording device.
 * - \ref agora::media::IAudioFrameObserverBase::onPlaybackAudioFrame "onPlaybackAudioFrame": Occurs when the SDK receives the remote audio data for playback.
 * - \ref agora::media::IAudioFrameObserverBase::onPlaybackAudioFrameBeforeMixing "onPlaybackAudioFrameBeforeMixing": Occurs when the SDK receives the remote audio data of a specified
 * remote user before mixing.
 * - \ref agora::media::IAudioFrameObserverBase::onMixedAudioFrame "onMixedAudioFrame": Occurs when the SDK receives the mixed data of recorded and playback audio.
 *
 * @param observer A pointer to the audio frame observer: \ref agora::media::IAudioFrameObserverBase "IAudioFrameObserverBase".
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) RegisterAudioFrameObserver(observer *IAudioFrameObserverBase) int {
	return int(C.C_ILocalUser_registerAudioFrameObserver(unsafe.Pointer(this_),
		(unsafe.Pointer(observer)),
	))
}

/**
 * Releases the audio frame observer.
 *
 * @param observer The pointer to the audio frame observer: \ref agora::media::IAudioFrameObserverBase "IAudioFrameObserverBase".
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) UnregisterAudioFrameObserver(observer *IAudioFrameObserverBase) int {
	return int(C.C_ILocalUser_unregisterAudioFrameObserver(unsafe.Pointer(this_),
		(unsafe.Pointer(observer)),
	))
}

//   /**
//    * Enable the audio spectrum monitor.
//    *
//    * @param intervalInMS Sets the time interval(ms) between two consecutive audio spectrum callback.
//    * The default value is 100. This param should be larger than 10.
//    *
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_enableAudioSpectrumMonitor(C_ILocalUser *this_, int intervalInMS);
//   /**
//    * Disalbe the audio spectrum monitor.
//    *
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_disableAudioSpectrumMonitor(C_ILocalUser *this_);

//   /**
//    * Registers an audio spectrum observer.
//    *
//    * You need to implement the `IAudioSpectrumObserver` class in this method, and register the following callbacks
//    * according to your scenario:
//    * - \ref agora::media::IAudioSpectrumObserver::onAudioSpectrumComputed "onAudioSpectrumComputed": Occurs when
//    * the SDK receives the audio data and at set intervals.
//    *
//    * @param observer A pointer to the audio spectrum observer: \ref agora::media::IAudioSpectrumObserver
//    * "IAudioSpectrumObserver".
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_registerAudioSpectrumObserver(C_ILocalUser *this_, C_IAudioSpectrumObserver *observer, void (*safeDeleter)(C_IAudioSpectrumObserver *));
//   /**
//    * Releases the audio spectrum observer.
//    *
//    * @param observer The pointer to the audio spectrum observer: \ref agora::media::IAudioSpectrumObserver
//    * "IAudioSpectrumObserver".
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_unregisterAudioSpectrumObserver(C_ILocalUser *this_, C_IAudioSpectrumObserver *observer);

//   /**
//    * Registers an \ref agora::media::IVideoEncodedFrameObserver "IVideoEncodedFrameObserver" object.
//    *
//    * You need to implement the `IVideoEncodedFrameObserver` class in this method. Once you successfully register
//    * the local encoded frame observer, the SDK triggers the \ref agora::media::IVideoEncodedFrameObserver::onEncodedVideoFrameReceived
//    * "onEncodedVideoFrameReceived" callback when it receives the encoded video image.
//    *
//    * @param observer The pointer to the `IVideoEncodedFrameObserver` object.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_registerLocalVideoEncodedFrameObserver(C_ILocalUser *this_, C_IVideoEncodedFrameObserver *observer);
//   /**
//    * Releases the \ref agora::media::IVideoEncodedFrameObserver "IVideoEncodedFrameObserver" object.
//    * @param observer The pointer to the `IVideoEncodedFrameObserver` object.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_unregisterLocalVideoEncodedFrameObserver(C_ILocalUser *this_, C_IVideoEncodedFrameObserver *observer);
//   /**
//    * Force trigger to intra-frame the next frame.
//    *
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_forceNextIntraFrame(C_ILocalUser *this_);
//   /**
//    * Registers an \ref agora::media::IVideoEncodedFrameObserver "IVideoEncodedFrameObserver" object.
//    *
//    * You need to implement the `IVideoEncodedFrameObserver` class in this method. Once you successfully register
//    * the encoded frame observer, the SDK triggers the \ref agora::media::IVideoEncodedFrameObserver::onEncodedVideoFrameReceived
//    * "onEncodedVideoFrameReceived" callback when it receives the encoded video image.
//    *
//    * @param observer The pointer to the `IVideoEncodedFrameObserver` object.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_registerVideoEncodedFrameObserver(C_ILocalUser *this_, C_IVideoEncodedFrameObserver *observer);
//   /**
//    * Releases the \ref agora::media::IVideoEncodedFrameObserver "IVideoEncodedFrameObserver" object.
//    * @param observer The pointer to the `IVideoEncodedFrameObserver` object.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_unregisterVideoEncodedFrameObserver(C_ILocalUser *this_, C_IVideoEncodedFrameObserver *observer);

//   /**
//    * Registers an \ref agora::rtc::IVideoFrameObserver2 "IVideoFrameObserver2" object.
//    *
//    * You need to implement the `IVideoFrameObserver2` class in this method. Once you successfully register
//    * the video frame observer, the SDK triggers the \ref agora::media::IVideoFrameObserver2::onFrame
//    * "onFrame" callback when it receives the video image.
//    *
//    * @param observer The pointer to the `IVideoFrameObserver2` object.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_registerVideoFrameObserver(C_ILocalUser *this_, C_IVideoFrameObserver2 *observer);
//   /**
//    * Releases the \ref agora::rtc::IVideoFrameObserver2 "IVideoFrameObserver2" object.
//    * @param observer The pointer to the `IVideoFrameObserver2` object.
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_unregisterVideoFrameObserver(C_ILocalUser *this_, C_IVideoFrameObserver2 *observer);

func (this_ *ILocalUser) SetVideoSubscriptionOptions(userId string, options *VideoSubscriptionOptions) int {
	cUserId := C.CString(userId)
	ret := int(C.C_ILocalUser_setVideoSubscriptionOptions(unsafe.Pointer(this_),
		cUserId,
		(*C.struct_C_VideoSubscriptionOptions)(options),
	))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

//   //   virtual int C_ILocalUser_setHighPriorityUserList(uid_t* vipList, int uidNum, int option) = 0;

//   //   virtual int C_ILocalUser_getHighPriorityUserList(std::vector<uid_t>& vipList, int& option) = 0;

//   /**
//    * Sets the blocklist of subscribe remote stream audio.
//    *
//    * @param userList The id list of users who do not subscribe to audio.
//    * @param userNumber The number of uid in uidList.
//    *
//    * @note
//    * If uid is in uidList, the remote user's audio will not be subscribed,
//    * even if subscribeAudio(uid) and subscribeAllAudio(true) are operated.
//    *
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_setSubscribeAudioBlocklist(C_ILocalUser *this_, C_user_id_t *userList, int userNumber);

//   /**
//    * Sets the allowlist of subscribe remote stream audio.
//    *
//    * @param userList The id list of users who do subscribe to audio.
//    * @param userNumber The number of uid in uidList.
//    *
//    * @note
//    * If uid is in uidList, the remote user's audio will be subscribed,
//    * even if unsubscribeAudio(uid) and unsubscribeAllAudio(true) are operated.
//    *
//    * If a user is in the blocklist and allowlist at the same time, the user will not subscribe to audio.
//    *
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_setSubscribeAudioAllowlist(C_ILocalUser *this_, C_user_id_t *userList, int userNumber);

//   /**
//    * Sets the blocklist of subscribe remote stream video.
//    *
//    * @param userList The id list of users who do not subscribe to video.
//    * @param userNumber The number of uid in uidList.
//    *
//    * @note
//    * If uid is in uidList, the remote user's video will not be subscribed,
//    * even if subscribeVideo(uid) and subscribeAllVideo(true) are operated.
//    *
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_setSubscribeVideoBlocklist(C_ILocalUser *this_, C_user_id_t *userList, int userNumber);

//   /**
//    * Sets the allowlist of subscribe remote stream video.
//    *
//    * @param userList The id list of users who do subscribe to video.
//    * @param userNumber The number of uid in uidList.
//    *
//    * @note
//    * If uid is in uidList, the remote user's video will be subscribed,
//    * even if unsubscribeVideo(uid) and unsubscribeAllVideo(true) are operated.
//    *
//    * If a user is in the blocklist and allowlist at the same time, the user will not subscribe to video.
//    *
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    */
//   int C_ILocalUser_setSubscribeVideoAllowlist(C_ILocalUser *this_, C_user_id_t *userList, int userNumber);

/**
 * Subscribes to the video of a specified remote user in the channel.
 *
 * @param userId The ID of the user whose video you want to subscribe to.
 * @param subscriptionOptions The reference to the video subscription options: \ref agora::rtc::VideoSubscriptionOptions "VideoSubscriptionOptions".
 * For example, subscribing to encoded video data only or subscribing to low-stream video.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 *   - -2(ERR_INVALID_ARGUMENT), if `userId` is invalid.
 */
func (this_ *ILocalUser) SubscribeVideo(userId string, subscriptionOptions *VideoSubscriptionOptions) int {
	cUserId := C.CString(userId)
	ret := int(C.C_ILocalUser_subscribeVideo(unsafe.Pointer(this_),
		cUserId,
		(*C.struct_C_VideoSubscriptionOptions)(subscriptionOptions),
	))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Subscribes to the video of all remote users in the channel.
 *
 * This method also automatically subscribes to the video of any subsequent remote user.
 *
 * @param subscriptionOptions The reference to the video subscription options: \ref agora::rtc::VideoSubscriptionOptions "VideoSubscriptionOptions".
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) SubscribeAllVideo(subscriptionOptions *VideoSubscriptionOptions) int {
	ret := int(C.C_ILocalUser_subscribeAllVideo(unsafe.Pointer(this_),
		(*C.struct_C_VideoSubscriptionOptions)(subscriptionOptions),
	))
	return ret
}

/**
 * Stops subscribing to the video of a specified remote user in the channel.
 *
 * @param userId The ID of the user whose video you want to stop subscribing to.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 *   - -2(ERR_INVALID_ARGUMENT), if `userId` is invalid.
 */
func (this_ *ILocalUser) UnsubscribeVideo(userId string) int {
	cUserId := C.CString(userId)
	ret := int(C.C_ILocalUser_unsubscribeVideo(unsafe.Pointer(this_), cUserId))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Stops subscribing to the video of all remote users in the channel.
 *
 * This method automatically stops subscribing to the video of any subsequent user, unless you explicitly
 * call \ref subscribeVideo "subscribeVideo" or \ref subscribeAllVideo "subscribeAllVideo".
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) UnsubscribeAllVideo() int {
	return int(C.C_ILocalUser_unsubscribeAllVideo(unsafe.Pointer(this_)))
}

//   /**
//    * Sets the time interval and the volume smoothing factor of the \ref agora::rtc::ILocalUserObserver::onAudioVolumeIndication "onAudioVolumeIndication" callback.
//    *
//    * By default, the SDK triggers the `onAudioVolumeIndication`
//    * callback once every 500 ms, with a smoothing factor of 3.
//    *
//    * @param intervalInMS Sets the time interval(ms) between two consecutive volume indications. The default
//    * value is 500.
//    * - &le; 10: Disables the volume indication.
//    * - > 10: The time interval (ms) between two consecutive callbacks.
//    *
//    * @param smooth The smoothing factor that sets the sensitivity of the audio volume indicator.
//    * The value range is [0,10]. The greater the value, the more sensitive the indicator.The default value is 3.
//    *
//    * @return
//    * - 0: Success.
//    * - < 0: Failure.
//    *   - -2(ERR_INVALID_ARGUMENT), if `intervalInMS` or `smooth` is out of range.
//    */
//   int C_ILocalUser_setAudioVolumeIndicationParameters(C_ILocalUser *this_, int intervalInMS, int smooth, bool reportVad);

/**
 * Registers a local user observer object.
 *
 * You need to implement the \ref agora::rtc::ILocalUserObserver "ILocalUserObserver" class in this method. Once registered, the
 * ILocalUserObserver receives events of the \ref agora::rtc::ILocalUser "ILocalUser" object.
 *
 * @param observer The pointer to the `ILocalUserObserver` object.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) RegisterLocalUserObserver(observer *ILocalUserObserver) int {
	return int(C.C_ILocalUser_registerLocalUserObserver(unsafe.Pointer(this_),
		(unsafe.Pointer(observer)),
		nil, //TODO
	))
}

/**
 * Releases the \ref agora::rtc::ILocalUserObserver "ILocalUserObserver" object.
 *
 * @param observer The pointer to the `ILocalUserObserver` object that you want to release.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *ILocalUser) UnregisterLocalUserObserver(observer *ILocalUserObserver) int {
	return int(C.C_ILocalUser_unregisterLocalUserObserver(unsafe.Pointer(this_),
		(unsafe.Pointer(observer)),
	))
}

//   // /**
//   //  * Gets the media control packet sender.
//   //  *
//   //  * @return
//   //  * - The pointer to the media control packet sender IMediaControlPacketSender: Success.
//   //  * - A null pointer: Failure.
//   //  */
//   // C_IMediaControlPacketSender* C_ILocalUser_getMediaControlPacketSender();

//   //   /**
//   //    * Registers a media control packet receiver.
//   //    *
//   //    * You need to implement the `IMediaControlPacketReceiver` class in this method. Once you successfully
//   //    * register the media control packet receiver, the SDK triggers the \ref agora::rtc::IMediaControlPacketReceiver::onMediaControlPacketReceived "onMediaControlPacketReceived"
//   //    * callback when it receives the media control packet.
//   //    *
//   //    * @param ctrlPacketReceiver The pointer to the IMediaControlPacketReceiver object.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_registerMediaControlPacketReceiver(IMediaControlPacketReceiver* ctrlPacketReceiver) = 0;

//   //   /**
//   //    * Releases the media control packet receiver.
//   //    * @param ctrlPacketReceiver The pointer to the `IMediaControlPacketReceiver` object.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_unregisterMediaControlPacketReceiver(IMediaControlPacketReceiver* ctrlPacketReceiver) = 0;

//   //   /**
//   //    * Sends intra-frame request to the broadcaster with a specified uid.
//   //    *
//   //    * The local user receives the `onIntraRequestReceived` callback when the broadcaster receives the request.
//   //    *
//   //    * @param userId The user ID of the target broadcaster .
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_sendIntraRequest(user_id_t userId) = 0;

//   //   /**
//   //    * Set local audio filterable by topn
//   //    *
//   //    * The local user receives the `onIntraRequestReceived` callback when the broadcaster receives the request.
//   //    *
//   //    * @param userId The user ID of the target broadcaster .
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */

//   //   virtual int C_ILocalUser_setAudioFilterable(bool filterable) = 0;

//   // /**
//   //   * Enable / Disable specified audio filter
//   //   * @param userId The ID of the remote user
//   //   * @param id id of the filter
//   //   * @param enable enable / disable the filter with given id
//   //   * @return
//   //   * - 0: success
//   //   * - <0: failure
//   //   */
//   //  virtual int C_ILocalUser_enableRemoteAudioTrackFilter(user_id_t userId, const char* id, bool enable) = 0;

//   //  /**
//   //   * set the properties of the specified audio filter
//   //   * @param userId The ID of the remote user
//   //   * @param id id of the filter
//   //   * @param key key of the property
//   //   * @param jsonValue json str value of the property
//   //   * @return
//   //   * - 0: success
//   //   * - <0: failure
//   //   */
//   //  virtual int C_ILocalUser_setRemoteAudioTrackFilterProperty(user_id_t userId, const char* id, const char* key, const char* jsonValue) = 0;

//   //  /**
//   //   * get the properties of the specified audio filter
//   //   * @param userId The ID of the remote user
//   //   * @param id id of the filter
//   //   * @param key key of the property
//   //   * @param jsonValue json str value of the property
//   //   * @param bufSize max length of the json value buffer
//   //   * @return
//   //   * - 0: success
//   //   * - <0: failure
//   //   */
//   //  virtual int C_ILocalUser_getRemoteAudioTrackFilterProperty(user_id_t userId, const char* id, const char* key, char* jsonValue, size_t bufSize) = 0;
//   //   /**
//   //    * Publishes a local data channel to the channel.
//   //    *
//   //    * @param channel The data stream to be published.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_publishDataChannel(agora_refptr<ILocalDataChannel> channel) = 0;
//   //   /**
//   //    * Stops publishing the data channel to the channel.
//   //    *
//   //    * @param channel The data channel that you want to stop publishing.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_unpublishDataChannel(agora_refptr<ILocalDataChannel> channel) = 0;
//   //   /**
//   //    * Subscribes to a specified data channel of a specified remote user in channel.
//   //    *
//   //    * @param userId The ID of the remote user whose data channel you want to subscribe to.
//   //    * @param channelId The ID of the data channel that you want to subscribe to.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_subscribeDataChannel(user_id_t userId, int channelId) = 0;
//   //   /**
//   //    * Stops subscribing to the data channel of a specified remote user in the channel.
//   //    *
//   //    * @param userId The ID of the remote user whose data channel you want to stop subscribing to.
//   //    * @param channelId The ID of the data channel that you want to stop subscribing to.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    *   - -2(ERR_INVALID_ARGUMENT), if no such user exists or `userId` is invalid.
//   //    */

//   //   virtual int C_ILocalUser_unsubscribeDataChannel(user_id_t userId, int channelId) = 0;
//   //   /**
//   //    * Registers an data channel observer.
//   //    *
//   //    * You need to implement the `IDataChannelObserver` class in this method
//   //    *
//   //    * @param observer A pointer to the data channel observer:
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_registerDataChannelObserver(IDataChannelObserver * observer) = 0;
//   //   /**
//   //    * Releases the data channel observer.
//   //    *
//   //    * @param observer The pointer to the data channel observer
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int C_ILocalUser_unregisterDataChannelObserver(IDataChannelObserver * observer) = 0;
//   //   /**
//   //    * set the profile of audio noise suppression module
//   //    *
//   //    * @param NsEnable enable ns or not
//   //    * @param NsMode type of ns
//   //    * @param NsLevel level of the suppression
//   //    * @param NsDelay algorithm delay
//   //    * @return
//   //    * - 0: success
//   //    * - <0: failure
//   //   */
//   //   virtual int C_ILocalUser_SetAudioNsMode(bool NsEnable, NS_MODE NsMode, NS_LEVEL NsLevel, NS_DELAY NsDelay) = 0;
//   //   /**
//   //    * register a mix source to capture the track of this user.
//   //    *
//   //    * @param mixSource The mix source that is the sink attach to track.
//   //    * @param mixPolicy Policy of whether mixed local or remote.
//   //    * @return
//   //    * - 0: success
//   //    * - <0: failure
//   //   */
//   //   virtual int C_ILocalUser_registerMixSource(agora_refptr<IAudioMixerSource> mixSource, int mixPolicy) = 0;
//   // };

// #endregion ILocalUser

// #endregion agora::rtc

// #endregion agora

// #ifdef __cplusplus
// }
// #endif // __cplusplus

// #endif // C_NGI_AGORA_LOCAL_USER_H
