#version 120

varying vec2 v_texcoord;
varying vec3 v_normal;
varying vec3 v_color;

void main() {
  gl_FragColor = vec4(v_color, 1.0);
}
