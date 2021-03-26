package infrastructure

import (
	"context"
	"time"

	"github.com/ari1021/hack-ios-server/hconstant"
	"github.com/ari1021/hack-ios-server/pkg/domain/entity"
	"github.com/ari1021/hack-ios-server/pkg/domain/repository"
	"gorm.io/gorm"
)

type Task struct {
	ID          string    `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"not null"`
	Description *string
	IsDone      bool   `gorm:"not null"`
	UserID      string `gorm:"size:255;not null"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewTask(taskID string, taskName string, taskDescription *string, isDone bool, userID string) *Task {
	return &Task{
		ID:          taskID,
		Name:        taskName,
		Description: taskDescription,
		IsDone:      isDone,
		UserID:      userID,
	}
}

func taskToEntity(t *Task) (task *entity.Task) {
	task = &entity.Task{
		ID:          t.ID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		DeletedAt:   t.DeletedAt.Time,
		Name:        t.Name,
		Description: t.Description,
		IsDone:      t.IsDone,
		UserID:      t.UserID,
	}
	return task
}

func entityToTask(task *entity.Task) (t *Task) {
	t = NewTask(task.ID, task.Name, task.Description, task.IsDone, task.UserID)
	return t
}

type TaskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.Task {
	return &TaskRepository{
		Conn: conn,
	}
}

func (tr *TaskRepository) CreateTask(ctx context.Context, task *entity.Task) (err error) {
	t := entityToTask(task)
	if err := tr.Conn.WithContext(ctx).Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) FindTasksByUserID(ctx context.Context, userID string) (tasks []*entity.Task, err error) {
	ts := make([]*Task, 0, hconstant.GetTaskLimit)
	if err := tr.Conn.WithContext(ctx).Order("created_at desc").Find(&ts, "user_id = ?", userID).Limit(hconstant.GetTaskLimit).Error; err != nil {
		return nil, err
	}
	for _, t := range ts {
		task := taskToEntity(t)
		tasks = append(tasks, task)
	}
	return tasks, nil
}
