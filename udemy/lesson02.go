package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//hello()
	//func_if()
	//func_for()
	//func_range()
	//func_switch()
	//func_defer()
	//func_log()
	//func_err()
	func_panic()
}

func thirdPartyConnectDB() {
	// panicで強制終了することができる
	// panicではなく、errで返すようなコードを推奨
	panic("Unable to connect database!")
}

func save() {
	// deferでこの関数が終わる前にrecoverすることで、プログラムの強制終了を防ぐことができる。
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func func_panic() {
	save()
	fmt.Println("OK!")
}

func func_err() {
	file, err := os.Open("lesson01.go")

	if err != nil {
		log.Fatalln("ERR!")
	}

	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("ERR")
	}
	fmt.Println(count, string(data))

	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("ERRoR")
	}
}

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func func_log() {
	LoggingSettings("test.log")

	_, err := os.Open("nonononono")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging!")
	log.Fatalln("error!!") // コードが終了するので注意
	log.Println("ok!")
}

// 遅れて実行される。関数がスタックされて、関数の最後の、deferの後ろから実行されるので注意
// ファイルを閉じるときなどに利用される
func func_defer() {
	defer fmt.Println("world")
	defer fmt.Println("world2")

	fmt.Println("hello")
	fmt.Println("hello2")
}

func getOsName() string {
	return "mac"
}

func func_switch() {
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("MAC!")
	case "windows":
		fmt.Println("WINDOWS!")
	default:
		fmt.Println("DEFAULT!", os)
	}
}

func func_range() {
	l := []string{"python", "go", "java"}

	for i, v := range l {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func func_for() {
	// 真ん中だけでも良い。なにもないと無限ループ
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue", i)
			continue
		}
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}

func hello() {
	fmt.Println("hello world")
}

func func_if() {
	num := 5
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}
}
