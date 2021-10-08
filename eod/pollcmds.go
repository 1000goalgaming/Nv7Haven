package eod

import (
	"fmt"
	"strings"

	"github.com/Nv7-Github/Nv7Haven/eod/types"
)

func (b *EoD) markCmd(elem string, mark string, m types.Msg, rsp types.Rsp) {
	lock.RLock()
	dat, exists := b.dat[m.GuildID]
	lock.RUnlock()
	if !exists {
		return
	}

	rsp.Acknowledge()

	el, res := dat.GetElement(elem)
	if !res.Exists {
		rsp.ErrorMessage(fmt.Sprintf("Element **%s** doesn't exist!", elem))
		return
	}

	inv, res := dat.GetInv(m.Author.ID, true)
	if !res.Exists {
		rsp.ErrorMessage(res.Message)
		return
	}
	_, exists = inv[strings.ToLower(el.Name)]
	if !exists {
		rsp.ErrorMessage(fmt.Sprintf("Element **%s** is not in your inventory!", el.Name))
		return
	}
	if len(mark) >= 2400 {
		rsp.ErrorMessage("Creator marks must be under 2400 characters!")
		return
	}

	if el.Creator == m.Author.ID {
		b.mark(m.GuildID, elem, mark, "", "")
		rsp.Message(fmt.Sprintf("You have signed **%s**! 🖋️", el.Name))
		return
	}

	err := b.createPoll(types.Poll{
		Channel: dat.VotingChannel,
		Guild:   m.GuildID,
		Kind:    types.PollSign,
		Value1:  el.Name,
		Value2:  mark,
		Value3:  el.Comment,
		Value4:  m.Author.ID,
	})
	if rsp.Error(err) {
		return
	}
	id := rsp.Message(fmt.Sprintf("Suggested a note for **%s** 🖊️", el.Name))
	dat.SetMsgElem(id, el.Name)

	lock.Lock()
	b.dat[m.GuildID] = dat
	lock.Unlock()
}

func (b *EoD) imageCmd(elem string, image string, m types.Msg, rsp types.Rsp) {
	lock.RLock()
	dat, exists := b.dat[m.GuildID]
	lock.RUnlock()
	if !exists {
		return
	}

	rsp.Acknowledge()

	el, res := dat.GetElement(elem)
	if !res.Exists {
		rsp.ErrorMessage(fmt.Sprintf("Element **%s** doesn't exist!", elem))
		return
	}

	inv, res := dat.GetInv(m.Author.ID, true)
	if !res.Exists {
		rsp.ErrorMessage(res.Message)
		return
	}
	_, exists = inv[strings.ToLower(el.Name)]
	if !exists {
		rsp.ErrorMessage(fmt.Sprintf("Element **%s** is not in your inventory!", el.Name))
		return
	}

	if el.Creator == m.Author.ID {
		b.image(m.GuildID, elem, image, "", "")
		rsp.Message(fmt.Sprintf("You added an image to **%s**! 📷", el.Name))
		return
	}

	err := b.createPoll(types.Poll{
		Channel: dat.VotingChannel,
		Guild:   m.GuildID,
		Kind:    types.PollImage,
		Value1:  el.Name,
		Value2:  image,
		Value3:  el.Image,
		Value4:  m.Author.ID,
	})
	if rsp.Error(err) {
		return
	}
	id := rsp.Message(fmt.Sprintf("Suggested an image for **%s** 📷", el.Name))
	dat.SetMsgElem(id, el.Name)

	lock.Lock()
	b.dat[m.GuildID] = dat
	lock.Unlock()
}

func (b *EoD) colorCmd(elem string, color int, m types.Msg, rsp types.Rsp) {
	lock.RLock()
	dat, exists := b.dat[m.GuildID]
	lock.RUnlock()
	if !exists {
		return
	}

	rsp.Acknowledge()

	el, res := dat.GetElement(elem)
	if !res.Exists {
		rsp.ErrorMessage(fmt.Sprintf("Element **%s** doesn't exist!", elem))
		return
	}

	inv, res := dat.GetInv(m.Author.ID, true)
	if !res.Exists {
		rsp.ErrorMessage(res.Message)
		return
	}
	_, exists = inv[strings.ToLower(el.Name)]
	if !exists {
		rsp.ErrorMessage(fmt.Sprintf("Element **%s** is not in your inventory!", el.Name))
		return
	}

	if el.Creator == m.Author.ID {
		b.color(m.GuildID, elem, color, "", "")
		rsp.Message(fmt.Sprintf("You have set the color of **%s**! 🖌️", el.Name))
		return
	}

	err := b.createPoll(types.Poll{
		Channel: dat.VotingChannel,
		Guild:   m.GuildID,
		Kind:    types.PollColor,
		Value1:  el.Name,
		Value3:  el.Image,
		Value4:  m.Author.ID,
		Data:    map[string]interface{}{"color": color},
	})
	if rsp.Error(err) {
		return
	}
	id := rsp.Message(fmt.Sprintf("Suggested a color for **%s** 🖌️", el.Name))
	dat.SetMsgElem(id, el.Name)

	lock.Lock()
	b.dat[m.GuildID] = dat
	lock.Unlock()
}

func (b *EoD) catImgCmd(catName string, url string, m types.Msg, rsp types.Rsp) {
	lock.RLock()
	dat, exists := b.dat[m.GuildID]
	lock.RUnlock()
	if !exists {
		return
	}

	cat, res := dat.GetCategory(catName)
	if !res.Exists {
		rsp.ErrorMessage(res.Message)
		return
	}

	err := b.createPoll(types.Poll{
		Channel: dat.VotingChannel,
		Guild:   m.GuildID,
		Kind:    types.PollCatImage,
		Value1:  cat.Name,
		Value2:  url,
		Value3:  cat.Image,
		Value4:  m.Author.ID,
	})
	if rsp.Error(err) {
		return
	}
	rsp.Message(fmt.Sprintf("Suggested an image for category **%s** 📷", cat.Name))
}

func (b *EoD) catColorCmd(catName string, color int, m types.Msg, rsp types.Rsp) {
	lock.RLock()
	dat, exists := b.dat[m.GuildID]
	lock.RUnlock()
	if !exists {
		return
	}

	cat, res := dat.GetCategory(catName)
	if !res.Exists {
		rsp.ErrorMessage(res.Message)
		return
	}

	err := b.createPoll(types.Poll{
		Channel: dat.VotingChannel,
		Guild:   m.GuildID,
		Kind:    types.PollCatColor,
		Value1:  cat.Name,
		Value4:  m.Author.ID,
		Data:    map[string]interface{}{"color": color},
	})
	if rsp.Error(err) {
		return
	}
	rsp.Message(fmt.Sprintf("Suggested a color for category **%s** 🖌️", cat.Name))
}
