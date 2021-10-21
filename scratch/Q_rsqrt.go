// Fast Inverse Square Root Algorithm
// Wikipedia : https://en.wikipedia.org/wiki/Fast_inverse_square_root
// Youtube : https://www.youtube.com/watch?v=p8u_k2LIZyo
// Description : "The algorithm is best known for its implementation in 1999 in the source code of Quake III Arena, a first-person shooter video game that made heavy use of 3D graphics."
// Playground : https://play.golang.org/p/C0KNGXBb14p
// Documentations :
// 高速逆平方根(fast inverse square root)のアルゴリズム解説 - https://lipoyang.hatenablog.com/entry/2021/02/06/194619
// 30のプログラミング言語でFast inverse square rootを実装してみました！- https://itchyny.hatenablog.com/entry/2016/07/25/100000
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

func Q_rsqrt(x float32) float32 {
	const threehalfs = float32(1.5)
	// 32bitの float と int をそれぞれポインタでキャスティングが必要になるので、 unsafe.Pointer が必要になります。32ビット同士の整数と単精度浮動少数点数のポインターキャスティングを行っています。
	i := *(*int32)(unsafe.Pointer(&x))
	// y = ​1⁄√x を求解するために、logを取ってテイラー展開を行いそれを一次の精度で打ち切ります。これによって得られた近似式に対して、得られた数というのが 0x5f3759df になるのです。
	i = 0x5f3759df - i>>1
	y := *(*float32)(unsafe.Pointer(&i))
	// ここで、ニュートン法の反復を一度行っています。ここで、最大発生する計算誤差についても0.175%程度しか起きないため、そこそこの精度を保ったまま、一度の反復で逆平方根の近似値を得ることができます。これ以上の精度が欲しい場合は同じ計算を再度行うことによって、二回目の反復を実行します。
	return y * (threehalfs - 0.5*x*y*y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if x, err := strconv.ParseFloat(scanner.Text(), 32); err == nil {
			fmt.Println(Q_rsqrt(float32(x)))
		}
	}
}
