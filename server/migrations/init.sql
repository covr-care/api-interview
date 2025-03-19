create table if not exists time_punches (
  id serial primary key,
  time_clock_id uuid,
  punch_in text,
  punch_out text
);

