package utils

import (
    "fmt"
    "os"
    "time"

    "bytes"
    "encoding/json"
    "net/http"

    "github.com/hashgraph/hedera-sdk-go/v2"
    "github.com/joho/godotenv"

)

func VerifyMain() {
    data := map[string]int{"input": 42}
    jsonData, _ := json.Marshal(data)

    resp, err := http.Post("http://localhost:5000/process", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    var result map[string]string
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&result)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Result:", result["output"])

    client := SetupHedera()

    topicID := createTopic(client)

    SubscribeToTopic(client, topicID)

    transmission, err := json.Marshal(result)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    TransmitMessage(string(transmission), client, topicID)
    TransmitMessage("Message 2", client, topicID)
    TransmitMessage("Message 3", client, topicID)

    time.Sleep(30 * time.Second) // Prevent the program from exiting to display the message from the mirror to the console
}

func SetupHedera() *hedera.Client {
    err := godotenv.Load(".env")
    if err != nil {
        panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
    }

    myAccountId, err := hedera.AccountIDFromString(os.Getenv("MY_ACCOUNT_ID"))
    if err != nil {
        panic(err)
    }

    myPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("MY_PRIVATE_KEY"))
    if err != nil {
        panic(err)
    }

    client := hedera.ClientForTestnet()
    client.SetOperator(myAccountId, myPrivateKey)

    return client

}

func createTopic(client *hedera.Client) hedera.TopicID {
    transactionResponse, err := hedera.NewTopicCreateTransaction().Execute(client)
    if err != nil {
        panic(fmt.Errorf("%v: error creating topic", err))
    }

    transactionReceipt, err := transactionResponse.GetReceipt(client)
    if err != nil {
        panic(fmt.Errorf("%v: error getting topic create receipt", err))
    }

    topicID := *transactionReceipt.TopicID
    fmt.Printf("topicID: %v\n", topicID)

    return topicID
}

func GetTopicInfo(topicID string) hedera.TopicID {
    ID, err := hedera.TopicIDFromString(topicID)
    if err != nil {
        panic(fmt.Errorf("%v: error retrieving topic info", err))
    }
    return ID
}

func SubscribeToTopic(client *hedera.Client, topicID hedera.TopicID) {
    _, err := hedera.NewTopicMessageQuery().
        SetTopicID(topicID).
        Subscribe(client, func(message hedera.TopicMessage) {
            fmt.Println(message.ConsensusTimestamp.String(), "received topic message ", string(message.Contents), "\r")
        })

    if err != nil {
        panic(fmt.Errorf("%v: error subscribing to topic", err))
    }
}

func TransmitMessage(input string, client *hedera.Client, topicID hedera.TopicID) {
    submitMessage, err := hedera.NewTopicMessageSubmitTransaction().
        SetMessage([]byte(input)).
        SetTopicID(topicID).
        Execute(client)

    if err != nil {
        panic(fmt.Errorf("%v: error submitting to topic", err))
    }

    receipt, err := submitMessage.GetReceipt(client)
    if err != nil {
        panic(fmt.Errorf("%v: error getting transaction receipt", err))
    }

    if receipt.Status.String() != "SUCCESS" {
        panic(fmt.Errorf("error submitting to topic"))
    }

    fmt.Println("Transaction message status " + receipt.Status.String())
}
