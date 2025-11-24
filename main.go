//go run main.go

package main

import (
	"fmt"
	"net/http"
	"os"
)

const targetFile = "/tmp/testpwd.txt"

func main() {

	http.HandleFunc("/attack/write", handleWriteAttack) // /attack/write으로 접속하면 핸들러 함수 실행
	http.HandleFunc("/attack/read", handleReadAttack)

	fmt.Println("공격대기중. . .\n")
	if err := http.ListenAndServe(":8008", nil); err != nil { // 8008포트로 서버 오픈
		panic(err)
	}
}

func handleWriteAttack(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("write 공격 발생\n")

	f, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		errMSG := fmt.Sprintf("파일 열기 실패: %v", err)
		http.Error(w, errMSG, 500)
		return
	}
	defer f.Close()

	_, err = f.WriteString("공격 발생\n")
	if err != nil {
		http.Error(w, fmt.Sprintf("쓰기 에러 발생 : %v", err), 500)
		return
	}

	msg := fmt.Sprintf("공격 성공 : %s 파일에 쓰기 완료", targetFile)
	fmt.Printf(msg)
	w.Write([]byte(msg + "\n")) //클라이언트에게 응답 전송
}

func handleReadAttack(w http.ResponseWriter, r *http.Request) { //w 응답, r 요청 인자
	fmt.Printf("read 공격 발생\n")

	bytes, err := os.ReadFile(targetFile)
	if err != nil {
		errMSG := fmt.Sprintf("파일 읽기 실패: %v", err)
		http.Error(w, errMSG, 500)
		return
	}
	msg := fmt.Sprintf("공격 성공 : %s 파일 내용 읽기 완료\n내용: %s", targetFile, string(bytes))
	fmt.Printf(msg)
	w.Write([]byte(msg + "\n")) //클라이언트에게 응답전송
}
