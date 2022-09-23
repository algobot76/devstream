package jenkins

import (
	_ "embed"

	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bndr/gojenkins"

	"github.com/devstream-io/devstream/pkg/util/log"
	"github.com/devstream-io/devstream/pkg/util/template"
)

//go:embed tpl/casc.tpl.groovy
var cascGroovyScript string

type scriptError struct {
	output   string
	errorMsg string
}

// all jenkins client return error format is host: error detail
// Error method wrap this type of error to only return error detail
func (e *scriptError) Error() string {
	errDetailMessage := e.errorMsg
	if strings.Contains(errDetailMessage, ":") {
		errDetailMessageArray := strings.Split(errDetailMessage, ":")
		errDetailMessage = strings.TrimSpace(errDetailMessageArray[len(errDetailMessageArray)-1])
	}
	return fmt.Sprintf("execute groovy script failed: %s(%s)", errDetailMessage, e.output)
}

func (jenkins *jenkins) ExecuteScript(script string) (string, error) {
	now := time.Now().Unix()
	verifier := fmt.Sprintf("verifier-%d", now)
	output := ""
	fullScript := fmt.Sprintf("%s\nprint println('%s')", script, verifier)

	data := url.Values{}
	data.Set("script", fullScript)

	ar := gojenkins.NewAPIRequest("POST", "/scriptText", bytes.NewBufferString(data.Encode()))
	if err := jenkins.Requester.SetCrumb(jenkins.ctx, ar); err != nil {
		return output, err
	}
	ar.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	ar.Suffix = ""

	r, err := jenkins.Requester.Do(jenkins.ctx, ar, &output, nil)
	if err != nil {
		return "", &scriptError{
			output:   output,
			errorMsg: err.Error(),
		}
	}
	log.Debugf("------> %s\n%s", output, script)
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return output, &scriptError{
			output:   output,
			errorMsg: fmt.Sprintf("invalid status code '%d'", r.StatusCode),
		}
	}

	if !strings.Contains(output, verifier) {
		return output, &scriptError{
			output:   output,
			errorMsg: "script verifier error",
		}
	}

	return output, nil
}

func (jenkins *jenkins) ConfigCasc(cascConfig string) error {
	groovyCascScript, err := template.Render(
		"jenkins casc", cascGroovyScript, map[string]string{
			"CascConfig": cascConfig,
		},
	)
	if err != nil {
		log.Debugf("jenkins render casc failed: %s", err)
		return err
	}
	_, err = jenkins.ExecuteScript(groovyCascScript)
	return err
}