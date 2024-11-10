package handler

type DbHandler interface{
	initDb()
	backupDb()
	reportDb(reportName string)
}