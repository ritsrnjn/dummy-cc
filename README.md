What are the plus points of this repo?
1. Separation of responsibilities
2. Extensible
3. Maintainaibility
4. Security aspects := middlewares + keeping configs in .env files and not in repo
5. Taken care of status codes
6. Covered edge cases


A brief summary of all the packages


Main: The service will start from here. Main function make sure the env configs are being loaded -> then db connection is made -> and then it start a local server at port 3000


Router: It handles the routing of incoming requests and routes them to appropriate handler functions.


Middleware: This package can be used to integrate middleware before the service does any processing for the request. For now, I have added a check to see if the request body contains an authorization header or not. Multiple other middlewares can be easily added.


Handler: This layer is where the request will land. This laer will do basic sanity check of request body. and will throw 400 in the case of any issue. And if the basic sanity check is passed, then it will call the service layer for further processing.


Service: This is where all the business logic will go. This will have most part of code. This layer will interact with Model layer to get/update offer/account from DB.


Model: This layer will have all the logic related to DB queries and execution. But the actual connection is separated from this layer as well. It helps in abstracting the core DB details, and different or multiple DBs can be easily integrated into this model layer. For this project, this layer interacts with the sqldb package.


Sqldb: This is the layer just above MySQL DB; this handles making the connection with DB and executing the queries on DB.


Constants: This is used to hold constants that are used throughout the repo.


Configuration: to hold sensitive information. This can be further broken into multiple files if the project grows for easier categorization of configs.


Utils: Common utility methog. Again, this package can be extended if categorization is needed for util methods.



Further technical improvements:
1. Adding logs
2. Recording metrics


Further Product Improvements:
1. Some corner cases are missing; what if an offer is accepted and at that moment some offers automatically become invalid? We can automatically reject those offers without any user input.
