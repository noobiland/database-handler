# TODO List:

* [x] Validation db existense code for expenses only
* [x] Backup code for expenses only
* [x] Expenses Database initialization
* [x] Family db initialization
* [x] Test new dbs on PI
* [X] ~~*data migration mechanism*~~ [2024-12]
* [X] ~~*Redirect budget bot on new db*~~ [2024-12]
* [ ] Configure report command
* [ ] Configure general validation command for schemas
* [ ] Configure general backup command
* [ ] Configure get and get-all for supported dbs
* [ ] Can't create required folders, if missing
* [ ] When table is empty, migration option returns error



# Docker commands
```
docker build -t firepand4/fortress:database-handler .
docker run -dit --rm --name database-handler -v "$(Get-Location)/databases:/databases" firepand4/fortress:database-handler
docker start database-handler



```

# Brainstorm for future
- add db versioning (migrations)
    - https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md


