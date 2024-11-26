package omicert

import (
	"embed"
	"fmt"
	"net/http"
	"os"
)

//go:embed credential/*
var certSource embed.FS

type Credential struct {
	CertFile string
	KeyFile  string
}

func ListenAndServeTLS(port string, credential *Credential) {
	fmt.Println("Starting HTTPS/2 server on port" + port)
	// 如果未提供证书，使用嵌入的证书
	if credential == nil {
		// 从 embed.FS 读取证书文件
		certFile, _ := certSource.ReadFile("credential/server.crt")

		keyFile, _ := certSource.ReadFile("credential/server.key")

		// 将嵌入的证书文件写入当前文件夹
		os.WriteFile("server.crt", certFile, 0644)

		os.WriteFile("server.key", keyFile, 0644)

		// 使用当前文件夹中的证书文件启动 HTTPS 服务器
		if err := http.ListenAndServeTLS(port, "server.crt", "server.key", nil); err != nil {
			fmt.Println("Server error:", err)
		}
	} else {
		// 使用传入的证书文件启动 HTTPS 服务器
		if err := http.ListenAndServeTLS(port, credential.CertFile, credential.KeyFile, nil); err != nil {
			fmt.Println("Server error:", err)
		}
	}
}
