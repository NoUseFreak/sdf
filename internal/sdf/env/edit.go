package env

import (
	"fmt"

	"github.com/NoUseFreak/sdf/internal/pkg/utils"
	"github.com/pkg/errors"
)

func Edit(name string) error {
	nextCmd := fmt.Sprintf("%s env _savetmpfile %s %s", utils.SdfBinaryPath(), name, "%s")
	m, err := readModelFile(modelPath(name))
	if err != nil {
		return errors.Wrapf(err, "Could not find %s", name)
	}

	return getEditorCmd(m.Data, nextCmd)
}
