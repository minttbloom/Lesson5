package main

import (
	"fmt"
	"math"
	"reflect"
)

type figure interface {
	area() float64
	perimeter() float64
	log(method string, parametres ...interface{})
	getHistory() []string
}

type allFigures struct {
	figures []figure
}

func (a *allFigures) addFigure(f figure) *allFigures {
	a.figures = append(a.figures, f)
	return a
}

func (b *allFigures) areaSum() float64 {
	sum := 0.0
	for _, v := range b.figures {
		sum += float64(v.area())
	}
	return float64(sum)
}

func (c *allFigures) perSum() float64 {
	perim := 0.0
	for _, v := range c.figures {
		perim += float64(v.perimeter())
	}
	return float64(perim)
}

func (a *allFigures) figuresCount() int {
	lenght := len(a.figures)
	return lenght
}

func (a *allFigures) callHistory() {
	for _, figure := range a.figures {
		printLog(figure.getHistory())
		fmt.Println("---")
	}
}

type square struct {
	squareSide int
	history    []string
}

func (s *square) log(method string, parametres ...interface{}) {
	row := ""
	for i, v := range parametres {
		row += fmt.Sprint(v)
		if i < len(parametres)-1 {
			row += ", "
		}
	}
	outputRow := fmt.Sprintf("%s(%s)", method, row)
	s.history = append(s.history, outputRow)
}

func (s *square) getHistory() []string {
	return s.history
}

func (s *square) xz(a ...string) {
	fmt.Println(s.history)
}

func newSquare(squareSide int) *square {
	return &square{squareSide: squareSide}
}

func (s *square) area() float64 {
	s.log("area", reflect.TypeOf(s), s.squareSide)
	return float64(s.squareSide * s.squareSide)
}

func (s *square) perimeter() float64 {
	s.log("perimeter", reflect.TypeOf(s), s.squareSide)
	return float64(s.squareSide * 4)
}

type circle struct {
	circleRadius float64
	history      []string
}

func (c *circle) log(method string, parametres ...interface{}) {
	row := ""
	for i, v := range parametres {
		row += fmt.Sprint(v)
		if i < len(parametres)-1 {
			row += ", "
		}
	}
	outputRow := fmt.Sprintf("%s(%s)", method, row)
	c.history = append(c.history, outputRow)
}

func (c *circle) getHistory() []string {
	return c.history
}

func newCircle(circleRadius float64) *circle {
	return &circle{circleRadius: circleRadius}
}

func (c *circle) area() float64 {
	c.log("area", reflect.TypeOf(c), c.circleRadius)
	return math.Pi * math.Pow(c.circleRadius, 2)
}

func (c *circle) perimeter() float64 {
	c.log("perimeter", reflect.TypeOf(c), c.circleRadius)
	return 2 * math.Pi * c.circleRadius
}

type parallelogram struct {
	parTopBottomSides int
	parAnotherSide    int
	history           []string
}

func (p *parallelogram) log(method string, parametres ...interface{}) {
	row := ""
	for i, v := range parametres {
		row += fmt.Sprint(v)
		if i < len(parametres)-1 {
			row += ", "
		}
	}
	outputRow := fmt.Sprintf("%s(%s)", method, row)
	p.history = append(p.history, outputRow)
}

func (p *parallelogram) getHistory() []string {
	return p.history
}

func newParallelogram(parTopBottomSides, parAnotherSide int) *parallelogram {
	return &parallelogram{parTopBottomSides: parTopBottomSides, parAnotherSide: parAnotherSide}
}

func (p *parallelogram) area() float64 {
	p.log("area", reflect.TypeOf(p), p.parTopBottomSides, p.parAnotherSide)
	return float64(p.parTopBottomSides * p.parAnotherSide)
}

func (p *parallelogram) perimeter() float64 {
	p.log("perimeter", reflect.TypeOf(p), p.parTopBottomSides, p.parAnotherSide)
	return float64((p.parTopBottomSides + p.parAnotherSide) * 2)
}

type rectangle struct {
	recLength int
	recWidth  int
	history   []string
}

func (r *rectangle) log(method string, parametres ...interface{}) {
	row := ""
	for i, v := range parametres {
		row += fmt.Sprint(v)
		if i < len(parametres)-1 {
			row += ", "
		}
	}
	outputRow := fmt.Sprintf("%s(%s)", method, row)
	r.history = append(r.history, outputRow)
}

func (r *rectangle) getHistory() []string {
	return r.history
}

func newRectangle(recWidth, recLength int) *rectangle {
	return &rectangle{recLength: recWidth, recWidth: recLength}
}

func (r *rectangle) area() float64 {
	r.log("area", reflect.TypeOf(r), r.recLength, r.recWidth)
	return float64(r.recLength * r.recWidth)
}

func (r *rectangle) perimeter() float64 {
	r.log("perimeter", reflect.TypeOf(r), r.recLength, r.recWidth)
	return float64((r.recWidth + r.recLength) * 2)
}

type cilinder struct {
	*circle
	cilHeight float64
}

func newCilinder(circle *circle, cilHeight float64) *cilinder {
	return &cilinder{circle, cilHeight}
}

func (c *cilinder) area() float64 {
	c.log("area", reflect.TypeOf(c), c)
	return float64(2 * math.Pi * c.circleRadius * (c.cilHeight + c.circleRadius))
}

func (c *cilinder) perimeter() float64 {
	c.log("perimeter", reflect.TypeOf(c), c)
	diametr := 2 * c.circleRadius
	return float64((2 * diametr)) + (2 * c.cilHeight)
}

func (c *cilinder) volume() float64 {
	return float64(math.Pi * math.Pow(c.circleRadius, 2) * c.cilHeight)
}

func printLog(history []string) {
	for _, v := range history {
		fmt.Println(v)
	}
}

func main() {

	fmt.Println("------------------------Added first figure in slice------------------------")
	all := allFigures{}
	all.addFigure(&square{squareSide: 15})

	for _, figure := range all.figures {
		fmt.Printf("%s %s%v %s%v\n", reflect.TypeOf(figure), "area:", figure.area(), "perimeter:", figure.perimeter())
	}

	fmt.Println(all.figuresCount())

	fmt.Println("------------------------Added all figures in slice------------------------")
	nCircle := &circle{circleRadius: 10}
	all.addFigure(nCircle)
	all.addFigure(&parallelogram{parTopBottomSides: 15, parAnotherSide: 10})
	all.addFigure(&rectangle{recLength: 15, recWidth: 10})

	for _, figure := range all.figures {
		fmt.Printf("%s %s%v %s%v\n", reflect.TypeOf(figure), "area:", figure.area(), "perimeter:", figure.perimeter())
	}

	fmt.Println(all.figuresCount())

	fmt.Println("------------------------History-------------------------------------------")
	all.callHistory()

	fmt.Println("------------------------All sums and perimeters---------------------------")
	fmt.Printf("%s %v\n", "All area sum:", all.areaSum())
	fmt.Printf("%s %v\n", "All perimeter sum:", all.perSum())

	fmt.Println("------------------------Cylinder-----------------------------------------")
	nCylinder := newCilinder(nCircle, 10)
	fmt.Printf("Cylinder volume: %v \nCylinder area: %v \nCylinder perimeter: %v \n", nCylinder.volume(), nCylinder.area(), nCylinder.perimeter())

}
