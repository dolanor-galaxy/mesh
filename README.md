# Mesh

* Install go
* Install SDL2

    brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config

* Install go bindings for SDL2
    To get the bindings, type:
    go get -v github.com/veandco/go-sdl2/sdl
    go get -v github.com/veandco/go-sdl2/img
    go get -v github.com/veandco/go-sdl2/mix
    go get -v github.com/veandco/go-sdl2/ttf
    go get -v github.com/veandco/go-sdl2/gfx

    or type this if you use Bash terminal:
    go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}

* Install binding for opengl (3.3)
    go get -u github.com/go-gl/gl/v{3.2,3.3,4.1,4.2,4.3,4.4,4.5,4.6}-{core,compatibility}/gl
    go get -u github.com/go-gl/gl/v3.1/gles2
    go get -u github.com/go-gl/gl/v2.1/gl
