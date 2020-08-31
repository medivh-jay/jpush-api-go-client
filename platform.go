package jpushclient

import (
	"errors"
)

const (
	IOS      = "ios"
	ANDROID  = "android"
	WINPHONE = "winphone"
)

type Platform struct {
	Os        interface{}
	Platforms []string
}

func (platform *Platform) All() {
	platform.Os = "all"
}

func (platform *Platform) Add(os string) error {
	if platform.Os == nil {
		platform.Platforms = make([]string, 0, 3)
	} else {
		switch platform.Os.(type) {
		case string:
			return errors.New("platform is all")
		default:
		}
	}

	//判断是否重复
	for _, value := range platform.Platforms {
		if os == value {
			return nil
		}
	}

	switch os {
	case IOS:
		fallthrough
	case ANDROID:
		fallthrough
	case WINPHONE:
		platform.Platforms = append(platform.Platforms, os)
		platform.Os = platform.Platforms
	default:
		return errors.New("platform error")
	}

	return nil
}

func (platform *Platform) AddIOS() {
	_ = platform.Add(IOS)
}

func (platform *Platform) AddAndroid() {
	_ = platform.Add(ANDROID)
}

func (platform *Platform) AddWinphone() {
	_ = platform.Add(WINPHONE)
}
