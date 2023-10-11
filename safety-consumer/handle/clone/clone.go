package clone

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/goodrain/rainbond-task-plug/util"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/client"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"net/http"
	"os/exec"
	"os/user"
	"runtime"

	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func GitClone(csi model.CodeDetectionModel, sourceDir string, timeout int, ctx context.Context) (*git.Repository, string, error) {
	logrus.Infof("begin clone %v", csi.RepositoryURL)
	if !strings.HasSuffix(csi.RepositoryURL, ".git") {
		csi.RepositoryURL = csi.RepositoryURL + ".git"
	}
	flag := true
Loop:
	ep, err := transport.NewEndpoint(csi.RepositoryURL)
	if err != nil {
		return nil, "", err
	}
	opts := &git.CloneOptions{
		URL:               csi.RepositoryURL,
		SingleBranch:      true,
		Progress:          os.Stdout,
		Tags:              git.NoTags,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Depth:             1,
	}
	if csi.Branch != "" {
		opts.ReferenceName = getBranch(csi.Branch)
	}
	var rs *git.Repository
	if ep.Protocol != "ssh" {
		customClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Minute * time.Duration(timeout),
		}
		if strings.Contains(csi.RepositoryURL, "github.com") && os.Getenv("GITHUB_PROXY") != "" {
			proxyURL, err := url.Parse(os.Getenv("GITHUB_PROXY"))
			if err == nil {
				customClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL), TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
			} else {
				logrus.Error(err)
			}
		}
		if csi.User != "" && csi.Password != "" {
			httpAuth := &githttp.BasicAuth{
				Username: csi.User,
				Password: csi.Password,
			}
			opts.Auth = httpAuth
		}
		client.InstallProtocol("https", githttp.NewClient(customClient))
		defer func() {
			client.InstallProtocol("https", githttp.DefaultClient)
		}()
		rs, err = git.PlainCloneContext(ctx, sourceDir, false, opts)
	}
	if err != nil {
		if err == transport.ErrAuthenticationRequired {
			errMsg := fmt.Sprintf("拉取代码发生错误，代码源需要授权访问。")
			logrus.Error(errMsg, map[string]string{"step": "clone-code", "status": "failure"})
			return rs, errMsg, err
		}
		if err == transport.ErrAuthorizationFailed {
			errMsg := fmt.Sprintf("拉取代码发生错误，代码源鉴权失败。")
			logrus.Error(errMsg, map[string]string{"step": "clone-code", "status": "failure"})
			return rs, errMsg, err
		}
		if err == transport.ErrRepositoryNotFound {
			errMsg := fmt.Sprintf("拉取代码发生错误，仓库不存在。")
			logrus.Error(fmt.Sprintf("拉取代码发生错误，仓库不存在。"), map[string]string{"step": "clone-code", "status": "failure"})
			return rs, errMsg, err
		}
		if err == transport.ErrEmptyRemoteRepository {
			errMsg := fmt.Sprintf("拉取代码发生错误，远程仓库为空。")
			logrus.Error(errMsg, map[string]string{"step": "clone-code", "status": "failure"})
			return rs, errMsg, err
		}
		if err == plumbing.ErrReferenceNotFound || strings.Contains(err.Error(), "couldn't find remote ref") {
			errMsg := fmt.Sprintf("代码分支(%s)不存在。", csi.Branch)
			logrus.Error(errMsg, map[string]string{"step": "clone-code", "status": "failure"})
			return rs, errMsg, fmt.Errorf("branch %s is not exist", csi.Branch)
		}
		if strings.Contains(err.Error(), "ssh: unable to authenticate") {
			if flag {
				flag = false
				goto Loop
			}
			errMsg := fmt.Sprintf("远程代码库需要配置SSH Key。")
			logrus.Error(errMsg, map[string]string{"step": "clone-code", "status": "failure"})
			return rs, errMsg, err
		}
		if strings.Contains(err.Error(), "context deadline exceeded") {
			errMsg := fmt.Sprintf("获取代码超时")
			logrus.Error(errMsg, map[string]string{"step": "clone-code", "status": "failure"})
			return rs, errMsg, err
		}
	}
	return rs, "", err
}

func getBranch(branch string) plumbing.ReferenceName {
	if strings.HasPrefix(branch, "tag:") {
		return plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", branch[4:]))
	}
	return plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch))
}

func GetPrivateFile(tenantID string) string {
	home, _ := Home()
	if home == "" {
		home = "/root"
	}
	if ok, _ := util.FileExists(path.Join(home, "/.ssh/"+tenantID)); ok {
		return path.Join(home, "/.ssh/"+tenantID)
	}
	if ok, _ := util.FileExists(path.Join(home, "/.ssh/builder_rsa")); ok {
		return path.Join(home, "/.ssh/builder_rsa")
	}
	return path.Join(home, "/.ssh/id_rsa")

}

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}
