package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xwb1989/sqlparser"
)

// Client interface abstracts the PostgreSQL client
type Client interface {
	Database(string) Database
	Disconnect(context.Context) error
	Ping(context.Context) error
}

// Database interface abstracts PostgreSQL database operations
type Database interface {
	Table(string) Table
	Client() Client
}

// Table interface abstracts PostgreSQL table operations
type Table interface {
	FindOne(context.Context, []string, string, ...interface{}) SingleResult
	FindMany(context.Context, []string, string, ...interface{}) (Cursor, error)
	InsertOne(context.Context, map[string]interface{}) error
	InsertMany(context.Context, []map[string]interface{}) error
	UpdateOne(context.Context, map[string]interface{}, map[string]interface{}) error
	UpdateMany(context.Context, map[string]interface{}, map[string]interface{}) error
	DeleteOne(context.Context, map[string]interface{}) error
	DeleteMany(context.Context, map[string]interface{}) error
}

// SingleResult abstracts a single query result
type SingleResult interface {
	Scan(dest ...interface{}) error
}

// Cursor abstracts query results
type Cursor interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close()
}

// postgresClient holds the pgx pool
type postgresClient struct {
	pool *pgxpool.Pool
}

// postgresDatabase holds the db name
type postgresDatabase struct {
	name   string
	client *postgresClient
}

// postgresTable represents a table
type postgresTable struct {
	db    *postgresDatabase
	table string
}

func (pt *postgresTable) UpdateMany(ctx context.Context, m map[string]interface{}, m2 map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (pt *postgresTable) DeleteMany(ctx context.Context, m map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

// postgresSingleResult represents a single result
type postgresSingleResult struct {
	err error
	row pgx.Row
}

// postgresCursor represents multiple rows
type postgresCursor struct {
	rows pgx.Row
}

func (pc *postgresCursor) Next() bool {
	//TODO implement me
	panic("implement me")
}

func (pc *postgresCursor) Close() {
	//TODO implement me
	panic("implement me")
}

// NewClient creates a new Postgres client
func NewClient(dsn string) (Client, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return &postgresClient{pool: pool}, nil
}

func (pc *postgresClient) Disconnect(ctx context.Context) error {
	pc.pool.Close()
	return nil
}

func (pc *postgresClient) Ping(ctx context.Context) error {
	return pc.pool.Ping(ctx)
}

func (pc *postgresClient) Database(name string) Database {
	return &postgresDatabase{name: name, client: pc}
}

func (pd *postgresDatabase) Table(tableName string) Table {
	if !isValidIdentifier(tableName) {
		log.Fatalf("Invalid table name detected: %s", tableName)
	}
	return &postgresTable{db: pd, table: tableName}
}

func (pd *postgresDatabase) Client() Client {
	return pd.client
}

func (pt *postgresTable) FindOne(ctx context.Context, columns []string, condition string, args ...interface{}) SingleResult {
	cols := "*"
	if len(columns) > 0 {
		for _, col := range columns {
			if !isValidIdentifier(col) {
				log.Fatalf("Invalid column name detected: %s", col)
			}
		}
		cols = strings.Join(columns, ", ")
	}
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT 1", cols, pt.table, condition)
	if err := validateSQL(query); err != nil {
		log.Fatalf("Potential SQL Injection detected: %v", err)
	}
	row := pt.db.client.pool.QueryRow(ctx, query, args...)
	return &postgresSingleResult{row: row}
}

func (pt *postgresTable) FindMany(ctx context.Context, columns []string, condition string, args ...interface{}) (Cursor, error) {
	cols := "*"
	if len(columns) > 0 {
		for _, col := range columns {
			if !isValidIdentifier(col) {
				return nil, fmt.Errorf("Invalid column name detected: %s", col)
			}
		}
		cols = strings.Join(columns, ", ")
	}
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", cols, pt.table, condition)
	if err := validateSQL(query); err != nil {
		return nil, err
	}
	rows, err := pt.db.client.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &postgresCursor{rows: rows}, nil
}

func (pt *postgresTable) InsertOne(ctx context.Context, data map[string]interface{}) error {
	columns := []string{}
	placeholders := []string{}
	values := []interface{}{}
	for col, val := range data {
		if !isValidIdentifier(col) {
			return errors.New("invalid column name detected")
		}
		columns = append(columns, col)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)+1))
		values = append(values, val)
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", pt.table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	if err := validateSQL(query); err != nil {
		return err
	}
	_, err := pt.db.client.pool.Exec(ctx, query, values...)
	return err
}

func (pt *postgresTable) InsertMany(ctx context.Context, data []map[string]interface{}) error {
	tx, err := pt.db.client.pool.Begin(ctx)
	if err != nil {
		return err
	}
	for _, row := range data {
		err = pt.InsertOne(ctx, row)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	return tx.Commit(ctx)
}

func (pt *postgresTable) UpdateOne(ctx context.Context, conditions map[string]interface{}, updates map[string]interface{}) error {
	setClauses := []string{}
	values := []interface{}{}
	for col, val := range updates {
		if !isValidIdentifier(col) {
			return errors.New("invalid column name in update")
		}
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, len(values)+1))
		values = append(values, val)
	}

	condClauses := []string{}
	for col, val := range conditions {
		if !isValidIdentifier(col) {
			return errors.New("invalid column name in condition")
		}
		condClauses = append(condClauses, fmt.Sprintf("%s = $%d", col, len(values)+1))
		values = append(values, val)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", pt.table, strings.Join(setClauses, ", "), strings.Join(condClauses, " AND "))
	if err := validateSQL(query); err != nil {
		return err
	}
	_, err := pt.db.client.pool.Exec(ctx, query, values...)
	return err
}

func (pt *postgresTable) DeleteOne(ctx context.Context, conditions map[string]interface{}) error {
	condClauses := []string{}
	values := []interface{}{}
	for col, val := range conditions {
		if !isValidIdentifier(col) {
			return errors.New("invalid column name in delete condition")
		}
		condClauses = append(condClauses, fmt.Sprintf("%s = $%d", col, len(values)+1))
		values = append(values, val)
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", pt.table, strings.Join(condClauses, " AND "))
	if err := validateSQL(query); err != nil {
		return err
	}
	_, err := pt.db.client.pool.Exec(ctx, query, values...)
	return err
}

func (psr *postgresSingleResult) Scan(dest ...interface{}) error {
	return psr.row.Scan(dest...)
}

func (pc *postgresCursor) Scan(dest ...interface{}) error {
	return pc.rows.Scan(dest...)
}

// Helper to validate SQL using sqlparser
func validateSQL(query string) error {
	_, err := sqlparser.Parse(query)
	return err
}

// isValidIdentifier checks table/column names to prevent injection
func isValidIdentifier(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return re.MatchString(name)
}
