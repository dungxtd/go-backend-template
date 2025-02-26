package postgres

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

var (
	ErrInvalidSQLInput = errors.New("invalid SQL input detected")
	// Regular expression to detect common SQL injection patterns
	sqlInjectionPattern = regexp.MustCompile(`(?i)(--|--%|;s*$|/\*|\*/|@@|sys\.|sysobjects\.|syscolumns\.|xp_\w+|sp_\w+|exec\s+\w+|execute\s+\w+|WAITFOR\s+DELAY\s+'|WAITFOR\s+TIME\s+'|;\s*SHUTDOWN\s+|;s*xp_\w+|'\s*OR\s*'\d+'\s*=\s*'\d+'|'\s*OR\s*'\w+'\s*=\s*'\w+'|'\s*OR\s*'.*?'\s*=\s*'.*?')`)
)

type Database interface {
	Exec(context.Context, string, ...interface{}) (sql.Result, error)
	Query(context.Context, string, ...interface{}) (Rows, error)
	QueryRow(context.Context, string, ...interface{}) Row
	Begin(context.Context) (Tx, error)
	Client() Client
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Rows interface {
	Close() error
	Next() bool
	Scan(dest ...interface{}) error
}

type Tx interface {
	Commit() error
	Rollback() error
	Exec(context.Context, string, ...interface{}) (sql.Result, error)
	Query(context.Context, string, ...interface{}) (Rows, error)
	QueryRow(context.Context, string, ...interface{}) Row
}

type Client interface {
	Connect(context.Context) error
	Close() error
	Ping(context.Context) error
	Database() Database
}

type postgresClient struct {
	db *sql.DB
}

type postgresDatabase struct {
	db *sql.DB
}

type postgresRow struct {
	row *sql.Row
}

type postgresRows struct {
	rows *sql.Rows
}

type postgresTx struct {
	tx *sql.Tx
}

// validateInput checks for potential SQL injection patterns in the input
func validateInput(input string) error {
	if sqlInjectionPattern.MatchString(input) {
		return ErrInvalidSQLInput
	}
	return nil
}

// validateQueryAndArgs validates the SQL query and its arguments
func validateQueryAndArgs(query string, args ...interface{}) error {
	// Validate the base query
	if err := validateInput(query); err != nil {
		return err
	}

	// Validate string arguments
	for _, arg := range args {
		if strArg, ok := arg.(string); ok {
			if err := validateInput(strArg); err != nil {
				return err
			}
		}
	}

	return nil
}

func NewClient(dsn string) (Client, error) {
	// Validate DSN string
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("empty database connection string")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &postgresClient{db: db}, nil
}

func (pc *postgresClient) Connect(ctx context.Context) error {
	return pc.db.PingContext(ctx)
}

func (pc *postgresClient) Close() error {
	return pc.db.Close()
}

func (pc *postgresClient) Ping(ctx context.Context) error {
	return pc.db.PingContext(ctx)
}

func (pc *postgresClient) Database() Database {
	return &postgresDatabase{db: pc.db}
}

func (pd *postgresDatabase) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// Validate query and arguments
	if err := validateQueryAndArgs(query, args...); err != nil {
		return nil, err
	}

	// Prepare statement to ensure proper escaping
	stmt, err := pd.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.ExecContext(ctx, args...)
}

func (pd *postgresDatabase) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	// Validate query and arguments
	if err := validateQueryAndArgs(query, args...); err != nil {
		return nil, err
	}

	// Prepare statement to ensure proper escaping
	stmt, err := pd.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return &postgresRows{rows: rows}, nil
}

func (pd *postgresDatabase) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	// Validate query and arguments
	if err := validateQueryAndArgs(query, args...); err != nil {
		return &postgresRow{row: pd.db.QueryRowContext(ctx, "SELECT NULL WHERE FALSE")}
	}

	// Prepare statement to ensure proper escaping
	stmt, err := pd.db.PrepareContext(ctx, query)
	if err != nil {
		return &postgresRow{row: pd.db.QueryRowContext(ctx, "SELECT NULL WHERE FALSE")}
	}
	defer stmt.Close()

	return &postgresRow{row: stmt.QueryRowContext(ctx, args...)}
}

func (pd *postgresDatabase) Begin(ctx context.Context) (Tx, error) {
	tx, err := pd.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &postgresTx{tx: tx}, nil
}

func (pd *postgresDatabase) Client() Client {
	return &postgresClient{db: pd.db}
}

func (pr *postgresRow) Scan(dest ...interface{}) error {
	return pr.row.Scan(dest...)
}

func (pr *postgresRows) Close() error {
	return pr.rows.Close()
}

func (pr *postgresRows) Next() bool {
	return pr.rows.Next()
}

func (pr *postgresRows) Scan(dest ...interface{}) error {
	return pr.rows.Scan(dest...)
}

func (pt *postgresTx) Commit() error {
	return pt.tx.Commit()
}

func (pt *postgresTx) Rollback() error {
	return pt.tx.Rollback()
}

func (pt *postgresTx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return pt.tx.ExecContext(ctx, query, args...)
}

func (pt *postgresTx) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rows, err := pt.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &postgresRows{rows: rows}, nil
}

func (pt *postgresTx) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	row := pt.tx.QueryRowContext(ctx, query, args...)
	return &postgresRow{row: row}
}
