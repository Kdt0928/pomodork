package entities

import "time"

// Task タスク
type Task struct {
	TaskId                string    // タスクID
	ScheduleDate          time.Time // 予定日
	DeleteFlag            string    // 削除フラグ
	TaskName              string    // タスク名
	PomodoroNum           int32     // ポモドーロ数
	PomodoroTime          int32     // ポモドーロ時間
	CompletionPomodoroNum float32   // 完了ポモドーロ数
	CompletionDate        time.Time // 完了
	Memo                  string    // メモ
}
