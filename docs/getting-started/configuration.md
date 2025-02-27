# Configuration

This guide explains the configuration options available for the Twilio Microphone Test App.

## Environment Variables

The application uses environment variables for configuration. These can be set directly in your environment or through a `.env` file (recommended for development).

### Required Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `TWILIO_ACCOUNT_SID` | Your Twilio account SID | `ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx` |
| `TWILIO_AUTH_TOKEN` | Your Twilio auth token | `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx` |
| `TWILIO_PHONE_NUMBER` | Your Twilio phone number | `+1234567890` |

### Optional Variables

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `PORT` | The port on which the server will listen | `8080` | `3000` |
| `APP_ENV` | The application environment | `development` | `production` |

## TwiML Customization

The application uses TwiML (Twilio Markup Language) to control call flow. You can customize various aspects of the call experience by modifying the TwiML responses in `main.go`.

### Recording Settings

In the `handleIncomingCall` function, you can adjust the `<Record>` element's attributes:

```xml
<Record
    action="/record"
    maxLength="30"
    finishOnKey="#"
    playBeep="true"
    trim="trim-silence"
/>
```

| Attribute | Description | Default |
|-----------|-------------|---------|
| `maxLength` | Maximum recording duration in seconds | `30` |
| `finishOnKey` | Key to press to end recording | `#` |
| `playBeep` | Whether to play a beep before recording | `true` |
| `trim` | Whether to trim silence from recordings | `trim-silence` |

### Voice Prompts

You can customize the voice prompts by modifying the text within the `<Say>` elements:

```go
const twiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Welcome to the microphone test service. After the beep, please speak to test your microphone. When finished, press the pound key.</Say>
    <!-- ... -->
</Response>`
```

## Advanced Configuration

### Voice Settings

To change the voice used for prompts, add a `voice` attribute to the `<Say>` elements:

```xml
<Say voice="woman">Welcome to the microphone test service.</Say>
```

Available voice options include:
- `man` (default)
- `woman`
- `alice` (enhanced voice)

### Language Settings

To change the language, add a `language` attribute to the `<Say>` elements:

```xml
<Say language="en-GB">Welcome to the microphone test service.</Say>
```

See [Twilio's documentation](https://www.twilio.com/docs/voice/twiml/say#attributes-language) for a list of supported languages.

## Security Considerations

- Keep your `.env` file secure and never commit it to version control
- If deploying to a public environment, ensure you have appropriate firewalls and rate limiting in place
- Consider implementing authentication for any administrative endpoints you add