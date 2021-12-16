package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var c = newCircle(10.0)

func TestSquareArea(t *testing.T) {
	s := newSquare(2)
	assert.Equal(t, 4.0, s.area())
}

func TestSquarePerimeter(t *testing.T) {
	s := newSquare(2)
	assert.Equal(t, 8.0, s.perimeter())
}

func TestCircleArea(t *testing.T) {
	assert.Equal(t, 314.1592653589793, c.area())
}

func TestCirclePerimeter(t *testing.T) {
	assert.Equal(t, 62.83185307179586, c.perimeter())
}

func TestParallelogramArea(t *testing.T) {
	p := newParallelogram(15, 10)
	assert.Equal(t, 150.0, p.area())
}

func TestParallelogramPerimeter(t *testing.T) {
	p := newParallelogram(15, 10)
	assert.Equal(t, 50.0, p.perimeter())
}

func TestRectangleleArea(t *testing.T) {
	r := newRectangle(15, 10)
	assert.Equal(t, 150.0, r.area())
}

func TestRectanglePerimeter(t *testing.T) {
	r := newRectangle(15, 10)
	assert.Equal(t, 50.0, r.perimeter())
}

func TestCilinderArea(t *testing.T) {
	c := newCilinder(c, 10.0)
	assert.Equal(t, 1256.6370614359173, c.area())
}

func TestCilinderPerimeter(t *testing.T) {
	c := newCilinder(c, 10.0)
	assert.Equal(t, 60.0, c.perimeter())
}

func TestCilinderVolume(t *testing.T) {
	c := newCilinder(c, 10.0)
	assert.Equal(t, 3141.5926535897934, c.volume())
}

func TestAllAreaSum(t *testing.T) {
	figures := allFigures{}
	figures.addFigure(newSquare(15))
	figures.addFigure(newCircle(10.0))
	figures.addFigure(newParallelogram(15.0, 10.0))
	figures.addFigure(newRectangle(15.0, 10.0))
	assert.Equal(t, 839.1592653589794, figures.areaSum())
}

func TestAllPerimeterSum(t *testing.T) {
	figures := allFigures{}
	figures.addFigure(newSquare(15))
	figures.addFigure(newCircle(10.0))
	figures.addFigure(newParallelogram(15.0, 10.0))
	figures.addFigure(newRectangle(15.0, 10.0))
	assert.Equal(t, 222.83185307179588, figures.perSum())
}

func TestFiguresCount(t *testing.T) {
	figures := allFigures{}
	figures.addFigure(newSquare(15))
	figures.addFigure(newCircle(10.0))
	figures.addFigure(newParallelogram(15.0, 10.0))
	figures.addFigure(newRectangle(15.0, 10.0))
	lenght := len(figures.figures)
	assert.Equal(t, 4, lenght)

}
