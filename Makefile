DEFAULT_GOAL = run

run:
	CONFIG_PATH=config/path.yaml go run cmd/main.go
migrator:
	CONFIG_PATH=config/path.yaml go run cmd/migrator/migrator.go