package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"strconv"
	"strings"
	"unicode"
)

func isInt(s string) int {
	var a = 0
	for _, c := range s {
		if unicode.IsDigit(c) {
			if a > 0 {
				a = a*10 + int(c-48)
			} else {
				a = a + int(c-48)
			}
		}
	}
	return a
}

func IsSquare(value int) bool {
	var a = float64(value)
	var int_root int = int(math.Sqrt(a))
	if (int_root * int_root) == int(a) {
		return true
	} else {
		return false
	}
}

func check(err error, message string) {
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", message)
}

func main() {
	data, err := ioutil.ReadFile("config.txt")
	if err != nil {
		fmt.Println("Cannot read from file ", err)
		return
	}

	var nrCifre, _ = strconv.Atoi(string(data))

	clientCount := 0
	allClients := make(map[net.Conn]int)

	ln, err := net.Listen("tcp", ":8080")
	check(err, "Server is ready.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		allClients[conn] = clientCount
		fmt.Printf("Client %d Conectat.\n", allClients[conn])

		clientCount += 1

		go func() {
			reader := bufio.NewReader(conn)

			for {

				incoming, err := reader.ReadString('\n')

				if err != nil {
					fmt.Printf("Clientul a fost deconectat.")
					break
				}

				fmt.Printf("Client %d a facut request cu datele: %s.\n", allClients[conn], incoming)
				conn.Write([]byte("Server a primit requestul.\n"))

				incoming = incoming[0 : len(incoming)-2]

				vect := strings.Split(incoming, ",")
				var patrare []int
				conn.Write([]byte("Server proceseaza datele.\n"))
				for i := 0; i < len(vect); i++ {
					// elem, _ := strconv.Atoi(vect[i])
					if len(vect[i]) < nrCifre {
						var temp = isInt(vect[i])
						if temp != 0 && IsSquare(temp) == true {
							patrare = append(patrare, temp)
						}
					} else {
						conn.Write([]byte("Eroare!Ati introdus un numar mai mare de " + strconv.Itoa(nrCifre) + " elemente.\n"))
						break
					}

				}
				nume := strconv.Itoa(allClients[conn])
				fmt.Printf("Server trimite:au fost gasit %d patrare perfecte:%v", len(patrare), patrare)
				var patrareperfecte = strings.Trim(strings.Replace(fmt.Sprint(patrare), " ", ",", -1), "[]")

				conn.Write([]byte("Client" + nume + " a primit raspunsul: Au fost gasite " + string(len(patrare)+48) + " patrare perfecte:" + patrareperfecte + "\n"))
			}
		}()
	}
}
