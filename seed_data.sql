select * from category;

insert into category (name, description)
values ('fitness', 'gym, workouts, suplements, etc.')

select * from expense;

insert into expense (amount, category_id, created)
values (1500, 1, now()::timestamp)

SELECT now()::timestamp;
