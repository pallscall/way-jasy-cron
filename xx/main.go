package main



func main(){

}

func setZeroes(matrix [][]int)  {
	for i, v := range matrix {
		for j, m := range v {
			if m == 0 {
				for k := 0; k < len(v); k++ {
					matrix[i][k] = -1
				}
				for k := 0; k < len(matrix); k++ {
					matrix[k][j] = -1
				}
			}
		}
	}

	for i, v := range matrix {
		for j, m := range v {
			if m == -1 {
				matrix[i][j] = 0
			}
		}
	}
}