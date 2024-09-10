This project demonstrates the Strategy Pattern in Go by implementing multiple storage strategies (MongoDB and Postgres) for handling numbers. The Strategy Pattern allows you to define a family of algorithms (or in this case, storage mechanisms), encapsulate each one as a strategy, and make them interchangeable at runtime.

### What is the Strategy Pattern?
The Strategy Pattern is a behavioral design pattern that allows you to define different algorithms, encapsulate them, and make them interchangeable at runtime. This pattern decouples the algorithm from the client, meaning that the client doesn’t need to know which specific implementation it’s using.

In this case, we have multiple strategies for storing numbers (e.g., MongoDB and Postgres), and the pattern allows us to easily switch between these different storage methods without modifying the client code (the ApiServer).

### How This Example Works
In this example:
- Strategy Interface: NumberStorer defines two methods: GetAll() and Put(int). This interface is the contract that all storage strategies must follow.
- Concrete Strategies: We have two concrete strategies that implement the NumberStorer interface:
  - MongoDBNumberStorer: Simulates storing numbers in a MongoDB database.
  - PostgresNumberStore: Simulates storing numbers in a PostgreSQL database.
- Context: The ApiServer struct holds a reference to a NumberStorer. This allows the ApiServer to use different storage strategies interchangeably, depending on which one is passed in.
