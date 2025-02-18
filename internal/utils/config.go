package utils

type ConfigRapidApi struct {
	Host string
	Key  string
}

type ConfigMTeam struct {
	URL             string
	ApiKey          string
	UploadImgApiKey string
}

type ConfigProxy struct {
	Enabled bool
	Addr    string
}

type Config struct {
	MacPicNum                 int    // 截图张数
	TorrentMapVideoDirEnabled bool   // 处理后视频存放目录功能开关
	TorrentMapVideoDir        string // 处理后视频存放目录
}

func (*Config) Read() bool {
	return false
}
