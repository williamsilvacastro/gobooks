package models

import (
    "crud-livros/config"
    "log"
)

type Livro struct {
    ID        int    `json:"id"`
    Titulo    string `json:"titulo"`
    Autor     string `json:"autor"`
    Descricao string `json:"descricao"`
}

// Obter todos os livros
func BuscarLivros() ([]Livro, error) {
    rows, err := config.DB.Query("SELECT id, titulo, autor, descricao FROM livros")
    if err != nil {
        log.Println("Erro ao buscar livros:", err)
        return nil, err
    }
    defer rows.Close()

    var livros []Livro
    for rows.Next() {
        var livro Livro
        if err := rows.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Descricao); err != nil {
            log.Println("Erro ao escanear livro:", err)
            return nil, err
        }
        livros = append(livros, livro)
    }
    return livros, nil
}

// Criar um novo livro
func CriarLivro(livro *Livro) error {
    query := "INSERT INTO livros (titulo, autor, descricao) VALUES (?, ?, ?)"
    _, err := config.DB.Exec(query, livro.Titulo, livro.Autor, livro.Descricao)
    if err != nil {
        log.Println("Erro ao criar livro:", err)
        return err
    }
    return nil
}

// Atualizar um livro existente
func AtualizarLivro(livro *Livro) error {
    query := "UPDATE livros SET titulo = ?, autor = ?, descricao = ? WHERE id = ?"
    _, err := config.DB.Exec(query, livro.Titulo, livro.Autor, livro.Descricao, livro.ID)
    if err != nil {
        log.Println("Erro ao atualizar livro:", err)
        return err
    }
    return nil
}

// Deletar um livro
func DeletarLivro(id int) error {
    query := "DELETE FROM livros WHERE id = ?"
    _, err := config.DB.Exec(query, id)
    if err != nil {
        log.Println("Erro ao deletar livro:", err)
        return err
    }
    return nil
}
