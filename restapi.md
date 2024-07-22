### REST Api

REST Stands for REpresentational State Transfer. It is an architecutre style and not a protocol.

#### HATEOAS

Hypermedia As The Engine Of Application State (HATEOAS) is a constraint of the REST Architecutre.

**Hypermedia**: referes to any content that contains links to other form of media like images, documents etc.

REST architecutre style allows hypermedia links as part of the API response.Navigating hypermedia links is conceptually the same as browsing through web pages by clicking the relevant hyperlinks to achieve a final goal.

Example of REST API response for the following request: HTTP GET http://api.domain.com/management/departments/10
```
{
    "departmentId": 10,
    "departmentName": "Administration",
    "locationId": 1700,
    "managerId": 200,
    "links": [
        {
            "href": "10/employees",
            "rel": "employees",
            "type" : "GET"
        }
    ]
}
```
If you look at links section above, it notifies the user(client) that to get to a further resource by using the `/employees` endpoint along with the current request. 

The advantage of such a mechanism is hypermedia links returned from the server side drive the application state and not the other way around. For a REST client with no prior information about the API this mechanism is useful because the REST client can hit an initial API endpoint and uses the server-provided links to access the resources it needs and discover available actions dynamically.

#### HTTP Methods

CRUD - Create, Retrieve, Update, Delete 

Basically any application that uses a data store you will eventually have to create the above functionality to handle the data store. HTTP provides the following methods for this:

1. HTTP GET

GET requests are used for retreiving the resource representation/information. Should not use GET to modify the resource in any way at all.
Since GET methods dont modify the state of the resource, GET methods are considered safe methods.
GET APIs should be idempotent. Making multiple identical requests must produce the same result every time until another API (POST or PUT) has changed the state of the resource on the server.

Response Codes:

a. For any given HTTP GET API, if the resource is found on the server, then it must return HTTP response code 200 (OK) – along with the response body, which is usually either XML or JSON content (due to their platform-independent nature).
b. In case the resource is NOT found on the server then API must return HTTP response code 404 (NOT FOUND).
c. Similarly, if it is determined that the GET request itself is not correctly formed then the server will return the HTTP response code 400 (BAD REQUEST).

2. HTTP POST

Use POST APIs to create new subordinate resources, e.g., a file is subordinate to a directory containing it or a row is subordinate to a database table.
Basically POST APIs are used to create new resources in a collection of resources (REST API terminology).
Responses to this method are not cacheable unless the response includes appropriate Cache-Control or Expires header fields.
Please note that POST is neither safe nor idempotent, and invoking two identical POST requests will result in two different resources containing the same information (except resource ids).

Response Codes:

a. Ideally, if a resource has been created on the origin server, the response SHOULD be HTTP response code 201 (Created) and contain an entity that describes the status of the request and refers to the new resource, and a Location header.
b. Many times, the action performed by the POST method might not result in a resource that can be identified by a URI. In this case, either HTTP response code 200 (OK) or 204 (No Content) is the appropriate response status.

3. HTTP PUT

PUT requests are used to update an existing resource (if the resource does not exist, then API may decide to create a new resource or not). 
If the request passes through a cache and the Request-URI identifies one or more currently cached entities, those entries SHOULD be treated as stale. Responses to PUT method are not cacheable.

Response Codes:

a. If a new resource has been created by the PUT API, the origin server MUST inform the user agent via the HTTP response code 201 (Created) response.
b. If an existing resource is modified, either the 200 (OK) or 204 (No Content) response codes SHOULD be sent to indicate successful completion of the request.

4. HTTP DELETE

DELETE APIs delete the resources (identified by the Request-URI). If the request passes through a cache and the Request-URI identifies one or more currently cached entities, those entries SHOULD be treated as stale. Responses to this method are not cacheable.

Response Codes:

a. A successful response of DELETE requests SHOULD be an HTTP response code 200 (OK) if the response includes an entity describing the status.
b. The status should be 202 (Accepted) if the action has been queued.
c. The status should be 204 (No Content) if the action has been performed but the response does not include an entity.

5. HTTP PATCH

Use this method to partially update a resource. Use POST to update a resource in its entirety.

#### What is a resource?

In REST the primary data representation is Resource. Any information that we can name is considered as a Resource. In other words, any concept that might be the target of an author’s hypertext reference must fit within the definition of a resource.
Examples of a resource: document, image, service, collection of other resources, or a non virtual object like a person.

For REST apis the state of the resource can change over time (more on states later). At any given time the state of the resource is known as **resource representation**. The resource representation consists of:

1. Data itself
2. Metadata describing the data
3. Hypermedia Links : this will help clients to trasition to the next desired state

**Resource Model**: The set of resources for the REST API. These resources may be interlinked or alone.

Every Resource should conform to the following: 

1. Resource Identifiers

REST uses resource identifiers to identify each resource involved in the interactions between the client and the server components.
Each resource is treated as a separate entity that can be accessed and manipulated independently. 
For example, a resource for a user might be represented by the URI "/users/123", where "123" is the identifier for a specific user.
URI: Uniform Resource Identifier

2. HypterText

Every Resource should be a hypertext or hypermedia(described in the HATEOAS section)

3. Self-Descriptive

Resource representations shall be self-descriptive: the client does not need to know if a resource is an employee or a device. It should act based on the media type associated with the resource.

#### The six guiding principles for REST Api's:

1. Uniform Interface

Every REST interface should have a uniform interface. The following constraints help achieve it: 

    - Every resource should have a uniform identifier
    - Manipulation of resources through representations: The resources should have uniform representations in the server response. API consumers should use these representations to modify the resource state in the server.
    - Self-descriptive messages: Each resource representation should carry enough information to describe how to process the message. It should also provide information of the additional actions that the client can perform on the resource.
    - Hypermedia as the engine of application state: The client should have only the initial URI of the application. The client application should dynamically drive all other resources and interactions with the use of hyperlinks.

2. Client-Server architecture

REST API based systems should follow this architecture. The reasoning is that: By separating the user interface concerns (client) from the data storage concerns (server), we improve the portability of the user interface across multiple platforms and improve scalability by simplifying the server components.

While the client and the server evolve, we have to make sure that the interface/contract between the client and the server does not break.

3. Stateless

It recommends making all the client-server interactions stateless. What this means is that the server will not store anything about the latest HTTP request the client made. It will treat every request as new.
Each HTTP request to a RESTful service must contain all the information needed to understand and process this request. This kind of statelessness makes it easier to scale, cache, and manage the service.
Stateless REST APIs do not establish or maintain client sessions. Clients are responsible for providing all necessary information in each request, such as authentication tokens, credentials, or context data. The server does not store client-specific session data.
The client is responsible for storing and handling the session-related information on its own side.This also means that the client is responsible for sending any state information to the server whenever it is needed.

Advantages of this approach:

a. Stateless APIs are often simpler to develop, test, and maintain because they do not require managing session state or tracking client interactions.
b. Statelessness helps in scaling the APIs to millions of concurrent users by deploying it to multiple servers. Any server can handle any request because there is no session-related dependency.
c. Being stateless makes REST APIs less complex – by removing all server-side state synchronization logic.
d. A stateless API is also easy to cache as well.
e. The server never loses track of “where” each client is in the application because the client sends all necessary information with each request.

4. Cacheable

The cacheable constraint requires that a response should implicitly or explicitly label itself as cacheable or non-cacheable. If the response is cacheable, the client application gets the right to reuse the response data later for equivalent requests and a specified period.

There can be levels of caches such as local cache, proxy cache, or reverse proxy and finally if none of these caches have the request response stored then the service handles it. By using HTTP headers, an origin server indicates whether a response can be cached and, if so, by whom, and for how long.

GET requests should be cachable by default – until a special condition arises. Usually, browsers treat all GET requests as cacheable.
POST requests are not cacheable by default but can be made cacheable if either an Expires header or a Cache-Control header with a directive, to explicitly allows caching, is added to the response.
Responses to PUT and DELETE requests are not cacheable at all.

5. Layered System

The layered system style allows an architecture to be composed of hierarchical layers by constraining component behavior. In a layered system, each component cannot see beyond the immediate layer they are interacting with.

Consider this:

```
@GET
public String myService() {
    return "<html><body><div>HELLO</div></body></html>";
}
```

This is an example of a not layered system. Here you have the service and presentation layer all mixed up. Instead, the service should just return "HELLO", while the client (which I assume here is a presentation layer) should be able to decide how to present the data. 

[link](https://stackoverflow.com/questions/30303116/layered-system-constraint-in-rest-api)

6. Code on Demand (Optional)

REST also allows client functionality to extend by downloading and executing code in the form of applets or scripts.

Now it's commonplace for JavaScript-powered web apps to be consuming REST APIs. This is an example of code on demand - the browser grabs an initial HTML document and supports <script> tags inside that document so that an application can be loaded on-demand.
