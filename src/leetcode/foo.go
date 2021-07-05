package leetcode

type Foo struct {
	ch1 chan int
	ch2 chan int
}

func (f Foo) first() {
	print("first")
	f.ch1 <- 1 // 存入任何值均可
}

func (f Foo) second() {
	<-f.ch1
	print("second")
	f.ch2 <- 2 // 存入任何值均可

}

func (f Foo) third() {
	<-f.ch2
	print("third")
}
