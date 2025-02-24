package ankaregistry

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/veertuinc/packer-plugin-veertu-anka/builder/anka"
	"github.com/veertuinc/packer-plugin-veertu-anka/client"
	mocks "github.com/veertuinc/packer-plugin-veertu-anka/mocks"
	"gotest.tools/v3/assert"
)

var templateList []client.RegistryListResponse
var registryRemote client.RegistryRemote
var registryRemoteArm64 client.RegistryRemoteArm64
var reposList client.RegistryListReposResponse

func TestAnkaRegistryPostProcessor(t *testing.T) {

	err := json.Unmarshal(json.RawMessage(`{"default": true, "host": "localhost", "scheme": "http", "port": "8080"}`), &registryRemote)
	if err != nil {
		t.Fail()
	}

	err = json.Unmarshal(json.RawMessage(`{"default": true, "url": "http://localhost:8080", "name": "go-mock"}`), &registryRemoteArm64)
	if err != nil {
		t.Fail()
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ankaClient := mocks.NewMockClient(mockCtrl)

	ui := packer.TestUi(t)

	artifact := &anka.Artifact{}

	reposList = client.RegistryListReposResponse{
		Default: "go-mock",
		Remotes: map[string]client.RegistryRemote{"go-mock": registryRemote},
	}

	t.Run("push to registry with defaults", func(t *testing.T) {
		config := Config{
			RemoteVM:    "foo",
			Tag:         "registry-push",
			Description: "mock for testing anka registry push",
			HostArch:    runtime.GOARCH,
		}

		pp := PostProcessor{
			config: config,
			client: ankaClient,
		}

		registryParams := client.RegistryParams{
			RegistryName: "go-mock",
			RegistryURL:  "http://localhost:8080",
			HostArch:     config.HostArch,
		}

		pushParams := client.RegistryPushParams{
			Tag:         config.Tag,
			Description: config.Description,
			RemoteVM:    config.RemoteVM,
			Local:       false,
			Force:       false,
		}

		if runtime.GOARCH == "amd64" {
			ankaClient.EXPECT().RegistryListRepos().Return(reposList, nil).Times(1)
		} else {
			ankaClient.EXPECT().RegistryListReposArm64().Return(reposList, nil).Times(1)
		}
		ankaClient.EXPECT().RegistryList(registryParams).Return([]client.RegistryListResponse{}, nil).Times(1)
		ankaClient.EXPECT().RegistryPush(registryParams, pushParams).Return(nil).Times(1)

		mockui := packer.MockUi{}
		mockui.Say(fmt.Sprintf("Pushing template to Anka Registry as %s with tag %s", config.RemoteVM, config.Tag))
		mockui.Say("Registry push successful")

		assert.Equal(t, mockui.SayMessages[0].Message, "Pushing template to Anka Registry as foo with tag registry-push")
		assert.Equal(t, mockui.SayMessages[1].Message, "Registry push successful")

		_, _, _, err := pp.PostProcess(context.Background(), ui, artifact)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("push to registry with registry name", func(t *testing.T) {
		config := Config{
			RegistryName: "go-mock",
			RemoteVM:     "foo",
			Tag:          "registry-push",
			Description:  "mock for testing anka registry push",
			HostArch:     runtime.GOARCH,
		}

		pp := PostProcessor{
			config: config,
			client: ankaClient,
		}

		registryParams := client.RegistryParams{
			RegistryName: "go-mock",
			RegistryURL:  "http://localhost:8080",
			HostArch:     config.HostArch,
		}

		pushParams := client.RegistryPushParams{
			Tag:         config.Tag,
			Description: config.Description,
			RemoteVM:    config.RemoteVM,
			Local:       false,
			Force:       false,
		}

		if runtime.GOARCH == "amd64" {
			ankaClient.EXPECT().RegistryListRepos().Return(reposList, nil).Times(1)
		} else {
			ankaClient.EXPECT().RegistryListReposArm64().Return(reposList, nil).Times(1)
		}
		ankaClient.EXPECT().RegistryList(registryParams).Return([]client.RegistryListResponse{}, nil).Times(1)
		ankaClient.EXPECT().RegistryPush(registryParams, pushParams).Return(nil).Times(1)

		mockui := packer.MockUi{}
		mockui.Say(fmt.Sprintf("Pushing template to Anka Registry as %s with tag %s", config.RemoteVM, config.Tag))

		assert.Equal(t, mockui.SayMessages[0].Message, "Pushing template to Anka Registry as foo with tag registry-push")

		_, _, _, err := pp.PostProcess(context.Background(), ui, artifact)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("push to registry with registry URL", func(t *testing.T) {
		config := Config{
			RegistryURL: "http://anka.example.test:8080",
			RemoteVM:    "foo",
			Tag:         "registry-push",
			Description: "mock for testing anka registry push",
			HostArch:    runtime.GOARCH,
		}

		pp := PostProcessor{
			config: config,
			client: ankaClient,
		}

		registryParams := client.RegistryParams{
			RegistryName: "",
			RegistryURL:  "http://anka.example.test:8080",
			HostArch:     config.HostArch,
		}

		pushParams := client.RegistryPushParams{
			Tag:         config.Tag,
			Description: config.Description,
			RemoteVM:    config.RemoteVM,
			Local:       false,
			Force:       false,
		}

		if runtime.GOARCH == "amd64" {
			ankaClient.EXPECT().RegistryListRepos().Return(reposList, nil).Times(1)
		} else {
			ankaClient.EXPECT().RegistryListReposArm64().Return(reposList, nil).Times(1)
		}
		ankaClient.EXPECT().RegistryList(registryParams).Return([]client.RegistryListResponse{}, nil).Times(1)
		ankaClient.EXPECT().RegistryPush(registryParams, pushParams).Return(nil).Times(1)

		mockui := packer.MockUi{}
		mockui.Say(fmt.Sprintf("Pushing template to Anka Registry as %s with tag %s", config.RemoteVM, config.Tag))

		assert.Equal(t, mockui.SayMessages[0].Message, "Pushing template to Anka Registry as foo with tag registry-push")

		_, _, _, err := pp.PostProcess(context.Background(), ui, artifact)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("push to registry with no existing templates", func(t *testing.T) {
		packerConfig := common.PackerConfig{
			PackerForce: true,
		}

		config := Config{
			PackerConfig: packerConfig,
			RemoteVM:     "foo",
			Tag:          "registry-push",
			Description:  "mock for testing anka registry push",
			HostArch:     runtime.GOARCH,
		}

		pp := PostProcessor{
			config: config,
			client: ankaClient,
		}

		registryParams := client.RegistryParams{
			RegistryName: "go-mock",
			RegistryURL:  "http://localhost:8080",
			HostArch:     config.HostArch,
		}

		pushParams := client.RegistryPushParams{
			Tag:         config.Tag,
			Description: config.Description,
			RemoteVM:    config.RemoteVM,
			Local:       false,
			Force:       false,
		}

		if runtime.GOARCH == "amd64" {
			ankaClient.EXPECT().RegistryListRepos().Return(reposList, nil).Times(1)
		} else {
			ankaClient.EXPECT().RegistryListReposArm64().Return(reposList, nil).Times(1)
		}
		ankaClient.EXPECT().RegistryList(registryParams).Return([]client.RegistryListResponse{}, nil).Times(1)
		ankaClient.EXPECT().RegistryPush(registryParams, pushParams).Return(nil).Times(1)

		mockui := packer.MockUi{}
		mockui.Say(fmt.Sprintf("Pushing template to Anka Registry as %s with tag %s", config.RemoteVM, config.Tag))

		assert.Equal(t, mockui.SayMessages[0].Message, "Pushing template to Anka Registry as foo with tag registry-push")

		_, _, _, err := pp.PostProcess(context.Background(), ui, artifact)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("push to registry with existing template and fail", func(t *testing.T) {
		err := json.Unmarshal(json.RawMessage(`[{ "id": "foo_id", "name": "foo", "latest": "foo_tag" }]`), &templateList)
		if err != nil {
			t.Fail()
		}

		packerConfig := common.PackerConfig{
			PackerForce: false,
		}

		config := Config{
			PackerConfig: packerConfig,
			RemoteVM:     "foo",
			Tag:          "registry-push",
			Description:  "mock for testing anka registry push",
			HostArch:     runtime.GOARCH,
		}

		pp := PostProcessor{
			config: config,
			client: ankaClient,
		}

		registryParams := client.RegistryParams{
			RegistryName: "go-mock",
			RegistryURL:  "http://localhost:8080",
			HostArch:     config.HostArch,
		}

		if runtime.GOARCH == "amd64" {
			ankaClient.EXPECT().RegistryListRepos().Return(reposList, nil).Times(1)
		} else {
			ankaClient.EXPECT().RegistryListReposArm64().Return(reposList, nil).Times(1)
		}
		ankaClient.EXPECT().RegistryList(registryParams).Return(templateList, nil).Times(1)

		mockui := packer.MockUi{}
		mockui.Say(fmt.Sprintf("Pushing template to Anka Registry as %s with tag %s", config.RemoteVM, config.Tag))
		mockui.Say(fmt.Sprintf("Found existing template %s on registry that matches name '%s'", templateList[0].ID, config.RemoteVM))

		assert.Equal(t, mockui.SayMessages[0].Message, "Pushing template to Anka Registry as foo with tag registry-push")
		assert.Equal(t, mockui.SayMessages[1].Message, "Found existing template foo_id on registry that matches name 'foo'")

		_, _, _, err = pp.PostProcess(context.Background(), ui, artifact)
		if err == nil {
			t.Fail()
		}
	})

	t.Run("push to registry with existing template and don't revert latest tag first [packer build -force]", func(t *testing.T) {
		err := json.Unmarshal(json.RawMessage(`[{ "id": "foo_id", "name": "foo", "latest": "foo_tag" }]`), &templateList)
		if err != nil {
			t.Fail()
		}

		packerConfig := common.PackerConfig{
			PackerForce: true,
		}

		config := Config{
			PackerConfig: packerConfig,
			RemoteVM:     "foo",
			Tag:          "registry-push",
			Description:  "mock for testing anka registry push",
			HostArch:     runtime.GOARCH,
		}

		pp := PostProcessor{
			config: config,
			client: ankaClient,
		}

		registryParams := client.RegistryParams{
			RegistryName: "go-mock",
			RegistryURL:  "http://localhost:8080",
			HostArch:     config.HostArch,
		}

		pushParams := client.RegistryPushParams{
			Tag:         config.Tag,
			Description: config.Description,
			RemoteVM:    config.RemoteVM,
			Local:       false,
			Force:       false,
		}

		if runtime.GOARCH == "amd64" {
			ankaClient.EXPECT().RegistryListRepos().Return(reposList, nil).Times(1)
		} else {
			ankaClient.EXPECT().RegistryListReposArm64().Return(reposList, nil).Times(1)
		}
		ankaClient.EXPECT().RegistryList(registryParams).Return(templateList, nil).Times(1)
		ankaClient.EXPECT().RegistryRevert(registryParams.RegistryURL, templateList[0].ID).Return(nil).Times(0)
		ankaClient.EXPECT().RegistryPush(registryParams, pushParams).Return(nil).Times(1)

		mockui := packer.MockUi{}
		mockui.Say(fmt.Sprintf("Pushing template to Anka Registry as %s with tag %s", config.RemoteVM, config.Tag))
		mockui.Say(fmt.Sprintf("Found existing template %s on registry that matches name '%s'", templateList[0].ID, config.RemoteVM))

		assert.Equal(t, mockui.SayMessages[0].Message, "Pushing template to Anka Registry as foo with tag registry-push")
		assert.Equal(t, mockui.SayMessages[1].Message, "Found existing template foo_id on registry that matches name 'foo'")

		_, _, _, err = pp.PostProcess(context.Background(), ui, artifact)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("push to registry with existing templates with latest tag match and revert tag first [packer build -force]", func(t *testing.T) {
		err := json.Unmarshal(json.RawMessage(`[{ "id": "foo_id", "name": "foo", "latest": "registry-push" }]`), &templateList)
		if err != nil {
			t.Fail()
		}

		packerConfig := common.PackerConfig{
			PackerForce: true,
		}

		config := Config{
			PackerConfig: packerConfig,
			RemoteVM:     "foo",
			Tag:          "registry-push",
			Description:  "mock for testing anka registry push",
			HostArch:     runtime.GOARCH,
		}

		pp := PostProcessor{
			config: config,
			client: ankaClient,
		}

		registryParams := client.RegistryParams{
			RegistryName: "go-mock",
			RegistryURL:  "http://localhost:8080",
			HostArch:     config.HostArch,
		}

		pushParams := client.RegistryPushParams{
			Tag:         config.Tag,
			Description: config.Description,
			RemoteVM:    config.RemoteVM,
			Local:       false,
			Force:       false,
		}

		if runtime.GOARCH == "amd64" {
			ankaClient.EXPECT().RegistryListRepos().Return(reposList, nil).Times(1)
		} else {
			ankaClient.EXPECT().RegistryListReposArm64().Return(reposList, nil).Times(1)
		}
		ankaClient.EXPECT().RegistryList(registryParams).Return(templateList, nil).Times(1)
		ankaClient.EXPECT().RegistryRevert(registryParams.RegistryURL, templateList[0].ID).Return(nil).Times(1)
		ankaClient.EXPECT().RegistryPush(registryParams, pushParams).Return(nil).Times(1)

		mockui := packer.MockUi{}
		mockui.Say(fmt.Sprintf("Pushing template to Anka Registry as %s with tag %s", config.RemoteVM, config.Tag))
		mockui.Say(fmt.Sprintf("Found existing template %s on registry that matches name '%s'", templateList[0].ID, config.RemoteVM))
		mockui.Say(fmt.Sprintf("Reverted latest tag for template '%s' on registry", templateList[0].ID))

		assert.Equal(t, mockui.SayMessages[0].Message, "Pushing template to Anka Registry as foo with tag registry-push")
		assert.Equal(t, mockui.SayMessages[1].Message, "Found existing template foo_id on registry that matches name 'foo'")
		assert.Equal(t, mockui.SayMessages[2].Message, "Reverted latest tag for template 'foo_id' on registry")

		_, _, _, err = pp.PostProcess(context.Background(), ui, artifact)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("force push to registry with existing template", func(t *testing.T) {
		err := json.Unmarshal(json.RawMessage(`[{ "id": "foo_id", "name": "foo", "latest": "registry-push" }]`), &templateList)
		if err != nil {
			t.Fail()
		}

		config := Config{
			RemoteVM:    "foo",
			Tag:         "registry-push",
			Description: "mock for testing anka registry push",
			HostArch:    runtime.GOARCH,
			Force:       true,
		}

		pp := PostProcessor{
			config: config,
			client: ankaClient,
		}

		registryParams := client.RegistryParams{
			RegistryName: "go-mock",
			RegistryURL:  "http://localhost:8080",
			HostArch:     config.HostArch,
		}

		pushParams := client.RegistryPushParams{
			Tag:         config.Tag,
			Description: config.Description,
			RemoteVM:    config.RemoteVM,
			Local:       false,
			Force:       true,
		}

		if runtime.GOARCH == "amd64" {
			ankaClient.EXPECT().RegistryListRepos().Return(reposList, nil).Times(1)
		} else {
			ankaClient.EXPECT().RegistryListReposArm64().Return(reposList, nil).Times(1)
		}
		ankaClient.EXPECT().RegistryList(registryParams).Return(templateList, nil).Times(1)
		ankaClient.EXPECT().RegistryRevert(registryParams.RegistryURL, templateList[0].ID).Return(nil).Times(0)
		ankaClient.EXPECT().RegistryPush(registryParams, pushParams).Return(nil).Times(1)

		mockui := packer.MockUi{}
		mockui.Say(fmt.Sprintf("Pushing template to Anka Registry as %s with tag %s", config.RemoteVM, config.Tag))
		mockui.Say(fmt.Sprintf("Found existing template %s on registry that matches name '%s'", templateList[0].ID, config.RemoteVM))

		assert.Equal(t, mockui.SayMessages[0].Message, "Pushing template to Anka Registry as foo with tag registry-push")
		assert.Equal(t, mockui.SayMessages[1].Message, "Found existing template foo_id on registry that matches name 'foo'")

		_, _, _, err = pp.PostProcess(context.Background(), ui, artifact)
		if err != nil {
			t.Fail()
		}
	})

}
