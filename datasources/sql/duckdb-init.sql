-- name: duckdb-tables
CREATE TABLE IF NOT EXISTS settings (
    id VARCHAR(50) UNIQUE,
    config JSON
);

CREATE TABLE IF NOT EXISTS logs (
    client_ip TEXT,
    timestamp TIMESTAMP,
    method TEXT,
    path TEXT,
    protocol TEXT,
    status_code INTEGER,
    latency TEXT,
    user_agent TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

-- name: mysql-table
CREATE TABLE IF NOT EXISTS mysql_db.logs (
    client_ip TEXT,
    timestamp TIMESTAMP,
    method TEXT,
    path TEXT,
    protocol TEXT,
    status_code INTEGER,
    latency TEXT,
    user_agent TEXT,
    created_at TIMESTAMP,
);

-- name: finish-init
INSERT INTO settings VALUES (
    'app',
    '{"database_initialized": true}' 
)