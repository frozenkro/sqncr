package projectsdb

import "context"

// DeleteDependenciesForTask deletes all dependency rows that reference the given task
// either as the task_id or depends_on column. This is needed because SQLite may not
// have foreign key enforcement enabled (no PRAGMA foreign_keys=ON), so we manually
// clean up dependencies before deleting a task.
func (q *Queries) DeleteDependenciesForTask(ctx context.Context, taskId int64) error {
	_, err := q.db.ExecContext(ctx, "DELETE FROM dependencies WHERE task_id = ? OR depends_on = ?", taskId, taskId)
	return err
}