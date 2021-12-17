clean:
	@rm -rf build

FMT_PATHS = ./

fmt-check:
	@unformatted=$$(gofmt -l $(FMT_PATHS)); [ -z "$$unformatted" ] && exit 0; echo "Unformatted:"; for fn in $$unformatted; do echo "  $$fn"; done; exit 1

smoke-test:
	@mkdir -p build
	tinygo build --target microbit -o ./build/test.hex -size short ./examples/epd/
	@md5sum ./build/test.hex
	tinygo build --target pybadge -o ./build/test.hex -size short ./examples/pybadge/
	@md5sum ./build/test.hex

test: clean fmt-check smoke-test
	go test
