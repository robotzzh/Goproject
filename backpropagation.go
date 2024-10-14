/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-14 15:44:14
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-14 21:47:37
 * @FilePath: /Goproject/Complex_learning/backpropagation.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main 

import (
	"fmt"
	"github.com/zzh/Goproject/backpropagation/common/Matrix"
)


func main(){
	fmt.Println("The project is about backpropagation")
	m1 := Matrix.NewMatrix(2,2)
	m2 := Matrix.NewMatrix(2,2)
	m1.Set(0,0,1)
	m1.Set(0,1,2)
	m1.Set(1,0,3)
	m1.Set(1,1,4)
	m2.Set(0,0,1)
	m2.Set(0,1,2)
	m2.Set(1,0,3)
	m2.Set(1,1,4)

	fmt.Println(m1)
	m1.T()
	fmt.Println(m1)
	fmt.Println(Matrix.Matrix_mul(m1,m2))
	Matrix.Matrix_scalar_mul(0.0,m1)
	fmt.Println(m1)
}