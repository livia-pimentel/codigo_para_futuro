CREATE TABLE oportunidades (
    id BIGSERIAL PRIMARY KEY,
    titulo TEXT NOT NULL,
    empresa TEXT NOT NULL,
    categoria TEXT NOT NULL,
    ativa BOOLEAN NOT NULL DEFAULT TRUE
)