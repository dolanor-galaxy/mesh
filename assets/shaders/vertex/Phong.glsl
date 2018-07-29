precision mediump float;

attribute vec3 position;
attribute vec2 texCoord;
attribute vec3 normal;
attribute vec3 tangent;

uniform mat4 mWorld, mView, mProj;

varying vec2 v_texcoord;
varying vec3 v_normal;
varying vec3 vertPos;

void main(){
    v_normal = (mWorld * vec4(normal, 0.0)).xyz;
    v_texcoord = texCoord;

    vec4 vertPos4 = mView * vec4(position, 1.0);
    vertPos = vec3(vertPos4) / vertPos4.w;

    gl_Position = mProj * mView * mWorld * vec4(position, 1.0);
}