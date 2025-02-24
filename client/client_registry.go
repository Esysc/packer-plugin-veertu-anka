package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/packer-plugin-sdk/net"
)

// Run command against the registry
type RegistryParams struct {
	RegistryName string
	RegistryURL  string
	NodeCertPath string
	NodeKeyPath  string
	CaRootPath   string
	IsInsecure   bool
	HostArch     string
}

// https://ankadocs.veertu.com/docs/anka-virtualization/command-reference/#registry-list
type RegistryListResponse struct {
	Latest string `json:"latest"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

func (c *AnkaClient) RegistryList(registryParams RegistryParams) ([]RegistryListResponse, error) {
	var response []RegistryListResponse

	output, err := runRegistryCommand(registryParams, "list")
	if err != nil {
		return nil, err
	}
	if output.Status != "OK" {
		log.Print("Error executing registry list command: ", output.ExceptionType, " ", output.Message)
		return nil, fmt.Errorf(output.Message)
	}

	err = json.Unmarshal(output.Body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

type RegistryRemote struct {
	Default bool   `json:"default"`
	Host    string `json:"host"`
	Scheme  string `json:"scheme"`
	Port    string `json:"port"`
}

type RegistryListReposResponse struct {
	Default string
	Remotes map[string]RegistryRemote
}

type RegistryRemoteArm64 struct {
	Default bool   `json:"default"`
	Url     string `json:"url"`
	Name    string `json:"name"`
}

type RegistryListReposResponseArm64 struct {
	Remotes []RegistryRemoteArm64
}

func (c *AnkaClient) RegistryListRepos() (RegistryListReposResponse, error) {
	var response RegistryListReposResponse

	output, err := runRegistryCommand(RegistryParams{}, "list-repos")
	if err != nil {
		return response, err
	}
	if output.Status != "OK" {
		log.Print("Error executing 'registry list-repos' command: ", output.ExceptionType, " ", output.Message)
		return response, fmt.Errorf(output.Message)
	}

	err = json.Unmarshal(output.Body, &response.Remotes)
	if err != nil {
		return response, err
	}

	for name, remote := range response.Remotes {
		if remote.Default {
			response.Default = name
		}
	}

	return response, nil
}

func (c *AnkaClient) RegistryListReposArm64() (RegistryListReposResponse, error) {
	var response RegistryListReposResponse
	var responseArm64 RegistryListReposResponseArm64

	output, err := runRegistryCommand(RegistryParams{}, "list-repos")
	if err != nil {
		return response, err
	}
	if output.Status != "OK" {
		log.Print("Error executing 'registry list-repos' command: ", output.ExceptionType, " ", output.Message)
		return response, fmt.Errorf(output.Message)
	}

	// Refactor the object in Anka 2 format.
	err = json.Unmarshal(output.Body, &responseArm64.Remotes)
	if err != nil {
		return response, err
	}
	tmpBody := make(map[string]RegistryRemote)
	for _, remote := range responseArm64.Remotes {
		u, err := url.Parse(remote.Url)
		if err != nil {
			return response, err
		}

		s := strings.Split(u.Host, ":")
		a := RegistryRemote{
			Scheme:  u.Scheme,
			Default: remote.Default,
			Host:    s[0],
			Port:    s[1],
		}
		tmpBody[remote.Name] = a
	}
	response = RegistryListReposResponse{Remotes: tmpBody}

	for name, remote := range response.Remotes {
		if remote.Default {
			response.Default = name
		}
	}

	return response, nil
}

type RegistryPullParams struct {
	VMID   string
	Tag    string
	Local  bool
	Shrink bool
}

func (c *AnkaClient) RegistryPull(registryParams RegistryParams, pullParams RegistryPullParams) error {
	cmdArgs := []string{"pull"}

	if pullParams.Tag != "" {
		cmdArgs = append(cmdArgs, "--tag", pullParams.Tag)
	}

	if pullParams.Local {
		cmdArgs = append(cmdArgs, "--local")

		if pullParams.Shrink {
			cmdArgs = append(cmdArgs, "--shrink")
		}
	}

	cmdArgs = append(cmdArgs, pullParams.VMID)

	output, err := runRegistryCommand(registryParams, cmdArgs...)
	if err != nil {
		return err
	}
	if output.Status != "OK" {
		return fmt.Errorf(output.Message)
	}

	return nil
}

// https://ankadocs.veertu.com/docs/anka-virtualization/command-reference/#registry-push
type RegistryPushParams struct {
	VMID        string
	Tag         string
	Description string
	RemoteVM    string
	Local       bool
	Force       bool
}

func (c *AnkaClient) RegistryPush(registryParams RegistryParams, pushParams RegistryPushParams) error {
	cmdArgs := []string{"push"}

	if pushParams.Tag != "" {
		cmdArgs = append(cmdArgs, "--tag", pushParams.Tag)
	}

	if pushParams.Description != "" {
		cmdArgs = append(cmdArgs, "--description", pushParams.Description)
	}

	if pushParams.RemoteVM != "" {
		cmdArgs = append(cmdArgs, "--remote-vm", pushParams.RemoteVM)
	}

	if pushParams.Local {
		cmdArgs = append(cmdArgs, "--local")
	}

	if pushParams.Force {
		cmdArgs = append(cmdArgs, "--force")
	}

	cmdArgs = append(cmdArgs, pushParams.VMID)

	output, err := runRegistryCommand(registryParams, cmdArgs...)
	if err != nil {
		return err
	}
	if output.Status != "OK" {
		return fmt.Errorf(output.Message)
	}

	return nil
}

// https://ankadocs.veertu.com/docs/anka-build-cloud/working-with-registry-and-api/#revert
func (c *AnkaClient) RegistryRevert(url string, id string) error {
	response, err := registryRESTRequest("DELETE", fmt.Sprintf("%s/registry/revert?id=%s", url, id), nil)
	if err != nil {
		return err
	}
	if response.Status != statusOK {
		return fmt.Errorf("failed to revert VM on registry: %s", response.Message)
	}

	return nil
}

func registryRESTRequest(method string, url string, body io.Reader) (MachineReadableOutput, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return MachineReadableOutput{}, err
	}

	httpClient := net.HttpClientWithEnvironmentProxy()

	resp, err := httpClient.Do(request)
	if err != nil {
		return MachineReadableOutput{}, err
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return MachineReadableOutput{}, err
		}

		return parseOutput(body)
	}

	return MachineReadableOutput{}, fmt.Errorf("unsupported http response code: %d", resp.StatusCode)
}
