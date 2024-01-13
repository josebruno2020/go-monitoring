package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const numberOfMonitoring = 3
const delay = 5 * time.Second

func main() {
	showIntro()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Exibir Logs....")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "José Bruno"
	versao := 1.0

	fmt.Println("Olá, sr(a).", name)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println()
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println()
	return command
}

func initMonitoring() {
	fmt.Println("Monitorando....")

	sites := readWebSitesFile()

	fmt.Printf("\n\nTenho %d sites para monitorar! \nCapacidade de %d no slice \n\n\n", len(sites), cap(sites))

	for i := 0; i < numberOfMonitoring; i++ {
		fmt.Println()
		fmt.Printf("===== TENTATIVA %d de %d =====\n", i+1, numberOfMonitoring)
		for _, site := range sites {
			testWebSite(site)
		}
		time.Sleep(delay)
	}
}

func readWebSitesFile() []string {
	file, err := os.Open("./data/sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	var sites []string

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(line))

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func testWebSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao visitar o site:", err)
		os.Exit(-1)
	}

	if res.StatusCode == 200 {
		fmt.Printf("Site %s funcionando legal \n", site)
	} else {
		fmt.Printf("Site %s retornou com Status Code: %d. \n", site, res.StatusCode)
	}

	saveSiteLog(site, res.StatusCode == 200)
}

func saveSiteLog(site string, status bool) {
	file, err := os.OpenFile("./data/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	log := fmt.Sprintf("site %s está %t \n", site, status)

	file.WriteString(log)

	file.Close()
}

func doisRetornosTeste() (string, int) {
	return "Legal", 10
}
