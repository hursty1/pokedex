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
	TextInput string
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
		TextInput: "",
		OldState: oldState,
		exitChan: exitChan,
	}
	fmt.Print("Pokedex > ")
	return cli, nil
}

func (c *CLI) Input() (string, error) {
	b, _ := c.Reader.ReadByte()
	switch b {
	case 3: 
		term.Restore(int(os.Stdin.Fd()), c.OldState)
		// os.Exit(0)
		c.exitChan <- struct{}{}
	case 13: //enter
		fmt.Println()
		fmt.Print("\r\033[K") 
		c.History = append(c.History, c.TextInput)
		c.HistoryIndex = len(c.History)
		return c.TextInput, nil
	case 9: //tab
		fmt.Println("Tab")
	case 27:
		b1, _ := c.Reader.ReadByte()
		b2, _ := c.Reader.ReadByte()
		if b1 == 91 {
			switch b2 {
			case 65:
				if c.HistoryIndex > 0 {
					c.HistoryIndex --
					c.TextInput = c.History[c.HistoryIndex]
					fmt.Print("\r\033[K") // clear line
					fmt.Print("Pokedex > " + c.TextInput)
				}
			case 66: // Down
            if c.HistoryIndex < len(c.History)-1 {
                c.HistoryIndex++
                c.TextInput = c.History[c.HistoryIndex]
            } else {
                c.HistoryIndex = len(c.History)
                c.TextInput = ""
            }
            fmt.Print("\r\033[K")
            fmt.Print("Pokedex > " + c.TextInput)
			}
		}		
		return "",nil
	case 127:
		if len(c.TextInput) > 0 {
			c.TextInput = c.TextInput[:len(c.TextInput)-1]
			fmt.Print("\rPokedex > " + c.TextInput + " \b")
		}
	default:
		c.TextInput += string(b)
		fmt.Print(string(b))
	}

	return "",nil
}

func (c *CLI) ResetInput() {
	fmt.Print("\r\033[K") // reset the line of text
	fmt.Print("Pokedex > ")
}

func (c *CLI) CleanUp() {
	term.Restore(int(os.Stdin.Fd()), c.OldState)
	fmt.Println("\nExiting... ")
}