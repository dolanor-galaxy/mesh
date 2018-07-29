#version 120

// vec4 outColor;
varying vec3 fragmentColor;

void main()
{
	vec4 outColor = vec4(fragmentColor, 1.0);
	gl_FragColor = outColor;
}