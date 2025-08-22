package console

import "flag"

type Flags struct {
    Mode     string
    Host     string
    Port     string
    Database string
    Protocol string
}

func NewFlags() *Flags {
    var mode, host, port, database, protocol string

    flag.StringVar(&mode, "mode", "debug", "Режим работы сервера (debug, release)")
    flag.StringVar(&host, "host", "127.0.0.1", "Хост для запуска сервера")
    flag.StringVar(&port, "port", "8080", "Порт на котором будет работать сервер")
    flag.StringVar(&database, "db", "test.db", "Путь к базе данных")
    flag.StringVar(&protocol, "protocol", "http", "Протокол (http, https, и т.д)")

    flag.Parse()

    return &Flags{
        Mode:     mode,
        Database: database,
        Host:     host,
        Port:     port,
        Protocol: protocol,
    }
}
