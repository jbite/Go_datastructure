package main

import "fmt"

type PazzleMap struct {
	Row int
	Col int
	Map [8][7]int
}

func NewPazzleMap() *PazzleMap {
	p := new(PazzleMap)
	p.Col = len(p.Map)
	p.Row = len(p.Map[0])
	return p
}

func (p *PazzleMap) ShowPazzle() {
	for i := 0; i < len(p.Map); i++ {
		for j := 0; j < len(p.Map[i]); j++ {
			fmt.Print(p.Map[i][j], " ")
		}
		fmt.Println()
	}
}

//編寫一個函數 讓老鼠找路
//傳入地圖及起始點
func (p *PazzleMap) SetWay(i int, j int) bool {
	if p.Map[6][5] == 2 {
		fmt.Println("找到路了")
		return true
	} else {
		//繼續找
		if p.Map[i][j] == 0 { //如果給定的點不是牆而且沒有探測過
			//假設此點可以通 先將其設為2 但需要探測 上下左右
			//改成下右上左
			p.Map[i][j] = 2
			if p.SetWay(i+1, j) {
				return true
			} else if p.SetWay(i, j+1) {
				return true
			} else if p.SetWay(i-1, j) {
				return true
			} else if p.SetWay(i, j-1) {
				return true
			} else {
				p.Map[i][j] = 3
				return false
			}

		} else { //這個點不能探測
			return false
		}
	}

}

func main() {
	//規則
	//1. 如果元素的值為1 就是牆
	//2. 如果元素為0 表示還沒走過的格子
	//3. 如果元素為2 表示是通路
	//4. 如果元素為3 表示走過 但是是死路
	p := NewPazzleMap()

	//先把地圖的最上和最下設置為1
	for i := 0; i < 7; i++ {
		p.Map[0][i] = 1
		p.Map[7][i] = 1
	}
	//把最左和最後設為1
	for i := 0; i < 8; i++ {
		p.Map[i][0] = 1
		p.Map[i][6] = 1
	}
	p.Map[3][1] = 1
	p.Map[3][2] = 1
	p.ShowPazzle()
	p.SetWay(1, 1)
	p.ShowPazzle()
}
