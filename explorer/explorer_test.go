package explorer

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"time"
)

func Test_RootDirectories_ShouldReturnASliceOfDirectoriesWithinTheExplorerRoot(t *testing.T) {
	os.Mkdir("rootDir", 777)
	os.Mkdir("rootDir/dir1", 777)
	os.Mkdir("rootDir/dir2", 777)
	os.Mkdir("rootDir/dir3", 777)
	file1, _ := os.Create("rootDir/file1.txt")
	file2, _ := os.Create("rootDir/file2.txt")
	file3, _ := os.Create("rootDir/file3.txt")
	defer func() {
		file1.Close()
		file2.Close()
		file3.Close()
		os.RemoveAll("rootDir")
	}()
	explorer := New("rootDir")
	expected := []Directory{
		Directory{
			Name: "dir1",
			Path: "rootDir/dir1",
		},
		Directory{
			Name: "dir2",
			Path: "rootDir/dir2",
		},
		Directory{
			Name: "dir3",
			Path: "rootDir/dir3",
		},
	}

	actual, _ := explorer.RootDirectories()

	assert.Equal(t, expected, actual)
}

func Test_RootFiles_ShouldReturnASliceOfFilesWithinTheExplorerRoot(t *testing.T) {
	os.Mkdir("rootDir", 777)
	os.Mkdir("rootDir/dir1", 777)
	os.Mkdir("rootDir/dir2", 777)
	os.Mkdir("rootDir/dir3", 777)
	file1, _ := os.Create("rootDir/file1.txt")
	file2, _ := os.Create("rootDir/file2.txt")
	file3, _ := os.Create("rootDir/file3.txt")
	defer func() {
		file1.Close()
		file2.Close()
		file3.Close()
		os.RemoveAll("rootDir")
	}()
	expected := []File{
		File{
			Name: "file1.txt",
			Path: "rootDir/file1.txt",
			Size: 0,
		},
		File{
			Name: "file2.txt",
			Path: "rootDir/file2.txt",
			Size: 0,
		},
		File{
			Name: "file3.txt",
			Path: "rootDir/file3.txt",
			Size: 0,
		},
	}
	explorer := New("rootDir")

	actual, _ := explorer.RootFiles()

	assert.Equal(t, expected, actual)
}

func Test_Directories_ShouldReturnASliceOfDirectoriesWithinTheProvidedPath(t *testing.T) {
	os.Mkdir("rootDir", 777)
	os.Mkdir("rootDir/dir1", 777)
	os.Mkdir("rootDir/dir2", 777)
	os.Mkdir("rootDir/dir3", 777)
	file1, _ := os.Create("rootDir/file1.txt")
	file2, _ := os.Create("rootDir/file2.txt")
	file3, _ := os.Create("rootDir/file3.txt")
	defer func() {
		file1.Close()
		file2.Close()
		file3.Close()
		os.RemoveAll("rootDir")
	}()
	explorer := New(".")
	expected := []Directory{
		Directory{
			Name: "dir1",
			Path: "./rootDir/dir1",
		},
		Directory{
			Name: "dir2",
			Path: "./rootDir/dir2",
		},
		Directory{
			Name: "dir3",
			Path: "./rootDir/dir3",
		},
	}

	actual, err := explorer.Directories("./rootDir")

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func Test_Directories_ShouldReturnErrorIfScannedDirectoryIsOutsideTheRoot(t *testing.T) {
	os.Mkdir("rootDir", 777)
	defer func() {
		os.RemoveAll("rootDir")
	}()
	explorer := New("./rootDir")
	var expected []Directory

	actual, err := explorer.Directories(".")

	assert.Equal(t, expected, actual)
	assert.Equal(t, ERR_OUT_OF_ROOT, err)
}

func Test_Directories_ShouldReturnErrorIfDirectoryNotExists(t *testing.T) {
	os.Mkdir("rootDir", 777)
	defer func() {
		os.RemoveAll("rootDir")
	}()
	explorer := New("./rootDir")
	var expected []Directory

	actual, err := explorer.Directories("./rootDir/unknownDir")

	assert.Equal(t, expected, actual)
	assert.Equal(t, ERR_CANNOT_SCAN, err)
}

func Test_Files_ShouldReturnASliceOfFilesWithinTheProvidedPath(t *testing.T) {
	os.Mkdir("rootDir", 777)
	os.Mkdir("rootDir/dir1", 777)
	os.Mkdir("rootDir/dir2", 777)
	os.Mkdir("rootDir/dir3", 777)
	file1, _ := os.Create("rootDir/file1.txt")
	file2, _ := os.Create("rootDir/file2.txt")
	file3, _ := os.Create("rootDir/file3.txt")
	defer func() {
		file1.Close()
		file2.Close()
		file3.Close()
		os.RemoveAll("rootDir")
	}()
	explorer := New(".")
	expected := []File{
		File{
			Name: "file1.txt",
			Path: "./rootDir/file1.txt",
			Size: 0,
		},
		File{
			Name: "file2.txt",
			Path: "./rootDir/file2.txt",
			Size: 0,
		},
		File{
			Name: "file3.txt",
			Path: "./rootDir/file3.txt",
			Size: 0,
		},
	}

	actual, err := explorer.Files("./rootDir")

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func Test_Files_ShouldReturnErrorIfScannedDirectoryIsOutsideTheRoot(t *testing.T) {
	os.Mkdir("rootDir", 777)
	defer func() {
		os.RemoveAll("rootDir")
	}()
	explorer := New("./rootDir")
	var expected []File

	actual, err := explorer.Files(".")

	assert.Equal(t, expected, actual)
	assert.Equal(t, ERR_OUT_OF_ROOT, err)
}

func Test_Files_ShouldReturnErrorIfDirectoryNotExists(t *testing.T) {
	os.Mkdir("rootDir", 777)
	defer func() {
		os.RemoveAll("rootDir")
	}()
	explorer := New("./rootDir")
	var expected []File

	actual, err := explorer.Files("./rootDir/unknownDir")

	assert.Equal(t, expected, actual)
	assert.Equal(t, ERR_CANNOT_SCAN, err)
}

func Test_checkLevel_ShouldNotReturnErrorIfPathIsUnderTheRoot(t *testing.T) {
	explorer := New("rootDir")

	err := explorer.checkLevel("rootDir/subDir")

	assert.Equal(t, nil, err)
}

func Test_checkLevel_ShouldReturnErrorIfPathIsNotUnderTheRoot(t *testing.T) {
	explorer := New("rootDir")

	err := explorer.checkLevel("notARootDir/subDir")

	assert.Equal(t, ERR_OUT_OF_ROOT, err)
}

// Implements interface FileInfo
type mockFileInfo struct {
	mock.Mock
}

func (m *mockFileInfo)  Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *mockFileInfo)  Size() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *mockFileInfo) Mode() os.FileMode {
	args := m.Called()
	return args.Get(0).(os.FileMode)
}

func (m *mockFileInfo) ModTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func (m *mockFileInfo) IsDir() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *mockFileInfo) Sys() interface{} {
	args := m.Called()
	return args.Get(0)
}

func Test_filterDirectories_ShouldReturnOnlyDirectories(t *testing.T) {
	dir1 := new(mockFileInfo)
	dir1.On("IsDir").Return(true)
	dir1.On("Name").Return("directory1")
	dir2 := new(mockFileInfo)
	dir2.On("IsDir").Return(true)
	dir2.On("Name").Return("directory2")
	notDir := new(mockFileInfo)
	notDir.On("IsDir").Return(false)
	entities := []os.FileInfo{
		dir1,
		dir2,
		notDir,
	}
	expected := []Directory{
		Directory{
			Name: "directory1",
			Path: "dummy-path/directory1",
		},
		Directory{
			Name: "directory2",
			Path: "dummy-path/directory2",
		},
	}

	actual := filterDirectories(entities, "dummy-path")

	assert.Equal(t, expected, actual)
}

func Test_filterFiles_ShouldReturnOnlyFiles(t *testing.T) {
	file1 := new(mockFileInfo)
	file1.On("IsDir").Return(false)
	file1.On("Name").Return("file1")
	file1.On("Size").Return(int64(1))
	file2 := new(mockFileInfo)
	file2.On("IsDir").Return(false)
	file2.On("Name").Return("file2")
	file2.On("Size").Return(int64(2))
	notFile := new(mockFileInfo)
	notFile.On("IsDir").Return(true)
	entities := []os.FileInfo{
		file1,
		file2,
		notFile,
	}
	expected := []File{
		File{
			Name: "file1",
			Path: "dummy-path/file1",
			Size: 1,
		},
		File{
			Name: "file2",
			Path: "dummy-path/file2",
			Size: 2,
		},
	}

	actual := filterFiles(entities, "dummy-path")

	assert.Equal(t, expected, actual)
}

func Test_buildPath_ShouldBuildEntityPathCorrectly(t *testing.T) {
	assert.Equal(t, "C:/subdir/file.txt", buildPath("C:/subdir", "file.txt"))
	assert.Equal(t, "C:/subdir/file.txt", buildPath("C:/subdir/", "file.txt"))
}