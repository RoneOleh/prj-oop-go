CREATE TABLE IF NOT EXISTS public.organizations
(
    id              serial PRIMARY KEY,
    user_id         varchar(50) NOT NULL REFERENCES public.users(id),
    name            varchar(250) NOT NULL,
    description      text,
    city            varchar(150) NOT NULL,
    addresss        varchar(100) NOT NULL,
    lat             float4 NOT NULL,
    lon             float4 NOT NULL,
    created_date    timestamptz NOT NULL,
    updated_date    timestamptz NOT NULL,
    deleted_date    timestamptz
);
