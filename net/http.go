package net

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/Pangjiping/terrafmtter/util"
	"golang.org/x/net/html"
)

const latest = "latest"

// GetDocFromGithubV2 fetch markdown doc from github.
// It will be saved in a tmp dir.
func GetDocFromGithubV2(version, file string, isResource bool) error {
	resp, err := http.Get(getHttpAddrV2(version, file, isResource))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code is %v. Please network connection and resource/data %s by version %s",
			resp.StatusCode, file, version)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	h := node2html(doc)
	filePath := fmt.Sprintf("%s%s.md", util.FILE_LOC_PREFIX, file)
	return util.ConvertHtml2Md(filePath, h)
}

// getHttpAddrV2 build http url to markdown.
func getHttpAddrV2(version, file string, isResource bool) string {
	if isResource {
		if version == latest {
			return fmt.Sprintf(`https://github.com/aliyun/terraform-provider-alicloud/blob/master/website/docs/r/%s.html.markdown`, file)
		} else {
			return fmt.Sprintf(`https://github.com/aliyun/terraform-provider-alicloud/blob/v%s/website/docs/r/%s.html.markdown`, version, file)
		}
	} else {
		if version == latest {
			return fmt.Sprintf(`https://github.com/aliyun/terraform-provider-alicloud/blob/master/website/docs/d/%s.html.markdown`, file)
		} else {
			return fmt.Sprintf(`https://github.com/aliyun/terraform-provider-alicloud/blob/v%s/website/docs/d/%s.html.markdown`, version, file)
		}
	}
}

// clear_dom clean html node recursively.
func clear_dom(pn *html.Node, isgb2312 bool) error {
	for nd := pn.FirstChild; nd != nil; {
		switch nd.Type {
		case html.ElementNode:
			tn := strings.ToLower(nd.Data)
			if tn == "script" || tn == "style" {
				tmp := nd
				nd = tmp.NextSibling
				pn.RemoveChild(tmp)
			} else if tn == "a" {
				nd = nd.NextSibling
			} else if tn == "span" {
				tmp := nd
				nd = nd.NextSibling
				clear_dom(tmp, isgb2312)
			} else {
				nd = nd.NextSibling
			}
		case html.CommentNode:
			tmp := nd
			nd = tmp.NextSibling
			pn.RemoveChild(tmp)
		case html.TextNode:
			nd = nd.NextSibling
		default:
			nd = nd.NextSibling
		}
	}
	return nil
}

// node2html convert html nodes to html.
func node2html(n *html.Node) string {
	var buf = bytes.NewBuffer([]byte{})
	html.Render(buf, n)
	return buf.String()
}

// node2text convert html nodes to text.
func node2text(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	} else if node.FirstChild != nil {
		var buf bytes.Buffer
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			buf.WriteString(node2text(c))
		}
		return buf.String()
	}
	return ""
}

// GetCodeFromGithub fetch html page from alicloud provider github.
// Then clean html page and parse it to string.
// Deprecated
func GetCodeFromGithub(version, file string) (string, error) {
	resp, err := http.Get(getHttpAddr(version, file))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http server error, status code is %v", resp.StatusCode)
	}

	// parse html
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	clear_dom(doc, false)
	str := node2text(doc)
	return str, nil
}

// GetDocFromGithub fetch html page from alicloud provider docs.
// Then clean html and parse it to string.
// Deprecated
func GetDocFromGithub(version, file string, isResource bool) (string, error) {
	resp, err := http.Get(getHttpAddrV2(version, file, isResource))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http server error, status code is %v", resp.StatusCode)
	}

	// parse html
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	clear_dom(doc, false)
	str := node2text(doc)
	return str, nil
}

// getHttpAddr build http url to go code.
// Deprecated
func getHttpAddr(version, file string) string {
	return fmt.Sprintf(`https://github.com/aliyun/terraform-provider-alicloud/blob/v%s/alicloud/%s`, version, file)
}
