# TODO List:

* [x] Validation db existense code for expenses only
* [x] Backup code for expenses only
* [x] Expenses Database initialization
* [x] Family db initialization
* [ ] Test new dbs on PI
* [ ] Redirect budget bot on new db
* [ ] Configure report command
* [ ] Configure general validation command for schemas
* [ ] Configure general backup command
* [ ] Configure get and get-all


# Docker commands
```
docker build -t firepand4/fortress:database-handler .
docker run -dit --rm --name database-handler -v "$(Get-Location)/databases:/databases" firepand4/fortress:database-handler
docker start database-handler



```

# Brainstorm for future
- add db versioning (migrations)
    - https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md


