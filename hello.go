package main
import (
  "github.com/mParticle/mparticle-go-sdk/events"
  "context"
  "fmt"
)

func main() {
client := events.NewAPIClient(events.NewConfiguration())

ctx := context.WithValue(
    context.Background(),
    events.ContextBasicAuth,
    events.BasicAuth{
        APIKey:    "APPKEY",
        APISecret: "APPSECRET",
    },
)
batch := events.Batch{Environment: events.DevelopmentEnvironment} //or "ProductionEnvironment"

//set user identities
batch.UserIdentities = &events.UserIdentities{
    CustomerID: "go1234",
    Email:      "go-example@foo.com",
}

//set device identities
batch.DeviceInfo = &events.DeviceInformation{
    IOSAdvertisingID: "607258d9-c28b-43ad-95ed-e9593025d5a1",
}

//set user attributes
batch.UserAttributes = make(map[string]interface{})
batch.UserAttributes["foo"] = "bar"
batch.UserAttributes["foo-array"] = []string{"bar1", "bar2"}

customEvent := events.NewCustomEvent()
customEvent.Data.EventName = "My Custom Event Name"
customEvent.Data.CustomEventType = events.OtherCustomEventType
customEvent.Data.CustomAttributes = make(map[string]string)
customEvent.Data.CustomAttributes["foo"] = "bar"

screenEvent := events.NewScreenViewEvent()
screenEvent.Data.ScreenName = "My Screen Name"

ttl := 123.12

product := events.Product{
    TotalProductAmount: &ttl,
    ID:                 "product-id",
    Name:               "product-name",
}

commerceEvent := events.NewCommerceEvent()
commerceEvent.Data.ProductAction = &events.ProductAction{
    Action:        events.PurchaseAction,
    TotalAmount:   &ttl,
    TransactionID: "foo-transaction-id",
    Products:      []events.Product{product},
}

batch.Events = []events.Event{customEvent, screenEvent, commerceEvent}

result, err := client.EventsAPI.UploadEvents(ctx, batch)
if result != nil && result.StatusCode == 202 {
    fmt.Println("Upload successful")
} else {
    fmt.Errorf(
        "Error while uploading!\nstatus: %v\nresponse body: %#v",
        err.(events.GenericError).Error(),
        err.(events.GenericError).Model())
}
}
