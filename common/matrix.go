

package common

import (
	"math"
	"fmt"
)

// @brief Matrix type is private and information about Matrix columns rows and datas
type Matrix struct {
	rows, cols int
	data []float64
} 

// normal matrix norm
type Normal_matrix_type int
const(
	FrobeniusN Normal_matrix_type = iota
	Norm_1
	InfinityNm
	SpectralN
)

// Part of Matrix's functions
func NewMatrix(rows, cols int) *Matrix {
    return &Matrix{rows, cols, make([]float64, rows*cols)}
}

// 矩阵加法
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

// 矩阵减法
func Matrix_sub(m1 *Matrix,m2 *Matrix) *Matrix {
	if(m1.rows != m2.rows)||(m1.cols != m2.cols){
		panic(fmt.Sprintf("Matrix_sub fatal \nm1.rows:%d ,m1.cols:%d\nm2.rows:%d,m2.rows:%d ",m1.rows,m1.cols,m2.rows,m2.cols))
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

// 矩阵数乘
func Matrix_scalar_mul(coefficient float64,m *Matrix)*Matrix{
	res := NewMatrix(m.rows,m.cols)
	for i:=0;i<m.rows;i++{
		for j:=0;j<m.cols;j++{
			res.Set(i,j,m.Get(i,j)*coefficient)
		}
	}
	return res
}

// 矩阵乘法
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
// 对矩阵中每个元素都进行特殊的函数处理
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

// 矩阵的模
func CalculateMatrixNorm(M *Matrix,Type Normal_matrix_type) float64 {
	switch Type {
    case FrobeniusN:
        sum := 0.0
        for i:=0;i<M.rows;i++{
            for j:=0;j<M.cols;j++{
                sum += M.Get(i,j)*M.Get(i,j)
            }
        }
        return math.Sqrt(sum)
    case Norm_1:
        sum := 0.0
		maxsum := 0.0
        for i:=0;i<M.rows;i++{
            for j:=0;j<M.cols;j++{
                sum += math.Abs(M.Get(i,j))
            }
			if sum > maxsum {
                maxsum = sum
            }
        }
        return maxsum
	// TODO: 完成各个范式的代码
    default:
		sum := 0.0
        for i:=0;i<M.rows;i++{
            for j:=0;j<M.cols;j++{
                sum += M.Get(i,j)*M.Get(i,j)
            }
        }
        return math.Sqrt(sum)
	}
}


// diagonal matrix cols must be 1
func ToDiagonalMatrix(M *Matrix) *Matrix {
	total_len := M.rows * M.cols
	res := NewMatrix(total_len, total_len)
	for i:=0;i<M.cols;i++{
		res.Set(i,i,M.Get(0,i))
	}
	return res
}

func FillDiagonalMatrix(M *Matrix) *Matrix {
	res := NewMatrix(M.rows, M.cols)
	for i := 0; i < M.rows; i++ {
		for j := 0; j < M.cols; j++{
			res.Set(i,j,M.Get(i,i))
		}
	}
	return res
}

func CombineMatrixByCol(M1 *Matrix, M2*Matrix)*Matrix {
	if M1.cols!= M2.cols {
        panic(fmt.Sprintf("Matrix_sub fatal \nm1.rows:%d ,m1.cols:%d\nm2.rows:%d,m2.cols:%d ",M1.rows,M1.cols,M2.rows,M2.cols))
    } else {
        res := NewMatrix(M1.rows + M2.rows, M1.cols)
        for i:=0;i<M1.rows;i++{
            for j:=0;j<M1.cols;j++{
                res.Set(i,j,M1.Get(i,j))
            }
        }
        for i:=0;i<M2.rows;i++{
            for j:=0;j<M2.cols;j++{
                res.Set(i+M1.rows,j,M2.Get(i,j))
            }
        }
        return res
    }
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