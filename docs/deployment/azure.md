# Deploying to Azure

This guide walks you through deploying the Twilio Microphone Test App to Azure using GitHub Actions and Azure Web App.

## Prerequisites

Before proceeding, ensure you have:

- An Azure account (you can [create one for free](https://azure.microsoft.com/free/))
- A GitHub account with your code repository
- The following resources created in Azure:
  - Resource Group
  - Azure Container Registry (ACR)
  - Azure Web App (configured for containers)

## Setting Up Azure Resources

### 1. Create a Resource Group

```bash
az group create --name twilio-mic-test-rg --location eastus
```

### 2. Create an Azure Container Registry

```bash
az acr create --resource-group twilio-mic-test-rg --name twilioregistry --sku Basic
```

### 3. Enable Admin Access to ACR

```bash
az acr update --name twilioregistry --admin-enabled true
```

### 4. Get ACR Credentials

```bash
az acr credential show --name twilioregistry
```

Note down the username and passwords from the output.

### 5. Create an Azure Web App

```bash
az appservice plan create --name twilio-mic-test-plan --resource-group twilio-mic-test-rg --sku B1 --is-linux
az webapp create --resource-group twilio-mic-test-rg --plan twilio-mic-test-plan --name twilio-mic-test-app --deployment-container-image-name twilioregistry.azurecr.io/twilio-mic-test:latest
```

### 6. Configure Environment Variables for the Web App

```bash
az webapp config appsettings set --resource-group twilio-mic-test-rg --name twilio-mic-test-app --settings TWILIO_ACCOUNT_SID=your_account_sid TWILIO_AUTH_TOKEN=your_auth_token TWILIO_PHONE_NUMBER=your_twilio_phone_number APP_ENV=production
```

## Setting Up GitHub Actions

To automate deployment using GitHub Actions, you need to add secrets to your GitHub repository.

### 1. Get Azure Credentials for GitHub Actions

```bash
az ad sp create-for-rbac --name "twilio-mic-test-ci" --role contributor \
  --scopes /subscriptions/{subscription-id}/resourceGroups/twilio-mic-test-rg \
  --sdk-auth
```

This will output a JSON object containing credentials.

### 2. Add GitHub Secrets

In your GitHub repository, go to **Settings > Secrets and variables > Actions** and add the following secrets:

- `AZURE_CREDENTIALS`: The entire JSON output from the previous step
- `AZURE_REGISTRY_URL`: Your ACR login server (e.g., `twilioregistry.azurecr.io`)
- `AZURE_REGISTRY_USERNAME`: ACR username
- `AZURE_REGISTRY_PASSWORD`: ACR password
- `AZURE_WEBAPP_NAME`: Web App name (e.g., `twilio-mic-test-app`)

### 3. Configure Workflow File

The `.github/workflows/deploy.yml` file in your repository should already be configured correctly, but verify it contains:

```yaml
name: Deploy to Azure

on:
  push:
    branches: [ main ]
  workflow_dispatch:

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    
    steps:
    - name: Checkout repository
      uses: actions/checkout@v3

    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v2

    - name: Login to Azure Container Registry
      uses: docker/login-action@v2
      with:
        registry: ${{ secrets.AZURE_REGISTRY_URL }}
        username: ${{ secrets.AZURE_REGISTRY_USERNAME }}
        password: ${{ secrets.AZURE_REGISTRY_PASSWORD }}

    - name: Build and push Docker image
      uses: docker/build-push-action@v4
      with:
        context: .
        push: true
        tags: ${{ secrets.AZURE_REGISTRY_URL }}/twilio-mic-test:${{ github.sha }}, ${{ secrets.AZURE_REGISTRY_URL }}/twilio-mic-test:latest

    - name: Login to Azure
      uses: azure/login@v1
      with:
        creds: ${{ secrets.AZURE_CREDENTIALS }}

    - name: Deploy to Azure Web App
      uses: azure/webapps-deploy@v2
      with:
        app-name: ${{ secrets.AZURE_WEBAPP_NAME }}
        images: ${{ secrets.AZURE_REGISTRY_URL }}/twilio-mic-test:${{ github.sha }}
        
    - name: Deploy to GitHub Pages (Documentation)
      uses: mhausenblas/mkdocs-deploy-gh-pages@master
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        CONFIG_FILE: mkdocs.yml
```

## Triggering Deployment

Push to the main branch of your repository to trigger the GitHub Actions workflow:

```bash
git add .
git commit -m "Update application for Azure deployment"
git push origin main
```

The workflow will:
1. Build your Docker image
2. Push it to Azure Container Registry
3. Deploy it to your Azure Web App
4. Deploy your documentation to GitHub Pages

## Configuring Twilio

After deployment, update your Twilio webhook URL to point to your Azure Web App:

```
https://twilio-mic-test-app.azurewebsites.net/voice
```

Follow the [Twilio Setup Guide](../twilio-setup.md) for detailed instructions.

## Monitoring

You can monitor your application using Azure's built-in tools:

- **Logs**: Check application logs in the Azure portal under your Web App > Monitoring > Log stream
- **Metrics**: View performance metrics under your Web App > Monitoring > Metrics
- **Alerts**: Set up alerts for important events under your Web App > Monitoring > Alerts

## Troubleshooting

If you encounter issues with your deployment:

1. Check GitHub Actions workflow runs for error messages
2. Verify that all secrets are correctly set in GitHub
3. Check the application logs in Azure
4. Ensure your Web App is properly configured for containers
5. Verify that Twilio can reach your Web App by checking the Twilio console logs

## Cost Management

The configuration described here uses:
- Azure Web App (Basic B1 tier): ~$13/month
- Azure Container Registry (Basic tier): ~$5/month

To reduce costs, you could:
- Use a Free or Shared tier App Service Plan for non-production workloads
- Delete resources when not in use