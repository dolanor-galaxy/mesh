.PHONEY: test

info:
	@echo "Task     Does"
	@echo "--------|-------------------------------"
	@echo "deps     installs dependencies (if on macos)"
	@echo "test     runs tests"
	@echo "start    runs the app in development mode"
	@echo "clean    cleans up doco and dist files"
	@echo "build    makes a binary"
	@echo "install  installs the app"
	@echo "docs     creates some documentation"

clean:
# go clean -i
	rm -rf doco
	rm -rf dist

test:
	go test ./...

start:
	CGO_ENABLED=1 go run cmd/mesh/main.go

build: clean
	mkdir dist
	CGO_ENABLED=1 go build -o dist/mesh cmd/mesh/main.go
	cp -R assets dist/assets

wasm: clean
	mkdir dist
	GOOS=js GOARCH=wasm go build -o dist/test.wasm cmd/mesh/main.go

deps:
	# SDL 2
	brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
	# SDL Bindings
	# go get -v github.com/veandco/go-sdl2/sdl
	# go get -v github.com/veandco/go-sdl2/img
	# go get -v github.com/veandco/go-sdl2/mix
	# go get -v github.com/veandco/go-sdl2/ttf
	# go get -v github.com/veandco/go-sdl2/gfx
	go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}
	# OpenGL
	go get -u github.com/go-gl/gl/v{3.2,3.3,4.1,4.2,4.3,4.4,4.5,4.6}-{core,compatibility}/gl
	# OpenGL ES is not supported on macos (only ios)
	# go get -u github.com/go-gl/gl/v3.1/gles2
	go get -u github.com/go-gl/gl/v2.1/gl
	go get -v github.com/chsc/gogl/gl33

install:
	go install

docs: clean
	mkdir doco
	godoc -analysis type,pointer -html ./internal/algebra > ./doco/algebra.html
	godoc -analysis type,pointer -html ./internal/geometry > ./doco/geometry.html
	godoc -analysis type,pointer -html ./internal/render > ./doco/render.html
