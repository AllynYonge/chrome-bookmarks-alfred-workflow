package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfile_AvatarIconURL(t *testing.T) {
	assert := assert.New(t)
	browser := Browser{Path: "testdata", IconURL: "img/chrome.png"}
	profile := Profile{
		AvatarURL:       "chrome://theme/IDR_PROFILE_AVATAR_12",
		CustomAvatarURL: "avatar.png",
		IsDefaultAvatar: false,
	}

	assert.Equal("img/chrome.png", profile.AvatarIconURL(browser, "Default"))

	profile.IsDefaultAvatar = true

	assert.Equal("testdata/Avatars/avatar_ninja.png", profile.AvatarIconURL(browser, "Default"))
}
