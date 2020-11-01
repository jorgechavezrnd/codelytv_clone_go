package godog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"

	moocFrontend "tv/codely/apps/mooc/frontend"

	"github.com/cucumber/godog"
)

type moocFrontendFeature struct {
	resp *http.Response
}

func (a *moocFrontendFeature) iSendARequestTo(method, endpoint string) (err error) {
	req := httptest.NewRequest(method, endpoint, nil)
	app := moocFrontend.NewMoocFrontendApplication()
	app.ConfigureServer()

	resp, _ := app.Test(req)

	a.resp = resp

	return nil
}

func (a *moocFrontendFeature) theResponseContentShouldBe(body *godog.DocString) (err error) {
	var expected, actual interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	// re-encode actual response too
	actualRespBytes, _ := ioutil.ReadAll(a.resp.Body)
	if err = json.Unmarshal(actualRespBytes, &actual); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("The output does not match!\n\n-- Expected:\n%s\n\n--Actual:\n%s\n\n",
			mapToJsonString(expected),
			mapToJsonString(actual),
		)
	}
	return nil
}

func mapToJsonString(mapWithData interface{}) string {
	stringJson, _ := json.Marshal(mapWithData)

	return string(stringJson)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	app := &moocFrontendFeature{}

	ctx.Step(`^I send a "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, app.iSendARequestTo)
	ctx.Step(`^the response content should be:$`, app.theResponseContentShouldBe)
}
