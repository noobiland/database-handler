# TODO List:

* [ ] Validation code for expenses only
* [ ] Backup code for expenses only
* [ ] Expenses Database initialization
* [ ] Family db initialization
* [ ] Redirect budget bot on new db
* [ ] Configure report command
* [ ] Configure general validation command
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


