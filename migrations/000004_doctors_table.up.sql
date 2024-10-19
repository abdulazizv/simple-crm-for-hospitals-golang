CREATE TABLE  IF NOT EXISTS doctors(
    id SERIAL NOT NULL PRIMARY KEY,
    clinic_id INT NOT NULL REFERENCES clinics(id),
    service_id INT NOT NULL REFERENCES services(id),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone_number VARCHAR(100),
    start_time VARCHAR(100),
    end_time VARCHAR(100),
    work_day VARCHAR(100),
    floor INT,
    room_number INT,
    image_link TEXT,
    experience INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    refresh_token TEXT
);