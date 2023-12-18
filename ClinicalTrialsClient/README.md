# StudyScoutService

Service that handles subscription and billing

## SETUP

Need to create a .env file at the project root level. It needs to contain the following:

```
AUTH0_DOMAIN =''
AUTH0_AUDIENCE =''
stripeKey = ""
AUTH0_CLIENT_ID = ""
AUTH0_CLIENT_SECRET = ""
baseFeePrice = ""
checkoutSessionSuccessURL =''
checkoutSessionCancelURL =''
billingPortalSessionReturnURL =''
```

> [!NOTE]
> Yes the values need to be filled in

## WINDOWS BUILD/RUN

In terminal:

```
go run main.go handler.go access-token-handler.go rest-api-prepare-response.go rest-api-struct.go subscription-type-handler.go
```

OR

- Open folder in VSCode
- Open the main.go file
- While selecting the main.go press F5

> [!NOTE]
> Make sure you dont have any other applications running on 8080 port

## DOCKER BUILD

[BUILD IMAGE]

```
docker build --rm -t studyscout-go:alpha .
```

[BUILD CONTAINER]

```
docker run -d -p 8080:8080 --name go-docker-api studyscout-go:alpha
```
