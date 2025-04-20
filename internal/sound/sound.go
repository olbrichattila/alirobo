// Sound play sounds
package sound

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const sampleRate = 44100

var audioContext *audio.Context = audio.NewContext(sampleRate)

// Play audio context, only one sound can be played
func Play(s *audio.Player) {
	if !s.IsPlaying() {
		s.Rewind()
		s.Play()
	}
}

// PlayNewFromData plays audio form byte raw data, multiple times can be played, sounds overlapping
func PlayNewFromData(soundData []byte) {
	audioStream, _ := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(soundData))
	sound, _ := audioContext.NewPlayer(audioStream)

	if !sound.IsPlaying() {
		sound.Rewind()
		sound.Play()
	}
}

func LoadMp3SoundData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Println("Cannot load " + url)
		return nil, fmt.Errorf("cannot load %s", url)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return bodyBytes, nil
}

func LoadMp3Sound(path string) (*audio.Player, error) {
	d, err := LoadMp3SoundData(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	audioStream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(d))

	return audioContext.NewPlayer(audioStream)
}
