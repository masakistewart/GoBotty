package router

import (
	"github.com/masakistewart/GoBotty/logging"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
)

type Handle func(http.ResponseWriter, *http.Request, url.Values)

type Router struct {
	tree        *node
	rootHandler Handle
}

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]Handle
}

func New(rootHandler Handle) *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]Handle)}
	return &Router{tree: &node, rootHandler: rootHandler}
}

func (r *Router) Handle(method, path string, hanlder Handle) {
	if path[0] != '/' {
		panic("path has to start with a '/'.")
	}
	r.tree.addNode(method, path, hanlder)
}

func (this *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    logging.Info.Printf(path)

    data, err := ioutil.ReadFile(string(path))

    if err == nil {
        w.Write(data)
    } else {
        w.WriteHeader(404)
        w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
    }
}

// func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	req.ParseForm()
// 	params := req.Form
// 	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)

// 	if handler := node.methods[req.Method]; handler != nil {
// 		handler(w, req, params)
// 	} else {
// 		r.rootHandler(w, req, params)
// 	}
// }

func (n *node) addNode(method, path string, handler Handle) {
	components := strings.Split(path, "/")[1:]
	count := len(components)

	for {
		aNode, component := n.traverse(components, nil)
		if aNode.component == component && count == 1 {
			aNode.methods[method] = handler
			return
		}

		newNode := node{component: component, isNamedParam: false, methods: make(map[string]Handle)}

		if len(component) > 0 && component[0] == ':' {
			newNode.isNamedParam = true
		}

		if count == 1 {
			newNode.methods[method] = handler
		}

		aNode.children = append(aNode.children, &newNode)
		count--
		if count == 0 {
			break
		}
	}
}

func (n *node) traverse(components []string, params url.Values) (*node, string) {
	component := components[0]
	if len(n.children) > 0 {
		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.component[1:], component)
				}

				next := components[1:]
				if len(next) > 0 {
					return child.traverse(next, params)
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}
