package wdid

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	parallel "github.com/NoUseFreak/go-parallel"
	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func PrintReport(window string) {
	out, err := exec.Command("git", "config", "--get", "user.email").Output()
	if err != nil {
		log.Fatal(err)
	}
	author := strings.TrimSpace(string(out))

	input := parallel.Input{}
	for _, project := range findProjects() {
		input = append(input, project)
	}

	output.Println("Looking for commits by %s in %d repos since %s", author, len(input), window)

	p := parallel.Processor{Threads: 10}
	result := p.Process(input, func(i interface{}) interface{} {
		project := i.(string)

		output.Debug("Checking", project, "for commits by", author)
		cmd := exec.Command("git", "log", "--reverse", "--no-merges", "--pretty=%ct %s", "--since='"+window+"'", "--author="+author, "--branches=*")
		cmd.Dir = project
		out, err := cmd.Output()
		if err != nil {
			output.Errorln(errors.Wrapf(err, "Failure on %s ", project).Error())
			return nil
		}
		s := strings.TrimSpace(string(out))
		if s == "" {
			return nil
		}

		var res []repoLog
		for _, l := range strings.Split(s, "\n") {
			res = append(res, repoLog{
				line:    l,
				project: project,
			})
		}
		return res
	})

	var full []repoLog
	for _, r := range result {
		if l, ok := r.([]repoLog); ok {
			full = append(full, l...)
		}
	}

	full = unique(full)
	sort.SliceStable(full, func(i, j int) bool {
		return full[i].timestamp() < full[j].timestamp()
	})

	output.Println("")
	output.Println("Update:")
	for _, r := range full {
		output.Println(" - %s", r.String())
	}
}

func unique(slice []repoLog) []repoLog {
	keys := make(map[string]bool)
	var list []repoLog
	for _, entry := range slice {
		k := fmt.Sprintf("%s-%s", entry.changeLine(), entry.project)
		if _, value := keys[k]; !value {
			keys[k] = true
			list = append(list, entry)
		}
	}

	return list
}

type repoLog struct {
	line    string
	project string
}

func (r *repoLog) String() string {
	return fmt.Sprintf("%s (%s)", r.changeLine(), r.shortRepo())
}

func (r *repoLog) shortRepo() string {
	return strings.ReplaceAll(r.project, viper.GetString("projectdir")+"/", "")
}
func (r *repoLog) changeLine() string {
	return r.line[11:]
}
func (r *repoLog) timestamp() int64 {
	if t, err := strconv.ParseInt(r.line[0:10], 10, 64); err == nil {
		return t
	}
	return 0
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

func findProjects() []string {
	s := fmt.Sprintf("%s/*/*/*/.git", viper.GetString("projectdir"))
	m, _ := filepath.Glob(s)
	var o []string
	for _, p := range m {
		o = append(o, filepath.Dir(p))
	}
	return o
}
