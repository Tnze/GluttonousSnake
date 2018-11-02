#version 330 core

in vec2 vert;
in float color;

out float fragColor;

void main() {
    gl_Position = vec4(vert, 0, 1);
    if (color <= 0)
        gl_PointSize = 35;
    else if (color > 0)
        gl_PointSize = max(40,min(30 + color,60) );
    fragColor = color;
}