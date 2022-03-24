package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/JimySheepman/go-rest-api/internal/helper"
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

// POST Endpoint (fetch data from mongodb)
func GetFetchDataHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var recordsRequestPayload RecordsRequestPayload

		// Read request body
		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode(RecordsResponsePayload{
				Code:    1,
				Message: "Error: could not complete read from request body",
				Records: []Record{},
			})
			log.Println("Could not complete read from request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Unmarshal request body
		err = json.Unmarshal(req, &recordsRequestPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RecordsResponsePayload{
				Code:    2,
				Message: "Error: could not complete unmarshal body",
				Records: []Record{},
			})
			log.Println("Could not complete unmarshal body")
			return
		}

		// time format validation
		startDateValidation := helper.TimeFormatValidator(recordsRequestPayload.StartDate)
		endDateValidation := helper.TimeFormatValidator(recordsRequestPayload.EndDate)
		if !startDateValidation || !endDateValidation {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(RecordsResponsePayload{
				Code:    3,
				Message: "Error: wrong time format ",
				Records: []Record{},
			})
			log.Println("Wrong time format ")
			return
		}

		// convert to time yyyy-mm-dd to time.RFC3339
		startDate := helper.TimeConverter(recordsRequestPayload.StartDate)
		endDate := helper.TimeConverter(recordsRequestPayload.EndDate)
		fmt.Println(startDate, endDate)

		/*create query part for mongo pipeline
		db.records.aggregate([
			{
			  $match: {
				createdAt: {
				  $gte: ISODate("2016-07-25T00:00:00Z"),
				  $lt: ISODate("2017-07-25T00:00:00Z"),

				}
			  }
			},
			{
			  $unwind: "$counts"
			},
			{
			  "$group": {
				"_id": {
				  "key": "$key",
				  "createdAt": "$createdAt"
				},
				"totalCount": {
				  "$sum": "$counts"
				}
			  }
			},
			{
			  $match: {
				totalCount: {
				  $gte: 1800,
				  $lt: 2800,

				}
			  }
			},
			{
			  $project: {
				"_id": 0,
				"key": "$_id.key",
				"createdAt": "$_id.createdAt",
				"totalCount": 1
			  }
			}
		  ])
		*/
		matchDate := bson.D{{"$match", bson.D{{"createdAt", bson.D{{"$gte", startDate}, {"$lt", endDate}}}}}}
		unwindCounts := bson.D{{"$unwind", "$counts"}}
		groupCount := bson.D{{"$group", bson.D{{"_id", bson.D{{"key", "$key"}, {"createdAt", "$createdAt"}}}, {"totalCount", bson.D{{"$sum", "$counts"}}}}}}
		matchCount := bson.D{{"$match", bson.D{{"totalCount", bson.D{{"$gte", recordsRequestPayload.MinCount}, {"$lt", recordsRequestPayload.MaxCount}}}}}}
		projectQuery := bson.D{{"$project", bson.D{{"_id", 0}, {"key", "$_id.key"}, {"createdAt", "$_id.createdAt"}, {"totalCount", 1}}}}

		coll := db.Collection("records")

		// add pipeline
		cursor, err := coll.Find(context.TODO(), mongo.Pipeline{})
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