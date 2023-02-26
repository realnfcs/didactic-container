package interfaces

type ImageInterface interface {
    DownloadImage(...string) error
    GetLocalImage(local, imgName, filename string, args ...string) error
    DeleteImage(id, name, path string)
}
