package internal

import "os"

const (
	WORKSPACE_FOLDER_PATH string = "./workspace"
	FS_FOLDER_PATH        string = "./workspace/fs"
)

// Function to inicialize the workspace and filesystem folders
func InitFolders() {
	err := os.Mkdir(WORKSPACE_FOLDER_PATH, 0755)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(FS_FOLDER_PATH, 0755)
	if err != nil {
		panic(err)
	}
}
