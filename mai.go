// author : Kimsoenghyeon
// title : sakurajimamaisay
// description : do not distribution this code

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

//go:embed mai.txt
var asciiArt string

func termWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("터미널 크기를 가져올 수 없습니다:", err)
		return -1
	}
	return width
}

func printAsciiArt() error {
	scanner := bufio.NewScanner(strings.NewReader(asciiArt))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}

func echo(cmd *cobra.Command, args []string) {
	var s1 string = args[0]

	var n1 int = runewidth.StringWidth(s1)

	var t1 int = termWidth()

	for i := 1; i <= n1; i++ {
		if i < t1 {
			fmt.Printf("-")
		}
	}
	fmt.Println()

	fmt.Println(s1)

	for i := 1; i <= n1; i++ {
		if i < t1 {
			fmt.Printf("-")
		}
	}
	fmt.Println()

	fmt.Println("\t\t\t\t\\")
	fmt.Println("\t\t\t\t \\")

	//ascii art print
	if err := printAsciiArt(); err != nil {
		fmt.Fprintf(os.Stderr, "파일 읽기 오류: %v\n", err)
		os.Exit(1)
	}

}

func main() {
	rootCmd := &cobra.Command{
		Use:   "maisay",
		Short: "최종 목표는 마이가 말하는 것 !!",
	}

	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "print pong",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pong")
		},
	}

	echoCmd := &cobra.Command{
		Use:   "echo",
		Short: "echo all input",
		Args:  cobra.ExactArgs(1),
		Run:   echo,
	}

	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(echoCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
