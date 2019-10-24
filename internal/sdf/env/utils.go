package env

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

const DefaultEditor = "vim"

func modelPath(name string) string {
	return path.Join(envDir(), name+".json")
}

func readModelFile(path string) (model, error) {
	var m model
	jsonFile, err := os.Open(path)
	if err != nil {
		return m, nil
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return m, err
	}

	if err := json.Unmarshal(byteValue, &m); err != nil {
		return m, err
	}

	m.Data = decodeData(m.Data)

	return m, nil
}

func writeModelFile(path string, m model) error {
	m.Data = encodeData(m.Data)
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, b, os.FileMode(0600)); err != nil {
		return err
	}

	return nil
}

func getEditorCmd(data, nextCmd string) error {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return err
	}

	filename := file.Name()

	if err = file.Close(); err != nil {
		return err
	}

	// Fill file with previous data
	if err := ioutil.WriteFile(filename, []byte(data), 0777); err != nil {
		return err
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	editorCmd := fmt.Sprintf("%s %s", executable, filename)
	output.Exec(editorCmd)
	if nextCmd != "" {
		output.Exec(nextCmd, filename)
	}
	return nil
}

func envDir() string {
	return path.Join(viper.GetString("configdir"), "env")
}

func yesNo() bool {
	prompt := promptui.Select{
		Label: "Select[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func encodeData(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func decodeData(data string) string {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}

	return string(b)
}
