# Go lang
Introdução à linguagem Go

## O que são tipos?
* Usa-se um Tipo (String) para escrever nosso Hello World!
* Exemplo: Cachorro Nelson
* Linguagem _estaticamente tipada_
* O tipo tem uma característica definida.

## Tipos: Número
### Inteiros;
* Números sem compoente decimal(-2, -1, 0, 1)
* Tem tamanho definido
* Tipos inteiros em Go: uint8, unint16, uint32, uint64, int8, int16, int32, int64
* Uint (unsigned integer - Inteiro sem sinal - somente +)
* Int (signed integer - inteiro com sinal - + e - )
### Curisidade
* Uint8 => Bite
* Int32 => rune

## Ponto Flutuante
* Números com componente decimal(1,523; 153,4)
* TEm tamanho definido
* Tipos ponto flutuante em go: float32, float64
* Em geral, usamos float64

## Tipo String
* representa um texto que é uma sequência de caracteres
* possui tamanho definido
* aparece entre "" ou ``
 ### Operações comuns com Strings
 * Len("Hello World") - Descobre o tamanho da String
 * "Hello World"[2] - Acrescenta um caractres específico
 * "Hello" + "World" - Concatena duas Strings

## Booleanos
* Homenagem ao matemático George Boole
* Representa verdadeiro ou falso (ligado ou desligado)
* Operadores lógicos: &&(e), ||(ou) !(negação)

## Inferência de Tipos
* Go consegue inferir(decifrar) o tipo dessas variáveis. Ele consegue entender que, se a variável começa e termina com as aspas, ela é uma string. Da mesma forma, se temos um número inteiro, sem casa decimal, A linguagem GO entenderá que a variável é do tipo inteiro.

## Palavras Reservadas
* São componentes da própria linguagem e não podem ser redefinidas, ou seja, denominar elementos criados pelo programador.
* Exemplo: _CONST_, palavra reservada para indicar uma constante, não pode ser usada para quaisquer outro identificador no programa.
### Go possui 25 palavras reservadas
| break   | case     | chan       | func    | if       | interface | return | type      | var      | 
|  :---:  |  :----:  |  :---:     | :---:   | :---:    |  :---:    | :---:  | :---:     | :---:    |
| const   | default  |fallthrough | go      |  if      | package   | struct |  map      |          |
| continue| else     | for        | goto    |  import  | range     | switch | select    |          |
