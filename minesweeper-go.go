package main

import ("fmt"
	"os"
	"math/rand"
	"strconv"
)

func main(){

	var s int
	fmt.Println("Set the map size")
	fmt.Fscanln(os.Stdin,&s)
	var mapi, showfield = make([][]int,s), make([][]int,s)
	var seeF = make([][]string,s)
	for i:=0;i<s;i++{
		mapi[i] = make([]int,s) 
		showfield[i] = make([]int,s) 
		seeF[i] = make([]string,s) 
		for key := range seeF[i]{
			seeF[i][key] = "*"
		}
	}
	fmt.Println(mapi)


	var mc int
	fmt.Println("Set the mine count")
	fmt.Fscanln(os.Stdin,&mc)
	setMines(&mapi, &mc)
	for key, value := range mapi {
		key = key
		fmt.Println(value)
	}
	fmt.Println()
	fmt.Println()
	setFieldShow(&showfield, mapi)
	for key, value := range showfield {
		key = key
		fmt.Println(value)
	}
	printSeeF(seeF)
	dead := false
	x,y:=0,0
	for dead != true {
		fmt.Println("Pick X")
		fmt.Fscanln(os.Stdin,&x)
		fmt.Println("Pick Y")
		fmt.Fscanln(os.Stdin,&y)
		seeF[x][y], dead = openCell(x,y,showfield,mapi)
		printSeeF(seeF)
		if dead {
			fmt.Println("You are ded!")
		}

	}


}

func printSeeF(sf [][]string){
	fmt.Println()
	fmt.Println()
	/////////
	for key, value := range sf {
		key = key
		fmt.Println(value)
	}
}

func setMines(m *[][]int, s *int){
	x, y := rand.Intn(len(*m)), rand.Intn(len(*m))
	//x, y := len(*m), cap(*m)
	// fmt.Println(x,y)
	(*m)[x][y] = 1;
	*s--
	if (*s!=0){
		setMines(&*m,&*s)
	}
}

func setFieldShow(sf *[][]int, m [][]int) {
	for key, value := range m {
		for key2, value2 := range value {
			value2 = value2
			(*sf)[key][key2]=setCellShow(m,key,key2)
		}
	}
	// return m
}

func setCellShow(m [][]int, x int, y int) int {
	count := 0
	if x!=0 {
		if m[x-1][y]==1{
			count++ //8
		}
		if y!=0 {
			if m[x-1][y-1]==1{
				count++ //1
			}
		}
		if y!=len(m[x])-1{
			if m[x-1][y+1]==1{
				count++ //7
			}
		}
	}
	if x!=len(m)-1 {
		if m[x+1][y]==1{
			count++ //4
		}
		if y!=0 {
			if m[x+1][y-1]==1{
				count++ //3
			}
		}
		if y!=len(m[x])-1{
			if m[x+1][y+1]==1{
				count++ //5
			}
		}
	}
	if y!=0 {
		if m[x][y-1] ==1 {
			count++ //2
		}
	}
	if y!=len(m[x])-1{
		if m[x][y+1]==1{
			count++ //6
		}
	}

	return count
}

func openCell(x int, y int, sf [][]int, m [][]int) (string, bool) {
	if m[x][y]==1{
		return "m",true
	}else {
		return strconv.Itoa(sf[x][y]), false
	}
}