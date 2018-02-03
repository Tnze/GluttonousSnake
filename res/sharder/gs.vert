#version 330

in vec2 vert;
in vec3 color;

out vec3 fragColor;

void main() {
    gl_Position = vec4(vert, 0, 1);
    gl_PointSize =30;
    fragColor = color;
}