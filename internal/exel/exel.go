package exel

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"src/golang_testWork2/internal/cache/record"
	"strconv"
	"time"
)

type ExelFile struct {
	filename string
	data     []byte
}

const sheetName = "Sheet1"

func Excel(rec []*record.Record) ExelFile {
	f := excelize.NewFile()
	err := f.SetCellValue(sheetName, "A1", "Ключ")
	if err != nil {
		log.Panicln(err)
		return ExelFile{filename: "", data: nil}
	}
	err = f.SetCellValue(sheetName, "B1", "Значение")
	if err != nil {
		log.Panicln(err)
		return ExelFile{filename: "", data: nil}
	}
	for i, r := range rec {
		err := f.SetCellValue(sheetName, makeAxis(0, i+1), r.Key)
		if err != nil {
			log.Panicln(err)
			return ExelFile{filename: "", data: nil}
		}
		err = f.SetCellValue(sheetName, makeAxis(1, i+1), r.Value)
		if err != nil {
			log.Panicln(err)
			return ExelFile{filename: "", data: nil}
		}
	}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	err = f.Write(writer)
	if err != nil {
		log.Panicln(err)
		return ExelFile{filename: "", data: nil}
	}
	fileName := "cache_data_" + time.Now().String() + ".xlsx"
	return ExelFile{filename: fileName, data: b.Bytes()}

}

func makeAxis(x, y int) string {
	var r []byte
	if x < 26 {
		r = make([]byte, 1, 5)
		r[0] = 'A' + byte(x)
	} else if x < 27*26 {
		r = make([]byte, 2, 5)
		r[0] = 'A' - 1 + byte(x/26)
		r[1] = 'A' + byte(x%26)
	} else if x < 16384 {
		r = make([]byte, 3, 5)
		r[2] = 'A' + byte(x%26)
		x /= 26
		r[0] = 'A' - 1 + byte(x/26)
		r[1] = 'A' - 1 + byte(x%26)
	} else {
		panic(fmt.Errorf("more than 16384 columns: %d", x))
	}
	return string(strconv.AppendUint(r, uint64(y+1), 10))
}
