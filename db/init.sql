CREATE TABLE IF NOT EXISTS ingredients (
    id                 serial primary key,
    name               varchar(150),
    type               varchar(50),
    serving_size       numeric,
    unit               varchar(20),
    calories           numeric,
    carbs              numeric,
    protein            numeric,
    fat                numeric,
    cholestorol        numeric
);