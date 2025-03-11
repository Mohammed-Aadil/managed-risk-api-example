I have followed service domain structure for the project. I tried to minimize layer overlapping in this structure.
This project is created with mindset of production code, so I used go work to support multiple microservices and one common pkg for all other services.
I tried to capture all possible config and DRY principle.
For simplicity, I have added any locking mechanism for inmemory storage consistency. If required it could be applied.
I used doubly linked list to maintain data in more space efficent manner in case if we introduce delete api.