CREATE  TABLE  IF NOT EXISTS services(
    id SERIAL NOT NULL PRIMARY KEY,
    clinic_id INT NOT NULL REFERENCES clinics(id),
    name TEXT,  
    price NUMERIC(10, 2)
);