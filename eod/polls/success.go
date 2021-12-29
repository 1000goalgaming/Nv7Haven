package polls

import (
	"fmt"

	"github.com/Nv7-Github/Nv7Haven/eod/types"
)

func (b *Polls) handlePollSuccess(p types.Poll) {
	db, res := b.GetDB(p.Guild)
	if !res.Exists {
		return
	}

	controversial := db.Config.VoteCount != 0 && float32(p.Downvotes)/float32(db.Config.VoteCount) >= 0.3
	controversialTxt := ""
	if controversial {
		controversialTxt = " 🌩️"
	}

	switch p.Kind {
	case types.PollCombo:
		b.elemCreate(p.PollComboData.Result, p.PollComboData.Elems, p.Suggestor, controversialTxt, p.Guild)
	case types.PollSign:
		b.mark(p.Guild, p.PollSignData.Elem, p.PollSignData.NewNote, p.Suggestor, controversialTxt)
	case types.PollImage:
		b.image(p.Guild, p.PollImageData.Elem, p.PollImageData.NewImage, p.Suggestor, p.PollImageData.Changed, controversialTxt)
	case types.PollCategorize:
		els := p.PollCategorizeData.Elems
		for _, val := range els {
			b.Categorize(val, p.PollCategorizeData.Category, p.Guild)
		}
		if len(els) == 1 {
			name, _ := db.GetElement(els[0])
			b.dg.ChannelMessageSend(db.Config.NewsChannel, fmt.Sprintf("🗃️ Added **%s** to **%s** (By <@%s>)%s", name.Name, p.PollCategorizeData.Category, p.Suggestor, controversialTxt))
		} else {
			b.dg.ChannelMessageSend(db.Config.NewsChannel, fmt.Sprintf("🗃️ Added **%d elements** to **%s** (By <@%s>)%s", len(els), p.PollCategorizeData.Category, p.Suggestor, controversialTxt))
		}
	case types.PollUnCategorize:
		els := p.PollCategorizeData.Elems
		for _, val := range els {
			b.UnCategorize(val, p.PollCategorizeData.Category, p.Guild)
		}
		if len(els) == 1 {
			name, _ := db.GetElement(els[0])
			b.dg.ChannelMessageSend(db.Config.NewsChannel, fmt.Sprintf("🗃️ Removed **%s** from **%s** (By <@%s>)%s", name.Name, p.PollCategorizeData.Category, p.Suggestor, controversialTxt))
		} else {
			b.dg.ChannelMessageSend(db.Config.NewsChannel, fmt.Sprintf("🗃️ Removed **%d elements** from **%s** (By <@%s>)%s", len(els), p.PollCategorizeData.Category, p.Suggestor, controversialTxt))
		}
	case types.PollCatImage:
		b.catImage(p.Guild, p.PollCatImageData.Category, p.PollCatImageData.NewImage, p.Suggestor, p.PollCatImageData.Changed, controversialTxt)
	case types.PollColor:
		b.color(p.Guild, p.PollColorData.Element, p.PollColorData.Color, p.Suggestor, controversialTxt)
	case types.PollCatColor:
		b.catColor(p.Guild, p.PollCatColorData.Category, p.PollCatColorData.Color, p.Suggestor, controversialTxt)
	}
}
