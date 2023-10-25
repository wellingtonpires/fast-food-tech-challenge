package main

import(
    "github.com/gin-gonic/gin"
    "github.com/wellingtonpires/fast-food-tech-challenge/interface/entrega_handler"
    "github.com/wellingtonpires/fast-food-tech-challenge/interface/identificacao_handler"
)

func main() {
    router := gin.Default()
    entrega_handler.Routes(router)
    identificacao_handler.Routes(router)
    router.Run()
}