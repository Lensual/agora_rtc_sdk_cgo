package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/ -lagora_rtc_sdk_c -lstdc++

//链接AgoraRTC SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtc_sdk_c/agora_rtc_sdk_download-prefix/src/agora_rtc_sdk_download/agora_sdk -lagora_rtc_sdk -lagora-ffmpeg -lagora-fdkaac

#include "C_NGIAgoraVideoTrack.h"
#include <stdlib.h>
*/
import "C"

// #ifndef C_NGI_AGORA_VIDEO_TRACK_H
// #define C_NGI_AGORA_VIDEO_TRACK_H

// #include "C_AgoraBase.h"

// // #ifndef OPTIONAL_OVERRIDE
// // #if __cplusplus >= 201103L || (defined(_MSC_VER) && _MSC_VER >= 1800)
// // #define OPTIONAL_OVERRIDE override
// // #else
// // #define OPTIONAL_OVERRIDE
// // #endif
// // #endif

// #ifdef __cplusplus
// extern "C"
// {
// #endif // __cplusplus

// #region agora

// #region agora::rtc

//   // class IVideoFilter;
//   // class IVideoEncodedFrameObserver;
//   // class IMediaPacketReceiver;
//   // class IVideoSinkBase;

type VideoTrackType C.enum_C_VideoTrackType

const (
	LOCAL_VIDEO_TRACK        VideoTrackType = C.LOCAL_VIDEO_TRACK
	REMOTE_VIDEO_TRACK       VideoTrackType = C.REMOTE_VIDEO_TRACK
	REMOTE_VIDEO_IMAGE_TRACK VideoTrackType = C.REMOTE_VIDEO_IMAGE_TRACK
)

//   /**
//    * The `IVideoTrack` class defines the behavior and status of a video track.
//    */
//   typedef void C_IVideoTrack;
// #region C_IVideoTrack

//   // class IVideoTrack : public RefCountInterface {
//   //  public:
//   //   /**
//   //    * Adds a video filter to the video track.
//   //    *
//   //    * Add a video filter in either of the following ways:
//   //    * - Use the \ref agora::rtc::IMediaNodeFactory "IMediaNodeFactory" object to create a built-in video filter.
//   //    * - Use a custom video filter by implementing the \ref agora::rtc::IVideoFilter "IVideoFilter" class.
//   //    *
//   //    * To add multiple filters, call this method multiple times. The order of the added filters depends on when
//   //    * the app successfully adds the filter.
//   //    *
//   //    * @param filter The video filter that you want to add to the video track.
//   //    * @param position The position where the filter is added.
//   //    * @param id id of the filter
//   //    * @return
//   //    * - `true`: The video filter is added successfully.
//   //    * - `false`: The video filter fails to be added.
//   //    */
//   //   virtual bool addVideoFilter(
//   //       agora_refptr<IVideoFilter> filter, media::base::VIDEO_MODULE_POSITION position = media::base::POSITION_POST_CAPTURER,
//   //       const char* id = NULL) = 0;

//   //   /**
//   //    * Removes the video filter added by `addVideoFilter` from the video track.
//   //    *
//   //    * @param filter The video filter that you want to remove: `IVideoFilter`.
//   //    * @param position The position of the filter.
//   //    * @id id of the filter
//   //    * @return
//   //    * - `true`: The video filter is removed successfully.
//   //    * - `false`: The video filter fails to be removed.
//   //    */
//   //   virtual bool removeVideoFilter(
//   //       agora_refptr<IVideoFilter> filter, media::base::VIDEO_MODULE_POSITION position = media::base::POSITION_POST_CAPTURER,
//   //       const char* id = NULL) = 0;

//   //   /**
//   //    * Whether a video filter exists
//   //    * @param id id of the filter
//   //    * @return
//   //    * - true: exist
//   //    * - false: not exist
//   //    */
//   //   virtual bool hasVideoFilter(const char* id, media::base::VIDEO_MODULE_POSITION position = media::base::POSITION_POST_CAPTURER) = 0;

//   //   /**
//   //    * Adds a video renderer to the video track.
//   //    *
//   //    * Add a video renderer in either of the following ways:
//   //    * - Use the built-in video renderer by implementing the `IVideoRenderer` in the `IMediaNodeFactory` class.
//   //    * - Use a custom video renderer by implementing the `IVideoSinkBase` class.
//   //    *
//   //    * @param videoRenderer The video renderer that you want to add: IVideoSinkBase.
//   //    * @param position The position where the renderer is added.
//   //    *
//   //    * @return
//   //    * - `true`: The video renderer is added successfully.
//   //    * - `false`: The video renderer fails to be added.
//   //    */
//   //   virtual bool addRenderer(agora_refptr<IVideoSinkBase> videoRenderer, media::base::VIDEO_MODULE_POSITION position = media::base::POSITION_PRE_RENDERER) = 0;
//   //   /**
//   //    * Removes the video renderer added by `addRenderer` from the video track.
//   //    *
//   //    * @param videoRenderer The video renderer that you want to remove: IVideoSinkBase.
//   //    * @param position The position where the renderer is removed: \ref media::base::VIDEO_MODULE_POSITION "VIDEO_MODULE_POSITION".
//   //    * @return
//   //    * - `true`: The video renderer is removed successfully.
//   //    * - `false`: The video renderer fails to be removed.
//   //    */
//   //   virtual bool removeRenderer(agora_refptr<IVideoSinkBase> videoRenderer, media::base::VIDEO_MODULE_POSITION position = media::base::POSITION_PRE_RENDERER) = 0;
//   //   /**
//   //    * Get the track type of the video track
//   //    * @return
//   //    * - VideoTrackType
//   //    */
//   //   virtual VideoTrackType getType() = 0;

//   //   /**
//   //    * Enable / Disable specified video filter
//   //    * @param id id of the filter
//   //    * @param enable enable / disable the filter with given id
//   //    * @return
//   //    * - 0: success
//   //    * - <0: failure
//   //    */
//   //   virtual int enableVideoFilter(const char* id, bool enable) { return -1; }

//   //   /**
//   //    * set the properties of the specified video filter
//   //    * @param id id of the filter
//   //    * @param key key of the property
//   //    * @param json_value json str value of the property
//   //    * @return
//   //    * - 0: success
//   //    * - <0: failure
//   //    */
//   //   virtual int setFilterProperty(const char* id, const char* key, const char* json_value) { return -1; }

//   //   /**
//   //    * get the properties of the specified video filter
//   //    * @param id id of the filter
//   //    * @param key key of the property
//   //    * @param json_value json str value of the property
//   //    * @return
//   //    * - 0: success
//   //    * - <0: failure
//   //    */
//   //   virtual int getFilterProperty(const char* id, const char* key, char* json_value, size_t buf_size) { return -1; }

//   //  protected:
//   //   ~IVideoTrack() {}
//   // };

// #endregion C_IVideoTrack

/**
 * The statistics of the local video track.
 */
type LocalVideoTrackStats C.struct_C_LocalVideoTrackStats

//   struct C_LocalVideoTrackStats
//   {
//     /**
//      * The number of streams.
//      */
//     uint64_t number_of_streams;
//     /**
//      * The number of bytes of the major stream.
//      */
//     uint64_t bytes_major_stream;
//     /**
//      * The number of bytes of the minor stream.
//      */
//     uint64_t bytes_minor_stream;
//     /**
//      * The number of encoded frames.
//      */
//     uint32_t frames_encoded;
//     /**
//      * The SSRC (synchronization source) of the major stream.
//      */
//     uint32_t ssrc_major_stream;
//     /**
//      * The SSRC (synchronization source) of the minor stream.
//      */
//     uint32_t ssrc_minor_stream;
//     /**
//      * The capture frame rate of the video.
//      */
//     int capture_frame_rate;
//     /**
//      * The regulated frame rate of capture frame rate according to video encoder configuration.
//      */
//     int regulated_capture_frame_rate;
//     /**
//      * The input frame rate of the encoder.
//      */
//     int input_frame_rate;
//     /**
//      * The output frame rate of the encoder.
//      */
//     int encode_frame_rate;
//     /**
//      * The rendering frame rate.
//      */
//     int render_frame_rate;
//     /**
//      * The target bitrate (bps).
//      */
//     int target_media_bitrate_bps;
//     /**
//      * The frame rate excluding FEC.
//      */
//     int media_bitrate_bps;
//     /**
//      * The total frame rate including FEC.
//      */
//     int total_bitrate_bps; // Include FEC
//     /**
//      * The capture frame width (pixel).
//      */
//     int capture_width;
//     /**
//      * The capture frame height (pixel).
//      */
//     int capture_height;
//     /**
//      * The regulated frame width (pixel) of capture frame width according to video encoder configuration.
//      */
//     int regulated_capture_width;
//     /**
//      * The regulated frame height (pixel) of capture frame height according to video encoder configuration.
//      */
//     int regulated_capture_height;
//     /**
//      * The frame width (pixel).
//      */
//     int width;
//     /**
//      * The frame height (pixel).
//      */
//     int height;
//     uint32_t encoder_type;
//     uint32_t hw_encoder_accelerating;
//     /**
//      * The average time diff between frame captured and framed encoded.
//      */
//     uint32_t uplink_cost_time_ms;
//     /** Quality change of the local video in terms of target frame rate and
//      * target bit rate in this reported interval. See #QUALITY_ADAPT_INDICATION.
//      */
//     enum C_QUALITY_ADAPT_INDICATION quality_adapt_indication;
//     /**
//      * The video packet loss rate (%) from the local client to the Agora edge server before applying the anti-packet loss strategies.
//      */
//     unsigned short txPacketLossRate;

//     /** The brightness level of the video image captured by the local camera. See #CAPTURE_BRIGHTNESS_LEVEL_TYPE.
//      */
//     enum C_CAPTURE_BRIGHTNESS_LEVEL_TYPE capture_brightness_level;
//   };
//   struct C_LocalVideoTrackStats *C_LocalVideoTrackStats_New();
//   void C_LocalVideoTrackStats_Delete(struct C_LocalVideoTrackStats *this_);

/**
 * `ILocalVideoTrack` is the basic class for local video tracks, providing the main methods of local video tracks.
 * You can create a local video track by calling one of the following methods:
 * - `createCameraVideoTrack`
 * - `createScreenVideoTrack`
 * - `createMixedVideoTrack`
 * - `createCustomVideoTrack`
 * - `createMediaPlayerVideoTrack`
 *
 * After creating local video tracks, you can publish one or more local video tracks by calling \ref agora::rtc::ILocalUser::publishVideo "publishVideo".
 */
type ILocalVideoTrack C.C_ILocalVideoTrack

//   // class ILocalVideoTrack : public IVideoTrack {
//   //  public:
//   //   /**
//   //    * Enables or disables the local video track.
//   //    *
//   //    * Once the local video track is enabled, the SDK allows for local video capturing, processing, and encoding.
//   //    *
//   //    * @param enable Determines whether to enable the local video track.
//   //    * - `true`: Enable the local video track.
//   //    * - `false`: Disable the local video track.
//   //    */
//   //   virtual void setEnabled(bool enable) = 0;

//   //   /**
//   //    * Sets the video encoder configuration.
//   //    *
//   //    * Each video encoder configuration corresponds to a set of video parameters, including the
//   //    * resolution, frame rate, bitrate, and video orientation.
//   //    *
//   //    * The configurations specified in this method are the maximum values under ideal network conditions. If
//   //    * the video engine cannot render the video using the specified parameters due to poor network
//   //    * conditions, the configurations further down the list are considered until a successful configuration
//   //    * is found.
//   //    *
//   //    * @param config The reference to the video encoder configuration. See \ref agora::rtc::VideoEncoderConfiguration "VideoEncoderConfiguration".
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int setVideoEncoderConfiguration(const VideoEncoderConfiguration& config) = 0;

//   //   /**
//   //    * Set simulcast stream mode, enable, disable or auto enable
//   //    *
//   //    * @param mode Determines simulcast stream mode. See \ref agora::rtc::SIMULCAST_STREAM_MODE "SIMULCAST_STREAM_MODE".
//   //    * @param config The reference to the configurations for the simulcast stream mode. See \ref agora::rtc::SimulcastStreamConfig "SimulcastStreamConfig".
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int setSimulcastStreamMode(SIMULCAST_STREAM_MODE mode, const SimulcastStreamConfig& config) = 0;

//   //   /**
//   //    * Gets the state of the local video stream.
//   //    *
//   //    * @return The current state of the local video stream.
//   //    */
//   //   virtual LOCAL_VIDEO_STREAM_STATE getState() = 0;

//   //   /**
//   //    * Gets the statistics of the local video track.
//   //    *
//   //    * @param[out] stats The reference to the statistics of the local video track.
//   //    * @return
//   //    * - `true`: Success.
//   //    * - `false`: Failure.
//   //    */
//   //   virtual bool getStatistics(LocalVideoTrackStats& stats) = 0;

//   //   virtual VideoTrackType getType() OPTIONAL_OVERRIDE { return LOCAL_VIDEO_TRACK; }

//   //  protected:
//   //   ~ILocalVideoTrack() {}
//   // };

/**
 * The struct of RemoteVideoTrackStats.
 */
type RemoteVideoTrackStats C.struct_C_RemoteVideoTrackStats

//   struct C_RemoteVideoTrackStats
//   {
//     /**
//      The ID of the remote user.
//      */
//     C_uid_t uid;
//     /**
//      * The overall delay (ms) of the video frames.
//      */
//     int delay;
//     /**
//      * End-to-end delay from video capturer to video renderer. Hardware capture or render delay is excluded.
//      */
//     int e2eDelay;
//     /**
//      * The width (pixel) of the remote video track.
//      */
//     int width;
//     /**
//      * The height (pixel) of the remote video track.
//      */
//     int height;
//     /**
//      * The bitrate (Kbps) received in the reported interval.
//      */
//     int receivedBitrate;
//     /** The decoder output frame rate (fps) of the remote video track.
//      */
//     int decoderOutputFrameRate;
//     /** The render output frame rate (fps) of the remote video track.
//      */
//     int rendererOutputFrameRate;
//     /** The video frame loss rate (%) of the remote video stream in the reported interval.
//      */
//     int frameLossRate;
//     /** The packet loss rate (%) of the remote video track after using the anti-packet-loss method.
//      */
//     int packetLossRate;
//     /**
//      * The remote video stream type: #VIDEO_STREAM_TYPE.
//      */
//     enum C_VIDEO_STREAM_TYPE rxStreamType;
//     /**
//      The total freeze time (ms) of the remote video track after the remote user joins the channel.
//      In a video session where the frame rate is set to no less than 5 fps, video freeze occurs when
//      the time interval between two adjacent renderable video frames is more than 500 ms.
//      */
//     int totalFrozenTime;
//     /**
//      The total video freeze time as a percentage (%) of the total time when the video is available.
//      */
//     int frozenRate;
//     /**
//      * The number of video bytes received.
//      */
//     uint32_t received_bytes;
//     /**
//      The total number of decoded video frames.
//      */
//     uint32_t totalDecodedFrames;
//     /**
//      The offset (ms) between audio and video stream. A positive value indicates the audio leads the
//      video, and a negative value indicates the audio lags the video.
//      */
//     int avSyncTimeMs;
//     /**
//      The average offset(ms) between receive first packet which composite the frame and  the frame
//      ready to render.
//      */
//     uint32_t downlink_process_time_ms;
//     /**
//      The average time cost in renderer.
//      */
//     uint32_t frame_render_delay_ms;
//     /**
//      The total time (ms) when the remote user neither stops sending the video
//      stream nor disables the video module after joining the channel.
//      */
//     uint64_t totalActiveTime;
//     /**
//      The total publish duration (ms) of the remote video stream.
//      */
//     uint64_t publishDuration;
//     /**
//      decoded frame vqa mos value after all filter.
//     */
//     int vqa_mos;
//     /**
//      vqa avg cost ms
//     */
//     int vqa_avg_cost_ms;
//   };
//   struct C_RemoteVideoTrackStats *C_RemoteVideoTrackStats_New();
//   void C_RemoteVideoTrackStats_Delete(struct C_RemoteVideoTrackStats *this_);

/**
 * The IRemoteVideoTrack class.
 */
type IRemoteVideoTrack C.C_IRemoteVideoTrack

// #region IVideoTrack

//   // class IRemoteVideoTrack : public IVideoTrack {
//   //  public:
//   //   /**
//   //    * Gets the statistics of the remote video track.
//   //    * @param[out] stats The reference to the statistics of the remote video track.
//   //    * @return
//   //    * - `true`: Success.
//   //    * - `false`: Failure.
//   //    */
//   //   virtual bool getStatistics(RemoteVideoTrackStats& stats) = 0;
//   //   /**
//   //    * Gets the state of the remote video track.
//   //    * @return The state of the remote video track.
//   //    */
//   //   virtual REMOTE_VIDEO_STATE getState() = 0;
//   //   /**
//   //    * Gets the information of the remote video track.
//   //    * @param[out] info The reference to the information of the remote video track.
//   //    * @return
//   //    * - `true`: Success.
//   //    * - `false`: Failure.
//   //    */
//   //   virtual bool getTrackInfo(VideoTrackInfo& info) = 0;
//   //   /**
//   //    * Registers an \ref agora::media::IVideoEncodedFrameObserver "IVideoEncodedFrameObserver" object.
//   //    *
//   //    * You need to implement the `IVideoEncodedFrameObserver` class in this method. Once you successfully register
//   //    * the encoded image receiver, the SDK triggers the \ref agora::rtc::IVideoEncodedFrameObserver::onEncodedVideoFrameReceived "onEncodedVideoFrameReceived" callback when it receives the
//   //    * encoded video image.
//   //    *
//   //    * @param encodedObserver The pointer to the `IVideoEncodedFrameObserver` object.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int registerVideoEncodedFrameObserver(agora::media::IVideoEncodedFrameObserver* encodedObserver) = 0;
//   //   /**
//   //    * Releases the \ref agora::media::IVideoEncodedFrameObserver "IVideoEncodedFrameObserver" object.
//   //    * @param encodedObserver The pointer to the `IVideoEncodedFrameObserver` object.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int unregisterVideoEncodedFrameObserver(agora::media::IVideoEncodedFrameObserver* encodedObserver) = 0;

//   //   /**
//   //    * Registers an \ref agora::rtc::IMediaPacketReceiver "IMediaPacketReceiver" object.
//   //    *
//   //    * You need to implement the `IMediaPacketReceiver` class in this method. Once you successfully register
//   //    * the media packet receiver, the SDK triggers the \ref agora::rtc::IMediaPacketReceiver::onMediaPacketReceived "onMediaPacketReceived" callback when it receives the
//   //    * video packet.
//   //    *
//   //    * @param videoReceiver The pointer to the `IMediaPacketReceiver` object.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int registerMediaPacketReceiver(IMediaPacketReceiver* videoReceiver) = 0;
//   //   /**
//   //    * Releases the \ref agora::rtc::IMediaPacketReceiver "IMediaPacketReceiver" object.
//   //    * @param videoReceiver The pointer to the `IMediaPacketReceiver` object.
//   //    * @return
//   //    * - 0: Success.
//   //    * - < 0: Failure.
//   //    */
//   //   virtual int unregisterMediaPacketReceiver(IMediaPacketReceiver* videoReceiver) = 0;

//   //   virtual VideoTrackType getType() OPTIONAL_OVERRIDE { return REMOTE_VIDEO_TRACK; }

//   //  protected:
//   //   ~IRemoteVideoTrack() {}
//   // };

// #endregion IVideoTrack

// #endregion agora::rtc

// #endregion agora

// #ifdef __cplusplus
// }
// #endif // __cplusplus

// #endif // C_NGI_AGORA_VIDEO_TRACK_H
