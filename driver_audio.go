// Copyright 2017 Eric Zhou. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base64Captcha

//DriverAudio captcha config for captcha-engine-audio.
type DriverAudio struct {
	// Length Default number of digits in captcha solution.
	Length int
	// Language possible values for lang are "en", "ja", "ru", "zh".
	Language string
}

//DefaultDriverAudio is a default audio driver
var DefaultDriverAudio = NewDriverAudio(6, "en")

//NewDriverAudio creates a driver of audio
func NewDriverAudio(length int, language string) *DriverAudio {
	return &DriverAudio{Length: length, Language: language}
}

//GenerateItem creates audio captcha item
func (d *DriverAudio) GenerateItem(content string) (item Item, err error) {
	digits := stringToFakeByte(content)
	audio := newAudio("", digits, d.Language)
	return audio, nil
}

//GenerateQuestionAnswer creates captcha content and answer
func (d *DriverAudio) GenerateQuestionAnswer() (q, a string) {
	digits := randomDigits(d.Length)
	a = parseDigitsToString(digits)
	return a, a
}
