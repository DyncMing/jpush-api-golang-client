package jpush

import "errors"

type PlatformType string

const (
	IOS      PlatformType = "ios"
	ANDROID  PlatformType = "android"
	WINPHONE PlatformType = "winphone"
	QUICKAPP PlatformType = "quickapp"
	HMOS     PlatformType = "hmos"
)

type Platform struct {
	Os      interface{}
	osArray []string
}

func (p *Platform) Interface() interface{} {
	switch p.Os.(type) {
	case string:
		return p.Os
	default:
	}
	return p.osArray
}

// All set all platforms
func (p *Platform) All() {
	p.Os = "all"
}

// Add platform
func (p *Platform) Add(os PlatformType) error {
	if p.osArray == nil {
		p.osArray = make([]string, 0)
	}

	switch p.Os.(type) {
	case string:
		return errors.New("platform already set all")
	default:
	}

	// check if already added
	for _, v := range p.osArray {
		if v == string(os) {
			return nil
		}
	}

	switch os {
	case IOS:
		fallthrough
	case ANDROID:
		fallthrough
	case WINPHONE:
		fallthrough
	case QUICKAPP:
		fallthrough
	case HMOS:
		p.osArray = append(p.osArray, string(os))
		p.Os = p.osArray
	default:
		return errors.New("invalid platform")
	}

	return nil
}

// AddIOS add ios platform
func (p *Platform) AddIOS() {
	_ = p.Add(IOS)
}

// AddAndroid add android platform
func (p *Platform) AddAndroid() {
	_ = p.Add(ANDROID)
}

// AddWinphone add winphone platform
func (p *Platform) AddWinphone() {
	_ = p.Add(WINPHONE)
}

// AddQuickApp add quickapp platform
func (p *Platform) AddQuickApp() {
	_ = p.Add(QUICKAPP)
}

// AddHmos add hmos (HarmonyOS) platform
func (p *Platform) AddHmos() {
	_ = p.Add(HMOS)
}

// Remove remove platform
func (p *Platform) Remove(os PlatformType) error {
	if p.osArray == nil {
		return errors.New("platform not set")
	}

	for i, v := range p.osArray {
		if v == string(os) {
			p.osArray = append(p.osArray[:i], p.osArray[i+1:]...)
			if len(p.osArray) == 0 {
				p.Os = nil
			} else {
				p.Os = p.osArray
			}
			return nil
		}
	}

	return errors.New("platform not found")
}
