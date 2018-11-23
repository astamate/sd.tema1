# sd.tema1
Cerințe generale:

Sistem de tipul Client - Server (mai mulți clienți și un server).

Metoda de procesare a datelor pe Server trebuie să fie făcută concurent - folosind go routines. 

Există un fișier de configurare în care există parametrii inițiali ai programului. 
Ex: câte elemente are array-ul de date pe care îl poate trimite clientul, de câte ori se apelează o routină go. Decideți voi ce anume puteți include în acest fișier.

Trebuie să existe mesaje între client și server de tipul: 
Client <Nume> Conectat.
Client <Nume> a facut request cu datele: <date>.
Server a primit requestul.
Server proceseaza datele.
Server trimite <raspuns> catre client.
Client <Nume> a primit raspunsul: <raspuns>.

Clientul trimite către server un array de strings. Un string poate conține atât caractere, cât și cifre, amestecate. 
Serverul returnează către client numărul de numere care sunt pătrate perfecte.
Exemplu: abd4g5, 1sdf6fd, fd2fdsf5 => 2 pătrate perfecte: 16 din 1sdf6fd, 25 dinfd2fdsf5.

go run server.go
telnet localhost 8080
abd4g5, 1sdf6fd, fd2fdsf5
