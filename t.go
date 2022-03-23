// retrieve single and multiple documents with a specified filter using FindOne() and Find()
// create a search filer
filter := bson.D{
	{"$and",
		bson.A{
			bson.D{
				{"age", bson.D{{"$gt", 25}}},
			},
		},
	},
}

// retrieve all the documents that match the filter
cursor, err := usersCollection.Find(context.TODO(), filter)
// check for errors in the finding
if err != nil {
	panic(err)
}

// convert the cursor result to bson
var results []bson.M
// check for errors in the conversion
if err = cursor.All(context.TODO(), &results); err != nil {
	panic(err)
}

// display the documents retrieved
fmt.Println("displaying all results from the search query")
for _, result := range results {
	fmt.Println(result)
}

// retrieving the first document that match the filter
var result bson.M
// check for errors in the finding
if err = usersCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
	panic(err)
}

// display the document retrieved
fmt.Println("displaying the first result from the search filter")
fmt.Println(result)