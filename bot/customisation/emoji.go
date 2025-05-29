package customisation

import (
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects"
	"github.com/TicketsBot-cloud/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1348678458078920765, false)
	EmojiOpen       = NewCustomEmoji("open", 1348678391704064000, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1348678499858382998, false)
	EmojiClose      = NewCustomEmoji("close", 1348678478379221042, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1348691867906211970, false)
	EmojiReason     = NewCustomEmoji("reason", 1348678534620778516, false)
	EmojiSubject    = NewCustomEmoji("subject", 1348691896964354098, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1348691929960808488, false)
	EmojiClaim      = NewCustomEmoji("claim", 1348678516044075149, false)
	EmojiPanel      = NewCustomEmoji("panel", 1348691957467189309, false)
	EmojiRating     = NewCustomEmoji("rating", 1348691973497688095, false)
	EmojiStaff      = NewCustomEmoji("staff", 1348691993273962577, false)
	EmojiThread     = NewCustomEmoji("thread", 1348692011829559347, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1348692032167608392, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1348692054179577888, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1348692073792012409, false)
	//EmojiTime       = NewCustomEmoji("time", 1348692096428544002, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
