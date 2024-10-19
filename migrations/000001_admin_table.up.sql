CREATE TABLE IF NOT EXISTS admins(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);


INSERT INTO admins(name, username, password, refresh_token) VALUES ('admin', 'first_admin', '123pass', 'refresh_token');