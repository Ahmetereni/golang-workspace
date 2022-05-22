package main
import "fmt"
import "time"

func hello(){
	numbers:=[]int{0,1,2,3,4,5,6,7,8,9}
	start:=time.Now()
	for _,number :=range numbers{
		if (number%2==0){
			fmt.Println("hi",number)
		}
	}
	t := time.Now()
	end :=t.Sub(start)
	println(end.Milliseconds())
}

