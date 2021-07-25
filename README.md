# url-shortening-service
## Definition of the system
- Provide shorter aliases for long URLs.
- Redirect to the original URL if hit the shortened links.

## Requirements of the system
- The user gives a URL as input -> our service generate a shorter and unique alias of that URL.
- Users hit the shorter link -> our service return original link.
- Links may expire after duration
- Users can specify the expiration time
- Users can't use custom links

## Language
- golang

## Data storage
- mongoDB
- redis
