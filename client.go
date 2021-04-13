package xredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// A Client defines all client Redis actions (Client or ClusterClient).
type Client interface {
	Append(ctx context.Context, key, value string) *redis.IntCmd
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd
	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd
	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd
	BgRewriteAOF(ctx context.Context) *redis.StatusCmd
	BgSave(ctx context.Context) *redis.StatusCmd
	BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd
	BitField(ctx context.Context, key string, args ...interface{}) *redis.IntSliceCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd
	ClientGetName(ctx context.Context) *redis.StringCmd
	ClientID(ctx context.Context) *redis.IntCmd
	ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd
	ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd
	ClientList(ctx context.Context) *redis.StringCmd
	ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd
	ClientUnblock(ctx context.Context, id int64) *redis.IntCmd
	ClientUnblockWithError(ctx context.Context, id int64) *redis.IntCmd
	Close() error
	ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd
	ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd
	ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd
	ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd
	ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd
	ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd
	ClusterFailover(ctx context.Context) *redis.StatusCmd
	ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd
	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd
	ClusterInfo(ctx context.Context) *redis.StringCmd
	ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd
	ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd
	ClusterNodes(ctx context.Context) *redis.StringCmd
	ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd
	ClusterResetHard(ctx context.Context) *redis.StatusCmd
	ClusterResetSoft(ctx context.Context) *redis.StatusCmd
	ClusterSaveConfig(ctx context.Context) *redis.StatusCmd
	ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd
	ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd
	Command(ctx context.Context) *redis.CommandsInfoCmd
	ConfigGet(ctx context.Context, parameter string) *redis.SliceCmd
	ConfigResetStat(ctx context.Context) *redis.StatusCmd
	ConfigRewrite(ctx context.Context) *redis.StatusCmd
	ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd
	// Conn(ctx context.Context) *redis.Conn
	Context() context.Context
	DBSize(ctx context.Context) *redis.IntCmd
	DebugObject(ctx context.Context, key string) *redis.StringCmd
	Decr(ctx context.Context, key string) *redis.IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Do(ctx context.Context, args ...interface{}) *redis.Cmd
	Dump(ctx context.Context, key string) *redis.StringCmd
	Echo(ctx context.Context, message interface{}) *redis.StringCmd
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
	FlushAllAsync(ctx context.Context) *redis.StatusCmd
	FlushDB(ctx context.Context) *redis.StatusCmd
	FlushDBAsync(ctx context.Context) *redis.StatusCmd
	GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd
	GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd
	GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd
	GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd
	HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd
	HExists(ctx context.Context, key, field string) *redis.BoolCmd
	HGet(ctx context.Context, key, field string) *redis.StringCmd
	HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd
	HKeys(ctx context.Context, key string) *redis.StringSliceCmd
	HLen(ctx context.Context, key string) *redis.IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd
	HVals(ctx context.Context, key string) *redis.StringSliceCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
	IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd
	Info(ctx context.Context, section ...string) *redis.StringCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
	LIndex(ctx context.Context, key string, index int64) *redis.StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd
	LLen(ctx context.Context, key string) *redis.IntCmd
	LPop(ctx context.Context, key string) *redis.StringCmd
	LPos(ctx context.Context, key string, value string, a redis.LPosArgs) *redis.IntCmd
	LPosCount(ctx context.Context, key string, value string, count int64, a redis.LPosArgs) *redis.IntSliceCmd
	LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd
	LastSave(ctx context.Context) *redis.IntCmd
	MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd
	MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd
	MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd
	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd
	Move(ctx context.Context, key string, db int) *redis.BoolCmd
	ObjectEncoding(ctx context.Context, key string) *redis.StringCmd
	ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd
	ObjectRefCount(ctx context.Context, key string) *redis.IntCmd
	// Options() *redis.Options
	PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd
	PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd
	PFCount(ctx context.Context, keys ...string) *redis.IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd
	PSubscribe(ctx context.Context, channels ...string) *redis.PubSub
	PTTL(ctx context.Context, key string) *redis.DurationCmd
	Persist(ctx context.Context, key string) *redis.BoolCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Pipeline() redis.Pipeliner
	Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	PoolStats() *redis.PoolStats
	Process(ctx context.Context, cmd redis.Cmder) error
	PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd
	PubSubNumPat(ctx context.Context) *redis.IntCmd
	PubSubNumSub(ctx context.Context, channels ...string) *redis.StringIntMapCmd
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	Quit(ctx context.Context) *redis.StatusCmd
	RPop(ctx context.Context, key string) *redis.StringCmd
	RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	RandomKey(ctx context.Context) *redis.StringCmd
	ReadOnly(ctx context.Context) *redis.StatusCmd
	ReadWrite(ctx context.Context) *redis.StatusCmd
	Rename(ctx context.Context, key, newkey string) *redis.StatusCmd
	RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd
	Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd
	SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	SCard(ctx context.Context, key string) *redis.IntCmd
	SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd
	SMembers(ctx context.Context, key string) *redis.StringSliceCmd
	SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd
	SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd
	SPop(ctx context.Context, key string) *redis.StringCmd
	SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd
	SRandMember(ctx context.Context, key string) *redis.StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd
	SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	Save(ctx context.Context) *redis.StatusCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
	ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd
	ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd
	ScriptFlush(ctx context.Context) *redis.StatusCmd
	ScriptKill(ctx context.Context) *redis.StatusCmd
	ScriptLoad(ctx context.Context, script string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetArgs(ctx context.Context, key string, value interface{}, a redis.SetArgs) *redis.StatusCmd
	SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Shutdown(ctx context.Context) *redis.StatusCmd
	ShutdownNoSave(ctx context.Context) *redis.StatusCmd
	ShutdownSave(ctx context.Context) *redis.StatusCmd
	SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd
	SlowLogGet(ctx context.Context, num int64) *redis.SlowLogCmd
	Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd
	SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd
	SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd
	StrLen(ctx context.Context, key string) *redis.IntCmd
	// String() string
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	Sync(ctx context.Context)
	TTL(ctx context.Context, key string) *redis.DurationCmd
	Time(ctx context.Context) *redis.TimeCmd
	Touch(ctx context.Context, keys ...string) *redis.IntCmd
	TxPipeline() redis.Pipeliner
	TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	Type(ctx context.Context, key string) *redis.StatusCmd
	Unlink(ctx context.Context, keys ...string) *redis.IntCmd
	Wait(ctx context.Context, numSlaves int, timeout time.Duration) *redis.IntCmd
	Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error
	// WithContext(ctx context.Context) *redis.Client
	// WithTimeout(timeout time.Duration) *redis.Client
	XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd
	XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd
	XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd
	XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd
	XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd
	XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd
	XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd
	XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XInfoConsumers(ctx context.Context, key string, group string) *redis.XInfoConsumersCmd
	XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd
	XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd
	XLen(ctx context.Context, stream string) *redis.IntCmd
	XPending(ctx context.Context, stream, group string) *redis.XPendingCmd
	XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd
	XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd
	XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd
	XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd
	XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd
	XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd
	XRevRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd
	XRevRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd
	XTrim(ctx context.Context, key string, maxLen int64) *redis.IntCmd
	XTrimApprox(ctx context.Context, key string, maxLen int64) *redis.IntCmd
	ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddNX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddXX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZCard(ctx context.Context, key string) *redis.IntCmd
	ZCount(ctx context.Context, key, min, max string) *redis.IntCmd
	ZIncr(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd
	ZIncrNX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZIncrXX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd
	ZMScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd
	ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd
	ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd
	ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd
	ZRank(ctx context.Context, key, member string) *redis.IntCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd
	ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *redis.IntCmd
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	ZScore(ctx context.Context, key, member string) *redis.FloatCmd
	ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd
}
