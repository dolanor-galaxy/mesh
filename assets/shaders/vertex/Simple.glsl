precision mediump float;

attribute vec3 position;  // verts
attribute vec2 texCoord;  // texture coords (on the vert)
attribute vec3 normal;    // which way is out (on the vert)
attribute vec3 tangent;   // has to do with light refraction

varying vec3 v_normal;    // Texture normal
varying vec2 v_texcoord;  // Texture coords (on the image)

uniform mat4 mWorld;      // model to world
uniform mat4 mView;       // view
uniform mat4 mProj;       // projection

void main() {
  v_normal = mat3(mWorld) * normalize(normal.xzy);
  // v_normal = vec3( mWorld * vec4(normal.x, normal.z, normal.y, 0.0) );
  v_texcoord = texCoord;

  gl_Position = mProj * mView * mWorld * vec4(position, 1.0);
}