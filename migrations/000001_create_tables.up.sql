CREATE TYPE "attend_type" AS ENUM (
  'attended',
  'absent',
  'late'
);

CREATE TABLE IF NOT EXISTS "schedules" (
  "id" UUID PRIMARY KEY,
  "group_id" UUID,
  "start_time" TIMESTAMPTZ NOT NULL,
  "end_time" TIMESTAMPTZ NOT NULL,
  "branch_id" UUID,
  "teacher_id" UUID,
  "support_teacher_id" UUID,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "lessons" (
  "id" UUID PRIMARY KEY,
  "schedule_id" UUID REFERENCES "schedules" ("id"),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "tasks" (
  "id" UUID PRIMARY KEY,
  "lesson_id" UUID REFERENCES "lessons" ("id"),
  "label" VARCHAR(255) NOT NULL,
  "deadline" TIMESTAMPTZ NOT NULL,
  "score" NUMERIC(3,1) DEFAULT 0,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE IF NOT EXISTS "attendances" (
  "id" UUID PRIMARY KEY,
  "lesson_id" UUID REFERENCES "lessons" ("id"),
  "student_id" UUID,
  "status" attend_type,
  "late_minute" INT DEFAULT 0
);