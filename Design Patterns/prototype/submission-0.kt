interface Shape {
    fun clone(): Shape
}

class Rectangle(private val width: Int, private val height: Int) : Shape {
    fun getWidth(): Int = width
    fun getHeight(): Int = height

    override fun clone(): Shape {
        return Rectangle(width, height)
    }
}

class Square(private val length: Int) : Shape {
    fun getLength(): Int = length

    override fun clone(): Shape {
        return Square(length)
    }
}

class Test {
    fun cloneShapes(shapes: List<Shape>): List<Shape> {
        return shapes.map { it.clone() }
    }
}
