package dependencies

import "github.com/realnfcs/didactic-container/internal/image"

func GetImageEngine() *image.Filesystem {
    return &image.Filesystem{}
}
