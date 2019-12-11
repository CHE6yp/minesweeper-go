package main

import ("fmt"
	"os"
	"math/rand"
)

func main(){

	var s int
	fmt.Println("Set the map size")
	fmt.Fscanln(os.Stdin,&s)
	var mapi = make([][]int,s)
	for i:=0;i<s;i++{
		mapi[i] = make([]int,s) 
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
	// fmt.Println(mapi)
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