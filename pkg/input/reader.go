package input

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func ReadSingleCharacter(buffer *rune) error {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stderr.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		return err
	}

	*buffer = char
	return nil
}

func PromptSingleCharacter(prompt string) (rune, error) {
	var char rune
	print(prompt)
	err := ReadSingleCharacter(&char)
	return char, err
}

func ReadString(buffer *[]byte) error {
	fmt.Scan()
	return nil
}
