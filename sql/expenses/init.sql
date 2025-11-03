CREATE TABLE expenses (
    "timestamp" INTEGER NOT NULL,
    "user" TEXT NOT NULL,
    "amount" INTEGER NOT NULL,
    "category" TEXT NOT NULL,
    "payment" TEXT NOT NULL,
    "comment" TEXT 
);
CREATE VIEW prev_month_stat(tr_date, user, amount, category, payment, comment) as
SELECT strftime(
        '%Y-%m-%d %H:%M:%S',
        datetime(e.timestamp, 'unixepoch', 'localtime')
    ) as tr_date,
    e.user,
    e.amount,
    e.category,
    e.payment,
    e.comment
FROM expenses e
WHERE tr_date >= date('now', 'start of month', '-1 month')
    AND tr_date < date('now', 'start of month');
CREATE VIEW cur_month_stat(tr_date, user, amount, category, payment, comment) as
SELECT strftime(
        '%Y-%m-%d %H:%M:%S',
        datetime(e.timestamp, 'unixepoch', 'localtime')
    ) as tr_date,
    e.user,
    e.amount,
    e.category,
    e.payment,
    e.comment
FROM expenses e
WHERE tr_date >= date('now', 'start of month')
    AND tr_date < date('now', 'start of month', '+1 month');