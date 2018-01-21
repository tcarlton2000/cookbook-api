CREATE TABLE IF NOT EXISTS ingredients (
    id                 serial primary key,
    name               varchar(150) UNIQUE,
    type               varchar(50),
    serving_size       numeric,
    unit               varchar(20),
    calories           numeric,
    carbs              numeric,
    protein            numeric,
    fat                numeric,
    cholestorol        numeric
);

CREATE TABLE IF NOT EXISTS recipes (
    id                 serial primary key,
    name               varchar(150) UNIQUE
);

CREATE TABLE IF NOT EXISTS recipe_ingredients (
    id                 serial primary key,
    recipe_id          integer references recipes,
    amount             numeric,
    unit               varchar(150)
);

CREATE TABLE IF NOT EXISTS recipe_steps (
    id                 serial primary key,
    recipe_id          integer references recipes,
    step_order         numeric,
    instruction        varchar(2000)
);