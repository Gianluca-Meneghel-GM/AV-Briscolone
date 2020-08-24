package row

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

func Hash(rowPipe chan Data, resultPipe chan Data, wg *sync.WaitGroup) {
	defer wg.Done()

	for inputData := range rowPipe {
		hashedRow, err := bcrypt.GenerateFromPassword([]byte(inputData.OriginalText), 6)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		resultPipe <- Data{
			RowNumber:    inputData.RowNumber,
			OriginalText: inputData.OriginalText,
			Hash:         hashedRow,
		}
	}

}
