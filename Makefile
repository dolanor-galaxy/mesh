clean:
	go clean -i
	rm -rf doco
	rm -rf dist

test:
	cd algebra; go test
	cd geometry; go test
	cd render; go test

start:
	go run main.go

build: clean
	mkdir dist
	go build -o dist/mesh
	cp -R assets dist/assets

wasm: clean
	mkdir dist
	# go get golang.org/dl/go1.11
	# go1.11 download
	GOOS=js GOARCH=wasm go1.11 build -o dist/test.wasm main.go

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
	godoc -analysis type,pointer -html ./algebra > ./doco/algebra.html
	godoc -analysis type,pointer -html ./geometry > ./doco/geometry.html
	godoc -analysis type,pointer -html ./render > ./doco/render.html
