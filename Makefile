build:
	./script/build.sh

make run: build
	./script/run.sh

help:
	@echo "build   - Build the web binary and docker image"
	@echo "run     - Run the web server locally using the build docker image"
	@echo "help    - Show help"