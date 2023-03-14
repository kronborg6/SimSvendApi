build:
	@go build -o bin/SimSvend

run: build
	@PUBLIC=August PRIVATE=Kronborg MYSQLUSER=root MYSQLPASSWORD=password MYSQLHOST=localhost MYSQLPORT=3306 MYSQLDATABASE=simsvend ./bin/SimSvend

test:
	@go test -v ./...