package connections

type Connection struct {
	UUID        string
	SendChannel chan []byte
}

var Connections = make(map[string]*Connection)

func AddConnection(uuid string, conn *Connection) {
	Connections[uuid] = conn
}

func RemoveConnection(uuid string) {
	delete(Connections, uuid)
}

func SendToAll(msg []byte) {
	for _, conn := range Connections {
		func() {
			defer func() {
				if r := recover(); r != nil {
					RemoveConnection(conn.UUID)
				}
			}()
			conn.SendChannel <- msg
		}()
	}
}
