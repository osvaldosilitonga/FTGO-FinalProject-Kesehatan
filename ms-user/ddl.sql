CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE,
    address VARCHAR(255),
    phone VARCHAR(20),
    birthdate DATE,
    gender VARCHAR(20),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE user_activities (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    method VARCHAR(10),
    description VARCHAR(255),
    date TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
