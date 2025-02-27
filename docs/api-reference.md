# API Reference

This document provides details about the API endpoints exposed by the Twilio Microphone Test App and the TwiML responses they generate.

## Endpoints

The application exposes three main endpoints that handle different stages of the call flow:

### POST /voice

Handles incoming calls and initiates the recording process.

**Request**:
- Method: `POST`
- Content-Type: `application/x-www-form-urlencoded`
- Body: Twilio's standard [voice webhook parameters](https://www.twilio.com/docs/voice/twiml/gather#attributes-action)

**Response**:
- Content-Type: `text/xml`
- Body: TwiML that welcomes the caller and starts recording

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Welcome to the microphone test service. After the beep, please speak to test your microphone. When finished, press the pound key.</Say>
    <Record
        action="/record"
        maxLength="30"
        finishOnKey="#"
        playBeep="true"
        trim="trim-silence"
    />
</Response>
```

**TwiML Elements Used**:
- [`<Say>`](https://www.twilio.com/docs/voice/twiml/say): Converts text to speech to play to the caller
- [`<Record>`](https://www.twilio.com/docs/voice/twiml/record): Records the caller's voice and sends it to the specified action URL

### POST /record

Receives the recording data and plays it back to the caller.

**Request**:
- Method: `POST`
- Content-Type: `application/x-www-form-urlencoded`
- Body: Twilio's standard [recording webhook parameters](https://www.twilio.com/docs/voice/twiml/record#attributes-action)
  - `RecordingUrl`: URL of the recording
  - `RecordingDuration`: Duration of the recording in seconds
  - Other standard parameters

**Response**:
- Content-Type: `text/xml`
- Body: TwiML that plays back the recording and offers to record again

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Here is your recording:</Say>
    <Play>https://api.twilio.com/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/REXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX</Play>
    <Gather numDigits="1" action="/playback" method="POST">
        <Say>Press 1 to record again, or hang up to end the call.</Say>
    </Gather>
    <Say>Thank you for using the microphone test service. Goodbye.</Say>
</Response>
```

**TwiML Elements Used**:
- [`<Say>`](https://www.twilio.com/docs/voice/twiml/say): Converts text to speech to play to the caller
- [`<Play>`](https://www.twilio.com/docs/voice/twiml/play): Plays an audio file to the caller
- [`<Gather>`](https://www.twilio.com/docs/voice/twiml/gather): Collects digits pressed by the caller

**Alternative Response** (if no recording was made):
```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>No recording was detected. Let's try again.</Say>
    <Redirect>/voice</Redirect>
</Response>
```

**Additional TwiML Elements Used**:
- [`<Redirect>`](https://www.twilio.com/docs/voice/twiml/redirect): Redirects to another TwiML URL

### POST /playback

Processes the caller's choice after playback (record again or end the call).

**Request**:
- Method: `POST`
- Content-Type: `application/x-www-form-urlencoded`
- Body: Twilio's standard [gather webhook parameters](https://www.twilio.com/docs/voice/twiml/gather#attributes-action)
  - `Digits`: The digits pressed by the caller (we expect `1` to record again)

**Response** (if `Digits` is `1`):
- Content-Type: `text/xml`
- Body: TwiML that redirects to the `/voice` endpoint to start another recording

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Redirect>/voice</Redirect>
</Response>
```

**Response** (for any other value or no input):
- Content-Type: `text/xml`
- Body: TwiML that ends the call with a goodbye message

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Thank you for using the microphone test service. Goodbye.</Say>
</Response>
```

## TwiML Quick Reference

The application uses the following [TwiML](https://www.twilio.com/docs/voice/twiml) elements:

| Element | Description | Documentation |
|---------|-------------|---------------|
| `<Response>` | Root element for all TwiML responses | [Docs](https://www.twilio.com/docs/voice/twiml) |
| `<Say>` | Converts text to speech | [Docs](https://www.twilio.com/docs/voice/twiml/say) |
| `<Record>` | Records caller's voice | [Docs](https://www.twilio.com/docs/voice/twiml/record) |
| `<Play>` | Plays an audio file | [Docs](https://www.twilio.com/docs/voice/twiml/play) |
| `<Gather>` | Collects caller's input | [Docs](https://www.twilio.com/docs/voice/twiml/gather) |
| `<Redirect>` | Redirects to another endpoint | [Docs](https://www.twilio.com/docs/voice/twiml/redirect) |

## Adding Custom Endpoints

If you want to extend the application with additional functionality, you can add new endpoints in `main.go`:

```go
// Example of adding a new endpoint for call statistics
router.GET("/stats", handleCallStats)

// Implementation
func handleCallStats(c *gin.Context) {
    // Your code here to retrieve and return call statistics
    c.JSON(http.StatusOK, gin.H{
        "total_calls": 100,
        "total_recordings": 250,
        "average_duration": 15.5,
    })
}
```

## Webhook Verification

For production deployments, it's recommended to verify that incoming requests are actually from Twilio. You can implement this by adding middleware to validate the request signature:

```go
// Simplified example of Twilio signature validation middleware
func twilioAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get Twilio signature from header
        twilioSignature := c.GetHeader("X-Twilio-Signature")
        
        // Get request URL and form values
        url := "https://" + c.Request.Host + c.Request.URL.Path
        
        // Validate the signature (simplified example)
        if !validateTwilioSignature(twilioSignature, url, c.Request.Form, os.Getenv("TWILIO_AUTH_TOKEN")) {
            c.AbortWithStatus(403)
            return
        }
        
        c.Next()
    }
}
```

See Twilio's [Security documentation](https://www.twilio.com/docs/usage/security) for more details on webhook validation.