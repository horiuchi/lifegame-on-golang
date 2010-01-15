package main

import (
	"rand"
	"time"
)

type LifeGame struct {
	board [][]bool
	row, col int
	time int
}

func (p *LifeGame) Init(row, col int) *LifeGame {
	//rand.Seed(0)
	p.row = row
	p.col = col
	p.board = make([][]bool, row)
	for r := 0; r < row; r++ {
		p.board[r] = make([]bool, col)
		for c := 0; c < col; c++ {
			if rand.Float() < 0.3 {
				p.board[r][c] = true
			}
		}
	}
	p.time = 0
	return p
}

func (p *LifeGame) Print() {
	print("time=")
	println(p.time)
	f := func(b bool) (s string) {
		if b {
			s = "[*]"
		} else {
			s = "[ ]"
		}
		return s
	}
	for r := 0; r < p.row; r++ {
		for c := 0; c < p.col; c++ {
			print(f(p.board[r][c]))
		}
		println()
	}
}


func (p *LifeGame) count_now_alive(r, c int) (i int) {
	if r < 0 {
		r += p.row
	}
	if p.row <= r {
		r -= p.row
	}
	if c < 0 {
		c += p.col
	}
	if p.col <= c {
		c -= p.col
	}
	if p.board[r][c] {
		i = 1
	} else {
		i = 0
	}
	return i
}
func (p *LifeGame) is_dead_or_alive(r, c int) (b bool) {
	count := p.count_now_alive(r - 1, c - 1) +
		 p.count_now_alive(r - 1, c    ) +
		 p.count_now_alive(r - 1, c + 1) +
		 p.count_now_alive(r    , c - 1) +
		 p.count_now_alive(r    , c    ) +
		 p.count_now_alive(r    , c + 1) +
		 p.count_now_alive(r + 1, c - 1) +
		 p.count_now_alive(r + 1, c    ) +
		 p.count_now_alive(r + 1, c + 1)
	switch count {
		case 3:
			b = true
		case 4:
			b = p.board[r][c]
		default:
			b = false
	}
	return b
}

func generate_gen(game *LifeGame, ch chan<- *LifeGame) {
	for true {
		next := new(LifeGame).Init(game.row, game.col)
		for r := 0; r < game.row; r++ {
			for c := 0; c < game.col; c++ {
				next.board[r][c] = game.is_dead_or_alive(r, c)
			}
		}
		next.time = game.time + 1
		ch <- next
		game = next
	}
}


func main() {
	game := new(LifeGame).Init(10, 10)
	game.Print()

	ch := make(chan *LifeGame)
	go generate_gen(game, ch)
	for next := range ch {
		next.Print()
		time.Sleep(1e9)
	}
}

