package area

import (
	"math"
)

//  AreaError significa um erro associado ao calculo de àrea.
type AreaError struct {
    Mensagem string
}

func (e *AreaError) Error() string {
    return e.Mensagem
}

// FIGURAS ESPACIAIS

// Calcular área da Pirâmide
func CalcularPiramideArea(ladoQuadradoBase, baseTrianguloFace, alturaTrianguloFace float32) (float32, error) {
	if ladoQuadradoBase <= 0 || baseTrianguloFace <= 0 || alturaTrianguloFace <= 0 {
        return 0, &AreaError{"Todos valores devem ser maior que 0."}
    }
	// Área das faces laterais somadas
	areaFace := (baseTrianguloFace * alturaTrianguloFace) / 2
	areaFaces := areaFace * 4

	// Área da base
	areaBase := ladoQuadradoBase * ladoQuadradoBase

	areaPiramide := areaFaces + areaBase

	return areaPiramide, nil
}

// Calcular área do cubo
func CalcularCuboArea(ladoQuadradoBase float32) (float32, error) {
	if ladoQuadradoBase <= 0 {
        return 0, &AreaError{"O valor do lado deve ser maior que zero."}
    }
	// Área da base
	areaBase := ladoQuadradoBase * ladoQuadradoBase

	areaCubo := areaBase * 6

	return areaCubo, nil
}

// Calcular área do paralelepípedo
func CalcularParalelepipedoArea(comprimentoBase, larguraBase, alturaBase float32) (float32, error) {
	if comprimentoBase <= 0 || larguraBase <= 0 || alturaBase <= 0 {
        return 0, &AreaError{"Nenhum valor deve ser menor que zero."}
    }
	area := (2 * comprimentoBase * larguraBase) + (2 * comprimentoBase * alturaBase) + (2 * alturaBase * larguraBase)
	return area, nil
}

// Calcular área da esfera
func CalcularEsferaArea(raio float32) (float32, error) {
	if raio <= 0 {
        return 0, &AreaError{"O valor do raio deve ser maior que zero."}
    }
	area := 4 * math.Pi * raio * raio
	return area, nil
}

// FIGURAS PLANAS

// Função para calcular a área do triângulo
func CalcularAreaTriangulo(base, altura float32) (float32, error) {
	if base <= 0 || altura <= 0 {
        return 0, &AreaError{"Os valores de base e altura devem ser maiores que zero."}
    }
	area := (base * altura) / 2
	return area, nil
}

// Função para calcular a área do Retângulo
func CalcularAreaRetangulo(base, altura float32) (float32, error) {
	if base <= 0 || altura <= 0 {
        return 0, &AreaError{"Os valores de base e altura devem ser maiores que zero."}
	}
	area := base * altura
	return area, nil
}

// Função para calcular a área do Quadrado
func CalcularAreaQuadrado(lado float32) (float32, error) {
	if lado <= 0 {
        return 0, &AreaError{"O valor do lado deve ser maior que zero."}
    }
	area := lado * lado
	return area, nil
}

// Função para calcular a área do Círculo
func CalcularAreaCirculo(raio float32) (float32, error) {
	if raio <= 0 {
        return 0, &AreaError{"O valor do raio deve ser maior que zero."}
    }
	area := math.Pi * raio * raio
	return area, nil
}
