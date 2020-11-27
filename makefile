run_01:
	mysql -u root -ppass < 01_simple_database_querying/mysql.sql
init_02:
	mysql -u root -ppass < 02_movie_searcher/files/int_mysql.sql
	cd 02_movie_searcher && go mod vendor
run_02:
	cd 02_movie_searcher && go run app.go 
run_03:
	go run 03_code_refactor/app.go
run_04:
	go run 04_logic_anagram/app.go
