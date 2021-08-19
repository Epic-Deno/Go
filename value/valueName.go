/*
 * @Description: It was written by Pony
 * @Version: 2.0
 * @Author: Pony
 * @Date: 2020-08-14 10:18:56
 * @LastEditors: ZHANG ZHEN
 * @LastEditTime: 2021-07-21 18:21:07
 */
package main
import "fmt"
func main() {
	/*声明一个变量初始化*/
	var a = "hello"
	fmt.Println(a) /*hello*/

	/*没有初始化就为零*/
	var b int
	fmt.Println(b) /*0*/

	/*bool 零为 false*/
	var c bool
	fmt.Println(c) /*false*/
}