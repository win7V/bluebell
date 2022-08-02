package redis

// redis key

// redis key 尽量使用命名空间方式,方便查询和拆分

const (
	KeyPrefix        = "bluebell:"   //通用前缀
	KeyPostTimeZSet  = "post:time"   //zset;帖子及发帖时间
	KeyPostScoreZSet = "post:score"  //zset;帖子及投票的分数
	KeyPostVotedPF   = "post:voted:" //zset;记录用户及投票的类型;参数是post id
)
