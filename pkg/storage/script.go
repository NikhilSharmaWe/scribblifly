package storage

import (
	"database/sql"
	"time"
)

type ScriptStorage interface {
	Create(Script) error
	Delete(username string) error
	Update(oldRecord, newRecord Script) error
	GetAll() ([]Script, error)
	Get(username, title string) (Script, error)
}

type ScriptModel struct {
	DB *sql.DB
}

type Script struct {
	Title     string    `json:"title"`
	Username  string    `json:"username"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *ScriptModel) Create(st Script) error {

	query := `insert into scripts
	(title, username, type, content, create_at)
	values ($1, $2, $3, $4, $5))`

	_, err := s.DB.Exec(query, st.Title, st.Username, st.Type, st.Content, st.CreatedAt)
	return err
}

func (s *ScriptModel) Delete(username string) error {
	query := `delete from scripts where username = $1`

	_, err := s.DB.Exec(query, username)
	return err
}

func (s *ScriptModel) Update(old, new Script) error {
	query := `update scripts 
	set type = $3, title = $4
	WHERE username = $1 and title = $2`

	_, err := s.DB.Exec(query, old.Username, old.Title, new.Type, new.Title)
	return err
}

func (s *ScriptModel) GetAll() ([]Script, error) {
	query := `select * from accounts`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	scripts := []Script{}

	for rows.Next() {
		script, err := scanIntoScript(rows)
		if err != nil {
			return nil, err
		}
		scripts = append(scripts, script)
	}
	return scripts, nil
}

func (s *ScriptModel) Get(username, title string) (Script, error) {
	st := Script{}
	query := `select username, title, type, content, createdAt from scripts where username = $1 and title = $2`
	err := s.DB.QueryRow(query, username, title).Scan(&st.Username, &st.Title, &st.Type, &st.Content, &st.CreatedAt)
	if err != nil {
		return st, err
	}

	return st, nil
}

func scanIntoScript(rows *sql.Rows) (Script, error) {
	script := Script{}
	if err := rows.Scan(
		&script.Username,
		&script.Title,
		&script.Type,
		&script.Content,
		&script.CreatedAt); err != nil {
		return script, err
	}
	return script, nil
}
