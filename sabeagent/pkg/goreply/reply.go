package goreply

import (
	"context"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const AllSender = ""

type (
	Params struct {
		Sender  string `json:"sender"`
		Message string `json:"message"`
	}

	ReplyCache struct {
		Msg     string   `json:"msg"`
		Replies []string `json:"replies"`
	}
	SyncReply struct {
		Sender  string   `json:"sender"`
		Msg     string   `json:"msg"`
		Replies []string `json:"replies"`
	}
)

var (
	rwMutex sync.RWMutex

	// repliesCache[sender][msg]ReplyCache
	repliesCache map[string]map[string]ReplyCache

	ErrNotMatch = errors.New("not match")
)

// Reply 简单回复
func Reply(params Params) (string, error) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	custom, ok := repliesCache[params.Sender]
	if ok {
		resp, err := rangeReply(custom, params, func(s, v string) bool {
			return strings.Contains(s, v)
		})
		if err == nil {
			return resp, nil
		}
		if !errors.Is(err, ErrNotMatch) {
			return "", err
		}
	}

	return rangeReply(repliesCache[AllSender], params, func(s, v string) bool {
		return strings.Contains(s, v)
	})
}

func rangeReply(replies map[string]ReplyCache, params Params, match func(s, v string) bool) (string, error) {
	for msg, reply := range replies {
		if match(params.Message, msg) {
			return randReply(reply.Replies), nil
		}
	}
	return "", errors.WithStack(ErrNotMatch)
}

func randReply(replies []string) string {
	rd := rand.Intn(len(replies))
	return replies[rd]
}

type GetReplies func(ctx context.Context) (row []SyncReply, err error)

// RegisterSyncReply 同步 回复
func RegisterSyncReply(ctx context.Context, getReplies GetReplies) {
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			delaySync(ctx, getReplies)
		}
	}()
}

func delaySync(ctx context.Context, getReplies GetReplies) {
	// 5分钟同步一次
	defer time.Sleep(time.Minute * 5)
	replies, err := getReplies(ctx)
	if err != nil {
		// TODO log
		return
	}

	rwMutex.Lock()
	defer rwMutex.Unlock()
	repliesCache = map[string]map[string]ReplyCache{}
	for _, reply := range replies {
		if _, ok := repliesCache[reply.Sender]; !ok {
			repliesCache[reply.Sender] = map[string]ReplyCache{}
		}
		repliesCache[reply.Sender][reply.Msg] = ReplyCache{
			Msg:     reply.Msg,
			Replies: reply.Replies,
		}
	}
}
