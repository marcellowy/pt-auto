package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/marcellowy/go-common/gogf/vlog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type MediaInfoDataTrack struct {
	Type                           string `json:"@type,omitempty"`
	VideoCount                     string `json:"VideoCount,omitempty"`
	FileExtension                  string `json:"FileExtension,omitempty"`
	Format                         string `json:"Format,omitempty"`
	FormatCommercialIfAny          string `json:"Format_Commercial_IfAny,omitempty"`
	FormatProfile                  string `json:"Format_Profile,omitempty"`
	CodecID                        string `json:"CodecID,omitempty"`
	CodecIDCompatible              string `json:"CodecID_Compatible,omitempty"`
	FileSize                       string `json:"FileSize,omitempty"`
	Duration                       string `json:"Duration,omitempty"`
	OverallBitRate                 string `json:"OverallBitRate,omitempty"`
	FrameRate                      string `json:"FrameRate,omitempty"`
	FrameCount                     string `json:"FrameCount,omitempty"`
	StreamSize                     string `json:"StreamSize,omitempty"`
	HeaderSize                     string `json:"HeaderSize,omitempty"`
	DataSize                       string `json:"DataSize,omitempty"`
	FooterSize                     string `json:"FooterSize,omitempty"`
	IsStreamable                   string `json:"IsStreamable,omitempty"`
	EncodedDate                    string `json:"Encoded_Date,omitempty"`
	TaggedDate                     string `json:"Tagged_Date,omitempty"`
	FileCreatedDate                string `json:"File_Created_Date,omitempty"`
	FileCreatedDateLocal           string `json:"File_Created_Date_Local,omitempty"`
	FileModifiedDate               string `json:"File_Modified_Date,omitempty"`
	FileModifiedDateLocal          string `json:"File_Modified_Date_Local,omitempty"`
	StreamOrder                    string `json:"StreamOrder,omitempty"`
	ID                             string `json:"ID,omitempty"`
	FormatLevel                    string `json:"Format_Level,omitempty"`
	FormatSettingsCABAC            string `json:"Format_Settings_CABAC,omitempty"`
	FormatSettingsRefFrames        string `json:"Format_Settings_RefFrames,omitempty"`
	BitRate                        string `json:"BitRate,omitempty"`
	Width                          string `json:"Width,omitempty"`
	Height                         string `json:"Height,omitempty"`
	SampledWidth                   string `json:"Sampled_Width,omitempty"`
	SampledHeight                  string `json:"Sampled_Height,omitempty"`
	PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
	DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
	Rotation                       string `json:"Rotation,omitempty"`
	FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
	FrameRateModeOriginal          string `json:"FrameRate_Mode_Original,omitempty"`
	FrameRateNum                   string `json:"FrameRate_Num,omitempty"`
	FrameRateDen                   string `json:"FrameRate_Den,omitempty"`
	ColorSpace                     string `json:"ColorSpace,omitempty"`
	ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
	BitDepth                       string `json:"BitDepth,omitempty"`
	ScanType                       string `json:"ScanType,omitempty"`
	EncodedLibrary                 string `json:"Encoded_Library,omitempty"`
	EncodedLibraryName             string `json:"Encoded_Library_Name,omitempty"`
	EncodedLibraryVersion          string `json:"Encoded_Library_Version,omitempty"`
	EncodedLibrarySettings         string `json:"Encoded_Library_Settings,omitempty"`
	ColourDescriptionPresent       string `json:"colour_description_present,omitempty"`
	ColourDescriptionPresentSource string `json:"colour_description_present_Source,omitempty"`
	ColourRange                    string `json:"colour_range,omitempty"`
	ColourRangeSource              string `json:"colour_range_Source,omitempty"`
	ColourPrimaries                string `json:"colour_primaries,omitempty"`
	ColourPrimariesSource          string `json:"colour_primaries_Source,omitempty"`
	TransferCharacteristics        string `json:"transfer_characteristics,omitempty"`
	TransferCharacteristicsSource  string `json:"transfer_characteristics_Source,omitempty"`
	MatrixCoefficients             string `json:"matrix_coefficients,omitempty"`
	MatrixCoefficientsSource       string `json:"matrix_coefficients_Source,omitempty"`
	Extra                          struct {
		CodecConfigurationBox string `json:"CodecConfigurationBox,omitempty"`
	} `json:"extra,omitempty"`

	FormatSettingsSBR        string `json:"Format_Settings_SBR,omitempty"`
	FormatAdditionalFeatures string `json:"Format_AdditionalFeatures,omitempty"`
	BitRateMode              string `json:"BitRate_Mode,omitempty"`
	Channels                 string `json:"Channels,omitempty"`
	ChannelPositions         string `json:"ChannelPositions,omitempty"`
	ChannelLayout            string `json:"ChannelLayout,omitempty"`
	SamplesPerFrame          string `json:"SamplesPerFrame,omitempty"`
	SamplingRate             string `json:"SamplingRate,omitempty"`
	SamplingCount            string `json:"SamplingCount,omitempty"`
	CompressionMode          string `json:"Compression_Mode,omitempty"`
	Default                  string `json:"Default,omitempty"`
	AlternateGroup           string `json:"AlternateGroup,omitempty"`
}

type MediaInfoData struct {
	CreatingLibrary struct {
		Name    string `json:"name,omitempty"`
		Version string `json:"version,omitempty"`
		URL     string `json:"url,omitempty"`
	} `json:"creatingLibrary,omitempty"`
	Media struct {
		Ref   string               `json:"@ref,omitempty"`
		Track []MediaInfoDataTrack `json:"track"`
	} `json:"media"`
}

type MediaInfo struct {
	Definition string // 1080i/1080p/2160p
	//Code             string // H264/H265
	VideoCodec           string
	AudioCodec           string
	MediaInfoContent     string         // 用于上传
	MediaInfoJSONContent string         // JSON格式,用于解析
	mid                  *MediaInfoData // JSON翻译后的
	mi                   []*mediaInfos
}

type MediaScanType string

var DefinitionStringMapId = map[string]string{
	DefinitionUHD: DefinitionUHDId,
	DefinitionHDP: DefinitionHDPId,
	DefinitionHDI: DefinitionHDIId,
	DefinitionHD:  DefinitionHDId,
	DefinitionSD:  DefinitionSDId,
	Definition8k:  Definition8kId,
}

var DefinitionStringMapTitle = map[string]string{
	DefinitionUHD: DefinitionUHDTitle,
	DefinitionHDP: DefinitionHDP,
	DefinitionHDI: DefinitionHDI,
	DefinitionHD:  DefinitionHD,
	DefinitionSD:  DefinitionTitle,
	Definition8k:  Definition8k,
}

var VideoCodecStringMapId = map[string]string{
	VideoCodecH264:  VideoCodecH264Id,
	VideoCodecH265:  VideoCodecH265Id,
	VideoCodecVC1:   VideoCodecVC1Id,
	VideoCodecMPEG2: VideoCodecMPEG2Id,
	VideoCodecXvid:  VideoCodecXvidId,
	VideoCodecAV1:   VideoCodecAV1Id,
	VideoCodecVP8:   VideoCodecVP8Id,
	VideoCodecAVS:   VideoCodecAVSId,
}

var VideoCodecStringMapTitle = map[string]string{
	VideoCodecH264: VideoCodecH264Title,
	VideoCodecH265: VideoCodecH265Title,
	VideoCodecVC1:  VideoCodecVC1Title,
	//VideoCodecMPEG2: VideoCodecMPEG2Id,
	//VideoCodecXvid:  VideoCodecXvidId,
	//VideoCodecAV1:   VideoCodecAV1Id,
	//VideoCodecVP8:   VideoCodecVP8Id,
	VideoCodecAVS: VideoCodecAVSTitle,
}

var MediaAudioCodecStringMapId = map[string]string{
	MediaAudioCodecAAC:         MediaAudioCodecAACId,
	MediaAudioCodecAC3DD:       MediaAudioCodecAC3DDId,
	MediaAudioCodecDTS:         MediaAudioCodecDTSId,
	MediaAudioCodecDTSHDMA:     MediaAudioCodecDTSHDMAId,
	MediaAudioCodecEAC3DDP:     MediaAudioCodecEAC3DDPId,
	MediaAudioCodecEAC3Atoms:   MediaAudioCodecEAC3AtomsId,
	MediaAudioCodecTrueHD:      MediaAudioCodecTrueHDId,
	MediaAudioCodecTrueHDAtoms: MediaAudioCodecTrueHDAtomsId,
	MediaAudioCodecLPCM:        MediaAudioCodecLPCMId,
	MediaAudioCodecWAV:         MediaAudioCodecWAVId,
	MediaAudioCodecFlac:        MediaAudioCodecFlacId,
	MediaAudioCodecAPE:         MediaAudioCodecAPEId,
	MediaAudioCodecMp2:         MediaAudioCodecMp2Id,
	MediaAudioCodecOgg:         MediaAudioCodecOggId,
	MediaAudioCodecOther:       MediaAudioCodecOtherId,
}

const (
	MediaScanTypeProgressive MediaScanType = "Progressive"
	MediaScanTypeInterlaced  MediaScanType = "Interlaced"
	MediaScanTypeMBAFF       MediaScanType = "MBAFF" // 1080p

	SelectionGenernal = "General"
	SelectionVideo    = "Video"
	SelectionAudio    = "Audio"

	DefinitionUHD      = "4k/2160p"
	DefinitionUHDTitle = "2160p"
	DefinitionUHDId    = "6"
	DefinitionHDP      = "1080p"
	DefinitionHDPId    = "1"
	DefinitionHDI      = "1080i"
	DefinitionHDIId    = "2"
	DefinitionHD       = "720p"
	DefinitionHDId     = "3"
	DefinitionSD       = "SD" // 480p
	DefinitionTitle    = "720p"
	DefinitionSDId     = "4"
	Definition8k       = "8k"
	Definition8kId     = "7"

	VideoCodecH264      = "H.264(x264/AVC)"
	VideoCodecH264Title = "H264"
	VideoCodecH264Id    = "1"
	VideoCodecH265      = "H.265(x265/HEVC)"
	VideoCodecH265Title = "H265"
	VideoCodecH265Id    = "16"
	VideoCodecVC1       = "VC-1"
	VideoCodecVC1Title  = "VC-1"
	VideoCodecVC1Id     = "2"
	VideoCodecMPEG2     = "MPEG-2"
	VideoCodecMPEG2Id   = "4"
	VideoCodecXvid      = "Xvid"
	VideoCodecXvidId    = "3"
	VideoCodecAV1       = "AV1"
	VideoCodecAV1Id     = "19"
	VideoCodecVP8       = "VP8/9"
	VideoCodecVP8Id     = "21"
	VideoCodecAVS       = "AVS"
	VideoCodecAVSTitle  = "AVS"
	VideoCodecAVSId     = "22"

	MediaAudioCodecAAC           = "AAC"
	MediaAudioCodecAACId         = "6"
	MediaAudioCodecAC3DD         = "AC3(DD)"
	MediaAudioCodecAC3DDId       = "8"
	MediaAudioCodecDTS           = "DTS"
	MediaAudioCodecDTSId         = "3"
	MediaAudioCodecDTSHDMA       = "DTS-HD MA"
	MediaAudioCodecDTSHDMAId     = "11"
	MediaAudioCodecEAC3DDP       = "E-AC3(DDP)"
	MediaAudioCodecEAC3DDPId     = "12"
	MediaAudioCodecEAC3Atoms     = "E-AC3 Atoms(DDP Atoms)"
	MediaAudioCodecEAC3AtomsId   = "13"
	MediaAudioCodecTrueHD        = "TrueHD"
	MediaAudioCodecTrueHDId      = "9"
	MediaAudioCodecTrueHDAtoms   = "TrueHD Atmos"
	MediaAudioCodecTrueHDAtomsId = "10"
	MediaAudioCodecLPCM          = "LPCM/PCM"
	MediaAudioCodecLPCMId        = "14"
	MediaAudioCodecWAV           = "WAV"
	MediaAudioCodecWAVId         = "15"
	MediaAudioCodecFlac          = "FLAC"
	MediaAudioCodecFlacId        = "1"
	MediaAudioCodecAPE           = "APE"
	MediaAudioCodecAPEId         = "2"
	MediaAudioCodecMp2           = "MP2/3"
	MediaAudioCodecMp2Id         = "4"
	MediaAudioCodecOgg           = "OGG"
	MediaAudioCodecOggId         = "5"
	MediaAudioCodecOther         = "Other"
	MediaAudioCodecOtherId       = "7"
)

type mediaInfos struct {
	Selection string             `json:"selection"`
	Options   []*mediaInfoOption `json:"options"`
}

type mediaInfoOption struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (*MediaInfo) getTmpDir() string {
	tmp, _ := filepath.Abs("tmp")
	_ = os.MkdirAll(tmp, os.ModePerm)
	return tmp
}

func (m *MediaInfo) getValue(ctx context.Context, selection, key string) string {
	for _, v := range m.mi {
		if v.Selection == selection {
			for _, opt := range v.Options {
				if opt.Key == key {
					return opt.Value.(string)
				}
			}
		}
	}
	return ""
}

func (m *MediaInfo) getValueV2(ctx context.Context, typ string) MediaInfoDataTrack {
	if m.mid == nil {
		return MediaInfoDataTrack{}
	}
	for _, track := range m.mid.Media.Track {
		if track.Type == typ {
			return track
		}
	}
	return MediaInfoDataTrack{}
}

func (m *MediaInfo) parseInfo(ctx context.Context, info string) ([]*mediaInfos, error) {
	if info == "" {
		return nil, fmt.Errorf("empty info")
	}

	infoArr := strings.Split(info, "\n")
	if len(infoArr) == 1 {
		return nil, fmt.Errorf("empty info")
	}
	var mis []*mediaInfos
	var mi *mediaInfos = nil
	for _, line := range infoArr {
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.ReplaceAll(line, "\r", "")

		// 假如是空行
		ifSpaceLine := strings.TrimSpace(line)
		if ifSpaceLine == "" {
			continue
		}

		if !strings.Contains(line, ":") {
			if mi != nil {
				mis = append(mis, mi)
			}
			mi = &mediaInfos{}
			mi.Selection = line
			continue
		}

		var options = strings.Split(line, ":")
		options[0] = strings.TrimSpace(options[0])
		//Complete name                            : E:\study\Auto\pt-auto\tools\永夜星河.mp4
		//Recorded date                            : 2024-11-07 19:08:50.7656294+08:00
		//Display aspect ratio                     : 2.39:1
		var newOptionsValue []string
		newOptionsValue = append(newOptionsValue, options[1:]...)
		if mi != nil {
			mi.Options = append(mi.Options, &mediaInfoOption{
				Key:   options[0],
				Value: strings.TrimSpace(strings.Join(newOptionsValue, ":")),
			})
		}
	}
	// last one
	if mi != nil {
		mis = append(mis, mi)
	}
	return mis, nil
}

func (m *MediaInfo) Open(ctx context.Context, video string) (err error) {
	m.Definition = DefinitionHDP  // 1080p
	m.VideoCodec = VideoCodecH264 // x264
	dir := m.getTmpDir()          // 临时目录

	// 获取常规文字版本
	{
		file := filepath.Join(dir, "mediaInfo.txt")
		shell := fmt.Sprintf("%s --LogFile=\"%s\" \"%s\"", "MediaInfo", file, video)
		var output string
		if output, err = Exec(ctx, shell); err != nil {
			vlog.Info(ctx, video)
			vlog.Info(ctx, dir)
			vlog.Info(ctx, output)
			vlog.Error(ctx, err)
			return err
		}

		var b []byte
		if b, err = os.ReadFile(file); err != nil {
			vlog.Error(ctx, err)
			return err
		}
		m.MediaInfoContent = string(b)
		if m.mi, err = m.parseInfo(ctx, string(b)); err != nil {
			vlog.Error(ctx, err)
			return err
		}
		defer func() {
			_ = os.Remove(file)
		}()
	}

	// 获取json版本
	{
		file := filepath.Join(dir, "mediaInfo_json.txt")
		shell := fmt.Sprintf("%s --LogFile=\"%s\" --Output=JSON \"%s\"", "MediaInfo", file, video)
		var output string
		if output, err = Exec(ctx, shell); err != nil {
			vlog.Info(ctx, video)
			vlog.Info(ctx, dir)
			vlog.Info(ctx, output)
			vlog.Error(ctx, err)
			return err
		}

		var b []byte
		if b, err = os.ReadFile(file); err != nil {
			vlog.Error(ctx, err)
			return err
		}
		m.mid = &MediaInfoData{}
		if err = json.Unmarshal(b, &m.mid); err != nil {
			vlog.Error(ctx, err)
			return
		}
		defer func() {
			_ = os.Remove(file)
		}()
	}

	// 如果height == 2160那直接就是2160p
	videoHeight := m.getValueV2(ctx, SelectionVideo).Height
	{
		//var re = regexp.MustCompile(`[0-9]{1,}`)
		//v := m.getValue(ctx, SelectionVideo, "Height")
		//height := re.FindString(videoHeight)
		var heightInt int64
		if heightInt, err = strconv.ParseInt(videoHeight, 10, 64); err != nil {
			vlog.Error(ctx, err)
			return err
		}
		if heightInt == 4320 {
			m.Definition = Definition8k
		} else if heightInt == 2160 {
			m.Definition = DefinitionUHD
		} else if heightInt > 720 && heightInt <= 1080 {
			scanType := m.getValueV2(ctx, SelectionVideo).ScanType
			switch scanType {
			case string(MediaScanTypeInterlaced):
				m.Definition = DefinitionHDI // 1080i
			case string(MediaScanTypeMBAFF), string(MediaScanTypeProgressive):
				m.Definition = DefinitionHDP // 1080p
			}
		} else if heightInt == 720 {
			m.Definition = DefinitionHD
		} else if heightInt == 480 {
			m.Definition = DefinitionSD
		}
	}

	// 处理视频编码
	{
		//format := m.getValue(ctx, SelectionVideo, "Format")
		format := m.getValueV2(ctx, SelectionVideo).Format
		switch format {
		case "AVC":
			m.VideoCodec = VideoCodecH264
		case "MPEGVideo":
			m.VideoCodec = VideoCodecMPEG2
		case "HEVC":
			m.VideoCodec = VideoCodecH265
		}
	}

	// 处理音频编码
	{
		// OGG ， MP3 ， WAV ， REALAUDIO ， AC3 ， DTS ， AAC ， M4A ， AU ，AIF， AIFF ， OPUS
		// format := m.getValue(ctx, SelectionAudio, "Format")
		format := m.getValueV2(ctx, SelectionAudio).Format
		commercialName := m.getValueV2(ctx, SelectionAudio).FormatCommercialIfAny
		// formatProfile := m.getValueV2(ctx, SelectionAudio).FormatProfile
		switch format {
		case "AAC LC", "AAC":
			m.AudioCodec = MediaAudioCodecAAC
		case "AC-3":
			switch commercialName {
			case "Dolby Digital":
				m.AudioCodec = MediaAudioCodecAC3DD
			}
		case "OGG":
			m.AudioCodec = MediaAudioCodecOgg
		case "MPEG Audio":
			m.AudioCodec = MediaAudioCodecMp2
		case "E-AC-3":
			switch commercialName {
			case "Dolby Digital Plus":
				m.AudioCodec = MediaAudioCodecEAC3DDP
			default:
				m.AudioCodec = MediaAudioCodecEAC3Atoms
			}
		case "DTS":
			switch commercialName {
			case "DTS-HD Master Audio":
				m.AudioCodec = MediaAudioCodecDTSHDMA
			default:
				m.AudioCodec = MediaAudioCodecDTS
			}

		}
	}

	return nil
}

// IsUHD 3840 * 2160
func (m *MediaInfo) IsUHD() bool {
	return m.Definition == DefinitionUHD
}
