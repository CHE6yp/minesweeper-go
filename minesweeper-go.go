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

	var mc int
	fmt.Println("Set the mine count")
	fmt.Fscanln(os.Stdin,&mc)
	setMines(&mapi, &mc)
	fmt.Println()
	fmt.Println()
	setFieldShow(&showfield, mapi)
	printSeeF(seeF)
	dead := false
	x,y:=0,0
		fmt.Println("Pick X")
	for dead != true {
		fmt.Fscanln(os.Stdin,&x)
		fmt.Println("Pick Y")
		fmt.Fscanln(os.Stdin,&y)
		dead = openCell(x,y,&showfield,&seeF,mapi)
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

func openCell(x int, y int, sf *[][]int, seeF *[][]string, m [][]int) (bool) {
	if m[x][y]==1{
		(*seeF)[x][y] = "m"
		return true
	}else {
		if (*sf)[x][y]==0 {
			openNeighbourCells(x,y,&*sf, &*seeF)
		}
		(*seeF)[x][y] = strconv.Itoa((*sf)[x][y])
		return false
	}
}

func openCellSafe(x int, y int, sf *[][]int, seeF *[][]string) {
	if (*seeF)[x][y]!="*"{
		return
	}
	(*seeF)[x][y] = strconv.Itoa((*sf)[x][y])
	if (*sf)[x][y]==0 {
		openNeighbourCells(x,y,&*sf,&*seeF)
	}
}

func openNeighbourCells(x int, y int, sf *[][]int, seeF *[][]string) {
 	if x!=0 {
		openCellSafe(x-1,y,&*sf,&*seeF)
		if y!=0 {
			openCellSafe(x-1,y-1,&*sf,&*seeF)
		}
		if y!=len((*sf)[x])-1{
			openCellSafe(x-1,y+1,&*sf,&*seeF)
		}
	}
	if x!=len(*sf)-1 {
		openCellSafe(x+1,y,&*sf,&*seeF)
		if y!=0 {
			openCellSafe(x+1,y-1,&*sf,&*seeF)
		}
		if y!=len((*sf)[x])-1{
			openCellSafe(x+1,y+1,&*sf,&*seeF)
		}
	}
	if y!=0 {
		openCellSafe(x,y-1,&*sf,&*seeF)
	}
	if y!=len((*sf)[x])-1{
		openCellSafe(x,y+1,&*sf,&*seeF)
	}
 }