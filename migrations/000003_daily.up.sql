
CREATE TABLE IF NOT EXISTS daily (
    "w_date" DATE NOT NULL DEFAULT CURRENT_DATE,
    "worker_id" UUID NOT NULL REFERENCES workers(id),
    "come_time" TIME NOT NULL,
    "leave_time" TIME ,
    "w_hours" FLOAT,
    "status" BOOLEAN,
    "late_min" FLOAT
);
