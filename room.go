package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

/*
var Redis *redis.Client

func CreateRedisClient() {
	opt, err := redis.ParseURL("redis://localhost:6364/0")
	if err != nil {
		panic(err)
	}

	redis := redis.NewClient(opt)
	Redis = redis
	log.Println("Create connection...")
}
*/
/*
func publishMessage(message []byte) {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}

	redis := redis.NewClient(opt)

	errs := redis.Publish(context.TODO(), "mensaje", message).Err()

	if errs != nil {
		log.Println(errs)
	}
}
*/
/*
func createTask(w http.ResponseWriter, r *http.Request) {
	requestAt := time.Now()
	duration := time.Since(requestAt)
	fmt.Fprintf(w, "Task scheduled in %+v", duration)
}
*/

func subscribeMessages() {

	log.Println("Connection Subscriber...")
	opt, err := redis.ParseURL("redis://34.123.108.198:6379/0")
	if err != nil {
		panic(err)
	}

	redis := redis.NewClient(opt)

	pubsub := redis.Subscribe(context.Background(), "mensaje")
	log.Println("subscriber listen on... ")
	ch := pubsub.Channel()
	a := 0
	for msg := range ch {
		log.Println(a)
		a += 1
		log.Println("Mensaje way1: ", string([]byte(msg.Payload)))
		post := []byte(msg.Payload)

		response := make(chan *http.Response)
		response1 := make(chan *http.Response)

		go SendPostAsync("http://34.66.140.170:8080/nuevoRegistro", post, response)
		go SendPostAsync("http://35.223.156.4:7019/nuevoRegistro", post, response1)

		Presponse := <-response
		defer Presponse.Body.Close()
		bytes, _ := ioutil.ReadAll(Presponse.Body)
		log.Println(string(bytes))

		Presponse1 := <-response1
		defer Presponse1.Body.Close()
		bytes1, _ := ioutil.ReadAll(Presponse1.Body)
		log.Println(string(bytes1))

		/*storeResponse1 := SendPostRequest("http://35.223.156.4:7019/nuevoRegistro", post)
		defer storeResponse1.Body.Close()
		n1, _ := ioutil.ReadAll(storeResponse1.Body)
		fmt.Println(string(n1))*/

	}
}

func SendPostAsync(url string, body []byte, rc chan *http.Response) {
	response, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	rc <- response
}

func SendPostRequest(url string, body []byte) *http.Response {
	response, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	return response
}

/*
func sendMsg(msg string) {
	log.Println("Mensaje way1: ", string([]byte(msg)))
	post := []byte(msg)                                                                                         //convertimos a una cadena de bytes
	req, err := http.Post("http://34.66.140.170:8080/nuevoRegistro", "application/json", bytes.NewBuffer(post)) //hacemos la peticion a la bd
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal("Post nuevo documento... ", err)
	}
	defer req.Body.Close() // cerramos el body

	//Leyendo la respuesta del cuerpo
	nuevo, err := ioutil.ReadAll(req.Body) //se convierte en cadena
	if err != nil {
		log.Fatal("Leyendo Respuesta desde el Post Http... ", err)
	}
	sb := string(nuevo) //lo transformamos en una cadena
	log.Printf(sb)
}

func sendMsg1(msg string) {
	log.Println("Mensaje way2: ", string([]byte(msg)))
	post1 := []byte(msg)
	//http://35.223.156.4:7019/nuevoRegistro
	req1, err1 := http.Post("http://35.223.156.4:7019/nuevoRegistro", "application/json", bytes.NewBuffer(post1)) //hacemos la peticion a la bd
	req1.Header.Set("Content-Type", "application/json")
	if err1 != nil {
		log.Fatal("Post nuevo documento... ", err1)
	}
	defer req1.Body.Close() // cerramos el body

	//Leyendo la respuesta del cuerpo
	nuevo1, err1 := ioutil.ReadAll(req1.Body) //se convierte en cadena
	if err1 != nil {
		log.Fatal("Leyendo Respuesta desde el Post Http... ", err1)
	}
	sb1 := string(nuevo1) //lo transformamos en una cadena
	log.Printf(sb1)

}
*/

/*
func createTask(w http.ResponseWriter, r *http.Request) {

	requestAt := time.Now()
	w.Header().Set("Content-Type", "application/json")
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	log.Println("Error Parseando JSON: ", err)
	data, err := json.Marshal(body
	og.Println("Error Reading Body: ", err)
	fmt.Println(string(data))
	publishMessage(data)
	duration := time.Since(requestAt)
	fmt.Fprintf(w, "Task scheduled in %+v", duration)
}*/

func main() {
	subscribeMessages()
}
