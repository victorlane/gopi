CREATE TABLE IF NOT EXISTS logs (
    client_ip TEXT,
    time_stamp TIMESTAMP,
    method TEXT,
    path TEXT,
    protocol TEXT,
    status_code INTEGER,
    latency TEXT,
    user_agent TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
)