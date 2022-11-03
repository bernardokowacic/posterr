<h1 align="center">
  Posterr
</h1>

# Setting up the project
## Pre-requisites
To run this application you need to have the following tools installed on your machine:

* [Git](https://git-scm.com)
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Running Posterr
Right now the project runs with docker-compose, so to execute the project you must run the command `docker-compose up` in your terminal inside the project folder. It will start the API in port 8080 and the database in port 5432.

# Critique
## ToDo
This project lacks unitary tests in the repository package. This was due to the lack of time to finish the project. It also needs an integration test, that in my point of view is more important to a project than the integration tests.

I would also create a middleware to handle the user validation, but since the project description said that no authentication should be done I didn't do it.

Another thing that must be done is to create a struct to simplify the validation of the parameters in the `/posts` end-point.

Due to the lack of time, I didn't do a validation of when the migrations and the seeds should be executed. I would put them inside a function that would only run when a specific flag is passed in the execution command.

## Scaling
Since this is a monolith, its scaling of it will be cost inefficient. This project separates the business rules from the handlers, so the services could be separated into different microservices to handle the increasing amount of requisitions, and then the API could work as a BFF.

The first point of attention in the project if it starts scaling is the post repository, specifically the method that gets the posts. Because of the way it was done, it has an logarithmic complexity `O(n)`.

# Disclaimer
The exercise asked to return the user creation date in the "March 25, 2021" format, but in my point of view, this is a front-end responsability, since the back-end doesn't know for what purpose the data will be used and the format changes depending on the region.

The end-points are in the `posterr.postman_end_points.json` files. To access and use it you can use Postman.
