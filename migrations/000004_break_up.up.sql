CREATE TABLE IF NOT EXISTS break_up (
    "date" DATE NOT NULL DEFAULT CURRENT_DATE,
    "worker_id" UUID NOT NULL REFERENCES workers(id),
    "exit" TIME,
    "back" TIME
);