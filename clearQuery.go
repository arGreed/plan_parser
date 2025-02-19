package main

/*
!	Пометка:
!		функции будут некорректно работать, если комментарий находится внутри текстовых строк (select 'aaaa--;' - будет просто select, что не верно)
!	В планах исправить, но потом
*/

import (
	"errors"
	"strings"
)

// ? Функция очистки от всех комментариев в запросах
func clearQueries(text *string, commSign string) error {
	deleteSingleLineComments(text, commSign)
	err := deleteMultiLineComments(text)
	if err != nil {
		return err
	}
	return nil
}

// ? Функция удаления однострочных комментариев
func deleteSingleLineComments(text *string, commSign string) {
	var clearLine string
	lines := strings.Split(*text, "\n")
	for i, line := range lines {
		clearLine = strings.Split(line, commSign)[0]
		lines[i] = clearLine
	}
	*text = strings.Join(lines, "\n")
}

// ? Функция удаления многострочных комментариев
func deleteMultiLineComments(text *string) error {
	var result strings.Builder
	input := *text
	commentFlag := false
	len := len(input)
	for i := 0; i < len; i++ {
		if commentFlag {
			if i+1 < len && input[i] == '*' && input[i+1] == '/' {
				commentFlag = false
				i++
			}
		} else {
			if i+1 < len && input[i] == '/' && input[i+1] == '*' {
				commentFlag = true
				i++
			} else {
				result.WriteByte(input[i])
			}
		}
	}
	if commentFlag {
		return errors.New("в запросе не закрыт многострочный комментарий")
	}

	*text = result.String()
	return nil
}
