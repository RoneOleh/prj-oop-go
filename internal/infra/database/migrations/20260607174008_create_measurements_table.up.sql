CREATE TABLE measurements (
    id BIGSERIAL PRIMARY KEY,
    device_id BIGINT NOT NULL,
    room_id BIGINT,
    value DOUBLE PRECISION NOT NULL,
    type VARCHAR(255) NOT NULL,
    created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP,
    deleted_date TIMESTAMP,

    
    CONSTRAINT fk_measurement_device 
        FOREIGN KEY (device_id) 
        REFERENCES devices(id) 
        ON DELETE CASCADE,

    CONSTRAINT fk_measurement_room 
        FOREIGN KEY (room_id) 
        REFERENCES rooms(id) 
        ON DELETE SET NULL
);


CREATE INDEX idx_measurements_device_time 
    ON measurements (device_id, created_date);