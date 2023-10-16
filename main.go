package main



import(
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"github.com/google/uuid"
)


type Cadastro struct{
	ID string
	Usuario string
	Email string
	Senha string

	

}

func main(){
	fmt.Println("Cadastro de usuarios")
	fmt.Println()

	c := Cadastro{

		ID: uuid.New().String(),
		//Usuario: "joao",
		//Email: "joao2@gmail.com",
	//	Senha: "dnuihduhsudhu",
		
		
	}

	//Entrada de dados do usuario

	fmt.Println("Digite seu usuario: ")
	fmt.Scan(&c.Usuario)
	fmt.Println("Digite seu Email: ")
	fmt.Scan(&c.Email)
	fmt.Println("Digite Sua senha: ")
	fmt.Scan(&c.Senha)


	//Abrindo a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cadastro")
	if err != nil{
		log.Println("Erro ao se conectar com DB", err)
	}

	//log.Println("Conexão feito com o banco de dados")

	defer db.Close()



	//Inserindo dados no banco de dados
	colocar, err := db.Prepare("insert into cadastro(id, usuario, email, senha) values(?,?,?,?)")
	if  err != nil{
		log.Println("Erro1 ao inserir dados no banco de dado", err)
	}

	_, err = colocar.Exec(c.ID, c.Usuario, c.Email, c.Senha)
	if err != nil{
		log.Println("Erro2 ao inserir dados no banco de dado", err)
	}

	fmt.Println("Muito bem" ,c.Usuario, "Seus dados foram cadastrado com sucesso!")

	defer colocar.Close()


}