CREATE extension IF NOT EXISTS "uuid-ossp";

CREATE TABLE api_log (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    request_url VARCHAR(255) not NULL,
    request_parameter VARCHAR(255)not NULL,
    status_code INT not NULL,
    response TEXT not NULL,
    created_at TIMESTAMP WITH TIME ZONE
);