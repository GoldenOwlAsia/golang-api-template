package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestMakeFilenameUnique(t *testing.T) {
	// Test MakeFilenameUnique with a sample file name
	fileName := "My File Name.jpeg"
	uniqueFileName := MakeFilenameUnique(fileName)

	// Check if the unique file name has the correct format
	expectedName := "my-file-name-" + fmt.Sprint(time.Now().Unix()) + ".jpeg"
	if uniqueFileName != expectedName {
		t.Errorf("MakeFilenameUnique(%q) = %q, want %q", fileName, uniqueFileName, expectedName)
	}
}

func TestCleanFileName(t *testing.T) {
	// Test CleanFileName with a sample file name
	fileName := "My File Name"
	cleanedName := CleanFileName(fileName)

	// Check if the cleaned file name has the correct format
	expectedName := "my-file-name"
	if cleanedName != expectedName {
		t.Errorf("CleanFileName(%q) = %q, want %q", fileName, cleanedName, expectedName)
	}

	// Test CleanFileName with another sample file name
	fileName2 := "My_File-Name_123"
	cleanedName2 := CleanFileName(fileName2)

	// Check if the cleaned file name has the correct format
	expectedName2 := "my_file_name_123"
	if cleanedName2 != expectedName2 {
		t.Errorf("CleanFileName(%q) = %q, want %q", fileName2, cleanedName2, expectedName2)
	}
}
