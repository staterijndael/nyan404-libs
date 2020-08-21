# nyan404-libs

Internal libs for nyan404 server
## Basic usage internal database
```go
	db := NewModelStorage()
	
    	testUserCaseModel := &models.UserCase{
		UserInfo: models.UserInfo{
			Name: "Василий",
		},
	}

	db.Model(testUserCaseModel).Set()

	outValue, err := db.Model(&models.UserCase{}).Field("Name").Equal("Василий").Get()
  
  
  
  ```
In this example we are creating new model and insert this into the internal database. After we trying to get model with that value with our orm and successfully receive.
