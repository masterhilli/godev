echo jira/config:
cd jira\Config
go test 
cd ..
cd HtmlConnection
go test
cd ..
cd Timeentry
go test
cd ..\..
