package draw

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	//"github.com/go-gl/mathgl/mgl32"
	"fmt"
	gs "gluttonous_snake"
	"io/ioutil"
	"runtime"
	"strings"
)

//read sharder source from file
func readSharderSource(dir string) (source string, err error) {
	bsource, err := ioutil.ReadFile(dir)
	source = string(bsource)
	source += "\x00"
	return
}

var snakeDrawable [gs.Weight][gs.Hight]uint32

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

const windowW, windowH = 1000, windowW / 16 * 9

//OpenWindow 打开一个窗口来绘制贪吃蛇
func OpenWindow(getsnake func() *gs.Snake, reciveKey func(direction int)) error {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		return err
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(windowW, windowH, "贪吃蛇", nil, nil)
	if err != nil {
		return err
	}
	window.MakeContextCurrent()
	//设置回掉函数
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press { //“动作”如果是按下
			switch key {
			case glfw.KeyUp:
				reciveKey(1)
			case glfw.KeyDown:
				reciveKey(2)
			case glfw.KeyLeft:
				reciveKey(3)
			case glfw.KeyRight:
				reciveKey(4)
			}
		}
	})
	// 初始化 GL
	if err := gl.Init(); err != nil {
		panic(err)
	}

	//打印OpenGL版本
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
	//从文件读取着色器源码
	vertexShaderSource, err := readSharderSource(`../res/sharder/gs.vert`)
	if err != nil {
		return fmt.Errorf("read vertex sharder source faild: %v", err)
	}
	fragmentShaderSource, err := readSharderSource(`../res/sharder/gs.frag`)
	if err != nil {
		return fmt.Errorf("read fragment sharder source faild: %v", err)
	}
	//编译、链接着色器
	program, err := newProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		panic(err)
	}
	gl.ClearColor(0, 0.25, 0.2, 1)

	for !window.ShouldClose() {
		s := getsnake()
		//计算点的位置
		var points [gs.Weight * gs.Hight * 5]float32
		const cubeW float32 = 2.0 / gs.Weight
		const cubeH float32 = 2.0 / gs.Hight
		for i := 0; i < gs.Weight; i++ {
			for j := 0; j < gs.Hight; j++ {
				var r, g, b float32
				if s.GetBlock([2]int{i, j}) > 0 {
					r, g, b = 1, 1, 1
				} else if s.GetBlock([2]int{i, j}) < 0 {
					r, g, b = 1, 0, 0
				}
				point := []float32{cubeW/2 + float32(i)*cubeW - 1, -cubeH/2 - float32(j)*cubeH + 1, r, g, b}
				//fmt.Printf("kkkk	%d   %d 	%d\n", i, j, gs.Weight*j+i*5)//查看下标是否计算正确
				p := points[(j*gs.Weight+i)*5:]
				copy(p, point)
			}
		}
		//fmt.Println(points)//查看顶点数组是否正确生成
		gl.Clear(gl.COLOR_BUFFER_BIT) //清除画布
		gl.UseProgram(program)

		//创建顶点数组对象
		var vao uint32
		gl.GenVertexArrays(1, &vao)
		gl.BindVertexArray(vao)
		//创建顶点缓存对象
		var vbo uint32
		gl.GenBuffers(1, &vbo)
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(points)*4, gl.Ptr(points[:]), gl.STATIC_DRAW)
		//取着色器传入值
		vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
		gl.EnableVertexAttribArray(vertAttrib)
		gl.VertexAttribPointer(vertAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
		colorAttrib := uint32(gl.GetAttribLocation(program, gl.Str("color\x00")))
		gl.EnableVertexAttribArray(colorAttrib)
		gl.VertexAttribPointer(colorAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(2*4))

		gl.BindVertexArray(vao)
		gl.Enable(gl.PROGRAM_POINT_SIZE)
		gl.DrawArrays(gl.POINTS, 0, gs.Weight*gs.Hight)
		window.SwapBuffers()
		glfw.PollEvents()
	}
	return nil
}
