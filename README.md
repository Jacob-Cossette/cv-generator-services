# cv-generator-services
# CV Generator Service

The `cv-generator-service` is a microservice that generates a CV for a given user based on a job offer. It takes user data and a job offer in JSON format and returns a CV in PDF format.

## Getting Started

To get started with the `cv-generator-service`, you'll need to have Go installed on your local machine. Follow these steps to run the service:

1. Clone the repository: `git clone https://github.com/your-username/cv-generator-service.git`.
2. Navigate to the project directory: `cd cv-generator-service`.
3. Build the service: `go build`.
4. Run the service: `./cv-generator-service`.

## API Usage

The `cv-generator-service` exposes an API that takes user data and a job offer in JSON format and returns a CV in PDF format. The API has the following endpoints:

* `/cv`: Generates a CV for a given user based on a job offer.

### Request Body

The request body should be in the following JSON format:

```json
{
  "user": {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "+1-555-123-4567",
    "address": "123 Main St, Anytown USA",
    "summary": "A software engineer with 5+ years of experience in web development.",
    "skills": [
      {
        "name": "JavaScript",
        "level": "expert"
      },
      {
        "name": "Python",
        "level": "advanced"
      },
      {
        "name": "SQL",
        "level": "intermediate"
      }
    ]
  },
  "jobOffer": {
    "title": "Software Engineer",
    "company": "Acme Inc.",
    "description": "We're looking for a software engineer to join our team.",
    "skills": [
      {
        "name": "JavaScript",
        "level": "required"
      },
      {
        "name": "React",
        "level": "required"
      },
      {
        "name": "Node.js",
        "level": "preferred"
      }
    ]
  }
}
```

### Response Body
The response body will be a PDF file containing the generated CV.

## Contributing
We welcome contributions to the cv-generator-service! To contribute, please follow these steps:

Fork the repository and clone it to your local machine.
Create a new branch for your feature or bug fix: git checkout -b feature/my-new-feature or git checkout -b bugfix/my-bug-fix.
Write your code and tests.
Commit your changes: git commit -am 'Add some feature' or git commit -am 'Fix some bug'.
Push your branch to your forked repository: git push origin feature/my-new-feature or git push origin bugfix/my-bug-fix.
Create a new pull request from your branch to the master branch of the cv-generator-service repository.
Wait for a code review and address any feedback.
Once your pull request is approved, it will be merged into the master branch and deployed to production.
## License
The cv-generator-service is licensed under the MIT license. See the LICENSE file for more details.

## Contact
If you have any questions or feedback, please contact us at 
