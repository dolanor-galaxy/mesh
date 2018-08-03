package render

import "github.com/therohans/mesh/algebra"

const (
	// IllumColorOnAmbientOff Color on and Ambient off
	IllumColorOnAmbientOff = iota
	// IllumColorOnAmbientOn Color on and Ambient on
	IllumColorOnAmbientOn
	// IllumHighlightOn Highlight on
	IllumHighlightOn
	// IllumReflectionOn Reflection on and Ray trace on
	IllumReflectionOn
	// IllumTransparencyGlass Transparency: Glass on, Reflection: Ray trace on
	IllumTransparencyGlass
	// IllumReflectionFresnel Reflection: Fresnel on and Ray trace on
	IllumReflectionFresnel
	// IllumTransparencyNoFresnel Transparency: Refraction on, Reflection: Fresnel off and Ray trace on
	IllumTransparencyNoFresnel
	// IllumTransparency Transparency: Refraction on, Reflection: Fresnel on and Ray trace on
	IllumTransparency
	// IllumReflectionNoRay Reflection on and Ray trace off
	IllumReflectionNoRay
	// IllumTransparencyNoRay Transparency: Glass on, Reflection: Ray trace off
	IllumTransparencyNoRay
	// IllumCastsShadows Casts shadows onto invisible surfaces
	IllumCastsShadows
)

// Material a skin for an entity
type Material struct {
	Name string
	// AmbientColor 'Ka' blender does not have per-object ambient color.
	// AmbientColor algebra.Vector
	// DiffuseColor 'Kd' is the most instinctive meaning of the color of an object.
	DiffuseColor algebra.Vector
	// SpecularColor 'Ks' the color of highlights on a shiny surface
	// SpecularColor algebra.Vector
	// SpecularColorWeight 'Ns'
	// SpecularColorWeight float32
	// Transparent 'd' 1=opaque (or Tr - inverted d)
	Transparent float32
	// Illumination specifies the illumination model to use in the material.
	Illumination uint8

	// Textures
	// AmbientTextureName string
	// AmbientTexture map_Ka
	// AmbientTexture Texture

	DiffuseTextureName string
	// DiffuseTexture map_Kd
	DiffuseTexture Texture

	// SpecularColorTextureName string
	// SpecularColorTexture map_Ks
	// SpecularColorTexture Texture

	// SpecularHighlightTextureName string
	// SpecularHighlightTexture map_Ns
	// SpecularHighlightTexture Texture

	// AlphaTextureName string
	// AlphaTexture map_d
	// AlphaTexture Texture

	// BumpTextureName string
	// BumpTexture map_bump or bump
	// BumpTexture Texture

	// TextureOrigin the origin of the texture
	TextureOrigin algebra.Vector
	// TextureOrigin the scale of the nexture
	TextureScale algebra.Vector

	// The shader will operate on the textures
	Shader Shader
}
