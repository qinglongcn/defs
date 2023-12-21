package defs

import (
	"context"

	"github.com/bpfs/defs/core/pool"
	"github.com/bpfs/defs/eventbus"

	"github.com/bpfs/defs/sqlites"
	"github.com/bpfs/dep2p"
	"github.com/bpfs/dep2p/pubsub"
	"github.com/dgraph-io/ristretto"
)

// FS提供了与DeFS交互所需的各种函数
type FS struct {
	ctx          context.Context     // 全局上下文
	opt          *Options            // 文件存储选项配置
	p2p          *dep2p.DeP2P        // DeP2P网络主机
	pubsub       *pubsub.DeP2PPubSub // DeP2P网络订阅
	db           *sqlites.SqliteDB   // sqlite数据库服务
	uploadChan   chan *uploadChan    // 用于刷新上传的通道
	downloadChan chan *downloadChan  // 用于刷新下载的通道

	registry *eventbus.EventRegistry // 事件总线
	cache    *ristretto.Cache        // 缓存实例
	pool     *pool.MemoryPool        // 内存池
}
