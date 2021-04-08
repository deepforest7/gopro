package main

import f "fmt"
type jiekou interface {
	add()
	jian()
}
type num1 struct {
	n1 int
	n2 int
}
type num2 struct {
	n1 int
	n2 int
}

func (n *num1)error()string  {
	str := `
	Cannot proceed, the divider is zero.
	jianshu: %d
	beijianshu: %d
`
	return f.Sprintf(str,n.n1,n.n2)
}

func (n *num1)add()int  {
	return n.n1+n.n2
}
func (n *num1)jian()int  {
	return n.n1-n.n2
}

func jian(n1 int,n2 int)(res int,errormsg string){


	if(n1-n2<0){
		date:=num1{
			n1:n1,
			n2:n2,

		}

		errormsg = date.error()
		return
	}else {
		return n1-n2,""

	}

}


func (n *num2)add()int {
	return n.n1+n.n2
}
func (n *num2)jian()int  {
	return n.n1-n.n2
}

func main()  {
	//nu1:=num1{5,2}
	//nu2:=num2{6,3}
	//
	//
	//f.Println("nu1.n1",nu1.n1,"nu1.n2",nu1.n2,"nu2.n1",nu2.n1,"nu2.n2",nu2.n2)
	//f.Println(nu1.add(),
	//nu1.jian(),
	//nu2.add(),
	//nu2.jian())


	if res,errormsg:=jian(5,2);errormsg ==""{
		f.Println(res)
	}
	if _,errormsg:=jian(2,5);errormsg !=""{
		f.Println(errormsg)
	}


}