precision mediump float;

struct DirectionalLight {
  vec3 direction;
  vec3 color;
};

varying vec3 v_normal;
varying vec3 vertPos;
varying vec2 v_texcoord;

uniform int mode;
uniform vec3 u_ambientLightIntensity;
uniform DirectionalLight u_sun;
uniform sampler2D u_texture;

vec3 lightPos = u_sun.direction;          // vec3(20.0, 20.0, 20.0);
vec3 diffuseColor = u_sun.color;          // vec3(1.0, 1.0, 1.0);
vec3 specColor = u_ambientLightIntensity; // vec3(1.0, 1.0, 1.0);

void main() {
  vec3 normal = normalize(v_normal);
  vec3 lightDir = normalize(lightPos - vertPos);

  float lambertian = max(dot(lightDir, normal), 0.0);
  float specular = 0.0;

  if(lambertian > 0.0) {

    vec3 reflectDir = reflect(-lightDir, normal);
    vec3 viewDir = normalize(-vertPos);

    float specAngle = max(dot(reflectDir, viewDir), 0.0);
    specular = pow(specAngle, 4.0);

    // the exponent controls the shininess (try mode 2)
    if(mode == 2)  specular = pow(specAngle, 16.0);

    // according to the rendering equation we would need to multiply
    // with the the "lambertian", but this has little visual effect
    if(mode == 3) specular *= lambertian;

    // switch to mode 4 to turn off the specular component
    if(mode == 4) specular *= 0.0;
  }

  vec4 texel = texture2D(u_texture, v_texcoord);
  gl_FragColor = vec4(
    texel.rgb * (lambertian*diffuseColor + specular*specColor),
    texel.a
  );
}
