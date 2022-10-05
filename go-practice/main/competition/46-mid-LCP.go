package competition

func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	// 逆向思维
	// 设置每一个场馆的人数都可以用kx+b表示
	// 可以逆推回去，根据方案
	k0 := []int64{}
	b0 := []int64{}

	k0 = append(k0, 1)
	b0 = append(b0, 0)
	for i:=0;i<len(finalCnt);i++ {
		k0 = append(k0, 0)
		b0 = append(b0, int64(finalCnt[i]))
	}

	edgeMap := map[int][]int{}
	for i:=0;i<len(edges);i++ {
			_,ok0 := edgeMap[edges[i][0]]
			if ok0 {
				edgeMap[edges[i][0]] = append(edgeMap[edges[i][0]], edges[i][1])
			} else {
				edgeMap[edges[i][0]] = []int{edges[i][1]}
			}

			_,ok1 := edgeMap[edges[i][1]]
			if ok1 {
				edgeMap[edges[i][1]] = append(edgeMap[edges[i][1]], edges[i][0])
			} else {
				edgeMap[edges[i][1]] = []int{edges[i][0]}
			}
	}

	for i:= len(plans)-1; i>=0; i-- {
		//println(fmt.Sprint("+v, +v, +v", k0, b0, edgeMap[plans[i][1]]))
		opt(plans[i][0], plans[i][1], k0, b0, edgeMap[plans[i][1]])
	}

	sumK := int64(0)
	sumB := int64(0)
	for i:= 0; i <len(k0); i++ {
		sumK += k0[i]
		sumB += b0[i]
	}

	x := (totalNum-sumB)/sumK

	res:= []int{}
	for i:= 0; i <len(k0); i++ {
		res = append(res, int(x*k0[i]+b0[i]))
	}

	//println(fmt.Sprint("+v", res))
	return res

}

func opt(num int, idx int, k []int64, b []int64, neighbors []int) (){
	if num == 1 {
		k[idx] = 2*k[idx]
		b[idx] = 2*b[idx]
	}
	if num ==2 {
		for i:=0; i < len(neighbors);i++ {
			k[neighbors[i]]-=k[idx]
			b[neighbors[i]]-=b[idx]
		}
	}
	if num == 3 {
		for i:=0; i < len(neighbors);i++ {
			k[neighbors[i]]+=k[idx]
			b[neighbors[i]]+=b[idx]
		}
	}
}


func TestVolunteerDeployment() {
	volunteerDeployment([]int{1,16}, 21, [][]int{{0,1}, {1,2}}, [][]int{{2,1},{1,0},{3,0}})
}