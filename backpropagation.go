/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-14 15:44:14
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-14 23:49:33
 * @FilePath: /Goproject/Complex_learning/backpropagation.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main 

import (
	"fmt"
	"github.com/zzh/Goproject/backpropagation/common"
)

func frontprop(Input *common.Matrix,Weight *common.Matrix) *common.Matrix {
	OneMatrix := common.NewMatrix(1,1)
	OneMatrix.Set(0,0,1)
	Input = common.AugmentMatrix(Input,OneMatrix)
	return common.Complex_Func(common.Matrix_mul(Input,Weight),common.Sigmoid)
}

func Loss(Input *common.Matrix,Output *common.Matrix) float64{
	O1 := common.Matrix_sub(Input,Output)
	return common.CalculateMatrixNorm(O1,common.FrobeniusN)
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
	
	O1:=frontprop(Input, Weight)
	fmt.Println(O1)
	fmt.Println(Loss(Input,Output))
}