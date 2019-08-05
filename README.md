# Go DDD
The repo to demonstrate DDD approach to construct apps.

# Architecture

![alt text](./layers.jpg)

DDD has 4 layers in the architecture:

**Interface:** 

This layer responsibles for the interaction with user, whether software presents information or recieves information from user.

**Application:**
 
This is a thin layer between interface and domain, it could call domain services to serve the application purposes.

**Domain:** 

The heart of the software, this layer holds domain logic and business knowledge.

**Infrastructure:** 

A supporting layer for the other layers. This layer contains supporting libraries or external services like database or UI supporting library.

In the domain layer we know several more terms:

**Entity:** 

an object that has an identity, identity could be a row in a database table or other type of identity that differs one entity and the other.

**Value object:**

an object which we are only interested in its attributes not the identity. Value objects do not have identities and creation/modification of these objects should not be a problem. Thus it is better to be implemented as immutable objects.

**Repository:** 

Entities are usually stored in database, but we would not want to expose our domain logic to database logic so we could have a ‘virtual’ storage for our entities — Repositories.
Service: Business logic sometimes would involve more than one different objects. When it doesn’t feel right to put a function in either object, we would use a service.