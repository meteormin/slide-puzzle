package main

import (
	"bufio"
	"fmt"
	"github.com/meteormin/slide-puzzle/internal/core"
	"github.com/meteormin/slide-puzzle/internal/counter"
	"github.com/meteormin/slide-puzzle/internal/logger"
	"github.com/meteormin/slide-puzzle/internal/printer"
	"github.com/meteormin/slide-puzzle/internal/timer"
	"golang.org/x/term"
	"os"
	"strconv"
	"strings"
	"time"
)

const shuffleMoves = 1024

func readKey() (byte, error) {
	// stdin을 raw 모드로 전환
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return 0, err
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			panic(err)
		}
	}(int(os.Stdin.Fd()), oldState) // 끝나면 원복

	buf := make([]byte, 1)
	_, err = os.Stdin.Read(buf)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	container, err := core.New(2)
	if err != nil {
		panic(err)
	}

	l := logger.New()
	ll := logger.NewListener(l)
	p := printer.NewPrinter()
	c := counter.New()
	t := timer.New()

	container.AddListener(ll)
	container.AddListener(p)
	container.AddListener(c)
	container.AddListener(t)
	container.ErrorHandler(func(err error) {
		l.Error(err)
	})

	fmt.Println("환영합니다! 슬라이드 퍼즐 게임을 시작합니다.")
	container.Shuffle(shuffleMoves) // 초기 퍼즐 셔플
	for {
		fmt.Print("\033[2J")   // 전체 화면 지움
		fmt.Print("\033[5;1H") // 커서를 5행 1열로 이동

		tiles := container.Snapshot()
		err = p.WriteTiles(tiles)
		if err != nil {
			l.Error(err)
			fmt.Println("⚠️ 퍼즐을 출력하는 중 오류가 발생했습니다.")
			return
		}

		if container.IsSolved() {
			fmt.Println("🎉 퍼즐을 완성했습니다!")
			fmt.Printf("퍼즐 크기: %d", len(tiles))
			fmt.Printf(" | 이동 횟수: %d", c.Count())
			fmt.Printf(" | 소요 시간: %.2fs\n", t.Elapsed().Seconds())
			fmt.Println("[s = 다시 셔플 | r = 크기 변경 리셋 | q = 종료]")
		}

		fmt.Print("[w,k/a,h/s,j/d,l = 이동 | r = 리셋 | q = 종료] > ")
		input, _ := readKey()
		switch input {
		case 'w', 'k':
			container.MoveBy(core.Up)
		case 's', 'j':
			if container.IsSolved() {
				container.Shuffle(shuffleMoves)
			} else {
				container.MoveBy(core.Down)
			}
		case 'a', 'h':
			container.MoveBy(core.Left)
		case 'd', 'l':
			container.MoveBy(core.Right)
		case 'r':
			fmt.Print("새 보드 크기 입력 (예: 3) > ")
			sizeInput, _ := reader.ReadString('\n')
			sizeInput = strings.TrimSpace(sizeInput)
			size, err := strconv.Atoi(sizeInput)
			if err != nil || size < 2 || size > 10 {
				fmt.Println("⚠️ 올바른 숫자를 입력하세요 (2~10).")
				time.Sleep(1 * time.Second)
				continue
			}
			err = container.Reset(size)
			if err != nil {
				l.Error(err)
				fmt.Println("⚠️ 퍼즐을 리셋하는 중 오류가 발생했습니다.")
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Println("✅ 퍼즐 크기를 변경했습니다.")
			container.Shuffle(shuffleMoves)
		case 'q':
			fmt.Println("👋 게임을 종료합니다.")
			return
		default:
			fmt.Println("❌ 유효하지 않은 입력입니다. w/a/s/d/r/q 중 하나를 입력하세요.")
			time.Sleep(1 * time.Second)
		}
	}
}
