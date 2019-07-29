// nx/service/websock will hold the translation of what is now scattered all
// around other nx projects in the form of wshub.go, wsmsg.go, wsclientmsg.go,
// etc. It will be isolated into this package and setup to use inversion of
// control in order to communicate with other packages it needs to at the
// discretion of the caller.
package websock
