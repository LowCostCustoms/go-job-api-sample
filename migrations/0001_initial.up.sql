begin;

create table jobs(
    id uuid primary key,
    name text
);

create table job_schedules(
    id uuid primary key,
    job_id uuid references jobs(id) on delete restrict,
    cron text
);

create index ix__job_schedules__job_id on job_schedules(job_id);

commit;
