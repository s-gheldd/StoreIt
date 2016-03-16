package DataStructure

import (
	"os"
	"strings"
	"fmt"
)

const descriptorFormatString string = "StoredFile { FileName=\"%s\" FileType=\".%s\" }"
const descriptorFormatStringVerbose string = "StoredFile { FileName=\"%s\" FileType=\".%s\" Data=%v}"

type StoredFile struct {
	fileName string
	fileType string
	data     []byte
}

func NewStoredFile(file *os.File) (*StoredFile, error) {

	fileStat, err := file.Stat();
	if err != nil {
		return nil, err
	}

	newFile := &StoredFile{};
	fullName := strings.Split(fileStat.Name(), ".")

	newFile.fileName = fullName[0]
	newFile.fileType = fullName[1]
	newFile.data = make([]byte, fileStat.Size())
	_, err = file.Read(newFile.data)
	if err != nil {
		return nil, err
	}
	return newFile, nil
}

func (s StoredFile) RestoreInDir() *os.File {

}

func (s StoredFile) String() string {
	return fmt.Sprintf(descriptorFormatString, s.fileName, s.fileType)
}

func (s StoredFile) VerboseString() string {
	return fmt.Sprintf(descriptorFormatStringVerbose, s.fileName, s.fileType, s.data)
}