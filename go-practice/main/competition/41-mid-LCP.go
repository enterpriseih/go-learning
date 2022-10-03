package competition

func flipChess(g []string) (ans int) {
	// 其实像是N皇后一样，需要判断行、列、对角线
	// 黑色的棋子放一次，然后不断判断是不是有白棋可以翻转，并且统计最多的数量
	ans = 0
	for i, row:= range g {
		for j:=0;j<len(row); j++{
			if row[j] == '.' {
				g[i]= row[0:j]+"X"+row[j+1:]
			}
			oneAns := countReverseWhiteChess(g)
			if oneAns > ans {
				ans = oneAns
			}
			//println(fmt.Sprint( g, oneAns))
			g[i] = row
		}
	}
	return ans
}

func countReverseWhiteChess(g []string) int{
	// 在与[x,y]连接着的地盘，如果周围存在被包围的白色，将其翻转成黑色
	// 可以判断是不是每个位置的白色可以进行翻转即可,每一个翻转，都需要重新check一次

	copyG := make([]string, len(g))
	copy(copyG, g)

	ans := 0
	for {
		oneReverse := 0
		for i, _:= range g {
			for j:=0;j<len(g[i]); j++{
				if g[i][j] == 'O' && judgeCanReverse(g, i,j) {
					//println(i,j)
					g[i]= g[i][0:j]+"X"+g[i][j+1:]
					oneReverse++
				}
			}
		}
		ans += oneReverse
		if oneReverse == 0 {
			break
		}
	}

	copy(g, copyG)
	return ans
}

func judgeCanReverse(g []string, i int, j int) bool{
	// 列
	for x1:=i-1; x1>=0; x1--{
		if g[x1][j]=='O' {
			continue
		}else if g[x1][j]=='X' {
			for x2:=i+1; x2<len(g);x2++{
				if g[x2][j]=='O' {
					continue
				} else if g[x2][j]=='X' {
					return true
				} else {
					break
				}
			}
		} else {
			break
		}
	}

	// 行
	for y1:=j-1; y1>=0; y1--{
		if g[i][y1]=='O' {
			continue
		}else if g[i][y1]=='X' {
			for y2:=j+1; y2<len(g[i]);y2++{
				if g[i][y2]=='O' {
					continue
				} else if g[i][y2]=='X' {
					return true
				} else {
					break
				}
			}
		} else {
			break
		}
	}

	// 对角线【左上-右下】x--，y--; x++,y++
	for m:=1; i-m>=0 && j-m>=0; m++ {
		if g[i-m][j-m]=='O' {
			continue
		}else if g[i-m][j-m]=='X' {
			for n:=1; i+n<len(g) && j+n<len(g[i]); n++{
				if g[i+n][j+n]=='O' {
					continue
				} else if g[i+n][j+n]=='X' {
					return true
				} else {
					break
				}
			}
		} else {
			break
		}
	}

	// 对角线【右上-左下】x--,y++; x++,y--
	for m:=1; i-m>=0 && j+m<len(g[i]); m++ {
		if g[i-m][j+m]=='O' {
			continue
		}else if g[i-m][j+m]=='X' {
			for n:=1; i+n<len(g) && j-n>=0; n++{
				if g[i+n][j-n]=='O' {
					continue
				} else if g[i+n][j-n]=='X' {
					return true
				} else {
					break
				}
			}
		} else {
			break
		}
	}
	return false
}

func TestFlipChess() {
	println(flipChess([]string{".......",".......",".......","X......",".O.....","..O....","....OOX"}))
}