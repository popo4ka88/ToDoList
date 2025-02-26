CREATE TABLE todos (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       completed BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);