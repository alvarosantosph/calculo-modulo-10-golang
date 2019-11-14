/*
 * Cálculo de dígito verificador no módulo 10 padrão febraban (Federação Brasileira de Bancos)
 * Linguagem: GoLang
 * Autor: Alvaro Santos
 * Data: 13/11/2019
 * GitHub: https://github.com/alvarosantosph
 * Email: alvaro_webmaster@hotmail.com
 */

func modulo10(num string) int {
	// Variaveis de controle
	tamanhoString := len(num) + 1
	soma := 0
	resto := 0
	dv := 0
	numeros := make([]string, tamanhoString)
	multiplicador := 2
	runes := []rune(num)
	for i := len(num); i > 0; i-- {
		// Multiplica da direita pra esquerda, alternando os algarismos 2 e 1
		if multiplicador%2 == 0 {
			// Pega cada numero isoladamente
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * 2
			numeros[i] = strconv.Itoa(calculo)
			multiplicador = 1
		} else {
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * 1
			numeros[i] = strconv.Itoa(calculo)
			multiplicador = 2
		}
	}
	for i := (len(numeros) - 1); i > 0; i-- {
		conteudo := numeros[i]
		conteudoInt, _ := strconv.Atoi(conteudo)
		aux := strconv.Itoa(conteudoInt)
		auxiliar := []rune(aux)

		if len(auxiliar) > 1 {
			aux2 := string(auxiliar[0 : len(auxiliar)-1])
			aux3 := string(auxiliar[len(auxiliar)-1 : len(auxiliar)])
			aux4, _ := strconv.Atoi(aux2)
			aux5, _ := strconv.Atoi(aux3)
			aux6 := aux4 + aux5
			variavel := strconv.Itoa(aux6)
			numeros[i] = variavel
		} else {
			numeros[i] = aux
		}
	}
	for i := len(numeros); i > 0; i-- {
		if len(numeros[i-1]) > 0 {
			conteudoSoma, _ := strconv.Atoi(numeros[i-1])
			soma = soma + conteudoSoma
		}
	}
	resto = soma % 10
	dv = 10 - resto
	if dv == 10 {
		dv = 0
	}
	// Retorna o digito verificador
	return dv
}