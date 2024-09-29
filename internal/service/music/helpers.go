package music

import (
	"strings"

	"github.com/meraiku/music_lib/internal/model"
)

func paginateText(text string) []model.Text {
	formatedText := strings.Split(text, "\n\n")

	out := make([]model.Text, len(formatedText))

	for i := range formatedText {
		out[i].Verse = formatedText[i]
		out[i].VerseNumber = i + 1
	}

	return out
}
