package main

import (
	"time"
	"fmt"
	"crypto/md5"
	"encoding/json"

	vapi "github.com/hashicorp/vault/api"
)


// Poll on a individual app for new secrets
// Expect the path VaultPrefix/App/{key:value}
func PollLoop(client *vapi.Logical, stop chan struct{}, appname string) {
	VaultPath := VaultPrefix + appname
	PollMap := map[string][16]byte{}
	ticker := time.NewTicker(PollInterval)

	for {
		select {
		case <-stop:
			ticker.Stop()
			return
		case <-ticker.C:

			// First thing we'll do is get a list of all secrets in the path.
			secret, err := client.List(VaultPath)
			if err != nil {
				panic(err)
				// Do some handling.
			}

			if secret == nil {
				continue
			}

			for _, s := range secret.Data["keys"].([]interface{}) {
				secretName := s.(string)
				completePath := VaultPath + "/" + secretName
				val, err := client.Read(completePath)
				if err != nil {
					panic(err)
				}

				// Yolo. Need to turn the data into something that can easily be evaluated for equality
				jsons, err := json.Marshal(val.Data)
				if err != nil {
					panic(err)
				}

				hash := md5.Sum([]byte(jsons))

				if _, ok := PollMap[secretName]; ok {
					if PollMap[secretName] == hash {
						// Our hash value has not changed,
						// So there is no need to update the dictionary
						fmt.Println("Not updating secret: ", secretName)
						continue
					}
				}
				// Key doesnt exist, or it does exist and the value needs updating.
				PollMap[secretName] = hash
				fmt.Println("UPDATING: ", secretName)
				content := map[string][]byte{}
				// Turn val.Data into a map[string][]byte
				for k, v := range val.Data {
					val, ok := v.(string)
					if !ok {
						fmt.Println("Couldnt convert to string.", v, val)
					}
					content[k] = []byte(val)
				}

				UpdateSecret(appname, secretName, content)
			}
		}
	}
}
