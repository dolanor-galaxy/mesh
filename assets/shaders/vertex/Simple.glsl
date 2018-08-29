#version 120

// precision mediump float;

attribute vec3 Pos;
attribute vec3 Color;
attribute vec2 TexCoord;
attribute vec3 Normal;
attribute vec3 Tangent;

// attribute vec3 position;  // verts
// attribute vec2 texCoord;  // texture coords (on the vert)
// attribute vec3 normal;    // which way is out (on the vert)
// attribute vec3 tangent;   // has to do with light refraction

varying vec3 v_normal;    // Texture normal
varying vec2 v_texcoord;  // Texture coords (on the image)
varying vec3 v_color;

uniform mat4 uWorld;      // model to world
uniform mat4 uView;       // view
uniform mat4 uProj;       // projection

void main() {
  // v_normal = mat3(uWorld) * normalize(Normal.xzy);
  // // v_normal = vec3( mWorld * vec4(normal.x, normal.z, normal.y, 0.0) );
  // // v_texcoord = texCoord;
  // v_texcoord = TexCoord;

  // gl_Position = mProj * mView * mWorld * vec4(position, 1.0);
  

  // v_color = iuWorld[0].xyz;
  v_color = Color;
  // gl_Position = vec4(Pos, 1.0);
  // gl_Position = vec4(Pos, 1.0) * uWorld;
  gl_Position = uProj * uView * uWorld * vec4(Pos, 1.0);
}