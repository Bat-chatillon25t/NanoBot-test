// Code generated by codegen/postopenapiof. DO NOT EDIT.

package nano

import (
	"io"

	"github.com/pkg/errors"
)

func (bot *Bot) postOpenAPIofChannel(ep, contenttype string, body io.Reader) (*Channel, error) {
	resp := &struct {
		CodeMessageBase
		Channel
	}{}
	err := bot.PostOpenAPI(ep, contenttype, resp, body)
	if err != nil {
		err = errors.Wrap(err, getCallerFuncName())
	}
	return &resp.Channel, err
}

func (bot *Bot) postOpenAPIofGuildRoleCreate(ep, contenttype string, body io.Reader) (*GuildRoleCreate, error) {
	resp := &struct {
		CodeMessageBase
		GuildRoleCreate
	}{}
	err := bot.PostOpenAPI(ep, contenttype, resp, body)
	if err != nil {
		err = errors.Wrap(err, getCallerFuncName())
	}
	return &resp.GuildRoleCreate, err
}

func (bot *Bot) postOpenAPIofMessage(ep, contenttype string, body io.Reader) (*Message, error) {
	resp := &struct {
		CodeMessageBase
		Message
	}{}
	err := bot.PostOpenAPI(ep, contenttype, resp, body)
	if err != nil {
		err = errors.Wrap(err, getCallerFuncName())
	}
	return &resp.Message, err
}
