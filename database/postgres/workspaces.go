package postgres

import (
	"github.com/bektosh/fiber-app/models/responses"
	"github.com/jmoiron/sqlx"
	"time"
)

func (p *Postgres) CreateWorkspace(id, name string) (string, error) {
	var createdAt time.Time

	row := p.db.QueryRowx(
		`INSERT INTO workspaces(id, name, created_at) VALUES ($1, $2, CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
				RETURNING created_at`,
		id,
		name,
	)
	err := row.Scan(&createdAt)
	if err != nil {
		return "", err
	}

	return createdAt.Format(time.RFC822), nil
}

func (p *Postgres) GetWorkspaces(page, limit int, name string) ([]responses.Workspace, int, error) {
	var (
		rows       *sqlx.Rows
		err        error
		workspaces = make([]responses.Workspace, 0, 10)
		offset     = (page - 1) * limit
		count      int
		createdAt  time.Time
	)

	if name == "" {
		err = p.db.Get(
			&count,
			`SELECT count(*) FROM workspaces`,
		)
		if err != nil {
			return nil, 0, err
		}
		rows, err = p.db.Queryx(
			`SELECT id,
       					  name,
       					  created_at
       		FROM workspaces
       		ORDER BY created_at
       		OFFSET $1 LIMIT $2`,
			offset, limit,
		)
	} else {
		name = `'%` + name + `%'`
		err = p.db.Get(
			&count,
			`SELECT count(*) FROM workspaces
				   WHERE name LIKE $1`,
			name,
		)
		if err != nil {
			return nil, 0, err
		}
		rows, err = p.db.Queryx(
			`SELECT id,
       					  name,
       					  created_at
       		FROM workspaces
       		WHERE name LIKE ` + name +
       		`OFFSET $1 LIMIT $2`,
			offset, limit,
		)
	}
	defer func() {
		_ = rows.Close()
	}()
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var workspace responses.Workspace
		err = rows.Scan(
			&workspace.ID,
			&workspace.Name,
			&createdAt,
		)
		if err != nil {
			return nil, 0, err
		}
		workspace.CreatedAt = createdAt.Format(time.RFC822)
		workspaces = append(workspaces, workspace)
	}
	return workspaces, count, nil
}

func (p *Postgres) UpdateWorkspace(id, name string) error {
	_, err := p.db.Exec(
		`UPDATE workspaces SET name = $1 WHERE id = $2`,
		name, id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteWorkspace(id string) error {
	return transact(p.db, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`DELETE FROM workspaces WHERE id = $1`,
			id,
		)
		if err != nil {
			return err
		}
		_, err = tx.Exec(`UPDATE users SET workspace_id = array_remove(workspace_id, $1) WHERE $1 IN (workspace_id)`,
			id,
		)
		if err != nil {
			return err
		}
		return nil
	})
}
