package server

import (
	"bytes"
	"vm/emulator"
	"vm/game"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"reflect"
)

type Server struct {
	Listener net.Listener
}

type Pool struct {
	Clients     map[*Client]bool
	Register    chan *Client
	Unregister  chan *Client
	NewByteCode chan *Client
	Win         chan *Client
	Loose       chan *Client
	Message     chan Message_t
}

type Message_t struct {
	Message string
	Client  *Client
}

type Client struct {
	Pool           *Pool
	Conn           *net.Conn
	NumberBytecode int
	Score          int
	BytecodeTMP    []byte
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

func (client *Client) Win() {
	if client.Score > client.NumberBytecode {
		if _, err := (*client.Conn).Write([]byte(fmt.Sprintf("GG ! You can validate this flag: %v\n", FLAG))); err != nil {
			return
		}
	}

}

func (c *Client) Loose(p *Pool) {
	if _, err := (*c.Conn).Write([]byte("Tu es un kikoo")); err != nil {
		return
	}
}

func (client *Client) NewByteCode() {
	fmt.Println(client.Score)
	if client.Score <= client.NumberBytecode {
		generated_bytecode := game.GenerateBytecode()
		generated_bytecode2 := generated_bytecode[rand.Intn(len(generated_bytecode))]

		if _, err := (*client.Conn).Write([]byte(fmt.Sprintf("%v\n", generated_bytecode2))); err != nil {
			return
		}

		client.BytecodeTMP = generated_bytecode2

	} else {
		client.Win()
	}
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

			client.NewByteCode()

		case client := <-p.Unregister:
			fmt.Println("Issoufre au chocolat")
			fmt.Println("Client close")

			delete(p.Clients, client)

			(*client.Conn).Close()
		case client := <-p.Win:
			client.Win()
		case message := <-p.Message:
			var bytecode_input map[emulator.Register]int

			if err := json.Unmarshal([]byte(bytes.Trim([]byte(message.Message), "\x00")), &bytecode_input); err != nil {
				if _, err := (*message.Client.Conn).Write([]byte(fmt.Sprintf("%v\n", err.Error()))); err != nil {
					break
				}
				break
			}

			if message.Client.BytecodeTMP != nil {
				runtime := emulator.Exec(message.Client.BytecodeTMP)
				fmt.Println(bytecode_input, runtime.Registers)
				if ok := reflect.DeepEqual(bytecode_input, runtime.Registers); ok {
					message.Client.Score += 1
					message.Client.NewByteCode()
				} else {
					message.Client.Loose(p)
				}
			}

		}
	}
}

func InitPool() *Pool {
	return &Pool{
		Clients:     make(map[*Client]bool),
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Message:     make(chan Message_t),
		NewByteCode: make(chan *Client),
		Win:         make(chan *Client),
		Loose:       make(chan *Client),
	}
}

func (client *Client) Read() {

	defer func() {
		client.Pool.Unregister <- client
	}()

	for {
		var data []byte = make([]byte, 1024)
		if _, err := (*client.Conn).Read(data); err != nil {
			fmt.Println("error receiving")
			break
		}

		var message Message_t = Message_t{
			Message: string(data),
			Client:  client,
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
			Pool:           pool,
			Conn:           &conn,
			NumberBytecode: rand.Intn(30),
			Score:          0,
		}

		pool.Register <- &client

		go client.Read()
	}

}
