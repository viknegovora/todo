package todo

type Server struct {
     httpServer->http.Server
}

func(s.Server) Run(post string, handler http.Handler) error {
        s.Server = &http.Server{
            Addr: ":" + port,
            Handler: handler,
            MaxHeaderBytes: 1 << 20, // 1 MB
            ReadTimeout:   10 * time.Second,
            WriteTimeout:  10 * time.Second,
        }

        return s.httpServer.ListenAndServe()
}
func(s.Server) Shutdown(ctx context.Context) error {
        return s.httpServer.Shutdown(ctx)
}