DROP TABLE IF EXISTS events;
CREATE TABLE events (
    id           BIGSERIAL PRIMARY KEY,
    device_id    BIGINT NOT NULL REFERENCES devices(id),
    room_id      BIGINT NOT NULL REFERENCES rooms(id),
    action       VARCHAR(255) NOT NULL,
    created_date TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_date TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_date TIMESTAMP
);