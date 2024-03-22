package gee

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

//封装自己的一个Web框架玩玩
//建立路由表
//访问指定路径时映射执行对应的处理函数
type HandlerFunc func(c *Context)

type Engine struct {
	//route map[string]HandlerFunc
	router *router
	*RouterGroup
	groups        []*RouterGroup     //store all
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render

}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc //support middleware
	parent      *RouterGroup  //support nesting 支持嵌套
	engine      *Engine       //all groups share a Engine instance
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
//定义组以创建新的路由器组
//记住所有组共享同一个引擎实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		//目的是为了支持嵌套
		parent: group,
		engine: engine,
	}
	fmt.Println(*newGroup.parent)
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
func (engine *Engine) Run(addr string) (err error) {
	//第二次参数本应该是接口类型的，为什么传结构体类型的也行呢？
	//因为实现了接口方法的 struct 都可以强制转换为接口类型
	//handler := (http.Handler)(engine) // 手动转换为接口类型
	return http.ListenAndServe(addr, engine)

}

// Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	//遍历路由组
	for _, group := range engine.groups {
		//根据请求的URL匹配路由组的前缀，说白了就是匹配对应前缀的路由组
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			//匹配到对应前缀的路由组将中间件取出加入到处理函数中
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	//这几步综合起来就是:1.我们把请求和响应的信息封装在Context中然后我们new了一个处理请求和响应的Context实例
	//2.中间件为其中一部分，将中间件注册到handlers处理切片中
	//3.根据请求的路由调用对应的处理函数
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	//在group.prefix表示的字符串和relativePath字符串之间加上 /
	absolutePath := path.Join(group.prefix, relativePath)
	//  http.StripPrefix函数的作用之一，就是在将请求定向到你通过参数指定的请求处理处之前，将特定的prefix从URL中过滤出去
	// FileServer 已经明确静态文件的根目录在"/tmp"，但是我们希望URL以"/tmpfiles/"开头
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

//
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	//一个帮助程序，如果变量初始化失败程序会崩溃
	//解析模版建议在初始化时加载一次，否则每次调用时都加载会比较浪费性能
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
