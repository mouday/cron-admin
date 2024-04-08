package enums

const (
	// === 启动阶段 ===
	// 启动失败
	TaskStatusStartError = 1
	// 启动成功
	TaskStatusStartSuccess = 2

	// === 运行阶段 ===
	// 开始运行
	TaskStatusStartRun = 3
	// 运行中
	TaskStatusRunning = 4

	// === 结束阶段 ===
	// 运行成功
	TaskStatusRunSuccess = 5
	// 运行失败
	TaskStatusRunError = 6
)
