/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-14 22:48:19
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-15 10:27:00
 * @FilePath: /Goproject/backpropagation/common/function.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
 package common

import (
    "math"
)



func Sigmoid(x float64) float64 {
    return 1 / (1 + math.Exp(-x))
}

func Derivative_Sigmoid(x float64) float64 {
	return Sigmoid(x)*(1-Sigmoid(x))
}