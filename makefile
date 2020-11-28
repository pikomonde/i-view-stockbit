01_run:
	mysql -u root -ppass < 01_simple_database_querying/mysql.sql
02_init:
	mysql -u root -ppass < 02_movie_searcher/files/int_mysql.sql
	cd 02_movie_searcher && go mod vendor
02_build_proto:
	cd 02_movie_searcher/delivery/grpchandler && protoc --go_out=plugins=grpc:moviesearch moviesearch.proto
02_run:
	cd 02_movie_searcher && go run app.go 
02_api_test:
	cd 02_movie_searcher/files && sh http_sample.sh
	cd 02_movie_searcher/files && go run app.go
03_run:
	go run 03_code_refactor/app.go
04_run:
	go run 04_logic_anagram/app.go
