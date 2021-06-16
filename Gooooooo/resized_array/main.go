package main
 
import (
	"reflect"
	"fmt"
)
func main(){
	arr := [9][9]int{}
	arr[0][1] = 1
	
	srcToSparse(arr)
}

func ShowMatrix(arg interface{}) {

	value := reflect.ValueOf(arg)

	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		rowNum := value.Len()
		colNum := value.Index(0).Len()
		for i := 0; i < rowNum; i++{
			for j := 0; j < colNum; j++{
				fmt.Printf("%v ", value.Index(i).Index(j))
			}
			fmt.Println()
		}
default:
		fmt.Println("非矩阵形式数据")
		return 

	}
}

func srcToSparse (arr [9][9]int) [][3]int {
	var sparse = make([][3]int, 1)
	dataCap := 0

	sparse[0] = [3]int{9,9,dataCap}

	for i, arr := range arr {
		for j, v := range arr {
			if v != 0 {
				dataCap ++
				sparse = append(sparse, [3]int{i,j,v})
			}
		}
	}
	sparse[0][2] = dataCap
	ShowMatrix(sparse)
	return sparse
}