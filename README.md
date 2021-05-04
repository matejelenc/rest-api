# User And Group Management API
- provide a way to list, add, modify and remove users and groups
- uses a postgresql database
# How To Run
- start the database: docker start postgres1
- then start the start the rest api: docker start rest-api1
# OpenAPI Documentation
- run localhost:8080/docs
# Routes
- /users GET(get all users) and CREATE(creates a user)
- /users/{id} GET(get a user), PATCH(update a user), DELETE(delete a user)
- /groups GET(get all groups) and CREATE(creates a group)
- /groups/{id} GET(get a group), PATCH(update a group), DELETE(delete a group)
- /groups/{id}/members GET(get all users in this group)
