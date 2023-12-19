// backend/fsnotify/watcher_test.go

package fsnotify

// go test -v -bench=. -benchtime=610s watcher.go watcher_test.go -run TestXxx
// func TestXxx(t *testing.T) {
// 	app := fx.New(
// 		fx.Provide(
// 			// 新观察者
// 			NewWatcher,
// 		),
// 		fx.Invoke(
// 			// 启动观察器
// 			StartWatcher,
// 		),
// 	)

// 	app.Run()
// }
