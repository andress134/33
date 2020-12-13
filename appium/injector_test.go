package appium

import "github.com/andress134/agouti"

func NewTestDevice(session mobileSession) *Device {
	return &Device{
		Page:    &agouti.Page{},
		session: session,
	}
}
