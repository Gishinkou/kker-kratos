package verificationcoderedis

import (
	"context"
	"fmt"
	"time"

	"github.com/Gishinkou/kker-kratos/backend/gopkgs/components/redisx"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redisCli *redis.Client
}

func New() *RedisRepository {
	return &RedisRepository{redisCli: redisx.GetClient(context.Background())}
}

func (r *RedisRepository) formatVerificationCodeKey(verificationCodeId int64) string {
	return fmt.Sprintf("VERIFICATION_CODE_ID_%d", verificationCodeId)
}

func (r *RedisRepository) Save(ctx context.Context, verificationCodeId, expireTime int64, code string) error {
	key := r.formatVerificationCodeKey(verificationCodeId)
	_, err := r.redisCli.Set(ctx, key, code, time.Duration(expireTime)*time.Millisecond).Result()
	return err
}

func (r *RedisRepository) Get(ctx context.Context, verificationCodeId int64) (string, error) {
	code, err := r.redisCli.Get(ctx, r.formatVerificationCodeKey(verificationCodeId)).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}

func (r *RedisRepository) Remove(ctx context.Context, verificationCodeId int64) error {
	_, err := r.redisCli.Del(ctx, r.formatVerificationCodeKey(verificationCodeId)).Result()
	if err != nil {
		return err
	}

	return nil
}
