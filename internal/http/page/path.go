package page

import (
	"path/filepath"
)

func getCommonPaths() map[string]string {
	commonPaths := make(map[string]string)
	commonPaths["footPath"], _ = filepath.Abs(filepath.Join(".", "view", "common", "footer.html"))
	commonPaths["headPath"], _ = filepath.Abs(filepath.Join(".", "view", "common", "header.html"))
	commonPaths["headAuthPath"], _ = filepath.Abs(filepath.Join(".", "view", "common", "headerAuth.html"))
	commonPaths["headAdminPath"], _ = filepath.Abs(filepath.Join(".", "view", "common", "headerAdmin.html"))
	commonPaths["headRegPath"], _ = filepath.Abs(filepath.Join(".", "view", "common", "headerReg.html"))
	return commonPaths
}
