package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"pwdkeeper/internal/app/crypter"
	"pwdkeeper/internal/app/initconfig"
	"pwdkeeper/internal/app/msgsender"

	pb "pwdkeeper/internal/app/proto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Greetings(text string){
	fmt.Print(text)
}

func main() {
	initconfig.SetinitVars()

	log.Logger.Info().Msg("Starting CLIENT...")
	log.Logger.Info().Msg("Connecting to Server localhost:3200...")
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err)
	}
	//defer conn.Close()
	// получаем переменную интерфейсного типа ActionsClient,
	// через которую будем отправлять сообщения
	c := pb.NewActionsClient(conn)
	log.Logger.Info().Msg("Connected.")
	Greetings("Starting UI...")

	StartUI(c)
	// функция, в которой будем отправлять сообщения
	//SendUserAuthmsg(c, "", "")	
}

func StartUI(c pb.ActionsClient) {
	menulevel := 1
	userisNew := true
	login := ""
	loginstatus := ""
	password := ""
	key1enc :=""
	recordisNew := false
	recordIDname :=""			
	somedata :=""			
	datatype :=""			

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Hello! Enter Login: ")



	for {				
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		consoleInput := scanner.Text()
		log.Debug().Msgf("consoleInput =%v", consoleInput)
		switch menulevel {
//!Read Login			
			case 1: 				
				login = consoleInput
				loginstatus, key1enc = msgsender.SendUserGetmsg(c, login)
				if (loginstatus == "200") {
					userisNew = false
					fmt.Print("Enter Password: ")
				} else {
					fmt.Print("Create Password for NEW user: ")
				}
				menulevel = 2

//!Read Password and Show All user Records
			case 2: 
				password = consoleInput
				if userisNew {			
					log.Info().Msgf(string(crypter.Key2build(password)))	
					EncryptedKey1 := crypter.EncryptKey1(crypter.Key1build(), crypter.Key2build(password))
					log.Info().Msgf("EncryptedKey1 %v", hex.EncodeToString(EncryptedKey1))
					log.Info().Msgf(string(crypter.DecryptKey1((EncryptedKey1), crypter.Key2build(password))))
					//if msgsender.SendUserStoremsg(c, login, "password", string(EncryptedKey1)) == "200" {
					if msgsender.SendUserStoremsg(c, login, "password", hex.EncodeToString(EncryptedKey1)) == "200" {
						log.Info().Msgf("User %v created and logged in!", login)
						fmt.Print("Enter NAME of new record to create: ")
						menulevel = 3
						} else {
							log.Error().Msg("Error!")
							fmt.Print("Enter Login: ")
							menulevel = 1
						}
					} else {
						log.Info().Msgf(string(crypter.DecryptKey1([]byte(key1enc), crypter.Key2build(password))))
						log.Info().Msgf(string(crypter.DecryptKey1([]byte(key1enc), crypter.Key2build(password))))
						log.Info().Msgf(string(crypter.DecryptKey1([]byte(key1enc), crypter.Key2build(password))))
							

						if msgsender.SendUserAuthmsg(c, login, password) == "200" {
							log.Info().Msgf("User %v logged in!", login)
							status, userRecordsJSON := msgsender.SendUserGetRecordsmsg(c, login)
							log.Debug().Msg(status)
							log.Info().Msg(userRecordsJSON)
							fmt.Print("Enter ID of existing record or NAME of new record to create: ")
							menulevel = 3
							} else {
								log.Error().Msg("Error! Wrong Password!")
								fmt.Print("Enter Login: ")
								menulevel = 1
							}
					}

//!Create new or Update existing Record 1)Read existing ID or new NAME
			case 3:
				recordIDname = consoleInput
				
				//TODO check if entered int ID xor text NAME
				recordisNew = true
				if recordisNew{
					fmt.Print("Enter somedata to store: ")
					menulevel = 31
					} else {
						fmt.Print("Enter somedata to update: ")
						menulevel = 41
					}				

//!Create new Record 2)Read new someDATA
			case 31:
				somedata = consoleInput			
				fmt.Print("Enter data type to store: ")
				menulevel = 32
				
//!Create new Record 3)Read new datatype and Store NEW record returning created id
			case 32:
				datatype = consoleInput
				status, recordID := msgsender.SendUserStoreRecordmsg(c, recordIDname, somedata, datatype, login)
				if status == "200" {
					log.Info().Msg("Created new record with ID=")
					log.Info().Msg(recordID)
				} else {
					log.Error().Msg("Error creating NEW data record!")
				}
				status, userRecordsJSON := msgsender.SendUserGetRecordsmsg(c, login)
				log.Debug().Msg(status)
				log.Info().Msg(userRecordsJSON)
				fmt.Print("Enter ID of existing record or NAME of new record to create: ")
				menulevel = 3				
			}					
	}
}

