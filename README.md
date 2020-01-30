# go-migrate

This is a simple script that helps keeps track of migrations.

NOT PRODUCTION READY! 

Install:
 ```bash
go get github.com/CommoDor64/go-migrate
cd go-migrate
go build
 ```
 
Usage:
 ```bash
./go-migrate -dir <migration_files_dir> -envfile <env_file_location> -dburl <connection_string_DB_in_env_file>
```

Notes:

- Works only with postgres-like dialect DB engines  
- Only Briefly tested and serves as a pet project


