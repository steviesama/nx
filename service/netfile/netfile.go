// nx/service/netfile facilitates file requests for storage and fetching between
// a TCP Listener server and any number of TCP Clients that may make requests.
// The netfile cli used this package to provide its functionality once the cli
// commands have been parsed.
package netfile

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

// DataDir is the Linux location which netfile stores it's data when
// netfile server is run.
const DataDir string = "/var/lib/netfile"
const FilesDir string = "/var/lib/netfile/files"
const TempDir string = "/var/lib/netfile/temp"

const BufferSize int = 1024

// CommandHandlerFunc is a type that is used to describe the functions that will
// handle the various commands netfile server can process.
type CommandHandlerFunc func(*bufio.ReadWriter)

// commandHandlers is a pool of handlers for associated commands.
var commandHandlers map[string]CommandHandlerFunc

// A mutex to protect async access to commandHandlers.
var mtxCommandHandlers sync.RWMutex

func init() {
	commandHandlers = make(map[string]CommandHandlerFunc)
	AddCommandHandler("client.fetch", handleClientFetch)
}

func handleClientFetch(rw *bufio.ReadWriter) {
	fmt.Println("Before readMsg()")
	msg, readErr := readMsg(rw)

	switch {
	case readErr == io.EOF:
		fmt.Println("netfile client closed connection.")
		return
	case readErr != nil:
		// fmt.Printf("netfile.ReadMsg() error: %s\n", readErr.Error())
		return
	}

	var noFileFlush = func() {
		rw.WriteString("server.fetch.nofile\n")
		rw.Flush()
	}

	// Ensure the current working directory is the server data files directory.
	os.Chdir(FilesDir)

	fileInfo, statErr := os.Stat(msg)

	if statErr != nil {
		fmt.Printf("netfile.handleClientFetch error: %s\n", statErr.Error())
		noFileFlush()
		return
	}

	if fileInfo.IsDir() {
		noFileFlush()
		return
	}

	fmt.Printf("Sending: '%s'\n", fileInfo.Name())

	// Write the next 2 messages and send.
	rw.WriteString("server.fetch.file\n")
	rw.WriteString(fmt.Sprintf("%d\n", fileInfo.Size()))
	rw.Flush()

	file, openErr := os.OpenFile(fileInfo.Name(), os.O_RDONLY, 0755)
	defer file.Close()

	if openErr != nil {
		noFileFlush()
		return
	}

	data := make([]byte, BufferSize)

	for {
		n, readErr := file.Read(data)

		_, _ = io.CopyN(rw, bytes.NewReader(data), int64(n))

		if readErr != nil {
			fmt.Printf("handleClientFetch(): netfile fetch file read error: %s\n", readErr.Error())
			break
		}
	}
}

// AddCommandHandler takes a command name and a command handler and add the handler
// to a command handler map using the command name as a key. When the command name
// command is received...the associated command handler will be executed.
// AddCommandHandler is thread safe.
func AddCommandHandler(cmd string, commandHandler CommandHandlerFunc) {
	mtxCommandHandlers.Lock()
	commandHandlers[cmd] = commandHandler
	mtxCommandHandlers.Unlock()
}

func onHandleCommand(cmd string, conn net.Conn) {
	mtxCommandHandlers.RLock()
	handleCommand, ok := commandHandlers[cmd]
	mtxCommandHandlers.RUnlock()

	if !ok {
		fmt.Printf("The command '%s' is not registered.\n", cmd)
		return
	}

	handleCommand(NewConnReadWriter(conn))
}

// ensureDataDir makes sure that the netfile data directory exists before the
// server starts listening.
// It returns an error if the dir doesn't exist and it can't create it.
// Note: If not running netfile as CLI...ensure the DataDir structure and
// 			 permissions are set correctly before the app that uses this code is run.
func ensureDataDir() error {
	if _, err := os.Stat(DataDir); os.IsNotExist(err) {
		fmt.Println("netfile data directory doesn't exist...")

		curUser := os.Getenv("USER")
		if curUser != "root" {
			return errors.New("unable to create netfile data directory...please run with sudo")
		}

		fmt.Printf("Creating '%s'...\n", DataDir)
		var mkdirErr error
		mkdirErr = os.MkdirAll(DataDir, os.ModePerm)
		if mkdirErr != nil {
			return errors.New(fmt.Sprintf("mkdir error: %s\n", mkdirErr.Error()))
		}
		fmt.Printf("Creating '%s'...\n", FilesDir)
		mkdirErr = os.MkdirAll(FilesDir, os.ModePerm)
		if mkdirErr != nil {
			return errors.New(fmt.Sprintf("mkdir error: %s\n", mkdirErr.Error()))
		}
		fmt.Printf("Creating '%s'...\n", TempDir)
		mkdirErr = os.MkdirAll(TempDir, os.ModePerm)
		if mkdirErr != nil {
			return errors.New(fmt.Sprintf("mkdir error: %s\n", mkdirErr.Error()))
		}

		fmt.Println("netfile data directory created.")

		user := os.Getenv("SUDO_USER")
		uid, _ := strconv.Atoi(os.Getenv("SUDO_UID"))
		gid, _ := strconv.Atoi(os.Getenv("SUDO_GID"))

		fmt.Println("Currently only the user that first runs netfile on a system will have access to store data.")
		fmt.Printf("Setting owner of data directory to %s...", user)

		os.Chown(DataDir, uid, gid)
		os.Chown(FilesDir, uid, gid)
		os.Chown(TempDir, uid, gid)

		fmt.Println("finished.")

		return errors.New("Please run netfile server again normally.")
	}

	return nil
}

// ServerListen takes a host address and port number which it uses to dial a tcp
// connection.
// It returns an error if a successful connection is not made...or it loops listening
// for client connections otherwise until killed.
func ServerListen(host string, port int) error {
	if dataDirErr := ensureDataDir(); dataDirErr != nil {
		return dataDirErr
	}
	addrString := fmt.Sprintf("%s:%d", host, port)
	server, serverErr := net.Listen("tcp", addrString)

	if serverErr != nil {
		return errors.New(fmt.Sprintf("netfile server error listening: %s\n", serverErr.Error()))
	}
	defer server.Close()

	fmt.Printf("netfile server listening on port: %d...\n", port)

	for {
		conn, connErr := server.Accept()

		if connErr != nil {
			return errors.New(fmt.Sprintf("netfile server client connection error: %s\n", connErr.Error()))
		}

		fmt.Printf("Client connected from: '%s:%s'\n", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
		go func() {
			defer conn.Close()
			// Send client ready message
			SendMsg(conn, "server.ready")
			for {
				cmd, readErr := ReadMsg(conn)

				switch {
				case readErr == io.EOF:
					fmt.Println("netfile server closed connection.")
					return
				case readErr != nil:
					fmt.Printf("netfile.ReadMsg() error: %s\n", readErr.Error())
					return
				}

				if cmd == "client.quit" {
					fmt.Println("client.quit msg received...")
					break
				}

				onHandleCommand(cmd, conn)
			}
		}()
	}

	return nil
}

func NewConnReadWriter(conn net.Conn) *bufio.ReadWriter {
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
}

func SendMsg(conn net.Conn, msg string) error {
	rw := NewConnReadWriter(conn)

	// n, writeErr := conn.Write([]byte(msg + "\n"))
	n, writeErr := rw.WriteString(msg + "\n")

	if writeErr != nil {
		return errors.New(fmt.Sprintf("netfile.SendMsg() write error: %s\n", writeErr.Error()))
	}

	flushErr := rw.Flush()

	if flushErr != nil {
		return errors.New(fmt.Sprintf("netfile.SendMsg() flush error: %s\n", flushErr.Error()))
	}

	fmt.Printf("netfile.SendMsg('%s') %d byte(s) written...\n", msg, n)

	return nil
}

func readMsg(rw *bufio.ReadWriter) (string, error) {
	delim := byte('\n')
	command, ioErr := rw.ReadString(delim)
	command = strings.Trim(command, string(delim)+" ")

	if ioErr != nil {
		return "", errors.New(fmt.Sprintf("netfile.ReadMsg() network io error: %s\n", ioErr.Error()))
	}

	return command, nil
}

func ReadMsg(conn net.Conn) (string, error) {
	return readMsg(NewConnReadWriter(conn))
}
