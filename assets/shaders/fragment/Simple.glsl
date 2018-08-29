#version 120

// precision mediump float;

// struct DirectionalLight {
//   vec3 direction;
//   vec3 color;
// };

varying vec2 v_texcoord;
varying vec3 v_normal;
varying vec3 v_color;

// uniform vec3 u_ambientLightIntensity;
// uniform DirectionalLight u_sun;
// uniform vec3 u_diffuseColor;
// uniform sampler2D u_texture;         // Main texture

void main() {
  // float cosTheta = clamp(dot(v_normal, normalize(u_sun.direction)), .0, 1.);
  // vec3 lightIntensity = u_ambientLightIntensity + u_sun.color * cosTheta;

  // vec4 texel = texture2D(u_texture, v_texcoord);

  // gl_FragColor = vec4(u_diffuseColor.rgb * texel.rgb * lightIntensity, texel.a);

  gl_FragColor = vec4(v_color, 1.0);
  // gl_FragColor = vec4(1.0, .5, .5, 1);

  // if(gl_FragColor.a < 0.1)
  //    discard;
}
