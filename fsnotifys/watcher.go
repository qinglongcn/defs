// backend/fsnotify/watcher.go

package fsnotify

import (
	"bpfs/backend/config"
	"bpfs/pkg/dir"
	"context"
	"fmt"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type StartWatcherInput struct {
	fx.In
	Ctx     context.Context
	Host    host.Host          // 网络主机
	IsState *config.IsAppState // 应用状态
	Watcher *fsnotify.Watcher  // 文件监听
}

// // 定义回调函数类型
// type FileChangedCallback func(path string)

// 启动观察器
func StartWatcher(input StartWatcherInput) error {
	path := filepath.Join(config.SlicePath, input.Host.ID().String())
	if err := dir.DirExistsAndMkdirAll(path); err != nil {
		logrus.Errorf("检查路径失败: %v", err)
		return err
	}

	// 添加路径
	if err := input.Watcher.Add(path); err != nil {
		logrus.Errorf("添加路径失败: %v", err)
		return err
	}

	// 启动监听文件对象事件协程
	fmt.Println("开始监听文件变化")
	go func() {
		for {
			select {
			case event, ok := <-input.Watcher.Events:
				if ok {
					// Has 报告此事件是否具有给定的操作。
					if event.Has(fsnotify.Create) {
						//	logrus.Println("创建文件:", event.Name)
						// 调用回调函数处理文件创建事件
						//callback(event.Name) // 调用回调函数处理逻辑
						// 发送时间
						// 但是发送到发送
						// input.UploadEvent.Publish("file:Upload", event.Name)
					}
				}

			case err, ok := <-input.Watcher.Errors:
				if !ok {
					logrus.Errorf("监听文件失败: %v", err)
					return
				}
			}
		}
	}()

	return nil
}

// 新观察者
func NewWatcher(lc fx.Lifecycle) (*fsnotify.Watcher, error) {
	// NewWatcher 与底层操作系统建立一个新的观察者并开始等待事件。
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return watcher.Close()
		},
	})

	return watcher, nil
}
