package server

import (
	"fmt"
	"net"
)

type Server struct {
	Listener net.Listener
}

type Pool struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Message    chan Message_t
}

type Message_t struct {
	Message string
}

type Client struct {
	Pool *Pool
	Conn *net.Conn
}

func NewServer(port string) (*Server, error) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, err
	}

	server := Server{
		Listener: listener,
	}

	return &server, nil

}

func (p *Pool) RunPool() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("New client")
			if _, err := (*client.Conn).Write([]byte("[ - Memulator VM bytecode tester - ]\n")); err != nil {
				return
			}
		case client := <-p.Unregister:
			delete(p.Clients, client)

			fmt.Println("Client close")
		case message := <-p.Message:
			fmt.Println("New message: ")
			fmt.Println(message)
		}
	}
}

func InitPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Message:    make(chan Message_t),
	}
}

func (client *Client) Read() {

	defer func() {
		client.Pool.Unregister <- client
		(*client.Conn).Close()
	}()

	for {
		var data []byte = make([]byte, 1024)
		if _, err := (*client.Conn).Read(data); err != nil {
			fmt.Println("error receiving")
			break
		}

		var message Message_t = Message_t{
			Message: string(data),
		}

		client.Pool.Message <- message
	}
}

func (serve *Server) Run() {
	defer serve.Listener.Close()

	pool := InitPool()

	go pool.RunPool()

	fmt.Println("Server started at 1337 port")
	for {
		conn, err := serve.Listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		client := Client{
			Pool: pool,
			Conn: &conn,
		}

		pool.Register <- &client

		go client.Read()
	}

}
