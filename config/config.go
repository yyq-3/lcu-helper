package config

var Cfg *Config

type Config struct {
	SpeakHelper SpeakHelperConfig
}
type SpeakHelperConfig struct {
	Enable bool
	Speed  uint8
	Volume uint8
}

func Init() {
	// 读取ini
	Cfg = &Config{}
}
