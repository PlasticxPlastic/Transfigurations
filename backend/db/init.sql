-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert some test users (you can modify these credentials)
INSERT INTO users (username, password) VALUES
    ('admin', 'admin123'),
    ('test', 'test123')
ON CONFLICT (username) DO NOTHING; 