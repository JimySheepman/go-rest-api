package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordsRequestPayload struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type RecordsResponsePayload struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Records []Record `json:"records"`
}

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

func GetDataHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var recordsRequestPayload RecordsRequestPayload

		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Could not complete read from body")
		}

		err = json.Unmarshal(req, &recordsRequestPayload)
		if err != nil {
			log.Fatal("Could not complete unmarsha body")
		}

		filter := bson.D{{"counts", bson.D{{"$lte", recordsRequestPayload.MaxCount}}}}

		coll := db.Collection("records")

		cursor, err := coll.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}
		fmt.Println(cursor)

		var results []bson.D
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		fmt.Println(results)

		for _, result := range results {
			fmt.Println(result)
		}

		json.NewEncoder(w).Encode(recordsRequestPayload)
	}
}

/*
json.NewDecoder(r.Body()).Decode(recordsRequestPayload)

recordsCollection := db.Collection("records")
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

//query, err := recordsCollection.Find(ctx, bson.M{})
query, err := recordsCollection.Aggregate(ctx, bson.D{})
if err != nil {
	log.Fatal(err)
}
var recordsQueryResult []bson.M
if err = query.All(ctx, &recordsQueryResult); err != nil {
	log.Fatal(err)
}

var tmp []int
var tmpStr []string

for i := 0; i < len(recordsQueryResult); i++ {
	jsonRecordCount, _ := json.Marshal(recordsQueryResult[i]["counts"])
	stringRecord := string(jsonRecordCount)
	splitRecord := strings.Split(stringRecord[1:len(stringRecord)-1], ",")
	tmp = append(tmp, helper.Add(splitRecord))

	jsonRecordCreatedAt, _ := json.Marshal(recordsQueryResult[i]["createdAt"])
	stringRecordCreatedAt := string(jsonRecordCreatedAt)

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	submatchall := re.FindAllString(stringRecordCreatedAt, -1)
	tmpStr = append(tmpStr, submatchall[0])
}
minC, maxC := helper.FindMaxAndMinValue(tmp)
startD, endD := helper.FindStartAndEndDate(tmpStr)

recordsRequestPayload := RecordsRequestPayload{
	StartDate: startD,
	EndDate:   endD,
	MinCount:  minC,
	MaxCount:  maxC,
}

json.NewEncoder(w).Encode(recordsRequestPayload) */
