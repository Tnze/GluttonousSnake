#version 330 core

in float fragColor;

out vec4 outputColor;

void main() {
    const vec4 color1 = vec4(0.0,0.0,1,1.0);
    vec2 temp = gl_PointCoord - vec2(0.5);
    float f = dot(temp, temp);
    if (f>0.25)
        discard;
    if (fragColor==0)
        discard;
        //outputColor = vec4(0,0,0,1);
    else if (fragColor>0)
        outputColor =mix( vec4(0.5,fragColor*0.2,1,1),color1,smoothstep(0.1,0.5,f));
    else
        outputColor = vec4(1,0,0,1);
}