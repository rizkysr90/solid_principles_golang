# Single Responsibility Principle (SRP) in Go

The Single Responsibility Principle (SRP) is one of the SOLID principles, and it states that a class should have only one reason to change. In Go, which is a statically typed language with support for interfaces and composition, you can apply SRP in a similar way as in other languages.

```

In this example:

- UserService is responsible for handling user-related operations.
- UserRepository is responsible for database interactions related to users.
- NotifyService is responsible for sending notifications.
- BusinessService orchestrates business operations involving users. It uses the UserRepository and NotifyService to create a user,    adhering to SRP by delegating responsibilities to specialized services.

By separating concerns and responsibilities into distinct types, you adhere to the Single Responsibility Principle. If you need to  change the way users are persisted, you modify UserRepository; if you need to change how notifications are sent, you modify NotifyService. The BusinessService remains focused on orchestrating high-level business operations.
