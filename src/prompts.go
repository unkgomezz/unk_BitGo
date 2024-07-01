package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// promptRangeNumber prompts the user to select a range number
func PromptRangeNumber(totalRanges int) int {
	reader := bufio.NewReader(os.Stdin)
	charReadline := '\n'

	if runtime.GOOS == "windows" {
		charReadline = '\r'
	}

	for {
		color.Cyan("\n")
		color.Cyan("╔════════════════════════════════════════════╗")
		color.Cyan("\n")
		color.Cyan(" + Selecione uma carteira: [1 ~ %d]", totalRanges)
		color.Cyan("\n")
		color.Cyan("╚════════════════════════════════════════════╝")
		color.Cyan("\n")
		input, _ := reader.ReadString(byte(charReadline))
		input = strings.TrimSpace(input)
		rangeNumber, err := strconv.Atoi(input)
		if err == nil && rangeNumber >= 1 && rangeNumber <= totalRanges {
			return rangeNumber
		}
		color.Red("Número inválido.")
	}
}

func PromptCPUNumber() int {
	reader := bufio.NewReader(os.Stdin)
	charReadline := '\n'

	if runtime.GOOS == "windows" {
		charReadline = '\r'
	}

	for {
		color.Cyan("\n")
		color.Cyan("╔════════════════════════════════════════════╗")
		color.Cyan("\n")
		color.Cyan(" + Selecione a quantidade de CPU: ")
		color.Cyan("\n")
		color.Cyan("╚════════════════════════════════════════════╝")
		color.Cyan("\n")
		input, _ := reader.ReadString(byte(charReadline))
		input = strings.TrimSpace(input)
		cpusNumber, err := strconv.Atoi(input)
		if err == nil && cpusNumber >= 1 && cpusNumber <= 50 {
			return cpusNumber
		}
		color.Red("Número inválido.")
	}
}

// PromptModos prompts the user to select a modo's
func PromptModos(totalModos int) int {
	reader := bufio.NewReader(os.Stdin)
	charReadline := '\n'

	if runtime.GOOS == "windows" {
		charReadline = '\r'
	}

	for {
		color.Cyan("\n")
		color.Cyan("╔════════════════════════════════════════════╗")
		color.Cyan("\n")
		color.Cyan(" + Selecione um modo: [1 ~ %d] \n\n [ 1 ] --> Search Start (Busca do início) \n [ 2 ] --> Search Process (Busca sequencial)", totalModos)
		color.Cyan("\n")
		color.Cyan("╚════════════════════════════════════════════╝")
		color.Cyan("\n")
		input, _ := reader.ReadString(byte(charReadline))
		input = strings.TrimSpace(input)
		modoSelecinado, err := strconv.Atoi(input)
		if err == nil && modoSelecinado >= 1 && modoSelecinado <= totalModos {
			return modoSelecinado
			//fmt.Println(modoSelecinado)
		}
		color.Red("Número inválido.")
	}
}

// PromptAuto solicita ao usuário a seleção de um número dentro de um intervalo específico.
func PromptAuto(pergunta string, totalnumbers int) int {
	reader := bufio.NewReader(os.Stdin)
	charReadline := '\n'

	if runtime.GOOS == "windows" {
		charReadline = '\r'
	}

	for {
		color.Cyan(pergunta)
		input, _ := reader.ReadString(byte(charReadline))
		input = strings.TrimSpace(input)
		resposta, err := strconv.Atoi(input)
		if err == nil && resposta >= 1 && resposta <= totalnumbers {
			return resposta
		}
		color.Red("Resposta inválido.")
	}
}

// HandleModoSelecionado - selecionar modos de incializacao
func HandleModoSelecionado(modoSelecionado int, ranges *Ranges, rangeNumber int, privKeyInt *big.Int, carteirasalva string) *big.Int {
	if modoSelecionado == 1 {
		// Initialize privKeyInt with the minimum value of the selected range
		privKeyHex := ranges.Ranges[rangeNumber-1].Min
		privKeyInt.SetString(privKeyHex[2:], 16)
	} else if modoSelecionado == 2 {
		verificaKey, err := LoadUltimaKeyWallet("lastkey.txt", carteirasalva)
		if err != nil || verificaKey == "" {
			// FAZER PERGUNTA SE DESEJA INFORMAR O NUMERO DE INCIO DO MODO SEQUENCIAL OU COMEÇAR DO INICIO
			msSequencialouInicio := PromptAuto("\n ╔════════════════════════════════════════════╗ \n\n + Selecione uma opção: [1 ~ 2]\n\n [ 1 ] --> Começar do início \n [ 2 ] --> Escolher entre o range da carteira \n\n ╚════════════════════════════════════════════╝ \n\n", 2)
			if msSequencialouInicio == 2 {
				// Definindo as variáveis privKeyMinInt e privKeyMaxInt como big.Int
				privKeyMinInt := new(big.Int)
				privKeyMaxInt := new(big.Int)
				privKeyMin := ranges.Ranges[rangeNumber-1].Min
				privKeyMax := ranges.Ranges[rangeNumber-1].Max
				privKeyMinInt.SetString(privKeyMin[2:], 16)
				privKeyMaxInt.SetString(privKeyMax[2:], 16)

				// Calculando a diferença entre privKeyMaxInt e privKeyMinInt
				rangeKey := new(big.Int).Sub(privKeyMaxInt, privKeyMinInt)

				// Solicitando a porcentagem do range da carteira como entrada
				var rangeCarteiraSequencialStr string
				color.Cyan("\n")
				color.Cyan("╔════════════════════════════════════════════╗")
				color.Cyan("\n")
				color.Cyan(" + Informe a porcentagem do range: [1 ~ 100]")
				color.Cyan("\n")
				color.Cyan("╚════════════════════════════════════════════╝")
				color.Cyan("\n")
				fmt.Scanln(&rangeCarteiraSequencialStr)

				// Substituindo vírgulas por pontos se necessário
				rangeCarteiraSequencialStr = strings.Replace(rangeCarteiraSequencialStr, ",", ".", -1)

				// Convertendo a porcentagem para um número decimal
				rangeCarteiraSequencial, err := strconv.ParseFloat(rangeCarteiraSequencialStr, 64)
				if err != nil {
					color.Red("Erro ao ler porcentagem:", err)
					return nil
				}

				// Verificando se a porcentagem está no intervalo válido
				if rangeCarteiraSequencial < 1 || rangeCarteiraSequencial > 100 {
					color.Red("Porcentagem fora do intervalo válido (1 a 100).")
					return nil
				}

				// Calculando o valor de rangeKey multiplicado pela porcentagem
				rangeMultiplier := new(big.Float).Mul(new(big.Float).SetInt(rangeKey), big.NewFloat(rangeCarteiraSequencial/100.0))

				// Convertendo o resultado para inteiro (arredondamento para baixo)
				min := new(big.Int)
				rangeMultiplier.Int(min)

				// Adicionando rangeMultiplier ao valor mínimo (privKeyMinInt)
				min.Add(privKeyMinInt, min)

				// Verificando o valor final como uma string hexadecimal
				verificaKey := min.Text(16)
				privKeyInt.SetString(verificaKey, 16)
				color.Green("\n[INFO] Informações selecionadas com sucesso.")
			} else {
				verificaKey = ranges.Ranges[rangeNumber-1].Min
				privKeyInt.SetString(verificaKey[2:], 16)
				color.Green("\n[INFO] Nenhuma chave privada salva encontrada, iniciando do começo. %s: %s\n", carteirasalva, verificaKey)
			}
		} else {
			color.Green("\n[INFO] Última chave da carteira %s: %s\n", carteirasalva, verificaKey)
			privKeyInt.SetString(verificaKey, 16)
		}
	}
	return privKeyInt
}
