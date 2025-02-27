package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Set up Gin router with release mode for production
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Routes for Twilio webhooks
	router.POST("/voice", handleFirstTimeCall)
	router.POST("/voice/repeat", handleRepeatCall)
	router.POST("/record", handleRecording)
	router.POST("/playback", handlePlayback)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// handleFirstTimeCall responds to first-time callers with detailed instructions
func handleFirstTimeCall(c *gin.Context) {
	const twiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Welcome to the microphone test service. This tool will help you test your microphone quality. After the beep, please speak normally to test your microphone. When you're finished recording, press the pound key. You'll then hear your recording played back, allowing you to evaluate your microphone's sound quality.</Say>
    <Record
        action="/record"
        maxLength="30"
        finishOnKey="#"
        playBeep="true"
        trim="trim-silence"
    />
</Response>`

	c.Header("Content-Type", "text/xml")
	c.String(http.StatusOK, twiml)
}

// handleRepeatCall responds to repeat callers with briefer instructions
func handleRepeatCall(c *gin.Context) {
	const twiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Speak after the beep. Press pound when done.</Say>
    <Record
        action="/record"
        maxLength="30"
        finishOnKey="#"
        playBeep="true"
        trim="trim-silence"
    />
</Response>`

	c.Header("Content-Type", "text/xml")
	c.String(http.StatusOK, twiml)
}

// handleRecording receives the recording URL and plays it back to the caller
func handleRecording(c *gin.Context) {
	recordingURL := c.PostForm("RecordingUrl")

	if recordingURL == "" {
		// If no recording was made, prompt to try again
		const twiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>No recording was detected. Let's try again.</Say>
    <Redirect>/voice/repeat</Redirect>
</Response>`
		c.Header("Content-Type", "text/xml")
		c.String(http.StatusOK, twiml)
		return
	}

	// Play back the recording and offer to record again
	twiml := `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Here is your recording:</Say>
    <Play>` + recordingURL + `</Play>
    <Gather numDigits="1" action="/playback" method="POST">
        <Say>Press 1 to record again, or hang up to end the call.</Say>
    </Gather>
    <Say>Thank you for using the microphone test service. Goodbye.</Say>
</Response>`

	c.Header("Content-Type", "text/xml")
	c.String(http.StatusOK, twiml)
}

// handlePlayback processes the caller's choice after playback
func handlePlayback(c *gin.Context) {
	digits := c.PostForm("Digits")

	if digits == "1" {
		// Redirect to the repeat voice handler for shorter instructions on subsequent recordings
		const twiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Redirect>/voice/repeat</Redirect>
</Response>`
		c.Header("Content-Type", "text/xml")
		c.String(http.StatusOK, twiml)
		return
	}

	// If any other key was pressed, or no key was pressed, end the call
	const twiml = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Thank you for using the microphone test service. Goodbye.</Say>
</Response>`
	c.Header("Content-Type", "text/xml")
	c.String(http.StatusOK, twiml)
}
