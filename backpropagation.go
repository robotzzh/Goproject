/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-14 15:44:14
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-15 12:31:42
 * @FilePath: /Goproject/Complex_learning/backpropagation.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main 

import (
	"fmt"
	"github.com/zzh/Goproject/backpropagation/common"
)

type D_state int
const (
	D_Sigmoid_1 D_state = iota
	D_Sigmoid_2
)

const Learning_ratio float64 = 0.3

func Linear_Change(Input *common.Matrix,Weight *common.Matrix) *common.Matrix {
	OneMatrix := common.NewMatrix(1,1)
	OneMatrix.Set(0,0,1)
	Input = common.AugmentMatrix(Input,OneMatrix)
	return common.Matrix_mul(Input,Weight)
}

func Sigmoid_Change(Input *common.Matrix) *common.Matrix {
	return common.Complex_Func(Input,common.Sigmoid)
}

func Frontprop(Input *common.Matrix,Weight *common.Matrix) *common.Matrix {
	T1 := Linear_Change(Input,Weight)
	return Sigmoid_Change(T1)
}

func Loss(Input *common.Matrix,Output *common.Matrix) float64 {
	O1 := common.Matrix_sub(Input,Output)
	O2 := common.Matrix_mul(O1,O1.T())
	return common.Matrix_scalar_mul(0.5,O2).Get(0,0)
}

// 返回更新后的权重矩阵
func Backprop(Input *common.Matrix,
			Weight *common.Matrix,
			Output *common.Matrix,
			TestoutPut *common.Matrix) *common.Matrix {
				
				D_E_TestoutPut := common.Matrix_sub(Output,TestoutPut) // 计算误差函数的矩阵值

				Z1 := Linear_Change(Input,Weight) // 计算输出层Z矩阵

				
				D_E_Z := common.Complex_Func(Z1,common.Derivative_Sigmoid) // Derivative_sigmoid Z1
				E_Z1 := common.ToDiagonalMatrix(D_E_Z) // 将Z矩阵转化为对角矩阵
				E_SIM := common.Matrix_mul(D_E_TestoutPut,E_Z1) // Z矩阵和误差函数想乘
				

				E2_SIM := common.ToDiagonalMatrix(E_SIM) // 将结果再度升纬
				E_I := common.ToDiagonalMatrix(Input) // 将input升纬度
				E_I2 := common.FillDiagonalMatrix(E_I)
				D_Weight := common.Matrix_mul(E_I2,E2_SIM) //得到目标矩阵
			
				D_1 := common.CombineMatrixByCol(D_Weight,E_SIM)
				D2 := common.Matrix_scalar_mul(Learning_ratio,D_1)
				D3 := common.Matrix_add(Weight,D2)
				return D3
}

func main(){
	fmt.Println("The project is about backpropagation")
	Input := common.NewMatrix(1,2)
	Output := common.NewMatrix(1,2)
	Weight := common.NewMatrix(3,2)
	

	Input.Set(0,0,0.05)
	Input.Set(0,1,0.1)

	Output.Set(0,0,0.99)
	Output.Set(0,1,0.99)

	Weight.Set(0,0,0.15)
	Weight.Set(0,1,0.2)
	Weight.Set(1,0,0.15)
	Weight.Set(1,1,0.2)
	Weight.Set(2,0,0.35)
	Weight.Set(2,1,0.35)
	
	O1:=Frontprop(Input, Weight)
	//fmt.Println(O1)
	fmt.Println(Loss(O1,Output))
	//fmt.Println(Weight)
	//fmt.Println(Backprop(Input,Weight,Output,O1))

	for i := 0;i<1000;i++{
		O1:=Frontprop(Input, Weight)
		Weight = Backprop(Input,Weight,Output,O1)
		fmt.Println(Loss(O1,Output))
	}
	fmt.Println(Weight)
}