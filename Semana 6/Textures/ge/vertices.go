package ge

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//CreateVAO ...
func CreateVAO(vertices []mgl32.Vec3, textureCoord []mgl32.Vec2) uint32 {

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)

	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	// copy vertices data into VBO (it needs to be bound first)
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4*3, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	if len(textureCoord) > 0 {
		var TBO uint32
		gl.GenBuffers(1, &TBO)
		gl.BindBuffer(gl.ARRAY_BUFFER, TBO)
		gl.BufferData(gl.ARRAY_BUFFER, len(textureCoord)*4*3, gl.Ptr(textureCoord), gl.STATIC_DRAW)
		gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 2*4, gl.PtrOffset(0))
		gl.EnableVertexAttribArray(1)
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	}
	gl.BindVertexArray(0)

	return VAO
}

//Mul defines multiplication of 2 vert3
func Mul(v1 mgl32.Vec3, v2 mgl32.Vec3) mgl32.Vec3 {
	return mgl32.Vec3{v1.X() * v2.X(), v1.Y() * v2.Y(), v1.Z() * v2.Z()}
}

//Translate defines sum of vert3 array and a ver3
func Translate(vertices []mgl32.Vec3, vertex mgl32.Vec3) (translated []mgl32.Vec3) {
	for _, ver := range vertices {
		translated = append(translated, ver.Add(vertex))
	}
	return
}

//Transform defines multiplication of vert3 array and a ver3
func Transform(vertices []mgl32.Vec3, vertex mgl32.Vec3) (translated []mgl32.Vec3) {
	for _, ver := range vertices {
		translated = append(translated, Mul(ver, vertex))
	}
	return
}
