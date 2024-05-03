package cache

import (
	"bibi/cmd/chat/service/ws"
	"bibi/pkg/constants"
	"context"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"time"
)

func SetMessage(ctx context.Context, targetId int64, marshalMsg []byte) error {
	tx := rMessage.TxPipeline()
	var m = map[string][]byte{
		"marshalMsg": marshalMsg,
		"timestamp":  []byte(time.Now().String()),
	}
	marshalRdsMsg, _ := sonic.Marshal(&m)
	if err := tx.ZAdd(ctx, i64ToStr(targetId)+constants.ReceiveSuffix, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: marshalRdsMsg,
	}).Err(); err != nil {
		return err
	}
	if err := tx.Expire(ctx, i64ToStr(targetId)+constants.ReceiveSuffix, constants.MessageExpTime).Err(); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func GetMessages(ctx context.Context, uid int64) (count int64, replyMsgs []ws.ReplyMsg, err error) {

	count, err = rMessage.ZCard(ctx, i64ToStr(uid)+constants.ReceiveSuffix).Result()
	if err != nil {
		return 0, nil, err
	}

	marshalRdsMsgs, err := rMessage.ZRevRange(ctx, i64ToStr(uid)+constants.ReceiveSuffix, 0, -1).Result()
	if err != nil {
		return 0, nil, err
	}
	if err = rMessage.ZRemRangeByRank(ctx, i64ToStr(uid)+constants.ReceiveSuffix, 0, -1).Err(); err != nil {
		return 0, nil, err
	}

	for _, marshalRdsMsg := range marshalRdsMsgs {
		um := make(map[string][]byte)
		_ = sonic.Unmarshal([]byte(marshalRdsMsg), &um)
		var replyMsg ws.ReplyMsg
		_, err = replyMsg.UnmarshalMsg(um["marshalMsg"])
		replyMsgs = append(replyMsgs, replyMsg)
	}

	return
}

func IsUserChattedByOthers(ctx context.Context, uid int64) (bool, error) {
	n, err := rMessage.Exists(ctx, i64ToStr(uid)+constants.ReceiveSuffix).Result()
	if err != nil {
		return true, err
	}
	if n > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
