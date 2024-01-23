package web

// Middleware 函数式的责任链模式
// 函数式的洋葱模式
type Middleware func(next HandleFunc) HandleFunc

// AOP 方案在不同的框架，不同的语言里面都有不同的叫法
// Middleware, Handler, Chain, Filter, Filter-Chain
// Interceptor, Wrapper
