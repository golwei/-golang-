# -golang-

帅
车
马
炮

棋盘

type 马 struct{
当前位置:x,y
移动()↑:（x，y）{x+2，y+1}
移动()→:(x,y){x+1,y+2}
}

package main
import "fmt"
func main(){
for y:=0;y<12;y++{
	for x:=0;x<9;x++{
fmt.Printf("*") 
}
fmt.Println()
}
 
}
