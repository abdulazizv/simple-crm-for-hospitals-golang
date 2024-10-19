CREATE TABLE IF NOT EXISTS korik
(
    id SERIAL NOT NULL PRIMARY KEY,
    client_id INT NOT NULL REFERENCES clients(id),
    doctor_id INT NOT NULL REFERENCES doctors(id),
    diagnosis_id INT NOT NULL REFERENCES diagnosis(id),
    file_url TEXT NOT NULL,
    count INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
)