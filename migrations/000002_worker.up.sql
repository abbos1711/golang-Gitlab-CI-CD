
CREATE TABLE IF NOT EXISTS "workers"(
    "id" UUID NOT NULL PRIMARY KEY,
    "img_url" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "position" VARCHAR(255) NOT NULL,
    "department" VARCHAR(255) NOT NULL,
    "gender" VARCHAR(255) NOT NULL,
    "contact" VARCHAR(255) NOT NULL,
    "birthday" DATE NOT NULL,
    "come_time" TIME NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE  
);
