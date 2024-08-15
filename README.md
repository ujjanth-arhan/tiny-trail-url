The aim of the project to create sharable URL shortners.

AIR: $(go env GOPATH)/bin/air

SQL Server Image: docker pull mcr.microsoft.com/mssql/server:2022-latest

docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=<YourStrong@Passw0rd>" \
   -p 1433:1433 --name sql1 --hostname sql1 \
   -d \
   mcr.microsoft.com/mssql/server:2022-latest

Note: Please substitute <YourStrong@Passw0rd> with a strong password
Note: Your password should follow the SQL Server default password policy, otherwise the container can't set up SQL Server, and stops working. By default, the password must be at least eight characters long and contain characters from three of the following four sets: uppercase letters, lowercase letters, base-10 digits, and symbols.

docker exec -it sql1 "bash"

sudo /opt/mssql-tools/bin/sqlcmd -S localhost -U <userid> -P "<YourNewStrong@Passw0rd>"


docker exec -it sql1 "bash"
/opt/mssql-tools/bin/sqlcmd -S localhost -U <userid> -P "<YourStrong@Passw0rd>"

CREATE DATABASE TestDB;

SELECT Name from sys.databases;

GO

Note: QUIT to quit sqlcmd
NOte: exit to exit the interactive terminal


Future Enchancements:
- Create utility method for read of body and unmarshall of json
- Support Database
- Support migration system
- Remove background context and use actual ones with support for time out
- Consider a use case where duplicate URLs might be needed
- Track remote address
- Create a distributed UUID generator
- Support swagger
- Support docker container deployment
- Have a shorter URL
- Support executable file versioning
- Covert all responses to JSON
- Change to use prepared context (don't forget to do the same in DB setup)
- Nice way to convert struct to json
- Change DB schema to prevent non null of created at
- Code cleanup of repo.setupdatabase
- Write all errors to a different file
- Quoted string for DB setup
- Redirect if opened from a browser