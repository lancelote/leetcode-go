package p901

type Pair struct {
	i int
	p int
}

type StockSpanner struct {
	i     int
	stack []Pair
}

func Constructor() StockSpanner {
	return StockSpanner{i: 0, stack: []Pair{}}
}

func (this *StockSpanner) Next(price int) int {
	this.i++

	for len(this.stack) > 0 && this.stack[len(this.stack)-1].p <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}

	var result int

	if len(this.stack) == 0 {
		result = this.i
	} else {
		result = this.i - this.stack[len(this.stack)-1].i
	}

	this.stack = append(this.stack, Pair{i: this.i, p: price})
	return result
}
