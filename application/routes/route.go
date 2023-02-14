package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Estrutura de rota.
type Route struct {
	ID        string      `json:"routeId"`
	ClientID  string      `json:"clientId"`
	Positions []Positions `json:"Position"`
}

// Estrutura de posições que serão baseadas em latitude e longitude. "Objeto" apenas para representar as posições em longitude e latitude.
type Positions struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

// Iremos criar um método para carregar as posições para dentro da rota. Vamos carregar todas as posições para dentro do array "Positions []Positions"
func (r *Route) LoadPositions() error {
	// validação se a rota realmente existe.
	if r.ID == "" {
		return errors.New("Route ID not informed.")
	}

	// Vamos abrir o arquivo de texto.
	f, err := os.Open("destinations/" + r.ID + ".txt")

	// Vamos verificar se o "err" possui um valor diferente de "vazio". Caso possua, significa que um erro ocorreu.
	if err != nil {
		return err
	}

	// Quando abrimos o arquivo, precisamos necessariamente fechá-lo após o uso para não ocorrer o erro de "memoru leaks". Mas, para não esquecermos de fechar e deixar a aplicação consumindo memória de forma contínua, usamos o "defer". O "defer" espera que o método inteiro seja completamente executado para fechar o processo em aberto e liberar memória.
	defer f.Close()

	// Estrutura responsável por ler o arquivo de texto. Aparentemente é uma leitura de baixo nível (bytes).
	scanner := bufio.NewScanner(f)

	// O método "Scan()" retorna um boolean. Esse loop irá chegar se existem linhas para serem lidas e, caso existam, um true será
	for scanner.Scan() {
		// O método split pega uma sequência de caracteres, separa-os por um determinado símbolo e junta-os em um array. Nesse caso a latitude e longitude vêm em apenas uma linha de informação. Após isso, será um array com dois elementos.
		data := strings.Split(scanner.Text(), ",")

		// Vamos converter os dados de latitude e longitude, que antes estavam em string, para float.
		lat, err := strconv.ParseFloat(data[0], 64) // O segundo parâmetro "64" são os bits que o float recebe.
		if err != nil {
			return err
		}

		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}

		// Caso não ocorra nenhum erro, vamos adicionar essas novas informações para dentro de "Positions". Quando usamos "Positions{}" dessa forma, estamos passando valores diretamente para dentro do struct de Positions e, esse struct, está sendo passado para o "r.Positions" que vem no primeiro parâmetro.
		r.Positions = append(r.Positions, Positions{
			Lat:  lat,
			Long: long,
		})

	}
	return nil
}

// Vamos sempre retornar os dados em formato Json e esse struct irá servir para isso. Ele irá receber cada linha de rota que os arquivos de texto possuem (Position []float64) e terá um atributo para definir se já chegou ao fim da rota. Esse atributo "finished" será "true" quando a última posição for enviada, que significa que a rota foi finalizada.
type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (r *Route) ExportJsonPositions() ([]string, error) {

	var route PartialRoutePosition
	var result []string

	total := len(r.Positions)

	// Iremos percorrer todas as posições "r.Positions" do struct de rota.
	for k, v := range r.Positions {
		// Vamos popular todas as informações do struct "PartialRoutePositions" com as informações do struct "Routes".
		route.ID = r.ID
		route.ClientID = r.ClientID
		// Esse atributo recebe um array de float64 com as informações sendo passadas diretamente para o array através da variável "v" (value).
		route.Position = []float64{
			v.Lat,
			v.Long,
		}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		// Conversão de dados para Json. Esse método "json.Marshal()" converte o argumento que ele recebeu em um array de bytes.
		jsonRoute, err := json.Marshal(route)

		if err != nil {
			return nil, err
		}
		// Caso nenhum erro tenha ocorrido, iremos jogar o "jsonRoute" transformado para string JSON dentro da variável "result".
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
