/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-14 16:33:37
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-14 21:40:30
 * @FilePath: /Goproject/backpropagation/common/Matrix/Matrix.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

package Matrix

type matrix struct {
	rows, cols int
	data []float64
} 

func NewMatrix(rows, cols int) *matrix {
    return &matrix{rows, cols, make([]float64, rows*cols)}
}

func (m *matrix) Set(row int, col int, val float64) {
    m.data[row*m.cols + col] = val
}

func (m *matrix) Get(row, col int) float64 {
    return m.data[row*m.cols + col]
}


func Matrix_add(m1 *matrix,m2 *matrix) *matrix {
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

func Matrix_sub(m1 *matrix,m2 *matrix) *matrix {
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

func Matrix_scalar_mul(coefficient float64,m *matrix){
	for i:=0;i<m.rows;i++{
		for j:=0;j<m.cols;j++{
			m.Set(i,j,m.Get(i,j)*coefficient)
		}
	}
}

func Matrix_mul(m1 *matrix,m2 *matrix) *matrix{
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
