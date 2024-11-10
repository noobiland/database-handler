CREATE VIEW prev_month_stat(tr_date, user, amount, category, payment) as
SELECT strftime(
        '%Y-%m-%d %H:%M:%S',
        datetime(e.timestamp, 'unixepoch', 'localtime')
    ) as tr_date,
    e.user,
    e.amount,
    e.category,
    e.payment
FROM expenses e
WHERE tr_date >= date('now', 'start of month', '-1 month')
    AND tr_date < date('now', 'start of month');
CREATE VIEW cur_month_stat(tr_date, user, amount, category, payment) as
SELECT strftime(
        '%Y-%m-%d %H:%M:%S',
        datetime(e.timestamp, 'unixepoch', 'localtime')
    ) as tr_date,
    e.user,
    e.amount,
    e.category,
    e.payment
FROM expenses e
WHERE tr_date >= date('now', 'start of month')
    AND tr_date < date('now', 'start of month', '+1 month');