{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}in {{.request}}{{if .stream}},stream {{.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) ({{if .hasReply}}resp {{.response}},{{end}}err error) {
	// todo: add your logic here and delete this line
	
	return
}
