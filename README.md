# CashierStation Backend
Backend Repository for Playstation Rental Cashier App called CashierStation. The source for the mobile app (frontend) itself is not available.

## Pre-requisites

-   TimescaleDB 15 instance
-   Linux instance (Tested on Ubuntu 20.04)
-   Auth0 account

## Getting Started

-   ### Clone the repository

    ```
    git clone https://github.com/CashierStation/cs-backend.git
    ```

-   ### Navigate to the project directory

    ```
    cd cs-backend
    ```

-   ### Install dependencies

    ```
    go get
    ```

-   ### Setup secrets and variables
    - Set secrets values within the `.env` file. Create one if theres no `.env` file

    - Set your postgres host and creds on `docker-compose.yml`

    - Change the image name across docker related and build related files.
    
-   ### Run the project for development

    ```
    go run .
    ```

    This command will start the server on http://localhost:8080. The server will automatically restart if you make any changes to the code.

-   ### Build the project for production
    Run this command on your local machine/CI server
    ```
    ./build_prod.sh
    ```

-   ### Deploy the project for production
    The deployment is automated by GitHub Workflows. However, if you havent set it up, run this command on your vm/host server. 
    ```
    ./deploy.sh
    ```

## Deployment

Deployment is triggered manually by running Github Workflows Action which will build and deploy the image to your host server via SSH.
