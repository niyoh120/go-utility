package fileutil

import "os"

func Exist(path string) (bool, error) {
    if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            return false, nil
        }
        return false, err
    }
    return true, nil
}
