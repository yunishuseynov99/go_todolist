package tasks_postgres_repository

import (
	"context"
	"fmt"
	core_errors "github.com/yunishuseynov99/go_todolist/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM todoapp.tasks 
		 	  WHERE id = $1;`

	cmdtag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query delete task: %w", err)
	}

	if cmdtag.RowsAffected() == 0 {
		return fmt.Errorf(
			"task with id=`%d` not found: %w",
			id,
			core_errors.ErrNotFound,
		)
	}

	return nil
}
