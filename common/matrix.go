

package common

// @brief Matrix type is private and information about Matrix columns rows and datas
type Matrix struct {
	rows, cols int
	data []float64
} 

// Part of Matrix's functions
func NewMatrix(rows, cols int) *Matrix {
    return &Matrix{rows, cols, make([]float64, rows*cols)}
}

func Matrix_add(m1 *Matrix,m2 *Matrix) *Matrix {
	if(m1.rows != m2.rows)||(m1.cols != m2.cols){
		panic("Matrix_add fatal")
	}else{
		res := NewMatrix(m1.rows,m1.cols)
		for i:=0;i<m1.rows;i++{
			for j:=0;j<m1.cols;j++{
				res.Set(i,j,m1.Get(i,j)+m2.Get(i,j))
			}
		}
		return res
	}
}

func Matrix_sub(m1 *Matrix,m2 *Matrix) *Matrix {
	if(m1.rows != m2.rows)||(m1.cols != m2.cols){
		panic("Matrix_add fatal")
	}else{
		res := NewMatrix(m1.rows,m1.cols)
		for i:=0;i<m1.rows;i++{
			for j:=0;j<m1.cols;j++{
				res.Set(i,j,m1.Get(i,j)-m2.Get(i,j))
			}
		}
		return res
	}
}

func Matrix_scalar_mul(coefficient float64,m *Matrix){
	for i:=0;i<m.rows;i++{
		for j:=0;j<m.cols;j++{
			m.Set(i,j,m.Get(i,j)*coefficient)
		}
	}
}

func Matrix_mul(m1 *Matrix,m2 *Matrix) *Matrix{
	if m1.cols != m2.rows{
		panic("Matrix_mul fatal")
	}else{
		res := NewMatrix(m1.rows,m2.cols)
		for i:=0;i<res.rows;i++{
			for j:=0;j<res.cols;j++{
				sum := 0.0
				for k:=0;k<m1.cols;k++{
                    sum += m1.Get(i,k) * m2.Get(k,j)
                }
				res.Set(i,j,sum)
			}
		}
		return res
	}
}

type operation func(float64) float64

func Complex_Func(M *Matrix ,CFUNC operation) *Matrix{
	res := NewMatrix(M.rows,M.cols)
    for i:=0;i<res.rows;i++{
        for j:=0;j<res.cols;j++{
            res.Set(i,j,CFUNC(M.Get(i,j)))
        }
    }
    return res
}

// 增广矩阵生成
func AugmentMatrix(M *Matrix, vec *Matrix) *Matrix{
	res := NewMatrix(M.rows, M.cols + vec.cols)
    for i:=0;i<M.rows;i++{
        for j:=0;j<M.cols;j++{
            res.Set(i,j,M.Get(i,j))
        }
        for j:=0;j<vec.cols;j++{
            res.Set(i,j+M.cols,vec.Get(i,j))
        }
    }
    return res
}

// Part of Matrix's methods
// @brief Matrix transformation
func (M *Matrix) T() *Matrix { 
	M1 := NewMatrix(M.cols,M.rows)
	for i:=0;i<M.rows;i++{
        for j:=0;j<M.cols;j++{
            M1.Set(j,i,M.Get(i,j))
        }
    }
	return M1
}

func (m *Matrix) Set(row int, col int, val float64) {
    m.data[row*m.cols + col] = val
}

func (m *Matrix) Get(row, col int) float64 {
    return m.data[row*m.cols + col]
}

func (m *Matrix) GetCols() int {
	return m.cols
}

func (m *Matrix) GetRows() int {
	return m.rows
}