package area

import (
	"math"
	"testing"
)

// TESTES DE CAMINHOS FELIZES
func TestCalcularPiramideArea(t *testing.T) {
	ladoQuadradoBase := float32(5)
	baseTrianguloFace := float32(6)
	alturaTrianguloFace := float32(8)

	expected := float32(121)
	result, err := CalcularPiramideArea(ladoQuadradoBase, baseTrianguloFace, alturaTrianguloFace)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalcularCuboArea(t *testing.T) {
	ladoQuadradoBase := float32(4)

	expected := float32(96)
	result, err := CalcularCuboArea(ladoQuadradoBase)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalcularParalelepipedoArea(t *testing.T) {
	comprimentoBase := float32(5)
	larguraBase := float32(3)
	alturaBase := float32(4)

	expected := float32(94)
	result, err := CalcularParalelepipedoArea(comprimentoBase, larguraBase, alturaBase)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalcularEsferaArea(t *testing.T) {
	raio := float32(2)

	expected := float32(math.Pi * 4 * 4)
	result, err := CalcularEsferaArea(raio)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalcularAreaTriangulo(t *testing.T) {
	base := float32(5)
	altura := float32(8)

	expected := float32(20)
	result, err := CalcularAreaTriangulo(base, altura)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalcularAreaRetangulo(t *testing.T) {
	base := float32(6)
	altura := float32(10)

	expected := float32(60)
	result, err := CalcularAreaRetangulo(base, altura)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalcularAreaQuadrado(t *testing.T) {
	lado := float32(4)

	expected := float32(16)
	result, err := CalcularAreaQuadrado(lado)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalcularAreaCirculo(t *testing.T) {
	raio := float32(3)

	expected := float32(math.Pi * 3 * 3)
	result, err := CalcularAreaCirculo(raio)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

// TESTES DE CAMINHOS INFELIZES

func TestErroCalcularPiramideArea(t *testing.T) {
	area, err := CalcularPiramideArea(2, 4, -3)
	if err == nil {
		t.Errorf("Esperava-se um erro, mas nenhum erro foi retornado.")
	}

	if area != 0 {
		t.Errorf("Esperava-se uma área de 0, mas o valor retornado foi %f.", area)
	}
}

func TestErroCalcularCuboArea(t *testing.T) {
	area, err := CalcularCuboArea(0)
	if err == nil {
		t.Errorf("Esperava-se um erro, mas nenhum erro foi retornado.")
	}

	if area != 0 {
		t.Errorf("Esperava-se uma área de 0, mas o valor retornado foi %f.", area)
	}
}

func TestErroCalcularParalelepipedoArea(t *testing.T) {
	area, err := CalcularParalelepipedoArea(2, -4, 3)
	if err == nil {
		t.Errorf("Esperava-se um erro, mas nenhum erro foi retornado.")
	}

	if area != 0 {
		t.Errorf("Esperava-se uma área de 0, mas o valor retornado foi %f.", area)
	}
}

func TestErroCalcularEsferaArea(t *testing.T) {
	area, err := CalcularEsferaArea(-2)
	if err == nil {
		t.Errorf("Esperava-se um erro, mas nenhum erro foi retornado.")
	}

	if area != 0 {
		t.Errorf("Esperava-se uma área de 0, mas o valor retornado foi %f.", area)
	}
}

func TestErroCalcularAreaTriangulo(t *testing.T) {
	area, err := CalcularAreaTriangulo(-2, 4)
	if err == nil {
		t.Errorf("Esperava-se um erro, mas nenhum erro foi retornado.")
	}

	if area != 0 {
		t.Errorf("Esperava-se uma área de 0, mas o valor retornado foi %f.", area)
	}
}

func TestErroCalcularAreaRetangulo(t *testing.T) {
	area, err := CalcularAreaRetangulo(0, 5)
	if err == nil {
		t.Errorf("Esperava-se um erro, mas nenhum erro foi retornado.")
	}

	if area != 0 {
		t.Errorf("Esperava-se uma área de 0, mas o valor retornado foi %f.", area)
	}
}

func TestErroCalcularAreaQuadrado(t *testing.T) {
	area, err := CalcularAreaQuadrado(-2)
	if err == nil {
		t.Errorf("Esperava-se um erro, mas nenhum erro foi retornado.")
	}

	if area != 0 {
		t.Errorf("Esperava-se uma área de 0, mas o valor retornado foi %f.", area)
	}
}

func TestErroCalcularAreaCirculo(t *testing.T) {
	area, err := CalcularAreaCirculo(0)
	if err == nil {
		t.Errorf("Esperava-se um erro, mas nenhum erro foi retornado.")
	}

	if area != 0 {
		t.Errorf("Esperava-se uma área de 0, mas o valor retornado foi %f.", area)
	}
}

