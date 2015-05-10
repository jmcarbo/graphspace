package main

import (
	"database/sql"
	"fmt"
	"io"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
create table if not exists graphs (
	id text not null primary key, 
	format text, 
	text text,
	width int, 
	height int, 
	output text,
	description text
);
`

type sqlGraphviz struct {
	dbpath      string
	db          *sql.DB
	stmt_insert *sql.Stmt
	stmt_select *sql.Stmt
	stmt_delete *sql.Stmt
}

// prepare the database for usage
func PrepareSqliteDb(dbpath string) error {

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func NewSqlGraphviz(dbpath string) (*sqlGraphviz, error) {
	err := PrepareSqliteDb(dbpath)
	if err != nil {
		return nil, fmt.Errorf("prepare %s: %s", dbpath, err)
	}

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	stmt_insert, err := db.Prepare("insert into graphs (id, format, text, width, height, output, description) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	stmt_select, err := db.Prepare("select format, text, width, height, output, description from graphs where id = ?")
	if err != nil {
		return nil, err
	}

	stmt_delete, err := db.Prepare("delete from graphs where id = ?")
	if err != nil {
		return nil, err
	}

	q := &sqlGraphviz{
		db:          db,
		dbpath:      dbpath,
		stmt_insert: stmt_insert,
		stmt_select: stmt_select,
		stmt_delete: stmt_delete,
	}
	return q, nil
}

func (q *sqlGraphviz) Create(g *Graph) (string, error) {

	id := g.GetId()
	log.Tracef("graph id=%s", id)

	_, err := q.stmt_insert.Exec(id, g.Format, g.Text, g.Width, g.Height, g.Output, g.Description)
	if err != nil {
		return "", err
	}
	log.Tracef("create id %s from graph %d bytes", id, len(g.Text))
	return id, nil
}

func (q *sqlGraphviz) Get(id string) (*Graph, error) {
	row := q.stmt_select.QueryRow(id)
	if row == nil {
		return nil, io.EOF
	}

	g := &Graph{}

	err := row.Scan(&g.Format, &g.Text, &g.Width, &g.Height, &g.Output, &g.Description)
	if err == sql.ErrNoRows {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	return g, nil
}

func (q *sqlGraphviz) Delete(id string) error {
	_, err := q.stmt_delete.Exec(id)
	if err != nil {
		return err
	}
	return err
}

func (q *sqlGraphviz) Close() error {
	return q.db.Close()
}
