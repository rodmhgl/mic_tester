# Twilio Microphone Test App

A simple Golang application that uses Twilio to let callers record and immediately play back audio for microphone testing purposes.

## Overview

This application provides a VoIP service that allows callers to:

1. Call a Twilio phone number
2. Record a short audio message
3. Hear the recording played back immediately
4. Repeat the process as many times as needed

This is particularly useful for testing microphone quality and audio settings during VoIP calls.

## How It Works

When a caller dials your Twilio phone number:

1. The caller hears a welcome message and instructions
2. After a beep, the caller can speak to test their microphone
3. The caller presses `#` to end the recording
4. The system plays back the recording
5. The caller can press `1` to record again or hang up to end the call

## Features

- **Simple Interface**: Clear voice prompts guide callers through the testing process
- **Immediate Feedback**: Recordings are played back immediately after completion
- **Repeat Testing**: Callers can make multiple recordings in a single call
- **Silent Trimming**: Automatically removes silence from the beginning and end of recordings
- **Configurable**: Easy to customize recording length and other parameters

## Technology Stack

- **Backend**: Golang with Gin web framework
- **Voice Services**: Twilio Voice API
- **Deployment**: Docker containerization for Azure Web App
- **CI/CD**: GitHub Actions workflow for automated deployment
- **Documentation**: MkDocs with Material theme

## Getting Started

See the [Installation](getting-started/installation.md) and [Configuration](getting-started/configuration.md) guides to get started.