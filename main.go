package main

import (
	"fmt"
	"bytes"
	"gopkg.in/redis.v4"
	"encoding/json"
	"time"
)

type FileForDownload struct {
	HasProgramming            bool  `json: "HasProgramming"`
	IndDiffer                 bool  `json: "IndDiffer"`
	DestVersionId             int   `json: "DestVersionId"`
	SwitchingPercentage       int   `json: "SwitchingPercentage"`
	ProgrammingId             int   `json: "ProgrammingId"`
	IndStartWifiOnly          bool  `json: "IndStartWifiOnly"`
	MaxTermDownloadingVersion int   `json: "MaxTermDownloadingVersion"`
}

func main() {
    fmt.Println("Go Redis Connection Test")

	// connect REDIS
	// -----------------------------------
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

    pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// -------------------------------------

	// tratamento para JSON - ESCREVER
	// -------------------------------------

	// struct que vai ser escrita pelo GO
	var fileForDownloadWrite FileForDownload

	fileForDownloadWrite.HasProgramming = true
	fileForDownloadWrite.IndDiffer = true
	fileForDownloadWrite.DestVersionId = 5
	fileForDownloadWrite.SwitchingPercentage = 1000
	fileForDownloadWrite.ProgrammingId = 6
	fileForDownloadWrite.IndStartWifiOnly = true
	fileForDownloadWrite.MaxTermDownloadingVersion = 2000


	fileForDownloadWriteJSON,  _ := json.Marshal(fileForDownloadWrite)
	fmt.Println("Escrevendo JSON: ", bytes.NewBuffer(fileForDownloadWriteJSON))

	client.Set("Teste", fileForDownloadWriteJSON, 10*time.Minute)
	// -------------------------------------

	// tratamento para JSON - LER (vamos ler o valor TESTE gravado acima no REDIS)	
	// -------------------------------------
	var fileForDownloadRead FileForDownload
	fReadRedisJSON, _ := client.Get("Teste").Result()

	if erro := json.Unmarshal([]byte(fReadRedisJSON), &fileForDownloadRead); erro != nil {
		fmt.Println(erro)
	}

	fmt.Println("Lendo struct ", fileForDownloadRead)
	// -------------------------------------


	// tratando com MAP[string][string]
	//result := client.HGetAll("Student:940198001");
	//fmt.Println(result);

	//fmt.Println(fileForDownload)

}