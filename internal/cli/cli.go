package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/term"
)

type CLI struct {
	Reader *bufio.Reader
	HistoryIndex int
	History []string
	input string
	OldState *term.State
	exitChan chan struct{}

}

func Setup() (CLI, error) {

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	// defer term.Restore(int(os.Stdin.Fd()), oldState)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	exitChan := make(chan struct{})
	go func () {
		<-sig
		<-exitChan
		term.Restore(int(os.Stdin.Fd()), oldState)
		fmt.Println("\nExiting... ")
		os.Exit(0)
	}()

	
	reader := bufio.NewReader(os.Stdin)
	// var input string
	history := []string{}
	
	cli := CLI{
		Reader: reader,
		History: history,
		HistoryIndex: -1,
		input: "",
		OldState: oldState,
		exitChan: exitChan,
	}
	fmt.Print("Pokedex > ")
	return cli, nil
}

func (c *CLI) Input() error {
	// var input string
	b, _ := c.Reader.ReadByte()
	switch b {
	case 3: 
		term.Restore(int(os.Stdin.Fd()), c.OldState)
		// os.Exit(0)
		c.exitChan <- struct{}{}
	case 13: //enter
		fmt.Println()
		fmt.Print("\r\033[K") 
		// fmt.Println("You typed:", c.input)
		c.History = append(c.History, c.input)
		c.HistoryIndex = len(c.History)
		c.input = ""
		fmt.Print("\r\033[K") // reset the line of text
		fmt.Print("Pokedex > ")
	case 9: //tab
		fmt.Println("Tab")
	case 27:
		fmt.Println("ArrowKey")
		return nil
	default:
		c.input += string(b)
		fmt.Print(string(b))
	}

	return nil
}

func (c *CLI) CleanUp() {
	term.Restore(int(os.Stdin.Fd()), c.OldState)
	fmt.Println("\nExiting... ")
}