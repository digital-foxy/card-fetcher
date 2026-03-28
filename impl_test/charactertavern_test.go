package impl_test

import (
	"testing"

	"github.com/digital-foxy/card-fetcher/source"
	"github.com/digital-foxy/card-parser/character"
)

func TestCharacterTavernImport_CCV3(t *testing.T) {
	t.Parallel()
	FetchAndAssert(t, "https://character-tavern.com/character/tidyup/beth_homeless_on_her_birthday_").
		AssertNoErr().
		Source(source.CharacterTavern).
		SheetRevision(character.RevisionV3).
		IsForked(false).
		Consistent().
		AssertImage()
}

func TestCharacterTavernImportFail(t *testing.T) {
	FetchAndAssert(t, "character-tavern.com/character/brian007/lara_s").
		AssertErr()
}

func TestCharacterTavernImportNotes(t *testing.T) {
	t.Parallel()
	FetchAndAssert(t, "https://character-tavern.com/character/animatedspell/Violete%20V4").
		AssertNoErr().
		Source(source.CharacterTavern).
		SheetCreator("animatedspell").
		SheetCreatorNotes("Your shapeshifting monster waifu—sweet as sin, twice as wicked. A devoted lover by day, an Encyclopedia-inspired hentai experiment by night.\n\nViolete is a bot with lots of kinky tag, if you don't like any of her suggestion during chat, just gently asked her she will comply (tentacle, Insect,monster, monster girl, etc)").
		IsForked(false).
		Consistent().
		AssertImage()
}

func TestCharacterTavernPngData(t *testing.T) {
	FetchAndAssert(t, "https://character-tavern.com/character/wicked_ali/Veronica").
		AssertNoErr().
		Source(source.CharacterTavern).
		IsForked(false).
		Consistent().
		AssertImage()
}
