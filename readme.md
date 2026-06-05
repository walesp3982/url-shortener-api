# Proyecto Acortador de URL

Es una pequeña api cuyo objetivo es generar un url acortada 
Endpoint importante

|  Endpoint    |   Request       |  Retorna     |
|--------------|-----------------|---------------|
| POST /url     | { url : "http://www.google.com"} | devuelve el code respectivo |
| GET /url      |  None           | Todas la urls con sus acortadores |
| DELETE /url/{code} | None       | No Content |
| GET /{code}  |  None | Redirección a url si es que existe |
