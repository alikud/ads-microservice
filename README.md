```ACTUAL STATUS - REFACTORING```

This is microservice for storing and posting advertisements. Advertisements are stored in a database. The service provide an API that works over HTTP in JSON format.

Requirements:

✅ - Simple instructions for running the service (ideally with the ability to run via docker-compose up, but this is not mandatory);
The project can be launched by cloning the repository and running make up.

✅ - Field validation: no more than 3 links to photos, description no more than 1000 characters, title no more than 200 characters.

Details
Method for getting a list of advertisements

✅ - Pagination: 10 advertisements should be displayed on one page (this parameter can be configured, currently set to 2).

✅ - Sorting: by price (ascending/descending) and by creation date (ascending/descending).

✅ - Fields in the response: title of the advertisement, link to the main photo (first in the list), price.

Method for getting a specific advertisement

✅ - Mandatory fields in the response: title of the advertisement, price, link to the main photo.

Optional fields (can be requested by passing the "fields" parameter): description, links to all photos.
Method for creating an advertisement:

✅ - Accepts all the fields listed above: title, description, several links to photos (the actual photos do not need to be uploaded anywhere), price.
✅ - Returns the ID of the created advertisement and the result code (success or failure).

Additional features

Unit tests: try to achieve test coverage of 70% or higher.
✅ - Containerization: the project can be launched using the docker-compose up command.

The service architecture is described in text and/or diagrams.
✅ - Documentation: there is a structured description of the service methods.
