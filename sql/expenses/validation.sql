SELECT name
FROM sqlite_master
WHERE type = 'table'
    AND name = 'expenses';

SELECT name
FROM sqlite_master
WHERE type = 'view'
    AND name = 'cur_month_stat';

SELECT name
FROM sqlite_master
WHERE type = 'view'
    AND name = 'prev_month_stat';