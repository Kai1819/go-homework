package homework02

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
)

var (
	failedQuestions []string
	totalQuestions  int
	mu              sync.Mutex
)

func recordResult(t *testing.T, name string) {
	mu.Lock()
	defer mu.Unlock()
	totalQuestions++
	if t.Failed() {
		failedQuestions = append(failedQuestions, name)
	}
}

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()

	// Print summary
	if totalQuestions > 0 {
		fmt.Println("\n---------------------------------------------------")
		fmt.Printf("Total Questions: %d\n", totalQuestions)
		fmt.Printf("Passed: %d\n", totalQuestions-len(failedQuestions))
		fmt.Printf("Failed: %d\n", len(failedQuestions))

		score := float64(totalQuestions-len(failedQuestions)) / float64(totalQuestions) * 100
		fmt.Printf("Score: %.2f%%\n", score)

		if len(failedQuestions) > 0 {
			fmt.Println("Failed Questions:")
			for _, q := range failedQuestions {
				fmt.Printf("- %s\n", q)
			}
		}
		fmt.Println("---------------------------------------------------")
	}

	os.Exit(code)
}

func TestCheckPoint1(t *testing.T) {
	defer recordResult(t, "TestCheckPoint1")
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"Example 1", 121, 131},
		{"Example 2", -121, -111},
		{"Example 3", 10, 20},
		{"Example 4", 0, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckPoint1(&tt.x)
			if tt.x != tt.want {
				t.Errorf("CheckPoint1() = %v, want %v", tt.x, tt.want)
			}
		})
	}
}

func TestCheckPoint2(t *testing.T) {
	defer recordResult(t, "TestCheckPoint2")
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, []int{2, 4, 6, 8, 10}},
		{"Example 2", []int{-1, -2, -3}, []int{-2, -4, -6}},
		{"Example 3", []int{0, 5, 10}, []int{0, 10, 20}},
		{"Example 4", []int{7}, []int{14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckPoint2(&tt.slice)
			if !reflect.DeepEqual(tt.slice, tt.want) {
				t.Errorf("CheckPoint2() = %v, want %v", tt.slice, tt.want)
			}
		})
	}
}

func TestCheckGoroutine1(t *testing.T) {
	defer recordResult(t, "TestCheckGoroutine1")

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	CheckGoroutine1()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	tests := []struct {
		name string
		want string
	}{
		{"奇数1", "奇数: 1"},
		{"奇数3", "奇数: 3"},
		{"奇数5", "奇数: 5"},
		{"奇数7", "奇数: 7"},
		{"奇数9", "奇数: 9"},
		{"偶数2", "偶数: 2"},
		{"偶数4", "偶数: 4"},
		{"偶数6", "偶数: 6"},
		{"偶数8", "偶数: 8"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(output, tt.want) {
				t.Errorf("输出中未找到: %v, 实际输出:\n%v", tt.want, output)
			}
		})
	}
}

func TestCheckGoroutine2(t *testing.T) {
	CheckGoroutine2()
}

func TestCheckStruct1(t *testing.T) {
	CheckStruct1()
}

func TestCheckStruct2(t *testing.T) {
	CheckStruct2()
}

func TestCheckChannel1(t *testing.T) {
	CheckChannel1()
}

func TestCheckChannel2(t *testing.T) {
	CheckChannel2()
}

func TestCheckLock1(t *testing.T) {
	CheckLock1()
}

func TestCheckLock2(t *testing.T) {
	CheckLock2()
}