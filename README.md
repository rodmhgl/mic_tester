# Twilio Microphone Test App

A Golang application that lets callers record short audio messages and immediately hear them played back. Perfect for testing microphone quality over VoIP calls.

## Features

- Callers can record a message after the beep
- The recording is immediately played back
- Callers can repeat the process as many times as needed
- Simple and clear voice prompts guide users
- Automatic silence trimming for better audio quality

## Technology Stack

- Golang with Gin web framework
- Twilio Voice API for call handling
- Docker for containerization
- Azure Web App for hosting (or other alternatives)
- GitHub Actions for CI/CD
- MkDocs for documentation

## Prerequisites

- Go 1.21 or higher
- Twilio account with a phone number
- Azure account (for deployment) or alternative hosting solution
- GitHub account for repository hosting and deployment

## Quick Start

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/twilio-mic-test.git
   cd twilio-mic-test
   ```

2. Copy the example environment file and fill in your details:
   ```
   cp .env.example .env
   ```

3. Run the application locally:
   ```
   go run main.go
   ```

4. Expose your local server using ngrok or a similar tool:
   ```
   ngrok http 8080
   ```

5. Set your Twilio phone number's voice webhook to point to your ngrok URL + "/voice"

6. Call your Twilio number to test the application!

## Deployment

See the [documentation](https://yourusername.github.io/twilio-mic-test/) for detailed deployment instructions to Azure and other platforms.

## Documentation

Full documentation is available at [https://yourusername.github.io/twilio-mic-test/](https://yourusername.github.io/twilio-mic-test/)

To build the documentation locally:
```
pip install mkdocs mkdocs-material
mkdocs serve
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.