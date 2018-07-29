precision lowp float;

struct DirectionalLight {
  vec3 direction;
  vec3 color;
};

uniform vec3 u_ambientLightIntensity;
uniform DirectionalLight u_sun;
uniform lowp sampler2D u_texture;
uniform lowp sampler2D b_texture;
uniform vec2 u_texU;       // texture
uniform vec2 u_texV;       // scale, origin

varying vec2 v_texcoord;
varying vec3 v_normal;

void main(void) {
  float cosTheta = clamp(dot(v_normal, normalize(u_sun.direction)), .0, 1.);
  vec3 lightIntensity = u_ambientLightIntensity + u_sun.color * cosTheta;

  ////////////////////////////////////////////////////////////////
  vec2 texCoord = v_texcoord.xy;
  vec2 paramU   = u_texU.xy;
  vec2 paramV   = u_texV.xy;
  // TODO: Modifying v_textcoord should be in vertex shader
  texCoord.x = fract(texCoord.x) * paramU.x + paramU.y;
  texCoord.y = fract(texCoord.y) * paramV.x + paramV.y;
  vec4 texel = texture2D(u_texture, texCoord);

  // // RGB of our normal map
  // // Local normal, in tangent space
  // // vec3 normalMap = normalize(texture2D(b_texture, texCoord).rgb * 2.0 - 1.0);

  gl_FragColor = vec4(texel.rgb * lightIntensity, texel.a);
}