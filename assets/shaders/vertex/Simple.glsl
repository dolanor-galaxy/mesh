#version 120

attribute vec3 Pos;       // verts
attribute vec3 Color;     // color
attribute vec2 TexCoord;  // texture coords (on the vert)
attribute vec3 Normal;    // which way is out (on the vert)
attribute vec3 Tangent;   // has to do with light refraction

varying vec3 v_color;
varying vec2 v_texcoord;  // Texture coords (on the image)
varying vec3 v_normal;    // Texture normal
varying vec3 v_tangent;

uniform mat4 uWorld;      // model to world
uniform mat4 uView;       // view
uniform mat4 uProj;       // projection

void main() {
  // v_normal = mat3(uWorld) * normalize(Normal.xzy);
  // // v_normal = vec3( mWorld * vec4(normal.x, normal.z, normal.y, 0.0) );
  // // v_texcoord = texCoord;
  
  v_color = Color;
  v_texcoord = TexCoord;
  v_normal = Normal;
  v_tangent = Tangent;
  
  gl_Position = uProj * uView * uWorld * vec4(Pos, 1.0);
}