package identificacao


import(
	"fmt"
	"net/http"
    "github.com/gin-gonic/gin"
    "github.com/wellingtonpires/fast-food-tech-challenge/infrastructure/persistence"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/cliente"
)

func Cadastro(c *gin.Context) {
	var cli cliente.Cliente
	err := c.ShouldBindJSON(&cli)
	if err != nil {
		fmt.Println(err.Error())
	}
	persistence.CadastraCliente(cli)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Cliente criado!"})
}