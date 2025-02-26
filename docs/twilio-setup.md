# Twilio Setup Guide

This guide walks you through setting up your Twilio account and configuring it to work with the Microphone Test App.

## Prerequisites

- A Twilio account (you can sign up for a free trial at [twilio.com](https://www.twilio.com))
- Your application deployed and accessible via a public URL

## Step 1: Create a Twilio Account

1. Go to [twilio.com](https://www.twilio.com) and sign up for an account if you don't already have one
2. Verify your email address and phone number
3. Note your **Account SID** and **Auth Token** from the Twilio Console dashboard - you'll need these for your app configuration

## Step 2: Get a Twilio Phone Number

1. In the Twilio Console, navigate to **Phone Numbers** > **Manage** > **Buy a Number**
2. Search for a phone number with voice capabilities
3. Purchase a phone number that suits your needs
4. Note down this phone number for your app configuration

## Step 3: Configure Your Twilio Phone Number

1. In the Twilio Console, go to **Phone Numbers** > **Manage** > **Active Numbers**
2. Click on the phone number you purchased
3. Under the **Voice & Fax** section, find the **A Call Comes In** setting
4. Select **Webhook** from the dropdown menu
5. In the URL field, enter your application's voice endpoint: `https://your-app-url.com/voice`
6. Make sure the HTTP method is set to **POST**
7. Click **Save** at the bottom of the page

## Step 4: Configure Your Application

1. Update your application's `.env` file with your Twilio credentials:
   ```
   TWILIO_ACCOUNT_SID=your_account_sid_here
   TWILIO_AUTH_TOKEN=your_auth_token_here
   TWILIO_PHONE_NUMBER=your_twilio_phone_number_here
   ```

2. Restart your application to apply the changes

## Step 5: Testing Your Setup

1. Call your Twilio phone number from any phone
2. You should hear the welcome message asking you to record after the beep
3. Record a short message and press `#`
4. You should hear your recording played back
5. You can press `1` to record again or hang up

## Troubleshooting

If your application isn't receiving calls properly:

1. **Check your webhooks**: In the Twilio Console, go to **Monitor** > **Logs** > **Debugger** to see if there are any errors when Twilio tries to connect to your webhook
2. **Verify your URL**: Make sure your application is publicly accessible at the URL you configured
3. **Check your application logs**: Look for any errors in your application logs that might indicate issues processing Twilio requests
4. **Test your endpoints**: Use a tool like Postman to send test POST requests to your `/voice` endpoint to verify it's responding correctly

## Next Steps

- Consider setting up [TwiML Bins](https://www.twilio.com/docs/runtime/tutorials/twiml-bins) for quick testing
- Explore [Twilio's Voice API documentation](https://www.twilio.com/docs/voice) for more advanced features
- Set up [call recording storage](https://www.twilio.com/docs/voice/tutorials/call-recording-php) if you want to save recordings for later analysis