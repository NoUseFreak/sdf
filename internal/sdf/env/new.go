package env

import (
	"fmt"

	"github.com/NoUseFreak/sdf/internal/pkg/utils"
)

func New(name string) error {
	nextCmd := fmt.Sprintf("%s env _savetmpfile %s %s", utils.SdfBinaryPath(), name, "%s")
	return getEditorCmd("", nextCmd)

}
