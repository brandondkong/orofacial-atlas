package input

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"golang.org/x/term"
)

var ErrInterrupted error = errors.New("interrupted")

func ReadSingleCharacter(buffer *string) error {
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
	if char == 0x03 {
		return ErrInterrupted
	}

	*buffer = string(char)
	return nil
}

func PromptSingleCharacter(prompt string) (string, error) {
	var char string
	print(prompt)
	err := ReadSingleCharacter(&char)
	return char, err
}

func PromptBooleanAnswer(prompt string, trueKey string, falseKey string) (bool, error) {
	var char string
	var choice bool

	trueKey = strings.ToLower(trueKey)
	falseKey = strings.ToLower(falseKey)

	print(prompt)
	for (char != trueKey && char != falseKey) {
		err := ReadSingleCharacter(&char)
		if err != nil {
			return false, err
		}
		char = strings.ToLower(char)
		if char == trueKey {
			choice = true
		}

		if char == falseKey {
			choice = false
		}
	}

	if choice {
		println(trueKey)
	} else {
		println(falseKey)
	}

	return choice, nil
}

func PromptYesNoAnswer(prompt string) (bool, error) {
	return PromptBooleanAnswer(prompt, "y", "n")
}

func ReadString(buffer *[]byte) error {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadLine()
	if err != nil {
		return err
	}

	*buffer = char
	return nil
}

func PromptString(prompt string) (string, error) {
	var buffer []byte
	print(prompt)

	err := ReadString(&buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
