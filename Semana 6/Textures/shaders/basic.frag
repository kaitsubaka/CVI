#version 410 core

in vec2 TexCoord;

out vec4 color;
uniform vec3 objectColor;
uniform vec3 lightColor;

uniform sampler2D material;

void main()
{
    // mix the two textures together (texture1 is colored with "ourColor")
    if (textureSize(material, 0).x > 1){
    color = texture(material, TexCoord)* vec4(objectColor*lightColor, 1.0f);
    }
    else {
    color = vec4(objectColor*lightColor, 1.0f);
    }
}
