# chirpy
Boot.dev project to build a web server in Golang for a mock social network.


## API Endpoints

Here's a list of the API endpoints:

*   **`GET /api/healthz`**: Checks the health status of the API. <br>
    Response Body:
    ```json
    {
        "status": "ok"
    }
    ```
*   **`GET /admin/metrics`**: Provides server metrics for administrative purposes.
*   **`POST /admin/reset`**: Resets server data (admin-only).
*   **`POST /api/users`**: Creates a new user account. <br>
    Expected Request Body:
    ```json
    {
        "email": "string",
        "password": "string"
    }
    ```
    Response Body:
    ```json
    {
        "id": 123,
        "created_at": "2025-08-01 16:50:13.793654 +0000 UTC",
        "updated_at": "2025-08-01 18:24:44.283421 +0000 UTC",
        "email": "user@example.com",
        "is_chirpy_red": false
    }
    ```
*   **`PUT /api/users`**: Updates an existing user's information. <br>
    **Authentication Required**: Bearer Token. <br>
    Expected Headers:
    ```
    Authorization: Bearer <your_jwt_token>
    ```
    Expected Request Body:
    ```json
    {
        "email": "string",
        "password": "string"
    }
    ```
    Response Body:
    ```json
    {
        "id": 123,
        "created_at": "2025-08-01 16:50:13.793654 +0000 UTC",
        "updated_at": "2025-08-01 18:24:44.283421 +0000 UTC",
        "email": "user@example.com",
        "is_chirpy_red": false
    }
    ```
*   **`POST /api/login`**: Authenticates a user and provides session tokens.    <br>
    Expected Request Body:
    ```json
    {
        "email": "string",
        "password": "string"
    }
    ```
    Response Body:
    ```json
    {
        "id": 123,
        "created_at": "2025-08-01 16:50:13.793654 +0000 UTC",
        "updated_at": "2025-08-01 18:24:44.283421 +0000 UTC",
        "email": "user@example.com",
        "token": "jwt_token_string",
        "refresh_token": "refresh_token_string",
        "is_chirpy_red": false
    }
    ```
*   **`POST /api/refresh`**: Refreshes an expired access token using a refresh token. <br>
    **Authentication Required**: Bearer Token. <br>
    Expected Headers:
    ```
    Authorization: Bearer <your_refresh_token>
    ```
    Response Body:
    ```json
    {
        "token": "jwt_token_string"
    }
    ```
*   **`POST /api/revoke`**: Revokes a user's refresh token, logging them out.
    <br>
    **Authentication Required**: Bearer Token. <br>
    Expected Headers:
    ```
    Authorization: Bearer <your_refresh_token>
    ```
*   **`POST /api/chirps`**: Creates a new chirp (post) for the social network.
    <br>
    **Authentication Required**: Bearer Token. <br>
    Expected Headers:
    ```
    Authorization: Bearer <your_jwt_token>
    ```
    Expected Request Body:
    ```json
    {
        "body": "string"
    }
    ```
    Response Body:
    ```json
     {
        "id": 1,
        "created_at": "2025-08-01 16:50:13.793654 +0000 UTC",
        "updated_at": "2025-08-01 18:24:44.283421 +0000 UTC",
        "body": "Hello world!",
        "user_id": 123
    }
    ```
*   **`GET /api/chirps?sort={sortingOrder}&author_id={userID}`**: Retrieves a list of chirps; optional query parameters are sorting order (either "asc" by default or "desc") and "author_id" to filter chirps by user. <br>
    Response Body:
    ```json
    [
        {
            "id": 1,
            "created_at": "2025-08-01 16:50:13.793654 +0000 UTC",
            "updated_at": "2025-08-01 18:24:44.283421 +0000 UTC",
            "body": "Hello world!",
            "user_id": 123
        },
        {
            "id": 2,
            "created_at": "2025-08-02 11:34:38.526116 +0000 UTC",
            "updated_at": "2025-08-04 09:22:48.262227 +0000 UTC",
            "body": "Another chirp here!",
            "user_id": 456
        }
    ]
    ```
*   **`GET /api/chirps/{chirpID}`**: Retrieves a specific chirp by its ID. <br>
    Response Body:
     ```json
     {
        "id": 1,
        "created_at": "2025-08-01 16:50:13.793654 +0000 UTC",
        "updated_at": "2025-08-01 18:24:44.283421 +0000 UTC",
        "body": "Hello world!",
        "user_id": 123
    }
    ```
*   **`DELETE /api/chirps/{chirpID}`**: Deletes a specific chirp by its ID.
    <br>
    **Authentication Required**: Bearer Token. <br>
    Expected Headers:
    ```
    Authorization: Bearer <your_jwt_token>
    ```
*   **`POST /api/polka/webhooks`**: Handles a webhook for user upgrades to "chirpy red" status. <br>
    **Authentication Required**: Api Key. <br>
    Expected Headers:
    ```
    Authorization: ApiKey <your_api_key>
    ```
    Expected Request Body:
    ```json
    {
        "event": "user.upgraded",
        "data": {
            "user_id": 123
        }
    }
    ```