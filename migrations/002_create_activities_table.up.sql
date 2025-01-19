-- Tabel activities
CREATE TABLE activities (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    activity_type VARCHAR(255) NOT NULL,
    done_at TIMESTAMP NOT NULL,
    duration_in_minutes INT NOT NULL,
    calories_burned INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);