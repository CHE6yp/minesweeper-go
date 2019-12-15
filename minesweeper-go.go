package main

import ("fmt"
	"os"
	"math/rand"
	//"strconv"
)
const (
        InfoColor    = "\033[1;34m%v\033[0m"
        NoticeColor  = "\033[1;36m%v\033[0m"
        WarningColor = "\033[1;33m%v\033[0m"
        ErrorColor   = "\033[1;31m%v\033[0m"
        DebugColor   = "\033[0;36m%v\033[0m"
)
func main(){

	var s int
	fmt.Println("Set the map size")
	fmt.Fscanln(os.Stdin,&s)
	var innerMap,outerMap = make([][]string,s), make([][]string,s)
	for i:=0;i<s;i++{
		innerMap[i], outerMap[i] = make([]string,s), make([]string,s) 
		for key := range outerMap[i]{
			outerMap[i][key] = "*"
		}
	}

	var mineCount int
	fmt.Println("Set the mine count")
	fmt.Fscanln(os.Stdin,&mineCount)
	setMines(&innerMap, &mineCount)

	printOuterMap(outerMap)

	dead := false
	x,y:=0,0
	for dead != true {
		fmt.Println("Pick X")
		fmt.Fscanln(os.Stdin,&x)
		fmt.Println("Pick Y")
		fmt.Fscanln(os.Stdin,&y)
		dead = openCell(x,y,innerMap,&outerMap)
		printOuterMap(outerMap)
		if dead {
			fmt.Println("You are ded!")
		}
	}
}

func printOuterMap(innerMap [][]string){
	fmt.Println()
	/////////
	fmt.Print(" ");
	for key, _ := range innerMap {
		fmt.Print(" ")
		fmt.Printf(NoticeColor,key);
	}
	fmt.Println();
	for key, value := range innerMap {
		fmt.Printf(NoticeColor,key);
		fmt.Println(value);
	}
}

func setMines(m *[][]string, s *int){
	x, y := rand.Intn(len(*m)), rand.Intn(len(*m))
	//x, y := len(*m), cap(*m)
	// fmt.Println(x,y)
	(*m)[x][y] = "m";
	*s--
	if (*s!=0){
		setMines(&*m,&*s)
	}
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

func openCell(x int, y int, innerMap [][]string, outerMap *[][]string) (bool) {
	if innerMap[x][y]=="m"{
		(*outerMap)[x][y] = "m"
		return true
	}else {
		if innerMap[x][y]=="0" {
			for _, value := range neighbourCells(x,y,&*outerMap){
				openCellSafe(x,y,innerMap,&value)
			}
		}
		(*outerMap)[x][y] = innerMap[x][y]
		return false
	}
}

func openCellSafe(x int, y int, innerMap [][]string, cell *string) {
	if *cell!="*"{
		return
	}
	*cell = (*innerMap)[x][y]
	if (*innerMap)[x][y]==0 {
		for _, value = range neighbourCells(x,y,&*innerMap){
			openCellSafe(x,y,&*innerMap,&*value)
		}
	}
}

func openNeighbourCells(x int, y int, innerMap *[][]int, outerMap *[][]string) {
 	if x!=0 {
		openCellSafe(x-1,y,&*innerMap,&*outerMap)
		if y!=0 {
			openCellSafe(x-1,y-1,&*innerMap,&*outerMap)
		}
		if y!=len((*innerMap)[x])-1{
			openCellSafe(x-1,y+1,&*innerMap,&*outerMap)
		}
	}
	if x!=len(*innerMap)-1 {
		openCellSafe(x+1,y,&*innerMap,&*outerMap)
		if y!=0 {
			openCellSafe(x+1,y-1,&*innerMap,&*outerMap)
		}
		if y!=len((*innerMap)[x])-1{
			openCellSafe(x+1,y+1,&*innerMap,&*outerMap)
		}
	}
	if y!=0 {
		openCellSafe(x,y-1,&*innerMap,&*outerMap)
	}
	if y!=len((*innerMap)[x])-1{
		openCellSafe(x,y+1,&*innerMap,&*outerMap)
	}
 }

 func neighbourCells(x int, y int, array *[][]string) []string{
 	result := make([]string, 0);
 	if x!=0 {
 		result = append(result,(*array)[x-1][y])
		if y!=0 {
			result = append(result,(*array)[x-1][y-1])
		}
		if y!=len((*array)[x])-1{
			result = append(result,(*array)[x-1][y+1])
		}
	}
	if x!=len(*array)-1 {
		result = append(result,(*array)[x+1][y])
		if y!=0 {
			result = append(result,(*array)[x+1][y-1])
		}
		if y!=len((*array)[x])-1{
			result = append(result,(*array)[x+1][y+1])
		}
	}
	if y!=0 {
		result = append(result,(*array)[x][y-1])
	}
	if y!=len((*array)[x])-1{
		result = append(result,(*array)[x][y+1])
	}
	return result;
 }