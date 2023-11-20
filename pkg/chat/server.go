package chat

type Server struct {
	Clients []*Client
	mu sync.mutex
}

func NewServer() *Server {
	return &Server{
		Clients: make([]*Client, 0),
		mu: sync.Mutex{},
	}
}

func (h *Server) AddClient(c *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Clients = append(s.Clients, c)

	s.broadcast(c, 1)	
}

func (h *Server) RemoveClient(c *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, client := range s.Clients {
		if client == c {
			s.Clients = append(s.Clients[:i], s.Clients[i+1:]...)
			s.broadcast(c, 2)
			return
		}
	}
}

func (h *Server) broadcast(c *Client, t int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	for _, client := range s.Clients {
		client.WriteMessage(c.Name, t)
	}
}

func (s *Server) Serve() func(connection *websocket.Conn) {
	return func(connection *websocket.Conn) {
		client := NewClient(connection.Request().Header.Get("Username"), connection, s)
		s.Join(client)

		client.Read()
	}
}