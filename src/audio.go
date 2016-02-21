package liggetest

import (
	"github.com/lietu/ligge"
)

func init() {
	ligge.RegisterAudioCategory("sfx", ligge.FloatToVolume(0.7))
	ligge.RegisterAudioCategory("music", ligge.FloatToVolume(0.5))
	ligge.RegisterAudioCategory("ambient", ligge.FloatToVolume(0.6))
}
