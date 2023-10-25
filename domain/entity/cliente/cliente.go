package cliente

type Cliente struct {
    Nome string  `json:"nome"`
    Cpf  string `json:"cpf"`
    Senha string `json:"senha"`
    Id string `json:"id"`
}