package setup

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/NoUseFreak/sdf/internal/pkg/utils"
)

func SetupShellFunc() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	for _, profileName := range []string{".bash_profile", ".zshrc"} {
		profile := path.Join(usr.HomeDir, profileName)
		needsSetup, err := needsShellSetup(profile)
		if err != nil {
			output.Error("Something went wrong\n")
		}

		if needsSetup {
			writeBashFunc(profile)
			output.Print("Need to install\n")
		}
	}

}

func needsShellSetup(bashProfile string) (bool, error) {
	if _, err := os.Stat(bashProfile); err == nil {
		f, err := os.Open(bashProfile)
		if err != nil {
			return false, err
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "sdf") {
				return false, nil
			}
		}
	}

	return true, nil
}

func writeBashFunc(bashProfile string) {
	f, err := os.OpenFile(bashProfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	sdfPath := utils.SdfBinaryPath()

	text := `
# Do not change this line
func sdf() { 
	local sdfFile=$(` + sdfPath + ` _tmpfile)
	` + sdfPath + ` "$@" > $sdfFile
	source $sdfFile || echo "Debug info can be found in ${sdfFile}"
}
`

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
	output.Println(text)
}
