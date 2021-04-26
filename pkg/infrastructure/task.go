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

func taskToEntity(t *Task) (task *entity.Task) {
	var deletedAt *time.Time
	if t.DeletedAt.Valid {
		deletedAt = &t.DeletedAt.Time
	} else {
		deletedAt = nil
	}
	task = &entity.Task{
		ID:          t.ID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		DeletedAt:   deletedAt,
		Name:        t.Name,
		Description: t.Description,
		IsDone:      t.IsDone,
		UserID:      t.UserID,
	}
	return task
}

func entityToTask(task *entity.Task) (t *Task) {
	t = &Task{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		IsDone:      task.IsDone,
		UserID:      task.UserID,
	}
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

func (tr *TaskRepository) UpdateTaskDoneByTaskIDs(ctx context.Context, taskIDs []string) (err error) {
	// taskIDをもとにis_doneをtrueに更新
	if err := tr.Conn.WithContext(ctx).Table("tasks").Where("id IN ?", taskIDs).Updates(map[string]interface{}{"is_done": 1}).Error; err != nil {
		return err
	}
	return err
}

func (tr *TaskRepository) FindTasksByTaskIDs(ctx context.Context, taskIDs []string) (tasks []*entity.Task, err error) {
	ts := make([]*Task, 0, len(taskIDs))
	if err := tr.Conn.WithContext(ctx).Table("tasks").Where("id IN ?", taskIDs).Find(&ts).Error; err != nil {
		return nil, err
	}
	for _, t := range ts {
		task := taskToEntity(t)
		tasks = append(tasks, task)
	}
	return tasks, nil
}
