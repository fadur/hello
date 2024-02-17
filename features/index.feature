Feature: Index Handler

  Scenario: Successful index request
    Given the user makes a GET request to the index endpoint
    When the server receives the request
    Then the server should respond with a 200 status code
    And the server should return the index page

